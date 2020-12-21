package main

import (
	"flag"
	"os"
	"os/exec"
)

// TODO: run and test

var (
	showVersion bool
)

var shellPath string

func init() {
	flag.BoolVar(&showVersion, "version", false, "show version info")

	shellPath = os.Getenv("SHELL")
	if shellPath == "" {
		shellPath = "/bin/sh"
	}
}

func main() {
	flag.Parse()
	args := flag.Args()
	run(args[0], args[1:])
}

func run(root string, args []string) error {
	if len(args) == 0 {
		args = []string{shellPath, "-i"}
	}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.SysProcAttr.Chroot = root

	return cmd.Run()
}
