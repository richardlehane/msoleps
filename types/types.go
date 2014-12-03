package types

import (
	"encoding/binary"
	"errors"
)

var (
	ErrType        = errors.New("msoleps: error coercing byte stream to type")
	ErrUnknownType = errors.New("msoleps: unknown type error")
)

func Evaluate(b []byte) (Type, error) {
	if len(b) < 4 {
		return Integer8(0), ErrType
	}
	id := TypeID(binary.LittleEndian.Uint16(b[:2]))
	f, ok := MakeTypes[id]
	if !ok {
		return Integer8(0), ErrType
	}
	return f(b[4:])
}

type Type interface {
	String() string
	Type() string
}

type TypeID uint16

const (
	VT_EMPTY TypeID = iota // 0x0000
	VT_NULL
	VT_I2
	VT_I4
	VT_R4
	VT_R8
	VT_CY
	VT_DATE
	VT_BSTR
	_
	VT_ERROR
	VT_BOOL
	VT_VARIANT
	_
	VT_DECIMAL
	_
	VT_I1
	VT_U1
	VT_UI2
	VT_UI4
	VT_I8
	VT_UI8
	VT_INT
	VT_UINT  //0x0017
	_        = iota + 5
	VT_LPSTR //0x001E
	VT_LPWSTR
	VT_FILETIME = iota + 0x25 // 0x0040
	VT_BLOB
	VT_STREAM
	VT_STORAGE
	VT_STREAMED_OBJECT
	VT_STORED_OBJECT
	VT_BLOB_OBJECT
	VT_CF
	VT_CLSID
	VT_VERSIONED_STREAM // 0x0049
)

type MakeType func([]byte) (Type, error)

var MakeTypes map[TypeID]MakeType = map[TypeID]MakeType{
	VT_CY:       MakeCurrency,
	VT_DATE:     MakeDate,
	VT_DECIMAL:  MakeDecimal,
	VT_FILETIME: MakeFileTime,
	VT_CLSID:    MakeGuid,
}
