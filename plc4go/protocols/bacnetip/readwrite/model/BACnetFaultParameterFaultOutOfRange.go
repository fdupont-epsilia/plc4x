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

// BACnetFaultParameterFaultOutOfRange is the corresponding interface of BACnetFaultParameterFaultOutOfRange
type BACnetFaultParameterFaultOutOfRange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetFaultParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetMinNormalValue returns MinNormalValue (property field)
	GetMinNormalValue() BACnetFaultParameterFaultOutOfRangeMinNormalValue
	// GetMaxNormalValue returns MaxNormalValue (property field)
	GetMaxNormalValue() BACnetFaultParameterFaultOutOfRangeMaxNormalValue
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetFaultParameterFaultOutOfRangeExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultOutOfRange.
// This is useful for switch cases.
type BACnetFaultParameterFaultOutOfRangeExactly interface {
	BACnetFaultParameterFaultOutOfRange
	isBACnetFaultParameterFaultOutOfRange() bool
}

// _BACnetFaultParameterFaultOutOfRange is the data-structure of this message
type _BACnetFaultParameterFaultOutOfRange struct {
	BACnetFaultParameterContract
	OpeningTag     BACnetOpeningTag
	MinNormalValue BACnetFaultParameterFaultOutOfRangeMinNormalValue
	MaxNormalValue BACnetFaultParameterFaultOutOfRangeMaxNormalValue
	ClosingTag     BACnetClosingTag
}

var _ BACnetFaultParameterFaultOutOfRange = (*_BACnetFaultParameterFaultOutOfRange)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetFaultParameterFaultOutOfRange) GetParent() BACnetFaultParameterContract {
	return m.BACnetFaultParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultOutOfRange) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetFaultParameterFaultOutOfRange) GetMinNormalValue() BACnetFaultParameterFaultOutOfRangeMinNormalValue {
	return m.MinNormalValue
}

func (m *_BACnetFaultParameterFaultOutOfRange) GetMaxNormalValue() BACnetFaultParameterFaultOutOfRangeMaxNormalValue {
	return m.MaxNormalValue
}

func (m *_BACnetFaultParameterFaultOutOfRange) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultOutOfRange factory function for _BACnetFaultParameterFaultOutOfRange
func NewBACnetFaultParameterFaultOutOfRange(openingTag BACnetOpeningTag, minNormalValue BACnetFaultParameterFaultOutOfRangeMinNormalValue, maxNormalValue BACnetFaultParameterFaultOutOfRangeMaxNormalValue, closingTag BACnetClosingTag, peekedTagHeader BACnetTagHeader) *_BACnetFaultParameterFaultOutOfRange {
	_result := &_BACnetFaultParameterFaultOutOfRange{
		BACnetFaultParameterContract: NewBACnetFaultParameter(peekedTagHeader),
		OpeningTag:                   openingTag,
		MinNormalValue:               minNormalValue,
		MaxNormalValue:               maxNormalValue,
		ClosingTag:                   closingTag,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultOutOfRange(structType any) BACnetFaultParameterFaultOutOfRange {
	if casted, ok := structType.(BACnetFaultParameterFaultOutOfRange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultOutOfRange) GetTypeName() string {
	return "BACnetFaultParameterFaultOutOfRange"
}

func (m *_BACnetFaultParameterFaultOutOfRange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetFaultParameterContract.(*_BACnetFaultParameter).getLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (minNormalValue)
	lengthInBits += m.MinNormalValue.GetLengthInBits(ctx)

	// Simple field (maxNormalValue)
	lengthInBits += m.MaxNormalValue.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultOutOfRange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetFaultParameterFaultOutOfRange) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetFaultParameter) (_bACnetFaultParameterFaultOutOfRange BACnetFaultParameterFaultOutOfRange, err error) {
	m.BACnetFaultParameterContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultOutOfRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultOutOfRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(6))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	minNormalValue, err := ReadSimpleField[BACnetFaultParameterFaultOutOfRangeMinNormalValue](ctx, "minNormalValue", ReadComplex[BACnetFaultParameterFaultOutOfRangeMinNormalValue](BACnetFaultParameterFaultOutOfRangeMinNormalValueParseWithBufferProducer[BACnetFaultParameterFaultOutOfRangeMinNormalValue]((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'minNormalValue' field"))
	}
	m.MinNormalValue = minNormalValue

	maxNormalValue, err := ReadSimpleField[BACnetFaultParameterFaultOutOfRangeMaxNormalValue](ctx, "maxNormalValue", ReadComplex[BACnetFaultParameterFaultOutOfRangeMaxNormalValue](BACnetFaultParameterFaultOutOfRangeMaxNormalValueParseWithBufferProducer[BACnetFaultParameterFaultOutOfRangeMaxNormalValue]((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxNormalValue' field"))
	}
	m.MaxNormalValue = maxNormalValue

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(6))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultOutOfRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultOutOfRange")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultOutOfRange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultOutOfRange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultOutOfRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultOutOfRange")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'openingTag' field")
		}

		if err := WriteSimpleField[BACnetFaultParameterFaultOutOfRangeMinNormalValue](ctx, "minNormalValue", m.GetMinNormalValue(), WriteComplex[BACnetFaultParameterFaultOutOfRangeMinNormalValue](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'minNormalValue' field")
		}

		if err := WriteSimpleField[BACnetFaultParameterFaultOutOfRangeMaxNormalValue](ctx, "maxNormalValue", m.GetMaxNormalValue(), WriteComplex[BACnetFaultParameterFaultOutOfRangeMaxNormalValue](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxNormalValue' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultOutOfRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultOutOfRange")
		}
		return nil
	}
	return m.BACnetFaultParameterContract.(*_BACnetFaultParameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultOutOfRange) isBACnetFaultParameterFaultOutOfRange() bool {
	return true
}

func (m *_BACnetFaultParameterFaultOutOfRange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
