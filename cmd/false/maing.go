package main

import "os"

/*
FALSE(1)                  BSD General Commands Manual                 FALSE(1)

NAME
     false -- Return false value.

SYNOPSIS
     false

DESCRIPTION
     The false utility always exits with a nonzero exit code.

SEE ALSO
     csh(1), sh(1), true(1)

STANDARDS
     The false utility conforms to IEEE Std 1003.2-1992 (``POSIX.2'').

4.2 Berkeley Distribution        July 24, 1991       4.2 Berkeley Distribution
*/

func main() {
	os.Exit(1)
}
