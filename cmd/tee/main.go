package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	appendMode bool
)

func init() {
	flag.BoolVar(&appendMode, "append", false, "append to the given FILES")
}

func main() {
	flag.Parse()
	files := flag.Args()

	var fileflag = os.O_RDWR | os.O_CREATE
	if appendMode {
		fileflag |= os.O_APPEND
	} else {
		fileflag |= os.O_TRUNC
	}

	writers := []io.Writer{os.Stdout}
	for _, fname := range files {
		f, err := os.OpenFile(fname, fileflag, 0666)
		if err != nil {
			fmt.Fprint(os.Stderr, "err:", err)
			return
		}
		defer f.Close()
		writers = append(writers, f)
	}
	mw := io.MultiWriter(writers...)
	io.Copy(mw, os.Stdin)

	// TODO: investigate flushing behavior
}
