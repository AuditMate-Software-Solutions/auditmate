//go:build linux
// +build linux

package collectors

import (
	"os"
	"os/exec"
	"strings"
)

func CollectFirewall() string {
	if isWSL() {
		return "not_applicable"
	}

	if status, err := checkUFW(); err == nil {
		return status
	}

	if status, err := checkNFTables(); err == nil {
		return status
	}

	if status, err := checkIPTables(); err == nil {
		return status
	}

	return "inactive"
}

func isWSL() bool {
	data, err := os.ReadFile("/proc/version")
	if err != nil {
		return false
	}
	s := strings.ToLower(string(data))
	return strings.Contains(s, "microsoft") || strings.Contains(s, "wsl")
}

func checkUFW() (string, error) {
	if _, err := exec.LookPath("ufw"); err != nil {
		return "", err
	}

	out, err := exec.Command("ufw", "status").Output()
	if err != nil {
		return "", err
	}

	if strings.Contains(strings.ToLower(string(out)), "status: active") {
		return "active", nil
	}
	return "inactive", nil
}

func checkNFTables() (string, error) {
	if _, err := exec.LookPath("nft"); err != nil {
		return "", err
	}

	out, err := exec.Command("nft", "list", "ruleset").Output()
	if err != nil {
		return "", err
	}

	if len(strings.TrimSpace(string(out))) > 0 {
		return "active", nil
	}
	return "inactive", nil
}

func checkIPTables() (string, error) {
	if _, err := exec.LookPath("iptables"); err != nil {
		return "", err
	}

	out, err := exec.Command("iptables", "-L").Output()
	if err != nil {
		return "", err
	}

	if strings.Contains(string(out), "Chain") {
		return "active", nil
	}
	return "inactive", nil
}


