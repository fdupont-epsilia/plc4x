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

// BACnetTimerStateChangeValueNull is the corresponding interface of BACnetTimerStateChangeValueNull
type BACnetTimerStateChangeValueNull interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetTimerStateChangeValue
	// GetNullValue returns NullValue (property field)
	GetNullValue() BACnetApplicationTagNull
}

// BACnetTimerStateChangeValueNullExactly can be used when we want exactly this type and not a type which fulfills BACnetTimerStateChangeValueNull.
// This is useful for switch cases.
type BACnetTimerStateChangeValueNullExactly interface {
	BACnetTimerStateChangeValueNull
	isBACnetTimerStateChangeValueNull() bool
}

// _BACnetTimerStateChangeValueNull is the data-structure of this message
type _BACnetTimerStateChangeValueNull struct {
	*_BACnetTimerStateChangeValue
	NullValue BACnetApplicationTagNull
}

var _ BACnetTimerStateChangeValueNull = (*_BACnetTimerStateChangeValueNull)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetTimerStateChangeValueNull) InitializeParent(parent BACnetTimerStateChangeValue, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetTimerStateChangeValueNull) GetParent() BACnetTimerStateChangeValue {
	return m._BACnetTimerStateChangeValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimerStateChangeValueNull) GetNullValue() BACnetApplicationTagNull {
	return m.NullValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTimerStateChangeValueNull factory function for _BACnetTimerStateChangeValueNull
func NewBACnetTimerStateChangeValueNull(nullValue BACnetApplicationTagNull, peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType) *_BACnetTimerStateChangeValueNull {
	_result := &_BACnetTimerStateChangeValueNull{
		NullValue:                    nullValue,
		_BACnetTimerStateChangeValue: NewBACnetTimerStateChangeValue(peekedTagHeader, objectTypeArgument),
	}
	_result._BACnetTimerStateChangeValue._BACnetTimerStateChangeValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetTimerStateChangeValueNull(structType any) BACnetTimerStateChangeValueNull {
	if casted, ok := structType.(BACnetTimerStateChangeValueNull); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimerStateChangeValueNull); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimerStateChangeValueNull) GetTypeName() string {
	return "BACnetTimerStateChangeValueNull"
}

func (m *_BACnetTimerStateChangeValueNull) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (nullValue)
	lengthInBits += m.NullValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetTimerStateChangeValueNull) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTimerStateChangeValueNullParse(ctx context.Context, theBytes []byte, objectTypeArgument BACnetObjectType) (BACnetTimerStateChangeValueNull, error) {
	return BACnetTimerStateChangeValueNullParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), objectTypeArgument)
}

func BACnetTimerStateChangeValueNullParseWithBufferProducer(objectTypeArgument BACnetObjectType) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTimerStateChangeValueNull, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTimerStateChangeValueNull, error) {
		return BACnetTimerStateChangeValueNullParseWithBuffer(ctx, readBuffer, objectTypeArgument)
	}
}

func BACnetTimerStateChangeValueNullParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (BACnetTimerStateChangeValueNull, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimerStateChangeValueNull"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimerStateChangeValueNull")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	nullValue, err := ReadSimpleField[BACnetApplicationTagNull](ctx, "nullValue", ReadComplex[BACnetApplicationTagNull](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagNull](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nullValue' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetTimerStateChangeValueNull"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimerStateChangeValueNull")
	}

	// Create a partially initialized instance
	_child := &_BACnetTimerStateChangeValueNull{
		_BACnetTimerStateChangeValue: &_BACnetTimerStateChangeValue{
			ObjectTypeArgument: objectTypeArgument,
		},
		NullValue: nullValue,
	}
	_child._BACnetTimerStateChangeValue._BACnetTimerStateChangeValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetTimerStateChangeValueNull) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimerStateChangeValueNull) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimerStateChangeValueNull"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimerStateChangeValueNull")
		}

		if err := WriteSimpleField[BACnetApplicationTagNull](ctx, "nullValue", m.GetNullValue(), WriteComplex[BACnetApplicationTagNull](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'nullValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimerStateChangeValueNull"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimerStateChangeValueNull")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetTimerStateChangeValueNull) isBACnetTimerStateChangeValueNull() bool {
	return true
}

func (m *_BACnetTimerStateChangeValueNull) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
