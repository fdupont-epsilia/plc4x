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

// RequestReset is the corresponding interface of RequestReset
type RequestReset interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	Request
	// GetTildePeek returns TildePeek (property field)
	GetTildePeek() RequestType
	// GetSecondTilde returns SecondTilde (property field)
	GetSecondTilde() *RequestType
	// GetTildePeek2 returns TildePeek2 (property field)
	GetTildePeek2() RequestType
	// GetThirdTilde returns ThirdTilde (property field)
	GetThirdTilde() *RequestType
	// IsRequestReset is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsRequestReset()
}

// _RequestReset is the data-structure of this message
type _RequestReset struct {
	RequestContract
	TildePeek   RequestType
	SecondTilde *RequestType
	TildePeek2  RequestType
	ThirdTilde  *RequestType
}

var _ RequestReset = (*_RequestReset)(nil)
var _ RequestRequirements = (*_RequestReset)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RequestReset) GetParent() RequestContract {
	return m.RequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RequestReset) GetTildePeek() RequestType {
	return m.TildePeek
}

func (m *_RequestReset) GetSecondTilde() *RequestType {
	return m.SecondTilde
}

func (m *_RequestReset) GetTildePeek2() RequestType {
	return m.TildePeek2
}

func (m *_RequestReset) GetThirdTilde() *RequestType {
	return m.ThirdTilde
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewRequestReset factory function for _RequestReset
func NewRequestReset(tildePeek RequestType, secondTilde *RequestType, tildePeek2 RequestType, thirdTilde *RequestType, peekedByte RequestType, startingCR *RequestType, resetMode *RequestType, secondPeek RequestType, termination RequestTermination, cBusOptions CBusOptions) *_RequestReset {
	_result := &_RequestReset{
		RequestContract: NewRequest(peekedByte, startingCR, resetMode, secondPeek, termination, cBusOptions),
		TildePeek:       tildePeek,
		SecondTilde:     secondTilde,
		TildePeek2:      tildePeek2,
		ThirdTilde:      thirdTilde,
	}
	_result.RequestContract.(*_Request)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastRequestReset(structType any) RequestReset {
	if casted, ok := structType.(RequestReset); ok {
		return casted
	}
	if casted, ok := structType.(*RequestReset); ok {
		return *casted
	}
	return nil
}

func (m *_RequestReset) GetTypeName() string {
	return "RequestReset"
}

func (m *_RequestReset) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.RequestContract.(*_Request).getLengthInBits(ctx))

	// Optional Field (secondTilde)
	if m.SecondTilde != nil {
		lengthInBits += 8
	}

	// Optional Field (thirdTilde)
	if m.ThirdTilde != nil {
		lengthInBits += 8
	}

	return lengthInBits
}

func (m *_RequestReset) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_RequestReset) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_Request, cBusOptions CBusOptions) (__requestReset RequestReset, err error) {
	m.RequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RequestReset"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RequestReset")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	tildePeek, err := ReadPeekField[RequestType](ctx, "tildePeek", ReadEnum(RequestTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'tildePeek' field"))
	}
	m.TildePeek = tildePeek

	var secondTilde *RequestType
	secondTilde, err = ReadOptionalField[RequestType](ctx, "secondTilde", ReadEnum(RequestTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))), bool((tildePeek) == (RequestType_RESET)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'secondTilde' field"))
	}
	m.SecondTilde = secondTilde

	tildePeek2, err := ReadPeekField[RequestType](ctx, "tildePeek2", ReadEnum(RequestTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'tildePeek2' field"))
	}
	m.TildePeek2 = tildePeek2

	var thirdTilde *RequestType
	thirdTilde, err = ReadOptionalField[RequestType](ctx, "thirdTilde", ReadEnum(RequestTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))), bool((tildePeek2) == (RequestType_RESET)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'thirdTilde' field"))
	}
	m.ThirdTilde = thirdTilde

	if closeErr := readBuffer.CloseContext("RequestReset"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RequestReset")
	}

	return m, nil
}

func (m *_RequestReset) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RequestReset) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RequestReset"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RequestReset")
		}

		if err := WriteOptionalEnumField[RequestType](ctx, "secondTilde", "RequestType", m.GetSecondTilde(), WriteEnum[RequestType, uint8](RequestType.GetValue, RequestType.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8)), bool((m.GetTildePeek()) == (RequestType_RESET))); err != nil {
			return errors.Wrap(err, "Error serializing 'secondTilde' field")
		}

		if err := WriteOptionalEnumField[RequestType](ctx, "thirdTilde", "RequestType", m.GetThirdTilde(), WriteEnum[RequestType, uint8](RequestType.GetValue, RequestType.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8)), bool((m.GetTildePeek2()) == (RequestType_RESET))); err != nil {
			return errors.Wrap(err, "Error serializing 'thirdTilde' field")
		}

		if popErr := writeBuffer.PopContext("RequestReset"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RequestReset")
		}
		return nil
	}
	return m.RequestContract.(*_Request).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RequestReset) IsRequestReset() {}

func (m *_RequestReset) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
