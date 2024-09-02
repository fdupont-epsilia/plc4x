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

// IdentifyReplyCommandDelays is the corresponding interface of IdentifyReplyCommandDelays
type IdentifyReplyCommandDelays interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	IdentifyReplyCommand
	// GetTerminalLevels returns TerminalLevels (property field)
	GetTerminalLevels() []byte
	// GetReStrikeDelay returns ReStrikeDelay (property field)
	GetReStrikeDelay() byte
	// IsIdentifyReplyCommandDelays is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsIdentifyReplyCommandDelays()
}

// _IdentifyReplyCommandDelays is the data-structure of this message
type _IdentifyReplyCommandDelays struct {
	IdentifyReplyCommandContract
	TerminalLevels []byte
	ReStrikeDelay  byte
}

var _ IdentifyReplyCommandDelays = (*_IdentifyReplyCommandDelays)(nil)
var _ IdentifyReplyCommandRequirements = (*_IdentifyReplyCommandDelays)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandDelays) GetAttribute() Attribute {
	return Attribute_Delays
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandDelays) GetParent() IdentifyReplyCommandContract {
	return m.IdentifyReplyCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandDelays) GetTerminalLevels() []byte {
	return m.TerminalLevels
}

func (m *_IdentifyReplyCommandDelays) GetReStrikeDelay() byte {
	return m.ReStrikeDelay
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewIdentifyReplyCommandDelays factory function for _IdentifyReplyCommandDelays
func NewIdentifyReplyCommandDelays(terminalLevels []byte, reStrikeDelay byte, numBytes uint8) *_IdentifyReplyCommandDelays {
	_result := &_IdentifyReplyCommandDelays{
		IdentifyReplyCommandContract: NewIdentifyReplyCommand(numBytes),
		TerminalLevels:               terminalLevels,
		ReStrikeDelay:                reStrikeDelay,
	}
	_result.IdentifyReplyCommandContract.(*_IdentifyReplyCommand)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandDelays(structType any) IdentifyReplyCommandDelays {
	if casted, ok := structType.(IdentifyReplyCommandDelays); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandDelays); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandDelays) GetTypeName() string {
	return "IdentifyReplyCommandDelays"
}

func (m *_IdentifyReplyCommandDelays) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).getLengthInBits(ctx))

	// Array field
	if len(m.TerminalLevels) > 0 {
		lengthInBits += 8 * uint16(len(m.TerminalLevels))
	}

	// Simple field (reStrikeDelay)
	lengthInBits += 8

	return lengthInBits
}

func (m *_IdentifyReplyCommandDelays) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_IdentifyReplyCommandDelays) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_IdentifyReplyCommand, attribute Attribute, numBytes uint8) (__identifyReplyCommandDelays IdentifyReplyCommandDelays, err error) {
	m.IdentifyReplyCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandDelays"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandDelays")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	terminalLevels, err := readBuffer.ReadByteArray("terminalLevels", int(int32(numBytes)-int32(int32(1))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'terminalLevels' field"))
	}
	m.TerminalLevels = terminalLevels

	reStrikeDelay, err := ReadSimpleField(ctx, "reStrikeDelay", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reStrikeDelay' field"))
	}
	m.ReStrikeDelay = reStrikeDelay

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandDelays"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandDelays")
	}

	return m, nil
}

func (m *_IdentifyReplyCommandDelays) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IdentifyReplyCommandDelays) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandDelays"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandDelays")
		}

		if err := WriteByteArrayField(ctx, "terminalLevels", m.GetTerminalLevels(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'terminalLevels' field")
		}

		if err := WriteSimpleField[byte](ctx, "reStrikeDelay", m.GetReStrikeDelay(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reStrikeDelay' field")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandDelays"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandDelays")
		}
		return nil
	}
	return m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_IdentifyReplyCommandDelays) IsIdentifyReplyCommandDelays() {}

func (m *_IdentifyReplyCommandDelays) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
