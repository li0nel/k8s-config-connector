package google

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func DataSourceDataprocMetastoreService() *schema.Resource {

	dsSchema := tpgresource.DatasourceSchemaFromResourceSchema(ResourceDataprocMetastoreService().Schema)
	tpgresource.AddRequiredFieldsToSchema(dsSchema, "service_id")
	tpgresource.AddRequiredFieldsToSchema(dsSchema, "location")
	tpgresource.AddOptionalFieldsToSchema(dsSchema, "project")

	return &schema.Resource{
		Read:   dataSourceDataprocMetastoreServiceRead,
		Schema: dsSchema,
	}
}

func dataSourceDataprocMetastoreServiceRead(d *schema.ResourceData, meta interface{}) error {
	id, err := tpgresource.ReplaceVars(d, meta.(*transport_tpg.Config), "projects/{{project}}/locations/{{location}}/services/{{service_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)
	return resourceDataprocMetastoreServiceRead(d, meta)
}
