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

// BACnetConstructedDataStructuredViewAll is the corresponding interface of BACnetConstructedDataStructuredViewAll
type BACnetConstructedDataStructuredViewAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// IsBACnetConstructedDataStructuredViewAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataStructuredViewAll()
}

// _BACnetConstructedDataStructuredViewAll is the data-structure of this message
type _BACnetConstructedDataStructuredViewAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataStructuredViewAll = (*_BACnetConstructedDataStructuredViewAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataStructuredViewAll)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataStructuredViewAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_STRUCTURED_VIEW
}

func (m *_BACnetConstructedDataStructuredViewAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataStructuredViewAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// NewBACnetConstructedDataStructuredViewAll factory function for _BACnetConstructedDataStructuredViewAll
func NewBACnetConstructedDataStructuredViewAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataStructuredViewAll {
	_result := &_BACnetConstructedDataStructuredViewAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataStructuredViewAll(structType any) BACnetConstructedDataStructuredViewAll {
	if casted, ok := structType.(BACnetConstructedDataStructuredViewAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataStructuredViewAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataStructuredViewAll) GetTypeName() string {
	return "BACnetConstructedDataStructuredViewAll"
}

func (m *_BACnetConstructedDataStructuredViewAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataStructuredViewAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataStructuredViewAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataStructuredViewAll BACnetConstructedDataStructuredViewAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataStructuredViewAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataStructuredViewAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataStructuredViewAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataStructuredViewAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataStructuredViewAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataStructuredViewAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataStructuredViewAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataStructuredViewAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataStructuredViewAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataStructuredViewAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataStructuredViewAll) IsBACnetConstructedDataStructuredViewAll() {}

func (m *_BACnetConstructedDataStructuredViewAll) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
