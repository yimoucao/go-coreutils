package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

// TODO: flags
// TODO: some output fmt isn't right. e.g. seq 1 0.5 4

var (
	separator string
)

func init() {
	flag.StringVar(&separator, "separater", "\n", "sep")
}

func main() {
	flag.Parse()
	args := flag.Args()

	var err error
	first, inc, last := 1.0, 1.0, 0.0
	if len(args) == 1 {
		last, err = strconv.ParseFloat(args[0], 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "seq: invalid floating point argument:", strconv.Quote(args[0]))
			os.Exit(1)
			return
		}
	}

	if len(args) == 2 {
		first, err = strconv.ParseFloat(args[0], 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "seq: invalid floating point argument:", strconv.Quote(args[0]))
			os.Exit(1)
			return
		}
		last, err = strconv.ParseFloat(args[1], 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "seq: invalid floating point argument:", strconv.Quote(args[0]))
			os.Exit(1)
			return
		}
	}

	if len(args) == 3 {
		first, err = strconv.ParseFloat(args[0], 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "seq: invalid floating point argument:", strconv.Quote(args[0]))
			os.Exit(1)
			return
		}
		inc, err = strconv.ParseFloat(args[1], 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "seq: invalid floating point argument:", strconv.Quote(args[0]))
			os.Exit(1)
			return
		}
		last, err = strconv.ParseFloat(args[2], 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "seq: invalid floating point argument:", strconv.Quote(args[0]))
			os.Exit(1)
			return
		}
	}

	for n := first; n <= last; n += inc {
		fmt.Print(n)
		fmt.Print(separator)
	}
}
