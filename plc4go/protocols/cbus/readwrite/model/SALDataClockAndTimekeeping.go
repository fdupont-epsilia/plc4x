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

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// SALDataClockAndTimekeeping is the corresponding interface of SALDataClockAndTimekeeping
type SALDataClockAndTimekeeping interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SALData
	// GetClockAndTimekeepingData returns ClockAndTimekeepingData (property field)
	GetClockAndTimekeepingData() ClockAndTimekeepingData
}

// SALDataClockAndTimekeepingExactly can be used when we want exactly this type and not a type which fulfills SALDataClockAndTimekeeping.
// This is useful for switch cases.
type SALDataClockAndTimekeepingExactly interface {
	SALDataClockAndTimekeeping
	isSALDataClockAndTimekeeping() bool
}

// _SALDataClockAndTimekeeping is the data-structure of this message
type _SALDataClockAndTimekeeping struct {
	*_SALData
	ClockAndTimekeepingData ClockAndTimekeepingData
}

var _ SALDataClockAndTimekeeping = (*_SALDataClockAndTimekeeping)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataClockAndTimekeeping) GetApplicationId() ApplicationId {
	return ApplicationId_CLOCK_AND_TIMEKEEPING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataClockAndTimekeeping) InitializeParent(parent SALData, salData SALData) {
	m.SalData = salData
}

func (m *_SALDataClockAndTimekeeping) GetParent() SALData {
	return m._SALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataClockAndTimekeeping) GetClockAndTimekeepingData() ClockAndTimekeepingData {
	return m.ClockAndTimekeepingData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSALDataClockAndTimekeeping factory function for _SALDataClockAndTimekeeping
func NewSALDataClockAndTimekeeping(clockAndTimekeepingData ClockAndTimekeepingData, salData SALData) *_SALDataClockAndTimekeeping {
	_result := &_SALDataClockAndTimekeeping{
		ClockAndTimekeepingData: clockAndTimekeepingData,
		_SALData:                NewSALData(salData),
	}
	_result._SALData._SALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSALDataClockAndTimekeeping(structType any) SALDataClockAndTimekeeping {
	if casted, ok := structType.(SALDataClockAndTimekeeping); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataClockAndTimekeeping); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataClockAndTimekeeping) GetTypeName() string {
	return "SALDataClockAndTimekeeping"
}

func (m *_SALDataClockAndTimekeeping) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (clockAndTimekeepingData)
	lengthInBits += m.ClockAndTimekeepingData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SALDataClockAndTimekeeping) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SALDataClockAndTimekeepingParse(ctx context.Context, theBytes []byte, applicationId ApplicationId) (SALDataClockAndTimekeeping, error) {
	return SALDataClockAndTimekeepingParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), applicationId)
}

func SALDataClockAndTimekeepingParseWithBufferProducer(applicationId ApplicationId) func(ctx context.Context, readBuffer utils.ReadBuffer) (SALDataClockAndTimekeeping, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (SALDataClockAndTimekeeping, error) {
		return SALDataClockAndTimekeepingParseWithBuffer(ctx, readBuffer, applicationId)
	}
}

func SALDataClockAndTimekeepingParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, applicationId ApplicationId) (SALDataClockAndTimekeeping, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataClockAndTimekeeping"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataClockAndTimekeeping")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	clockAndTimekeepingData, err := ReadSimpleField[ClockAndTimekeepingData](ctx, "clockAndTimekeepingData", ReadComplex[ClockAndTimekeepingData](ClockAndTimekeepingDataParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'clockAndTimekeepingData' field"))
	}

	if closeErr := readBuffer.CloseContext("SALDataClockAndTimekeeping"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataClockAndTimekeeping")
	}

	// Create a partially initialized instance
	_child := &_SALDataClockAndTimekeeping{
		_SALData:                &_SALData{},
		ClockAndTimekeepingData: clockAndTimekeepingData,
	}
	_child._SALData._SALDataChildRequirements = _child
	return _child, nil
}

func (m *_SALDataClockAndTimekeeping) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataClockAndTimekeeping) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataClockAndTimekeeping"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataClockAndTimekeeping")
		}

		if err := WriteSimpleField[ClockAndTimekeepingData](ctx, "clockAndTimekeepingData", m.GetClockAndTimekeepingData(), WriteComplex[ClockAndTimekeepingData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'clockAndTimekeepingData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataClockAndTimekeeping"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataClockAndTimekeeping")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SALDataClockAndTimekeeping) isSALDataClockAndTimekeeping() bool {
	return true
}

func (m *_SALDataClockAndTimekeeping) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
