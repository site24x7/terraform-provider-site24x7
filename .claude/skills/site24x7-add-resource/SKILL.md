---
name: site24x7-add-resource
description: >
  Add a new resource or data source to the terraform-provider-site24x7 codebase from start
  to finish — API type, REST endpoint client, fixtures and endpoint test, testify fake,
  Client interface wiring, schema and CRUD with an Importer, resource unit tests, provider
  registration, docs page and example — following this repo's conventions for 404 handling,
  secrets and sets. Also covers retrofitting a missing piece (Importer, data source, 404
  handling) onto an existing resource. Use whenever a developer wants to add, scaffold or
  implement a new site24x7_* resource, monitor type, integration or data source, expose an
  existing api/endpoints client to Terraform, add import support, or asks "what files do I
  need to touch to add a resource" — even if they don't say "skill".
user-invocable: true
argument-hint: "<what to add, e.g. 'site24x7_smtp_monitor resource' or 'data source for business hours' or 'Importer for site24x7_amazon_monitor'>"
allowed-tools:
  - Read
  - Write
  - Edit
  - Glob
  - Grep
  - Bash
  - PowerShell
---

# Add a resource to the Site24x7 provider

This is a **developer** skill: it changes the provider's Go code. For writing HCL that *uses*
the provider, use `site24x7-monitor-authoring` or `site24x7-alerting-and-profiles`.

A resource touches **up to 13 places**. Most defects in this repo came from skipping one of
them, or from copying a neighbour and not renaming everything (the credential profile tests
were once named `TestRestApiMonitors` and pointed at monitor fixtures). Work through the
checklist in order, because every layer compiles against the one before it.

The provider is on **`terraform-plugin-sdk` v1** (`github.com/hashicorp/terraform-plugin-sdk
v1.1.1`). Use `Create`/`Read`/`Update`/`Delete` returning `error`, not the SDKv2
`CreateContext` / `diag.Diagnostics` forms. They won't compile here.

## Step 0: Pin down the API before writing code

Read the endpoint in the Site24x7 API docs (`https://www.site24x7.com/help/api/`) and write
down, before touching Go:

1. **Paths for each verb.** Site24x7 often uses the **singular for get/create/update/delete
   and the plural for list**: `/credential_profile/{id}` vs `/credential_profiles`,
   `apminsight/agent_config_profile` vs `.../agent_config_profiles`. Don't assume.
2. **The ID field name** in the response (`provider_id`, `profile_id`, `monitor_id`,
   `credential_profile_id` …). It becomes `d.Id()`.
3. **Which fields the API never returns, masks or encrypts** (passwords, tokens, secrets).
   That decides the secret pattern in Step 8.
4. **Which fields the server defaults** when you omit them. Those become `Optional +
   Computed`, not `Default`.
5. **Which list fields are unordered** (ID lists: groups, tags, user groups). Those become
   `TypeSet`.
6. **Whether create exists at all.** `site24x7_apm_application` is adopt-only because
   applications appear when an agent reports in. If there is no POST, the resource's Create
   has to look up an existing object, and that needs its own design.
7. **What DELETE really does.** `site24x7_customer`'s delete is a no-op. Document it if so.

If a similar resource already exists, pick the **closest one with good patterns** as the
template:

| Pattern | Best reference in repo |
|---|---|
| Plain CRUD, 404 handling, logging, validation | [site24x7/apm/apm_agent_config_profile.go](../../../site24x7/apm/apm_agent_config_profile.go) |
| Nested blocks (`TypeList` + `MaxItems: 1`, list of objects) | [site24x7/common/oauth2_provider.go](../../../site24x7/common/oauth2_provider.go) |
| Monitor with profiles, user groups, tags | [site24x7/monitors/website.go](../../../site24x7/monitors/website.go) |
| Integration | [site24x7/integration/slack.go](../../../site24x7/integration/slack.go) |
| `name_regex` data source | [site24x7/common/credential_profile_data_source.go](../../../site24x7/common/credential_profile_data_source.go) |
| Endpoint test | [api/endpoints/common/credential_profiles_impl_test.go](../../../api/endpoints/common/credential_profiles_impl_test.go) |
| Resource test | [site24x7/apm/apm_agent_config_profile_test.go](../../../site24x7/apm/apm_agent_config_profile_test.go) |

