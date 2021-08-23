package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/yimoucao/go-coreutils/pkg/exit"
	"github.com/yimoucao/go-coreutils/pkg/ranges"
)

/*
b: bytes=LIST
c: characters=LIST
d: set delimiter
f: fields=LIST

LIST: one or more ranges separated by commas
N
N-
N-M
-M
*/

var (
	mode     string
	rangeStr string
	delim    string
)

func init() {
	flag.StringVar(&mode, "m", "", "mode can be: byte, character, field")
	flag.StringVar(&rangeStr, "r", "", "comma separated ranges")
	flag.StringVar(&delim, "d", "\t", "use DELIM instead of TAB for field delimiter")
}

func main() {
	flag.Parse()

	range_, err := ranges.Parse(rangeStr)
	if err != nil {
		exit.Error(err)
		return
	}
	cutOpts := &cutOpts{Mode: mode, Range: range_, Delim: delim}
	for _, f := range flag.Args() {
		if err := processFile(f, cutOpts); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

type cutOpts struct {
	Mode  string
	Range *ranges.Range
	Delim string
}

func processFile(filename string, opts *cutOpts) error {
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
		line := sc.Text()
		cutLine(line, opts)
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "read data err:", err)
		return err
	}
	return nil
}

func cutLine(line string, opts *cutOpts) {
	switch opts.Mode {
	case "bytes":
		// echo "你好啊世界" | cut -b1-3 -
		// it only prints "你"
		for i, b := range []byte(line) {
			if opts.Range.Has(i + 1) {
				fmt.Printf("%c", b)
			}
		}
		fmt.Println()
		return
	case "chars", "characters":
		/*
			echo "你好啊世界" | cut -c1-3 -
			it only prints "你"
			From: https://www.gnu.org/software/coreutils/manual/html_node/cut-invocation.html
			Select for printing only the characters in positions listed in character-list.
			The same as -b for now, but internationalization will change that.
		*/
		for i, r := range []rune(line) {
			if opts.Range.Has(i + 1) {
				fmt.Printf("%c", r)
			}
		}
		fmt.Println()
	case "fields":
		for i, field := range strings.Split(line, opts.Delim) {
			if opts.Range.Has(i + 1) {
				fmt.Print(field)
			}
		}
		fmt.Println()
	default:
		errExit("invalid mode:", opts.Mode)
	}
}

func errExit(v ...interface{}) {
	fmt.Fprintln(os.Stderr, v...)
	os.Exit(1)
}
