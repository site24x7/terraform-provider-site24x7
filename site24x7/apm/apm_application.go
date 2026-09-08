package apm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	log "github.com/sirupsen/logrus"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

func apmApplicationResourceSchema() map[string]*schema.Schema {
	resourceSchema := map[string]*schema.Schema{
		"application_id": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "ID of the APM Insight application to adopt. The APM Insight API has no create endpoint - applications come into existence when an agent first reports in - so this resource takes over an application that already exists.",
		},
		"managed": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Whether the application should be managed (true) or suspended (false). Changing this calls the manage or unmanage endpoint.",
		},
		"delete_on_destroy": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
			Description: "Whether destroying this resource should also delete the application in Site24x7. " +
				"Defaults to false, which only removes the application from Terraform state and leaves it untouched. " +
				"Set to true only if you intend `terraform destroy` to permanently delete the application and its monitoring history, which cannot be undone.",
		},
		"time_window": timeWindowSchema,
	}

	for name, attributeSchema := range applicationComputedSchema() {
		resourceSchema[name] = attributeSchema
	}

	return resourceSchema
}

func ResourceSite24x7APMApplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceSite24x7APMApplicationCreate,
		Read:   resourceSite24x7APMApplicationRead,
		Update: resourceSite24x7APMApplicationUpdate,
		Delete: resourceSite24x7APMApplicationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: apmApplicationResourceSchema(),
	}
}

// resourceSite24x7APMApplicationCreate adopts an existing application rather
// than creating one. It reads the application first so that a typo in
// application_id fails with a clear message instead of a bare 404 from a
// manage call.
func resourceSite24x7APMApplicationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	applicationID := d.Get("application_id").(string)
	timeWindow := d.Get("time_window").(string)

	application, err := client.APMApplications().Get(applicationID, timeWindow)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return fmt.Errorf(
				"no APM Insight application found with application_id %q. This resource adopts an application that already exists; deploy an APM Insight agent and let it report in first",
				applicationID,
			)
		}
		return err
	}

	desiredManaged := d.Get("managed").(bool)
	if err := reconcileManagedState(client, applicationID, desiredManaged, application.AvailabilityHealthInfo.ManagedState); err != nil {
		return err
	}

	d.SetId(applicationID)

	return resourceSite24x7APMApplicationRead(d, meta)
}

func resourceSite24x7APMApplicationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	timeWindow := d.Get("time_window").(string)

	application, err := client.APMApplications().Get(d.Id(), timeWindow)
	if err != nil {
		// The application was deleted outside Terraform. Drop it from state so
		// that the next plan proposes adopting it again rather than failing.
		if apierrors.IsNotFound(err) {
			log.Warnf("APM application %s no longer exists, removing it from state", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("application_id", application.ApplicationInfo.ApplicationID)
	d.Set("managed", application.AvailabilityHealthInfo.ManagedState)
	// Recorded so that an imported resource, which starts with no time_window in
	// state, does not leave "" behind for the next plan to correct.
	d.Set("time_window", timeWindowOrDefault(timeWindow))
	setApplicationData(d, application)

	return nil
}

func resourceSite24x7APMApplicationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	// application_id is ForceNew, and delete_on_destroy and time_window are
	// local-only, so managed is the only attribute with a remote effect.
	if d.HasChange("managed") {
		currentManaged, desiredManaged := d.GetChange("managed")
		if err := reconcileManagedState(client, d.Id(), desiredManaged.(bool), currentManaged.(bool)); err != nil {
			return err
		}
	}

	return resourceSite24x7APMApplicationRead(d, meta)
}

func resourceSite24x7APMApplicationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	if !d.Get("delete_on_destroy").(bool) {
		log.Infof(
			"Removing APM application %s from Terraform state without deleting it, because delete_on_destroy is false",
			d.Id(),
		)
		d.SetId("")
		return nil
	}

	log.Warnf("Permanently deleting APM application %s and its monitoring history", d.Id())

	err := client.APMApplications().Delete(d.Id())
	if err != nil && !apierrors.IsNotFound(err) {
		return err
	}

	d.SetId("")
	return nil
}

func reconcileManagedState(client site24x7.Client, applicationID string, desiredManaged, currentManaged bool) error {
	if desiredManaged == currentManaged {
		return nil
	}

	if desiredManaged {
		return client.APMApplications().Manage(applicationID)
	}

	return client.APMApplications().Unmanage(applicationID)
}
