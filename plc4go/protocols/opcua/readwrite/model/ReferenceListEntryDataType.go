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

// ReferenceListEntryDataType is the corresponding interface of ReferenceListEntryDataType
type ReferenceListEntryDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetReferenceType returns ReferenceType (property field)
	GetReferenceType() NodeId
	// GetIsForward returns IsForward (property field)
	GetIsForward() bool
	// GetTargetNode returns TargetNode (property field)
	GetTargetNode() ExpandedNodeId
	// IsReferenceListEntryDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsReferenceListEntryDataType()
}

// _ReferenceListEntryDataType is the data-structure of this message
type _ReferenceListEntryDataType struct {
	ExtensionObjectDefinitionContract
	ReferenceType NodeId
	IsForward     bool
	TargetNode    ExpandedNodeId
	// Reserved Fields
	reservedField0 *uint8
}

var _ ReferenceListEntryDataType = (*_ReferenceListEntryDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ReferenceListEntryDataType)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ReferenceListEntryDataType) GetIdentifier() string {
	return "32662"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ReferenceListEntryDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ReferenceListEntryDataType) GetReferenceType() NodeId {
	return m.ReferenceType
}

func (m *_ReferenceListEntryDataType) GetIsForward() bool {
	return m.IsForward
}

func (m *_ReferenceListEntryDataType) GetTargetNode() ExpandedNodeId {
	return m.TargetNode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewReferenceListEntryDataType factory function for _ReferenceListEntryDataType
func NewReferenceListEntryDataType(referenceType NodeId, isForward bool, targetNode ExpandedNodeId) *_ReferenceListEntryDataType {
	_result := &_ReferenceListEntryDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ReferenceType:                     referenceType,
		IsForward:                         isForward,
		TargetNode:                        targetNode,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastReferenceListEntryDataType(structType any) ReferenceListEntryDataType {
	if casted, ok := structType.(ReferenceListEntryDataType); ok {
		return casted
	}
	if casted, ok := structType.(*ReferenceListEntryDataType); ok {
		return *casted
	}
	return nil
}

func (m *_ReferenceListEntryDataType) GetTypeName() string {
	return "ReferenceListEntryDataType"
}

func (m *_ReferenceListEntryDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (referenceType)
	lengthInBits += m.ReferenceType.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (isForward)
	lengthInBits += 1

	// Simple field (targetNode)
	lengthInBits += m.TargetNode.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ReferenceListEntryDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ReferenceListEntryDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__referenceListEntryDataType ReferenceListEntryDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ReferenceListEntryDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ReferenceListEntryDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	referenceType, err := ReadSimpleField[NodeId](ctx, "referenceType", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'referenceType' field"))
	}
	m.ReferenceType = referenceType

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

	targetNode, err := ReadSimpleField[ExpandedNodeId](ctx, "targetNode", ReadComplex[ExpandedNodeId](ExpandedNodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'targetNode' field"))
	}
	m.TargetNode = targetNode

	if closeErr := readBuffer.CloseContext("ReferenceListEntryDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ReferenceListEntryDataType")
	}

	return m, nil
}

func (m *_ReferenceListEntryDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ReferenceListEntryDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ReferenceListEntryDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ReferenceListEntryDataType")
		}

		if err := WriteSimpleField[NodeId](ctx, "referenceType", m.GetReferenceType(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'referenceType' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "isForward", m.GetIsForward(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'isForward' field")
		}

		if err := WriteSimpleField[ExpandedNodeId](ctx, "targetNode", m.GetTargetNode(), WriteComplex[ExpandedNodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'targetNode' field")
		}

		if popErr := writeBuffer.PopContext("ReferenceListEntryDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ReferenceListEntryDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ReferenceListEntryDataType) IsReferenceListEntryDataType() {}

func (m *_ReferenceListEntryDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
