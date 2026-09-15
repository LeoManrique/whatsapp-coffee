# Run `just` to list recipes.

app := "app"
bin := "app/bin/wacoffee"

default:
    @just --list

# Build the binary into app/bin
build:
    cd {{app}} && go build -o bin/wacoffee ./cmd/wacoffee

# One manual tick
tick: build
    {{bin}} tick

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
