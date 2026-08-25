package collector

import (
	"context"
	"errors"
)

// collectProcesses must enumerate a point-in-time process snapshot and derive
// parent/child links from observed PID/PPID values.
// TODO(local-agent): use read-only Windows process/query APIs. Do not request
// PROCESS_VM_WRITE, PROCESS_TERMINATE, PROCESS_CREATE_THREAD or equivalent
// write-capable rights. Treat access-denied fields as nil + Limitation.
// TODO(local-agent): signature/hash inspection must read the image only and
// must never execute it or load it into the target process.
func collectProcesses(ctx context.Context, cfg Config) ([]ProcessEvidence, []Limitation, error) {
	_ = ctx
	_ = cfg
	return nil, nil, errors.New("process collection TODO")
}
