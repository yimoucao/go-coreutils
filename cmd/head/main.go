package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	nLines int
	nBytes int
)

func init() {
	flag.IntVar(&nLines, "n", 10, "display sthe first count lines")
}

func main() {
	flag.Parse()

	// determine mode, LineMode or ByteMode

	files := flag.Args()
	for i, f := range files {
		if i > 0 {
			os.Stdout.WriteString("\n")
		}
		head(LineMode, nLines, f, os.Stdout)
	}

}

type Mode uint8

const (
	LineMode Mode = iota
	ByteMode
)

func head(m Mode, c int, file string, w io.Writer) error {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return err
	}
	fmt.Fprintln(w, "===>", file, "<===")
	scanner := bufio.NewScanner(f)
	if m == ByteMode {
		scanner.Split(bufio.ScanBytes)
	}
	for ; c > 0; c-- {
		if !scanner.Scan() {
			return scanner.Err()
		}
		w.Write(scanner.Bytes())
		if m == LineMode {
			w.Write([]byte{'\n'})
		}
	}
	return nil
}
