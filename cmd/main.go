package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"auditmate/pkg/audit"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(os.Stderr, "panic:", r)
			os.Exit(2)
		}
	}()

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer cancel()

	_, code := audit.RunCLIWithContext(ctx)

	os.Exit(code)
}