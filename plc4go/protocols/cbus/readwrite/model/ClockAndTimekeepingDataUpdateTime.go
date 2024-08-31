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

// ClockAndTimekeepingDataUpdateTime is the corresponding interface of ClockAndTimekeepingDataUpdateTime
type ClockAndTimekeepingDataUpdateTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ClockAndTimekeepingData
	// GetHours returns Hours (property field)
	GetHours() uint8
	// GetMinute returns Minute (property field)
	GetMinute() uint8
	// GetSecond returns Second (property field)
	GetSecond() uint8
	// GetDaylightSaving returns DaylightSaving (property field)
	GetDaylightSaving() byte
	// GetIsNoDaylightSavings returns IsNoDaylightSavings (virtual field)
	GetIsNoDaylightSavings() bool
	// GetIsAdvancedBy1Hour returns IsAdvancedBy1Hour (virtual field)
	GetIsAdvancedBy1Hour() bool
	// GetIsReserved returns IsReserved (virtual field)
	GetIsReserved() bool
	// GetIsUnknown returns IsUnknown (virtual field)
	GetIsUnknown() bool
}

// ClockAndTimekeepingDataUpdateTimeExactly can be used when we want exactly this type and not a type which fulfills ClockAndTimekeepingDataUpdateTime.
// This is useful for switch cases.
type ClockAndTimekeepingDataUpdateTimeExactly interface {
	ClockAndTimekeepingDataUpdateTime
	isClockAndTimekeepingDataUpdateTime() bool
}

// _ClockAndTimekeepingDataUpdateTime is the data-structure of this message
type _ClockAndTimekeepingDataUpdateTime struct {
	*_ClockAndTimekeepingData
	Hours          uint8
	Minute         uint8
	Second         uint8
	DaylightSaving byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ClockAndTimekeepingDataUpdateTime) InitializeParent(parent ClockAndTimekeepingData, commandTypeContainer ClockAndTimekeepingCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_ClockAndTimekeepingDataUpdateTime) GetParent() ClockAndTimekeepingData {
	return m._ClockAndTimekeepingData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ClockAndTimekeepingDataUpdateTime) GetHours() uint8 {
	return m.Hours
}

func (m *_ClockAndTimekeepingDataUpdateTime) GetMinute() uint8 {
	return m.Minute
}

func (m *_ClockAndTimekeepingDataUpdateTime) GetSecond() uint8 {
	return m.Second
}

func (m *_ClockAndTimekeepingDataUpdateTime) GetDaylightSaving() byte {
	return m.DaylightSaving
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_ClockAndTimekeepingDataUpdateTime) GetIsNoDaylightSavings() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetDaylightSaving()) == (0x00)))
}

func (m *_ClockAndTimekeepingDataUpdateTime) GetIsAdvancedBy1Hour() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetDaylightSaving()) == (0x01)))
}

func (m *_ClockAndTimekeepingDataUpdateTime) GetIsReserved() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool(bool((m.GetDaylightSaving()) > (0x01))) && bool(bool((m.GetDaylightSaving()) <= (0xFE))))
}

