# license-checker

This is a simple utility that recursively checks all files in a directory to ensure they contain the
Apache license header at the beginning.

By default, it checks the following prefixes:

- Dockerfile (e.g. Dockerfile-v2)
- Makefile

By default, it checks the following extensions:

- .go
- .c
- .cpp
- .py
- .sh
- .rb
- .yaml

By default, it skips the following directories:

- .git
- vendor
- generated

## Building

You can just `go get github.com/heptio/license-checker` to get the binary installed locally.

## Running

Running is easy - just do `license-checker` to check the current directory. You can also pass
`-check` with a comma-separated list of prefixes and extensions to check (e.g. `-check
Dockerfile,.go`). You can also pass `-skip` with a comman-separated list of directory names to skip.



