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

func TestAccCertificateManagerDnsAuthorization_certificateManagerDnsAuthorizationBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCertificateManagerDnsAuthorizationDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCertificateManagerDnsAuthorization_certificateManagerDnsAuthorizationBasicExample(context),
			},
			{
				ResourceName:            "google_certificate_manager_dns_authorization.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name"},
			},
		},
	})
}

func testAccCertificateManagerDnsAuthorization_certificateManagerDnsAuthorizationBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_certificate_manager_dns_authorization" "default" {
  name        = "tf-test-dns-auth%{random_suffix}"
  description = "The default dnss"
  domain      = "%{random_suffix}.hashicorptest.com"
}

output "record_name_to_insert" {
 value = google_certificate_manager_dns_authorization.default.dns_resource_record.0.name
}

output "record_type_to_insert" {
 value = google_certificate_manager_dns_authorization.default.dns_resource_record.0.type
}

output "record_data_to_insert" {
 value = google_certificate_manager_dns_authorization.default.dns_resource_record.0.data
}
`, context)
}

func testAccCheckCertificateManagerDnsAuthorizationDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_certificate_manager_dns_authorization" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{CertificateManagerBasePath}}projects/{{project}}/locations/global/dnsAuthorizations/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("CertificateManagerDnsAuthorization still exists at %s", url)
			}
		}

		return nil
	}
}
