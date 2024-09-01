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

// BACnetEventParameterAccessEvent is the corresponding interface of BACnetEventParameterAccessEvent
type BACnetEventParameterAccessEvent interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfAccessEvents returns ListOfAccessEvents (property field)
	GetListOfAccessEvents() BACnetEventParameterAccessEventListOfAccessEvents
	// GetAccessEventTimeReference returns AccessEventTimeReference (property field)
	GetAccessEventTimeReference() BACnetDeviceObjectPropertyReferenceEnclosed
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetEventParameterAccessEventExactly can be used when we want exactly this type and not a type which fulfills BACnetEventParameterAccessEvent.
// This is useful for switch cases.
type BACnetEventParameterAccessEventExactly interface {
	BACnetEventParameterAccessEvent
	isBACnetEventParameterAccessEvent() bool
}

// _BACnetEventParameterAccessEvent is the data-structure of this message
type _BACnetEventParameterAccessEvent struct {
	*_BACnetEventParameter
	OpeningTag               BACnetOpeningTag
	ListOfAccessEvents       BACnetEventParameterAccessEventListOfAccessEvents
	AccessEventTimeReference BACnetDeviceObjectPropertyReferenceEnclosed
	ClosingTag               BACnetClosingTag
}

var _ BACnetEventParameterAccessEvent = (*_BACnetEventParameterAccessEvent)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetEventParameterAccessEvent) InitializeParent(parent BACnetEventParameter, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetEventParameterAccessEvent) GetParent() BACnetEventParameter {
	return m._BACnetEventParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterAccessEvent) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterAccessEvent) GetListOfAccessEvents() BACnetEventParameterAccessEventListOfAccessEvents {
	return m.ListOfAccessEvents
}

func (m *_BACnetEventParameterAccessEvent) GetAccessEventTimeReference() BACnetDeviceObjectPropertyReferenceEnclosed {
	return m.AccessEventTimeReference
}

func (m *_BACnetEventParameterAccessEvent) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterAccessEvent factory function for _BACnetEventParameterAccessEvent
func NewBACnetEventParameterAccessEvent(openingTag BACnetOpeningTag, listOfAccessEvents BACnetEventParameterAccessEventListOfAccessEvents, accessEventTimeReference BACnetDeviceObjectPropertyReferenceEnclosed, closingTag BACnetClosingTag, peekedTagHeader BACnetTagHeader) *_BACnetEventParameterAccessEvent {
	_result := &_BACnetEventParameterAccessEvent{
		OpeningTag:               openingTag,
		ListOfAccessEvents:       listOfAccessEvents,
		AccessEventTimeReference: accessEventTimeReference,
		ClosingTag:               closingTag,
		_BACnetEventParameter:    NewBACnetEventParameter(peekedTagHeader),
	}
	_result._BACnetEventParameter._BACnetEventParameterChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterAccessEvent(structType any) BACnetEventParameterAccessEvent {
	if casted, ok := structType.(BACnetEventParameterAccessEvent); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterAccessEvent); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterAccessEvent) GetTypeName() string {
	return "BACnetEventParameterAccessEvent"
}

func (m *_BACnetEventParameterAccessEvent) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (listOfAccessEvents)
	lengthInBits += m.ListOfAccessEvents.GetLengthInBits(ctx)

	// Simple field (accessEventTimeReference)
	lengthInBits += m.AccessEventTimeReference.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterAccessEvent) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventParameterAccessEventParse(ctx context.Context, theBytes []byte) (BACnetEventParameterAccessEvent, error) {
	return BACnetEventParameterAccessEventParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetEventParameterAccessEventParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterAccessEvent, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterAccessEvent, error) {
		return BACnetEventParameterAccessEventParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetEventParameterAccessEventParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterAccessEvent, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterAccessEvent"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterAccessEvent")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(13))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}

	listOfAccessEvents, err := ReadSimpleField[BACnetEventParameterAccessEventListOfAccessEvents](ctx, "listOfAccessEvents", ReadComplex[BACnetEventParameterAccessEventListOfAccessEvents](BACnetEventParameterAccessEventListOfAccessEventsParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfAccessEvents' field"))
	}

	accessEventTimeReference, err := ReadSimpleField[BACnetDeviceObjectPropertyReferenceEnclosed](ctx, "accessEventTimeReference", ReadComplex[BACnetDeviceObjectPropertyReferenceEnclosed](BACnetDeviceObjectPropertyReferenceEnclosedParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'accessEventTimeReference' field"))
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(13))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterAccessEvent"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterAccessEvent")
	}

	// Create a partially initialized instance
	_child := &_BACnetEventParameterAccessEvent{
		_BACnetEventParameter:    &_BACnetEventParameter{},
		OpeningTag:               openingTag,
		ListOfAccessEvents:       listOfAccessEvents,
		AccessEventTimeReference: accessEventTimeReference,
		ClosingTag:               closingTag,
	}
	_child._BACnetEventParameter._BACnetEventParameterChildRequirements = _child
	return _child, nil
}

func (m *_BACnetEventParameterAccessEvent) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterAccessEvent) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterAccessEvent"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterAccessEvent")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'openingTag' field")
		}

		if err := WriteSimpleField[BACnetEventParameterAccessEventListOfAccessEvents](ctx, "listOfAccessEvents", m.GetListOfAccessEvents(), WriteComplex[BACnetEventParameterAccessEventListOfAccessEvents](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'listOfAccessEvents' field")
		}

		if err := WriteSimpleField[BACnetDeviceObjectPropertyReferenceEnclosed](ctx, "accessEventTimeReference", m.GetAccessEventTimeReference(), WriteComplex[BACnetDeviceObjectPropertyReferenceEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'accessEventTimeReference' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventParameterAccessEvent"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterAccessEvent")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetEventParameterAccessEvent) isBACnetEventParameterAccessEvent() bool {
	return true
}

func (m *_BACnetEventParameterAccessEvent) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
