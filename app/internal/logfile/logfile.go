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

	// TODO 1: Read the whole file at once with os.ReadFile, passing path, and
	//         keep the two results as data and err.

	// TODO 2: If err says the file is not there yet, return "" and no error.
	//         errors.Is(err, os.ErrNotExist) reports that, following the chain
	//         of wrapped errors instead of comparing err directly.

	// TODO 3: If err is anything else, return "" and err.

	// TODO 4: Turn data into a string and cut the trailing newline the logger
	//         leaves, with strings.TrimRight, whose second argument is the set
	//         of characters to cut from the end. Keep the result as text.
	//         Example, with other names: strings.TrimRight(name, "?!")

	// TODO 5: Find where the last line starts: strings.LastIndex(text, "\n")
	//         gives the position of the last newline, or -1 when there is none.
	//         Keep it as cut.

	// TODO 6: If cut is not -1, return everything in text after that newline,
	//         which is text[cut+1:], and no error.

	return text, nil
}
