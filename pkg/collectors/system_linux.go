package collectors

import (
	"os"
	"os/exec"
	"strings"

	"auditmate/pkg/models"
)

func FillSystem(s *models.Snapshot) {
	s.Hostname = normalizeValue(runCmd("hostname"))
	s.OS = normalizeValue(getOSDescription())
	s.Kernel = normalizeValue(runCmd("uname", "-r"))
	s.Uptime = normalizeValue(runCmd("uptime", "-p"))

	s.FirewallState = normalizeValue(CollectFirewall())
}

func runCmd(cmd string, args ...string) string {
	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return "unknown"
	}
	return strings.TrimSpace(string(out))
}

func getOSDescription() string {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return "unknown"
	}

	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "PRETTY_NAME=") {
			v := strings.TrimPrefix(line, "PRETTY_NAME=")
			return strings.Trim(strings.ToLower(v), "\"")
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