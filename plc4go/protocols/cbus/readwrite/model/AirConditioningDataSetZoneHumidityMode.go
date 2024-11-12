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

// AirConditioningDataSetZoneHumidityMode is the corresponding interface of AirConditioningDataSetZoneHumidityMode
type AirConditioningDataSetZoneHumidityMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AirConditioningData
	// GetZoneGroup returns ZoneGroup (property field)
	GetZoneGroup() byte
	// GetZoneList returns ZoneList (property field)
	GetZoneList() HVACZoneList
	// GetHumidityModeAndFlags returns HumidityModeAndFlags (property field)
	GetHumidityModeAndFlags() HVACHumidityModeAndFlags
	// GetHumidityType returns HumidityType (property field)
	GetHumidityType() HVACHumidityType
	// GetLevel returns Level (property field)
	GetLevel() HVACHumidity
	// GetRawLevel returns RawLevel (property field)
	GetRawLevel() HVACRawLevels
	// GetAuxLevel returns AuxLevel (property field)
	GetAuxLevel() HVACAuxiliaryLevel
	// IsAirConditioningDataSetZoneHumidityMode is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAirConditioningDataSetZoneHumidityMode()
	// CreateBuilder creates a AirConditioningDataSetZoneHumidityModeBuilder
	CreateAirConditioningDataSetZoneHumidityModeBuilder() AirConditioningDataSetZoneHumidityModeBuilder
}

// _AirConditioningDataSetZoneHumidityMode is the data-structure of this message
type _AirConditioningDataSetZoneHumidityMode struct {
	AirConditioningDataContract
	ZoneGroup            byte
	ZoneList             HVACZoneList
	HumidityModeAndFlags HVACHumidityModeAndFlags
	HumidityType         HVACHumidityType
	Level                HVACHumidity
	RawLevel             HVACRawLevels
	AuxLevel             HVACAuxiliaryLevel
}

var _ AirConditioningDataSetZoneHumidityMode = (*_AirConditioningDataSetZoneHumidityMode)(nil)
var _ AirConditioningDataRequirements = (*_AirConditioningDataSetZoneHumidityMode)(nil)

