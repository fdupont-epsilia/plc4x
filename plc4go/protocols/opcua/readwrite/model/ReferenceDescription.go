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

// ReferenceDescription is the corresponding interface of ReferenceDescription
type ReferenceDescription interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetReferenceTypeId returns ReferenceTypeId (property field)
	GetReferenceTypeId() NodeId
	// GetIsForward returns IsForward (property field)
	GetIsForward() bool
	// GetNodeId returns NodeId (property field)
	GetNodeId() ExpandedNodeId
	// GetBrowseName returns BrowseName (property field)
	GetBrowseName() QualifiedName
	// GetDisplayName returns DisplayName (property field)
	GetDisplayName() LocalizedText
	// GetNodeClass returns NodeClass (property field)
	GetNodeClass() NodeClass
	// GetTypeDefinition returns TypeDefinition (property field)
	GetTypeDefinition() ExpandedNodeId
	// IsReferenceDescription is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsReferenceDescription()
}

// _ReferenceDescription is the data-structure of this message
type _ReferenceDescription struct {
	ExtensionObjectDefinitionContract
	ReferenceTypeId NodeId
	IsForward       bool
	NodeId          ExpandedNodeId
	BrowseName      QualifiedName
	DisplayName     LocalizedText
	NodeClass       NodeClass
	TypeDefinition  ExpandedNodeId
	// Reserved Fields
	reservedField0 *uint8
}

var _ ReferenceDescription = (*_ReferenceDescription)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ReferenceDescription)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ReferenceDescription) GetIdentifier() string {
	return "520"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ReferenceDescription) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ReferenceDescription) GetReferenceTypeId() NodeId {
	return m.ReferenceTypeId
}

func (m *_ReferenceDescription) GetIsForward() bool {
	return m.IsForward
}

func (m *_ReferenceDescription) GetNodeId() ExpandedNodeId {
	return m.NodeId
}

func (m *_ReferenceDescription) GetBrowseName() QualifiedName {
	return m.BrowseName
}

func (m *_ReferenceDescription) GetDisplayName() LocalizedText {
	return m.DisplayName
}

func (m *_ReferenceDescription) GetNodeClass() NodeClass {
	return m.NodeClass
}

func (m *_ReferenceDescription) GetTypeDefinition() ExpandedNodeId {
	return m.TypeDefinition
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewReferenceDescription factory function for _ReferenceDescription
func NewReferenceDescription(referenceTypeId NodeId, isForward bool, nodeId ExpandedNodeId, browseName QualifiedName, displayName LocalizedText, nodeClass NodeClass, typeDefinition ExpandedNodeId) *_ReferenceDescription {
	_result := &_ReferenceDescription{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ReferenceTypeId:                   referenceTypeId,
		IsForward:                         isForward,
		NodeId:                            nodeId,
		BrowseName:                        browseName,
		DisplayName:                       displayName,
		NodeClass:                         nodeClass,
		TypeDefinition:                    typeDefinition,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastReferenceDescription(structType any) ReferenceDescription {
	if casted, ok := structType.(ReferenceDescription); ok {
		return casted
	}
	if casted, ok := structType.(*ReferenceDescription); ok {
		return *casted
	}
	return nil
}

func (m *_ReferenceDescription) GetTypeName() string {
	return "ReferenceDescription"
}

func (m *_ReferenceDescription) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (referenceTypeId)
	lengthInBits += m.ReferenceTypeId.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (isForward)
	lengthInBits += 1

	// Simple field (nodeId)
	lengthInBits += m.NodeId.GetLengthInBits(ctx)

	// Simple field (browseName)
	lengthInBits += m.BrowseName.GetLengthInBits(ctx)

	// Simple field (displayName)
	lengthInBits += m.DisplayName.GetLengthInBits(ctx)

	// Simple field (nodeClass)
	lengthInBits += 32

	// Simple field (typeDefinition)
	lengthInBits += m.TypeDefinition.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ReferenceDescription) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ReferenceDescription) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__referenceDescription ReferenceDescription, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ReferenceDescription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ReferenceDescription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

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

	nodeId, err := ReadSimpleField[ExpandedNodeId](ctx, "nodeId", ReadComplex[ExpandedNodeId](ExpandedNodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodeId' field"))
	}
	m.NodeId = nodeId

	browseName, err := ReadSimpleField[QualifiedName](ctx, "browseName", ReadComplex[QualifiedName](QualifiedNameParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'browseName' field"))
	}
	m.BrowseName = browseName

	displayName, err := ReadSimpleField[LocalizedText](ctx, "displayName", ReadComplex[LocalizedText](LocalizedTextParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'displayName' field"))
	}
	m.DisplayName = displayName

	nodeClass, err := ReadEnumField[NodeClass](ctx, "nodeClass", "NodeClass", ReadEnum(NodeClassByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodeClass' field"))
	}
	m.NodeClass = nodeClass

	typeDefinition, err := ReadSimpleField[ExpandedNodeId](ctx, "typeDefinition", ReadComplex[ExpandedNodeId](ExpandedNodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'typeDefinition' field"))
	}
	m.TypeDefinition = typeDefinition

	if closeErr := readBuffer.CloseContext("ReferenceDescription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ReferenceDescription")
	}

	return m, nil
}

func (m *_ReferenceDescription) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ReferenceDescription) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ReferenceDescription"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ReferenceDescription")
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

		if err := WriteSimpleField[ExpandedNodeId](ctx, "nodeId", m.GetNodeId(), WriteComplex[ExpandedNodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'nodeId' field")
		}

		if err := WriteSimpleField[QualifiedName](ctx, "browseName", m.GetBrowseName(), WriteComplex[QualifiedName](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'browseName' field")
		}

		if err := WriteSimpleField[LocalizedText](ctx, "displayName", m.GetDisplayName(), WriteComplex[LocalizedText](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'displayName' field")
		}

		if err := WriteSimpleEnumField[NodeClass](ctx, "nodeClass", "NodeClass", m.GetNodeClass(), WriteEnum[NodeClass, uint32](NodeClass.GetValue, NodeClass.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'nodeClass' field")
		}

		if err := WriteSimpleField[ExpandedNodeId](ctx, "typeDefinition", m.GetTypeDefinition(), WriteComplex[ExpandedNodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'typeDefinition' field")
		}

		if popErr := writeBuffer.PopContext("ReferenceDescription"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ReferenceDescription")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ReferenceDescription) IsReferenceDescription() {}

func (m *_ReferenceDescription) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
