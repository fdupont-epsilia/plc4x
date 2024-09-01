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

// S7PayloadUserDataItemClkFResponse is the corresponding interface of S7PayloadUserDataItemClkFResponse
type S7PayloadUserDataItemClkFResponse interface {
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
}

// S7PayloadUserDataItemClkFResponseExactly can be used when we want exactly this type and not a type which fulfills S7PayloadUserDataItemClkFResponse.
// This is useful for switch cases.
type S7PayloadUserDataItemClkFResponseExactly interface {
	S7PayloadUserDataItemClkFResponse
	isS7PayloadUserDataItemClkFResponse() bool
}

// _S7PayloadUserDataItemClkFResponse is the data-structure of this message
type _S7PayloadUserDataItemClkFResponse struct {
	*_S7PayloadUserDataItem
	Res       uint8
	Year1     uint8
	TimeStamp DateAndTime
}

var _ S7PayloadUserDataItemClkFResponse = (*_S7PayloadUserDataItemClkFResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemClkFResponse) GetCpuFunctionGroup() uint8 {
	return 0x07
}

func (m *_S7PayloadUserDataItemClkFResponse) GetCpuFunctionType() uint8 {
	return 0x08
}

func (m *_S7PayloadUserDataItemClkFResponse) GetCpuSubfunction() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemClkFResponse) InitializeParent(parent S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) {
	m.ReturnCode = returnCode
	m.TransportSize = transportSize
	m.DataLength = dataLength
}

func (m *_S7PayloadUserDataItemClkFResponse) GetParent() S7PayloadUserDataItem {
	return m._S7PayloadUserDataItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadUserDataItemClkFResponse) GetRes() uint8 {
	return m.Res
}

func (m *_S7PayloadUserDataItemClkFResponse) GetYear1() uint8 {
	return m.Year1
}

func (m *_S7PayloadUserDataItemClkFResponse) GetTimeStamp() DateAndTime {
	return m.TimeStamp
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadUserDataItemClkFResponse factory function for _S7PayloadUserDataItemClkFResponse
func NewS7PayloadUserDataItemClkFResponse(res uint8, year1 uint8, timeStamp DateAndTime, returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) *_S7PayloadUserDataItemClkFResponse {
	_result := &_S7PayloadUserDataItemClkFResponse{
		Res:                    res,
		Year1:                  year1,
		TimeStamp:              timeStamp,
		_S7PayloadUserDataItem: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
	}
	_result._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemClkFResponse(structType any) S7PayloadUserDataItemClkFResponse {
	if casted, ok := structType.(S7PayloadUserDataItemClkFResponse); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemClkFResponse); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemClkFResponse) GetTypeName() string {
	return "S7PayloadUserDataItemClkFResponse"
}

func (m *_S7PayloadUserDataItemClkFResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (res)
	lengthInBits += 8

	// Simple field (year1)
	lengthInBits += 8

	// Simple field (timeStamp)
	lengthInBits += m.TimeStamp.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_S7PayloadUserDataItemClkFResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7PayloadUserDataItemClkFResponseParse(ctx context.Context, theBytes []byte, dataLength uint16, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadUserDataItemClkFResponse, error) {
	return S7PayloadUserDataItemClkFResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), dataLength, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
}

func S7PayloadUserDataItemClkFResponseParseWithBufferProducer(dataLength uint16, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (S7PayloadUserDataItemClkFResponse, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (S7PayloadUserDataItemClkFResponse, error) {
		return S7PayloadUserDataItemClkFResponseParseWithBuffer(ctx, readBuffer, dataLength, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
	}
}

func S7PayloadUserDataItemClkFResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, dataLength uint16, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadUserDataItemClkFResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemClkFResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemClkFResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	res, err := ReadSimpleField(ctx, "res", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'res' field"))
	}

	year1, err := ReadSimpleField(ctx, "year1", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'year1' field"))
	}

	timeStamp, err := ReadSimpleField[DateAndTime](ctx, "timeStamp", ReadComplex[DateAndTime](DateAndTimeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeStamp' field"))
	}

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemClkFResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemClkFResponse")
	}

	// Create a partially initialized instance
	_child := &_S7PayloadUserDataItemClkFResponse{
		_S7PayloadUserDataItem: &_S7PayloadUserDataItem{},
		Res:                    res,
		Year1:                  year1,
		TimeStamp:              timeStamp,
	}
	_child._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _child
	return _child, nil
}

func (m *_S7PayloadUserDataItemClkFResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemClkFResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemClkFResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemClkFResponse")
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

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemClkFResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemClkFResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemClkFResponse) isS7PayloadUserDataItemClkFResponse() bool {
	return true
}

func (m *_S7PayloadUserDataItemClkFResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
