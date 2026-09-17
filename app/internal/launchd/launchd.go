// Package launchd schedules wacoffee as a launchd user agent.
package launchd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Label is the job's name in launchd and the name of its plist file.
const Label = "com.leomanrique.wacoffee"

// plistTemplate runs "<binary> tick" every 300 seconds and once when loaded,
// with stderr going to the log. It takes Label, the binary path and the log path.
const plistTemplate = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>Label</key>
	<string>%s</string>
	<key>ProgramArguments</key>
	<array>
		<string>%s</string>
		<string>tick</string>
	</array>
	<key>StartInterval</key>
	<integer>300</integer>
	<key>RunAtLoad</key>
	<true/>
	<key>StandardErrorPath</key>
	<string>%s</string>
</dict>
</plist>
`

// domain is the logged-in user's GUI session in launchd, such as "gui/501".
func domain() string {
	return fmt.Sprintf("gui/%d", os.Getuid())
}

// service is the job's full name in launchd.
func service() string {
	return domain() + "/" + Label
}

// plistPath returns where launchd looks for the job's plist.
func plistPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "Library", "LaunchAgents", Label+".plist"), nil
}

// launchctl runs launchctl with args, putting its output into the error.
func launchctl(args ...string) error {
	cmd := exec.Command("launchctl", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf(
			"launchctl %s: %s: %w",
			strings.Join(args, " "),
			strings.TrimSpace(string(out)),
			err,
		)
	}
	return nil
}

// notLoaded reports whether err is launchctl saying the job is not loaded.
func notLoaded(err error) bool {
	exit, ok := errors.AsType[*exec.ExitError](err)
	if !ok {
		return false
	}

	return exit.ExitCode() == 3 || exit.ExitCode() == 113

}

// unload removes the job from launchd. A job that is not loaded is not an error.
func unload() error {
	err := launchctl("bootout", service())
	if err != nil && !notLoaded(err) {
		return err
	}
	return nil
}

// Install writes the plist for binary and loads it, replacing any earlier copy.
func Install(binary, logPath string) error {
	path, err := plistPath()

	if err != nil {
		return err
	}

	content := fmt.Sprintf(plistTemplate, Label, binary, logPath)

	err = os.WriteFile(path, []byte(content), 0o644)
	if err != nil {
		return err
	}

	err = unload()
	if err != nil {
		return err
	}

	return launchctl("bootstrap", domain(), path)
}

// Loaded reports whether launchd still has the job.
func Loaded() (bool, error) {
	err := launchctl("print", service())
	if err == nil {
		return true, nil
	}

	if notLoaded(err) {
		return false, nil
	}

	return false, err
}

// Kill unloads the job and deletes its plist. WhatsApp is left running.
func Kill() error {
	path, err := plistPath()
	if err != nil {
		return err
	}

	err = unload()
	if err != nil {
		return err
	}

	err = os.Remove(path)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	return nil
}
