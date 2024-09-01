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

// ByteStringNodeId is the corresponding interface of ByteStringNodeId
type ByteStringNodeId interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetNamespaceIndex returns NamespaceIndex (property field)
	GetNamespaceIndex() uint16
	// GetIdentifier returns Identifier (property field)
	GetIdentifier() PascalByteString
}

// ByteStringNodeIdExactly can be used when we want exactly this type and not a type which fulfills ByteStringNodeId.
// This is useful for switch cases.
type ByteStringNodeIdExactly interface {
	ByteStringNodeId
	isByteStringNodeId() bool
}

// _ByteStringNodeId is the data-structure of this message
type _ByteStringNodeId struct {
	NamespaceIndex uint16
	Identifier     PascalByteString
}

var _ ByteStringNodeId = (*_ByteStringNodeId)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ByteStringNodeId) GetNamespaceIndex() uint16 {
	return m.NamespaceIndex
}

func (m *_ByteStringNodeId) GetIdentifier() PascalByteString {
	return m.Identifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewByteStringNodeId factory function for _ByteStringNodeId
func NewByteStringNodeId(namespaceIndex uint16, identifier PascalByteString) *_ByteStringNodeId {
	return &_ByteStringNodeId{NamespaceIndex: namespaceIndex, Identifier: identifier}
}

// Deprecated: use the interface for direct cast
func CastByteStringNodeId(structType any) ByteStringNodeId {
	if casted, ok := structType.(ByteStringNodeId); ok {
		return casted
	}
	if casted, ok := structType.(*ByteStringNodeId); ok {
		return *casted
	}
	return nil
}

func (m *_ByteStringNodeId) GetTypeName() string {
	return "ByteStringNodeId"
}

func (m *_ByteStringNodeId) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (namespaceIndex)
	lengthInBits += 16

	// Simple field (identifier)
	lengthInBits += m.Identifier.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ByteStringNodeId) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ByteStringNodeIdParse(ctx context.Context, theBytes []byte) (ByteStringNodeId, error) {
	return ByteStringNodeIdParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ByteStringNodeIdParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (ByteStringNodeId, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ByteStringNodeId, error) {
		return ByteStringNodeIdParseWithBuffer(ctx, readBuffer)
	}
}

func ByteStringNodeIdParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ByteStringNodeId, error) {
	v, err := (&_ByteStringNodeId{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_ByteStringNodeId) parse(ctx context.Context, readBuffer utils.ReadBuffer) (_byteStringNodeId ByteStringNodeId, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ByteStringNodeId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ByteStringNodeId")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	namespaceIndex, err := ReadSimpleField(ctx, "namespaceIndex", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'namespaceIndex' field"))
	}
	m.NamespaceIndex = namespaceIndex

	identifier, err := ReadSimpleField[PascalByteString](ctx, "identifier", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'identifier' field"))
	}
	m.Identifier = identifier

	if closeErr := readBuffer.CloseContext("ByteStringNodeId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ByteStringNodeId")
	}

	return m, nil
}

func (m *_ByteStringNodeId) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ByteStringNodeId) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ByteStringNodeId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ByteStringNodeId")
	}

	if err := WriteSimpleField[uint16](ctx, "namespaceIndex", m.GetNamespaceIndex(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'namespaceIndex' field")
	}

	if err := WriteSimpleField[PascalByteString](ctx, "identifier", m.GetIdentifier(), WriteComplex[PascalByteString](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'identifier' field")
	}

	if popErr := writeBuffer.PopContext("ByteStringNodeId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ByteStringNodeId")
	}
	return nil
}

func (m *_ByteStringNodeId) isByteStringNodeId() bool {
	return true
}

func (m *_ByteStringNodeId) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
