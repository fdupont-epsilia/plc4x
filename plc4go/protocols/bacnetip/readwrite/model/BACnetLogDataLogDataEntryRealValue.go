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

// BACnetLogDataLogDataEntryRealValue is the corresponding interface of BACnetLogDataLogDataEntryRealValue
type BACnetLogDataLogDataEntryRealValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetLogDataLogDataEntry
	// GetRealValue returns RealValue (property field)
	GetRealValue() BACnetContextTagReal
}

// BACnetLogDataLogDataEntryRealValueExactly can be used when we want exactly this type and not a type which fulfills BACnetLogDataLogDataEntryRealValue.
// This is useful for switch cases.
type BACnetLogDataLogDataEntryRealValueExactly interface {
	BACnetLogDataLogDataEntryRealValue
	isBACnetLogDataLogDataEntryRealValue() bool
}

// _BACnetLogDataLogDataEntryRealValue is the data-structure of this message
type _BACnetLogDataLogDataEntryRealValue struct {
	*_BACnetLogDataLogDataEntry
	RealValue BACnetContextTagReal
}

var _ BACnetLogDataLogDataEntryRealValue = (*_BACnetLogDataLogDataEntryRealValue)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetLogDataLogDataEntryRealValue) InitializeParent(parent BACnetLogDataLogDataEntry, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetLogDataLogDataEntryRealValue) GetParent() BACnetLogDataLogDataEntry {
	return m._BACnetLogDataLogDataEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogDataLogDataEntryRealValue) GetRealValue() BACnetContextTagReal {
	return m.RealValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogDataLogDataEntryRealValue factory function for _BACnetLogDataLogDataEntryRealValue
func NewBACnetLogDataLogDataEntryRealValue(realValue BACnetContextTagReal, peekedTagHeader BACnetTagHeader) *_BACnetLogDataLogDataEntryRealValue {
	_result := &_BACnetLogDataLogDataEntryRealValue{
		RealValue:                  realValue,
		_BACnetLogDataLogDataEntry: NewBACnetLogDataLogDataEntry(peekedTagHeader),
	}
	_result._BACnetLogDataLogDataEntry._BACnetLogDataLogDataEntryChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetLogDataLogDataEntryRealValue(structType any) BACnetLogDataLogDataEntryRealValue {
	if casted, ok := structType.(BACnetLogDataLogDataEntryRealValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogDataLogDataEntryRealValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogDataLogDataEntryRealValue) GetTypeName() string {
	return "BACnetLogDataLogDataEntryRealValue"
}

func (m *_BACnetLogDataLogDataEntryRealValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (realValue)
	lengthInBits += m.RealValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLogDataLogDataEntryRealValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLogDataLogDataEntryRealValueParse(ctx context.Context, theBytes []byte) (BACnetLogDataLogDataEntryRealValue, error) {
	return BACnetLogDataLogDataEntryRealValueParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetLogDataLogDataEntryRealValueParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLogDataLogDataEntryRealValue, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLogDataLogDataEntryRealValue, error) {
		return BACnetLogDataLogDataEntryRealValueParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetLogDataLogDataEntryRealValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLogDataLogDataEntryRealValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogDataLogDataEntryRealValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogDataLogDataEntryRealValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	realValue, err := ReadSimpleField[BACnetContextTagReal](ctx, "realValue", ReadComplex[BACnetContextTagReal](BACnetContextTagParseWithBufferProducer[BACnetContextTagReal]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_REAL)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'realValue' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetLogDataLogDataEntryRealValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogDataLogDataEntryRealValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetLogDataLogDataEntryRealValue{
		_BACnetLogDataLogDataEntry: &_BACnetLogDataLogDataEntry{},
		RealValue:                  realValue,
	}
	_child._BACnetLogDataLogDataEntry._BACnetLogDataLogDataEntryChildRequirements = _child
	return _child, nil
}

func (m *_BACnetLogDataLogDataEntryRealValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogDataLogDataEntryRealValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogDataLogDataEntryRealValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogDataLogDataEntryRealValue")
		}

		if err := WriteSimpleField[BACnetContextTagReal](ctx, "realValue", m.GetRealValue(), WriteComplex[BACnetContextTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'realValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogDataLogDataEntryRealValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogDataLogDataEntryRealValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetLogDataLogDataEntryRealValue) isBACnetLogDataLogDataEntryRealValue() bool {
	return true
}

func (m *_BACnetLogDataLogDataEntryRealValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
