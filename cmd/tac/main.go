package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// TODO: flags

func main() {
	flag.Parse()
	files := flag.Args()
	var err error
	for _, fname := range files {
		var f io.Reader

		if fname == "-" {
			f = os.Stdin
		} else {
			f, err = os.Open(fname)
			if err != nil {
				fmt.Fprintln(os.Stderr, "err:", err)
				continue
			}
		}
		tac(f)
	}
}

// TODO: too much memory cost for a large file.
// research orginal tac for a better solution
func tac(r io.Reader) error {
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		defer fmt.Println(line)
	}
	if err := sc.Err(); err != nil {
		return err
	}
	return nil
}
