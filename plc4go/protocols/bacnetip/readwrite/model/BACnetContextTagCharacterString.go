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

// BACnetContextTagCharacterString is the corresponding interface of BACnetContextTagCharacterString
type BACnetContextTagCharacterString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetContextTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadCharacterString
	// GetValue returns Value (virtual field)
	GetValue() string
	// IsBACnetContextTagCharacterString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetContextTagCharacterString()
}

// _BACnetContextTagCharacterString is the data-structure of this message
type _BACnetContextTagCharacterString struct {
	BACnetContextTagContract
	Payload BACnetTagPayloadCharacterString
}

var _ BACnetContextTagCharacterString = (*_BACnetContextTagCharacterString)(nil)
var _ BACnetContextTagRequirements = (*_BACnetContextTagCharacterString)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetContextTagCharacterString) GetDataType() BACnetDataType {
	return BACnetDataType_CHARACTER_STRING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetContextTagCharacterString) GetParent() BACnetContextTagContract {
	return m.BACnetContextTagContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetContextTagCharacterString) GetPayload() BACnetTagPayloadCharacterString {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetContextTagCharacterString) GetValue() string {
	ctx := context.Background()
	_ = ctx
	return fmt.Sprintf("%v", m.GetPayload().GetValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetContextTagCharacterString factory function for _BACnetContextTagCharacterString
func NewBACnetContextTagCharacterString(payload BACnetTagPayloadCharacterString, header BACnetTagHeader, tagNumberArgument uint8) *_BACnetContextTagCharacterString {
	_result := &_BACnetContextTagCharacterString{
		BACnetContextTagContract: NewBACnetContextTag(header, tagNumberArgument),
		Payload:                  payload,
	}
	_result.BACnetContextTagContract.(*_BACnetContextTag)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetContextTagCharacterString(structType any) BACnetContextTagCharacterString {
	if casted, ok := structType.(BACnetContextTagCharacterString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetContextTagCharacterString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetContextTagCharacterString) GetTypeName() string {
	return "BACnetContextTagCharacterString"
}

func (m *_BACnetContextTagCharacterString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetContextTagContract.(*_BACnetContextTag).getLengthInBits(ctx))

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetContextTagCharacterString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetContextTagCharacterString) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetContextTag, header BACnetTagHeader, tagNumberArgument uint8, dataType BACnetDataType) (__bACnetContextTagCharacterString BACnetContextTagCharacterString, err error) {
	m.BACnetContextTagContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetContextTagCharacterString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTagCharacterString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	payload, err := ReadSimpleField[BACnetTagPayloadCharacterString](ctx, "payload", ReadComplex[BACnetTagPayloadCharacterString](BACnetTagPayloadCharacterStringParseWithBufferProducer((uint32)(header.GetActualLength())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	value, err := ReadVirtualField[string](ctx, "value", (*string)(nil), payload.GetValue())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	_ = value

	if closeErr := readBuffer.CloseContext("BACnetContextTagCharacterString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTagCharacterString")
	}

	return m, nil
}

func (m *_BACnetContextTagCharacterString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetContextTagCharacterString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagCharacterString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetContextTagCharacterString")
		}

		if err := WriteSimpleField[BACnetTagPayloadCharacterString](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'payload' field")
		}
		// Virtual field
		value := m.GetValue()
		_ = value
		if _valueErr := writeBuffer.WriteVirtual(ctx, "value", m.GetValue()); _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagCharacterString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetContextTagCharacterString")
		}
		return nil
	}
	return m.BACnetContextTagContract.(*_BACnetContextTag).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetContextTagCharacterString) IsBACnetContextTagCharacterString() {}

func (m *_BACnetContextTagCharacterString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
