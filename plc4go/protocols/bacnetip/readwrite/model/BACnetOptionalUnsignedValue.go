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

// BACnetOptionalUnsignedValue is the corresponding interface of BACnetOptionalUnsignedValue
type BACnetOptionalUnsignedValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetOptionalUnsigned
	// GetUnsignedValue returns UnsignedValue (property field)
	GetUnsignedValue() BACnetApplicationTagUnsignedInteger
}

// BACnetOptionalUnsignedValueExactly can be used when we want exactly this type and not a type which fulfills BACnetOptionalUnsignedValue.
// This is useful for switch cases.
type BACnetOptionalUnsignedValueExactly interface {
	BACnetOptionalUnsignedValue
	isBACnetOptionalUnsignedValue() bool
}

// _BACnetOptionalUnsignedValue is the data-structure of this message
type _BACnetOptionalUnsignedValue struct {
	*_BACnetOptionalUnsigned
	UnsignedValue BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetOptionalUnsignedValue) InitializeParent(parent BACnetOptionalUnsigned, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetOptionalUnsignedValue) GetParent() BACnetOptionalUnsigned {
	return m._BACnetOptionalUnsigned
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetOptionalUnsignedValue) GetUnsignedValue() BACnetApplicationTagUnsignedInteger {
	return m.UnsignedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetOptionalUnsignedValue factory function for _BACnetOptionalUnsignedValue
func NewBACnetOptionalUnsignedValue(unsignedValue BACnetApplicationTagUnsignedInteger, peekedTagHeader BACnetTagHeader) *_BACnetOptionalUnsignedValue {
	_result := &_BACnetOptionalUnsignedValue{
		UnsignedValue:           unsignedValue,
		_BACnetOptionalUnsigned: NewBACnetOptionalUnsigned(peekedTagHeader),
	}
	_result._BACnetOptionalUnsigned._BACnetOptionalUnsignedChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetOptionalUnsignedValue(structType any) BACnetOptionalUnsignedValue {
	if casted, ok := structType.(BACnetOptionalUnsignedValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetOptionalUnsignedValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetOptionalUnsignedValue) GetTypeName() string {
	return "BACnetOptionalUnsignedValue"
}

func (m *_BACnetOptionalUnsignedValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (unsignedValue)
	lengthInBits += m.UnsignedValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetOptionalUnsignedValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetOptionalUnsignedValueParse(ctx context.Context, theBytes []byte) (BACnetOptionalUnsignedValue, error) {
	return BACnetOptionalUnsignedValueParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetOptionalUnsignedValueParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetOptionalUnsignedValue, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetOptionalUnsignedValue, error) {
		return BACnetOptionalUnsignedValueParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetOptionalUnsignedValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetOptionalUnsignedValue, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetOptionalUnsignedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetOptionalUnsignedValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	unsignedValue, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "unsignedValue", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer(), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unsignedValue' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetOptionalUnsignedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetOptionalUnsignedValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetOptionalUnsignedValue{
		_BACnetOptionalUnsigned: &_BACnetOptionalUnsigned{},
		UnsignedValue:           unsignedValue,
	}
	_child._BACnetOptionalUnsigned._BACnetOptionalUnsignedChildRequirements = _child
	return _child, nil
}

func (m *_BACnetOptionalUnsignedValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetOptionalUnsignedValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetOptionalUnsignedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetOptionalUnsignedValue")
		}

		// Simple Field (unsignedValue)
		if pushErr := writeBuffer.PushContext("unsignedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for unsignedValue")
		}
		_unsignedValueErr := writeBuffer.WriteSerializable(ctx, m.GetUnsignedValue())
		if popErr := writeBuffer.PopContext("unsignedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for unsignedValue")
		}
		if _unsignedValueErr != nil {
			return errors.Wrap(_unsignedValueErr, "Error serializing 'unsignedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetOptionalUnsignedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetOptionalUnsignedValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetOptionalUnsignedValue) isBACnetOptionalUnsignedValue() bool {
	return true
}

func (m *_BACnetOptionalUnsignedValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
