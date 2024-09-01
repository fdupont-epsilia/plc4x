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

// HVACHumidityStatusFlags is the corresponding interface of HVACHumidityStatusFlags
type HVACHumidityStatusFlags interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetExpansion returns Expansion (property field)
	GetExpansion() bool
	// GetError returns Error (property field)
	GetError() bool
	// GetBusy returns Busy (property field)
	GetBusy() bool
	// GetDamperState returns DamperState (property field)
	GetDamperState() bool
	// GetFanActive returns FanActive (property field)
	GetFanActive() bool
	// GetDehumidifyingPlant returns DehumidifyingPlant (property field)
	GetDehumidifyingPlant() bool
	// GetHumidifyingPlant returns HumidifyingPlant (property field)
	GetHumidifyingPlant() bool
	// GetIsDamperStateClosed returns IsDamperStateClosed (virtual field)
	GetIsDamperStateClosed() bool
	// GetIsDamperStateOpen returns IsDamperStateOpen (virtual field)
	GetIsDamperStateOpen() bool
}

// HVACHumidityStatusFlagsExactly can be used when we want exactly this type and not a type which fulfills HVACHumidityStatusFlags.
// This is useful for switch cases.
type HVACHumidityStatusFlagsExactly interface {
	HVACHumidityStatusFlags
	isHVACHumidityStatusFlags() bool
}

// _HVACHumidityStatusFlags is the data-structure of this message
type _HVACHumidityStatusFlags struct {
	Expansion          bool
	Error              bool
	Busy               bool
	DamperState        bool
	FanActive          bool
	DehumidifyingPlant bool
	HumidifyingPlant   bool
	// Reserved Fields
	reservedField0 *bool
}

var _ HVACHumidityStatusFlags = (*_HVACHumidityStatusFlags)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HVACHumidityStatusFlags) GetExpansion() bool {
	return m.Expansion
}

func (m *_HVACHumidityStatusFlags) GetError() bool {
	return m.Error
}

func (m *_HVACHumidityStatusFlags) GetBusy() bool {
	return m.Busy
}

func (m *_HVACHumidityStatusFlags) GetDamperState() bool {
	return m.DamperState
}

func (m *_HVACHumidityStatusFlags) GetFanActive() bool {
	return m.FanActive
}

func (m *_HVACHumidityStatusFlags) GetDehumidifyingPlant() bool {
	return m.DehumidifyingPlant
}

func (m *_HVACHumidityStatusFlags) GetHumidifyingPlant() bool {
	return m.HumidifyingPlant
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_HVACHumidityStatusFlags) GetIsDamperStateClosed() bool {
	ctx := context.Background()
	_ = ctx
	return bool(!(m.GetDamperState()))
}

func (m *_HVACHumidityStatusFlags) GetIsDamperStateOpen() bool {
	ctx := context.Background()
	_ = ctx
	return bool(m.GetDamperState())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewHVACHumidityStatusFlags factory function for _HVACHumidityStatusFlags
func NewHVACHumidityStatusFlags(expansion bool, error bool, busy bool, damperState bool, fanActive bool, dehumidifyingPlant bool, humidifyingPlant bool) *_HVACHumidityStatusFlags {
	return &_HVACHumidityStatusFlags{Expansion: expansion, Error: error, Busy: busy, DamperState: damperState, FanActive: fanActive, DehumidifyingPlant: dehumidifyingPlant, HumidifyingPlant: humidifyingPlant}
}

// Deprecated: use the interface for direct cast
func CastHVACHumidityStatusFlags(structType any) HVACHumidityStatusFlags {
	if casted, ok := structType.(HVACHumidityStatusFlags); ok {
		return casted
	}
	if casted, ok := structType.(*HVACHumidityStatusFlags); ok {
		return *casted
	}
	return nil
}

func (m *_HVACHumidityStatusFlags) GetTypeName() string {
	return "HVACHumidityStatusFlags"
}

func (m *_HVACHumidityStatusFlags) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (expansion)
	lengthInBits += 1

	// Simple field (error)
	lengthInBits += 1

	// Simple field (busy)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (damperState)
	lengthInBits += 1

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (fanActive)
	lengthInBits += 1

	// Simple field (dehumidifyingPlant)
	lengthInBits += 1

	// Simple field (humidifyingPlant)
	lengthInBits += 1

	return lengthInBits
}

