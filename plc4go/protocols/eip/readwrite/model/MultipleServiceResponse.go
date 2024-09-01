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

// MultipleServiceResponse is the corresponding interface of MultipleServiceResponse
type MultipleServiceResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CipService
	// GetStatus returns Status (property field)
	GetStatus() uint8
	// GetExtStatus returns ExtStatus (property field)
	GetExtStatus() uint8
	// GetServiceNb returns ServiceNb (property field)
	GetServiceNb() uint16
	// GetOffsets returns Offsets (property field)
	GetOffsets() []uint16
	// GetServicesData returns ServicesData (property field)
	GetServicesData() []byte
}

// MultipleServiceResponseExactly can be used when we want exactly this type and not a type which fulfills MultipleServiceResponse.
// This is useful for switch cases.
type MultipleServiceResponseExactly interface {
	MultipleServiceResponse
	isMultipleServiceResponse() bool
}

// _MultipleServiceResponse is the data-structure of this message
type _MultipleServiceResponse struct {
	*_CipService
	Status       uint8
	ExtStatus    uint8
	ServiceNb    uint16
	Offsets      []uint16
	ServicesData []byte
	// Reserved Fields
	reservedField0 *uint8
}

var _ MultipleServiceResponse = (*_MultipleServiceResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MultipleServiceResponse) GetService() uint8 {
	return 0x0A
}

func (m *_MultipleServiceResponse) GetResponse() bool {
	return bool(true)
}

func (m *_MultipleServiceResponse) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MultipleServiceResponse) InitializeParent(parent CipService) {}

func (m *_MultipleServiceResponse) GetParent() CipService {
	return m._CipService
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MultipleServiceResponse) GetStatus() uint8 {
	return m.Status
}

func (m *_MultipleServiceResponse) GetExtStatus() uint8 {
	return m.ExtStatus
}

func (m *_MultipleServiceResponse) GetServiceNb() uint16 {
	return m.ServiceNb
}

func (m *_MultipleServiceResponse) GetOffsets() []uint16 {
	return m.Offsets
}

func (m *_MultipleServiceResponse) GetServicesData() []byte {
	return m.ServicesData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMultipleServiceResponse factory function for _MultipleServiceResponse
func NewMultipleServiceResponse(status uint8, extStatus uint8, serviceNb uint16, offsets []uint16, servicesData []byte, serviceLen uint16) *_MultipleServiceResponse {
	_result := &_MultipleServiceResponse{
		Status:       status,
		ExtStatus:    extStatus,
		ServiceNb:    serviceNb,
		Offsets:      offsets,
		ServicesData: servicesData,
		_CipService:  NewCipService(serviceLen),
	}
	_result._CipService._CipServiceChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMultipleServiceResponse(structType any) MultipleServiceResponse {
	if casted, ok := structType.(MultipleServiceResponse); ok {
		return casted
	}
	if casted, ok := structType.(*MultipleServiceResponse); ok {
		return *casted
	}
	return nil
}

func (m *_MultipleServiceResponse) GetTypeName() string {
	return "MultipleServiceResponse"
}

func (m *_MultipleServiceResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	// Simple field (extStatus)
	lengthInBits += 8

	// Simple field (serviceNb)
	lengthInBits += 16

	// Array field
	if len(m.Offsets) > 0 {
		lengthInBits += 16 * uint16(len(m.Offsets))
	}

	// Array field
	if len(m.ServicesData) > 0 {
		lengthInBits += 8 * uint16(len(m.ServicesData))
	}

	return lengthInBits
}

func (m *_MultipleServiceResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MultipleServiceResponseParse(ctx context.Context, theBytes []byte, connected bool, serviceLen uint16) (MultipleServiceResponse, error) {
	return MultipleServiceResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), connected, serviceLen)
}

func MultipleServiceResponseParseWithBufferProducer(connected bool, serviceLen uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (MultipleServiceResponse, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (MultipleServiceResponse, error) {
		return MultipleServiceResponseParseWithBuffer(ctx, readBuffer, connected, serviceLen)
	}
}

func MultipleServiceResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, connected bool, serviceLen uint16) (MultipleServiceResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MultipleServiceResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MultipleServiceResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	status, err := ReadSimpleField(ctx, "status", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'status' field"))
	}

	extStatus, err := ReadSimpleField(ctx, "extStatus", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'extStatus' field"))
	}

	serviceNb, err := ReadSimpleField(ctx, "serviceNb", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serviceNb' field"))
	}

	offsets, err := ReadCountArrayField[uint16](ctx, "offsets", ReadUnsignedShort(readBuffer, uint8(16)), uint64(serviceNb))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'offsets' field"))
	}

	servicesData, err := readBuffer.ReadByteArray("servicesData", int(int32(int32(serviceLen)-int32(int32(6)))-int32((int32(int32(2))*int32(serviceNb)))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'servicesData' field"))
	}

	if closeErr := readBuffer.CloseContext("MultipleServiceResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MultipleServiceResponse")
	}

	// Create a partially initialized instance
	_child := &_MultipleServiceResponse{
		_CipService: &_CipService{
			ServiceLen: serviceLen,
		},
		Status:         status,
		ExtStatus:      extStatus,
		ServiceNb:      serviceNb,
		Offsets:        offsets,
		ServicesData:   servicesData,
		reservedField0: reservedField0,
	}
	_child._CipService._CipServiceChildRequirements = _child
	return _child, nil
}

func (m *_MultipleServiceResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MultipleServiceResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MultipleServiceResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MultipleServiceResponse")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x0), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[uint8](ctx, "status", m.GetStatus(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'status' field")
		}

		if err := WriteSimpleField[uint8](ctx, "extStatus", m.GetExtStatus(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'extStatus' field")
		}

		if err := WriteSimpleField[uint16](ctx, "serviceNb", m.GetServiceNb(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'serviceNb' field")
		}

		if err := WriteSimpleTypeArrayField(ctx, "offsets", m.GetOffsets(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'offsets' field")
		}

		if err := WriteByteArrayField(ctx, "servicesData", m.GetServicesData(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'servicesData' field")
		}

		if popErr := writeBuffer.PopContext("MultipleServiceResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MultipleServiceResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MultipleServiceResponse) isMultipleServiceResponse() bool {
	return true
}

func (m *_MultipleServiceResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
