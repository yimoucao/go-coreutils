package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	prefix  = "xx"
	nDigits = 2
)

func init() {

	flag.Parse()
}

func main() {
	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		return
	}
	if len(args) == 1 {
		fmt.Fprintf(os.Stderr, "csplit: missing operand after '%s'", args[0])
		return
	}
	csplitCmd(args[0], args[1:])
}

func nameFmt() string {
	return fmt.Sprintf("%s%%%dd", prefix, nDigits)
}

func csplitCmd(filename string, patterns []string) {

}
