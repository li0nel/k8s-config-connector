// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceNetworkSecurityTlsInspectionPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkSecurityTlsInspectionPolicyCreate,
		Read:   resourceNetworkSecurityTlsInspectionPolicyRead,
		Update: resourceNetworkSecurityTlsInspectionPolicyUpdate,
		Delete: resourceNetworkSecurityTlsInspectionPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNetworkSecurityTlsInspectionPolicyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"ca_pool": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `A CA pool resource used to issue interception certificates.`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Short name of the TlsInspectionPolicy resource to be created.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Free-text description of the resource.`,
			},
			"exclude_public_ca_set": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `If FALSE (the default), use our default set of public CAs in addition to any CAs specified in trustConfig. These public CAs are currently based on the Mozilla Root Program and are subject to change over time. If TRUE, do not accept our default set of public CAs. Only CAs specified in trustConfig will be accepted.`,
			},
			"location": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The location of the tls inspection policy.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp when the resource was created.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp when the resource was updated.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceNetworkSecurityTlsInspectionPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkSecurityTlsInspectionPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	caPoolProp, err := expandNetworkSecurityTlsInspectionPolicyCaPool(d.Get("ca_pool"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ca_pool"); !tpgresource.IsEmptyValue(reflect.ValueOf(caPoolProp)) && (ok || !reflect.DeepEqual(v, caPoolProp)) {
		obj["caPool"] = caPoolProp
	}
	excludePublicCaSetProp, err := expandNetworkSecurityTlsInspectionPolicyExcludePublicCaSet(d.Get("exclude_public_ca_set"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("exclude_public_ca_set"); !tpgresource.IsEmptyValue(reflect.ValueOf(excludePublicCaSetProp)) && (ok || !reflect.DeepEqual(v, excludePublicCaSetProp)) {
		obj["excludePublicCaSet"] = excludePublicCaSetProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/tlsInspectionPolicies?tlsInspectionPolicyId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new TlsInspectionPolicy: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TlsInspectionPolicy: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating TlsInspectionPolicy: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/tlsInspectionPolicies/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = NetworkSecurityOperationWaitTime(
		config, res, project, "Creating TlsInspectionPolicy", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create TlsInspectionPolicy: %s", err)
	}

	log.Printf("[DEBUG] Finished creating TlsInspectionPolicy %q: %#v", d.Id(), res)

	return resourceNetworkSecurityTlsInspectionPolicyRead(d, meta)
}

func resourceNetworkSecurityTlsInspectionPolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/tlsInspectionPolicies/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TlsInspectionPolicy: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("NetworkSecurityTlsInspectionPolicy %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading TlsInspectionPolicy: %s", err)
	}

	if err := d.Set("create_time", flattenNetworkSecurityTlsInspectionPolicyCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading TlsInspectionPolicy: %s", err)
	}
	if err := d.Set("update_time", flattenNetworkSecurityTlsInspectionPolicyUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading TlsInspectionPolicy: %s", err)
	}
	if err := d.Set("description", flattenNetworkSecurityTlsInspectionPolicyDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading TlsInspectionPolicy: %s", err)
	}
	if err := d.Set("ca_pool", flattenNetworkSecurityTlsInspectionPolicyCaPool(res["caPool"], d, config)); err != nil {
		return fmt.Errorf("Error reading TlsInspectionPolicy: %s", err)
	}
	if err := d.Set("exclude_public_ca_set", flattenNetworkSecurityTlsInspectionPolicyExcludePublicCaSet(res["excludePublicCaSet"], d, config)); err != nil {
		return fmt.Errorf("Error reading TlsInspectionPolicy: %s", err)
	}

	return nil
}

func resourceNetworkSecurityTlsInspectionPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TlsInspectionPolicy: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkSecurityTlsInspectionPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	caPoolProp, err := expandNetworkSecurityTlsInspectionPolicyCaPool(d.Get("ca_pool"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ca_pool"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, caPoolProp)) {
		obj["caPool"] = caPoolProp
	}
	excludePublicCaSetProp, err := expandNetworkSecurityTlsInspectionPolicyExcludePublicCaSet(d.Get("exclude_public_ca_set"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("exclude_public_ca_set"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, excludePublicCaSetProp)) {
		obj["excludePublicCaSet"] = excludePublicCaSetProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/tlsInspectionPolicies/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating TlsInspectionPolicy %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("ca_pool") {
		updateMask = append(updateMask, "caPool")
	}

	if d.HasChange("exclude_public_ca_set") {
		updateMask = append(updateMask, "excludePublicCaSet")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating TlsInspectionPolicy %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating TlsInspectionPolicy %q: %#v", d.Id(), res)
	}

	err = NetworkSecurityOperationWaitTime(
		config, res, project, "Updating TlsInspectionPolicy", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceNetworkSecurityTlsInspectionPolicyRead(d, meta)
}

func resourceNetworkSecurityTlsInspectionPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TlsInspectionPolicy: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/tlsInspectionPolicies/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting TlsInspectionPolicy %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "TlsInspectionPolicy")
	}

	err = NetworkSecurityOperationWaitTime(
		config, res, project, "Deleting TlsInspectionPolicy", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting TlsInspectionPolicy %q: %#v", d.Id(), res)
	return nil
}

func resourceNetworkSecurityTlsInspectionPolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/tlsInspectionPolicies/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/tlsInspectionPolicies/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNetworkSecurityTlsInspectionPolicyCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityTlsInspectionPolicyUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityTlsInspectionPolicyDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityTlsInspectionPolicyCaPool(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityTlsInspectionPolicyExcludePublicCaSet(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNetworkSecurityTlsInspectionPolicyDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityTlsInspectionPolicyCaPool(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityTlsInspectionPolicyExcludePublicCaSet(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
