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

// BACnetFaultParameterFaultExtendedParametersEntryEnumerated is the corresponding interface of BACnetFaultParameterFaultExtendedParametersEntryEnumerated
type BACnetFaultParameterFaultExtendedParametersEntryEnumerated interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetFaultParameterFaultExtendedParametersEntry
	// GetEnumeratedValue returns EnumeratedValue (property field)
	GetEnumeratedValue() BACnetApplicationTagEnumerated
	// IsBACnetFaultParameterFaultExtendedParametersEntryEnumerated is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetFaultParameterFaultExtendedParametersEntryEnumerated()
}

// _BACnetFaultParameterFaultExtendedParametersEntryEnumerated is the data-structure of this message
type _BACnetFaultParameterFaultExtendedParametersEntryEnumerated struct {
	BACnetFaultParameterFaultExtendedParametersEntryContract
	EnumeratedValue BACnetApplicationTagEnumerated
}

var _ BACnetFaultParameterFaultExtendedParametersEntryEnumerated = (*_BACnetFaultParameterFaultExtendedParametersEntryEnumerated)(nil)
var _ BACnetFaultParameterFaultExtendedParametersEntryRequirements = (*_BACnetFaultParameterFaultExtendedParametersEntryEnumerated)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetFaultParameterFaultExtendedParametersEntryEnumerated) GetParent() BACnetFaultParameterFaultExtendedParametersEntryContract {
	return m.BACnetFaultParameterFaultExtendedParametersEntryContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultExtendedParametersEntryEnumerated) GetEnumeratedValue() BACnetApplicationTagEnumerated {
	return m.EnumeratedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultExtendedParametersEntryEnumerated factory function for _BACnetFaultParameterFaultExtendedParametersEntryEnumerated
func NewBACnetFaultParameterFaultExtendedParametersEntryEnumerated(enumeratedValue BACnetApplicationTagEnumerated, peekedTagHeader BACnetTagHeader) *_BACnetFaultParameterFaultExtendedParametersEntryEnumerated {
	_result := &_BACnetFaultParameterFaultExtendedParametersEntryEnumerated{
		BACnetFaultParameterFaultExtendedParametersEntryContract: NewBACnetFaultParameterFaultExtendedParametersEntry(peekedTagHeader),
		EnumeratedValue: enumeratedValue,
	}
	_result.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultExtendedParametersEntryEnumerated(structType any) BACnetFaultParameterFaultExtendedParametersEntryEnumerated {
	if casted, ok := structType.(BACnetFaultParameterFaultExtendedParametersEntryEnumerated); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtendedParametersEntryEnumerated); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryEnumerated) GetTypeName() string {
	return "BACnetFaultParameterFaultExtendedParametersEntryEnumerated"
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryEnumerated) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry).getLengthInBits(ctx))

	// Simple field (enumeratedValue)
	lengthInBits += m.EnumeratedValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryEnumerated) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryEnumerated) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetFaultParameterFaultExtendedParametersEntry) (__bACnetFaultParameterFaultExtendedParametersEntryEnumerated BACnetFaultParameterFaultExtendedParametersEntryEnumerated, err error) {
	m.BACnetFaultParameterFaultExtendedParametersEntryContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultExtendedParametersEntryEnumerated"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultExtendedParametersEntryEnumerated")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	enumeratedValue, err := ReadSimpleField[BACnetApplicationTagEnumerated](ctx, "enumeratedValue", ReadComplex[BACnetApplicationTagEnumerated](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagEnumerated](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'enumeratedValue' field"))
	}
	m.EnumeratedValue = enumeratedValue

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultExtendedParametersEntryEnumerated"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultExtendedParametersEntryEnumerated")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryEnumerated) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryEnumerated) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultExtendedParametersEntryEnumerated"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultExtendedParametersEntryEnumerated")
		}

		if err := WriteSimpleField[BACnetApplicationTagEnumerated](ctx, "enumeratedValue", m.GetEnumeratedValue(), WriteComplex[BACnetApplicationTagEnumerated](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'enumeratedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultExtendedParametersEntryEnumerated"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultExtendedParametersEntryEnumerated")
		}
		return nil
	}
	return m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryEnumerated) IsBACnetFaultParameterFaultExtendedParametersEntryEnumerated() {
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryEnumerated) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
