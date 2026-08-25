package collector

import (
	"bytes"
	"encoding/json"
	"errors"
	"regexp"
	"strings"
)

// RedactionPolicy describes the user-approved data mode before API submission.
type RedactionPolicy struct {
	Mode                  string
	RedactHostnames       bool
	RedactUserNames       bool
	RedactProfileSegments bool
	RedactPrivateIPs      bool
	RedactLongArguments   bool
}

var (
	// Matches standard private IPv4 ranges
	privateIPRegex = regexp.MustCompile(`(?:10\.\d{1,3}\.\d{1,3}\.\d{1,3})|(?:172\.(?:1[6-9]|2\d|3[0-1])\.\d{1,3}\.\d{1,3})|(?:192\.168\.\d{1,3}\.\d{1,3})`)
)

// ApplyRedaction returns a new package view suitable for preview/submission.
func ApplyRedaction(pkg EvidencePackage, policy RedactionPolicy) (EvidencePackage, error) {
	if policy.Mode == "full" {
		return pkg, nil
	}
	if policy.Mode != "redacted" {
		return EvidencePackage{}, errors.New("invalid redaction mode")
	}

	// The simplest and most robust way to redact arbitrarily nested data 
	// (including []map[string]any in CategoryResult) is to serialize, string-replace, and deserialize.
	raw, err := json.Marshal(pkg)
	if err != nil {
		return EvidencePackage{}, err
	}

	redactedRaw := raw

	if policy.RedactHostnames && pkg.Host.Hostname != "" {
		redactedRaw = bytes.ReplaceAll(redactedRaw, []byte(pkg.Host.Hostname), []byte("REDACTED-HOST"))
	}

	if policy.RedactUserNames && pkg.Host.InteractiveUser != "" {
		redactedRaw = bytes.ReplaceAll(redactedRaw, []byte(pkg.Host.InteractiveUser), []byte("REDACTED-USER"))
		// Sometimes username is represented with domain e.g. DOMAIN\User
		parts := strings.Split(pkg.Host.InteractiveUser, "\\")
		if len(parts) == 2 {
			redactedRaw = bytes.ReplaceAll(redactedRaw, []byte(parts[1]), []byte("REDACTED-USER"))
		}
	}

	if policy.RedactProfileSegments {
		// Replace C:\Users\Username with C:\Users\REDACTED-USER (case insensitive in paths)
		// We handle the JSON escaped version: C:\\Users\\...
		// Just a simple heuristic for now.
		redactedRaw = regexp.MustCompile(`(?i)C:\\\\Users\\\\[^\\]+`).ReplaceAll(redactedRaw, []byte("C:\\\\Users\\\\REDACTED-USER"))
	}

	if policy.RedactPrivateIPs {
		redactedRaw = privateIPRegex.ReplaceAll(redactedRaw, []byte("[REDACTED-IP]"))
	}

	var redactedPkg EvidencePackage
	if err := json.Unmarshal(redactedRaw, &redactedPkg); err != nil {
		return EvidencePackage{}, err
	}

	return redactedPkg, nil
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
