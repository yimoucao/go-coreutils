package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yimoucao/go-coreutils/pkg/checksum"
)

var (
	readText   bool
	readBinary bool
)

func init() {
	flag.BoolVar(&readText, "t", true, "read in text mode")
	flag.BoolVar(&readText, "text", true, "read in text mode")
	flag.BoolVar(&readBinary, "b", false, "read in binary mode")
	flag.BoolVar(&readBinary, "binary", false, "read in binary mode")
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, `Usage: sha512sum [OPTION]... [FILE]...
Print checksum and block counts for each FILE.

With no FILE, or when FILE is -, read standard input.
`)
		flag.PrintDefaults()
	}
}

// TODO: binary mode???

func main() {
	flag.Parse()
	args := flag.Args()

	algoFunc := checksum.SHA512Sum
	if len(args) == 0 {
		sum, err := algoFunc(os.Stdin)
		if err != nil {
			fmt.Println("stdin: ", err)
		} else {
			fmt.Printf("%x\n", sum)
		}
		return
	}

	for _, fname := range args {
		sum, err := checksum.SumFile(fname, algoFunc)
		if err != nil {
			fmt.Printf("%s: %v", fname, err)
		} else {
			fmt.Printf("%x\t%s\n", sum, fname)
		}
	}
}
