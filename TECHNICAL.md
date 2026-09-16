# Technical

## Stack

- **Go 1.27**, standard library only. One binary: `wacoffee`.
- **osascript** (AppleScript, System Events) for launching, focusing, and reading the WhatsApp window.
- **launchd** user agent (`~/Library/LaunchAgents/com.leomanrique.wacoffee.plist`) runs `wacoffee tick` every 300 seconds and once when loaded, with stderr going to the log.
- Target: macOS 27, WhatsApp Desktop 26.x (native Mac Catalyst app).

## Layout

```
justfile                         build, tick, window, schedule, kill, status, log, vet, test
app/                             Go project
  cmd/wacoffee/                  main: subcommands tick, window, schedule, kill, status
  internal/whatsapp/             running, launch, focus, quit, window text (osascript wrappers)
  internal/launchd/              write and load / unload the plist
  internal/logfile/              log path and logger
  internal/tick/                 the tick sequence from DESIGN.md, strike counting
  bin/                           build output, ignored by git
docs/initial-implementation/     analysis, experiment results, slice plan
```

## Reading the window

WhatsApp is a Mac Catalyst app: its text sits in the `description` attribute of the accessibility
elements, not in `value`, and it is wrapped in invisible direction marks (U+200E, U+200F,
U+2066 to U+2069) that are stripped before matching. One AppleScript handler walks the window tree to
depth 12 in bulk calls per container (about 8 seconds), which covers the title bar, sidebar and
banners and stops before the chat rows, so message content is never read. The window does not need to
be frontmost for the read.

## Accessibility permission

macOS grants permissions to the first non-system program in the chain: the terminal app for manual
runs, the `wacoffee` binary itself under launchd. The binary needs Automation for System Events and
WhatsApp, which macOS asks about on screen the first time it is missing, and Accessibility for reading
the window, added by hand in System Settings > Privacy & Security > Accessibility. Without
Accessibility, reads fail with `osascript is not allowed assistive access. (-25211)`. The first
scheduled tick launched and focused WhatsApp with no prompt.

Grants are tied to the code signature and the path. `just build` signs the binary with the Apple
Development identity and the identifier `com.leomanrique.wacoffee`, so grants survive rebuilds as
long as it stays at `app/bin/wacoffee`.

## State and logs

- Log: `~/Library/Logs/wacoffee.log`.
- Strike count: `~/Library/Application Support/wacoffee/state.json`.

## Mac prerequisites

- Never sleeps or locks (either from Settings or third party apps like Amphetamine).
- Display sleep: never.
- macOS automatic updates off.
