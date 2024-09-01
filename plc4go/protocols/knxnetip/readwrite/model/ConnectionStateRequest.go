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
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// ConnectionStateRequest is the corresponding interface of ConnectionStateRequest
type ConnectionStateRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	KnxNetIpMessage
	// GetCommunicationChannelId returns CommunicationChannelId (property field)
	GetCommunicationChannelId() uint8
	// GetHpaiControlEndpoint returns HpaiControlEndpoint (property field)
	GetHpaiControlEndpoint() HPAIControlEndpoint
}

// ConnectionStateRequestExactly can be used when we want exactly this type and not a type which fulfills ConnectionStateRequest.
// This is useful for switch cases.
type ConnectionStateRequestExactly interface {
	ConnectionStateRequest
	isConnectionStateRequest() bool
}

// _ConnectionStateRequest is the data-structure of this message
type _ConnectionStateRequest struct {
	KnxNetIpMessageContract
	CommunicationChannelId uint8
	HpaiControlEndpoint    HPAIControlEndpoint
	// Reserved Fields
	reservedField0 *uint8
}

var _ ConnectionStateRequest = (*_ConnectionStateRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ConnectionStateRequest) GetMsgType() uint16 {
	return 0x0207
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ConnectionStateRequest) GetParent() KnxNetIpMessageContract {
	return m.KnxNetIpMessageContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ConnectionStateRequest) GetCommunicationChannelId() uint8 {
	return m.CommunicationChannelId
}

func (m *_ConnectionStateRequest) GetHpaiControlEndpoint() HPAIControlEndpoint {
	return m.HpaiControlEndpoint
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewConnectionStateRequest factory function for _ConnectionStateRequest
func NewConnectionStateRequest(communicationChannelId uint8, hpaiControlEndpoint HPAIControlEndpoint) *_ConnectionStateRequest {
	_result := &_ConnectionStateRequest{
		KnxNetIpMessageContract: NewKnxNetIpMessage(),
		CommunicationChannelId:  communicationChannelId,
		HpaiControlEndpoint:     hpaiControlEndpoint,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastConnectionStateRequest(structType any) ConnectionStateRequest {
	if casted, ok := structType.(ConnectionStateRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ConnectionStateRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ConnectionStateRequest) GetTypeName() string {
	return "ConnectionStateRequest"
}

func (m *_ConnectionStateRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.KnxNetIpMessageContract.(*_KnxNetIpMessage).getLengthInBits(ctx))

	// Simple field (communicationChannelId)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (hpaiControlEndpoint)
	lengthInBits += m.HpaiControlEndpoint.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ConnectionStateRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ConnectionStateRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_KnxNetIpMessage) (_connectionStateRequest ConnectionStateRequest, err error) {
	m.KnxNetIpMessageContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ConnectionStateRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ConnectionStateRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	communicationChannelId, err := ReadSimpleField(ctx, "communicationChannelId", ReadUnsignedByte(readBuffer, uint8(8)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'communicationChannelId' field"))
	}
	m.CommunicationChannelId = communicationChannelId

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	hpaiControlEndpoint, err := ReadSimpleField[HPAIControlEndpoint](ctx, "hpaiControlEndpoint", ReadComplex[HPAIControlEndpoint](HPAIControlEndpointParseWithBuffer, readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hpaiControlEndpoint' field"))
	}
	m.HpaiControlEndpoint = hpaiControlEndpoint

	if closeErr := readBuffer.CloseContext("ConnectionStateRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ConnectionStateRequest")
	}

	return m, nil
}

func (m *_ConnectionStateRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ConnectionStateRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ConnectionStateRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ConnectionStateRequest")
		}

		if err := WriteSimpleField[uint8](ctx, "communicationChannelId", m.GetCommunicationChannelId(), WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'communicationChannelId' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[HPAIControlEndpoint](ctx, "hpaiControlEndpoint", m.GetHpaiControlEndpoint(), WriteComplex[HPAIControlEndpoint](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'hpaiControlEndpoint' field")
		}

		if popErr := writeBuffer.PopContext("ConnectionStateRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ConnectionStateRequest")
		}
		return nil
	}
	return m.KnxNetIpMessageContract.(*_KnxNetIpMessage).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ConnectionStateRequest) isConnectionStateRequest() bool {
	return true
}

func (m *_ConnectionStateRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
