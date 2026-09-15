# Roadmap

- [x] Analysis (`docs/initial-implementation/analysis/linked-devices.md`)
- [ ] Experiment: phone in airplane mode, second account checks presence and ticks
- [x] Slice 1: `wacoffee tick` ensures WhatsApp is running and focused, logs the result
- [x] Slice 2a: `wacoffee window` prints the window content
- [ ] Slice 3: `schedule` / `kill` / `status` through launchd, Accessibility granted to the binary
- [ ] Slice 2b: offline window check, classify the window state, strike counting, relaunch when stuck
- [ ] Run supervised
- [ ] Run unsupervised
- [ ] V2: run automatically on login or reboot
- [ ] V2: advanced tick schedule (e.g. better distributed online times, no ticks at night)
