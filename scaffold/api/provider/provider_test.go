package provider

import (
	"context"
	"testing"
	"time"
)

func samplePackage() EvidencePackage {
	return EvidencePackage{
		SchemaVersion: "1.0",
		CollectedAt:   time.Now(),
		Items: []Evidence{
			{ID: "ev-001", Kind: "process", Level: LevelObserved, Value: "powershell.exe"},
			{ID: "ev-002", Kind: "run_key", Level: LevelObserved, Value: "HKCU\\...\\Run"},
		},
	}
}

func TestValidateVerdict_OK(t *testing.T) {
	pkg := samplePackage()
	raw := []byte(`{
		"state": "suspicious",
		"summary": "teste",
		"claims": [
			{"what":"x","who":"y","when":"z","where":"w","why":"q","how":"h","impact":"i","evidence_ids":["ev-001"]}
		],
		"mitre_attack": ["T1059.001"]
	}`)
	v, err := ValidateVerdict(raw, pkg)
	if err != nil {
		t.Fatalf("esperava sucesso, veio erro: %v", err)
	}
	if v.State != StateSuspicious {
		t.Fatalf("state incorreto: %v", v.State)
	}
}

func TestValidateVerdict_EstadoInvalido(t *testing.T) {
	pkg := samplePackage()
	raw := []byte(`{"state": "definitely_malware", "summary": "x", "claims": []}`)
	_, err := ValidateVerdict(raw, pkg)
	if err != ErrSchemaInvalid {
		t.Fatalf("esperava ErrSchemaInvalid, veio: %v", err)
	}
}

func TestValidateVerdict_EvidenceIdOrfao(t *testing.T) {
	pkg := samplePackage()
	raw := []byte(`{
		"state": "likely_safe",
		"summary": "x",
		"claims": [
			{"what":"x","who":"y","when":"z","where":"w","why":"q","how":"h","impact":"i","evidence_ids":["ev-999"]}
		]
	}`)
	_, err := ValidateVerdict(raw, pkg)
	if err != ErrSchemaInvalid {
		t.Fatalf("esperava ErrSchemaInvalid para evidence_id inexistente, veio: %v", err)
	}
}

func TestLocalOllamaProvider_Name(t *testing.T) {
	p := NewLocalOllamaProvider("http://localhost:11434", "qwen2.5:7b-instruct", "sys")
	if p.Kind() != KindLocal {
		t.Fatalf("kind incorreto")
	}
	if p.Name() == "" {
		t.Fatalf("nome vazio")
	}
}

func TestCloudProvider_AvailableSemChave(t *testing.T) {
	p := NewCloudProvider(CloudCredentials{Vendor: VendorAnthropic, Model: "claude-haiku-4-5"}, "sys")
	if err := p.Available(context.Background()); err == nil {
		t.Fatalf("esperava erro por falta de chave")
	}
}
