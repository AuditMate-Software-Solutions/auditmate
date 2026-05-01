package collectors

import (
	"fmt"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

type portEntryWindows struct {
	port int
	raw  string
}

func CollectPorts() ([]string, error) {
	if _, err := exec.LookPath("netstat"); err != nil {
		return nil, fmt.Errorf("netstat not available: %w", err)
	}

	out, err := exec.Command("netstat", "-ano").Output()
	if err != nil {
		return nil, fmt.Errorf("netstat command failed: %w", err)
	}

	lines := strings.Split(string(out), "\n")

	seen := make(map[string]struct{})
	entries := make([]portEntryWindows, 0)

	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}

		l = strings.ToUpper(l)

		if !strings.Contains(l, "LISTENING") {
			continue
		}

		fields := strings.Fields(l)
		if len(fields) < 2 {
			continue
		}

		addr := fields[1]
		portStr := extractWinPort(addr)

		if portStr == "" {
			continue
		}

		port, err := strconv.Atoi(portStr)
		if err != nil {
			continue
		}

		key := portStr + "/tcp"
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}

		entries = append(entries, portEntryWindows{
			port: port,
			raw:  key,
		})
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].port < entries[j].port
	})

	result := make([]string, 0, len(entries))
	for _, e := range entries {
		result = append(result, e.raw)
	}

	return result, nil
}

func extractWinPort(addr string) string {
	parts := strings.Split(addr, ":")
	if len(parts) == 0 {
		return ""
	}

	p := parts[len(parts)-1]
	if p == "" {
		return ""
	}

	if _, err := strconv.Atoi(p); err != nil {
		return ""
	}

	return p
}


