# [CLI Task Tracker](github.com/gomahmoud/roadmap.sh-github-user-activity)

A small command-line application to fetch the latest GitHub user activity. Activities are fetched in a JSON format, and then be parsed, formatted and presented in the CLI.

## Requirements

- Go 1.27.1 or later

## Run the application

From the project directory, run commands with `go run .`:

```sh
go run . gomahmoud
```

To build an executable:

```sh
go build -o github-activity .
```

Then run it using the same commands, replacing `go run .` with `./github-activity` (on Windows, use `github-activity.exe`).

## Errors

Invalid commands, incorrect argument counts and Not-Found errors are reported on the command line.
