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

// AdsNotificationSample is the corresponding interface of AdsNotificationSample
type AdsNotificationSample interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetNotificationHandle returns NotificationHandle (property field)
	GetNotificationHandle() uint32
	// GetSampleSize returns SampleSize (property field)
	GetSampleSize() uint32
	// GetData returns Data (property field)
	GetData() []byte
}

// AdsNotificationSampleExactly can be used when we want exactly this type and not a type which fulfills AdsNotificationSample.
// This is useful for switch cases.
type AdsNotificationSampleExactly interface {
	AdsNotificationSample
	isAdsNotificationSample() bool
}

// _AdsNotificationSample is the data-structure of this message
type _AdsNotificationSample struct {
	NotificationHandle uint32
	SampleSize         uint32
	Data               []byte
}

var _ AdsNotificationSample = (*_AdsNotificationSample)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsNotificationSample) GetNotificationHandle() uint32 {
	return m.NotificationHandle
}

func (m *_AdsNotificationSample) GetSampleSize() uint32 {
	return m.SampleSize
}

func (m *_AdsNotificationSample) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsNotificationSample factory function for _AdsNotificationSample
func NewAdsNotificationSample(notificationHandle uint32, sampleSize uint32, data []byte) *_AdsNotificationSample {
	return &_AdsNotificationSample{NotificationHandle: notificationHandle, SampleSize: sampleSize, Data: data}
}

// Deprecated: use the interface for direct cast
func CastAdsNotificationSample(structType any) AdsNotificationSample {
	if casted, ok := structType.(AdsNotificationSample); ok {
		return casted
	}
	if casted, ok := structType.(*AdsNotificationSample); ok {
		return *casted
	}
	return nil
}

func (m *_AdsNotificationSample) GetTypeName() string {
	return "AdsNotificationSample"
}

func (m *_AdsNotificationSample) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (notificationHandle)
	lengthInBits += 32

	// Simple field (sampleSize)
	lengthInBits += 32

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_AdsNotificationSample) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsNotificationSampleParse(ctx context.Context, theBytes []byte) (AdsNotificationSample, error) {
	return AdsNotificationSampleParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AdsNotificationSampleParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (AdsNotificationSample, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (AdsNotificationSample, error) {
		return AdsNotificationSampleParseWithBuffer(ctx, readBuffer)
	}
}

func AdsNotificationSampleParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AdsNotificationSample, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsNotificationSample"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsNotificationSample")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	notificationHandle, err := ReadSimpleField(ctx, "notificationHandle", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'notificationHandle' field"))
	}

	sampleSize, err := ReadSimpleField(ctx, "sampleSize", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sampleSize' field"))
	}

	data, err := readBuffer.ReadByteArray("data", int(sampleSize))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}

	if closeErr := readBuffer.CloseContext("AdsNotificationSample"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsNotificationSample")
	}

	// Create the instance
	return &_AdsNotificationSample{
		NotificationHandle: notificationHandle,
		SampleSize:         sampleSize,
		Data:               data,
	}, nil
}

func (m *_AdsNotificationSample) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsNotificationSample) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AdsNotificationSample"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AdsNotificationSample")
	}

	if err := WriteSimpleField[uint32](ctx, "notificationHandle", m.GetNotificationHandle(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
		return errors.Wrap(err, "Error serializing 'notificationHandle' field")
	}

	if err := WriteSimpleField[uint32](ctx, "sampleSize", m.GetSampleSize(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
		return errors.Wrap(err, "Error serializing 'sampleSize' field")
	}

	if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'data' field")
	}

	if popErr := writeBuffer.PopContext("AdsNotificationSample"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AdsNotificationSample")
	}
	return nil
}

func (m *_AdsNotificationSample) isAdsNotificationSample() bool {
	return true
}

func (m *_AdsNotificationSample) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
