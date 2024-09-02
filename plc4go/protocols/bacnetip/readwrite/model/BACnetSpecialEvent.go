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

// BACnetSpecialEvent is the corresponding interface of BACnetSpecialEvent
type BACnetSpecialEvent interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetPeriod returns Period (property field)
	GetPeriod() BACnetSpecialEventPeriod
	// GetListOfTimeValues returns ListOfTimeValues (property field)
	GetListOfTimeValues() BACnetSpecialEventListOfTimeValues
	// GetEventPriority returns EventPriority (property field)
	GetEventPriority() BACnetContextTagUnsignedInteger
	// IsBACnetSpecialEvent is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetSpecialEvent()
}

// _BACnetSpecialEvent is the data-structure of this message
type _BACnetSpecialEvent struct {
	Period           BACnetSpecialEventPeriod
	ListOfTimeValues BACnetSpecialEventListOfTimeValues
	EventPriority    BACnetContextTagUnsignedInteger
}

var _ BACnetSpecialEvent = (*_BACnetSpecialEvent)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetSpecialEvent) GetPeriod() BACnetSpecialEventPeriod {
	return m.Period
}

func (m *_BACnetSpecialEvent) GetListOfTimeValues() BACnetSpecialEventListOfTimeValues {
	return m.ListOfTimeValues
}

func (m *_BACnetSpecialEvent) GetEventPriority() BACnetContextTagUnsignedInteger {
	return m.EventPriority
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetSpecialEvent factory function for _BACnetSpecialEvent
func NewBACnetSpecialEvent(period BACnetSpecialEventPeriod, listOfTimeValues BACnetSpecialEventListOfTimeValues, eventPriority BACnetContextTagUnsignedInteger) *_BACnetSpecialEvent {
	return &_BACnetSpecialEvent{Period: period, ListOfTimeValues: listOfTimeValues, EventPriority: eventPriority}
}

// Deprecated: use the interface for direct cast
func CastBACnetSpecialEvent(structType any) BACnetSpecialEvent {
	if casted, ok := structType.(BACnetSpecialEvent); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetSpecialEvent); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetSpecialEvent) GetTypeName() string {
	return "BACnetSpecialEvent"
}

func (m *_BACnetSpecialEvent) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (period)
	lengthInBits += m.Period.GetLengthInBits(ctx)

	// Simple field (listOfTimeValues)
	lengthInBits += m.ListOfTimeValues.GetLengthInBits(ctx)

	// Simple field (eventPriority)
	lengthInBits += m.EventPriority.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetSpecialEvent) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetSpecialEventParse(ctx context.Context, theBytes []byte) (BACnetSpecialEvent, error) {
	return BACnetSpecialEventParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetSpecialEventParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetSpecialEvent, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetSpecialEvent, error) {
		return BACnetSpecialEventParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetSpecialEventParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetSpecialEvent, error) {
	v, err := (&_BACnetSpecialEvent{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetSpecialEvent) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetSpecialEvent BACnetSpecialEvent, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetSpecialEvent"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetSpecialEvent")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	period, err := ReadSimpleField[BACnetSpecialEventPeriod](ctx, "period", ReadComplex[BACnetSpecialEventPeriod](BACnetSpecialEventPeriodParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'period' field"))
	}
	m.Period = period

	listOfTimeValues, err := ReadSimpleField[BACnetSpecialEventListOfTimeValues](ctx, "listOfTimeValues", ReadComplex[BACnetSpecialEventListOfTimeValues](BACnetSpecialEventListOfTimeValuesParseWithBufferProducer((uint8)(uint8(2))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfTimeValues' field"))
	}
	m.ListOfTimeValues = listOfTimeValues

	eventPriority, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "eventPriority", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventPriority' field"))
	}
	m.EventPriority = eventPriority

	if closeErr := readBuffer.CloseContext("BACnetSpecialEvent"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetSpecialEvent")
	}

	return m, nil
}

func (m *_BACnetSpecialEvent) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetSpecialEvent) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetSpecialEvent"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetSpecialEvent")
	}

	if err := WriteSimpleField[BACnetSpecialEventPeriod](ctx, "period", m.GetPeriod(), WriteComplex[BACnetSpecialEventPeriod](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'period' field")
	}

	if err := WriteSimpleField[BACnetSpecialEventListOfTimeValues](ctx, "listOfTimeValues", m.GetListOfTimeValues(), WriteComplex[BACnetSpecialEventListOfTimeValues](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'listOfTimeValues' field")
	}

	if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "eventPriority", m.GetEventPriority(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'eventPriority' field")
	}

	if popErr := writeBuffer.PopContext("BACnetSpecialEvent"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetSpecialEvent")
	}
	return nil
}

func (m *_BACnetSpecialEvent) IsBACnetSpecialEvent() {}

func (m *_BACnetSpecialEvent) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
