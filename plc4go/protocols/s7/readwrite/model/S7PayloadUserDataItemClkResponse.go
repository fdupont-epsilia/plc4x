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

// S7PayloadUserDataItemClkResponse is the corresponding interface of S7PayloadUserDataItemClkResponse
type S7PayloadUserDataItemClkResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7PayloadUserDataItem
	// GetRes returns Res (property field)
	GetRes() uint8
	// GetYear1 returns Year1 (property field)
	GetYear1() uint8
	// GetTimeStamp returns TimeStamp (property field)
	GetTimeStamp() DateAndTime
	// IsS7PayloadUserDataItemClkResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7PayloadUserDataItemClkResponse()
}

// _S7PayloadUserDataItemClkResponse is the data-structure of this message
type _S7PayloadUserDataItemClkResponse struct {
	S7PayloadUserDataItemContract
	Res       uint8
	Year1     uint8
	TimeStamp DateAndTime
}

var _ S7PayloadUserDataItemClkResponse = (*_S7PayloadUserDataItemClkResponse)(nil)
var _ S7PayloadUserDataItemRequirements = (*_S7PayloadUserDataItemClkResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemClkResponse) GetCpuFunctionGroup() uint8 {
	return 0x07
}

func (m *_S7PayloadUserDataItemClkResponse) GetCpuFunctionType() uint8 {
	return 0x08
}

func (m *_S7PayloadUserDataItemClkResponse) GetCpuSubfunction() uint8 {
	return 0x01
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemClkResponse) GetParent() S7PayloadUserDataItemContract {
	return m.S7PayloadUserDataItemContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadUserDataItemClkResponse) GetRes() uint8 {
	return m.Res
}

func (m *_S7PayloadUserDataItemClkResponse) GetYear1() uint8 {
	return m.Year1
}

func (m *_S7PayloadUserDataItemClkResponse) GetTimeStamp() DateAndTime {
	return m.TimeStamp
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadUserDataItemClkResponse factory function for _S7PayloadUserDataItemClkResponse
func NewS7PayloadUserDataItemClkResponse(res uint8, year1 uint8, timeStamp DateAndTime, returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) *_S7PayloadUserDataItemClkResponse {
	_result := &_S7PayloadUserDataItemClkResponse{
		S7PayloadUserDataItemContract: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
		Res:                           res,
		Year1:                         year1,
		TimeStamp:                     timeStamp,
	}
	_result.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemClkResponse(structType any) S7PayloadUserDataItemClkResponse {
	if casted, ok := structType.(S7PayloadUserDataItemClkResponse); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemClkResponse); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemClkResponse) GetTypeName() string {
	return "S7PayloadUserDataItemClkResponse"
}

func (m *_S7PayloadUserDataItemClkResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).getLengthInBits(ctx))

	// Simple field (res)
	lengthInBits += 8

	// Simple field (year1)
	lengthInBits += 8

	// Simple field (timeStamp)
	lengthInBits += m.TimeStamp.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_S7PayloadUserDataItemClkResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7PayloadUserDataItemClkResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7PayloadUserDataItem, dataLength uint16, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (__s7PayloadUserDataItemClkResponse S7PayloadUserDataItemClkResponse, err error) {
	m.S7PayloadUserDataItemContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemClkResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemClkResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	res, err := ReadSimpleField(ctx, "res", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'res' field"))
	}
	m.Res = res

	year1, err := ReadSimpleField(ctx, "year1", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'year1' field"))
	}
	m.Year1 = year1

	timeStamp, err := ReadSimpleField[DateAndTime](ctx, "timeStamp", ReadComplex[DateAndTime](DateAndTimeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeStamp' field"))
	}
	m.TimeStamp = timeStamp

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemClkResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemClkResponse")
	}

	return m, nil
}

func (m *_S7PayloadUserDataItemClkResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemClkResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemClkResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemClkResponse")
		}

		if err := WriteSimpleField[uint8](ctx, "res", m.GetRes(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'res' field")
		}

		if err := WriteSimpleField[uint8](ctx, "year1", m.GetYear1(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'year1' field")
		}

		if err := WriteSimpleField[DateAndTime](ctx, "timeStamp", m.GetTimeStamp(), WriteComplex[DateAndTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeStamp' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemClkResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemClkResponse")
		}
		return nil
	}
	return m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemClkResponse) IsS7PayloadUserDataItemClkResponse() {}

func (m *_S7PayloadUserDataItemClkResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
