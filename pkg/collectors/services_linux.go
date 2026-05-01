package collectors

import (
	"context"
	"fmt"
	"os/exec"
	"sort"
	"strings"
	"time"
)

const serviceTimeout = 5 * time.Second

func CollectServices() ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), serviceTimeout)
	defer cancel()

	cmd := exec.CommandContext(
		ctx,
		"systemctl",
		"list-units",
		"--type=service",
		"--state=running",
		"--no-pager",
		"--no-legend",
	)

	out, err := cmd.Output()
	if ctx.Err() == context.DeadlineExceeded {
		return nil, fmt.Errorf("services collection timed out after %s", serviceTimeout)
	}
	if err != nil {
		return nil, fmt.Errorf("systemctl failed: %w", err)
	}

	lines := strings.Split(string(out), "\n")

	seen := map[string]struct{}{}
	services := make([]string, 0)

	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}

		fields := strings.Fields(l)
		if len(fields) == 0 {
			continue
		}

		svc := fields[0]

		// normalize
		svc = strings.ToLower(strings.TrimSuffix(svc, ".service"))

		// strip templates safely
		if i := strings.Index(svc, "@"); i != -1 {
			svc = svc[:i]
		}

		if svc == "" {
			continue
		}

		if _, ok := seen[svc]; ok {
			continue
		}

		seen[svc] = struct{}{}
		services = append(services, svc)
	}

	sort.Strings(services)
	return services, nil
}