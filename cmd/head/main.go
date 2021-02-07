package main

import (
	"bufio"
	"container/list"
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

// Mode can be LineMode or ByteMode
type Mode uint8

const (
	LineMode Mode = iota
	ByteMode
)

func head(m Mode, n int, file string, w io.Writer) error {
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
	if n >= 0 {
		return scanWrite(n, scanner, w, m == LineMode)
	}
	return slidingScanWrite(-n, scanner, w, m == LineMode)
}

func scanWrite(n int, sc *bufio.Scanner, w io.Writer, newline bool) error {
	for ; n > 0; n-- {
		if !sc.Scan() {
			return sc.Err()
		}
		w.Write(sc.Bytes())
		if newline {
			w.Write([]byte{'\n'})
		}
	}
	return nil
}

func slidingScanWrite(width int, sc *bufio.Scanner, w io.Writer, newline bool) error {
	bufl := list.New() // front oldes, back newest
	i := 0
	for ; sc.Scan(); i++ {
		if i >= width {
			v := bufl.Remove(bufl.Front())
			w.Write(v.([]byte))
			if newline {
				w.Write([]byte{'\n'})
			}
		}
		bufl.PushBack(sc.Bytes())
	}
	return sc.Err()
}
