package main

import (
	"flag"
	"io"
	"os"

	"github.com/yimoucao/go-coreutils/pkg/exit"
)

//flags

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
		exit.Error(err)
	}
	defer f.Close()

	cat(f)
}

func cat(reader io.Reader) {
	io.Copy(os.Stdout, reader)
}
