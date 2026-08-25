package collector

import (
	"context"
	"errors"
)

// collectWinlogon reads the explicitly allow-listed Winlogon values that can
// explain startup behavior, such as shell/userinit-related metadata.
// TODO(local-agent): define the exact allow-list in tests before implementation.
// TODO(local-agent): open the registry with query/read-only access and never
// write, import, delete or normalize values in place.
// TODO(local-agent): avoid treating a non-default value as malicious without
// preserving the observed value, signature context and evidence_id.
func collectWinlogon(ctx context.Context, cfg Config) (CategoryResult, []Limitation, error) {
	_ = ctx
	_ = cfg
	return CategoryResult{Status: "not_supported", Items: []map[string]any{}}, nil, errors.New("Winlogon collection TODO")
}
