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

// BACnetConstructedDataDeviceMaxInfoFrames is the corresponding interface of BACnetConstructedDataDeviceMaxInfoFrames
type BACnetConstructedDataDeviceMaxInfoFrames interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetMaxInfoFrames returns MaxInfoFrames (property field)
	GetMaxInfoFrames() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataDeviceMaxInfoFramesExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDeviceMaxInfoFrames.
// This is useful for switch cases.
type BACnetConstructedDataDeviceMaxInfoFramesExactly interface {
	BACnetConstructedDataDeviceMaxInfoFrames
	isBACnetConstructedDataDeviceMaxInfoFrames() bool
}

// _BACnetConstructedDataDeviceMaxInfoFrames is the data-structure of this message
type _BACnetConstructedDataDeviceMaxInfoFrames struct {
	BACnetConstructedDataContract
	MaxInfoFrames BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataDeviceMaxInfoFrames = (*_BACnetConstructedDataDeviceMaxInfoFrames)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDeviceMaxInfoFrames) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_DEVICE
}

func (m *_BACnetConstructedDataDeviceMaxInfoFrames) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAX_INFO_FRAMES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDeviceMaxInfoFrames) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDeviceMaxInfoFrames) GetMaxInfoFrames() BACnetApplicationTagUnsignedInteger {
	return m.MaxInfoFrames
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDeviceMaxInfoFrames) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetMaxInfoFrames())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDeviceMaxInfoFrames factory function for _BACnetConstructedDataDeviceMaxInfoFrames
func NewBACnetConstructedDataDeviceMaxInfoFrames(maxInfoFrames BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDeviceMaxInfoFrames {
	_result := &_BACnetConstructedDataDeviceMaxInfoFrames{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		MaxInfoFrames:                 maxInfoFrames,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDeviceMaxInfoFrames(structType any) BACnetConstructedDataDeviceMaxInfoFrames {
	if casted, ok := structType.(BACnetConstructedDataDeviceMaxInfoFrames); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDeviceMaxInfoFrames); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDeviceMaxInfoFrames) GetTypeName() string {
	return "BACnetConstructedDataDeviceMaxInfoFrames"
}

func (m *_BACnetConstructedDataDeviceMaxInfoFrames) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (maxInfoFrames)
	lengthInBits += m.MaxInfoFrames.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDeviceMaxInfoFrames) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDeviceMaxInfoFrames) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (_bACnetConstructedDataDeviceMaxInfoFrames BACnetConstructedDataDeviceMaxInfoFrames, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDeviceMaxInfoFrames"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDeviceMaxInfoFrames")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	maxInfoFrames, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "maxInfoFrames", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxInfoFrames' field"))
	}
	m.MaxInfoFrames = maxInfoFrames

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), maxInfoFrames)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDeviceMaxInfoFrames"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDeviceMaxInfoFrames")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDeviceMaxInfoFrames) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDeviceMaxInfoFrames) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDeviceMaxInfoFrames"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDeviceMaxInfoFrames")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "maxInfoFrames", m.GetMaxInfoFrames(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxInfoFrames' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDeviceMaxInfoFrames"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDeviceMaxInfoFrames")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDeviceMaxInfoFrames) isBACnetConstructedDataDeviceMaxInfoFrames() bool {
	return true
}

func (m *_BACnetConstructedDataDeviceMaxInfoFrames) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
