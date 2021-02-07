package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"syscall"
)

var (
	sig   int
	list  int
	table int
)

func init() {
	flag.IntVar(&sig, "signal", 0, "specify signal to be sent.")
	flag.IntVar(&list, "list", 0, "List signal names.")
	flag.IntVar(&table, "table", 0, "List signal names in a nice table.")
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "kill: usage: kill [-s sigspec | -n signum | -sigspec] pid | jobspec ... or kill -l [sigspec]")
		os.Exit(1)
	}
	for _, pid := range args {
		signal(pid, syscall.SIGTERM)
	}
}

func signal(pidstr string, sig os.Signal) error {
	pid, err := strconv.Atoi(pidstr)
	if err != nil {
		return errors.New("not int id")
	}
	proc, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	return proc.Signal(sig)
}
