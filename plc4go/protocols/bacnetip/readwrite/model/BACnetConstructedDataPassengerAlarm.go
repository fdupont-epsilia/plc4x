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

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataPassengerAlarm is the corresponding interface of BACnetConstructedDataPassengerAlarm
type BACnetConstructedDataPassengerAlarm interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetPassengerAlarm returns PassengerAlarm (property field)
	GetPassengerAlarm() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataPassengerAlarmExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataPassengerAlarm.
// This is useful for switch cases.
type BACnetConstructedDataPassengerAlarmExactly interface {
	BACnetConstructedDataPassengerAlarm
	isBACnetConstructedDataPassengerAlarm() bool
}

// _BACnetConstructedDataPassengerAlarm is the data-structure of this message
type _BACnetConstructedDataPassengerAlarm struct {
	*_BACnetConstructedData
	PassengerAlarm BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataPassengerAlarm = (*_BACnetConstructedDataPassengerAlarm)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataPassengerAlarm) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataPassengerAlarm) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PASSENGER_ALARM
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataPassengerAlarm) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataPassengerAlarm) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataPassengerAlarm) GetPassengerAlarm() BACnetApplicationTagBoolean {
	return m.PassengerAlarm
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataPassengerAlarm) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetPassengerAlarm())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataPassengerAlarm factory function for _BACnetConstructedDataPassengerAlarm
func NewBACnetConstructedDataPassengerAlarm(passengerAlarm BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataPassengerAlarm {
	_result := &_BACnetConstructedDataPassengerAlarm{
		PassengerAlarm:         passengerAlarm,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataPassengerAlarm(structType any) BACnetConstructedDataPassengerAlarm {
	if casted, ok := structType.(BACnetConstructedDataPassengerAlarm); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPassengerAlarm); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataPassengerAlarm) GetTypeName() string {
	return "BACnetConstructedDataPassengerAlarm"
}

func (m *_BACnetConstructedDataPassengerAlarm) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (passengerAlarm)
	lengthInBits += m.PassengerAlarm.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataPassengerAlarm) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataPassengerAlarmParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataPassengerAlarm, error) {
	return BACnetConstructedDataPassengerAlarmParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataPassengerAlarmParseWithBufferProducer(tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataPassengerAlarm, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataPassengerAlarm, error) {
		return BACnetConstructedDataPassengerAlarmParseWithBuffer(ctx, readBuffer, tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
	}
}

func BACnetConstructedDataPassengerAlarmParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataPassengerAlarm, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPassengerAlarm"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPassengerAlarm")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	passengerAlarm, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "passengerAlarm", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'passengerAlarm' field"))
	}

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), passengerAlarm)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPassengerAlarm"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPassengerAlarm")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataPassengerAlarm{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		PassengerAlarm: passengerAlarm,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataPassengerAlarm) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataPassengerAlarm) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPassengerAlarm"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPassengerAlarm")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "passengerAlarm", m.GetPassengerAlarm(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'passengerAlarm' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPassengerAlarm"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPassengerAlarm")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataPassengerAlarm) isBACnetConstructedDataPassengerAlarm() bool {
	return true
}

func (m *_BACnetConstructedDataPassengerAlarm) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
