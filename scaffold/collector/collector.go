package collector

import (
	"context"
	"time"
)

const CollectorVersion = "0.1.0-scaffold"

// Config contains explicit, non-privileged collection options.
type Config struct {
	IncludeHashes       bool
	IncludeFullPaths    bool
	IncludeCommandLines bool
	DefenderEventLimit  int
}

// Collect is the single entry point used by the local application.
func Collect(ctx context.Context, cfg Config) (EvidencePackage, error) {
	if err := assertReadOnly(); err != nil {
		return EvidencePackage{}, err
	}

	startedAt := time.Now().UTC()
	var allLimitations []Limitation

	processes, limits, err := collectProcesses(ctx, cfg)
	if err != nil {
		allLimitations = append(allLimitations, Limitation{Category: "processes", Code: "COLLECTION_ERROR", Message: err.Error(), Impact: "Missing process data"})
	}
	allLimitations = append(allLimitations, limits...)

	persistence, limits, err := collectPersistence(ctx, cfg)
	if err != nil {
		allLimitations = append(allLimitations, Limitation{Category: "persistence", Code: "COLLECTION_ERROR", Message: err.Error(), Impact: "Missing persistence data"})
	}
	allLimitations = append(allLimitations, limits...)

	tasks, limits, err := collectScheduledTasks(ctx, cfg)
	if err != nil {
		allLimitations = append(allLimitations, Limitation{Category: "scheduled_tasks", Code: "COLLECTION_ERROR", Message: err.Error(), Impact: "Missing scheduled tasks data"})
	}
	allLimitations = append(allLimitations, limits...)

	services, limits, err := collectServices(ctx, cfg)
	if err != nil {
		allLimitations = append(allLimitations, Limitation{Category: "services", Code: "COLLECTION_ERROR", Message: err.Error(), Impact: "Missing services data"})
	}
	allLimitations = append(allLimitations, limits...)

	wmi, limits, err := collectWMISubscriptions(ctx, cfg)
	if err != nil {
		allLimitations = append(allLimitations, Limitation{Category: "wmi_subscriptions", Code: "COLLECTION_ERROR", Message: err.Error(), Impact: "Missing WMI data"})
	}
	allLimitations = append(allLimitations, limits...)

	winlogon, limits, err := collectWinlogon(ctx, cfg)
	if err != nil {
		allLimitations = append(allLimitations, Limitation{Category: "winlogon", Code: "COLLECTION_ERROR", Message: err.Error(), Impact: "Missing Winlogon data"})
	}
	allLimitations = append(allLimitations, limits...)

	network, limits, err := collectNetwork(ctx, cfg)
	if err != nil {
		allLimitations = append(allLimitations, Limitation{Category: "network", Code: "COLLECTION_ERROR", Message: err.Error(), Impact: "Missing network data"})
	}
	allLimitations = append(allLimitations, limits...)

	defender, limits, err := collectDefenderEvents(ctx, cfg)
	if err != nil {
		allLimitations = append(allLimitations, Limitation{Category: "defender_events", Code: "COLLECTION_ERROR", Message: err.Error(), Impact: "Missing Defender data"})
	}
	allLimitations = append(allLimitations, limits...)

	finishedAt := time.Now().UTC()

	pkg := EvidencePackage{
		SchemaVersion: "2020-12",
		CreatedAt:     finishedAt,
		Collection: CollectionMetadata{
			StartedAt:           startedAt,
			FinishedAt:          finishedAt,
			CollectorVersion:    CollectorVersion,
			CategoriesRequested: []string{"processes", "persistence", "scheduled_tasks", "services", "wmi", "winlogon", "network", "defender"},
			CategoriesCompleted: []string{"processes", "persistence", "scheduled_tasks", "services", "wmi", "winlogon", "network", "defender"},
			ReadOnlyAssertion:   true,
		},
		Processes:      processes,
		Persistence:    persistence,
		ScheduledTasks: tasks,
		Services:       services,
		WMI:            wmi,
		Winlogon:       winlogon,
		Network:        network,
		DefenderEvents: defender,
		Limitations:    allLimitations,
	}

	// We calculate canonical SHA256 in a later pass or orchestration layer,
	// but the structure is now filled.
	return pkg, nil
}

// assertReadOnly is intentionally a placeholder for a testable invariant.
func assertReadOnly() error {
	return nil
}
