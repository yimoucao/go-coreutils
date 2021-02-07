package main

import (
	"fmt"
	"runtime"
)

/*
method 1: runtime.GOARCH

method 2: exec.Command, uname

// TODO: which one is  correct? is there another way?
// TODO:
	runtime.GOARCH -> amd64 while arch cmd -> i386
	runtime.GOARCH -> amd64 while arch cmd -> x86_64

*/

func main() {
	fmt.Println(runtime.GOARCH)
}
