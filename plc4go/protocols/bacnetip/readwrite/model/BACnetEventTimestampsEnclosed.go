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

// BACnetEventTimestampsEnclosed is the corresponding interface of BACnetEventTimestampsEnclosed
type BACnetEventTimestampsEnclosed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetEventTimestamps returns EventTimestamps (property field)
	GetEventTimestamps() BACnetEventTimestamps
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetEventTimestampsEnclosed is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventTimestampsEnclosed()
}

// _BACnetEventTimestampsEnclosed is the data-structure of this message
type _BACnetEventTimestampsEnclosed struct {
	OpeningTag      BACnetOpeningTag
	EventTimestamps BACnetEventTimestamps
	ClosingTag      BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetEventTimestampsEnclosed = (*_BACnetEventTimestampsEnclosed)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventTimestampsEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventTimestampsEnclosed) GetEventTimestamps() BACnetEventTimestamps {
	return m.EventTimestamps
}

func (m *_BACnetEventTimestampsEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventTimestampsEnclosed factory function for _BACnetEventTimestampsEnclosed
func NewBACnetEventTimestampsEnclosed(openingTag BACnetOpeningTag, eventTimestamps BACnetEventTimestamps, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetEventTimestampsEnclosed {
	return &_BACnetEventTimestampsEnclosed{OpeningTag: openingTag, EventTimestamps: eventTimestamps, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetEventTimestampsEnclosed(structType any) BACnetEventTimestampsEnclosed {
	if casted, ok := structType.(BACnetEventTimestampsEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventTimestampsEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventTimestampsEnclosed) GetTypeName() string {
	return "BACnetEventTimestampsEnclosed"
}

func (m *_BACnetEventTimestampsEnclosed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (eventTimestamps)
	lengthInBits += m.EventTimestamps.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventTimestampsEnclosed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventTimestampsEnclosedParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetEventTimestampsEnclosed, error) {
	return BACnetEventTimestampsEnclosedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetEventTimestampsEnclosedParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventTimestampsEnclosed, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventTimestampsEnclosed, error) {
		return BACnetEventTimestampsEnclosedParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetEventTimestampsEnclosedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetEventTimestampsEnclosed, error) {
	v, err := (&_BACnetEventTimestampsEnclosed{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetEventTimestampsEnclosed) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetEventTimestampsEnclosed BACnetEventTimestampsEnclosed, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventTimestampsEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventTimestampsEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	eventTimestamps, err := ReadSimpleField[BACnetEventTimestamps](ctx, "eventTimestamps", ReadComplex[BACnetEventTimestamps](BACnetEventTimestampsParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventTimestamps' field"))
	}
	m.EventTimestamps = eventTimestamps

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetEventTimestampsEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventTimestampsEnclosed")
	}

	return m, nil
}

func (m *_BACnetEventTimestampsEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventTimestampsEnclosed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetEventTimestampsEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventTimestampsEnclosed")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteSimpleField[BACnetEventTimestamps](ctx, "eventTimestamps", m.GetEventTimestamps(), WriteComplex[BACnetEventTimestamps](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'eventTimestamps' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventTimestampsEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventTimestampsEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetEventTimestampsEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetEventTimestampsEnclosed) IsBACnetEventTimestampsEnclosed() {}

func (m *_BACnetEventTimestampsEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
