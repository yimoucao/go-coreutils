package main

import (
	"fmt"
	"os"
)

/*
YES(1)                    BSD General Commands Manual                   YES(1)

NAME
     yes -- be repetitively affirmative

SYNOPSIS
     yes [expletive]

DESCRIPTION
     yes outputs expletive, or, by default, ``y'', forever.

HISTORY
     The yes command appeared in 4.0BSD.

4th Berkeley Distribution        June 6, 1993        4th Berkeley Distribution
*/

func main() {
	var expletive = "y"
	if len(os.Args) >= 2 {
		expletive = os.Args[1]
	}
	for {
		fmt.Println(expletive)
	}
}
