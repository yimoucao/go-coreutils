package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
)

// TODO: handler stdin

func main() {
	flag.Parse()
	filenames := flag.Args()
	if len(filenames) > 0 {
		for _, filename := range filenames {
			sum, err := md5sum(filename)
			if err != nil {
				fmt.Printf("%s: %s\n", filename, err)
				continue
			}
			fmt.Printf("%x  %s\n", sum, filename)
		}
	}
}

func md5sum(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, f); err != nil {
		return nil, err
	}
	return hash.Sum(nil), nil
}
