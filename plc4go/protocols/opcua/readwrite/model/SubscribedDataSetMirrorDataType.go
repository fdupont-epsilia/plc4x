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

// SubscribedDataSetMirrorDataType is the corresponding interface of SubscribedDataSetMirrorDataType
type SubscribedDataSetMirrorDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetParentNodeName returns ParentNodeName (property field)
	GetParentNodeName() PascalString
	// GetRolePermissions returns RolePermissions (property field)
	GetRolePermissions() []RolePermissionType
	// IsSubscribedDataSetMirrorDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSubscribedDataSetMirrorDataType()
	// CreateBuilder creates a SubscribedDataSetMirrorDataTypeBuilder
	CreateSubscribedDataSetMirrorDataTypeBuilder() SubscribedDataSetMirrorDataTypeBuilder
}

// _SubscribedDataSetMirrorDataType is the data-structure of this message
type _SubscribedDataSetMirrorDataType struct {
	ExtensionObjectDefinitionContract
	ParentNodeName  PascalString
	RolePermissions []RolePermissionType
}

var _ SubscribedDataSetMirrorDataType = (*_SubscribedDataSetMirrorDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_SubscribedDataSetMirrorDataType)(nil)

// NewSubscribedDataSetMirrorDataType factory function for _SubscribedDataSetMirrorDataType
func NewSubscribedDataSetMirrorDataType(parentNodeName PascalString, rolePermissions []RolePermissionType) *_SubscribedDataSetMirrorDataType {
	if parentNodeName == nil {
		panic("parentNodeName of type PascalString for SubscribedDataSetMirrorDataType must not be nil")
	}
	_result := &_SubscribedDataSetMirrorDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ParentNodeName:                    parentNodeName,
		RolePermissions:                   rolePermissions,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SubscribedDataSetMirrorDataTypeBuilder is a builder for SubscribedDataSetMirrorDataType
type SubscribedDataSetMirrorDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(parentNodeName PascalString, rolePermissions []RolePermissionType) SubscribedDataSetMirrorDataTypeBuilder
	// WithParentNodeName adds ParentNodeName (property field)
	WithParentNodeName(PascalString) SubscribedDataSetMirrorDataTypeBuilder
	// WithParentNodeNameBuilder adds ParentNodeName (property field) which is build by the builder
	WithParentNodeNameBuilder(func(PascalStringBuilder) PascalStringBuilder) SubscribedDataSetMirrorDataTypeBuilder
	// WithRolePermissions adds RolePermissions (property field)
	WithRolePermissions(...RolePermissionType) SubscribedDataSetMirrorDataTypeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the SubscribedDataSetMirrorDataType or returns an error if something is wrong
	Build() (SubscribedDataSetMirrorDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SubscribedDataSetMirrorDataType
}

// NewSubscribedDataSetMirrorDataTypeBuilder() creates a SubscribedDataSetMirrorDataTypeBuilder
func NewSubscribedDataSetMirrorDataTypeBuilder() SubscribedDataSetMirrorDataTypeBuilder {
	return &_SubscribedDataSetMirrorDataTypeBuilder{_SubscribedDataSetMirrorDataType: new(_SubscribedDataSetMirrorDataType)}
}

type _SubscribedDataSetMirrorDataTypeBuilder struct {
	*_SubscribedDataSetMirrorDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (SubscribedDataSetMirrorDataTypeBuilder) = (*_SubscribedDataSetMirrorDataTypeBuilder)(nil)

func (b *_SubscribedDataSetMirrorDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._SubscribedDataSetMirrorDataType
}

func (b *_SubscribedDataSetMirrorDataTypeBuilder) WithMandatoryFields(parentNodeName PascalString, rolePermissions []RolePermissionType) SubscribedDataSetMirrorDataTypeBuilder {
	return b.WithParentNodeName(parentNodeName).WithRolePermissions(rolePermissions...)
}

func (b *_SubscribedDataSetMirrorDataTypeBuilder) WithParentNodeName(parentNodeName PascalString) SubscribedDataSetMirrorDataTypeBuilder {
	b.ParentNodeName = parentNodeName
	return b
}

func (b *_SubscribedDataSetMirrorDataTypeBuilder) WithParentNodeNameBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) SubscribedDataSetMirrorDataTypeBuilder {
	builder := builderSupplier(b.ParentNodeName.CreatePascalStringBuilder())
	var err error
	b.ParentNodeName, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_SubscribedDataSetMirrorDataTypeBuilder) WithRolePermissions(rolePermissions ...RolePermissionType) SubscribedDataSetMirrorDataTypeBuilder {
	b.RolePermissions = rolePermissions
	return b
}

func (b *_SubscribedDataSetMirrorDataTypeBuilder) Build() (SubscribedDataSetMirrorDataType, error) {
	if b.ParentNodeName == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'parentNodeName' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SubscribedDataSetMirrorDataType.deepCopy(), nil
}

func (b *_SubscribedDataSetMirrorDataTypeBuilder) MustBuild() SubscribedDataSetMirrorDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SubscribedDataSetMirrorDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_SubscribedDataSetMirrorDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_SubscribedDataSetMirrorDataTypeBuilder) DeepCopy() any {
	_copy := b.CreateSubscribedDataSetMirrorDataTypeBuilder().(*_SubscribedDataSetMirrorDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSubscribedDataSetMirrorDataTypeBuilder creates a SubscribedDataSetMirrorDataTypeBuilder
func (b *_SubscribedDataSetMirrorDataType) CreateSubscribedDataSetMirrorDataTypeBuilder() SubscribedDataSetMirrorDataTypeBuilder {
	if b == nil {
		return NewSubscribedDataSetMirrorDataTypeBuilder()
	}
	return &_SubscribedDataSetMirrorDataTypeBuilder{_SubscribedDataSetMirrorDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SubscribedDataSetMirrorDataType) GetExtensionId() int32 {
	return int32(15637)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SubscribedDataSetMirrorDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SubscribedDataSetMirrorDataType) GetParentNodeName() PascalString {
	return m.ParentNodeName
}

func (m *_SubscribedDataSetMirrorDataType) GetRolePermissions() []RolePermissionType {
	return m.RolePermissions
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSubscribedDataSetMirrorDataType(structType any) SubscribedDataSetMirrorDataType {
	if casted, ok := structType.(SubscribedDataSetMirrorDataType); ok {
		return casted
	}
	if casted, ok := structType.(*SubscribedDataSetMirrorDataType); ok {
		return *casted
	}
	return nil
}

func (m *_SubscribedDataSetMirrorDataType) GetTypeName() string {
	return "SubscribedDataSetMirrorDataType"
}

func (m *_SubscribedDataSetMirrorDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (parentNodeName)
	lengthInBits += m.ParentNodeName.GetLengthInBits(ctx)

	// Implicit Field (noOfRolePermissions)
	lengthInBits += 32

	// Array field
	if len(m.RolePermissions) > 0 {
		for _curItem, element := range m.RolePermissions {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.RolePermissions), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_SubscribedDataSetMirrorDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SubscribedDataSetMirrorDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__subscribedDataSetMirrorDataType SubscribedDataSetMirrorDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SubscribedDataSetMirrorDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SubscribedDataSetMirrorDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	parentNodeName, err := ReadSimpleField[PascalString](ctx, "parentNodeName", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'parentNodeName' field"))
	}
	m.ParentNodeName = parentNodeName

	noOfRolePermissions, err := ReadImplicitField[int32](ctx, "noOfRolePermissions", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfRolePermissions' field"))
	}
	_ = noOfRolePermissions

	rolePermissions, err := ReadCountArrayField[RolePermissionType](ctx, "rolePermissions", ReadComplex[RolePermissionType](ExtensionObjectDefinitionParseWithBufferProducer[RolePermissionType]((int32)(int32(98))), readBuffer), uint64(noOfRolePermissions))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'rolePermissions' field"))
	}
	m.RolePermissions = rolePermissions

	if closeErr := readBuffer.CloseContext("SubscribedDataSetMirrorDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SubscribedDataSetMirrorDataType")
	}

	return m, nil
}

func (m *_SubscribedDataSetMirrorDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SubscribedDataSetMirrorDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SubscribedDataSetMirrorDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SubscribedDataSetMirrorDataType")
		}

		if err := WriteSimpleField[PascalString](ctx, "parentNodeName", m.GetParentNodeName(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'parentNodeName' field")
		}
		noOfRolePermissions := int32(utils.InlineIf(bool((m.GetRolePermissions()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetRolePermissions()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfRolePermissions", noOfRolePermissions, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfRolePermissions' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "rolePermissions", m.GetRolePermissions(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'rolePermissions' field")
		}

		if popErr := writeBuffer.PopContext("SubscribedDataSetMirrorDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SubscribedDataSetMirrorDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SubscribedDataSetMirrorDataType) IsSubscribedDataSetMirrorDataType() {}

func (m *_SubscribedDataSetMirrorDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SubscribedDataSetMirrorDataType) deepCopy() *_SubscribedDataSetMirrorDataType {
	if m == nil {
		return nil
	}
	_SubscribedDataSetMirrorDataTypeCopy := &_SubscribedDataSetMirrorDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[PascalString](m.ParentNodeName),
		utils.DeepCopySlice[RolePermissionType, RolePermissionType](m.RolePermissions),
	}
	_SubscribedDataSetMirrorDataTypeCopy.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _SubscribedDataSetMirrorDataTypeCopy
}

func (m *_SubscribedDataSetMirrorDataType) String() string {
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
