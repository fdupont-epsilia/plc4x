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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// TrimmedString is the corresponding interface of TrimmedString
type TrimmedString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// TrimmedStringExactly can be used when we want exactly this type and not a type which fulfills TrimmedString.
// This is useful for switch cases.
type TrimmedStringExactly interface {
	TrimmedString
	isTrimmedString() bool
}

// _TrimmedString is the data-structure of this message
type _TrimmedString struct {
}

var _ TrimmedString = (*_TrimmedString)(nil)

// NewTrimmedString factory function for _TrimmedString
func NewTrimmedString() *_TrimmedString {
	return &_TrimmedString{}
}

// Deprecated: use the interface for direct cast
func CastTrimmedString(structType any) TrimmedString {
	if casted, ok := structType.(TrimmedString); ok {
		return casted
	}
	if casted, ok := structType.(*TrimmedString); ok {
		return *casted
	}
	return nil
}

func (m *_TrimmedString) GetTypeName() string {
	return "TrimmedString"
}

func (m *_TrimmedString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_TrimmedString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TrimmedStringParse(ctx context.Context, theBytes []byte) (TrimmedString, error) {
	return TrimmedStringParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func TrimmedStringParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (TrimmedString, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (TrimmedString, error) {
		return TrimmedStringParseWithBuffer(ctx, readBuffer)
	}
}

func TrimmedStringParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TrimmedString, error) {
	v, err := (&_TrimmedString{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_TrimmedString) parse(ctx context.Context, readBuffer utils.ReadBuffer) (_trimmedString TrimmedString, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TrimmedString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TrimmedString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("TrimmedString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TrimmedString")
	}

	return m, nil
}

func (m *_TrimmedString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TrimmedString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("TrimmedString"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for TrimmedString")
	}

	if popErr := writeBuffer.PopContext("TrimmedString"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for TrimmedString")
	}
	return nil
}

func (m *_TrimmedString) isTrimmedString() bool {
	return true
}

func (m *_TrimmedString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
