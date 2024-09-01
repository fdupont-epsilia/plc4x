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

// BACnetPropertyStatesDoorSecuredStatus is the corresponding interface of BACnetPropertyStatesDoorSecuredStatus
type BACnetPropertyStatesDoorSecuredStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetDoorSecuredStatus returns DoorSecuredStatus (property field)
	GetDoorSecuredStatus() BACnetDoorSecuredStatusTagged
}

// BACnetPropertyStatesDoorSecuredStatusExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesDoorSecuredStatus.
// This is useful for switch cases.
type BACnetPropertyStatesDoorSecuredStatusExactly interface {
	BACnetPropertyStatesDoorSecuredStatus
	isBACnetPropertyStatesDoorSecuredStatus() bool
}

// _BACnetPropertyStatesDoorSecuredStatus is the data-structure of this message
type _BACnetPropertyStatesDoorSecuredStatus struct {
	BACnetPropertyStatesContract
	DoorSecuredStatus BACnetDoorSecuredStatusTagged
}

var _ BACnetPropertyStatesDoorSecuredStatus = (*_BACnetPropertyStatesDoorSecuredStatus)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesDoorSecuredStatus) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesDoorSecuredStatus) GetDoorSecuredStatus() BACnetDoorSecuredStatusTagged {
	return m.DoorSecuredStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesDoorSecuredStatus factory function for _BACnetPropertyStatesDoorSecuredStatus
func NewBACnetPropertyStatesDoorSecuredStatus(doorSecuredStatus BACnetDoorSecuredStatusTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesDoorSecuredStatus {
	_result := &_BACnetPropertyStatesDoorSecuredStatus{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		DoorSecuredStatus:            doorSecuredStatus,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesDoorSecuredStatus(structType any) BACnetPropertyStatesDoorSecuredStatus {
	if casted, ok := structType.(BACnetPropertyStatesDoorSecuredStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesDoorSecuredStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesDoorSecuredStatus) GetTypeName() string {
	return "BACnetPropertyStatesDoorSecuredStatus"
}

func (m *_BACnetPropertyStatesDoorSecuredStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (doorSecuredStatus)
	lengthInBits += m.DoorSecuredStatus.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesDoorSecuredStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesDoorSecuredStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (_bACnetPropertyStatesDoorSecuredStatus BACnetPropertyStatesDoorSecuredStatus, err error) {
	m.BACnetPropertyStatesContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesDoorSecuredStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesDoorSecuredStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	doorSecuredStatus, err := ReadSimpleField[BACnetDoorSecuredStatusTagged](ctx, "doorSecuredStatus", ReadComplex[BACnetDoorSecuredStatusTagged](BACnetDoorSecuredStatusTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'doorSecuredStatus' field"))
	}
	m.DoorSecuredStatus = doorSecuredStatus

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesDoorSecuredStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesDoorSecuredStatus")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesDoorSecuredStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesDoorSecuredStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesDoorSecuredStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesDoorSecuredStatus")
		}

		if err := WriteSimpleField[BACnetDoorSecuredStatusTagged](ctx, "doorSecuredStatus", m.GetDoorSecuredStatus(), WriteComplex[BACnetDoorSecuredStatusTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'doorSecuredStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesDoorSecuredStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesDoorSecuredStatus")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesDoorSecuredStatus) isBACnetPropertyStatesDoorSecuredStatus() bool {
	return true
}

func (m *_BACnetPropertyStatesDoorSecuredStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
