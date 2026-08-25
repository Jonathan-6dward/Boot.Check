package collector

import (
	"context"
	"errors"
)

// collectDefenderEvents reads a bounded, recent slice of Defender events that
// is available to the current user and relevant to startup/process activity.
// TODO(local-agent): query Event Log in read-only mode with a strict time and
// count bound. Do not clear logs, change policy, trigger scans or disable
// protection. Record event-channel access errors as limitations.
// TODO(local-agent): minimize message content and redact usernames/paths before
// it reaches the API preview.
func collectDefenderEvents(ctx context.Context, cfg Config) (CategoryResult, []Limitation, error) {
	_ = ctx
	_ = cfg
	return CategoryResult{Status: "not_supported", Items: []map[string]any{}}, nil, errors.New("Defender event collection TODO")
}
