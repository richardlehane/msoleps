package msoleps

import (
	"errors"
	"io"

	"github.com/richardlehane/msoleps/types"
)

var (
	ErrFormat = errors.New("msoleps: not a valid msoleps stream")
	ErrRead   = errors.New("msoleps: error reading msoleps stream")
)

// check the first uint16 of an MSCFB name to see if this is a MSOLEPS stream
func IsMSOLEPS(i uint16) bool {
	if i == 0x0005 {
		return true
	}
	return false
}

type Reader struct {
	r io.Reader
	// indexes
	i    int // place in stream
	psi  int // place in Property Set
	ps   int // index of property set
	prop int // index of current property
	// common data
	dict map[uint32]string
	code types.CodePageID
	// loaded data
	pSetStream *PropertySetStream
	pSet       *PropertySet
	idsOffs    []PropertyIdentifierAndOffset
}

func New(rdr io.Reader) (*Reader, error) {
	r := &Reader{r: rdr}
	r.pSetStream = &PropertySetStream{}
	if err := binary.Read(r.r, binary.LittleEndian, r.pSetStream); err != nil {
		return nil, ErrRead
	}
	switch {
	case r.pSetStream.ByteOrder != 0xFFFE, r.pSetStream.Version > 0x0001, r.pSetStream.NumPropertySets > 0x00000002:
		return nil, ErrFormat
	}
	r.i = 17 * 4
	if err := r.setIdsOffs(); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *Reader) setIdsOffs() error {
	r.pSet = &PropertySet{}
	if err := binary.Read(r.r, binary.LittleEndian, r.pSet); err != nil {
		return ErrRead
	}
	r.idsOffs = make([]PropertyIdentifierAndOffset, int(r.pSet.NumProperties))
	if err := binary.Read(r.r, binary.LittleEndian, r.idsOffs); err != nil {
		return ErrRead
	}
	r.psi = len(r.idsOffs)*8 + 8
	r.i += r.psi
	return nil
}

func (r *Reader) Read() (*Property, error) {
	if r.prop >= len(r.idsOffs) {
		if int(r.pSetStream.NumPropertySets)-1 > r.ps {

		}
	}
}
