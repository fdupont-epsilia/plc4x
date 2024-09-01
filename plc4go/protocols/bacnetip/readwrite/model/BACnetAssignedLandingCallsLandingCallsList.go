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

// BACnetAssignedLandingCallsLandingCallsList is the corresponding interface of BACnetAssignedLandingCallsLandingCallsList
type BACnetAssignedLandingCallsLandingCallsList interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetLandingCalls returns LandingCalls (property field)
	GetLandingCalls() []BACnetAssignedLandingCallsLandingCallsListEntry
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetAssignedLandingCallsLandingCallsListExactly can be used when we want exactly this type and not a type which fulfills BACnetAssignedLandingCallsLandingCallsList.
// This is useful for switch cases.
type BACnetAssignedLandingCallsLandingCallsListExactly interface {
	BACnetAssignedLandingCallsLandingCallsList
	isBACnetAssignedLandingCallsLandingCallsList() bool
}

// _BACnetAssignedLandingCallsLandingCallsList is the data-structure of this message
type _BACnetAssignedLandingCallsLandingCallsList struct {
	OpeningTag   BACnetOpeningTag
	LandingCalls []BACnetAssignedLandingCallsLandingCallsListEntry
	ClosingTag   BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetAssignedLandingCallsLandingCallsList = (*_BACnetAssignedLandingCallsLandingCallsList)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAssignedLandingCallsLandingCallsList) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetAssignedLandingCallsLandingCallsList) GetLandingCalls() []BACnetAssignedLandingCallsLandingCallsListEntry {
	return m.LandingCalls
}

func (m *_BACnetAssignedLandingCallsLandingCallsList) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetAssignedLandingCallsLandingCallsList factory function for _BACnetAssignedLandingCallsLandingCallsList
func NewBACnetAssignedLandingCallsLandingCallsList(openingTag BACnetOpeningTag, landingCalls []BACnetAssignedLandingCallsLandingCallsListEntry, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetAssignedLandingCallsLandingCallsList {
	return &_BACnetAssignedLandingCallsLandingCallsList{OpeningTag: openingTag, LandingCalls: landingCalls, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetAssignedLandingCallsLandingCallsList(structType any) BACnetAssignedLandingCallsLandingCallsList {
	if casted, ok := structType.(BACnetAssignedLandingCallsLandingCallsList); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAssignedLandingCallsLandingCallsList); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAssignedLandingCallsLandingCallsList) GetTypeName() string {
	return "BACnetAssignedLandingCallsLandingCallsList"
}

func (m *_BACnetAssignedLandingCallsLandingCallsList) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.LandingCalls) > 0 {
		for _, element := range m.LandingCalls {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetAssignedLandingCallsLandingCallsList) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAssignedLandingCallsLandingCallsListParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetAssignedLandingCallsLandingCallsList, error) {
	return BACnetAssignedLandingCallsLandingCallsListParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetAssignedLandingCallsLandingCallsListParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAssignedLandingCallsLandingCallsList, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAssignedLandingCallsLandingCallsList, error) {
		return BACnetAssignedLandingCallsLandingCallsListParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetAssignedLandingCallsLandingCallsListParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetAssignedLandingCallsLandingCallsList, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAssignedLandingCallsLandingCallsList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAssignedLandingCallsLandingCallsList")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}

	landingCalls, err := ReadTerminatedArrayField[BACnetAssignedLandingCallsLandingCallsListEntry](ctx, "landingCalls", ReadComplex[BACnetAssignedLandingCallsLandingCallsListEntry](BACnetAssignedLandingCallsLandingCallsListEntryParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'landingCalls' field"))
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetAssignedLandingCallsLandingCallsList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAssignedLandingCallsLandingCallsList")
	}

	// Create the instance
	return &_BACnetAssignedLandingCallsLandingCallsList{
		TagNumber:    tagNumber,
		OpeningTag:   openingTag,
		LandingCalls: landingCalls,
		ClosingTag:   closingTag,
	}, nil
}

func (m *_BACnetAssignedLandingCallsLandingCallsList) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAssignedLandingCallsLandingCallsList) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAssignedLandingCallsLandingCallsList"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAssignedLandingCallsLandingCallsList")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "landingCalls", m.GetLandingCalls(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'landingCalls' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAssignedLandingCallsLandingCallsList"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAssignedLandingCallsLandingCallsList")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetAssignedLandingCallsLandingCallsList) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetAssignedLandingCallsLandingCallsList) isBACnetAssignedLandingCallsLandingCallsList() bool {
	return true
}

func (m *_BACnetAssignedLandingCallsLandingCallsList) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
