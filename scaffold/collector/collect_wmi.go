package collector

import (
	"context"
	"errors"
)

// collectWMISubscriptions inspects permanent event filters, consumers and
// bindings as metadata only.
// TODO(local-agent): query the relevant WMI namespaces with read-only access;
// do not create, modify, bind, unbind or delete filters/consumers.
// TODO(local-agent): capture creator/namespace/timestamps only when available
// and represent provider/permission gaps explicitly.
func collectWMISubscriptions(ctx context.Context, cfg Config) (CategoryResult, []Limitation, error) {
	_ = ctx
	_ = cfg
	return CategoryResult{Status: "not_supported", Items: []map[string]any{}}, nil, errors.New("WMI subscription collection TODO")
}
