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

// APDUSegmentAck is the corresponding interface of APDUSegmentAck
type APDUSegmentAck interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	APDU
	// GetNegativeAck returns NegativeAck (property field)
	GetNegativeAck() bool
	// GetServer returns Server (property field)
	GetServer() bool
	// GetOriginalInvokeId returns OriginalInvokeId (property field)
	GetOriginalInvokeId() uint8
	// GetSequenceNumber returns SequenceNumber (property field)
	GetSequenceNumber() uint8
	// GetActualWindowSize returns ActualWindowSize (property field)
	GetActualWindowSize() uint8
	// IsAPDUSegmentAck is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAPDUSegmentAck()
	// CreateBuilder creates a APDUSegmentAckBuilder
	CreateAPDUSegmentAckBuilder() APDUSegmentAckBuilder
}

// _APDUSegmentAck is the data-structure of this message
type _APDUSegmentAck struct {
	APDUContract
	NegativeAck      bool
	Server           bool
	OriginalInvokeId uint8
	SequenceNumber   uint8
	ActualWindowSize uint8
	// Reserved Fields
	reservedField0 *uint8
}

var _ APDUSegmentAck = (*_APDUSegmentAck)(nil)
var _ APDURequirements = (*_APDUSegmentAck)(nil)

// NewAPDUSegmentAck factory function for _APDUSegmentAck
func NewAPDUSegmentAck(negativeAck bool, server bool, originalInvokeId uint8, sequenceNumber uint8, actualWindowSize uint8, apduLength uint16) *_APDUSegmentAck {
	_result := &_APDUSegmentAck{
		APDUContract:     NewAPDU(apduLength),
		NegativeAck:      negativeAck,
		Server:           server,
		OriginalInvokeId: originalInvokeId,
		SequenceNumber:   sequenceNumber,
		ActualWindowSize: actualWindowSize,
	}
	_result.APDUContract.(*_APDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// APDUSegmentAckBuilder is a builder for APDUSegmentAck
type APDUSegmentAckBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(negativeAck bool, server bool, originalInvokeId uint8, sequenceNumber uint8, actualWindowSize uint8) APDUSegmentAckBuilder
	// WithNegativeAck adds NegativeAck (property field)
	WithNegativeAck(bool) APDUSegmentAckBuilder
	// WithServer adds Server (property field)
	WithServer(bool) APDUSegmentAckBuilder
	// WithOriginalInvokeId adds OriginalInvokeId (property field)
	WithOriginalInvokeId(uint8) APDUSegmentAckBuilder
	// WithSequenceNumber adds SequenceNumber (property field)
	WithSequenceNumber(uint8) APDUSegmentAckBuilder
	// WithActualWindowSize adds ActualWindowSize (property field)
	WithActualWindowSize(uint8) APDUSegmentAckBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() APDUBuilder
	// Build builds the APDUSegmentAck or returns an error if something is wrong
	Build() (APDUSegmentAck, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() APDUSegmentAck
}

// NewAPDUSegmentAckBuilder() creates a APDUSegmentAckBuilder
func NewAPDUSegmentAckBuilder() APDUSegmentAckBuilder {
	return &_APDUSegmentAckBuilder{_APDUSegmentAck: new(_APDUSegmentAck)}
}

type _APDUSegmentAckBuilder struct {
	*_APDUSegmentAck

	parentBuilder *_APDUBuilder

	err *utils.MultiError
}

var _ (APDUSegmentAckBuilder) = (*_APDUSegmentAckBuilder)(nil)

func (b *_APDUSegmentAckBuilder) setParent(contract APDUContract) {
	b.APDUContract = contract
	contract.(*_APDU)._SubType = b._APDUSegmentAck
}

func (b *_APDUSegmentAckBuilder) WithMandatoryFields(negativeAck bool, server bool, originalInvokeId uint8, sequenceNumber uint8, actualWindowSize uint8) APDUSegmentAckBuilder {
	return b.WithNegativeAck(negativeAck).WithServer(server).WithOriginalInvokeId(originalInvokeId).WithSequenceNumber(sequenceNumber).WithActualWindowSize(actualWindowSize)
}

func (b *_APDUSegmentAckBuilder) WithNegativeAck(negativeAck bool) APDUSegmentAckBuilder {
	b.NegativeAck = negativeAck
	return b
}

func (b *_APDUSegmentAckBuilder) WithServer(server bool) APDUSegmentAckBuilder {
	b.Server = server
	return b
}

func (b *_APDUSegmentAckBuilder) WithOriginalInvokeId(originalInvokeId uint8) APDUSegmentAckBuilder {
	b.OriginalInvokeId = originalInvokeId
	return b
}

func (b *_APDUSegmentAckBuilder) WithSequenceNumber(sequenceNumber uint8) APDUSegmentAckBuilder {
	b.SequenceNumber = sequenceNumber
	return b
}

func (b *_APDUSegmentAckBuilder) WithActualWindowSize(actualWindowSize uint8) APDUSegmentAckBuilder {
	b.ActualWindowSize = actualWindowSize
	return b
}

func (b *_APDUSegmentAckBuilder) Build() (APDUSegmentAck, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._APDUSegmentAck.deepCopy(), nil
}

func (b *_APDUSegmentAckBuilder) MustBuild() APDUSegmentAck {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_APDUSegmentAckBuilder) Done() APDUBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewAPDUBuilder().(*_APDUBuilder)
	}
	return b.parentBuilder
}

