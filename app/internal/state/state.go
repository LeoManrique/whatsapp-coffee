// Package state keeps what one tick leaves for the next, in a small JSON file.
package state

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

// State is what survives between ticks.
type State struct {
	// Strikes counts the consecutive ticks that found WhatsApp stuck.
	Strikes int `json:"strikes"`
}

// Path returns the absolute path of the state file,
// ~/Library/Application Support/wacoffee/state.json.
func Path() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, "Library", "Application Support", "wacoffee", "state.json"), nil
}

// Load reads the state file. A missing file is the zero State.
func Load() (State, error) {
	var s State

	path, err := Path()
	if err != nil {
		return s, err
	}

	data, err := os.ReadFile(path)

	if errors.Is(err, os.ErrNotExist) {
		return s, nil
	}

	if err != nil {
		return s, err
	}

	err = json.Unmarshal(data, &s)

	return s, err
}

// Save writes s to the state file, creating its folder on the first run.
func Save(s State) error {
	path, err := Path()
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(path), 0o755)

	if err != nil {
		return err
	}

	data, err := json.Marshal(s)

	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, 0o644)

	return err
}
