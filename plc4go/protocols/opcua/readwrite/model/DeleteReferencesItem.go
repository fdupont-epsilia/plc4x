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

// DeleteReferencesItem is the corresponding interface of DeleteReferencesItem
type DeleteReferencesItem interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetSourceNodeId returns SourceNodeId (property field)
	GetSourceNodeId() NodeId
	// GetReferenceTypeId returns ReferenceTypeId (property field)
	GetReferenceTypeId() NodeId
	// GetIsForward returns IsForward (property field)
	GetIsForward() bool
	// GetTargetNodeId returns TargetNodeId (property field)
	GetTargetNodeId() ExpandedNodeId
	// GetDeleteBidirectional returns DeleteBidirectional (property field)
	GetDeleteBidirectional() bool
}

// DeleteReferencesItemExactly can be used when we want exactly this type and not a type which fulfills DeleteReferencesItem.
// This is useful for switch cases.
type DeleteReferencesItemExactly interface {
	DeleteReferencesItem
	isDeleteReferencesItem() bool
}

// _DeleteReferencesItem is the data-structure of this message
type _DeleteReferencesItem struct {
	*_ExtensionObjectDefinition
	SourceNodeId        NodeId
	ReferenceTypeId     NodeId
	IsForward           bool
	TargetNodeId        ExpandedNodeId
	DeleteBidirectional bool
	// Reserved Fields
	reservedField0 *uint8
	reservedField1 *uint8
}

var _ DeleteReferencesItem = (*_DeleteReferencesItem)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DeleteReferencesItem) GetIdentifier() string {
	return "387"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DeleteReferencesItem) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_DeleteReferencesItem) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeleteReferencesItem) GetSourceNodeId() NodeId {
	return m.SourceNodeId
}

func (m *_DeleteReferencesItem) GetReferenceTypeId() NodeId {
	return m.ReferenceTypeId
}

func (m *_DeleteReferencesItem) GetIsForward() bool {
	return m.IsForward
}

func (m *_DeleteReferencesItem) GetTargetNodeId() ExpandedNodeId {
	return m.TargetNodeId
}

func (m *_DeleteReferencesItem) GetDeleteBidirectional() bool {
	return m.DeleteBidirectional
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDeleteReferencesItem factory function for _DeleteReferencesItem
func NewDeleteReferencesItem(sourceNodeId NodeId, referenceTypeId NodeId, isForward bool, targetNodeId ExpandedNodeId, deleteBidirectional bool) *_DeleteReferencesItem {
	_result := &_DeleteReferencesItem{
		SourceNodeId:               sourceNodeId,
		ReferenceTypeId:            referenceTypeId,
		IsForward:                  isForward,
		TargetNodeId:               targetNodeId,
		DeleteBidirectional:        deleteBidirectional,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDeleteReferencesItem(structType any) DeleteReferencesItem {
	if casted, ok := structType.(DeleteReferencesItem); ok {
		return casted
	}
	if casted, ok := structType.(*DeleteReferencesItem); ok {
		return *casted
	}
	return nil
}

func (m *_DeleteReferencesItem) GetTypeName() string {
	return "DeleteReferencesItem"
}

func (m *_DeleteReferencesItem) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (sourceNodeId)
	lengthInBits += m.SourceNodeId.GetLengthInBits(ctx)

	// Simple field (referenceTypeId)
	lengthInBits += m.ReferenceTypeId.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (isForward)
	lengthInBits += 1

	// Simple field (targetNodeId)
	lengthInBits += m.TargetNodeId.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (deleteBidirectional)
	lengthInBits += 1

	return lengthInBits
}

func (m *_DeleteReferencesItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DeleteReferencesItemParse(ctx context.Context, theBytes []byte, identifier string) (DeleteReferencesItem, error) {
	return DeleteReferencesItemParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func DeleteReferencesItemParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (DeleteReferencesItem, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (DeleteReferencesItem, error) {
		return DeleteReferencesItemParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func DeleteReferencesItemParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (DeleteReferencesItem, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DeleteReferencesItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeleteReferencesItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	sourceNodeId, err := ReadSimpleField[NodeId](ctx, "sourceNodeId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sourceNodeId' field"))
	}

	referenceTypeId, err := ReadSimpleField[NodeId](ctx, "referenceTypeId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'referenceTypeId' field"))
	}

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	isForward, err := ReadSimpleField(ctx, "isForward", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isForward' field"))
	}

	targetNodeId, err := ReadSimpleField[ExpandedNodeId](ctx, "targetNodeId", ReadComplex[ExpandedNodeId](ExpandedNodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'targetNodeId' field"))
	}

	reservedField1, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	deleteBidirectional, err := ReadSimpleField(ctx, "deleteBidirectional", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deleteBidirectional' field"))
	}

	if closeErr := readBuffer.CloseContext("DeleteReferencesItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeleteReferencesItem")
	}

	// Create a partially initialized instance
	_child := &_DeleteReferencesItem{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		SourceNodeId:               sourceNodeId,
		ReferenceTypeId:            referenceTypeId,
		IsForward:                  isForward,
		TargetNodeId:               targetNodeId,
		DeleteBidirectional:        deleteBidirectional,
		reservedField0:             reservedField0,
		reservedField1:             reservedField1,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_DeleteReferencesItem) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DeleteReferencesItem) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DeleteReferencesItem"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DeleteReferencesItem")
		}

		if err := WriteSimpleField[NodeId](ctx, "sourceNodeId", m.GetSourceNodeId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'sourceNodeId' field")
		}

		if err := WriteSimpleField[NodeId](ctx, "referenceTypeId", m.GetReferenceTypeId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'referenceTypeId' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "isForward", m.GetIsForward(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'isForward' field")
		}

		if err := WriteSimpleField[ExpandedNodeId](ctx, "targetNodeId", m.GetTargetNodeId(), WriteComplex[ExpandedNodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'targetNodeId' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 2")
		}

		if err := WriteSimpleField[bool](ctx, "deleteBidirectional", m.GetDeleteBidirectional(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'deleteBidirectional' field")
		}

		if popErr := writeBuffer.PopContext("DeleteReferencesItem"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DeleteReferencesItem")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DeleteReferencesItem) isDeleteReferencesItem() bool {
	return true
}

func (m *_DeleteReferencesItem) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
