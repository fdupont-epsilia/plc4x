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
	utils.Copyable
	TypeId
	// GetConnectionId returns ConnectionId (property field)
	GetConnectionId() uint32
	// IsConnectedAddressItem is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsConnectedAddressItem()
	// CreateBuilder creates a ConnectedAddressItemBuilder
	CreateConnectedAddressItemBuilder() ConnectedAddressItemBuilder
}

// _ConnectedAddressItem is the data-structure of this message
type _ConnectedAddressItem struct {
	TypeIdContract
	ConnectionId uint32
	// Reserved Fields
	reservedField0 *uint16
}

var _ ConnectedAddressItem = (*_ConnectedAddressItem)(nil)
var _ TypeIdRequirements = (*_ConnectedAddressItem)(nil)

// NewConnectedAddressItem factory function for _ConnectedAddressItem
func NewConnectedAddressItem(connectionId uint32) *_ConnectedAddressItem {
	_result := &_ConnectedAddressItem{
		TypeIdContract: NewTypeId(),
		ConnectionId:   connectionId,
	}
	_result.TypeIdContract.(*_TypeId)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ConnectedAddressItemBuilder is a builder for ConnectedAddressItem
type ConnectedAddressItemBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(connectionId uint32) ConnectedAddressItemBuilder
	// WithConnectionId adds ConnectionId (property field)
	WithConnectionId(uint32) ConnectedAddressItemBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() TypeIdBuilder
	// Build builds the ConnectedAddressItem or returns an error if something is wrong
	Build() (ConnectedAddressItem, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ConnectedAddressItem
}

// NewConnectedAddressItemBuilder() creates a ConnectedAddressItemBuilder
func NewConnectedAddressItemBuilder() ConnectedAddressItemBuilder {
	return &_ConnectedAddressItemBuilder{_ConnectedAddressItem: new(_ConnectedAddressItem)}
}

type _ConnectedAddressItemBuilder struct {
	*_ConnectedAddressItem

	parentBuilder *_TypeIdBuilder

	err *utils.MultiError
}

var _ (ConnectedAddressItemBuilder) = (*_ConnectedAddressItemBuilder)(nil)

func (b *_ConnectedAddressItemBuilder) setParent(contract TypeIdContract) {
	b.TypeIdContract = contract
	contract.(*_TypeId)._SubType = b._ConnectedAddressItem
}

func (b *_ConnectedAddressItemBuilder) WithMandatoryFields(connectionId uint32) ConnectedAddressItemBuilder {
	return b.WithConnectionId(connectionId)
}

func (b *_ConnectedAddressItemBuilder) WithConnectionId(connectionId uint32) ConnectedAddressItemBuilder {
	b.ConnectionId = connectionId
	return b
}

func (b *_ConnectedAddressItemBuilder) Build() (ConnectedAddressItem, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ConnectedAddressItem.deepCopy(), nil
}

func (b *_ConnectedAddressItemBuilder) MustBuild() ConnectedAddressItem {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ConnectedAddressItemBuilder) Done() TypeIdBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewTypeIdBuilder().(*_TypeIdBuilder)
	}
	return b.parentBuilder
}

func (b *_ConnectedAddressItemBuilder) buildForTypeId() (TypeId, error) {
	return b.Build()
}

func (b *_ConnectedAddressItemBuilder) DeepCopy() any {
	_copy := b.CreateConnectedAddressItemBuilder().(*_ConnectedAddressItemBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateConnectedAddressItemBuilder creates a ConnectedAddressItemBuilder
func (b *_ConnectedAddressItem) CreateConnectedAddressItemBuilder() ConnectedAddressItemBuilder {
	if b == nil {
		return NewConnectedAddressItemBuilder()
	}
	return &_ConnectedAddressItemBuilder{_ConnectedAddressItem: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

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

func (m *_ConnectedAddressItem) GetParent() TypeIdContract {
	return m.TypeIdContract
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
	lengthInBits := uint16(m.TypeIdContract.(*_TypeId).getLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 16

	// Simple field (connectionId)
	lengthInBits += 32

	return lengthInBits
}

func (m *_ConnectedAddressItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ConnectedAddressItem) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_TypeId) (__connectedAddressItem ConnectedAddressItem, err error) {
	m.TypeIdContract = parent
	parent._SubType = m
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
	m.reservedField0 = reservedField0

	connectionId, err := ReadSimpleField(ctx, "connectionId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'connectionId' field"))
	}
	m.ConnectionId = connectionId

	if closeErr := readBuffer.CloseContext("ConnectedAddressItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ConnectedAddressItem")
	}

	return m, nil
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
	return m.TypeIdContract.(*_TypeId).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ConnectedAddressItem) IsConnectedAddressItem() {}

func (m *_ConnectedAddressItem) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ConnectedAddressItem) deepCopy() *_ConnectedAddressItem {
	if m == nil {
		return nil
	}
	_ConnectedAddressItemCopy := &_ConnectedAddressItem{
		m.TypeIdContract.(*_TypeId).deepCopy(),
		m.ConnectionId,
		m.reservedField0,
	}
	_ConnectedAddressItemCopy.TypeIdContract.(*_TypeId)._SubType = m
	return _ConnectedAddressItemCopy
}

func (m *_ConnectedAddressItem) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
