package main

import (
	"errors"
	"os"
	"syscall"
	"time"
)

func touchOne(opts Options, fname string) error {
	fi, err := os.Stat(fname)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
		// file not exist
		if opts.NoCreate {
			return nil
		}
		_, err := os.Create(fname)
		return err
	}
	mtime := fi.ModTime()
	atime := time.Unix(fi.Sys().(*syscall.Stat_t).Atim.Unix())

	if opts.ChangeMode == "atime" {
		atime = opts.Time
	} else if opts.ChangeMode == "mtime" {
		mtime = opts.Time
	} else {
		atime = opts.Time
		mtime = opts.Time
	}

	return os.Chtimes(fname, atime, mtime)
}
