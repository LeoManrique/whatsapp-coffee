package whatsapp

import "strings"

// Status is what the window text says about the session.
type Status int

const (
	// Unknown matched no marker: the window is not a screen we know.
	Unknown Status = iota
	// Connected shows the chat list with no connection warning.
	Connected
	// Stuck shows the chat list with a connection warning in the Chats header.
	Stuck
	// LoggedOut shows the screen to link the device again.
	LoggedOut
)

// stuckMarkers are the subtitles the Chats header shows without a connection:
// "Connecting..." when WhatsApp's servers are unreachable, "Waiting for
// network" when the Mac has no network at all. "Connecting" is matched without
// the dots so that both the ASCII dots and the single ellipsis character work.
var stuckMarkers = []string{"Connecting", "Waiting for network"}

// loggedOutMarkers are texts of the link screen. Best guess from WhatsApp Web,
// not yet confirmed on this app version.
var loggedOutMarkers = []string{
	"Link a Device",
	"Linked Devices",
	"Log in with phone number",
	"QR code",
}

// connectedMarker is the sidebar entry every linked session shows.
const connectedMarker = "Chats"

// statusNames are the words used in logs and in the status command.
var statusNames = map[Status]string{
	Unknown:   "unknown",
	Connected: "connected",
	Stuck:     "stuck",
	LoggedOut: "logged out",
}

// String returns the status name used in logs.
func (s Status) String() string {
	return statusNames[s]
}

// containsAny reports whether text contains at least one of markers.
func containsAny(text string, markers []string) bool {
	for _, marker := range markers {
		if strings.Contains(text, marker) {
			return true
		}
	}
	return false
}

// Classify tells which screen the window text is. The order matters: a stuck
// window still shows the chat list, so its markers are checked before the
// connected one.
func Classify(text string) Status {
	if containsAny(text, loggedOutMarkers) {
		return LoggedOut
	}

	if containsAny(text, stuckMarkers) {
		return Stuck
	}

	if strings.Contains(text, connectedMarker) {
		return Connected
	}

	return Unknown
}
