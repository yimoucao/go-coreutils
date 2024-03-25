package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"
)

var (
	flags struct {
		preserveStatus bool
		foreground     bool
		killAfter      time.Duration
		signal         string
	}
)

func init() {
	flag.BoolVar(&flags.preserveStatus, "preserve-status", false, "exit with the same status as COMMAND")
	flag.BoolVar(&flags.foreground, "foreground", false, "..")
	flag.DurationVar(&flags.killAfter, "kill-after", 0, "also send a KILL signal")
	flag.StringVar(&flags.signal, "signal", "", "specify the signal")
}

// NOTE: go run ignores interrupt https://github.com/golang/go/issues/40467
// TODO: on windows program cannot be finished properly.

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "try 'timeout --help' for more information")
		return
	}
	duration, err := time.ParseDuration(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, "try 'timeout --help' for more information")
		return
	}
	var sig os.Signal
	if flags.signal != "" {
		var err error
		sig, err = parseSignal(flags.signal)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
			return
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	cmd := exec.CommandContext(ctx, args[1], args[2:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Cancel = func() error {
		return cmd.Process.Signal(sig)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGABRT, syscall.SIGTERM)
	go func() {
		fmt.Println("listen to interrupt")
		s := <-sigCh
		fmt.Println("got signal")
		cancel()
		if runtime.GOOS == "windows" {
			// Signal(): Sending Interrupt on Windows is not implemented.
			cmd.Process.Signal(os.Kill)
		} else {
			cmd.Process.Signal(s)
		}
		os.Exit(1)
	}()

	if err := cmd.Start(); err != nil {
		fmt.Println("err:", err)
		return
	}
	if flags.killAfter > 0 {
		go time.AfterFunc(flags.killAfter, func() {
			cmd.Process.Signal(os.Kill)
		})
	}

	cmd.Wait()

}

func parseSignal(s string) (os.Signal, error) {
	if num, ok := isAllUnsignedDigit(s); ok {
		if num > 64 {
			return nil, fmt.Errorf("invalid signal: %q", s)
		}
		return syscall.Signal(num), nil
	}
	if sig, ok := sigTable[strings.ToUpper(s)]; ok {
		return sig, nil
	}
	return nil, fmt.Errorf("invalid signal: %q", s)
}

func isAllUnsignedDigit(s string) (uint, bool) {
	var num uint
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0, false
		}
		num = num*10 + (uint(r) - '0')
	}
	return num, true
}

var sigTable = map[string]syscall.Signal{
	"HUP":  syscall.SIGHUP,
	"INT":  syscall.SIGINT,
	"QUIT": syscall.SIGQUIT,
	"ILL":  syscall.SIGILL,
	"TRAP": syscall.SIGTRAP,
	"ABRT": syscall.SIGABRT,
	"BUS":  syscall.SIGBUS,
	"FPE":  syscall.SIGFPE,
	"KILL": syscall.SIGKILL,
	"SEGV": syscall.SIGSEGV,
	"PIPE": syscall.SIGPIPE,
	"ALRM": syscall.SIGALRM,
	"TERM": syscall.SIGTERM,
}
