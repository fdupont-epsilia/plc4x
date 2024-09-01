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

// DeviceDescriptorType2 is the corresponding interface of DeviceDescriptorType2
type DeviceDescriptorType2 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetManufacturerId returns ManufacturerId (property field)
	GetManufacturerId() uint16
	// GetDeviceType returns DeviceType (property field)
	GetDeviceType() uint16
	// GetVersion returns Version (property field)
	GetVersion() uint8
	// GetReadSupported returns ReadSupported (property field)
	GetReadSupported() bool
	// GetWriteSupported returns WriteSupported (property field)
	GetWriteSupported() bool
	// GetLogicalTagBase returns LogicalTagBase (property field)
	GetLogicalTagBase() uint8
	// GetChannelInfo1 returns ChannelInfo1 (property field)
	GetChannelInfo1() ChannelInformation
	// GetChannelInfo2 returns ChannelInfo2 (property field)
	GetChannelInfo2() ChannelInformation
	// GetChannelInfo3 returns ChannelInfo3 (property field)
	GetChannelInfo3() ChannelInformation
	// GetChannelInfo4 returns ChannelInfo4 (property field)
	GetChannelInfo4() ChannelInformation
}

// DeviceDescriptorType2Exactly can be used when we want exactly this type and not a type which fulfills DeviceDescriptorType2.
// This is useful for switch cases.
type DeviceDescriptorType2Exactly interface {
	DeviceDescriptorType2
	isDeviceDescriptorType2() bool
}

// _DeviceDescriptorType2 is the data-structure of this message
type _DeviceDescriptorType2 struct {
	ManufacturerId uint16
	DeviceType     uint16
	Version        uint8
	ReadSupported  bool
	WriteSupported bool
	LogicalTagBase uint8
	ChannelInfo1   ChannelInformation
	ChannelInfo2   ChannelInformation
	ChannelInfo3   ChannelInformation
	ChannelInfo4   ChannelInformation
}

var _ DeviceDescriptorType2 = (*_DeviceDescriptorType2)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeviceDescriptorType2) GetManufacturerId() uint16 {
	return m.ManufacturerId
}

func (m *_DeviceDescriptorType2) GetDeviceType() uint16 {
	return m.DeviceType
}

func (m *_DeviceDescriptorType2) GetVersion() uint8 {
	return m.Version
}

func (m *_DeviceDescriptorType2) GetReadSupported() bool {
	return m.ReadSupported
}

func (m *_DeviceDescriptorType2) GetWriteSupported() bool {
	return m.WriteSupported
}

func (m *_DeviceDescriptorType2) GetLogicalTagBase() uint8 {
	return m.LogicalTagBase
}

func (m *_DeviceDescriptorType2) GetChannelInfo1() ChannelInformation {
	return m.ChannelInfo1
}

func (m *_DeviceDescriptorType2) GetChannelInfo2() ChannelInformation {
	return m.ChannelInfo2
}

func (m *_DeviceDescriptorType2) GetChannelInfo3() ChannelInformation {
	return m.ChannelInfo3
}

