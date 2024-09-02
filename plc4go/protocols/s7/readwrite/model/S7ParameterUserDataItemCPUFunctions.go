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

// S7ParameterUserDataItemCPUFunctions is the corresponding interface of S7ParameterUserDataItemCPUFunctions
type S7ParameterUserDataItemCPUFunctions interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7ParameterUserDataItem
	// GetMethod returns Method (property field)
	GetMethod() uint8
	// GetCpuFunctionType returns CpuFunctionType (property field)
	GetCpuFunctionType() uint8
	// GetCpuFunctionGroup returns CpuFunctionGroup (property field)
	GetCpuFunctionGroup() uint8
	// GetCpuSubfunction returns CpuSubfunction (property field)
	GetCpuSubfunction() uint8
	// GetSequenceNumber returns SequenceNumber (property field)
	GetSequenceNumber() uint8
	// GetDataUnitReferenceNumber returns DataUnitReferenceNumber (property field)
	GetDataUnitReferenceNumber() *uint8
	// GetLastDataUnit returns LastDataUnit (property field)
	GetLastDataUnit() *uint8
	// GetErrorCode returns ErrorCode (property field)
	GetErrorCode() *uint16
	// IsS7ParameterUserDataItemCPUFunctions is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7ParameterUserDataItemCPUFunctions()
}

// _S7ParameterUserDataItemCPUFunctions is the data-structure of this message
type _S7ParameterUserDataItemCPUFunctions struct {
	S7ParameterUserDataItemContract
	Method                  uint8
	CpuFunctionType         uint8
	CpuFunctionGroup        uint8
	CpuSubfunction          uint8
	SequenceNumber          uint8
	DataUnitReferenceNumber *uint8
	LastDataUnit            *uint8
	ErrorCode               *uint16
}

var _ S7ParameterUserDataItemCPUFunctions = (*_S7ParameterUserDataItemCPUFunctions)(nil)
var _ S7ParameterUserDataItemRequirements = (*_S7ParameterUserDataItemCPUFunctions)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7ParameterUserDataItemCPUFunctions) GetItemType() uint8 {
	return 0x12
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7ParameterUserDataItemCPUFunctions) GetParent() S7ParameterUserDataItemContract {
	return m.S7ParameterUserDataItemContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7ParameterUserDataItemCPUFunctions) GetMethod() uint8 {
	return m.Method
}

func (m *_S7ParameterUserDataItemCPUFunctions) GetCpuFunctionType() uint8 {
	return m.CpuFunctionType
}

func (m *_S7ParameterUserDataItemCPUFunctions) GetCpuFunctionGroup() uint8 {
	return m.CpuFunctionGroup
}

func (m *_S7ParameterUserDataItemCPUFunctions) GetCpuSubfunction() uint8 {
	return m.CpuSubfunction
}

func (m *_S7ParameterUserDataItemCPUFunctions) GetSequenceNumber() uint8 {
	return m.SequenceNumber
}

func (m *_S7ParameterUserDataItemCPUFunctions) GetDataUnitReferenceNumber() *uint8 {
	return m.DataUnitReferenceNumber
}

func (m *_S7ParameterUserDataItemCPUFunctions) GetLastDataUnit() *uint8 {
	return m.LastDataUnit
}

func (m *_S7ParameterUserDataItemCPUFunctions) GetErrorCode() *uint16 {
	return m.ErrorCode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7ParameterUserDataItemCPUFunctions factory function for _S7ParameterUserDataItemCPUFunctions
func NewS7ParameterUserDataItemCPUFunctions(method uint8, cpuFunctionType uint8, cpuFunctionGroup uint8, cpuSubfunction uint8, sequenceNumber uint8, dataUnitReferenceNumber *uint8, lastDataUnit *uint8, errorCode *uint16) *_S7ParameterUserDataItemCPUFunctions {
	_result := &_S7ParameterUserDataItemCPUFunctions{
		S7ParameterUserDataItemContract: NewS7ParameterUserDataItem(),
		Method:                          method,
		CpuFunctionType:                 cpuFunctionType,
		CpuFunctionGroup:                cpuFunctionGroup,
		CpuSubfunction:                  cpuSubfunction,
		SequenceNumber:                  sequenceNumber,
		DataUnitReferenceNumber:         dataUnitReferenceNumber,
		LastDataUnit:                    lastDataUnit,
		ErrorCode:                       errorCode,
	}
	_result.S7ParameterUserDataItemContract.(*_S7ParameterUserDataItem)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7ParameterUserDataItemCPUFunctions(structType any) S7ParameterUserDataItemCPUFunctions {
	if casted, ok := structType.(S7ParameterUserDataItemCPUFunctions); ok {
		return casted
	}
	if casted, ok := structType.(*S7ParameterUserDataItemCPUFunctions); ok {
		return *casted
	}
	return nil
}

func (m *_S7ParameterUserDataItemCPUFunctions) GetTypeName() string {
	return "S7ParameterUserDataItemCPUFunctions"
}

func (m *_S7ParameterUserDataItemCPUFunctions) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7ParameterUserDataItemContract.(*_S7ParameterUserDataItem).getLengthInBits(ctx))

	// Implicit Field (itemLength)
	lengthInBits += 8

	// Simple field (method)
	lengthInBits += 8

	// Simple field (cpuFunctionType)
	lengthInBits += 4

	// Simple field (cpuFunctionGroup)
	lengthInBits += 4

	// Simple field (cpuSubfunction)
	lengthInBits += 8

	// Simple field (sequenceNumber)
	lengthInBits += 8

	// Optional Field (dataUnitReferenceNumber)
	if m.DataUnitReferenceNumber != nil {
		lengthInBits += 8
	}

	// Optional Field (lastDataUnit)
	if m.LastDataUnit != nil {
		lengthInBits += 8
	}

	// Optional Field (errorCode)
	if m.ErrorCode != nil {
		lengthInBits += 16
	}

	return lengthInBits
}

