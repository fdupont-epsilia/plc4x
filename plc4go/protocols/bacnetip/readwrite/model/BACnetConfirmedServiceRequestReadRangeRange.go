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

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConfirmedServiceRequestReadRangeRange is the corresponding interface of BACnetConfirmedServiceRequestReadRangeRange
type BACnetConfirmedServiceRequestReadRangeRange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetConfirmedServiceRequestReadRangeRangeExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestReadRangeRange.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestReadRangeRangeExactly interface {
	BACnetConfirmedServiceRequestReadRangeRange
	isBACnetConfirmedServiceRequestReadRangeRange() bool
}

// _BACnetConfirmedServiceRequestReadRangeRange is the data-structure of this message
type _BACnetConfirmedServiceRequestReadRangeRange struct {
	_BACnetConfirmedServiceRequestReadRangeRangeChildRequirements
	PeekedTagHeader BACnetTagHeader
	OpeningTag      BACnetOpeningTag
	ClosingTag      BACnetClosingTag
}

type _BACnetConfirmedServiceRequestReadRangeRangeChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetPeekedTagNumber() uint8
}

type BACnetConfirmedServiceRequestReadRangeRangeParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetConfirmedServiceRequestReadRangeRange, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetConfirmedServiceRequestReadRangeRangeChild interface {
	utils.Serializable
	InitializeParent(parent BACnetConfirmedServiceRequestReadRangeRange, peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, closingTag BACnetClosingTag)
	GetParent() *BACnetConfirmedServiceRequestReadRangeRange

	GetTypeName() string
	BACnetConfirmedServiceRequestReadRangeRange
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestReadRangeRange) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *_BACnetConfirmedServiceRequestReadRangeRange) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetConfirmedServiceRequestReadRangeRange) GetClosingTag() BACnetClosingTag {
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

func (m *_BACnetConfirmedServiceRequestReadRangeRange) GetPeekedTagNumber() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestReadRangeRange factory function for _BACnetConfirmedServiceRequestReadRangeRange
func NewBACnetConfirmedServiceRequestReadRangeRange(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, closingTag BACnetClosingTag) *_BACnetConfirmedServiceRequestReadRangeRange {
	return &_BACnetConfirmedServiceRequestReadRangeRange{PeekedTagHeader: peekedTagHeader, OpeningTag: openingTag, ClosingTag: closingTag}
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestReadRangeRange(structType any) BACnetConfirmedServiceRequestReadRangeRange {
	if casted, ok := structType.(BACnetConfirmedServiceRequestReadRangeRange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestReadRangeRange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestReadRangeRange) GetTypeName() string {
	return "BACnetConfirmedServiceRequestReadRangeRange"
}

func (m *_BACnetConfirmedServiceRequestReadRangeRange) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestReadRangeRange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestReadRangeRangeParse(ctx context.Context, theBytes []byte) (BACnetConfirmedServiceRequestReadRangeRange, error) {
	return BACnetConfirmedServiceRequestReadRangeRangeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetConfirmedServiceRequestReadRangeRangeParseWithBufferProducer[T BACnetConfirmedServiceRequestReadRangeRange]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		buffer, err := BACnetConfirmedServiceRequestReadRangeRangeParseWithBuffer(ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return buffer.(T), err
	}
}

func BACnetConfirmedServiceRequestReadRangeRangeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestReadRangeRange, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestReadRangeRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestReadRangeRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (peekedTagHeader)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for peekedTagHeader")
	}
	peekedTagHeader, _ := BACnetTagHeaderParseWithBuffer(ctx, readBuffer)
	readBuffer.Reset(currentPos)

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(peekedTagHeader.GetActualTagNumber())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetConfirmedServiceRequestReadRangeRangeChildSerializeRequirement interface {
		BACnetConfirmedServiceRequestReadRangeRange
		InitializeParent(BACnetConfirmedServiceRequestReadRangeRange, BACnetTagHeader, BACnetOpeningTag, BACnetClosingTag)
		GetParent() BACnetConfirmedServiceRequestReadRangeRange
	}
	var _childTemp any
	var _child BACnetConfirmedServiceRequestReadRangeRangeChildSerializeRequirement
	var typeSwitchError error
	switch {
	case peekedTagNumber == 0x3: // BACnetConfirmedServiceRequestReadRangeRangeByPosition
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestReadRangeRangeByPositionParseWithBuffer(ctx, readBuffer)
	case peekedTagNumber == 0x6: // BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberParseWithBuffer(ctx, readBuffer)
	case peekedTagNumber == 0x7: // BACnetConfirmedServiceRequestReadRangeRangeByTime
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestReadRangeRangeByTimeParseWithBuffer(ctx, readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetConfirmedServiceRequestReadRangeRange")
	}
	_child = _childTemp.(BACnetConfirmedServiceRequestReadRangeRangeChildSerializeRequirement)

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(peekedTagHeader.GetActualTagNumber())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestReadRangeRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestReadRangeRange")
	}

	// Finish initializing
	_child.InitializeParent(_child, peekedTagHeader, openingTag, closingTag)
	return _child, nil
}

func (pm *_BACnetConfirmedServiceRequestReadRangeRange) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetConfirmedServiceRequestReadRangeRange, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestReadRangeRange"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestReadRangeRange")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(ctx, m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
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

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(ctx, m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestReadRangeRange"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestReadRangeRange")
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestReadRangeRange) isBACnetConfirmedServiceRequestReadRangeRange() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestReadRangeRange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
