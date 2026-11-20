package monitors

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	log "github.com/sirupsen/logrus"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

var databricksWorkspaceURIRegexp = regexp.MustCompile(`^https://\S+$`)

// defaultDatabricksMonitoredService is the service code the Site24x7 UI sends when a
// Databricks workspace is added.
const defaultDatabricksMonitoredService = 5002

// Attributes whose change must be flagged to the API with update_auth_details=true,
// otherwise the server keeps the previously validated credentials.
var databricksAuthAttributes = []string{"workspace_uri", "access_key", "secret_key", "auth_type", "token_profile"}

var DatabricksMonitorSchema = map[string]*schema.Schema{
	"display_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name for the Databricks monitor.",
	},
	"workspace_uri": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringMatch(databricksWorkspaceURIRegexp, "must be an https:// URL, e.g. https://adb-1234567890123456.7.azuredatabricks.net/"),
		Description:  "URL of the Databricks workspace, e.g. https://adb-1234567890123456.7.azuredatabricks.net/.",
	},
	"access_key": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Client ID of the Databricks service principal.",
	},
	"secret_key": {
		Type:        schema.TypeString,
		Required:    true,
		Sensitive:   true,
		Description: "OAuth secret of the Databricks service principal. Write-only: it is never read back from Site24x7.",
	},
	"auth_type": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     4,
		Description: "Authentication type. 4 is OAuth machine-to-machine (service principal).",
	},
	"token_profile": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "databricks_oidc",
		Description: "Token profile used to obtain the access token.",
	},
	"check_frequency": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "30",
		Description: "Polling interval in minutes.",
	},
	"monitored_services": {
		Type:        schema.TypeSet,
		Elem:        &schema.Schema{Type: schema.TypeInt},
		Optional:    true,
		Computed:    true,
		Description: "Databricks service codes to discover. Defaults to [5002] when omitted.",
	},
	"disable_discovery": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     false,
		Description: "Stop auto-discovering new Databricks resources.",
	},
	"threshold_profile_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Threshold profile to be associated with the monitor.",
	},
	"notification_profile_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Notification profile to be associated with the monitor.",
	},
	"notification_profile_name": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Name of the notification profile to be associated with the monitor.",
	},
	"user_group_ids": {
		Type:        schema.TypeSet,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Optional:    true,
		Computed:    true,
		Description: "User group IDs to be notified on down.",
	},
	"user_group_names": {
		Type:        schema.TypeList,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Optional:    true,
		Description: "Names of the user groups to be notified on down.",
	},
	"tag_ids": {
		Type:        schema.TypeSet,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Optional:    true,
		Computed:    true,
		Description: "Tag IDs to be associated with the monitor.",
	},
	"tag_names": {
		Type:        schema.TypeList,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Optional:    true,
		Description: "Names of the tags to be associated with the monitor.",
	},
}

func ResourceSite24x7DatabricksMonitor() *schema.Resource {
	return &schema.Resource{
		Create: databricksMonitorCreate,
		Read:   databricksMonitorRead,
		Update: databricksMonitorUpdate,
		Delete: databricksMonitorDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: DatabricksMonitorSchema,
	}
}

func databricksMonitorCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	monitor, err := resourceDataToDatabricksMonitor(d, client)
	if err != nil {
		return err
	}

	created, err := client.DatabricksMonitors().Create(monitor)
	if err != nil {
		return err
	}

	d.SetId(created.MonitorID)
	return databricksMonitorRead(d, meta)
}

func databricksMonitorRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	monitor, err := client.DatabricksMonitors().Get(d.Id())
	if err != nil {
		// Deleted outside Terraform: drop from state so the next plan recreates it
		// instead of failing the refresh.
		if apierrors.IsNotFound(err) {
			log.Warnf("Databricks monitor %s no longer exists, removing it from state", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}

	updateDatabricksMonitorResourceData(d, monitor)
	return nil
}

func databricksMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	monitor, err := resourceDataToDatabricksMonitor(d, client)
	if err != nil {
		return err
	}

	if _, err := client.DatabricksMonitors().Update(monitor); err != nil {
		return err
	}

	return databricksMonitorRead(d, meta)
}

func databricksMonitorDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	err := client.DatabricksMonitors().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}
	return err
}

