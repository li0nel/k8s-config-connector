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

func TestAccComputeRegionTargetHttpProxy_regionTargetHttpProxyBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionTargetHttpProxyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionTargetHttpProxy_regionTargetHttpProxyBasicExample(context),
			},
			{
				ResourceName:            "google_compute_region_target_http_proxy.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"url_map", "region"},
			},
		},
	})
}

func testAccComputeRegionTargetHttpProxy_regionTargetHttpProxyBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_target_http_proxy" "default" {
  region  = "us-central1"
  name    = "tf-test-test-proxy%{random_suffix}"
  url_map = google_compute_region_url_map.default.id
}

resource "google_compute_region_url_map" "default" {
  region          = "us-central1"
  name            = "tf-test-url-map%{random_suffix}"
  default_service = google_compute_region_backend_service.default.id

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name            = "allpaths"
    default_service = google_compute_region_backend_service.default.id

    path_rule {
      paths   = ["/*"]
      service = google_compute_region_backend_service.default.id
    }
  }
}

resource "google_compute_region_backend_service" "default" {
  region                = "us-central1"
  name                  = "tf-test-backend-service%{random_suffix}"
  protocol              = "HTTP"
  timeout_sec           = 10
  load_balancing_scheme = "INTERNAL_MANAGED"

  health_checks = [google_compute_region_health_check.default.id]
}

resource "google_compute_region_health_check" "default" {
  region = "us-central1"
  name   = "tf-test-http-health-check%{random_suffix}"
  http_health_check {
    port = 80
  }
}
`, context)
}

func TestAccComputeRegionTargetHttpProxy_regionTargetHttpProxyHttpsRedirectExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionTargetHttpProxyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionTargetHttpProxy_regionTargetHttpProxyHttpsRedirectExample(context),
			},
			{
				ResourceName:            "google_compute_region_target_http_proxy.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"url_map", "region"},
			},
		},
	})
}

func testAccComputeRegionTargetHttpProxy_regionTargetHttpProxyHttpsRedirectExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_target_http_proxy" "default" {
  region  = "us-central1"
  name    = "tf-test-test-https-redirect-proxy%{random_suffix}"
  url_map = google_compute_region_url_map.default.id
}

resource "google_compute_region_url_map" "default" {
  region          = "us-central1"
  name            = "tf-test-url-map%{random_suffix}"
  default_url_redirect {
    https_redirect = true
    strip_query    = false
  }
}
`, context)
}

func testAccCheckComputeRegionTargetHttpProxyDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_region_target_http_proxy" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/targetHttpProxies/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("ComputeRegionTargetHttpProxy still exists at %s", url)
			}
		}

		return nil
	}
}
