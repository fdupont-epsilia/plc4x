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

// AccessControlDataValidAccessRequest is the corresponding interface of AccessControlDataValidAccessRequest
type AccessControlDataValidAccessRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AccessControlData
	// GetAccessControlDirection returns AccessControlDirection (property field)
	GetAccessControlDirection() AccessControlDirection
	// GetData returns Data (property field)
	GetData() []byte
	// IsAccessControlDataValidAccessRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAccessControlDataValidAccessRequest()
}

// _AccessControlDataValidAccessRequest is the data-structure of this message
type _AccessControlDataValidAccessRequest struct {
	AccessControlDataContract
	AccessControlDirection AccessControlDirection
	Data                   []byte
}

var _ AccessControlDataValidAccessRequest = (*_AccessControlDataValidAccessRequest)(nil)
var _ AccessControlDataRequirements = (*_AccessControlDataValidAccessRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AccessControlDataValidAccessRequest) GetParent() AccessControlDataContract {
	return m.AccessControlDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AccessControlDataValidAccessRequest) GetAccessControlDirection() AccessControlDirection {
	return m.AccessControlDirection
}

func (m *_AccessControlDataValidAccessRequest) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAccessControlDataValidAccessRequest factory function for _AccessControlDataValidAccessRequest
func NewAccessControlDataValidAccessRequest(accessControlDirection AccessControlDirection, data []byte, commandTypeContainer AccessControlCommandTypeContainer, networkId byte, accessPointId byte) *_AccessControlDataValidAccessRequest {
	_result := &_AccessControlDataValidAccessRequest{
		AccessControlDataContract: NewAccessControlData(commandTypeContainer, networkId, accessPointId),
		AccessControlDirection:    accessControlDirection,
		Data:                      data,
	}
	_result.AccessControlDataContract.(*_AccessControlData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAccessControlDataValidAccessRequest(structType any) AccessControlDataValidAccessRequest {
	if casted, ok := structType.(AccessControlDataValidAccessRequest); ok {
		return casted
	}
	if casted, ok := structType.(*AccessControlDataValidAccessRequest); ok {
		return *casted
	}
	return nil
}

func (m *_AccessControlDataValidAccessRequest) GetTypeName() string {
	return "AccessControlDataValidAccessRequest"
}

func (m *_AccessControlDataValidAccessRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AccessControlDataContract.(*_AccessControlData).getLengthInBits(ctx))

	// Simple field (accessControlDirection)
	lengthInBits += 8

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_AccessControlDataValidAccessRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AccessControlDataValidAccessRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AccessControlData, commandTypeContainer AccessControlCommandTypeContainer) (__accessControlDataValidAccessRequest AccessControlDataValidAccessRequest, err error) {
	m.AccessControlDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AccessControlDataValidAccessRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AccessControlDataValidAccessRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	accessControlDirection, err := ReadEnumField[AccessControlDirection](ctx, "accessControlDirection", "AccessControlDirection", ReadEnum(AccessControlDirectionByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'accessControlDirection' field"))
	}
	m.AccessControlDirection = accessControlDirection

	data, err := readBuffer.ReadByteArray("data", int(int32(commandTypeContainer.NumBytes())-int32(int32(3))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("AccessControlDataValidAccessRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AccessControlDataValidAccessRequest")
	}

	return m, nil
}

func (m *_AccessControlDataValidAccessRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AccessControlDataValidAccessRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AccessControlDataValidAccessRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AccessControlDataValidAccessRequest")
		}

		if err := WriteSimpleEnumField[AccessControlDirection](ctx, "accessControlDirection", "AccessControlDirection", m.GetAccessControlDirection(), WriteEnum[AccessControlDirection, uint8](AccessControlDirection.GetValue, AccessControlDirection.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'accessControlDirection' field")
		}

		if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("AccessControlDataValidAccessRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AccessControlDataValidAccessRequest")
		}
		return nil
	}
	return m.AccessControlDataContract.(*_AccessControlData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AccessControlDataValidAccessRequest) IsAccessControlDataValidAccessRequest() {}

func (m *_AccessControlDataValidAccessRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
