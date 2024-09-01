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

// BACnetConstructedDataEventMessageTexts is the corresponding interface of BACnetConstructedDataEventMessageTexts
type BACnetConstructedDataEventMessageTexts interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetEventMessageTexts returns EventMessageTexts (property field)
	GetEventMessageTexts() []BACnetOptionalCharacterString
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// GetToOffnormalText returns ToOffnormalText (virtual field)
	GetToOffnormalText() BACnetOptionalCharacterString
	// GetToFaultText returns ToFaultText (virtual field)
	GetToFaultText() BACnetOptionalCharacterString
	// GetToNormalText returns ToNormalText (virtual field)
	GetToNormalText() BACnetOptionalCharacterString
}

// BACnetConstructedDataEventMessageTextsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataEventMessageTexts.
// This is useful for switch cases.
type BACnetConstructedDataEventMessageTextsExactly interface {
	BACnetConstructedDataEventMessageTexts
	isBACnetConstructedDataEventMessageTexts() bool
}

// _BACnetConstructedDataEventMessageTexts is the data-structure of this message
type _BACnetConstructedDataEventMessageTexts struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	EventMessageTexts    []BACnetOptionalCharacterString
}

var _ BACnetConstructedDataEventMessageTexts = (*_BACnetConstructedDataEventMessageTexts)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEventMessageTexts) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataEventMessageTexts) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EVENT_MESSAGE_TEXTS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEventMessageTexts) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEventMessageTexts) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataEventMessageTexts) GetEventMessageTexts() []BACnetOptionalCharacterString {
	return m.EventMessageTexts
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataEventMessageTexts) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return uint64(uint64(0))
}

func (m *_BACnetConstructedDataEventMessageTexts) GetToOffnormalText() BACnetOptionalCharacterString {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetOptionalCharacterString(CastBACnetOptionalCharacterString(utils.InlineIf(bool((len(m.GetEventMessageTexts())) == (3)), func() any { return CastBACnetOptionalCharacterString(m.GetEventMessageTexts()[0]) }, func() any { return CastBACnetOptionalCharacterString(nil) })))
}

func (m *_BACnetConstructedDataEventMessageTexts) GetToFaultText() BACnetOptionalCharacterString {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetOptionalCharacterString(CastBACnetOptionalCharacterString(utils.InlineIf(bool((len(m.GetEventMessageTexts())) == (3)), func() any { return CastBACnetOptionalCharacterString(m.GetEventMessageTexts()[1]) }, func() any { return CastBACnetOptionalCharacterString(nil) })))
}

