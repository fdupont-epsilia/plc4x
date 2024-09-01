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

// BACnetNotificationParametersBufferReady is the corresponding interface of BACnetNotificationParametersBufferReady
type BACnetNotificationParametersBufferReady interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() BACnetOpeningTag
	// GetBufferProperty returns BufferProperty (property field)
	GetBufferProperty() BACnetDeviceObjectPropertyReferenceEnclosed
	// GetPreviousNotification returns PreviousNotification (property field)
	GetPreviousNotification() BACnetContextTagUnsignedInteger
	// GetCurrentNotification returns CurrentNotification (property field)
	GetCurrentNotification() BACnetContextTagUnsignedInteger
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() BACnetClosingTag
}

// BACnetNotificationParametersBufferReadyExactly can be used when we want exactly this type and not a type which fulfills BACnetNotificationParametersBufferReady.
// This is useful for switch cases.
type BACnetNotificationParametersBufferReadyExactly interface {
	BACnetNotificationParametersBufferReady
	isBACnetNotificationParametersBufferReady() bool
}

// _BACnetNotificationParametersBufferReady is the data-structure of this message
type _BACnetNotificationParametersBufferReady struct {
	BACnetNotificationParametersContract
	InnerOpeningTag      BACnetOpeningTag
	BufferProperty       BACnetDeviceObjectPropertyReferenceEnclosed
	PreviousNotification BACnetContextTagUnsignedInteger
	CurrentNotification  BACnetContextTagUnsignedInteger
	InnerClosingTag      BACnetClosingTag
}

var _ BACnetNotificationParametersBufferReady = (*_BACnetNotificationParametersBufferReady)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetNotificationParametersBufferReady) GetParent() BACnetNotificationParametersContract {
	return m.BACnetNotificationParametersContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersBufferReady) GetInnerOpeningTag() BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *_BACnetNotificationParametersBufferReady) GetBufferProperty() BACnetDeviceObjectPropertyReferenceEnclosed {
	return m.BufferProperty
}

func (m *_BACnetNotificationParametersBufferReady) GetPreviousNotification() BACnetContextTagUnsignedInteger {
	return m.PreviousNotification
}

func (m *_BACnetNotificationParametersBufferReady) GetCurrentNotification() BACnetContextTagUnsignedInteger {
	return m.CurrentNotification
}

func (m *_BACnetNotificationParametersBufferReady) GetInnerClosingTag() BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersBufferReady factory function for _BACnetNotificationParametersBufferReady
func NewBACnetNotificationParametersBufferReady(innerOpeningTag BACnetOpeningTag, bufferProperty BACnetDeviceObjectPropertyReferenceEnclosed, previousNotification BACnetContextTagUnsignedInteger, currentNotification BACnetContextTagUnsignedInteger, innerClosingTag BACnetClosingTag, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *_BACnetNotificationParametersBufferReady {
	_result := &_BACnetNotificationParametersBufferReady{
		BACnetNotificationParametersContract: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument),
		InnerOpeningTag:                      innerOpeningTag,
		BufferProperty:                       bufferProperty,
		PreviousNotification:                 previousNotification,
		CurrentNotification:                  currentNotification,
		InnerClosingTag:                      innerClosingTag,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersBufferReady(structType any) BACnetNotificationParametersBufferReady {
	if casted, ok := structType.(BACnetNotificationParametersBufferReady); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersBufferReady); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersBufferReady) GetTypeName() string {
	return "BACnetNotificationParametersBufferReady"
}

func (m *_BACnetNotificationParametersBufferReady) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).getLengthInBits(ctx))

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits(ctx)

	// Simple field (bufferProperty)
	lengthInBits += m.BufferProperty.GetLengthInBits(ctx)

	// Simple field (previousNotification)
	lengthInBits += m.PreviousNotification.GetLengthInBits(ctx)

	// Simple field (currentNotification)
	lengthInBits += m.CurrentNotification.GetLengthInBits(ctx)

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersBufferReady) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetNotificationParametersBufferReady) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetNotificationParameters, peekedTagNumber uint8, tagNumber uint8, objectTypeArgument BACnetObjectType) (_bACnetNotificationParametersBufferReady BACnetNotificationParametersBufferReady, err error) {
	m.BACnetNotificationParametersContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersBufferReady"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersBufferReady")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	innerOpeningTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(peekedTagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerOpeningTag' field"))
	}
	m.InnerOpeningTag = innerOpeningTag

	bufferProperty, err := ReadSimpleField[BACnetDeviceObjectPropertyReferenceEnclosed](ctx, "bufferProperty", ReadComplex[BACnetDeviceObjectPropertyReferenceEnclosed](BACnetDeviceObjectPropertyReferenceEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bufferProperty' field"))
	}
	m.BufferProperty = bufferProperty

	previousNotification, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "previousNotification", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'previousNotification' field"))
	}
	m.PreviousNotification = previousNotification

	currentNotification, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "currentNotification", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'currentNotification' field"))
	}
	m.CurrentNotification = currentNotification

	innerClosingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "innerClosingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(peekedTagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerClosingTag' field"))
	}
	m.InnerClosingTag = innerClosingTag

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersBufferReady"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersBufferReady")
	}

	return m, nil
}

func (m *_BACnetNotificationParametersBufferReady) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersBufferReady) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersBufferReady"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersBufferReady")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", m.GetInnerOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'innerOpeningTag' field")
		}

		if err := WriteSimpleField[BACnetDeviceObjectPropertyReferenceEnclosed](ctx, "bufferProperty", m.GetBufferProperty(), WriteComplex[BACnetDeviceObjectPropertyReferenceEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'bufferProperty' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "previousNotification", m.GetPreviousNotification(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'previousNotification' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "currentNotification", m.GetCurrentNotification(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'currentNotification' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "innerClosingTag", m.GetInnerClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersBufferReady"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersBufferReady")
		}
		return nil
	}
	return m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersBufferReady) isBACnetNotificationParametersBufferReady() bool {
	return true
}

func (m *_BACnetNotificationParametersBufferReady) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