Don't copy `oauth2_provider.go`'s Read as-is. It writes `client_secret` and `auth_pass` back
from the API response, which is the perpetual-diff pattern Step 8 warns against.

## Step 1: Pick the package

The Go package decides three directories at once. Keep them aligned:

| Kind | Resource package | Endpoint package | Client accessor type |
|---|---|---|---|
| Monitor | `site24x7/monitors` | `api/endpoints/monitors` | `monitors.X` |
| Integration | `site24x7/integration` | `api/endpoints/integration` | `integration.X` |
| Account-wide object (credential, schedule, SLA, OAuth2 …) | `site24x7/common` | `api/endpoints/common` | `common.X` |
| APM Insight | `site24x7/apm` | `api/endpoints/apm` | `apm.X` |
| MSP | `site24x7/msp` | `api/endpoints/msp` | `msp.X` |
| AWS helpers | `site24x7/aws` | `api/endpoints/aws` | `aws.X` |
| Core profiles/groups (legacy location) | `site24x7` | `api/endpoints` | `endpoints.X` |

Put new work in a sub-package, not the root `site24x7` package. Sub-packages get the client
with `meta.(site24x7.Client)`. Only root-package files use the bare `meta.(Client)`.

## Step 2: API type (`api/`)

Add the struct to `api/<name>.go`, or to `api/monitor_types.go` for a monitor.

```go
// Widget represents a Widget in Site24x7.
type Widget struct {
	_          struct{} `type:"structure"` // Enforces key based initialization.
	WidgetID   string   `json:"widget_id,omitempty"`
	WidgetName string   `json:"widget_name"`
	Secret     string   `json:"secret,omitempty"`
	GroupIDs   []string `json:"group_ids,omitempty"`
}
```

- The **ID field is `omitempty`**, so create requests don't send `"widget_id": ""`.
- Required fields are **not** `omitempty`, so the API gets an explicit value.
- Optional booleans: `omitempty` drops `false`. If `false` is meaningful and differs from
  the server default, use `*bool` or no `omitempty`.
- **Monitors only:**
  - Add the type code to the `MonitorType` constants in `api/constants.go`.
  - Implement the `Site24x7Monitor` interface on the struct (`Get/SetLocationProfileID`,
    `Get/SetNotificationProfileID`, `Get/SetUserGroupIDs`, `Get/SetTagIDs`, `String`),
    copying from an existing monitor. The profile helpers in
    [site24x7/monitor_defaults.go](../../../site24x7/monitor_defaults.go) need it.

## Step 3: Endpoint client (`api/endpoints/<pkg>/<name>_impl.go`)

```go
type Widgets interface {
	Get(widgetID string) (*api.Widget, error)
	Create(widget *api.Widget) (*api.Widget, error)
	Update(widget *api.Widget) (*api.Widget, error)
	Delete(widgetID string) error
	List() ([]*api.Widget, error)
}

type widgets struct{ client rest.Client }

func NewWidgets(client rest.Client) Widgets { return &widgets{client: client} }

func (c *widgets) Create(widget *api.Widget) (*api.Widget, error) {
	created := &api.Widget{}
	err := c.client.
		Post().
		Resource("widget").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(widget).
		Do().
		Parse(created)
	return created, err
}
```

- `Parse` unwraps the `{"code":0,"message":"success","data":…}` envelope for you. Pass the
  inner type.
- `Delete` returns `.Do().Err()`, with no `Parse`.
- Add `List()` even if the resource doesn't need it. Data sources and import tooling will.

## Step 4: Fixtures and endpoint test

Fixtures live in **one shared directory**, whatever the package:
`api/endpoints/testdata/fixtures/requests/` and `.../responses/`. `validation.Fixture`
resolves paths relative to it. Response fixtures carry the full envelope:

