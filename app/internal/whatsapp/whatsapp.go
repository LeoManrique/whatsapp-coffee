// Package whatsapp drives the official WhatsApp Desktop app through osascript.
// It never opens a chat and never sends a message.
package whatsapp

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// appName is the name macOS knows the app by, in both the process list and AppleScript.
const appName = "WhatsApp"

// quitTimeout is how long Relaunch waits for the WhatsApp process to go away,
// and quitPoll is the pause between two looks at the process list.
const (
	quitTimeout = 30 * time.Second
	quitPoll    = time.Second
)

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

// Quit asks WhatsApp to close, the same as Cmd+Q.
func Quit() error {
	return tell("quit")
}

// Relaunch quits WhatsApp, waits until its process is gone, and launches it again.
func Relaunch() error {
	if err := Quit(); err != nil {
		return err
	}

	deadline := time.Now().Add(quitTimeout)
	for {
		running, err := Running()
		if err != nil {
			return err
		}

		if !running {
			return Launch()
		}

		if time.Now().After(deadline) {
			return fmt.Errorf("WhatsApp did not quit within %s", quitTimeout)
		}

		time.Sleep(quitPoll)
	}
}

// tell sends one verb to the WhatsApp application object, such as "launch" or "activate".
func tell(verb string) error {
	line := fmt.Sprintf(`tell application %q to %s`, appName, verb)

	_, err := osascript(line)
	return err
}

// windowTextScript is the AppleScript that collects the text of the WhatsApp
// window. It is a format string: the %q receives appName.
const windowTextScript = `on grab(el, lvl)
	set acc to {}
	if lvl > 12 then return acc
	tell application "System Events"
		try
			set ds to description of UI elements of el
		on error
			set ds to {}
		end try
		try
			set vs to value of UI elements of el
		on error
			set vs to {}
		end try
		try
			set kids to UI elements of el
		on error
			set kids to {}
		end try
	end tell
	repeat with x in (ds & vs)
		if (contents of x) is not missing value then copy ((contents of x) as text) to end of acc
	end repeat
	repeat with k in kids
		set acc to acc & grab(contents of k, lvl + 1)
	end repeat
	return acc
end grab
tell application "System Events" to tell process %q to set w to window 1
set r to grab(w, 0)
set o to ""
repeat with x in r
	set o to o & (contents of x) & " | "
end repeat
return o`

// marks removes the invisible Unicode direction characters WhatsApp wraps its
// text in (left-to-right and right-to-left marks, and the isolate marks), so
// that substring matching on the window text works.
var marks = strings.NewReplacer(
	"\u200e", "",
	"\u200f", "",
	"\u2066", "",
	"\u2067", "",
	"\u2068", "",
	"\u2069", "",
)

// WindowText returns all the text the WhatsApp window shows, joined by " | ",
// with the direction marks removed. It takes several seconds.
func WindowText() (string, error) {
	script := fmt.Sprintf(windowTextScript, appName)

	out, err := osascript(script)
	if err != nil {
		return "", err
	}

	return marks.Replace(out), nil
}
