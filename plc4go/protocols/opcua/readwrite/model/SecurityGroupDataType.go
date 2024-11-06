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

// SecurityGroupDataType is the corresponding interface of SecurityGroupDataType
type SecurityGroupDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetName returns Name (property field)
	GetName() PascalString
	// GetSecurityGroupFolder returns SecurityGroupFolder (property field)
	GetSecurityGroupFolder() []PascalString
	// GetKeyLifetime returns KeyLifetime (property field)
	GetKeyLifetime() float64
	// GetSecurityPolicyUri returns SecurityPolicyUri (property field)
	GetSecurityPolicyUri() PascalString
	// GetMaxFutureKeyCount returns MaxFutureKeyCount (property field)
	GetMaxFutureKeyCount() uint32
	// GetMaxPastKeyCount returns MaxPastKeyCount (property field)
	GetMaxPastKeyCount() uint32
	// GetSecurityGroupId returns SecurityGroupId (property field)
	GetSecurityGroupId() PascalString
	// GetRolePermissions returns RolePermissions (property field)
	GetRolePermissions() []RolePermissionType
	// GetGroupProperties returns GroupProperties (property field)
	GetGroupProperties() []KeyValuePair
	// IsSecurityGroupDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityGroupDataType()
	// CreateBuilder creates a SecurityGroupDataTypeBuilder
	CreateSecurityGroupDataTypeBuilder() SecurityGroupDataTypeBuilder
}

// _SecurityGroupDataType is the data-structure of this message
type _SecurityGroupDataType struct {
	ExtensionObjectDefinitionContract
	Name                PascalString
	SecurityGroupFolder []PascalString
	KeyLifetime         float64
	SecurityPolicyUri   PascalString
	MaxFutureKeyCount   uint32
	MaxPastKeyCount     uint32
	SecurityGroupId     PascalString
	RolePermissions     []RolePermissionType
	GroupProperties     []KeyValuePair
}

var _ SecurityGroupDataType = (*_SecurityGroupDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_SecurityGroupDataType)(nil)

