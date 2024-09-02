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

// BACnetConstructedDataVendorIdentifier is the corresponding interface of BACnetConstructedDataVendorIdentifier
type BACnetConstructedDataVendorIdentifier interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetVendorIdentifier returns VendorIdentifier (property field)
	GetVendorIdentifier() BACnetVendorIdTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetVendorIdTagged
	// IsBACnetConstructedDataVendorIdentifier is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataVendorIdentifier()
}

// _BACnetConstructedDataVendorIdentifier is the data-structure of this message
type _BACnetConstructedDataVendorIdentifier struct {
	BACnetConstructedDataContract
	VendorIdentifier BACnetVendorIdTagged
}

var _ BACnetConstructedDataVendorIdentifier = (*_BACnetConstructedDataVendorIdentifier)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataVendorIdentifier)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataVendorIdentifier) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataVendorIdentifier) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_VENDOR_IDENTIFIER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataVendorIdentifier) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataVendorIdentifier) GetVendorIdentifier() BACnetVendorIdTagged {
	return m.VendorIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataVendorIdentifier) GetActualValue() BACnetVendorIdTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetVendorIdTagged(m.GetVendorIdentifier())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataVendorIdentifier factory function for _BACnetConstructedDataVendorIdentifier
func NewBACnetConstructedDataVendorIdentifier(vendorIdentifier BACnetVendorIdTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataVendorIdentifier {
	_result := &_BACnetConstructedDataVendorIdentifier{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		VendorIdentifier:              vendorIdentifier,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataVendorIdentifier(structType any) BACnetConstructedDataVendorIdentifier {
	if casted, ok := structType.(BACnetConstructedDataVendorIdentifier); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataVendorIdentifier); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataVendorIdentifier) GetTypeName() string {
	return "BACnetConstructedDataVendorIdentifier"
}

func (m *_BACnetConstructedDataVendorIdentifier) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (vendorIdentifier)
	lengthInBits += m.VendorIdentifier.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataVendorIdentifier) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataVendorIdentifier) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataVendorIdentifier BACnetConstructedDataVendorIdentifier, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataVendorIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataVendorIdentifier")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	vendorIdentifier, err := ReadSimpleField[BACnetVendorIdTagged](ctx, "vendorIdentifier", ReadComplex[BACnetVendorIdTagged](BACnetVendorIdTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'vendorIdentifier' field"))
	}
	m.VendorIdentifier = vendorIdentifier

	actualValue, err := ReadVirtualField[BACnetVendorIdTagged](ctx, "actualValue", (*BACnetVendorIdTagged)(nil), vendorIdentifier)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataVendorIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataVendorIdentifier")
	}

	return m, nil
}

func (m *_BACnetConstructedDataVendorIdentifier) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataVendorIdentifier) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataVendorIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataVendorIdentifier")
		}

		if err := WriteSimpleField[BACnetVendorIdTagged](ctx, "vendorIdentifier", m.GetVendorIdentifier(), WriteComplex[BACnetVendorIdTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'vendorIdentifier' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataVendorIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataVendorIdentifier")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataVendorIdentifier) IsBACnetConstructedDataVendorIdentifier() {}

func (m *_BACnetConstructedDataVendorIdentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
