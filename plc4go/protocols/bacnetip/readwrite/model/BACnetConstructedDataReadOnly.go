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

// BACnetConstructedDataReadOnly is the corresponding interface of BACnetConstructedDataReadOnly
type BACnetConstructedDataReadOnly interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetReadOnly returns ReadOnly (property field)
	GetReadOnly() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
	// IsBACnetConstructedDataReadOnly is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataReadOnly()
}

// _BACnetConstructedDataReadOnly is the data-structure of this message
type _BACnetConstructedDataReadOnly struct {
	BACnetConstructedDataContract
	ReadOnly BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataReadOnly = (*_BACnetConstructedDataReadOnly)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataReadOnly)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataReadOnly) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataReadOnly) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_READ_ONLY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataReadOnly) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataReadOnly) GetReadOnly() BACnetApplicationTagBoolean {
	return m.ReadOnly
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataReadOnly) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetReadOnly())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataReadOnly factory function for _BACnetConstructedDataReadOnly
func NewBACnetConstructedDataReadOnly(readOnly BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataReadOnly {
	_result := &_BACnetConstructedDataReadOnly{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ReadOnly:                      readOnly,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataReadOnly(structType any) BACnetConstructedDataReadOnly {
	if casted, ok := structType.(BACnetConstructedDataReadOnly); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataReadOnly); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataReadOnly) GetTypeName() string {
	return "BACnetConstructedDataReadOnly"
}

func (m *_BACnetConstructedDataReadOnly) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (readOnly)
	lengthInBits += m.ReadOnly.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataReadOnly) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataReadOnly) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataReadOnly BACnetConstructedDataReadOnly, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataReadOnly"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataReadOnly")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	readOnly, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "readOnly", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'readOnly' field"))
	}
	m.ReadOnly = readOnly

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), readOnly)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataReadOnly"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataReadOnly")
	}

	return m, nil
}

func (m *_BACnetConstructedDataReadOnly) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataReadOnly) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataReadOnly"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataReadOnly")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "readOnly", m.GetReadOnly(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'readOnly' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataReadOnly"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataReadOnly")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataReadOnly) IsBACnetConstructedDataReadOnly() {}

func (m *_BACnetConstructedDataReadOnly) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
