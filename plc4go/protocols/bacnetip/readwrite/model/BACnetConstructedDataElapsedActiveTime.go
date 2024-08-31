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

// BACnetConstructedDataElapsedActiveTime is the corresponding interface of BACnetConstructedDataElapsedActiveTime
type BACnetConstructedDataElapsedActiveTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetElapsedActiveTime returns ElapsedActiveTime (property field)
	GetElapsedActiveTime() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataElapsedActiveTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataElapsedActiveTime.
// This is useful for switch cases.
type BACnetConstructedDataElapsedActiveTimeExactly interface {
	BACnetConstructedDataElapsedActiveTime
	isBACnetConstructedDataElapsedActiveTime() bool
}

// _BACnetConstructedDataElapsedActiveTime is the data-structure of this message
type _BACnetConstructedDataElapsedActiveTime struct {
	*_BACnetConstructedData
	ElapsedActiveTime BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataElapsedActiveTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataElapsedActiveTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ELAPSED_ACTIVE_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataElapsedActiveTime) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataElapsedActiveTime) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataElapsedActiveTime) GetElapsedActiveTime() BACnetApplicationTagUnsignedInteger {
	return m.ElapsedActiveTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataElapsedActiveTime) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetElapsedActiveTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataElapsedActiveTime factory function for _BACnetConstructedDataElapsedActiveTime
func NewBACnetConstructedDataElapsedActiveTime(elapsedActiveTime BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataElapsedActiveTime {
	_result := &_BACnetConstructedDataElapsedActiveTime{
		ElapsedActiveTime:      elapsedActiveTime,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataElapsedActiveTime(structType any) BACnetConstructedDataElapsedActiveTime {
	if casted, ok := structType.(BACnetConstructedDataElapsedActiveTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataElapsedActiveTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataElapsedActiveTime) GetTypeName() string {
	return "BACnetConstructedDataElapsedActiveTime"
}

func (m *_BACnetConstructedDataElapsedActiveTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (elapsedActiveTime)
	lengthInBits += m.ElapsedActiveTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataElapsedActiveTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataElapsedActiveTimeParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataElapsedActiveTime, error) {
	return BACnetConstructedDataElapsedActiveTimeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataElapsedActiveTimeParseWithBufferProducer(tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataElapsedActiveTime, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataElapsedActiveTime, error) {
		return BACnetConstructedDataElapsedActiveTimeParseWithBuffer(ctx, readBuffer, tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
	}
}

func BACnetConstructedDataElapsedActiveTimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataElapsedActiveTime, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataElapsedActiveTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataElapsedActiveTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	elapsedActiveTime, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "elapsedActiveTime", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer(), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'elapsedActiveTime' field"))
	}

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), elapsedActiveTime)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataElapsedActiveTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataElapsedActiveTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataElapsedActiveTime{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ElapsedActiveTime: elapsedActiveTime,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataElapsedActiveTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataElapsedActiveTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataElapsedActiveTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataElapsedActiveTime")
		}

		// Simple Field (elapsedActiveTime)
		if pushErr := writeBuffer.PushContext("elapsedActiveTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for elapsedActiveTime")
		}
		_elapsedActiveTimeErr := writeBuffer.WriteSerializable(ctx, m.GetElapsedActiveTime())
		if popErr := writeBuffer.PopContext("elapsedActiveTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for elapsedActiveTime")
		}
		if _elapsedActiveTimeErr != nil {
			return errors.Wrap(_elapsedActiveTimeErr, "Error serializing 'elapsedActiveTime' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataElapsedActiveTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataElapsedActiveTime")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataElapsedActiveTime) isBACnetConstructedDataElapsedActiveTime() bool {
	return true
}

func (m *_BACnetConstructedDataElapsedActiveTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
