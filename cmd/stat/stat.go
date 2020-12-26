package main

import "os"

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
