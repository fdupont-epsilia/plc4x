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

// SysexCommandPinStateQuery is the corresponding interface of SysexCommandPinStateQuery
type SysexCommandPinStateQuery interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SysexCommand
	// GetPin returns Pin (property field)
	GetPin() uint8
}

// SysexCommandPinStateQueryExactly can be used when we want exactly this type and not a type which fulfills SysexCommandPinStateQuery.
// This is useful for switch cases.
type SysexCommandPinStateQueryExactly interface {
	SysexCommandPinStateQuery
	isSysexCommandPinStateQuery() bool
}

// _SysexCommandPinStateQuery is the data-structure of this message
type _SysexCommandPinStateQuery struct {
	SysexCommandContract
	Pin uint8
}

var _ SysexCommandPinStateQuery = (*_SysexCommandPinStateQuery)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SysexCommandPinStateQuery) GetCommandType() uint8 {
	return 0x6D
}

func (m *_SysexCommandPinStateQuery) GetResponse() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SysexCommandPinStateQuery) GetParent() SysexCommandContract {
	return m.SysexCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SysexCommandPinStateQuery) GetPin() uint8 {
	return m.Pin
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSysexCommandPinStateQuery factory function for _SysexCommandPinStateQuery
func NewSysexCommandPinStateQuery(pin uint8) *_SysexCommandPinStateQuery {
	_result := &_SysexCommandPinStateQuery{
		SysexCommandContract: NewSysexCommand(),
		Pin:                  pin,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastSysexCommandPinStateQuery(structType any) SysexCommandPinStateQuery {
	if casted, ok := structType.(SysexCommandPinStateQuery); ok {
		return casted
	}
	if casted, ok := structType.(*SysexCommandPinStateQuery); ok {
		return *casted
	}
	return nil
}

func (m *_SysexCommandPinStateQuery) GetTypeName() string {
	return "SysexCommandPinStateQuery"
}

func (m *_SysexCommandPinStateQuery) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SysexCommandContract.(*_SysexCommand).getLengthInBits(ctx))

	// Simple field (pin)
	lengthInBits += 8

	return lengthInBits
}

func (m *_SysexCommandPinStateQuery) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SysexCommandPinStateQuery) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SysexCommand, response bool) (_sysexCommandPinStateQuery SysexCommandPinStateQuery, err error) {
	m.SysexCommandContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommandPinStateQuery"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandPinStateQuery")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	pin, err := ReadSimpleField(ctx, "pin", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'pin' field"))
	}
	m.Pin = pin

	if closeErr := readBuffer.CloseContext("SysexCommandPinStateQuery"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandPinStateQuery")
	}

	return m, nil
}

func (m *_SysexCommandPinStateQuery) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SysexCommandPinStateQuery) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandPinStateQuery"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandPinStateQuery")
		}

		if err := WriteSimpleField[uint8](ctx, "pin", m.GetPin(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'pin' field")
		}

		if popErr := writeBuffer.PopContext("SysexCommandPinStateQuery"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandPinStateQuery")
		}
		return nil
	}
	return m.SysexCommandContract.(*_SysexCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SysexCommandPinStateQuery) isSysexCommandPinStateQuery() bool {
	return true
}

func (m *_SysexCommandPinStateQuery) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
