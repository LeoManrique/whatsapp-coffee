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

## Slice 3: schedule

Comes before 2b so the offline check can run while nobody is connected to the Mac.

- **3a (done):** `schedule` writes the plist for the signed binary and loads it. The tick it runs at
  load time worked with no permission prompt. Proves: `just schedule`, then `launchctl print` exits 0
  and the log gets a new tick line.
- **3b (done):** `kill` unloads the job and deletes the plist, `status` prints whether it is loaded
  and the last log line. Proves: `just status`, `just kill`, `just status`.

## Slice 2b: stuck detection and recovery

The Chats header shows "Connecting..." when WhatsApp's servers are unreachable and "Waiting for
network" when the Mac has no network at all (seen on a plane). Both count as stuck. A quiet disconnect
with no header text is still an open question.

- **2b-i (done):** `whatsapp.Classify` turns the window text into connected, stuck, logged out or
  unknown, and `wacoffee window` prints the verdict. A `state` package keeps the strike count in
  `state.json`, and `wacoffee status` prints it. Proves: `just window` says `status: connected`, and
  `status: stuck` with the network off; `just status` prints `strikes: 0`.
- **2b-ii (next):** the tick reads the window after focusing, counts a strike when stuck or unknown, quits
  and relaunches WhatsApp after 2, resets the count when connected, and logs and stops on the link
  screen. Proves: WhatsApp blocked in the firewall, two ticks, the log shows the strikes and the
  relaunch, `status` shows the count going back to 0.

The link screen markers are a guess until the screen is read once: log WhatsApp Desktop out on a Mac
that is not the always-on one, run `just window`, and relink from the phone.
