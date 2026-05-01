package diff

import (
	"fmt"
	"sort"
	"strings"

	"auditmate/pkg/models"
)

type Diff struct {
	Category string   `json:"category"`
	Added    []string `json:"added"`
	Removed  []string `json:"removed"`
}

func DiffSnapshot(oldS, newS *models.Snapshot) []string {
	structured := diffSnapshotStructured(*oldS, *newS)

	out := make([]string, 0, len(structured))
	for _, d := range structured {
		if s := formatDiff(d); s != "" {
			out = append(out, s)
		}
	}
	return out
}

func diffSnapshotStructured(oldS, newS models.Snapshot) []Diff {
	var diffs []Diff

	diffs = append(diffs, sliceDiff("Admins", oldS.Admins, newS.Admins)...)
	diffs = append(diffs, sliceDiff("Regular users", oldS.RegularUsers, newS.RegularUsers)...)
	diffs = append(diffs, sliceDiff("Services", oldS.Services, newS.Services)...)
	diffs = append(diffs, sliceDiff("Packages", oldS.Packages, newS.Packages)...)
	diffs = append(diffs, sliceDiff("Open ports", oldS.OpenPorts, newS.OpenPorts)...)

	if oldS.FirewallState != newS.FirewallState {
		diffs = append(diffs, Diff{
			Category: "Firewall state",
			Added:    []string{newS.FirewallState},
			Removed:  []string{oldS.FirewallState},
		})
	}

	return diffs
}

func sliceDiff(name string, oldSlice, newSlice []string) []Diff {
	oldSlice = normalize(oldSlice)
	newSlice = normalize(newSlice)

	added, removed := diffSlice(oldSlice, newSlice)

	if len(added) == 0 && len(removed) == 0 {
		return nil
	}

	sort.Strings(added)
	sort.Strings(removed)

	return []Diff{{
		Category: name,
		Added:    added,
		Removed:  removed,
	}}
}

func normalize(in []string) []string {
	set := map[string]struct{}{}

	for _, v := range in {
		v = strings.ToLower(strings.TrimSpace(v))
		if v != "" {
			set[v] = struct{}{}
		}
	}

	out := make([]string, 0, len(set))
	for v := range set {
		out = append(out, v)
	}

	sort.Strings(out)
	return out
}

func diffSlice(oldSlice, newSlice []string) (added, removed []string) {
	oldMap := map[string]struct{}{}
	newMap := map[string]struct{}{}

	for _, v := range oldSlice {
		oldMap[v] = struct{}{}
	}
	for _, v := range newSlice {
		newMap[v] = struct{}{}
	}

	for v := range newMap {
		if _, ok := oldMap[v]; !ok {
			added = append(added, v)
		}
	}

	for v := range oldMap {
		if _, ok := newMap[v]; !ok {
			removed = append(removed, v)
		}
	}

	return
}

func formatDiff(d Diff) string {
	if len(d.Added) == 0 && len(d.Removed) == 0 {
		return ""
	}

	var b strings.Builder
	b.WriteString(d.Category)
	b.WriteString(":\n")

	for _, a := range d.Added {
		b.WriteString(fmt.Sprintf("  + %s\n", a))
	}
	for _, r := range d.Removed {
		b.WriteString(fmt.Sprintf("  - %s\n", r))
	}

	return b.String()
}