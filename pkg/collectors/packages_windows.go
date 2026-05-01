package collectors

import (
	"fmt"
	"os/exec"
	"sort"
	"strings"
)

// CollectPackages returns deterministic installed package list on Windows.
func CollectPackages() ([]string, error) {
	ps := `
$names = @()

$keys = @(
  "HKLM:\Software\Microsoft\Windows\CurrentVersion\Uninstall\*",
  "HKLM:\Software\WOW6432Node\Microsoft\Windows\CurrentVersion\Uninstall\*",
  "HKCU:\Software\Microsoft\Windows\CurrentVersion\Uninstall\*"
)

foreach ($k in $keys) {
  Get-ItemProperty -Path $k -ErrorAction SilentlyContinue |
    Where-Object { $_.DisplayName } |
    ForEach-Object { $names += $_.DisplayName }
}

Get-AppxPackage |
  ForEach-Object { $names += $_.Name }

$names |
  ForEach-Object { $_.Trim().ToLower() } |
  Where-Object { $_ -ne "" } |
  Sort-Object -Unique
`

	out, err := exec.Command(
		"powershell",
		"-NoProfile",
		"-ExecutionPolicy",
		"Bypass",
		"-Command",
		ps,
	).Output()

	if err != nil {
		return nil, fmt.Errorf("packages collection failed: %w", err)
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")

	seen := make(map[string]struct{})
	pkgs := make([]string, 0)

	for _, v := range lines {
		v = strings.TrimSpace(strings.ToLower(v))
		if v == "" {
			continue
		}

		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			pkgs = append(pkgs, v)
		}
	}

	sort.Strings(pkgs)
	return pkgs, nil
}


