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

// ApduDataExtDomainAddressResponse is the corresponding interface of ApduDataExtDomainAddressResponse
type ApduDataExtDomainAddressResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ApduDataExt
	// IsApduDataExtDomainAddressResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataExtDomainAddressResponse()
}

// _ApduDataExtDomainAddressResponse is the data-structure of this message
type _ApduDataExtDomainAddressResponse struct {
	ApduDataExtContract
}

var _ ApduDataExtDomainAddressResponse = (*_ApduDataExtDomainAddressResponse)(nil)
var _ ApduDataExtRequirements = (*_ApduDataExtDomainAddressResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtDomainAddressResponse) GetExtApciType() uint8 {
	return 0x22
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtDomainAddressResponse) GetParent() ApduDataExtContract {
	return m.ApduDataExtContract
}

// NewApduDataExtDomainAddressResponse factory function for _ApduDataExtDomainAddressResponse
func NewApduDataExtDomainAddressResponse(length uint8) *_ApduDataExtDomainAddressResponse {
	_result := &_ApduDataExtDomainAddressResponse{
		ApduDataExtContract: NewApduDataExt(length),
	}
	_result.ApduDataExtContract.(*_ApduDataExt)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtDomainAddressResponse(structType any) ApduDataExtDomainAddressResponse {
	if casted, ok := structType.(ApduDataExtDomainAddressResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtDomainAddressResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtDomainAddressResponse) GetTypeName() string {
	return "ApduDataExtDomainAddressResponse"
}

func (m *_ApduDataExtDomainAddressResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataExtContract.(*_ApduDataExt).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtDomainAddressResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataExtDomainAddressResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduDataExt, length uint8) (__apduDataExtDomainAddressResponse ApduDataExtDomainAddressResponse, err error) {
	m.ApduDataExtContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtDomainAddressResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtDomainAddressResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtDomainAddressResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtDomainAddressResponse")
	}

	return m, nil
}

func (m *_ApduDataExtDomainAddressResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtDomainAddressResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtDomainAddressResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtDomainAddressResponse")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtDomainAddressResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtDomainAddressResponse")
		}
		return nil
	}
	return m.ApduDataExtContract.(*_ApduDataExt).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtDomainAddressResponse) IsApduDataExtDomainAddressResponse() {}

func (m *_ApduDataExtDomainAddressResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
