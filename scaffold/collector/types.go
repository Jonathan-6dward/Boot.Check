package collector

import "time"

// EvidencePackage is the only artifact that leaves the collector boundary.
// TODO(local-agent): keep this structure aligned with evidence.schema.json and
// reject unknown security-sensitive fields before serialization.
type EvidencePackage struct {
	SchemaVersion  string              `json:"schema_version"`
	CollectionID   string              `json:"collection_id"`
	CreatedAt      time.Time           `json:"created_at"`
	Consent        ConsentRecord       `json:"consent"`
	Host           HostSnapshot        `json:"host"`
	Collection     CollectionMetadata  `json:"collection"`
	Processes      []ProcessEvidence   `json:"processes"`
	Persistence    PersistenceEvidence `json:"persistence"`
	ScheduledTasks CategoryResult      `json:"scheduled_tasks"`
	Services       CategoryResult      `json:"services"`
	WMI            CategoryResult      `json:"wmi_subscriptions"`
	Winlogon       CategoryResult      `json:"winlogon"`
	Network        NetworkEvidence     `json:"network"`
	DefenderEvents CategoryResult      `json:"defender_events"`
	Limitations    []Limitation        `json:"limitations"`
	Integrity      IntegrityMetadata   `json:"integrity"`
}

type ConsentRecord struct {
	LocalCollectionConfirmed bool      `json:"local_collection_confirmed"`
	LLMSubmissionConfirmed   bool      `json:"llm_submission_confirmed"`
	Purpose                  string    `json:"purpose"`
	ProviderName             string    `json:"provider_name"`
	DataMode                 string    `json:"data_mode"`
	RetentionAcknowledged    bool      `json:"retention_acknowledged,omitempty"`
	RecordedAt               time.Time `json:"recorded_at"`
}

type HostSnapshot struct {
	OSFamily         string `json:"os_family"`
	OSVersion        string `json:"os_version"`
	BuildNumber      string `json:"build_number,omitempty"`
	Architecture     string `json:"architecture"`
	Hostname         string `json:"hostname"`
	InteractiveUser  string `json:"interactive_user"`
	PrivilegeContext string `json:"privilege_context"`
	DefenderPresent  bool   `json:"defender_present,omitempty"`
}

type CollectionMetadata struct {
	StartedAt           time.Time `json:"started_at"`
	FinishedAt          time.Time `json:"finished_at"`
	CollectorVersion    string    `json:"collector_version"`
	CategoriesRequested []string  `json:"categories_requested"`
	CategoriesCompleted []string  `json:"categories_completed"`
	ReadOnlyAssertion   bool      `json:"read_only_assertion"`
}

type EvidenceBase struct {
	EvidenceID      string    `json:"evidence_id"`
	Source          string    `json:"source"`
	ObservedAt      time.Time `json:"observed_at"`
	Sensitivity     string    `json:"sensitivity"`
	RedactionStatus string    `json:"redaction_status"`
	Notes           string    `json:"notes,omitempty"`
}

type ProcessEvidence struct {
	EvidenceBase
	PID              uint32        `json:"pid"`
	ParentPID        *uint32       `json:"parent_pid"`
	ImageName        string        `json:"image_name"`
	ImagePath        *string       `json:"image_path"`
	CommandLine      *string       `json:"command_line"`
	StartTime        *time.Time    `json:"start_time"`
	IntegrityLevel   string        `json:"integrity_level"`
	User             *string       `json:"user,omitempty"`
	Signature        SignatureInfo `json:"signature"`
	SHA256           *string       `json:"sha256,omitempty"`
	HashStatus       string        `json:"hash_status"`
	ParentEvidenceID *string       `json:"parent_evidence_id,omitempty"`
}

type PersistenceEvidence struct {
	Run     []RegistryAutorunEvidence `json:"run"`
	RunOnce []RegistryAutorunEvidence `json:"run_once"`
}

type RegistryAutorunEvidence struct {
	EvidenceBase
	Hive      string  `json:"hive"`
	KeyPath   string  `json:"key_path"`
	ValueName string  `json:"value_name"`
	ValueData *string `json:"value_data"`
	Scope     string  `json:"scope"`
}

type CategoryResult struct {
	Status       string           `json:"status"`
	Items        []map[string]any `json:"items"`
	ErrorCode    *string          `json:"error_code,omitempty"`
	ErrorMessage *string          `json:"error_message,omitempty"`
}

type NetworkEvidence struct {
	Status         string               `json:"status"`
	TCPConnections []ConnectionEvidence `json:"tcp_connections"`
	UDPEndpoints   []ConnectionEvidence `json:"udp_endpoints"`
}

type ConnectionEvidence struct {
	EvidenceBase
	Protocol      string  `json:"protocol"`
	LocalAddress  string  `json:"local_address"`
	LocalPort     uint16  `json:"local_port"`
	RemoteAddress *string `json:"remote_address"`
	RemotePort    *uint16 `json:"remote_port"`
	State         string  `json:"state"`
	OwningPID     *uint32 `json:"owning_pid"`
	OwningImage   *string `json:"owning_image,omitempty"`
}

type Limitation struct {
	Category string `json:"category"`
	Code     string `json:"code"`
	Message  string `json:"message"`
	Impact   string `json:"impact"`
}

type IntegrityMetadata struct {
	Canonicalization  string  `json:"canonicalization"`
	SHA256            string  `json:"sha256"`
	ReadOnlyAssertion bool    `json:"read_only_assertion"`
	SigningKeyID      *string `json:"signing_key_id,omitempty"`
}

type SignatureInfo struct {
	Status             string     `json:"status"`
	Publisher          *string    `json:"publisher,omitempty"`
	CertificateSubject *string    `json:"certificate_subject,omitempty"`
	CheckedAt          *time.Time `json:"checked_at,omitempty"`
}
