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

// Constant values.
const CycServiceItemType_FUNCTIONID uint8 = 0x12

// CycServiceItemType is the corresponding interface of CycServiceItemType
type CycServiceItemType interface {
	CycServiceItemTypeContract
	CycServiceItemTypeRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// CycServiceItemTypeContract provides a set of functions which can be overwritten by a sub struct
type CycServiceItemTypeContract interface {
	// GetByteLength returns ByteLength (property field)
	GetByteLength() uint8
	// GetSyntaxId returns SyntaxId (property field)
	GetSyntaxId() uint8
}

// CycServiceItemTypeRequirements provides a set of functions which need to be implemented by a sub struct
type CycServiceItemTypeRequirements interface {
}

// CycServiceItemTypeExactly can be used when we want exactly this type and not a type which fulfills CycServiceItemType.
// This is useful for switch cases.
type CycServiceItemTypeExactly interface {
	CycServiceItemType
	isCycServiceItemType() bool
}

// _CycServiceItemType is the data-structure of this message
type _CycServiceItemType struct {
	_CycServiceItemTypeChildRequirements
	ByteLength uint8
	SyntaxId   uint8
}

var _ CycServiceItemTypeContract = (*_CycServiceItemType)(nil)

type _CycServiceItemTypeChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetSyntaxId() uint8
}

type CycServiceItemTypeParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child CycServiceItemType, serializeChildFunction func() error) error
	GetTypeName() string
}

type CycServiceItemTypeChild interface {
	utils.Serializable
	InitializeParent(parent CycServiceItemType, byteLength uint8, syntaxId uint8)
	GetParent() *CycServiceItemType

	GetTypeName() string
	CycServiceItemType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CycServiceItemType) GetByteLength() uint8 {
	return m.ByteLength
}

func (m *_CycServiceItemType) GetSyntaxId() uint8 {
	return m.SyntaxId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_CycServiceItemType) GetFunctionId() uint8 {
	return CycServiceItemType_FUNCTIONID
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCycServiceItemType factory function for _CycServiceItemType
func NewCycServiceItemType(byteLength uint8, syntaxId uint8) *_CycServiceItemType {
	return &_CycServiceItemType{ByteLength: byteLength, SyntaxId: syntaxId}
}

// Deprecated: use the interface for direct cast
func CastCycServiceItemType(structType any) CycServiceItemType {
	if casted, ok := structType.(CycServiceItemType); ok {
		return casted
	}
	if casted, ok := structType.(*CycServiceItemType); ok {
		return *casted
	}
	return nil
}

func (m *_CycServiceItemType) GetTypeName() string {
	return "CycServiceItemType"
}

func (m *_CycServiceItemType) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Const Field (functionId)
	lengthInBits += 8

	// Simple field (byteLength)
	lengthInBits += 8

	// Simple field (syntaxId)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CycServiceItemType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CycServiceItemTypeParse(ctx context.Context, theBytes []byte) (CycServiceItemType, error) {
	return CycServiceItemTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func CycServiceItemTypeParseWithBufferProducer[T CycServiceItemType]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		buffer, err := CycServiceItemTypeParseWithBuffer(ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return buffer.(T), err
	}
}

func CycServiceItemTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (CycServiceItemType, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CycServiceItemType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CycServiceItemType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	functionId, err := ReadConstField[uint8](ctx, "functionId", ReadUnsignedByte(readBuffer, uint8(8)), CycServiceItemType_FUNCTIONID)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'functionId' field"))
	}
	_ = functionId

	byteLength, err := ReadSimpleField(ctx, "byteLength", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'byteLength' field"))
	}

	syntaxId, err := ReadSimpleField(ctx, "syntaxId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'syntaxId' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type CycServiceItemTypeChildSerializeRequirement interface {
		CycServiceItemType
		InitializeParent(CycServiceItemType, uint8, uint8)
		GetParent() CycServiceItemType
	}
	var _childTemp any
	var _child CycServiceItemTypeChildSerializeRequirement
	var typeSwitchError error
	switch {
	case syntaxId == 0x10: // CycServiceItemAnyType
		_childTemp, typeSwitchError = CycServiceItemAnyTypeParseWithBuffer(ctx, readBuffer)
	case syntaxId == 0xb0: // CycServiceItemDbReadType
		_childTemp, typeSwitchError = CycServiceItemDbReadTypeParseWithBuffer(ctx, readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [syntaxId=%v]", syntaxId)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of CycServiceItemType")
	}
	_child = _childTemp.(CycServiceItemTypeChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("CycServiceItemType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CycServiceItemType")
	}

	// Finish initializing
	_child.InitializeParent(_child, byteLength, syntaxId)
	return _child, nil
}

func (pm *_CycServiceItemType) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child CycServiceItemType, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("CycServiceItemType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CycServiceItemType")
	}

	if err := WriteConstField(ctx, "functionId", CycServiceItemType_FUNCTIONID, WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'functionId' field")
	}

	if err := WriteSimpleField[uint8](ctx, "byteLength", m.GetByteLength(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'byteLength' field")
	}

	if err := WriteSimpleField[uint8](ctx, "syntaxId", m.GetSyntaxId(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'syntaxId' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("CycServiceItemType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CycServiceItemType")
	}
	return nil
}

func (m *_CycServiceItemType) isCycServiceItemType() bool {
	return true
}

func (m *_CycServiceItemType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
