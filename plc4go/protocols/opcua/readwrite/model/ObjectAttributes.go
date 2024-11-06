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

// ObjectAttributes is the corresponding interface of ObjectAttributes
type ObjectAttributes interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetSpecifiedAttributes returns SpecifiedAttributes (property field)
	GetSpecifiedAttributes() uint32
	// GetDisplayName returns DisplayName (property field)
	GetDisplayName() LocalizedText
	// GetDescription returns Description (property field)
	GetDescription() LocalizedText
	// GetWriteMask returns WriteMask (property field)
	GetWriteMask() uint32
	// GetUserWriteMask returns UserWriteMask (property field)
	GetUserWriteMask() uint32
	// GetEventNotifier returns EventNotifier (property field)
	GetEventNotifier() uint8
	// IsObjectAttributes is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsObjectAttributes()
	// CreateBuilder creates a ObjectAttributesBuilder
	CreateObjectAttributesBuilder() ObjectAttributesBuilder
}

// _ObjectAttributes is the data-structure of this message
type _ObjectAttributes struct {
	ExtensionObjectDefinitionContract
	SpecifiedAttributes uint32
	DisplayName         LocalizedText
	Description         LocalizedText
	WriteMask           uint32
	UserWriteMask       uint32
	EventNotifier       uint8
}

var _ ObjectAttributes = (*_ObjectAttributes)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ObjectAttributes)(nil)

// NewObjectAttributes factory function for _ObjectAttributes
func NewObjectAttributes(specifiedAttributes uint32, displayName LocalizedText, description LocalizedText, writeMask uint32, userWriteMask uint32, eventNotifier uint8) *_ObjectAttributes {
	if displayName == nil {
		panic("displayName of type LocalizedText for ObjectAttributes must not be nil")
	}
	if description == nil {
		panic("description of type LocalizedText for ObjectAttributes must not be nil")
	}
	_result := &_ObjectAttributes{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		SpecifiedAttributes:               specifiedAttributes,
		DisplayName:                       displayName,
		Description:                       description,
		WriteMask:                         writeMask,
		UserWriteMask:                     userWriteMask,
		EventNotifier:                     eventNotifier,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ObjectAttributesBuilder is a builder for ObjectAttributes
type ObjectAttributesBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(specifiedAttributes uint32, displayName LocalizedText, description LocalizedText, writeMask uint32, userWriteMask uint32, eventNotifier uint8) ObjectAttributesBuilder
	// WithSpecifiedAttributes adds SpecifiedAttributes (property field)
	WithSpecifiedAttributes(uint32) ObjectAttributesBuilder
	// WithDisplayName adds DisplayName (property field)
	WithDisplayName(LocalizedText) ObjectAttributesBuilder
	// WithDisplayNameBuilder adds DisplayName (property field) which is build by the builder
	WithDisplayNameBuilder(func(LocalizedTextBuilder) LocalizedTextBuilder) ObjectAttributesBuilder
	// WithDescription adds Description (property field)
	WithDescription(LocalizedText) ObjectAttributesBuilder
	// WithDescriptionBuilder adds Description (property field) which is build by the builder
	WithDescriptionBuilder(func(LocalizedTextBuilder) LocalizedTextBuilder) ObjectAttributesBuilder
	// WithWriteMask adds WriteMask (property field)
	WithWriteMask(uint32) ObjectAttributesBuilder
	// WithUserWriteMask adds UserWriteMask (property field)
	WithUserWriteMask(uint32) ObjectAttributesBuilder
	// WithEventNotifier adds EventNotifier (property field)
	WithEventNotifier(uint8) ObjectAttributesBuilder
	// Build builds the ObjectAttributes or returns an error if something is wrong
	Build() (ObjectAttributes, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ObjectAttributes
}

// NewObjectAttributesBuilder() creates a ObjectAttributesBuilder
func NewObjectAttributesBuilder() ObjectAttributesBuilder {
	return &_ObjectAttributesBuilder{_ObjectAttributes: new(_ObjectAttributes)}
}

type _ObjectAttributesBuilder struct {
	*_ObjectAttributes

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (ObjectAttributesBuilder) = (*_ObjectAttributesBuilder)(nil)

func (b *_ObjectAttributesBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_ObjectAttributesBuilder) WithMandatoryFields(specifiedAttributes uint32, displayName LocalizedText, description LocalizedText, writeMask uint32, userWriteMask uint32, eventNotifier uint8) ObjectAttributesBuilder {
	return b.WithSpecifiedAttributes(specifiedAttributes).WithDisplayName(displayName).WithDescription(description).WithWriteMask(writeMask).WithUserWriteMask(userWriteMask).WithEventNotifier(eventNotifier)
}

func (b *_ObjectAttributesBuilder) WithSpecifiedAttributes(specifiedAttributes uint32) ObjectAttributesBuilder {
	b.SpecifiedAttributes = specifiedAttributes
	return b
}

func (b *_ObjectAttributesBuilder) WithDisplayName(displayName LocalizedText) ObjectAttributesBuilder {
	b.DisplayName = displayName
	return b
}

func (b *_ObjectAttributesBuilder) WithDisplayNameBuilder(builderSupplier func(LocalizedTextBuilder) LocalizedTextBuilder) ObjectAttributesBuilder {
	builder := builderSupplier(b.DisplayName.CreateLocalizedTextBuilder())
	var err error
	b.DisplayName, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "LocalizedTextBuilder failed"))
	}
	return b
}

