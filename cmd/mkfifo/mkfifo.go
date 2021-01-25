package main

import (
	"flag"
	"fmt"
	"syscall"
)

func main() {
	flag.Parse()

	filenames := flag.Args()

	for _, filename := range filenames {
		if err := syscall.Mkfifo(filename, 0666); err != nil {
			fmt.Printf("mkfifo: cannot create fifo '%s': %s\n", filename, err)
		}

	}
}
