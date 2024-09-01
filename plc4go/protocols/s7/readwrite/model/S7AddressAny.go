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

// S7AddressAny is the corresponding interface of S7AddressAny
type S7AddressAny interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7Address
	// GetTransportSize returns TransportSize (property field)
	GetTransportSize() TransportSize
	// GetNumberOfElements returns NumberOfElements (property field)
	GetNumberOfElements() uint16
	// GetDbNumber returns DbNumber (property field)
	GetDbNumber() uint16
	// GetArea returns Area (property field)
	GetArea() MemoryArea
	// GetByteAddress returns ByteAddress (property field)
	GetByteAddress() uint16
	// GetBitAddress returns BitAddress (property field)
	GetBitAddress() uint8
}

// S7AddressAnyExactly can be used when we want exactly this type and not a type which fulfills S7AddressAny.
// This is useful for switch cases.
type S7AddressAnyExactly interface {
	S7AddressAny
	isS7AddressAny() bool
}

// _S7AddressAny is the data-structure of this message
type _S7AddressAny struct {
	*_S7Address
	TransportSize    TransportSize
	NumberOfElements uint16
	DbNumber         uint16
	Area             MemoryArea
	ByteAddress      uint16
	BitAddress       uint8
	// Reserved Fields
	reservedField0 *uint8
}

var _ S7AddressAny = (*_S7AddressAny)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7AddressAny) GetAddressType() uint8 {
	return 0x10
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7AddressAny) InitializeParent(parent S7Address) {}

func (m *_S7AddressAny) GetParent() S7Address {
	return m._S7Address
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7AddressAny) GetTransportSize() TransportSize {
	return m.TransportSize
}

func (m *_S7AddressAny) GetNumberOfElements() uint16 {
	return m.NumberOfElements
}

func (m *_S7AddressAny) GetDbNumber() uint16 {
	return m.DbNumber
}

func (m *_S7AddressAny) GetArea() MemoryArea {
	return m.Area
}

func (m *_S7AddressAny) GetByteAddress() uint16 {
	return m.ByteAddress
}

func (m *_S7AddressAny) GetBitAddress() uint8 {
	return m.BitAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7AddressAny factory function for _S7AddressAny
func NewS7AddressAny(transportSize TransportSize, numberOfElements uint16, dbNumber uint16, area MemoryArea, byteAddress uint16, bitAddress uint8) *_S7AddressAny {
	_result := &_S7AddressAny{
		TransportSize:    transportSize,
		NumberOfElements: numberOfElements,
		DbNumber:         dbNumber,
		Area:             area,
		ByteAddress:      byteAddress,
		BitAddress:       bitAddress,
		_S7Address:       NewS7Address(),
	}
	_result._S7Address._S7AddressChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7AddressAny(structType any) S7AddressAny {
	if casted, ok := structType.(S7AddressAny); ok {
		return casted
	}
	if casted, ok := structType.(*S7AddressAny); ok {
		return *casted
	}
	return nil
}

func (m *_S7AddressAny) GetTypeName() string {
	return "S7AddressAny"
}

func (m *_S7AddressAny) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Enum Field (transportSize)
	lengthInBits += 8

	// Simple field (numberOfElements)
	lengthInBits += 16

	// Simple field (dbNumber)
	lengthInBits += 16

	// Simple field (area)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 5

	// Simple field (byteAddress)
	lengthInBits += 16

	// Simple field (bitAddress)
	lengthInBits += 3

	return lengthInBits
}

func (m *_S7AddressAny) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7AddressAnyParse(ctx context.Context, theBytes []byte) (S7AddressAny, error) {
	return S7AddressAnyParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func S7AddressAnyParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (S7AddressAny, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (S7AddressAny, error) {
		return S7AddressAnyParseWithBuffer(ctx, readBuffer)
	}
}

func S7AddressAnyParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (S7AddressAny, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7AddressAny"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7AddressAny")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	transportSize, err := ReadEnumField[TransportSize](ctx, "transportSize", "TransportSize", ReadEnum[TransportSize, uint8](TransportSizeFirstEnumForFieldCode, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transportSize' field"))
	}

	numberOfElements, err := ReadSimpleField(ctx, "numberOfElements", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfElements' field"))
	}

	dbNumber, err := ReadSimpleField(ctx, "dbNumber", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dbNumber' field"))
	}

	area, err := ReadEnumField[MemoryArea](ctx, "area", "MemoryArea", ReadEnum(MemoryAreaByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'area' field"))
	}

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(5)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	byteAddress, err := ReadSimpleField(ctx, "byteAddress", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'byteAddress' field"))
	}

	bitAddress, err := ReadSimpleField(ctx, "bitAddress", ReadUnsignedByte(readBuffer, uint8(3)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bitAddress' field"))
	}

	if closeErr := readBuffer.CloseContext("S7AddressAny"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7AddressAny")
	}

	// Create a partially initialized instance
	_child := &_S7AddressAny{
		_S7Address:       &_S7Address{},
		TransportSize:    transportSize,
		NumberOfElements: numberOfElements,
		DbNumber:         dbNumber,
		Area:             area,
		ByteAddress:      byteAddress,
		BitAddress:       bitAddress,
		reservedField0:   reservedField0,
	}
	_child._S7Address._S7AddressChildRequirements = _child
	return _child, nil
}

func (m *_S7AddressAny) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7AddressAny) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7AddressAny"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7AddressAny")
		}

		if err := WriteEnumField(ctx, "transportSize", "TransportSize", m.GetTransportSize(), WriteEnum[TransportSize, uint8](TransportSize.GetCode, TransportSize.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'transportSize' field")
		}

		if err := WriteSimpleField[uint16](ctx, "numberOfElements", m.GetNumberOfElements(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'numberOfElements' field")
		}

		if err := WriteSimpleField[uint16](ctx, "dbNumber", m.GetDbNumber(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'dbNumber' field")
		}

		if err := WriteSimpleEnumField[MemoryArea](ctx, "area", "MemoryArea", m.GetArea(), WriteEnum[MemoryArea, uint8](MemoryArea.GetValue, MemoryArea.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'area' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 5)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[uint16](ctx, "byteAddress", m.GetByteAddress(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'byteAddress' field")
		}

		if err := WriteSimpleField[uint8](ctx, "bitAddress", m.GetBitAddress(), WriteUnsignedByte(writeBuffer, 3)); err != nil {
			return errors.Wrap(err, "Error serializing 'bitAddress' field")
		}

		if popErr := writeBuffer.PopContext("S7AddressAny"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7AddressAny")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7AddressAny) isS7AddressAny() bool {
	return true
}

func (m *_S7AddressAny) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
