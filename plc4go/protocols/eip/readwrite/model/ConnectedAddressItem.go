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

// ConnectedAddressItem is the corresponding interface of ConnectedAddressItem
type ConnectedAddressItem interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	TypeId
	// GetConnectionId returns ConnectionId (property field)
	GetConnectionId() uint32
}

// ConnectedAddressItemExactly can be used when we want exactly this type and not a type which fulfills ConnectedAddressItem.
// This is useful for switch cases.
type ConnectedAddressItemExactly interface {
	ConnectedAddressItem
	isConnectedAddressItem() bool
}

// _ConnectedAddressItem is the data-structure of this message
type _ConnectedAddressItem struct {
	*_TypeId
	ConnectionId uint32
	// Reserved Fields
	reservedField0 *uint16
}

var _ ConnectedAddressItem = (*_ConnectedAddressItem)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ConnectedAddressItem) GetId() uint16 {
	return 0x00A1
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ConnectedAddressItem) InitializeParent(parent TypeId) {}

func (m *_ConnectedAddressItem) GetParent() TypeId {
	return m._TypeId
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ConnectedAddressItem) GetConnectionId() uint32 {
	return m.ConnectionId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewConnectedAddressItem factory function for _ConnectedAddressItem
func NewConnectedAddressItem(connectionId uint32) *_ConnectedAddressItem {
	_result := &_ConnectedAddressItem{
		ConnectionId: connectionId,
		_TypeId:      NewTypeId(),
	}
	_result._TypeId._TypeIdChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastConnectedAddressItem(structType any) ConnectedAddressItem {
	if casted, ok := structType.(ConnectedAddressItem); ok {
		return casted
	}
	if casted, ok := structType.(*ConnectedAddressItem); ok {
		return *casted
	}
	return nil
}

func (m *_ConnectedAddressItem) GetTypeName() string {
	return "ConnectedAddressItem"
}

func (m *_ConnectedAddressItem) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 16

	// Simple field (connectionId)
	lengthInBits += 32

	return lengthInBits
}

func (m *_ConnectedAddressItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ConnectedAddressItemParse(ctx context.Context, theBytes []byte) (ConnectedAddressItem, error) {
	return ConnectedAddressItemParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ConnectedAddressItemParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (ConnectedAddressItem, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ConnectedAddressItem, error) {
		return ConnectedAddressItemParseWithBuffer(ctx, readBuffer)
	}
}

func ConnectedAddressItemParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ConnectedAddressItem, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ConnectedAddressItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ConnectedAddressItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedShort(readBuffer, uint8(16)), uint16(0x0004))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	connectionId, err := ReadSimpleField(ctx, "connectionId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'connectionId' field"))
	}

	if closeErr := readBuffer.CloseContext("ConnectedAddressItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ConnectedAddressItem")
	}

	// Create a partially initialized instance
	_child := &_ConnectedAddressItem{
		_TypeId:        &_TypeId{},
		ConnectionId:   connectionId,
		reservedField0: reservedField0,
	}
	_child._TypeId._TypeIdChildRequirements = _child
	return _child, nil
}

func (m *_ConnectedAddressItem) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ConnectedAddressItem) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ConnectedAddressItem"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ConnectedAddressItem")
		}

		if err := WriteReservedField[uint16](ctx, "reserved", uint16(0x0004), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[uint32](ctx, "connectionId", m.GetConnectionId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'connectionId' field")
		}

		if popErr := writeBuffer.PopContext("ConnectedAddressItem"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ConnectedAddressItem")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ConnectedAddressItem) isConnectedAddressItem() bool {
	return true
}

func (m *_ConnectedAddressItem) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
