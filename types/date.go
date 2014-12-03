package types

import (
	"encoding/binary"
	"time"
)

// http://msdn.microsoft.com/en-us/library/cc237601.aspx
type Date float64

func (d Date) Time() time.Time {
	start := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC)
	day := float64(time.Hour * 24)
	dur := time.Duration(day * float64(d))
	return start.Add(dur)
}

func (d Date) String() string {
	return d.Time().String()
}

func (d Date) Type() string {
	return "Date"
}

func MakeDate(b []byte) (Type, error) {
	if len(b) < 8 {
		return Date(0), ErrType
	}
	return Date(binary.LittleEndian.Uint64(b[:8])), nil
}
