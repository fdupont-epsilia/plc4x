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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetEventParameterSignedOutOfRange is the corresponding interface of BACnetEventParameterSignedOutOfRange
type BACnetEventParameterSignedOutOfRange interface {
	utils.LengthAware
	utils.Serializable
	BACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetTimeDelay returns TimeDelay (property field)
	GetTimeDelay() BACnetContextTagUnsignedInteger
	// GetLowLimit returns LowLimit (property field)
	GetLowLimit() BACnetContextTagSignedInteger
	// GetHighLimit returns HighLimit (property field)
	GetHighLimit() BACnetContextTagSignedInteger
	// GetDeadband returns Deadband (property field)
	GetDeadband() BACnetContextTagUnsignedInteger
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetEventParameterSignedOutOfRangeExactly can be used when we want exactly this type and not a type which fulfills BACnetEventParameterSignedOutOfRange.
// This is useful for switch cases.
type BACnetEventParameterSignedOutOfRangeExactly interface {
	BACnetEventParameterSignedOutOfRange
	isBACnetEventParameterSignedOutOfRange() bool
}

// _BACnetEventParameterSignedOutOfRange is the data-structure of this message
type _BACnetEventParameterSignedOutOfRange struct {
	*_BACnetEventParameter
	OpeningTag BACnetOpeningTag
	TimeDelay  BACnetContextTagUnsignedInteger
	LowLimit   BACnetContextTagSignedInteger
	HighLimit  BACnetContextTagSignedInteger
	Deadband   BACnetContextTagUnsignedInteger
	ClosingTag BACnetClosingTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetEventParameterSignedOutOfRange) InitializeParent(parent BACnetEventParameter, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetEventParameterSignedOutOfRange) GetParent() BACnetEventParameter {
	return m._BACnetEventParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterSignedOutOfRange) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterSignedOutOfRange) GetTimeDelay() BACnetContextTagUnsignedInteger {
	return m.TimeDelay
}

func (m *_BACnetEventParameterSignedOutOfRange) GetLowLimit() BACnetContextTagSignedInteger {
	return m.LowLimit
}

func (m *_BACnetEventParameterSignedOutOfRange) GetHighLimit() BACnetContextTagSignedInteger {
	return m.HighLimit
}

func (m *_BACnetEventParameterSignedOutOfRange) GetDeadband() BACnetContextTagUnsignedInteger {
	return m.Deadband
}

func (m *_BACnetEventParameterSignedOutOfRange) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterSignedOutOfRange factory function for _BACnetEventParameterSignedOutOfRange
func NewBACnetEventParameterSignedOutOfRange(openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, lowLimit BACnetContextTagSignedInteger, highLimit BACnetContextTagSignedInteger, deadband BACnetContextTagUnsignedInteger, closingTag BACnetClosingTag, peekedTagHeader BACnetTagHeader) *_BACnetEventParameterSignedOutOfRange {
	_result := &_BACnetEventParameterSignedOutOfRange{
		OpeningTag:            openingTag,
		TimeDelay:             timeDelay,
		LowLimit:              lowLimit,
		HighLimit:             highLimit,
		Deadband:              deadband,
		ClosingTag:            closingTag,
		_BACnetEventParameter: NewBACnetEventParameter(peekedTagHeader),
	}
	_result._BACnetEventParameter._BACnetEventParameterChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterSignedOutOfRange(structType interface{}) BACnetEventParameterSignedOutOfRange {
	if casted, ok := structType.(BACnetEventParameterSignedOutOfRange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterSignedOutOfRange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterSignedOutOfRange) GetTypeName() string {
	return "BACnetEventParameterSignedOutOfRange"
}

func (m *_BACnetEventParameterSignedOutOfRange) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetEventParameterSignedOutOfRange) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (timeDelay)
	lengthInBits += m.TimeDelay.GetLengthInBits()

	// Simple field (lowLimit)
	lengthInBits += m.LowLimit.GetLengthInBits()

	// Simple field (highLimit)
	lengthInBits += m.HighLimit.GetLengthInBits()

	// Simple field (deadband)
	lengthInBits += m.Deadband.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetEventParameterSignedOutOfRange) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventParameterSignedOutOfRangeParse(readBuffer utils.ReadBuffer) (BACnetEventParameterSignedOutOfRange, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterSignedOutOfRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterSignedOutOfRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(uint8(15)))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetEventParameterSignedOutOfRange")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (timeDelay)
	if pullErr := readBuffer.PullContext("timeDelay"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeDelay")
	}
	_timeDelay, _timeDelayErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _timeDelayErr != nil {
		return nil, errors.Wrap(_timeDelayErr, "Error parsing 'timeDelay' field of BACnetEventParameterSignedOutOfRange")
	}
	timeDelay := _timeDelay.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("timeDelay"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeDelay")
	}

	// Simple Field (lowLimit)
	if pullErr := readBuffer.PullContext("lowLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lowLimit")
	}
	_lowLimit, _lowLimitErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_SIGNED_INTEGER))
	if _lowLimitErr != nil {
		return nil, errors.Wrap(_lowLimitErr, "Error parsing 'lowLimit' field of BACnetEventParameterSignedOutOfRange")
	}
	lowLimit := _lowLimit.(BACnetContextTagSignedInteger)
	if closeErr := readBuffer.CloseContext("lowLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lowLimit")
	}

	// Simple Field (highLimit)
	if pullErr := readBuffer.PullContext("highLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for highLimit")
	}
	_highLimit, _highLimitErr := BACnetContextTagParse(readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_SIGNED_INTEGER))
	if _highLimitErr != nil {
		return nil, errors.Wrap(_highLimitErr, "Error parsing 'highLimit' field of BACnetEventParameterSignedOutOfRange")
	}
	highLimit := _highLimit.(BACnetContextTagSignedInteger)
	if closeErr := readBuffer.CloseContext("highLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for highLimit")
	}

	// Simple Field (deadband)
	if pullErr := readBuffer.PullContext("deadband"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for deadband")
	}
	_deadband, _deadbandErr := BACnetContextTagParse(readBuffer, uint8(uint8(3)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _deadbandErr != nil {
		return nil, errors.Wrap(_deadbandErr, "Error parsing 'deadband' field of BACnetEventParameterSignedOutOfRange")
	}
	deadband := _deadband.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("deadband"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for deadband")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(uint8(15)))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetEventParameterSignedOutOfRange")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterSignedOutOfRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterSignedOutOfRange")
	}

	// Create a partially initialized instance
	_child := &_BACnetEventParameterSignedOutOfRange{
		_BACnetEventParameter: &_BACnetEventParameter{},
		OpeningTag:            openingTag,
		TimeDelay:             timeDelay,
		LowLimit:              lowLimit,
		HighLimit:             highLimit,
		Deadband:              deadband,
		ClosingTag:            closingTag,
	}
	_child._BACnetEventParameter._BACnetEventParameterChildRequirements = _child
	return _child, nil
}

