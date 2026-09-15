// Package whatsapp drives the official WhatsApp Desktop app through osascript.
// It never opens a chat and never sends a message.
package whatsapp

import (
	"fmt"
	"os/exec"
	"strings"
)

// appName is the name macOS knows the app by, in both the process list and AppleScript.
const appName = "WhatsApp"

// osascript runs one line of AppleScript and returns its output without surrounding whitespace.
func osascript(script string) (string, error) {
	cmd := exec.Command("osascript", "-e", script)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s: %w", out, err)
	}
	return strings.TrimSpace(string(out)), nil
}

// Running reports whether WhatsApp has a process, according to System Events.
func Running() (bool, error) {
	script := fmt.Sprintf(`tell application "System Events" to (name of processes) contains %q`, appName)
	out, err := osascript(script)

	if err != nil {
		return false, err
	}

	return out == "true", nil
}

// Launch starts WhatsApp without bringing it to the front.
func Launch() error {
	return tell("launch")
}

// Focus brings the WhatsApp window to the front, starting the app if it is not running.
func Focus() error {
	return tell("activate")
}

// tell sends one verb to the WhatsApp application object, such as "launch" or "activate".
func tell(verb string) error {
	line := fmt.Sprintf(`tell application %q to %s`, appName, verb)

	_, err := osascript(line)
	return err
}
