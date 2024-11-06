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

// TransmitQosPriorityDataType is the corresponding interface of TransmitQosPriorityDataType
type TransmitQosPriorityDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetPriorityLabel returns PriorityLabel (property field)
	GetPriorityLabel() PascalString
	// IsTransmitQosPriorityDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTransmitQosPriorityDataType()
	// CreateBuilder creates a TransmitQosPriorityDataTypeBuilder
	CreateTransmitQosPriorityDataTypeBuilder() TransmitQosPriorityDataTypeBuilder
}

// _TransmitQosPriorityDataType is the data-structure of this message
type _TransmitQosPriorityDataType struct {
	ExtensionObjectDefinitionContract
	PriorityLabel PascalString
}

var _ TransmitQosPriorityDataType = (*_TransmitQosPriorityDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_TransmitQosPriorityDataType)(nil)

// NewTransmitQosPriorityDataType factory function for _TransmitQosPriorityDataType
func NewTransmitQosPriorityDataType(priorityLabel PascalString) *_TransmitQosPriorityDataType {
	if priorityLabel == nil {
		panic("priorityLabel of type PascalString for TransmitQosPriorityDataType must not be nil")
	}
	_result := &_TransmitQosPriorityDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		PriorityLabel:                     priorityLabel,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// TransmitQosPriorityDataTypeBuilder is a builder for TransmitQosPriorityDataType
type TransmitQosPriorityDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(priorityLabel PascalString) TransmitQosPriorityDataTypeBuilder
	// WithPriorityLabel adds PriorityLabel (property field)
	WithPriorityLabel(PascalString) TransmitQosPriorityDataTypeBuilder
	// WithPriorityLabelBuilder adds PriorityLabel (property field) which is build by the builder
	WithPriorityLabelBuilder(func(PascalStringBuilder) PascalStringBuilder) TransmitQosPriorityDataTypeBuilder
	// Build builds the TransmitQosPriorityDataType or returns an error if something is wrong
	Build() (TransmitQosPriorityDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() TransmitQosPriorityDataType
}

// NewTransmitQosPriorityDataTypeBuilder() creates a TransmitQosPriorityDataTypeBuilder
func NewTransmitQosPriorityDataTypeBuilder() TransmitQosPriorityDataTypeBuilder {
	return &_TransmitQosPriorityDataTypeBuilder{_TransmitQosPriorityDataType: new(_TransmitQosPriorityDataType)}
}

type _TransmitQosPriorityDataTypeBuilder struct {
	*_TransmitQosPriorityDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (TransmitQosPriorityDataTypeBuilder) = (*_TransmitQosPriorityDataTypeBuilder)(nil)

func (b *_TransmitQosPriorityDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_TransmitQosPriorityDataTypeBuilder) WithMandatoryFields(priorityLabel PascalString) TransmitQosPriorityDataTypeBuilder {
	return b.WithPriorityLabel(priorityLabel)
}

func (b *_TransmitQosPriorityDataTypeBuilder) WithPriorityLabel(priorityLabel PascalString) TransmitQosPriorityDataTypeBuilder {
	b.PriorityLabel = priorityLabel
	return b
}

func (b *_TransmitQosPriorityDataTypeBuilder) WithPriorityLabelBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) TransmitQosPriorityDataTypeBuilder {
	builder := builderSupplier(b.PriorityLabel.CreatePascalStringBuilder())
	var err error
	b.PriorityLabel, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_TransmitQosPriorityDataTypeBuilder) Build() (TransmitQosPriorityDataType, error) {
	if b.PriorityLabel == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'priorityLabel' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._TransmitQosPriorityDataType.deepCopy(), nil
}

func (b *_TransmitQosPriorityDataTypeBuilder) MustBuild() TransmitQosPriorityDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_TransmitQosPriorityDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_TransmitQosPriorityDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_TransmitQosPriorityDataTypeBuilder) DeepCopy() any {
	_copy := b.CreateTransmitQosPriorityDataTypeBuilder().(*_TransmitQosPriorityDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateTransmitQosPriorityDataTypeBuilder creates a TransmitQosPriorityDataTypeBuilder
func (b *_TransmitQosPriorityDataType) CreateTransmitQosPriorityDataTypeBuilder() TransmitQosPriorityDataTypeBuilder {
	if b == nil {
		return NewTransmitQosPriorityDataTypeBuilder()
	}
	return &_TransmitQosPriorityDataTypeBuilder{_TransmitQosPriorityDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_TransmitQosPriorityDataType) GetExtensionId() int32 {
	return int32(23607)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TransmitQosPriorityDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TransmitQosPriorityDataType) GetPriorityLabel() PascalString {
	return m.PriorityLabel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastTransmitQosPriorityDataType(structType any) TransmitQosPriorityDataType {
	if casted, ok := structType.(TransmitQosPriorityDataType); ok {
		return casted
	}
	if casted, ok := structType.(*TransmitQosPriorityDataType); ok {
		return *casted
	}
	return nil
}

func (m *_TransmitQosPriorityDataType) GetTypeName() string {
	return "TransmitQosPriorityDataType"
}

func (m *_TransmitQosPriorityDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (priorityLabel)
	lengthInBits += m.PriorityLabel.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_TransmitQosPriorityDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_TransmitQosPriorityDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__transmitQosPriorityDataType TransmitQosPriorityDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TransmitQosPriorityDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TransmitQosPriorityDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	priorityLabel, err := ReadSimpleField[PascalString](ctx, "priorityLabel", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priorityLabel' field"))
	}
	m.PriorityLabel = priorityLabel

	if closeErr := readBuffer.CloseContext("TransmitQosPriorityDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TransmitQosPriorityDataType")
	}

	return m, nil
}

func (m *_TransmitQosPriorityDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TransmitQosPriorityDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TransmitQosPriorityDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TransmitQosPriorityDataType")
		}

		if err := WriteSimpleField[PascalString](ctx, "priorityLabel", m.GetPriorityLabel(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'priorityLabel' field")
		}

		if popErr := writeBuffer.PopContext("TransmitQosPriorityDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TransmitQosPriorityDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TransmitQosPriorityDataType) IsTransmitQosPriorityDataType() {}

func (m *_TransmitQosPriorityDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_TransmitQosPriorityDataType) deepCopy() *_TransmitQosPriorityDataType {
	if m == nil {
		return nil
	}
	_TransmitQosPriorityDataTypeCopy := &_TransmitQosPriorityDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.PriorityLabel.DeepCopy().(PascalString),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _TransmitQosPriorityDataTypeCopy
}

func (m *_TransmitQosPriorityDataType) String() string {
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