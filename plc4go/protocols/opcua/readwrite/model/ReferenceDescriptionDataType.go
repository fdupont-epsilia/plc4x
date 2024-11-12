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

// ReferenceDescriptionDataType is the corresponding interface of ReferenceDescriptionDataType
type ReferenceDescriptionDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetSourceNode returns SourceNode (property field)
	GetSourceNode() NodeId
	// GetReferenceType returns ReferenceType (property field)
	GetReferenceType() NodeId
	// GetIsForward returns IsForward (property field)
	GetIsForward() bool
	// GetTargetNode returns TargetNode (property field)
	GetTargetNode() ExpandedNodeId
	// IsReferenceDescriptionDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsReferenceDescriptionDataType()
	// CreateBuilder creates a ReferenceDescriptionDataTypeBuilder
	CreateReferenceDescriptionDataTypeBuilder() ReferenceDescriptionDataTypeBuilder
}

// _ReferenceDescriptionDataType is the data-structure of this message
type _ReferenceDescriptionDataType struct {
	ExtensionObjectDefinitionContract
	SourceNode    NodeId
	ReferenceType NodeId
	IsForward     bool
	TargetNode    ExpandedNodeId
	// Reserved Fields
	reservedField0 *uint8
}

var _ ReferenceDescriptionDataType = (*_ReferenceDescriptionDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ReferenceDescriptionDataType)(nil)

// NewReferenceDescriptionDataType factory function for _ReferenceDescriptionDataType
func NewReferenceDescriptionDataType(sourceNode NodeId, referenceType NodeId, isForward bool, targetNode ExpandedNodeId) *_ReferenceDescriptionDataType {
	if sourceNode == nil {
		panic("sourceNode of type NodeId for ReferenceDescriptionDataType must not be nil")
	}
	if referenceType == nil {
		panic("referenceType of type NodeId for ReferenceDescriptionDataType must not be nil")
	}
	if targetNode == nil {
		panic("targetNode of type ExpandedNodeId for ReferenceDescriptionDataType must not be nil")
	}
	_result := &_ReferenceDescriptionDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		SourceNode:                        sourceNode,
		ReferenceType:                     referenceType,
		IsForward:                         isForward,
		TargetNode:                        targetNode,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ReferenceDescriptionDataTypeBuilder is a builder for ReferenceDescriptionDataType
type ReferenceDescriptionDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(sourceNode NodeId, referenceType NodeId, isForward bool, targetNode ExpandedNodeId) ReferenceDescriptionDataTypeBuilder
	// WithSourceNode adds SourceNode (property field)
	WithSourceNode(NodeId) ReferenceDescriptionDataTypeBuilder
	// WithSourceNodeBuilder adds SourceNode (property field) which is build by the builder
	WithSourceNodeBuilder(func(NodeIdBuilder) NodeIdBuilder) ReferenceDescriptionDataTypeBuilder
	// WithReferenceType adds ReferenceType (property field)
	WithReferenceType(NodeId) ReferenceDescriptionDataTypeBuilder
	// WithReferenceTypeBuilder adds ReferenceType (property field) which is build by the builder
	WithReferenceTypeBuilder(func(NodeIdBuilder) NodeIdBuilder) ReferenceDescriptionDataTypeBuilder
	// WithIsForward adds IsForward (property field)
	WithIsForward(bool) ReferenceDescriptionDataTypeBuilder
	// WithTargetNode adds TargetNode (property field)
	WithTargetNode(ExpandedNodeId) ReferenceDescriptionDataTypeBuilder
	// WithTargetNodeBuilder adds TargetNode (property field) which is build by the builder
	WithTargetNodeBuilder(func(ExpandedNodeIdBuilder) ExpandedNodeIdBuilder) ReferenceDescriptionDataTypeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the ReferenceDescriptionDataType or returns an error if something is wrong
	Build() (ReferenceDescriptionDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ReferenceDescriptionDataType
}

// NewReferenceDescriptionDataTypeBuilder() creates a ReferenceDescriptionDataTypeBuilder
func NewReferenceDescriptionDataTypeBuilder() ReferenceDescriptionDataTypeBuilder {
	return &_ReferenceDescriptionDataTypeBuilder{_ReferenceDescriptionDataType: new(_ReferenceDescriptionDataType)}
}

type _ReferenceDescriptionDataTypeBuilder struct {
	*_ReferenceDescriptionDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (ReferenceDescriptionDataTypeBuilder) = (*_ReferenceDescriptionDataTypeBuilder)(nil)

func (b *_ReferenceDescriptionDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._ReferenceDescriptionDataType
}

func (b *_ReferenceDescriptionDataTypeBuilder) WithMandatoryFields(sourceNode NodeId, referenceType NodeId, isForward bool, targetNode ExpandedNodeId) ReferenceDescriptionDataTypeBuilder {
	return b.WithSourceNode(sourceNode).WithReferenceType(referenceType).WithIsForward(isForward).WithTargetNode(targetNode)
}

func (b *_ReferenceDescriptionDataTypeBuilder) WithSourceNode(sourceNode NodeId) ReferenceDescriptionDataTypeBuilder {
	b.SourceNode = sourceNode
	return b
}

func (b *_ReferenceDescriptionDataTypeBuilder) WithSourceNodeBuilder(builderSupplier func(NodeIdBuilder) NodeIdBuilder) ReferenceDescriptionDataTypeBuilder {
	builder := builderSupplier(b.SourceNode.CreateNodeIdBuilder())
	var err error
	b.SourceNode, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "NodeIdBuilder failed"))
	}
	return b
}

