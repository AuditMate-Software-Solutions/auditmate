package audit

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

const Version = "1.0.0"

var RegisterExtraFlags func(*flag.FlagSet)
var RunExtraFeatures func(*RunResult)
var RegisterHelpLines func() []string

func RunCLI() int {
	_, code := RunCLIWithContext(context.Background())
	return code
}

func RunCLIWithContext(ctx context.Context) (*RunResult, int) {
	fs := flag.NewFlagSet("auditmate", flag.ContinueOnError)

	outDir := fs.String("out", "auditmate-output", "output directory")
	reset := fs.Bool("reset", false, "replace stored snapshot with current snapshot")
	quiet := fs.Bool("quiet", false, "suppress output")
	jsonOut := fs.Bool("json", false, "json output only")
	help := fs.Bool("help", false, "show help")
	version := fs.Bool("version", false, "show version")

	if RegisterExtraFlags != nil {
		RegisterExtraFlags(fs)
	}

	if err := fs.Parse(os.Args[1:]); err != nil {
		return nil, 2
	}

	if *help {
		printHelp()
		return nil, 0
	}

	if *version {
		fmt.Printf("AuditMate v%s\n", Version)
		return nil, 0
	}

	select {
	case <-ctx.Done():
		fmt.Fprintln(os.Stderr, "cancelled")
		return nil, 130
	default:
	}

	res, err := Run(*outDir, *reset)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		return nil, 20
	}

	if RunExtraFeatures != nil {
		RunExtraFeatures(res)
	}

	if *jsonOut {
		b, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			fmt.Fprintln(os.Stderr, "json error:", err)
			return res, 21
		}
		fmt.Println(string(b))
		return res, exitCodeFromResult(res)
	}

	if !*quiet {
		fmt.Println("AuditMate")
		fmt.Println("-----------------------------------")
		fmt.Println("Hostname :", res.Hostname)
		fmt.Println("OS       :", res.OS)
		fmt.Println("Uptime   :", res.Uptime)
		fmt.Println()
		fmt.Println("Users    :", res.Users)
		fmt.Println("Services :", res.Services)
		fmt.Println("Packages :", res.Packages)
		fmt.Println("Ports    :", res.Ports)
		fmt.Println()
		fmt.Println("Status   :", res.Status)

		if len(res.Diffs) > 0 {
			fmt.Println()
			fmt.Println("Changes:")
			for _, d := range res.Diffs {
				fmt.Println(" +", d)
			}
		}

		fmt.Println("-----------------------------------")
		fmt.Println("Warnings :", len(res.Warnings))
		fmt.Println("Errors   :", res.Errors)
		fmt.Println("Reset    :", res.Reset)
		fmt.Println("Report   :", res.ReportPath)
		fmt.Println("Duration :", res.Duration)
	}

	return res, exitCodeFromResult(res)
}

func exitCodeFromResult(res *RunResult) int {
	switch {
	case res.Status == "TAMPER DETECTED":
		return 50 // 🔥 critical security failure
	case res.Errors && len(res.Diffs) > 0:
		return 11
	case res.Errors:
		return 1
	case len(res.Diffs) > 0:
		return 10
	default:
		return 0
	}
}

func printHelp() {
	fmt.Println(`AuditMate - System Audit Tool

USAGE:
  auditmate [options]

OPTIONS:
  --out <dir>     Output directory
  --reset         Replace stored snapshot with current snapshot
  --quiet         Suppress output
  --json          JSON output only
  --help          Show help
  --version       Show version
`)

	if RegisterHelpLines != nil {
		for _, l := range RegisterHelpLines() {
			fmt.Println(" ", l)
		}
	}
}