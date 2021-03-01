package main

import (
	"flag"
	"fmt"
	"time"
)

// update the access and modification times of each FILE to the current time

// flags
var (
	changeOnlyAccessTime bool
	changeOnlyModTime    bool
	timeWord             string // --time=WORD, WORD is access/atime, modify/mtime
	date                 string
	stamp                string
	reference            string
	noCreate             bool
	noDereference        bool
	showHelp             bool
	showVersion          bool
)

//Options are options parsed from flags
type Options struct {
	NoCreate      bool
	NoDereference bool
	ChangeMode    string // can be "atime", "mtime", "both"
	Time          time.Time
}

func init() {
	flag.BoolVar(&changeOnlyAccessTime, "a", false, "change only access time")
	flag.BoolVar(&changeOnlyModTime, "m", false, "change only modification time")
	flag.StringVar(&timeWord, "time", "", "access/atime, modify/mtime")
	flag.StringVar(&stamp, "t", "", "use [[CC]YY]MMDDhhmm[.ss] instead of current time")
	flag.StringVar(&date, "d", "", "parse STRING and use it instead of current time")
	flag.StringVar(&reference, "reference", "", "use this file's times instead of current time")
	flag.BoolVar(&noCreate, "c", false, "do not create any file")
	flag.BoolVar(&noDereference, "no-dereference", false, "affect each symbolic link instead of any referenced file")

	flag.BoolVar(&showVersion, "version", false, "show version")
}

func parseOptions() Options {
	opts := Options{
		NoCreate:      noCreate,
		NoDereference: noDereference,
		ChangeMode:    "both",
		Time:          time.Now(),
	}
	// changeOnlyAccessTime, changeOnlyModTime, timeWord
	// TODO:

	// date, stamp, reference
	// TODO:
	return opts
}

func main() {
	flag.Parse()
	opts := parseOptions()
	fnames := flag.Args()
	touchMany(opts, fnames...)
}

func touchMany(opts Options, fnames ...string) {
	for _, fname := range fnames {
		err := touchOne(opts, fname)
		if err != nil {
			fmt.Printf("%s: %v\n", fname, err)
		}
	}
}
