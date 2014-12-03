package types

import "strconv"

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
