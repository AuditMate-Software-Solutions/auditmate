//go:build windows

package collectors

import (
	"os/exec"
	"strings"
)

func CollectFirewall() string {
	out, err := exec.Command(
		"netsh",
		"advfirewall",
		"show",
		"allprofiles",
	).Output()

	if err != nil {
		return "unknown"
	}

	s := strings.ToLower(string(out))

	if strings.Contains(s, "state on") {
		return "active"
	}

	if strings.Contains(s, "state off") {
		return "inactive"
	}

	return "unknown"
}


