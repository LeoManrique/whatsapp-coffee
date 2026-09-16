# Run `just` to list recipes.

app := "app"
bin := "app/bin/wacoffee"

default:
    @just --list

# Build the binary into app/bin, signed so macOS permissions survive rebuilds
build:
    cd {{app}} && go build -o bin/wacoffee ./cmd/wacoffee
    codesign -f -s "Apple Development" --identifier com.leomanrique.wacoffee {{bin}}

# One manual tick
tick: build
    {{bin}} tick

# Print the content of the WhatsApp window (takes a few seconds)
window: build
    {{bin}} window

# Schedule the tick with launchd
schedule: build
    {{bin}} schedule

# Kill switch: unload and remove the launchd job
kill: build
    {{bin}} kill

# Is the job loaded, and the last log line
status: build
    {{bin}} status

# Follow the log
log:
    tail -f ~/Library/Logs/wacoffee.log

vet:
    cd {{app}} && go vet ./...

test:
    cd {{app}} && go test ./...
