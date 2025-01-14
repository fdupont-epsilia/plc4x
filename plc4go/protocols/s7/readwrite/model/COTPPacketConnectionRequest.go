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

// COTPPacketConnectionRequest is the corresponding interface of COTPPacketConnectionRequest
type COTPPacketConnectionRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	COTPPacket
	// GetDestinationReference returns DestinationReference (property field)
	GetDestinationReference() uint16
	// GetSourceReference returns SourceReference (property field)
	GetSourceReference() uint16
	// GetProtocolClass returns ProtocolClass (property field)
	GetProtocolClass() COTPProtocolClass
	// IsCOTPPacketConnectionRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCOTPPacketConnectionRequest()
	// CreateBuilder creates a COTPPacketConnectionRequestBuilder
	CreateCOTPPacketConnectionRequestBuilder() COTPPacketConnectionRequestBuilder
}

// _COTPPacketConnectionRequest is the data-structure of this message
type _COTPPacketConnectionRequest struct {
	COTPPacketContract
	DestinationReference uint16
	SourceReference      uint16
	ProtocolClass        COTPProtocolClass
}

var _ COTPPacketConnectionRequest = (*_COTPPacketConnectionRequest)(nil)
var _ COTPPacketRequirements = (*_COTPPacketConnectionRequest)(nil)

// NewCOTPPacketConnectionRequest factory function for _COTPPacketConnectionRequest
func NewCOTPPacketConnectionRequest(parameters []COTPParameter, payload S7Message, destinationReference uint16, sourceReference uint16, protocolClass COTPProtocolClass, cotpLen uint16) *_COTPPacketConnectionRequest {
	_result := &_COTPPacketConnectionRequest{
		COTPPacketContract:   NewCOTPPacket(parameters, payload, cotpLen),
		DestinationReference: destinationReference,
		SourceReference:      sourceReference,
		ProtocolClass:        protocolClass,
	}
	_result.COTPPacketContract.(*_COTPPacket)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// COTPPacketConnectionRequestBuilder is a builder for COTPPacketConnectionRequest
type COTPPacketConnectionRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(destinationReference uint16, sourceReference uint16, protocolClass COTPProtocolClass) COTPPacketConnectionRequestBuilder
	// WithDestinationReference adds DestinationReference (property field)
	WithDestinationReference(uint16) COTPPacketConnectionRequestBuilder
	// WithSourceReference adds SourceReference (property field)
	WithSourceReference(uint16) COTPPacketConnectionRequestBuilder
	// WithProtocolClass adds ProtocolClass (property field)
	WithProtocolClass(COTPProtocolClass) COTPPacketConnectionRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() COTPPacketBuilder
	// Build builds the COTPPacketConnectionRequest or returns an error if something is wrong
	Build() (COTPPacketConnectionRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() COTPPacketConnectionRequest
}

// NewCOTPPacketConnectionRequestBuilder() creates a COTPPacketConnectionRequestBuilder
func NewCOTPPacketConnectionRequestBuilder() COTPPacketConnectionRequestBuilder {
	return &_COTPPacketConnectionRequestBuilder{_COTPPacketConnectionRequest: new(_COTPPacketConnectionRequest)}
}

type _COTPPacketConnectionRequestBuilder struct {
	*_COTPPacketConnectionRequest

	parentBuilder *_COTPPacketBuilder

	err *utils.MultiError
}

var _ (COTPPacketConnectionRequestBuilder) = (*_COTPPacketConnectionRequestBuilder)(nil)

func (b *_COTPPacketConnectionRequestBuilder) setParent(contract COTPPacketContract) {
	b.COTPPacketContract = contract
	contract.(*_COTPPacket)._SubType = b._COTPPacketConnectionRequest
}

func (b *_COTPPacketConnectionRequestBuilder) WithMandatoryFields(destinationReference uint16, sourceReference uint16, protocolClass COTPProtocolClass) COTPPacketConnectionRequestBuilder {
	return b.WithDestinationReference(destinationReference).WithSourceReference(sourceReference).WithProtocolClass(protocolClass)
}

func (b *_COTPPacketConnectionRequestBuilder) WithDestinationReference(destinationReference uint16) COTPPacketConnectionRequestBuilder {
	b.DestinationReference = destinationReference
	return b
}

func (b *_COTPPacketConnectionRequestBuilder) WithSourceReference(sourceReference uint16) COTPPacketConnectionRequestBuilder {
	b.SourceReference = sourceReference
	return b
}

func (b *_COTPPacketConnectionRequestBuilder) WithProtocolClass(protocolClass COTPProtocolClass) COTPPacketConnectionRequestBuilder {
	b.ProtocolClass = protocolClass
	return b
}

func (b *_COTPPacketConnectionRequestBuilder) Build() (COTPPacketConnectionRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._COTPPacketConnectionRequest.deepCopy(), nil
}

func (b *_COTPPacketConnectionRequestBuilder) MustBuild() COTPPacketConnectionRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_COTPPacketConnectionRequestBuilder) Done() COTPPacketBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewCOTPPacketBuilder().(*_COTPPacketBuilder)
	}
	return b.parentBuilder
}