func (m *_S7ParameterUserDataItemCPUFunctions) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7ParameterUserDataItemCPUFunctions) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7ParameterUserDataItem) (__s7ParameterUserDataItemCPUFunctions S7ParameterUserDataItemCPUFunctions, err error) {
	m.S7ParameterUserDataItemContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7ParameterUserDataItemCPUFunctions"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7ParameterUserDataItemCPUFunctions")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	itemLength, err := ReadImplicitField[uint8](ctx, "itemLength", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemLength' field"))
	}
	_ = itemLength

	method, err := ReadSimpleField(ctx, "method", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'method' field"))
	}
	m.Method = method

	cpuFunctionType, err := ReadSimpleField(ctx, "cpuFunctionType", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'cpuFunctionType' field"))
	}
	m.CpuFunctionType = cpuFunctionType

	cpuFunctionGroup, err := ReadSimpleField(ctx, "cpuFunctionGroup", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'cpuFunctionGroup' field"))
	}
	m.CpuFunctionGroup = cpuFunctionGroup

	cpuSubfunction, err := ReadSimpleField(ctx, "cpuSubfunction", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'cpuSubfunction' field"))
	}
	m.CpuSubfunction = cpuSubfunction

	sequenceNumber, err := ReadSimpleField(ctx, "sequenceNumber", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sequenceNumber' field"))
	}
	m.SequenceNumber = sequenceNumber

	var dataUnitReferenceNumber *uint8
	dataUnitReferenceNumber, err = ReadOptionalField[uint8](ctx, "dataUnitReferenceNumber", ReadUnsignedByte(readBuffer, uint8(8)), bool((bool((cpuFunctionType) == (8)))) || bool((bool((bool((cpuFunctionType) == (0)))) && bool((bool((cpuFunctionGroup) == (2)))))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataUnitReferenceNumber' field"))
	}
	m.DataUnitReferenceNumber = dataUnitReferenceNumber

	var lastDataUnit *uint8
	lastDataUnit, err = ReadOptionalField[uint8](ctx, "lastDataUnit", ReadUnsignedByte(readBuffer, uint8(8)), bool((bool((cpuFunctionType) == (8)))) || bool((bool((bool((cpuFunctionType) == (0)))) && bool((bool((cpuFunctionGroup) == (2)))))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastDataUnit' field"))
	}
	m.LastDataUnit = lastDataUnit

	var errorCode *uint16
	errorCode, err = ReadOptionalField[uint16](ctx, "errorCode", ReadUnsignedShort(readBuffer, uint8(16)), bool((bool((cpuFunctionType) == (8)))) || bool((bool((bool((cpuFunctionType) == (0)))) && bool((bool((cpuFunctionGroup) == (2)))))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'errorCode' field"))
	}
	m.ErrorCode = errorCode

	if closeErr := readBuffer.CloseContext("S7ParameterUserDataItemCPUFunctions"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7ParameterUserDataItemCPUFunctions")
	}

	return m, nil
}

func (m *_S7ParameterUserDataItemCPUFunctions) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7ParameterUserDataItemCPUFunctions) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7ParameterUserDataItemCPUFunctions"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7ParameterUserDataItemCPUFunctions")
		}
		itemLength := uint8(uint8(uint8(m.GetLengthInBytes(ctx))) - uint8(uint8(2)))
		if err := WriteImplicitField(ctx, "itemLength", itemLength, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'itemLength' field")
		}

		if err := WriteSimpleField[uint8](ctx, "method", m.GetMethod(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'method' field")
		}

		if err := WriteSimpleField[uint8](ctx, "cpuFunctionType", m.GetCpuFunctionType(), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'cpuFunctionType' field")
		}

		if err := WriteSimpleField[uint8](ctx, "cpuFunctionGroup", m.GetCpuFunctionGroup(), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'cpuFunctionGroup' field")
		}

		if err := WriteSimpleField[uint8](ctx, "cpuSubfunction", m.GetCpuSubfunction(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'cpuSubfunction' field")
		}

		if err := WriteSimpleField[uint8](ctx, "sequenceNumber", m.GetSequenceNumber(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'sequenceNumber' field")
		}

		if err := WriteOptionalField[uint8](ctx, "dataUnitReferenceNumber", m.GetDataUnitReferenceNumber(), WriteUnsignedByte(writeBuffer, 8), true); err != nil {
			return errors.Wrap(err, "Error serializing 'dataUnitReferenceNumber' field")
		}

		if err := WriteOptionalField[uint8](ctx, "lastDataUnit", m.GetLastDataUnit(), WriteUnsignedByte(writeBuffer, 8), true); err != nil {
			return errors.Wrap(err, "Error serializing 'lastDataUnit' field")
		}

		if err := WriteOptionalField[uint16](ctx, "errorCode", m.GetErrorCode(), WriteUnsignedShort(writeBuffer, 16), true); err != nil {
			return errors.Wrap(err, "Error serializing 'errorCode' field")
		}

		if popErr := writeBuffer.PopContext("S7ParameterUserDataItemCPUFunctions"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7ParameterUserDataItemCPUFunctions")
		}
		return nil
	}
	return m.S7ParameterUserDataItemContract.(*_S7ParameterUserDataItem).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7ParameterUserDataItemCPUFunctions) IsS7ParameterUserDataItemCPUFunctions() {}

func (m *_S7ParameterUserDataItemCPUFunctions) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
