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

// BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass is the corresponding interface of BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
type BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass interface {
	BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract
	BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract provides a set of functions which can be overwritten by a sub struct
type BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract interface {
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassRequirements interface {
}

// BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassExactly interface {
	BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
	isBACnetConfirmedServiceRequestConfirmedTextMessageMessageClass() bool
}

// _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass is the data-structure of this message
type _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass struct {
	_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassChildRequirements
	OpeningTag      BACnetOpeningTag
	PeekedTagHeader BACnetTagHeader
	ClosingTag      BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract = (*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass)(nil)

type _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetPeekedTagNumber() uint8
}

type BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassChild interface {
	utils.Serializable

	GetParent() *BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass

	GetTypeName() string
	BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass) GetPeekedTagNumber() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClass factory function for _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
func NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClass(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass {
	return &_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass{OpeningTag: openingTag, PeekedTagHeader: peekedTagHeader, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestConfirmedTextMessageMessageClass(structType any) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass {
	if casted, ok := structType.(BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass) GetTypeName() string {
	return "BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass"
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassParse[T BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass](ctx context.Context, theBytes []byte, tagNumber uint8) (T, error) {
	return BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassParseWithBufferProducer[T BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass](tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassParseWithBuffer[T](ctx, readBuffer, tagNumber)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassParseWithBuffer[T BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass](ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (T, error) {
	v, err := (&_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass{}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (_bACnetConfirmedServiceRequestConfirmedTextMessageMessageClass BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

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
	var _child BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
	switch {
	case peekedTagNumber == uint8(0): // BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric
		if _child, err = (&_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric{}).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric for type-switch of BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass")
		}
	case peekedTagNumber == uint8(1): // BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter
		if _child, err = (&_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter{}).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter for type-switch of BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass")
	}

	return _child, nil
}

func (pm *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
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

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass) isBACnetConfirmedServiceRequestConfirmedTextMessageMessageClass() bool {
	return true
}
