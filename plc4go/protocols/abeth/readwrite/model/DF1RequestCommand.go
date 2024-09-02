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

// DF1RequestCommand is the corresponding interface of DF1RequestCommand
type DF1RequestCommand interface {
	DF1RequestCommandContract
	DF1RequestCommandRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// IsDF1RequestCommand is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDF1RequestCommand()
}

// DF1RequestCommandContract provides a set of functions which can be overwritten by a sub struct
type DF1RequestCommandContract interface {
	// IsDF1RequestCommand is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDF1RequestCommand()
}

// DF1RequestCommandRequirements provides a set of functions which need to be implemented by a sub struct
type DF1RequestCommandRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetFunctionCode returns FunctionCode (discriminator field)
	GetFunctionCode() uint8
}

// _DF1RequestCommand is the data-structure of this message
type _DF1RequestCommand struct {
	_SubType DF1RequestCommand
}

var _ DF1RequestCommandContract = (*_DF1RequestCommand)(nil)

// NewDF1RequestCommand factory function for _DF1RequestCommand
func NewDF1RequestCommand() *_DF1RequestCommand {
	return &_DF1RequestCommand{}
}

// Deprecated: use the interface for direct cast
func CastDF1RequestCommand(structType any) DF1RequestCommand {
	if casted, ok := structType.(DF1RequestCommand); ok {
		return casted
	}
	if casted, ok := structType.(*DF1RequestCommand); ok {
		return *casted
	}
	return nil
}

func (m *_DF1RequestCommand) GetTypeName() string {
	return "DF1RequestCommand"
}

func (m *_DF1RequestCommand) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (functionCode)
	lengthInBits += 8

	return lengthInBits
}

func (m *_DF1RequestCommand) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func DF1RequestCommandParse[T DF1RequestCommand](ctx context.Context, theBytes []byte) (T, error) {
	return DF1RequestCommandParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func DF1RequestCommandParseWithBufferProducer[T DF1RequestCommand]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := DF1RequestCommandParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func DF1RequestCommandParseWithBuffer[T DF1RequestCommand](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_DF1RequestCommand{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_DF1RequestCommand) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__dF1RequestCommand DF1RequestCommand, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DF1RequestCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DF1RequestCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	functionCode, err := ReadDiscriminatorField[uint8](ctx, "functionCode", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'functionCode' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child DF1RequestCommand
	switch {
	case functionCode == 0xA2: // DF1RequestProtectedTypedLogicalRead
		if _child, err = (&_DF1RequestProtectedTypedLogicalRead{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type DF1RequestProtectedTypedLogicalRead for type-switch of DF1RequestCommand")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [functionCode=%v]", functionCode)
	}

	if closeErr := readBuffer.CloseContext("DF1RequestCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DF1RequestCommand")
	}

	return _child, nil
}

func (pm *_DF1RequestCommand) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child DF1RequestCommand, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("DF1RequestCommand"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for DF1RequestCommand")
	}

	if err := WriteDiscriminatorField(ctx, "functionCode", m.GetFunctionCode(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'functionCode' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("DF1RequestCommand"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for DF1RequestCommand")
	}
	return nil
}

func (m *_DF1RequestCommand) IsDF1RequestCommand() {}
