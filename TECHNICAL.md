# Technical

## Stack

- **Go 1.27**, standard library only. One binary: `wacoffee`.
- **osascript** (AppleScript, System Events) for launching, focusing, and reading the WhatsApp window. Needs Accessibility permission for the terminal/binary.
- **launchd** user agent (`~/Library/LaunchAgents/com.leomanrique.wacoffee.plist`, `StartInterval` 300) runs `wacoffee tick`.
- Target: macOS 27, WhatsApp Desktop 26.x.

## Layout

```
justfile                 build, tick, schedule, kill, status, log, vet, test
app/                     Go project
  cmd/wacoffee/          main: subcommands tick, schedule, kill, status
  internal/whatsapp/     running, launch, focus, quit, read window state (osascript wrappers)
  internal/launchd/      write and load / unload the plist
  internal/tick/         the tick sequence from DESIGN.md, strike counting
  bin/                   build output, ignored by git
docs/analysis/           research and experiment results
docs/plans/              slice plan
```

## State and logs

- Log: `~/Library/Logs/wacoffee.log`.
- Strike count: `~/Library/Application Support/wacoffee/state.json`.

## Mac prerequisites

- Never sleeps or locks (either from Settings or third party apps like Amphetamine).
- Display sleep: never.
- macOS automatic updates off.
