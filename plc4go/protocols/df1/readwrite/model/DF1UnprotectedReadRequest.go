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

// DF1UnprotectedReadRequest is the corresponding interface of DF1UnprotectedReadRequest
type DF1UnprotectedReadRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	DF1Command
	// GetAddress returns Address (property field)
	GetAddress() uint16
	// GetSize returns Size (property field)
	GetSize() uint8
}

// DF1UnprotectedReadRequestExactly can be used when we want exactly this type and not a type which fulfills DF1UnprotectedReadRequest.
// This is useful for switch cases.
type DF1UnprotectedReadRequestExactly interface {
	DF1UnprotectedReadRequest
	isDF1UnprotectedReadRequest() bool
}

// _DF1UnprotectedReadRequest is the data-structure of this message
type _DF1UnprotectedReadRequest struct {
	DF1CommandContract
	Address uint16
	Size    uint8
}

var _ DF1UnprotectedReadRequest = (*_DF1UnprotectedReadRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DF1UnprotectedReadRequest) GetCommandCode() uint8 {
	return 0x01
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DF1UnprotectedReadRequest) GetParent() DF1CommandContract {
	return m.DF1CommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DF1UnprotectedReadRequest) GetAddress() uint16 {
	return m.Address
}

func (m *_DF1UnprotectedReadRequest) GetSize() uint8 {
	return m.Size
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDF1UnprotectedReadRequest factory function for _DF1UnprotectedReadRequest
func NewDF1UnprotectedReadRequest(address uint16, size uint8, status uint8, transactionCounter uint16) *_DF1UnprotectedReadRequest {
	_result := &_DF1UnprotectedReadRequest{
		DF1CommandContract: NewDF1Command(status, transactionCounter),
		Address:            address,
		Size:               size,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastDF1UnprotectedReadRequest(structType any) DF1UnprotectedReadRequest {
	if casted, ok := structType.(DF1UnprotectedReadRequest); ok {
		return casted
	}
	if casted, ok := structType.(*DF1UnprotectedReadRequest); ok {
		return *casted
	}
	return nil
}

func (m *_DF1UnprotectedReadRequest) GetTypeName() string {
	return "DF1UnprotectedReadRequest"
}

func (m *_DF1UnprotectedReadRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.DF1CommandContract.(*_DF1Command).getLengthInBits(ctx))

	// Simple field (address)
	lengthInBits += 16

	// Simple field (size)
	lengthInBits += 8

	return lengthInBits
}

func (m *_DF1UnprotectedReadRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DF1UnprotectedReadRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_DF1Command) (_dF1UnprotectedReadRequest DF1UnprotectedReadRequest, err error) {
	m.DF1CommandContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DF1UnprotectedReadRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DF1UnprotectedReadRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	address, err := ReadSimpleField(ctx, "address", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'address' field"))
	}
	m.Address = address

	size, err := ReadSimpleField(ctx, "size", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'size' field"))
	}
	m.Size = size

	if closeErr := readBuffer.CloseContext("DF1UnprotectedReadRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DF1UnprotectedReadRequest")
	}

	return m, nil
}

func (m *_DF1UnprotectedReadRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DF1UnprotectedReadRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DF1UnprotectedReadRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DF1UnprotectedReadRequest")
		}

		if err := WriteSimpleField[uint16](ctx, "address", m.GetAddress(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'address' field")
		}

		if err := WriteSimpleField[uint8](ctx, "size", m.GetSize(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'size' field")
		}

		if popErr := writeBuffer.PopContext("DF1UnprotectedReadRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DF1UnprotectedReadRequest")
		}
		return nil
	}
	return m.DF1CommandContract.(*_DF1Command).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DF1UnprotectedReadRequest) isDF1UnprotectedReadRequest() bool {
	return true
}

func (m *_DF1UnprotectedReadRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
