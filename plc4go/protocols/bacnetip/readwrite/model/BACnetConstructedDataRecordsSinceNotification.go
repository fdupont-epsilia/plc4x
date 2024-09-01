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

// BACnetConstructedDataRecordsSinceNotification is the corresponding interface of BACnetConstructedDataRecordsSinceNotification
type BACnetConstructedDataRecordsSinceNotification interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetRecordsSinceNotifications returns RecordsSinceNotifications (property field)
	GetRecordsSinceNotifications() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataRecordsSinceNotificationExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataRecordsSinceNotification.
// This is useful for switch cases.
type BACnetConstructedDataRecordsSinceNotificationExactly interface {
	BACnetConstructedDataRecordsSinceNotification
	isBACnetConstructedDataRecordsSinceNotification() bool
}

// _BACnetConstructedDataRecordsSinceNotification is the data-structure of this message
type _BACnetConstructedDataRecordsSinceNotification struct {
	BACnetConstructedDataContract
	RecordsSinceNotifications BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataRecordsSinceNotification = (*_BACnetConstructedDataRecordsSinceNotification)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataRecordsSinceNotification) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataRecordsSinceNotification) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RECORDS_SINCE_NOTIFICATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataRecordsSinceNotification) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataRecordsSinceNotification) GetRecordsSinceNotifications() BACnetApplicationTagUnsignedInteger {
	return m.RecordsSinceNotifications
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataRecordsSinceNotification) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetRecordsSinceNotifications())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataRecordsSinceNotification factory function for _BACnetConstructedDataRecordsSinceNotification
func NewBACnetConstructedDataRecordsSinceNotification(recordsSinceNotifications BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataRecordsSinceNotification {
	_result := &_BACnetConstructedDataRecordsSinceNotification{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		RecordsSinceNotifications:     recordsSinceNotifications,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataRecordsSinceNotification(structType any) BACnetConstructedDataRecordsSinceNotification {
	if casted, ok := structType.(BACnetConstructedDataRecordsSinceNotification); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataRecordsSinceNotification); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataRecordsSinceNotification) GetTypeName() string {
	return "BACnetConstructedDataRecordsSinceNotification"
}

func (m *_BACnetConstructedDataRecordsSinceNotification) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (recordsSinceNotifications)
	lengthInBits += m.RecordsSinceNotifications.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataRecordsSinceNotification) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataRecordsSinceNotification) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (_bACnetConstructedDataRecordsSinceNotification BACnetConstructedDataRecordsSinceNotification, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataRecordsSinceNotification"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataRecordsSinceNotification")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	recordsSinceNotifications, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "recordsSinceNotifications", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'recordsSinceNotifications' field"))
	}
	m.RecordsSinceNotifications = recordsSinceNotifications

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), recordsSinceNotifications)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataRecordsSinceNotification"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataRecordsSinceNotification")
	}

	return m, nil
}

func (m *_BACnetConstructedDataRecordsSinceNotification) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataRecordsSinceNotification) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataRecordsSinceNotification"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataRecordsSinceNotification")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "recordsSinceNotifications", m.GetRecordsSinceNotifications(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'recordsSinceNotifications' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataRecordsSinceNotification"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataRecordsSinceNotification")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataRecordsSinceNotification) isBACnetConstructedDataRecordsSinceNotification() bool {
	return true
}

func (m *_BACnetConstructedDataRecordsSinceNotification) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
