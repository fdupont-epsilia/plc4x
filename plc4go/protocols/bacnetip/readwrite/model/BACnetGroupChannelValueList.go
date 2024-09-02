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

// BACnetGroupChannelValueList is the corresponding interface of BACnetGroupChannelValueList
type BACnetGroupChannelValueList interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfEventSummaries returns ListOfEventSummaries (property field)
	GetListOfEventSummaries() []BACnetEventSummary
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetGroupChannelValueList is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetGroupChannelValueList()
}

// _BACnetGroupChannelValueList is the data-structure of this message
type _BACnetGroupChannelValueList struct {
	OpeningTag           BACnetOpeningTag
	ListOfEventSummaries []BACnetEventSummary
	ClosingTag           BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetGroupChannelValueList = (*_BACnetGroupChannelValueList)(nil)

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
func CastBACnetGroupChannelValueList(structType any) BACnetGroupChannelValueList {
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

func (m *_BACnetGroupChannelValueList) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.ListOfEventSummaries) > 0 {
		for _, element := range m.ListOfEventSummaries {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetGroupChannelValueList) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetGroupChannelValueListParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetGroupChannelValueList, error) {
	return BACnetGroupChannelValueListParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetGroupChannelValueListParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetGroupChannelValueList, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetGroupChannelValueList, error) {
		return BACnetGroupChannelValueListParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetGroupChannelValueListParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetGroupChannelValueList, error) {
	v, err := (&_BACnetGroupChannelValueList{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetGroupChannelValueList) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetGroupChannelValueList BACnetGroupChannelValueList, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetGroupChannelValueList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetGroupChannelValueList")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	listOfEventSummaries, err := ReadTerminatedArrayField[BACnetEventSummary](ctx, "listOfEventSummaries", ReadComplex[BACnetEventSummary](BACnetEventSummaryParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfEventSummaries' field"))
	}
	m.ListOfEventSummaries = listOfEventSummaries

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetGroupChannelValueList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetGroupChannelValueList")
	}

	return m, nil
}

func (m *_BACnetGroupChannelValueList) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetGroupChannelValueList) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetGroupChannelValueList"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetGroupChannelValueList")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "listOfEventSummaries", m.GetListOfEventSummaries(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'listOfEventSummaries' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
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

func (m *_BACnetGroupChannelValueList) IsBACnetGroupChannelValueList() {}

func (m *_BACnetGroupChannelValueList) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
