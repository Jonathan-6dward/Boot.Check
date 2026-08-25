package collector

import (
	"context"
	"errors"
)

// collectNetwork enumerates active TCP/UDP endpoints from local Windows APIs.
// TODO(local-agent): use GetExtendedTcpTable/GetExtendedUdpTable or equivalent
// read APIs, covering IPv4/IPv6 and owner PID where the OS exposes it.
// TODO(local-agent): do not make DNS queries, port scans, sockets, firewall
// changes or connection attempts. A missing owner PID is a limitation, not a
// reason to omit the endpoint silently.
func collectNetwork(ctx context.Context, cfg Config) (NetworkEvidence, []Limitation, error) {
	_ = ctx
	_ = cfg
	return NetworkEvidence{Status: "failed", TCPConnections: []ConnectionEvidence{}, UDPEndpoints: []ConnectionEvidence{}}, nil, errors.New("network collection TODO")
}
