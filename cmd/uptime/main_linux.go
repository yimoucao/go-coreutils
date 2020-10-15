package main

import (
	"fmt"
	"syscall"
)

func main() {
	if err := syscall.Sysinfo(); err != nil {
		fmt.Println(err)
	}
	syscall.synn
}
