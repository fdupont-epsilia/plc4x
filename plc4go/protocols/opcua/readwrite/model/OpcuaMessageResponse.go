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

// OpcuaMessageResponse is the corresponding interface of OpcuaMessageResponse
type OpcuaMessageResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	MessagePDU
	// GetSecurityHeader returns SecurityHeader (property field)
	GetSecurityHeader() SecurityHeader
	// GetMessage returns Message (property field)
	GetMessage() Payload
	// IsOpcuaMessageResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsOpcuaMessageResponse()
	// CreateBuilder creates a OpcuaMessageResponseBuilder
	CreateOpcuaMessageResponseBuilder() OpcuaMessageResponseBuilder
}

// _OpcuaMessageResponse is the data-structure of this message
type _OpcuaMessageResponse struct {
	MessagePDUContract
	SecurityHeader SecurityHeader
	Message        Payload

	// Arguments.
	TotalLength uint32
}

var _ OpcuaMessageResponse = (*_OpcuaMessageResponse)(nil)
var _ MessagePDURequirements = (*_OpcuaMessageResponse)(nil)

// NewOpcuaMessageResponse factory function for _OpcuaMessageResponse
func NewOpcuaMessageResponse(chunk ChunkType, securityHeader SecurityHeader, message Payload, totalLength uint32, binary bool) *_OpcuaMessageResponse {
	if securityHeader == nil {
		panic("securityHeader of type SecurityHeader for OpcuaMessageResponse must not be nil")
	}
	if message == nil {
		panic("message of type Payload for OpcuaMessageResponse must not be nil")
	}
	_result := &_OpcuaMessageResponse{
		MessagePDUContract: NewMessagePDU(chunk, binary),
		SecurityHeader:     securityHeader,
		Message:            message,
	}
	_result.MessagePDUContract.(*_MessagePDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// OpcuaMessageResponseBuilder is a builder for OpcuaMessageResponse
type OpcuaMessageResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(securityHeader SecurityHeader, message Payload) OpcuaMessageResponseBuilder
	// WithSecurityHeader adds SecurityHeader (property field)
	WithSecurityHeader(SecurityHeader) OpcuaMessageResponseBuilder
	// WithSecurityHeaderBuilder adds SecurityHeader (property field) which is build by the builder
	WithSecurityHeaderBuilder(func(SecurityHeaderBuilder) SecurityHeaderBuilder) OpcuaMessageResponseBuilder
	// WithMessage adds Message (property field)
	WithMessage(Payload) OpcuaMessageResponseBuilder
	// WithMessageBuilder adds Message (property field) which is build by the builder
	WithMessageBuilder(func(PayloadBuilder) PayloadBuilder) OpcuaMessageResponseBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() MessagePDUBuilder
	// Build builds the OpcuaMessageResponse or returns an error if something is wrong
	Build() (OpcuaMessageResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() OpcuaMessageResponse
}

// NewOpcuaMessageResponseBuilder() creates a OpcuaMessageResponseBuilder
func NewOpcuaMessageResponseBuilder() OpcuaMessageResponseBuilder {
	return &_OpcuaMessageResponseBuilder{_OpcuaMessageResponse: new(_OpcuaMessageResponse)}
}

type _OpcuaMessageResponseBuilder struct {
	*_OpcuaMessageResponse

	parentBuilder *_MessagePDUBuilder

	err *utils.MultiError
}

var _ (OpcuaMessageResponseBuilder) = (*_OpcuaMessageResponseBuilder)(nil)

func (b *_OpcuaMessageResponseBuilder) setParent(contract MessagePDUContract) {
	b.MessagePDUContract = contract
	contract.(*_MessagePDU)._SubType = b._OpcuaMessageResponse
}

func (b *_OpcuaMessageResponseBuilder) WithMandatoryFields(securityHeader SecurityHeader, message Payload) OpcuaMessageResponseBuilder {
	return b.WithSecurityHeader(securityHeader).WithMessage(message)
}

func (b *_OpcuaMessageResponseBuilder) WithSecurityHeader(securityHeader SecurityHeader) OpcuaMessageResponseBuilder {
	b.SecurityHeader = securityHeader
	return b
}

func (b *_OpcuaMessageResponseBuilder) WithSecurityHeaderBuilder(builderSupplier func(SecurityHeaderBuilder) SecurityHeaderBuilder) OpcuaMessageResponseBuilder {
	builder := builderSupplier(b.SecurityHeader.CreateSecurityHeaderBuilder())
	var err error
	b.SecurityHeader, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "SecurityHeaderBuilder failed"))
	}
	return b
}

func (b *_OpcuaMessageResponseBuilder) WithMessage(message Payload) OpcuaMessageResponseBuilder {
	b.Message = message
	return b
}

func (b *_OpcuaMessageResponseBuilder) WithMessageBuilder(builderSupplier func(PayloadBuilder) PayloadBuilder) OpcuaMessageResponseBuilder {
	builder := builderSupplier(b.Message.CreatePayloadBuilder())
	var err error
	b.Message, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PayloadBuilder failed"))
	}
	return b
}

