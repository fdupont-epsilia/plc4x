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

// S7VarPayloadStatusItem is the corresponding interface of S7VarPayloadStatusItem
type S7VarPayloadStatusItem interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetReturnCode returns ReturnCode (property field)
	GetReturnCode() DataTransportErrorCode
}

// S7VarPayloadStatusItemExactly can be used when we want exactly this type and not a type which fulfills S7VarPayloadStatusItem.
// This is useful for switch cases.
type S7VarPayloadStatusItemExactly interface {
	S7VarPayloadStatusItem
	isS7VarPayloadStatusItem() bool
}

// _S7VarPayloadStatusItem is the data-structure of this message
type _S7VarPayloadStatusItem struct {
	ReturnCode DataTransportErrorCode
}

var _ S7VarPayloadStatusItem = (*_S7VarPayloadStatusItem)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7VarPayloadStatusItem) GetReturnCode() DataTransportErrorCode {
	return m.ReturnCode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7VarPayloadStatusItem factory function for _S7VarPayloadStatusItem
func NewS7VarPayloadStatusItem(returnCode DataTransportErrorCode) *_S7VarPayloadStatusItem {
	return &_S7VarPayloadStatusItem{ReturnCode: returnCode}
}

// Deprecated: use the interface for direct cast
func CastS7VarPayloadStatusItem(structType any) S7VarPayloadStatusItem {
	if casted, ok := structType.(S7VarPayloadStatusItem); ok {
		return casted
	}
	if casted, ok := structType.(*S7VarPayloadStatusItem); ok {
		return *casted
	}
	return nil
}

func (m *_S7VarPayloadStatusItem) GetTypeName() string {
	return "S7VarPayloadStatusItem"
}

func (m *_S7VarPayloadStatusItem) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (returnCode)
	lengthInBits += 8

	return lengthInBits
}

func (m *_S7VarPayloadStatusItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7VarPayloadStatusItemParse(ctx context.Context, theBytes []byte) (S7VarPayloadStatusItem, error) {
	return S7VarPayloadStatusItemParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func S7VarPayloadStatusItemParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (S7VarPayloadStatusItem, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (S7VarPayloadStatusItem, error) {
		return S7VarPayloadStatusItemParseWithBuffer(ctx, readBuffer)
	}
}

func S7VarPayloadStatusItemParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (S7VarPayloadStatusItem, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7VarPayloadStatusItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7VarPayloadStatusItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	returnCode, err := ReadEnumField[DataTransportErrorCode](ctx, "returnCode", "DataTransportErrorCode", ReadEnum(DataTransportErrorCodeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'returnCode' field"))
	}

	if closeErr := readBuffer.CloseContext("S7VarPayloadStatusItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7VarPayloadStatusItem")
	}

	// Create the instance
	return &_S7VarPayloadStatusItem{
		ReturnCode: returnCode,
	}, nil
}

func (m *_S7VarPayloadStatusItem) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7VarPayloadStatusItem) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("S7VarPayloadStatusItem"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for S7VarPayloadStatusItem")
	}

	if err := WriteSimpleEnumField[DataTransportErrorCode](ctx, "returnCode", "DataTransportErrorCode", m.GetReturnCode(), WriteEnum[DataTransportErrorCode, uint8](DataTransportErrorCode.GetValue, DataTransportErrorCode.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'returnCode' field")
	}

	if popErr := writeBuffer.PopContext("S7VarPayloadStatusItem"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for S7VarPayloadStatusItem")
	}
	return nil
}

func (m *_S7VarPayloadStatusItem) isS7VarPayloadStatusItem() bool {
	return true
}

func (m *_S7VarPayloadStatusItem) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
