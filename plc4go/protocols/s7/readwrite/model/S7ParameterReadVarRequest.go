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

// S7ParameterReadVarRequest is the corresponding interface of S7ParameterReadVarRequest
type S7ParameterReadVarRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7Parameter
	// GetItems returns Items (property field)
	GetItems() []S7VarRequestParameterItem
}

// S7ParameterReadVarRequestExactly can be used when we want exactly this type and not a type which fulfills S7ParameterReadVarRequest.
// This is useful for switch cases.
type S7ParameterReadVarRequestExactly interface {
	S7ParameterReadVarRequest
	isS7ParameterReadVarRequest() bool
}

// _S7ParameterReadVarRequest is the data-structure of this message
type _S7ParameterReadVarRequest struct {
	S7ParameterContract
	Items []S7VarRequestParameterItem
}

var _ S7ParameterReadVarRequest = (*_S7ParameterReadVarRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7ParameterReadVarRequest) GetParameterType() uint8 {
	return 0x04
}

func (m *_S7ParameterReadVarRequest) GetMessageType() uint8 {
	return 0x01
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7ParameterReadVarRequest) GetParent() S7ParameterContract {
	return m.S7ParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7ParameterReadVarRequest) GetItems() []S7VarRequestParameterItem {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7ParameterReadVarRequest factory function for _S7ParameterReadVarRequest
func NewS7ParameterReadVarRequest(items []S7VarRequestParameterItem) *_S7ParameterReadVarRequest {
	_result := &_S7ParameterReadVarRequest{
		S7ParameterContract: NewS7Parameter(),
		Items:               items,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7ParameterReadVarRequest(structType any) S7ParameterReadVarRequest {
	if casted, ok := structType.(S7ParameterReadVarRequest); ok {
		return casted
	}
	if casted, ok := structType.(*S7ParameterReadVarRequest); ok {
		return *casted
	}
	return nil
}

func (m *_S7ParameterReadVarRequest) GetTypeName() string {
	return "S7ParameterReadVarRequest"
}

func (m *_S7ParameterReadVarRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7ParameterContract.(*_S7Parameter).getLengthInBits(ctx))

	// Implicit Field (numItems)
	lengthInBits += 8

	// Array field
	if len(m.Items) > 0 {
		for _curItem, element := range m.Items {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Items), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_S7ParameterReadVarRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7ParameterReadVarRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7Parameter, messageType uint8) (_s7ParameterReadVarRequest S7ParameterReadVarRequest, err error) {
	m.S7ParameterContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7ParameterReadVarRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7ParameterReadVarRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	numItems, err := ReadImplicitField[uint8](ctx, "numItems", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numItems' field"))
	}
	_ = numItems

	items, err := ReadCountArrayField[S7VarRequestParameterItem](ctx, "items", ReadComplex[S7VarRequestParameterItem](S7VarRequestParameterItemParseWithBuffer, readBuffer), uint64(numItems))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'items' field"))
	}
	m.Items = items

	if closeErr := readBuffer.CloseContext("S7ParameterReadVarRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7ParameterReadVarRequest")
	}

	return m, nil
}

func (m *_S7ParameterReadVarRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7ParameterReadVarRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7ParameterReadVarRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7ParameterReadVarRequest")
		}
		numItems := uint8(uint8(len(m.GetItems())))
		if err := WriteImplicitField(ctx, "numItems", numItems, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'numItems' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "items", m.GetItems(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'items' field")
		}

		if popErr := writeBuffer.PopContext("S7ParameterReadVarRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7ParameterReadVarRequest")
		}
		return nil
	}
	return m.S7ParameterContract.(*_S7Parameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7ParameterReadVarRequest) isS7ParameterReadVarRequest() bool {
	return true
}

func (m *_S7ParameterReadVarRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