func (b *_OpcuaMessageResponseBuilder) Build() (OpcuaMessageResponse, error) {
	if b.SecurityHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'securityHeader' not set"))
	}
	if b.Message == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'message' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._OpcuaMessageResponse.deepCopy(), nil
}

func (b *_OpcuaMessageResponseBuilder) MustBuild() OpcuaMessageResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_OpcuaMessageResponseBuilder) Done() MessagePDUBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewMessagePDUBuilder().(*_MessagePDUBuilder)
	}
	return b.parentBuilder
}

func (b *_OpcuaMessageResponseBuilder) buildForMessagePDU() (MessagePDU, error) {
	return b.Build()
}

func (b *_OpcuaMessageResponseBuilder) DeepCopy() any {
	_copy := b.CreateOpcuaMessageResponseBuilder().(*_OpcuaMessageResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateOpcuaMessageResponseBuilder creates a OpcuaMessageResponseBuilder
func (b *_OpcuaMessageResponse) CreateOpcuaMessageResponseBuilder() OpcuaMessageResponseBuilder {
	if b == nil {
		return NewOpcuaMessageResponseBuilder()
	}
	return &_OpcuaMessageResponseBuilder{_OpcuaMessageResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_OpcuaMessageResponse) GetMessageType() string {
	return "MSG"
}

func (m *_OpcuaMessageResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_OpcuaMessageResponse) GetParent() MessagePDUContract {
	return m.MessagePDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_OpcuaMessageResponse) GetSecurityHeader() SecurityHeader {
	return m.SecurityHeader
}

func (m *_OpcuaMessageResponse) GetMessage() Payload {
	return m.Message
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastOpcuaMessageResponse(structType any) OpcuaMessageResponse {
	if casted, ok := structType.(OpcuaMessageResponse); ok {
		return casted
	}
	if casted, ok := structType.(*OpcuaMessageResponse); ok {
		return *casted
	}
	return nil
}

func (m *_OpcuaMessageResponse) GetTypeName() string {
	return "OpcuaMessageResponse"
}

func (m *_OpcuaMessageResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MessagePDUContract.(*_MessagePDU).getLengthInBits(ctx))

	// Simple field (securityHeader)
	lengthInBits += m.SecurityHeader.GetLengthInBits(ctx)

	// Simple field (message)
	lengthInBits += m.Message.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_OpcuaMessageResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_OpcuaMessageResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MessagePDU, totalLength uint32, response bool, binary bool) (__opcuaMessageResponse OpcuaMessageResponse, err error) {
	m.MessagePDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("OpcuaMessageResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for OpcuaMessageResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	securityHeader, err := ReadSimpleField[SecurityHeader](ctx, "securityHeader", ReadComplex[SecurityHeader](SecurityHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityHeader' field"))
	}
	m.SecurityHeader = securityHeader

	message, err := ReadSimpleField[Payload](ctx, "message", ReadComplex[Payload](PayloadParseWithBufferProducer[Payload]((bool)(binary), (uint32)(uint32(uint32(totalLength)-uint32(securityHeader.GetLengthInBytes(ctx)))-uint32(uint32(16)))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'message' field"))
	}
	m.Message = message

	if closeErr := readBuffer.CloseContext("OpcuaMessageResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for OpcuaMessageResponse")
	}

	return m, nil
}

func (m *_OpcuaMessageResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_OpcuaMessageResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("OpcuaMessageResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for OpcuaMessageResponse")
		}

		if err := WriteSimpleField[SecurityHeader](ctx, "securityHeader", m.GetSecurityHeader(), WriteComplex[SecurityHeader](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'securityHeader' field")
		}

		if err := WriteSimpleField[Payload](ctx, "message", m.GetMessage(), WriteComplex[Payload](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'message' field")
		}

		if popErr := writeBuffer.PopContext("OpcuaMessageResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for OpcuaMessageResponse")
		}
		return nil
	}
	return m.MessagePDUContract.(*_MessagePDU).serializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_OpcuaMessageResponse) GetTotalLength() uint32 {
	return m.TotalLength
}

//
////

func (m *_OpcuaMessageResponse) IsOpcuaMessageResponse() {}

func (m *_OpcuaMessageResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_OpcuaMessageResponse) deepCopy() *_OpcuaMessageResponse {
	if m == nil {
		return nil
	}
	_OpcuaMessageResponseCopy := &_OpcuaMessageResponse{
		m.MessagePDUContract.(*_MessagePDU).deepCopy(),
		utils.DeepCopy[SecurityHeader](m.SecurityHeader),
		utils.DeepCopy[Payload](m.Message),
		m.TotalLength,
	}
	_OpcuaMessageResponseCopy.MessagePDUContract.(*_MessagePDU)._SubType = m
	return _OpcuaMessageResponseCopy
}

func (m *_OpcuaMessageResponse) String() string {
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
