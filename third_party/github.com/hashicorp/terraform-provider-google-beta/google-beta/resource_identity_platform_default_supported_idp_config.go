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

func ResourceIdentityPlatformDefaultSupportedIdpConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceIdentityPlatformDefaultSupportedIdpConfigCreate,
		Read:   resourceIdentityPlatformDefaultSupportedIdpConfigRead,
		Update: resourceIdentityPlatformDefaultSupportedIdpConfigUpdate,
		Delete: resourceIdentityPlatformDefaultSupportedIdpConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceIdentityPlatformDefaultSupportedIdpConfigImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"client_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `OAuth client ID`,
			},
			"client_secret": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `OAuth client secret`,
			},
			"idp_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `ID of the IDP. Possible values include:

* 'apple.com'

* 'facebook.com'

* 'gc.apple.com'

* 'github.com'

* 'google.com'

* 'linkedin.com'

* 'microsoft.com'

* 'playgames.google.com'

* 'twitter.com'

* 'yahoo.com'`,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `If this IDP allows the user to sign in`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The name of the DefaultSupportedIdpConfig resource`,
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

func resourceIdentityPlatformDefaultSupportedIdpConfigCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	clientIdProp, err := expandIdentityPlatformDefaultSupportedIdpConfigClientId(d.Get("client_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("client_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(clientIdProp)) && (ok || !reflect.DeepEqual(v, clientIdProp)) {
		obj["clientId"] = clientIdProp
	}
	clientSecretProp, err := expandIdentityPlatformDefaultSupportedIdpConfigClientSecret(d.Get("client_secret"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("client_secret"); !tpgresource.IsEmptyValue(reflect.ValueOf(clientSecretProp)) && (ok || !reflect.DeepEqual(v, clientSecretProp)) {
		obj["clientSecret"] = clientSecretProp
	}
	enabledProp, err := expandIdentityPlatformDefaultSupportedIdpConfigEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(enabledProp)) && (ok || !reflect.DeepEqual(v, enabledProp)) {
		obj["enabled"] = enabledProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/defaultSupportedIdpConfigs?idpId={{idp_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new DefaultSupportedIdpConfig: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DefaultSupportedIdpConfig: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating DefaultSupportedIdpConfig: %s", err)
	}
	if err := d.Set("name", flattenIdentityPlatformDefaultSupportedIdpConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/defaultSupportedIdpConfigs/{{idp_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating DefaultSupportedIdpConfig %q: %#v", d.Id(), res)

	return resourceIdentityPlatformDefaultSupportedIdpConfigRead(d, meta)
}

func resourceIdentityPlatformDefaultSupportedIdpConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/defaultSupportedIdpConfigs/{{idp_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DefaultSupportedIdpConfig: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("IdentityPlatformDefaultSupportedIdpConfig %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading DefaultSupportedIdpConfig: %s", err)
	}

	if err := d.Set("name", flattenIdentityPlatformDefaultSupportedIdpConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading DefaultSupportedIdpConfig: %s", err)
	}
	if err := d.Set("client_id", flattenIdentityPlatformDefaultSupportedIdpConfigClientId(res["clientId"], d, config)); err != nil {
		return fmt.Errorf("Error reading DefaultSupportedIdpConfig: %s", err)
	}
	if err := d.Set("client_secret", flattenIdentityPlatformDefaultSupportedIdpConfigClientSecret(res["clientSecret"], d, config)); err != nil {
		return fmt.Errorf("Error reading DefaultSupportedIdpConfig: %s", err)
	}
	if err := d.Set("enabled", flattenIdentityPlatformDefaultSupportedIdpConfigEnabled(res["enabled"], d, config)); err != nil {
		return fmt.Errorf("Error reading DefaultSupportedIdpConfig: %s", err)
	}

	return nil
}

func resourceIdentityPlatformDefaultSupportedIdpConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DefaultSupportedIdpConfig: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	clientIdProp, err := expandIdentityPlatformDefaultSupportedIdpConfigClientId(d.Get("client_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("client_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, clientIdProp)) {
		obj["clientId"] = clientIdProp
	}
	clientSecretProp, err := expandIdentityPlatformDefaultSupportedIdpConfigClientSecret(d.Get("client_secret"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("client_secret"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, clientSecretProp)) {
		obj["clientSecret"] = clientSecretProp
	}
	enabledProp, err := expandIdentityPlatformDefaultSupportedIdpConfigEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enabledProp)) {
		obj["enabled"] = enabledProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/defaultSupportedIdpConfigs/{{idp_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating DefaultSupportedIdpConfig %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("client_id") {
		updateMask = append(updateMask, "clientId")
	}

	if d.HasChange("client_secret") {
		updateMask = append(updateMask, "clientSecret")
	}

	if d.HasChange("enabled") {
		updateMask = append(updateMask, "enabled")
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
		return fmt.Errorf("Error updating DefaultSupportedIdpConfig %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating DefaultSupportedIdpConfig %q: %#v", d.Id(), res)
	}

	return resourceIdentityPlatformDefaultSupportedIdpConfigRead(d, meta)
}

func resourceIdentityPlatformDefaultSupportedIdpConfigDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DefaultSupportedIdpConfig: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/defaultSupportedIdpConfigs/{{idp_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting DefaultSupportedIdpConfig %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "DefaultSupportedIdpConfig")
	}

	log.Printf("[DEBUG] Finished deleting DefaultSupportedIdpConfig %q: %#v", d.Id(), res)
	return nil
}

func resourceIdentityPlatformDefaultSupportedIdpConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/defaultSupportedIdpConfigs/(?P<idp_id>[^/]+)",
		"(?P<project>[^/]+)/(?P<idp_id>[^/]+)",
		"(?P<idp_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/defaultSupportedIdpConfigs/{{idp_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenIdentityPlatformDefaultSupportedIdpConfigName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformDefaultSupportedIdpConfigClientId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformDefaultSupportedIdpConfigClientSecret(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformDefaultSupportedIdpConfigEnabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandIdentityPlatformDefaultSupportedIdpConfigClientId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformDefaultSupportedIdpConfigClientSecret(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformDefaultSupportedIdpConfigEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
