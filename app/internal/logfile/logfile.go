// Package logfile owns wacoffee's log: where it lives and how it is opened.
package logfile

import (
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Path returns the absolute path of the log, ~/Library/Logs/wacoffee.log.
func Path() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, "Library", "Logs", "wacoffee.log"), nil
}

// Open opens the log for appending and returns a logger that writes both there
// and to stdout, plus a function that closes the file.
func Open() (*log.Logger, func(), error) {
	path, err := Path()
	if err != nil {
		return nil, nil, err
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return nil, nil, err
	}

	logger := log.New(io.MultiWriter(os.Stdout, f), "", log.LstdFlags)
	return logger, func() { f.Close() }, nil
}

// LastLine returns the last line of the log, or "" when there is nothing to show.
func LastLine() (string, error) {
	path, err := Path()
	if err != nil {
		return "", err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return "", err
		}
		return "", nil
	}

	text := strings.TrimRight(string(data), "\n")

	cut := strings.LastIndex(text, "\n")
	if cut != -1 {
		return text[cut+1:], nil
	}

	return text, nil
}
