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

// BACnetTimeValue is the corresponding interface of BACnetTimeValue
type BACnetTimeValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetTimeValue returns TimeValue (property field)
	GetTimeValue() BACnetApplicationTagTime
	// GetValue returns Value (property field)
	GetValue() BACnetConstructedDataElement
}

// BACnetTimeValueExactly can be used when we want exactly this type and not a type which fulfills BACnetTimeValue.
// This is useful for switch cases.
type BACnetTimeValueExactly interface {
	BACnetTimeValue
	isBACnetTimeValue() bool
}

// _BACnetTimeValue is the data-structure of this message
type _BACnetTimeValue struct {
	TimeValue BACnetApplicationTagTime
	Value     BACnetConstructedDataElement
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimeValue) GetTimeValue() BACnetApplicationTagTime {
	return m.TimeValue
}

func (m *_BACnetTimeValue) GetValue() BACnetConstructedDataElement {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTimeValue factory function for _BACnetTimeValue
func NewBACnetTimeValue(timeValue BACnetApplicationTagTime, value BACnetConstructedDataElement) *_BACnetTimeValue {
	return &_BACnetTimeValue{TimeValue: timeValue, Value: value}
}

// Deprecated: use the interface for direct cast
func CastBACnetTimeValue(structType any) BACnetTimeValue {
	if casted, ok := structType.(BACnetTimeValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimeValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimeValue) GetTypeName() string {
	return "BACnetTimeValue"
}

func (m *_BACnetTimeValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (timeValue)
	lengthInBits += m.TimeValue.GetLengthInBits(ctx)

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetTimeValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTimeValueParse(ctx context.Context, theBytes []byte) (BACnetTimeValue, error) {
	return BACnetTimeValueParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetTimeValueParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTimeValue, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTimeValue, error) {
		return BACnetTimeValueParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetTimeValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTimeValue, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetTimeValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimeValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timeValue, err := ReadSimpleField[BACnetApplicationTagTime](ctx, "timeValue", ReadComplex[BACnetApplicationTagTime](BACnetApplicationTagParseWithBufferProducer(), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeValue' field"))
	}

	value, err := ReadSimpleField[BACnetConstructedDataElement](ctx, "value", ReadComplex[BACnetConstructedDataElement](BACnetConstructedDataElementParseWithBufferProducer((BACnetObjectType)(BACnetObjectType_VENDOR_PROPRIETARY_VALUE), (BACnetPropertyIdentifier)(BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE), (BACnetTagPayloadUnsignedInteger)(nil)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetTimeValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimeValue")
	}

	// Create the instance
	return &_BACnetTimeValue{
		TimeValue: timeValue,
		Value:     value,
	}, nil
}

func (m *_BACnetTimeValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimeValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetTimeValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTimeValue")
	}

	// Simple Field (timeValue)
	if pushErr := writeBuffer.PushContext("timeValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for timeValue")
	}
	_timeValueErr := writeBuffer.WriteSerializable(ctx, m.GetTimeValue())
	if popErr := writeBuffer.PopContext("timeValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for timeValue")
	}
	if _timeValueErr != nil {
		return errors.Wrap(_timeValueErr, "Error serializing 'timeValue' field")
	}

	// Simple Field (value)
	if pushErr := writeBuffer.PushContext("value"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for value")
	}
	_valueErr := writeBuffer.WriteSerializable(ctx, m.GetValue())
	if popErr := writeBuffer.PopContext("value"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for value")
	}
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTimeValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTimeValue")
	}
	return nil
}

func (m *_BACnetTimeValue) isBACnetTimeValue() bool {
	return true
}

func (m *_BACnetTimeValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
