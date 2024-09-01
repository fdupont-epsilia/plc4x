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

// BACnetConstructedDataAccessTransactionEvents is the corresponding interface of BACnetConstructedDataAccessTransactionEvents
type BACnetConstructedDataAccessTransactionEvents interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAccessTransactionEvents returns AccessTransactionEvents (property field)
	GetAccessTransactionEvents() []BACnetAccessEventTagged
}

// BACnetConstructedDataAccessTransactionEventsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAccessTransactionEvents.
// This is useful for switch cases.
type BACnetConstructedDataAccessTransactionEventsExactly interface {
	BACnetConstructedDataAccessTransactionEvents
	isBACnetConstructedDataAccessTransactionEvents() bool
}

// _BACnetConstructedDataAccessTransactionEvents is the data-structure of this message
type _BACnetConstructedDataAccessTransactionEvents struct {
	*_BACnetConstructedData
	AccessTransactionEvents []BACnetAccessEventTagged
}

var _ BACnetConstructedDataAccessTransactionEvents = (*_BACnetConstructedDataAccessTransactionEvents)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccessTransactionEvents) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAccessTransactionEvents) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACCESS_TRANSACTION_EVENTS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccessTransactionEvents) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAccessTransactionEvents) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccessTransactionEvents) GetAccessTransactionEvents() []BACnetAccessEventTagged {
	return m.AccessTransactionEvents
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAccessTransactionEvents factory function for _BACnetConstructedDataAccessTransactionEvents
func NewBACnetConstructedDataAccessTransactionEvents(accessTransactionEvents []BACnetAccessEventTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccessTransactionEvents {
	_result := &_BACnetConstructedDataAccessTransactionEvents{
		AccessTransactionEvents: accessTransactionEvents,
		_BACnetConstructedData:  NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccessTransactionEvents(structType any) BACnetConstructedDataAccessTransactionEvents {
	if casted, ok := structType.(BACnetConstructedDataAccessTransactionEvents); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccessTransactionEvents); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccessTransactionEvents) GetTypeName() string {
	return "BACnetConstructedDataAccessTransactionEvents"
}

func (m *_BACnetConstructedDataAccessTransactionEvents) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.AccessTransactionEvents) > 0 {
		for _, element := range m.AccessTransactionEvents {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataAccessTransactionEvents) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataAccessTransactionEventsParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAccessTransactionEvents, error) {
	return BACnetConstructedDataAccessTransactionEventsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataAccessTransactionEventsParseWithBufferProducer(tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataAccessTransactionEvents, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataAccessTransactionEvents, error) {
		return BACnetConstructedDataAccessTransactionEventsParseWithBuffer(ctx, readBuffer, tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
	}
}

func BACnetConstructedDataAccessTransactionEventsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAccessTransactionEvents, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccessTransactionEvents"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccessTransactionEvents")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	accessTransactionEvents, err := ReadTerminatedArrayField[BACnetAccessEventTagged](ctx, "accessTransactionEvents", ReadComplex[BACnetAccessEventTagged](BACnetAccessEventTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'accessTransactionEvents' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccessTransactionEvents"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccessTransactionEvents")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAccessTransactionEvents{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		AccessTransactionEvents: accessTransactionEvents,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAccessTransactionEvents) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccessTransactionEvents) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccessTransactionEvents"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccessTransactionEvents")
		}

		if err := WriteComplexTypeArrayField(ctx, "accessTransactionEvents", m.GetAccessTransactionEvents(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'accessTransactionEvents' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccessTransactionEvents"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccessTransactionEvents")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccessTransactionEvents) isBACnetConstructedDataAccessTransactionEvents() bool {
	return true
}

func (m *_BACnetConstructedDataAccessTransactionEvents) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