func (b *_APDUSegmentAckBuilder) buildForAPDU() (APDU, error) {
	return b.Build()
}

func (b *_APDUSegmentAckBuilder) DeepCopy() any {
	_copy := b.CreateAPDUSegmentAckBuilder().(*_APDUSegmentAckBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAPDUSegmentAckBuilder creates a APDUSegmentAckBuilder
func (b *_APDUSegmentAck) CreateAPDUSegmentAckBuilder() APDUSegmentAckBuilder {
	if b == nil {
		return NewAPDUSegmentAckBuilder()
	}
	return &_APDUSegmentAckBuilder{_APDUSegmentAck: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_APDUSegmentAck) GetApduType() ApduType {
	return ApduType_SEGMENT_ACK_PDU
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_APDUSegmentAck) GetParent() APDUContract {
	return m.APDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_APDUSegmentAck) GetNegativeAck() bool {
	return m.NegativeAck
}

func (m *_APDUSegmentAck) GetServer() bool {
	return m.Server
}

func (m *_APDUSegmentAck) GetOriginalInvokeId() uint8 {
	return m.OriginalInvokeId
}

func (m *_APDUSegmentAck) GetSequenceNumber() uint8 {
	return m.SequenceNumber
}

func (m *_APDUSegmentAck) GetActualWindowSize() uint8 {
	return m.ActualWindowSize
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAPDUSegmentAck(structType any) APDUSegmentAck {
	if casted, ok := structType.(APDUSegmentAck); ok {
		return casted
	}
	if casted, ok := structType.(*APDUSegmentAck); ok {
		return *casted
	}
	return nil
}

func (m *_APDUSegmentAck) GetTypeName() string {
	return "APDUSegmentAck"
}

func (m *_APDUSegmentAck) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.APDUContract.(*_APDU).getLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 2

	// Simple field (negativeAck)
	lengthInBits += 1

	// Simple field (server)
	lengthInBits += 1

	// Simple field (originalInvokeId)
	lengthInBits += 8

	// Simple field (sequenceNumber)
	lengthInBits += 8

	// Simple field (actualWindowSize)
	lengthInBits += 8

	return lengthInBits
}

func (m *_APDUSegmentAck) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_APDUSegmentAck) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_APDU, apduLength uint16) (__aPDUSegmentAck APDUSegmentAck, err error) {
	m.APDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("APDUSegmentAck"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for APDUSegmentAck")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(2)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	negativeAck, err := ReadSimpleField(ctx, "negativeAck", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'negativeAck' field"))
	}
	m.NegativeAck = negativeAck

	server, err := ReadSimpleField(ctx, "server", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'server' field"))
	}
	m.Server = server

	originalInvokeId, err := ReadSimpleField(ctx, "originalInvokeId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'originalInvokeId' field"))
	}
	m.OriginalInvokeId = originalInvokeId

	sequenceNumber, err := ReadSimpleField(ctx, "sequenceNumber", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sequenceNumber' field"))
	}
	m.SequenceNumber = sequenceNumber

	actualWindowSize, err := ReadSimpleField(ctx, "actualWindowSize", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualWindowSize' field"))
	}
	m.ActualWindowSize = actualWindowSize

	if closeErr := readBuffer.CloseContext("APDUSegmentAck"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for APDUSegmentAck")
	}

	return m, nil
}

func (m *_APDUSegmentAck) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_APDUSegmentAck) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUSegmentAck"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for APDUSegmentAck")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 2)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "negativeAck", m.GetNegativeAck(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'negativeAck' field")
		}

		if err := WriteSimpleField[bool](ctx, "server", m.GetServer(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'server' field")
		}

		if err := WriteSimpleField[uint8](ctx, "originalInvokeId", m.GetOriginalInvokeId(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'originalInvokeId' field")
		}

		if err := WriteSimpleField[uint8](ctx, "sequenceNumber", m.GetSequenceNumber(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'sequenceNumber' field")
		}

		if err := WriteSimpleField[uint8](ctx, "actualWindowSize", m.GetActualWindowSize(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'actualWindowSize' field")
		}

		if popErr := writeBuffer.PopContext("APDUSegmentAck"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for APDUSegmentAck")
		}
		return nil
	}
	return m.APDUContract.(*_APDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_APDUSegmentAck) IsAPDUSegmentAck() {}

func (m *_APDUSegmentAck) DeepCopy() any {
	return m.deepCopy()
}

func (m *_APDUSegmentAck) deepCopy() *_APDUSegmentAck {
	if m == nil {
		return nil
	}
	_APDUSegmentAckCopy := &_APDUSegmentAck{
		m.APDUContract.(*_APDU).deepCopy(),
		m.NegativeAck,
		m.Server,
		m.OriginalInvokeId,
		m.SequenceNumber,
		m.ActualWindowSize,
		m.reservedField0,
	}
	_APDUSegmentAckCopy.APDUContract.(*_APDU)._SubType = m
	return _APDUSegmentAckCopy
}

func (m *_APDUSegmentAck) String() string {
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
