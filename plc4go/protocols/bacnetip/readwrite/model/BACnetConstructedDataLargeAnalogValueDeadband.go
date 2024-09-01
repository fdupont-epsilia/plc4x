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

// BACnetConstructedDataLargeAnalogValueDeadband is the corresponding interface of BACnetConstructedDataLargeAnalogValueDeadband
type BACnetConstructedDataLargeAnalogValueDeadband interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetDeadband returns Deadband (property field)
	GetDeadband() BACnetApplicationTagDouble
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagDouble
}

// BACnetConstructedDataLargeAnalogValueDeadbandExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLargeAnalogValueDeadband.
// This is useful for switch cases.
type BACnetConstructedDataLargeAnalogValueDeadbandExactly interface {
	BACnetConstructedDataLargeAnalogValueDeadband
	isBACnetConstructedDataLargeAnalogValueDeadband() bool
}

// _BACnetConstructedDataLargeAnalogValueDeadband is the data-structure of this message
type _BACnetConstructedDataLargeAnalogValueDeadband struct {
	*_BACnetConstructedData
	Deadband BACnetApplicationTagDouble
}

var _ BACnetConstructedDataLargeAnalogValueDeadband = (*_BACnetConstructedDataLargeAnalogValueDeadband)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueDeadband) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LARGE_ANALOG_VALUE
}

func (m *_BACnetConstructedDataLargeAnalogValueDeadband) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DEADBAND
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueDeadband) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLargeAnalogValueDeadband) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueDeadband) GetDeadband() BACnetApplicationTagDouble {
	return m.Deadband
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueDeadband) GetActualValue() BACnetApplicationTagDouble {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagDouble(m.GetDeadband())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLargeAnalogValueDeadband factory function for _BACnetConstructedDataLargeAnalogValueDeadband
func NewBACnetConstructedDataLargeAnalogValueDeadband(deadband BACnetApplicationTagDouble, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLargeAnalogValueDeadband {
	_result := &_BACnetConstructedDataLargeAnalogValueDeadband{
		Deadband:               deadband,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLargeAnalogValueDeadband(structType any) BACnetConstructedDataLargeAnalogValueDeadband {
	if casted, ok := structType.(BACnetConstructedDataLargeAnalogValueDeadband); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLargeAnalogValueDeadband); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLargeAnalogValueDeadband) GetTypeName() string {
	return "BACnetConstructedDataLargeAnalogValueDeadband"
}

func (m *_BACnetConstructedDataLargeAnalogValueDeadband) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (deadband)
	lengthInBits += m.Deadband.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLargeAnalogValueDeadband) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataLargeAnalogValueDeadbandParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLargeAnalogValueDeadband, error) {
	return BACnetConstructedDataLargeAnalogValueDeadbandParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataLargeAnalogValueDeadbandParseWithBufferProducer(tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataLargeAnalogValueDeadband, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataLargeAnalogValueDeadband, error) {
		return BACnetConstructedDataLargeAnalogValueDeadbandParseWithBuffer(ctx, readBuffer, tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
	}
}

func BACnetConstructedDataLargeAnalogValueDeadbandParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLargeAnalogValueDeadband, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLargeAnalogValueDeadband"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLargeAnalogValueDeadband")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	deadband, err := ReadSimpleField[BACnetApplicationTagDouble](ctx, "deadband", ReadComplex[BACnetApplicationTagDouble](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDouble](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deadband' field"))
	}

	actualValue, err := ReadVirtualField[BACnetApplicationTagDouble](ctx, "actualValue", (*BACnetApplicationTagDouble)(nil), deadband)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLargeAnalogValueDeadband"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLargeAnalogValueDeadband")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLargeAnalogValueDeadband{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		Deadband: deadband,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLargeAnalogValueDeadband) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLargeAnalogValueDeadband) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLargeAnalogValueDeadband"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLargeAnalogValueDeadband")
		}

		if err := WriteSimpleField[BACnetApplicationTagDouble](ctx, "deadband", m.GetDeadband(), WriteComplex[BACnetApplicationTagDouble](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'deadband' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLargeAnalogValueDeadband"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLargeAnalogValueDeadband")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLargeAnalogValueDeadband) isBACnetConstructedDataLargeAnalogValueDeadband() bool {
	return true
}

func (m *_BACnetConstructedDataLargeAnalogValueDeadband) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
