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

// SysexCommandAnalogMappingQueryResponse is the corresponding interface of SysexCommandAnalogMappingQueryResponse
type SysexCommandAnalogMappingQueryResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SysexCommand
	// GetPin returns Pin (property field)
	GetPin() uint8
	// IsSysexCommandAnalogMappingQueryResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSysexCommandAnalogMappingQueryResponse()
}

// _SysexCommandAnalogMappingQueryResponse is the data-structure of this message
type _SysexCommandAnalogMappingQueryResponse struct {
	SysexCommandContract
	Pin uint8
}

var _ SysexCommandAnalogMappingQueryResponse = (*_SysexCommandAnalogMappingQueryResponse)(nil)
var _ SysexCommandRequirements = (*_SysexCommandAnalogMappingQueryResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SysexCommandAnalogMappingQueryResponse) GetCommandType() uint8 {
	return 0x69
}

func (m *_SysexCommandAnalogMappingQueryResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SysexCommandAnalogMappingQueryResponse) GetParent() SysexCommandContract {
	return m.SysexCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SysexCommandAnalogMappingQueryResponse) GetPin() uint8 {
	return m.Pin
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSysexCommandAnalogMappingQueryResponse factory function for _SysexCommandAnalogMappingQueryResponse
func NewSysexCommandAnalogMappingQueryResponse(pin uint8) *_SysexCommandAnalogMappingQueryResponse {
	_result := &_SysexCommandAnalogMappingQueryResponse{
		SysexCommandContract: NewSysexCommand(),
		Pin:                  pin,
	}
	_result.SysexCommandContract.(*_SysexCommand)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSysexCommandAnalogMappingQueryResponse(structType any) SysexCommandAnalogMappingQueryResponse {
	if casted, ok := structType.(SysexCommandAnalogMappingQueryResponse); ok {
		return casted
	}
	if casted, ok := structType.(*SysexCommandAnalogMappingQueryResponse); ok {
		return *casted
	}
	return nil
}

func (m *_SysexCommandAnalogMappingQueryResponse) GetTypeName() string {
	return "SysexCommandAnalogMappingQueryResponse"
}

func (m *_SysexCommandAnalogMappingQueryResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SysexCommandContract.(*_SysexCommand).getLengthInBits(ctx))

	// Simple field (pin)
	lengthInBits += 8

	return lengthInBits
}

func (m *_SysexCommandAnalogMappingQueryResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SysexCommandAnalogMappingQueryResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SysexCommand, response bool) (__sysexCommandAnalogMappingQueryResponse SysexCommandAnalogMappingQueryResponse, err error) {
	m.SysexCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommandAnalogMappingQueryResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandAnalogMappingQueryResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	pin, err := ReadSimpleField(ctx, "pin", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'pin' field"))
	}
	m.Pin = pin

	if closeErr := readBuffer.CloseContext("SysexCommandAnalogMappingQueryResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandAnalogMappingQueryResponse")
	}

	return m, nil
}

func (m *_SysexCommandAnalogMappingQueryResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SysexCommandAnalogMappingQueryResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandAnalogMappingQueryResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandAnalogMappingQueryResponse")
		}

		if err := WriteSimpleField[uint8](ctx, "pin", m.GetPin(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'pin' field")
		}

		if popErr := writeBuffer.PopContext("SysexCommandAnalogMappingQueryResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandAnalogMappingQueryResponse")
		}
		return nil
	}
	return m.SysexCommandContract.(*_SysexCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SysexCommandAnalogMappingQueryResponse) IsSysexCommandAnalogMappingQueryResponse() {}

func (m *_SysexCommandAnalogMappingQueryResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
