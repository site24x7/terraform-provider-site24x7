package apm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	log "github.com/sirupsen/logrus"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

// apmAgentConfigProfileSchema flattens the API's nested agent_config object
// into top-level attributes, naming each one after its wire key with the dots
// replaced by underscores - "transaction.trace.enabled" becomes
// "transaction_trace_enabled" - so the mapping stays predictable as the API
// grows new settings.
//
// Every agent_config key except cloud.instance.cleanup.threshold is mandatory
// on create and update, so each one is Optional with a Default rather than
// Required: the defaults mirror the values Site24x7 ships in its own default
// profiles, and a value is always sent.
var apmAgentConfigProfileSchema = map[string]*schema.Schema{
	"profile_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name of the configuration profile.",
	},
	"agent_type": {
		Type:     schema.TypeString,
		Required: true,
		Description: "Type of APM Insight agent the profile configures, for example JAVA, DOTNET, PHP, RUBY, NODEJS or PYTHON. " +
			"Not validated locally, so agent types added by Site24x7 after this provider release can still be used.",
	},
	"is_default": {
		Type:     schema.TypeBool,
		Optional: true,
		Default:  false,
		Description: "Whether this profile is the default configuration for its agent type. " +
			"Only one profile per agent type can be the default, so setting this to true demotes whichever profile was default before, " +
			"a change Terraform cannot see until the other profile is refreshed.",
	},
	"transaction_trace_enabled": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     true,
		Description: "Whether transaction traces are collected. Maps to transaction.trace.enabled.",
	},
	"transaction_trace_threshold": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     2,
		Description: "Tracing threshold in seconds: transactions slower than this are traced. Maps to transaction.trace.threshold.",
	},
	"transaction_trace_sql_parametrize": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     true,
		Description: "Whether SQL queries in traces are obfuscated by replacing literal values with parameters. Maps to transaction.trace.sql.parametrize.",
	},
	"transaction_trace_sql_stacktrace_threshold": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     3,
		Description: "Slow SQL query threshold in seconds: queries slower than this get a stack trace. Maps to transaction.trace.sql.stacktrace.threshold.",
	},
	"transaction_tracking_request_interval": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     1,
		Description: "Web transaction sampling factor: 1 tracks every request, 5 tracks one in five. Maps to transaction.tracking.request.interval.",
	},
	"sql_capture_enabled": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     true,
		Description: "Whether SQL queries are captured. Maps to sql.capture.enabled.",
	},
	"autoupgrade_enabled": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     false,
		Description: "Whether agents using this profile upgrade themselves automatically. Maps to autoupgrade.enabled.",
	},
	"show_instance_port_number": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     true,
		Description: "Whether instance names include the port number. Maps to show.instance.port.number.",
	},
	"apdex_threshold": {
		Type:        schema.TypeFloat,
		Optional:    true,
		Default:     0.5,
		Description: "Apdex threshold in seconds, the response time below which a request counts as satisfying. Maps to apdex.threshold.",
	},
	"cloud_instance_cleanup_threshold": {
		Type:         schema.TypeInt,
		Optional:     true,
		Default:      7,
		ValidateFunc: validation.IntBetween(1, 15),
		Description:  "Number of inactive days after which auto-suspended cloud instances are deleted. Must be between 1 and 15. Maps to cloud.instance.cleanup.threshold.",
	},
	"last_modified_time": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "When the profile was last updated, in milliseconds since the epoch. Maps to last.modified.time.",
	},
}

func ResourceSite24x7APMAgentConfigProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceSite24x7APMAgentConfigProfileCreate,
		Read:   resourceSite24x7APMAgentConfigProfileRead,
		Update: resourceSite24x7APMAgentConfigProfileUpdate,
		Delete: resourceSite24x7APMAgentConfigProfileDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: apmAgentConfigProfileSchema,
	}
}

func resourceSite24x7APMAgentConfigProfileCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	created, err := client.APMAgentConfigProfiles().Create(resourceDataToAPMAgentConfigProfile(d))
	if err != nil {
		return err
	}

	d.SetId(created.ProfileID)
	setAgentConfigProfileData(d, created)

	return nil
}

func resourceSite24x7APMAgentConfigProfileRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	profile, err := client.APMAgentConfigProfiles().Get(d.Id())
	if err != nil {
		// The profile was deleted outside Terraform. Drop it from state so the
		// next plan proposes recreating it rather than failing the refresh.
		if apierrors.IsNotFound(err) {
			log.Warnf("APM agent configuration profile %s no longer exists, removing it from state", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}

	setAgentConfigProfileData(d, profile)

	return nil
}

func resourceSite24x7APMAgentConfigProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	updated, err := client.APMAgentConfigProfiles().Update(resourceDataToAPMAgentConfigProfile(d))
	if err != nil {
		return err
	}

	d.SetId(updated.ProfileID)
	setAgentConfigProfileData(d, updated)

	return nil
}

func resourceSite24x7APMAgentConfigProfileDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	err := client.APMAgentConfigProfiles().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

// resourceDataToAPMAgentConfigProfile builds the request body. LastModifiedTime
// is deliberately left unset: the API maintains it, and echoing a stale value
// back would be at best ignored and at worst wrong.
func resourceDataToAPMAgentConfigProfile(d *schema.ResourceData) *api.APMAgentConfigProfile {
	return &api.APMAgentConfigProfile{
		ProfileID:   d.Id(),
		ProfileName: d.Get("profile_name").(string),
		AgentType:   d.Get("agent_type").(string),
		IsDefault:   d.Get("is_default").(bool),
		AgentConfig: api.APMAgentConfig{
			TransactionTraceEnabled:       d.Get("transaction_trace_enabled").(bool),
			TransactionTraceThreshold:     d.Get("transaction_trace_threshold").(int),
			ParametrizeSQLQuery:           d.Get("transaction_trace_sql_parametrize").(bool),
			SQLStackTraceThreshold:        d.Get("transaction_trace_sql_stacktrace_threshold").(int),
			RequestTrackingInterval:       d.Get("transaction_tracking_request_interval").(int),
			SQLCaptureEnabled:             d.Get("sql_capture_enabled").(bool),
			AutoUpgradeEnabled:            d.Get("autoupgrade_enabled").(bool),
			ShowInstancePortNumber:        d.Get("show_instance_port_number").(bool),
			ApdexThreshold:                d.Get("apdex_threshold").(float64),
			CloudInstanceCleanupThreshold: d.Get("cloud_instance_cleanup_threshold").(int),
		},
	}
}
