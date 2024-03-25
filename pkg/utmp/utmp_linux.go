package utmp

import (
	"bytes"
	binary "encoding/binary"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/yimoucao/go-coreutils/pkg/endian"
)

// References:
// https://en.wikipedia.org/wiki/Utmp
// In sum, depends on linux/unix flavor.utmpx is specified in POSIX however linux
// variants chose different implementations/locations.
// TODO: find out what each linux distro does...
// https://github.com/ericlagergren/go-gnulib/tree/master/utmp
// https://github.com/patrickToca/go-coreutils/blob/master/who/who.go

const (
	UtmpxFile   = "/var/run/utmp"
	Wtmpxfile   = "/var/log/wtmp"
	LastLogFile = "/var/log/lastlog"
)

type ExitStatus struct {
	X__e_termination int16
	X__e_exit        int16
}

type TimeVal struct {
	Sec  int32
	Usec int32
}

type Utmp struct {
	Type              int16
	Pad_cgo_0         [2]byte
	Pid               int32
	Line              [32]byte
	Id                [4]byte
	User              [32]byte
	Host              [256]byte
	Exit              ExitStatus
	Session           int32
	Tv                TimeVal
	Addr_v6           [4]int32
	X__glibc_reserved [20]byte
}

func GetUtmp() []Utmp {
	bs, err := ioutil.ReadFile(UtmpxFile)
	if err != nil {
		//// comment this out because `users` cmd show no error on wsl where /var/run/utmp doens't exist
		// panic(err)
		return nil
	}
	lines := bytes.Split(bs, []byte{'\n'})
	fmt.Println("lines: ", len(lines))
	// var res []Utmp
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

func readAll(reader io.Reader) ([]*Utmp, error) {
	var us []*Utmp

	for {
		u, readErr := readEntry(reader)
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

func readEntry(reader io.Reader) (*Utmp, error) {
	var u Utmp
	if err := binary.Read(reader, endian.NativeEndian, &u); err != nil {
		return nil, err
	}
	return &u, nil
}
