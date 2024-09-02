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

// S7PayloadUserDataItemCyclicServicesSubscribeRequest is the corresponding interface of S7PayloadUserDataItemCyclicServicesSubscribeRequest
type S7PayloadUserDataItemCyclicServicesSubscribeRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7PayloadUserDataItem
	// GetItemsCount returns ItemsCount (property field)
	GetItemsCount() uint16
	// GetTimeBase returns TimeBase (property field)
	GetTimeBase() TimeBase
	// GetTimeFactor returns TimeFactor (property field)
	GetTimeFactor() uint8
	// GetItem returns Item (property field)
	GetItem() []CycServiceItemType
	// IsS7PayloadUserDataItemCyclicServicesSubscribeRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7PayloadUserDataItemCyclicServicesSubscribeRequest()
}

// _S7PayloadUserDataItemCyclicServicesSubscribeRequest is the data-structure of this message
type _S7PayloadUserDataItemCyclicServicesSubscribeRequest struct {
	S7PayloadUserDataItemContract
	ItemsCount uint16
	TimeBase   TimeBase
	TimeFactor uint8
	Item       []CycServiceItemType
}

var _ S7PayloadUserDataItemCyclicServicesSubscribeRequest = (*_S7PayloadUserDataItemCyclicServicesSubscribeRequest)(nil)
var _ S7PayloadUserDataItemRequirements = (*_S7PayloadUserDataItemCyclicServicesSubscribeRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemCyclicServicesSubscribeRequest) GetCpuFunctionGroup() uint8 {
	return 0x02
}

func (m *_S7PayloadUserDataItemCyclicServicesSubscribeRequest) GetCpuFunctionType() uint8 {
	return 0x04
}

func (m *_S7PayloadUserDataItemCyclicServicesSubscribeRequest) GetCpuSubfunction() uint8 {
	return 0x01
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemCyclicServicesSubscribeRequest) GetParent() S7PayloadUserDataItemContract {
	return m.S7PayloadUserDataItemContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadUserDataItemCyclicServicesSubscribeRequest) GetItemsCount() uint16 {
	return m.ItemsCount
}

func (m *_S7PayloadUserDataItemCyclicServicesSubscribeRequest) GetTimeBase() TimeBase {
	return m.TimeBase
}

func (m *_S7PayloadUserDataItemCyclicServicesSubscribeRequest) GetTimeFactor() uint8 {
	return m.TimeFactor
}

func (m *_S7PayloadUserDataItemCyclicServicesSubscribeRequest) GetItem() []CycServiceItemType {
	return m.Item
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadUserDataItemCyclicServicesSubscribeRequest factory function for _S7PayloadUserDataItemCyclicServicesSubscribeRequest
func NewS7PayloadUserDataItemCyclicServicesSubscribeRequest(itemsCount uint16, timeBase TimeBase, timeFactor uint8, item []CycServiceItemType, returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) *_S7PayloadUserDataItemCyclicServicesSubscribeRequest {
	_result := &_S7PayloadUserDataItemCyclicServicesSubscribeRequest{
		S7PayloadUserDataItemContract: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
		ItemsCount:                    itemsCount,
		TimeBase:                      timeBase,
		TimeFactor:                    timeFactor,
		Item:                          item,
	}
	_result.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemCyclicServicesSubscribeRequest(structType any) S7PayloadUserDataItemCyclicServicesSubscribeRequest {
	if casted, ok := structType.(S7PayloadUserDataItemCyclicServicesSubscribeRequest); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCyclicServicesSubscribeRequest); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemCyclicServicesSubscribeRequest) GetTypeName() string {
	return "S7PayloadUserDataItemCyclicServicesSubscribeRequest"
}

func (m *_S7PayloadUserDataItemCyclicServicesSubscribeRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).getLengthInBits(ctx))

	// Simple field (itemsCount)
	lengthInBits += 16

	// Simple field (timeBase)
	lengthInBits += 8

	// Simple field (timeFactor)
	lengthInBits += 8

	// Array field
	if len(m.Item) > 0 {
		for _curItem, element := range m.Item {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Item), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_S7PayloadUserDataItemCyclicServicesSubscribeRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7PayloadUserDataItemCyclicServicesSubscribeRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7PayloadUserDataItem, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (__s7PayloadUserDataItemCyclicServicesSubscribeRequest S7PayloadUserDataItemCyclicServicesSubscribeRequest, err error) {
	m.S7PayloadUserDataItemContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCyclicServicesSubscribeRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemCyclicServicesSubscribeRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	itemsCount, err := ReadSimpleField(ctx, "itemsCount", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemsCount' field"))
	}
	m.ItemsCount = itemsCount

	timeBase, err := ReadEnumField[TimeBase](ctx, "timeBase", "TimeBase", ReadEnum(TimeBaseByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeBase' field"))
	}
	m.TimeBase = timeBase

	timeFactor, err := ReadSimpleField(ctx, "timeFactor", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeFactor' field"))
	}
	m.TimeFactor = timeFactor

	item, err := ReadCountArrayField[CycServiceItemType](ctx, "item", ReadComplex[CycServiceItemType](CycServiceItemTypeParseWithBuffer, readBuffer), uint64(itemsCount))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'item' field"))
	}
	m.Item = item

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCyclicServicesSubscribeRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemCyclicServicesSubscribeRequest")
	}

	return m, nil
}

func (m *_S7PayloadUserDataItemCyclicServicesSubscribeRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemCyclicServicesSubscribeRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCyclicServicesSubscribeRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemCyclicServicesSubscribeRequest")
		}

		if err := WriteSimpleField[uint16](ctx, "itemsCount", m.GetItemsCount(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'itemsCount' field")
		}

		if err := WriteSimpleEnumField[TimeBase](ctx, "timeBase", "TimeBase", m.GetTimeBase(), WriteEnum[TimeBase, uint8](TimeBase.GetValue, TimeBase.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'timeBase' field")
		}

		if err := WriteSimpleField[uint8](ctx, "timeFactor", m.GetTimeFactor(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeFactor' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "item", m.GetItem(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'item' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCyclicServicesSubscribeRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemCyclicServicesSubscribeRequest")
		}
		return nil
	}
	return m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemCyclicServicesSubscribeRequest) IsS7PayloadUserDataItemCyclicServicesSubscribeRequest() {
}

func (m *_S7PayloadUserDataItemCyclicServicesSubscribeRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