func (m *_BACnetEventParameterSignedOutOfRange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterSignedOutOfRange) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterSignedOutOfRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterSignedOutOfRange")
		}

		// Simple Field (openingTag)
		if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for openingTag")
		}
		_openingTagErr := writeBuffer.WriteSerializable(m.GetOpeningTag())
		if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for openingTag")
		}
		if _openingTagErr != nil {
			return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
		}

		// Simple Field (timeDelay)
		if pushErr := writeBuffer.PushContext("timeDelay"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timeDelay")
		}
		_timeDelayErr := writeBuffer.WriteSerializable(m.GetTimeDelay())
		if popErr := writeBuffer.PopContext("timeDelay"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timeDelay")
		}
		if _timeDelayErr != nil {
			return errors.Wrap(_timeDelayErr, "Error serializing 'timeDelay' field")
		}

		// Simple Field (lowLimit)
		if pushErr := writeBuffer.PushContext("lowLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lowLimit")
		}
		_lowLimitErr := writeBuffer.WriteSerializable(m.GetLowLimit())
		if popErr := writeBuffer.PopContext("lowLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lowLimit")
		}
		if _lowLimitErr != nil {
			return errors.Wrap(_lowLimitErr, "Error serializing 'lowLimit' field")
		}

		// Simple Field (highLimit)
		if pushErr := writeBuffer.PushContext("highLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for highLimit")
		}
		_highLimitErr := writeBuffer.WriteSerializable(m.GetHighLimit())
		if popErr := writeBuffer.PopContext("highLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for highLimit")
		}
		if _highLimitErr != nil {
			return errors.Wrap(_highLimitErr, "Error serializing 'highLimit' field")
		}

		// Simple Field (deadband)
		if pushErr := writeBuffer.PushContext("deadband"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for deadband")
		}
		_deadbandErr := writeBuffer.WriteSerializable(m.GetDeadband())
		if popErr := writeBuffer.PopContext("deadband"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for deadband")
		}
		if _deadbandErr != nil {
			return errors.Wrap(_deadbandErr, "Error serializing 'deadband' field")
		}

		// Simple Field (closingTag)
		if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for closingTag")
		}
		_closingTagErr := writeBuffer.WriteSerializable(m.GetClosingTag())
		if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for closingTag")
		}
		if _closingTagErr != nil {
			return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventParameterSignedOutOfRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterSignedOutOfRange")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetEventParameterSignedOutOfRange) isBACnetEventParameterSignedOutOfRange() bool {
	return true
}

func (m *_BACnetEventParameterSignedOutOfRange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}