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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetLightingInProgressTagged is the corresponding interface of BACnetLightingInProgressTagged
type BACnetLightingInProgressTagged interface {
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetLightingInProgress
}

// BACnetLightingInProgressTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetLightingInProgressTagged.
// This is useful for switch cases.
type BACnetLightingInProgressTaggedExactly interface {
	BACnetLightingInProgressTagged
	isBACnetLightingInProgressTagged() bool
}

// _BACnetLightingInProgressTagged is the data-structure of this message
type _BACnetLightingInProgressTagged struct {
	Header BACnetTagHeader
	Value  BACnetLightingInProgress

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLightingInProgressTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetLightingInProgressTagged) GetValue() BACnetLightingInProgress {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLightingInProgressTagged factory function for _BACnetLightingInProgressTagged
func NewBACnetLightingInProgressTagged(header BACnetTagHeader, value BACnetLightingInProgress, tagNumber uint8, tagClass TagClass) *_BACnetLightingInProgressTagged {
	return &_BACnetLightingInProgressTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetLightingInProgressTagged(structType interface{}) BACnetLightingInProgressTagged {
	if casted, ok := structType.(BACnetLightingInProgressTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLightingInProgressTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLightingInProgressTagged) GetTypeName() string {
	return "BACnetLightingInProgressTagged"
}

func (m *_BACnetLightingInProgressTagged) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetLightingInProgressTagged) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits()

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetLightingInProgressTagged) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLightingInProgressTaggedParse(readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetLightingInProgressTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLightingInProgressTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLightingInProgressTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
	_header, _headerErr := BACnetTagHeaderParse(readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field of BACnetLightingInProgressTagged")
	}
	header := _header.(BACnetTagHeader)
	if closeErr := readBuffer.CloseContext("header"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for header")
	}

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{"tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{"tagnumber doesn't match"})
	}

	// Manual Field (value)
	_value, _valueErr := ReadEnumGenericFailing(readBuffer, header.GetActualLength(), BACnetLightingInProgress_IDLE)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of BACnetLightingInProgressTagged")
	}
	var value BACnetLightingInProgress
	if _value != nil {
		value = _value.(BACnetLightingInProgress)
	}

	if closeErr := readBuffer.CloseContext("BACnetLightingInProgressTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLightingInProgressTagged")
	}

	// Create the instance
	return &_BACnetLightingInProgressTagged{
		TagNumber: tagNumber,
		TagClass:  tagClass,
		Header:    header,
		Value:     value,
	}, nil
}

func (m *_BACnetLightingInProgressTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLightingInProgressTagged) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetLightingInProgressTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLightingInProgressTagged")
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for header")
	}
	_headerErr := writeBuffer.WriteSerializable(m.GetHeader())
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for header")
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
	}

	// Manual Field (value)
	_valueErr := WriteEnumGeneric(writeBuffer, m.GetValue())
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLightingInProgressTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLightingInProgressTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetLightingInProgressTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetLightingInProgressTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetLightingInProgressTagged) isBACnetLightingInProgressTagged() bool {
	return true
}

func (m *_BACnetLightingInProgressTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}