func resourceDataToDatabricksMonitor(d *schema.ResourceData, client site24x7.Client) (*api.DatabricksMonitor, error) {
	userGroupIDs := []string{}
	for _, id := range d.Get("user_group_ids").(*schema.Set).List() {
		if id != nil {
			userGroupIDs = append(userGroupIDs, id.(string))
		}
	}

	tagIDs := []string{}
	for _, id := range d.Get("tag_ids").(*schema.Set).List() {
		if id != nil {
			tagIDs = append(tagIDs, id.(string))
		}
	}

	monitoredServices := []int{}
	for _, s := range d.Get("monitored_services").(*schema.Set).List() {
		monitoredServices = append(monitoredServices, s.(int))
	}
	if len(monitoredServices) == 0 {
		monitoredServices = []int{defaultDatabricksMonitoredService}
	}

	authType := d.Get("auth_type").(int)
	disableDiscovery := 0
	if d.Get("disable_discovery").(bool) {
		disableDiscovery = 1
	}

	// Credentials are re-validated by the server only when this flag is set, so a
	// rotated secret must raise it. It is always false on create.
	updateAuthDetails := false
	if d.Id() != "" {
		for _, attr := range databricksAuthAttributes {
			if d.HasChange(attr) {
				updateAuthDetails = true
				break
			}
		}
	}

	monitor := &api.DatabricksMonitor{
		MonitorID:             d.Id(),
		DisplayName:           d.Get("display_name").(string),
		Type:                  string(api.DBRICKS_WS),
		CheckFrequency:        d.Get("check_frequency").(string),
		AuthType:              &authType,
		TokenProfile:          d.Get("token_profile").(string),
		WorkspaceURI:          d.Get("workspace_uri").(string),
		AccessKey:             d.Get("access_key").(string),
		SecretKey:             d.Get("secret_key").(string),
		MonitoredServices:     monitoredServices,
		DisableDiscovery:      &disableDiscovery,
		ThresholdProfileID:    d.Get("threshold_profile_id").(string),
		NotificationProfileID: d.Get("notification_profile_id").(string),
		UserGroupIDs:          userGroupIDs,
		TagIDs:                tagIDs,
		UpdateAuthDetails:     &updateAuthDetails,
	}

	if _, err := site24x7.SetNotificationProfile(client, d, monitor); err != nil {
		return nil, err
	}

	if _, err := site24x7.SetUserGroup(client, d, monitor); err != nil {
		return nil, err
	}

	if _, err := site24x7.SetTags(client, d, monitor); err != nil {
		return nil, err
	}

	if monitor.ThresholdProfileID == "" {
		profile, err := site24x7.DefaultThresholdProfile(client, api.DBRICKS_WS)
		if err != nil {
			// No DBRICKS_WS threshold profile exists in the account: send none and let
			// the server apply its own default instead of failing the plan here.
			log.Warnf("No default threshold profile for %s: %s", api.DBRICKS_WS, err)
		} else {
			monitor.ThresholdProfileID = profile.ProfileID
			d.Set("threshold_profile_id", profile.ProfileID)
		}
	}

	return monitor, nil
}

func updateDatabricksMonitorResourceData(d *schema.ResourceData, m *api.DatabricksMonitor) {
	d.Set("display_name", m.DisplayName)
	d.Set("threshold_profile_id", m.ThresholdProfileID)
	d.Set("notification_profile_id", m.NotificationProfileID)

	// The fields below are not always echoed by the API. Only overwrite state with
	// values the API actually returned, otherwise every refresh would wipe them and
	// produce a perpetual diff.
	if m.CheckFrequency != "" {
		d.Set("check_frequency", m.CheckFrequency)
	}
	if m.WorkspaceURI != "" {
		d.Set("workspace_uri", m.WorkspaceURI)
	}
	if m.AccessKey != "" {
		d.Set("access_key", m.AccessKey)
	}
	if m.TokenProfile != "" {
		d.Set("token_profile", m.TokenProfile)
	}
	if m.AuthType != nil {
		d.Set("auth_type", *m.AuthType)
	}
	if m.DisableDiscovery != nil {
		d.Set("disable_discovery", *m.DisableDiscovery == 1)
	}
	if m.MonitoredServices != nil {
		d.Set("monitored_services", []int(m.MonitoredServices))
	}
	if m.UserGroupIDs != nil {
		d.Set("user_group_ids", m.UserGroupIDs)
	}
	if m.TagIDs != nil {
		d.Set("tag_ids", m.TagIDs)
	}
	// secret_key is write-only and deliberately never set from the API response.
}
