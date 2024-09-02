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

// Alarm8MessageQueryType is the corresponding interface of Alarm8MessageQueryType
type Alarm8MessageQueryType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetFunctionId returns FunctionId (property field)
	GetFunctionId() uint8
	// GetNumberOfObjects returns NumberOfObjects (property field)
	GetNumberOfObjects() uint8
	// GetReturnCode returns ReturnCode (property field)
	GetReturnCode() DataTransportErrorCode
	// GetTransportSize returns TransportSize (property field)
	GetTransportSize() DataTransportSize
	// GetByteCount returns ByteCount (property field)
	GetByteCount() uint16
	// GetMessageObjects returns MessageObjects (property field)
	GetMessageObjects() []AlarmMessageObjectQueryType
	// IsAlarm8MessageQueryType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAlarm8MessageQueryType()
}

// _Alarm8MessageQueryType is the data-structure of this message
type _Alarm8MessageQueryType struct {
	FunctionId      uint8
	NumberOfObjects uint8
	ReturnCode      DataTransportErrorCode
	TransportSize   DataTransportSize
	ByteCount       uint16
	MessageObjects  []AlarmMessageObjectQueryType
}

var _ Alarm8MessageQueryType = (*_Alarm8MessageQueryType)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_Alarm8MessageQueryType) GetFunctionId() uint8 {
	return m.FunctionId
}

func (m *_Alarm8MessageQueryType) GetNumberOfObjects() uint8 {
	return m.NumberOfObjects
}

func (m *_Alarm8MessageQueryType) GetReturnCode() DataTransportErrorCode {
	return m.ReturnCode
}

func (m *_Alarm8MessageQueryType) GetTransportSize() DataTransportSize {
	return m.TransportSize
}

func (m *_Alarm8MessageQueryType) GetByteCount() uint16 {
	return m.ByteCount
}

func (m *_Alarm8MessageQueryType) GetMessageObjects() []AlarmMessageObjectQueryType {
	return m.MessageObjects
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAlarm8MessageQueryType factory function for _Alarm8MessageQueryType
func NewAlarm8MessageQueryType(functionId uint8, numberOfObjects uint8, returnCode DataTransportErrorCode, transportSize DataTransportSize, byteCount uint16, messageObjects []AlarmMessageObjectQueryType) *_Alarm8MessageQueryType {
	return &_Alarm8MessageQueryType{FunctionId: functionId, NumberOfObjects: numberOfObjects, ReturnCode: returnCode, TransportSize: transportSize, ByteCount: byteCount, MessageObjects: messageObjects}
}

// Deprecated: use the interface for direct cast
func CastAlarm8MessageQueryType(structType any) Alarm8MessageQueryType {
	if casted, ok := structType.(Alarm8MessageQueryType); ok {
		return casted
	}
	if casted, ok := structType.(*Alarm8MessageQueryType); ok {
		return *casted
	}
	return nil
}

func (m *_Alarm8MessageQueryType) GetTypeName() string {
	return "Alarm8MessageQueryType"
}

func (m *_Alarm8MessageQueryType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (functionId)
	lengthInBits += 8

	// Simple field (numberOfObjects)
	lengthInBits += 8

	// Simple field (returnCode)
	lengthInBits += 8

	// Simple field (transportSize)
	lengthInBits += 8

	// Simple field (byteCount)
	lengthInBits += 16

	// Array field
	if len(m.MessageObjects) > 0 {
		for _curItem, element := range m.MessageObjects {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.MessageObjects), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_Alarm8MessageQueryType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func Alarm8MessageQueryTypeParse(ctx context.Context, theBytes []byte) (Alarm8MessageQueryType, error) {
	return Alarm8MessageQueryTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func Alarm8MessageQueryTypeParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (Alarm8MessageQueryType, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (Alarm8MessageQueryType, error) {
		return Alarm8MessageQueryTypeParseWithBuffer(ctx, readBuffer)
	}
}

func Alarm8MessageQueryTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (Alarm8MessageQueryType, error) {
	v, err := (&_Alarm8MessageQueryType{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_Alarm8MessageQueryType) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__alarm8MessageQueryType Alarm8MessageQueryType, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("Alarm8MessageQueryType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Alarm8MessageQueryType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	functionId, err := ReadSimpleField(ctx, "functionId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'functionId' field"))
	}
	m.FunctionId = functionId

	numberOfObjects, err := ReadSimpleField(ctx, "numberOfObjects", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfObjects' field"))
	}
	m.NumberOfObjects = numberOfObjects

	returnCode, err := ReadEnumField[DataTransportErrorCode](ctx, "returnCode", "DataTransportErrorCode", ReadEnum(DataTransportErrorCodeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'returnCode' field"))
	}
	m.ReturnCode = returnCode

	transportSize, err := ReadEnumField[DataTransportSize](ctx, "transportSize", "DataTransportSize", ReadEnum(DataTransportSizeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transportSize' field"))
	}
	m.TransportSize = transportSize

	byteCount, err := ReadSimpleField(ctx, "byteCount", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'byteCount' field"))
	}
	m.ByteCount = byteCount

	messageObjects, err := ReadCountArrayField[AlarmMessageObjectQueryType](ctx, "messageObjects", ReadComplex[AlarmMessageObjectQueryType](AlarmMessageObjectQueryTypeParseWithBuffer, readBuffer), uint64(int32(byteCount)/int32(int32(12))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'messageObjects' field"))
	}
	m.MessageObjects = messageObjects

	if closeErr := readBuffer.CloseContext("Alarm8MessageQueryType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Alarm8MessageQueryType")
	}

	return m, nil
}

func (m *_Alarm8MessageQueryType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_Alarm8MessageQueryType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("Alarm8MessageQueryType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for Alarm8MessageQueryType")
	}

	if err := WriteSimpleField[uint8](ctx, "functionId", m.GetFunctionId(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'functionId' field")
	}

	if err := WriteSimpleField[uint8](ctx, "numberOfObjects", m.GetNumberOfObjects(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'numberOfObjects' field")
	}

	if err := WriteSimpleEnumField[DataTransportErrorCode](ctx, "returnCode", "DataTransportErrorCode", m.GetReturnCode(), WriteEnum[DataTransportErrorCode, uint8](DataTransportErrorCode.GetValue, DataTransportErrorCode.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'returnCode' field")
	}

	if err := WriteSimpleEnumField[DataTransportSize](ctx, "transportSize", "DataTransportSize", m.GetTransportSize(), WriteEnum[DataTransportSize, uint8](DataTransportSize.GetValue, DataTransportSize.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'transportSize' field")
	}

	if err := WriteSimpleField[uint16](ctx, "byteCount", m.GetByteCount(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'byteCount' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "messageObjects", m.GetMessageObjects(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'messageObjects' field")
	}

	if popErr := writeBuffer.PopContext("Alarm8MessageQueryType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for Alarm8MessageQueryType")
	}
	return nil
}

func (m *_Alarm8MessageQueryType) IsAlarm8MessageQueryType() {}

func (m *_Alarm8MessageQueryType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
