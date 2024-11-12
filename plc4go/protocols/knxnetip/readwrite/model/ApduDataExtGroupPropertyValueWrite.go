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

// ApduDataExtGroupPropertyValueWrite is the corresponding interface of ApduDataExtGroupPropertyValueWrite
type ApduDataExtGroupPropertyValueWrite interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduDataExt
	// IsApduDataExtGroupPropertyValueWrite is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataExtGroupPropertyValueWrite()
	// CreateBuilder creates a ApduDataExtGroupPropertyValueWriteBuilder
	CreateApduDataExtGroupPropertyValueWriteBuilder() ApduDataExtGroupPropertyValueWriteBuilder
}

// _ApduDataExtGroupPropertyValueWrite is the data-structure of this message
type _ApduDataExtGroupPropertyValueWrite struct {
	ApduDataExtContract
}

var _ ApduDataExtGroupPropertyValueWrite = (*_ApduDataExtGroupPropertyValueWrite)(nil)
var _ ApduDataExtRequirements = (*_ApduDataExtGroupPropertyValueWrite)(nil)

// NewApduDataExtGroupPropertyValueWrite factory function for _ApduDataExtGroupPropertyValueWrite
func NewApduDataExtGroupPropertyValueWrite(length uint8) *_ApduDataExtGroupPropertyValueWrite {
	_result := &_ApduDataExtGroupPropertyValueWrite{
		ApduDataExtContract: NewApduDataExt(length),
	}
	_result.ApduDataExtContract.(*_ApduDataExt)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ApduDataExtGroupPropertyValueWriteBuilder is a builder for ApduDataExtGroupPropertyValueWrite
type ApduDataExtGroupPropertyValueWriteBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ApduDataExtGroupPropertyValueWriteBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ApduDataExtBuilder
	// Build builds the ApduDataExtGroupPropertyValueWrite or returns an error if something is wrong
	Build() (ApduDataExtGroupPropertyValueWrite, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ApduDataExtGroupPropertyValueWrite
}

// NewApduDataExtGroupPropertyValueWriteBuilder() creates a ApduDataExtGroupPropertyValueWriteBuilder
func NewApduDataExtGroupPropertyValueWriteBuilder() ApduDataExtGroupPropertyValueWriteBuilder {
	return &_ApduDataExtGroupPropertyValueWriteBuilder{_ApduDataExtGroupPropertyValueWrite: new(_ApduDataExtGroupPropertyValueWrite)}
}

type _ApduDataExtGroupPropertyValueWriteBuilder struct {
	*_ApduDataExtGroupPropertyValueWrite

	parentBuilder *_ApduDataExtBuilder

	err *utils.MultiError
}

var _ (ApduDataExtGroupPropertyValueWriteBuilder) = (*_ApduDataExtGroupPropertyValueWriteBuilder)(nil)

func (b *_ApduDataExtGroupPropertyValueWriteBuilder) setParent(contract ApduDataExtContract) {
	b.ApduDataExtContract = contract
	contract.(*_ApduDataExt)._SubType = b._ApduDataExtGroupPropertyValueWrite
}

func (b *_ApduDataExtGroupPropertyValueWriteBuilder) WithMandatoryFields() ApduDataExtGroupPropertyValueWriteBuilder {
	return b
}

func (b *_ApduDataExtGroupPropertyValueWriteBuilder) Build() (ApduDataExtGroupPropertyValueWrite, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ApduDataExtGroupPropertyValueWrite.deepCopy(), nil
}

func (b *_ApduDataExtGroupPropertyValueWriteBuilder) MustBuild() ApduDataExtGroupPropertyValueWrite {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ApduDataExtGroupPropertyValueWriteBuilder) Done() ApduDataExtBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewApduDataExtBuilder().(*_ApduDataExtBuilder)
	}
	return b.parentBuilder
}

func (b *_ApduDataExtGroupPropertyValueWriteBuilder) buildForApduDataExt() (ApduDataExt, error) {
	return b.Build()
}

func (b *_ApduDataExtGroupPropertyValueWriteBuilder) DeepCopy() any {
	_copy := b.CreateApduDataExtGroupPropertyValueWriteBuilder().(*_ApduDataExtGroupPropertyValueWriteBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateApduDataExtGroupPropertyValueWriteBuilder creates a ApduDataExtGroupPropertyValueWriteBuilder
func (b *_ApduDataExtGroupPropertyValueWrite) CreateApduDataExtGroupPropertyValueWriteBuilder() ApduDataExtGroupPropertyValueWriteBuilder {
	if b == nil {
		return NewApduDataExtGroupPropertyValueWriteBuilder()
	}
	return &_ApduDataExtGroupPropertyValueWriteBuilder{_ApduDataExtGroupPropertyValueWrite: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtGroupPropertyValueWrite) GetExtApciType() uint8 {
	return 0x2A
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtGroupPropertyValueWrite) GetParent() ApduDataExtContract {
	return m.ApduDataExtContract
}

// Deprecated: use the interface for direct cast
func CastApduDataExtGroupPropertyValueWrite(structType any) ApduDataExtGroupPropertyValueWrite {
	if casted, ok := structType.(ApduDataExtGroupPropertyValueWrite); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtGroupPropertyValueWrite); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtGroupPropertyValueWrite) GetTypeName() string {
	return "ApduDataExtGroupPropertyValueWrite"
}

func (m *_ApduDataExtGroupPropertyValueWrite) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataExtContract.(*_ApduDataExt).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtGroupPropertyValueWrite) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataExtGroupPropertyValueWrite) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduDataExt, length uint8) (__apduDataExtGroupPropertyValueWrite ApduDataExtGroupPropertyValueWrite, err error) {
	m.ApduDataExtContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtGroupPropertyValueWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtGroupPropertyValueWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtGroupPropertyValueWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtGroupPropertyValueWrite")
	}

	return m, nil
}

func (m *_ApduDataExtGroupPropertyValueWrite) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtGroupPropertyValueWrite) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtGroupPropertyValueWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtGroupPropertyValueWrite")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtGroupPropertyValueWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtGroupPropertyValueWrite")
		}
		return nil
	}
	return m.ApduDataExtContract.(*_ApduDataExt).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtGroupPropertyValueWrite) IsApduDataExtGroupPropertyValueWrite() {}

func (m *_ApduDataExtGroupPropertyValueWrite) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataExtGroupPropertyValueWrite) deepCopy() *_ApduDataExtGroupPropertyValueWrite {
	if m == nil {
		return nil
	}
	_ApduDataExtGroupPropertyValueWriteCopy := &_ApduDataExtGroupPropertyValueWrite{
		m.ApduDataExtContract.(*_ApduDataExt).deepCopy(),
	}
	_ApduDataExtGroupPropertyValueWriteCopy.ApduDataExtContract.(*_ApduDataExt)._SubType = m
	return _ApduDataExtGroupPropertyValueWriteCopy
}

func (m *_ApduDataExtGroupPropertyValueWrite) String() string {
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
