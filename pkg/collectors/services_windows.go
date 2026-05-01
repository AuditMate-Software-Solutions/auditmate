//go:build windows

package collectors

import (
	"context"
	"fmt"
	"os/exec"
	"sort"
	"strings"
	"time"
)

// FIX: prevent PowerShell hang
const windowsServiceTimeout = 6 * time.Second

// CollectServices returns running services on Windows in a stable format.
func CollectServices() ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), windowsServiceTimeout)
	defer cancel()

	ps := "Get-Service | Where-Object {$_.Status -eq 'Running'} | Select-Object -ExpandProperty Name"

	cmd := exec.CommandContext(
		ctx,
		"powershell",
		"-NoProfile",
		"-Command",
		ps,
	)

	out, err := cmd.Output()

	if ctx.Err() == context.DeadlineExceeded {
		return nil, fmt.Errorf("services collection timed out after %s", windowsServiceTimeout)
	}
	if err != nil {
		return nil, fmt.Errorf("powershell services failed: %w", err)
	}

	raw := strings.TrimSpace(string(out))
	if raw == "" {
		return []string{}, nil
	}

	lines := strings.Split(raw, "\n")

	seen := make(map[string]struct{})
	services := make([]string, 0)

	for _, s := range lines {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}

		s = strings.ToLower(s)

		if _, ok := seen[s]; ok {
			continue
		}

		seen[s] = struct{}{}
		services = append(services, s)
	}

	sort.Strings(services)
	return services, nil
}