// NewSecurityGroupDataType factory function for _SecurityGroupDataType
func NewSecurityGroupDataType(name PascalString, securityGroupFolder []PascalString, keyLifetime float64, securityPolicyUri PascalString, maxFutureKeyCount uint32, maxPastKeyCount uint32, securityGroupId PascalString, rolePermissions []RolePermissionType, groupProperties []KeyValuePair) *_SecurityGroupDataType {
	if name == nil {
		panic("name of type PascalString for SecurityGroupDataType must not be nil")
	}
	if securityPolicyUri == nil {
		panic("securityPolicyUri of type PascalString for SecurityGroupDataType must not be nil")
	}
	if securityGroupId == nil {
		panic("securityGroupId of type PascalString for SecurityGroupDataType must not be nil")
	}
	_result := &_SecurityGroupDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Name:                              name,
		SecurityGroupFolder:               securityGroupFolder,
		KeyLifetime:                       keyLifetime,
		SecurityPolicyUri:                 securityPolicyUri,
		MaxFutureKeyCount:                 maxFutureKeyCount,
		MaxPastKeyCount:                   maxPastKeyCount,
		SecurityGroupId:                   securityGroupId,
		RolePermissions:                   rolePermissions,
		GroupProperties:                   groupProperties,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SecurityGroupDataTypeBuilder is a builder for SecurityGroupDataType
type SecurityGroupDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(name PascalString, securityGroupFolder []PascalString, keyLifetime float64, securityPolicyUri PascalString, maxFutureKeyCount uint32, maxPastKeyCount uint32, securityGroupId PascalString, rolePermissions []RolePermissionType, groupProperties []KeyValuePair) SecurityGroupDataTypeBuilder
	// WithName adds Name (property field)
	WithName(PascalString) SecurityGroupDataTypeBuilder
	// WithNameBuilder adds Name (property field) which is build by the builder
	WithNameBuilder(func(PascalStringBuilder) PascalStringBuilder) SecurityGroupDataTypeBuilder
	// WithSecurityGroupFolder adds SecurityGroupFolder (property field)
	WithSecurityGroupFolder(...PascalString) SecurityGroupDataTypeBuilder
	// WithKeyLifetime adds KeyLifetime (property field)
	WithKeyLifetime(float64) SecurityGroupDataTypeBuilder
	// WithSecurityPolicyUri adds SecurityPolicyUri (property field)
	WithSecurityPolicyUri(PascalString) SecurityGroupDataTypeBuilder
	// WithSecurityPolicyUriBuilder adds SecurityPolicyUri (property field) which is build by the builder
	WithSecurityPolicyUriBuilder(func(PascalStringBuilder) PascalStringBuilder) SecurityGroupDataTypeBuilder
	// WithMaxFutureKeyCount adds MaxFutureKeyCount (property field)
	WithMaxFutureKeyCount(uint32) SecurityGroupDataTypeBuilder
	// WithMaxPastKeyCount adds MaxPastKeyCount (property field)
	WithMaxPastKeyCount(uint32) SecurityGroupDataTypeBuilder
	// WithSecurityGroupId adds SecurityGroupId (property field)
	WithSecurityGroupId(PascalString) SecurityGroupDataTypeBuilder
	// WithSecurityGroupIdBuilder adds SecurityGroupId (property field) which is build by the builder
	WithSecurityGroupIdBuilder(func(PascalStringBuilder) PascalStringBuilder) SecurityGroupDataTypeBuilder
	// WithRolePermissions adds RolePermissions (property field)
	WithRolePermissions(...RolePermissionType) SecurityGroupDataTypeBuilder
	// WithGroupProperties adds GroupProperties (property field)
	WithGroupProperties(...KeyValuePair) SecurityGroupDataTypeBuilder
	// Build builds the SecurityGroupDataType or returns an error if something is wrong
	Build() (SecurityGroupDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SecurityGroupDataType
}

// NewSecurityGroupDataTypeBuilder() creates a SecurityGroupDataTypeBuilder
func NewSecurityGroupDataTypeBuilder() SecurityGroupDataTypeBuilder {
	return &_SecurityGroupDataTypeBuilder{_SecurityGroupDataType: new(_SecurityGroupDataType)}
}

type _SecurityGroupDataTypeBuilder struct {
	*_SecurityGroupDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (SecurityGroupDataTypeBuilder) = (*_SecurityGroupDataTypeBuilder)(nil)

func (b *_SecurityGroupDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_SecurityGroupDataTypeBuilder) WithMandatoryFields(name PascalString, securityGroupFolder []PascalString, keyLifetime float64, securityPolicyUri PascalString, maxFutureKeyCount uint32, maxPastKeyCount uint32, securityGroupId PascalString, rolePermissions []RolePermissionType, groupProperties []KeyValuePair) SecurityGroupDataTypeBuilder {
	return b.WithName(name).WithSecurityGroupFolder(securityGroupFolder...).WithKeyLifetime(keyLifetime).WithSecurityPolicyUri(securityPolicyUri).WithMaxFutureKeyCount(maxFutureKeyCount).WithMaxPastKeyCount(maxPastKeyCount).WithSecurityGroupId(securityGroupId).WithRolePermissions(rolePermissions...).WithGroupProperties(groupProperties...)
}

func (b *_SecurityGroupDataTypeBuilder) WithName(name PascalString) SecurityGroupDataTypeBuilder {
	b.Name = name
	return b
}

func (b *_SecurityGroupDataTypeBuilder) WithNameBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) SecurityGroupDataTypeBuilder {
	builder := builderSupplier(b.Name.CreatePascalStringBuilder())
	var err error
	b.Name, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_SecurityGroupDataTypeBuilder) WithSecurityGroupFolder(securityGroupFolder ...PascalString) SecurityGroupDataTypeBuilder {
	b.SecurityGroupFolder = securityGroupFolder
	return b
}

func (b *_SecurityGroupDataTypeBuilder) WithKeyLifetime(keyLifetime float64) SecurityGroupDataTypeBuilder {
	b.KeyLifetime = keyLifetime
	return b
}

func (b *_SecurityGroupDataTypeBuilder) WithSecurityPolicyUri(securityPolicyUri PascalString) SecurityGroupDataTypeBuilder {
	b.SecurityPolicyUri = securityPolicyUri
	return b
}

func (b *_SecurityGroupDataTypeBuilder) WithSecurityPolicyUriBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) SecurityGroupDataTypeBuilder {
	builder := builderSupplier(b.SecurityPolicyUri.CreatePascalStringBuilder())
	var err error
	b.SecurityPolicyUri, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_SecurityGroupDataTypeBuilder) WithMaxFutureKeyCount(maxFutureKeyCount uint32) SecurityGroupDataTypeBuilder {
	b.MaxFutureKeyCount = maxFutureKeyCount
	return b
}

func (b *_SecurityGroupDataTypeBuilder) WithMaxPastKeyCount(maxPastKeyCount uint32) SecurityGroupDataTypeBuilder {
	b.MaxPastKeyCount = maxPastKeyCount
	return b
}

func (b *_SecurityGroupDataTypeBuilder) WithSecurityGroupId(securityGroupId PascalString) SecurityGroupDataTypeBuilder {
	b.SecurityGroupId = securityGroupId
	return b
}

