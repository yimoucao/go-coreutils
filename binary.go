package coreutils

import (
	"encoding/binary"
	"errors"
	"unsafe"
)

// ErrDetermEndian not able to get native endian
var ErrDetermEndian = errors.New("could not determine native endianness")

// NativeEndian represents system's native endianess
var NativeEndian binary.ByteOrder = func() binary.ByteOrder {
	endian, err := GetNativeEndian()
	if err != nil {
		panic(err)
	}
	return endian
}()

// GetNativeEndian return the system's native endianess
func GetNativeEndian() (binary.ByteOrder, error) {
	var nativeEndian binary.ByteOrder
	buf := [2]byte{}
	*(*uint16)(unsafe.Pointer(&buf[0])) = uint16(0xABCD)

	switch buf {
	case [2]byte{0xCD, 0xAB}:
		nativeEndian = binary.LittleEndian
	case [2]byte{0xAB, 0xCD}:
		nativeEndian = binary.BigEndian
	default:
		return nil, ErrDetermEndian
	}
	return nativeEndian, nil
}
