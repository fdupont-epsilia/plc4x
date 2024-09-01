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

// TelephonyDataRinging is the corresponding interface of TelephonyDataRinging
type TelephonyDataRinging interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	TelephonyData
	// GetNumber returns Number (property field)
	GetNumber() string
}

// TelephonyDataRingingExactly can be used when we want exactly this type and not a type which fulfills TelephonyDataRinging.
// This is useful for switch cases.
type TelephonyDataRingingExactly interface {
	TelephonyDataRinging
	isTelephonyDataRinging() bool
}

// _TelephonyDataRinging is the data-structure of this message
type _TelephonyDataRinging struct {
	TelephonyDataContract
	Number string
	// Reserved Fields
	reservedField0 *byte
}

var _ TelephonyDataRinging = (*_TelephonyDataRinging)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TelephonyDataRinging) GetParent() TelephonyDataContract {
	return m.TelephonyDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TelephonyDataRinging) GetNumber() string {
	return m.Number
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTelephonyDataRinging factory function for _TelephonyDataRinging
func NewTelephonyDataRinging(number string, commandTypeContainer TelephonyCommandTypeContainer, argument byte) *_TelephonyDataRinging {
	_result := &_TelephonyDataRinging{
		TelephonyDataContract: NewTelephonyData(commandTypeContainer, argument),
		Number:                number,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastTelephonyDataRinging(structType any) TelephonyDataRinging {
	if casted, ok := structType.(TelephonyDataRinging); ok {
		return casted
	}
	if casted, ok := structType.(*TelephonyDataRinging); ok {
		return *casted
	}
	return nil
}

func (m *_TelephonyDataRinging) GetTypeName() string {
	return "TelephonyDataRinging"
}

func (m *_TelephonyDataRinging) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.TelephonyDataContract.(*_TelephonyData).getLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (number)
	lengthInBits += uint16(int32((int32(m.GetCommandTypeContainer().NumBytes()) - int32(int32(2)))) * int32(int32(8)))

	return lengthInBits
}

func (m *_TelephonyDataRinging) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_TelephonyDataRinging) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_TelephonyData, commandTypeContainer TelephonyCommandTypeContainer) (_telephonyDataRinging TelephonyDataRinging, err error) {
	m.TelephonyDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TelephonyDataRinging"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TelephonyDataRinging")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadByte(readBuffer, 8), byte(0x01))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	number, err := ReadSimpleField(ctx, "number", ReadString(readBuffer, uint32(int32((int32(commandTypeContainer.NumBytes())-int32(int32(2))))*int32(int32(8)))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'number' field"))
	}
	m.Number = number

	if closeErr := readBuffer.CloseContext("TelephonyDataRinging"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TelephonyDataRinging")
	}

	return m, nil
}

func (m *_TelephonyDataRinging) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TelephonyDataRinging) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TelephonyDataRinging"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TelephonyDataRinging")
		}

		if err := WriteReservedField[byte](ctx, "reserved", byte(0x01), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[string](ctx, "number", m.GetNumber(), WriteString(writeBuffer, int32(int32((int32(m.GetCommandTypeContainer().NumBytes())-int32(int32(2))))*int32(int32(8))))); err != nil {
			return errors.Wrap(err, "Error serializing 'number' field")
		}

		if popErr := writeBuffer.PopContext("TelephonyDataRinging"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TelephonyDataRinging")
		}
		return nil
	}
	return m.TelephonyDataContract.(*_TelephonyData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TelephonyDataRinging) isTelephonyDataRinging() bool {
	return true
}

func (m *_TelephonyDataRinging) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
