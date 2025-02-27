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

func TestAccIdentityPlatformOauthIdpConfig_identityPlatformOauthIdpConfigBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"name":          "oidc.oauth-idp-config-" + RandString(t, 10),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckIdentityPlatformOauthIdpConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIdentityPlatformOauthIdpConfig_identityPlatformOauthIdpConfigBasicExample(context),
			},
			{
				ResourceName:      "google_identity_platform_oauth_idp_config.oauth_idp_config",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccIdentityPlatformOauthIdpConfig_identityPlatformOauthIdpConfigBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_identity_platform_oauth_idp_config" "oauth_idp_config" {
  name          = "%{name}"
  display_name  = "Display Name"
  client_id     = "client-id"
  issuer        = "issuer"
  enabled       = true
  client_secret = "secret"
}
`, context)
}

func testAccCheckIdentityPlatformOauthIdpConfigDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_identity_platform_oauth_idp_config" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{IdentityPlatformBasePath}}projects/{{project}}/oauthIdpConfigs/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("IdentityPlatformOauthIdpConfig still exists at %s", url)
			}
		}

		return nil
	}
}
