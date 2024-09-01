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

// Constant values.
const S7Message_PROTOCOLID uint8 = 0x32

// S7Message is the corresponding interface of S7Message
type S7Message interface {
	S7MessageContract
	S7MessageRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// S7MessageContract provides a set of functions which can be overwritten by a sub struct
type S7MessageContract interface {
	// GetTpduReference returns TpduReference (property field)
	GetTpduReference() uint16
	// GetParameter returns Parameter (property field)
	GetParameter() S7Parameter
	// GetPayload returns Payload (property field)
	GetPayload() S7Payload
}

// S7MessageRequirements provides a set of functions which need to be implemented by a sub struct
type S7MessageRequirements interface {
	// GetMessageType returns MessageType (discriminator field)
	GetMessageType() uint8
}

// S7MessageExactly can be used when we want exactly this type and not a type which fulfills S7Message.
// This is useful for switch cases.
type S7MessageExactly interface {
	S7Message
	isS7Message() bool
}

// _S7Message is the data-structure of this message
type _S7Message struct {
	_S7MessageChildRequirements
	TpduReference uint16
	Parameter     S7Parameter
	Payload       S7Payload
	// Reserved Fields
	reservedField0 *uint16
}

var _ S7MessageContract = (*_S7Message)(nil)

type _S7MessageChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetMessageType() uint8
}

type S7MessageParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child S7Message, serializeChildFunction func() error) error
	GetTypeName() string
}

type S7MessageChild interface {
	utils.Serializable
	InitializeParent(parent S7Message, tpduReference uint16, parameter S7Parameter, payload S7Payload)
	GetParent() *S7Message

	GetTypeName() string
	S7Message
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7Message) GetTpduReference() uint16 {
	return m.TpduReference
}

func (m *_S7Message) GetParameter() S7Parameter {
	return m.Parameter
}

