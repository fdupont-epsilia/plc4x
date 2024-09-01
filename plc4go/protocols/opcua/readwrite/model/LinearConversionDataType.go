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

// LinearConversionDataType is the corresponding interface of LinearConversionDataType
type LinearConversionDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetInitialAddend returns InitialAddend (property field)
	GetInitialAddend() float32
	// GetMultiplicand returns Multiplicand (property field)
	GetMultiplicand() float32
	// GetDivisor returns Divisor (property field)
	GetDivisor() float32
	// GetFinalAddend returns FinalAddend (property field)
	GetFinalAddend() float32
}

// LinearConversionDataTypeExactly can be used when we want exactly this type and not a type which fulfills LinearConversionDataType.
// This is useful for switch cases.
type LinearConversionDataTypeExactly interface {
	LinearConversionDataType
	isLinearConversionDataType() bool
}

// _LinearConversionDataType is the data-structure of this message
type _LinearConversionDataType struct {
	*_ExtensionObjectDefinition
	InitialAddend float32
	Multiplicand  float32
	Divisor       float32
	FinalAddend   float32
}

var _ LinearConversionDataType = (*_LinearConversionDataType)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_LinearConversionDataType) GetIdentifier() string {
	return "32437"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LinearConversionDataType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_LinearConversionDataType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LinearConversionDataType) GetInitialAddend() float32 {
	return m.InitialAddend
}

func (m *_LinearConversionDataType) GetMultiplicand() float32 {
	return m.Multiplicand
}

func (m *_LinearConversionDataType) GetDivisor() float32 {
	return m.Divisor
}

func (m *_LinearConversionDataType) GetFinalAddend() float32 {
	return m.FinalAddend
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewLinearConversionDataType factory function for _LinearConversionDataType
func NewLinearConversionDataType(initialAddend float32, multiplicand float32, divisor float32, finalAddend float32) *_LinearConversionDataType {
	_result := &_LinearConversionDataType{
		InitialAddend:              initialAddend,
		Multiplicand:               multiplicand,
		Divisor:                    divisor,
		FinalAddend:                finalAddend,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastLinearConversionDataType(structType any) LinearConversionDataType {
	if casted, ok := structType.(LinearConversionDataType); ok {
		return casted
	}
	if casted, ok := structType.(*LinearConversionDataType); ok {
		return *casted
	}
	return nil
}

func (m *_LinearConversionDataType) GetTypeName() string {
	return "LinearConversionDataType"
}

func (m *_LinearConversionDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (initialAddend)
	lengthInBits += 32

	// Simple field (multiplicand)
	lengthInBits += 32

	// Simple field (divisor)
	lengthInBits += 32

	// Simple field (finalAddend)
	lengthInBits += 32

	return lengthInBits
}

func (m *_LinearConversionDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func LinearConversionDataTypeParse(ctx context.Context, theBytes []byte, identifier string) (LinearConversionDataType, error) {
	return LinearConversionDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func LinearConversionDataTypeParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (LinearConversionDataType, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (LinearConversionDataType, error) {
		return LinearConversionDataTypeParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func LinearConversionDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (LinearConversionDataType, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LinearConversionDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LinearConversionDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	initialAddend, err := ReadSimpleField(ctx, "initialAddend", ReadFloat(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'initialAddend' field"))
	}

	multiplicand, err := ReadSimpleField(ctx, "multiplicand", ReadFloat(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'multiplicand' field"))
	}

	divisor, err := ReadSimpleField(ctx, "divisor", ReadFloat(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'divisor' field"))
	}

	finalAddend, err := ReadSimpleField(ctx, "finalAddend", ReadFloat(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'finalAddend' field"))
	}

	if closeErr := readBuffer.CloseContext("LinearConversionDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LinearConversionDataType")
	}

	// Create a partially initialized instance
	_child := &_LinearConversionDataType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		InitialAddend:              initialAddend,
		Multiplicand:               multiplicand,
		Divisor:                    divisor,
		FinalAddend:                finalAddend,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_LinearConversionDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LinearConversionDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LinearConversionDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LinearConversionDataType")
		}

		if err := WriteSimpleField[float32](ctx, "initialAddend", m.GetInitialAddend(), WriteFloat(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'initialAddend' field")
		}

		if err := WriteSimpleField[float32](ctx, "multiplicand", m.GetMultiplicand(), WriteFloat(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'multiplicand' field")
		}

		if err := WriteSimpleField[float32](ctx, "divisor", m.GetDivisor(), WriteFloat(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'divisor' field")
		}

		if err := WriteSimpleField[float32](ctx, "finalAddend", m.GetFinalAddend(), WriteFloat(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'finalAddend' field")
		}

		if popErr := writeBuffer.PopContext("LinearConversionDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LinearConversionDataType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_LinearConversionDataType) isLinearConversionDataType() bool {
	return true
}

func (m *_LinearConversionDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
