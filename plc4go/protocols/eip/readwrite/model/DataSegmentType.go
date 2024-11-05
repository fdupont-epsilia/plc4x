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

// DataSegmentType is the corresponding interface of DataSegmentType
type DataSegmentType interface {
	DataSegmentTypeContract
	DataSegmentTypeRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsDataSegmentType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDataSegmentType()
	// CreateBuilder creates a DataSegmentTypeBuilder
	CreateDataSegmentTypeBuilder() DataSegmentTypeBuilder
}

// DataSegmentTypeContract provides a set of functions which can be overwritten by a sub struct
type DataSegmentTypeContract interface {
	// IsDataSegmentType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDataSegmentType()
	// CreateBuilder creates a DataSegmentTypeBuilder
	CreateDataSegmentTypeBuilder() DataSegmentTypeBuilder
}

// DataSegmentTypeRequirements provides a set of functions which need to be implemented by a sub struct
type DataSegmentTypeRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetDataSegmentType returns DataSegmentType (discriminator field)
	GetDataSegmentType() uint8
}

// _DataSegmentType is the data-structure of this message
type _DataSegmentType struct {
	_SubType interface {
		DataSegmentTypeContract
		DataSegmentTypeRequirements
	}
}

var _ DataSegmentTypeContract = (*_DataSegmentType)(nil)

// NewDataSegmentType factory function for _DataSegmentType
func NewDataSegmentType() *_DataSegmentType {
	return &_DataSegmentType{}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DataSegmentTypeBuilder is a builder for DataSegmentType
type DataSegmentTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() DataSegmentTypeBuilder
	// AsAnsiExtendedSymbolSegment converts this build to a subType of DataSegmentType. It is always possible to return to current builder using Done()
	AsAnsiExtendedSymbolSegment() interface {
		AnsiExtendedSymbolSegmentBuilder
		Done() DataSegmentTypeBuilder
	}
	// Build builds the DataSegmentType or returns an error if something is wrong
	PartialBuild() (DataSegmentTypeContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() DataSegmentTypeContract
	// Build builds the DataSegmentType or returns an error if something is wrong
	Build() (DataSegmentType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DataSegmentType
}

// NewDataSegmentTypeBuilder() creates a DataSegmentTypeBuilder
func NewDataSegmentTypeBuilder() DataSegmentTypeBuilder {
	return &_DataSegmentTypeBuilder{_DataSegmentType: new(_DataSegmentType)}
}

type _DataSegmentTypeChildBuilder interface {
	utils.Copyable
	setParent(DataSegmentTypeContract)
	buildForDataSegmentType() (DataSegmentType, error)
}

type _DataSegmentTypeBuilder struct {
	*_DataSegmentType

	childBuilder _DataSegmentTypeChildBuilder

	err *utils.MultiError
}

var _ (DataSegmentTypeBuilder) = (*_DataSegmentTypeBuilder)(nil)

func (b *_DataSegmentTypeBuilder) WithMandatoryFields() DataSegmentTypeBuilder {
	return b
}

func (b *_DataSegmentTypeBuilder) PartialBuild() (DataSegmentTypeContract, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._DataSegmentType.deepCopy(), nil
}

func (b *_DataSegmentTypeBuilder) PartialMustBuild() DataSegmentTypeContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_DataSegmentTypeBuilder) AsAnsiExtendedSymbolSegment() interface {
	AnsiExtendedSymbolSegmentBuilder
	Done() DataSegmentTypeBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		AnsiExtendedSymbolSegmentBuilder
		Done() DataSegmentTypeBuilder
	}); ok {
		return cb
	}
	cb := NewAnsiExtendedSymbolSegmentBuilder().(*_AnsiExtendedSymbolSegmentBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_DataSegmentTypeBuilder) Build() (DataSegmentType, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForDataSegmentType()
}

func (b *_DataSegmentTypeBuilder) MustBuild() DataSegmentType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_DataSegmentTypeBuilder) DeepCopy() any {
	_copy := b.CreateDataSegmentTypeBuilder().(*_DataSegmentTypeBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_DataSegmentTypeChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateDataSegmentTypeBuilder creates a DataSegmentTypeBuilder
func (b *_DataSegmentType) CreateDataSegmentTypeBuilder() DataSegmentTypeBuilder {
	if b == nil {
		return NewDataSegmentTypeBuilder()
	}
	return &_DataSegmentTypeBuilder{_DataSegmentType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDataSegmentType(structType any) DataSegmentType {
	if casted, ok := structType.(DataSegmentType); ok {
		return casted
	}
	if casted, ok := structType.(*DataSegmentType); ok {
		return *casted
	}
	return nil
}

func (m *_DataSegmentType) GetTypeName() string {
	return "DataSegmentType"
}

func (m *_DataSegmentType) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (dataSegmentType)
	lengthInBits += 5

	return lengthInBits
}

func (m *_DataSegmentType) GetLengthInBits(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx)
}

func (m *_DataSegmentType) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func DataSegmentTypeParse[T DataSegmentType](ctx context.Context, theBytes []byte) (T, error) {
	return DataSegmentTypeParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func DataSegmentTypeParseWithBufferProducer[T DataSegmentType]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := DataSegmentTypeParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func DataSegmentTypeParseWithBuffer[T DataSegmentType](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_DataSegmentType{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_DataSegmentType) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__dataSegmentType DataSegmentType, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DataSegmentType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DataSegmentType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dataSegmentType, err := ReadDiscriminatorField[uint8](ctx, "dataSegmentType", ReadUnsignedByte(readBuffer, uint8(5)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataSegmentType' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child DataSegmentType
	switch {
	case dataSegmentType == 0x11: // AnsiExtendedSymbolSegment
		if _child, err = new(_AnsiExtendedSymbolSegment).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AnsiExtendedSymbolSegment for type-switch of DataSegmentType")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [dataSegmentType=%v]", dataSegmentType)
	}

	if closeErr := readBuffer.CloseContext("DataSegmentType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DataSegmentType")
	}

	return _child, nil
}

func (pm *_DataSegmentType) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child DataSegmentType, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("DataSegmentType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for DataSegmentType")
	}

	if err := WriteDiscriminatorField(ctx, "dataSegmentType", m.GetDataSegmentType(), WriteUnsignedByte(writeBuffer, 5)); err != nil {
		return errors.Wrap(err, "Error serializing 'dataSegmentType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("DataSegmentType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for DataSegmentType")
	}
	return nil
}

func (m *_DataSegmentType) IsDataSegmentType() {}

func (m *_DataSegmentType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DataSegmentType) deepCopy() *_DataSegmentType {
	if m == nil {
		return nil
	}
	_DataSegmentTypeCopy := &_DataSegmentType{
		nil, // will be set by child
	}
	return _DataSegmentTypeCopy
}
