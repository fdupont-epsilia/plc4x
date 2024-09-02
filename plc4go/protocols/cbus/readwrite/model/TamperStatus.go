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
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// TamperStatus is the corresponding interface of TamperStatus
type TamperStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetStatus returns Status (property field)
	GetStatus() uint8
	// GetIsNoTamper returns IsNoTamper (virtual field)
	GetIsNoTamper() bool
	// GetIsReserved returns IsReserved (virtual field)
	GetIsReserved() bool
	// GetIsTamperActive returns IsTamperActive (virtual field)
	GetIsTamperActive() bool
	// IsTamperStatus is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTamperStatus()
}

// _TamperStatus is the data-structure of this message
type _TamperStatus struct {
	Status uint8
}

var _ TamperStatus = (*_TamperStatus)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TamperStatus) GetStatus() uint8 {
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

func (m *_TamperStatus) GetIsNoTamper() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetStatus()) == (0x00)))
}

func (m *_TamperStatus) GetIsReserved() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool(bool((m.GetStatus()) >= (0x01))) && bool(bool((m.GetStatus()) <= (0xFE))))
}

func (m *_TamperStatus) GetIsTamperActive() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetStatus()) > (0xFE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTamperStatus factory function for _TamperStatus
func NewTamperStatus(status uint8) *_TamperStatus {
	return &_TamperStatus{Status: status}
}

// Deprecated: use the interface for direct cast
func CastTamperStatus(structType any) TamperStatus {
	if casted, ok := structType.(TamperStatus); ok {
		return casted
	}
	if casted, ok := structType.(*TamperStatus); ok {
		return *casted
	}
	return nil
}

func (m *_TamperStatus) GetTypeName() string {
	return "TamperStatus"
}

func (m *_TamperStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (status)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_TamperStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TamperStatusParse(ctx context.Context, theBytes []byte) (TamperStatus, error) {
	return TamperStatusParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func TamperStatusParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (TamperStatus, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (TamperStatus, error) {
		return TamperStatusParseWithBuffer(ctx, readBuffer)
	}
}

func TamperStatusParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TamperStatus, error) {
	v, err := (&_TamperStatus{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_TamperStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__tamperStatus TamperStatus, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TamperStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TamperStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	status, err := ReadSimpleField(ctx, "status", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'status' field"))
	}
	m.Status = status

	isNoTamper, err := ReadVirtualField[bool](ctx, "isNoTamper", (*bool)(nil), bool((status) == (0x00)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isNoTamper' field"))
	}
	_ = isNoTamper

	isReserved, err := ReadVirtualField[bool](ctx, "isReserved", (*bool)(nil), bool(bool((status) >= (0x01))) && bool(bool((status) <= (0xFE))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isReserved' field"))
	}
	_ = isReserved

	isTamperActive, err := ReadVirtualField[bool](ctx, "isTamperActive", (*bool)(nil), bool((status) > (0xFE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isTamperActive' field"))
	}
	_ = isTamperActive

	if closeErr := readBuffer.CloseContext("TamperStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TamperStatus")
	}

	return m, nil
}

func (m *_TamperStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TamperStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("TamperStatus"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for TamperStatus")
	}

	if err := WriteSimpleField[uint8](ctx, "status", m.GetStatus(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'status' field")
	}
	// Virtual field
	isNoTamper := m.GetIsNoTamper()
	_ = isNoTamper
	if _isNoTamperErr := writeBuffer.WriteVirtual(ctx, "isNoTamper", m.GetIsNoTamper()); _isNoTamperErr != nil {
		return errors.Wrap(_isNoTamperErr, "Error serializing 'isNoTamper' field")
	}
	// Virtual field
	isReserved := m.GetIsReserved()
	_ = isReserved
	if _isReservedErr := writeBuffer.WriteVirtual(ctx, "isReserved", m.GetIsReserved()); _isReservedErr != nil {
		return errors.Wrap(_isReservedErr, "Error serializing 'isReserved' field")
	}
	// Virtual field
	isTamperActive := m.GetIsTamperActive()
	_ = isTamperActive
	if _isTamperActiveErr := writeBuffer.WriteVirtual(ctx, "isTamperActive", m.GetIsTamperActive()); _isTamperActiveErr != nil {
		return errors.Wrap(_isTamperActiveErr, "Error serializing 'isTamperActive' field")
	}

	if popErr := writeBuffer.PopContext("TamperStatus"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for TamperStatus")
	}
	return nil
}

func (m *_TamperStatus) IsTamperStatus() {}

func (m *_TamperStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
