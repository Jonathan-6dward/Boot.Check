package collector

import (
	"context"
	"errors"
)

// collectScheduledTasks enumerates Task Scheduler 2.0 definitions and emits
// the actions/triggers as data; it never starts, stops, registers, updates or
// deletes a task.
// TODO(local-agent): use the Task Scheduler 2.0 API/COM in read-only mode.
// TODO(local-agent): ensure COM initialization/cleanup is bounded to this
// goroutine and that no action object is invoked. Capture inaccessible paths.
// TODO(local-agent): redact user profile segments and secret-like arguments
// before the preview; preserve a local correlation identifier.
func collectScheduledTasks(ctx context.Context, cfg Config) (CategoryResult, []Limitation, error) {
	_ = ctx
	_ = cfg
	return CategoryResult{Status: "not_supported", Items: []map[string]any{}}, nil, errors.New("Task Scheduler collection TODO")
}
