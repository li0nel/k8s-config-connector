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

func TestAccLoggingLinkedDataset_loggingLinkedDatasetBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       acctest.GetTestProjectFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckLoggingLinkedDatasetDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccLoggingLinkedDataset_loggingLinkedDatasetBasicExample(context),
			},
			{
				ResourceName:            "google_logging_linked_dataset.logging_linked_dataset",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"link_id", "parent", "location", "bucket"},
			},
		},
	})
}

func testAccLoggingLinkedDataset_loggingLinkedDatasetBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_logging_project_bucket_config" "logging_linked_dataset" {
  location         = "global"
  project          = "%{project}"
  enable_analytics = true
  bucket_id        = "tf-test-my-bucket%{random_suffix}"
}

resource "google_logging_linked_dataset" "logging_linked_dataset" {
  link_id     = "mylink%{random_suffix}"
  bucket      = google_logging_project_bucket_config.logging_linked_dataset.id
  description = "Linked dataset test"
}
`, context)
}

func TestAccLoggingLinkedDataset_loggingLinkedDatasetAllParamsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       acctest.GetTestProjectFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckLoggingLinkedDatasetDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccLoggingLinkedDataset_loggingLinkedDatasetAllParamsExample(context),
			},
			{
				ResourceName:            "google_logging_linked_dataset.logging_linked_dataset",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"link_id", "parent", "location", "bucket"},
			},
		},
	})
}

func testAccLoggingLinkedDataset_loggingLinkedDatasetAllParamsExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_logging_project_bucket_config" "logging_linked_dataset" {
  location         = "global"
  project          = "%{project}"
  enable_analytics = true
  bucket_id        = "tf-test-my-bucket%{random_suffix}"
}

resource "google_logging_linked_dataset" "logging_linked_dataset" {
  link_id     = "mylink%{random_suffix}"
  bucket      = "tf-test-my-bucket%{random_suffix}"
  parent      = "projects/%{project}"
  location    = "global"
  description = "Linked dataset test"

  depends_on = ["google_logging_project_bucket_config.logging_linked_dataset"]
}
`, context)
}

func testAccCheckLoggingLinkedDatasetDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_logging_linked_dataset" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{LoggingBasePath}}{{parent}}/locations/{{location}}/buckets/{{bucket}}/links/{{link_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("LoggingLinkedDataset still exists at %s", url)
			}
		}

		return nil
	}
}
