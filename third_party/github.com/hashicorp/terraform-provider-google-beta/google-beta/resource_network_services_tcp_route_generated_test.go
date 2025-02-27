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

func TestAccNetworkServicesTcpRoute_networkServicesTcpRouteBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckNetworkServicesTcpRouteDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesTcpRoute_networkServicesTcpRouteBasicExample(context),
			},
			{
				ResourceName:            "google_network_services_tcp_route.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name"},
			},
		},
	})
}

func testAccNetworkServicesTcpRoute_networkServicesTcpRouteBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_backend_service" "default" {
  provider               = google-beta
  name          = "tf-test-my-backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  provider               = google-beta
  name               = "tf-test-backend-service-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_network_services_tcp_route" "default" {
  provider               = google-beta
  name                   = "tf-test-my-tcp-route%{random_suffix}"
  labels                 = {
    foo = "bar"
  }
  description             = "my description"
  rules                   {
    matches {
      address = "10.0.0.1/32"
      port = "8081"
    }
    action {
      destinations {
        service_name = google_compute_backend_service.default.id
        weight = 1
      }
      original_destination = false
    }
  }
}
`, context)
}

func TestAccNetworkServicesTcpRoute_networkServicesTcpRouteActionsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckNetworkServicesTcpRouteDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesTcpRoute_networkServicesTcpRouteActionsExample(context),
			},
			{
				ResourceName:            "google_network_services_tcp_route.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name"},
			},
		},
	})
}

func testAccNetworkServicesTcpRoute_networkServicesTcpRouteActionsExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_backend_service" "default" {
  provider               = google-beta
  name          = "tf-test-my-backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  provider               = google-beta
  name               = "tf-test-backend-service-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_network_services_tcp_route" "default" {
  provider               = google-beta
  name                   = "tf-test-my-tcp-route%{random_suffix}"
  labels                 = {
    foo = "bar"
  }
  description             = "my description"
  rules                   {
    action {
      destinations {
        service_name = google_compute_backend_service.default.id
        weight = 1
      }
      original_destination = false
    }
  }
}
`, context)
}

func TestAccNetworkServicesTcpRoute_networkServicesTcpRouteMeshBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckNetworkServicesTcpRouteDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesTcpRoute_networkServicesTcpRouteMeshBasicExample(context),
			},
			{
				ResourceName:            "google_network_services_tcp_route.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name"},
			},
		},
	})
}

func testAccNetworkServicesTcpRoute_networkServicesTcpRouteMeshBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_backend_service" "default" {
  provider               = google-beta
  name          = "tf-test-my-backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  provider               = google-beta
  name               = "tf-test-backend-service-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_network_services_mesh" "default" {
  provider    = google-beta
  name        = "tf-test-my-tcp-route%{random_suffix}"
  labels      = {
    foo = "bar"
  }
  description = "my description"
}


resource "google_network_services_tcp_route" "default" {
  provider               = google-beta
  name                   = "tf-test-my-tcp-route%{random_suffix}"
  labels                 = {
    foo = "bar"
  }
  description             = "my description"
  meshes = [
    google_network_services_mesh.default.id
  ]
  rules                   {
    matches {
      address = "10.0.0.1/32"
      port = "8081"
    }
    action {
      destinations {
        service_name = google_compute_backend_service.default.id
        weight = 1
      }
      original_destination = false
    }
  }
}
`, context)
}

func TestAccNetworkServicesTcpRoute_networkServicesTcpRouteGatewayBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckNetworkServicesTcpRouteDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesTcpRoute_networkServicesTcpRouteGatewayBasicExample(context),
			},
			{
				ResourceName:            "google_network_services_tcp_route.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name"},
			},
		},
	})
}

func testAccNetworkServicesTcpRoute_networkServicesTcpRouteGatewayBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_backend_service" "default" {
  provider               = google-beta
  name          = "tf-test-my-backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  provider               = google-beta
  name               = "tf-test-backend-service-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_network_services_gateway" "default" {
  provider    = google-beta
  name        = "tf-test-my-tcp-route%{random_suffix}"
  labels      = {
    foo = "bar"
  }
  description = "my description"
  scope = "my-scope"
  type = "OPEN_MESH"
  ports = [443]
}


resource "google_network_services_tcp_route" "default" {
  provider               = google-beta
  name                   = "tf-test-my-tcp-route%{random_suffix}"
  labels                 = {
    foo = "bar"
  }
  description             = "my description"
  gateways = [
    google_network_services_gateway.default.id
  ]
  rules                   {
    matches {
      address = "10.0.0.1/32"
      port = "8081"
    }
    action {
      destinations {
        service_name = google_compute_backend_service.default.id
        weight = 1
      }
      original_destination = false
    }
  }
}
`, context)
}

func testAccCheckNetworkServicesTcpRouteDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_network_services_tcp_route" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{NetworkServicesBasePath}}projects/{{project}}/locations/global/tcpRoutes/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("NetworkServicesTcpRoute still exists at %s", url)
			}
		}

		return nil
	}
}