func (b *_SecurityGroupDataTypeBuilder) WithSecurityGroupIdBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) SecurityGroupDataTypeBuilder {
	builder := builderSupplier(b.SecurityGroupId.CreatePascalStringBuilder())
	var err error
	b.SecurityGroupId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_SecurityGroupDataTypeBuilder) WithRolePermissions(rolePermissions ...RolePermissionType) SecurityGroupDataTypeBuilder {
	b.RolePermissions = rolePermissions
	return b
}

func (b *_SecurityGroupDataTypeBuilder) WithGroupProperties(groupProperties ...KeyValuePair) SecurityGroupDataTypeBuilder {
	b.GroupProperties = groupProperties
	return b
}

func (b *_SecurityGroupDataTypeBuilder) Build() (SecurityGroupDataType, error) {
	if b.Name == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'name' not set"))
	}
	if b.SecurityPolicyUri == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'securityPolicyUri' not set"))
	}
	if b.SecurityGroupId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'securityGroupId' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SecurityGroupDataType.deepCopy(), nil
}

func (b *_SecurityGroupDataTypeBuilder) MustBuild() SecurityGroupDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_SecurityGroupDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_SecurityGroupDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_SecurityGroupDataTypeBuilder) DeepCopy() any {
	_copy := b.CreateSecurityGroupDataTypeBuilder().(*_SecurityGroupDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSecurityGroupDataTypeBuilder creates a SecurityGroupDataTypeBuilder
func (b *_SecurityGroupDataType) CreateSecurityGroupDataTypeBuilder() SecurityGroupDataTypeBuilder {
	if b == nil {
		return NewSecurityGroupDataTypeBuilder()
	}
	return &_SecurityGroupDataTypeBuilder{_SecurityGroupDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SecurityGroupDataType) GetExtensionId() int32 {
	return int32(23603)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityGroupDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityGroupDataType) GetName() PascalString {
	return m.Name
}

func (m *_SecurityGroupDataType) GetSecurityGroupFolder() []PascalString {
	return m.SecurityGroupFolder
}

func (m *_SecurityGroupDataType) GetKeyLifetime() float64 {
	return m.KeyLifetime
}

func (m *_SecurityGroupDataType) GetSecurityPolicyUri() PascalString {
	return m.SecurityPolicyUri
}

func (m *_SecurityGroupDataType) GetMaxFutureKeyCount() uint32 {
	return m.MaxFutureKeyCount
}

func (m *_SecurityGroupDataType) GetMaxPastKeyCount() uint32 {
	return m.MaxPastKeyCount
}

func (m *_SecurityGroupDataType) GetSecurityGroupId() PascalString {
	return m.SecurityGroupId
}

func (m *_SecurityGroupDataType) GetRolePermissions() []RolePermissionType {
	return m.RolePermissions
}

func (m *_SecurityGroupDataType) GetGroupProperties() []KeyValuePair {
	return m.GroupProperties
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSecurityGroupDataType(structType any) SecurityGroupDataType {
	if casted, ok := structType.(SecurityGroupDataType); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityGroupDataType); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityGroupDataType) GetTypeName() string {
	return "SecurityGroupDataType"
}

func (m *_SecurityGroupDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (name)
	lengthInBits += m.Name.GetLengthInBits(ctx)

	// Implicit Field (noOfSecurityGroupFolder)
	lengthInBits += 32

	// Array field
	if len(m.SecurityGroupFolder) > 0 {
		for _curItem, element := range m.SecurityGroupFolder {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.SecurityGroupFolder), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (keyLifetime)
	lengthInBits += 64

	// Simple field (securityPolicyUri)
	lengthInBits += m.SecurityPolicyUri.GetLengthInBits(ctx)

	// Simple field (maxFutureKeyCount)
	lengthInBits += 32

	// Simple field (maxPastKeyCount)
	lengthInBits += 32

	// Simple field (securityGroupId)
	lengthInBits += m.SecurityGroupId.GetLengthInBits(ctx)

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

	// Implicit Field (noOfGroupProperties)
	lengthInBits += 32

	// Array field
	if len(m.GroupProperties) > 0 {
		for _curItem, element := range m.GroupProperties {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GroupProperties), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_SecurityGroupDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityGroupDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__securityGroupDataType SecurityGroupDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityGroupDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityGroupDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	name, err := ReadSimpleField[PascalString](ctx, "name", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'name' field"))
	}
	m.Name = name

	noOfSecurityGroupFolder, err := ReadImplicitField[int32](ctx, "noOfSecurityGroupFolder", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfSecurityGroupFolder' field"))
	}
	_ = noOfSecurityGroupFolder

	securityGroupFolder, err := ReadCountArrayField[PascalString](ctx, "securityGroupFolder", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer), uint64(noOfSecurityGroupFolder))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityGroupFolder' field"))
	}
	m.SecurityGroupFolder = securityGroupFolder

	keyLifetime, err := ReadSimpleField(ctx, "keyLifetime", ReadDouble(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'keyLifetime' field"))
	}
	m.KeyLifetime = keyLifetime

	securityPolicyUri, err := ReadSimpleField[PascalString](ctx, "securityPolicyUri", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityPolicyUri' field"))
	}
	m.SecurityPolicyUri = securityPolicyUri

	maxFutureKeyCount, err := ReadSimpleField(ctx, "maxFutureKeyCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxFutureKeyCount' field"))
	}
	m.MaxFutureKeyCount = maxFutureKeyCount

	maxPastKeyCount, err := ReadSimpleField(ctx, "maxPastKeyCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxPastKeyCount' field"))
	}
	m.MaxPastKeyCount = maxPastKeyCount

	securityGroupId, err := ReadSimpleField[PascalString](ctx, "securityGroupId", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityGroupId' field"))
	}
	m.SecurityGroupId = securityGroupId

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

	noOfGroupProperties, err := ReadImplicitField[int32](ctx, "noOfGroupProperties", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfGroupProperties' field"))
	}
	_ = noOfGroupProperties

	groupProperties, err := ReadCountArrayField[KeyValuePair](ctx, "groupProperties", ReadComplex[KeyValuePair](ExtensionObjectDefinitionParseWithBufferProducer[KeyValuePair]((int32)(int32(14535))), readBuffer), uint64(noOfGroupProperties))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'groupProperties' field"))
	}
	m.GroupProperties = groupProperties

	if closeErr := readBuffer.CloseContext("SecurityGroupDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityGroupDataType")
	}

	return m, nil
}

