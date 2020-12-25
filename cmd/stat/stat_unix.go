// +build !windows
// +build !plan9

package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

// TODO: get username for UID, groupname for GID

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
	fmt.Printf("Uid: %v\t", sys.Uid)
	fmt.Printf("Gid: %v\n", sys.Gid)

	fmt.Printf("Access: %v\n", time.Unix(sys.Atim.Unix()))
	fmt.Printf("Modify: %v\n", stat.ModTime())
	fmt.Printf("Change: %v\n", time.Unix(sys.Ctim.Unix()))
	fmt.Printf("Birth: -\n")
	return nil
}

func modeTypeString(mode os.FileMode) string {
	if mode&os.ModeDir != 0 {
		return "directory"
	}
	if mode&os.ModeSymlink != 0 {
		return "symbolic link"
	}
	if mode&os.ModeNamedPipe != 0 {
		return "fifo"
	}
	if mode&os.ModeSocket != 0 {
		return "socket"
	}
	if mode&os.ModeDevice != 0 {
		return "block special file"
	}
	if mode&os.ModeCharDevice != 0 {
		return "character special file"
	}
	if mode&os.ModeIrregular != 0 {
		return "irregular file"
	}
	return "regular file"
}
