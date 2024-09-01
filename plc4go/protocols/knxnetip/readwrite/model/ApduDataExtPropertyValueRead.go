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

// ApduDataExtPropertyValueRead is the corresponding interface of ApduDataExtPropertyValueRead
type ApduDataExtPropertyValueRead interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ApduDataExt
	// GetObjectIndex returns ObjectIndex (property field)
	GetObjectIndex() uint8
	// GetPropertyId returns PropertyId (property field)
	GetPropertyId() uint8
	// GetCount returns Count (property field)
	GetCount() uint8
	// GetIndex returns Index (property field)
	GetIndex() uint16
}

// ApduDataExtPropertyValueReadExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtPropertyValueRead.
// This is useful for switch cases.
type ApduDataExtPropertyValueReadExactly interface {
	ApduDataExtPropertyValueRead
	isApduDataExtPropertyValueRead() bool
}

// _ApduDataExtPropertyValueRead is the data-structure of this message
type _ApduDataExtPropertyValueRead struct {
	*_ApduDataExt
	ObjectIndex uint8
	PropertyId  uint8
	Count       uint8
	Index       uint16
}

var _ ApduDataExtPropertyValueRead = (*_ApduDataExtPropertyValueRead)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtPropertyValueRead) GetExtApciType() uint8 {
	return 0x15
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtPropertyValueRead) InitializeParent(parent ApduDataExt) {}

func (m *_ApduDataExtPropertyValueRead) GetParent() ApduDataExt {
	return m._ApduDataExt
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ApduDataExtPropertyValueRead) GetObjectIndex() uint8 {
	return m.ObjectIndex
}

func (m *_ApduDataExtPropertyValueRead) GetPropertyId() uint8 {
	return m.PropertyId
}

func (m *_ApduDataExtPropertyValueRead) GetCount() uint8 {
	return m.Count
}

func (m *_ApduDataExtPropertyValueRead) GetIndex() uint16 {
	return m.Index
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewApduDataExtPropertyValueRead factory function for _ApduDataExtPropertyValueRead
func NewApduDataExtPropertyValueRead(objectIndex uint8, propertyId uint8, count uint8, index uint16, length uint8) *_ApduDataExtPropertyValueRead {
	_result := &_ApduDataExtPropertyValueRead{
		ObjectIndex:  objectIndex,
		PropertyId:   propertyId,
		Count:        count,
		Index:        index,
		_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtPropertyValueRead(structType any) ApduDataExtPropertyValueRead {
	if casted, ok := structType.(ApduDataExtPropertyValueRead); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtPropertyValueRead); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtPropertyValueRead) GetTypeName() string {
	return "ApduDataExtPropertyValueRead"
}

func (m *_ApduDataExtPropertyValueRead) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (objectIndex)
	lengthInBits += 8

	// Simple field (propertyId)
	lengthInBits += 8

	// Simple field (count)
	lengthInBits += 4

	// Simple field (index)
	lengthInBits += 12

	return lengthInBits
}

func (m *_ApduDataExtPropertyValueRead) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ApduDataExtPropertyValueReadParse(ctx context.Context, theBytes []byte, length uint8) (ApduDataExtPropertyValueRead, error) {
	return ApduDataExtPropertyValueReadParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), length)
}

func ApduDataExtPropertyValueReadParseWithBufferProducer(length uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (ApduDataExtPropertyValueRead, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ApduDataExtPropertyValueRead, error) {
		return ApduDataExtPropertyValueReadParseWithBuffer(ctx, readBuffer, length)
	}
}

func ApduDataExtPropertyValueReadParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, length uint8) (ApduDataExtPropertyValueRead, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtPropertyValueRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtPropertyValueRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectIndex, err := ReadSimpleField(ctx, "objectIndex", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectIndex' field"))
	}

	propertyId, err := ReadSimpleField(ctx, "propertyId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyId' field"))
	}

	count, err := ReadSimpleField(ctx, "count", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'count' field"))
	}

	index, err := ReadSimpleField(ctx, "index", ReadUnsignedShort(readBuffer, uint8(12)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'index' field"))
	}

	if closeErr := readBuffer.CloseContext("ApduDataExtPropertyValueRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtPropertyValueRead")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtPropertyValueRead{
		_ApduDataExt: &_ApduDataExt{
			Length: length,
		},
		ObjectIndex: objectIndex,
		PropertyId:  propertyId,
		Count:       count,
		Index:       index,
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtPropertyValueRead) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtPropertyValueRead) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtPropertyValueRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtPropertyValueRead")
		}

		if err := WriteSimpleField[uint8](ctx, "objectIndex", m.GetObjectIndex(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'objectIndex' field")
		}

		if err := WriteSimpleField[uint8](ctx, "propertyId", m.GetPropertyId(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'propertyId' field")
		}

		if err := WriteSimpleField[uint8](ctx, "count", m.GetCount(), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'count' field")
		}

		if err := WriteSimpleField[uint16](ctx, "index", m.GetIndex(), WriteUnsignedShort(writeBuffer, 12)); err != nil {
			return errors.Wrap(err, "Error serializing 'index' field")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtPropertyValueRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtPropertyValueRead")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtPropertyValueRead) isApduDataExtPropertyValueRead() bool {
	return true
}

func (m *_ApduDataExtPropertyValueRead) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
