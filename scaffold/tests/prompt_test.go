package tests

import (
	"context"
	"testing"
	
	"github.com/Jonathan-6dward/Boot.Check/scaffold/api"
	"github.com/Jonathan-6dward/Boot.Check/scaffold/api/provider"
)

type MockProvider struct {
	Response provider.Verdict
	Err      error
}

func (m *MockProvider) Kind() provider.Kind { return provider.KindLocal }
func (m *MockProvider) Name() string        { return "mock" }
func (m *MockProvider) Available(ctx context.Context) error { return m.Err }
func (m *MockProvider) Analyze(ctx context.Context, pkg provider.EvidencePackage) (provider.Verdict, error) {
	return m.Response, m.Err
}

func TestAnalyzeMapsProviderResponse(t *testing.T) {
	mockProv := &MockProvider{
		Response: provider.Verdict{
			State:        provider.StateLikelySafe,
			Summary:      "Tudo seguro.",
			ProviderKind: provider.KindLocal,
			ModelName:    "mock",
			MitreATTACK:  []string{"T1059"},
		},
	}
	
	ctx := context.Background()
	vr, err := api.Analyze(ctx, mockProv, provider.EvidencePackage{}, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	if vr.Verdict != "likely_safe" {
		t.Errorf("expected likely_safe, got %s", vr.Verdict)
	}
	if vr.HeadlinePlain != "Tudo seguro." {
		t.Errorf("expected Tudo seguro., got %s", vr.HeadlinePlain)
	}
	if len(vr.MITRE) != 1 || vr.MITRE[0].TechniqueID != "T1059" {
		t.Errorf("MITRE mapping failed")
	}
}

func TestAnalyzeMapsSchemaErrorToInconclusive(t *testing.T) {
	mockProv := &MockProvider{
		Err: provider.ErrSchemaInvalid,
	}
	
	ctx := context.Background()
	vr, err := api.Analyze(ctx, mockProv, provider.EvidencePackage{}, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	if vr.Verdict != "inconclusive" {
		t.Errorf("expected inconclusive on schema error, got %s", vr.Verdict)
	}
}
