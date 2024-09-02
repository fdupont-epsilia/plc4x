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

// BACnetChannelValueEnumerated is the corresponding interface of BACnetChannelValueEnumerated
type BACnetChannelValueEnumerated interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetChannelValue
	// GetEnumeratedValue returns EnumeratedValue (property field)
	GetEnumeratedValue() BACnetApplicationTagEnumerated
	// IsBACnetChannelValueEnumerated is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetChannelValueEnumerated()
}

// _BACnetChannelValueEnumerated is the data-structure of this message
type _BACnetChannelValueEnumerated struct {
	BACnetChannelValueContract
	EnumeratedValue BACnetApplicationTagEnumerated
}

var _ BACnetChannelValueEnumerated = (*_BACnetChannelValueEnumerated)(nil)
var _ BACnetChannelValueRequirements = (*_BACnetChannelValueEnumerated)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetChannelValueEnumerated) GetParent() BACnetChannelValueContract {
	return m.BACnetChannelValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetChannelValueEnumerated) GetEnumeratedValue() BACnetApplicationTagEnumerated {
	return m.EnumeratedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetChannelValueEnumerated factory function for _BACnetChannelValueEnumerated
func NewBACnetChannelValueEnumerated(enumeratedValue BACnetApplicationTagEnumerated, peekedTagHeader BACnetTagHeader) *_BACnetChannelValueEnumerated {
	_result := &_BACnetChannelValueEnumerated{
		BACnetChannelValueContract: NewBACnetChannelValue(peekedTagHeader),
		EnumeratedValue:            enumeratedValue,
	}
	_result.BACnetChannelValueContract.(*_BACnetChannelValue)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetChannelValueEnumerated(structType any) BACnetChannelValueEnumerated {
	if casted, ok := structType.(BACnetChannelValueEnumerated); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetChannelValueEnumerated); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetChannelValueEnumerated) GetTypeName() string {
	return "BACnetChannelValueEnumerated"
}

func (m *_BACnetChannelValueEnumerated) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetChannelValueContract.(*_BACnetChannelValue).getLengthInBits(ctx))

	// Simple field (enumeratedValue)
	lengthInBits += m.EnumeratedValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetChannelValueEnumerated) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetChannelValueEnumerated) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetChannelValue) (__bACnetChannelValueEnumerated BACnetChannelValueEnumerated, err error) {
	m.BACnetChannelValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValueEnumerated"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetChannelValueEnumerated")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	enumeratedValue, err := ReadSimpleField[BACnetApplicationTagEnumerated](ctx, "enumeratedValue", ReadComplex[BACnetApplicationTagEnumerated](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagEnumerated](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'enumeratedValue' field"))
	}
	m.EnumeratedValue = enumeratedValue

	if closeErr := readBuffer.CloseContext("BACnetChannelValueEnumerated"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetChannelValueEnumerated")
	}

	return m, nil
}

func (m *_BACnetChannelValueEnumerated) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetChannelValueEnumerated) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueEnumerated"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetChannelValueEnumerated")
		}

		if err := WriteSimpleField[BACnetApplicationTagEnumerated](ctx, "enumeratedValue", m.GetEnumeratedValue(), WriteComplex[BACnetApplicationTagEnumerated](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'enumeratedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetChannelValueEnumerated"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetChannelValueEnumerated")
		}
		return nil
	}
	return m.BACnetChannelValueContract.(*_BACnetChannelValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetChannelValueEnumerated) IsBACnetChannelValueEnumerated() {}

func (m *_BACnetChannelValueEnumerated) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
