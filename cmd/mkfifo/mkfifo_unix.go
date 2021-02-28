// +build !windows
// +build !plan9

package main

import (
	"os"
	"syscall"
)

func mkfifo(fname string, mode os.FileMode) error {
	return syscall.Mkfifo(fname, 0666)
}
