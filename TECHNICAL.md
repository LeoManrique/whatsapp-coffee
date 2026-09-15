# Technical

## Stack

- **Go 1.27**, standard library only. One binary: `wacoffee`.
- **osascript** (AppleScript, System Events) for launching, focusing, and reading the WhatsApp window.
- **launchd** user agent (`~/Library/LaunchAgents/com.leomanrique.wacoffee.plist`, `StartInterval` 300) runs `wacoffee tick`.
- Target: macOS 27, WhatsApp Desktop 26.x (native Mac Catalyst app).

## Layout

```
justfile                         build, tick, window, schedule, kill, status, log, vet, test
app/                             Go project
  cmd/wacoffee/                  main: subcommands tick, window, schedule, kill, status
  internal/whatsapp/             running, launch, focus, quit, window text (osascript wrappers)
  internal/launchd/              write and load / unload the plist
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

System Events refuses UI reads with `osascript is not allowed assistive access. (-25211)` when the
caller is not granted. macOS grants to the first non-system program in the chain: in a terminal that is
the terminal app, under launchd it is the `wacoffee` binary itself, so the binary must be added in
System Settings > Privacy & Security > Accessibility. The grant is tied to the code signature, and an
unsigned build gets a new identity on every rebuild, so grant it again after the final build, or sign
the binary with a stable identity.

## State and logs

- Log: `~/Library/Logs/wacoffee.log`.
- Strike count: `~/Library/Application Support/wacoffee/state.json`.

## Mac prerequisites

- Never sleeps or locks (either from Settings or third party apps like Amphetamine).
- Display sleep: never.
- macOS automatic updates off.
