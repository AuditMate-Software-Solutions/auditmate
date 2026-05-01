//go:build windows

package collectors

import (
	"os"
	"strings"

	"golang.org/x/sys/windows/registry"
)

func CollectHostID() (string, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE,
		`SOFTWARE\Microsoft\Cryptography`,
		registry.QUERY_VALUE)

	if err == nil {
		defer k.Close()
		g, _, err := k.GetStringValue("MachineGuid")
		if err == nil && g != "" {
			return sanitize(g), nil
		}
	}

	h, err := os.Hostname()
	if err != nil {
		return "", err
	}
	return sanitize(h), nil
}

func sanitize(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	return strings.ReplaceAll(s, "-", "")
}



