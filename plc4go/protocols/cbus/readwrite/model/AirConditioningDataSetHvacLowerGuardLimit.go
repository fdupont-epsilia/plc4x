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

// AirConditioningDataSetHvacLowerGuardLimit is the corresponding interface of AirConditioningDataSetHvacLowerGuardLimit
type AirConditioningDataSetHvacLowerGuardLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AirConditioningData
	// GetZoneGroup returns ZoneGroup (property field)
	GetZoneGroup() byte
	// GetZoneList returns ZoneList (property field)
	GetZoneList() HVACZoneList
	// GetLimit returns Limit (property field)
	GetLimit() HVACTemperature
	// GetHvacModeAndFlags returns HvacModeAndFlags (property field)
	GetHvacModeAndFlags() HVACModeAndFlags
}

// AirConditioningDataSetHvacLowerGuardLimitExactly can be used when we want exactly this type and not a type which fulfills AirConditioningDataSetHvacLowerGuardLimit.
// This is useful for switch cases.
type AirConditioningDataSetHvacLowerGuardLimitExactly interface {
	AirConditioningDataSetHvacLowerGuardLimit
	isAirConditioningDataSetHvacLowerGuardLimit() bool
}

// _AirConditioningDataSetHvacLowerGuardLimit is the data-structure of this message
type _AirConditioningDataSetHvacLowerGuardLimit struct {
	*_AirConditioningData
	ZoneGroup        byte
	ZoneList         HVACZoneList
	Limit            HVACTemperature
	HvacModeAndFlags HVACModeAndFlags
}

var _ AirConditioningDataSetHvacLowerGuardLimit = (*_AirConditioningDataSetHvacLowerGuardLimit)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AirConditioningDataSetHvacLowerGuardLimit) InitializeParent(parent AirConditioningData, commandTypeContainer AirConditioningCommandTypeContainer) {
	m.CommandTypeContainer = commandTypeContainer
}

func (m *_AirConditioningDataSetHvacLowerGuardLimit) GetParent() AirConditioningData {
	return m._AirConditioningData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataSetHvacLowerGuardLimit) GetZoneGroup() byte {
	return m.ZoneGroup
}

func (m *_AirConditioningDataSetHvacLowerGuardLimit) GetZoneList() HVACZoneList {
	return m.ZoneList
}

func (m *_AirConditioningDataSetHvacLowerGuardLimit) GetLimit() HVACTemperature {
	return m.Limit
}

func (m *_AirConditioningDataSetHvacLowerGuardLimit) GetHvacModeAndFlags() HVACModeAndFlags {
	return m.HvacModeAndFlags
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAirConditioningDataSetHvacLowerGuardLimit factory function for _AirConditioningDataSetHvacLowerGuardLimit
func NewAirConditioningDataSetHvacLowerGuardLimit(zoneGroup byte, zoneList HVACZoneList, limit HVACTemperature, hvacModeAndFlags HVACModeAndFlags, commandTypeContainer AirConditioningCommandTypeContainer) *_AirConditioningDataSetHvacLowerGuardLimit {
	_result := &_AirConditioningDataSetHvacLowerGuardLimit{
		ZoneGroup:            zoneGroup,
		ZoneList:             zoneList,
		Limit:                limit,
		HvacModeAndFlags:     hvacModeAndFlags,
		_AirConditioningData: NewAirConditioningData(commandTypeContainer),
	}
	_result._AirConditioningData._AirConditioningDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAirConditioningDataSetHvacLowerGuardLimit(structType any) AirConditioningDataSetHvacLowerGuardLimit {
	if casted, ok := structType.(AirConditioningDataSetHvacLowerGuardLimit); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataSetHvacLowerGuardLimit); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataSetHvacLowerGuardLimit) GetTypeName() string {
	return "AirConditioningDataSetHvacLowerGuardLimit"
}

func (m *_AirConditioningDataSetHvacLowerGuardLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (zoneGroup)
	lengthInBits += 8

	// Simple field (zoneList)
	lengthInBits += m.ZoneList.GetLengthInBits(ctx)

	// Simple field (limit)
	lengthInBits += m.Limit.GetLengthInBits(ctx)

	// Simple field (hvacModeAndFlags)
	lengthInBits += m.HvacModeAndFlags.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_AirConditioningDataSetHvacLowerGuardLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AirConditioningDataSetHvacLowerGuardLimitParse(ctx context.Context, theBytes []byte) (AirConditioningDataSetHvacLowerGuardLimit, error) {
	return AirConditioningDataSetHvacLowerGuardLimitParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AirConditioningDataSetHvacLowerGuardLimitParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (AirConditioningDataSetHvacLowerGuardLimit, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (AirConditioningDataSetHvacLowerGuardLimit, error) {
		return AirConditioningDataSetHvacLowerGuardLimitParseWithBuffer(ctx, readBuffer)
	}
}

func AirConditioningDataSetHvacLowerGuardLimitParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AirConditioningDataSetHvacLowerGuardLimit, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AirConditioningDataSetHvacLowerGuardLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataSetHvacLowerGuardLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zoneGroup, err := ReadSimpleField(ctx, "zoneGroup", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneGroup' field"))
	}

	zoneList, err := ReadSimpleField[HVACZoneList](ctx, "zoneList", ReadComplex[HVACZoneList](HVACZoneListParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneList' field"))
	}

	limit, err := ReadSimpleField[HVACTemperature](ctx, "limit", ReadComplex[HVACTemperature](HVACTemperatureParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'limit' field"))
	}

	hvacModeAndFlags, err := ReadSimpleField[HVACModeAndFlags](ctx, "hvacModeAndFlags", ReadComplex[HVACModeAndFlags](HVACModeAndFlagsParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hvacModeAndFlags' field"))
	}

	if closeErr := readBuffer.CloseContext("AirConditioningDataSetHvacLowerGuardLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataSetHvacLowerGuardLimit")
	}

	// Create a partially initialized instance
	_child := &_AirConditioningDataSetHvacLowerGuardLimit{
		_AirConditioningData: &_AirConditioningData{},
		ZoneGroup:            zoneGroup,
		ZoneList:             zoneList,
		Limit:                limit,
		HvacModeAndFlags:     hvacModeAndFlags,
	}
	_child._AirConditioningData._AirConditioningDataChildRequirements = _child
	return _child, nil
}

func (m *_AirConditioningDataSetHvacLowerGuardLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AirConditioningDataSetHvacLowerGuardLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataSetHvacLowerGuardLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataSetHvacLowerGuardLimit")
		}

		if err := WriteSimpleField[byte](ctx, "zoneGroup", m.GetZoneGroup(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneGroup' field")
		}

		if err := WriteSimpleField[HVACZoneList](ctx, "zoneList", m.GetZoneList(), WriteComplex[HVACZoneList](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneList' field")
		}

		if err := WriteSimpleField[HVACTemperature](ctx, "limit", m.GetLimit(), WriteComplex[HVACTemperature](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'limit' field")
		}

		if err := WriteSimpleField[HVACModeAndFlags](ctx, "hvacModeAndFlags", m.GetHvacModeAndFlags(), WriteComplex[HVACModeAndFlags](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'hvacModeAndFlags' field")
		}

		if popErr := writeBuffer.PopContext("AirConditioningDataSetHvacLowerGuardLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataSetHvacLowerGuardLimit")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AirConditioningDataSetHvacLowerGuardLimit) isAirConditioningDataSetHvacLowerGuardLimit() bool {
	return true
}

func (m *_AirConditioningDataSetHvacLowerGuardLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
