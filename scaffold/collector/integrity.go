package collector

import (
	"errors"
)

// CanonicalSHA256 computes the digest that binds a package to the report and
// consent record.
// TODO(local-agent): use deterministic JSON canonicalization compatible with
// the documented profile, exclude the integrity.sha256 field while hashing,
// and test that map ordering does not change the digest.
func CanonicalSHA256(pkg EvidencePackage) (string, error) {
	_ = pkg
	return "", errors.New("integrity scaffold: not implemented")
}

// ValidatePackage must enforce schema_version, required categories, the
// read_only_assertion and hash format before any API client can submit data.
// TODO(local-agent): call a JSON Schema validator and return actionable errors.
func ValidatePackage(pkg EvidencePackage) error {
	_ = pkg
	return errors.New("package validation scaffold: not implemented")
}
