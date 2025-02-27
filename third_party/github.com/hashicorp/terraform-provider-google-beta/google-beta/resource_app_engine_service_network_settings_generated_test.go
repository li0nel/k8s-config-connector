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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccAppEngineServiceNetworkSettings_appEngineServiceNetworkSettingsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccAppEngineServiceNetworkSettings_appEngineServiceNetworkSettingsExample(context),
			},
			{
				ResourceName:      "google_app_engine_service_network_settings.internalapp",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAppEngineServiceNetworkSettings_appEngineServiceNetworkSettingsExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "bucket" {
	name     = "tf-test-appengine-static-content%{random_suffix}"
  location = "US"
}

resource "google_storage_bucket_object" "object" {
	name   = "hello-world.zip"
	bucket = google_storage_bucket.bucket.name
	source = "./test-fixtures/appengine/hello-world.zip"
}

resource "google_app_engine_standard_app_version" "internalapp" {
  version_id = "v1"
  service = "internalapp"
  delete_service_on_destroy = true

  runtime = "nodejs10"
  entrypoint {
    shell = "node ./app.js"
  }
  deployment {
    zip {
      source_url = "https://storage.googleapis.com/${google_storage_bucket.bucket.name}/${google_storage_bucket_object.object.name}"
    }  
  }
  env_variables = {
    port = "8080"
  }
}

resource "google_app_engine_service_network_settings" "internalapp" {
  service = google_app_engine_standard_app_version.internalapp.service
  network_settings {
    ingress_traffic_allowed = "INGRESS_TRAFFIC_ALLOWED_INTERNAL_ONLY"
  }
}
`, context)
}
