package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
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
	ownerAndGroup, files := args[0], args[1:]

	for _, fname := range files {
		err := chown(ownerAndGroup, fname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "chown: %s: %v", fname, err)
		}
	}
}

func chown(ownerAndGroup, fname string) (err error) {
	var gid = -1
	var uid = -1

	og := strings.Split(ownerAndGroup, ":")
	if len(og) == 2 && og[1] != "" {
		gid, err = lookupGID(og[1])
		if err != nil {
			return err
		}
	}
	uid, err = lookupUID(og[0])
	if err != nil {
		return err
	}
	return os.Chown(fname, uid, gid)
}

// group can be either groupname or numeric value
func lookupGID(group string) (int, error) {
	gid, err := strconv.Atoi(group)
	if err == nil {
		return gid, nil
	}
	grp, err := user.LookupGroup(group)
	if err != nil {
		return 0, err
	}
	gid, err = strconv.Atoi(grp.Gid)
	if err != nil {
		return 0, err
	}
	return gid, nil
}

// usr can be either groupname or numeric value
func lookupUID(userstr string) (int, error) {
	uid, err := strconv.Atoi(userstr)
	if err == nil {
		return uid, nil
	}
	usr, err := user.Lookup(userstr)
	if err != nil {
		return 0, err
	}
	uid, err = strconv.Atoi(usr.Uid)
	if err != nil {
		return 0, err
	}
	return uid, nil
}
