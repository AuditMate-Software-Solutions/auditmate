//go:build linux
// +build linux

package collectors

import (
	"os"
	"strings"
)

// CollectHostID returns a unique host identifier for Linux.
// Prefers machine-id files, falls back to hostname.
func CollectHostID() (string, error) {
	paths := []string{
		"/etc/machine-id",
		"/var/lib/dbus/machine-id",
	}

	for _, p := range paths {
		if data, err := os.ReadFile(p); err == nil {
			id := strings.ToLower(strings.TrimSpace(string(data)))
			if id != "" {
				return id, nil
			}
		}
	}

	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}

	return strings.ToLower(strings.TrimSpace(hostname)), nil
}



