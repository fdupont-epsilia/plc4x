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

// BACnetConstructedDataAccessZoneAdjustValue is the corresponding interface of BACnetConstructedDataAccessZoneAdjustValue
type BACnetConstructedDataAccessZoneAdjustValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAdjustValue returns AdjustValue (property field)
	GetAdjustValue() BACnetApplicationTagSignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagSignedInteger
	// IsBACnetConstructedDataAccessZoneAdjustValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAccessZoneAdjustValue()
}

// _BACnetConstructedDataAccessZoneAdjustValue is the data-structure of this message
type _BACnetConstructedDataAccessZoneAdjustValue struct {
	BACnetConstructedDataContract
	AdjustValue BACnetApplicationTagSignedInteger
}

var _ BACnetConstructedDataAccessZoneAdjustValue = (*_BACnetConstructedDataAccessZoneAdjustValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAccessZoneAdjustValue)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccessZoneAdjustValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ACCESS_ZONE
}

func (m *_BACnetConstructedDataAccessZoneAdjustValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ADJUST_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccessZoneAdjustValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccessZoneAdjustValue) GetAdjustValue() BACnetApplicationTagSignedInteger {
	return m.AdjustValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAccessZoneAdjustValue) GetActualValue() BACnetApplicationTagSignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagSignedInteger(m.GetAdjustValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAccessZoneAdjustValue factory function for _BACnetConstructedDataAccessZoneAdjustValue
func NewBACnetConstructedDataAccessZoneAdjustValue(adjustValue BACnetApplicationTagSignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccessZoneAdjustValue {
	_result := &_BACnetConstructedDataAccessZoneAdjustValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		AdjustValue:                   adjustValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccessZoneAdjustValue(structType any) BACnetConstructedDataAccessZoneAdjustValue {
	if casted, ok := structType.(BACnetConstructedDataAccessZoneAdjustValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccessZoneAdjustValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccessZoneAdjustValue) GetTypeName() string {
	return "BACnetConstructedDataAccessZoneAdjustValue"
}

func (m *_BACnetConstructedDataAccessZoneAdjustValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (adjustValue)
	lengthInBits += m.AdjustValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAccessZoneAdjustValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAccessZoneAdjustValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAccessZoneAdjustValue BACnetConstructedDataAccessZoneAdjustValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccessZoneAdjustValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccessZoneAdjustValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	adjustValue, err := ReadSimpleField[BACnetApplicationTagSignedInteger](ctx, "adjustValue", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagSignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'adjustValue' field"))
	}
	m.AdjustValue = adjustValue

	actualValue, err := ReadVirtualField[BACnetApplicationTagSignedInteger](ctx, "actualValue", (*BACnetApplicationTagSignedInteger)(nil), adjustValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccessZoneAdjustValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccessZoneAdjustValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAccessZoneAdjustValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccessZoneAdjustValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccessZoneAdjustValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccessZoneAdjustValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagSignedInteger](ctx, "adjustValue", m.GetAdjustValue(), WriteComplex[BACnetApplicationTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'adjustValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccessZoneAdjustValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccessZoneAdjustValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccessZoneAdjustValue) IsBACnetConstructedDataAccessZoneAdjustValue() {
}

func (m *_BACnetConstructedDataAccessZoneAdjustValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
