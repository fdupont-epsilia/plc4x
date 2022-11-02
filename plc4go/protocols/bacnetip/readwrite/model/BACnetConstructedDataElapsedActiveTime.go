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

// BACnetConstructedDataElapsedActiveTime is the corresponding interface of BACnetConstructedDataElapsedActiveTime
type BACnetConstructedDataElapsedActiveTime interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetElapsedActiveTime returns ElapsedActiveTime (property field)
	GetElapsedActiveTime() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataElapsedActiveTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataElapsedActiveTime.
// This is useful for switch cases.
type BACnetConstructedDataElapsedActiveTimeExactly interface {
	BACnetConstructedDataElapsedActiveTime
	isBACnetConstructedDataElapsedActiveTime() bool
}

// _BACnetConstructedDataElapsedActiveTime is the data-structure of this message
type _BACnetConstructedDataElapsedActiveTime struct {
	*_BACnetConstructedData
	ElapsedActiveTime BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataElapsedActiveTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataElapsedActiveTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ELAPSED_ACTIVE_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataElapsedActiveTime) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataElapsedActiveTime) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataElapsedActiveTime) GetElapsedActiveTime() BACnetApplicationTagUnsignedInteger {
	return m.ElapsedActiveTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataElapsedActiveTime) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetElapsedActiveTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataElapsedActiveTime factory function for _BACnetConstructedDataElapsedActiveTime
func NewBACnetConstructedDataElapsedActiveTime(elapsedActiveTime BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataElapsedActiveTime {
	_result := &_BACnetConstructedDataElapsedActiveTime{
		ElapsedActiveTime:      elapsedActiveTime,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataElapsedActiveTime(structType interface{}) BACnetConstructedDataElapsedActiveTime {
	if casted, ok := structType.(BACnetConstructedDataElapsedActiveTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataElapsedActiveTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataElapsedActiveTime) GetTypeName() string {
	return "BACnetConstructedDataElapsedActiveTime"
}

func (m *_BACnetConstructedDataElapsedActiveTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataElapsedActiveTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (elapsedActiveTime)
	lengthInBits += m.ElapsedActiveTime.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataElapsedActiveTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataElapsedActiveTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataElapsedActiveTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataElapsedActiveTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataElapsedActiveTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (elapsedActiveTime)
	if pullErr := readBuffer.PullContext("elapsedActiveTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for elapsedActiveTime")
	}
	_elapsedActiveTime, _elapsedActiveTimeErr := BACnetApplicationTagParse(readBuffer)
	if _elapsedActiveTimeErr != nil {
		return nil, errors.Wrap(_elapsedActiveTimeErr, "Error parsing 'elapsedActiveTime' field of BACnetConstructedDataElapsedActiveTime")
	}
	elapsedActiveTime := _elapsedActiveTime.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("elapsedActiveTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for elapsedActiveTime")
	}

	// Virtual field
	_actualValue := elapsedActiveTime
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataElapsedActiveTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataElapsedActiveTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataElapsedActiveTime{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ElapsedActiveTime: elapsedActiveTime,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataElapsedActiveTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataElapsedActiveTime) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataElapsedActiveTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataElapsedActiveTime")
		}

		// Simple Field (elapsedActiveTime)
		if pushErr := writeBuffer.PushContext("elapsedActiveTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for elapsedActiveTime")
		}
		_elapsedActiveTimeErr := writeBuffer.WriteSerializable(m.GetElapsedActiveTime())
		if popErr := writeBuffer.PopContext("elapsedActiveTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for elapsedActiveTime")
		}
		if _elapsedActiveTimeErr != nil {
			return errors.Wrap(_elapsedActiveTimeErr, "Error serializing 'elapsedActiveTime' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataElapsedActiveTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataElapsedActiveTime")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataElapsedActiveTime) isBACnetConstructedDataElapsedActiveTime() bool {
	return true
}

func (m *_BACnetConstructedDataElapsedActiveTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}