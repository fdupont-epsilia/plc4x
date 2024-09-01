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

// BACnetPropertyStatesSystemStatus is the corresponding interface of BACnetPropertyStatesSystemStatus
type BACnetPropertyStatesSystemStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetSystemStatus returns SystemStatus (property field)
	GetSystemStatus() BACnetDeviceStatusTagged
}

// BACnetPropertyStatesSystemStatusExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesSystemStatus.
// This is useful for switch cases.
type BACnetPropertyStatesSystemStatusExactly interface {
	BACnetPropertyStatesSystemStatus
	isBACnetPropertyStatesSystemStatus() bool
}

// _BACnetPropertyStatesSystemStatus is the data-structure of this message
type _BACnetPropertyStatesSystemStatus struct {
	BACnetPropertyStatesContract
	SystemStatus BACnetDeviceStatusTagged
}

var _ BACnetPropertyStatesSystemStatus = (*_BACnetPropertyStatesSystemStatus)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesSystemStatus) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesSystemStatus) GetSystemStatus() BACnetDeviceStatusTagged {
	return m.SystemStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesSystemStatus factory function for _BACnetPropertyStatesSystemStatus
func NewBACnetPropertyStatesSystemStatus(systemStatus BACnetDeviceStatusTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesSystemStatus {
	_result := &_BACnetPropertyStatesSystemStatus{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		SystemStatus:                 systemStatus,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesSystemStatus(structType any) BACnetPropertyStatesSystemStatus {
	if casted, ok := structType.(BACnetPropertyStatesSystemStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesSystemStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesSystemStatus) GetTypeName() string {
	return "BACnetPropertyStatesSystemStatus"
}

func (m *_BACnetPropertyStatesSystemStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (systemStatus)
	lengthInBits += m.SystemStatus.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesSystemStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesSystemStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (_bACnetPropertyStatesSystemStatus BACnetPropertyStatesSystemStatus, err error) {
	m.BACnetPropertyStatesContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesSystemStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesSystemStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	systemStatus, err := ReadSimpleField[BACnetDeviceStatusTagged](ctx, "systemStatus", ReadComplex[BACnetDeviceStatusTagged](BACnetDeviceStatusTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'systemStatus' field"))
	}
	m.SystemStatus = systemStatus

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesSystemStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesSystemStatus")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesSystemStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesSystemStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesSystemStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesSystemStatus")
		}

		if err := WriteSimpleField[BACnetDeviceStatusTagged](ctx, "systemStatus", m.GetSystemStatus(), WriteComplex[BACnetDeviceStatusTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'systemStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesSystemStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesSystemStatus")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesSystemStatus) isBACnetPropertyStatesSystemStatus() bool {
	return true
}

func (m *_BACnetPropertyStatesSystemStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
