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

// BACnetPolarityTagged is the corresponding interface of BACnetPolarityTagged
type BACnetPolarityTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetPolarity
}

// BACnetPolarityTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetPolarityTagged.
// This is useful for switch cases.
type BACnetPolarityTaggedExactly interface {
	BACnetPolarityTagged
	isBACnetPolarityTagged() bool
}

// _BACnetPolarityTagged is the data-structure of this message
type _BACnetPolarityTagged struct {
	Header BACnetTagHeader
	Value  BACnetPolarity

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetPolarityTagged = (*_BACnetPolarityTagged)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPolarityTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetPolarityTagged) GetValue() BACnetPolarity {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPolarityTagged factory function for _BACnetPolarityTagged
func NewBACnetPolarityTagged(header BACnetTagHeader, value BACnetPolarity, tagNumber uint8, tagClass TagClass) *_BACnetPolarityTagged {
	return &_BACnetPolarityTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetPolarityTagged(structType any) BACnetPolarityTagged {
	if casted, ok := structType.(BACnetPolarityTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPolarityTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPolarityTagged) GetTypeName() string {
	return "BACnetPolarityTagged"
}

func (m *_BACnetPolarityTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetPolarityTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPolarityTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetPolarityTagged, error) {
	return BACnetPolarityTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetPolarityTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPolarityTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPolarityTagged, error) {
		return BACnetPolarityTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetPolarityTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetPolarityTagged, error) {
	v, err := (&_BACnetPolarityTagged{}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetPolarityTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (_bACnetPolarityTagged BACnetPolarityTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPolarityTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPolarityTagged")
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

	value, err := ReadManualField[BACnetPolarity](ctx, "value", readBuffer, EnsureType[BACnetPolarity](ReadEnumGenericFailing(ctx, readBuffer, header.GetActualLength(), BACnetPolarity_NORMAL)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("BACnetPolarityTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPolarityTagged")
	}

	return m, nil
}

func (m *_BACnetPolarityTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPolarityTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetPolarityTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetPolarityTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetPolarity](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetPolarityTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetPolarityTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetPolarityTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetPolarityTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetPolarityTagged) isBACnetPolarityTagged() bool {
	return true
}

func (m *_BACnetPolarityTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