func (b *_ObjectAttributesBuilder) WithDescription(description LocalizedText) ObjectAttributesBuilder {
	b.Description = description
	return b
}

func (b *_ObjectAttributesBuilder) WithDescriptionBuilder(builderSupplier func(LocalizedTextBuilder) LocalizedTextBuilder) ObjectAttributesBuilder {
	builder := builderSupplier(b.Description.CreateLocalizedTextBuilder())
	var err error
	b.Description, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "LocalizedTextBuilder failed"))
	}
	return b
}

func (b *_ObjectAttributesBuilder) WithWriteMask(writeMask uint32) ObjectAttributesBuilder {
	b.WriteMask = writeMask
	return b
}

func (b *_ObjectAttributesBuilder) WithUserWriteMask(userWriteMask uint32) ObjectAttributesBuilder {
	b.UserWriteMask = userWriteMask
	return b
}

func (b *_ObjectAttributesBuilder) WithEventNotifier(eventNotifier uint8) ObjectAttributesBuilder {
	b.EventNotifier = eventNotifier
	return b
}

func (b *_ObjectAttributesBuilder) Build() (ObjectAttributes, error) {
	if b.DisplayName == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'displayName' not set"))
	}
	if b.Description == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'description' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ObjectAttributes.deepCopy(), nil
}