func (m *_BACnetConstructedDataEventMessageTexts) GetToNormalText() BACnetOptionalCharacterString {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetOptionalCharacterString(CastBACnetOptionalCharacterString(utils.InlineIf(bool((len(m.GetEventMessageTexts())) == (3)), func() any { return CastBACnetOptionalCharacterString(m.GetEventMessageTexts()[2]) }, func() any { return CastBACnetOptionalCharacterString(nil) })))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEventMessageTexts factory function for _BACnetConstructedDataEventMessageTexts
func NewBACnetConstructedDataEventMessageTexts(numberOfDataElements BACnetApplicationTagUnsignedInteger, eventMessageTexts []BACnetOptionalCharacterString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEventMessageTexts {
	_result := &_BACnetConstructedDataEventMessageTexts{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		EventMessageTexts:             eventMessageTexts,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEventMessageTexts(structType any) BACnetConstructedDataEventMessageTexts {
	if casted, ok := structType.(BACnetConstructedDataEventMessageTexts); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEventMessageTexts); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEventMessageTexts) GetTypeName() string {
	return "BACnetConstructedDataEventMessageTexts"
}

func (m *_BACnetConstructedDataEventMessageTexts) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.EventMessageTexts) > 0 {
		for _, element := range m.EventMessageTexts {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataEventMessageTexts) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataEventMessageTexts) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (_bACnetConstructedDataEventMessageTexts BACnetConstructedDataEventMessageTexts, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEventMessageTexts"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEventMessageTexts")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zero, err := ReadVirtualField[uint64](ctx, "zero", (*uint64)(nil), uint64(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zero' field"))
	}
	_ = zero

	_numberOfDataElements, err := ReadOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfDataElements' field"))
	}
	var numberOfDataElements BACnetApplicationTagUnsignedInteger
	if _numberOfDataElements != nil {
		numberOfDataElements = *_numberOfDataElements
	}
	m.NumberOfDataElements = numberOfDataElements

	eventMessageTexts, err := ReadTerminatedArrayField[BACnetOptionalCharacterString](ctx, "eventMessageTexts", ReadComplex[BACnetOptionalCharacterString](BACnetOptionalCharacterStringParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventMessageTexts' field"))
	}
	m.EventMessageTexts = eventMessageTexts

	toOffnormalText, err := ReadVirtualField[BACnetOptionalCharacterString](ctx, "toOffnormalText", (*BACnetOptionalCharacterString)(nil), CastBACnetOptionalCharacterString(utils.InlineIf(bool((len(eventMessageTexts)) == (3)), func() any { return CastBACnetOptionalCharacterString(eventMessageTexts[0]) }, func() any { return CastBACnetOptionalCharacterString(nil) })))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toOffnormalText' field"))
	}
	_ = toOffnormalText

	toFaultText, err := ReadVirtualField[BACnetOptionalCharacterString](ctx, "toFaultText", (*BACnetOptionalCharacterString)(nil), CastBACnetOptionalCharacterString(utils.InlineIf(bool((len(eventMessageTexts)) == (3)), func() any { return CastBACnetOptionalCharacterString(eventMessageTexts[1]) }, func() any { return CastBACnetOptionalCharacterString(nil) })))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toFaultText' field"))
	}
	_ = toFaultText

	toNormalText, err := ReadVirtualField[BACnetOptionalCharacterString](ctx, "toNormalText", (*BACnetOptionalCharacterString)(nil), CastBACnetOptionalCharacterString(utils.InlineIf(bool((len(eventMessageTexts)) == (3)), func() any { return CastBACnetOptionalCharacterString(eventMessageTexts[2]) }, func() any { return CastBACnetOptionalCharacterString(nil) })))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toNormalText' field"))
	}
	_ = toNormalText

	// Validation
	if !(bool(bool((arrayIndexArgument) != (nil))) || bool(bool((len(eventMessageTexts)) == (3)))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "eventMessageTexts should have exactly 3 values"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEventMessageTexts"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEventMessageTexts")
	}

	return m, nil
}

func (m *_BACnetConstructedDataEventMessageTexts) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEventMessageTexts) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEventMessageTexts"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEventMessageTexts")
		}
		// Virtual field
		zero := m.GetZero()
		_ = zero
		if _zeroErr := writeBuffer.WriteVirtual(ctx, "zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		if err := WriteOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", GetRef(m.GetNumberOfDataElements()), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'numberOfDataElements' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "eventMessageTexts", m.GetEventMessageTexts(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'eventMessageTexts' field")
		}
		// Virtual field
		toOffnormalText := m.GetToOffnormalText()
		_ = toOffnormalText
		if _toOffnormalTextErr := writeBuffer.WriteVirtual(ctx, "toOffnormalText", m.GetToOffnormalText()); _toOffnormalTextErr != nil {
			return errors.Wrap(_toOffnormalTextErr, "Error serializing 'toOffnormalText' field")
		}
		// Virtual field
		toFaultText := m.GetToFaultText()
		_ = toFaultText
		if _toFaultTextErr := writeBuffer.WriteVirtual(ctx, "toFaultText", m.GetToFaultText()); _toFaultTextErr != nil {
			return errors.Wrap(_toFaultTextErr, "Error serializing 'toFaultText' field")
		}
		// Virtual field
		toNormalText := m.GetToNormalText()
		_ = toNormalText
		if _toNormalTextErr := writeBuffer.WriteVirtual(ctx, "toNormalText", m.GetToNormalText()); _toNormalTextErr != nil {
			return errors.Wrap(_toNormalTextErr, "Error serializing 'toNormalText' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEventMessageTexts"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEventMessageTexts")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEventMessageTexts) isBACnetConstructedDataEventMessageTexts() bool {
	return true
}

func (m *_BACnetConstructedDataEventMessageTexts) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
