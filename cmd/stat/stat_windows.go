package main

import (
	"fmt"
	"os"
	"os/user"
	"syscall"
	"time"
)

func stat(filename string) error {
	stat, err := os.Stat(filename)
	if err != nil {
		return err
	}
	sys := stat.Sys().(*syscall.Win32FileAttributeData)
	fmt.Printf("%+v\n", sys)

	fmt.Printf("File: %v\n", stat.Name())

	fmt.Printf("Size: %v\t", stat.Size())
	// fmt.Printf("Blocks: %v\t", sys.Blocks)
	// fmt.Printf("IO Block: %v\t", sys.Blksize)
	fmt.Printf("%v\n", modeTypeString(stat.Mode()))

	// fmt.Printf("Device: %xh/%dd\t", sys.Dev, sys.Dev)
	// fmt.Printf("Inode: %d\t", sys.Ino)
	// fmt.Printf("Links: %d\n", sys.Nlink)

	fmt.Printf("Access: (%#o/%v)\n", stat.Mode(), stat.Mode())
	u, _ := user.Current()
	fmt.Printf("Uid: (%v/%v)\n", u.Uid, u.Username)
	grp, err := user.LookupGroupId(u.Gid)
	if err != nil {
		fmt.Printf("Gid: (%v/%v)\n", u.Gid, "UNKNOWN")
	} else {
		fmt.Printf("Gid: (%v/%v)\n", grp.Gid, grp.Name)
	}

	fmt.Printf("Access: %v\n", time.Unix(0, sys.LastAccessTime.Nanoseconds()))
	fmt.Printf("Modify: %v\n", stat.ModTime())
	fmt.Printf("Change: %v\n", stat.ModTime()) // see os/types_windows.go:136 ModTime()
	fmt.Printf("Birth: %v\n", time.Unix(0, sys.CreationTime.Nanoseconds()))
	return nil
}