func (m *_SecurityGroupDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityGroupDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityGroupDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityGroupDataType")
		}

		if err := WriteSimpleField[PascalString](ctx, "name", m.GetName(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'name' field")
		}
		noOfSecurityGroupFolder := int32(utils.InlineIf(bool((m.GetSecurityGroupFolder()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetSecurityGroupFolder()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfSecurityGroupFolder", noOfSecurityGroupFolder, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfSecurityGroupFolder' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "securityGroupFolder", m.GetSecurityGroupFolder(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'securityGroupFolder' field")
		}

		if err := WriteSimpleField[float64](ctx, "keyLifetime", m.GetKeyLifetime(), WriteDouble(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'keyLifetime' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "securityPolicyUri", m.GetSecurityPolicyUri(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'securityPolicyUri' field")
		}

		if err := WriteSimpleField[uint32](ctx, "maxFutureKeyCount", m.GetMaxFutureKeyCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxFutureKeyCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "maxPastKeyCount", m.GetMaxPastKeyCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxPastKeyCount' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "securityGroupId", m.GetSecurityGroupId(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'securityGroupId' field")
		}
		noOfRolePermissions := int32(utils.InlineIf(bool((m.GetRolePermissions()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetRolePermissions()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfRolePermissions", noOfRolePermissions, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfRolePermissions' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "rolePermissions", m.GetRolePermissions(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'rolePermissions' field")
		}
		noOfGroupProperties := int32(utils.InlineIf(bool((m.GetGroupProperties()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetGroupProperties()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfGroupProperties", noOfGroupProperties, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfGroupProperties' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "groupProperties", m.GetGroupProperties(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'groupProperties' field")
		}

		if popErr := writeBuffer.PopContext("SecurityGroupDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityGroupDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityGroupDataType) IsSecurityGroupDataType() {}

func (m *_SecurityGroupDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SecurityGroupDataType) deepCopy() *_SecurityGroupDataType {
	if m == nil {
		return nil
	}
	_SecurityGroupDataTypeCopy := &_SecurityGroupDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.Name.DeepCopy().(PascalString),
		utils.DeepCopySlice[PascalString, PascalString](m.SecurityGroupFolder),
		m.KeyLifetime,
		m.SecurityPolicyUri.DeepCopy().(PascalString),
		m.MaxFutureKeyCount,
		m.MaxPastKeyCount,
		m.SecurityGroupId.DeepCopy().(PascalString),
		utils.DeepCopySlice[RolePermissionType, RolePermissionType](m.RolePermissions),
		utils.DeepCopySlice[KeyValuePair, KeyValuePair](m.GroupProperties),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _SecurityGroupDataTypeCopy
}

func (m *_SecurityGroupDataType) String() string {
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
