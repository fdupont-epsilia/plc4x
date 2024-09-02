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

// NodeReference is the corresponding interface of NodeReference
type NodeReference interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNodeId returns NodeId (property field)
	GetNodeId() NodeId
	// GetReferenceTypeId returns ReferenceTypeId (property field)
	GetReferenceTypeId() NodeId
	// GetIsForward returns IsForward (property field)
	GetIsForward() bool
	// GetNoOfReferencedNodeIds returns NoOfReferencedNodeIds (property field)
	GetNoOfReferencedNodeIds() int32
	// GetReferencedNodeIds returns ReferencedNodeIds (property field)
	GetReferencedNodeIds() []NodeId
	// IsNodeReference is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNodeReference()
}

// _NodeReference is the data-structure of this message
type _NodeReference struct {
	ExtensionObjectDefinitionContract
	NodeId                NodeId
	ReferenceTypeId       NodeId
	IsForward             bool
	NoOfReferencedNodeIds int32
	ReferencedNodeIds     []NodeId
	// Reserved Fields
	reservedField0 *uint8
}

var _ NodeReference = (*_NodeReference)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_NodeReference)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NodeReference) GetIdentifier() string {
	return "582"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NodeReference) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NodeReference) GetNodeId() NodeId {
	return m.NodeId
}

func (m *_NodeReference) GetReferenceTypeId() NodeId {
	return m.ReferenceTypeId
}

func (m *_NodeReference) GetIsForward() bool {
	return m.IsForward
}

func (m *_NodeReference) GetNoOfReferencedNodeIds() int32 {
	return m.NoOfReferencedNodeIds
}

func (m *_NodeReference) GetReferencedNodeIds() []NodeId {
	return m.ReferencedNodeIds
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNodeReference factory function for _NodeReference
func NewNodeReference(nodeId NodeId, referenceTypeId NodeId, isForward bool, noOfReferencedNodeIds int32, referencedNodeIds []NodeId) *_NodeReference {
	_result := &_NodeReference{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		NodeId:                            nodeId,
		ReferenceTypeId:                   referenceTypeId,
		IsForward:                         isForward,
		NoOfReferencedNodeIds:             noOfReferencedNodeIds,
		ReferencedNodeIds:                 referencedNodeIds,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNodeReference(structType any) NodeReference {
	if casted, ok := structType.(NodeReference); ok {
		return casted
	}
	if casted, ok := structType.(*NodeReference); ok {
		return *casted
	}
	return nil
}

func (m *_NodeReference) GetTypeName() string {
	return "NodeReference"
}

func (m *_NodeReference) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (nodeId)
	lengthInBits += m.NodeId.GetLengthInBits(ctx)

	// Simple field (referenceTypeId)
	lengthInBits += m.ReferenceTypeId.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (isForward)
	lengthInBits += 1

	// Simple field (noOfReferencedNodeIds)
	lengthInBits += 32

	// Array field
	if len(m.ReferencedNodeIds) > 0 {
		for _curItem, element := range m.ReferencedNodeIds {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ReferencedNodeIds), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_NodeReference) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NodeReference) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__nodeReference NodeReference, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NodeReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NodeReference")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	nodeId, err := ReadSimpleField[NodeId](ctx, "nodeId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodeId' field"))
	}
	m.NodeId = nodeId

	referenceTypeId, err := ReadSimpleField[NodeId](ctx, "referenceTypeId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'referenceTypeId' field"))
	}
	m.ReferenceTypeId = referenceTypeId

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	isForward, err := ReadSimpleField(ctx, "isForward", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isForward' field"))
	}
	m.IsForward = isForward

	noOfReferencedNodeIds, err := ReadSimpleField(ctx, "noOfReferencedNodeIds", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfReferencedNodeIds' field"))
	}
	m.NoOfReferencedNodeIds = noOfReferencedNodeIds

	referencedNodeIds, err := ReadCountArrayField[NodeId](ctx, "referencedNodeIds", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer), uint64(noOfReferencedNodeIds))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'referencedNodeIds' field"))
	}
	m.ReferencedNodeIds = referencedNodeIds

	if closeErr := readBuffer.CloseContext("NodeReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NodeReference")
	}

	return m, nil
}

func (m *_NodeReference) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NodeReference) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NodeReference"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NodeReference")
		}

		if err := WriteSimpleField[NodeId](ctx, "nodeId", m.GetNodeId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'nodeId' field")
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

		if err := WriteSimpleField[int32](ctx, "noOfReferencedNodeIds", m.GetNoOfReferencedNodeIds(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfReferencedNodeIds' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "referencedNodeIds", m.GetReferencedNodeIds(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'referencedNodeIds' field")
		}

		if popErr := writeBuffer.PopContext("NodeReference"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NodeReference")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NodeReference) IsNodeReference() {}

func (m *_NodeReference) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
