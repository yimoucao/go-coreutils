package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
)

var (
	initial bool // means only initial. default is converting all
	tabs    uint
)

func init() {
	flag.BoolVar(&initial, "initial", false, "do not convert tabs after non blanks")
	flag.UintVar(&tabs, "tabs", 8, "have tabs N characters apart")
}

func main() {
	flag.Parse()

	opts := &expandOpts{
		ConvertOnlyInitial: initial,
		NumberOfSpaces:     tabs,
	}
	files := flag.Args()
	if len(files) == 0 {
		files = append(files, "-")
	}
	for _, f := range files {
		if err := processFile(f, opts); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

type expandOpts struct {
	ConvertOnlyInitial bool
	NumberOfSpaces     uint
}

func processFile(filename string, opts *expandOpts) error {
	f := os.Stdin
	if filename != "-" {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, "open file err:", err)
			return err
		}
		defer f.Close()
	}

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Bytes()
		fmt.Fprintf(os.Stdout, "%s\n", expandLine(line, opts))
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "read data err:", err)
		return err
	}
	return nil
}

func expandLine(line []byte, opts *expandOpts) []byte {
	var bs bytes.Buffer
	isInitial := true
	for _, b := range line {
		if b == '\t' {
			if !isInitial && opts.ConvertOnlyInitial {
				bs.WriteByte(b)
			} else {
				bs.Write(makeNByte(' ', opts.NumberOfSpaces))
			}
		} else {
			isInitial = false
			bs.WriteByte(b)
		}
	}
	return bs.Bytes()
}

func makeNByte(b byte, N uint) []byte {
	bs := make([]byte, 0, N)
	for i := uint(0); i < N; i++ {
		bs = append(bs, b)
	}
	return bs
}