```json
{ "code": 0, "message": "success", "data": { "widget_id": "123", "widget_name": "w" } }
```

Test file: `api/endpoints/<pkg>/<name>_impl_test.go`. Use `validation.RunTests` with one
`EndpointTest` per verb, and assert `ExpectedVerb`, `ExpectedPath` and `ExpectedBody`. The
path assertion is what catches singular/plural mistakes, so never leave it loose.

## Step 5: Fake (`api/endpoints/fake/<name>.go`)

A testify mock with a compile-time assertion. Mirror every interface method:

```go
var _ common.Widgets = &Widgets{}

type Widgets struct{ mock.Mock }

func (e *Widgets) Get(widgetID string) (*api.Widget, error) {
	args := e.Called(widgetID)
	if obj, ok := args.Get(0).(*api.Widget); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
```

The `ok` type check matters. Tests return `nil` for the object on error paths, and a direct
type assertion would panic.

## Step 6: Wire the client in two places

1. **[site24x7/client.go](../../../site24x7/client.go)**: add the method to the `Client`
   interface, and an implementation `func (c *client) Widgets() common.Widgets { return
   common.NewWidgets(c.restClient) }`.
2. **[fake/client.go](../../../fake/client.go)**: add three things, or `fake.Client` stops
   satisfying `site24x7.Client` and every resource test in the repo fails to compile:
   - a `FakeWidgets *fake.Widgets` field
   - its initialiser in `NewClient()`
   - the accessor method

## Step 7: Resource (`site24x7/<pkg>/<name>.go`)

The skeleton, following the conventions the repo has converged on:

```go
var widgetSchema = map[string]*schema.Schema{ /* Step 8 */ }

func ResourceSite24x7Widget() *schema.Resource {
	return &schema.Resource{
		Create: resourceSite24x7WidgetCreate,
		Read:   resourceSite24x7WidgetRead,
		Update: resourceSite24x7WidgetUpdate,
		Delete: resourceSite24x7WidgetDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: widgetSchema,
	}
}

func resourceSite24x7WidgetRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	widget, err := client.Widgets().Get(d.Id())
	if err != nil {
		// Deleted outside Terraform: drop from state so the next plan recreates it
		// instead of failing the refresh.
		if apierrors.IsNotFound(err) {
			log.Warnf("Widget %s no longer exists, removing it from state", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}

	setWidgetResourceData(d, widget)
	return nil
}

func resourceSite24x7WidgetDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	err := client.Widgets().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}
	return err
}
```

Non-negotiables. Each has a known offender in the repo:

- **Always add `Importer`.** `amazon`, `azure`, `gcp` and `milestone_marker` have none,
  and users can't adopt them. `ImportStatePassthrough` is
  enough whenever Read rebuilds the full state from the ID.
- **Read handles 404 by clearing the ID.** `credential_profiles`, `oauth2_provider`,
  `sla_setting`, `attribute_alert_group` and `milestone_marker` don't, so a resource deleted
  in the UI breaks every later `terraform plan`.
- **Delete treats 404 as success**, so `terraform destroy` doesn't fail on already-gone
  objects.
- **Create calls `d.SetId` and then refreshes state**, either by calling Read or by setting
  state from the create response. `tags.go` does neither (its `return tagRead(...)` is
  commented out), so computed values stay empty until the next refresh.
- **Update sets the ID on the request object** (`widget.WidgetID = d.Id()`). The expand
  function builds from config, which carries no ID.
- **Check every `d.Set` value's type.** SDK v1 returns an error from `d.Set` that this repo
  ignores everywhere, so a wrong type silently does nothing. `website.go` has
  `d.Set("threshold_profile_id", profile)` with a struct, and it never takes effect.
- **Logging:** `log "github.com/sirupsen/logrus"`.
- **Validation:** `github.com/hashicorp/terraform-plugin-sdk/helper/validation`
  (`IntBetween`, `StringInSlice` …) on enums and ranges. Plan-time errors beat 400s at
  apply time.
