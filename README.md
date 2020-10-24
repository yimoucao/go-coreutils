# go-coreutils
Reinventing GNU coreutils using Go

## Requirements
- Implement GNU coreutils
- ONLY use Golang standard library
- No cheating, for example, use exec.Command() directly

## Priority
- core functionality
- output format comformant
- input flags comformant, bsd/posix compliant
- wildcards, blobs

## Difficulty

|module | difficulty | comments |
|-------|------------|----------|
|uptime | ** | macOS, windows syscall API, implement utmpx |
