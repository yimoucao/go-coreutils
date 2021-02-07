# go-coreutils
Reinventing GNU coreutils using Go

## Requirements
- Implement GNU coreutils
- ONLY use Golang standard library (std, golang.org/x)
- No cheating, for example, use exec.Command() directly

## Priority
- core functionality
- output format comformant
- input flags comformant, bsd/posix compliant
- wildcards, globs (incoming go1.16 [io/fs](https://tip.golang.org/doc/go1.16#fs))

## Difficulty

|module | difficulty | comments |
|-------|------------|----------|
|uptime | ** | macOS, windows syscall API, implement utmpx |
|b2sum| ** | crypto hashing|
