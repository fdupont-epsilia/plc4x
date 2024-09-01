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

// AdsAddDeviceNotificationRequest is the corresponding interface of AdsAddDeviceNotificationRequest
type AdsAddDeviceNotificationRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AmsPacket
	// GetIndexGroup returns IndexGroup (property field)
	GetIndexGroup() uint32
	// GetIndexOffset returns IndexOffset (property field)
	GetIndexOffset() uint32
	// GetLength returns Length (property field)
	GetLength() uint32
	// GetTransmissionMode returns TransmissionMode (property field)
	GetTransmissionMode() AdsTransMode
	// GetMaxDelayInMs returns MaxDelayInMs (property field)
	GetMaxDelayInMs() uint32
	// GetCycleTimeInMs returns CycleTimeInMs (property field)
	GetCycleTimeInMs() uint32
}

// AdsAddDeviceNotificationRequestExactly can be used when we want exactly this type and not a type which fulfills AdsAddDeviceNotificationRequest.
// This is useful for switch cases.
type AdsAddDeviceNotificationRequestExactly interface {
	AdsAddDeviceNotificationRequest
	isAdsAddDeviceNotificationRequest() bool
}

// _AdsAddDeviceNotificationRequest is the data-structure of this message
type _AdsAddDeviceNotificationRequest struct {
	AmsPacketContract
	IndexGroup       uint32
	IndexOffset      uint32
	Length           uint32
	TransmissionMode AdsTransMode
	MaxDelayInMs     uint32
	CycleTimeInMs    uint32
	// Reserved Fields
	reservedField0 *uint64
	reservedField1 *uint64
}

var _ AdsAddDeviceNotificationRequest = (*_AdsAddDeviceNotificationRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsAddDeviceNotificationRequest) GetCommandId() CommandId {
	return CommandId_ADS_ADD_DEVICE_NOTIFICATION
}

func (m *_AdsAddDeviceNotificationRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsAddDeviceNotificationRequest) GetParent() AmsPacketContract {
	return m.AmsPacketContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsAddDeviceNotificationRequest) GetIndexGroup() uint32 {
	return m.IndexGroup
}

func (m *_AdsAddDeviceNotificationRequest) GetIndexOffset() uint32 {
	return m.IndexOffset
}

func (m *_AdsAddDeviceNotificationRequest) GetLength() uint32 {
	return m.Length
}

func (m *_AdsAddDeviceNotificationRequest) GetTransmissionMode() AdsTransMode {
	return m.TransmissionMode
}

func (m *_AdsAddDeviceNotificationRequest) GetMaxDelayInMs() uint32 {
	return m.MaxDelayInMs
}

