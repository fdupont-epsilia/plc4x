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

// BACnetTagPayloadOctetString is the corresponding interface of BACnetTagPayloadOctetString
type BACnetTagPayloadOctetString interface {
	utils.LengthAware
	utils.Serializable
	// GetOctets returns Octets (property field)
	GetOctets() []byte
}

// BACnetTagPayloadOctetStringExactly can be used when we want exactly this type and not a type which fulfills BACnetTagPayloadOctetString.
// This is useful for switch cases.
type BACnetTagPayloadOctetStringExactly interface {
	BACnetTagPayloadOctetString
	isBACnetTagPayloadOctetString() bool
}

// _BACnetTagPayloadOctetString is the data-structure of this message
type _BACnetTagPayloadOctetString struct {
	Octets []byte

	// Arguments.
	ActualLength uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTagPayloadOctetString) GetOctets() []byte {
	return m.Octets
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTagPayloadOctetString factory function for _BACnetTagPayloadOctetString
func NewBACnetTagPayloadOctetString(octets []byte, actualLength uint32) *_BACnetTagPayloadOctetString {
	return &_BACnetTagPayloadOctetString{Octets: octets, ActualLength: actualLength}
}

// Deprecated: use the interface for direct cast
func CastBACnetTagPayloadOctetString(structType interface{}) BACnetTagPayloadOctetString {
	if casted, ok := structType.(BACnetTagPayloadOctetString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTagPayloadOctetString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTagPayloadOctetString) GetTypeName() string {
	return "BACnetTagPayloadOctetString"
}

func (m *_BACnetTagPayloadOctetString) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetTagPayloadOctetString) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Array field
	if len(m.Octets) > 0 {
		lengthInBits += 8 * uint16(len(m.Octets))
	}

	return lengthInBits
}

func (m *_BACnetTagPayloadOctetString) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTagPayloadOctetStringParse(readBuffer utils.ReadBuffer, actualLength uint32) (BACnetTagPayloadOctetString, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTagPayloadOctetString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTagPayloadOctetString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	// Byte Array field (octets)
	numberOfBytesoctets := int(actualLength)
	octets, _readArrayErr := readBuffer.ReadByteArray("octets", numberOfBytesoctets)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'octets' field of BACnetTagPayloadOctetString")
	}

	if closeErr := readBuffer.CloseContext("BACnetTagPayloadOctetString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTagPayloadOctetString")
	}

	// Create the instance
	return &_BACnetTagPayloadOctetString{
		ActualLength: actualLength,
		Octets:       octets,
	}, nil
}

func (m *_BACnetTagPayloadOctetString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTagPayloadOctetString) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetTagPayloadOctetString"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTagPayloadOctetString")
	}

	// Array Field (octets)
	// Byte Array field (octets)
	if err := writeBuffer.WriteByteArray("octets", m.GetOctets()); err != nil {
		return errors.Wrap(err, "Error serializing 'octets' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagPayloadOctetString"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTagPayloadOctetString")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetTagPayloadOctetString) GetActualLength() uint32 {
	return m.ActualLength
}

//
////

func (m *_BACnetTagPayloadOctetString) isBACnetTagPayloadOctetString() bool {
	return true
}

func (m *_BACnetTagPayloadOctetString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}