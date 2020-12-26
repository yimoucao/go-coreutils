// +build !windows
// +build !plan9

package main

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"syscall"
	"time"
)

func stat(filename string) error {
	stat, err := os.Stat(filename)
	if err != nil {
		return err
	}
	sys := stat.Sys().(*syscall.Stat_t)
	fmt.Printf("File: %v\n", stat.Name())

	fmt.Printf("Size: %v\t", stat.Size())
	fmt.Printf("Blocks: %v\t", sys.Blocks)
	fmt.Printf("IO Block: %v\t", sys.Blksize)
	fmt.Printf("%v\n", modeTypeString(stat.Mode()))

	fmt.Printf("Device: %xh/%dd\t", sys.Dev, sys.Dev)
	fmt.Printf("Inode: %d\t", sys.Ino)
	fmt.Printf("Links: %d\n", sys.Nlink)

	fmt.Printf("Access: (%#o/%v)\t", stat.Mode(), stat.Mode())
	userid, username := lookupUser(sys.Uid)
	fmt.Printf("Uid: (%v/%v)\t", userid, username)
	groupid, groupname := lookupGroup(sys.Gid)
	fmt.Printf("Gid: (%v/%v)\n", groupid, groupname)

	fmt.Printf("Access: %v\n", time.Unix(sys.Atim.Unix()))
	fmt.Printf("Modify: %v\n", stat.ModTime())
	fmt.Printf("Change: %v\n", time.Unix(sys.Ctim.Unix()))
	fmt.Printf("Birth: -\n")
	return nil
}

func lookupUser(uid uint32) (id, name string) {
	uidStr := strconv.Itoa(int(uid))
	usr, err := user.LookupId(uidStr)
	if err != nil {
		return uidStr, "unknown"
	}
	return usr.Uid, usr.Username
}

func lookupGroup(gid uint32) (id, name string) {
	gidStr := strconv.Itoa(int(gid))
	grp, err := user.LookupGroupId(gidStr)
	if err != nil {
		return gidStr, "unknown"
	}
	return grp.Gid, grp.Name
}
