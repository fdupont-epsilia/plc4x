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

// VariantStatusCode is the corresponding interface of VariantStatusCode
type VariantStatusCode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	Variant
	// GetArrayLength returns ArrayLength (property field)
	GetArrayLength() *int32
	// GetValue returns Value (property field)
	GetValue() []StatusCode
	// IsVariantStatusCode is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsVariantStatusCode()
	// CreateBuilder creates a VariantStatusCodeBuilder
	CreateVariantStatusCodeBuilder() VariantStatusCodeBuilder
}

// _VariantStatusCode is the data-structure of this message
type _VariantStatusCode struct {
	VariantContract
	ArrayLength *int32
	Value       []StatusCode
}

var _ VariantStatusCode = (*_VariantStatusCode)(nil)
var _ VariantRequirements = (*_VariantStatusCode)(nil)

// NewVariantStatusCode factory function for _VariantStatusCode
func NewVariantStatusCode(arrayLengthSpecified bool, arrayDimensionsSpecified bool, noOfArrayDimensions *int32, arrayDimensions []bool, arrayLength *int32, value []StatusCode) *_VariantStatusCode {
	_result := &_VariantStatusCode{
		VariantContract: NewVariant(arrayLengthSpecified, arrayDimensionsSpecified, noOfArrayDimensions, arrayDimensions),
		ArrayLength:     arrayLength,
		Value:           value,
	}
	_result.VariantContract.(*_Variant)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// VariantStatusCodeBuilder is a builder for VariantStatusCode
type VariantStatusCodeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(value []StatusCode) VariantStatusCodeBuilder
	// WithArrayLength adds ArrayLength (property field)
	WithOptionalArrayLength(int32) VariantStatusCodeBuilder
	// WithValue adds Value (property field)
	WithValue(...StatusCode) VariantStatusCodeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() VariantBuilder
	// Build builds the VariantStatusCode or returns an error if something is wrong
	Build() (VariantStatusCode, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() VariantStatusCode
}

// NewVariantStatusCodeBuilder() creates a VariantStatusCodeBuilder
func NewVariantStatusCodeBuilder() VariantStatusCodeBuilder {
	return &_VariantStatusCodeBuilder{_VariantStatusCode: new(_VariantStatusCode)}
}

type _VariantStatusCodeBuilder struct {
	*_VariantStatusCode

	parentBuilder *_VariantBuilder

	err *utils.MultiError
}

var _ (VariantStatusCodeBuilder) = (*_VariantStatusCodeBuilder)(nil)

func (b *_VariantStatusCodeBuilder) setParent(contract VariantContract) {
	b.VariantContract = contract
	contract.(*_Variant)._SubType = b._VariantStatusCode
}

func (b *_VariantStatusCodeBuilder) WithMandatoryFields(value []StatusCode) VariantStatusCodeBuilder {
	return b.WithValue(value...)
}

func (b *_VariantStatusCodeBuilder) WithOptionalArrayLength(arrayLength int32) VariantStatusCodeBuilder {
	b.ArrayLength = &arrayLength
	return b
}

func (b *_VariantStatusCodeBuilder) WithValue(value ...StatusCode) VariantStatusCodeBuilder {
	b.Value = value
	return b
}

func (b *_VariantStatusCodeBuilder) Build() (VariantStatusCode, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._VariantStatusCode.deepCopy(), nil
}

func (b *_VariantStatusCodeBuilder) MustBuild() VariantStatusCode {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_VariantStatusCodeBuilder) Done() VariantBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewVariantBuilder().(*_VariantBuilder)
	}
	return b.parentBuilder
}

func (b *_VariantStatusCodeBuilder) buildForVariant() (Variant, error) {
	return b.Build()
}

func (b *_VariantStatusCodeBuilder) DeepCopy() any {
	_copy := b.CreateVariantStatusCodeBuilder().(*_VariantStatusCodeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateVariantStatusCodeBuilder creates a VariantStatusCodeBuilder
func (b *_VariantStatusCode) CreateVariantStatusCodeBuilder() VariantStatusCodeBuilder {
	if b == nil {
		return NewVariantStatusCodeBuilder()
	}
	return &_VariantStatusCodeBuilder{_VariantStatusCode: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_VariantStatusCode) GetVariantType() uint8 {
	return uint8(19)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_VariantStatusCode) GetParent() VariantContract {
	return m.VariantContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_VariantStatusCode) GetArrayLength() *int32 {
	return m.ArrayLength
}

func (m *_VariantStatusCode) GetValue() []StatusCode {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastVariantStatusCode(structType any) VariantStatusCode {
	if casted, ok := structType.(VariantStatusCode); ok {
		return casted
	}
	if casted, ok := structType.(*VariantStatusCode); ok {
		return *casted
	}
	return nil
}

func (m *_VariantStatusCode) GetTypeName() string {
	return "VariantStatusCode"
}

func (m *_VariantStatusCode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.VariantContract.(*_Variant).getLengthInBits(ctx))

	// Optional Field (arrayLength)
	if m.ArrayLength != nil {
		lengthInBits += 32
	}

	// Array field
	if len(m.Value) > 0 {
		for _curItem, element := range m.Value {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Value), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_VariantStatusCode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_VariantStatusCode) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_Variant, arrayLengthSpecified bool) (__variantStatusCode VariantStatusCode, err error) {
	m.VariantContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("VariantStatusCode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for VariantStatusCode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var arrayLength *int32
	arrayLength, err = ReadOptionalField[int32](ctx, "arrayLength", ReadSignedInt(readBuffer, uint8(32)), arrayLengthSpecified)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayLength' field"))
	}
	m.ArrayLength = arrayLength

	value, err := ReadCountArrayField[StatusCode](ctx, "value", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer), uint64(utils.InlineIf(bool((arrayLength) == (nil)), func() any { return int32(int32(1)) }, func() any { return int32((*arrayLength)) }).(int32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("VariantStatusCode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for VariantStatusCode")
	}

	return m, nil
}

func (m *_VariantStatusCode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_VariantStatusCode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("VariantStatusCode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for VariantStatusCode")
		}

		if err := WriteOptionalField[int32](ctx, "arrayLength", m.GetArrayLength(), WriteSignedInt(writeBuffer, 32), true); err != nil {
			return errors.Wrap(err, "Error serializing 'arrayLength' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "value", m.GetValue(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("VariantStatusCode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for VariantStatusCode")
		}
		return nil
	}
	return m.VariantContract.(*_Variant).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_VariantStatusCode) IsVariantStatusCode() {}

func (m *_VariantStatusCode) DeepCopy() any {
	return m.deepCopy()
}

func (m *_VariantStatusCode) deepCopy() *_VariantStatusCode {
	if m == nil {
		return nil
	}
	_VariantStatusCodeCopy := &_VariantStatusCode{
		m.VariantContract.(*_Variant).deepCopy(),
		utils.CopyPtr[int32](m.ArrayLength),
		utils.DeepCopySlice[StatusCode, StatusCode](m.Value),
	}
	_VariantStatusCodeCopy.VariantContract.(*_Variant)._SubType = m
	return _VariantStatusCodeCopy
}

func (m *_VariantStatusCode) String() string {
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
