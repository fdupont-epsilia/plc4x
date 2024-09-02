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

// S7PayloadUserDataItemClkSetResponse is the corresponding interface of S7PayloadUserDataItemClkSetResponse
type S7PayloadUserDataItemClkSetResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7PayloadUserDataItem
	// IsS7PayloadUserDataItemClkSetResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7PayloadUserDataItemClkSetResponse()
}

// _S7PayloadUserDataItemClkSetResponse is the data-structure of this message
type _S7PayloadUserDataItemClkSetResponse struct {
	S7PayloadUserDataItemContract
}

var _ S7PayloadUserDataItemClkSetResponse = (*_S7PayloadUserDataItemClkSetResponse)(nil)
var _ S7PayloadUserDataItemRequirements = (*_S7PayloadUserDataItemClkSetResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemClkSetResponse) GetCpuFunctionGroup() uint8 {
	return 0x07
}

func (m *_S7PayloadUserDataItemClkSetResponse) GetCpuFunctionType() uint8 {
	return 0x08
}

func (m *_S7PayloadUserDataItemClkSetResponse) GetCpuSubfunction() uint8 {
	return 0x04
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemClkSetResponse) GetParent() S7PayloadUserDataItemContract {
	return m.S7PayloadUserDataItemContract
}

// NewS7PayloadUserDataItemClkSetResponse factory function for _S7PayloadUserDataItemClkSetResponse
func NewS7PayloadUserDataItemClkSetResponse(returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) *_S7PayloadUserDataItemClkSetResponse {
	_result := &_S7PayloadUserDataItemClkSetResponse{
		S7PayloadUserDataItemContract: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
	}
	_result.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemClkSetResponse(structType any) S7PayloadUserDataItemClkSetResponse {
	if casted, ok := structType.(S7PayloadUserDataItemClkSetResponse); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemClkSetResponse); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemClkSetResponse) GetTypeName() string {
	return "S7PayloadUserDataItemClkSetResponse"
}

func (m *_S7PayloadUserDataItemClkSetResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_S7PayloadUserDataItemClkSetResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7PayloadUserDataItemClkSetResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7PayloadUserDataItem, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (__s7PayloadUserDataItemClkSetResponse S7PayloadUserDataItemClkSetResponse, err error) {
	m.S7PayloadUserDataItemContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemClkSetResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemClkSetResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemClkSetResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemClkSetResponse")
	}

	return m, nil
}

func (m *_S7PayloadUserDataItemClkSetResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemClkSetResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemClkSetResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemClkSetResponse")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemClkSetResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemClkSetResponse")
		}
		return nil
	}
	return m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemClkSetResponse) IsS7PayloadUserDataItemClkSetResponse() {}

func (m *_S7PayloadUserDataItemClkSetResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
