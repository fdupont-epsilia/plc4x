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

// ChannelInformation is the corresponding interface of ChannelInformation
type ChannelInformation interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetNumChannels returns NumChannels (property field)
	GetNumChannels() uint8
	// GetChannelCode returns ChannelCode (property field)
	GetChannelCode() uint16
	// IsChannelInformation is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsChannelInformation()
}

// _ChannelInformation is the data-structure of this message
type _ChannelInformation struct {
	NumChannels uint8
	ChannelCode uint16
}

var _ ChannelInformation = (*_ChannelInformation)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ChannelInformation) GetNumChannels() uint8 {
	return m.NumChannels
}

func (m *_ChannelInformation) GetChannelCode() uint16 {
	return m.ChannelCode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewChannelInformation factory function for _ChannelInformation
func NewChannelInformation(numChannels uint8, channelCode uint16) *_ChannelInformation {
	return &_ChannelInformation{NumChannels: numChannels, ChannelCode: channelCode}
}

// Deprecated: use the interface for direct cast
func CastChannelInformation(structType any) ChannelInformation {
	if casted, ok := structType.(ChannelInformation); ok {
		return casted
	}
	if casted, ok := structType.(*ChannelInformation); ok {
		return *casted
	}
	return nil
}

func (m *_ChannelInformation) GetTypeName() string {
	return "ChannelInformation"
}

func (m *_ChannelInformation) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (numChannels)
	lengthInBits += 3

	// Simple field (channelCode)
	lengthInBits += 13

	return lengthInBits
}

func (m *_ChannelInformation) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ChannelInformationParse(ctx context.Context, theBytes []byte) (ChannelInformation, error) {
	return ChannelInformationParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ChannelInformationParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (ChannelInformation, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ChannelInformation, error) {
		return ChannelInformationParseWithBuffer(ctx, readBuffer)
	}
}

func ChannelInformationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ChannelInformation, error) {
	v, err := (&_ChannelInformation{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_ChannelInformation) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__channelInformation ChannelInformation, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ChannelInformation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ChannelInformation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	numChannels, err := ReadSimpleField(ctx, "numChannels", ReadUnsignedByte(readBuffer, uint8(3)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numChannels' field"))
	}
	m.NumChannels = numChannels

	channelCode, err := ReadSimpleField(ctx, "channelCode", ReadUnsignedShort(readBuffer, uint8(13)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'channelCode' field"))
	}
	m.ChannelCode = channelCode

	if closeErr := readBuffer.CloseContext("ChannelInformation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ChannelInformation")
	}

	return m, nil
}

func (m *_ChannelInformation) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ChannelInformation) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ChannelInformation"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ChannelInformation")
	}

	if err := WriteSimpleField[uint8](ctx, "numChannels", m.GetNumChannels(), WriteUnsignedByte(writeBuffer, 3)); err != nil {
		return errors.Wrap(err, "Error serializing 'numChannels' field")
	}

	if err := WriteSimpleField[uint16](ctx, "channelCode", m.GetChannelCode(), WriteUnsignedShort(writeBuffer, 13)); err != nil {
		return errors.Wrap(err, "Error serializing 'channelCode' field")
	}

	if popErr := writeBuffer.PopContext("ChannelInformation"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ChannelInformation")
	}
	return nil
}

func (m *_ChannelInformation) IsChannelInformation() {}

func (m *_ChannelInformation) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
