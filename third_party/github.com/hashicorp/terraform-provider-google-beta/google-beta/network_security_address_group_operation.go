package google

import (
	"time"

	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// NetworkSecurityAddressGroupOperationWaitTime is specific for address group resource because the only difference is that it does not need project param.
func NetworkSecurityAddressGroupOperationWaitTime(config *transport_tpg.Config, op map[string]interface{}, activity, userAgent string, timeout time.Duration) error {
	// project is not necessary for this operation.
	return NetworkSecurityOperationWaitTime(config, op, "", activity, userAgent, timeout)
}