// NewAirConditioningDataSetZoneHumidityMode factory function for _AirConditioningDataSetZoneHumidityMode
func NewAirConditioningDataSetZoneHumidityMode(commandTypeContainer AirConditioningCommandTypeContainer, zoneGroup byte, zoneList HVACZoneList, humidityModeAndFlags HVACHumidityModeAndFlags, humidityType HVACHumidityType, level HVACHumidity, rawLevel HVACRawLevels, auxLevel HVACAuxiliaryLevel) *_AirConditioningDataSetZoneHumidityMode {
	if zoneList == nil {
		panic("zoneList of type HVACZoneList for AirConditioningDataSetZoneHumidityMode must not be nil")
	}
	if humidityModeAndFlags == nil {
		panic("humidityModeAndFlags of type HVACHumidityModeAndFlags for AirConditioningDataSetZoneHumidityMode must not be nil")
	}
	_result := &_AirConditioningDataSetZoneHumidityMode{
		AirConditioningDataContract: NewAirConditioningData(commandTypeContainer),
		ZoneGroup:                   zoneGroup,
		ZoneList:                    zoneList,
		HumidityModeAndFlags:        humidityModeAndFlags,
		HumidityType:                humidityType,
		Level:                       level,
		RawLevel:                    rawLevel,
		AuxLevel:                    auxLevel,
	}
	_result.AirConditioningDataContract.(*_AirConditioningData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AirConditioningDataSetZoneHumidityModeBuilder is a builder for AirConditioningDataSetZoneHumidityMode
type AirConditioningDataSetZoneHumidityModeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(zoneGroup byte, zoneList HVACZoneList, humidityModeAndFlags HVACHumidityModeAndFlags, humidityType HVACHumidityType) AirConditioningDataSetZoneHumidityModeBuilder
	// WithZoneGroup adds ZoneGroup (property field)
	WithZoneGroup(byte) AirConditioningDataSetZoneHumidityModeBuilder
	// WithZoneList adds ZoneList (property field)
	WithZoneList(HVACZoneList) AirConditioningDataSetZoneHumidityModeBuilder
	// WithZoneListBuilder adds ZoneList (property field) which is build by the builder
	WithZoneListBuilder(func(HVACZoneListBuilder) HVACZoneListBuilder) AirConditioningDataSetZoneHumidityModeBuilder
	// WithHumidityModeAndFlags adds HumidityModeAndFlags (property field)
	WithHumidityModeAndFlags(HVACHumidityModeAndFlags) AirConditioningDataSetZoneHumidityModeBuilder
	// WithHumidityModeAndFlagsBuilder adds HumidityModeAndFlags (property field) which is build by the builder
	WithHumidityModeAndFlagsBuilder(func(HVACHumidityModeAndFlagsBuilder) HVACHumidityModeAndFlagsBuilder) AirConditioningDataSetZoneHumidityModeBuilder
	// WithHumidityType adds HumidityType (property field)
	WithHumidityType(HVACHumidityType) AirConditioningDataSetZoneHumidityModeBuilder
	// WithLevel adds Level (property field)
	WithOptionalLevel(HVACHumidity) AirConditioningDataSetZoneHumidityModeBuilder
	// WithOptionalLevelBuilder adds Level (property field) which is build by the builder
	WithOptionalLevelBuilder(func(HVACHumidityBuilder) HVACHumidityBuilder) AirConditioningDataSetZoneHumidityModeBuilder
	// WithRawLevel adds RawLevel (property field)
	WithOptionalRawLevel(HVACRawLevels) AirConditioningDataSetZoneHumidityModeBuilder
	// WithOptionalRawLevelBuilder adds RawLevel (property field) which is build by the builder
	WithOptionalRawLevelBuilder(func(HVACRawLevelsBuilder) HVACRawLevelsBuilder) AirConditioningDataSetZoneHumidityModeBuilder
	// WithAuxLevel adds AuxLevel (property field)
	WithOptionalAuxLevel(HVACAuxiliaryLevel) AirConditioningDataSetZoneHumidityModeBuilder
	// WithOptionalAuxLevelBuilder adds AuxLevel (property field) which is build by the builder
	WithOptionalAuxLevelBuilder(func(HVACAuxiliaryLevelBuilder) HVACAuxiliaryLevelBuilder) AirConditioningDataSetZoneHumidityModeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() AirConditioningDataBuilder
	// Build builds the AirConditioningDataSetZoneHumidityMode or returns an error if something is wrong
	Build() (AirConditioningDataSetZoneHumidityMode, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AirConditioningDataSetZoneHumidityMode
}

// NewAirConditioningDataSetZoneHumidityModeBuilder() creates a AirConditioningDataSetZoneHumidityModeBuilder
func NewAirConditioningDataSetZoneHumidityModeBuilder() AirConditioningDataSetZoneHumidityModeBuilder {
	return &_AirConditioningDataSetZoneHumidityModeBuilder{_AirConditioningDataSetZoneHumidityMode: new(_AirConditioningDataSetZoneHumidityMode)}
}

type _AirConditioningDataSetZoneHumidityModeBuilder struct {
	*_AirConditioningDataSetZoneHumidityMode

	parentBuilder *_AirConditioningDataBuilder

	err *utils.MultiError
}

var _ (AirConditioningDataSetZoneHumidityModeBuilder) = (*_AirConditioningDataSetZoneHumidityModeBuilder)(nil)

func (b *_AirConditioningDataSetZoneHumidityModeBuilder) setParent(contract AirConditioningDataContract) {
	b.AirConditioningDataContract = contract
	contract.(*_AirConditioningData)._SubType = b._AirConditioningDataSetZoneHumidityMode
}

func (b *_AirConditioningDataSetZoneHumidityModeBuilder) WithMandatoryFields(zoneGroup byte, zoneList HVACZoneList, humidityModeAndFlags HVACHumidityModeAndFlags, humidityType HVACHumidityType) AirConditioningDataSetZoneHumidityModeBuilder {
	return b.WithZoneGroup(zoneGroup).WithZoneList(zoneList).WithHumidityModeAndFlags(humidityModeAndFlags).WithHumidityType(humidityType)
}

func (b *_AirConditioningDataSetZoneHumidityModeBuilder) WithZoneGroup(zoneGroup byte) AirConditioningDataSetZoneHumidityModeBuilder {
	b.ZoneGroup = zoneGroup
	return b
}

func (b *_AirConditioningDataSetZoneHumidityModeBuilder) WithZoneList(zoneList HVACZoneList) AirConditioningDataSetZoneHumidityModeBuilder {
	b.ZoneList = zoneList
	return b
}

func (b *_AirConditioningDataSetZoneHumidityModeBuilder) WithZoneListBuilder(builderSupplier func(HVACZoneListBuilder) HVACZoneListBuilder) AirConditioningDataSetZoneHumidityModeBuilder {
	builder := builderSupplier(b.ZoneList.CreateHVACZoneListBuilder())
	var err error
	b.ZoneList, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "HVACZoneListBuilder failed"))
	}
	return b
}

