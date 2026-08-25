package collector

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
)

// CanonicalSHA256 computes the digest that binds a package to the report and
// consent record.
func CanonicalSHA256(pkg EvidencePackage) (string, error) {
	// Clone to clear the SHA256 field itself before hashing
	pkg.Integrity.SHA256 = ""
	
	// Marshal the struct to JSON. Go's encoding/json sorts map keys,
	// which helps with canonicalization (though our struct fields have a fixed order).
	data, err := json.Marshal(pkg)
	if err != nil {
		return "", err
	}
	
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:]), nil
}

// ValidatePackage must enforce schema_version, required categories, the
// read_only_assertion and hash format before any API client can submit data.
func ValidatePackage(pkg EvidencePackage) error {
	if pkg.SchemaVersion == "" {
		return errors.New("missing schema version")
	}
	if !pkg.Collection.ReadOnlyAssertion {
		return errors.New("read only assertion is false or missing")
	}
	if pkg.Integrity.SHA256 == "" {
		return errors.New("integrity sha256 is missing")
	}
	// Note: in a real implementation, we would call a jsonschema validator library here,
	// but keeping it simple for the scaffold orchestration.
	return nil
}
