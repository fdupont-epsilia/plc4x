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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// Handle is the corresponding interface of Handle
type Handle interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// HandleExactly can be used when we want exactly this type and not a type which fulfills Handle.
// This is useful for switch cases.
type HandleExactly interface {
	Handle
	isHandle() bool
}

// _Handle is the data-structure of this message
type _Handle struct {
}

var _ Handle = (*_Handle)(nil)

// NewHandle factory function for _Handle
func NewHandle() *_Handle {
	return &_Handle{}
}

// Deprecated: use the interface for direct cast
func CastHandle(structType any) Handle {
	if casted, ok := structType.(Handle); ok {
		return casted
	}
	if casted, ok := structType.(*Handle); ok {
		return *casted
	}
	return nil
}

func (m *_Handle) GetTypeName() string {
	return "Handle"
}

func (m *_Handle) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_Handle) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func HandleParse(ctx context.Context, theBytes []byte) (Handle, error) {
	return HandleParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func HandleParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (Handle, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (Handle, error) {
		return HandleParseWithBuffer(ctx, readBuffer)
	}
}

func HandleParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (Handle, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("Handle"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Handle")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("Handle"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Handle")
	}

	// Create the instance
	return &_Handle{}, nil
}

func (m *_Handle) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_Handle) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("Handle"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for Handle")
	}

	if popErr := writeBuffer.PopContext("Handle"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for Handle")
	}
	return nil
}

func (m *_Handle) isHandle() bool {
	return true
}

func (m *_Handle) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
