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

// NullAddressItem is the corresponding interface of NullAddressItem
type NullAddressItem interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	TypeId
	// IsNullAddressItem is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNullAddressItem()
}

// _NullAddressItem is the data-structure of this message
type _NullAddressItem struct {
	TypeIdContract
	// Reserved Fields
	reservedField0 *uint16
}

var _ NullAddressItem = (*_NullAddressItem)(nil)
var _ TypeIdRequirements = (*_NullAddressItem)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NullAddressItem) GetId() uint16 {
	return 0x0000
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NullAddressItem) GetParent() TypeIdContract {
	return m.TypeIdContract
}

// NewNullAddressItem factory function for _NullAddressItem
func NewNullAddressItem() *_NullAddressItem {
	_result := &_NullAddressItem{
		TypeIdContract: NewTypeId(),
	}
	_result.TypeIdContract.(*_TypeId)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNullAddressItem(structType any) NullAddressItem {
	if casted, ok := structType.(NullAddressItem); ok {
		return casted
	}
	if casted, ok := structType.(*NullAddressItem); ok {
		return *casted
	}
	return nil
}

func (m *_NullAddressItem) GetTypeName() string {
	return "NullAddressItem"
}

func (m *_NullAddressItem) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.TypeIdContract.(*_TypeId).getLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 16

	return lengthInBits
}

func (m *_NullAddressItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NullAddressItem) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_TypeId) (__nullAddressItem NullAddressItem, err error) {
	m.TypeIdContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NullAddressItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NullAddressItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedShort(readBuffer, uint8(16)), uint16(0x0000))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	if closeErr := readBuffer.CloseContext("NullAddressItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NullAddressItem")
	}

	return m, nil
}

func (m *_NullAddressItem) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NullAddressItem) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NullAddressItem"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NullAddressItem")
		}

		if err := WriteReservedField[uint16](ctx, "reserved", uint16(0x0000), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if popErr := writeBuffer.PopContext("NullAddressItem"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NullAddressItem")
		}
		return nil
	}
	return m.TypeIdContract.(*_TypeId).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NullAddressItem) IsNullAddressItem() {}

func (m *_NullAddressItem) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
