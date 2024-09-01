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

// ModbusPDUReadWriteMultipleHoldingRegistersRequest is the corresponding interface of ModbusPDUReadWriteMultipleHoldingRegistersRequest
type ModbusPDUReadWriteMultipleHoldingRegistersRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetReadStartingAddress returns ReadStartingAddress (property field)
	GetReadStartingAddress() uint16
	// GetReadQuantity returns ReadQuantity (property field)
	GetReadQuantity() uint16
	// GetWriteStartingAddress returns WriteStartingAddress (property field)
	GetWriteStartingAddress() uint16
	// GetWriteQuantity returns WriteQuantity (property field)
	GetWriteQuantity() uint16
	// GetValue returns Value (property field)
	GetValue() []byte
}

// ModbusPDUReadWriteMultipleHoldingRegistersRequestExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUReadWriteMultipleHoldingRegistersRequest.
// This is useful for switch cases.
type ModbusPDUReadWriteMultipleHoldingRegistersRequestExactly interface {
	ModbusPDUReadWriteMultipleHoldingRegistersRequest
	isModbusPDUReadWriteMultipleHoldingRegistersRequest() bool
}

// _ModbusPDUReadWriteMultipleHoldingRegistersRequest is the data-structure of this message
type _ModbusPDUReadWriteMultipleHoldingRegistersRequest struct {
	*_ModbusPDU
	ReadStartingAddress  uint16
	ReadQuantity         uint16
	WriteStartingAddress uint16
	WriteQuantity        uint16
	Value                []byte
}

var _ ModbusPDUReadWriteMultipleHoldingRegistersRequest = (*_ModbusPDUReadWriteMultipleHoldingRegistersRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetFunctionFlag() uint8 {
	return 0x17
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) InitializeParent(parent ModbusPDU) {}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetParent() ModbusPDU {
	return m._ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetReadStartingAddress() uint16 {
	return m.ReadStartingAddress
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetReadQuantity() uint16 {
	return m.ReadQuantity
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetWriteStartingAddress() uint16 {
	return m.WriteStartingAddress
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetWriteQuantity() uint16 {
	return m.WriteQuantity
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetValue() []byte {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUReadWriteMultipleHoldingRegistersRequest factory function for _ModbusPDUReadWriteMultipleHoldingRegistersRequest
func NewModbusPDUReadWriteMultipleHoldingRegistersRequest(readStartingAddress uint16, readQuantity uint16, writeStartingAddress uint16, writeQuantity uint16, value []byte) *_ModbusPDUReadWriteMultipleHoldingRegistersRequest {
	_result := &_ModbusPDUReadWriteMultipleHoldingRegistersRequest{
		ReadStartingAddress:  readStartingAddress,
		ReadQuantity:         readQuantity,
		WriteStartingAddress: writeStartingAddress,
		WriteQuantity:        writeQuantity,
		Value:                value,
		_ModbusPDU:           NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUReadWriteMultipleHoldingRegistersRequest(structType any) ModbusPDUReadWriteMultipleHoldingRegistersRequest {
	if casted, ok := structType.(ModbusPDUReadWriteMultipleHoldingRegistersRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadWriteMultipleHoldingRegistersRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetTypeName() string {
	return "ModbusPDUReadWriteMultipleHoldingRegistersRequest"
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (readStartingAddress)
	lengthInBits += 16

	// Simple field (readQuantity)
	lengthInBits += 16

	// Simple field (writeStartingAddress)
	lengthInBits += 16

	// Simple field (writeQuantity)
	lengthInBits += 16

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 8 * uint16(len(m.Value))
	}

	return lengthInBits
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusPDUReadWriteMultipleHoldingRegistersRequestParse(ctx context.Context, theBytes []byte, response bool) (ModbusPDUReadWriteMultipleHoldingRegistersRequest, error) {
	return ModbusPDUReadWriteMultipleHoldingRegistersRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func ModbusPDUReadWriteMultipleHoldingRegistersRequestParseWithBufferProducer(response bool) func(ctx context.Context, readBuffer utils.ReadBuffer) (ModbusPDUReadWriteMultipleHoldingRegistersRequest, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ModbusPDUReadWriteMultipleHoldingRegistersRequest, error) {
		return ModbusPDUReadWriteMultipleHoldingRegistersRequestParseWithBuffer(ctx, readBuffer, response)
	}
}

func ModbusPDUReadWriteMultipleHoldingRegistersRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (ModbusPDUReadWriteMultipleHoldingRegistersRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadWriteMultipleHoldingRegistersRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadWriteMultipleHoldingRegistersRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	readStartingAddress, err := ReadSimpleField(ctx, "readStartingAddress", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'readStartingAddress' field"))
	}

	readQuantity, err := ReadSimpleField(ctx, "readQuantity", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'readQuantity' field"))
	}

	writeStartingAddress, err := ReadSimpleField(ctx, "writeStartingAddress", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'writeStartingAddress' field"))
	}

	writeQuantity, err := ReadSimpleField(ctx, "writeQuantity", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'writeQuantity' field"))
	}

	byteCount, err := ReadImplicitField[uint8](ctx, "byteCount", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'byteCount' field"))
	}
	_ = byteCount

	value, err := readBuffer.ReadByteArray("value", int(byteCount))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUReadWriteMultipleHoldingRegistersRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadWriteMultipleHoldingRegistersRequest")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUReadWriteMultipleHoldingRegistersRequest{
		_ModbusPDU:           &_ModbusPDU{},
		ReadStartingAddress:  readStartingAddress,
		ReadQuantity:         readQuantity,
		WriteStartingAddress: writeStartingAddress,
		WriteQuantity:        writeQuantity,
		Value:                value,
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadWriteMultipleHoldingRegistersRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadWriteMultipleHoldingRegistersRequest")
		}

		if err := WriteSimpleField[uint16](ctx, "readStartingAddress", m.GetReadStartingAddress(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'readStartingAddress' field")
		}

		if err := WriteSimpleField[uint16](ctx, "readQuantity", m.GetReadQuantity(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'readQuantity' field")
		}

		if err := WriteSimpleField[uint16](ctx, "writeStartingAddress", m.GetWriteStartingAddress(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'writeStartingAddress' field")
		}

		if err := WriteSimpleField[uint16](ctx, "writeQuantity", m.GetWriteQuantity(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'writeQuantity' field")
		}
		byteCount := uint8(uint8(len(m.GetValue())))
		if err := WriteImplicitField(ctx, "byteCount", byteCount, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'byteCount' field")
		}

		if err := WriteByteArrayField(ctx, "value", m.GetValue(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadWriteMultipleHoldingRegistersRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReadWriteMultipleHoldingRegistersRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) isModbusPDUReadWriteMultipleHoldingRegistersRequest() bool {
	return true
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
