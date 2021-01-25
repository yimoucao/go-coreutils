package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var (
	showHelp  bool
	targetDir string
)

func init() {
	flag.BoolVar(&showHelp, "help", false, "show help")
	flag.StringVar(&targetDir, "t", "", "copy all SOURCE into target directory")

	flag.Parse()
}

func main() {
	args := flag.Args()
	lastIdx := len(args) - 1
	if len(targetDir) != 0 {
		cpCmd(args, targetDir)
	} else {
		cpCmd(args[:lastIdx], args[lastIdx])
	}

}

func cpCmd(src []string, target string) {
	var n = len(src)
	if n == 0 {
		// log
	} else if n == 1 {
		cpToFile(src[0], target)
	} else {
		cpToDir(target, src...)
	}
}

func cpToFile(src string, target string) {
	dstF, err := os.Create(target)
	if err != nil {
		// log err
		return
	}
	defer dstF.Close()

	srcF, err := os.Open(src)
	if err != nil {
		// log err
		return
	}
	defer srcF.Close()

	_, err = io.Copy(dstF, srcF)
	if err != nil {
		//log err
		return
	}
	return
}

func cpToDir(target string, src ...string) {
	// mkdir all for target
	// TODO: opts like -R
	for _, name := range src {
		srcF, err := os.Open(name)
		if err != nil {
			// log err
			continue
		}
		defer srcF.Close()

		basename := srcF.Name()
		fmt.Print("file basename:", basename)

		dst := filepath.Join(target, basename)
		dstF, err := os.Create(dst) // TODO: check file existence, overwrite or not
		if err != nil {
			//
		}
		defer dstF.Close()

		_, err = io.Copy(dstF, srcF)
		if err != nil {
			//
		}
	}
}
