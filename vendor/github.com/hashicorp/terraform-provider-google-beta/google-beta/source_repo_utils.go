package google

import (
	"regexp"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func expandSourceRepoRepositoryPubsubConfigsTopic(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (string, error) {
	// short-circuit if the topic is a full uri so we don't need to GetProject
	ok, err := regexp.MatchString(PubsubTopicRegex, v.(string))
	if err != nil {
		return "", err
	}

	if ok {
		return v.(string), nil
	}

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return "", err
	}

	return getComputedTopicName(project, v.(string)), err
}
