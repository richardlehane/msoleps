package msoleps

import "github.com/richardlehane/msoleps/types"

type Property struct {
	Name string
	T    types.Type
}

func (p *Property) String() string {
	return p.T.String()
}

func (p *Property) Type() string {
	return p.T.Type()
}

type propertySetStream struct {
	ByteOrder       uint16
	Version         uint16
	SystemID        uint32
	CLSID           types.Guid
	NumPropertySets uint32
	FmtidA          types.Guid
	OffsetA         uint32
	FmtidB          types.Guid
	OffsetB         uint32
}

type propertySet struct {
	size          uint32
	numProperties uint32
	idsOffs       []propertyIDandOffset
	dict          map[uint32]string
	code          types.CodePageID
}

type propertyIDandOffset struct {
	id     uint32
	offset uint32
}
