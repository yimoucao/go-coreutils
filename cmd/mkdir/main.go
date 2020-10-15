package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		parent  bool
		verbose bool
		mode    = os.ModePerm // default mode
	)
	if len(os.Args) < 2 {
		usageAndExit(1)
	}
	var i = 1
	for ; i < len(os.Args); i++ {
		arg := os.Args[i]
		if !strings.HasPrefix(arg, "-") {
			break
		}
		switch arg {
		case "-p":
			parent = true
		case "-v":
			verbose = true
		case "-m":
			i++
			setMode(i, &mode)
		case "-pv", "-vp":
			parent = true
			verbose = true
		case "-pm", "-mp":
			parent = true
			i++
			setMode(i, &mode)
		case "-vm", "-mv":
			verbose = true
			i++
			setMode(i, &mode)
		case "-vpm", "-pvm", "-mvp", "-vmp", "-pmv", "-mpv":
			parent = true
			verbose = true
			i++
			setMode(i, &mode)
		default:
			fmt.Printf("mkdir: illegal option -- %s\n", arg)
			usageAndExit(1)
		}
	}

	if i == len(os.Args) {
		usageAndExit(1)
	}

	var err error
	for ; i < len(os.Args); i++ {
		arg := os.Args[i]
		if err = mkdir(arg, os.FileMode(mode), parent); err != nil {
			fmt.Println(err)
			continue
		}
		if verbose {
			fmt.Printf("mkdir: created directory '%s'\n", arg)
		}
	}
	if err != nil {
		os.Exit(1)
	}
}

func setMode(argIdx int, mode *os.FileMode) {
	if argIdx == len(os.Args) {
		fmt.Println("mkdir: option requires an argument -- m")
		usageAndExit(1)
	}
	mode64, err := strconv.ParseUint(os.Args[argIdx], 10, 32)
	if err != nil {
		// fmt.Println(err) // GNU mkdir doesn't print out err but only usage
		usageAndExit(1)
	}
	*mode = os.FileMode(mode64)
}

// syscall internally handles umask
func mkdir(path string, mode os.FileMode, intermediate bool) error {
	if intermediate {
		return os.MkdirAll(path, mode)
	}
	return os.Mkdir(path, mode)
}

func usageAndExit(code int) {
	printUsage()
	os.Exit(code)
}

func printUsage() {
	fmt.Println("usage: mkdir [-pv] [-m mode] directory ...")
}