func (b *_COTPPacketConnectionRequestBuilder) buildForCOTPPacket() (COTPPacket, error) {
	return b.Build()
}

func (b *_COTPPacketConnectionRequestBuilder) DeepCopy() any {
	_copy := b.CreateCOTPPacketConnectionRequestBuilder().(*_COTPPacketConnectionRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCOTPPacketConnectionRequestBuilder creates a COTPPacketConnectionRequestBuilder
func (b *_COTPPacketConnectionRequest) CreateCOTPPacketConnectionRequestBuilder() COTPPacketConnectionRequestBuilder {
	if b == nil {
		return NewCOTPPacketConnectionRequestBuilder()
	}
	return &_COTPPacketConnectionRequestBuilder{_COTPPacketConnectionRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_COTPPacketConnectionRequest) GetTpduCode() uint8 {
	return 0xE0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_COTPPacketConnectionRequest) GetParent() COTPPacketContract {
	return m.COTPPacketContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_COTPPacketConnectionRequest) GetDestinationReference() uint16 {
	return m.DestinationReference
}

func (m *_COTPPacketConnectionRequest) GetSourceReference() uint16 {
	return m.SourceReference
}

func (m *_COTPPacketConnectionRequest) GetProtocolClass() COTPProtocolClass {
	return m.ProtocolClass
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCOTPPacketConnectionRequest(structType any) COTPPacketConnectionRequest {
	if casted, ok := structType.(COTPPacketConnectionRequest); ok {
		return casted
	}
	if casted, ok := structType.(*COTPPacketConnectionRequest); ok {
		return *casted
	}
	return nil
}

func (m *_COTPPacketConnectionRequest) GetTypeName() string {
	return "COTPPacketConnectionRequest"
}

func (m *_COTPPacketConnectionRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.COTPPacketContract.(*_COTPPacket).getLengthInBits(ctx))

	// Simple field (destinationReference)
	lengthInBits += 16

	// Simple field (sourceReference)
	lengthInBits += 16

	// Simple field (protocolClass)
	lengthInBits += 8

	return lengthInBits
}

func (m *_COTPPacketConnectionRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_COTPPacketConnectionRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_COTPPacket, cotpLen uint16) (__cOTPPacketConnectionRequest COTPPacketConnectionRequest, err error) {
	m.COTPPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("COTPPacketConnectionRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for COTPPacketConnectionRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	destinationReference, err := ReadSimpleField(ctx, "destinationReference", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'destinationReference' field"))
	}
	m.DestinationReference = destinationReference

	sourceReference, err := ReadSimpleField(ctx, "sourceReference", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sourceReference' field"))
	}
	m.SourceReference = sourceReference

	protocolClass, err := ReadEnumField[COTPProtocolClass](ctx, "protocolClass", "COTPProtocolClass", ReadEnum(COTPProtocolClassByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'protocolClass' field"))
	}
	m.ProtocolClass = protocolClass

	if closeErr := readBuffer.CloseContext("COTPPacketConnectionRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for COTPPacketConnectionRequest")
	}

	return m, nil
}

func (m *_COTPPacketConnectionRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_COTPPacketConnectionRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("COTPPacketConnectionRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for COTPPacketConnectionRequest")
		}

		if err := WriteSimpleField[uint16](ctx, "destinationReference", m.GetDestinationReference(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'destinationReference' field")
		}

		if err := WriteSimpleField[uint16](ctx, "sourceReference", m.GetSourceReference(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'sourceReference' field")
		}

		if err := WriteSimpleEnumField[COTPProtocolClass](ctx, "protocolClass", "COTPProtocolClass", m.GetProtocolClass(), WriteEnum[COTPProtocolClass, uint8](COTPProtocolClass.GetValue, COTPProtocolClass.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'protocolClass' field")
		}

		if popErr := writeBuffer.PopContext("COTPPacketConnectionRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for COTPPacketConnectionRequest")
		}
		return nil
	}
	return m.COTPPacketContract.(*_COTPPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_COTPPacketConnectionRequest) IsCOTPPacketConnectionRequest() {}

func (m *_COTPPacketConnectionRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_COTPPacketConnectionRequest) deepCopy() *_COTPPacketConnectionRequest {
	if m == nil {
		return nil
	}
	_COTPPacketConnectionRequestCopy := &_COTPPacketConnectionRequest{
		m.COTPPacketContract.(*_COTPPacket).deepCopy(),
		m.DestinationReference,
		m.SourceReference,
		m.ProtocolClass,
	}
	_COTPPacketConnectionRequestCopy.COTPPacketContract.(*_COTPPacket)._SubType = m
	return _COTPPacketConnectionRequestCopy
}

func (m *_COTPPacketConnectionRequest) String() string {
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
