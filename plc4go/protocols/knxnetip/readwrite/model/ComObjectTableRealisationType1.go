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

// ComObjectTableRealisationType1 is the corresponding interface of ComObjectTableRealisationType1
type ComObjectTableRealisationType1 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ComObjectTable
	// GetNumEntries returns NumEntries (property field)
	GetNumEntries() uint8
	// GetRamFlagsTablePointer returns RamFlagsTablePointer (property field)
	GetRamFlagsTablePointer() uint8
	// GetComObjectDescriptors returns ComObjectDescriptors (property field)
	GetComObjectDescriptors() []GroupObjectDescriptorRealisationType1
}

// ComObjectTableRealisationType1Exactly can be used when we want exactly this type and not a type which fulfills ComObjectTableRealisationType1.
// This is useful for switch cases.
type ComObjectTableRealisationType1Exactly interface {
	ComObjectTableRealisationType1
	isComObjectTableRealisationType1() bool
}

// _ComObjectTableRealisationType1 is the data-structure of this message
type _ComObjectTableRealisationType1 struct {
	ComObjectTableContract
	NumEntries           uint8
	RamFlagsTablePointer uint8
	ComObjectDescriptors []GroupObjectDescriptorRealisationType1
}

var _ ComObjectTableRealisationType1 = (*_ComObjectTableRealisationType1)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ComObjectTableRealisationType1) GetFirmwareType() FirmwareType {
	return FirmwareType_SYSTEM_1
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ComObjectTableRealisationType1) GetParent() ComObjectTableContract {
	return m.ComObjectTableContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ComObjectTableRealisationType1) GetNumEntries() uint8 {
	return m.NumEntries
}

func (m *_ComObjectTableRealisationType1) GetRamFlagsTablePointer() uint8 {
	return m.RamFlagsTablePointer
}

func (m *_ComObjectTableRealisationType1) GetComObjectDescriptors() []GroupObjectDescriptorRealisationType1 {
	return m.ComObjectDescriptors
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewComObjectTableRealisationType1 factory function for _ComObjectTableRealisationType1
func NewComObjectTableRealisationType1(numEntries uint8, ramFlagsTablePointer uint8, comObjectDescriptors []GroupObjectDescriptorRealisationType1) *_ComObjectTableRealisationType1 {
	_result := &_ComObjectTableRealisationType1{
		ComObjectTableContract: NewComObjectTable(),
		NumEntries:             numEntries,
		RamFlagsTablePointer:   ramFlagsTablePointer,
		ComObjectDescriptors:   comObjectDescriptors,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastComObjectTableRealisationType1(structType any) ComObjectTableRealisationType1 {
	if casted, ok := structType.(ComObjectTableRealisationType1); ok {
		return casted
	}
	if casted, ok := structType.(*ComObjectTableRealisationType1); ok {
		return *casted
	}
	return nil
}

func (m *_ComObjectTableRealisationType1) GetTypeName() string {
	return "ComObjectTableRealisationType1"
}

func (m *_ComObjectTableRealisationType1) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ComObjectTableContract.(*_ComObjectTable).getLengthInBits(ctx))

	// Simple field (numEntries)
	lengthInBits += 8

	// Simple field (ramFlagsTablePointer)
	lengthInBits += 8

	// Array field
	if len(m.ComObjectDescriptors) > 0 {
		for _curItem, element := range m.ComObjectDescriptors {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ComObjectDescriptors), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_ComObjectTableRealisationType1) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ComObjectTableRealisationType1) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ComObjectTable, firmwareType FirmwareType) (_comObjectTableRealisationType1 ComObjectTableRealisationType1, err error) {
	m.ComObjectTableContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ComObjectTableRealisationType1"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ComObjectTableRealisationType1")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	numEntries, err := ReadSimpleField(ctx, "numEntries", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numEntries' field"))
	}
	m.NumEntries = numEntries

	ramFlagsTablePointer, err := ReadSimpleField(ctx, "ramFlagsTablePointer", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ramFlagsTablePointer' field"))
	}
	m.RamFlagsTablePointer = ramFlagsTablePointer

	comObjectDescriptors, err := ReadCountArrayField[GroupObjectDescriptorRealisationType1](ctx, "comObjectDescriptors", ReadComplex[GroupObjectDescriptorRealisationType1](GroupObjectDescriptorRealisationType1ParseWithBuffer, readBuffer), uint64(numEntries))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'comObjectDescriptors' field"))
	}
	m.ComObjectDescriptors = comObjectDescriptors

	if closeErr := readBuffer.CloseContext("ComObjectTableRealisationType1"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ComObjectTableRealisationType1")
	}

	return m, nil
}

func (m *_ComObjectTableRealisationType1) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ComObjectTableRealisationType1) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ComObjectTableRealisationType1"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ComObjectTableRealisationType1")
		}

		if err := WriteSimpleField[uint8](ctx, "numEntries", m.GetNumEntries(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'numEntries' field")
		}

		if err := WriteSimpleField[uint8](ctx, "ramFlagsTablePointer", m.GetRamFlagsTablePointer(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'ramFlagsTablePointer' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "comObjectDescriptors", m.GetComObjectDescriptors(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'comObjectDescriptors' field")
		}

		if popErr := writeBuffer.PopContext("ComObjectTableRealisationType1"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ComObjectTableRealisationType1")
		}
		return nil
	}
	return m.ComObjectTableContract.(*_ComObjectTable).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ComObjectTableRealisationType1) isComObjectTableRealisationType1() bool {
	return true
}

func (m *_ComObjectTableRealisationType1) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
