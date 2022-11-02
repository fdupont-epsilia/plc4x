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

// AccessControlDataCloseAccessPoint is the corresponding interface of AccessControlDataCloseAccessPoint
type AccessControlDataCloseAccessPoint interface {
	utils.LengthAware
	utils.Serializable
	AccessControlData
}

// AccessControlDataCloseAccessPointExactly can be used when we want exactly this type and not a type which fulfills AccessControlDataCloseAccessPoint.
// This is useful for switch cases.
type AccessControlDataCloseAccessPointExactly interface {
	AccessControlDataCloseAccessPoint
	isAccessControlDataCloseAccessPoint() bool
}

// _AccessControlDataCloseAccessPoint is the data-structure of this message
type _AccessControlDataCloseAccessPoint struct {
	*_AccessControlData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AccessControlDataCloseAccessPoint) InitializeParent(parent AccessControlData, commandTypeContainer AccessControlCommandTypeContainer, networkId byte, accessPointId byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.NetworkId = networkId
	m.AccessPointId = accessPointId
}

func (m *_AccessControlDataCloseAccessPoint) GetParent() AccessControlData {
	return m._AccessControlData
}

// NewAccessControlDataCloseAccessPoint factory function for _AccessControlDataCloseAccessPoint
func NewAccessControlDataCloseAccessPoint(commandTypeContainer AccessControlCommandTypeContainer, networkId byte, accessPointId byte) *_AccessControlDataCloseAccessPoint {
	_result := &_AccessControlDataCloseAccessPoint{
		_AccessControlData: NewAccessControlData(commandTypeContainer, networkId, accessPointId),
	}
	_result._AccessControlData._AccessControlDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAccessControlDataCloseAccessPoint(structType interface{}) AccessControlDataCloseAccessPoint {
	if casted, ok := structType.(AccessControlDataCloseAccessPoint); ok {
		return casted
	}
	if casted, ok := structType.(*AccessControlDataCloseAccessPoint); ok {
		return *casted
	}
	return nil
}

func (m *_AccessControlDataCloseAccessPoint) GetTypeName() string {
	return "AccessControlDataCloseAccessPoint"
}

func (m *_AccessControlDataCloseAccessPoint) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_AccessControlDataCloseAccessPoint) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_AccessControlDataCloseAccessPoint) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AccessControlDataCloseAccessPointParse(readBuffer utils.ReadBuffer) (AccessControlDataCloseAccessPoint, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AccessControlDataCloseAccessPoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AccessControlDataCloseAccessPoint")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("AccessControlDataCloseAccessPoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AccessControlDataCloseAccessPoint")
	}

	// Create a partially initialized instance
	_child := &_AccessControlDataCloseAccessPoint{
		_AccessControlData: &_AccessControlData{},
	}
	_child._AccessControlData._AccessControlDataChildRequirements = _child
	return _child, nil
}

func (m *_AccessControlDataCloseAccessPoint) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AccessControlDataCloseAccessPoint) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AccessControlDataCloseAccessPoint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AccessControlDataCloseAccessPoint")
		}

		if popErr := writeBuffer.PopContext("AccessControlDataCloseAccessPoint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AccessControlDataCloseAccessPoint")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_AccessControlDataCloseAccessPoint) isAccessControlDataCloseAccessPoint() bool {
	return true
}

func (m *_AccessControlDataCloseAccessPoint) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}