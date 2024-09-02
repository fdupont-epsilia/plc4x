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

// ApduDataExtReadRouterStatusResponse is the corresponding interface of ApduDataExtReadRouterStatusResponse
type ApduDataExtReadRouterStatusResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ApduDataExt
	// IsApduDataExtReadRouterStatusResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataExtReadRouterStatusResponse()
}

// _ApduDataExtReadRouterStatusResponse is the data-structure of this message
type _ApduDataExtReadRouterStatusResponse struct {
	ApduDataExtContract
}

var _ ApduDataExtReadRouterStatusResponse = (*_ApduDataExtReadRouterStatusResponse)(nil)
var _ ApduDataExtRequirements = (*_ApduDataExtReadRouterStatusResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtReadRouterStatusResponse) GetExtApciType() uint8 {
	return 0x0E
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtReadRouterStatusResponse) GetParent() ApduDataExtContract {
	return m.ApduDataExtContract
}

// NewApduDataExtReadRouterStatusResponse factory function for _ApduDataExtReadRouterStatusResponse
func NewApduDataExtReadRouterStatusResponse(length uint8) *_ApduDataExtReadRouterStatusResponse {
	_result := &_ApduDataExtReadRouterStatusResponse{
		ApduDataExtContract: NewApduDataExt(length),
	}
	_result.ApduDataExtContract.(*_ApduDataExt)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtReadRouterStatusResponse(structType any) ApduDataExtReadRouterStatusResponse {
	if casted, ok := structType.(ApduDataExtReadRouterStatusResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtReadRouterStatusResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtReadRouterStatusResponse) GetTypeName() string {
	return "ApduDataExtReadRouterStatusResponse"
}

func (m *_ApduDataExtReadRouterStatusResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataExtContract.(*_ApduDataExt).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtReadRouterStatusResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataExtReadRouterStatusResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduDataExt, length uint8) (__apduDataExtReadRouterStatusResponse ApduDataExtReadRouterStatusResponse, err error) {
	m.ApduDataExtContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtReadRouterStatusResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtReadRouterStatusResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtReadRouterStatusResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtReadRouterStatusResponse")
	}

	return m, nil
}

func (m *_ApduDataExtReadRouterStatusResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtReadRouterStatusResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtReadRouterStatusResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtReadRouterStatusResponse")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtReadRouterStatusResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtReadRouterStatusResponse")
		}
		return nil
	}
	return m.ApduDataExtContract.(*_ApduDataExt).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtReadRouterStatusResponse) IsApduDataExtReadRouterStatusResponse() {}

func (m *_ApduDataExtReadRouterStatusResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
