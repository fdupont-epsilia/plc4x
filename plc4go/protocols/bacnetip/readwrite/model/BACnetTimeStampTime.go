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

// BACnetTimeStampTime is the corresponding interface of BACnetTimeStampTime
type BACnetTimeStampTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetTimeStamp
	// GetTimeValue returns TimeValue (property field)
	GetTimeValue() BACnetContextTagTime
	// IsBACnetTimeStampTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTimeStampTime()
}

// _BACnetTimeStampTime is the data-structure of this message
type _BACnetTimeStampTime struct {
	BACnetTimeStampContract
	TimeValue BACnetContextTagTime
}

var _ BACnetTimeStampTime = (*_BACnetTimeStampTime)(nil)
var _ BACnetTimeStampRequirements = (*_BACnetTimeStampTime)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetTimeStampTime) GetParent() BACnetTimeStampContract {
	return m.BACnetTimeStampContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimeStampTime) GetTimeValue() BACnetContextTagTime {
	return m.TimeValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTimeStampTime factory function for _BACnetTimeStampTime
func NewBACnetTimeStampTime(timeValue BACnetContextTagTime, peekedTagHeader BACnetTagHeader) *_BACnetTimeStampTime {
	_result := &_BACnetTimeStampTime{
		BACnetTimeStampContract: NewBACnetTimeStamp(peekedTagHeader),
		TimeValue:               timeValue,
	}
	_result.BACnetTimeStampContract.(*_BACnetTimeStamp)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetTimeStampTime(structType any) BACnetTimeStampTime {
	if casted, ok := structType.(BACnetTimeStampTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimeStampTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimeStampTime) GetTypeName() string {
	return "BACnetTimeStampTime"
}

func (m *_BACnetTimeStampTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetTimeStampContract.(*_BACnetTimeStamp).getLengthInBits(ctx))

	// Simple field (timeValue)
	lengthInBits += m.TimeValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetTimeStampTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetTimeStampTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetTimeStamp) (__bACnetTimeStampTime BACnetTimeStampTime, err error) {
	m.BACnetTimeStampContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimeStampTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimeStampTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timeValue, err := ReadSimpleField[BACnetContextTagTime](ctx, "timeValue", ReadComplex[BACnetContextTagTime](BACnetContextTagParseWithBufferProducer[BACnetContextTagTime]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_TIME)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeValue' field"))
	}
	m.TimeValue = timeValue

	if closeErr := readBuffer.CloseContext("BACnetTimeStampTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimeStampTime")
	}

	return m, nil
}

func (m *_BACnetTimeStampTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimeStampTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimeStampTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimeStampTime")
		}

		if err := WriteSimpleField[BACnetContextTagTime](ctx, "timeValue", m.GetTimeValue(), WriteComplex[BACnetContextTagTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimeStampTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimeStampTime")
		}
		return nil
	}
	return m.BACnetTimeStampContract.(*_BACnetTimeStamp).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetTimeStampTime) IsBACnetTimeStampTime() {}

func (m *_BACnetTimeStampTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
