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

// BACnetConstructedDataEventLogLogBuffer is the corresponding interface of BACnetConstructedDataEventLogLogBuffer
type BACnetConstructedDataEventLogLogBuffer interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFloorText returns FloorText (property field)
	GetFloorText() []BACnetEventLogRecord
}

// BACnetConstructedDataEventLogLogBufferExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataEventLogLogBuffer.
// This is useful for switch cases.
type BACnetConstructedDataEventLogLogBufferExactly interface {
	BACnetConstructedDataEventLogLogBuffer
	isBACnetConstructedDataEventLogLogBuffer() bool
}

// _BACnetConstructedDataEventLogLogBuffer is the data-structure of this message
type _BACnetConstructedDataEventLogLogBuffer struct {
	*_BACnetConstructedData
	FloorText []BACnetEventLogRecord
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEventLogLogBuffer) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_EVENT_LOG
}

func (m *_BACnetConstructedDataEventLogLogBuffer) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOG_BUFFER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEventLogLogBuffer) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataEventLogLogBuffer) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEventLogLogBuffer) GetFloorText() []BACnetEventLogRecord {
	return m.FloorText
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEventLogLogBuffer factory function for _BACnetConstructedDataEventLogLogBuffer
func NewBACnetConstructedDataEventLogLogBuffer(floorText []BACnetEventLogRecord, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEventLogLogBuffer {
	_result := &_BACnetConstructedDataEventLogLogBuffer{
		FloorText:              floorText,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEventLogLogBuffer(structType interface{}) BACnetConstructedDataEventLogLogBuffer {
	if casted, ok := structType.(BACnetConstructedDataEventLogLogBuffer); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEventLogLogBuffer); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEventLogLogBuffer) GetTypeName() string {
	return "BACnetConstructedDataEventLogLogBuffer"
}

func (m *_BACnetConstructedDataEventLogLogBuffer) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataEventLogLogBuffer) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.FloorText) > 0 {
		for _, element := range m.FloorText {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataEventLogLogBuffer) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataEventLogLogBufferParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataEventLogLogBuffer, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEventLogLogBuffer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEventLogLogBuffer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (floorText)
	if pullErr := readBuffer.PullContext("floorText", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for floorText")
	}
	// Terminated array
	var floorText []BACnetEventLogRecord
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetEventLogRecordParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'floorText' field of BACnetConstructedDataEventLogLogBuffer")
			}
			floorText = append(floorText, _item.(BACnetEventLogRecord))

		}
	}
	if closeErr := readBuffer.CloseContext("floorText", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for floorText")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEventLogLogBuffer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEventLogLogBuffer")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataEventLogLogBuffer{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		FloorText: floorText,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataEventLogLogBuffer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEventLogLogBuffer) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEventLogLogBuffer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEventLogLogBuffer")
		}

		// Array Field (floorText)
		if pushErr := writeBuffer.PushContext("floorText", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for floorText")
		}
		for _, _element := range m.GetFloorText() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'floorText' field")
			}
		}
		if popErr := writeBuffer.PopContext("floorText", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for floorText")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEventLogLogBuffer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEventLogLogBuffer")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEventLogLogBuffer) isBACnetConstructedDataEventLogLogBuffer() bool {
	return true
}

func (m *_BACnetConstructedDataEventLogLogBuffer) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}