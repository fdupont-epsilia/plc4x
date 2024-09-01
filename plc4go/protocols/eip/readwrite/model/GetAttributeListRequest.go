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

// GetAttributeListRequest is the corresponding interface of GetAttributeListRequest
type GetAttributeListRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CipService
}

// GetAttributeListRequestExactly can be used when we want exactly this type and not a type which fulfills GetAttributeListRequest.
// This is useful for switch cases.
type GetAttributeListRequestExactly interface {
	GetAttributeListRequest
	isGetAttributeListRequest() bool
}

// _GetAttributeListRequest is the data-structure of this message
type _GetAttributeListRequest struct {
	CipServiceContract
}

var _ GetAttributeListRequest = (*_GetAttributeListRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_GetAttributeListRequest) GetService() uint8 {
	return 0x03
}

func (m *_GetAttributeListRequest) GetResponse() bool {
	return bool(false)
}

func (m *_GetAttributeListRequest) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_GetAttributeListRequest) GetParent() CipServiceContract {
	return m.CipServiceContract
}

// NewGetAttributeListRequest factory function for _GetAttributeListRequest
func NewGetAttributeListRequest(serviceLen uint16) *_GetAttributeListRequest {
	_result := &_GetAttributeListRequest{
		CipServiceContract: NewCipService(serviceLen),
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastGetAttributeListRequest(structType any) GetAttributeListRequest {
	if casted, ok := structType.(GetAttributeListRequest); ok {
		return casted
	}
	if casted, ok := structType.(*GetAttributeListRequest); ok {
		return *casted
	}
	return nil
}

func (m *_GetAttributeListRequest) GetTypeName() string {
	return "GetAttributeListRequest"
}

func (m *_GetAttributeListRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CipServiceContract.(*_CipService).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_GetAttributeListRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_GetAttributeListRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CipService, connected bool, serviceLen uint16) (_getAttributeListRequest GetAttributeListRequest, err error) {
	m.CipServiceContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("GetAttributeListRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for GetAttributeListRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("GetAttributeListRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for GetAttributeListRequest")
	}

	return m, nil
}

func (m *_GetAttributeListRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_GetAttributeListRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("GetAttributeListRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for GetAttributeListRequest")
		}

		if popErr := writeBuffer.PopContext("GetAttributeListRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for GetAttributeListRequest")
		}
		return nil
	}
	return m.CipServiceContract.(*_CipService).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_GetAttributeListRequest) isGetAttributeListRequest() bool {
	return true
}

func (m *_GetAttributeListRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
