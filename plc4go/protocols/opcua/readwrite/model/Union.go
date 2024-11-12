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

// Union is the corresponding interface of Union
type Union interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// IsUnion is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsUnion()
	// CreateBuilder creates a UnionBuilder
	CreateUnionBuilder() UnionBuilder
}

// _Union is the data-structure of this message
type _Union struct {
	ExtensionObjectDefinitionContract
}

var _ Union = (*_Union)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_Union)(nil)

// NewUnion factory function for _Union
func NewUnion() *_Union {
	_result := &_Union{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// UnionBuilder is a builder for Union
type UnionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() UnionBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the Union or returns an error if something is wrong
	Build() (Union, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() Union
}

// NewUnionBuilder() creates a UnionBuilder
func NewUnionBuilder() UnionBuilder {
	return &_UnionBuilder{_Union: new(_Union)}
}

type _UnionBuilder struct {
	*_Union

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (UnionBuilder) = (*_UnionBuilder)(nil)

func (b *_UnionBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._Union
}

func (b *_UnionBuilder) WithMandatoryFields() UnionBuilder {
	return b
}

func (b *_UnionBuilder) Build() (Union, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._Union.deepCopy(), nil
}

func (b *_UnionBuilder) MustBuild() Union {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_UnionBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_UnionBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_UnionBuilder) DeepCopy() any {
	_copy := b.CreateUnionBuilder().(*_UnionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateUnionBuilder creates a UnionBuilder
func (b *_Union) CreateUnionBuilder() UnionBuilder {
	if b == nil {
		return NewUnionBuilder()
	}
	return &_UnionBuilder{_Union: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_Union) GetExtensionId() int32 {
	return int32(12758)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_Union) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

// Deprecated: use the interface for direct cast
func CastUnion(structType any) Union {
	if casted, ok := structType.(Union); ok {
		return casted
	}
	if casted, ok := structType.(*Union); ok {
		return *casted
	}
	return nil
}

func (m *_Union) GetTypeName() string {
	return "Union"
}

func (m *_Union) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_Union) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_Union) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__union Union, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("Union"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Union")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("Union"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Union")
	}

	return m, nil
}

func (m *_Union) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_Union) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("Union"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for Union")
		}

		if popErr := writeBuffer.PopContext("Union"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for Union")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_Union) IsUnion() {}

func (m *_Union) DeepCopy() any {
	return m.deepCopy()
}

func (m *_Union) deepCopy() *_Union {
	if m == nil {
		return nil
	}
	_UnionCopy := &_Union{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
	}
	_UnionCopy.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _UnionCopy
}

func (m *_Union) String() string {
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
