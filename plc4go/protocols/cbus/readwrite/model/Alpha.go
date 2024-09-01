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

// Alpha is the corresponding interface of Alpha
type Alpha interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetCharacter returns Character (property field)
	GetCharacter() byte
}

// AlphaExactly can be used when we want exactly this type and not a type which fulfills Alpha.
// This is useful for switch cases.
type AlphaExactly interface {
	Alpha
	isAlpha() bool
}

// _Alpha is the data-structure of this message
type _Alpha struct {
	Character byte
}

var _ Alpha = (*_Alpha)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_Alpha) GetCharacter() byte {
	return m.Character
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAlpha factory function for _Alpha
func NewAlpha(character byte) *_Alpha {
	return &_Alpha{Character: character}
}

// Deprecated: use the interface for direct cast
func CastAlpha(structType any) Alpha {
	if casted, ok := structType.(Alpha); ok {
		return casted
	}
	if casted, ok := structType.(*Alpha); ok {
		return *casted
	}
	return nil
}

func (m *_Alpha) GetTypeName() string {
	return "Alpha"
}

func (m *_Alpha) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (character)
	lengthInBits += 8

	return lengthInBits
}

func (m *_Alpha) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AlphaParse(ctx context.Context, theBytes []byte) (Alpha, error) {
	return AlphaParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AlphaParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (Alpha, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (Alpha, error) {
		return AlphaParseWithBuffer(ctx, readBuffer)
	}
}

func AlphaParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (Alpha, error) {
	v, err := (&_Alpha{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_Alpha) parse(ctx context.Context, readBuffer utils.ReadBuffer) (_alpha Alpha, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("Alpha"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Alpha")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	character, err := ReadSimpleField(ctx, "character", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'character' field"))
	}
	m.Character = character

	// Validation
	if !(bool((bool((character) >= (0x67)))) && bool((bool((character) <= (0x7A))))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "character not in alpha space"})
	}

	if closeErr := readBuffer.CloseContext("Alpha"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Alpha")
	}

	return m, nil
}

func (m *_Alpha) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_Alpha) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("Alpha"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for Alpha")
	}

	if err := WriteSimpleField[byte](ctx, "character", m.GetCharacter(), WriteByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'character' field")
	}

	if popErr := writeBuffer.PopContext("Alpha"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for Alpha")
	}
	return nil
}

func (m *_Alpha) isAlpha() bool {
	return true
}

func (m *_Alpha) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
