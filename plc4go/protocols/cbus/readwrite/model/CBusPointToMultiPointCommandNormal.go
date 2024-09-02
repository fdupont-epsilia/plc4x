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

// CBusPointToMultiPointCommandNormal is the corresponding interface of CBusPointToMultiPointCommandNormal
type CBusPointToMultiPointCommandNormal interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CBusPointToMultiPointCommand
	// GetApplication returns Application (property field)
	GetApplication() ApplicationIdContainer
	// GetSalData returns SalData (property field)
	GetSalData() SALData
	// IsCBusPointToMultiPointCommandNormal is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCBusPointToMultiPointCommandNormal()
}

// _CBusPointToMultiPointCommandNormal is the data-structure of this message
type _CBusPointToMultiPointCommandNormal struct {
	CBusPointToMultiPointCommandContract
	Application ApplicationIdContainer
	SalData     SALData
	// Reserved Fields
	reservedField0 *byte
}

var _ CBusPointToMultiPointCommandNormal = (*_CBusPointToMultiPointCommandNormal)(nil)
var _ CBusPointToMultiPointCommandRequirements = (*_CBusPointToMultiPointCommandNormal)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CBusPointToMultiPointCommandNormal) GetParent() CBusPointToMultiPointCommandContract {
	return m.CBusPointToMultiPointCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusPointToMultiPointCommandNormal) GetApplication() ApplicationIdContainer {
	return m.Application
}

func (m *_CBusPointToMultiPointCommandNormal) GetSalData() SALData {
	return m.SalData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCBusPointToMultiPointCommandNormal factory function for _CBusPointToMultiPointCommandNormal
func NewCBusPointToMultiPointCommandNormal(application ApplicationIdContainer, salData SALData, peekedApplication byte, cBusOptions CBusOptions) *_CBusPointToMultiPointCommandNormal {
	_result := &_CBusPointToMultiPointCommandNormal{
		CBusPointToMultiPointCommandContract: NewCBusPointToMultiPointCommand(peekedApplication, cBusOptions),
		Application:                          application,
		SalData:                              salData,
	}
	_result.CBusPointToMultiPointCommandContract.(*_CBusPointToMultiPointCommand)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCBusPointToMultiPointCommandNormal(structType any) CBusPointToMultiPointCommandNormal {
	if casted, ok := structType.(CBusPointToMultiPointCommandNormal); ok {
		return casted
	}
	if casted, ok := structType.(*CBusPointToMultiPointCommandNormal); ok {
		return *casted
	}
	return nil
}

func (m *_CBusPointToMultiPointCommandNormal) GetTypeName() string {
	return "CBusPointToMultiPointCommandNormal"
}

func (m *_CBusPointToMultiPointCommandNormal) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CBusPointToMultiPointCommandContract.(*_CBusPointToMultiPointCommand).getLengthInBits(ctx))

	// Simple field (application)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (salData)
	lengthInBits += m.SalData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CBusPointToMultiPointCommandNormal) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CBusPointToMultiPointCommandNormal) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CBusPointToMultiPointCommand, cBusOptions CBusOptions) (__cBusPointToMultiPointCommandNormal CBusPointToMultiPointCommandNormal, err error) {
	m.CBusPointToMultiPointCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusPointToMultiPointCommandNormal"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusPointToMultiPointCommandNormal")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	application, err := ReadEnumField[ApplicationIdContainer](ctx, "application", "ApplicationIdContainer", ReadEnum(ApplicationIdContainerByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'application' field"))
	}
	m.Application = application

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadByte(readBuffer, 8), byte(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	salData, err := ReadSimpleField[SALData](ctx, "salData", ReadComplex[SALData](SALDataParseWithBufferProducer[SALData]((ApplicationId)(application.ApplicationId())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'salData' field"))
	}
	m.SalData = salData

	if closeErr := readBuffer.CloseContext("CBusPointToMultiPointCommandNormal"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusPointToMultiPointCommandNormal")
	}

	return m, nil
}

func (m *_CBusPointToMultiPointCommandNormal) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CBusPointToMultiPointCommandNormal) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CBusPointToMultiPointCommandNormal"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CBusPointToMultiPointCommandNormal")
		}

		if err := WriteSimpleEnumField[ApplicationIdContainer](ctx, "application", "ApplicationIdContainer", m.GetApplication(), WriteEnum[ApplicationIdContainer, uint8](ApplicationIdContainer.GetValue, ApplicationIdContainer.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'application' field")
		}

		if err := WriteReservedField[byte](ctx, "reserved", byte(0x00), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[SALData](ctx, "salData", m.GetSalData(), WriteComplex[SALData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'salData' field")
		}

		if popErr := writeBuffer.PopContext("CBusPointToMultiPointCommandNormal"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CBusPointToMultiPointCommandNormal")
		}
		return nil
	}
	return m.CBusPointToMultiPointCommandContract.(*_CBusPointToMultiPointCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CBusPointToMultiPointCommandNormal) IsCBusPointToMultiPointCommandNormal() {}

func (m *_CBusPointToMultiPointCommandNormal) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
