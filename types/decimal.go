package types

import (
	"encoding/binary"
	"math"
	"math/big"
)

// http://msdn.microsoft.com/en-us/library/cc237603.aspx
type Decimal struct {
	res    [2]byte
	scale  byte
	sign   byte
	high32 uint32
	low64  uint64
}

func (d Decimal) Type() string {
	return "Decimal"
}

func (d Decimal) String() string {
	h, l, b := new(big.Int), new(big.Int), new(big.Int)
	l.SetUint64(d.low64)
	h.Lsh(big.NewInt(int64(d.high32)), 64)
	b.Add(h, l)
	q, f, r := new(big.Rat), new(big.Rat), new(big.Rat)
	q.SetFloat64(math.Pow10(int(d.scale)))
	r.Quo(f.SetInt(b), q)
	if d.sign == 0x80 {
		r.Neg(r)
	}
	return r.FloatString(20)
}

func MakeDecimal(b []byte) (Type, error) {
	if len(b) < 16 {
		return Decimal{}, ErrType
	}
	return Decimal{
		res:    [2]byte{b[0], b[1]},
		scale:  b[2],
		sign:   b[3],
		high32: binary.LittleEndian.Uint32(b[4:8]),
		low64:  binary.LittleEndian.Uint64(b[8:16]),
	}, nil
}
