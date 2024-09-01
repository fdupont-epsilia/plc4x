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

// SetAttributeAllResponse is the corresponding interface of SetAttributeAllResponse
type SetAttributeAllResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CipService
}

// SetAttributeAllResponseExactly can be used when we want exactly this type and not a type which fulfills SetAttributeAllResponse.
// This is useful for switch cases.
type SetAttributeAllResponseExactly interface {
	SetAttributeAllResponse
	isSetAttributeAllResponse() bool
}

// _SetAttributeAllResponse is the data-structure of this message
type _SetAttributeAllResponse struct {
	*_CipService
}

var _ SetAttributeAllResponse = (*_SetAttributeAllResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SetAttributeAllResponse) GetService() uint8 {
	return 0x02
}

func (m *_SetAttributeAllResponse) GetResponse() bool {
	return bool(true)
}

func (m *_SetAttributeAllResponse) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SetAttributeAllResponse) InitializeParent(parent CipService) {}

func (m *_SetAttributeAllResponse) GetParent() CipService {
	return m._CipService
}

// NewSetAttributeAllResponse factory function for _SetAttributeAllResponse
func NewSetAttributeAllResponse(serviceLen uint16) *_SetAttributeAllResponse {
	_result := &_SetAttributeAllResponse{
		_CipService: NewCipService(serviceLen),
	}
	_result._CipService._CipServiceChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSetAttributeAllResponse(structType any) SetAttributeAllResponse {
	if casted, ok := structType.(SetAttributeAllResponse); ok {
		return casted
	}
	if casted, ok := structType.(*SetAttributeAllResponse); ok {
		return *casted
	}
	return nil
}

func (m *_SetAttributeAllResponse) GetTypeName() string {
	return "SetAttributeAllResponse"
}

func (m *_SetAttributeAllResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_SetAttributeAllResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SetAttributeAllResponseParse(ctx context.Context, theBytes []byte, connected bool, serviceLen uint16) (SetAttributeAllResponse, error) {
	return SetAttributeAllResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), connected, serviceLen)
}

func SetAttributeAllResponseParseWithBufferProducer(connected bool, serviceLen uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (SetAttributeAllResponse, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (SetAttributeAllResponse, error) {
		return SetAttributeAllResponseParseWithBuffer(ctx, readBuffer, connected, serviceLen)
	}
}

func SetAttributeAllResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, connected bool, serviceLen uint16) (SetAttributeAllResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SetAttributeAllResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SetAttributeAllResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SetAttributeAllResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SetAttributeAllResponse")
	}

	// Create a partially initialized instance
	_child := &_SetAttributeAllResponse{
		_CipService: &_CipService{
			ServiceLen: serviceLen,
		},
	}
	_child._CipService._CipServiceChildRequirements = _child
	return _child, nil
}

func (m *_SetAttributeAllResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SetAttributeAllResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SetAttributeAllResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SetAttributeAllResponse")
		}

		if popErr := writeBuffer.PopContext("SetAttributeAllResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SetAttributeAllResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SetAttributeAllResponse) isSetAttributeAllResponse() bool {
	return true
}

func (m *_SetAttributeAllResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
