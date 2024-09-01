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

// BACnetRecipientProcessEnclosed is the corresponding interface of BACnetRecipientProcessEnclosed
type BACnetRecipientProcessEnclosed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetRecipientProcess returns RecipientProcess (property field)
	GetRecipientProcess() BACnetRecipientProcess
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetRecipientProcessEnclosedExactly can be used when we want exactly this type and not a type which fulfills BACnetRecipientProcessEnclosed.
// This is useful for switch cases.
type BACnetRecipientProcessEnclosedExactly interface {
	BACnetRecipientProcessEnclosed
	isBACnetRecipientProcessEnclosed() bool
}

// _BACnetRecipientProcessEnclosed is the data-structure of this message
type _BACnetRecipientProcessEnclosed struct {
	OpeningTag       BACnetOpeningTag
	RecipientProcess BACnetRecipientProcess
	ClosingTag       BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetRecipientProcessEnclosed = (*_BACnetRecipientProcessEnclosed)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetRecipientProcessEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetRecipientProcessEnclosed) GetRecipientProcess() BACnetRecipientProcess {
	return m.RecipientProcess
}

func (m *_BACnetRecipientProcessEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetRecipientProcessEnclosed factory function for _BACnetRecipientProcessEnclosed
func NewBACnetRecipientProcessEnclosed(openingTag BACnetOpeningTag, recipientProcess BACnetRecipientProcess, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetRecipientProcessEnclosed {
	return &_BACnetRecipientProcessEnclosed{OpeningTag: openingTag, RecipientProcess: recipientProcess, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetRecipientProcessEnclosed(structType any) BACnetRecipientProcessEnclosed {
	if casted, ok := structType.(BACnetRecipientProcessEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetRecipientProcessEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetRecipientProcessEnclosed) GetTypeName() string {
	return "BACnetRecipientProcessEnclosed"
}

func (m *_BACnetRecipientProcessEnclosed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (recipientProcess)
	lengthInBits += m.RecipientProcess.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetRecipientProcessEnclosed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetRecipientProcessEnclosedParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetRecipientProcessEnclosed, error) {
	return BACnetRecipientProcessEnclosedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetRecipientProcessEnclosedParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetRecipientProcessEnclosed, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetRecipientProcessEnclosed, error) {
		return BACnetRecipientProcessEnclosedParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetRecipientProcessEnclosedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetRecipientProcessEnclosed, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetRecipientProcessEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetRecipientProcessEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}

	recipientProcess, err := ReadSimpleField[BACnetRecipientProcess](ctx, "recipientProcess", ReadComplex[BACnetRecipientProcess](BACnetRecipientProcessParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'recipientProcess' field"))
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetRecipientProcessEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetRecipientProcessEnclosed")
	}

	// Create the instance
	return &_BACnetRecipientProcessEnclosed{
		TagNumber:        tagNumber,
		OpeningTag:       openingTag,
		RecipientProcess: recipientProcess,
		ClosingTag:       closingTag,
	}, nil
}

func (m *_BACnetRecipientProcessEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetRecipientProcessEnclosed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetRecipientProcessEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetRecipientProcessEnclosed")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteSimpleField[BACnetRecipientProcess](ctx, "recipientProcess", m.GetRecipientProcess(), WriteComplex[BACnetRecipientProcess](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'recipientProcess' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetRecipientProcessEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetRecipientProcessEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetRecipientProcessEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetRecipientProcessEnclosed) isBACnetRecipientProcessEnclosed() bool {
	return true
}

func (m *_BACnetRecipientProcessEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
