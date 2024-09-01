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

// BACnetOptionalREALValue is the corresponding interface of BACnetOptionalREALValue
type BACnetOptionalREALValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetOptionalREAL
	// GetRealValue returns RealValue (property field)
	GetRealValue() BACnetApplicationTagReal
}

// BACnetOptionalREALValueExactly can be used when we want exactly this type and not a type which fulfills BACnetOptionalREALValue.
// This is useful for switch cases.
type BACnetOptionalREALValueExactly interface {
	BACnetOptionalREALValue
	isBACnetOptionalREALValue() bool
}

// _BACnetOptionalREALValue is the data-structure of this message
type _BACnetOptionalREALValue struct {
	BACnetOptionalREALContract
	RealValue BACnetApplicationTagReal
}

var _ BACnetOptionalREALValue = (*_BACnetOptionalREALValue)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetOptionalREALValue) GetParent() BACnetOptionalREALContract {
	return m.BACnetOptionalREALContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetOptionalREALValue) GetRealValue() BACnetApplicationTagReal {
	return m.RealValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetOptionalREALValue factory function for _BACnetOptionalREALValue
func NewBACnetOptionalREALValue(realValue BACnetApplicationTagReal, peekedTagHeader BACnetTagHeader) *_BACnetOptionalREALValue {
	_result := &_BACnetOptionalREALValue{
		BACnetOptionalREALContract: NewBACnetOptionalREAL(peekedTagHeader),
		RealValue:                  realValue,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetOptionalREALValue(structType any) BACnetOptionalREALValue {
	if casted, ok := structType.(BACnetOptionalREALValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetOptionalREALValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetOptionalREALValue) GetTypeName() string {
	return "BACnetOptionalREALValue"
}

func (m *_BACnetOptionalREALValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetOptionalREALContract.(*_BACnetOptionalREAL).getLengthInBits(ctx))

	// Simple field (realValue)
	lengthInBits += m.RealValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetOptionalREALValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetOptionalREALValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetOptionalREAL) (_bACnetOptionalREALValue BACnetOptionalREALValue, err error) {
	m.BACnetOptionalREALContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetOptionalREALValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetOptionalREALValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	realValue, err := ReadSimpleField[BACnetApplicationTagReal](ctx, "realValue", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'realValue' field"))
	}
	m.RealValue = realValue

	if closeErr := readBuffer.CloseContext("BACnetOptionalREALValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetOptionalREALValue")
	}

	return m, nil
}

func (m *_BACnetOptionalREALValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetOptionalREALValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetOptionalREALValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetOptionalREALValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagReal](ctx, "realValue", m.GetRealValue(), WriteComplex[BACnetApplicationTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'realValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetOptionalREALValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetOptionalREALValue")
		}
		return nil
	}
	return m.BACnetOptionalREALContract.(*_BACnetOptionalREAL).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetOptionalREALValue) isBACnetOptionalREALValue() bool {
	return true
}

func (m *_BACnetOptionalREALValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
