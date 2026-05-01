package models

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"sort"
	"strings"
	"time"
)

const SchemaVersion = 2

type Snapshot struct {
	SchemaVersion int    `json:"schema_version"`
	Timestamp     string `json:"timestamp"`

	HostID     string `json:"host_id"`
	SnapshotID string `json:"snapshot_id"`

	Hostname string `json:"hostname"`
	OS       string `json:"os"`
	Kernel   string `json:"kernel"`
	Uptime   string `json:"uptime"`

	CoreHash      string   `json:"core_hash,omitempty"`
	FirewallState string   `json:"firewall_state"`
	OpenPorts     []string `json:"open_ports"`

	Admins       []string `json:"admins"`
	RegularUsers []string `json:"regular_users"`
	Services     []string `json:"services"`

	Packages []string `json:"packages"`
}

func NewSnapshot() *Snapshot {
	return &Snapshot{
		SchemaVersion: SchemaVersion,
		SnapshotID:    generateSnapshotID(),
	}
}

func generateSnapshotID() string {
	h := sha256.Sum256([]byte(time.Now().UTC().Format(time.RFC3339Nano)))
	return hex.EncodeToString(h[:])
}

func (s *Snapshot) Normalize() {
	if s == nil {
		return
	}

	s.Admins = normalize(s.Admins)
	s.RegularUsers = normalize(s.RegularUsers)
	s.Services = normalize(s.Services)
	s.Packages = normalize(s.Packages)
	s.OpenPorts = normalize(s.OpenPorts)
}

func (s *Snapshot) GenerateCoreHash() string {
	if s == nil {
		return ""
	}

	// IMPORTANT FIX: fully deterministic ordering
	parts := []string{
		s.HostID,
		s.Hostname,
		s.OS,
		s.Kernel,
		s.FirewallState,
	}

	parts = append(parts,
		joinSorted(s.Services),
		joinSorted(s.Packages),
		joinSorted(s.OpenPorts),
	)

	sort.Strings(parts)

	h := sha256.Sum256([]byte(strings.Join(parts, "|")))
	return hex.EncodeToString(h[:])
}

func joinSorted(in []string) string {
	cp := append([]string{}, in...)
	sort.Strings(cp)
	return strings.Join(cp, ",")
}

func (s *Snapshot) Finalize() {
	if s == nil {
		return
	}
	s.Normalize()
	s.CoreHash = s.GenerateCoreHash()
}

func normalize(in []string) []string {
	set := map[string]struct{}{}

	for _, v := range in {
		v = strings.ToLower(strings.TrimSpace(v))
		if v != "" {
			set[v] = struct{}{}
		}
	}

	out := make([]string, 0, len(set))
	for v := range set {
		out = append(out, v)
	}

	sort.Strings(out)
	return out
}

func (s *Snapshot) SetTimestamp() {
	s.Timestamp = time.Now().UTC().Format(time.RFC3339)
}

func LoadSnapshot(path string) (*Snapshot, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var out Snapshot
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, err
	}

	return &out, nil
}