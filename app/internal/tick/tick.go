// Package tick runs the sequence to keep WhatsApp Desktop running.
package tick

import (
	"wacoffee/internal/logfile"
	"wacoffee/internal/whatsapp"
)

// Run performs one tick: ensure WhatsApp is running, bring it to the front, log the outcome.
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

	logger.Println("tick ok: running and focused")
	return nil
}
