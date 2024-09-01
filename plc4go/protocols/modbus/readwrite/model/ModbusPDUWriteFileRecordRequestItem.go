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

// ModbusPDUWriteFileRecordRequestItem is the corresponding interface of ModbusPDUWriteFileRecordRequestItem
type ModbusPDUWriteFileRecordRequestItem interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetReferenceType returns ReferenceType (property field)
	GetReferenceType() uint8
	// GetFileNumber returns FileNumber (property field)
	GetFileNumber() uint16
	// GetRecordNumber returns RecordNumber (property field)
	GetRecordNumber() uint16
	// GetRecordData returns RecordData (property field)
	GetRecordData() []byte
}

// ModbusPDUWriteFileRecordRequestItemExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUWriteFileRecordRequestItem.
// This is useful for switch cases.
type ModbusPDUWriteFileRecordRequestItemExactly interface {
	ModbusPDUWriteFileRecordRequestItem
	isModbusPDUWriteFileRecordRequestItem() bool
}

// _ModbusPDUWriteFileRecordRequestItem is the data-structure of this message
type _ModbusPDUWriteFileRecordRequestItem struct {
	ReferenceType uint8
	FileNumber    uint16
	RecordNumber  uint16
	RecordData    []byte
}

var _ ModbusPDUWriteFileRecordRequestItem = (*_ModbusPDUWriteFileRecordRequestItem)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUWriteFileRecordRequestItem) GetReferenceType() uint8 {
	return m.ReferenceType
}

func (m *_ModbusPDUWriteFileRecordRequestItem) GetFileNumber() uint16 {
	return m.FileNumber
}

func (m *_ModbusPDUWriteFileRecordRequestItem) GetRecordNumber() uint16 {
	return m.RecordNumber
}

func (m *_ModbusPDUWriteFileRecordRequestItem) GetRecordData() []byte {
	return m.RecordData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUWriteFileRecordRequestItem factory function for _ModbusPDUWriteFileRecordRequestItem
func NewModbusPDUWriteFileRecordRequestItem(referenceType uint8, fileNumber uint16, recordNumber uint16, recordData []byte) *_ModbusPDUWriteFileRecordRequestItem {
	return &_ModbusPDUWriteFileRecordRequestItem{ReferenceType: referenceType, FileNumber: fileNumber, RecordNumber: recordNumber, RecordData: recordData}
}

// Deprecated: use the interface for direct cast
func CastModbusPDUWriteFileRecordRequestItem(structType any) ModbusPDUWriteFileRecordRequestItem {
	if casted, ok := structType.(ModbusPDUWriteFileRecordRequestItem); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUWriteFileRecordRequestItem); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUWriteFileRecordRequestItem) GetTypeName() string {
	return "ModbusPDUWriteFileRecordRequestItem"
}

func (m *_ModbusPDUWriteFileRecordRequestItem) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (referenceType)
	lengthInBits += 8

	// Simple field (fileNumber)
	lengthInBits += 16

	// Simple field (recordNumber)
	lengthInBits += 16

	// Implicit Field (recordLength)
	lengthInBits += 16

	// Array field
	if len(m.RecordData) > 0 {
		lengthInBits += 8 * uint16(len(m.RecordData))
	}

	return lengthInBits
}

func (m *_ModbusPDUWriteFileRecordRequestItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusPDUWriteFileRecordRequestItemParse(ctx context.Context, theBytes []byte) (ModbusPDUWriteFileRecordRequestItem, error) {
	return ModbusPDUWriteFileRecordRequestItemParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ModbusPDUWriteFileRecordRequestItemParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (ModbusPDUWriteFileRecordRequestItem, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ModbusPDUWriteFileRecordRequestItem, error) {
		return ModbusPDUWriteFileRecordRequestItemParseWithBuffer(ctx, readBuffer)
	}
}

func ModbusPDUWriteFileRecordRequestItemParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ModbusPDUWriteFileRecordRequestItem, error) {
	v, err := (&_ModbusPDUWriteFileRecordRequestItem{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_ModbusPDUWriteFileRecordRequestItem) parse(ctx context.Context, readBuffer utils.ReadBuffer) (_modbusPDUWriteFileRecordRequestItem ModbusPDUWriteFileRecordRequestItem, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUWriteFileRecordRequestItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUWriteFileRecordRequestItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	referenceType, err := ReadSimpleField(ctx, "referenceType", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'referenceType' field"))
	}
	m.ReferenceType = referenceType

	fileNumber, err := ReadSimpleField(ctx, "fileNumber", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fileNumber' field"))
	}
	m.FileNumber = fileNumber

	recordNumber, err := ReadSimpleField(ctx, "recordNumber", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'recordNumber' field"))
	}
	m.RecordNumber = recordNumber

	recordLength, err := ReadImplicitField[uint16](ctx, "recordLength", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'recordLength' field"))
	}
	_ = recordLength

	recordData, err := readBuffer.ReadByteArray("recordData", int(int32(recordLength)*int32(int32(2))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'recordData' field"))
	}
	m.RecordData = recordData

	if closeErr := readBuffer.CloseContext("ModbusPDUWriteFileRecordRequestItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUWriteFileRecordRequestItem")
	}

	return m, nil
}

func (m *_ModbusPDUWriteFileRecordRequestItem) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUWriteFileRecordRequestItem) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ModbusPDUWriteFileRecordRequestItem"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ModbusPDUWriteFileRecordRequestItem")
	}

	if err := WriteSimpleField[uint8](ctx, "referenceType", m.GetReferenceType(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'referenceType' field")
	}

	if err := WriteSimpleField[uint16](ctx, "fileNumber", m.GetFileNumber(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'fileNumber' field")
	}

	if err := WriteSimpleField[uint16](ctx, "recordNumber", m.GetRecordNumber(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'recordNumber' field")
	}
	recordLength := uint16(uint16(uint16(len(m.GetRecordData()))) / uint16(uint16(2)))
	if err := WriteImplicitField(ctx, "recordLength", recordLength, WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'recordLength' field")
	}

	if err := WriteByteArrayField(ctx, "recordData", m.GetRecordData(), WriteByteArray(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'recordData' field")
	}

	if popErr := writeBuffer.PopContext("ModbusPDUWriteFileRecordRequestItem"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ModbusPDUWriteFileRecordRequestItem")
	}
	return nil
}

func (m *_ModbusPDUWriteFileRecordRequestItem) isModbusPDUWriteFileRecordRequestItem() bool {
	return true
}

func (m *_ModbusPDUWriteFileRecordRequestItem) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
