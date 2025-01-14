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

// NLMReserved is the corresponding interface of NLMReserved
type NLMReserved interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	NLM
	// GetUnknownBytes returns UnknownBytes (property field)
	GetUnknownBytes() []byte
	// IsNLMReserved is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNLMReserved()
	// CreateBuilder creates a NLMReservedBuilder
	CreateNLMReservedBuilder() NLMReservedBuilder
}

// _NLMReserved is the data-structure of this message
type _NLMReserved struct {
	NLMContract
	UnknownBytes []byte
}

var _ NLMReserved = (*_NLMReserved)(nil)
var _ NLMRequirements = (*_NLMReserved)(nil)

// NewNLMReserved factory function for _NLMReserved
func NewNLMReserved(unknownBytes []byte, apduLength uint16) *_NLMReserved {
	_result := &_NLMReserved{
		NLMContract:  NewNLM(apduLength),
		UnknownBytes: unknownBytes,
	}
	_result.NLMContract.(*_NLM)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// NLMReservedBuilder is a builder for NLMReserved
type NLMReservedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(unknownBytes []byte) NLMReservedBuilder
	// WithUnknownBytes adds UnknownBytes (property field)
	WithUnknownBytes(...byte) NLMReservedBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() NLMBuilder
	// Build builds the NLMReserved or returns an error if something is wrong
	Build() (NLMReserved, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() NLMReserved
}

// NewNLMReservedBuilder() creates a NLMReservedBuilder
func NewNLMReservedBuilder() NLMReservedBuilder {
	return &_NLMReservedBuilder{_NLMReserved: new(_NLMReserved)}
}

type _NLMReservedBuilder struct {
	*_NLMReserved

	parentBuilder *_NLMBuilder

	err *utils.MultiError
}

var _ (NLMReservedBuilder) = (*_NLMReservedBuilder)(nil)

func (b *_NLMReservedBuilder) setParent(contract NLMContract) {
	b.NLMContract = contract
	contract.(*_NLM)._SubType = b._NLMReserved
}

func (b *_NLMReservedBuilder) WithMandatoryFields(unknownBytes []byte) NLMReservedBuilder {
	return b.WithUnknownBytes(unknownBytes...)
}

func (b *_NLMReservedBuilder) WithUnknownBytes(unknownBytes ...byte) NLMReservedBuilder {
	b.UnknownBytes = unknownBytes
	return b
}

func (b *_NLMReservedBuilder) Build() (NLMReserved, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._NLMReserved.deepCopy(), nil
}

func (b *_NLMReservedBuilder) MustBuild() NLMReserved {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_NLMReservedBuilder) Done() NLMBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewNLMBuilder().(*_NLMBuilder)
	}
	return b.parentBuilder
}

func (b *_NLMReservedBuilder) buildForNLM() (NLM, error) {
	return b.Build()
}

func (b *_NLMReservedBuilder) DeepCopy() any {
	_copy := b.CreateNLMReservedBuilder().(*_NLMReservedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateNLMReservedBuilder creates a NLMReservedBuilder
func (b *_NLMReserved) CreateNLMReservedBuilder() NLMReservedBuilder {
	if b == nil {
		return NewNLMReservedBuilder()
	}
	return &_NLMReservedBuilder{_NLMReserved: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMReserved) GetMessageType() uint8 {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMReserved) GetParent() NLMContract {
	return m.NLMContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMReserved) GetUnknownBytes() []byte {
	return m.UnknownBytes
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastNLMReserved(structType any) NLMReserved {
	if casted, ok := structType.(NLMReserved); ok {
		return casted
	}
	if casted, ok := structType.(*NLMReserved); ok {
		return *casted
	}
	return nil
}

func (m *_NLMReserved) GetTypeName() string {
	return "NLMReserved"
}

func (m *_NLMReserved) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.NLMContract.(*_NLM).getLengthInBits(ctx))

	// Array field
	if len(m.UnknownBytes) > 0 {
		lengthInBits += 8 * uint16(len(m.UnknownBytes))
	}

	return lengthInBits
}

func (m *_NLMReserved) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NLMReserved) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_NLM, apduLength uint16) (__nLMReserved NLMReserved, err error) {
	m.NLMContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMReserved"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMReserved")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	unknownBytes, err := readBuffer.ReadByteArray("unknownBytes", int(utils.InlineIf((bool((apduLength) > (0))), func() any { return int32((int32(apduLength) - int32(int32(1)))) }, func() any { return int32(int32(0)) }).(int32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unknownBytes' field"))
	}
	m.UnknownBytes = unknownBytes

	if closeErr := readBuffer.CloseContext("NLMReserved"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMReserved")
	}

	return m, nil
}

func (m *_NLMReserved) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMReserved) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMReserved"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMReserved")
		}

		if err := WriteByteArrayField(ctx, "unknownBytes", m.GetUnknownBytes(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'unknownBytes' field")
		}

		if popErr := writeBuffer.PopContext("NLMReserved"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMReserved")
		}
		return nil
	}
	return m.NLMContract.(*_NLM).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NLMReserved) IsNLMReserved() {}

func (m *_NLMReserved) DeepCopy() any {
	return m.deepCopy()
}

func (m *_NLMReserved) deepCopy() *_NLMReserved {
	if m == nil {
		return nil
	}
	_NLMReservedCopy := &_NLMReserved{
		m.NLMContract.(*_NLM).deepCopy(),
		utils.DeepCopySlice[byte, byte](m.UnknownBytes),
	}
	_NLMReservedCopy.NLMContract.(*_NLM)._SubType = m
	return _NLMReservedCopy
}

func (m *_NLMReserved) String() string {
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
