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

// ApduControl is the corresponding interface of ApduControl
type ApduControl interface {
	ApduControlContract
	ApduControlRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// IsApduControl is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduControl()
}

// ApduControlContract provides a set of functions which can be overwritten by a sub struct
type ApduControlContract interface {
	// IsApduControl is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduControl()
}

// ApduControlRequirements provides a set of functions which need to be implemented by a sub struct
type ApduControlRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetControlType returns ControlType (discriminator field)
	GetControlType() uint8
}

// _ApduControl is the data-structure of this message
type _ApduControl struct {
	_SubType ApduControl
}

var _ ApduControlContract = (*_ApduControl)(nil)

// NewApduControl factory function for _ApduControl
func NewApduControl() *_ApduControl {
	return &_ApduControl{}
}

// Deprecated: use the interface for direct cast
func CastApduControl(structType any) ApduControl {
	if casted, ok := structType.(ApduControl); ok {
		return casted
	}
	if casted, ok := structType.(*ApduControl); ok {
		return *casted
	}
	return nil
}

func (m *_ApduControl) GetTypeName() string {
	return "ApduControl"
}

func (m *_ApduControl) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (controlType)
	lengthInBits += 2

	return lengthInBits
}

func (m *_ApduControl) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func ApduControlParse[T ApduControl](ctx context.Context, theBytes []byte) (T, error) {
	return ApduControlParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func ApduControlParseWithBufferProducer[T ApduControl]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := ApduControlParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func ApduControlParseWithBuffer[T ApduControl](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_ApduControl{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_ApduControl) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__apduControl ApduControl, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduControl"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduControl")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	controlType, err := ReadDiscriminatorField[uint8](ctx, "controlType", ReadUnsignedByte(readBuffer, uint8(2)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'controlType' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child ApduControl
	switch {
	case controlType == 0x0: // ApduControlConnect
		if _child, err = (&_ApduControlConnect{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduControlConnect for type-switch of ApduControl")
		}
	case controlType == 0x1: // ApduControlDisconnect
		if _child, err = (&_ApduControlDisconnect{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduControlDisconnect for type-switch of ApduControl")
		}
	case controlType == 0x2: // ApduControlAck
		if _child, err = (&_ApduControlAck{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduControlAck for type-switch of ApduControl")
		}
	case controlType == 0x3: // ApduControlNack
		if _child, err = (&_ApduControlNack{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduControlNack for type-switch of ApduControl")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [controlType=%v]", controlType)
	}

	if closeErr := readBuffer.CloseContext("ApduControl"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduControl")
	}

	return _child, nil
}

func (pm *_ApduControl) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ApduControl, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ApduControl"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ApduControl")
	}

	if err := WriteDiscriminatorField(ctx, "controlType", m.GetControlType(), WriteUnsignedByte(writeBuffer, 2)); err != nil {
		return errors.Wrap(err, "Error serializing 'controlType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ApduControl"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ApduControl")
	}
	return nil
}

func (m *_ApduControl) IsApduControl() {}
