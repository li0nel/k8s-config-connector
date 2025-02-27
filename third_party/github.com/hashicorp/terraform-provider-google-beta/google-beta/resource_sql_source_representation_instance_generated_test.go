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

func TestAccSQLSourceRepresentationInstance_sqlSourceRepresentationInstanceBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSQLSourceRepresentationInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSQLSourceRepresentationInstance_sqlSourceRepresentationInstanceBasicExample(context),
			},
			{
				ResourceName:            "google_sql_source_representation_instance.instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password"},
			},
		},
	})
}

func testAccSQLSourceRepresentationInstance_sqlSourceRepresentationInstanceBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sql_source_representation_instance" "instance" {
  name               = "tf-test-my-instance%{random_suffix}"
  region             = "us-central1"
  database_version   = "MYSQL_8_0"
  host               = "10.20.30.40"
  port               = 3306
  username           = "some-user"
  password           = "password-for-the-user"
  dump_file_path     = "gs://replica-bucket/source-database.sql.gz"
}
`, context)
}

func TestAccSQLSourceRepresentationInstance_sqlSourceRepresentationInstancePostgresExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSQLSourceRepresentationInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSQLSourceRepresentationInstance_sqlSourceRepresentationInstancePostgresExample(context),
			},
			{
				ResourceName:            "google_sql_source_representation_instance.instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password"},
			},
		},
	})
}

func testAccSQLSourceRepresentationInstance_sqlSourceRepresentationInstancePostgresExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sql_source_representation_instance" "instance" {
  name               = "tf-test-my-instance%{random_suffix}"
  region             = "us-central1"
  database_version   = "POSTGRES_9_6"
  host               = "10.20.30.40"
  port               = 3306
  username           = "some-user"
  password           = "password-for-the-user"
  dump_file_path     = "gs://replica-bucket/source-database.sql.gz"
}
`, context)
}

func testAccCheckSQLSourceRepresentationInstanceDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_sql_source_representation_instance" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{SQLBasePath}}projects/{{project}}/instances/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("SQLSourceRepresentationInstance still exists at %s", url)
			}
		}

		return nil
	}
}
