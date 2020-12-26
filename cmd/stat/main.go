package main

import (
	"flag"
	"fmt"
	"os"
)

func init() {

}

func main() {
	flag.Parse()

	files := flag.Args()
	if len(files) == 0 {
		os.Stderr.WriteString("stat: missing operand\n")
		os.Stderr.WriteString("Try 'stat --help' for more information.\n")
		os.Exit(1)
		return
	}

	for _, f := range files {
		if err := stat(f); err != nil {
			fmt.Println(err)
		}
	}
}
