package main

import (
	"fmt"
	"os"
)

// const (
// 	TTY_STDIN_NOTTY = 1 + iota
// 	TTY_FAILURE
// 	TTY_WRITE_ERROR
// )

const fd0 = "/proc/self/fd/0"

func main() {
	dest, err := os.Readlink(fd0)
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println(dest)
	}
}
