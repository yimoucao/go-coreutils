package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	showHelp = flag.Bool("help", false, "display help text")
)

func init() {
	flag.Parse()
}

func main() {
	args := flag.Args()
	fmt.Println(args)
	mode, files := args[0], args[1:]

	for _, fname := range files {
		err := chmod(mode, fname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "chmod: %s: %v", fname, err)
		}
	}
}

func chmod(modstr, fname string) error {
	fi, err := os.Stat(fname)
	if err != nil {
		return err
	}
	mod, err := parseFileMode(modstr, fi.Mode())
	if err != nil {
		return err
	}
	return os.Chmod(fname, mod)
}

func parseFileMode(modstr string, old os.FileMode) (os.FileMode, error) {
	// chmod 777 file
	// chmod +x file
	// chmod u+x
	// symbolic: [ugoa][[-+=][perms...]...] perms: rwxXst
	// octal: 0777, 0666
	modUint, err := strconv.ParseUint(modstr, 8, 0)
	if err == nil {
		return os.FileMode(modUint), nil
	}
	if !errors.Is(err, strconv.ErrSyntax) {
		return old, err
	}

	// TODO: parse symbolic mod
	return old, nil
}
