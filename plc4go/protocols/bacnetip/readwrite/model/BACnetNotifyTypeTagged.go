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

// BACnetNotifyTypeTagged is the corresponding interface of BACnetNotifyTypeTagged
type BACnetNotifyTypeTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetNotifyType
}

// BACnetNotifyTypeTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetNotifyTypeTagged.
// This is useful for switch cases.
type BACnetNotifyTypeTaggedExactly interface {
	BACnetNotifyTypeTagged
	isBACnetNotifyTypeTagged() bool
}

// _BACnetNotifyTypeTagged is the data-structure of this message
type _BACnetNotifyTypeTagged struct {
	Header BACnetTagHeader
	Value  BACnetNotifyType

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetNotifyTypeTagged = (*_BACnetNotifyTypeTagged)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotifyTypeTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetNotifyTypeTagged) GetValue() BACnetNotifyType {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotifyTypeTagged factory function for _BACnetNotifyTypeTagged
func NewBACnetNotifyTypeTagged(header BACnetTagHeader, value BACnetNotifyType, tagNumber uint8, tagClass TagClass) *_BACnetNotifyTypeTagged {
	return &_BACnetNotifyTypeTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetNotifyTypeTagged(structType any) BACnetNotifyTypeTagged {
	if casted, ok := structType.(BACnetNotifyTypeTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotifyTypeTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotifyTypeTagged) GetTypeName() string {
	return "BACnetNotifyTypeTagged"
}

func (m *_BACnetNotifyTypeTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetNotifyTypeTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetNotifyTypeTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetNotifyTypeTagged, error) {
	return BACnetNotifyTypeTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetNotifyTypeTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetNotifyTypeTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetNotifyTypeTagged, error) {
		return BACnetNotifyTypeTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetNotifyTypeTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetNotifyTypeTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotifyTypeTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotifyTypeTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := ReadSimpleField[BACnetTagHeader](ctx, "header", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "tagnumber doesn't match"})
	}

	value, err := ReadManualField[BACnetNotifyType](ctx, "value", readBuffer, EnsureType[BACnetNotifyType](ReadEnumGenericFailing(ctx, readBuffer, header.GetActualLength(), BACnetNotifyType_ALARM)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetNotifyTypeTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotifyTypeTagged")
	}

	// Create the instance
	return &_BACnetNotifyTypeTagged{
		TagNumber: tagNumber,
		TagClass:  tagClass,
		Header:    header,
		Value:     value,
	}, nil
}

func (m *_BACnetNotifyTypeTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotifyTypeTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetNotifyTypeTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetNotifyTypeTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetNotifyType](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetNotifyTypeTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetNotifyTypeTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetNotifyTypeTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetNotifyTypeTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetNotifyTypeTagged) isBACnetNotifyTypeTagged() bool {
	return true
}

func (m *_BACnetNotifyTypeTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
