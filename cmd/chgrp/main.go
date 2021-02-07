package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
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
	group, files := args[0], args[1:]

	for _, fname := range files {
		err := chgrp(group, fname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "chgrp: %s: %v", fname, err)
		}
	}
}

func chgrp(grpname, fname string) error {
	group, err := user.LookupGroup(grpname)
	if err != nil {
		return err
	}
	gid, err := strconv.Atoi(group.Gid)
	if err != nil {
		return err
	}
	return os.Chown(fname, -1, gid)
}
