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

// BACnetConstructedDataActiveCOVMultipleSubscriptions is the corresponding interface of BACnetConstructedDataActiveCOVMultipleSubscriptions
type BACnetConstructedDataActiveCOVMultipleSubscriptions interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetActiveCOVMultipleSubscriptions returns ActiveCOVMultipleSubscriptions (property field)
	GetActiveCOVMultipleSubscriptions() []BACnetCOVMultipleSubscription
}

// BACnetConstructedDataActiveCOVMultipleSubscriptionsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataActiveCOVMultipleSubscriptions.
// This is useful for switch cases.
type BACnetConstructedDataActiveCOVMultipleSubscriptionsExactly interface {
	BACnetConstructedDataActiveCOVMultipleSubscriptions
	isBACnetConstructedDataActiveCOVMultipleSubscriptions() bool
}

// _BACnetConstructedDataActiveCOVMultipleSubscriptions is the data-structure of this message
type _BACnetConstructedDataActiveCOVMultipleSubscriptions struct {
	BACnetConstructedDataContract
	ActiveCOVMultipleSubscriptions []BACnetCOVMultipleSubscription
}

var _ BACnetConstructedDataActiveCOVMultipleSubscriptions = (*_BACnetConstructedDataActiveCOVMultipleSubscriptions)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACTIVE_COV_MULTIPLE_SUBSCRIPTIONS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) GetActiveCOVMultipleSubscriptions() []BACnetCOVMultipleSubscription {
	return m.ActiveCOVMultipleSubscriptions
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataActiveCOVMultipleSubscriptions factory function for _BACnetConstructedDataActiveCOVMultipleSubscriptions
func NewBACnetConstructedDataActiveCOVMultipleSubscriptions(activeCOVMultipleSubscriptions []BACnetCOVMultipleSubscription, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataActiveCOVMultipleSubscriptions {
	_result := &_BACnetConstructedDataActiveCOVMultipleSubscriptions{
		BACnetConstructedDataContract:  NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ActiveCOVMultipleSubscriptions: activeCOVMultipleSubscriptions,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataActiveCOVMultipleSubscriptions(structType any) BACnetConstructedDataActiveCOVMultipleSubscriptions {
	if casted, ok := structType.(BACnetConstructedDataActiveCOVMultipleSubscriptions); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataActiveCOVMultipleSubscriptions); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) GetTypeName() string {
	return "BACnetConstructedDataActiveCOVMultipleSubscriptions"
}

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.ActiveCOVMultipleSubscriptions) > 0 {
		for _, element := range m.ActiveCOVMultipleSubscriptions {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (_bACnetConstructedDataActiveCOVMultipleSubscriptions BACnetConstructedDataActiveCOVMultipleSubscriptions, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataActiveCOVMultipleSubscriptions"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataActiveCOVMultipleSubscriptions")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	activeCOVMultipleSubscriptions, err := ReadTerminatedArrayField[BACnetCOVMultipleSubscription](ctx, "activeCOVMultipleSubscriptions", ReadComplex[BACnetCOVMultipleSubscription](BACnetCOVMultipleSubscriptionParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'activeCOVMultipleSubscriptions' field"))
	}
	m.ActiveCOVMultipleSubscriptions = activeCOVMultipleSubscriptions

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataActiveCOVMultipleSubscriptions"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataActiveCOVMultipleSubscriptions")
	}

	return m, nil
}

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataActiveCOVMultipleSubscriptions"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataActiveCOVMultipleSubscriptions")
		}

		if err := WriteComplexTypeArrayField(ctx, "activeCOVMultipleSubscriptions", m.GetActiveCOVMultipleSubscriptions(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'activeCOVMultipleSubscriptions' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataActiveCOVMultipleSubscriptions"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataActiveCOVMultipleSubscriptions")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) isBACnetConstructedDataActiveCOVMultipleSubscriptions() bool {
	return true
}

func (m *_BACnetConstructedDataActiveCOVMultipleSubscriptions) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
