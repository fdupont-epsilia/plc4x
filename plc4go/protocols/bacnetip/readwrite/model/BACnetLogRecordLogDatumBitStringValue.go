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

// BACnetLogRecordLogDatumBitStringValue is the corresponding interface of BACnetLogRecordLogDatumBitStringValue
type BACnetLogRecordLogDatumBitStringValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetLogRecordLogDatum
	// GetBitStringValue returns BitStringValue (property field)
	GetBitStringValue() BACnetContextTagBitString
}

// BACnetLogRecordLogDatumBitStringValueExactly can be used when we want exactly this type and not a type which fulfills BACnetLogRecordLogDatumBitStringValue.
// This is useful for switch cases.
type BACnetLogRecordLogDatumBitStringValueExactly interface {
	BACnetLogRecordLogDatumBitStringValue
	isBACnetLogRecordLogDatumBitStringValue() bool
}

// _BACnetLogRecordLogDatumBitStringValue is the data-structure of this message
type _BACnetLogRecordLogDatumBitStringValue struct {
	BACnetLogRecordLogDatumContract
	BitStringValue BACnetContextTagBitString
}

var _ BACnetLogRecordLogDatumBitStringValue = (*_BACnetLogRecordLogDatumBitStringValue)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetLogRecordLogDatumBitStringValue) GetParent() BACnetLogRecordLogDatumContract {
	return m.BACnetLogRecordLogDatumContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogRecordLogDatumBitStringValue) GetBitStringValue() BACnetContextTagBitString {
	return m.BitStringValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogRecordLogDatumBitStringValue factory function for _BACnetLogRecordLogDatumBitStringValue
func NewBACnetLogRecordLogDatumBitStringValue(bitStringValue BACnetContextTagBitString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetLogRecordLogDatumBitStringValue {
	_result := &_BACnetLogRecordLogDatumBitStringValue{
		BACnetLogRecordLogDatumContract: NewBACnetLogRecordLogDatum(openingTag, peekedTagHeader, closingTag, tagNumber),
		BitStringValue:                  bitStringValue,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetLogRecordLogDatumBitStringValue(structType any) BACnetLogRecordLogDatumBitStringValue {
	if casted, ok := structType.(BACnetLogRecordLogDatumBitStringValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogRecordLogDatumBitStringValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogRecordLogDatumBitStringValue) GetTypeName() string {
	return "BACnetLogRecordLogDatumBitStringValue"
}

func (m *_BACnetLogRecordLogDatumBitStringValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum).getLengthInBits(ctx))

	// Simple field (bitStringValue)
	lengthInBits += m.BitStringValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLogRecordLogDatumBitStringValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetLogRecordLogDatumBitStringValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetLogRecordLogDatum, tagNumber uint8) (_bACnetLogRecordLogDatumBitStringValue BACnetLogRecordLogDatumBitStringValue, err error) {
	m.BACnetLogRecordLogDatumContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogRecordLogDatumBitStringValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogRecordLogDatumBitStringValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	bitStringValue, err := ReadSimpleField[BACnetContextTagBitString](ctx, "bitStringValue", ReadComplex[BACnetContextTagBitString](BACnetContextTagParseWithBufferProducer[BACnetContextTagBitString]((uint8)(uint8(6)), (BACnetDataType)(BACnetDataType_BIT_STRING)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bitStringValue' field"))
	}
	m.BitStringValue = bitStringValue

	if closeErr := readBuffer.CloseContext("BACnetLogRecordLogDatumBitStringValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogRecordLogDatumBitStringValue")
	}

	return m, nil
}

func (m *_BACnetLogRecordLogDatumBitStringValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogRecordLogDatumBitStringValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogRecordLogDatumBitStringValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogRecordLogDatumBitStringValue")
		}

		if err := WriteSimpleField[BACnetContextTagBitString](ctx, "bitStringValue", m.GetBitStringValue(), WriteComplex[BACnetContextTagBitString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'bitStringValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogRecordLogDatumBitStringValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogRecordLogDatumBitStringValue")
		}
		return nil
	}
	return m.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetLogRecordLogDatumBitStringValue) isBACnetLogRecordLogDatumBitStringValue() bool {
	return true
}

func (m *_BACnetLogRecordLogDatumBitStringValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
