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

// COTPParameterDisconnectAdditionalInformation is the corresponding interface of COTPParameterDisconnectAdditionalInformation
type COTPParameterDisconnectAdditionalInformation interface {
	utils.LengthAware
	utils.Serializable
	COTPParameter
	// GetData returns Data (property field)
	GetData() []byte
}

// COTPParameterDisconnectAdditionalInformationExactly can be used when we want exactly this type and not a type which fulfills COTPParameterDisconnectAdditionalInformation.
// This is useful for switch cases.
type COTPParameterDisconnectAdditionalInformationExactly interface {
	COTPParameterDisconnectAdditionalInformation
	isCOTPParameterDisconnectAdditionalInformation() bool
}

// _COTPParameterDisconnectAdditionalInformation is the data-structure of this message
type _COTPParameterDisconnectAdditionalInformation struct {
	*_COTPParameter
	Data []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_COTPParameterDisconnectAdditionalInformation) GetParameterType() uint8 {
	return 0xE0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_COTPParameterDisconnectAdditionalInformation) InitializeParent(parent COTPParameter) {}

func (m *_COTPParameterDisconnectAdditionalInformation) GetParent() COTPParameter {
	return m._COTPParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_COTPParameterDisconnectAdditionalInformation) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCOTPParameterDisconnectAdditionalInformation factory function for _COTPParameterDisconnectAdditionalInformation
func NewCOTPParameterDisconnectAdditionalInformation(data []byte, rest uint8) *_COTPParameterDisconnectAdditionalInformation {
	_result := &_COTPParameterDisconnectAdditionalInformation{
		Data:           data,
		_COTPParameter: NewCOTPParameter(rest),
	}
	_result._COTPParameter._COTPParameterChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCOTPParameterDisconnectAdditionalInformation(structType interface{}) COTPParameterDisconnectAdditionalInformation {
	if casted, ok := structType.(COTPParameterDisconnectAdditionalInformation); ok {
		return casted
	}
	if casted, ok := structType.(*COTPParameterDisconnectAdditionalInformation); ok {
		return *casted
	}
	return nil
}

func (m *_COTPParameterDisconnectAdditionalInformation) GetTypeName() string {
	return "COTPParameterDisconnectAdditionalInformation"
}

func (m *_COTPParameterDisconnectAdditionalInformation) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_COTPParameterDisconnectAdditionalInformation) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_COTPParameterDisconnectAdditionalInformation) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func COTPParameterDisconnectAdditionalInformationParse(readBuffer utils.ReadBuffer, rest uint8) (COTPParameterDisconnectAdditionalInformation, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("COTPParameterDisconnectAdditionalInformation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for COTPParameterDisconnectAdditionalInformation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	// Byte Array field (data)
	numberOfBytesdata := int(rest)
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field of COTPParameterDisconnectAdditionalInformation")
	}

	if closeErr := readBuffer.CloseContext("COTPParameterDisconnectAdditionalInformation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for COTPParameterDisconnectAdditionalInformation")
	}

	// Create a partially initialized instance
	_child := &_COTPParameterDisconnectAdditionalInformation{
		_COTPParameter: &_COTPParameter{
			Rest: rest,
		},
		Data: data,
	}
	_child._COTPParameter._COTPParameterChildRequirements = _child
	return _child, nil
}

func (m *_COTPParameterDisconnectAdditionalInformation) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_COTPParameterDisconnectAdditionalInformation) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("COTPParameterDisconnectAdditionalInformation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for COTPParameterDisconnectAdditionalInformation")
		}

		// Array Field (data)
		// Byte Array field (data)
		if err := writeBuffer.WriteByteArray("data", m.GetData()); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("COTPParameterDisconnectAdditionalInformation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for COTPParameterDisconnectAdditionalInformation")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_COTPParameterDisconnectAdditionalInformation) isCOTPParameterDisconnectAdditionalInformation() bool {
	return true
}

func (m *_COTPParameterDisconnectAdditionalInformation) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}