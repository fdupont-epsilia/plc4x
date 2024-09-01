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

// BACnetTimeStampEnclosed is the corresponding interface of BACnetTimeStampEnclosed
type BACnetTimeStampEnclosed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() BACnetTimeStamp
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetTimeStampEnclosedExactly can be used when we want exactly this type and not a type which fulfills BACnetTimeStampEnclosed.
// This is useful for switch cases.
type BACnetTimeStampEnclosedExactly interface {
	BACnetTimeStampEnclosed
	isBACnetTimeStampEnclosed() bool
}

// _BACnetTimeStampEnclosed is the data-structure of this message
type _BACnetTimeStampEnclosed struct {
	OpeningTag BACnetOpeningTag
	Timestamp  BACnetTimeStamp
	ClosingTag BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetTimeStampEnclosed = (*_BACnetTimeStampEnclosed)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimeStampEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetTimeStampEnclosed) GetTimestamp() BACnetTimeStamp {
	return m.Timestamp
}

func (m *_BACnetTimeStampEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTimeStampEnclosed factory function for _BACnetTimeStampEnclosed
func NewBACnetTimeStampEnclosed(openingTag BACnetOpeningTag, timestamp BACnetTimeStamp, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetTimeStampEnclosed {
	return &_BACnetTimeStampEnclosed{OpeningTag: openingTag, Timestamp: timestamp, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetTimeStampEnclosed(structType any) BACnetTimeStampEnclosed {
	if casted, ok := structType.(BACnetTimeStampEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimeStampEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimeStampEnclosed) GetTypeName() string {
	return "BACnetTimeStampEnclosed"
}

func (m *_BACnetTimeStampEnclosed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (timestamp)
	lengthInBits += m.Timestamp.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetTimeStampEnclosed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTimeStampEnclosedParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetTimeStampEnclosed, error) {
	return BACnetTimeStampEnclosedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetTimeStampEnclosedParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTimeStampEnclosed, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTimeStampEnclosed, error) {
		return BACnetTimeStampEnclosedParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetTimeStampEnclosedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetTimeStampEnclosed, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimeStampEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimeStampEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}

	timestamp, err := ReadSimpleField[BACnetTimeStamp](ctx, "timestamp", ReadComplex[BACnetTimeStamp](BACnetTimeStampParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timestamp' field"))
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetTimeStampEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimeStampEnclosed")
	}

	// Create the instance
	return &_BACnetTimeStampEnclosed{
		TagNumber:  tagNumber,
		OpeningTag: openingTag,
		Timestamp:  timestamp,
		ClosingTag: closingTag,
	}, nil
}

func (m *_BACnetTimeStampEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimeStampEnclosed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetTimeStampEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTimeStampEnclosed")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteSimpleField[BACnetTimeStamp](ctx, "timestamp", m.GetTimestamp(), WriteComplex[BACnetTimeStamp](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'timestamp' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTimeStampEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTimeStampEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetTimeStampEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetTimeStampEnclosed) isBACnetTimeStampEnclosed() bool {
	return true
}

func (m *_BACnetTimeStampEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