func (b *_ObjectAttributesBuilder) MustBuild() ObjectAttributes {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_ObjectAttributesBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_ObjectAttributesBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_ObjectAttributesBuilder) DeepCopy() any {
	_copy := b.CreateObjectAttributesBuilder().(*_ObjectAttributesBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateObjectAttributesBuilder creates a ObjectAttributesBuilder
func (b *_ObjectAttributes) CreateObjectAttributesBuilder() ObjectAttributesBuilder {
	if b == nil {
		return NewObjectAttributesBuilder()
	}
	return &_ObjectAttributesBuilder{_ObjectAttributes: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ObjectAttributes) GetExtensionId() int32 {
	return int32(354)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ObjectAttributes) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ObjectAttributes) GetSpecifiedAttributes() uint32 {
	return m.SpecifiedAttributes
}

func (m *_ObjectAttributes) GetDisplayName() LocalizedText {
	return m.DisplayName
}

func (m *_ObjectAttributes) GetDescription() LocalizedText {
	return m.Description
}

func (m *_ObjectAttributes) GetWriteMask() uint32 {
	return m.WriteMask
}

func (m *_ObjectAttributes) GetUserWriteMask() uint32 {
	return m.UserWriteMask
}

func (m *_ObjectAttributes) GetEventNotifier() uint8 {
	return m.EventNotifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastObjectAttributes(structType any) ObjectAttributes {
	if casted, ok := structType.(ObjectAttributes); ok {
		return casted
	}
	if casted, ok := structType.(*ObjectAttributes); ok {
		return *casted
	}
	return nil
}

func (m *_ObjectAttributes) GetTypeName() string {
	return "ObjectAttributes"
}

func (m *_ObjectAttributes) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (specifiedAttributes)
	lengthInBits += 32

	// Simple field (displayName)
	lengthInBits += m.DisplayName.GetLengthInBits(ctx)

	// Simple field (description)
	lengthInBits += m.Description.GetLengthInBits(ctx)

	// Simple field (writeMask)
	lengthInBits += 32

	// Simple field (userWriteMask)
	lengthInBits += 32

	// Simple field (eventNotifier)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ObjectAttributes) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ObjectAttributes) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__objectAttributes ObjectAttributes, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ObjectAttributes"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ObjectAttributes")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	specifiedAttributes, err := ReadSimpleField(ctx, "specifiedAttributes", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'specifiedAttributes' field"))
	}
	m.SpecifiedAttributes = specifiedAttributes

	displayName, err := ReadSimpleField[LocalizedText](ctx, "displayName", ReadComplex[LocalizedText](LocalizedTextParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'displayName' field"))
	}
	m.DisplayName = displayName

	description, err := ReadSimpleField[LocalizedText](ctx, "description", ReadComplex[LocalizedText](LocalizedTextParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'description' field"))
	}
	m.Description = description

	writeMask, err := ReadSimpleField(ctx, "writeMask", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'writeMask' field"))
	}
	m.WriteMask = writeMask

	userWriteMask, err := ReadSimpleField(ctx, "userWriteMask", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'userWriteMask' field"))
	}
	m.UserWriteMask = userWriteMask

	eventNotifier, err := ReadSimpleField(ctx, "eventNotifier", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventNotifier' field"))
	}
	m.EventNotifier = eventNotifier

	if closeErr := readBuffer.CloseContext("ObjectAttributes"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ObjectAttributes")
	}

	return m, nil
}

func (m *_ObjectAttributes) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ObjectAttributes) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ObjectAttributes"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ObjectAttributes")
		}

		if err := WriteSimpleField[uint32](ctx, "specifiedAttributes", m.GetSpecifiedAttributes(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'specifiedAttributes' field")
		}

		if err := WriteSimpleField[LocalizedText](ctx, "displayName", m.GetDisplayName(), WriteComplex[LocalizedText](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'displayName' field")
		}

		if err := WriteSimpleField[LocalizedText](ctx, "description", m.GetDescription(), WriteComplex[LocalizedText](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'description' field")
		}

		if err := WriteSimpleField[uint32](ctx, "writeMask", m.GetWriteMask(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'writeMask' field")
		}

		if err := WriteSimpleField[uint32](ctx, "userWriteMask", m.GetUserWriteMask(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'userWriteMask' field")
		}

		if err := WriteSimpleField[uint8](ctx, "eventNotifier", m.GetEventNotifier(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'eventNotifier' field")
		}

		if popErr := writeBuffer.PopContext("ObjectAttributes"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ObjectAttributes")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ObjectAttributes) IsObjectAttributes() {}

func (m *_ObjectAttributes) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ObjectAttributes) deepCopy() *_ObjectAttributes {
	if m == nil {
		return nil
	}
	_ObjectAttributesCopy := &_ObjectAttributes{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.SpecifiedAttributes,
		m.DisplayName.DeepCopy().(LocalizedText),
		m.Description.DeepCopy().(LocalizedText),
		m.WriteMask,
		m.UserWriteMask,
		m.EventNotifier,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _ObjectAttributesCopy
}

func (m *_ObjectAttributes) String() string {
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
