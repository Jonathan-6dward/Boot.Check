package api

// VerdictJSONSchema is supplied to providers that support strict structured
// output. Keep it synchronized with verdict.schema.json.
// TODO(local-agent): add a startup test that compares this representation with
// the checked-in JSON file and with the report renderer's required fields.
func VerdictJSONSchema() map[string]any {
	stringArray := map[string]any{"type": "array", "items": map[string]any{"type": "string"}}
	return map[string]any{
		"type":                 "object",
		"additionalProperties": false,
		"required":             []string{"verdict", "confidence", "headline_plain", "plain_language_summary", "five_w_two_h", "mitre", "supporting_evidence", "limitations", "recommended_next_steps", "technical_appendix", "safety_notice"},
		"properties": map[string]any{
			"verdict":                map[string]any{"type": "string", "enum": []string{"likely_safe", "suspicious", "inconclusive"}},
			"confidence":             map[string]any{"type": "number", "minimum": 0, "maximum": 1},
			"headline_plain":         map[string]any{"type": "string", "minLength": 1, "maxLength": 500},
			"plain_language_summary": map[string]any{"type": "string", "minLength": 1, "maxLength": 5000},
			"five_w_two_h": map[string]any{
				"type": "object", "additionalProperties": false,
				"required": []string{"what", "who", "when", "where", "why", "how", "impact", "evidence_ids"},
				"properties": map[string]any{
					"what": map[string]any{"type": "string"}, "who": map[string]any{"type": "string"}, "when": map[string]any{"type": "string"}, "where": map[string]any{"type": "string"}, "why": map[string]any{"type": "string"}, "how": map[string]any{"type": "string"}, "impact": map[string]any{"type": "string"}, "evidence_ids": stringArray,
				},
			},
			"mitre": map[string]any{"type": "array", "items": map[string]any{
				"type": "object", "additionalProperties": false,
				"required":   []string{"technique_id", "name", "rationale", "evidence_ids"},
				"properties": map[string]any{"technique_id": map[string]any{"type": "string"}, "name": map[string]any{"type": "string"}, "rationale": map[string]any{"type": "string"}, "evidence_ids": stringArray},
			}},
			"supporting_evidence": map[string]any{"type": "array", "items": map[string]any{
				"type": "object", "additionalProperties": false,
				"required":   []string{"evidence_id", "role", "explanation"},
				"properties": map[string]any{"evidence_id": map[string]any{"type": "string"}, "role": map[string]any{"type": "string", "enum": []string{"supports", "contradicts", "context"}}, "explanation": map[string]any{"type": "string"}},
			}},
			"limitations":            stringArray,
			"recommended_next_steps": stringArray,
			"technical_appendix": map[string]any{
				"type": "object", "additionalProperties": false,
				"required":   []string{"processes", "persistence", "network", "notes"},
				"properties": map[string]any{"processes": map[string]any{"type": "array", "items": map[string]any{}}, "persistence": map[string]any{"type": "array", "items": map[string]any{}}, "network": map[string]any{"type": "array", "items": map[string]any{}}, "notes": stringArray},
			},
			"safety_notice": map[string]any{"type": "string", "minLength": 1, "maxLength": 2000},
		},
	}
}
