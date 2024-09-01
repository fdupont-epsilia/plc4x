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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// FirmataMessageAnalogIO is the corresponding interface of FirmataMessageAnalogIO
type FirmataMessageAnalogIO interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	FirmataMessage
	// GetPin returns Pin (property field)
	GetPin() uint8
	// GetData returns Data (property field)
	GetData() []int8
}

// FirmataMessageAnalogIOExactly can be used when we want exactly this type and not a type which fulfills FirmataMessageAnalogIO.
// This is useful for switch cases.
type FirmataMessageAnalogIOExactly interface {
	FirmataMessageAnalogIO
	isFirmataMessageAnalogIO() bool
}

// _FirmataMessageAnalogIO is the data-structure of this message
type _FirmataMessageAnalogIO struct {
	*_FirmataMessage
	Pin  uint8
	Data []int8
}

var _ FirmataMessageAnalogIO = (*_FirmataMessageAnalogIO)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FirmataMessageAnalogIO) GetMessageType() uint8 {
	return 0xE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FirmataMessageAnalogIO) InitializeParent(parent FirmataMessage) {}

func (m *_FirmataMessageAnalogIO) GetParent() FirmataMessage {
	return m._FirmataMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_FirmataMessageAnalogIO) GetPin() uint8 {
	return m.Pin
}

func (m *_FirmataMessageAnalogIO) GetData() []int8 {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewFirmataMessageAnalogIO factory function for _FirmataMessageAnalogIO
func NewFirmataMessageAnalogIO(pin uint8, data []int8, response bool) *_FirmataMessageAnalogIO {
	_result := &_FirmataMessageAnalogIO{
		Pin:             pin,
		Data:            data,
		_FirmataMessage: NewFirmataMessage(response),
	}
	_result._FirmataMessage._FirmataMessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastFirmataMessageAnalogIO(structType any) FirmataMessageAnalogIO {
	if casted, ok := structType.(FirmataMessageAnalogIO); ok {
		return casted
	}
	if casted, ok := structType.(*FirmataMessageAnalogIO); ok {
		return *casted
	}
	return nil
}

func (m *_FirmataMessageAnalogIO) GetTypeName() string {
	return "FirmataMessageAnalogIO"
}

func (m *_FirmataMessageAnalogIO) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (pin)
	lengthInBits += 4

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_FirmataMessageAnalogIO) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func FirmataMessageAnalogIOParse(ctx context.Context, theBytes []byte, response bool) (FirmataMessageAnalogIO, error) {
	return FirmataMessageAnalogIOParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), response)
}

func FirmataMessageAnalogIOParseWithBufferProducer(response bool) func(ctx context.Context, readBuffer utils.ReadBuffer) (FirmataMessageAnalogIO, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (FirmataMessageAnalogIO, error) {
		return FirmataMessageAnalogIOParseWithBuffer(ctx, readBuffer, response)
	}
}

func FirmataMessageAnalogIOParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (FirmataMessageAnalogIO, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FirmataMessageAnalogIO"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FirmataMessageAnalogIO")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	pin, err := ReadSimpleField(ctx, "pin", ReadUnsignedByte(readBuffer, uint8(4)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'pin' field"))
	}

	data, err := ReadCountArrayField[int8](ctx, "data", ReadSignedByte(readBuffer, uint8(8)), uint64(int32(2)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}

	if closeErr := readBuffer.CloseContext("FirmataMessageAnalogIO"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FirmataMessageAnalogIO")
	}

	// Create a partially initialized instance
	_child := &_FirmataMessageAnalogIO{
		_FirmataMessage: &_FirmataMessage{
			Response: response,
		},
		Pin:  pin,
		Data: data,
	}
	_child._FirmataMessage._FirmataMessageChildRequirements = _child
	return _child, nil
}

func (m *_FirmataMessageAnalogIO) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FirmataMessageAnalogIO) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FirmataMessageAnalogIO"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FirmataMessageAnalogIO")
		}

		if err := WriteSimpleField[uint8](ctx, "pin", m.GetPin(), WriteUnsignedByte(writeBuffer, 4), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'pin' field")
		}

		if err := WriteSimpleTypeArrayField(ctx, "data", m.GetData(), WriteSignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("FirmataMessageAnalogIO"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FirmataMessageAnalogIO")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_FirmataMessageAnalogIO) isFirmataMessageAnalogIO() bool {
	return true
}

func (m *_FirmataMessageAnalogIO) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