- **Monitors:** use `site24x7.SetLocationProfile`, `SetNotificationProfile`,
  `DefaultThresholdProfile`, `SetUserGroup` and `SetTags` from `monitor_defaults.go` rather
  than reimplementing profile resolution. Offer both `location_profile_id` and
  `location_profile_name` (both `Optional + Computed`), like the other monitors. The
  substring-matching caveats of those helpers are documented in `site24x7-monitor-authoring`.

## Step 8: Schema decisions that prevent drift

Get these right at design time. Every wrong call here becomes a ticket for the
`site24x7-drift-troubleshooting` skill.

| Situation | Schema |
|---|---|
| User must provide | `Required: true` |
| Optional, server fills a default when omitted | `Optional: true, Computed: true`, and no `Default` |
| Optional, default is fixed and documented | `Optional: true, Default: <value>`, where the value is **exactly** what the API returns (same type and casing) |
| Server-generated (IDs, timestamps, tokens) | `Computed: true` only |
| Unordered ID collection (groups, tags, user groups, monitors) | `TypeSet` of `TypeString` |
| Order is meaningful (steps, ordered rules), or the API preserves order | `TypeList` |
| Single nested object | `TypeList` + `MaxItems: 1` + `Elem: &schema.Resource{…}` |
| Free-form key/value | `TypeMap` |

In the expand function, read a `TypeSet` with `d.Get("x").(*schema.Set).List()`. Writing a
`[]string` back with `d.Set` works fine.

### Secrets: pick exactly one pattern

First find out, in Step 0, what the API returns for the field. Then:

| API returns | Pattern |
|---|---|
| The plaintext value | `Sensitive: true`, and set it in Read as normal |
| Nothing, masked (`****`), or encrypted (`ENC(...)`) | `Sensitive: true`, and **don't `d.Set` it in Read**. State keeps the last applied value, and a changed config value produces a real diff. |

**Don't** add `DiffSuppressFunc: func(...) bool { return true }`. `website.go`
(`auth_pass`, `client_certificate_password`) and `webhook.go` (`password`) do, and the
result is that **rotating the secret in config never reaches the API**: the diff is
suppressed, so Update is never called. It only goes out if some other attribute changes in
the same apply.

Every secret must be `Sensitive: true`. That covers the data source schema too, which is
easy to forget.

## Step 9: Resource unit tests (`site24x7/<pkg>/<name>_test.go`)

The pattern uses the fake client and calls the CRUD functions directly:

```go
func TestWidgetCreate(t *testing.T) {
	d := widgetTestResourceData(t)
	c := fake.NewClient()

	expected := &api.Widget{WidgetName: "w"}
	c.FakeWidgets.On("Create", expected).Return(&api.Widget{WidgetID: "123", WidgetName: "w"}, nil).Once()
	c.FakeWidgets.On("Get", "123").Return(&api.Widget{WidgetID: "123", WidgetName: "w"}, nil).Once()

	require.NoError(t, resourceSite24x7WidgetCreate(d, c))
	assert.Equal(t, "123", d.Id())
}
```

Cover these at minimum:

- Create, Update and Delete, each for success and a 500.
- Read success.
- **Read on 404 clears the ID.**
- **Delete on 404 returns nil.**
- For secrets: **Read doesn't overwrite them**.

Test gotchas:

- `schema.TestResourceDataRaw` **doesn't apply schema defaults**. Set every defaulted
  attribute explicitly in the helper, or the expanded struct won't match the mock
  expectation.
- Mock arguments are compared with `ObjectsAreEqual`. The expected struct must match what
  the expand function builds **field for field**, including empty slices vs `nil`.
- Create calling Read means the test needs a `Get` expectation too.
- Name tests after the resource. Grep your file for the name of the resource you copied
  from before committing.

## Step 10: Register in [provider/provider.go](../../../provider/provider.go)

