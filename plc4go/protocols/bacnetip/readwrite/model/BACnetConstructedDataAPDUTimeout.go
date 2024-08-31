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

// BACnetConstructedDataAPDUTimeout is the corresponding interface of BACnetConstructedDataAPDUTimeout
type BACnetConstructedDataAPDUTimeout interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetApduTimeout returns ApduTimeout (property field)
	GetApduTimeout() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataAPDUTimeoutExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAPDUTimeout.
// This is useful for switch cases.
type BACnetConstructedDataAPDUTimeoutExactly interface {
	BACnetConstructedDataAPDUTimeout
	isBACnetConstructedDataAPDUTimeout() bool
}

// _BACnetConstructedDataAPDUTimeout is the data-structure of this message
type _BACnetConstructedDataAPDUTimeout struct {
	*_BACnetConstructedData
	ApduTimeout BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAPDUTimeout) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAPDUTimeout) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_APDU_TIMEOUT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAPDUTimeout) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAPDUTimeout) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAPDUTimeout) GetApduTimeout() BACnetApplicationTagUnsignedInteger {
	return m.ApduTimeout
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAPDUTimeout) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetApduTimeout())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAPDUTimeout factory function for _BACnetConstructedDataAPDUTimeout
func NewBACnetConstructedDataAPDUTimeout(apduTimeout BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAPDUTimeout {
	_result := &_BACnetConstructedDataAPDUTimeout{
		ApduTimeout:            apduTimeout,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAPDUTimeout(structType any) BACnetConstructedDataAPDUTimeout {
	if casted, ok := structType.(BACnetConstructedDataAPDUTimeout); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAPDUTimeout); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAPDUTimeout) GetTypeName() string {
	return "BACnetConstructedDataAPDUTimeout"
}

func (m *_BACnetConstructedDataAPDUTimeout) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (apduTimeout)
	lengthInBits += m.ApduTimeout.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAPDUTimeout) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataAPDUTimeoutParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAPDUTimeout, error) {
	return BACnetConstructedDataAPDUTimeoutParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataAPDUTimeoutParseWithBufferProducer(tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataAPDUTimeout, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataAPDUTimeout, error) {
		return BACnetConstructedDataAPDUTimeoutParseWithBuffer(ctx, readBuffer, tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
	}
}

func BACnetConstructedDataAPDUTimeoutParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAPDUTimeout, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAPDUTimeout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAPDUTimeout")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	apduTimeout, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "apduTimeout", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer(), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'apduTimeout' field"))
	}

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), apduTimeout)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAPDUTimeout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAPDUTimeout")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAPDUTimeout{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ApduTimeout: apduTimeout,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAPDUTimeout) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAPDUTimeout) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAPDUTimeout"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAPDUTimeout")
		}

		// Simple Field (apduTimeout)
		if pushErr := writeBuffer.PushContext("apduTimeout"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for apduTimeout")
		}
		_apduTimeoutErr := writeBuffer.WriteSerializable(ctx, m.GetApduTimeout())
		if popErr := writeBuffer.PopContext("apduTimeout"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for apduTimeout")
		}
		if _apduTimeoutErr != nil {
			return errors.Wrap(_apduTimeoutErr, "Error serializing 'apduTimeout' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAPDUTimeout"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAPDUTimeout")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAPDUTimeout) isBACnetConstructedDataAPDUTimeout() bool {
	return true
}

func (m *_BACnetConstructedDataAPDUTimeout) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
