package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"hash/crc32"
	"io"
	"os"
	// "github.com/yimoucao/go-coreutils"
)

func main() {
	flag.Parse()
	filenames := flag.Args()
	if len(filenames) > 0 {
		for _, filename := range filenames {
			var (
				sum    []byte
				n      int64
				err    error
				sumU32 uint32
			)

			// sum, n, err = crc32go(filename)
			// if err != nil {
			// 	fmt.Printf("%s: %s\n", filename, err)
			// 	continue
			// }
			// sumU32 = binary.BigEndian.Uint32(sum)
			// fmt.Printf("%d %d %s\n", sumU32, n, filename)

			sum, n, err = unixCksum32(filename)
			if err != nil {
				fmt.Printf("%s: %s\n", filename, err)
				continue
			}
			sumU32 = binary.BigEndian.Uint32(sum)
			fmt.Printf("%d %d %s\n", sumU32, n, filename)
		}
	}
}

/*
https://en.wikipedia.org/wiki/Cksum#Algorithm
polynomial 0x04C11DB7 and little endian
golang's crc32 use bit-reversed polynomial and output big endian
IEEE: 0xedb88320 is reversed
TODO: find out why go's CRC different than cksum
*/
func crc32go(filename string) ([]byte, int64, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, 0, err
	}
	defer f.Close()
	hash := crc32.NewIEEE()
	n, err := io.Copy(hash, f)
	if err != nil {
		return nil, 0, err
	}

	// for c := uint64(n); c != 0; c >>= 8 {
	// 	b := c & 0xff
	// 	fmt.Println(b)
	// 	fmt.Println(byte(b))
	// 	hash.Write([]byte{byte(b)})
	// }

	return hash.Sum(nil), n, nil
}

func unixCksum32(filename string) ([]byte, int64, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, 0, err
	}
	defer f.Close()
	hash := New()
	n, err := io.Copy(hash, f)
	if err != nil {
		return nil, 0, err
	}

	return hash.Sum(nil), n, nil
}