Add `"site24x7_widget": common.ResourceSite24x7Widget(),` to `ResourcesMap`, and
`DataSourcesMap` if relevant. Keep the column alignment (`gofmt` handles it). A resource
that isn't registered compiles, passes every unit test, and doesn't exist for users, so this
is the step most often missed.

## Step 11: Data source (optional, but usually wanted)

Follow the `name_regex` pattern: list all, filter with `regexp.MustCompile`, `d.SetId` on
the match, and fail with a clear `Unable to find ... matching the name` error.

**Stop at the first match, or error when there are several.** Only `monitor_group` and
`subgroup` break out of the loop. The other `name_regex` data sources silently return the
*last* match, which varies with API ordering.

## Step 12: Docs and example

- `docs/resources/<name>.md`, or `docs/data-sources/<name>.md`. **Docs are hand-written, not
  generated.** Copy the front-matter block (`layout`, `page_title`, `sidebar_current`,
  `description`) from a sibling. Include an Example Usage block and an Attributes Reference
  split into Required / Optional / Read-Only, and end with the API doc link. Several
  existing resources have no doc page (`oauth2_provider`, `sla_setting`,
  `milestone_marker`, `attribute_alert_group`, `schedule_report`). Don't add another.
- `examples/<name>_us.tf`, or `examples/data-sources/<name>_data_source_us.tf`. Start with
  the standard `terraform {}` and `provider "site24x7" {}` header from any sibling example.
  Annotate each attribute with `// (Required)` or `// (Optional)`. **No real credentials**.
  Use placeholders or variables for secrets.

## Step 13: Keep the other skills true

The user-facing skills state counts and lists that change with each resource. Update:

- `site24x7-alerting-and-profiles`: the resource/data-source map, the destroy-semantics
  list, and the secrets table if the resource has credentials.
- `site24x7-import-existing`: the importable count ("42 of the 46"), the no-importer list,
  and the dependency order.
- `site24x7-monitor-authoring`: the resource table and type codes, for a monitor.
- `site24x7-drift-troubleshooting`: its offender lists, if you fixed one.

Then run the name-validation snippet in [.claude/skills/README.md](../README.md).

## Verify

Run these from the repo root, in order. Fix each failure before moving on:

```sh
gofmt -l ./api ./site24x7 ./fake ./provider   # must print nothing
go build ./...
go vet ./...
go test ./api/... ./site24x7/... ./fake/... ./provider/...
```

`make test` works too, but the Makefile is written for Windows `cmd` (`findstr`), so call
`go test` directly from Git Bash or a POSIX shell.

Then run an end-to-end check against a real account with the locally built provider (`make
install`, or a `dev_overrides` block in `~/.terraformrc`):

1. `terraform apply` the example. It must create cleanly.
2. `terraform plan` again. It **must say "No changes."** A diff here is drift you built
   in, so go back to Step 8.
3. Change one attribute and apply. Only that attribute should change.
4. `terraform state rm` the resource, re-import it with an `import` block, and plan. It
   must be clean apart from any write-only secrets.
5. Delete the object in the Site24x7 UI and plan. It must propose a **create**, not fail.
6. `terraform destroy`, then destroy again after deleting in the UI. Neither may error.

## Done checklist

- [ ] API type, endpoint client, fixtures and endpoint test (verb, path, body asserted)
- [ ] Fake with `var _ Interface = &Fake{}`, and `fake/client.go` updated in all 3 spots
- [ ] `site24x7/client.go` interface and implementation
- [ ] Resource with `Importer`, 404-tolerant Read and Delete, and Create that refreshes state
- [ ] Secrets are `Sensitive`, not set from the API when write-only, and have no always-true `DiffSuppressFunc`
- [ ] Unordered ID lists are `TypeSet`; server defaults are `Optional + Computed`
- [ ] Unit tests, including the 404 cases
- [ ] Registered in `provider/provider.go`
- [ ] Doc page and example file
- [ ] Other skills updated
- [ ] `gofmt`, `go build`, `go vet` and `go test` are clean; a second `terraform plan` shows no changes
