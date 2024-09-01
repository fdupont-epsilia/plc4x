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

// APDUSimpleAck is the corresponding interface of APDUSimpleAck
type APDUSimpleAck interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	APDU
	// GetOriginalInvokeId returns OriginalInvokeId (property field)
	GetOriginalInvokeId() uint8
	// GetServiceChoice returns ServiceChoice (property field)
	GetServiceChoice() BACnetConfirmedServiceChoice
}

// APDUSimpleAckExactly can be used when we want exactly this type and not a type which fulfills APDUSimpleAck.
// This is useful for switch cases.
type APDUSimpleAckExactly interface {
	APDUSimpleAck
	isAPDUSimpleAck() bool
}

// _APDUSimpleAck is the data-structure of this message
type _APDUSimpleAck struct {
	*_APDU
	OriginalInvokeId uint8
	ServiceChoice    BACnetConfirmedServiceChoice
	// Reserved Fields
	reservedField0 *uint8
}

var _ APDUSimpleAck = (*_APDUSimpleAck)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_APDUSimpleAck) GetApduType() ApduType {
	return ApduType_SIMPLE_ACK_PDU
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_APDUSimpleAck) InitializeParent(parent APDU) {}

func (m *_APDUSimpleAck) GetParent() APDU {
	return m._APDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_APDUSimpleAck) GetOriginalInvokeId() uint8 {
	return m.OriginalInvokeId
}

func (m *_APDUSimpleAck) GetServiceChoice() BACnetConfirmedServiceChoice {
	return m.ServiceChoice
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAPDUSimpleAck factory function for _APDUSimpleAck
func NewAPDUSimpleAck(originalInvokeId uint8, serviceChoice BACnetConfirmedServiceChoice, apduLength uint16) *_APDUSimpleAck {
	_result := &_APDUSimpleAck{
		OriginalInvokeId: originalInvokeId,
		ServiceChoice:    serviceChoice,
		_APDU:            NewAPDU(apduLength),
	}
	_result._APDU._APDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAPDUSimpleAck(structType any) APDUSimpleAck {
	if casted, ok := structType.(APDUSimpleAck); ok {
		return casted
	}
	if casted, ok := structType.(*APDUSimpleAck); ok {
		return *casted
	}
	return nil
}

func (m *_APDUSimpleAck) GetTypeName() string {
	return "APDUSimpleAck"
}

func (m *_APDUSimpleAck) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 4

	// Simple field (originalInvokeId)
	lengthInBits += 8

	// Simple field (serviceChoice)
	lengthInBits += 8

	return lengthInBits
}

func (m *_APDUSimpleAck) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func APDUSimpleAckParse(ctx context.Context, theBytes []byte, apduLength uint16) (APDUSimpleAck, error) {
	return APDUSimpleAckParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), apduLength)
}

func APDUSimpleAckParseWithBufferProducer(apduLength uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (APDUSimpleAck, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (APDUSimpleAck, error) {
		return APDUSimpleAckParseWithBuffer(ctx, readBuffer, apduLength)
	}
}

func APDUSimpleAckParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, apduLength uint16) (APDUSimpleAck, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("APDUSimpleAck"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for APDUSimpleAck")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(4)), uint8(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	originalInvokeId, err := ReadSimpleField(ctx, "originalInvokeId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'originalInvokeId' field"))
	}

	serviceChoice, err := ReadEnumField[BACnetConfirmedServiceChoice](ctx, "serviceChoice", "BACnetConfirmedServiceChoice", ReadEnum(BACnetConfirmedServiceChoiceByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serviceChoice' field"))
	}

	if closeErr := readBuffer.CloseContext("APDUSimpleAck"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for APDUSimpleAck")
	}

	// Create a partially initialized instance
	_child := &_APDUSimpleAck{
		_APDU: &_APDU{
			ApduLength: apduLength,
		},
		OriginalInvokeId: originalInvokeId,
		ServiceChoice:    serviceChoice,
		reservedField0:   reservedField0,
	}
	_child._APDU._APDUChildRequirements = _child
	return _child, nil
}

func (m *_APDUSimpleAck) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_APDUSimpleAck) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUSimpleAck"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for APDUSimpleAck")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[uint8](ctx, "originalInvokeId", m.GetOriginalInvokeId(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'originalInvokeId' field")
		}

		if err := WriteSimpleEnumField[BACnetConfirmedServiceChoice](ctx, "serviceChoice", "BACnetConfirmedServiceChoice", m.GetServiceChoice(), WriteEnum[BACnetConfirmedServiceChoice, uint8](BACnetConfirmedServiceChoice.GetValue, BACnetConfirmedServiceChoice.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'serviceChoice' field")
		}

		if popErr := writeBuffer.PopContext("APDUSimpleAck"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for APDUSimpleAck")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_APDUSimpleAck) isAPDUSimpleAck() bool {
	return true
}

func (m *_APDUSimpleAck) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
