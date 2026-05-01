package collectors

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"

	"auditmate/pkg/models"
)

const humanUIDMin = 1000

var excludedUsers = map[string]bool{
	"nobody":   true,
	"daemon":   true,
	"syslog":   true,
	"games":    true,
	"sync":     true,
	"halt":     true,
	"shutdown": true,
}

func readSudoUsers() map[string]bool {
	admins := make(map[string]bool)

	file, err := os.Open("/etc/group")
	if err != nil {
		return admins
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Split(line, ":")
		if len(parts) < 4 {
			continue
		}

		group := parts[0]
		if group != "sudo" && group != "wheel" {
			continue
		}

		users := strings.Split(parts[3], ",")
		for _, u := range users {
			u = strings.TrimSpace(u)
			if u != "" {
				admins[u] = true
			}
		}
	}

	return admins
}

func FillUsers(s *models.Snapshot) {
	if s == nil {
		return
	}

	file, err := os.Open("/etc/passwd")
	if err != nil {
		s.Admins = []string{"root"}
		s.RegularUsers = []string{}
		return
	}
	defer file.Close()

	sudoUsers := readSudoUsers()

	adminSet := map[string]bool{}
	regularSet := map[string]bool{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Split(line, ":")
		if len(parts) < 3 {
			continue
		}

		username := parts[0]

		if excludedUsers[username] {
			continue
		}

		uid, err := strconv.Atoi(parts[2])
		if err != nil {
			continue
		}

		isRoot := uid == 0
		isSudo := sudoUsers[username]

		if isRoot || isSudo {
			adminSet[username] = true
			delete(regularSet, username)
			continue
		}

		if uid >= humanUIDMin {
			if !adminSet[username] {
				regularSet[username] = true
			}
		}
	}

	admins := make([]string, 0, len(adminSet))
	regular := make([]string, 0, len(regularSet))

	for u := range adminSet {
		admins = append(admins, u)
	}
	for u := range regularSet {
		regular = append(regular, u)
	}

	if len(admins) == 0 {
		admins = []string{"root"}
	}

	sort.Strings(admins)
	sort.Strings(regular)

	s.Admins = admins
	s.RegularUsers = regular
}