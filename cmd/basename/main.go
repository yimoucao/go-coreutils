package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	useAll bool
	suffix string
)

func init() {
	flag.BoolVar(&useAll, "a", false, "???")
	flag.StringVar(&suffix, "s", "", "suffix")
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println(filepath.Base(flag.Arg(0)))
}
