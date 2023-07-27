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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// GuidNodeId is the corresponding interface of GuidNodeId
type GuidNodeId interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetNamespaceIndex returns NamespaceIndex (property field)
	GetNamespaceIndex() uint16
	// GetIdentifier returns Identifier (property field)
	GetIdentifier() GuidValue
}

// GuidNodeIdExactly can be used when we want exactly this type and not a type which fulfills GuidNodeId.
// This is useful for switch cases.
type GuidNodeIdExactly interface {
	GuidNodeId
	isGuidNodeId() bool
}

// _GuidNodeId is the data-structure of this message
type _GuidNodeId struct {
	NamespaceIndex uint16
	Identifier     GuidValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_GuidNodeId) GetNamespaceIndex() uint16 {
	return m.NamespaceIndex
}

func (m *_GuidNodeId) GetIdentifier() GuidValue {
	return m.Identifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewGuidNodeId factory function for _GuidNodeId
func NewGuidNodeId(namespaceIndex uint16, identifier GuidValue) *_GuidNodeId {
	return &_GuidNodeId{NamespaceIndex: namespaceIndex, Identifier: identifier}
}

// Deprecated: use the interface for direct cast
func CastGuidNodeId(structType any) GuidNodeId {
	if casted, ok := structType.(GuidNodeId); ok {
		return casted
	}
	if casted, ok := structType.(*GuidNodeId); ok {
		return *casted
	}
	return nil
}

func (m *_GuidNodeId) GetTypeName() string {
	return "GuidNodeId"
}

func (m *_GuidNodeId) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (namespaceIndex)
	lengthInBits += 16

	// Simple field (identifier)
	lengthInBits += m.Identifier.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_GuidNodeId) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func GuidNodeIdParse(ctx context.Context, theBytes []byte) (GuidNodeId, error) {
	return GuidNodeIdParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func GuidNodeIdParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (GuidNodeId, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("GuidNodeId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for GuidNodeId")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (namespaceIndex)
	_namespaceIndex, _namespaceIndexErr := readBuffer.ReadUint16("namespaceIndex", 16)
	if _namespaceIndexErr != nil {
		return nil, errors.Wrap(_namespaceIndexErr, "Error parsing 'namespaceIndex' field of GuidNodeId")
	}
	namespaceIndex := _namespaceIndex

	// Simple Field (identifier)
	if pullErr := readBuffer.PullContext("identifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for identifier")
	}
	_identifier, _identifierErr := GuidValueParseWithBuffer(ctx, readBuffer)
	if _identifierErr != nil {
		return nil, errors.Wrap(_identifierErr, "Error parsing 'identifier' field of GuidNodeId")
	}
	identifier := _identifier.(GuidValue)
	if closeErr := readBuffer.CloseContext("identifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for identifier")
	}

	if closeErr := readBuffer.CloseContext("GuidNodeId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for GuidNodeId")
	}

	// Create the instance
	return &_GuidNodeId{
		NamespaceIndex: namespaceIndex,
		Identifier:     identifier,
	}, nil
}

func (m *_GuidNodeId) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_GuidNodeId) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("GuidNodeId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for GuidNodeId")
	}

	// Simple Field (namespaceIndex)
	namespaceIndex := uint16(m.GetNamespaceIndex())
	_namespaceIndexErr := writeBuffer.WriteUint16("namespaceIndex", 16, (namespaceIndex))
	if _namespaceIndexErr != nil {
		return errors.Wrap(_namespaceIndexErr, "Error serializing 'namespaceIndex' field")
	}

	// Simple Field (identifier)
	if pushErr := writeBuffer.PushContext("identifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for identifier")
	}
	_identifierErr := writeBuffer.WriteSerializable(ctx, m.GetIdentifier())
	if popErr := writeBuffer.PopContext("identifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for identifier")
	}
	if _identifierErr != nil {
		return errors.Wrap(_identifierErr, "Error serializing 'identifier' field")
	}

	if popErr := writeBuffer.PopContext("GuidNodeId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for GuidNodeId")
	}
	return nil
}

func (m *_GuidNodeId) isGuidNodeId() bool {
	return true
}

func (m *_GuidNodeId) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}