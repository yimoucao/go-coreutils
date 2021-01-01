package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"golang.org/x/crypto/blake2b"
)

// TODO: handler stdin

func main() {
	flag.Parse()
	filenames := flag.Args()
	if len(filenames) > 0 {
		for _, filename := range filenames {
			sum, err := b2sum(filename)
			if err != nil {
				fmt.Printf("%s: %s\n", filename, err)
				continue
			}
			fmt.Printf("%x  %s\n", sum, filename)
		}
	}
}

// TODO: looks like linux 64 is using blake2b 512 bit?
func b2sum(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	hash, err := blake2b.New512(nil)
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(hash, f); err != nil {
		return nil, err
	}
	return hash.Sum(nil), nil
}

/*Further Reading
blake2b: BLAKE2b is optimized for 64-bit platforms—including NEON-enabled ARMs—and produces digests of any size between 1 and 64 bytes.
blake2s: BLAKE2s is optimized for 8- to 32-bit platforms and produces digests of any size between 1 and 32 bytes
*/
