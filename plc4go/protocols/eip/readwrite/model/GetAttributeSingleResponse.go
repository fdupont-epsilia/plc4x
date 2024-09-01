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

// GetAttributeSingleResponse is the corresponding interface of GetAttributeSingleResponse
type GetAttributeSingleResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CipService
}

// GetAttributeSingleResponseExactly can be used when we want exactly this type and not a type which fulfills GetAttributeSingleResponse.
// This is useful for switch cases.
type GetAttributeSingleResponseExactly interface {
	GetAttributeSingleResponse
	isGetAttributeSingleResponse() bool
}

// _GetAttributeSingleResponse is the data-structure of this message
type _GetAttributeSingleResponse struct {
	*_CipService
}

var _ GetAttributeSingleResponse = (*_GetAttributeSingleResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_GetAttributeSingleResponse) GetService() uint8 {
	return 0x0E
}

func (m *_GetAttributeSingleResponse) GetResponse() bool {
	return bool(true)
}

func (m *_GetAttributeSingleResponse) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_GetAttributeSingleResponse) InitializeParent(parent CipService) {}

func (m *_GetAttributeSingleResponse) GetParent() CipService {
	return m._CipService
}

// NewGetAttributeSingleResponse factory function for _GetAttributeSingleResponse
func NewGetAttributeSingleResponse(serviceLen uint16) *_GetAttributeSingleResponse {
	_result := &_GetAttributeSingleResponse{
		_CipService: NewCipService(serviceLen),
	}
	_result._CipService._CipServiceChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastGetAttributeSingleResponse(structType any) GetAttributeSingleResponse {
	if casted, ok := structType.(GetAttributeSingleResponse); ok {
		return casted
	}
	if casted, ok := structType.(*GetAttributeSingleResponse); ok {
		return *casted
	}
	return nil
}

func (m *_GetAttributeSingleResponse) GetTypeName() string {
	return "GetAttributeSingleResponse"
}

func (m *_GetAttributeSingleResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_GetAttributeSingleResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func GetAttributeSingleResponseParse(ctx context.Context, theBytes []byte, connected bool, serviceLen uint16) (GetAttributeSingleResponse, error) {
	return GetAttributeSingleResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), connected, serviceLen)
}

func GetAttributeSingleResponseParseWithBufferProducer(connected bool, serviceLen uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (GetAttributeSingleResponse, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (GetAttributeSingleResponse, error) {
		return GetAttributeSingleResponseParseWithBuffer(ctx, readBuffer, connected, serviceLen)
	}
}

func GetAttributeSingleResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, connected bool, serviceLen uint16) (GetAttributeSingleResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("GetAttributeSingleResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for GetAttributeSingleResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("GetAttributeSingleResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for GetAttributeSingleResponse")
	}

	// Create a partially initialized instance
	_child := &_GetAttributeSingleResponse{
		_CipService: &_CipService{
			ServiceLen: serviceLen,
		},
	}
	_child._CipService._CipServiceChildRequirements = _child
	return _child, nil
}

func (m *_GetAttributeSingleResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_GetAttributeSingleResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("GetAttributeSingleResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for GetAttributeSingleResponse")
		}

		if popErr := writeBuffer.PopContext("GetAttributeSingleResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for GetAttributeSingleResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_GetAttributeSingleResponse) isGetAttributeSingleResponse() bool {
	return true
}

func (m *_GetAttributeSingleResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
