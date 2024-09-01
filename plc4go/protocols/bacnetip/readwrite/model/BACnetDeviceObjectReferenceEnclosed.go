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

// BACnetDeviceObjectReferenceEnclosed is the corresponding interface of BACnetDeviceObjectReferenceEnclosed
type BACnetDeviceObjectReferenceEnclosed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetObjectReference returns ObjectReference (property field)
	GetObjectReference() BACnetDeviceObjectReference
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetDeviceObjectReferenceEnclosedExactly can be used when we want exactly this type and not a type which fulfills BACnetDeviceObjectReferenceEnclosed.
// This is useful for switch cases.
type BACnetDeviceObjectReferenceEnclosedExactly interface {
	BACnetDeviceObjectReferenceEnclosed
	isBACnetDeviceObjectReferenceEnclosed() bool
}

// _BACnetDeviceObjectReferenceEnclosed is the data-structure of this message
type _BACnetDeviceObjectReferenceEnclosed struct {
	OpeningTag      BACnetOpeningTag
	ObjectReference BACnetDeviceObjectReference
	ClosingTag      BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetDeviceObjectReferenceEnclosed = (*_BACnetDeviceObjectReferenceEnclosed)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetDeviceObjectReferenceEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetDeviceObjectReferenceEnclosed) GetObjectReference() BACnetDeviceObjectReference {
	return m.ObjectReference
}

func (m *_BACnetDeviceObjectReferenceEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetDeviceObjectReferenceEnclosed factory function for _BACnetDeviceObjectReferenceEnclosed
func NewBACnetDeviceObjectReferenceEnclosed(openingTag BACnetOpeningTag, objectReference BACnetDeviceObjectReference, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetDeviceObjectReferenceEnclosed {
	return &_BACnetDeviceObjectReferenceEnclosed{OpeningTag: openingTag, ObjectReference: objectReference, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetDeviceObjectReferenceEnclosed(structType any) BACnetDeviceObjectReferenceEnclosed {
	if casted, ok := structType.(BACnetDeviceObjectReferenceEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetDeviceObjectReferenceEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetDeviceObjectReferenceEnclosed) GetTypeName() string {
	return "BACnetDeviceObjectReferenceEnclosed"
}

func (m *_BACnetDeviceObjectReferenceEnclosed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (objectReference)
	lengthInBits += m.ObjectReference.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetDeviceObjectReferenceEnclosed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetDeviceObjectReferenceEnclosedParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetDeviceObjectReferenceEnclosed, error) {
	return BACnetDeviceObjectReferenceEnclosedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetDeviceObjectReferenceEnclosedParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDeviceObjectReferenceEnclosed, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDeviceObjectReferenceEnclosed, error) {
		return BACnetDeviceObjectReferenceEnclosedParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetDeviceObjectReferenceEnclosedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetDeviceObjectReferenceEnclosed, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetDeviceObjectReferenceEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetDeviceObjectReferenceEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}

	objectReference, err := ReadSimpleField[BACnetDeviceObjectReference](ctx, "objectReference", ReadComplex[BACnetDeviceObjectReference](BACnetDeviceObjectReferenceParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectReference' field"))
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetDeviceObjectReferenceEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetDeviceObjectReferenceEnclosed")
	}

	// Create the instance
	return &_BACnetDeviceObjectReferenceEnclosed{
		TagNumber:       tagNumber,
		OpeningTag:      openingTag,
		ObjectReference: objectReference,
		ClosingTag:      closingTag,
	}, nil
}

func (m *_BACnetDeviceObjectReferenceEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetDeviceObjectReferenceEnclosed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetDeviceObjectReferenceEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetDeviceObjectReferenceEnclosed")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteSimpleField[BACnetDeviceObjectReference](ctx, "objectReference", m.GetObjectReference(), WriteComplex[BACnetDeviceObjectReference](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'objectReference' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetDeviceObjectReferenceEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetDeviceObjectReferenceEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetDeviceObjectReferenceEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetDeviceObjectReferenceEnclosed) isBACnetDeviceObjectReferenceEnclosed() bool {
	return true
}

func (m *_BACnetDeviceObjectReferenceEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
