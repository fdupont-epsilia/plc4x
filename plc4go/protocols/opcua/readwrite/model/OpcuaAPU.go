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
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// OpcuaAPU is the corresponding interface of OpcuaAPU
type OpcuaAPU interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetMessage returns Message (property field)
	GetMessage() MessagePDU
}

// OpcuaAPUExactly can be used when we want exactly this type and not a type which fulfills OpcuaAPU.
// This is useful for switch cases.
type OpcuaAPUExactly interface {
	OpcuaAPU
	isOpcuaAPU() bool
}

// _OpcuaAPU is the data-structure of this message
type _OpcuaAPU struct {
	Message MessagePDU

	// Arguments.
	Response bool
}

var _ OpcuaAPU = (*_OpcuaAPU)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_OpcuaAPU) GetMessage() MessagePDU {
	return m.Message
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewOpcuaAPU factory function for _OpcuaAPU
func NewOpcuaAPU(message MessagePDU, response bool) *_OpcuaAPU {
	return &_OpcuaAPU{Message: message, Response: response}
}

// Deprecated: use the interface for direct cast
func CastOpcuaAPU(structType any) OpcuaAPU {
	if casted, ok := structType.(OpcuaAPU); ok {
		return casted
	}
	if casted, ok := structType.(*OpcuaAPU); ok {
		return *casted
	}
	return nil
}

func (m *_OpcuaAPU) GetTypeName() string {
	return "OpcuaAPU"
}

func (m *_OpcuaAPU) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (message)
	lengthInBits += m.Message.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_OpcuaAPU) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaAPUParse(ctx context.Context, theBytes []byte, response bool) (OpcuaAPU, error) {
	return OpcuaAPUParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.LittleEndian)), response)
}

func OpcuaAPUParseWithBufferProducer(response bool) func(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaAPU, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaAPU, error) {
		return OpcuaAPUParseWithBuffer(ctx, readBuffer, response)
	}
}

func OpcuaAPUParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (OpcuaAPU, error) {
	v, err := (&_OpcuaAPU{}).parse(ctx, readBuffer, response)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_OpcuaAPU) parse(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (_opcuaAPU OpcuaAPU, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("OpcuaAPU"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for OpcuaAPU")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	message, err := ReadSimpleField[MessagePDU](ctx, "message", ReadComplex[MessagePDU](MessagePDUParseWithBufferProducer[MessagePDU]((bool)(response)), readBuffer), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'message' field"))
	}
	m.Message = message

	if closeErr := readBuffer.CloseContext("OpcuaAPU"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for OpcuaAPU")
	}

	return m, nil
}

func (m *_OpcuaAPU) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.LittleEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_OpcuaAPU) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("OpcuaAPU"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for OpcuaAPU")
	}

	if err := WriteSimpleField[MessagePDU](ctx, "message", m.GetMessage(), WriteComplex[MessagePDU](writeBuffer), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'message' field")
	}

	if popErr := writeBuffer.PopContext("OpcuaAPU"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for OpcuaAPU")
	}
	return nil
}

////
// Arguments Getter

func (m *_OpcuaAPU) GetResponse() bool {
	return m.Response
}

//
////

func (m *_OpcuaAPU) isOpcuaAPU() bool {
	return true
}

func (m *_OpcuaAPU) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
