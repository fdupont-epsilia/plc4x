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

// FourByteNodeId is the corresponding interface of FourByteNodeId
type FourByteNodeId interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetNamespaceIndex returns NamespaceIndex (property field)
	GetNamespaceIndex() uint8
	// GetIdentifier returns Identifier (property field)
	GetIdentifier() uint16
	// IsFourByteNodeId is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsFourByteNodeId()
}

// _FourByteNodeId is the data-structure of this message
type _FourByteNodeId struct {
	NamespaceIndex uint8
	Identifier     uint16
}

var _ FourByteNodeId = (*_FourByteNodeId)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_FourByteNodeId) GetNamespaceIndex() uint8 {
	return m.NamespaceIndex
}

func (m *_FourByteNodeId) GetIdentifier() uint16 {
	return m.Identifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewFourByteNodeId factory function for _FourByteNodeId
func NewFourByteNodeId(namespaceIndex uint8, identifier uint16) *_FourByteNodeId {
	return &_FourByteNodeId{NamespaceIndex: namespaceIndex, Identifier: identifier}
}

// Deprecated: use the interface for direct cast
func CastFourByteNodeId(structType any) FourByteNodeId {
	if casted, ok := structType.(FourByteNodeId); ok {
		return casted
	}
	if casted, ok := structType.(*FourByteNodeId); ok {
		return *casted
	}
	return nil
}

func (m *_FourByteNodeId) GetTypeName() string {
	return "FourByteNodeId"
}

func (m *_FourByteNodeId) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (namespaceIndex)
	lengthInBits += 8

	// Simple field (identifier)
	lengthInBits += 16

	return lengthInBits
}

func (m *_FourByteNodeId) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func FourByteNodeIdParse(ctx context.Context, theBytes []byte) (FourByteNodeId, error) {
	return FourByteNodeIdParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func FourByteNodeIdParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (FourByteNodeId, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (FourByteNodeId, error) {
		return FourByteNodeIdParseWithBuffer(ctx, readBuffer)
	}
}

func FourByteNodeIdParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (FourByteNodeId, error) {
	v, err := (&_FourByteNodeId{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_FourByteNodeId) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__fourByteNodeId FourByteNodeId, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FourByteNodeId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FourByteNodeId")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	namespaceIndex, err := ReadSimpleField(ctx, "namespaceIndex", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'namespaceIndex' field"))
	}
	m.NamespaceIndex = namespaceIndex

	identifier, err := ReadSimpleField(ctx, "identifier", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'identifier' field"))
	}
	m.Identifier = identifier

	if closeErr := readBuffer.CloseContext("FourByteNodeId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FourByteNodeId")
	}

	return m, nil
}

func (m *_FourByteNodeId) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FourByteNodeId) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("FourByteNodeId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for FourByteNodeId")
	}

	if err := WriteSimpleField[uint8](ctx, "namespaceIndex", m.GetNamespaceIndex(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'namespaceIndex' field")
	}

	if err := WriteSimpleField[uint16](ctx, "identifier", m.GetIdentifier(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'identifier' field")
	}

	if popErr := writeBuffer.PopContext("FourByteNodeId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for FourByteNodeId")
	}
	return nil
}

func (m *_FourByteNodeId) IsFourByteNodeId() {}

func (m *_FourByteNodeId) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
