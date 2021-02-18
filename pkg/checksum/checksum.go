package checksum

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"io"
	"os"
)

// SumFile returns checksum with the given Algorithm.
// if fname == "-", data will be read from stdin
func SumFile(fname string, algoFunc AlgoFunc) (sum []byte, err error) {
	var f io.ReadCloser
	if fname == "-" {
		f = os.Stdin
	} else {
		f, err = os.Open(fname)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	defer f.Close()
	return algoFunc(f)
}

//AlgoFunc defines signature of a checksum function
type AlgoFunc func(reader io.Reader) (sum []byte, err error)

//SHA1Sum do sha1 checksum
var SHA1Sum AlgoFunc = func(reader io.Reader) (sum []byte, err error) {
	return hashSum(sha1.New(), reader)
}

//SHA224Sum do sha224 checksum
var SHA224Sum AlgoFunc = func(reader io.Reader) (sum []byte, err error) {
	return hashSum(sha256.New224(), reader)
}

//SHA256Sum do sha256 checksum
var SHA256Sum AlgoFunc = func(reader io.Reader) (sum []byte, err error) {
	return hashSum(sha256.New(), reader)
}

//SHA384Sum do sha384 checksum
var SHA384Sum AlgoFunc = func(reader io.Reader) (sum []byte, err error) {
	return hashSum(sha512.New384(), reader)
}

// SHA512Sum TODO: sha-512/224, sha-512/256 ???
var SHA512Sum AlgoFunc = func(reader io.Reader) (sum []byte, err error) {
	return hashSum(sha512.New(), reader)
}

// hashSum copy data into hash and do sum. NO RESET happens for hash state!
func hashSum(h hash.Hash, r io.Reader) (sum []byte, err error) {
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}
