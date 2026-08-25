package tests

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestSyntheticFixturesAreJSONAndHaveExpectedShape(t *testing.T) {
	fixtures := map[string]string{
		"benign_obvious.json":    "likely_safe",
		"malicious_obvious.json": "suspicious",
		"ambiguous.json":         "inconclusive",
	}
	for filename, expected := range fixtures {
		filename, expected := filename, expected
		t.Run(filename, func(t *testing.T) {
			content, err := os.ReadFile(filepath.Join(".", filename))
			if err != nil {
				t.Fatal(err)
			}
			var packageData map[string]any
			if err := json.Unmarshal(content, &packageData); err != nil {
				t.Fatalf("fixture is not JSON: %v", err)
			}
			for _, key := range []string{"schema_version", "collection_id", "processes", "persistence", "scheduled_tasks", "services", "wmi_subscriptions", "winlogon", "network", "defender_events", "limitations", "integrity"} {
				if _, ok := packageData[key]; !ok {
					t.Errorf("missing required top-level key %q", key)
				}
			}
			// The expected value is metadata for the test plan. The LLM output
			// must still be evaluated with evidence references and limitations.
			if expected == "" {
				t.Fatal("expected verdict label is empty")
			}
		})
	}
}
