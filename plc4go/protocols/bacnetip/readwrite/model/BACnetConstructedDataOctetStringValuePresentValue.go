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

// BACnetConstructedDataOctetStringValuePresentValue is the corresponding interface of BACnetConstructedDataOctetStringValuePresentValue
type BACnetConstructedDataOctetStringValuePresentValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() BACnetApplicationTagOctetString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagOctetString
}

// BACnetConstructedDataOctetStringValuePresentValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataOctetStringValuePresentValue.
// This is useful for switch cases.
type BACnetConstructedDataOctetStringValuePresentValueExactly interface {
	BACnetConstructedDataOctetStringValuePresentValue
	isBACnetConstructedDataOctetStringValuePresentValue() bool
}

// _BACnetConstructedDataOctetStringValuePresentValue is the data-structure of this message
type _BACnetConstructedDataOctetStringValuePresentValue struct {
	*_BACnetConstructedData
	PresentValue BACnetApplicationTagOctetString
}

var _ BACnetConstructedDataOctetStringValuePresentValue = (*_BACnetConstructedDataOctetStringValuePresentValue)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataOctetStringValuePresentValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_OCTETSTRING_VALUE
}

func (m *_BACnetConstructedDataOctetStringValuePresentValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRESENT_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataOctetStringValuePresentValue) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataOctetStringValuePresentValue) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataOctetStringValuePresentValue) GetPresentValue() BACnetApplicationTagOctetString {
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

func (m *_BACnetConstructedDataOctetStringValuePresentValue) GetActualValue() BACnetApplicationTagOctetString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagOctetString(m.GetPresentValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataOctetStringValuePresentValue factory function for _BACnetConstructedDataOctetStringValuePresentValue
func NewBACnetConstructedDataOctetStringValuePresentValue(presentValue BACnetApplicationTagOctetString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataOctetStringValuePresentValue {
	_result := &_BACnetConstructedDataOctetStringValuePresentValue{
		PresentValue:           presentValue,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataOctetStringValuePresentValue(structType any) BACnetConstructedDataOctetStringValuePresentValue {
	if casted, ok := structType.(BACnetConstructedDataOctetStringValuePresentValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataOctetStringValuePresentValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataOctetStringValuePresentValue) GetTypeName() string {
	return "BACnetConstructedDataOctetStringValuePresentValue"
}

func (m *_BACnetConstructedDataOctetStringValuePresentValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataOctetStringValuePresentValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataOctetStringValuePresentValueParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataOctetStringValuePresentValue, error) {
	return BACnetConstructedDataOctetStringValuePresentValueParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataOctetStringValuePresentValueParseWithBufferProducer(tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataOctetStringValuePresentValue, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataOctetStringValuePresentValue, error) {
		return BACnetConstructedDataOctetStringValuePresentValueParseWithBuffer(ctx, readBuffer, tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
	}
}

func BACnetConstructedDataOctetStringValuePresentValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataOctetStringValuePresentValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataOctetStringValuePresentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataOctetStringValuePresentValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	presentValue, err := ReadSimpleField[BACnetApplicationTagOctetString](ctx, "presentValue", ReadComplex[BACnetApplicationTagOctetString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagOctetString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'presentValue' field"))
	}

	actualValue, err := ReadVirtualField[BACnetApplicationTagOctetString](ctx, "actualValue", (*BACnetApplicationTagOctetString)(nil), presentValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataOctetStringValuePresentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataOctetStringValuePresentValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataOctetStringValuePresentValue{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		PresentValue: presentValue,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataOctetStringValuePresentValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataOctetStringValuePresentValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataOctetStringValuePresentValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataOctetStringValuePresentValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagOctetString](ctx, "presentValue", m.GetPresentValue(), WriteComplex[BACnetApplicationTagOctetString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'presentValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataOctetStringValuePresentValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataOctetStringValuePresentValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataOctetStringValuePresentValue) isBACnetConstructedDataOctetStringValuePresentValue() bool {
	return true
}

func (m *_BACnetConstructedDataOctetStringValuePresentValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
