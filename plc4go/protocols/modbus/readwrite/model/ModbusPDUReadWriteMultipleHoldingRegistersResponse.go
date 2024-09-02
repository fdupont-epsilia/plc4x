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

// ModbusPDUReadWriteMultipleHoldingRegistersResponse is the corresponding interface of ModbusPDUReadWriteMultipleHoldingRegistersResponse
type ModbusPDUReadWriteMultipleHoldingRegistersResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetValue returns Value (property field)
	GetValue() []byte
	// IsModbusPDUReadWriteMultipleHoldingRegistersResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusPDUReadWriteMultipleHoldingRegistersResponse()
}

// _ModbusPDUReadWriteMultipleHoldingRegistersResponse is the data-structure of this message
type _ModbusPDUReadWriteMultipleHoldingRegistersResponse struct {
	ModbusPDUContract
	Value []byte
}

var _ ModbusPDUReadWriteMultipleHoldingRegistersResponse = (*_ModbusPDUReadWriteMultipleHoldingRegistersResponse)(nil)
var _ ModbusPDURequirements = (*_ModbusPDUReadWriteMultipleHoldingRegistersResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersResponse) GetFunctionFlag() uint8 {
	return 0x17
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersResponse) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersResponse) GetValue() []byte {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUReadWriteMultipleHoldingRegistersResponse factory function for _ModbusPDUReadWriteMultipleHoldingRegistersResponse
func NewModbusPDUReadWriteMultipleHoldingRegistersResponse(value []byte) *_ModbusPDUReadWriteMultipleHoldingRegistersResponse {
	_result := &_ModbusPDUReadWriteMultipleHoldingRegistersResponse{
		ModbusPDUContract: NewModbusPDU(),
		Value:             value,
	}
	_result.ModbusPDUContract.(*_ModbusPDU)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUReadWriteMultipleHoldingRegistersResponse(structType any) ModbusPDUReadWriteMultipleHoldingRegistersResponse {
	if casted, ok := structType.(ModbusPDUReadWriteMultipleHoldingRegistersResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadWriteMultipleHoldingRegistersResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersResponse) GetTypeName() string {
	return "ModbusPDUReadWriteMultipleHoldingRegistersResponse"
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 8 * uint16(len(m.Value))
	}

	return lengthInBits
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (__modbusPDUReadWriteMultipleHoldingRegistersResponse ModbusPDUReadWriteMultipleHoldingRegistersResponse, err error) {
	m.ModbusPDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadWriteMultipleHoldingRegistersResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadWriteMultipleHoldingRegistersResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	byteCount, err := ReadImplicitField[uint8](ctx, "byteCount", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'byteCount' field"))
	}
	_ = byteCount

	value, err := readBuffer.ReadByteArray("value", int(byteCount))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("ModbusPDUReadWriteMultipleHoldingRegistersResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadWriteMultipleHoldingRegistersResponse")
	}

	return m, nil
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadWriteMultipleHoldingRegistersResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadWriteMultipleHoldingRegistersResponse")
		}
		byteCount := uint8(uint8(len(m.GetValue())))
		if err := WriteImplicitField(ctx, "byteCount", byteCount, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'byteCount' field")
		}

		if err := WriteByteArrayField(ctx, "value", m.GetValue(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadWriteMultipleHoldingRegistersResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReadWriteMultipleHoldingRegistersResponse")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersResponse) IsModbusPDUReadWriteMultipleHoldingRegistersResponse() {
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
