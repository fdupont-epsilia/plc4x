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

// EipListIdentityRequest is the corresponding interface of EipListIdentityRequest
type EipListIdentityRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	EipPacket
}

// EipListIdentityRequestExactly can be used when we want exactly this type and not a type which fulfills EipListIdentityRequest.
// This is useful for switch cases.
type EipListIdentityRequestExactly interface {
	EipListIdentityRequest
	isEipListIdentityRequest() bool
}

// _EipListIdentityRequest is the data-structure of this message
type _EipListIdentityRequest struct {
	*_EipPacket
}

var _ EipListIdentityRequest = (*_EipListIdentityRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_EipListIdentityRequest) GetCommand() uint16 {
	return 0x0063
}

func (m *_EipListIdentityRequest) GetResponse() bool {
	return bool(false)
}

func (m *_EipListIdentityRequest) GetPacketLength() uint16 {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_EipListIdentityRequest) InitializeParent(parent EipPacket, sessionHandle uint32, status uint32, senderContext []byte, options uint32) {
	m.SessionHandle = sessionHandle
	m.Status = status
	m.SenderContext = senderContext
	m.Options = options
}

func (m *_EipListIdentityRequest) GetParent() EipPacket {
	return m._EipPacket
}

// NewEipListIdentityRequest factory function for _EipListIdentityRequest
func NewEipListIdentityRequest(sessionHandle uint32, status uint32, senderContext []byte, options uint32) *_EipListIdentityRequest {
	_result := &_EipListIdentityRequest{
		_EipPacket: NewEipPacket(sessionHandle, status, senderContext, options),
	}
	_result._EipPacket._EipPacketChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastEipListIdentityRequest(structType any) EipListIdentityRequest {
	if casted, ok := structType.(EipListIdentityRequest); ok {
		return casted
	}
	if casted, ok := structType.(*EipListIdentityRequest); ok {
		return *casted
	}
	return nil
}

func (m *_EipListIdentityRequest) GetTypeName() string {
	return "EipListIdentityRequest"
}

func (m *_EipListIdentityRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_EipListIdentityRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func EipListIdentityRequestParse(ctx context.Context, theBytes []byte, response bool) (EipListIdentityRequest, error) {
	return EipListIdentityRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func EipListIdentityRequestParseWithBufferProducer(response bool) func(ctx context.Context, readBuffer utils.ReadBuffer) (EipListIdentityRequest, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (EipListIdentityRequest, error) {
		return EipListIdentityRequestParseWithBuffer(ctx, readBuffer, response)
	}
}

func EipListIdentityRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (EipListIdentityRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("EipListIdentityRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EipListIdentityRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("EipListIdentityRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EipListIdentityRequest")
	}

	// Create a partially initialized instance
	_child := &_EipListIdentityRequest{
		_EipPacket: &_EipPacket{},
	}
	_child._EipPacket._EipPacketChildRequirements = _child
	return _child, nil
}

func (m *_EipListIdentityRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_EipListIdentityRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("EipListIdentityRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for EipListIdentityRequest")
		}

		if popErr := writeBuffer.PopContext("EipListIdentityRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for EipListIdentityRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_EipListIdentityRequest) isEipListIdentityRequest() bool {
	return true
}

func (m *_EipListIdentityRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
