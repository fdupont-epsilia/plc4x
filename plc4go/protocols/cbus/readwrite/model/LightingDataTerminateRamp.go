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

// LightingDataTerminateRamp is the corresponding interface of LightingDataTerminateRamp
type LightingDataTerminateRamp interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	LightingData
	// GetGroup returns Group (property field)
	GetGroup() byte
}

// LightingDataTerminateRampExactly can be used when we want exactly this type and not a type which fulfills LightingDataTerminateRamp.
// This is useful for switch cases.
type LightingDataTerminateRampExactly interface {
	LightingDataTerminateRamp
	isLightingDataTerminateRamp() bool
}

// _LightingDataTerminateRamp is the data-structure of this message
type _LightingDataTerminateRamp struct {
	LightingDataContract
	Group byte
}

var _ LightingDataTerminateRamp = (*_LightingDataTerminateRamp)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LightingDataTerminateRamp) GetParent() LightingDataContract {
	return m.LightingDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LightingDataTerminateRamp) GetGroup() byte {
	return m.Group
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewLightingDataTerminateRamp factory function for _LightingDataTerminateRamp
func NewLightingDataTerminateRamp(group byte, commandTypeContainer LightingCommandTypeContainer) *_LightingDataTerminateRamp {
	_result := &_LightingDataTerminateRamp{
		LightingDataContract: NewLightingData(commandTypeContainer),
		Group:                group,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastLightingDataTerminateRamp(structType any) LightingDataTerminateRamp {
	if casted, ok := structType.(LightingDataTerminateRamp); ok {
		return casted
	}
	if casted, ok := structType.(*LightingDataTerminateRamp); ok {
		return *casted
	}
	return nil
}

func (m *_LightingDataTerminateRamp) GetTypeName() string {
	return "LightingDataTerminateRamp"
}

func (m *_LightingDataTerminateRamp) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.LightingDataContract.(*_LightingData).getLengthInBits(ctx))

	// Simple field (group)
	lengthInBits += 8

	return lengthInBits
}

func (m *_LightingDataTerminateRamp) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_LightingDataTerminateRamp) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_LightingData) (_lightingDataTerminateRamp LightingDataTerminateRamp, err error) {
	m.LightingDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LightingDataTerminateRamp"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LightingDataTerminateRamp")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	group, err := ReadSimpleField(ctx, "group", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'group' field"))
	}
	m.Group = group

	if closeErr := readBuffer.CloseContext("LightingDataTerminateRamp"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LightingDataTerminateRamp")
	}

	return m, nil
}

func (m *_LightingDataTerminateRamp) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LightingDataTerminateRamp) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LightingDataTerminateRamp"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LightingDataTerminateRamp")
		}

		if err := WriteSimpleField[byte](ctx, "group", m.GetGroup(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'group' field")
		}

		if popErr := writeBuffer.PopContext("LightingDataTerminateRamp"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LightingDataTerminateRamp")
		}
		return nil
	}
	return m.LightingDataContract.(*_LightingData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_LightingDataTerminateRamp) isLightingDataTerminateRamp() bool {
	return true
}

func (m *_LightingDataTerminateRamp) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
