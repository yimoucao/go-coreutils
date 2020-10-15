package coreutils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"syscall"
)

// https://github.com/freebsd/freebsd/blob/a3bbb67b61eb7fd57b00e602672944da810e7417/include/utmpx.h#L43
// https://opensource.apple.com/source/Libc/Libc-391/gen/NetBSD/utmpx.c.auto.html
// https://github.com/ericlagergren/go-gnulib/tree/master/utmp
// https://github.com/patrickToca/go-coreutils/blob/master/who/who.go

type Utmpx struct {
	Type    int16
	Timeval syscall.Timeval
	ID      [8]byte
	PID     int32
	User    [32]byte
	Line    [16]byte
	Host    [128]byte
	_       [64]byte
}

func GetUtmpx() []Utmpx {
	bs, err := ioutil.ReadFile("/var/run/utmpx")
	if err != nil {
		panic(err)
	}
	lines := bytes.Split(bs, []byte{'\n'})
	fmt.Println("lines: ", len(lines))
	// var res []Utmpx
	for i, l := range lines {
		fmt.Printf("size: %d, %X\n", len(l), l)
		if i == 0 {
			fmt.Printf("%s\n", l)
			continue
		}
		records, err := readAll(bytes.NewBuffer(l))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(records)
		continue
	}
	return nil
}

func readAll(reader io.Reader) ([]*Utmpx, error) {
	var us []*Utmpx

	for {
		u, readErr := readEnt(reader)
		if readErr != nil {
			if readErr == io.EOF {
				break
			}
			return nil, readErr
		}
		us = append(us, u)
	}

	return us, nil
}

func readEnt(reader io.Reader) (*Utmpx, error) {
	var u Utmpx
	if err := binary.Read(reader, nativeEndian, &u); err != nil {
		return nil, err
	}
	return &u, nil
}
