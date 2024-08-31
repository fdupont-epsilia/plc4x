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

// BACnetConstructedDataIntervalOffset is the corresponding interface of BACnetConstructedDataIntervalOffset
type BACnetConstructedDataIntervalOffset interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetIntervalOffset returns IntervalOffset (property field)
	GetIntervalOffset() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataIntervalOffsetExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataIntervalOffset.
// This is useful for switch cases.
type BACnetConstructedDataIntervalOffsetExactly interface {
	BACnetConstructedDataIntervalOffset
	isBACnetConstructedDataIntervalOffset() bool
}

// _BACnetConstructedDataIntervalOffset is the data-structure of this message
type _BACnetConstructedDataIntervalOffset struct {
	*_BACnetConstructedData
	IntervalOffset BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIntervalOffset) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataIntervalOffset) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_INTERVAL_OFFSET
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIntervalOffset) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataIntervalOffset) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIntervalOffset) GetIntervalOffset() BACnetApplicationTagUnsignedInteger {
	return m.IntervalOffset
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIntervalOffset) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetIntervalOffset())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataIntervalOffset factory function for _BACnetConstructedDataIntervalOffset
func NewBACnetConstructedDataIntervalOffset(intervalOffset BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIntervalOffset {
	_result := &_BACnetConstructedDataIntervalOffset{
		IntervalOffset:         intervalOffset,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIntervalOffset(structType any) BACnetConstructedDataIntervalOffset {
	if casted, ok := structType.(BACnetConstructedDataIntervalOffset); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIntervalOffset); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIntervalOffset) GetTypeName() string {
	return "BACnetConstructedDataIntervalOffset"
}

func (m *_BACnetConstructedDataIntervalOffset) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (intervalOffset)
	lengthInBits += m.IntervalOffset.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIntervalOffset) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataIntervalOffsetParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataIntervalOffset, error) {
	return BACnetConstructedDataIntervalOffsetParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataIntervalOffsetParseWithBufferProducer(tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataIntervalOffset, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataIntervalOffset, error) {
		return BACnetConstructedDataIntervalOffsetParseWithBuffer(ctx, readBuffer, tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
	}
}

func BACnetConstructedDataIntervalOffsetParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataIntervalOffset, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIntervalOffset"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIntervalOffset")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	intervalOffset, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "intervalOffset", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer(), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'intervalOffset' field"))
	}

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), intervalOffset)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIntervalOffset"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIntervalOffset")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataIntervalOffset{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		IntervalOffset: intervalOffset,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataIntervalOffset) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIntervalOffset) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIntervalOffset"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIntervalOffset")
		}

		// Simple Field (intervalOffset)
		if pushErr := writeBuffer.PushContext("intervalOffset"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for intervalOffset")
		}
		_intervalOffsetErr := writeBuffer.WriteSerializable(ctx, m.GetIntervalOffset())
		if popErr := writeBuffer.PopContext("intervalOffset"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for intervalOffset")
		}
		if _intervalOffsetErr != nil {
			return errors.Wrap(_intervalOffsetErr, "Error serializing 'intervalOffset' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIntervalOffset"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIntervalOffset")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIntervalOffset) isBACnetConstructedDataIntervalOffset() bool {
	return true
}

func (m *_BACnetConstructedDataIntervalOffset) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