func (m *_DeviceDescriptorType2) GetChannelInfo4() ChannelInformation {
	return m.ChannelInfo4
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDeviceDescriptorType2 factory function for _DeviceDescriptorType2
func NewDeviceDescriptorType2(manufacturerId uint16, deviceType uint16, version uint8, readSupported bool, writeSupported bool, logicalTagBase uint8, channelInfo1 ChannelInformation, channelInfo2 ChannelInformation, channelInfo3 ChannelInformation, channelInfo4 ChannelInformation) *_DeviceDescriptorType2 {
	return &_DeviceDescriptorType2{ManufacturerId: manufacturerId, DeviceType: deviceType, Version: version, ReadSupported: readSupported, WriteSupported: writeSupported, LogicalTagBase: logicalTagBase, ChannelInfo1: channelInfo1, ChannelInfo2: channelInfo2, ChannelInfo3: channelInfo3, ChannelInfo4: channelInfo4}
}

// Deprecated: use the interface for direct cast
func CastDeviceDescriptorType2(structType any) DeviceDescriptorType2 {
	if casted, ok := structType.(DeviceDescriptorType2); ok {
		return casted
	}
	if casted, ok := structType.(*DeviceDescriptorType2); ok {
		return *casted
	}
	return nil
}

func (m *_DeviceDescriptorType2) GetTypeName() string {
	return "DeviceDescriptorType2"
}

func (m *_DeviceDescriptorType2) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (manufacturerId)
	lengthInBits += 16

	// Simple field (deviceType)
	lengthInBits += 16

	// Simple field (version)
	lengthInBits += 8

	// Simple field (readSupported)
	lengthInBits += 1

	// Simple field (writeSupported)
	lengthInBits += 1

	// Simple field (logicalTagBase)
	lengthInBits += 6

	// Simple field (channelInfo1)
	lengthInBits += m.ChannelInfo1.GetLengthInBits(ctx)

	// Simple field (channelInfo2)
	lengthInBits += m.ChannelInfo2.GetLengthInBits(ctx)

	// Simple field (channelInfo3)
	lengthInBits += m.ChannelInfo3.GetLengthInBits(ctx)

	// Simple field (channelInfo4)
	lengthInBits += m.ChannelInfo4.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_DeviceDescriptorType2) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DeviceDescriptorType2Parse(ctx context.Context, theBytes []byte) (DeviceDescriptorType2, error) {
	return DeviceDescriptorType2ParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func DeviceDescriptorType2ParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (DeviceDescriptorType2, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (DeviceDescriptorType2, error) {
		return DeviceDescriptorType2ParseWithBuffer(ctx, readBuffer)
	}
}

func DeviceDescriptorType2ParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DeviceDescriptorType2, error) {
	v, err := (&_DeviceDescriptorType2{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_DeviceDescriptorType2) parse(ctx context.Context, readBuffer utils.ReadBuffer) (_deviceDescriptorType2 DeviceDescriptorType2, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DeviceDescriptorType2"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeviceDescriptorType2")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	manufacturerId, err := ReadSimpleField(ctx, "manufacturerId", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'manufacturerId' field"))
	}
	m.ManufacturerId = manufacturerId

	deviceType, err := ReadSimpleField(ctx, "deviceType", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deviceType' field"))
	}
	m.DeviceType = deviceType

	version, err := ReadSimpleField(ctx, "version", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'version' field"))
	}
	m.Version = version

	readSupported, err := ReadSimpleField(ctx, "readSupported", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'readSupported' field"))
	}
	m.ReadSupported = readSupported

	writeSupported, err := ReadSimpleField(ctx, "writeSupported", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'writeSupported' field"))
	}
	m.WriteSupported = writeSupported

	logicalTagBase, err := ReadSimpleField(ctx, "logicalTagBase", ReadUnsignedByte(readBuffer, uint8(6)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'logicalTagBase' field"))
	}
	m.LogicalTagBase = logicalTagBase

	channelInfo1, err := ReadSimpleField[ChannelInformation](ctx, "channelInfo1", ReadComplex[ChannelInformation](ChannelInformationParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'channelInfo1' field"))
	}
	m.ChannelInfo1 = channelInfo1

	channelInfo2, err := ReadSimpleField[ChannelInformation](ctx, "channelInfo2", ReadComplex[ChannelInformation](ChannelInformationParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'channelInfo2' field"))
	}
	m.ChannelInfo2 = channelInfo2

	channelInfo3, err := ReadSimpleField[ChannelInformation](ctx, "channelInfo3", ReadComplex[ChannelInformation](ChannelInformationParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'channelInfo3' field"))
	}
	m.ChannelInfo3 = channelInfo3

	channelInfo4, err := ReadSimpleField[ChannelInformation](ctx, "channelInfo4", ReadComplex[ChannelInformation](ChannelInformationParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'channelInfo4' field"))
	}
	m.ChannelInfo4 = channelInfo4

	if closeErr := readBuffer.CloseContext("DeviceDescriptorType2"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeviceDescriptorType2")
	}

	return m, nil
}

func (m *_DeviceDescriptorType2) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DeviceDescriptorType2) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("DeviceDescriptorType2"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for DeviceDescriptorType2")
	}

	if err := WriteSimpleField[uint16](ctx, "manufacturerId", m.GetManufacturerId(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'manufacturerId' field")
	}

	if err := WriteSimpleField[uint16](ctx, "deviceType", m.GetDeviceType(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'deviceType' field")
	}

	if err := WriteSimpleField[uint8](ctx, "version", m.GetVersion(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'version' field")
	}

	if err := WriteSimpleField[bool](ctx, "readSupported", m.GetReadSupported(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'readSupported' field")
	}

	if err := WriteSimpleField[bool](ctx, "writeSupported", m.GetWriteSupported(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'writeSupported' field")
	}

	if err := WriteSimpleField[uint8](ctx, "logicalTagBase", m.GetLogicalTagBase(), WriteUnsignedByte(writeBuffer, 6)); err != nil {
		return errors.Wrap(err, "Error serializing 'logicalTagBase' field")
	}

	if err := WriteSimpleField[ChannelInformation](ctx, "channelInfo1", m.GetChannelInfo1(), WriteComplex[ChannelInformation](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'channelInfo1' field")
	}

	if err := WriteSimpleField[ChannelInformation](ctx, "channelInfo2", m.GetChannelInfo2(), WriteComplex[ChannelInformation](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'channelInfo2' field")
	}

	if err := WriteSimpleField[ChannelInformation](ctx, "channelInfo3", m.GetChannelInfo3(), WriteComplex[ChannelInformation](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'channelInfo3' field")
	}

	if err := WriteSimpleField[ChannelInformation](ctx, "channelInfo4", m.GetChannelInfo4(), WriteComplex[ChannelInformation](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'channelInfo4' field")
	}

	if popErr := writeBuffer.PopContext("DeviceDescriptorType2"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for DeviceDescriptorType2")
	}
	return nil
}

func (m *_DeviceDescriptorType2) isDeviceDescriptorType2() bool {
	return true
}

func (m *_DeviceDescriptorType2) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
