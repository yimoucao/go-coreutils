package sysinfo

import (
	"bytes"
	"encoding/binary"
	"syscall"
	"time"

	"github.com/yimoucao/go-coreutils/pkg/endian"
)

// TODO
type Sysinfo_t struct {
	Uptime    time.Duration
	Loads     interface{} // TODO
	totalram  int         // TODO
	totalswap int         // TODO
}

// TODO: return sysinfo
func Sysinfo() {

}

var timeZero = time.Time{}

func BootTime() (time.Time, error) {
	v, err := syscall.Sysctl("kern.boottime")
	if err != nil {
		return timeZero, err
	}

	buf := bytes.NewBufferString(v)
	for buf.Len() < 16 {
		buf.WriteByte(0x0)
	}

	var tv syscall.Timeval
	if err := binary.Read(buf, endian.NativeEndian, &tv); err != nil {
		return timeZero, err
	}

	return time.Unix(tv.Unix()), nil
}

// ref: https://github.com/lattera/freebsd/blob/401a161083850a9a4ce916f37520c084cff1543b/lib/libc/gen/getloadavg.c#L51

// TODO: use big number for fscale???
type loadavg struct {
	Ldavg  [3]uint32
	_      uint32
	Fscale uint32
	_      uint32
}

func LoadAvg() ([3]float32, error) {
	var res [3]float32
	v, err := syscall.Sysctl("vm.loadavg")
	if err != nil {
		return res, err
	}

	buf := bytes.NewBufferString(v)
	for buf.Len() < 24 {
		buf.WriteByte(0x0)
	}

	var ldavg loadavg
	if err := binary.Read(buf, endian.NativeEndian, &ldavg); err != nil {
		return res, err
	}

	for i := 0; i < 3; i++ {
		res[i] = float32(ldavg.Ldavg[i]) / float32(ldavg.Fscale)
	}

	return res, nil
}

// Totalram returns the size of memmory utilizing sysctl
func Totalram() uint64 {
	v, err := syscall.Sysctl("hw.memsize")
	if err != nil {
		return 0
	}

	buf := bytes.NewBufferString(v)
	for buf.Len() < 8 {
		buf.WriteByte(0x0)
	}

	var res uint64
	if err := binary.Read(buf, endian.NativeEndian, &res); err != nil {
		return 0
	}

	return res
}

// sysctl get data with given name and write into data
func sysctl(name string, data interface{}) error {
	// TODO: too much reflect code. performance?
	return nil
}
