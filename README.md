# go-coreutils
Reinventing GNU coreutils using Go

## Requirements
- Implement GNU coreutils
- ONLY use Golang standard library (std, golang.org/x)
- No cheating, for example, use exec.Command() directly

## Priority
1. core functionality
2. output format comformant
3. wildcards, globs (go1.16 [io/fs](https://tip.golang.org/doc/go1.16#fs))
4. ~~input flags comformant, bsd/posix compliant~~

## Difficulty

|module | difficulty | comments |
|-------|------------|----------|
|uptime | ** | macOS, windows syscall API, implement utmpx |
|b2sum| ** | crypto hashing|


### References
[GNU coreutils github mirror](//https://github.com/coreutils/coreutils)