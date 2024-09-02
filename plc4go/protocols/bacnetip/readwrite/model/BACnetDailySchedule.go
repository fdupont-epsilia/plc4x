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

// BACnetDailySchedule is the corresponding interface of BACnetDailySchedule
type BACnetDailySchedule interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetDaySchedule returns DaySchedule (property field)
	GetDaySchedule() []BACnetTimeValue
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetDailySchedule is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetDailySchedule()
}

// _BACnetDailySchedule is the data-structure of this message
type _BACnetDailySchedule struct {
	OpeningTag  BACnetOpeningTag
	DaySchedule []BACnetTimeValue
	ClosingTag  BACnetClosingTag
}

var _ BACnetDailySchedule = (*_BACnetDailySchedule)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetDailySchedule) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetDailySchedule) GetDaySchedule() []BACnetTimeValue {
	return m.DaySchedule
}

func (m *_BACnetDailySchedule) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetDailySchedule factory function for _BACnetDailySchedule
func NewBACnetDailySchedule(openingTag BACnetOpeningTag, daySchedule []BACnetTimeValue, closingTag BACnetClosingTag) *_BACnetDailySchedule {
	return &_BACnetDailySchedule{OpeningTag: openingTag, DaySchedule: daySchedule, ClosingTag: closingTag}
}

// Deprecated: use the interface for direct cast
func CastBACnetDailySchedule(structType any) BACnetDailySchedule {
	if casted, ok := structType.(BACnetDailySchedule); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetDailySchedule); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetDailySchedule) GetTypeName() string {
	return "BACnetDailySchedule"
}

func (m *_BACnetDailySchedule) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.DaySchedule) > 0 {
		for _, element := range m.DaySchedule {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetDailySchedule) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetDailyScheduleParse(ctx context.Context, theBytes []byte) (BACnetDailySchedule, error) {
	return BACnetDailyScheduleParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetDailyScheduleParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDailySchedule, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDailySchedule, error) {
		return BACnetDailyScheduleParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetDailyScheduleParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDailySchedule, error) {
	v, err := (&_BACnetDailySchedule{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetDailySchedule) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetDailySchedule BACnetDailySchedule, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetDailySchedule"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetDailySchedule")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	daySchedule, err := ReadTerminatedArrayField[BACnetTimeValue](ctx, "daySchedule", ReadComplex[BACnetTimeValue](BACnetTimeValueParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, 0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'daySchedule' field"))
	}
	m.DaySchedule = daySchedule

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetDailySchedule"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetDailySchedule")
	}

	return m, nil
}

func (m *_BACnetDailySchedule) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetDailySchedule) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetDailySchedule"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetDailySchedule")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "daySchedule", m.GetDaySchedule(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'daySchedule' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetDailySchedule"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetDailySchedule")
	}
	return nil
}

func (m *_BACnetDailySchedule) IsBACnetDailySchedule() {}

func (m *_BACnetDailySchedule) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
