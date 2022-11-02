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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// ApduDataAdcResponse is the corresponding interface of ApduDataAdcResponse
type ApduDataAdcResponse interface {
	utils.LengthAware
	utils.Serializable
	ApduData
}

// ApduDataAdcResponseExactly can be used when we want exactly this type and not a type which fulfills ApduDataAdcResponse.
// This is useful for switch cases.
type ApduDataAdcResponseExactly interface {
	ApduDataAdcResponse
	isApduDataAdcResponse() bool
}

// _ApduDataAdcResponse is the data-structure of this message
type _ApduDataAdcResponse struct {
	*_ApduData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataAdcResponse) GetApciType() uint8 {
	return 0x7
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataAdcResponse) InitializeParent(parent ApduData) {}

func (m *_ApduDataAdcResponse) GetParent() ApduData {
	return m._ApduData
}

// NewApduDataAdcResponse factory function for _ApduDataAdcResponse
func NewApduDataAdcResponse(dataLength uint8) *_ApduDataAdcResponse {
	_result := &_ApduDataAdcResponse{
		_ApduData: NewApduData(dataLength),
	}
	_result._ApduData._ApduDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataAdcResponse(structType interface{}) ApduDataAdcResponse {
	if casted, ok := structType.(ApduDataAdcResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataAdcResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataAdcResponse) GetTypeName() string {
	return "ApduDataAdcResponse"
}

func (m *_ApduDataAdcResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ApduDataAdcResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_ApduDataAdcResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataAdcResponseParse(readBuffer utils.ReadBuffer, dataLength uint8) (ApduDataAdcResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataAdcResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataAdcResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataAdcResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataAdcResponse")
	}

	// Create a partially initialized instance
	_child := &_ApduDataAdcResponse{
		_ApduData: &_ApduData{
			DataLength: dataLength,
		},
	}
	_child._ApduData._ApduDataChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataAdcResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataAdcResponse) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataAdcResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataAdcResponse")
		}

		if popErr := writeBuffer.PopContext("ApduDataAdcResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataAdcResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ApduDataAdcResponse) isApduDataAdcResponse() bool {
	return true
}

func (m *_ApduDataAdcResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}