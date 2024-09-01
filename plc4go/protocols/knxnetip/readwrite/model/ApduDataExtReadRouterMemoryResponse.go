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

// ApduDataExtReadRouterMemoryResponse is the corresponding interface of ApduDataExtReadRouterMemoryResponse
type ApduDataExtReadRouterMemoryResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ApduDataExt
}

// ApduDataExtReadRouterMemoryResponseExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtReadRouterMemoryResponse.
// This is useful for switch cases.
type ApduDataExtReadRouterMemoryResponseExactly interface {
	ApduDataExtReadRouterMemoryResponse
	isApduDataExtReadRouterMemoryResponse() bool
}

// _ApduDataExtReadRouterMemoryResponse is the data-structure of this message
type _ApduDataExtReadRouterMemoryResponse struct {
	ApduDataExtContract
}

var _ ApduDataExtReadRouterMemoryResponse = (*_ApduDataExtReadRouterMemoryResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtReadRouterMemoryResponse) GetExtApciType() uint8 {
	return 0x09
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtReadRouterMemoryResponse) GetParent() ApduDataExtContract {
	return m.ApduDataExtContract
}

// NewApduDataExtReadRouterMemoryResponse factory function for _ApduDataExtReadRouterMemoryResponse
func NewApduDataExtReadRouterMemoryResponse(length uint8) *_ApduDataExtReadRouterMemoryResponse {
	_result := &_ApduDataExtReadRouterMemoryResponse{
		ApduDataExtContract: NewApduDataExt(length),
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtReadRouterMemoryResponse(structType any) ApduDataExtReadRouterMemoryResponse {
	if casted, ok := structType.(ApduDataExtReadRouterMemoryResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtReadRouterMemoryResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtReadRouterMemoryResponse) GetTypeName() string {
	return "ApduDataExtReadRouterMemoryResponse"
}

func (m *_ApduDataExtReadRouterMemoryResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataExtContract.(*_ApduDataExt).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtReadRouterMemoryResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataExtReadRouterMemoryResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduDataExt, length uint8) (_apduDataExtReadRouterMemoryResponse ApduDataExtReadRouterMemoryResponse, err error) {
	m.ApduDataExtContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtReadRouterMemoryResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtReadRouterMemoryResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtReadRouterMemoryResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtReadRouterMemoryResponse")
	}

	return m, nil
}

func (m *_ApduDataExtReadRouterMemoryResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtReadRouterMemoryResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtReadRouterMemoryResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtReadRouterMemoryResponse")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtReadRouterMemoryResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtReadRouterMemoryResponse")
		}
		return nil
	}
	return m.ApduDataExtContract.(*_ApduDataExt).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtReadRouterMemoryResponse) isApduDataExtReadRouterMemoryResponse() bool {
	return true
}

func (m *_ApduDataExtReadRouterMemoryResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
