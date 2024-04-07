// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
)

var (
	ErrNilDest          = errors.New("nil dest")
	ErrTypeNotSupported = errors.New("type is not supported")
)

// Marshal v into []byte. It returns an error when v is an invalid value.
func Marshal(v Value, bo binary.ByteOrder) (b []byte, err error) {
	if err := MarshalTo(&b, v, bo); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo is a zero-alloc marshal function that will marshal v into []byte and append it to given dest.
// It returns an error when v is an invalid value.
func MarshalTo(dest *[]byte, value Value, bo binary.ByteOrder) error {
	if dest == nil {
		return fmt.Errorf("dest could not be: %w", ErrNilDest)
	}

	switch value.Type() { // Fast path
	case TypeBool:
		var boolean byte
		if value.Bool() {
			boolean = 1
		}
		*dest = append(*dest, boolean)
		return nil
	case TypeSliceBool:
		val := value.SliceBool()
		for i := range val {
			if val[i] {
				*dest = append(*dest, 1)
			} else {
				*dest = append(*dest, 0)
			}
		}
		return nil
	case TypeInt8:
		*dest = append(*dest, uint8(value.Int8()))
		return nil
	case TypeSliceInt8:
		val := value.SliceInt8()
		for i := range val {
			*dest = append(*dest, uint8(val[i]))
		}
		return nil
	case TypeUint8:
		*dest = append(*dest, uint8(value.Uint8()))
		return nil
	case TypeSliceUint8:
		*dest = append(*dest, value.SliceUint8()...)
		return nil
	case TypeInt16:
		cur := len(*dest)
		*dest = append(*dest, 0, 0)
		bo.PutUint16((*dest)[cur:], uint16(value.Int16()))
		return nil
	case TypeSliceUint16:
		val := value.SliceUint16()
		const n = 2
		cur := len(*dest)
		for i := range val {
			*dest = append(*dest, 0, 0)
			bo.PutUint16((*dest)[cur:cur+n], uint16(val[i]))
			cur += n
		}
		return nil
	case TypeUint16:
		cur := len(*dest)
		*dest = append(*dest, 0, 0)
		bo.PutUint16((*dest)[cur:], value.Uint16())
		return nil
	case TypeSliceInt16:
		val := value.SliceInt16()
		const n = 2
		cur := len(*dest)
		for i := range val {
			*dest = append(*dest, 0, 0)
			bo.PutUint16((*dest)[cur:cur+n], uint16(val[i]))
			cur += n
		}
		return nil
	case TypeInt32:
		cur := len(*dest)
		*dest = append(*dest, 0, 0, 0, 0)
		bo.PutUint32((*dest)[cur:], uint32(value.Int32()))
		return nil
	case TypeSliceInt32:
		val := value.SliceInt32()
		const n = 4
		cur := len(*dest)
		for i := range val {
			*dest = append(*dest, 0, 0, 0, 0)
			bo.PutUint32((*dest)[cur:cur+n], uint32(val[i]))
			cur += n
		}
		return nil
	case TypeUint32:
		cur := len(*dest)
		*dest = append(*dest, 0, 0, 0, 0)
		bo.PutUint32((*dest)[cur:], value.Uint32())
		return nil
	case TypeSliceUint32:
		val := value.SliceUint32()
		const n = 4
		cur := len(*dest)
		for i := range val {
			*dest = append(*dest, 0, 0, 0, 0)
			bo.PutUint32((*dest)[cur:cur+n], val[i])
			cur += n
		}
		return nil
	case TypeInt64:
		cur := len(*dest)
		*dest = append(*dest, 0, 0, 0, 0, 0, 0, 0, 0)
		bo.PutUint64((*dest)[cur:], uint64(value.Int64()))
		return nil
	case TypeSliceInt64:
		val := value.SliceInt64()
		const n = 8
		cur := len(*dest)
		for i := range val {
			*dest = append(*dest, 0, 0, 0, 0, 0, 0, 0, 0)
			bo.PutUint64((*dest)[cur:cur+n], uint64(val[i]))
			cur += n
		}
		return nil
	case TypeUint64:
		cur := len(*dest)
		*dest = append(*dest, 0, 0, 0, 0, 0, 0, 0, 0)
		bo.PutUint64((*dest)[cur:], value.Uint64())
		return nil
	case TypeSliceUint64:
		val := value.SliceUint64()
		const n = 8
		cur := len(*dest)
		for i := range val {
			*dest = append(*dest, 0, 0, 0, 0, 0, 0, 0, 0)
			bo.PutUint64((*dest)[cur:cur+n], val[i])
			cur += n
		}
		return nil
	case TypeFloat32:
		cur := len(*dest)
		v := math.Float32bits(value.Float32())
		*dest = append(*dest, 0, 0, 0, 0)
		bo.PutUint32((*dest)[cur:], v)
		return nil
	case TypeSliceFloat32:
		val := value.SliceFloat32()
		const n = 4
		cur := len(*dest)
		for i := range val {
			*dest = append(*dest, 0, 0, 0, 0)
			bo.PutUint32((*dest)[cur:cur+n], math.Float32bits(val[i]))
			cur += n
		}
		return nil
	case TypeFloat64:
		v := math.Float64bits(value.Float64())
		cur := len(*dest)
		*dest = append(*dest, 0, 0, 0, 0, 0, 0, 0, 0)
		bo.PutUint64((*dest)[cur:], v)
		return nil
	case TypeSliceFloat64:
		val := value.SliceFloat64()
		const n = 8
		cur := len(*dest)
		for i := range val {
			*dest = append(*dest, 0, 0, 0, 0, 0, 0, 0, 0)
			bo.PutUint64((*dest)[cur:cur+n], math.Float64bits(val[i]))
			cur += n
		}
		return nil
	case TypeString:
		val := value.String()
		if len(val) == 0 {
			*dest = append(*dest, 0x00)
			return nil
		}
		*dest = append(*dest, val...)
		if val[len(val)-1] != '\x00' {
			*dest = append(*dest, '\x00') // add utf-8 null-terminated string
		}
		return nil
	case TypeSliceString:
		val := value.SliceString()
		for i := range val {
			if len(val[i]) == 0 {
				*dest = append(*dest, '\x00')
				continue
			}
			*dest = append(*dest, val[i]...)
			if val[i][len(val[i])-1] != '\x00' {
				*dest = append(*dest, '\x00')
			}
		}
		return nil
	default:
		return fmt.Errorf("type Value(%T) is not supported: %w", value.Type(), ErrTypeNotSupported)
	}
}