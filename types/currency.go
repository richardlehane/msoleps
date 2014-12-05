package types

import (
	"encoding/binary"
	"strconv"
)

//The CURRENCY type specifies currency information. It is represented as an 8-byte integer, scaled by 10,000, to give a fixed-point number with 15 digits to the left of the decimal point, and four digits to the right. This representation provides a range of 922337203685477.5807 to –922337203685477.5808. For example, $5.25 is stored as the value 52500.

type Currency int64

func (c Currency) String() string {
	return "$" + strconv.FormatFloat(float64(c)/10000, 'f', -1, 64)
}

func (c Currency) Type() string {
	return "Currency"
}

func MakeCurrency(b []byte) (Type, error) {
	if len(b) < 8 {
		return Currency(0), ErrType
	}
	return Currency(binary.LittleEndian.Uint64(b[:8])), nil
}