func (m *_AdsAddDeviceNotificationRequest) GetCycleTimeInMs() uint32 {
	return m.CycleTimeInMs
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsAddDeviceNotificationRequest factory function for _AdsAddDeviceNotificationRequest
func NewAdsAddDeviceNotificationRequest(indexGroup uint32, indexOffset uint32, length uint32, transmissionMode AdsTransMode, maxDelayInMs uint32, cycleTimeInMs uint32, targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32) *_AdsAddDeviceNotificationRequest {
	_result := &_AdsAddDeviceNotificationRequest{
		AmsPacketContract: NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId),
		IndexGroup:        indexGroup,
		IndexOffset:       indexOffset,
		Length:            length,
		TransmissionMode:  transmissionMode,
		MaxDelayInMs:      maxDelayInMs,
		CycleTimeInMs:     cycleTimeInMs,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsAddDeviceNotificationRequest(structType any) AdsAddDeviceNotificationRequest {
	if casted, ok := structType.(AdsAddDeviceNotificationRequest); ok {
		return casted
	}
	if casted, ok := structType.(*AdsAddDeviceNotificationRequest); ok {
		return *casted
	}
	return nil
}

func (m *_AdsAddDeviceNotificationRequest) GetTypeName() string {
	return "AdsAddDeviceNotificationRequest"
}

func (m *_AdsAddDeviceNotificationRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AmsPacketContract.(*_AmsPacket).getLengthInBits(ctx))

	// Simple field (indexGroup)
	lengthInBits += 32

	// Simple field (indexOffset)
	lengthInBits += 32

	// Simple field (length)
	lengthInBits += 32

	// Simple field (transmissionMode)
	lengthInBits += 32

	// Simple field (maxDelayInMs)
	lengthInBits += 32

	// Simple field (cycleTimeInMs)
	lengthInBits += 32

	// Reserved Field (reserved)
	lengthInBits += 64

	// Reserved Field (reserved)
	lengthInBits += 64

	return lengthInBits
}

func (m *_AdsAddDeviceNotificationRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AdsAddDeviceNotificationRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AmsPacket) (_adsAddDeviceNotificationRequest AdsAddDeviceNotificationRequest, err error) {
	m.AmsPacketContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsAddDeviceNotificationRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsAddDeviceNotificationRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	indexGroup, err := ReadSimpleField(ctx, "indexGroup", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'indexGroup' field"))
	}
	m.IndexGroup = indexGroup

	indexOffset, err := ReadSimpleField(ctx, "indexOffset", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'indexOffset' field"))
	}
	m.IndexOffset = indexOffset

	length, err := ReadSimpleField(ctx, "length", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'length' field"))
	}
	m.Length = length

	transmissionMode, err := ReadEnumField[AdsTransMode](ctx, "transmissionMode", "AdsTransMode", ReadEnum(AdsTransModeByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transmissionMode' field"))
	}
	m.TransmissionMode = transmissionMode

	maxDelayInMs, err := ReadSimpleField(ctx, "maxDelayInMs", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxDelayInMs' field"))
	}
	m.MaxDelayInMs = maxDelayInMs

	cycleTimeInMs, err := ReadSimpleField(ctx, "cycleTimeInMs", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'cycleTimeInMs' field"))
	}
	m.CycleTimeInMs = cycleTimeInMs

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedLong(readBuffer, uint8(64)), uint64(0x0000))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	reservedField1, err := ReadReservedField(ctx, "reserved", ReadUnsignedLong(readBuffer, uint8(64)), uint64(0x0000))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField1 = reservedField1

	if closeErr := readBuffer.CloseContext("AdsAddDeviceNotificationRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsAddDeviceNotificationRequest")
	}

	return m, nil
}

func (m *_AdsAddDeviceNotificationRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsAddDeviceNotificationRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsAddDeviceNotificationRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsAddDeviceNotificationRequest")
		}

		if err := WriteSimpleField[uint32](ctx, "indexGroup", m.GetIndexGroup(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'indexGroup' field")
		}

		if err := WriteSimpleField[uint32](ctx, "indexOffset", m.GetIndexOffset(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'indexOffset' field")
		}

		if err := WriteSimpleField[uint32](ctx, "length", m.GetLength(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'length' field")
		}

		if err := WriteSimpleEnumField[AdsTransMode](ctx, "transmissionMode", "AdsTransMode", m.GetTransmissionMode(), WriteEnum[AdsTransMode, uint32](AdsTransMode.GetValue, AdsTransMode.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'transmissionMode' field")
		}

		if err := WriteSimpleField[uint32](ctx, "maxDelayInMs", m.GetMaxDelayInMs(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxDelayInMs' field")
		}

		if err := WriteSimpleField[uint32](ctx, "cycleTimeInMs", m.GetCycleTimeInMs(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'cycleTimeInMs' field")
		}

		if err := WriteReservedField[uint64](ctx, "reserved", uint64(0x0000), WriteUnsignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteReservedField[uint64](ctx, "reserved", uint64(0x0000), WriteUnsignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 2")
		}

		if popErr := writeBuffer.PopContext("AdsAddDeviceNotificationRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsAddDeviceNotificationRequest")
		}
		return nil
	}
	return m.AmsPacketContract.(*_AmsPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsAddDeviceNotificationRequest) isAdsAddDeviceNotificationRequest() bool {
	return true
}

func (m *_AdsAddDeviceNotificationRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
