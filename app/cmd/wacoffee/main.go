// Command wacoffee keeps WhatsApp Desktop running and focused on an always-on Mac.
//
// Usage:
//
//	wacoffee tick
//	wacoffee window
//	wacoffee schedule
package main

import (
	"fmt"
	"os"

	"wacoffee/internal/launchd"
	"wacoffee/internal/logfile"
	"wacoffee/internal/tick"
	"wacoffee/internal/whatsapp"
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
	case "window":
		return checkWindow()
	case "schedule":
		return schedule()
	default:
		usage()
		return fmt.Errorf("unknown command %q", cmd)
	}
}

// checkWindow prints everything the WhatsApp window shows
func checkWindow() error {
	text, err := whatsapp.WindowText()

	if err != nil {
		return err
	}

	fmt.Println(text)
	return nil
}

// schedule has launchd run this binary's tick every 5 minutes.
func schedule() error {
	binary, err := os.Executable()
	if err != nil {
		return err
	}

	logPath, err := logfile.Path()
	if err != nil {
		return err
	}

	err = launchd.Install(binary, logPath)
	if err != nil {
		return err
	}

	fmt.Println("scheduled", launchd.Label, "every 5 minutes")
	return nil
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: wacoffee tick|window|schedule")
}
