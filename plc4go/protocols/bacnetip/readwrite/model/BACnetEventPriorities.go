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

// BACnetEventPriorities is the corresponding interface of BACnetEventPriorities
type BACnetEventPriorities interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetToOffnormal returns ToOffnormal (property field)
	GetToOffnormal() BACnetApplicationTagUnsignedInteger
	// GetToFault returns ToFault (property field)
	GetToFault() BACnetApplicationTagUnsignedInteger
	// GetToNormal returns ToNormal (property field)
	GetToNormal() BACnetApplicationTagUnsignedInteger
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetEventPriorities is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventPriorities()
}

// _BACnetEventPriorities is the data-structure of this message
type _BACnetEventPriorities struct {
	OpeningTag  BACnetOpeningTag
	ToOffnormal BACnetApplicationTagUnsignedInteger
	ToFault     BACnetApplicationTagUnsignedInteger
	ToNormal    BACnetApplicationTagUnsignedInteger
	ClosingTag  BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetEventPriorities = (*_BACnetEventPriorities)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventPriorities) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventPriorities) GetToOffnormal() BACnetApplicationTagUnsignedInteger {
	return m.ToOffnormal
}

func (m *_BACnetEventPriorities) GetToFault() BACnetApplicationTagUnsignedInteger {
	return m.ToFault
}

func (m *_BACnetEventPriorities) GetToNormal() BACnetApplicationTagUnsignedInteger {
	return m.ToNormal
}

func (m *_BACnetEventPriorities) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventPriorities factory function for _BACnetEventPriorities
func NewBACnetEventPriorities(openingTag BACnetOpeningTag, toOffnormal BACnetApplicationTagUnsignedInteger, toFault BACnetApplicationTagUnsignedInteger, toNormal BACnetApplicationTagUnsignedInteger, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetEventPriorities {
	return &_BACnetEventPriorities{OpeningTag: openingTag, ToOffnormal: toOffnormal, ToFault: toFault, ToNormal: toNormal, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetEventPriorities(structType any) BACnetEventPriorities {
	if casted, ok := structType.(BACnetEventPriorities); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventPriorities); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventPriorities) GetTypeName() string {
	return "BACnetEventPriorities"
}

func (m *_BACnetEventPriorities) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (toOffnormal)
	lengthInBits += m.ToOffnormal.GetLengthInBits(ctx)

	// Simple field (toFault)
	lengthInBits += m.ToFault.GetLengthInBits(ctx)

	// Simple field (toNormal)
	lengthInBits += m.ToNormal.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventPriorities) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventPrioritiesParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetEventPriorities, error) {
	return BACnetEventPrioritiesParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetEventPrioritiesParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventPriorities, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventPriorities, error) {
		return BACnetEventPrioritiesParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetEventPrioritiesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetEventPriorities, error) {
	v, err := (&_BACnetEventPriorities{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetEventPriorities) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetEventPriorities BACnetEventPriorities, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventPriorities"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventPriorities")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	toOffnormal, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "toOffnormal", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toOffnormal' field"))
	}
	m.ToOffnormal = toOffnormal

	toFault, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "toFault", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toFault' field"))
	}
	m.ToFault = toFault

	toNormal, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "toNormal", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toNormal' field"))
	}
	m.ToNormal = toNormal

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetEventPriorities"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventPriorities")
	}

	return m, nil
}

func (m *_BACnetEventPriorities) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventPriorities) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetEventPriorities"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventPriorities")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "toOffnormal", m.GetToOffnormal(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'toOffnormal' field")
	}

	if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "toFault", m.GetToFault(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'toFault' field")
	}

	if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "toNormal", m.GetToNormal(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'toNormal' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventPriorities"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventPriorities")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetEventPriorities) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetEventPriorities) IsBACnetEventPriorities() {}

func (m *_BACnetEventPriorities) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
