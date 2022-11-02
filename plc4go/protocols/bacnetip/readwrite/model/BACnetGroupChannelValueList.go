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

// BACnetGroupChannelValueList is the corresponding interface of BACnetGroupChannelValueList
type BACnetGroupChannelValueList interface {
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfEventSummaries returns ListOfEventSummaries (property field)
	GetListOfEventSummaries() []BACnetEventSummary
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetGroupChannelValueListExactly can be used when we want exactly this type and not a type which fulfills BACnetGroupChannelValueList.
// This is useful for switch cases.
type BACnetGroupChannelValueListExactly interface {
	BACnetGroupChannelValueList
	isBACnetGroupChannelValueList() bool
}

// _BACnetGroupChannelValueList is the data-structure of this message
type _BACnetGroupChannelValueList struct {
	OpeningTag           BACnetOpeningTag
	ListOfEventSummaries []BACnetEventSummary
	ClosingTag           BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetGroupChannelValueList) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetGroupChannelValueList) GetListOfEventSummaries() []BACnetEventSummary {
	return m.ListOfEventSummaries
}

func (m *_BACnetGroupChannelValueList) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetGroupChannelValueList factory function for _BACnetGroupChannelValueList
func NewBACnetGroupChannelValueList(openingTag BACnetOpeningTag, listOfEventSummaries []BACnetEventSummary, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetGroupChannelValueList {
	return &_BACnetGroupChannelValueList{OpeningTag: openingTag, ListOfEventSummaries: listOfEventSummaries, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetGroupChannelValueList(structType interface{}) BACnetGroupChannelValueList {
	if casted, ok := structType.(BACnetGroupChannelValueList); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetGroupChannelValueList); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetGroupChannelValueList) GetTypeName() string {
	return "BACnetGroupChannelValueList"
}

func (m *_BACnetGroupChannelValueList) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetGroupChannelValueList) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Array field
	if len(m.ListOfEventSummaries) > 0 {
		for _, element := range m.ListOfEventSummaries {
			lengthInBits += element.GetLengthInBits()
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetGroupChannelValueList) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetGroupChannelValueListParse(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetGroupChannelValueList, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetGroupChannelValueList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetGroupChannelValueList")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetGroupChannelValueList")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Array field (listOfEventSummaries)
	if pullErr := readBuffer.PullContext("listOfEventSummaries", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for listOfEventSummaries")
	}
	// Terminated array
	var listOfEventSummaries []BACnetEventSummary
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetEventSummaryParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'listOfEventSummaries' field of BACnetGroupChannelValueList")
			}
			listOfEventSummaries = append(listOfEventSummaries, _item.(BACnetEventSummary))

		}
	}
	if closeErr := readBuffer.CloseContext("listOfEventSummaries", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for listOfEventSummaries")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetGroupChannelValueList")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetGroupChannelValueList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetGroupChannelValueList")
	}

	// Create the instance
	return &_BACnetGroupChannelValueList{
		TagNumber:            tagNumber,
		OpeningTag:           openingTag,
		ListOfEventSummaries: listOfEventSummaries,
		ClosingTag:           closingTag,
	}, nil
}

func (m *_BACnetGroupChannelValueList) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetGroupChannelValueList) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetGroupChannelValueList"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetGroupChannelValueList")
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

	// Array Field (listOfEventSummaries)
	if pushErr := writeBuffer.PushContext("listOfEventSummaries", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for listOfEventSummaries")
	}
	for _, _element := range m.GetListOfEventSummaries() {
		_elementErr := writeBuffer.WriteSerializable(_element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'listOfEventSummaries' field")
		}
	}
	if popErr := writeBuffer.PopContext("listOfEventSummaries", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for listOfEventSummaries")
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

	if popErr := writeBuffer.PopContext("BACnetGroupChannelValueList"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetGroupChannelValueList")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetGroupChannelValueList) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetGroupChannelValueList) isBACnetGroupChannelValueList() bool {
	return true
}

func (m *_BACnetGroupChannelValueList) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}