package collector

import (
	"context"
	"errors"
)

const CollectorVersion = "0.1.0-scaffold"

// Config contains explicit, non-privileged collection options.
// TODO(local-agent): add only options that preserve the v1 symptom scope.
// Do not add flags for remediation, execution, evasion or continuous monitoring.
type Config struct {
	IncludeHashes       bool
	IncludeFullPaths    bool
	IncludeCommandLines bool
	DefenderEventLimit  int
}

// Collect is the single entry point used by the local application.
//
// SECURITY CONTRACT:
//   - local collection must be read-only;
//   - no child process may be spawned;
//   - no registry/service/task/WMI/file mutation is allowed;
//   - no network request may be made from this package;
//   - permission failures must become explicit Limitations, not empty "safe" results.
//
// TODO(local-agent): implement orchestration after the synthetic fixtures and
// read-only verification harness are in place. Return a partial package when
// independent categories fail, and fail closed if the read-only assertion
// cannot be established.
func Collect(ctx context.Context, cfg Config) (EvidencePackage, error) {
	_ = ctx
	_ = cfg
	return EvidencePackage{}, errors.New("collector scaffold: local Windows collection is not implemented")
}

// assertReadOnly is intentionally a placeholder for a testable invariant.
// TODO(local-agent): make the production code impossible to call with a
// write-capable handle or mutating API. Prefer narrow wrappers around Windows
// read APIs and reject any unsupported operation at compile/review time.
func assertReadOnly() error {
	return nil
}
