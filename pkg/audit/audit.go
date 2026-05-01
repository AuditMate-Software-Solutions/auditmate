package audit

import (
        "encoding/json"
        "fmt"
        "os"
        "path/filepath"
        "time"

        "auditmate/pkg/collectors"
        "auditmate/pkg/diff"
        "auditmate/pkg/models"
)

type AuditState struct {
        Snapshot models.Snapshot `json:"snapshot"`
        Time     string          `json:"time"`
}

type RunResult struct {
        Hostname string
        OS       string
        Uptime   string

        Users    int
        Services int
        Packages int
        Ports    int

        Status string

        Diffs    []string
        Warnings []string
        Errors   bool
        Reset    bool
        FirstRun bool

        Duration   time.Duration
        ReportPath string
}

func Run(outputDir string, reset bool) (*RunResult, error) {
        start := time.Now()

        if err := os.MkdirAll(outputDir, 0755); err != nil {
                return nil, fmt.Errorf("failed to create output dir: %w", err)
        }

        reportPath := filepath.Join(outputDir, "audit.json")

        // -------------------------
        // LOAD BASELINE
        // -------------------------
        var baseline AuditState
        data, err := os.ReadFile(reportPath)
        exists := err == nil && len(data) > 0

        if exists {
                _ = json.Unmarshal(data, &baseline)
        }

        firstRun := !exists

        // -------------------------
        // INTEGRITY CHECK
        // -------------------------
        integrityViolation := false
        if exists {
                expected := baseline.Snapshot.CoreHash
                actual := baseline.Snapshot.GenerateCoreHash()

                if expected != "" && expected != actual {
                        integrityViolation = true
                }
        }

        // -------------------------
        // CURRENT SNAPSHOT
        // -------------------------
        s := models.NewSnapshot()
        s.SetTimestamp()

        sys := models.Snapshot{}
        collectors.FillSystem(&sys)

        hostID, _ := collectors.CollectHostID()
        s.HostID = hostID

        s.Hostname = sys.Hostname
        s.OS = sys.OS
        s.Kernel = sys.Kernel
        s.Uptime = sys.Uptime
        s.FirewallState = sys.FirewallState

        services, _ := collectors.CollectServices()
        pkgs, _ := collectors.CollectPackages()
        ports, _ := collectors.CollectPorts()

        s.Services = services
        s.Packages = pkgs
        s.OpenPorts = ports

        collectors.FillUsers(s)

        s.Normalize()
        s.Finalize()

        // -------------------------
        // DECISION LOGIC
        // -------------------------
        var diffs []string
        status := "CLEAN"
        errors := false

        switch {
        case firstRun:
                status = "INITIAL SNAPSHOT"

        case reset:
                status = "SNAPSHOT RESET"

        case integrityViolation:
                // 🔥 HARD FAIL
                status = "TAMPER DETECTED"
                errors = true

        default:
                diffs = diff.DiffSnapshot(&baseline.Snapshot, s)

                if len(diffs) > 0 {
                        status = "DRIFT DETECTED"
                }
        }

        // -------------------------
        // BASELINE UPDATE RULE
        // -------------------------
        if firstRun || reset {
                baseline = AuditState{
                        Snapshot: *s,
                        Time:     time.Now().UTC().Format(time.RFC3339),
                }
        }

        // ALWAYS persist baseline
        b, err := json.MarshalIndent(baseline, "", "  ")
        if err != nil {
                return nil, fmt.Errorf("failed to marshal state: %w", err)
        }

        if err := os.WriteFile(reportPath, b, 0644); err != nil {
                return nil, fmt.Errorf("failed to write report: %w", err)
        }

        return &RunResult{
                Hostname: s.Hostname,
                OS:       s.OS,
                Uptime:   s.Uptime,

                Users:    len(s.Admins) + len(s.RegularUsers),
                Services: len(s.Services),
                Packages: len(s.Packages),
                Ports:    len(s.OpenPorts),

                Status: status,

                Diffs:    diffs,
                Warnings: nil,
                Errors:   errors,
                Reset:    reset,
                FirstRun: firstRun,

                Duration:   time.Since(start),
                ReportPath: reportPath,
        }, nil
}