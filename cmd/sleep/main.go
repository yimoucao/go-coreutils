package main

import (
	"flag"
	"os"
	"time"
)

var (
	showHelp    bool
	showVersion bool
)

func init() {
	flag.BoolVar(&showHelp, "help", false, "show help")
	flag.BoolVar(&showVersion, "version", false, "show version")
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		return
	}
	d, err := time.ParseDuration(args[0])
	if err != nil {
		os.Exit(1)
		return
	}
	if d < 0 {
		os.Exit(1)
		return
	}
	time.Sleep(d)
}
