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

// ClockAndTimekeepingDataRequestRefresh is the corresponding interface of ClockAndTimekeepingDataRequestRefresh
type ClockAndTimekeepingDataRequestRefresh interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ClockAndTimekeepingData
	// IsClockAndTimekeepingDataRequestRefresh is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsClockAndTimekeepingDataRequestRefresh()
}

// _ClockAndTimekeepingDataRequestRefresh is the data-structure of this message
type _ClockAndTimekeepingDataRequestRefresh struct {
	ClockAndTimekeepingDataContract
}

var _ ClockAndTimekeepingDataRequestRefresh = (*_ClockAndTimekeepingDataRequestRefresh)(nil)
var _ ClockAndTimekeepingDataRequirements = (*_ClockAndTimekeepingDataRequestRefresh)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ClockAndTimekeepingDataRequestRefresh) GetParent() ClockAndTimekeepingDataContract {
	return m.ClockAndTimekeepingDataContract
}

// NewClockAndTimekeepingDataRequestRefresh factory function for _ClockAndTimekeepingDataRequestRefresh
func NewClockAndTimekeepingDataRequestRefresh(commandTypeContainer ClockAndTimekeepingCommandTypeContainer, argument byte) *_ClockAndTimekeepingDataRequestRefresh {
	_result := &_ClockAndTimekeepingDataRequestRefresh{
		ClockAndTimekeepingDataContract: NewClockAndTimekeepingData(commandTypeContainer, argument),
	}
	_result.ClockAndTimekeepingDataContract.(*_ClockAndTimekeepingData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastClockAndTimekeepingDataRequestRefresh(structType any) ClockAndTimekeepingDataRequestRefresh {
	if casted, ok := structType.(ClockAndTimekeepingDataRequestRefresh); ok {
		return casted
	}
	if casted, ok := structType.(*ClockAndTimekeepingDataRequestRefresh); ok {
		return *casted
	}
	return nil
}

func (m *_ClockAndTimekeepingDataRequestRefresh) GetTypeName() string {
	return "ClockAndTimekeepingDataRequestRefresh"
}

func (m *_ClockAndTimekeepingDataRequestRefresh) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ClockAndTimekeepingDataContract.(*_ClockAndTimekeepingData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ClockAndTimekeepingDataRequestRefresh) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ClockAndTimekeepingDataRequestRefresh) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ClockAndTimekeepingData) (__clockAndTimekeepingDataRequestRefresh ClockAndTimekeepingDataRequestRefresh, err error) {
	m.ClockAndTimekeepingDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ClockAndTimekeepingDataRequestRefresh"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ClockAndTimekeepingDataRequestRefresh")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ClockAndTimekeepingDataRequestRefresh"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ClockAndTimekeepingDataRequestRefresh")
	}

	return m, nil
}

func (m *_ClockAndTimekeepingDataRequestRefresh) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ClockAndTimekeepingDataRequestRefresh) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ClockAndTimekeepingDataRequestRefresh"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ClockAndTimekeepingDataRequestRefresh")
		}

		if popErr := writeBuffer.PopContext("ClockAndTimekeepingDataRequestRefresh"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ClockAndTimekeepingDataRequestRefresh")
		}
		return nil
	}
	return m.ClockAndTimekeepingDataContract.(*_ClockAndTimekeepingData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ClockAndTimekeepingDataRequestRefresh) IsClockAndTimekeepingDataRequestRefresh() {}

func (m *_ClockAndTimekeepingDataRequestRefresh) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
