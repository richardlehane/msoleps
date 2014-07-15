package msoleps

import "github.com/richardlehane/msoleps/types"

type PropertySetStream struct {
	ByteOrder        uint16
	Version          uint16
	SystemIdentifier uint32
	CLSID            types.Guid
	NumPropertySets  uint32
	FmtidA           types.Guid
	OffsetA          uint32
	FmtidB           types.Guid
	OffsetB          uint32
}

type PropertySet struct {
	Size          uint32
	NumProperties uint32
}

type PropertyIdentifierAndOffset struct {
	PropertyIdentifier uint32
	Offset             uint32
}

type Property struct {
	name string
	sz   int
	raw  []byte
	types.Type
}

func (p *Property) FMTID() string {
	return p.fmtid.String()
}

func (p *Property) Name() string {
	return p.name
}

func (p *Property) Bytes() []byte {
	return p.raw[:p.sz]
}
