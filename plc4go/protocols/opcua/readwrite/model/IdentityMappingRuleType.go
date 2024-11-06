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

// IdentityMappingRuleType is the corresponding interface of IdentityMappingRuleType
type IdentityMappingRuleType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetCriteriaType returns CriteriaType (property field)
	GetCriteriaType() IdentityCriteriaType
	// GetCriteria returns Criteria (property field)
	GetCriteria() PascalString
	// IsIdentityMappingRuleType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsIdentityMappingRuleType()
	// CreateBuilder creates a IdentityMappingRuleTypeBuilder
	CreateIdentityMappingRuleTypeBuilder() IdentityMappingRuleTypeBuilder
}

// _IdentityMappingRuleType is the data-structure of this message
type _IdentityMappingRuleType struct {
	ExtensionObjectDefinitionContract
	CriteriaType IdentityCriteriaType
	Criteria     PascalString
}

var _ IdentityMappingRuleType = (*_IdentityMappingRuleType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_IdentityMappingRuleType)(nil)

// NewIdentityMappingRuleType factory function for _IdentityMappingRuleType
func NewIdentityMappingRuleType(criteriaType IdentityCriteriaType, criteria PascalString) *_IdentityMappingRuleType {
	if criteria == nil {
		panic("criteria of type PascalString for IdentityMappingRuleType must not be nil")
	}
	_result := &_IdentityMappingRuleType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		CriteriaType:                      criteriaType,
		Criteria:                          criteria,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// IdentityMappingRuleTypeBuilder is a builder for IdentityMappingRuleType
type IdentityMappingRuleTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(criteriaType IdentityCriteriaType, criteria PascalString) IdentityMappingRuleTypeBuilder
	// WithCriteriaType adds CriteriaType (property field)
	WithCriteriaType(IdentityCriteriaType) IdentityMappingRuleTypeBuilder
	// WithCriteria adds Criteria (property field)
	WithCriteria(PascalString) IdentityMappingRuleTypeBuilder
	// WithCriteriaBuilder adds Criteria (property field) which is build by the builder
	WithCriteriaBuilder(func(PascalStringBuilder) PascalStringBuilder) IdentityMappingRuleTypeBuilder
	// Build builds the IdentityMappingRuleType or returns an error if something is wrong
	Build() (IdentityMappingRuleType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() IdentityMappingRuleType
}

// NewIdentityMappingRuleTypeBuilder() creates a IdentityMappingRuleTypeBuilder
func NewIdentityMappingRuleTypeBuilder() IdentityMappingRuleTypeBuilder {
	return &_IdentityMappingRuleTypeBuilder{_IdentityMappingRuleType: new(_IdentityMappingRuleType)}
}

type _IdentityMappingRuleTypeBuilder struct {
	*_IdentityMappingRuleType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (IdentityMappingRuleTypeBuilder) = (*_IdentityMappingRuleTypeBuilder)(nil)

func (b *_IdentityMappingRuleTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_IdentityMappingRuleTypeBuilder) WithMandatoryFields(criteriaType IdentityCriteriaType, criteria PascalString) IdentityMappingRuleTypeBuilder {
	return b.WithCriteriaType(criteriaType).WithCriteria(criteria)
}

func (b *_IdentityMappingRuleTypeBuilder) WithCriteriaType(criteriaType IdentityCriteriaType) IdentityMappingRuleTypeBuilder {
	b.CriteriaType = criteriaType
	return b
}

func (b *_IdentityMappingRuleTypeBuilder) WithCriteria(criteria PascalString) IdentityMappingRuleTypeBuilder {
	b.Criteria = criteria
	return b
}

func (b *_IdentityMappingRuleTypeBuilder) WithCriteriaBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) IdentityMappingRuleTypeBuilder {
	builder := builderSupplier(b.Criteria.CreatePascalStringBuilder())
	var err error
	b.Criteria, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_IdentityMappingRuleTypeBuilder) Build() (IdentityMappingRuleType, error) {
	if b.Criteria == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'criteria' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._IdentityMappingRuleType.deepCopy(), nil
}

func (b *_IdentityMappingRuleTypeBuilder) MustBuild() IdentityMappingRuleType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_IdentityMappingRuleTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_IdentityMappingRuleTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_IdentityMappingRuleTypeBuilder) DeepCopy() any {
	_copy := b.CreateIdentityMappingRuleTypeBuilder().(*_IdentityMappingRuleTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateIdentityMappingRuleTypeBuilder creates a IdentityMappingRuleTypeBuilder
func (b *_IdentityMappingRuleType) CreateIdentityMappingRuleTypeBuilder() IdentityMappingRuleTypeBuilder {
	if b == nil {
		return NewIdentityMappingRuleTypeBuilder()
	}
	return &_IdentityMappingRuleTypeBuilder{_IdentityMappingRuleType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentityMappingRuleType) GetExtensionId() int32 {
	return int32(15636)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentityMappingRuleType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentityMappingRuleType) GetCriteriaType() IdentityCriteriaType {
	return m.CriteriaType
}

func (m *_IdentityMappingRuleType) GetCriteria() PascalString {
	return m.Criteria
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastIdentityMappingRuleType(structType any) IdentityMappingRuleType {
	if casted, ok := structType.(IdentityMappingRuleType); ok {
		return casted
	}
	if casted, ok := structType.(*IdentityMappingRuleType); ok {
		return *casted
	}
	return nil
}

func (m *_IdentityMappingRuleType) GetTypeName() string {
	return "IdentityMappingRuleType"
}

func (m *_IdentityMappingRuleType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (criteriaType)
	lengthInBits += 32

	// Simple field (criteria)
	lengthInBits += m.Criteria.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_IdentityMappingRuleType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_IdentityMappingRuleType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__identityMappingRuleType IdentityMappingRuleType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentityMappingRuleType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentityMappingRuleType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	criteriaType, err := ReadEnumField[IdentityCriteriaType](ctx, "criteriaType", "IdentityCriteriaType", ReadEnum(IdentityCriteriaTypeByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'criteriaType' field"))
	}
	m.CriteriaType = criteriaType

	criteria, err := ReadSimpleField[PascalString](ctx, "criteria", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'criteria' field"))
	}
	m.Criteria = criteria

	if closeErr := readBuffer.CloseContext("IdentityMappingRuleType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentityMappingRuleType")
	}

	return m, nil
}

func (m *_IdentityMappingRuleType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IdentityMappingRuleType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentityMappingRuleType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentityMappingRuleType")
		}

		if err := WriteSimpleEnumField[IdentityCriteriaType](ctx, "criteriaType", "IdentityCriteriaType", m.GetCriteriaType(), WriteEnum[IdentityCriteriaType, uint32](IdentityCriteriaType.GetValue, IdentityCriteriaType.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'criteriaType' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "criteria", m.GetCriteria(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'criteria' field")
		}

		if popErr := writeBuffer.PopContext("IdentityMappingRuleType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentityMappingRuleType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_IdentityMappingRuleType) IsIdentityMappingRuleType() {}

func (m *_IdentityMappingRuleType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_IdentityMappingRuleType) deepCopy() *_IdentityMappingRuleType {
	if m == nil {
		return nil
	}
	_IdentityMappingRuleTypeCopy := &_IdentityMappingRuleType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.CriteriaType,
		m.Criteria.DeepCopy().(PascalString),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _IdentityMappingRuleTypeCopy
}

func (m *_IdentityMappingRuleType) String() string {
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
