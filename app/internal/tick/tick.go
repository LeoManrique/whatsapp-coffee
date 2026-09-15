// Package tick runs the sequence to keep WhatsApp Desktop running.
package tick

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"wacoffee/internal/whatsapp"
)

// Run performs one tick: ensure WhatsApp is running, bring it to the front, log the outcome.
func Run() error {
	logger, closeLog, err := openLog()
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

// openLog opens ~/Library/Logs/wacoffee.log for appending and returns a logger
// that writes both there and to stdout, plus a function that closes the file.
func openLog() (*log.Logger, func(), error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, nil, err
	}

	path := filepath.Join(home, "Library", "Logs", "wacoffee.log")

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return nil, nil, err
	}

	logger := log.New(io.MultiWriter(os.Stdout, f), "", log.LstdFlags)
	return logger, func() { f.Close() }, nil
}
