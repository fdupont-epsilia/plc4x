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

// GetAttributeAllResponse is the corresponding interface of GetAttributeAllResponse
type GetAttributeAllResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CipService
	// GetStatus returns Status (property field)
	GetStatus() uint8
	// GetExtStatus returns ExtStatus (property field)
	GetExtStatus() uint8
	// GetAttributes returns Attributes (property field)
	GetAttributes() CIPAttributes
}

// GetAttributeAllResponseExactly can be used when we want exactly this type and not a type which fulfills GetAttributeAllResponse.
// This is useful for switch cases.
type GetAttributeAllResponseExactly interface {
	GetAttributeAllResponse
	isGetAttributeAllResponse() bool
}

// _GetAttributeAllResponse is the data-structure of this message
type _GetAttributeAllResponse struct {
	*_CipService
	Status     uint8
	ExtStatus  uint8
	Attributes CIPAttributes
	// Reserved Fields
	reservedField0 *uint8
}

var _ GetAttributeAllResponse = (*_GetAttributeAllResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_GetAttributeAllResponse) GetService() uint8 {
	return 0x01
}

func (m *_GetAttributeAllResponse) GetResponse() bool {
	return bool(true)
}

func (m *_GetAttributeAllResponse) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_GetAttributeAllResponse) InitializeParent(parent CipService) {}

func (m *_GetAttributeAllResponse) GetParent() CipService {
	return m._CipService
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_GetAttributeAllResponse) GetStatus() uint8 {
	return m.Status
}

func (m *_GetAttributeAllResponse) GetExtStatus() uint8 {
	return m.ExtStatus
}

func (m *_GetAttributeAllResponse) GetAttributes() CIPAttributes {
	return m.Attributes
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewGetAttributeAllResponse factory function for _GetAttributeAllResponse
func NewGetAttributeAllResponse(status uint8, extStatus uint8, attributes CIPAttributes, serviceLen uint16) *_GetAttributeAllResponse {
	_result := &_GetAttributeAllResponse{
		Status:      status,
		ExtStatus:   extStatus,
		Attributes:  attributes,
		_CipService: NewCipService(serviceLen),
	}
	_result._CipService._CipServiceChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastGetAttributeAllResponse(structType any) GetAttributeAllResponse {
	if casted, ok := structType.(GetAttributeAllResponse); ok {
		return casted
	}
	if casted, ok := structType.(*GetAttributeAllResponse); ok {
		return *casted
	}
	return nil
}

func (m *_GetAttributeAllResponse) GetTypeName() string {
	return "GetAttributeAllResponse"
}

func (m *_GetAttributeAllResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	// Simple field (extStatus)
	lengthInBits += 8

	// Optional Field (attributes)
	if m.Attributes != nil {
		lengthInBits += m.Attributes.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_GetAttributeAllResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func GetAttributeAllResponseParse(ctx context.Context, theBytes []byte, connected bool, serviceLen uint16) (GetAttributeAllResponse, error) {
	return GetAttributeAllResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), connected, serviceLen)
}

func GetAttributeAllResponseParseWithBufferProducer(connected bool, serviceLen uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (GetAttributeAllResponse, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (GetAttributeAllResponse, error) {
		return GetAttributeAllResponseParseWithBuffer(ctx, readBuffer, connected, serviceLen)
	}
}

func GetAttributeAllResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, connected bool, serviceLen uint16) (GetAttributeAllResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("GetAttributeAllResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for GetAttributeAllResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00))
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

	_attributes, err := ReadOptionalField[CIPAttributes](ctx, "attributes", ReadComplex[CIPAttributes](CIPAttributesParseWithBufferProducer((uint16)(uint16(serviceLen)-uint16(uint16(4)))), readBuffer), bool(((serviceLen)-(4)) > (0)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'attributes' field"))
	}
	var attributes CIPAttributes
	if _attributes != nil {
		attributes = *_attributes
	}

	if closeErr := readBuffer.CloseContext("GetAttributeAllResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for GetAttributeAllResponse")
	}

	// Create a partially initialized instance
	_child := &_GetAttributeAllResponse{
		_CipService: &_CipService{
			ServiceLen: serviceLen,
		},
		Status:         status,
		ExtStatus:      extStatus,
		Attributes:     attributes,
		reservedField0: reservedField0,
	}
	_child._CipService._CipServiceChildRequirements = _child
	return _child, nil
}

func (m *_GetAttributeAllResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_GetAttributeAllResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("GetAttributeAllResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for GetAttributeAllResponse")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[uint8](ctx, "status", m.GetStatus(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'status' field")
		}

		if err := WriteSimpleField[uint8](ctx, "extStatus", m.GetExtStatus(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'extStatus' field")
		}

		if err := WriteOptionalField[CIPAttributes](ctx, "attributes", GetRef(m.GetAttributes()), WriteComplex[CIPAttributes](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'attributes' field")
		}

		if popErr := writeBuffer.PopContext("GetAttributeAllResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for GetAttributeAllResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_GetAttributeAllResponse) isGetAttributeAllResponse() bool {
	return true
}

func (m *_GetAttributeAllResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
