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

// BACnetConstructedDataLastAccessPoint is the corresponding interface of BACnetConstructedDataLastAccessPoint
type BACnetConstructedDataLastAccessPoint interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLastAccessPoint returns LastAccessPoint (property field)
	GetLastAccessPoint() BACnetDeviceObjectReference
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDeviceObjectReference
}

// BACnetConstructedDataLastAccessPointExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLastAccessPoint.
// This is useful for switch cases.
type BACnetConstructedDataLastAccessPointExactly interface {
	BACnetConstructedDataLastAccessPoint
	isBACnetConstructedDataLastAccessPoint() bool
}

// _BACnetConstructedDataLastAccessPoint is the data-structure of this message
type _BACnetConstructedDataLastAccessPoint struct {
	*_BACnetConstructedData
	LastAccessPoint BACnetDeviceObjectReference
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLastAccessPoint) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLastAccessPoint) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LAST_ACCESS_POINT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLastAccessPoint) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLastAccessPoint) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLastAccessPoint) GetLastAccessPoint() BACnetDeviceObjectReference {
	return m.LastAccessPoint
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLastAccessPoint) GetActualValue() BACnetDeviceObjectReference {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDeviceObjectReference(m.GetLastAccessPoint())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLastAccessPoint factory function for _BACnetConstructedDataLastAccessPoint
func NewBACnetConstructedDataLastAccessPoint(lastAccessPoint BACnetDeviceObjectReference, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLastAccessPoint {
	_result := &_BACnetConstructedDataLastAccessPoint{
		LastAccessPoint:        lastAccessPoint,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLastAccessPoint(structType any) BACnetConstructedDataLastAccessPoint {
	if casted, ok := structType.(BACnetConstructedDataLastAccessPoint); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLastAccessPoint); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLastAccessPoint) GetTypeName() string {
	return "BACnetConstructedDataLastAccessPoint"
}

func (m *_BACnetConstructedDataLastAccessPoint) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (lastAccessPoint)
	lengthInBits += m.LastAccessPoint.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLastAccessPoint) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataLastAccessPointParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLastAccessPoint, error) {
	return BACnetConstructedDataLastAccessPointParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataLastAccessPointParseWithBufferProducer(tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataLastAccessPoint, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataLastAccessPoint, error) {
		return BACnetConstructedDataLastAccessPointParseWithBuffer(ctx, readBuffer, tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
	}
}

func BACnetConstructedDataLastAccessPointParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLastAccessPoint, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLastAccessPoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLastAccessPoint")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lastAccessPoint, err := ReadSimpleField[BACnetDeviceObjectReference](ctx, "lastAccessPoint", ReadComplex[BACnetDeviceObjectReference](BACnetDeviceObjectReferenceParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastAccessPoint' field"))
	}

	actualValue, err := ReadVirtualField[BACnetDeviceObjectReference](ctx, "actualValue", (*BACnetDeviceObjectReference)(nil), lastAccessPoint)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLastAccessPoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLastAccessPoint")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLastAccessPoint{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		LastAccessPoint: lastAccessPoint,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLastAccessPoint) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLastAccessPoint) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLastAccessPoint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLastAccessPoint")
		}

		// Simple Field (lastAccessPoint)
		if pushErr := writeBuffer.PushContext("lastAccessPoint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lastAccessPoint")
		}
		_lastAccessPointErr := writeBuffer.WriteSerializable(ctx, m.GetLastAccessPoint())
		if popErr := writeBuffer.PopContext("lastAccessPoint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lastAccessPoint")
		}
		if _lastAccessPointErr != nil {
			return errors.Wrap(_lastAccessPointErr, "Error serializing 'lastAccessPoint' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLastAccessPoint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLastAccessPoint")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLastAccessPoint) isBACnetConstructedDataLastAccessPoint() bool {
	return true
}

func (m *_BACnetConstructedDataLastAccessPoint) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
