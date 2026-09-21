// Package tick runs the sequence to keep WhatsApp Desktop running and connected.
package tick

import (
	"log"

	"wacoffee/internal/logfile"
	"wacoffee/internal/state"
	"wacoffee/internal/whatsapp"
)

// maxStrikes is how many ticks in a row may find WhatsApp stuck before it is relaunched.
const maxStrikes = 2

// snippetLength is how much of an unrecognized window goes to the log, in characters.
const snippetLength = 300

// Run performs one tick: ensure WhatsApp is running, bring it to the front,
// check its connection, log the outcome.
func Run() error {
	logger, closeLog, err := logfile.Open()
	if err != nil {
		return err
	}
	defer closeLog()
	running, err := whatsapp.Running()

	if err != nil {
		return err
	}

	if !running {
		logger.Println("WhatsApp not running, launching")
		err := whatsapp.Launch()
		if err != nil {
			return err
		}
	}

	if err = whatsapp.Focus(); err != nil {
		logger.Println("Error while focusing WhatsApp")
		return err
	}

	// A window that is still loading would read as stuck, so the check waits for the next tick.
	if !running {
		logger.Println("tick ok: launched and focused, connection check at the next tick")
		return nil
	}

	return checkConnection(logger)
}

// checkConnection reads the window, keeps the strike count, and relaunches
// WhatsApp when it has been stuck for maxStrikes ticks in a row.
func checkConnection(logger *log.Logger) error {
	text, err := whatsapp.WindowText()
	if err != nil {
		logger.Println("Error while reading the WhatsApp window")
		return err
	}
	status := whatsapp.Classify(text)

	// A state file that cannot be read must not stop the check: start from zero and overwrite it.
	s, err := state.Load()
	if err != nil {
		logger.Println("state file unreadable, counting from 0:", err)
		s = state.State{}
	}

	switch status {
	case whatsapp.Connected:
		s.Strikes = 0

		logger.Println("tick ok: running, focused and connected")

	case whatsapp.LoggedOut:
		// A relaunch cannot fix a logged-out session, so the strike count stays as it is.
		logger.Println("logged out: link this Mac again from the phone")
		return nil

	default:
		// Stuck, or a screen nobody has seen yet.
		s.Strikes++

		logger.Printf("strike %d of %d: window is %s", s.Strikes, maxStrikes, status)
		if status == whatsapp.Unknown {
			logger.Println("unrecognized window:", snippet(text))
		}

		if s.Strikes < maxStrikes {
			break
		}

		if err := relaunch(logger); err != nil {
			return err
		}

		s.Strikes = 0
	}

	return state.Save(s)
}

// relaunch quits and restarts WhatsApp and brings it back to the front.
func relaunch(logger *log.Logger) error {
	logger.Println("relaunching WhatsApp")
	if err := whatsapp.Relaunch(); err != nil {
		logger.Println("Error while relaunching WhatsApp")
		return err
	}

	return whatsapp.Focus()
}

// snippet shortens the window text to what fits on one log line.
func snippet(text string) string {
	runes := []rune(text)
	if len(runes) <= snippetLength {
		return text
	}

	return string(runes[:snippetLength]) + "..."
}
