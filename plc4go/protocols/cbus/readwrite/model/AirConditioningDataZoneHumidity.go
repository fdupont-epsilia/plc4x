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

// AirConditioningDataZoneHumidity is the corresponding interface of AirConditioningDataZoneHumidity
type AirConditioningDataZoneHumidity interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AirConditioningData
	// GetZoneGroup returns ZoneGroup (property field)
	GetZoneGroup() byte
	// GetZoneList returns ZoneList (property field)
	GetZoneList() HVACZoneList
	// GetHumidity returns Humidity (property field)
	GetHumidity() HVACHumidity
	// GetSensorStatus returns SensorStatus (property field)
	GetSensorStatus() HVACSensorStatus
	// IsAirConditioningDataZoneHumidity is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAirConditioningDataZoneHumidity()
	// CreateBuilder creates a AirConditioningDataZoneHumidityBuilder
	CreateAirConditioningDataZoneHumidityBuilder() AirConditioningDataZoneHumidityBuilder
}

// _AirConditioningDataZoneHumidity is the data-structure of this message
type _AirConditioningDataZoneHumidity struct {
	AirConditioningDataContract
	ZoneGroup    byte
	ZoneList     HVACZoneList
	Humidity     HVACHumidity
	SensorStatus HVACSensorStatus
}

var _ AirConditioningDataZoneHumidity = (*_AirConditioningDataZoneHumidity)(nil)
var _ AirConditioningDataRequirements = (*_AirConditioningDataZoneHumidity)(nil)

// NewAirConditioningDataZoneHumidity factory function for _AirConditioningDataZoneHumidity
func NewAirConditioningDataZoneHumidity(commandTypeContainer AirConditioningCommandTypeContainer, zoneGroup byte, zoneList HVACZoneList, humidity HVACHumidity, sensorStatus HVACSensorStatus) *_AirConditioningDataZoneHumidity {
	if zoneList == nil {
		panic("zoneList of type HVACZoneList for AirConditioningDataZoneHumidity must not be nil")
	}
	if humidity == nil {
		panic("humidity of type HVACHumidity for AirConditioningDataZoneHumidity must not be nil")
	}
	_result := &_AirConditioningDataZoneHumidity{
		AirConditioningDataContract: NewAirConditioningData(commandTypeContainer),
		ZoneGroup:                   zoneGroup,
		ZoneList:                    zoneList,
		Humidity:                    humidity,
		SensorStatus:                sensorStatus,
	}
	_result.AirConditioningDataContract.(*_AirConditioningData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AirConditioningDataZoneHumidityBuilder is a builder for AirConditioningDataZoneHumidity
type AirConditioningDataZoneHumidityBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(zoneGroup byte, zoneList HVACZoneList, humidity HVACHumidity, sensorStatus HVACSensorStatus) AirConditioningDataZoneHumidityBuilder
	// WithZoneGroup adds ZoneGroup (property field)
	WithZoneGroup(byte) AirConditioningDataZoneHumidityBuilder
	// WithZoneList adds ZoneList (property field)
	WithZoneList(HVACZoneList) AirConditioningDataZoneHumidityBuilder
	// WithZoneListBuilder adds ZoneList (property field) which is build by the builder
	WithZoneListBuilder(func(HVACZoneListBuilder) HVACZoneListBuilder) AirConditioningDataZoneHumidityBuilder
	// WithHumidity adds Humidity (property field)
	WithHumidity(HVACHumidity) AirConditioningDataZoneHumidityBuilder
	// WithHumidityBuilder adds Humidity (property field) which is build by the builder
	WithHumidityBuilder(func(HVACHumidityBuilder) HVACHumidityBuilder) AirConditioningDataZoneHumidityBuilder
	// WithSensorStatus adds SensorStatus (property field)
	WithSensorStatus(HVACSensorStatus) AirConditioningDataZoneHumidityBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() AirConditioningDataBuilder
	// Build builds the AirConditioningDataZoneHumidity or returns an error if something is wrong
	Build() (AirConditioningDataZoneHumidity, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AirConditioningDataZoneHumidity
}

// NewAirConditioningDataZoneHumidityBuilder() creates a AirConditioningDataZoneHumidityBuilder
func NewAirConditioningDataZoneHumidityBuilder() AirConditioningDataZoneHumidityBuilder {
	return &_AirConditioningDataZoneHumidityBuilder{_AirConditioningDataZoneHumidity: new(_AirConditioningDataZoneHumidity)}
}

type _AirConditioningDataZoneHumidityBuilder struct {
	*_AirConditioningDataZoneHumidity

	parentBuilder *_AirConditioningDataBuilder

	err *utils.MultiError
}

var _ (AirConditioningDataZoneHumidityBuilder) = (*_AirConditioningDataZoneHumidityBuilder)(nil)

func (b *_AirConditioningDataZoneHumidityBuilder) setParent(contract AirConditioningDataContract) {
	b.AirConditioningDataContract = contract
	contract.(*_AirConditioningData)._SubType = b._AirConditioningDataZoneHumidity
}

func (b *_AirConditioningDataZoneHumidityBuilder) WithMandatoryFields(zoneGroup byte, zoneList HVACZoneList, humidity HVACHumidity, sensorStatus HVACSensorStatus) AirConditioningDataZoneHumidityBuilder {
	return b.WithZoneGroup(zoneGroup).WithZoneList(zoneList).WithHumidity(humidity).WithSensorStatus(sensorStatus)
}

func (b *_AirConditioningDataZoneHumidityBuilder) WithZoneGroup(zoneGroup byte) AirConditioningDataZoneHumidityBuilder {
	b.ZoneGroup = zoneGroup
	return b
}

func (b *_AirConditioningDataZoneHumidityBuilder) WithZoneList(zoneList HVACZoneList) AirConditioningDataZoneHumidityBuilder {
	b.ZoneList = zoneList
	return b
}

func (b *_AirConditioningDataZoneHumidityBuilder) WithZoneListBuilder(builderSupplier func(HVACZoneListBuilder) HVACZoneListBuilder) AirConditioningDataZoneHumidityBuilder {
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

func (b *_AirConditioningDataZoneHumidityBuilder) WithHumidity(humidity HVACHumidity) AirConditioningDataZoneHumidityBuilder {
	b.Humidity = humidity
	return b
}

func (b *_AirConditioningDataZoneHumidityBuilder) WithHumidityBuilder(builderSupplier func(HVACHumidityBuilder) HVACHumidityBuilder) AirConditioningDataZoneHumidityBuilder {
	builder := builderSupplier(b.Humidity.CreateHVACHumidityBuilder())
	var err error
	b.Humidity, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "HVACHumidityBuilder failed"))
	}
	return b
}

