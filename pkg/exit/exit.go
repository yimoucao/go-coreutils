package exit

import (
	"fmt"
	"os"
)

func Error(v ...interface{}) {
	fmt.Fprintln(os.Stderr, v...)
	os.Exit(1)
}

func ErrorCode(code int, v ...interface{}) {
	fmt.Fprintln(os.Stderr, v...)
	os.Exit(code)
}

func Errorf(format string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, format, v...)
	os.Exit(1)
}

func ErrorfCode(code int, format string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, format, v...)
	os.Exit(code)
}
