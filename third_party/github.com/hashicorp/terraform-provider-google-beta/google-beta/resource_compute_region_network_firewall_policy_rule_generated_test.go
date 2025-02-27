// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: DCL     ***
//
// ----------------------------------------------------------------------------
//
//     This file is managed by Magic Modules (https://github.com/GoogleCloudPlatform/magic-modules)
//     and is based on the DCL (https://github.com/GoogleCloudPlatform/declarative-resource-client-library).
//     Changes will need to be made to the DCL or Magic Modules instead of here.
//
//     We are not currently able to accept contributions to this file. If changes
//     are required, please file an issue at https://github.com/hashicorp/terraform-provider-google/issues/new/choose
//
// ----------------------------------------------------------------------------

package google

import (
	"context"
	"fmt"
	dcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	compute "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccComputeRegionNetworkFirewallPolicyRule_RegionalHandWritten(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        acctest.GetTestOrgFromEnv(t),
		"project_name":  acctest.GetTestProjectFromEnv(),
		"region":        acctest.GetTestRegionFromEnv(),
		"service_acct":  acctest.GetTestServiceAccountFromEnv(t),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionNetworkFirewallPolicyRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionNetworkFirewallPolicyRule_RegionalHandWritten(context),
			},
			{
				ResourceName:      "google_compute_region_network_firewall_policy_rule.primary",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionNetworkFirewallPolicyRule_RegionalHandWrittenUpdate0(context),
			},
			{
				ResourceName:      "google_compute_region_network_firewall_policy_rule.primary",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeRegionNetworkFirewallPolicyRule_RegionalHandWritten(context map[string]interface{}) string {
	return Nprintf(`
resource "google_network_security_address_group" "basic_regional_networksecurity_address_group" {
  name        = "tf-test-policy%{random_suffix}"
  parent      = "projects/%{project_name}"
  description = "Sample regional networksecurity_address_group"
  location    = "%{region}"
  items       = ["208.80.154.224/32"]
  type        = "IPV4"
  capacity    = 100
}

resource "google_compute_region_network_firewall_policy" "basic_regional_network_firewall_policy" {
  name        = "tf-test-policy%{random_suffix}"
  description = "Sample regional network firewall policy"
  project     = "%{project_name}"
  region      = "%{region}"
}

resource "google_compute_region_network_firewall_policy_rule" "primary" {
  action                  = "allow"
  description             = "This is a simple rule description"
  direction               = "INGRESS"
  disabled                = false
  enable_logging          = true
  firewall_policy         = google_compute_region_network_firewall_policy.basic_regional_network_firewall_policy.name
  priority                = 1000
  region                  = "%{region}"
  rule_name               = "test-rule"
  target_service_accounts = ["%{service_acct}"]

  match {
    src_ip_ranges = ["10.100.0.1/32"]
    src_fqdns = ["example.com"]
    src_region_codes = ["US"]
    src_threat_intelligences = ["iplist-known-malicious-ips"]

    layer4_configs {
      ip_protocol = "all"
    }

    src_secure_tags {
      name = "tagValues/${google_tags_tag_value.basic_value.name}"
    }
    
    src_address_groups = [google_network_security_address_group.basic_regional_networksecurity_address_group.id]
  }
}

resource "google_compute_network" "basic_network" {
  name = "tf-test-network%{random_suffix}"
}

resource "google_tags_tag_key" "basic_key" {
  description = "For keyname resources."
  parent      = "organizations/%{org_id}"
  purpose     = "GCE_FIREWALL"
  short_name  = "tf-test-tagkey%{random_suffix}"

  purpose_data = {
    network = "%{project_name}/${google_compute_network.basic_network.name}"
  }
}

resource "google_tags_tag_value" "basic_value" {
  description = "For valuename resources."
  parent      = "tagKeys/${google_tags_tag_key.basic_key.name}"
  short_name  = "tf-test-tagvalue%{random_suffix}"
}

`, context)
}

func testAccComputeRegionNetworkFirewallPolicyRule_RegionalHandWrittenUpdate0(context map[string]interface{}) string {
	return Nprintf(`
resource "google_network_security_address_group" "basic_regional_networksecurity_address_group" {
  name        = "tf-test-policy%{random_suffix}"
  parent      = "projects/%{project_name}"
  description = "Sample regional networksecurity_address_group. Update"
  location    = "%{region}"
  items       = ["208.80.154.224/32"]
  type        = "IPV4"
  capacity    = 100
}

resource "google_compute_region_network_firewall_policy" "basic_regional_network_firewall_policy" {
  name        = "tf-test-policy%{random_suffix}"
  description = "Sample regional network firewall policy"
  project     = "%{project_name}"
  region      = "%{region}"
}

resource "google_compute_region_network_firewall_policy_rule" "primary" {
  action          = "deny"
  description     = "This is an updated rule description"
  direction       = "EGRESS"
  disabled        = true
  enable_logging  = false
  firewall_policy = google_compute_region_network_firewall_policy.basic_regional_network_firewall_policy.name
  priority        = 1000
  region          = "%{region}"
  rule_name       = "updated-test-rule"

  match {
    dest_ip_ranges = ["0.0.0.0/0"]
    dest_fqdns = ["example.com"]
    dest_region_codes = ["US"]
    dest_threat_intelligences = ["iplist-known-malicious-ips"]

    layer4_configs {
      ip_protocol = "tcp"
      ports       = ["123"]
    }
    
    dest_address_groups = [google_network_security_address_group.basic_regional_networksecurity_address_group.id]
  }

  target_secure_tags {
    name = "tagValues/${google_tags_tag_value.basic_value.name}"
  }
}

resource "google_compute_network" "basic_network" {
  name = "tf-test-network%{random_suffix}"
}

resource "google_tags_tag_key" "basic_key" {
  description = "For keyname resources."
  parent      = "organizations/%{org_id}"
  purpose     = "GCE_FIREWALL"
  short_name  = "tf-test-tagkey%{random_suffix}"

  purpose_data = {
    network = "%{project_name}/${google_compute_network.basic_network.name}"
  }
}

resource "google_tags_tag_value" "basic_value" {
  description = "For valuename resources."
  parent      = "tagKeys/${google_tags_tag_key.basic_key.name}"
  short_name  = "tf-test-tagvalue%{random_suffix}"
}

`, context)
}

func testAccCheckComputeRegionNetworkFirewallPolicyRuleDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "rs.google_compute_region_network_firewall_policy_rule" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			billingProject := ""
			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			obj := &compute.NetworkFirewallPolicyRule{
				Action:         dcl.String(rs.Primary.Attributes["action"]),
				Direction:      compute.NetworkFirewallPolicyRuleDirectionEnumRef(rs.Primary.Attributes["direction"]),
				FirewallPolicy: dcl.String(rs.Primary.Attributes["firewall_policy"]),
				Description:    dcl.String(rs.Primary.Attributes["description"]),
				Disabled:       dcl.Bool(rs.Primary.Attributes["disabled"] == "true"),
				EnableLogging:  dcl.Bool(rs.Primary.Attributes["enable_logging"] == "true"),
				Project:        dcl.StringOrNil(rs.Primary.Attributes["project"]),
				Location:       dcl.StringOrNil(rs.Primary.Attributes["region"]),
				RuleName:       dcl.String(rs.Primary.Attributes["rule_name"]),
				Kind:           dcl.StringOrNil(rs.Primary.Attributes["kind"]),
			}

			client := transport_tpg.NewDCLComputeClient(config, config.UserAgent, billingProject, 0)
			_, err := client.GetNetworkFirewallPolicyRule(context.Background(), obj)
			if err == nil {
				return fmt.Errorf("google_compute_region_network_firewall_policy_rule still exists %v", obj)
			}
		}
		return nil
	}
}
