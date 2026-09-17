# whatsapp-coffee

Keeps WhatsApp Desktop on an always-on Mac, looking online while the phone is offline or fully off.

Each run brings WhatsApp to the front, checks it is connected, and relaunches it if it is stuck.

## Usage

```sh
just tick       # one manual run
just window     # print the content of the WhatsApp window
just schedule   # schedule it with launchd
just kill       # kill switch: stops and removes the schedule
just status     # is it scheduled, and the last log line
```

Log: `~/Library/Logs/wacoffee.log`.

See [DESIGN.md](DESIGN.md), [TECHNICAL.md](TECHNICAL.md), [ROADMAP.md](ROADMAP.md).
