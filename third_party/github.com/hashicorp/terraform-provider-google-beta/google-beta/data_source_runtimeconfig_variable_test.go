package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccRuntimeconfigVariableDatasource_basic(t *testing.T) {
	t.Parallel()

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccRuntimeconfigDatasourceVariable(RandString(t, 10), RandString(t, 10), RandString(t, 10)),
				Check: resource.ComposeTestCheckFunc(
					acctest.CheckDataSourceStateMatchesResourceState("data.google_runtimeconfig_variable.default", "google_runtimeconfig_variable.default"),
				),
			},
		},
	})
}

func testAccRuntimeconfigDatasourceVariable(suffix string, name string, text string) string {
	return fmt.Sprintf(`
	resource "google_runtimeconfig_config" "default" {
		name        = "runtime-%s"
		description = "runtime-%s"
	}

	resource "google_runtimeconfig_variable" "default" {
		parent  = google_runtimeconfig_config.default.name
		name    = "%s"
		text    = "%s"
	}

	data "google_runtimeconfig_variable" "default" {
		name    = google_runtimeconfig_variable.default.name
		parent  = google_runtimeconfig_config.default.name
	}
`, suffix, suffix, name, text)
}
