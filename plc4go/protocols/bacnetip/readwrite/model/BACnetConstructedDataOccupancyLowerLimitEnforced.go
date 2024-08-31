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

// BACnetConstructedDataOccupancyLowerLimitEnforced is the corresponding interface of BACnetConstructedDataOccupancyLowerLimitEnforced
type BACnetConstructedDataOccupancyLowerLimitEnforced interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetOccupancyLowerLimitEnforced returns OccupancyLowerLimitEnforced (property field)
	GetOccupancyLowerLimitEnforced() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataOccupancyLowerLimitEnforcedExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataOccupancyLowerLimitEnforced.
// This is useful for switch cases.
type BACnetConstructedDataOccupancyLowerLimitEnforcedExactly interface {
	BACnetConstructedDataOccupancyLowerLimitEnforced
	isBACnetConstructedDataOccupancyLowerLimitEnforced() bool
}

// _BACnetConstructedDataOccupancyLowerLimitEnforced is the data-structure of this message
type _BACnetConstructedDataOccupancyLowerLimitEnforced struct {
	*_BACnetConstructedData
	OccupancyLowerLimitEnforced BACnetApplicationTagBoolean
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataOccupancyLowerLimitEnforced) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataOccupancyLowerLimitEnforced) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_OCCUPANCY_LOWER_LIMIT_ENFORCED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataOccupancyLowerLimitEnforced) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataOccupancyLowerLimitEnforced) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataOccupancyLowerLimitEnforced) GetOccupancyLowerLimitEnforced() BACnetApplicationTagBoolean {
	return m.OccupancyLowerLimitEnforced
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataOccupancyLowerLimitEnforced) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetOccupancyLowerLimitEnforced())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataOccupancyLowerLimitEnforced factory function for _BACnetConstructedDataOccupancyLowerLimitEnforced
func NewBACnetConstructedDataOccupancyLowerLimitEnforced(occupancyLowerLimitEnforced BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataOccupancyLowerLimitEnforced {
	_result := &_BACnetConstructedDataOccupancyLowerLimitEnforced{
		OccupancyLowerLimitEnforced: occupancyLowerLimitEnforced,
		_BACnetConstructedData:      NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataOccupancyLowerLimitEnforced(structType any) BACnetConstructedDataOccupancyLowerLimitEnforced {
	if casted, ok := structType.(BACnetConstructedDataOccupancyLowerLimitEnforced); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataOccupancyLowerLimitEnforced); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataOccupancyLowerLimitEnforced) GetTypeName() string {
	return "BACnetConstructedDataOccupancyLowerLimitEnforced"
}

func (m *_BACnetConstructedDataOccupancyLowerLimitEnforced) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (occupancyLowerLimitEnforced)
	lengthInBits += m.OccupancyLowerLimitEnforced.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataOccupancyLowerLimitEnforced) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataOccupancyLowerLimitEnforcedParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataOccupancyLowerLimitEnforced, error) {
	return BACnetConstructedDataOccupancyLowerLimitEnforcedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataOccupancyLowerLimitEnforcedParseWithBufferProducer(tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataOccupancyLowerLimitEnforced, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataOccupancyLowerLimitEnforced, error) {
		return BACnetConstructedDataOccupancyLowerLimitEnforcedParseWithBuffer(ctx, readBuffer, tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
	}
}

func BACnetConstructedDataOccupancyLowerLimitEnforcedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataOccupancyLowerLimitEnforced, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataOccupancyLowerLimitEnforced"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataOccupancyLowerLimitEnforced")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	occupancyLowerLimitEnforced, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "occupancyLowerLimitEnforced", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer(), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'occupancyLowerLimitEnforced' field"))
	}

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), occupancyLowerLimitEnforced)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataOccupancyLowerLimitEnforced"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataOccupancyLowerLimitEnforced")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataOccupancyLowerLimitEnforced{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		OccupancyLowerLimitEnforced: occupancyLowerLimitEnforced,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataOccupancyLowerLimitEnforced) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataOccupancyLowerLimitEnforced) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataOccupancyLowerLimitEnforced"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataOccupancyLowerLimitEnforced")
		}

		// Simple Field (occupancyLowerLimitEnforced)
		if pushErr := writeBuffer.PushContext("occupancyLowerLimitEnforced"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for occupancyLowerLimitEnforced")
		}
		_occupancyLowerLimitEnforcedErr := writeBuffer.WriteSerializable(ctx, m.GetOccupancyLowerLimitEnforced())
		if popErr := writeBuffer.PopContext("occupancyLowerLimitEnforced"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for occupancyLowerLimitEnforced")
		}
		if _occupancyLowerLimitEnforcedErr != nil {
			return errors.Wrap(_occupancyLowerLimitEnforcedErr, "Error serializing 'occupancyLowerLimitEnforced' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataOccupancyLowerLimitEnforced"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataOccupancyLowerLimitEnforced")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataOccupancyLowerLimitEnforced) isBACnetConstructedDataOccupancyLowerLimitEnforced() bool {
	return true
}

func (m *_BACnetConstructedDataOccupancyLowerLimitEnforced) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
