package collector

import (
	"context"
	"errors"
)

// collectServices reads service metadata relevant to boot execution: name,
// state, start type, account, image path, description and signature metadata.
// TODO(local-agent): use Service Control Manager enumeration with query-only
// access. Never call StartService, ControlService, ChangeServiceConfig or
// delete APIs. Record service access failures as limitations.
// TODO(local-agent): do not load a service image; signature/hash inspection
// must read bytes and remain outside any execution path.
func collectServices(ctx context.Context, cfg Config) (CategoryResult, []Limitation, error) {
	_ = ctx
	_ = cfg
	return CategoryResult{Status: "not_supported", Items: []map[string]any{}}, nil, errors.New("service collection TODO")
}
