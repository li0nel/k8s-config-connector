package google

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"

	"google.golang.org/api/dataproc/v1"
)

type DataprocClusterOperationWaiter struct {
	Service *dataproc.Service
	tpgresource.CommonOperationWaiter
}

func (w *DataprocClusterOperationWaiter) QueryOp() (interface{}, error) {
	if w == nil {
		return nil, fmt.Errorf("Cannot query operation, it's unset or nil.")
	}
	return w.Service.Projects.Regions.Operations.Get(w.Op.Name).Do()
}

func dataprocClusterOperationWait(config *transport_tpg.Config, op *dataproc.Operation, activity, userAgent string, timeout time.Duration) error {
	w := &DataprocClusterOperationWaiter{
		Service: config.NewDataprocClient(userAgent),
	}
	if err := w.SetOp(op); err != nil {
		return err
	}
	return tpgresource.OperationWait(w, activity, timeout, config.PollInterval)
}
