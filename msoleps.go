package msoleps

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"

	"github.com/richardlehane/msoleps/types"
)

var (
	ErrFormat = errors.New("msoleps: not a valid msoleps stream")
	ErrRead   = errors.New("msoleps: error reading msoleps stream")
	ErrSeek   = errors.New("msoleps: can't seek backwards")
)

// check the first uint16 of an MSCFB name to see if this is a MSOLEPS stream
func IsMSOLEPS(i uint16) bool {
	if i == 0x0005 {
		return true
	}
	return false
}

type Reader struct {
	Property []*Property
	CLSID    types.Guid
	SystemID uint32

	b          *bytes.Buffer
	buf        []byte
	pSetStream *propertySetStream
	pSets      [2]*propertySet
}

func New() *Reader {
	r := &Reader{}
	r.b = &bytes.Buffer{}
	return r
}

func (r *Reader) Reset(rdr io.Reader) error {
	r.b.Reset()
	return r.start(rdr)
}

func NewFrom(rdr io.Reader) (*Reader, error) {
	r := &Reader{}
	r.b = &bytes.Buffer{}
	return r, r.start(rdr)
}

func (r *Reader) start(rdr io.Reader) error {
	if _, err := r.b.ReadFrom(rdr); err != nil {
		return ErrRead
	}
	r.buf = r.b.Bytes()
	// read the header (property stream details)
	r.pSetStream = &propertySetStream{}
	if err := binary.Read(r.b, binary.LittleEndian, r.pSetStream); err != nil {
		return ErrRead
	}
	// sanity checks to find obvious errors
	switch {
	case r.pSetStream.ByteOrder != 0xFFFE, r.pSetStream.Version > 0x0001, r.pSetStream.NumPropertySets > 0x00000002:
		return ErrFormat
	}
	// identify the property identifiers and offsets
	ps, err := r.getPropertySet(r.pSetStream.OffsetA)
	if err != nil {
		return err
	}
	plen := len(ps.idsOffs)
	r.pSets[0] = ps
	if r.pSetStream.NumPropertySets == 2 {
		psb, err := r.getPropertySet(r.pSetStream.OffsetB)
		if err != nil {
			return err
		}
		r.pSets[1] = psb
		plen += len(psb.idsOffs)
	}
	r.Property = make([]*Property, plen)
	return nil
}

func (r *Reader) getPropertySet(o uint32) (*propertySet, error) {
	pSet := &propertySet{}
	pSet.size = binary.LittleEndian.Uint32(r.buf[int(o) : int(o)+4])
	pSet.numProperties = binary.LittleEndian.Uint32(r.buf[int(o)+4 : int(o)+8])
	pSet.idsOffs = make([]propertyIDandOffset, int(pSet.numProperties))
	var dictOff uint32
	for i := range pSet.idsOffs {
		this := i*8 + 8 + int(o)
		pSet.idsOffs[i].id = binary.LittleEndian.Uint32(r.buf[this : this+4])
		pSet.idsOffs[i].offset = binary.LittleEndian.Uint32(r.buf[this+4 : this+8])
		switch pSet.idsOffs[i].id {
		case 0x00000000:
			dictOff = pSet.idsOffs[i].offset
		case 0x00000001:
			off := int(pSet.idsOffs[i].offset + o)
			pSet.code = types.CodePageID(binary.LittleEndian.Uint16(r.buf[off : off+2]))
		}
	}
	if dictOff > 0 {
		dictOff++ // just letting it compile - unfinished bit
	}
	return pSet, nil
}