func (b *_ReferenceDescriptionDataTypeBuilder) WithReferenceType(referenceType NodeId) ReferenceDescriptionDataTypeBuilder {
	b.ReferenceType = referenceType
	return b
}

func (b *_ReferenceDescriptionDataTypeBuilder) WithReferenceTypeBuilder(builderSupplier func(NodeIdBuilder) NodeIdBuilder) ReferenceDescriptionDataTypeBuilder {
	builder := builderSupplier(b.ReferenceType.CreateNodeIdBuilder())
	var err error
	b.ReferenceType, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "NodeIdBuilder failed"))
	}
	return b
}

func (b *_ReferenceDescriptionDataTypeBuilder) WithIsForward(isForward bool) ReferenceDescriptionDataTypeBuilder {
	b.IsForward = isForward
	return b
}

func (b *_ReferenceDescriptionDataTypeBuilder) WithTargetNode(targetNode ExpandedNodeId) ReferenceDescriptionDataTypeBuilder {
	b.TargetNode = targetNode
	return b
}

func (b *_ReferenceDescriptionDataTypeBuilder) WithTargetNodeBuilder(builderSupplier func(ExpandedNodeIdBuilder) ExpandedNodeIdBuilder) ReferenceDescriptionDataTypeBuilder {
	builder := builderSupplier(b.TargetNode.CreateExpandedNodeIdBuilder())
	var err error
	b.TargetNode, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExpandedNodeIdBuilder failed"))
	}
	return b
}

func (b *_ReferenceDescriptionDataTypeBuilder) Build() (ReferenceDescriptionDataType, error) {
	if b.SourceNode == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'sourceNode' not set"))
	}
	if b.ReferenceType == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'referenceType' not set"))
	}
	if b.TargetNode == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'targetNode' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ReferenceDescriptionDataType.deepCopy(), nil
}

func (b *_ReferenceDescriptionDataTypeBuilder) MustBuild() ReferenceDescriptionDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ReferenceDescriptionDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_ReferenceDescriptionDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_ReferenceDescriptionDataTypeBuilder) DeepCopy() any {
	_copy := b.CreateReferenceDescriptionDataTypeBuilder().(*_ReferenceDescriptionDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateReferenceDescriptionDataTypeBuilder creates a ReferenceDescriptionDataTypeBuilder
func (b *_ReferenceDescriptionDataType) CreateReferenceDescriptionDataTypeBuilder() ReferenceDescriptionDataTypeBuilder {
	if b == nil {
		return NewReferenceDescriptionDataTypeBuilder()
	}
	return &_ReferenceDescriptionDataTypeBuilder{_ReferenceDescriptionDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ReferenceDescriptionDataType) GetExtensionId() int32 {
	return int32(32661)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ReferenceDescriptionDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ReferenceDescriptionDataType) GetSourceNode() NodeId {
	return m.SourceNode
}

func (m *_ReferenceDescriptionDataType) GetReferenceType() NodeId {
	return m.ReferenceType
}

func (m *_ReferenceDescriptionDataType) GetIsForward() bool {
	return m.IsForward
}

func (m *_ReferenceDescriptionDataType) GetTargetNode() ExpandedNodeId {
	return m.TargetNode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastReferenceDescriptionDataType(structType any) ReferenceDescriptionDataType {
	if casted, ok := structType.(ReferenceDescriptionDataType); ok {
		return casted
	}
	if casted, ok := structType.(*ReferenceDescriptionDataType); ok {
		return *casted
	}
	return nil
}

func (m *_ReferenceDescriptionDataType) GetTypeName() string {
	return "ReferenceDescriptionDataType"
}

func (m *_ReferenceDescriptionDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (sourceNode)
	lengthInBits += m.SourceNode.GetLengthInBits(ctx)

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

func (m *_ReferenceDescriptionDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ReferenceDescriptionDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__referenceDescriptionDataType ReferenceDescriptionDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ReferenceDescriptionDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ReferenceDescriptionDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	sourceNode, err := ReadSimpleField[NodeId](ctx, "sourceNode", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sourceNode' field"))
	}
	m.SourceNode = sourceNode

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

	if closeErr := readBuffer.CloseContext("ReferenceDescriptionDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ReferenceDescriptionDataType")
	}

	return m, nil
}

func (m *_ReferenceDescriptionDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ReferenceDescriptionDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ReferenceDescriptionDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ReferenceDescriptionDataType")
		}

		if err := WriteSimpleField[NodeId](ctx, "sourceNode", m.GetSourceNode(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'sourceNode' field")
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

		if popErr := writeBuffer.PopContext("ReferenceDescriptionDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ReferenceDescriptionDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ReferenceDescriptionDataType) IsReferenceDescriptionDataType() {}

func (m *_ReferenceDescriptionDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ReferenceDescriptionDataType) deepCopy() *_ReferenceDescriptionDataType {
	if m == nil {
		return nil
	}
	_ReferenceDescriptionDataTypeCopy := &_ReferenceDescriptionDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[NodeId](m.SourceNode),
		utils.DeepCopy[NodeId](m.ReferenceType),
		m.IsForward,
		utils.DeepCopy[ExpandedNodeId](m.TargetNode),
		m.reservedField0,
	}
	_ReferenceDescriptionDataTypeCopy.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _ReferenceDescriptionDataTypeCopy
}

func (m *_ReferenceDescriptionDataType) String() string {
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
