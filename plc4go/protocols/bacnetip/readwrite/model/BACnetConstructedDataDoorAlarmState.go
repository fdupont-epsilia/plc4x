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

// BACnetConstructedDataDoorAlarmState is the corresponding interface of BACnetConstructedDataDoorAlarmState
type BACnetConstructedDataDoorAlarmState interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetDoorAlarmState returns DoorAlarmState (property field)
	GetDoorAlarmState() BACnetDoorAlarmStateTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDoorAlarmStateTagged
}

// BACnetConstructedDataDoorAlarmStateExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDoorAlarmState.
// This is useful for switch cases.
type BACnetConstructedDataDoorAlarmStateExactly interface {
	BACnetConstructedDataDoorAlarmState
	isBACnetConstructedDataDoorAlarmState() bool
}

// _BACnetConstructedDataDoorAlarmState is the data-structure of this message
type _BACnetConstructedDataDoorAlarmState struct {
	BACnetConstructedDataContract
	DoorAlarmState BACnetDoorAlarmStateTagged
}

var _ BACnetConstructedDataDoorAlarmState = (*_BACnetConstructedDataDoorAlarmState)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDoorAlarmState) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDoorAlarmState) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DOOR_ALARM_STATE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDoorAlarmState) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDoorAlarmState) GetDoorAlarmState() BACnetDoorAlarmStateTagged {
	return m.DoorAlarmState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDoorAlarmState) GetActualValue() BACnetDoorAlarmStateTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDoorAlarmStateTagged(m.GetDoorAlarmState())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDoorAlarmState factory function for _BACnetConstructedDataDoorAlarmState
func NewBACnetConstructedDataDoorAlarmState(doorAlarmState BACnetDoorAlarmStateTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDoorAlarmState {
	_result := &_BACnetConstructedDataDoorAlarmState{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		DoorAlarmState:                doorAlarmState,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDoorAlarmState(structType any) BACnetConstructedDataDoorAlarmState {
	if casted, ok := structType.(BACnetConstructedDataDoorAlarmState); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDoorAlarmState); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDoorAlarmState) GetTypeName() string {
	return "BACnetConstructedDataDoorAlarmState"
}

func (m *_BACnetConstructedDataDoorAlarmState) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (doorAlarmState)
	lengthInBits += m.DoorAlarmState.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDoorAlarmState) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDoorAlarmState) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (_bACnetConstructedDataDoorAlarmState BACnetConstructedDataDoorAlarmState, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDoorAlarmState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDoorAlarmState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	doorAlarmState, err := ReadSimpleField[BACnetDoorAlarmStateTagged](ctx, "doorAlarmState", ReadComplex[BACnetDoorAlarmStateTagged](BACnetDoorAlarmStateTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'doorAlarmState' field"))
	}
	m.DoorAlarmState = doorAlarmState

	actualValue, err := ReadVirtualField[BACnetDoorAlarmStateTagged](ctx, "actualValue", (*BACnetDoorAlarmStateTagged)(nil), doorAlarmState)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDoorAlarmState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDoorAlarmState")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDoorAlarmState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDoorAlarmState) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDoorAlarmState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDoorAlarmState")
		}

		if err := WriteSimpleField[BACnetDoorAlarmStateTagged](ctx, "doorAlarmState", m.GetDoorAlarmState(), WriteComplex[BACnetDoorAlarmStateTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'doorAlarmState' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDoorAlarmState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDoorAlarmState")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDoorAlarmState) isBACnetConstructedDataDoorAlarmState() bool {
	return true
}

func (m *_BACnetConstructedDataDoorAlarmState) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
