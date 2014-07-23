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
	b *bytes.Buffer
	r *bytes.Reader
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
	r := &Reader{}
	r.b = &bytes.Buffer{}
	return r, r.start(rdr)
}

func (r *Reader) Reset(rdr io.Reader) error {
	r.b.Reset()
	return r.start(rdr)
}

func (r *Reader) start(rdr io.Reader) error {
	if _, err := r.b.ReadFrom(rdr); err != nil {
		return ErrRead
	}
	r.r = bytes.NewReader(r.b.Bytes())
	// read the header (property stream details)
	r.pSetStream = &PropertySetStream{}
	if err := binary.Read(r.r, binary.LittleEndian, r.pSetStream); err != nil {
		return ErrRead
	}
	// sanity checks to find obvious errors
	switch {
	case r.pSetStream.ByteOrder != 0xFFFE, r.pSetStream.Version > 0x0001, r.pSetStream.NumPropertySets > 0x00000002:
		return ErrFormat
	}
	r.i = 17 * 4
	// identify the property identifiers and offsets
	if err := r.setIdsOffs(); err != nil {
		return err
	}
	return nil
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
	// set dictionary and codepage if present
	var dictIdx, codeIdx int
	var dictPres, codePres bool
	for i, v := range r.idsOffs {
		if v.PropertyIdentifier == 0x00000000 {
			dictIdx = i
			dictPres = true
		}
		if v.PropertyIdentifier == 0x00000001 {
			codeIdx = i
			codePres = true
		}
	}
	if codePres {
		codeIdx++ // just letting it compile
	}
	if dictPres {
		dictIdx++ // just letting it compile - unfinished bit
	}
	// increment the read index by length of the IdsOffs slice
	r.psi = len(r.idsOffs)*8 + 8
	r.i += r.psi
	return nil
}

func (r *Reader) Read() (*Property, error) {
	// check if we have we reached the end of our list of properties
	if r.prop >= len(r.idsOffs) {
		// check if there is a second property set in the stream
		if int(r.pSetStream.NumPropertySets)-1 > r.ps {
			r.ps++
			r.prop = 0
			r.i = int(r.pSetStream.OffsetB)
			if err := r.setIdsOffs(); err != nil {
				return nil, err
			}
		} else {
			return nil, io.EOF
		}
	}

	return nil, nil
}
