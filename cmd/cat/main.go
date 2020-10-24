package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

//flags

func errExit(a ...interface{}) {
	fmt.Fprint(os.Stderr, a...)
	os.Exit(1)
}

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		cat(os.Stdin)
	} else {
		for _, fname := range flag.Args() {
			catFile(fname)
		}
	}
}

func catFile(fname string) {
	f, err := os.Open(fname)
	if err != nil {
		errExit(err)
	}
	defer f.Close()

	cat(f)
}

func cat(reader io.Reader) {
	io.Copy(os.Stdout, reader)
}
