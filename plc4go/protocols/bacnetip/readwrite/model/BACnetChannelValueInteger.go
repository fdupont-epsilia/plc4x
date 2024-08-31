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

// BACnetChannelValueInteger is the corresponding interface of BACnetChannelValueInteger
type BACnetChannelValueInteger interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetChannelValue
	// GetIntegerValue returns IntegerValue (property field)
	GetIntegerValue() BACnetApplicationTagSignedInteger
}

// BACnetChannelValueIntegerExactly can be used when we want exactly this type and not a type which fulfills BACnetChannelValueInteger.
// This is useful for switch cases.
type BACnetChannelValueIntegerExactly interface {
	BACnetChannelValueInteger
	isBACnetChannelValueInteger() bool
}

// _BACnetChannelValueInteger is the data-structure of this message
type _BACnetChannelValueInteger struct {
	*_BACnetChannelValue
	IntegerValue BACnetApplicationTagSignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetChannelValueInteger) InitializeParent(parent BACnetChannelValue, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetChannelValueInteger) GetParent() BACnetChannelValue {
	return m._BACnetChannelValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetChannelValueInteger) GetIntegerValue() BACnetApplicationTagSignedInteger {
	return m.IntegerValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetChannelValueInteger factory function for _BACnetChannelValueInteger
func NewBACnetChannelValueInteger(integerValue BACnetApplicationTagSignedInteger, peekedTagHeader BACnetTagHeader) *_BACnetChannelValueInteger {
	_result := &_BACnetChannelValueInteger{
		IntegerValue:        integerValue,
		_BACnetChannelValue: NewBACnetChannelValue(peekedTagHeader),
	}
	_result._BACnetChannelValue._BACnetChannelValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetChannelValueInteger(structType any) BACnetChannelValueInteger {
	if casted, ok := structType.(BACnetChannelValueInteger); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetChannelValueInteger); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetChannelValueInteger) GetTypeName() string {
	return "BACnetChannelValueInteger"
}

func (m *_BACnetChannelValueInteger) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (integerValue)
	lengthInBits += m.IntegerValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetChannelValueInteger) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetChannelValueIntegerParse(ctx context.Context, theBytes []byte) (BACnetChannelValueInteger, error) {
	return BACnetChannelValueIntegerParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetChannelValueIntegerParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetChannelValueInteger, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetChannelValueInteger, error) {
		return BACnetChannelValueIntegerParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetChannelValueIntegerParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetChannelValueInteger, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetChannelValueInteger"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetChannelValueInteger")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	integerValue, err := ReadSimpleField[BACnetApplicationTagSignedInteger](ctx, "integerValue", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer(), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'integerValue' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetChannelValueInteger"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetChannelValueInteger")
	}

	// Create a partially initialized instance
	_child := &_BACnetChannelValueInteger{
		_BACnetChannelValue: &_BACnetChannelValue{},
		IntegerValue:        integerValue,
	}
	_child._BACnetChannelValue._BACnetChannelValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetChannelValueInteger) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetChannelValueInteger) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueInteger"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetChannelValueInteger")
		}

		// Simple Field (integerValue)
		if pushErr := writeBuffer.PushContext("integerValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for integerValue")
		}
		_integerValueErr := writeBuffer.WriteSerializable(ctx, m.GetIntegerValue())
		if popErr := writeBuffer.PopContext("integerValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for integerValue")
		}
		if _integerValueErr != nil {
			return errors.Wrap(_integerValueErr, "Error serializing 'integerValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetChannelValueInteger"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetChannelValueInteger")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetChannelValueInteger) isBACnetChannelValueInteger() bool {
	return true
}

func (m *_BACnetChannelValueInteger) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
