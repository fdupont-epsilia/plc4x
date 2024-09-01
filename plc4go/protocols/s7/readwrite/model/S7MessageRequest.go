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

// S7MessageRequest is the corresponding interface of S7MessageRequest
type S7MessageRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7Message
}

// S7MessageRequestExactly can be used when we want exactly this type and not a type which fulfills S7MessageRequest.
// This is useful for switch cases.
type S7MessageRequestExactly interface {
	S7MessageRequest
	isS7MessageRequest() bool
}

// _S7MessageRequest is the data-structure of this message
type _S7MessageRequest struct {
	*_S7Message
}

var _ S7MessageRequest = (*_S7MessageRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7MessageRequest) GetMessageType() uint8 {
	return 0x01
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7MessageRequest) InitializeParent(parent S7Message, tpduReference uint16, parameter S7Parameter, payload S7Payload) {
	m.TpduReference = tpduReference
	m.Parameter = parameter
	m.Payload = payload
}

func (m *_S7MessageRequest) GetParent() S7Message {
	return m._S7Message
}

// NewS7MessageRequest factory function for _S7MessageRequest
func NewS7MessageRequest(tpduReference uint16, parameter S7Parameter, payload S7Payload) *_S7MessageRequest {
	_result := &_S7MessageRequest{
		_S7Message: NewS7Message(tpduReference, parameter, payload),
	}
	_result._S7Message._S7MessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7MessageRequest(structType any) S7MessageRequest {
	if casted, ok := structType.(S7MessageRequest); ok {
		return casted
	}
	if casted, ok := structType.(*S7MessageRequest); ok {
		return *casted
	}
	return nil
}

func (m *_S7MessageRequest) GetTypeName() string {
	return "S7MessageRequest"
}

func (m *_S7MessageRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_S7MessageRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7MessageRequestParse(ctx context.Context, theBytes []byte) (S7MessageRequest, error) {
	return S7MessageRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func S7MessageRequestParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (S7MessageRequest, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (S7MessageRequest, error) {
		return S7MessageRequestParseWithBuffer(ctx, readBuffer)
	}
}

func S7MessageRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (S7MessageRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7MessageRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7MessageRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("S7MessageRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7MessageRequest")
	}

	// Create a partially initialized instance
	_child := &_S7MessageRequest{
		_S7Message: &_S7Message{},
	}
	_child._S7Message._S7MessageChildRequirements = _child
	return _child, nil
}

func (m *_S7MessageRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7MessageRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7MessageRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7MessageRequest")
		}

		if popErr := writeBuffer.PopContext("S7MessageRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7MessageRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7MessageRequest) isS7MessageRequest() bool {
	return true
}

func (m *_S7MessageRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