func (b *_AirConditioningDataSetZoneHumidityModeBuilder) WithHumidityModeAndFlags(humidityModeAndFlags HVACHumidityModeAndFlags) AirConditioningDataSetZoneHumidityModeBuilder {
	b.HumidityModeAndFlags = humidityModeAndFlags
	return b
}

func (b *_AirConditioningDataSetZoneHumidityModeBuilder) WithHumidityModeAndFlagsBuilder(builderSupplier func(HVACHumidityModeAndFlagsBuilder) HVACHumidityModeAndFlagsBuilder) AirConditioningDataSetZoneHumidityModeBuilder {
	builder := builderSupplier(b.HumidityModeAndFlags.CreateHVACHumidityModeAndFlagsBuilder())
	var err error
	b.HumidityModeAndFlags, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "HVACHumidityModeAndFlagsBuilder failed"))
	}
	return b
}

func (b *_AirConditioningDataSetZoneHumidityModeBuilder) WithHumidityType(humidityType HVACHumidityType) AirConditioningDataSetZoneHumidityModeBuilder {
	b.HumidityType = humidityType
	return b
}

func (b *_AirConditioningDataSetZoneHumidityModeBuilder) WithOptionalLevel(level HVACHumidity) AirConditioningDataSetZoneHumidityModeBuilder {
	b.Level = level
	return b
}

func (b *_AirConditioningDataSetZoneHumidityModeBuilder) WithOptionalLevelBuilder(builderSupplier func(HVACHumidityBuilder) HVACHumidityBuilder) AirConditioningDataSetZoneHumidityModeBuilder {
	builder := builderSupplier(b.Level.CreateHVACHumidityBuilder())
	var err error
	b.Level, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "HVACHumidityBuilder failed"))
	}
	return b
}

func (b *_AirConditioningDataSetZoneHumidityModeBuilder) WithOptionalRawLevel(rawLevel HVACRawLevels) AirConditioningDataSetZoneHumidityModeBuilder {
	b.RawLevel = rawLevel
	return b
}

func (b *_AirConditioningDataSetZoneHumidityModeBuilder) WithOptionalRawLevelBuilder(builderSupplier func(HVACRawLevelsBuilder) HVACRawLevelsBuilder) AirConditioningDataSetZoneHumidityModeBuilder {
	builder := builderSupplier(b.RawLevel.CreateHVACRawLevelsBuilder())
	var err error
	b.RawLevel, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "HVACRawLevelsBuilder failed"))
	}
	return b
}

func (b *_AirConditioningDataSetZoneHumidityModeBuilder) WithOptionalAuxLevel(auxLevel HVACAuxiliaryLevel) AirConditioningDataSetZoneHumidityModeBuilder {
	b.AuxLevel = auxLevel
	return b
}

func (b *_AirConditioningDataSetZoneHumidityModeBuilder) WithOptionalAuxLevelBuilder(builderSupplier func(HVACAuxiliaryLevelBuilder) HVACAuxiliaryLevelBuilder) AirConditioningDataSetZoneHumidityModeBuilder {
	builder := builderSupplier(b.AuxLevel.CreateHVACAuxiliaryLevelBuilder())
	var err error
	b.AuxLevel, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "HVACAuxiliaryLevelBuilder failed"))
	}
	return b
}

func (b *_AirConditioningDataSetZoneHumidityModeBuilder) Build() (AirConditioningDataSetZoneHumidityMode, error) {
	if b.ZoneList == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'zoneList' not set"))
	}
	if b.HumidityModeAndFlags == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'humidityModeAndFlags' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AirConditioningDataSetZoneHumidityMode.deepCopy(), nil
}

func (b *_AirConditioningDataSetZoneHumidityModeBuilder) MustBuild() AirConditioningDataSetZoneHumidityMode {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_AirConditioningDataSetZoneHumidityModeBuilder) Done() AirConditioningDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewAirConditioningDataBuilder().(*_AirConditioningDataBuilder)
	}
	return b.parentBuilder
}

func (b *_AirConditioningDataSetZoneHumidityModeBuilder) buildForAirConditioningData() (AirConditioningData, error) {
	return b.Build()
}

