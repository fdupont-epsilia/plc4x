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

// BACnetFaultParameterFaultCharacterStringListOfFaultValues is the corresponding interface of BACnetFaultParameterFaultCharacterStringListOfFaultValues
type BACnetFaultParameterFaultCharacterStringListOfFaultValues interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfFaultValues returns ListOfFaultValues (property field)
	GetListOfFaultValues() []BACnetApplicationTagCharacterString
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetFaultParameterFaultCharacterStringListOfFaultValuesExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultCharacterStringListOfFaultValues.
// This is useful for switch cases.
type BACnetFaultParameterFaultCharacterStringListOfFaultValuesExactly interface {
	BACnetFaultParameterFaultCharacterStringListOfFaultValues
	isBACnetFaultParameterFaultCharacterStringListOfFaultValues() bool
}

// _BACnetFaultParameterFaultCharacterStringListOfFaultValues is the data-structure of this message
type _BACnetFaultParameterFaultCharacterStringListOfFaultValues struct {
	OpeningTag        BACnetOpeningTag
	ListOfFaultValues []BACnetApplicationTagCharacterString
	ClosingTag        BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultCharacterStringListOfFaultValues) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetFaultParameterFaultCharacterStringListOfFaultValues) GetListOfFaultValues() []BACnetApplicationTagCharacterString {
	return m.ListOfFaultValues
}

func (m *_BACnetFaultParameterFaultCharacterStringListOfFaultValues) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultCharacterStringListOfFaultValues factory function for _BACnetFaultParameterFaultCharacterStringListOfFaultValues
func NewBACnetFaultParameterFaultCharacterStringListOfFaultValues(openingTag BACnetOpeningTag, listOfFaultValues []BACnetApplicationTagCharacterString, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetFaultParameterFaultCharacterStringListOfFaultValues {
	return &_BACnetFaultParameterFaultCharacterStringListOfFaultValues{OpeningTag: openingTag, ListOfFaultValues: listOfFaultValues, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultCharacterStringListOfFaultValues(structType any) BACnetFaultParameterFaultCharacterStringListOfFaultValues {
	if casted, ok := structType.(BACnetFaultParameterFaultCharacterStringListOfFaultValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultCharacterStringListOfFaultValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultCharacterStringListOfFaultValues) GetTypeName() string {
	return "BACnetFaultParameterFaultCharacterStringListOfFaultValues"
}

func (m *_BACnetFaultParameterFaultCharacterStringListOfFaultValues) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.ListOfFaultValues) > 0 {
		for _, element := range m.ListOfFaultValues {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultCharacterStringListOfFaultValues) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetFaultParameterFaultCharacterStringListOfFaultValuesParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetFaultParameterFaultCharacterStringListOfFaultValues, error) {
	return BACnetFaultParameterFaultCharacterStringListOfFaultValuesParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetFaultParameterFaultCharacterStringListOfFaultValuesParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetFaultParameterFaultCharacterStringListOfFaultValues, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetFaultParameterFaultCharacterStringListOfFaultValues, error) {
		return BACnetFaultParameterFaultCharacterStringListOfFaultValuesParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetFaultParameterFaultCharacterStringListOfFaultValuesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetFaultParameterFaultCharacterStringListOfFaultValues, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultCharacterStringListOfFaultValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultCharacterStringListOfFaultValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}

	listOfFaultValues, err := ReadTerminatedArrayField[BACnetApplicationTagCharacterString](ctx, "listOfFaultValues", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer(), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfFaultValues' field"))
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultCharacterStringListOfFaultValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultCharacterStringListOfFaultValues")
	}

	// Create the instance
	return &_BACnetFaultParameterFaultCharacterStringListOfFaultValues{
		TagNumber:         tagNumber,
		OpeningTag:        openingTag,
		ListOfFaultValues: listOfFaultValues,
		ClosingTag:        closingTag,
	}, nil
}

func (m *_BACnetFaultParameterFaultCharacterStringListOfFaultValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultCharacterStringListOfFaultValues) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultCharacterStringListOfFaultValues"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultCharacterStringListOfFaultValues")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(ctx, m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Array Field (listOfFaultValues)
	if pushErr := writeBuffer.PushContext("listOfFaultValues", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for listOfFaultValues")
	}
	for _curItem, _element := range m.GetListOfFaultValues() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetListOfFaultValues()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'listOfFaultValues' field")
		}
	}
	if popErr := writeBuffer.PopContext("listOfFaultValues", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for listOfFaultValues")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(ctx, m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultCharacterStringListOfFaultValues"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultCharacterStringListOfFaultValues")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetFaultParameterFaultCharacterStringListOfFaultValues) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetFaultParameterFaultCharacterStringListOfFaultValues) isBACnetFaultParameterFaultCharacterStringListOfFaultValues() bool {
	return true
}

func (m *_BACnetFaultParameterFaultCharacterStringListOfFaultValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
