# Roadmap

- [x] Analysis (`docs/initial-implementation/analysis/linked-devices.md`)
- [ ] Experiment: phone in airplane mode, second account checks presence and ticks
- [x] Slice 1: `wacoffee tick` ensures WhatsApp is running and focused, logs the result
- [x] Slice 2a: `wacoffee window` prints the window content
- [x] Slice 3a: `schedule` loads the launchd job, permissions granted to the signed binary
- [x] Slice 3b: `kill` / `status`
- [x] Slice 2b-i: classify the window as connected, stuck, logged out or unknown; strike count file
- [x] Slice 2b-ii: the tick checks the window, counts strikes, relaunches when stuck
- [ ] Read the link screen once on a spare Mac to confirm the logged-out markers
- [ ] Find out what hides the windows on the always-on Mac (only the desktop shows), and whether presence survives it
- [x] Run supervised
- [ ] Run unsupervised
- [ ] V2: force quit when WhatsApp does not answer the quit request
- [ ] V2: run automatically on login or reboot
- [ ] V2: advanced tick schedule (e.g. better distributed online times, no ticks at night)
