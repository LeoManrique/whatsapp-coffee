# Slice plan

Each slice builds and has a command that proves it. Details are added when a slice starts.

## Slice 1: running and focused

`wacoffee tick` launches WhatsApp if needed, brings it to the front, appends one log line.
Proves: `just tick` then `tail -1 ~/Library/Logs/wacoffee.log`.

## Slice 2: stuck detection and recovery

Read the window text through System Events, keep a strike count in `state.json`, relaunch after 2 strikes.
Proves: quit WhatsApp's network (go offline or block it in the firewall), run two ticks, see the relaunch in the log.

## Slice 3: schedule

`schedule` writes and loads the launchd plist, `kill` unloads and deletes it, `status` prints whether it is loaded and the last log line.
Proves: `just schedule` then `just status`.
