/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"context"
	"fmt"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetOptionalCharacterString is the corresponding interface of BACnetOptionalCharacterString
type BACnetOptionalCharacterString interface {
	BACnetOptionalCharacterStringContract
	BACnetOptionalCharacterStringRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// IsBACnetOptionalCharacterString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetOptionalCharacterString()
}

// BACnetOptionalCharacterStringContract provides a set of functions which can be overwritten by a sub struct
type BACnetOptionalCharacterStringContract interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// IsBACnetOptionalCharacterString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetOptionalCharacterString()
}

// BACnetOptionalCharacterStringRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetOptionalCharacterStringRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// _BACnetOptionalCharacterString is the data-structure of this message
type _BACnetOptionalCharacterString struct {
	_SubType        BACnetOptionalCharacterString
	PeekedTagHeader BACnetTagHeader
}

var _ BACnetOptionalCharacterStringContract = (*_BACnetOptionalCharacterString)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetOptionalCharacterString) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_BACnetOptionalCharacterString) GetPeekedTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetOptionalCharacterString factory function for _BACnetOptionalCharacterString
func NewBACnetOptionalCharacterString(peekedTagHeader BACnetTagHeader) *_BACnetOptionalCharacterString {
	return &_BACnetOptionalCharacterString{PeekedTagHeader: peekedTagHeader}
}

// Deprecated: use the interface for direct cast
func CastBACnetOptionalCharacterString(structType any) BACnetOptionalCharacterString {
	if casted, ok := structType.(BACnetOptionalCharacterString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetOptionalCharacterString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetOptionalCharacterString) GetTypeName() string {
	return "BACnetOptionalCharacterString"
}

func (m *_BACnetOptionalCharacterString) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetOptionalCharacterString) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetOptionalCharacterStringParse[T BACnetOptionalCharacterString](ctx context.Context, theBytes []byte) (T, error) {
	return BACnetOptionalCharacterStringParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetOptionalCharacterStringParseWithBufferProducer[T BACnetOptionalCharacterString]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetOptionalCharacterStringParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func BACnetOptionalCharacterStringParseWithBuffer[T BACnetOptionalCharacterString](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_BACnetOptionalCharacterString{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_BACnetOptionalCharacterString) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetOptionalCharacterString BACnetOptionalCharacterString, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetOptionalCharacterString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetOptionalCharacterString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	peekedTagHeader, err := ReadPeekField[BACnetTagHeader](ctx, "peekedTagHeader", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagHeader' field"))
	}
	m.PeekedTagHeader = peekedTagHeader

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetOptionalCharacterString
	switch {
	case peekedTagNumber == uint8(0): // BACnetOptionalCharacterStringNull
		if _child, err = (&_BACnetOptionalCharacterStringNull{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetOptionalCharacterStringNull for type-switch of BACnetOptionalCharacterString")
		}
	case 0 == 0: // BACnetOptionalCharacterStringValue
		if _child, err = (&_BACnetOptionalCharacterStringValue{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetOptionalCharacterStringValue for type-switch of BACnetOptionalCharacterString")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	if closeErr := readBuffer.CloseContext("BACnetOptionalCharacterString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetOptionalCharacterString")
	}

	return _child, nil
}

func (pm *_BACnetOptionalCharacterString) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetOptionalCharacterString, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetOptionalCharacterString"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetOptionalCharacterString")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetOptionalCharacterString"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetOptionalCharacterString")
	}
	return nil
}

func (m *_BACnetOptionalCharacterString) IsBACnetOptionalCharacterString() {}
