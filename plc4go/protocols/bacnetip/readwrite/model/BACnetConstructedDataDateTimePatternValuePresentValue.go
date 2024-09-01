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

// BACnetConstructedDataDateTimePatternValuePresentValue is the corresponding interface of BACnetConstructedDataDateTimePatternValuePresentValue
type BACnetConstructedDataDateTimePatternValuePresentValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() BACnetDateTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDateTime
}

// BACnetConstructedDataDateTimePatternValuePresentValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDateTimePatternValuePresentValue.
// This is useful for switch cases.
type BACnetConstructedDataDateTimePatternValuePresentValueExactly interface {
	BACnetConstructedDataDateTimePatternValuePresentValue
	isBACnetConstructedDataDateTimePatternValuePresentValue() bool
}

// _BACnetConstructedDataDateTimePatternValuePresentValue is the data-structure of this message
type _BACnetConstructedDataDateTimePatternValuePresentValue struct {
	BACnetConstructedDataContract
	PresentValue BACnetDateTime
}

var _ BACnetConstructedDataDateTimePatternValuePresentValue = (*_BACnetConstructedDataDateTimePatternValuePresentValue)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDateTimePatternValuePresentValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_DATETIMEPATTERN_VALUE
}

func (m *_BACnetConstructedDataDateTimePatternValuePresentValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRESENT_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDateTimePatternValuePresentValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDateTimePatternValuePresentValue) GetPresentValue() BACnetDateTime {
	return m.PresentValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDateTimePatternValuePresentValue) GetActualValue() BACnetDateTime {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDateTime(m.GetPresentValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDateTimePatternValuePresentValue factory function for _BACnetConstructedDataDateTimePatternValuePresentValue
func NewBACnetConstructedDataDateTimePatternValuePresentValue(presentValue BACnetDateTime, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDateTimePatternValuePresentValue {
	_result := &_BACnetConstructedDataDateTimePatternValuePresentValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		PresentValue:                  presentValue,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDateTimePatternValuePresentValue(structType any) BACnetConstructedDataDateTimePatternValuePresentValue {
	if casted, ok := structType.(BACnetConstructedDataDateTimePatternValuePresentValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDateTimePatternValuePresentValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDateTimePatternValuePresentValue) GetTypeName() string {
	return "BACnetConstructedDataDateTimePatternValuePresentValue"
}

func (m *_BACnetConstructedDataDateTimePatternValuePresentValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDateTimePatternValuePresentValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDateTimePatternValuePresentValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (_bACnetConstructedDataDateTimePatternValuePresentValue BACnetConstructedDataDateTimePatternValuePresentValue, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDateTimePatternValuePresentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDateTimePatternValuePresentValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	presentValue, err := ReadSimpleField[BACnetDateTime](ctx, "presentValue", ReadComplex[BACnetDateTime](BACnetDateTimeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'presentValue' field"))
	}
	m.PresentValue = presentValue

	actualValue, err := ReadVirtualField[BACnetDateTime](ctx, "actualValue", (*BACnetDateTime)(nil), presentValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDateTimePatternValuePresentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDateTimePatternValuePresentValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDateTimePatternValuePresentValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDateTimePatternValuePresentValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDateTimePatternValuePresentValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDateTimePatternValuePresentValue")
		}

		if err := WriteSimpleField[BACnetDateTime](ctx, "presentValue", m.GetPresentValue(), WriteComplex[BACnetDateTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'presentValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDateTimePatternValuePresentValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDateTimePatternValuePresentValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDateTimePatternValuePresentValue) isBACnetConstructedDataDateTimePatternValuePresentValue() bool {
	return true
}

func (m *_BACnetConstructedDataDateTimePatternValuePresentValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
