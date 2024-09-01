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

// ApduData is the corresponding interface of ApduData
type ApduData interface {
	ApduDataContract
	ApduDataRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// ApduDataContract provides a set of functions which can be overwritten by a sub struct
type ApduDataContract interface {
}

// ApduDataRequirements provides a set of functions which need to be implemented by a sub struct
type ApduDataRequirements interface {
	// GetApciType returns ApciType (discriminator field)
	GetApciType() uint8
}

// ApduDataExactly can be used when we want exactly this type and not a type which fulfills ApduData.
// This is useful for switch cases.
type ApduDataExactly interface {
	ApduData
	isApduData() bool
}

// _ApduData is the data-structure of this message
type _ApduData struct {
	_ApduDataChildRequirements

	// Arguments.
	DataLength uint8
}

var _ ApduDataContract = (*_ApduData)(nil)

type _ApduDataChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetApciType() uint8
}

type ApduDataParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ApduData, serializeChildFunction func() error) error
	GetTypeName() string
}

type ApduDataChild interface {
	utils.Serializable
	InitializeParent(parent ApduData)
	GetParent() *ApduData

	GetTypeName() string
	ApduData
}

// NewApduData factory function for _ApduData
func NewApduData(dataLength uint8) *_ApduData {
	return &_ApduData{DataLength: dataLength}
}

// Deprecated: use the interface for direct cast
func CastApduData(structType any) ApduData {
	if casted, ok := structType.(ApduData); ok {
		return casted
	}
	if casted, ok := structType.(*ApduData); ok {
		return *casted
	}
	return nil
}

func (m *_ApduData) GetTypeName() string {
	return "ApduData"
}

func (m *_ApduData) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (apciType)
	lengthInBits += 4

	return lengthInBits
}

func (m *_ApduData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ApduDataParse(ctx context.Context, theBytes []byte, dataLength uint8) (ApduData, error) {
	return ApduDataParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), dataLength)
}

func ApduDataParseWithBufferProducer[T ApduData](dataLength uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		buffer, err := ApduDataParseWithBuffer(ctx, readBuffer, dataLength)
		if err != nil {
			var zero T
			return zero, err
		}
		return buffer.(T), err
	}
}

func ApduDataParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, dataLength uint8) (ApduData, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	apciType, err := ReadDiscriminatorField[uint8](ctx, "apciType", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'apciType' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type ApduDataChildSerializeRequirement interface {
		ApduData
		InitializeParent(ApduData)
		GetParent() ApduData
	}
	var _childTemp any
	var _child ApduDataChildSerializeRequirement
	var typeSwitchError error
	switch {
	case apciType == 0x0: // ApduDataGroupValueRead
		_childTemp, typeSwitchError = ApduDataGroupValueReadParseWithBuffer(ctx, readBuffer, dataLength)
	case apciType == 0x1: // ApduDataGroupValueResponse
		_childTemp, typeSwitchError = ApduDataGroupValueResponseParseWithBuffer(ctx, readBuffer, dataLength)
	case apciType == 0x2: // ApduDataGroupValueWrite
		_childTemp, typeSwitchError = ApduDataGroupValueWriteParseWithBuffer(ctx, readBuffer, dataLength)
	case apciType == 0x3: // ApduDataIndividualAddressWrite
		_childTemp, typeSwitchError = ApduDataIndividualAddressWriteParseWithBuffer(ctx, readBuffer, dataLength)
	case apciType == 0x4: // ApduDataIndividualAddressRead
		_childTemp, typeSwitchError = ApduDataIndividualAddressReadParseWithBuffer(ctx, readBuffer, dataLength)
	case apciType == 0x5: // ApduDataIndividualAddressResponse
		_childTemp, typeSwitchError = ApduDataIndividualAddressResponseParseWithBuffer(ctx, readBuffer, dataLength)
	case apciType == 0x6: // ApduDataAdcRead
		_childTemp, typeSwitchError = ApduDataAdcReadParseWithBuffer(ctx, readBuffer, dataLength)
	case apciType == 0x7: // ApduDataAdcResponse
		_childTemp, typeSwitchError = ApduDataAdcResponseParseWithBuffer(ctx, readBuffer, dataLength)
	case apciType == 0x8: // ApduDataMemoryRead
		_childTemp, typeSwitchError = ApduDataMemoryReadParseWithBuffer(ctx, readBuffer, dataLength)
	case apciType == 0x9: // ApduDataMemoryResponse
		_childTemp, typeSwitchError = ApduDataMemoryResponseParseWithBuffer(ctx, readBuffer, dataLength)
	case apciType == 0xA: // ApduDataMemoryWrite
		_childTemp, typeSwitchError = ApduDataMemoryWriteParseWithBuffer(ctx, readBuffer, dataLength)
	case apciType == 0xB: // ApduDataUserMessage
		_childTemp, typeSwitchError = ApduDataUserMessageParseWithBuffer(ctx, readBuffer, dataLength)
	case apciType == 0xC: // ApduDataDeviceDescriptorRead
		_childTemp, typeSwitchError = ApduDataDeviceDescriptorReadParseWithBuffer(ctx, readBuffer, dataLength)
	case apciType == 0xD: // ApduDataDeviceDescriptorResponse
		_childTemp, typeSwitchError = ApduDataDeviceDescriptorResponseParseWithBuffer(ctx, readBuffer, dataLength)
	case apciType == 0xE: // ApduDataRestart
		_childTemp, typeSwitchError = ApduDataRestartParseWithBuffer(ctx, readBuffer, dataLength)
	case apciType == 0xF: // ApduDataOther
		_childTemp, typeSwitchError = ApduDataOtherParseWithBuffer(ctx, readBuffer, dataLength)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [apciType=%v]", apciType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of ApduData")
	}
	_child = _childTemp.(ApduDataChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("ApduData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduData")
	}

	// Finish initializing
	_child.InitializeParent(_child)
	return _child, nil
}

func (pm *_ApduData) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ApduData, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ApduData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ApduData")
	}

	if err := WriteDiscriminatorField(ctx, "apciType", m.GetApciType(), WriteUnsignedByte(writeBuffer, 4)); err != nil {
		return errors.Wrap(err, "Error serializing 'apciType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ApduData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ApduData")
	}
	return nil
}

////
// Arguments Getter

func (m *_ApduData) GetDataLength() uint8 {
	return m.DataLength
}

//
////

func (m *_ApduData) isApduData() bool {
	return true
}

func (m *_ApduData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
