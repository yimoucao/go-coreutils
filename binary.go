package coreutils

import (
	"encoding/binary"
	"errors"
	"unsafe"
)

var DetermEndianErr = errors.New(("Could not determine native endianness."))

// NativeEndian return the system's native endianess
func NativeEndian() (binary.ByteOrder, error) {
	var nativeEndian binary.ByteOrder
	buf := [2]byte{}
	*(*uint16)(unsafe.Pointer(&buf[0])) = uint16(0xABCD)

	switch buf {
	case [2]byte{0xCD, 0xAB}:
		nativeEndian = binary.LittleEndian
	case [2]byte{0xAB, 0xCD}:
		nativeEndian = binary.BigEndian
	default:
		return nil, DetermEndianErr
	}
	return nativeEndian, nil
}
