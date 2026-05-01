package collectors

import (
	"os/exec"
	"sort"
	"strings"

	"auditmate/pkg/models"
)

// FillUsers populates Snapshot with Windows users
func FillUsers(s *models.Snapshot) {
	s.Admins = collectWindowsUsers(
		"Get-LocalGroupMember -Group 'Administrators' | Select-Object -ExpandProperty Name",
	)

	s.RegularUsers = collectWindowsUsers(
		"Get-LocalUser | Where-Object {$_.Enabled} | Select-Object -ExpandProperty Name",
	)
}

func collectWindowsUsers(cmd string) []string {
	out, err := exec.Command("powershell", "-NoProfile", "-Command", cmd).Output()
	if err != nil {
		return []string{}
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")

	seen := make(map[string]struct{})
	result := make([]string, 0)

	for _, l := range lines {
		l = strings.TrimSpace(strings.ToLower(l))
		if l == "" {
			continue
		}

		if _, ok := seen[l]; ok {
			continue
		}

		seen[l] = struct{}{}
		result = append(result, l)
	}

	sort.Strings(result)
	return result
}