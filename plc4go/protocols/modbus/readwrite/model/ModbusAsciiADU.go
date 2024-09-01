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
	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// ModbusAsciiADU is the corresponding interface of ModbusAsciiADU
type ModbusAsciiADU interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusADU
	// GetAddress returns Address (property field)
	GetAddress() uint8
	// GetPdu returns Pdu (property field)
	GetPdu() ModbusPDU
}

// ModbusAsciiADUExactly can be used when we want exactly this type and not a type which fulfills ModbusAsciiADU.
// This is useful for switch cases.
type ModbusAsciiADUExactly interface {
	ModbusAsciiADU
	isModbusAsciiADU() bool
}

// _ModbusAsciiADU is the data-structure of this message
type _ModbusAsciiADU struct {
	ModbusADUContract
	Address uint8
	Pdu     ModbusPDU
}

var _ ModbusAsciiADU = (*_ModbusAsciiADU)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusAsciiADU) GetDriverType() DriverType {
	return DriverType_MODBUS_ASCII
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusAsciiADU) GetParent() ModbusADUContract {
	return m.ModbusADUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusAsciiADU) GetAddress() uint8 {
	return m.Address
}

func (m *_ModbusAsciiADU) GetPdu() ModbusPDU {
	return m.Pdu
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusAsciiADU factory function for _ModbusAsciiADU
func NewModbusAsciiADU(address uint8, pdu ModbusPDU, response bool) *_ModbusAsciiADU {
	_result := &_ModbusAsciiADU{
		ModbusADUContract: NewModbusADU(response),
		Address:           address,
		Pdu:               pdu,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusAsciiADU(structType any) ModbusAsciiADU {
	if casted, ok := structType.(ModbusAsciiADU); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusAsciiADU); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusAsciiADU) GetTypeName() string {
	return "ModbusAsciiADU"
}

func (m *_ModbusAsciiADU) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusADUContract.(*_ModbusADU).getLengthInBits(ctx))

	// Simple field (address)
	lengthInBits += 8

	// Simple field (pdu)
	lengthInBits += m.Pdu.GetLengthInBits(ctx)

	// Checksum Field (checksum)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ModbusAsciiADU) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusAsciiADU) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusADU, driverType DriverType, response bool) (_modbusAsciiADU ModbusAsciiADU, err error) {
	m.ModbusADUContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusAsciiADU"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusAsciiADU")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	address, err := ReadSimpleField(ctx, "address", ReadUnsignedByte(readBuffer, uint8(8)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'address' field"))
	}
	m.Address = address

	pdu, err := ReadSimpleField[ModbusPDU](ctx, "pdu", ReadComplex[ModbusPDU](ModbusPDUParseWithBufferProducer[ModbusPDU]((bool)(response)), readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'pdu' field"))
	}
	m.Pdu = pdu

	crc, err := ReadChecksumField[uint8](ctx, "crc", ReadUnsignedByte(readBuffer, uint8(8)), AsciiLrcCheck(ctx, address, pdu), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'crc' field"))
	}
	_ = crc

	if closeErr := readBuffer.CloseContext("ModbusAsciiADU"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusAsciiADU")
	}

	return m, nil
}

func (m *_ModbusAsciiADU) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusAsciiADU) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusAsciiADU"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusAsciiADU")
		}

		if err := WriteSimpleField[uint8](ctx, "address", m.GetAddress(), WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'address' field")
		}

		if err := WriteSimpleField[ModbusPDU](ctx, "pdu", m.GetPdu(), WriteComplex[ModbusPDU](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'pdu' field")
		}

		if err := WriteChecksumField[uint8](ctx, "crc", AsciiLrcCheck(ctx, m.GetAddress(), m.GetPdu()), WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'crc' field")
		}

		if popErr := writeBuffer.PopContext("ModbusAsciiADU"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusAsciiADU")
		}
		return nil
	}
	return m.ModbusADUContract.(*_ModbusADU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusAsciiADU) isModbusAsciiADU() bool {
	return true
}

func (m *_ModbusAsciiADU) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
