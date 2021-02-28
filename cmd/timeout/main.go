package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"
)

var (
	preserveStatus bool
	foreground     bool
	killAfter      time.Duration
	signal         string
)

func init() {
	flag.BoolVar(&preserveStatus, "preserve-status", false, "exit with the same status as COMMAND")
	flag.BoolVar(&foreground, "foreground", false, "..")
	flag.DurationVar(&killAfter, "kill-after", 0, "also send a KILL signal")
	flag.StringVar(&signal, "signal", "", "specify the signal")
}

func main() {
	flag.Parse()
	if signal != "" {
		// TODO: parse signal
		// syscall.SIGABRT
	}
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
	ctx, _ := context.WithTimeout(context.Background(), duration)
	cmd := exec.CommandContext(ctx, args[1], args[2:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		fmt.Println("err:", err)
		return
	}
	if killAfter > 0 {
		time.AfterFunc(killAfter, func() {
			// TODO: parse signal and send signal
			// cmd.Process.Signal()
		})
	}

	cmd.Wait()

}
