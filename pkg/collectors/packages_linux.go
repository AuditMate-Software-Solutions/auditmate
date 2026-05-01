package collectors

import (
	"os/exec"
	"sort"
	"strings"
)

func CollectPackages() ([]string, error) {
	var out []byte
	var err error

	if _, err = exec.LookPath("dpkg-query"); err == nil {
		out, err = exec.Command("dpkg-query", "-W", "-f=${binary:Package}\n").Output()
	} else if _, err = exec.LookPath("rpm"); err == nil {
		out, err = exec.Command("rpm", "-qa", "--qf", "%{NAME}\n").Output()
	} else {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(out), "\n")

	seen := map[string]struct{}{}
	pkgs := make([]string, 0)

	for _, p := range lines {
		p = strings.ToLower(strings.TrimSpace(p))
		if p == "" {
			continue
		}

		if _, ok := seen[p]; !ok {
			seen[p] = struct{}{}
			pkgs = append(pkgs, p)
		}
	}

	sort.Strings(pkgs)
	return pkgs, nil
}