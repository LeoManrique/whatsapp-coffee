# Functional design of the core

## Goal

While the phone is offline, messages must keep getting received and delivered, contacts must keep seeing the account as online or recently seen.

## Tick every n minutes

1. **Ensure running.** If WhatsApp is not running, warn and launch it.
2. **Refocus.** Bring the WhatsApp window to the front. Presence ("online") is only broadcast while the window is focused.
3. **Check connection.** Read the window through Accessibility. If the Chats header says "Connecting..." or "Waiting for network", or the screen is one the tick does not recognize, count a strike. A connected window puts the count back to 0. The tick that had to launch WhatsApp skips the check, because a window that is still loading would read as stuck. If the window shows the QR/login screen, the tick logs that a relink by hand is needed and stops: a relaunch cannot fix a logged-out session, and every later tick logs the same until someone relinks.
4. **Recover.** After 2 consecutive strikes, quit and relaunch WhatsApp, bring it back to the front, and start counting again. During a long outage that is one relaunch every 2 ticks, so a reconnect in progress is not interrupted by the next tick.
5. **Log.** One line per tick with the outcome.

The tick is the only process. There is no separate watchdog.

## Kill switch and status

`wacoffee kill` unloads the launchd job and deletes its plist, so it does not come back at the next login. Nothing else keeps running. WhatsApp itself stays open. Running it when the job is already gone is not an error.

`wacoffee status` prints whether launchd still has the job, the current strike count, and the last line of the log. The log is kept by `kill`, so the last tick stays visible afterwards.

## Rules

- Never open a chat or send a message: blue ticks or auto-messages would give the automation away or break the terms of service.
- Only the official app is used, through AppleScript. No unofficial protocol libraries.

## Decisions

- The window stays focused and the display awake, whether or not presence needs them.

## Open questions

- A lost connection can look normal: no banner, sent messages stay on the clock icon, and "Connecting..." only shows after a relaunch. How does the tick notice it?
- What does the window show when WhatsApp's servers are down? A screen the tick does not recognize counts as a strike, so the worst case is a relaunch.
- What does the link screen say in this app version? The words the tick looks for come from WhatsApp Web.