func (b *_AirConditioningDataZoneHumidityBuilder) WithSensorStatus(sensorStatus HVACSensorStatus) AirConditioningDataZoneHumidityBuilder {
	b.SensorStatus = sensorStatus
	return b
}

func (b *_AirConditioningDataZoneHumidityBuilder) Build() (AirConditioningDataZoneHumidity, error) {
	if b.ZoneList == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'zoneList' not set"))
	}
	if b.Humidity == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'humidity' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AirConditioningDataZoneHumidity.deepCopy(), nil
}

func (b *_AirConditioningDataZoneHumidityBuilder) MustBuild() AirConditioningDataZoneHumidity {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_AirConditioningDataZoneHumidityBuilder) Done() AirConditioningDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewAirConditioningDataBuilder().(*_AirConditioningDataBuilder)
	}
	return b.parentBuilder
}

func (b *_AirConditioningDataZoneHumidityBuilder) buildForAirConditioningData() (AirConditioningData, error) {
	return b.Build()
}

func (b *_AirConditioningDataZoneHumidityBuilder) DeepCopy() any {
	_copy := b.CreateAirConditioningDataZoneHumidityBuilder().(*_AirConditioningDataZoneHumidityBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAirConditioningDataZoneHumidityBuilder creates a AirConditioningDataZoneHumidityBuilder
func (b *_AirConditioningDataZoneHumidity) CreateAirConditioningDataZoneHumidityBuilder() AirConditioningDataZoneHumidityBuilder {
	if b == nil {
		return NewAirConditioningDataZoneHumidityBuilder()
	}
	return &_AirConditioningDataZoneHumidityBuilder{_AirConditioningDataZoneHumidity: b.deepCopy()}
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

func (m *_AirConditioningDataZoneHumidity) GetParent() AirConditioningDataContract {
	return m.AirConditioningDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataZoneHumidity) GetZoneGroup() byte {
	return m.ZoneGroup
}

func (m *_AirConditioningDataZoneHumidity) GetZoneList() HVACZoneList {
	return m.ZoneList
}

func (m *_AirConditioningDataZoneHumidity) GetHumidity() HVACHumidity {
	return m.Humidity
}

func (m *_AirConditioningDataZoneHumidity) GetSensorStatus() HVACSensorStatus {
	return m.SensorStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAirConditioningDataZoneHumidity(structType any) AirConditioningDataZoneHumidity {
	if casted, ok := structType.(AirConditioningDataZoneHumidity); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataZoneHumidity); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataZoneHumidity) GetTypeName() string {
	return "AirConditioningDataZoneHumidity"
}

func (m *_AirConditioningDataZoneHumidity) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AirConditioningDataContract.(*_AirConditioningData).getLengthInBits(ctx))

	// Simple field (zoneGroup)
	lengthInBits += 8

	// Simple field (zoneList)
	lengthInBits += m.ZoneList.GetLengthInBits(ctx)

	// Simple field (humidity)
	lengthInBits += m.Humidity.GetLengthInBits(ctx)

	// Simple field (sensorStatus)
	lengthInBits += 8

	return lengthInBits
}

func (m *_AirConditioningDataZoneHumidity) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AirConditioningDataZoneHumidity) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AirConditioningData) (__airConditioningDataZoneHumidity AirConditioningDataZoneHumidity, err error) {
	m.AirConditioningDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AirConditioningDataZoneHumidity"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataZoneHumidity")
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

	humidity, err := ReadSimpleField[HVACHumidity](ctx, "humidity", ReadComplex[HVACHumidity](HVACHumidityParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'humidity' field"))
	}
	m.Humidity = humidity

	sensorStatus, err := ReadEnumField[HVACSensorStatus](ctx, "sensorStatus", "HVACSensorStatus", ReadEnum(HVACSensorStatusByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sensorStatus' field"))
	}
	m.SensorStatus = sensorStatus

	if closeErr := readBuffer.CloseContext("AirConditioningDataZoneHumidity"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataZoneHumidity")
	}

	return m, nil
}

func (m *_AirConditioningDataZoneHumidity) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AirConditioningDataZoneHumidity) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataZoneHumidity"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataZoneHumidity")
		}

		if err := WriteSimpleField[byte](ctx, "zoneGroup", m.GetZoneGroup(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneGroup' field")
		}

		if err := WriteSimpleField[HVACZoneList](ctx, "zoneList", m.GetZoneList(), WriteComplex[HVACZoneList](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneList' field")
		}

		if err := WriteSimpleField[HVACHumidity](ctx, "humidity", m.GetHumidity(), WriteComplex[HVACHumidity](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'humidity' field")
		}

		if err := WriteSimpleEnumField[HVACSensorStatus](ctx, "sensorStatus", "HVACSensorStatus", m.GetSensorStatus(), WriteEnum[HVACSensorStatus, uint8](HVACSensorStatus.GetValue, HVACSensorStatus.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'sensorStatus' field")
		}

		if popErr := writeBuffer.PopContext("AirConditioningDataZoneHumidity"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataZoneHumidity")
		}
		return nil
	}
	return m.AirConditioningDataContract.(*_AirConditioningData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AirConditioningDataZoneHumidity) IsAirConditioningDataZoneHumidity() {}

func (m *_AirConditioningDataZoneHumidity) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AirConditioningDataZoneHumidity) deepCopy() *_AirConditioningDataZoneHumidity {
	if m == nil {
		return nil
	}
	_AirConditioningDataZoneHumidityCopy := &_AirConditioningDataZoneHumidity{
		m.AirConditioningDataContract.(*_AirConditioningData).deepCopy(),
		m.ZoneGroup,
		utils.DeepCopy[HVACZoneList](m.ZoneList),
		utils.DeepCopy[HVACHumidity](m.Humidity),
		m.SensorStatus,
	}
	_AirConditioningDataZoneHumidityCopy.AirConditioningDataContract.(*_AirConditioningData)._SubType = m
	return _AirConditioningDataZoneHumidityCopy
}

func (m *_AirConditioningDataZoneHumidity) String() string {
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
