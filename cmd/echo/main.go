package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	for i, arg := range flag.Args() {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(arg)
	}
	fmt.Println()
}
