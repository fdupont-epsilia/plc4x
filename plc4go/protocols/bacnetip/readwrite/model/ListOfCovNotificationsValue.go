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

// ListOfCovNotificationsValue is the corresponding interface of ListOfCovNotificationsValue
type ListOfCovNotificationsValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() BACnetPropertyIdentifierTagged
	// GetArrayIndex returns ArrayIndex (property field)
	GetArrayIndex() BACnetContextTagUnsignedInteger
	// GetPropertyValue returns PropertyValue (property field)
	GetPropertyValue() BACnetConstructedData
	// GetTimeOfChange returns TimeOfChange (property field)
	GetTimeOfChange() BACnetContextTagTime
}

// ListOfCovNotificationsValueExactly can be used when we want exactly this type and not a type which fulfills ListOfCovNotificationsValue.
// This is useful for switch cases.
type ListOfCovNotificationsValueExactly interface {
	ListOfCovNotificationsValue
	isListOfCovNotificationsValue() bool
}

// _ListOfCovNotificationsValue is the data-structure of this message
type _ListOfCovNotificationsValue struct {
	PropertyIdentifier BACnetPropertyIdentifierTagged
	ArrayIndex         BACnetContextTagUnsignedInteger
	PropertyValue      BACnetConstructedData
	TimeOfChange       BACnetContextTagTime

	// Arguments.
	ObjectTypeArgument BACnetObjectType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ListOfCovNotificationsValue) GetPropertyIdentifier() BACnetPropertyIdentifierTagged {
	return m.PropertyIdentifier
}

func (m *_ListOfCovNotificationsValue) GetArrayIndex() BACnetContextTagUnsignedInteger {
	return m.ArrayIndex
}

func (m *_ListOfCovNotificationsValue) GetPropertyValue() BACnetConstructedData {
	return m.PropertyValue
}

func (m *_ListOfCovNotificationsValue) GetTimeOfChange() BACnetContextTagTime {
	return m.TimeOfChange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewListOfCovNotificationsValue factory function for _ListOfCovNotificationsValue
func NewListOfCovNotificationsValue(propertyIdentifier BACnetPropertyIdentifierTagged, arrayIndex BACnetContextTagUnsignedInteger, propertyValue BACnetConstructedData, timeOfChange BACnetContextTagTime, objectTypeArgument BACnetObjectType) *_ListOfCovNotificationsValue {
	return &_ListOfCovNotificationsValue{PropertyIdentifier: propertyIdentifier, ArrayIndex: arrayIndex, PropertyValue: propertyValue, TimeOfChange: timeOfChange, ObjectTypeArgument: objectTypeArgument}
}

// Deprecated: use the interface for direct cast
func CastListOfCovNotificationsValue(structType any) ListOfCovNotificationsValue {
	if casted, ok := structType.(ListOfCovNotificationsValue); ok {
		return casted
	}
	if casted, ok := structType.(*ListOfCovNotificationsValue); ok {
		return *casted
	}
	return nil
}

func (m *_ListOfCovNotificationsValue) GetTypeName() string {
	return "ListOfCovNotificationsValue"
}

func (m *_ListOfCovNotificationsValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits(ctx)

	// Optional Field (arrayIndex)
	if m.ArrayIndex != nil {
		lengthInBits += m.ArrayIndex.GetLengthInBits(ctx)
	}

	// Simple field (propertyValue)
	lengthInBits += m.PropertyValue.GetLengthInBits(ctx)

	// Optional Field (timeOfChange)
	if m.TimeOfChange != nil {
		lengthInBits += m.TimeOfChange.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_ListOfCovNotificationsValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ListOfCovNotificationsValueParse(ctx context.Context, theBytes []byte, objectTypeArgument BACnetObjectType) (ListOfCovNotificationsValue, error) {
	return ListOfCovNotificationsValueParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), objectTypeArgument)
}

func ListOfCovNotificationsValueParseWithBufferProducer(objectTypeArgument BACnetObjectType) func(ctx context.Context, readBuffer utils.ReadBuffer) (ListOfCovNotificationsValue, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ListOfCovNotificationsValue, error) {
		return ListOfCovNotificationsValueParseWithBuffer(ctx, readBuffer, objectTypeArgument)
	}
}

func ListOfCovNotificationsValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (ListOfCovNotificationsValue, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ListOfCovNotificationsValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ListOfCovNotificationsValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	propertyIdentifier, err := ReadSimpleField[BACnetPropertyIdentifierTagged](ctx, "propertyIdentifier", ReadComplex[BACnetPropertyIdentifierTagged](BACnetPropertyIdentifierTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyIdentifier' field"))
	}

	_arrayIndex, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "arrayIndex", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayIndex' field"))
	}
	var arrayIndex BACnetContextTagUnsignedInteger
	if _arrayIndex != nil {
		arrayIndex = *_arrayIndex
	}

	propertyValue, err := ReadSimpleField[BACnetConstructedData](ctx, "propertyValue", ReadComplex[BACnetConstructedData](BACnetConstructedDataParseWithBufferProducer[BACnetConstructedData]((uint8)(uint8(2)), (BACnetObjectType)(objectTypeArgument), (BACnetPropertyIdentifier)(propertyIdentifier.GetValue()), (BACnetTagPayloadUnsignedInteger)((CastBACnetTagPayloadUnsignedInteger(utils.InlineIf(bool((arrayIndex) != (nil)), func() any { return CastBACnetTagPayloadUnsignedInteger((arrayIndex).GetPayload()) }, func() any { return CastBACnetTagPayloadUnsignedInteger(nil) }))))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyValue' field"))
	}

	_timeOfChange, err := ReadOptionalField[BACnetContextTagTime](ctx, "timeOfChange", ReadComplex[BACnetContextTagTime](BACnetContextTagParseWithBufferProducer((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_TIME)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeOfChange' field"))
	}
	var timeOfChange BACnetContextTagTime
	if _timeOfChange != nil {
		timeOfChange = *_timeOfChange
	}

	if closeErr := readBuffer.CloseContext("ListOfCovNotificationsValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ListOfCovNotificationsValue")
	}

	// Create the instance
	return &_ListOfCovNotificationsValue{
		ObjectTypeArgument: objectTypeArgument,
		PropertyIdentifier: propertyIdentifier,
		ArrayIndex:         arrayIndex,
		PropertyValue:      propertyValue,
		TimeOfChange:       timeOfChange,
	}, nil
}

func (m *_ListOfCovNotificationsValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ListOfCovNotificationsValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ListOfCovNotificationsValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ListOfCovNotificationsValue")
	}

	// Simple Field (propertyIdentifier)
	if pushErr := writeBuffer.PushContext("propertyIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for propertyIdentifier")
	}
	_propertyIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetPropertyIdentifier())
	if popErr := writeBuffer.PopContext("propertyIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for propertyIdentifier")
	}
	if _propertyIdentifierErr != nil {
		return errors.Wrap(_propertyIdentifierErr, "Error serializing 'propertyIdentifier' field")
	}

	// Optional Field (arrayIndex) (Can be skipped, if the value is null)
	var arrayIndex BACnetContextTagUnsignedInteger = nil
	if m.GetArrayIndex() != nil {
		if pushErr := writeBuffer.PushContext("arrayIndex"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for arrayIndex")
		}
		arrayIndex = m.GetArrayIndex()
		_arrayIndexErr := writeBuffer.WriteSerializable(ctx, arrayIndex)
		if popErr := writeBuffer.PopContext("arrayIndex"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for arrayIndex")
		}
		if _arrayIndexErr != nil {
			return errors.Wrap(_arrayIndexErr, "Error serializing 'arrayIndex' field")
		}
	}

	// Simple Field (propertyValue)
	if pushErr := writeBuffer.PushContext("propertyValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for propertyValue")
	}
	_propertyValueErr := writeBuffer.WriteSerializable(ctx, m.GetPropertyValue())
	if popErr := writeBuffer.PopContext("propertyValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for propertyValue")
	}
	if _propertyValueErr != nil {
		return errors.Wrap(_propertyValueErr, "Error serializing 'propertyValue' field")
	}

	// Optional Field (timeOfChange) (Can be skipped, if the value is null)
	var timeOfChange BACnetContextTagTime = nil
	if m.GetTimeOfChange() != nil {
		if pushErr := writeBuffer.PushContext("timeOfChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timeOfChange")
		}
		timeOfChange = m.GetTimeOfChange()
		_timeOfChangeErr := writeBuffer.WriteSerializable(ctx, timeOfChange)
		if popErr := writeBuffer.PopContext("timeOfChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timeOfChange")
		}
		if _timeOfChangeErr != nil {
			return errors.Wrap(_timeOfChangeErr, "Error serializing 'timeOfChange' field")
		}
	}

	if popErr := writeBuffer.PopContext("ListOfCovNotificationsValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ListOfCovNotificationsValue")
	}
	return nil
}

////
// Arguments Getter

func (m *_ListOfCovNotificationsValue) GetObjectTypeArgument() BACnetObjectType {
	return m.ObjectTypeArgument
}

//
////

func (m *_ListOfCovNotificationsValue) isListOfCovNotificationsValue() bool {
	return true
}

func (m *_ListOfCovNotificationsValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
