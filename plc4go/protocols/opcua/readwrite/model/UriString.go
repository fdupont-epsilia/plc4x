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

// UriString is the corresponding interface of UriString
type UriString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// UriStringExactly can be used when we want exactly this type and not a type which fulfills UriString.
// This is useful for switch cases.
type UriStringExactly interface {
	UriString
	isUriString() bool
}

// _UriString is the data-structure of this message
type _UriString struct {
}

var _ UriString = (*_UriString)(nil)

// NewUriString factory function for _UriString
func NewUriString() *_UriString {
	return &_UriString{}
}

// Deprecated: use the interface for direct cast
func CastUriString(structType any) UriString {
	if casted, ok := structType.(UriString); ok {
		return casted
	}
	if casted, ok := structType.(*UriString); ok {
		return *casted
	}
	return nil
}

func (m *_UriString) GetTypeName() string {
	return "UriString"
}

func (m *_UriString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_UriString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func UriStringParse(ctx context.Context, theBytes []byte) (UriString, error) {
	return UriStringParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func UriStringParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (UriString, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (UriString, error) {
		return UriStringParseWithBuffer(ctx, readBuffer)
	}
}

func UriStringParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (UriString, error) {
	v, err := (&_UriString{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_UriString) parse(ctx context.Context, readBuffer utils.ReadBuffer) (_uriString UriString, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("UriString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for UriString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("UriString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for UriString")
	}

	return m, nil
}

func (m *_UriString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_UriString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("UriString"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for UriString")
	}

	if popErr := writeBuffer.PopContext("UriString"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for UriString")
	}
	return nil
}

func (m *_UriString) isUriString() bool {
	return true
}

func (m *_UriString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
