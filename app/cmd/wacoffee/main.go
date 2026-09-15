// Command wacoffee keeps WhatsApp Desktop running and focused on an always-on Mac.
//
// Usage:
//
//	wacoffee tick
package main

import (
	"fmt"
	"os"

	"wacoffee/internal/tick"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}
	if err := run(os.Args[1]); err != nil {
		fmt.Fprintln(os.Stderr, "wacoffee:", err)
		os.Exit(1)
	}
}

// run dispatches one subcommand.
func run(cmd string) error {
	switch cmd {
	case "tick":
		return tick.Run()
	default:
		usage()
		return fmt.Errorf("unknown command %q", cmd)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: wacoffee tick")
}
