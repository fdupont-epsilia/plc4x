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

// CipConnectedResponse is the corresponding interface of CipConnectedResponse
type CipConnectedResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CipService
	// GetStatus returns Status (property field)
	GetStatus() uint8
	// GetAdditionalStatusWords returns AdditionalStatusWords (property field)
	GetAdditionalStatusWords() uint8
	// GetData returns Data (property field)
	GetData() CIPDataConnected
	// IsCipConnectedResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCipConnectedResponse()
}

// _CipConnectedResponse is the data-structure of this message
type _CipConnectedResponse struct {
	CipServiceContract
	Status                uint8
	AdditionalStatusWords uint8
	Data                  CIPDataConnected
	// Reserved Fields
	reservedField0 *uint8
}

var _ CipConnectedResponse = (*_CipConnectedResponse)(nil)
var _ CipServiceRequirements = (*_CipConnectedResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CipConnectedResponse) GetService() uint8 {
	return 0x52
}

func (m *_CipConnectedResponse) GetResponse() bool {
	return bool(true)
}

func (m *_CipConnectedResponse) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CipConnectedResponse) GetParent() CipServiceContract {
	return m.CipServiceContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CipConnectedResponse) GetStatus() uint8 {
	return m.Status
}

func (m *_CipConnectedResponse) GetAdditionalStatusWords() uint8 {
	return m.AdditionalStatusWords
}

func (m *_CipConnectedResponse) GetData() CIPDataConnected {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCipConnectedResponse factory function for _CipConnectedResponse
func NewCipConnectedResponse(status uint8, additionalStatusWords uint8, data CIPDataConnected, serviceLen uint16) *_CipConnectedResponse {
	_result := &_CipConnectedResponse{
		CipServiceContract:    NewCipService(serviceLen),
		Status:                status,
		AdditionalStatusWords: additionalStatusWords,
		Data:                  data,
	}
	_result.CipServiceContract.(*_CipService)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCipConnectedResponse(structType any) CipConnectedResponse {
	if casted, ok := structType.(CipConnectedResponse); ok {
		return casted
	}
	if casted, ok := structType.(*CipConnectedResponse); ok {
		return *casted
	}
	return nil
}

func (m *_CipConnectedResponse) GetTypeName() string {
	return "CipConnectedResponse"
}

func (m *_CipConnectedResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CipServiceContract.(*_CipService).getLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	// Simple field (additionalStatusWords)
	lengthInBits += 8

	// Optional Field (data)
	if m.Data != nil {
		lengthInBits += m.Data.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_CipConnectedResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CipConnectedResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CipService, connected bool, serviceLen uint16) (__cipConnectedResponse CipConnectedResponse, err error) {
	m.CipServiceContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CipConnectedResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CipConnectedResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	status, err := ReadSimpleField(ctx, "status", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'status' field"))
	}
	m.Status = status

	additionalStatusWords, err := ReadSimpleField(ctx, "additionalStatusWords", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'additionalStatusWords' field"))
	}
	m.AdditionalStatusWords = additionalStatusWords

	var data CIPDataConnected
	_data, err := ReadOptionalField[CIPDataConnected](ctx, "data", ReadComplex[CIPDataConnected](CIPDataConnectedParseWithBuffer, readBuffer), bool(((serviceLen)-(4)) > (0)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	if _data != nil {
		data = *_data
		m.Data = data
	}

	if closeErr := readBuffer.CloseContext("CipConnectedResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CipConnectedResponse")
	}

	return m, nil
}

func (m *_CipConnectedResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CipConnectedResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipConnectedResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CipConnectedResponse")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[uint8](ctx, "status", m.GetStatus(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'status' field")
		}

		if err := WriteSimpleField[uint8](ctx, "additionalStatusWords", m.GetAdditionalStatusWords(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'additionalStatusWords' field")
		}

		if err := WriteOptionalField[CIPDataConnected](ctx, "data", GetRef(m.GetData()), WriteComplex[CIPDataConnected](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("CipConnectedResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CipConnectedResponse")
		}
		return nil
	}
	return m.CipServiceContract.(*_CipService).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CipConnectedResponse) IsCipConnectedResponse() {}

func (m *_CipConnectedResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
