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

// ModbusPDUReadFileRecordResponse is the corresponding interface of ModbusPDUReadFileRecordResponse
type ModbusPDUReadFileRecordResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetItems returns Items (property field)
	GetItems() []ModbusPDUReadFileRecordResponseItem
}

// ModbusPDUReadFileRecordResponseExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUReadFileRecordResponse.
// This is useful for switch cases.
type ModbusPDUReadFileRecordResponseExactly interface {
	ModbusPDUReadFileRecordResponse
	isModbusPDUReadFileRecordResponse() bool
}

// _ModbusPDUReadFileRecordResponse is the data-structure of this message
type _ModbusPDUReadFileRecordResponse struct {
	ModbusPDUContract
	Items []ModbusPDUReadFileRecordResponseItem
}

var _ ModbusPDUReadFileRecordResponse = (*_ModbusPDUReadFileRecordResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReadFileRecordResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUReadFileRecordResponse) GetFunctionFlag() uint8 {
	return 0x14
}

func (m *_ModbusPDUReadFileRecordResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReadFileRecordResponse) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReadFileRecordResponse) GetItems() []ModbusPDUReadFileRecordResponseItem {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUReadFileRecordResponse factory function for _ModbusPDUReadFileRecordResponse
func NewModbusPDUReadFileRecordResponse(items []ModbusPDUReadFileRecordResponseItem) *_ModbusPDUReadFileRecordResponse {
	_result := &_ModbusPDUReadFileRecordResponse{
		ModbusPDUContract: NewModbusPDU(),
		Items:             items,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUReadFileRecordResponse(structType any) ModbusPDUReadFileRecordResponse {
	if casted, ok := structType.(ModbusPDUReadFileRecordResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadFileRecordResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadFileRecordResponse) GetTypeName() string {
	return "ModbusPDUReadFileRecordResponse"
}

func (m *_ModbusPDUReadFileRecordResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.Items) > 0 {
		for _, element := range m.Items {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_ModbusPDUReadFileRecordResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUReadFileRecordResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (_modbusPDUReadFileRecordResponse ModbusPDUReadFileRecordResponse, err error) {
	m.ModbusPDUContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadFileRecordResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadFileRecordResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	byteCount, err := ReadImplicitField[uint8](ctx, "byteCount", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'byteCount' field"))
	}
	_ = byteCount

	items, err := ReadLengthArrayField[ModbusPDUReadFileRecordResponseItem](ctx, "items", ReadComplex[ModbusPDUReadFileRecordResponseItem](ModbusPDUReadFileRecordResponseItemParseWithBuffer, readBuffer), int(byteCount))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'items' field"))
	}
	m.Items = items

	if closeErr := readBuffer.CloseContext("ModbusPDUReadFileRecordResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadFileRecordResponse")
	}

	return m, nil
}

func (m *_ModbusPDUReadFileRecordResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReadFileRecordResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	itemsArraySizeInBytes := func(items []ModbusPDUReadFileRecordResponseItem) uint32 {
		var sizeInBytes uint32 = 0
		for _, v := range items {
			sizeInBytes += uint32(v.GetLengthInBytes(ctx))
		}
		return sizeInBytes
	}
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadFileRecordResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadFileRecordResponse")
		}
		byteCount := uint8(uint8(itemsArraySizeInBytes(m.GetItems())))
		if err := WriteImplicitField(ctx, "byteCount", byteCount, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'byteCount' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "items", m.GetItems(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'items' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadFileRecordResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReadFileRecordResponse")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUReadFileRecordResponse) isModbusPDUReadFileRecordResponse() bool {
	return true
}

func (m *_ModbusPDUReadFileRecordResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
