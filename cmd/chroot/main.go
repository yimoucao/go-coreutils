package main

import (
	"flag"
	"os"
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
	chroot(args[0], args[1:])
}
