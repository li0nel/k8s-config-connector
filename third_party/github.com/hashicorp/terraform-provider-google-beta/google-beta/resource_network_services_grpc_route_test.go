package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetworkServicesGrpcRoute_update(t *testing.T) {
	t.Parallel()

	grpcRouteName := fmt.Sprintf("tf-test-grpc-route-%s", RandString(t, 10))

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkServicesGrpcRouteDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesGrpcRoute_basic(grpcRouteName),
			},
			{
				ResourceName:      "google_network_services_grpc_route.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccNetworkServicesGrpcRoute_update(grpcRouteName),
			},
			{
				ResourceName:      "google_network_services_grpc_route.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccNetworkServicesGrpcRoute_basic(grpcRouteName string) string {
	return fmt.Sprintf(`
  resource "google_network_services_grpc_route" "foobar" {
    name                   = "%s"
    labels                 = {
      foo = "bar"
    }
    description             = "my description"
    hostnames               = ["example"]
    rules                   {
      matches {
        headers {
          key = "key"
          value = "value"
        }
      }
      action {
        retry_policy {
            retry_conditions = ["cancelled"]
            num_retries = 1
        }
      }
    }
    rules                   {
      matches {
        headers {
          key = "key"
          value = "value"
        }
      }
      action {
        fault_injection_policy {
          delay {
            fixed_delay = "1s"
            percentage = 1
          }
          abort {
            http_status = 500
            percentage = 1
          }
        }
      }
    }
  }
`, grpcRouteName)
}

func testAccNetworkServicesGrpcRoute_update(grpcRouteName string) string {
	return fmt.Sprintf(`
  resource "google_network_services_grpc_route" "foobar" {
    name                   = "%s"
    labels                 = {
      foo = "bar"
    }
    description             = "updated description"
    hostnames               = ["example"]
    rules                   {
      matches {
        headers {
          key = "key"
          value = "value"
        }
      }
      action {
        retry_policy {
            retry_conditions = ["cancelled"]
            num_retries = 2
        }
      }
    }
    rules                   {
      matches {
        headers {
          key = "key1"
          value = "value1"
        }
      }
      action {
        retry_policy {
            retry_conditions = ["connect-failure"]
            num_retries = 1
        }
      }
    }
  }
`, grpcRouteName)
}
