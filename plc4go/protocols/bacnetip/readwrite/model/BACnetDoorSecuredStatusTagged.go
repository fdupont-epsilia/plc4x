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

// BACnetDoorSecuredStatusTagged is the corresponding interface of BACnetDoorSecuredStatusTagged
type BACnetDoorSecuredStatusTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetDoorSecuredStatus
	// IsBACnetDoorSecuredStatusTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetDoorSecuredStatusTagged()
}

// _BACnetDoorSecuredStatusTagged is the data-structure of this message
type _BACnetDoorSecuredStatusTagged struct {
	Header BACnetTagHeader
	Value  BACnetDoorSecuredStatus

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetDoorSecuredStatusTagged = (*_BACnetDoorSecuredStatusTagged)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetDoorSecuredStatusTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetDoorSecuredStatusTagged) GetValue() BACnetDoorSecuredStatus {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetDoorSecuredStatusTagged factory function for _BACnetDoorSecuredStatusTagged
func NewBACnetDoorSecuredStatusTagged(header BACnetTagHeader, value BACnetDoorSecuredStatus, tagNumber uint8, tagClass TagClass) *_BACnetDoorSecuredStatusTagged {
	return &_BACnetDoorSecuredStatusTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetDoorSecuredStatusTagged(structType any) BACnetDoorSecuredStatusTagged {
	if casted, ok := structType.(BACnetDoorSecuredStatusTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetDoorSecuredStatusTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetDoorSecuredStatusTagged) GetTypeName() string {
	return "BACnetDoorSecuredStatusTagged"
}

func (m *_BACnetDoorSecuredStatusTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetDoorSecuredStatusTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetDoorSecuredStatusTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetDoorSecuredStatusTagged, error) {
	return BACnetDoorSecuredStatusTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetDoorSecuredStatusTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDoorSecuredStatusTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDoorSecuredStatusTagged, error) {
		return BACnetDoorSecuredStatusTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetDoorSecuredStatusTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetDoorSecuredStatusTagged, error) {
	v, err := (&_BACnetDoorSecuredStatusTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetDoorSecuredStatusTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetDoorSecuredStatusTagged BACnetDoorSecuredStatusTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetDoorSecuredStatusTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetDoorSecuredStatusTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := ReadSimpleField[BACnetTagHeader](ctx, "header", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}
	m.Header = header

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "tagnumber doesn't match"})
	}

	value, err := ReadManualField[BACnetDoorSecuredStatus](ctx, "value", readBuffer, EnsureType[BACnetDoorSecuredStatus](ReadEnumGenericFailing(ctx, readBuffer, header.GetActualLength(), BACnetDoorSecuredStatus_SECURED)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("BACnetDoorSecuredStatusTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetDoorSecuredStatusTagged")
	}

	return m, nil
}

func (m *_BACnetDoorSecuredStatusTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetDoorSecuredStatusTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetDoorSecuredStatusTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetDoorSecuredStatusTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetDoorSecuredStatus](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetDoorSecuredStatusTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetDoorSecuredStatusTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetDoorSecuredStatusTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetDoorSecuredStatusTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetDoorSecuredStatusTagged) IsBACnetDoorSecuredStatusTagged() {}

func (m *_BACnetDoorSecuredStatusTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
