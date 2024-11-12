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

// COTPParameterTpduSize is the corresponding interface of COTPParameterTpduSize
type COTPParameterTpduSize interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	COTPParameter
	// GetTpduSize returns TpduSize (property field)
	GetTpduSize() COTPTpduSize
	// IsCOTPParameterTpduSize is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCOTPParameterTpduSize()
	// CreateBuilder creates a COTPParameterTpduSizeBuilder
	CreateCOTPParameterTpduSizeBuilder() COTPParameterTpduSizeBuilder
}

// _COTPParameterTpduSize is the data-structure of this message
type _COTPParameterTpduSize struct {
	COTPParameterContract
	TpduSize COTPTpduSize
}

var _ COTPParameterTpduSize = (*_COTPParameterTpduSize)(nil)
var _ COTPParameterRequirements = (*_COTPParameterTpduSize)(nil)

// NewCOTPParameterTpduSize factory function for _COTPParameterTpduSize
func NewCOTPParameterTpduSize(tpduSize COTPTpduSize, rest uint8) *_COTPParameterTpduSize {
	_result := &_COTPParameterTpduSize{
		COTPParameterContract: NewCOTPParameter(rest),
		TpduSize:              tpduSize,
	}
	_result.COTPParameterContract.(*_COTPParameter)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// COTPParameterTpduSizeBuilder is a builder for COTPParameterTpduSize
type COTPParameterTpduSizeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(tpduSize COTPTpduSize) COTPParameterTpduSizeBuilder
	// WithTpduSize adds TpduSize (property field)
	WithTpduSize(COTPTpduSize) COTPParameterTpduSizeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() COTPParameterBuilder
	// Build builds the COTPParameterTpduSize or returns an error if something is wrong
	Build() (COTPParameterTpduSize, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() COTPParameterTpduSize
}

// NewCOTPParameterTpduSizeBuilder() creates a COTPParameterTpduSizeBuilder
func NewCOTPParameterTpduSizeBuilder() COTPParameterTpduSizeBuilder {
	return &_COTPParameterTpduSizeBuilder{_COTPParameterTpduSize: new(_COTPParameterTpduSize)}
}

type _COTPParameterTpduSizeBuilder struct {
	*_COTPParameterTpduSize

	parentBuilder *_COTPParameterBuilder

	err *utils.MultiError
}

var _ (COTPParameterTpduSizeBuilder) = (*_COTPParameterTpduSizeBuilder)(nil)

func (b *_COTPParameterTpduSizeBuilder) setParent(contract COTPParameterContract) {
	b.COTPParameterContract = contract
	contract.(*_COTPParameter)._SubType = b._COTPParameterTpduSize
}

func (b *_COTPParameterTpduSizeBuilder) WithMandatoryFields(tpduSize COTPTpduSize) COTPParameterTpduSizeBuilder {
	return b.WithTpduSize(tpduSize)
}

func (b *_COTPParameterTpduSizeBuilder) WithTpduSize(tpduSize COTPTpduSize) COTPParameterTpduSizeBuilder {
	b.TpduSize = tpduSize
	return b
}

func (b *_COTPParameterTpduSizeBuilder) Build() (COTPParameterTpduSize, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._COTPParameterTpduSize.deepCopy(), nil
}

func (b *_COTPParameterTpduSizeBuilder) MustBuild() COTPParameterTpduSize {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_COTPParameterTpduSizeBuilder) Done() COTPParameterBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewCOTPParameterBuilder().(*_COTPParameterBuilder)
	}
	return b.parentBuilder
}

func (b *_COTPParameterTpduSizeBuilder) buildForCOTPParameter() (COTPParameter, error) {
	return b.Build()
}

func (b *_COTPParameterTpduSizeBuilder) DeepCopy() any {
	_copy := b.CreateCOTPParameterTpduSizeBuilder().(*_COTPParameterTpduSizeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCOTPParameterTpduSizeBuilder creates a COTPParameterTpduSizeBuilder
func (b *_COTPParameterTpduSize) CreateCOTPParameterTpduSizeBuilder() COTPParameterTpduSizeBuilder {
	if b == nil {
		return NewCOTPParameterTpduSizeBuilder()
	}
	return &_COTPParameterTpduSizeBuilder{_COTPParameterTpduSize: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_COTPParameterTpduSize) GetParameterType() uint8 {
	return 0xC0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_COTPParameterTpduSize) GetParent() COTPParameterContract {
	return m.COTPParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_COTPParameterTpduSize) GetTpduSize() COTPTpduSize {
	return m.TpduSize
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCOTPParameterTpduSize(structType any) COTPParameterTpduSize {
	if casted, ok := structType.(COTPParameterTpduSize); ok {
		return casted
	}
	if casted, ok := structType.(*COTPParameterTpduSize); ok {
		return *casted
	}
	return nil
}

func (m *_COTPParameterTpduSize) GetTypeName() string {
	return "COTPParameterTpduSize"
}

func (m *_COTPParameterTpduSize) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.COTPParameterContract.(*_COTPParameter).getLengthInBits(ctx))

	// Simple field (tpduSize)
	lengthInBits += 8

	return lengthInBits
}

func (m *_COTPParameterTpduSize) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_COTPParameterTpduSize) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_COTPParameter, rest uint8) (__cOTPParameterTpduSize COTPParameterTpduSize, err error) {
	m.COTPParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("COTPParameterTpduSize"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for COTPParameterTpduSize")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	tpduSize, err := ReadEnumField[COTPTpduSize](ctx, "tpduSize", "COTPTpduSize", ReadEnum(COTPTpduSizeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'tpduSize' field"))
	}
	m.TpduSize = tpduSize

	if closeErr := readBuffer.CloseContext("COTPParameterTpduSize"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for COTPParameterTpduSize")
	}

	return m, nil
}

func (m *_COTPParameterTpduSize) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_COTPParameterTpduSize) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("COTPParameterTpduSize"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for COTPParameterTpduSize")
		}

		if err := WriteSimpleEnumField[COTPTpduSize](ctx, "tpduSize", "COTPTpduSize", m.GetTpduSize(), WriteEnum[COTPTpduSize, uint8](COTPTpduSize.GetValue, COTPTpduSize.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'tpduSize' field")
		}

		if popErr := writeBuffer.PopContext("COTPParameterTpduSize"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for COTPParameterTpduSize")
		}
		return nil
	}
	return m.COTPParameterContract.(*_COTPParameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_COTPParameterTpduSize) IsCOTPParameterTpduSize() {}

func (m *_COTPParameterTpduSize) DeepCopy() any {
	return m.deepCopy()
}

func (m *_COTPParameterTpduSize) deepCopy() *_COTPParameterTpduSize {
	if m == nil {
		return nil
	}
	_COTPParameterTpduSizeCopy := &_COTPParameterTpduSize{
		m.COTPParameterContract.(*_COTPParameter).deepCopy(),
		m.TpduSize,
	}
	_COTPParameterTpduSizeCopy.COTPParameterContract.(*_COTPParameter)._SubType = m
	return _COTPParameterTpduSizeCopy
}

func (m *_COTPParameterTpduSize) String() string {
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
