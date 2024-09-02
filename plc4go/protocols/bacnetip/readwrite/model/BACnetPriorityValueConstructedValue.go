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

// BACnetPriorityValueConstructedValue is the corresponding interface of BACnetPriorityValueConstructedValue
type BACnetPriorityValueConstructedValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPriorityValue
	// GetConstructedValue returns ConstructedValue (property field)
	GetConstructedValue() BACnetConstructedData
	// IsBACnetPriorityValueConstructedValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPriorityValueConstructedValue()
}

// _BACnetPriorityValueConstructedValue is the data-structure of this message
type _BACnetPriorityValueConstructedValue struct {
	BACnetPriorityValueContract
	ConstructedValue BACnetConstructedData
}

var _ BACnetPriorityValueConstructedValue = (*_BACnetPriorityValueConstructedValue)(nil)
var _ BACnetPriorityValueRequirements = (*_BACnetPriorityValueConstructedValue)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPriorityValueConstructedValue) GetParent() BACnetPriorityValueContract {
	return m.BACnetPriorityValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPriorityValueConstructedValue) GetConstructedValue() BACnetConstructedData {
	return m.ConstructedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPriorityValueConstructedValue factory function for _BACnetPriorityValueConstructedValue
func NewBACnetPriorityValueConstructedValue(constructedValue BACnetConstructedData, peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType) *_BACnetPriorityValueConstructedValue {
	_result := &_BACnetPriorityValueConstructedValue{
		BACnetPriorityValueContract: NewBACnetPriorityValue(peekedTagHeader, objectTypeArgument),
		ConstructedValue:            constructedValue,
	}
	_result.BACnetPriorityValueContract.(*_BACnetPriorityValue)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPriorityValueConstructedValue(structType any) BACnetPriorityValueConstructedValue {
	if casted, ok := structType.(BACnetPriorityValueConstructedValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPriorityValueConstructedValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPriorityValueConstructedValue) GetTypeName() string {
	return "BACnetPriorityValueConstructedValue"
}

func (m *_BACnetPriorityValueConstructedValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPriorityValueContract.(*_BACnetPriorityValue).getLengthInBits(ctx))

	// Simple field (constructedValue)
	lengthInBits += m.ConstructedValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPriorityValueConstructedValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPriorityValueConstructedValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPriorityValue, objectTypeArgument BACnetObjectType) (__bACnetPriorityValueConstructedValue BACnetPriorityValueConstructedValue, err error) {
	m.BACnetPriorityValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityValueConstructedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValueConstructedValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	constructedValue, err := ReadSimpleField[BACnetConstructedData](ctx, "constructedValue", ReadComplex[BACnetConstructedData](BACnetConstructedDataParseWithBufferProducer[BACnetConstructedData]((uint8)(uint8(0)), (BACnetObjectType)(objectTypeArgument), (BACnetPropertyIdentifier)(BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE), (BACnetTagPayloadUnsignedInteger)(nil)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'constructedValue' field"))
	}
	m.ConstructedValue = constructedValue

	if closeErr := readBuffer.CloseContext("BACnetPriorityValueConstructedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValueConstructedValue")
	}

	return m, nil
}

func (m *_BACnetPriorityValueConstructedValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPriorityValueConstructedValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPriorityValueConstructedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValueConstructedValue")
		}

		if err := WriteSimpleField[BACnetConstructedData](ctx, "constructedValue", m.GetConstructedValue(), WriteComplex[BACnetConstructedData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'constructedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPriorityValueConstructedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPriorityValueConstructedValue")
		}
		return nil
	}
	return m.BACnetPriorityValueContract.(*_BACnetPriorityValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPriorityValueConstructedValue) IsBACnetPriorityValueConstructedValue() {}

func (m *_BACnetPriorityValueConstructedValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
