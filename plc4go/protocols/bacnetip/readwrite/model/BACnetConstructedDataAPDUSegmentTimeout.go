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

// BACnetConstructedDataAPDUSegmentTimeout is the corresponding interface of BACnetConstructedDataAPDUSegmentTimeout
type BACnetConstructedDataAPDUSegmentTimeout interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetApduSegmentTimeout returns ApduSegmentTimeout (property field)
	GetApduSegmentTimeout() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataAPDUSegmentTimeout is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAPDUSegmentTimeout()
}

// _BACnetConstructedDataAPDUSegmentTimeout is the data-structure of this message
type _BACnetConstructedDataAPDUSegmentTimeout struct {
	BACnetConstructedDataContract
	ApduSegmentTimeout BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataAPDUSegmentTimeout = (*_BACnetConstructedDataAPDUSegmentTimeout)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAPDUSegmentTimeout)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAPDUSegmentTimeout) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAPDUSegmentTimeout) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_APDU_SEGMENT_TIMEOUT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAPDUSegmentTimeout) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAPDUSegmentTimeout) GetApduSegmentTimeout() BACnetApplicationTagUnsignedInteger {
	return m.ApduSegmentTimeout
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAPDUSegmentTimeout) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetApduSegmentTimeout())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAPDUSegmentTimeout factory function for _BACnetConstructedDataAPDUSegmentTimeout
func NewBACnetConstructedDataAPDUSegmentTimeout(apduSegmentTimeout BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAPDUSegmentTimeout {
	_result := &_BACnetConstructedDataAPDUSegmentTimeout{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ApduSegmentTimeout:            apduSegmentTimeout,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAPDUSegmentTimeout(structType any) BACnetConstructedDataAPDUSegmentTimeout {
	if casted, ok := structType.(BACnetConstructedDataAPDUSegmentTimeout); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAPDUSegmentTimeout); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAPDUSegmentTimeout) GetTypeName() string {
	return "BACnetConstructedDataAPDUSegmentTimeout"
}

func (m *_BACnetConstructedDataAPDUSegmentTimeout) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (apduSegmentTimeout)
	lengthInBits += m.ApduSegmentTimeout.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAPDUSegmentTimeout) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAPDUSegmentTimeout) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAPDUSegmentTimeout BACnetConstructedDataAPDUSegmentTimeout, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAPDUSegmentTimeout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAPDUSegmentTimeout")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	apduSegmentTimeout, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "apduSegmentTimeout", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'apduSegmentTimeout' field"))
	}
	m.ApduSegmentTimeout = apduSegmentTimeout

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), apduSegmentTimeout)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAPDUSegmentTimeout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAPDUSegmentTimeout")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAPDUSegmentTimeout) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAPDUSegmentTimeout) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAPDUSegmentTimeout"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAPDUSegmentTimeout")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "apduSegmentTimeout", m.GetApduSegmentTimeout(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'apduSegmentTimeout' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAPDUSegmentTimeout"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAPDUSegmentTimeout")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAPDUSegmentTimeout) IsBACnetConstructedDataAPDUSegmentTimeout() {}

func (m *_BACnetConstructedDataAPDUSegmentTimeout) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