func (m *_HVACHumidityStatusFlags) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func HVACHumidityStatusFlagsParse(ctx context.Context, theBytes []byte) (HVACHumidityStatusFlags, error) {
	return HVACHumidityStatusFlagsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func HVACHumidityStatusFlagsParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (HVACHumidityStatusFlags, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (HVACHumidityStatusFlags, error) {
		return HVACHumidityStatusFlagsParseWithBuffer(ctx, readBuffer)
	}
}

func HVACHumidityStatusFlagsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (HVACHumidityStatusFlags, error) {
	v, err := (&_HVACHumidityStatusFlags{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_HVACHumidityStatusFlags) parse(ctx context.Context, readBuffer utils.ReadBuffer) (_hVACHumidityStatusFlags HVACHumidityStatusFlags, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("HVACHumidityStatusFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HVACHumidityStatusFlags")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	expansion, err := ReadSimpleField(ctx, "expansion", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'expansion' field"))
	}
	m.Expansion = expansion

	error, err := ReadSimpleField(ctx, "error", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'error' field"))
	}
	m.Error = error

	busy, err := ReadSimpleField(ctx, "busy", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'busy' field"))
	}
	m.Busy = busy

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	damperState, err := ReadSimpleField(ctx, "damperState", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'damperState' field"))
	}
	m.DamperState = damperState

	isDamperStateClosed, err := ReadVirtualField[bool](ctx, "isDamperStateClosed", (*bool)(nil), !(damperState))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isDamperStateClosed' field"))
	}
	_ = isDamperStateClosed

	isDamperStateOpen, err := ReadVirtualField[bool](ctx, "isDamperStateOpen", (*bool)(nil), damperState)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isDamperStateOpen' field"))
	}
	_ = isDamperStateOpen

	fanActive, err := ReadSimpleField(ctx, "fanActive", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fanActive' field"))
	}
	m.FanActive = fanActive

	dehumidifyingPlant, err := ReadSimpleField(ctx, "dehumidifyingPlant", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dehumidifyingPlant' field"))
	}
	m.DehumidifyingPlant = dehumidifyingPlant

	humidifyingPlant, err := ReadSimpleField(ctx, "humidifyingPlant", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'humidifyingPlant' field"))
	}
	m.HumidifyingPlant = humidifyingPlant

	if closeErr := readBuffer.CloseContext("HVACHumidityStatusFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HVACHumidityStatusFlags")
	}

	return m, nil
}

func (m *_HVACHumidityStatusFlags) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HVACHumidityStatusFlags) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("HVACHumidityStatusFlags"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for HVACHumidityStatusFlags")
	}

	if err := WriteSimpleField[bool](ctx, "expansion", m.GetExpansion(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'expansion' field")
	}

	if err := WriteSimpleField[bool](ctx, "error", m.GetError(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'error' field")
	}

	if err := WriteSimpleField[bool](ctx, "busy", m.GetBusy(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'busy' field")
	}

	if err := WriteReservedField[bool](ctx, "reserved", bool(false), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 1")
	}

	if err := WriteSimpleField[bool](ctx, "damperState", m.GetDamperState(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'damperState' field")
	}
	// Virtual field
	isDamperStateClosed := m.GetIsDamperStateClosed()
	_ = isDamperStateClosed
	if _isDamperStateClosedErr := writeBuffer.WriteVirtual(ctx, "isDamperStateClosed", m.GetIsDamperStateClosed()); _isDamperStateClosedErr != nil {
		return errors.Wrap(_isDamperStateClosedErr, "Error serializing 'isDamperStateClosed' field")
	}
	// Virtual field
	isDamperStateOpen := m.GetIsDamperStateOpen()
	_ = isDamperStateOpen
	if _isDamperStateOpenErr := writeBuffer.WriteVirtual(ctx, "isDamperStateOpen", m.GetIsDamperStateOpen()); _isDamperStateOpenErr != nil {
		return errors.Wrap(_isDamperStateOpenErr, "Error serializing 'isDamperStateOpen' field")
	}

	if err := WriteSimpleField[bool](ctx, "fanActive", m.GetFanActive(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'fanActive' field")
	}

	if err := WriteSimpleField[bool](ctx, "dehumidifyingPlant", m.GetDehumidifyingPlant(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'dehumidifyingPlant' field")
	}

	if err := WriteSimpleField[bool](ctx, "humidifyingPlant", m.GetHumidifyingPlant(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'humidifyingPlant' field")
	}

	if popErr := writeBuffer.PopContext("HVACHumidityStatusFlags"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for HVACHumidityStatusFlags")
	}
	return nil
}

func (m *_HVACHumidityStatusFlags) isHVACHumidityStatusFlags() bool {
	return true
}

func (m *_HVACHumidityStatusFlags) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
