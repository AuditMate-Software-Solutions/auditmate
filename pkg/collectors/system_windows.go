package collectors

import (
	"os"
	"os/exec"
	"strings"

	"auditmate/pkg/models"
)

func FillSystem(s *models.Snapshot) {
	host, _ := os.Hostname()
	s.Hostname = normalizeValue(host)

	out, err := exec.Command("cmd", "/c", "wmic os get Caption").Output()
	if err != nil {
		s.OS = "unknown"
	} else {
		lines := strings.Split(strings.TrimSpace(string(out)), "\n")
		if len(lines) > 1 {
			s.OS = normalizeValue(lines[1])
		} else {
			s.OS = "unknown"
		}
	}

	s.Uptime = normalizeValue(getWindowsUptime())
	s.FirewallState = normalizeValue(CollectFirewall())
}

func getWindowsUptime() string {
	out, err := exec.Command("net", "stats", "srv").Output()
	if err != nil {
		return "unknown"
	}

	lines := strings.Split(string(out), "\n")
	for _, l := range lines {
		l = strings.ToLower(l)
		if strings.Contains(l, "statistics since") {
			return strings.TrimSpace(l)
		}
	}
	return "unknown"
}

func normalizeValue(v string) string {
	v = strings.ToLower(strings.TrimSpace(v))
	if v == "" {
		return "unknown"
	}
	return v
}