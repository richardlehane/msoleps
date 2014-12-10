// Copyright 2014 Richard Lehane. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

func (g Guid) Type() string {
	return "Guid"
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
	return makeGuid(buf, binary.BigEndian), nil
}

func makeGuid(b []byte, order binary.ByteOrder) Guid {
	g := Guid{
		DataA: order.Uint32(b[:4]),
		DataB: order.Uint16(b[4:6]),
		DataC: order.Uint16(b[6:8]),
		DataD: [8]byte{},
	}
	copy(g.DataD[:], b[8:])
	return g
}

func MustGuidFromString(str string) Guid {
	g, err := GuidFromString(str)
	if err != nil {
		panic(err)
	}
	return g
}

func MakeGuid(b []byte) (Type, error) {
	if len(b) < 16 {
		return Guid{}, ErrType
	}
	return makeGuid(b, binary.LittleEndian), nil
}

func MustGuid(b []byte) Guid {
	return makeGuid(b, binary.LittleEndian)
}
