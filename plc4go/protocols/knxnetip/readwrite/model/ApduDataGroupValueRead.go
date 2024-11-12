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

// ApduDataGroupValueRead is the corresponding interface of ApduDataGroupValueRead
type ApduDataGroupValueRead interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduData
	// IsApduDataGroupValueRead is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataGroupValueRead()
	// CreateBuilder creates a ApduDataGroupValueReadBuilder
	CreateApduDataGroupValueReadBuilder() ApduDataGroupValueReadBuilder
}

// _ApduDataGroupValueRead is the data-structure of this message
type _ApduDataGroupValueRead struct {
	ApduDataContract
	// Reserved Fields
	reservedField0 *uint8
}

var _ ApduDataGroupValueRead = (*_ApduDataGroupValueRead)(nil)
var _ ApduDataRequirements = (*_ApduDataGroupValueRead)(nil)

// NewApduDataGroupValueRead factory function for _ApduDataGroupValueRead
func NewApduDataGroupValueRead(dataLength uint8) *_ApduDataGroupValueRead {
	_result := &_ApduDataGroupValueRead{
		ApduDataContract: NewApduData(dataLength),
	}
	_result.ApduDataContract.(*_ApduData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ApduDataGroupValueReadBuilder is a builder for ApduDataGroupValueRead
type ApduDataGroupValueReadBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ApduDataGroupValueReadBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ApduDataBuilder
	// Build builds the ApduDataGroupValueRead or returns an error if something is wrong
	Build() (ApduDataGroupValueRead, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ApduDataGroupValueRead
}

// NewApduDataGroupValueReadBuilder() creates a ApduDataGroupValueReadBuilder
func NewApduDataGroupValueReadBuilder() ApduDataGroupValueReadBuilder {
	return &_ApduDataGroupValueReadBuilder{_ApduDataGroupValueRead: new(_ApduDataGroupValueRead)}
}

type _ApduDataGroupValueReadBuilder struct {
	*_ApduDataGroupValueRead

	parentBuilder *_ApduDataBuilder

	err *utils.MultiError
}

var _ (ApduDataGroupValueReadBuilder) = (*_ApduDataGroupValueReadBuilder)(nil)

func (b *_ApduDataGroupValueReadBuilder) setParent(contract ApduDataContract) {
	b.ApduDataContract = contract
	contract.(*_ApduData)._SubType = b._ApduDataGroupValueRead
}

func (b *_ApduDataGroupValueReadBuilder) WithMandatoryFields() ApduDataGroupValueReadBuilder {
	return b
}

func (b *_ApduDataGroupValueReadBuilder) Build() (ApduDataGroupValueRead, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ApduDataGroupValueRead.deepCopy(), nil
}

func (b *_ApduDataGroupValueReadBuilder) MustBuild() ApduDataGroupValueRead {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ApduDataGroupValueReadBuilder) Done() ApduDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewApduDataBuilder().(*_ApduDataBuilder)
	}
	return b.parentBuilder
}

func (b *_ApduDataGroupValueReadBuilder) buildForApduData() (ApduData, error) {
	return b.Build()
}

func (b *_ApduDataGroupValueReadBuilder) DeepCopy() any {
	_copy := b.CreateApduDataGroupValueReadBuilder().(*_ApduDataGroupValueReadBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateApduDataGroupValueReadBuilder creates a ApduDataGroupValueReadBuilder
func (b *_ApduDataGroupValueRead) CreateApduDataGroupValueReadBuilder() ApduDataGroupValueReadBuilder {
	if b == nil {
		return NewApduDataGroupValueReadBuilder()
	}
	return &_ApduDataGroupValueReadBuilder{_ApduDataGroupValueRead: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataGroupValueRead) GetApciType() uint8 {
	return 0x0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataGroupValueRead) GetParent() ApduDataContract {
	return m.ApduDataContract
}

// Deprecated: use the interface for direct cast
func CastApduDataGroupValueRead(structType any) ApduDataGroupValueRead {
	if casted, ok := structType.(ApduDataGroupValueRead); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataGroupValueRead); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataGroupValueRead) GetTypeName() string {
	return "ApduDataGroupValueRead"
}

func (m *_ApduDataGroupValueRead) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataContract.(*_ApduData).getLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 6

	return lengthInBits
}

func (m *_ApduDataGroupValueRead) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataGroupValueRead) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduData, dataLength uint8) (__apduDataGroupValueRead ApduDataGroupValueRead, err error) {
	m.ApduDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataGroupValueRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataGroupValueRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(6)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	if closeErr := readBuffer.CloseContext("ApduDataGroupValueRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataGroupValueRead")
	}

	return m, nil
}

func (m *_ApduDataGroupValueRead) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataGroupValueRead) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataGroupValueRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataGroupValueRead")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 6)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if popErr := writeBuffer.PopContext("ApduDataGroupValueRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataGroupValueRead")
		}
		return nil
	}
	return m.ApduDataContract.(*_ApduData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataGroupValueRead) IsApduDataGroupValueRead() {}

func (m *_ApduDataGroupValueRead) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataGroupValueRead) deepCopy() *_ApduDataGroupValueRead {
	if m == nil {
		return nil
	}
	_ApduDataGroupValueReadCopy := &_ApduDataGroupValueRead{
		m.ApduDataContract.(*_ApduData).deepCopy(),
		m.reservedField0,
	}
	_ApduDataGroupValueReadCopy.ApduDataContract.(*_ApduData)._SubType = m
	return _ApduDataGroupValueReadCopy
}

func (m *_ApduDataGroupValueRead) String() string {
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
