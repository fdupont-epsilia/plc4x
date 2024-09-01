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

// ModbusPDUReadFifoQueueRequest is the corresponding interface of ModbusPDUReadFifoQueueRequest
type ModbusPDUReadFifoQueueRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetFifoPointerAddress returns FifoPointerAddress (property field)
	GetFifoPointerAddress() uint16
}

// ModbusPDUReadFifoQueueRequestExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUReadFifoQueueRequest.
// This is useful for switch cases.
type ModbusPDUReadFifoQueueRequestExactly interface {
	ModbusPDUReadFifoQueueRequest
	isModbusPDUReadFifoQueueRequest() bool
}

// _ModbusPDUReadFifoQueueRequest is the data-structure of this message
type _ModbusPDUReadFifoQueueRequest struct {
	ModbusPDUContract
	FifoPointerAddress uint16
}

var _ ModbusPDUReadFifoQueueRequest = (*_ModbusPDUReadFifoQueueRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReadFifoQueueRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUReadFifoQueueRequest) GetFunctionFlag() uint8 {
	return 0x18
}

func (m *_ModbusPDUReadFifoQueueRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReadFifoQueueRequest) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReadFifoQueueRequest) GetFifoPointerAddress() uint16 {
	return m.FifoPointerAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUReadFifoQueueRequest factory function for _ModbusPDUReadFifoQueueRequest
func NewModbusPDUReadFifoQueueRequest(fifoPointerAddress uint16) *_ModbusPDUReadFifoQueueRequest {
	_result := &_ModbusPDUReadFifoQueueRequest{
		ModbusPDUContract:  NewModbusPDU(),
		FifoPointerAddress: fifoPointerAddress,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUReadFifoQueueRequest(structType any) ModbusPDUReadFifoQueueRequest {
	if casted, ok := structType.(ModbusPDUReadFifoQueueRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadFifoQueueRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadFifoQueueRequest) GetTypeName() string {
	return "ModbusPDUReadFifoQueueRequest"
}

func (m *_ModbusPDUReadFifoQueueRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	// Simple field (fifoPointerAddress)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ModbusPDUReadFifoQueueRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUReadFifoQueueRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (_modbusPDUReadFifoQueueRequest ModbusPDUReadFifoQueueRequest, err error) {
	m.ModbusPDUContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadFifoQueueRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadFifoQueueRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	fifoPointerAddress, err := ReadSimpleField(ctx, "fifoPointerAddress", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fifoPointerAddress' field"))
	}
	m.FifoPointerAddress = fifoPointerAddress

	if closeErr := readBuffer.CloseContext("ModbusPDUReadFifoQueueRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadFifoQueueRequest")
	}

	return m, nil
}

func (m *_ModbusPDUReadFifoQueueRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReadFifoQueueRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadFifoQueueRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadFifoQueueRequest")
		}

		if err := WriteSimpleField[uint16](ctx, "fifoPointerAddress", m.GetFifoPointerAddress(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'fifoPointerAddress' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadFifoQueueRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReadFifoQueueRequest")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUReadFifoQueueRequest) isModbusPDUReadFifoQueueRequest() bool {
	return true
}

func (m *_ModbusPDUReadFifoQueueRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
