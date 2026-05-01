package collectors

import (
	"fmt"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

func CollectPorts() ([]string, error) {
	if _, err := exec.LookPath("ss"); err != nil {
		return nil, fmt.Errorf("ss not available: %w", err)
	}

	out, err := exec.Command("ss", "-tuln").Output()
	if err != nil {
		return nil, fmt.Errorf("ss command failed: %w", err)
	}

	lines := strings.Split(string(out), "\n")

	seen := map[string]struct{}{}
	ports := make([]string, 0)

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" || strings.HasPrefix(line, "Netid") {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) < 5 {
			continue
		}

		proto := strings.ToLower(fields[0])
		addr := fields[len(fields)-1]

		port := extractPort(addr)
		if port == "" {
			continue
		}

		key := port + "/" + proto
		if _, ok := seen[key]; ok {
			continue
		}

		seen[key] = struct{}{}
		ports = append(ports, key)
	}

	sort.Strings(ports)
	return ports, nil
}

func extractPort(addr string) string {
	// IPv6: [::]:443
	if strings.Contains(addr, "]:") {
		parts := strings.Split(addr, "]:")
		if len(parts) == 2 {
			return parts[1]
		}
	}

	// IPv4: 0.0.0.0:22
	if strings.Contains(addr, ":") {
		i := strings.LastIndex(addr, ":")
		if i != -1 {
			p := addr[i+1:]
			if _, err := strconv.Atoi(p); err == nil {
				return p
			}
		}
	}

	return ""
}