func (m *_ClockAndTimekeepingDataUpdateTime) GetIsUnknown() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetDaylightSaving()) > (0xFE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewClockAndTimekeepingDataUpdateTime factory function for _ClockAndTimekeepingDataUpdateTime
func NewClockAndTimekeepingDataUpdateTime(hours uint8, minute uint8, second uint8, daylightSaving byte, commandTypeContainer ClockAndTimekeepingCommandTypeContainer, argument byte) *_ClockAndTimekeepingDataUpdateTime {
	_result := &_ClockAndTimekeepingDataUpdateTime{
		Hours:                    hours,
		Minute:                   minute,
		Second:                   second,
		DaylightSaving:           daylightSaving,
		_ClockAndTimekeepingData: NewClockAndTimekeepingData(commandTypeContainer, argument),
	}
	_result._ClockAndTimekeepingData._ClockAndTimekeepingDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastClockAndTimekeepingDataUpdateTime(structType any) ClockAndTimekeepingDataUpdateTime {
	if casted, ok := structType.(ClockAndTimekeepingDataUpdateTime); ok {
		return casted
	}
	if casted, ok := structType.(*ClockAndTimekeepingDataUpdateTime); ok {
		return *casted
	}
	return nil
}

func (m *_ClockAndTimekeepingDataUpdateTime) GetTypeName() string {
	return "ClockAndTimekeepingDataUpdateTime"
}

func (m *_ClockAndTimekeepingDataUpdateTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (hours)
	lengthInBits += 8

	// Simple field (minute)
	lengthInBits += 8

	// Simple field (second)
	lengthInBits += 8

	// Simple field (daylightSaving)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_ClockAndTimekeepingDataUpdateTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ClockAndTimekeepingDataUpdateTimeParse(ctx context.Context, theBytes []byte) (ClockAndTimekeepingDataUpdateTime, error) {
	return ClockAndTimekeepingDataUpdateTimeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ClockAndTimekeepingDataUpdateTimeParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (ClockAndTimekeepingDataUpdateTime, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ClockAndTimekeepingDataUpdateTime, error) {
		return ClockAndTimekeepingDataUpdateTimeParseWithBuffer(ctx, readBuffer)
	}
}

func ClockAndTimekeepingDataUpdateTimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ClockAndTimekeepingDataUpdateTime, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ClockAndTimekeepingDataUpdateTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ClockAndTimekeepingDataUpdateTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	hours, err := ReadSimpleField(ctx, "hours", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hours' field"))
	}

	minute, err := ReadSimpleField(ctx, "minute", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'minute' field"))
	}

	second, err := ReadSimpleField(ctx, "second", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'second' field"))
	}

	daylightSaving, err := ReadSimpleField(ctx, "daylightSaving", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'daylightSaving' field"))
	}

	isNoDaylightSavings, err := ReadVirtualField[bool](ctx, "isNoDaylightSavings", (*bool)(nil), bool((daylightSaving) == (0x00)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isNoDaylightSavings' field"))
	}

	isAdvancedBy1Hour, err := ReadVirtualField[bool](ctx, "isAdvancedBy1Hour", (*bool)(nil), bool((daylightSaving) == (0x01)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isAdvancedBy1Hour' field"))
	}

	isReserved, err := ReadVirtualField[bool](ctx, "isReserved", (*bool)(nil), bool(bool((daylightSaving) > (0x01))) && bool(bool((daylightSaving) <= (0xFE))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isReserved' field"))
	}

	isUnknown, err := ReadVirtualField[bool](ctx, "isUnknown", (*bool)(nil), bool((daylightSaving) > (0xFE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isUnknown' field"))
	}

	if closeErr := readBuffer.CloseContext("ClockAndTimekeepingDataUpdateTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ClockAndTimekeepingDataUpdateTime")
	}

	// Create a partially initialized instance
	_child := &_ClockAndTimekeepingDataUpdateTime{
		_ClockAndTimekeepingData: &_ClockAndTimekeepingData{},
		Hours:                    hours,
		Minute:                   minute,
		Second:                   second,
		DaylightSaving:           daylightSaving,
	}
	_child._ClockAndTimekeepingData._ClockAndTimekeepingDataChildRequirements = _child
	return _child, nil
}

func (m *_ClockAndTimekeepingDataUpdateTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ClockAndTimekeepingDataUpdateTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ClockAndTimekeepingDataUpdateTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ClockAndTimekeepingDataUpdateTime")
		}

		// Simple Field (hours)
		hours := uint8(m.GetHours())
		_hoursErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("hours", 8, uint8((hours)))
		if _hoursErr != nil {
			return errors.Wrap(_hoursErr, "Error serializing 'hours' field")
		}

		// Simple Field (minute)
		minute := uint8(m.GetMinute())
		_minuteErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("minute", 8, uint8((minute)))
		if _minuteErr != nil {
			return errors.Wrap(_minuteErr, "Error serializing 'minute' field")
		}

		// Simple Field (second)
		second := uint8(m.GetSecond())
		_secondErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("second", 8, uint8((second)))
		if _secondErr != nil {
			return errors.Wrap(_secondErr, "Error serializing 'second' field")
		}

		// Simple Field (daylightSaving)
		daylightSaving := byte(m.GetDaylightSaving())
		_daylightSavingErr := /*TODO: migrate me*/ writeBuffer.WriteByte("daylightSaving", (daylightSaving))
		if _daylightSavingErr != nil {
			return errors.Wrap(_daylightSavingErr, "Error serializing 'daylightSaving' field")
		}
		// Virtual field
		isNoDaylightSavings := m.GetIsNoDaylightSavings()
		_ = isNoDaylightSavings
		if _isNoDaylightSavingsErr := writeBuffer.WriteVirtual(ctx, "isNoDaylightSavings", m.GetIsNoDaylightSavings()); _isNoDaylightSavingsErr != nil {
			return errors.Wrap(_isNoDaylightSavingsErr, "Error serializing 'isNoDaylightSavings' field")
		}
		// Virtual field
		isAdvancedBy1Hour := m.GetIsAdvancedBy1Hour()
		_ = isAdvancedBy1Hour
		if _isAdvancedBy1HourErr := writeBuffer.WriteVirtual(ctx, "isAdvancedBy1Hour", m.GetIsAdvancedBy1Hour()); _isAdvancedBy1HourErr != nil {
			return errors.Wrap(_isAdvancedBy1HourErr, "Error serializing 'isAdvancedBy1Hour' field")
		}
		// Virtual field
		isReserved := m.GetIsReserved()
		_ = isReserved
		if _isReservedErr := writeBuffer.WriteVirtual(ctx, "isReserved", m.GetIsReserved()); _isReservedErr != nil {
			return errors.Wrap(_isReservedErr, "Error serializing 'isReserved' field")
		}
		// Virtual field
		isUnknown := m.GetIsUnknown()
		_ = isUnknown
		if _isUnknownErr := writeBuffer.WriteVirtual(ctx, "isUnknown", m.GetIsUnknown()); _isUnknownErr != nil {
			return errors.Wrap(_isUnknownErr, "Error serializing 'isUnknown' field")
		}

		if popErr := writeBuffer.PopContext("ClockAndTimekeepingDataUpdateTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ClockAndTimekeepingDataUpdateTime")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ClockAndTimekeepingDataUpdateTime) isClockAndTimekeepingDataUpdateTime() bool {
	return true
}

func (m *_ClockAndTimekeepingDataUpdateTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
