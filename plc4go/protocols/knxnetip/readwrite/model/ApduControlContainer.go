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

// ApduControlContainer is the corresponding interface of ApduControlContainer
type ApduControlContainer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	Apdu
	// GetControlApdu returns ControlApdu (property field)
	GetControlApdu() ApduControl
	// IsApduControlContainer is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduControlContainer()
}

// _ApduControlContainer is the data-structure of this message
type _ApduControlContainer struct {
	ApduContract
	ControlApdu ApduControl
}

var _ ApduControlContainer = (*_ApduControlContainer)(nil)
var _ ApduRequirements = (*_ApduControlContainer)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduControlContainer) GetControl() uint8 {
	return uint8(1)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduControlContainer) GetParent() ApduContract {
	return m.ApduContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ApduControlContainer) GetControlApdu() ApduControl {
	return m.ControlApdu
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewApduControlContainer factory function for _ApduControlContainer
func NewApduControlContainer(controlApdu ApduControl, numbered bool, counter uint8, dataLength uint8) *_ApduControlContainer {
	_result := &_ApduControlContainer{
		ApduContract: NewApdu(numbered, counter, dataLength),
		ControlApdu:  controlApdu,
	}
	_result.ApduContract.(*_Apdu)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduControlContainer(structType any) ApduControlContainer {
	if casted, ok := structType.(ApduControlContainer); ok {
		return casted
	}
	if casted, ok := structType.(*ApduControlContainer); ok {
		return *casted
	}
	return nil
}

func (m *_ApduControlContainer) GetTypeName() string {
	return "ApduControlContainer"
}

func (m *_ApduControlContainer) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduContract.(*_Apdu).getLengthInBits(ctx))

	// Simple field (controlApdu)
	lengthInBits += m.ControlApdu.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ApduControlContainer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduControlContainer) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_Apdu, dataLength uint8) (__apduControlContainer ApduControlContainer, err error) {
	m.ApduContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduControlContainer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduControlContainer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	controlApdu, err := ReadSimpleField[ApduControl](ctx, "controlApdu", ReadComplex[ApduControl](ApduControlParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'controlApdu' field"))
	}
	m.ControlApdu = controlApdu

	if closeErr := readBuffer.CloseContext("ApduControlContainer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduControlContainer")
	}

	return m, nil
}

func (m *_ApduControlContainer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduControlContainer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduControlContainer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduControlContainer")
		}

		if err := WriteSimpleField[ApduControl](ctx, "controlApdu", m.GetControlApdu(), WriteComplex[ApduControl](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'controlApdu' field")
		}

		if popErr := writeBuffer.PopContext("ApduControlContainer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduControlContainer")
		}
		return nil
	}
	return m.ApduContract.(*_Apdu).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduControlContainer) IsApduControlContainer() {}

func (m *_ApduControlContainer) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
