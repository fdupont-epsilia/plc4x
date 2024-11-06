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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// WriterGroupMessageDataType is the corresponding interface of WriterGroupMessageDataType
type WriterGroupMessageDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// IsWriterGroupMessageDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsWriterGroupMessageDataType()
	// CreateBuilder creates a WriterGroupMessageDataTypeBuilder
	CreateWriterGroupMessageDataTypeBuilder() WriterGroupMessageDataTypeBuilder
}

// _WriterGroupMessageDataType is the data-structure of this message
type _WriterGroupMessageDataType struct {
	ExtensionObjectDefinitionContract
}

var _ WriterGroupMessageDataType = (*_WriterGroupMessageDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_WriterGroupMessageDataType)(nil)

// NewWriterGroupMessageDataType factory function for _WriterGroupMessageDataType
func NewWriterGroupMessageDataType() *_WriterGroupMessageDataType {
	_result := &_WriterGroupMessageDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// WriterGroupMessageDataTypeBuilder is a builder for WriterGroupMessageDataType
type WriterGroupMessageDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() WriterGroupMessageDataTypeBuilder
	// Build builds the WriterGroupMessageDataType or returns an error if something is wrong
	Build() (WriterGroupMessageDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() WriterGroupMessageDataType
}

// NewWriterGroupMessageDataTypeBuilder() creates a WriterGroupMessageDataTypeBuilder
func NewWriterGroupMessageDataTypeBuilder() WriterGroupMessageDataTypeBuilder {
	return &_WriterGroupMessageDataTypeBuilder{_WriterGroupMessageDataType: new(_WriterGroupMessageDataType)}
}

type _WriterGroupMessageDataTypeBuilder struct {
	*_WriterGroupMessageDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (WriterGroupMessageDataTypeBuilder) = (*_WriterGroupMessageDataTypeBuilder)(nil)

func (b *_WriterGroupMessageDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_WriterGroupMessageDataTypeBuilder) WithMandatoryFields() WriterGroupMessageDataTypeBuilder {
	return b
}

func (b *_WriterGroupMessageDataTypeBuilder) Build() (WriterGroupMessageDataType, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._WriterGroupMessageDataType.deepCopy(), nil
}

func (b *_WriterGroupMessageDataTypeBuilder) MustBuild() WriterGroupMessageDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_WriterGroupMessageDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_WriterGroupMessageDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_WriterGroupMessageDataTypeBuilder) DeepCopy() any {
	_copy := b.CreateWriterGroupMessageDataTypeBuilder().(*_WriterGroupMessageDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateWriterGroupMessageDataTypeBuilder creates a WriterGroupMessageDataTypeBuilder
func (b *_WriterGroupMessageDataType) CreateWriterGroupMessageDataTypeBuilder() WriterGroupMessageDataTypeBuilder {
	if b == nil {
		return NewWriterGroupMessageDataTypeBuilder()
	}
	return &_WriterGroupMessageDataTypeBuilder{_WriterGroupMessageDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_WriterGroupMessageDataType) GetExtensionId() int32 {
	return int32(15618)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_WriterGroupMessageDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

// Deprecated: use the interface for direct cast
func CastWriterGroupMessageDataType(structType any) WriterGroupMessageDataType {
	if casted, ok := structType.(WriterGroupMessageDataType); ok {
		return casted
	}
	if casted, ok := structType.(*WriterGroupMessageDataType); ok {
		return *casted
	}
	return nil
}

func (m *_WriterGroupMessageDataType) GetTypeName() string {
	return "WriterGroupMessageDataType"
}

func (m *_WriterGroupMessageDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_WriterGroupMessageDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_WriterGroupMessageDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__writerGroupMessageDataType WriterGroupMessageDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("WriterGroupMessageDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for WriterGroupMessageDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("WriterGroupMessageDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for WriterGroupMessageDataType")
	}

	return m, nil
}

func (m *_WriterGroupMessageDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_WriterGroupMessageDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("WriterGroupMessageDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for WriterGroupMessageDataType")
		}

		if popErr := writeBuffer.PopContext("WriterGroupMessageDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for WriterGroupMessageDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_WriterGroupMessageDataType) IsWriterGroupMessageDataType() {}

func (m *_WriterGroupMessageDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_WriterGroupMessageDataType) deepCopy() *_WriterGroupMessageDataType {
	if m == nil {
		return nil
	}
	_WriterGroupMessageDataTypeCopy := &_WriterGroupMessageDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _WriterGroupMessageDataTypeCopy
}

func (m *_WriterGroupMessageDataType) String() string {
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
