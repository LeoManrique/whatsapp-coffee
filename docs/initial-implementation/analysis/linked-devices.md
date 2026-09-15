# WhatsApp linked devices: what we know (2026-09-15)

## Confirmed (official docs)

- Linked devices keep their own connection. Messages are delivered to the Mac while the phone is off.
- Linked devices log out if the phone stays unused for 14 days, or after 30 days of inactivity on the linked device itself.
- Terms of service forbid auto-messaging on the consumer app. We do not send anything.

## Widely reported, not official

- A linked desktop session drives "online" and "last seen", but only while its window is focused. Minimized or in the background, contacts see "last seen".
- WhatsApp Desktop on macOS sometimes gets stuck in "Reconnecting" after long idle periods and needs a relaunch.
- App updates and reboots normally keep the session, but there is no official statement.

## Experiment (needs a second account)

1. Phone in airplane mode. WhatsApp focused on the Mac. Second account checks "online" and "last seen" after 30 min.
2. Second account sends a message: double grey ticks?
3. Quit and relaunch WhatsApp. Reboot the Mac. Still linked?

Record the results below.

## Results

- (pending)

## Sources

- https://faq.whatsapp.com/378279804439436
- https://faq.whatsapp.com/451924530376167
- https://www.whatsapp.com/legal/terms-of-service
- https://engineering.fb.com/2021/07/14/security/whatsapp-multi-device/
- https://wabetainfo.com/whatsapp-is-working-on-hiding-the-online-status-on-desktop-beta/
- https://discussions.apple.com/thread/253870026
