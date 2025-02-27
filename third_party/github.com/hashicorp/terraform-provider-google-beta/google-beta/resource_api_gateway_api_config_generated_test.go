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

func TestAccApiGatewayApiConfig_apigatewayApiConfigBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckApiGatewayApiConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApiGatewayApiConfig_apigatewayApiConfigBasicExample(context),
			},
			{
				ResourceName:            "google_api_gateway_api_config.api_cfg",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"gateway_config", "grpc_services", "api", "api_config_id"},
			},
		},
	})
}

func testAccApiGatewayApiConfig_apigatewayApiConfigBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_api_gateway_api" "api_cfg" {
  provider = google-beta
  api_id = "tf-test-my-api%{random_suffix}"
}

resource "google_api_gateway_api_config" "api_cfg" {
  provider = google-beta
  api = google_api_gateway_api.api_cfg.api_id
  api_config_id = "tf-test-my-config%{random_suffix}"

  openapi_documents {
    document {
      path = "spec.yaml"
      contents = filebase64("test-fixtures/apigateway/openapi.yaml")
    }
  }
  lifecycle {
    create_before_destroy = true
  }
}
`, context)
}

func TestAccApiGatewayApiConfig_apigatewayApiConfigFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckApiGatewayApiConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApiGatewayApiConfig_apigatewayApiConfigFullExample(context),
			},
			{
				ResourceName:            "google_api_gateway_api_config.api_cfg",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"gateway_config", "grpc_services", "api", "api_config_id"},
			},
		},
	})
}

func testAccApiGatewayApiConfig_apigatewayApiConfigFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_api_gateway_api" "api_cfg" {
  provider = google-beta
  api_id = "tf-test-my-api%{random_suffix}"
}

resource "google_api_gateway_api_config" "api_cfg" {
  provider = google-beta
  api = google_api_gateway_api.api_cfg.api_id
  api_config_id = "tf-test-my-config%{random_suffix}"
  display_name = "MM Dev API Config"
  labels = {
    environment = "dev"
  }

  openapi_documents {
    document {
      path = "spec.yaml"
      contents = filebase64("test-fixtures/apigateway/openapi.yaml")
    }
  }
}
`, context)
}

func TestAccApiGatewayApiConfig_apigatewayApiConfigGrpcExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckApiGatewayApiConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApiGatewayApiConfig_apigatewayApiConfigGrpcExample(context),
			},
			{
				ResourceName:            "google_api_gateway_api_config.api_cfg",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"gateway_config", "grpc_services", "api", "api_config_id", "grpc_services.0.file_descriptor_set"},
			},
		},
	})
}

func testAccApiGatewayApiConfig_apigatewayApiConfigGrpcExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_api_gateway_api" "api_cfg" {
  provider = google-beta
  api_id = "tf-test-my-api%{random_suffix}"
}

resource "google_api_gateway_api_config" "api_cfg" {
  provider = google-beta
  api = google_api_gateway_api.api_cfg.api_id
  api_config_id = "tf-test-my-config%{random_suffix}"

  grpc_services {
    file_descriptor_set {
      path = "api_descriptor.pb"
      contents = filebase64("test-fixtures/apigateway/api_descriptor.pb")
    }
  }
  managed_service_configs {
    path = "api_config.yaml"
    contents = base64encode(<<-EOF
      type: google.api.Service
      config_version: 3
      name: ${google_api_gateway_api.api_cfg.managed_service}
      title: gRPC API example
      apis:
        - name: endpoints.examples.bookstore.Bookstore
      usage:
        rules:
        - selector: endpoints.examples.bookstore.Bookstore.ListShelves
          allow_unregistered_calls: true
      backend:
        rules:
          - selector: "*"
            address: grpcs://example.com
            disable_auth: true

    EOF
    )
  }
  lifecycle {
    create_before_destroy = true
  }
}
`, context)
}

func TestAccApiGatewayApiConfig_apigatewayApiConfigGrpcFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckApiGatewayApiConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApiGatewayApiConfig_apigatewayApiConfigGrpcFullExample(context),
			},
			{
				ResourceName:            "google_api_gateway_api_config.api_cfg",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"gateway_config", "grpc_services", "api", "api_config_id"},
			},
		},
	})
}

func testAccApiGatewayApiConfig_apigatewayApiConfigGrpcFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_api_gateway_api" "api_cfg" {
  provider = google-beta
  api_id = "tf-test-my-api%{random_suffix}"
}

resource "google_api_gateway_api_config" "api_cfg" {
  provider = google-beta
  api = google_api_gateway_api.api_cfg.api_id
  api_config_id = "tf-test-my-config%{random_suffix}"

  grpc_services {
    file_descriptor_set {
      path = "api_descriptor.pb"
      contents = filebase64("test-fixtures/apigateway/api_descriptor.pb")
    }
    source {
      path = "bookstore.proto"
      contents = filebase64("test-fixtures/apigateway/bookstore.proto")
    }
  }
  managed_service_configs {
    path = "api_config.yaml"
    contents = base64encode(<<-EOF
      type: google.api.Service
      config_version: 3
      name: ${google_api_gateway_api.api_cfg.managed_service}
      title: gRPC API example
      apis:
        - name: endpoints.examples.bookstore.Bookstore
      usage:
        rules:
        - selector: endpoints.examples.bookstore.Bookstore.ListShelves
          allow_unregistered_calls: true
      backend:
        rules:
          - selector: "*"
            address: grpcs://example.com
            disable_auth: true

    EOF
    )
  }
  lifecycle {
    create_before_destroy = true
  }
}
`, context)
}

func testAccCheckApiGatewayApiConfigDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_api_gateway_api_config" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ApiGatewayBasePath}}projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("ApiGatewayApiConfig still exists at %s", url)
			}
		}

		return nil
	}
}
