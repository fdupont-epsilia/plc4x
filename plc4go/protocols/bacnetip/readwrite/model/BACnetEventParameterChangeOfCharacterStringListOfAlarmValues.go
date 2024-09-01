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

// BACnetEventParameterChangeOfCharacterStringListOfAlarmValues is the corresponding interface of BACnetEventParameterChangeOfCharacterStringListOfAlarmValues
type BACnetEventParameterChangeOfCharacterStringListOfAlarmValues interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfAlarmValues returns ListOfAlarmValues (property field)
	GetListOfAlarmValues() []BACnetApplicationTagCharacterString
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesExactly can be used when we want exactly this type and not a type which fulfills BACnetEventParameterChangeOfCharacterStringListOfAlarmValues.
// This is useful for switch cases.
type BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesExactly interface {
	BACnetEventParameterChangeOfCharacterStringListOfAlarmValues
	isBACnetEventParameterChangeOfCharacterStringListOfAlarmValues() bool
}

// _BACnetEventParameterChangeOfCharacterStringListOfAlarmValues is the data-structure of this message
type _BACnetEventParameterChangeOfCharacterStringListOfAlarmValues struct {
	OpeningTag        BACnetOpeningTag
	ListOfAlarmValues []BACnetApplicationTagCharacterString
	ClosingTag        BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetEventParameterChangeOfCharacterStringListOfAlarmValues = (*_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) GetListOfAlarmValues() []BACnetApplicationTagCharacterString {
	return m.ListOfAlarmValues
}

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterChangeOfCharacterStringListOfAlarmValues factory function for _BACnetEventParameterChangeOfCharacterStringListOfAlarmValues
func NewBACnetEventParameterChangeOfCharacterStringListOfAlarmValues(openingTag BACnetOpeningTag, listOfAlarmValues []BACnetApplicationTagCharacterString, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues {
	return &_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues{OpeningTag: openingTag, ListOfAlarmValues: listOfAlarmValues, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterChangeOfCharacterStringListOfAlarmValues(structType any) BACnetEventParameterChangeOfCharacterStringListOfAlarmValues {
	if casted, ok := structType.(BACnetEventParameterChangeOfCharacterStringListOfAlarmValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfCharacterStringListOfAlarmValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) GetTypeName() string {
	return "BACnetEventParameterChangeOfCharacterStringListOfAlarmValues"
}

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.ListOfAlarmValues) > 0 {
		for _, element := range m.ListOfAlarmValues {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetEventParameterChangeOfCharacterStringListOfAlarmValues, error) {
	return BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterChangeOfCharacterStringListOfAlarmValues, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterChangeOfCharacterStringListOfAlarmValues, error) {
		return BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetEventParameterChangeOfCharacterStringListOfAlarmValues, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfCharacterStringListOfAlarmValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterChangeOfCharacterStringListOfAlarmValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}

	listOfAlarmValues, err := ReadTerminatedArrayField[BACnetApplicationTagCharacterString](ctx, "listOfAlarmValues", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfAlarmValues' field"))
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfCharacterStringListOfAlarmValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterChangeOfCharacterStringListOfAlarmValues")
	}

	// Create the instance
	return &_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues{
		TagNumber:         tagNumber,
		OpeningTag:        openingTag,
		ListOfAlarmValues: listOfAlarmValues,
		ClosingTag:        closingTag,
	}, nil
}

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfCharacterStringListOfAlarmValues"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterChangeOfCharacterStringListOfAlarmValues")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "listOfAlarmValues", m.GetListOfAlarmValues(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'listOfAlarmValues' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfCharacterStringListOfAlarmValues"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventParameterChangeOfCharacterStringListOfAlarmValues")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) isBACnetEventParameterChangeOfCharacterStringListOfAlarmValues() bool {
	return true
}

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
