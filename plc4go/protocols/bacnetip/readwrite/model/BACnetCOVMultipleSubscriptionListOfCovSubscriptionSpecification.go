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

// BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification is the corresponding interface of BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification
type BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfCovSubscriptionSpecificationEntry returns ListOfCovSubscriptionSpecificationEntry (property field)
	GetListOfCovSubscriptionSpecificationEntry() []BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationExactly can be used when we want exactly this type and not a type which fulfills BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification.
// This is useful for switch cases.
type BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationExactly interface {
	BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification
	isBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification() bool
}

// _BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification is the data-structure of this message
type _BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification struct {
	OpeningTag                              BACnetOpeningTag
	ListOfCovSubscriptionSpecificationEntry []BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry
	ClosingTag                              BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification = (*_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) GetListOfCovSubscriptionSpecificationEntry() []BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry {
	return m.ListOfCovSubscriptionSpecificationEntry
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification factory function for _BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification
func NewBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification(openingTag BACnetOpeningTag, listOfCovSubscriptionSpecificationEntry []BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification {
	return &_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification{OpeningTag: openingTag, ListOfCovSubscriptionSpecificationEntry: listOfCovSubscriptionSpecificationEntry, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification(structType any) BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification {
	if casted, ok := structType.(BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) GetTypeName() string {
	return "BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification"
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.ListOfCovSubscriptionSpecificationEntry) > 0 {
		for _, element := range m.ListOfCovSubscriptionSpecificationEntry {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification, error) {
	return BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification, error) {
		return BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}

	listOfCovSubscriptionSpecificationEntry, err := ReadTerminatedArrayField[BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry](ctx, "listOfCovSubscriptionSpecificationEntry", ReadComplex[BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry](BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfCovSubscriptionSpecificationEntry' field"))
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification")
	}

	// Create the instance
	return &_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification{
		TagNumber:                               tagNumber,
		OpeningTag:                              openingTag,
		ListOfCovSubscriptionSpecificationEntry: listOfCovSubscriptionSpecificationEntry,
		ClosingTag:                              closingTag,
	}, nil
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "listOfCovSubscriptionSpecificationEntry", m.GetListOfCovSubscriptionSpecificationEntry(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'listOfCovSubscriptionSpecificationEntry' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) isBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification() bool {
	return true
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
