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

// AdsDeleteDeviceNotificationResponse is the corresponding interface of AdsDeleteDeviceNotificationResponse
type AdsDeleteDeviceNotificationResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AmsPacket
	// GetResult returns Result (property field)
	GetResult() ReturnCode
}

// AdsDeleteDeviceNotificationResponseExactly can be used when we want exactly this type and not a type which fulfills AdsDeleteDeviceNotificationResponse.
// This is useful for switch cases.
type AdsDeleteDeviceNotificationResponseExactly interface {
	AdsDeleteDeviceNotificationResponse
	isAdsDeleteDeviceNotificationResponse() bool
}

// _AdsDeleteDeviceNotificationResponse is the data-structure of this message
type _AdsDeleteDeviceNotificationResponse struct {
	AmsPacketContract
	Result ReturnCode
}

var _ AdsDeleteDeviceNotificationResponse = (*_AdsDeleteDeviceNotificationResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsDeleteDeviceNotificationResponse) GetCommandId() CommandId {
	return CommandId_ADS_DELETE_DEVICE_NOTIFICATION
}

func (m *_AdsDeleteDeviceNotificationResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsDeleteDeviceNotificationResponse) GetParent() AmsPacketContract {
	return m.AmsPacketContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsDeleteDeviceNotificationResponse) GetResult() ReturnCode {
	return m.Result
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsDeleteDeviceNotificationResponse factory function for _AdsDeleteDeviceNotificationResponse
func NewAdsDeleteDeviceNotificationResponse(result ReturnCode, targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32) *_AdsDeleteDeviceNotificationResponse {
	_result := &_AdsDeleteDeviceNotificationResponse{
		AmsPacketContract: NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId),
		Result:            result,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsDeleteDeviceNotificationResponse(structType any) AdsDeleteDeviceNotificationResponse {
	if casted, ok := structType.(AdsDeleteDeviceNotificationResponse); ok {
		return casted
	}
	if casted, ok := structType.(*AdsDeleteDeviceNotificationResponse); ok {
		return *casted
	}
	return nil
}

func (m *_AdsDeleteDeviceNotificationResponse) GetTypeName() string {
	return "AdsDeleteDeviceNotificationResponse"
}

func (m *_AdsDeleteDeviceNotificationResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AmsPacketContract.(*_AmsPacket).getLengthInBits(ctx))

	// Simple field (result)
	lengthInBits += 32

	return lengthInBits
}

func (m *_AdsDeleteDeviceNotificationResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AdsDeleteDeviceNotificationResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AmsPacket) (_adsDeleteDeviceNotificationResponse AdsDeleteDeviceNotificationResponse, err error) {
	m.AmsPacketContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsDeleteDeviceNotificationResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsDeleteDeviceNotificationResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	result, err := ReadEnumField[ReturnCode](ctx, "result", "ReturnCode", ReadEnum(ReturnCodeByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'result' field"))
	}
	m.Result = result

	if closeErr := readBuffer.CloseContext("AdsDeleteDeviceNotificationResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsDeleteDeviceNotificationResponse")
	}

	return m, nil
}

func (m *_AdsDeleteDeviceNotificationResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsDeleteDeviceNotificationResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsDeleteDeviceNotificationResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsDeleteDeviceNotificationResponse")
		}

		if err := WriteSimpleEnumField[ReturnCode](ctx, "result", "ReturnCode", m.GetResult(), WriteEnum[ReturnCode, uint32](ReturnCode.GetValue, ReturnCode.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'result' field")
		}

		if popErr := writeBuffer.PopContext("AdsDeleteDeviceNotificationResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsDeleteDeviceNotificationResponse")
		}
		return nil
	}
	return m.AmsPacketContract.(*_AmsPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsDeleteDeviceNotificationResponse) isAdsDeleteDeviceNotificationResponse() bool {
	return true
}

func (m *_AdsDeleteDeviceNotificationResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
