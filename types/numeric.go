package types

import (
	"encoding/binary"
	"strconv"
)

type Integer8 byte

func (i Integer8) Type() string {
	return "Int8"
}

func (i Integer8) String() string {
	return strconv.Itoa(int(i))
}

func MakeInteger8(b []byte) (Type, error) {
	if len(b) < 1 {
		return Integer8(0), ErrType
	}
	return Integer8(b[0]), nil
}

type Integer16 int16

func (i Integer16) Type() string {
	return "Int16"
}

func (i Integer16) String() string {
	return strconv.Itoa(int(i))
}

func MakeInteger16(b []byte) (Type, error) {
	if len(b) < 2 {
		return Integer16(0), ErrType
	}
	return Integer16(binary.LittleEndian.Uint16(b[:2])), nil
}
