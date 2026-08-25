package collector

import "errors"

// RedactionPolicy describes the user-approved data mode before API submission.
type RedactionPolicy struct {
	Mode                  string
	RedactHostnames       bool
	RedactUserNames       bool
	RedactProfileSegments bool
	RedactPrivateIPs      bool
	RedactLongArguments   bool
}

// ApplyRedaction returns a new package view suitable for preview/submission.
// TODO(local-agent): implement deterministic, reversible-in-memory-only
// redaction. Never overwrite the original local evidence artifact and never
// log unredacted values. The default must be "redacted".
func ApplyRedaction(pkg EvidencePackage, policy RedactionPolicy) (EvidencePackage, error) {
	_ = pkg
	_ = policy
	return EvidencePackage{}, errors.New("redaction scaffold: not implemented")
}

// ValidateSubmissionPolicy must reject a full-data submission unless the UI
// has a fresh, explicit, versioned consent record for the exact provider and
// purpose displayed to the user.
// TODO(local-agent): connect this to the consent screen and audit record.
func ValidateSubmissionPolicy(pkg EvidencePackage, policy RedactionPolicy) error {
	if policy.Mode != "redacted" && policy.Mode != "full" {
		return errors.New("invalid data mode")
	}
	if !pkg.Consent.LocalCollectionConfirmed {
		return errors.New("local collection consent is missing")
	}
	if policy.Mode == "full" && !pkg.Consent.LLMSubmissionConfirmed {
		return errors.New("full submission requires explicit LLM consent")
	}
	return nil
}