func (b *_AirConditioningDataSetZoneHumidityModeBuilder) DeepCopy() any {
	_copy := b.CreateAirConditioningDataSetZoneHumidityModeBuilder().(*_AirConditioningDataSetZoneHumidityModeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAirConditioningDataSetZoneHumidityModeBuilder creates a AirConditioningDataSetZoneHumidityModeBuilder
func (b *_AirConditioningDataSetZoneHumidityMode) CreateAirConditioningDataSetZoneHumidityModeBuilder() AirConditioningDataSetZoneHumidityModeBuilder {
	if b == nil {
		return NewAirConditioningDataSetZoneHumidityModeBuilder()
	}
	return &_AirConditioningDataSetZoneHumidityModeBuilder{_AirConditioningDataSetZoneHumidityMode: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AirConditioningDataSetZoneHumidityMode) GetParent() AirConditioningDataContract {
	return m.AirConditioningDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataSetZoneHumidityMode) GetZoneGroup() byte {
	return m.ZoneGroup
}

func (m *_AirConditioningDataSetZoneHumidityMode) GetZoneList() HVACZoneList {
	return m.ZoneList
}

func (m *_AirConditioningDataSetZoneHumidityMode) GetHumidityModeAndFlags() HVACHumidityModeAndFlags {
	return m.HumidityModeAndFlags
}

func (m *_AirConditioningDataSetZoneHumidityMode) GetHumidityType() HVACHumidityType {
	return m.HumidityType
}

func (m *_AirConditioningDataSetZoneHumidityMode) GetLevel() HVACHumidity {
	return m.Level
}

func (m *_AirConditioningDataSetZoneHumidityMode) GetRawLevel() HVACRawLevels {
	return m.RawLevel
}

func (m *_AirConditioningDataSetZoneHumidityMode) GetAuxLevel() HVACAuxiliaryLevel {
	return m.AuxLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAirConditioningDataSetZoneHumidityMode(structType any) AirConditioningDataSetZoneHumidityMode {
	if casted, ok := structType.(AirConditioningDataSetZoneHumidityMode); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataSetZoneHumidityMode); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataSetZoneHumidityMode) GetTypeName() string {
	return "AirConditioningDataSetZoneHumidityMode"
}

func (m *_AirConditioningDataSetZoneHumidityMode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AirConditioningDataContract.(*_AirConditioningData).getLengthInBits(ctx))

	// Simple field (zoneGroup)
	lengthInBits += 8

	// Simple field (zoneList)
	lengthInBits += m.ZoneList.GetLengthInBits(ctx)

	// Simple field (humidityModeAndFlags)
	lengthInBits += m.HumidityModeAndFlags.GetLengthInBits(ctx)

	// Simple field (humidityType)
	lengthInBits += 8

	// Optional Field (level)
	if m.Level != nil {
		lengthInBits += m.Level.GetLengthInBits(ctx)
	}

	// Optional Field (rawLevel)
	if m.RawLevel != nil {
		lengthInBits += m.RawLevel.GetLengthInBits(ctx)
	}

	// Optional Field (auxLevel)
	if m.AuxLevel != nil {
		lengthInBits += m.AuxLevel.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_AirConditioningDataSetZoneHumidityMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AirConditioningDataSetZoneHumidityMode) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AirConditioningData) (__airConditioningDataSetZoneHumidityMode AirConditioningDataSetZoneHumidityMode, err error) {
	m.AirConditioningDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AirConditioningDataSetZoneHumidityMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataSetZoneHumidityMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zoneGroup, err := ReadSimpleField(ctx, "zoneGroup", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneGroup' field"))
	}
	m.ZoneGroup = zoneGroup

	zoneList, err := ReadSimpleField[HVACZoneList](ctx, "zoneList", ReadComplex[HVACZoneList](HVACZoneListParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneList' field"))
	}
	m.ZoneList = zoneList

	humidityModeAndFlags, err := ReadSimpleField[HVACHumidityModeAndFlags](ctx, "humidityModeAndFlags", ReadComplex[HVACHumidityModeAndFlags](HVACHumidityModeAndFlagsParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'humidityModeAndFlags' field"))
	}
	m.HumidityModeAndFlags = humidityModeAndFlags

	humidityType, err := ReadEnumField[HVACHumidityType](ctx, "humidityType", "HVACHumidityType", ReadEnum(HVACHumidityTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'humidityType' field"))
	}
	m.HumidityType = humidityType

	var level HVACHumidity
	_level, err := ReadOptionalField[HVACHumidity](ctx, "level", ReadComplex[HVACHumidity](HVACHumidityParseWithBuffer, readBuffer), humidityModeAndFlags.GetIsLevelHumidity())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'level' field"))
	}
	if _level != nil {
		level = *_level
		m.Level = level
	}

	var rawLevel HVACRawLevels
	_rawLevel, err := ReadOptionalField[HVACRawLevels](ctx, "rawLevel", ReadComplex[HVACRawLevels](HVACRawLevelsParseWithBuffer, readBuffer), humidityModeAndFlags.GetIsLevelRaw())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'rawLevel' field"))
	}
	if _rawLevel != nil {
		rawLevel = *_rawLevel
		m.RawLevel = rawLevel
	}

	var auxLevel HVACAuxiliaryLevel
	_auxLevel, err := ReadOptionalField[HVACAuxiliaryLevel](ctx, "auxLevel", ReadComplex[HVACAuxiliaryLevel](HVACAuxiliaryLevelParseWithBuffer, readBuffer), humidityModeAndFlags.GetIsAuxLevelUsed())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'auxLevel' field"))
	}
	if _auxLevel != nil {
		auxLevel = *_auxLevel
		m.AuxLevel = auxLevel
	}

	if closeErr := readBuffer.CloseContext("AirConditioningDataSetZoneHumidityMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataSetZoneHumidityMode")
	}

	return m, nil
}

func (m *_AirConditioningDataSetZoneHumidityMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AirConditioningDataSetZoneHumidityMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataSetZoneHumidityMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataSetZoneHumidityMode")
		}

		if err := WriteSimpleField[byte](ctx, "zoneGroup", m.GetZoneGroup(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneGroup' field")
		}

		if err := WriteSimpleField[HVACZoneList](ctx, "zoneList", m.GetZoneList(), WriteComplex[HVACZoneList](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneList' field")
		}

		if err := WriteSimpleField[HVACHumidityModeAndFlags](ctx, "humidityModeAndFlags", m.GetHumidityModeAndFlags(), WriteComplex[HVACHumidityModeAndFlags](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'humidityModeAndFlags' field")
		}

		if err := WriteSimpleEnumField[HVACHumidityType](ctx, "humidityType", "HVACHumidityType", m.GetHumidityType(), WriteEnum[HVACHumidityType, uint8](HVACHumidityType.GetValue, HVACHumidityType.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'humidityType' field")
		}

		if err := WriteOptionalField[HVACHumidity](ctx, "level", GetRef(m.GetLevel()), WriteComplex[HVACHumidity](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'level' field")
		}

		if err := WriteOptionalField[HVACRawLevels](ctx, "rawLevel", GetRef(m.GetRawLevel()), WriteComplex[HVACRawLevels](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'rawLevel' field")
		}

		if err := WriteOptionalField[HVACAuxiliaryLevel](ctx, "auxLevel", GetRef(m.GetAuxLevel()), WriteComplex[HVACAuxiliaryLevel](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'auxLevel' field")
		}

		if popErr := writeBuffer.PopContext("AirConditioningDataSetZoneHumidityMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataSetZoneHumidityMode")
		}
		return nil
	}
	return m.AirConditioningDataContract.(*_AirConditioningData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AirConditioningDataSetZoneHumidityMode) IsAirConditioningDataSetZoneHumidityMode() {}

func (m *_AirConditioningDataSetZoneHumidityMode) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AirConditioningDataSetZoneHumidityMode) deepCopy() *_AirConditioningDataSetZoneHumidityMode {
	if m == nil {
		return nil
	}
	_AirConditioningDataSetZoneHumidityModeCopy := &_AirConditioningDataSetZoneHumidityMode{
		m.AirConditioningDataContract.(*_AirConditioningData).deepCopy(),
		m.ZoneGroup,
		utils.DeepCopy[HVACZoneList](m.ZoneList),
		utils.DeepCopy[HVACHumidityModeAndFlags](m.HumidityModeAndFlags),
		m.HumidityType,
		utils.DeepCopy[HVACHumidity](m.Level),
		utils.DeepCopy[HVACRawLevels](m.RawLevel),
		utils.DeepCopy[HVACAuxiliaryLevel](m.AuxLevel),
	}
	_AirConditioningDataSetZoneHumidityModeCopy.AirConditioningDataContract.(*_AirConditioningData)._SubType = m
	return _AirConditioningDataSetZoneHumidityModeCopy
}

func (m *_AirConditioningDataSetZoneHumidityMode) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
