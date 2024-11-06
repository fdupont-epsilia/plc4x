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

// Range is the corresponding interface of Range
type Range interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetLow returns Low (property field)
	GetLow() float64
	// GetHigh returns High (property field)
	GetHigh() float64
	// IsRange is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsRange()
	// CreateBuilder creates a RangeBuilder
	CreateRangeBuilder() RangeBuilder
}

// _Range is the data-structure of this message
type _Range struct {
	ExtensionObjectDefinitionContract
	Low  float64
	High float64
}

var _ Range = (*_Range)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_Range)(nil)

// NewRange factory function for _Range
func NewRange(low float64, high float64) *_Range {
	_result := &_Range{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Low:                               low,
		High:                              high,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// RangeBuilder is a builder for Range
type RangeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(low float64, high float64) RangeBuilder
	// WithLow adds Low (property field)
	WithLow(float64) RangeBuilder
	// WithHigh adds High (property field)
	WithHigh(float64) RangeBuilder
	// Build builds the Range or returns an error if something is wrong
	Build() (Range, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() Range
}

// NewRangeBuilder() creates a RangeBuilder
func NewRangeBuilder() RangeBuilder {
	return &_RangeBuilder{_Range: new(_Range)}
}

type _RangeBuilder struct {
	*_Range

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (RangeBuilder) = (*_RangeBuilder)(nil)

func (b *_RangeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_RangeBuilder) WithMandatoryFields(low float64, high float64) RangeBuilder {
	return b.WithLow(low).WithHigh(high)
}

func (b *_RangeBuilder) WithLow(low float64) RangeBuilder {
	b.Low = low
	return b
}

func (b *_RangeBuilder) WithHigh(high float64) RangeBuilder {
	b.High = high
	return b
}

func (b *_RangeBuilder) Build() (Range, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._Range.deepCopy(), nil
}

func (b *_RangeBuilder) MustBuild() Range {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_RangeBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_RangeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_RangeBuilder) DeepCopy() any {
	_copy := b.CreateRangeBuilder().(*_RangeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateRangeBuilder creates a RangeBuilder
func (b *_Range) CreateRangeBuilder() RangeBuilder {
	if b == nil {
		return NewRangeBuilder()
	}
	return &_RangeBuilder{_Range: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_Range) GetExtensionId() int32 {
	return int32(886)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_Range) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_Range) GetLow() float64 {
	return m.Low
}

func (m *_Range) GetHigh() float64 {
	return m.High
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastRange(structType any) Range {
	if casted, ok := structType.(Range); ok {
		return casted
	}
	if casted, ok := structType.(*Range); ok {
		return *casted
	}
	return nil
}

func (m *_Range) GetTypeName() string {
	return "Range"
}

func (m *_Range) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (low)
	lengthInBits += 64

	// Simple field (high)
	lengthInBits += 64

	return lengthInBits
}

func (m *_Range) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_Range) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__range Range, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("Range"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Range")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	low, err := ReadSimpleField(ctx, "low", ReadDouble(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'low' field"))
	}
	m.Low = low

	high, err := ReadSimpleField(ctx, "high", ReadDouble(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'high' field"))
	}
	m.High = high

	if closeErr := readBuffer.CloseContext("Range"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Range")
	}

	return m, nil
}

func (m *_Range) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_Range) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("Range"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for Range")
		}

		if err := WriteSimpleField[float64](ctx, "low", m.GetLow(), WriteDouble(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'low' field")
		}

		if err := WriteSimpleField[float64](ctx, "high", m.GetHigh(), WriteDouble(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'high' field")
		}

		if popErr := writeBuffer.PopContext("Range"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for Range")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_Range) IsRange() {}

func (m *_Range) DeepCopy() any {
	return m.deepCopy()
}

func (m *_Range) deepCopy() *_Range {
	if m == nil {
		return nil
	}
	_RangeCopy := &_Range{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.Low,
		m.High,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _RangeCopy
}

func (m *_Range) String() string {
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
