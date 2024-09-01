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

// SecurityDataPasswordEntryStatus is the corresponding interface of SecurityDataPasswordEntryStatus
type SecurityDataPasswordEntryStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SecurityData
	// GetCode returns Code (property field)
	GetCode() byte
	// GetIsPasswordEntrySucceeded returns IsPasswordEntrySucceeded (virtual field)
	GetIsPasswordEntrySucceeded() bool
	// GetIsPasswordEntryFailed returns IsPasswordEntryFailed (virtual field)
	GetIsPasswordEntryFailed() bool
	// GetIsPasswordEntryDisabled returns IsPasswordEntryDisabled (virtual field)
	GetIsPasswordEntryDisabled() bool
	// GetIsPasswordEntryEnabledAgain returns IsPasswordEntryEnabledAgain (virtual field)
	GetIsPasswordEntryEnabledAgain() bool
	// GetIsReserved returns IsReserved (virtual field)
	GetIsReserved() bool
}

// SecurityDataPasswordEntryStatusExactly can be used when we want exactly this type and not a type which fulfills SecurityDataPasswordEntryStatus.
// This is useful for switch cases.
type SecurityDataPasswordEntryStatusExactly interface {
	SecurityDataPasswordEntryStatus
	isSecurityDataPasswordEntryStatus() bool
}

// _SecurityDataPasswordEntryStatus is the data-structure of this message
type _SecurityDataPasswordEntryStatus struct {
	SecurityDataContract
	Code byte
}

var _ SecurityDataPasswordEntryStatus = (*_SecurityDataPasswordEntryStatus)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataPasswordEntryStatus) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataPasswordEntryStatus) GetCode() byte {
	return m.Code
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_SecurityDataPasswordEntryStatus) GetIsPasswordEntrySucceeded() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetCode()) == (0x01)))
}

func (m *_SecurityDataPasswordEntryStatus) GetIsPasswordEntryFailed() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetCode()) == (0x02)))
}

func (m *_SecurityDataPasswordEntryStatus) GetIsPasswordEntryDisabled() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetCode()) == (0x03)))
}

func (m *_SecurityDataPasswordEntryStatus) GetIsPasswordEntryEnabledAgain() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetCode()) == (0x04)))
}

func (m *_SecurityDataPasswordEntryStatus) GetIsReserved() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetCode()) >= (0x05)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSecurityDataPasswordEntryStatus factory function for _SecurityDataPasswordEntryStatus
func NewSecurityDataPasswordEntryStatus(code byte, commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataPasswordEntryStatus {
	_result := &_SecurityDataPasswordEntryStatus{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
		Code:                 code,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataPasswordEntryStatus(structType any) SecurityDataPasswordEntryStatus {
	if casted, ok := structType.(SecurityDataPasswordEntryStatus); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataPasswordEntryStatus); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataPasswordEntryStatus) GetTypeName() string {
	return "SecurityDataPasswordEntryStatus"
}

func (m *_SecurityDataPasswordEntryStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	// Simple field (code)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_SecurityDataPasswordEntryStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataPasswordEntryStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (_securityDataPasswordEntryStatus SecurityDataPasswordEntryStatus, err error) {
	m.SecurityDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataPasswordEntryStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataPasswordEntryStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	code, err := ReadSimpleField(ctx, "code", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'code' field"))
	}
	m.Code = code

	isPasswordEntrySucceeded, err := ReadVirtualField[bool](ctx, "isPasswordEntrySucceeded", (*bool)(nil), bool((code) == (0x01)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isPasswordEntrySucceeded' field"))
	}
	_ = isPasswordEntrySucceeded

	isPasswordEntryFailed, err := ReadVirtualField[bool](ctx, "isPasswordEntryFailed", (*bool)(nil), bool((code) == (0x02)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isPasswordEntryFailed' field"))
	}
	_ = isPasswordEntryFailed

	isPasswordEntryDisabled, err := ReadVirtualField[bool](ctx, "isPasswordEntryDisabled", (*bool)(nil), bool((code) == (0x03)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isPasswordEntryDisabled' field"))
	}
	_ = isPasswordEntryDisabled

	isPasswordEntryEnabledAgain, err := ReadVirtualField[bool](ctx, "isPasswordEntryEnabledAgain", (*bool)(nil), bool((code) == (0x04)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isPasswordEntryEnabledAgain' field"))
	}
	_ = isPasswordEntryEnabledAgain

	isReserved, err := ReadVirtualField[bool](ctx, "isReserved", (*bool)(nil), bool((code) >= (0x05)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isReserved' field"))
	}
	_ = isReserved

	if closeErr := readBuffer.CloseContext("SecurityDataPasswordEntryStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataPasswordEntryStatus")
	}

	return m, nil
}

func (m *_SecurityDataPasswordEntryStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataPasswordEntryStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataPasswordEntryStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataPasswordEntryStatus")
		}

		if err := WriteSimpleField[byte](ctx, "code", m.GetCode(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'code' field")
		}
		// Virtual field
		isPasswordEntrySucceeded := m.GetIsPasswordEntrySucceeded()
		_ = isPasswordEntrySucceeded
		if _isPasswordEntrySucceededErr := writeBuffer.WriteVirtual(ctx, "isPasswordEntrySucceeded", m.GetIsPasswordEntrySucceeded()); _isPasswordEntrySucceededErr != nil {
			return errors.Wrap(_isPasswordEntrySucceededErr, "Error serializing 'isPasswordEntrySucceeded' field")
		}
		// Virtual field
		isPasswordEntryFailed := m.GetIsPasswordEntryFailed()
		_ = isPasswordEntryFailed
		if _isPasswordEntryFailedErr := writeBuffer.WriteVirtual(ctx, "isPasswordEntryFailed", m.GetIsPasswordEntryFailed()); _isPasswordEntryFailedErr != nil {
			return errors.Wrap(_isPasswordEntryFailedErr, "Error serializing 'isPasswordEntryFailed' field")
		}
		// Virtual field
		isPasswordEntryDisabled := m.GetIsPasswordEntryDisabled()
		_ = isPasswordEntryDisabled
		if _isPasswordEntryDisabledErr := writeBuffer.WriteVirtual(ctx, "isPasswordEntryDisabled", m.GetIsPasswordEntryDisabled()); _isPasswordEntryDisabledErr != nil {
			return errors.Wrap(_isPasswordEntryDisabledErr, "Error serializing 'isPasswordEntryDisabled' field")
		}
		// Virtual field
		isPasswordEntryEnabledAgain := m.GetIsPasswordEntryEnabledAgain()
		_ = isPasswordEntryEnabledAgain
		if _isPasswordEntryEnabledAgainErr := writeBuffer.WriteVirtual(ctx, "isPasswordEntryEnabledAgain", m.GetIsPasswordEntryEnabledAgain()); _isPasswordEntryEnabledAgainErr != nil {
			return errors.Wrap(_isPasswordEntryEnabledAgainErr, "Error serializing 'isPasswordEntryEnabledAgain' field")
		}
		// Virtual field
		isReserved := m.GetIsReserved()
		_ = isReserved
		if _isReservedErr := writeBuffer.WriteVirtual(ctx, "isReserved", m.GetIsReserved()); _isReservedErr != nil {
			return errors.Wrap(_isReservedErr, "Error serializing 'isReserved' field")
		}

		if popErr := writeBuffer.PopContext("SecurityDataPasswordEntryStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataPasswordEntryStatus")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataPasswordEntryStatus) isSecurityDataPasswordEntryStatus() bool {
	return true
}

func (m *_SecurityDataPasswordEntryStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
