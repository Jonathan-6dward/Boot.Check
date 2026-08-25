package collector

import (
	"context"
	"errors"
)

// collectPersistence reads only the documented Run and RunOnce locations for
// the current v1 scope. It must preserve hive, key path, value name and the
// original/redacted value for correlation with cmd.exe or PowerShell.
// TODO(local-agent): open registry keys with query/read-only access only.
// TODO(local-agent): cover HKCU/HKLM and the supported 32/64-bit views without
// changing either view. Record which views were inaccessible.
// TODO(local-agent): never expand environment variables by executing them;
// expansion may be represented as raw text plus a deterministic note.
func collectPersistence(ctx context.Context, cfg Config) (PersistenceEvidence, []Limitation, error) {
	_ = ctx
	_ = cfg
	return PersistenceEvidence{}, nil, errors.New("registry Run/RunOnce collection TODO")
}
