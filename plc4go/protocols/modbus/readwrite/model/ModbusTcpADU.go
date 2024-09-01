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

// Constant values.
const ModbusTcpADU_PROTOCOLIDENTIFIER uint16 = 0x0000

// ModbusTcpADU is the corresponding interface of ModbusTcpADU
type ModbusTcpADU interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusADU
	// GetTransactionIdentifier returns TransactionIdentifier (property field)
	GetTransactionIdentifier() uint16
	// GetUnitIdentifier returns UnitIdentifier (property field)
	GetUnitIdentifier() uint8
	// GetPdu returns Pdu (property field)
	GetPdu() ModbusPDU
}

// ModbusTcpADUExactly can be used when we want exactly this type and not a type which fulfills ModbusTcpADU.
// This is useful for switch cases.
type ModbusTcpADUExactly interface {
	ModbusTcpADU
	isModbusTcpADU() bool
}

// _ModbusTcpADU is the data-structure of this message
type _ModbusTcpADU struct {
	*_ModbusADU
	TransactionIdentifier uint16
	UnitIdentifier        uint8
	Pdu                   ModbusPDU
}

var _ ModbusTcpADU = (*_ModbusTcpADU)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusTcpADU) GetDriverType() DriverType {
	return DriverType_MODBUS_TCP
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusTcpADU) InitializeParent(parent ModbusADU) {}

func (m *_ModbusTcpADU) GetParent() ModbusADU {
	return m._ModbusADU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusTcpADU) GetTransactionIdentifier() uint16 {
	return m.TransactionIdentifier
}

func (m *_ModbusTcpADU) GetUnitIdentifier() uint8 {
	return m.UnitIdentifier
}

func (m *_ModbusTcpADU) GetPdu() ModbusPDU {
	return m.Pdu
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_ModbusTcpADU) GetProtocolIdentifier() uint16 {
	return ModbusTcpADU_PROTOCOLIDENTIFIER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusTcpADU factory function for _ModbusTcpADU
func NewModbusTcpADU(transactionIdentifier uint16, unitIdentifier uint8, pdu ModbusPDU, response bool) *_ModbusTcpADU {
	_result := &_ModbusTcpADU{
		TransactionIdentifier: transactionIdentifier,
		UnitIdentifier:        unitIdentifier,
		Pdu:                   pdu,
		_ModbusADU:            NewModbusADU(response),
	}
	_result._ModbusADU._ModbusADUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusTcpADU(structType any) ModbusTcpADU {
	if casted, ok := structType.(ModbusTcpADU); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusTcpADU); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusTcpADU) GetTypeName() string {
	return "ModbusTcpADU"
}

func (m *_ModbusTcpADU) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (transactionIdentifier)
	lengthInBits += 16

	// Const Field (protocolIdentifier)
	lengthInBits += 16

	// Implicit Field (length)
	lengthInBits += 16

	// Simple field (unitIdentifier)
	lengthInBits += 8

	// Simple field (pdu)
	lengthInBits += m.Pdu.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ModbusTcpADU) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusTcpADUParse(ctx context.Context, theBytes []byte, driverType DriverType, response bool) (ModbusTcpADU, error) {
	return ModbusTcpADUParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), driverType, response)
}

func ModbusTcpADUParseWithBufferProducer(driverType DriverType, response bool) func(ctx context.Context, readBuffer utils.ReadBuffer) (ModbusTcpADU, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ModbusTcpADU, error) {
		return ModbusTcpADUParseWithBuffer(ctx, readBuffer, driverType, response)
	}
}

func ModbusTcpADUParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, driverType DriverType, response bool) (ModbusTcpADU, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusTcpADU"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusTcpADU")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	transactionIdentifier, err := ReadSimpleField(ctx, "transactionIdentifier", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transactionIdentifier' field"))
	}

	protocolIdentifier, err := ReadConstField[uint16](ctx, "protocolIdentifier", ReadUnsignedShort(readBuffer, uint8(16)), ModbusTcpADU_PROTOCOLIDENTIFIER, codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'protocolIdentifier' field"))
	}
	_ = protocolIdentifier

	length, err := ReadImplicitField[uint16](ctx, "length", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'length' field"))
	}
	_ = length

	unitIdentifier, err := ReadSimpleField(ctx, "unitIdentifier", ReadUnsignedByte(readBuffer, uint8(8)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unitIdentifier' field"))
	}

	pdu, err := ReadSimpleField[ModbusPDU](ctx, "pdu", ReadComplex[ModbusPDU](ModbusPDUParseWithBufferProducer[ModbusPDU]((bool)(response)), readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'pdu' field"))
	}

	if closeErr := readBuffer.CloseContext("ModbusTcpADU"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusTcpADU")
	}

	// Create a partially initialized instance
	_child := &_ModbusTcpADU{
		_ModbusADU: &_ModbusADU{
			Response: response,
		},
		TransactionIdentifier: transactionIdentifier,
		UnitIdentifier:        unitIdentifier,
		Pdu:                   pdu,
	}
	_child._ModbusADU._ModbusADUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusTcpADU) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusTcpADU) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusTcpADU"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusTcpADU")
		}

		if err := WriteSimpleField[uint16](ctx, "transactionIdentifier", m.GetTransactionIdentifier(), WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'transactionIdentifier' field")
		}

		if err := WriteConstField(ctx, "protocolIdentifier", ModbusTcpADU_PROTOCOLIDENTIFIER, WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'protocolIdentifier' field")
		}
		length := uint16(uint16(m.GetPdu().GetLengthInBytes(ctx)) + uint16(uint16(1)))
		if err := WriteImplicitField(ctx, "length", length, WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'length' field")
		}

		if err := WriteSimpleField[uint8](ctx, "unitIdentifier", m.GetUnitIdentifier(), WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'unitIdentifier' field")
		}

		if err := WriteSimpleField[ModbusPDU](ctx, "pdu", m.GetPdu(), WriteComplex[ModbusPDU](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'pdu' field")
		}

		if popErr := writeBuffer.PopContext("ModbusTcpADU"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusTcpADU")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusTcpADU) isModbusTcpADU() bool {
	return true
}

func (m *_ModbusTcpADU) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
