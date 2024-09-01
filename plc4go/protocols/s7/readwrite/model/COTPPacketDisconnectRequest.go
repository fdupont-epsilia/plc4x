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

// COTPPacketDisconnectRequest is the corresponding interface of COTPPacketDisconnectRequest
type COTPPacketDisconnectRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	COTPPacket
	// GetDestinationReference returns DestinationReference (property field)
	GetDestinationReference() uint16
	// GetSourceReference returns SourceReference (property field)
	GetSourceReference() uint16
	// GetProtocolClass returns ProtocolClass (property field)
	GetProtocolClass() COTPProtocolClass
}

// COTPPacketDisconnectRequestExactly can be used when we want exactly this type and not a type which fulfills COTPPacketDisconnectRequest.
// This is useful for switch cases.
type COTPPacketDisconnectRequestExactly interface {
	COTPPacketDisconnectRequest
	isCOTPPacketDisconnectRequest() bool
}

// _COTPPacketDisconnectRequest is the data-structure of this message
type _COTPPacketDisconnectRequest struct {
	*_COTPPacket
	DestinationReference uint16
	SourceReference      uint16
	ProtocolClass        COTPProtocolClass
}

var _ COTPPacketDisconnectRequest = (*_COTPPacketDisconnectRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_COTPPacketDisconnectRequest) GetTpduCode() uint8 {
	return 0x80
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_COTPPacketDisconnectRequest) InitializeParent(parent COTPPacket, parameters []COTPParameter, payload S7Message) {
	m.Parameters = parameters
	m.Payload = payload
}

func (m *_COTPPacketDisconnectRequest) GetParent() COTPPacket {
	return m._COTPPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_COTPPacketDisconnectRequest) GetDestinationReference() uint16 {
	return m.DestinationReference
}

func (m *_COTPPacketDisconnectRequest) GetSourceReference() uint16 {
	return m.SourceReference
}

func (m *_COTPPacketDisconnectRequest) GetProtocolClass() COTPProtocolClass {
	return m.ProtocolClass
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCOTPPacketDisconnectRequest factory function for _COTPPacketDisconnectRequest
func NewCOTPPacketDisconnectRequest(destinationReference uint16, sourceReference uint16, protocolClass COTPProtocolClass, parameters []COTPParameter, payload S7Message, cotpLen uint16) *_COTPPacketDisconnectRequest {
	_result := &_COTPPacketDisconnectRequest{
		DestinationReference: destinationReference,
		SourceReference:      sourceReference,
		ProtocolClass:        protocolClass,
		_COTPPacket:          NewCOTPPacket(parameters, payload, cotpLen),
	}
	_result._COTPPacket._COTPPacketChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCOTPPacketDisconnectRequest(structType any) COTPPacketDisconnectRequest {
	if casted, ok := structType.(COTPPacketDisconnectRequest); ok {
		return casted
	}
	if casted, ok := structType.(*COTPPacketDisconnectRequest); ok {
		return *casted
	}
	return nil
}

func (m *_COTPPacketDisconnectRequest) GetTypeName() string {
	return "COTPPacketDisconnectRequest"
}

func (m *_COTPPacketDisconnectRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (destinationReference)
	lengthInBits += 16

	// Simple field (sourceReference)
	lengthInBits += 16

	// Simple field (protocolClass)
	lengthInBits += 8

	return lengthInBits
}

func (m *_COTPPacketDisconnectRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func COTPPacketDisconnectRequestParse(ctx context.Context, theBytes []byte, cotpLen uint16) (COTPPacketDisconnectRequest, error) {
	return COTPPacketDisconnectRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), cotpLen)
}

func COTPPacketDisconnectRequestParseWithBufferProducer(cotpLen uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (COTPPacketDisconnectRequest, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (COTPPacketDisconnectRequest, error) {
		return COTPPacketDisconnectRequestParseWithBuffer(ctx, readBuffer, cotpLen)
	}
}

func COTPPacketDisconnectRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cotpLen uint16) (COTPPacketDisconnectRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("COTPPacketDisconnectRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for COTPPacketDisconnectRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	destinationReference, err := ReadSimpleField(ctx, "destinationReference", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'destinationReference' field"))
	}

	sourceReference, err := ReadSimpleField(ctx, "sourceReference", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sourceReference' field"))
	}

	protocolClass, err := ReadEnumField[COTPProtocolClass](ctx, "protocolClass", "COTPProtocolClass", ReadEnum(COTPProtocolClassByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'protocolClass' field"))
	}

	if closeErr := readBuffer.CloseContext("COTPPacketDisconnectRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for COTPPacketDisconnectRequest")
	}

	// Create a partially initialized instance
	_child := &_COTPPacketDisconnectRequest{
		_COTPPacket: &_COTPPacket{
			CotpLen: cotpLen,
		},
		DestinationReference: destinationReference,
		SourceReference:      sourceReference,
		ProtocolClass:        protocolClass,
	}
	_child._COTPPacket._COTPPacketChildRequirements = _child
	return _child, nil
}

func (m *_COTPPacketDisconnectRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_COTPPacketDisconnectRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("COTPPacketDisconnectRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for COTPPacketDisconnectRequest")
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

		if popErr := writeBuffer.PopContext("COTPPacketDisconnectRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for COTPPacketDisconnectRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_COTPPacketDisconnectRequest) isCOTPPacketDisconnectRequest() bool {
	return true
}

func (m *_COTPPacketDisconnectRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
