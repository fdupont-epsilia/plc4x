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

// BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues is the corresponding interface of BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues
type BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfLifeSavetyAlarmValues returns ListOfLifeSavetyAlarmValues (property field)
	GetListOfLifeSavetyAlarmValues() []BACnetLifeSafetyStateTagged
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValuesExactly can be used when we want exactly this type and not a type which fulfills BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues.
// This is useful for switch cases.
type BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValuesExactly interface {
	BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues
	isBACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues() bool
}

// _BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues is the data-structure of this message
type _BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues struct {
	OpeningTag                  BACnetOpeningTag
	ListOfLifeSavetyAlarmValues []BACnetLifeSafetyStateTagged
	ClosingTag                  BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues = (*_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) GetListOfLifeSavetyAlarmValues() []BACnetLifeSafetyStateTagged {
	return m.ListOfLifeSavetyAlarmValues
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues factory function for _BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues
func NewBACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues(openingTag BACnetOpeningTag, listOfLifeSavetyAlarmValues []BACnetLifeSafetyStateTagged, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues {
	return &_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues{OpeningTag: openingTag, ListOfLifeSavetyAlarmValues: listOfLifeSavetyAlarmValues, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues(structType any) BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues {
	if casted, ok := structType.(BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) GetTypeName() string {
	return "BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues"
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.ListOfLifeSavetyAlarmValues) > 0 {
		for _, element := range m.ListOfLifeSavetyAlarmValues {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValuesParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues, error) {
	return BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValuesParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValuesParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues, error) {
		return BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValuesParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValuesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}

	listOfLifeSavetyAlarmValues, err := ReadTerminatedArrayField[BACnetLifeSafetyStateTagged](ctx, "listOfLifeSavetyAlarmValues", ReadComplex[BACnetLifeSafetyStateTagged](BACnetLifeSafetyStateTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfLifeSavetyAlarmValues' field"))
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues")
	}

	// Create the instance
	return &_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues{
		TagNumber:                   tagNumber,
		OpeningTag:                  openingTag,
		ListOfLifeSavetyAlarmValues: listOfLifeSavetyAlarmValues,
		ClosingTag:                  closingTag,
	}, nil
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "listOfLifeSavetyAlarmValues", m.GetListOfLifeSavetyAlarmValues(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'listOfLifeSavetyAlarmValues' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) isBACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues() bool {
	return true
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
