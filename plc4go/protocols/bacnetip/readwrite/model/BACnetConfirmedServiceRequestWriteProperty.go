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

// BACnetConfirmedServiceRequestWriteProperty is the corresponding interface of BACnetConfirmedServiceRequestWriteProperty
type BACnetConfirmedServiceRequestWriteProperty interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() BACnetPropertyIdentifierTagged
	// GetArrayIndex returns ArrayIndex (property field)
	GetArrayIndex() BACnetContextTagUnsignedInteger
	// GetPropertyValue returns PropertyValue (property field)
	GetPropertyValue() BACnetConstructedData
	// GetPriority returns Priority (property field)
	GetPriority() BACnetContextTagUnsignedInteger
}

// BACnetConfirmedServiceRequestWritePropertyExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestWriteProperty.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestWritePropertyExactly interface {
	BACnetConfirmedServiceRequestWriteProperty
	isBACnetConfirmedServiceRequestWriteProperty() bool
}

// _BACnetConfirmedServiceRequestWriteProperty is the data-structure of this message
type _BACnetConfirmedServiceRequestWriteProperty struct {
	*_BACnetConfirmedServiceRequest
	ObjectIdentifier   BACnetContextTagObjectIdentifier
	PropertyIdentifier BACnetPropertyIdentifierTagged
	ArrayIndex         BACnetContextTagUnsignedInteger
	PropertyValue      BACnetConstructedData
	Priority           BACnetContextTagUnsignedInteger
}

var _ BACnetConfirmedServiceRequestWriteProperty = (*_BACnetConfirmedServiceRequestWriteProperty)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestWriteProperty) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_WRITE_PROPERTY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestWriteProperty) InitializeParent(parent BACnetConfirmedServiceRequest) {
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestWriteProperty) GetObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) GetPropertyIdentifier() BACnetPropertyIdentifierTagged {
	return m.PropertyIdentifier
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) GetArrayIndex() BACnetContextTagUnsignedInteger {
	return m.ArrayIndex
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) GetPropertyValue() BACnetConstructedData {
	return m.PropertyValue
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) GetPriority() BACnetContextTagUnsignedInteger {
	return m.Priority
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestWriteProperty factory function for _BACnetConfirmedServiceRequestWriteProperty
func NewBACnetConfirmedServiceRequestWriteProperty(objectIdentifier BACnetContextTagObjectIdentifier, propertyIdentifier BACnetPropertyIdentifierTagged, arrayIndex BACnetContextTagUnsignedInteger, propertyValue BACnetConstructedData, priority BACnetContextTagUnsignedInteger, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestWriteProperty {
	_result := &_BACnetConfirmedServiceRequestWriteProperty{
		ObjectIdentifier:               objectIdentifier,
		PropertyIdentifier:             propertyIdentifier,
		ArrayIndex:                     arrayIndex,
		PropertyValue:                  propertyValue,
		Priority:                       priority,
		_BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestWriteProperty(structType any) BACnetConfirmedServiceRequestWriteProperty {
	if casted, ok := structType.(BACnetConfirmedServiceRequestWriteProperty); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestWriteProperty); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) GetTypeName() string {
	return "BACnetConfirmedServiceRequestWriteProperty"
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits(ctx)

	// Optional Field (arrayIndex)
	if m.ArrayIndex != nil {
		lengthInBits += m.ArrayIndex.GetLengthInBits(ctx)
	}

	// Simple field (propertyValue)
	lengthInBits += m.PropertyValue.GetLengthInBits(ctx)

	// Optional Field (priority)
	if m.Priority != nil {
		lengthInBits += m.Priority.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestWritePropertyParse(ctx context.Context, theBytes []byte, serviceRequestLength uint32) (BACnetConfirmedServiceRequestWriteProperty, error) {
	return BACnetConfirmedServiceRequestWritePropertyParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), serviceRequestLength)
}

func BACnetConfirmedServiceRequestWritePropertyParseWithBufferProducer(serviceRequestLength uint32) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestWriteProperty, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestWriteProperty, error) {
		return BACnetConfirmedServiceRequestWritePropertyParseWithBuffer(ctx, readBuffer, serviceRequestLength)
	}
}

func BACnetConfirmedServiceRequestWritePropertyParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceRequestLength uint32) (BACnetConfirmedServiceRequestWriteProperty, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestWriteProperty"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestWriteProperty")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "objectIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectIdentifier' field"))
	}

	propertyIdentifier, err := ReadSimpleField[BACnetPropertyIdentifierTagged](ctx, "propertyIdentifier", ReadComplex[BACnetPropertyIdentifierTagged](BACnetPropertyIdentifierTaggedParseWithBufferProducer((uint8)(uint8(1)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyIdentifier' field"))
	}

	_arrayIndex, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "arrayIndex", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayIndex' field"))
	}
	var arrayIndex BACnetContextTagUnsignedInteger
	if _arrayIndex != nil {
		arrayIndex = *_arrayIndex
	}

	propertyValue, err := ReadSimpleField[BACnetConstructedData](ctx, "propertyValue", ReadComplex[BACnetConstructedData](BACnetConstructedDataParseWithBufferProducer[BACnetConstructedData]((uint8)(uint8(3)), (BACnetObjectType)(objectIdentifier.GetObjectType()), (BACnetPropertyIdentifier)(propertyIdentifier.GetValue()), (BACnetTagPayloadUnsignedInteger)((CastBACnetTagPayloadUnsignedInteger(utils.InlineIf(bool((arrayIndex) != (nil)), func() any { return CastBACnetTagPayloadUnsignedInteger((arrayIndex).GetPayload()) }, func() any { return CastBACnetTagPayloadUnsignedInteger(nil) }))))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyValue' field"))
	}

	_priority, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "priority", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(4)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priority' field"))
	}
	var priority BACnetContextTagUnsignedInteger
	if _priority != nil {
		priority = *_priority
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestWriteProperty"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestWriteProperty")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestWriteProperty{
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		ObjectIdentifier:   objectIdentifier,
		PropertyIdentifier: propertyIdentifier,
		ArrayIndex:         arrayIndex,
		PropertyValue:      propertyValue,
		Priority:           priority,
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestWriteProperty"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestWriteProperty")
		}

		if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "objectIdentifier", m.GetObjectIdentifier(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'objectIdentifier' field")
		}

		if err := WriteSimpleField[BACnetPropertyIdentifierTagged](ctx, "propertyIdentifier", m.GetPropertyIdentifier(), WriteComplex[BACnetPropertyIdentifierTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'propertyIdentifier' field")
		}

		if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "arrayIndex", GetRef(m.GetArrayIndex()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'arrayIndex' field")
		}

		if err := WriteSimpleField[BACnetConstructedData](ctx, "propertyValue", m.GetPropertyValue(), WriteComplex[BACnetConstructedData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'propertyValue' field")
		}

		if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "priority", GetRef(m.GetPriority()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'priority' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestWriteProperty"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestWriteProperty")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) isBACnetConfirmedServiceRequestWriteProperty() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
