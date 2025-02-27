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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceNetworkServicesServiceBinding() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkServicesServiceBindingCreate,
		Read:   resourceNetworkServicesServiceBindingRead,
		Delete: resourceNetworkServicesServiceBindingDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNetworkServicesServiceBindingImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Name of the ServiceBinding resource.`,
			},
			"service": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The full Service Directory Service name of the format
projects/*/locations/*/namespaces/*/services/*`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `A free-text description of the resource. Max length 1024 characters.`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: `Set of label tags associated with the ServiceBinding resource.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the ServiceBinding was created in UTC.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the ServiceBinding was updated in UTC.`,
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

func resourceNetworkServicesServiceBindingCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	labelsProp, err := expandNetworkServicesServiceBindingLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	descriptionProp, err := expandNetworkServicesServiceBindingDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	serviceProp, err := expandNetworkServicesServiceBindingService(d.Get("service"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("service"); !tpgresource.IsEmptyValue(reflect.ValueOf(serviceProp)) && (ok || !reflect.DeepEqual(v, serviceProp)) {
		obj["service"] = serviceProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/global/serviceBindings?serviceBindingId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ServiceBinding: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ServiceBinding: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating ServiceBinding: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/serviceBindings/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = NetworkServicesOperationWaitTime(
		config, res, project, "Creating ServiceBinding", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create ServiceBinding: %s", err)
	}

	log.Printf("[DEBUG] Finished creating ServiceBinding %q: %#v", d.Id(), res)

	return resourceNetworkServicesServiceBindingRead(d, meta)
}

func resourceNetworkServicesServiceBindingRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/global/serviceBindings/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ServiceBinding: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("NetworkServicesServiceBinding %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ServiceBinding: %s", err)
	}

	if err := d.Set("create_time", flattenNetworkServicesServiceBindingCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceBinding: %s", err)
	}
	if err := d.Set("update_time", flattenNetworkServicesServiceBindingUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceBinding: %s", err)
	}
	if err := d.Set("labels", flattenNetworkServicesServiceBindingLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceBinding: %s", err)
	}
	if err := d.Set("description", flattenNetworkServicesServiceBindingDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceBinding: %s", err)
	}
	if err := d.Set("service", flattenNetworkServicesServiceBindingService(res["service"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceBinding: %s", err)
	}

	return nil
}

func resourceNetworkServicesServiceBindingDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ServiceBinding: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/global/serviceBindings/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting ServiceBinding %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "ServiceBinding")
	}

	err = NetworkServicesOperationWaitTime(
		config, res, project, "Deleting ServiceBinding", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting ServiceBinding %q: %#v", d.Id(), res)
	return nil
}

func resourceNetworkServicesServiceBindingImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/global/serviceBindings/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/serviceBindings/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNetworkServicesServiceBindingCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesServiceBindingUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesServiceBindingLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesServiceBindingDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesServiceBindingService(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNetworkServicesServiceBindingLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandNetworkServicesServiceBindingDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesServiceBindingService(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
