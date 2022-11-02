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

// AccessControlDataValidAccessRequest is the corresponding interface of AccessControlDataValidAccessRequest
type AccessControlDataValidAccessRequest interface {
	utils.LengthAware
	utils.Serializable
	AccessControlData
	// GetAccessControlDirection returns AccessControlDirection (property field)
	GetAccessControlDirection() AccessControlDirection
	// GetData returns Data (property field)
	GetData() []byte
}

// AccessControlDataValidAccessRequestExactly can be used when we want exactly this type and not a type which fulfills AccessControlDataValidAccessRequest.
// This is useful for switch cases.
type AccessControlDataValidAccessRequestExactly interface {
	AccessControlDataValidAccessRequest
	isAccessControlDataValidAccessRequest() bool
}

// _AccessControlDataValidAccessRequest is the data-structure of this message
type _AccessControlDataValidAccessRequest struct {
	*_AccessControlData
	AccessControlDirection AccessControlDirection
	Data                   []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AccessControlDataValidAccessRequest) InitializeParent(parent AccessControlData, commandTypeContainer AccessControlCommandTypeContainer, networkId byte, accessPointId byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.NetworkId = networkId
	m.AccessPointId = accessPointId
}

func (m *_AccessControlDataValidAccessRequest) GetParent() AccessControlData {
	return m._AccessControlData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AccessControlDataValidAccessRequest) GetAccessControlDirection() AccessControlDirection {
	return m.AccessControlDirection
}

func (m *_AccessControlDataValidAccessRequest) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAccessControlDataValidAccessRequest factory function for _AccessControlDataValidAccessRequest
func NewAccessControlDataValidAccessRequest(accessControlDirection AccessControlDirection, data []byte, commandTypeContainer AccessControlCommandTypeContainer, networkId byte, accessPointId byte) *_AccessControlDataValidAccessRequest {
	_result := &_AccessControlDataValidAccessRequest{
		AccessControlDirection: accessControlDirection,
		Data:                   data,
		_AccessControlData:     NewAccessControlData(commandTypeContainer, networkId, accessPointId),
	}
	_result._AccessControlData._AccessControlDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAccessControlDataValidAccessRequest(structType interface{}) AccessControlDataValidAccessRequest {
	if casted, ok := structType.(AccessControlDataValidAccessRequest); ok {
		return casted
	}
	if casted, ok := structType.(*AccessControlDataValidAccessRequest); ok {
		return *casted
	}
	return nil
}

func (m *_AccessControlDataValidAccessRequest) GetTypeName() string {
	return "AccessControlDataValidAccessRequest"
}

func (m *_AccessControlDataValidAccessRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_AccessControlDataValidAccessRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (accessControlDirection)
	lengthInBits += 8

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_AccessControlDataValidAccessRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AccessControlDataValidAccessRequestParse(readBuffer utils.ReadBuffer, commandTypeContainer AccessControlCommandTypeContainer) (AccessControlDataValidAccessRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AccessControlDataValidAccessRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AccessControlDataValidAccessRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (accessControlDirection)
	if pullErr := readBuffer.PullContext("accessControlDirection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for accessControlDirection")
	}
	_accessControlDirection, _accessControlDirectionErr := AccessControlDirectionParse(readBuffer)
	if _accessControlDirectionErr != nil {
		return nil, errors.Wrap(_accessControlDirectionErr, "Error parsing 'accessControlDirection' field of AccessControlDataValidAccessRequest")
	}
	accessControlDirection := _accessControlDirection
	if closeErr := readBuffer.CloseContext("accessControlDirection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for accessControlDirection")
	}
	// Byte Array field (data)
	numberOfBytesdata := int(uint16(commandTypeContainer.NumBytes()) - uint16(uint16(3)))
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field of AccessControlDataValidAccessRequest")
	}

	if closeErr := readBuffer.CloseContext("AccessControlDataValidAccessRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AccessControlDataValidAccessRequest")
	}

	// Create a partially initialized instance
	_child := &_AccessControlDataValidAccessRequest{
		_AccessControlData:     &_AccessControlData{},
		AccessControlDirection: accessControlDirection,
		Data:                   data,
	}
	_child._AccessControlData._AccessControlDataChildRequirements = _child
	return _child, nil
}

func (m *_AccessControlDataValidAccessRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AccessControlDataValidAccessRequest) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AccessControlDataValidAccessRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AccessControlDataValidAccessRequest")
		}

		// Simple Field (accessControlDirection)
		if pushErr := writeBuffer.PushContext("accessControlDirection"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for accessControlDirection")
		}
		_accessControlDirectionErr := writeBuffer.WriteSerializable(m.GetAccessControlDirection())
		if popErr := writeBuffer.PopContext("accessControlDirection"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for accessControlDirection")
		}
		if _accessControlDirectionErr != nil {
			return errors.Wrap(_accessControlDirectionErr, "Error serializing 'accessControlDirection' field")
		}

		// Array Field (data)
		// Byte Array field (data)
		if err := writeBuffer.WriteByteArray("data", m.GetData()); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("AccessControlDataValidAccessRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AccessControlDataValidAccessRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_AccessControlDataValidAccessRequest) isAccessControlDataValidAccessRequest() bool {
	return true
}

func (m *_AccessControlDataValidAccessRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}