package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	useBSD  bool
	useSysV bool
)

func init() {
	flag.BoolVar(&useBSD, "r", true, "use BSD sum algo, use 1k blocks")
	flag.BoolVar(&useSysV, "s", false, "use System V sum algo, use 512 bytes blocks")
	flag.BoolVar(&useSysV, "sysv", false, "use System V sum algo, use 512 bytes blocks")
	// flag.BoolVar(&useSysV, "help", false, "use System V sum algo, use 512 bytes blocks")
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, `Usage: sum [OPTION]... [FILE]...
Print checksum and block counts for each FILE.

With no FILE, or when FILE is -, read standard input.
`)
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	args := flag.Args()

	algoFunc := bsdSum
	if useSysV {
		algoFunc = sysVSum
	}

	if len(args) == 0 {
		sum, blocks, err := algoFunc(os.Stdin)
		if err != nil {
			fmt.Println("stdin: ", err)
		} else {
			fmt.Printf("%d\t%d\n", sum, blocks)
		}
	}

	for _, fname := range args {
		err := sumFile(fname, algoFunc)
		if err != nil {
			fmt.Printf("%s: %v", fname, err)
		}
	}
}

func sumFile(fname string, algoFunc AlgoFunc) (err error) {
	var f io.ReadCloser
	if fname == "-" {
		f = os.Stdin
	} else {
		f, err = os.Open(fname)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	defer f.Close()
	sum, blocks, err := algoFunc(f)
	if err != nil {
		return err
	}
	fmt.Printf("%d\t%d\t%s\n", sum, blocks, fname)
	return nil
}

type AlgoFunc func(reader io.Reader) (sum int, blocks int, err error)

func bsdSum(reader io.Reader) (sum int, blocks int, err error) {
	for {
		buf := make([]byte, 1024)
		n, err := reader.Read(buf)
		if n > 0 {
			blocks++
			// https://en.wikipedia.org/wiki/BSD_checksum
			for _, b := range buf[:n] {
				sum = (sum >> 1) + ((sum & 1) << 15)
				sum += int(b)
				sum &= int(0xffff)
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, 0, err
		}
	}
	return sum, blocks, nil
}

func sysVSum(reader io.Reader) (sum int, blocks int, err error) {
	for {
		buf := make([]byte, 512)
		n, err := reader.Read(buf)
		if n > 0 {
			blocks++
			for _, b := range buf[:n] {
				sum += int(b)
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, 0, err
		}
	}
	r := (sum & 0xffff) + ((sum & 0xffffffff) >> 16)
	sum = (r & 0xffff) + (r >> 16)
	return sum, blocks, nil
}
