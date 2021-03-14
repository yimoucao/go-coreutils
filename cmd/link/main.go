package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	showHelp    bool
	showVersion bool
)

func init() {
	flag.BoolVar(&showHelp, "help", false, "display this help and exit")
	flag.BoolVar(&showVersion, "version", false, "ouput version information and exit")
}

func main() {
	flag.Parse()

	// using this if-else because we only want to show one if help and version
	// are both specified
	if showHelp {
		flag.Usage()
		return
	} else if showVersion {
		fmt.Println("dev")
		return
	}

	if flag.NArg() == 0 {
		fmt.Println("link: missing operand")
		fmt.Println("Try 'link --help' for more information.")
		os.Exit(1)
	}

	if flag.NArg() == 1 {
		fmt.Printf("link: missing operand after '%s'\n", flag.Arg(0))
		fmt.Println("Try 'link --help' for more information.")
		os.Exit(1)
	}

	if flag.NArg() > 2 {
		fmt.Printf("link: extra operand '%s'\n", flag.Arg(2))
		fmt.Println("Try 'link --help' for more information.")
		os.Exit(1)
	}
	if err := os.Link(flag.Arg(0), flag.Arg(1)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
