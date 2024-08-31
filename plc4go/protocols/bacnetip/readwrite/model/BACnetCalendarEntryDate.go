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

// BACnetCalendarEntryDate is the corresponding interface of BACnetCalendarEntryDate
type BACnetCalendarEntryDate interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetCalendarEntry
	// GetDateValue returns DateValue (property field)
	GetDateValue() BACnetContextTagDate
}

// BACnetCalendarEntryDateExactly can be used when we want exactly this type and not a type which fulfills BACnetCalendarEntryDate.
// This is useful for switch cases.
type BACnetCalendarEntryDateExactly interface {
	BACnetCalendarEntryDate
	isBACnetCalendarEntryDate() bool
}

// _BACnetCalendarEntryDate is the data-structure of this message
type _BACnetCalendarEntryDate struct {
	*_BACnetCalendarEntry
	DateValue BACnetContextTagDate
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetCalendarEntryDate) InitializeParent(parent BACnetCalendarEntry, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetCalendarEntryDate) GetParent() BACnetCalendarEntry {
	return m._BACnetCalendarEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetCalendarEntryDate) GetDateValue() BACnetContextTagDate {
	return m.DateValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetCalendarEntryDate factory function for _BACnetCalendarEntryDate
func NewBACnetCalendarEntryDate(dateValue BACnetContextTagDate, peekedTagHeader BACnetTagHeader) *_BACnetCalendarEntryDate {
	_result := &_BACnetCalendarEntryDate{
		DateValue:            dateValue,
		_BACnetCalendarEntry: NewBACnetCalendarEntry(peekedTagHeader),
	}
	_result._BACnetCalendarEntry._BACnetCalendarEntryChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetCalendarEntryDate(structType any) BACnetCalendarEntryDate {
	if casted, ok := structType.(BACnetCalendarEntryDate); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetCalendarEntryDate); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetCalendarEntryDate) GetTypeName() string {
	return "BACnetCalendarEntryDate"
}

func (m *_BACnetCalendarEntryDate) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (dateValue)
	lengthInBits += m.DateValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetCalendarEntryDate) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetCalendarEntryDateParse(ctx context.Context, theBytes []byte) (BACnetCalendarEntryDate, error) {
	return BACnetCalendarEntryDateParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetCalendarEntryDateParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCalendarEntryDate, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCalendarEntryDate, error) {
		return BACnetCalendarEntryDateParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetCalendarEntryDateParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCalendarEntryDate, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetCalendarEntryDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetCalendarEntryDate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dateValue, err := ReadSimpleField[BACnetContextTagDate](ctx, "dateValue", ReadComplex[BACnetContextTagDate](BACnetContextTagParseWithBufferProducer((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_DATE)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dateValue' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetCalendarEntryDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetCalendarEntryDate")
	}

	// Create a partially initialized instance
	_child := &_BACnetCalendarEntryDate{
		_BACnetCalendarEntry: &_BACnetCalendarEntry{},
		DateValue:            dateValue,
	}
	_child._BACnetCalendarEntry._BACnetCalendarEntryChildRequirements = _child
	return _child, nil
}

func (m *_BACnetCalendarEntryDate) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetCalendarEntryDate) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetCalendarEntryDate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetCalendarEntryDate")
		}

		// Simple Field (dateValue)
		if pushErr := writeBuffer.PushContext("dateValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for dateValue")
		}
		_dateValueErr := writeBuffer.WriteSerializable(ctx, m.GetDateValue())
		if popErr := writeBuffer.PopContext("dateValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for dateValue")
		}
		if _dateValueErr != nil {
			return errors.Wrap(_dateValueErr, "Error serializing 'dateValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetCalendarEntryDate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetCalendarEntryDate")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetCalendarEntryDate) isBACnetCalendarEntryDate() bool {
	return true
}

func (m *_BACnetCalendarEntryDate) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
