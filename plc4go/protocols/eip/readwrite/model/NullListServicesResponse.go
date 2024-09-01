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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// NullListServicesResponse is the corresponding interface of NullListServicesResponse
type NullListServicesResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	EipPacket
}

// NullListServicesResponseExactly can be used when we want exactly this type and not a type which fulfills NullListServicesResponse.
// This is useful for switch cases.
type NullListServicesResponseExactly interface {
	NullListServicesResponse
	isNullListServicesResponse() bool
}

// _NullListServicesResponse is the data-structure of this message
type _NullListServicesResponse struct {
	EipPacketContract
}

var _ NullListServicesResponse = (*_NullListServicesResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NullListServicesResponse) GetCommand() uint16 {
	return 0x0004
}

func (m *_NullListServicesResponse) GetResponse() bool {
	return bool(true)
}

func (m *_NullListServicesResponse) GetPacketLength() uint16 {
	return uint16(0)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NullListServicesResponse) GetParent() EipPacketContract {
	return m.EipPacketContract
}

// NewNullListServicesResponse factory function for _NullListServicesResponse
func NewNullListServicesResponse(sessionHandle uint32, status uint32, senderContext []byte, options uint32) *_NullListServicesResponse {
	_result := &_NullListServicesResponse{
		EipPacketContract: NewEipPacket(sessionHandle, status, senderContext, options),
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastNullListServicesResponse(structType any) NullListServicesResponse {
	if casted, ok := structType.(NullListServicesResponse); ok {
		return casted
	}
	if casted, ok := structType.(*NullListServicesResponse); ok {
		return *casted
	}
	return nil
}

func (m *_NullListServicesResponse) GetTypeName() string {
	return "NullListServicesResponse"
}

func (m *_NullListServicesResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.EipPacketContract.(*_EipPacket).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_NullListServicesResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NullListServicesResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_EipPacket, response bool) (_nullListServicesResponse NullListServicesResponse, err error) {
	m.EipPacketContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NullListServicesResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NullListServicesResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("NullListServicesResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NullListServicesResponse")
	}

	return m, nil
}

func (m *_NullListServicesResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NullListServicesResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NullListServicesResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NullListServicesResponse")
		}

		if popErr := writeBuffer.PopContext("NullListServicesResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NullListServicesResponse")
		}
		return nil
	}
	return m.EipPacketContract.(*_EipPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NullListServicesResponse) isNullListServicesResponse() bool {
	return true
}

func (m *_NullListServicesResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
