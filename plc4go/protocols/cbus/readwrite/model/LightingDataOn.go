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

// LightingDataOn is the corresponding interface of LightingDataOn
type LightingDataOn interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	LightingData
	// GetGroup returns Group (property field)
	GetGroup() byte
}

// LightingDataOnExactly can be used when we want exactly this type and not a type which fulfills LightingDataOn.
// This is useful for switch cases.
type LightingDataOnExactly interface {
	LightingDataOn
	isLightingDataOn() bool
}

// _LightingDataOn is the data-structure of this message
type _LightingDataOn struct {
	LightingDataContract
	Group byte
}

var _ LightingDataOn = (*_LightingDataOn)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LightingDataOn) GetParent() LightingDataContract {
	return m.LightingDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LightingDataOn) GetGroup() byte {
	return m.Group
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewLightingDataOn factory function for _LightingDataOn
func NewLightingDataOn(group byte, commandTypeContainer LightingCommandTypeContainer) *_LightingDataOn {
	_result := &_LightingDataOn{
		LightingDataContract: NewLightingData(commandTypeContainer),
		Group:                group,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastLightingDataOn(structType any) LightingDataOn {
	if casted, ok := structType.(LightingDataOn); ok {
		return casted
	}
	if casted, ok := structType.(*LightingDataOn); ok {
		return *casted
	}
	return nil
}

func (m *_LightingDataOn) GetTypeName() string {
	return "LightingDataOn"
}

func (m *_LightingDataOn) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.LightingDataContract.(*_LightingData).getLengthInBits(ctx))

	// Simple field (group)
	lengthInBits += 8

	return lengthInBits
}

func (m *_LightingDataOn) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_LightingDataOn) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_LightingData) (_lightingDataOn LightingDataOn, err error) {
	m.LightingDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LightingDataOn"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LightingDataOn")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	group, err := ReadSimpleField(ctx, "group", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'group' field"))
	}
	m.Group = group

	if closeErr := readBuffer.CloseContext("LightingDataOn"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LightingDataOn")
	}

	return m, nil
}

func (m *_LightingDataOn) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LightingDataOn) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LightingDataOn"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LightingDataOn")
		}

		if err := WriteSimpleField[byte](ctx, "group", m.GetGroup(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'group' field")
		}

		if popErr := writeBuffer.PopContext("LightingDataOn"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LightingDataOn")
		}
		return nil
	}
	return m.LightingDataContract.(*_LightingData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_LightingDataOn) isLightingDataOn() bool {
	return true
}

func (m *_LightingDataOn) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
