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

// PascalString is the corresponding interface of PascalString
type PascalString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetStringValue returns StringValue (property field)
	GetStringValue() string
	// GetStringLength returns StringLength (virtual field)
	GetStringLength() int32
}

// PascalStringExactly can be used when we want exactly this type and not a type which fulfills PascalString.
// This is useful for switch cases.
type PascalStringExactly interface {
	PascalString
	isPascalString() bool
}

// _PascalString is the data-structure of this message
type _PascalString struct {
	StringValue string
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PascalString) GetStringValue() string {
	return m.StringValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_PascalString) GetStringLength() int32 {
	ctx := context.Background()
	_ = ctx
	return int32(PascalLengthToUtf8Length(ctx, Utf8LengthToPascalLength(ctx, m.GetStringValue())))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewPascalString factory function for _PascalString
func NewPascalString(stringValue string) *_PascalString {
	return &_PascalString{StringValue: stringValue}
}

// Deprecated: use the interface for direct cast
func CastPascalString(structType any) PascalString {
	if casted, ok := structType.(PascalString); ok {
		return casted
	}
	if casted, ok := structType.(*PascalString); ok {
		return *casted
	}
	return nil
}

func (m *_PascalString) GetTypeName() string {
	return "PascalString"
}

func (m *_PascalString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (sLength)
	lengthInBits += 32

	// A virtual field doesn't have any in- or output.

	// Simple field (stringValue)
	lengthInBits += uint16(int32(m.GetStringLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_PascalString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func PascalStringParse(ctx context.Context, theBytes []byte) (PascalString, error) {
	return PascalStringParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func PascalStringParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (PascalString, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (PascalString, error) {
		return PascalStringParseWithBuffer(ctx, readBuffer)
	}
}

func PascalStringParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (PascalString, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("PascalString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PascalString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	sLength, err := ReadImplicitField[int32](ctx, "sLength", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sLength' field"))
	}
	_ = sLength

	stringLength, err := ReadVirtualField[int32](ctx, "stringLength", (*int32)(nil), PascalLengthToUtf8Length(ctx, sLength))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'stringLength' field"))
	}

	stringValue, err := ReadSimpleField(ctx, "stringValue", ReadString(readBuffer, uint32(int32(stringLength)*int32(int32(8)))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'stringValue' field"))
	}

	if closeErr := readBuffer.CloseContext("PascalString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PascalString")
	}

	// Create the instance
	return &_PascalString{
		StringValue: stringValue,
	}, nil
}

func (m *_PascalString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PascalString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("PascalString"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for PascalString")
	}

	// Implicit Field (sLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	sLength := int32(Utf8LengthToPascalLength(ctx, m.GetStringValue()))
	_sLengthErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("sLength", 32, int32((sLength)))
	if _sLengthErr != nil {
		return errors.Wrap(_sLengthErr, "Error serializing 'sLength' field")
	}
	// Virtual field
	stringLength := m.GetStringLength()
	_ = stringLength
	if _stringLengthErr := writeBuffer.WriteVirtual(ctx, "stringLength", m.GetStringLength()); _stringLengthErr != nil {
		return errors.Wrap(_stringLengthErr, "Error serializing 'stringLength' field")
	}

	// Simple Field (stringValue)
	stringValue := string(m.GetStringValue())
	_stringValueErr := /*TODO: migrate me*/ writeBuffer.WriteString("stringValue", uint32((stringLength)*(8)), (stringValue), utils.WithEncoding("UTF-8)"))
	if _stringValueErr != nil {
		return errors.Wrap(_stringValueErr, "Error serializing 'stringValue' field")
	}

	if popErr := writeBuffer.PopContext("PascalString"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for PascalString")
	}
	return nil
}

func (m *_PascalString) isPascalString() bool {
	return true
}

func (m *_PascalString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
