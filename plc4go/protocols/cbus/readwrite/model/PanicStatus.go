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

// PanicStatus is the corresponding interface of PanicStatus
type PanicStatus interface {
	utils.LengthAware
	utils.Serializable
	// GetStatus returns Status (property field)
	GetStatus() uint8
	// GetIsNoPanic returns IsNoPanic (virtual field)
	GetIsNoPanic() bool
	// GetIsReserved returns IsReserved (virtual field)
	GetIsReserved() bool
	// GetIsPanicCurrentlyActive returns IsPanicCurrentlyActive (virtual field)
	GetIsPanicCurrentlyActive() bool
}

// PanicStatusExactly can be used when we want exactly this type and not a type which fulfills PanicStatus.
// This is useful for switch cases.
type PanicStatusExactly interface {
	PanicStatus
	isPanicStatus() bool
}

// _PanicStatus is the data-structure of this message
type _PanicStatus struct {
	Status uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PanicStatus) GetStatus() uint8 {
	return m.Status
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_PanicStatus) GetIsNoPanic() bool {
	return bool(bool((m.GetStatus()) == (0x00)))
}

func (m *_PanicStatus) GetIsReserved() bool {
	return bool(bool(bool((m.GetStatus()) >= (0x01))) && bool(bool((m.GetStatus()) <= (0xFE))))
}

func (m *_PanicStatus) GetIsPanicCurrentlyActive() bool {
	return bool(bool((m.GetStatus()) > (0xFE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewPanicStatus factory function for _PanicStatus
func NewPanicStatus(status uint8) *_PanicStatus {
	return &_PanicStatus{Status: status}
}

// Deprecated: use the interface for direct cast
func CastPanicStatus(structType interface{}) PanicStatus {
	if casted, ok := structType.(PanicStatus); ok {
		return casted
	}
	if casted, ok := structType.(*PanicStatus); ok {
		return *casted
	}
	return nil
}

func (m *_PanicStatus) GetTypeName() string {
	return "PanicStatus"
}

func (m *_PanicStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_PanicStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (status)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_PanicStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func PanicStatusParse(readBuffer utils.ReadBuffer) (PanicStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("PanicStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PanicStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (status)
	_status, _statusErr := readBuffer.ReadUint8("status", 8)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field of PanicStatus")
	}
	status := _status

	// Virtual field
	_isNoPanic := bool((status) == (0x00))
	isNoPanic := bool(_isNoPanic)
	_ = isNoPanic

	// Virtual field
	_isReserved := bool(bool((status) >= (0x01))) && bool(bool((status) <= (0xFE)))
	isReserved := bool(_isReserved)
	_ = isReserved

	// Virtual field
	_isPanicCurrentlyActive := bool((status) > (0xFE))
	isPanicCurrentlyActive := bool(_isPanicCurrentlyActive)
	_ = isPanicCurrentlyActive

	if closeErr := readBuffer.CloseContext("PanicStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PanicStatus")
	}

	// Create the instance
	return &_PanicStatus{
		Status: status,
	}, nil
}

func (m *_PanicStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PanicStatus) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("PanicStatus"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for PanicStatus")
	}

	// Simple Field (status)
	status := uint8(m.GetStatus())
	_statusErr := writeBuffer.WriteUint8("status", 8, (status))
	if _statusErr != nil {
		return errors.Wrap(_statusErr, "Error serializing 'status' field")
	}
	// Virtual field
	if _isNoPanicErr := writeBuffer.WriteVirtual("isNoPanic", m.GetIsNoPanic()); _isNoPanicErr != nil {
		return errors.Wrap(_isNoPanicErr, "Error serializing 'isNoPanic' field")
	}
	// Virtual field
	if _isReservedErr := writeBuffer.WriteVirtual("isReserved", m.GetIsReserved()); _isReservedErr != nil {
		return errors.Wrap(_isReservedErr, "Error serializing 'isReserved' field")
	}
	// Virtual field
	if _isPanicCurrentlyActiveErr := writeBuffer.WriteVirtual("isPanicCurrentlyActive", m.GetIsPanicCurrentlyActive()); _isPanicCurrentlyActiveErr != nil {
		return errors.Wrap(_isPanicCurrentlyActiveErr, "Error serializing 'isPanicCurrentlyActive' field")
	}

	if popErr := writeBuffer.PopContext("PanicStatus"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for PanicStatus")
	}
	return nil
}

func (m *_PanicStatus) isPanicStatus() bool {
	return true
}

func (m *_PanicStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}