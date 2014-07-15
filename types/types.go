package types

type Type interface {
	String() string
	Type() string
}

const (
	VT_EMPTY uint16 = iota // 0x0000
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

type TypedProp struct {
	Val     uint16
	padding uint16
}