func (m *_S7Message) GetPayload() S7Payload {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_S7Message) GetProtocolId() uint8 {
	return S7Message_PROTOCOLID
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7Message factory function for _S7Message
func NewS7Message(tpduReference uint16, parameter S7Parameter, payload S7Payload) *_S7Message {
	return &_S7Message{TpduReference: tpduReference, Parameter: parameter, Payload: payload}
}

// Deprecated: use the interface for direct cast
func CastS7Message(structType any) S7Message {
	if casted, ok := structType.(S7Message); ok {
		return casted
	}
	if casted, ok := structType.(*S7Message); ok {
		return *casted
	}
	return nil
}

func (m *_S7Message) GetTypeName() string {
	return "S7Message"
}

func (m *_S7Message) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Const Field (protocolId)
	lengthInBits += 8
	// Discriminator Field (messageType)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 16

	// Simple field (tpduReference)
	lengthInBits += 16

	// Implicit Field (parameterLength)
	lengthInBits += 16

	// Implicit Field (payloadLength)
	lengthInBits += 16

	// Optional Field (parameter)
	if m.Parameter != nil {
		lengthInBits += m.Parameter.GetLengthInBits(ctx)
	}

	// Optional Field (payload)
	if m.Payload != nil {
		lengthInBits += m.Payload.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_S7Message) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7MessageParse(ctx context.Context, theBytes []byte) (S7Message, error) {
	return S7MessageParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func S7MessageParseWithBufferProducer[T S7Message]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		buffer, err := S7MessageParseWithBuffer(ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return buffer.(T), err
	}
}

func S7MessageParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (S7Message, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7Message"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7Message")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	protocolId, err := ReadConstField[uint8](ctx, "protocolId", ReadUnsignedByte(readBuffer, uint8(8)), S7Message_PROTOCOLID)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'protocolId' field"))
	}
	_ = protocolId

	messageType, err := ReadDiscriminatorField[uint8](ctx, "messageType", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'messageType' field"))
	}

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedShort(readBuffer, uint8(16)), uint16(0x0000))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	tpduReference, err := ReadSimpleField(ctx, "tpduReference", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'tpduReference' field"))
	}

	parameterLength, err := ReadImplicitField[uint16](ctx, "parameterLength", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'parameterLength' field"))
	}
	_ = parameterLength

	payloadLength, err := ReadImplicitField[uint16](ctx, "payloadLength", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payloadLength' field"))
	}
	_ = payloadLength

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type S7MessageChildSerializeRequirement interface {
		S7Message
		InitializeParent(S7Message, uint16, S7Parameter, S7Payload)
		GetParent() S7Message
	}
	var _childTemp any
	var _child S7MessageChildSerializeRequirement
	var typeSwitchError error
	switch {
	case messageType == 0x01: // S7MessageRequest
		_childTemp, typeSwitchError = S7MessageRequestParseWithBuffer(ctx, readBuffer)
	case messageType == 0x02: // S7MessageResponse
		_childTemp, typeSwitchError = S7MessageResponseParseWithBuffer(ctx, readBuffer)
	case messageType == 0x03: // S7MessageResponseData
		_childTemp, typeSwitchError = S7MessageResponseDataParseWithBuffer(ctx, readBuffer)
	case messageType == 0x07: // S7MessageUserData
		_childTemp, typeSwitchError = S7MessageUserDataParseWithBuffer(ctx, readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [messageType=%v]", messageType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of S7Message")
	}
	_child = _childTemp.(S7MessageChildSerializeRequirement)

	_parameter, err := ReadOptionalField[S7Parameter](ctx, "parameter", ReadComplex[S7Parameter](S7ParameterParseWithBufferProducer[S7Parameter]((uint8)(messageType)), readBuffer), bool((parameterLength) > (0)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'parameter' field"))
	}
	var parameter S7Parameter
	if _parameter != nil {
		parameter = *_parameter
	}

	_payload, err := ReadOptionalField[S7Payload](ctx, "payload", ReadComplex[S7Payload](S7PayloadParseWithBufferProducer[S7Payload]((uint8)(messageType), (S7Parameter)((parameter))), readBuffer), bool((payloadLength) > (0)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	var payload S7Payload
	if _payload != nil {
		payload = *_payload
	}

	if closeErr := readBuffer.CloseContext("S7Message"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7Message")
	}

	// Finish initializing
	_child.InitializeParent(_child, tpduReference, parameter, payload)
	_child.GetParent().(*_S7Message).reservedField0 = reservedField0
	return _child, nil
}

func (pm *_S7Message) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child S7Message, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("S7Message"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for S7Message")
	}

	if err := WriteConstField(ctx, "protocolId", S7Message_PROTOCOLID, WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'protocolId' field")
	}

	if err := WriteDiscriminatorField(ctx, "messageType", m.GetMessageType(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'messageType' field")
	}

	if err := WriteReservedField[uint16](ctx, "reserved", uint16(0x0000), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 1")
	}

	if err := WriteSimpleField[uint16](ctx, "tpduReference", m.GetTpduReference(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'tpduReference' field")
	}
	parameterLength := uint16(utils.InlineIf(bool((m.GetParameter()) != (nil)), func() any { return uint16((m.GetParameter()).GetLengthInBytes(ctx)) }, func() any { return uint16(uint16(0)) }).(uint16))
	if err := WriteImplicitField(ctx, "parameterLength", parameterLength, WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'parameterLength' field")
	}
	payloadLength := uint16(utils.InlineIf(bool((m.GetPayload()) != (nil)), func() any { return uint16((m.GetPayload()).GetLengthInBytes(ctx)) }, func() any { return uint16(uint16(0)) }).(uint16))
	if err := WriteImplicitField(ctx, "payloadLength", payloadLength, WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'payloadLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if err := WriteOptionalField[S7Parameter](ctx, "parameter", GetRef(m.GetParameter()), WriteComplex[S7Parameter](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'parameter' field")
	}

	if err := WriteOptionalField[S7Payload](ctx, "payload", GetRef(m.GetPayload()), WriteComplex[S7Payload](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'payload' field")
	}

	if popErr := writeBuffer.PopContext("S7Message"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for S7Message")
	}
	return nil
}

func (m *_S7Message) isS7Message() bool {
	return true
}

func (m *_S7Message) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
