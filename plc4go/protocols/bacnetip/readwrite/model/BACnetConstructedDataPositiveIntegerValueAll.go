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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataPositiveIntegerValueAll is the corresponding interface of BACnetConstructedDataPositiveIntegerValueAll
type BACnetConstructedDataPositiveIntegerValueAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
}

// BACnetConstructedDataPositiveIntegerValueAllExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataPositiveIntegerValueAll.
// This is useful for switch cases.
type BACnetConstructedDataPositiveIntegerValueAllExactly interface {
	BACnetConstructedDataPositiveIntegerValueAll
	isBACnetConstructedDataPositiveIntegerValueAll() bool
}

// _BACnetConstructedDataPositiveIntegerValueAll is the data-structure of this message
type _BACnetConstructedDataPositiveIntegerValueAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataPositiveIntegerValueAll = (*_BACnetConstructedDataPositiveIntegerValueAll)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataPositiveIntegerValueAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_POSITIVE_INTEGER_VALUE
}

func (m *_BACnetConstructedDataPositiveIntegerValueAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataPositiveIntegerValueAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// NewBACnetConstructedDataPositiveIntegerValueAll factory function for _BACnetConstructedDataPositiveIntegerValueAll
func NewBACnetConstructedDataPositiveIntegerValueAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataPositiveIntegerValueAll {
	_result := &_BACnetConstructedDataPositiveIntegerValueAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataPositiveIntegerValueAll(structType any) BACnetConstructedDataPositiveIntegerValueAll {
	if casted, ok := structType.(BACnetConstructedDataPositiveIntegerValueAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPositiveIntegerValueAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataPositiveIntegerValueAll) GetTypeName() string {
	return "BACnetConstructedDataPositiveIntegerValueAll"
}

func (m *_BACnetConstructedDataPositiveIntegerValueAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataPositiveIntegerValueAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataPositiveIntegerValueAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (_bACnetConstructedDataPositiveIntegerValueAll BACnetConstructedDataPositiveIntegerValueAll, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPositiveIntegerValueAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPositiveIntegerValueAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPositiveIntegerValueAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPositiveIntegerValueAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataPositiveIntegerValueAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataPositiveIntegerValueAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPositiveIntegerValueAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPositiveIntegerValueAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPositiveIntegerValueAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPositiveIntegerValueAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataPositiveIntegerValueAll) isBACnetConstructedDataPositiveIntegerValueAll() bool {
	return true
}

func (m *_BACnetConstructedDataPositiveIntegerValueAll) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
