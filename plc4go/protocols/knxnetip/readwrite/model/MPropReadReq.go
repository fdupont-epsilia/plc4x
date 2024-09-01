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

// MPropReadReq is the corresponding interface of MPropReadReq
type MPropReadReq interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CEMI
	// GetInterfaceObjectType returns InterfaceObjectType (property field)
	GetInterfaceObjectType() uint16
	// GetObjectInstance returns ObjectInstance (property field)
	GetObjectInstance() uint8
	// GetPropertyId returns PropertyId (property field)
	GetPropertyId() uint8
	// GetNumberOfElements returns NumberOfElements (property field)
	GetNumberOfElements() uint8
	// GetStartIndex returns StartIndex (property field)
	GetStartIndex() uint16
}

// MPropReadReqExactly can be used when we want exactly this type and not a type which fulfills MPropReadReq.
// This is useful for switch cases.
type MPropReadReqExactly interface {
	MPropReadReq
	isMPropReadReq() bool
}

// _MPropReadReq is the data-structure of this message
type _MPropReadReq struct {
	CEMIContract
	InterfaceObjectType uint16
	ObjectInstance      uint8
	PropertyId          uint8
	NumberOfElements    uint8
	StartIndex          uint16
}

var _ MPropReadReq = (*_MPropReadReq)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MPropReadReq) GetMessageCode() uint8 {
	return 0xFC
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MPropReadReq) GetParent() CEMIContract {
	return m.CEMIContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MPropReadReq) GetInterfaceObjectType() uint16 {
	return m.InterfaceObjectType
}

func (m *_MPropReadReq) GetObjectInstance() uint8 {
	return m.ObjectInstance
}

func (m *_MPropReadReq) GetPropertyId() uint8 {
	return m.PropertyId
}

func (m *_MPropReadReq) GetNumberOfElements() uint8 {
	return m.NumberOfElements
}

func (m *_MPropReadReq) GetStartIndex() uint16 {
	return m.StartIndex
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMPropReadReq factory function for _MPropReadReq
func NewMPropReadReq(interfaceObjectType uint16, objectInstance uint8, propertyId uint8, numberOfElements uint8, startIndex uint16, size uint16) *_MPropReadReq {
	_result := &_MPropReadReq{
		CEMIContract:        NewCEMI(size),
		InterfaceObjectType: interfaceObjectType,
		ObjectInstance:      objectInstance,
		PropertyId:          propertyId,
		NumberOfElements:    numberOfElements,
		StartIndex:          startIndex,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastMPropReadReq(structType any) MPropReadReq {
	if casted, ok := structType.(MPropReadReq); ok {
		return casted
	}
	if casted, ok := structType.(*MPropReadReq); ok {
		return *casted
	}
	return nil
}

func (m *_MPropReadReq) GetTypeName() string {
	return "MPropReadReq"
}

func (m *_MPropReadReq) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CEMIContract.(*_CEMI).getLengthInBits(ctx))

	// Simple field (interfaceObjectType)
	lengthInBits += 16

	// Simple field (objectInstance)
	lengthInBits += 8

	// Simple field (propertyId)
	lengthInBits += 8

	// Simple field (numberOfElements)
	lengthInBits += 4

	// Simple field (startIndex)
	lengthInBits += 12

	return lengthInBits
}

func (m *_MPropReadReq) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MPropReadReq) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CEMI, size uint16) (_mPropReadReq MPropReadReq, err error) {
	m.CEMIContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MPropReadReq"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MPropReadReq")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	interfaceObjectType, err := ReadSimpleField(ctx, "interfaceObjectType", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'interfaceObjectType' field"))
	}
	m.InterfaceObjectType = interfaceObjectType

	objectInstance, err := ReadSimpleField(ctx, "objectInstance", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectInstance' field"))
	}
	m.ObjectInstance = objectInstance

	propertyId, err := ReadSimpleField(ctx, "propertyId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyId' field"))
	}
	m.PropertyId = propertyId

	numberOfElements, err := ReadSimpleField(ctx, "numberOfElements", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfElements' field"))
	}
	m.NumberOfElements = numberOfElements

	startIndex, err := ReadSimpleField(ctx, "startIndex", ReadUnsignedShort(readBuffer, uint8(12)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'startIndex' field"))
	}
	m.StartIndex = startIndex

	if closeErr := readBuffer.CloseContext("MPropReadReq"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MPropReadReq")
	}

	return m, nil
}

func (m *_MPropReadReq) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MPropReadReq) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MPropReadReq"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MPropReadReq")
		}

		if err := WriteSimpleField[uint16](ctx, "interfaceObjectType", m.GetInterfaceObjectType(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'interfaceObjectType' field")
		}

		if err := WriteSimpleField[uint8](ctx, "objectInstance", m.GetObjectInstance(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'objectInstance' field")
		}

		if err := WriteSimpleField[uint8](ctx, "propertyId", m.GetPropertyId(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'propertyId' field")
		}

		if err := WriteSimpleField[uint8](ctx, "numberOfElements", m.GetNumberOfElements(), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'numberOfElements' field")
		}

		if err := WriteSimpleField[uint16](ctx, "startIndex", m.GetStartIndex(), WriteUnsignedShort(writeBuffer, 12)); err != nil {
			return errors.Wrap(err, "Error serializing 'startIndex' field")
		}

		if popErr := writeBuffer.PopContext("MPropReadReq"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MPropReadReq")
		}
		return nil
	}
	return m.CEMIContract.(*_CEMI).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MPropReadReq) isMPropReadReq() bool {
	return true
}

func (m *_MPropReadReq) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
