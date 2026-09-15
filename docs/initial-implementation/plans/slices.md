# Slice plan

Each slice builds and has a command that proves it. Details are added when a slice starts.

## Slice 1: running and focused (done)

`wacoffee tick` launches WhatsApp if needed, brings it to the front, appends one log line.
Proves: `just tick` then `tail -1 ~/Library/Logs/wacoffee.log`.

## Slice 2a: read the window (done)

`whatsapp.WindowText` collects every text the WhatsApp window shows through System Events, and
`wacoffee window` prints it. With the servers unreachable it showed "Connecting..." after a relaunch,
but no banner before it.
Proves: `just window`.

## Slice 3: schedule (next)

`schedule` writes and loads the launchd plist, `kill` unloads and deletes it, `status` prints whether
it is loaded and the last log line. The binary run by launchd must be granted Accessibility itself.
Comes before 2b so the offline check can run while nobody is connected to the Mac.
Proves: `just schedule` then `just status`.

## Slice 2b: stuck detection and recovery

Take the Mac fully offline with the tick scheduled and see what the window shows. Then classify the
window text into connected, stuck, or logged out, keep a strike count in `state.json`, quit and relaunch
after 2 strikes, log and stop on the QR screen. Also decide how a quiet disconnect with no banner gets noticed.
Proves: Mac offline, two ticks, the log shows the strikes and the relaunch.
