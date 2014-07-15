package types

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"strings"
)

// Win GUID and UUID type
// http://msdn.microsoft.com/en-us/library/cc230326.aspx
type Guid struct {
	DataA uint32
	DataB uint16
	DataC uint16
	DataD [8]byte
}

func (g Guid) String() string {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint32(buf[:4], g.DataA)
	binary.BigEndian.PutUint16(buf[4:6], g.DataB)
	binary.BigEndian.PutUint16(buf[6:], g.DataC)
	return strings.ToUpper("{" +
		hex.EncodeToString(buf[:4]) +
		"-" +
		hex.EncodeToString(buf[4:6]) +
		"-" +
		hex.EncodeToString(buf[6:]) +
		"-" +
		hex.EncodeToString(g.DataD[:2]) +
		"-" +
		hex.EncodeToString(g.DataD[2:]) +
		"}")
}

func GuidFromString(str string) (Guid, error) {
	gerr := "Invalid GUID: expecting in format {F29F85E0-4FF9-1068-AB91-08002B27B3D9}, got " + str
	if len(str) != 38 {
		return Guid{}, errors.New(gerr + "; bad length, should be 38 chars")
	}
	trimmed := strings.Trim(str, "{}")
	parts := strings.Split(trimmed, "-")
	if len(parts) != 5 {
		return Guid{}, errors.New(gerr + "; expecting should five '-' separators")
	}
	buf, err := hex.DecodeString(strings.Join(parts, ""))
	if err != nil {
		return Guid{}, errors.New(gerr + "; error decoding hex: " + err.Error())
	}
	g := Guid{
		binary.BigEndian.Uint32(buf[:4]),
		binary.BigEndian.Uint16(buf[4:6]),
		binary.BigEndian.Uint16(buf[6:8]),
		[8]byte{},
	}
	copy(g.DataD[:], buf[8:])
	return g, nil
}

func MustGuidFromString(str string) Guid {
	g, err := GuidFromString(str)
	if err != nil {
		panic(err)
	}
	return g
}
