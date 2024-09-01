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

// Constant values.
const RequestSmartConnectShortcut_PIPE byte = 0x7C

// RequestSmartConnectShortcut is the corresponding interface of RequestSmartConnectShortcut
type RequestSmartConnectShortcut interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	Request
	// GetPipePeek returns PipePeek (property field)
	GetPipePeek() RequestType
	// GetSecondPipe returns SecondPipe (property field)
	GetSecondPipe() *byte
}

// RequestSmartConnectShortcutExactly can be used when we want exactly this type and not a type which fulfills RequestSmartConnectShortcut.
// This is useful for switch cases.
type RequestSmartConnectShortcutExactly interface {
	RequestSmartConnectShortcut
	isRequestSmartConnectShortcut() bool
}

// _RequestSmartConnectShortcut is the data-structure of this message
type _RequestSmartConnectShortcut struct {
	RequestContract
	PipePeek   RequestType
	SecondPipe *byte
}

var _ RequestSmartConnectShortcut = (*_RequestSmartConnectShortcut)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RequestSmartConnectShortcut) GetParent() RequestContract {
	return m.RequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RequestSmartConnectShortcut) GetPipePeek() RequestType {
	return m.PipePeek
}

func (m *_RequestSmartConnectShortcut) GetSecondPipe() *byte {
	return m.SecondPipe
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_RequestSmartConnectShortcut) GetPipe() byte {
	return RequestSmartConnectShortcut_PIPE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewRequestSmartConnectShortcut factory function for _RequestSmartConnectShortcut
func NewRequestSmartConnectShortcut(pipePeek RequestType, secondPipe *byte, peekedByte RequestType, startingCR *RequestType, resetMode *RequestType, secondPeek RequestType, termination RequestTermination, cBusOptions CBusOptions) *_RequestSmartConnectShortcut {
	_result := &_RequestSmartConnectShortcut{
		RequestContract: NewRequest(peekedByte, startingCR, resetMode, secondPeek, termination, cBusOptions),
		PipePeek:        pipePeek,
		SecondPipe:      secondPipe,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastRequestSmartConnectShortcut(structType any) RequestSmartConnectShortcut {
	if casted, ok := structType.(RequestSmartConnectShortcut); ok {
		return casted
	}
	if casted, ok := structType.(*RequestSmartConnectShortcut); ok {
		return *casted
	}
	return nil
}

func (m *_RequestSmartConnectShortcut) GetTypeName() string {
	return "RequestSmartConnectShortcut"
}

func (m *_RequestSmartConnectShortcut) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.RequestContract.(*_Request).getLengthInBits(ctx))

	// Const Field (pipe)
	lengthInBits += 8

	// Optional Field (secondPipe)
	if m.SecondPipe != nil {
		lengthInBits += 8
	}

	return lengthInBits
}

func (m *_RequestSmartConnectShortcut) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_RequestSmartConnectShortcut) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_Request, cBusOptions CBusOptions) (_requestSmartConnectShortcut RequestSmartConnectShortcut, err error) {
	m.RequestContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RequestSmartConnectShortcut"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RequestSmartConnectShortcut")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	pipe, err := ReadConstField[byte](ctx, "pipe", ReadByte(readBuffer, 8), RequestSmartConnectShortcut_PIPE)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'pipe' field"))
	}
	_ = pipe

	pipePeek, err := ReadPeekField[RequestType](ctx, "pipePeek", ReadEnum(RequestTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'pipePeek' field"))
	}
	m.PipePeek = pipePeek

	secondPipe, err := ReadOptionalField[byte](ctx, "secondPipe", ReadByte(readBuffer, 8), bool((pipePeek) == (RequestType_SMART_CONNECT_SHORTCUT)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'secondPipe' field"))
	}
	m.SecondPipe = secondPipe

	if closeErr := readBuffer.CloseContext("RequestSmartConnectShortcut"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RequestSmartConnectShortcut")
	}

	return m, nil
}

func (m *_RequestSmartConnectShortcut) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RequestSmartConnectShortcut) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RequestSmartConnectShortcut"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RequestSmartConnectShortcut")
		}

		if err := WriteConstField(ctx, "pipe", RequestSmartConnectShortcut_PIPE, WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'pipe' field")
		}

		if err := WriteOptionalField[byte](ctx, "secondPipe", m.GetSecondPipe(), WriteByte(writeBuffer, 8), true); err != nil {
			return errors.Wrap(err, "Error serializing 'secondPipe' field")
		}

		if popErr := writeBuffer.PopContext("RequestSmartConnectShortcut"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RequestSmartConnectShortcut")
		}
		return nil
	}
	return m.RequestContract.(*_Request).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RequestSmartConnectShortcut) isRequestSmartConnectShortcut() bool {
	return true
}

func (m *_RequestSmartConnectShortcut) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
