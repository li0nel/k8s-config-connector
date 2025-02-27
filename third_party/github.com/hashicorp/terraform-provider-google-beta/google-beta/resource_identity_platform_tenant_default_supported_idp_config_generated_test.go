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
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccIdentityPlatformTenantDefaultSupportedIdpConfig_identityPlatformTenantDefaultSupportedIdpConfigBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckIdentityPlatformTenantDefaultSupportedIdpConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIdentityPlatformTenantDefaultSupportedIdpConfig_identityPlatformTenantDefaultSupportedIdpConfigBasicExample(context),
			},
			{
				ResourceName:            "google_identity_platform_tenant_default_supported_idp_config.idp_config",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"idp_id", "tenant"},
			},
		},
	})
}

func testAccIdentityPlatformTenantDefaultSupportedIdpConfig_identityPlatformTenantDefaultSupportedIdpConfigBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_identity_platform_tenant" "tenant" {
  display_name  = "tenant"
}

resource "google_identity_platform_tenant_default_supported_idp_config" "idp_config" {
  enabled       = true
  tenant        = google_identity_platform_tenant.tenant.name
  idp_id        = "playgames.google.com"
  client_id     = "my-client-id"
  client_secret = "secret"
}
`, context)
}

func testAccCheckIdentityPlatformTenantDefaultSupportedIdpConfigDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_identity_platform_tenant_default_supported_idp_config" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{IdentityPlatformBasePath}}projects/{{project}}/tenants/{{tenant}}/defaultSupportedIdpConfigs/{{idp_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("IdentityPlatformTenantDefaultSupportedIdpConfig still exists at %s", url)
			}
		}

		return nil
	}
}
