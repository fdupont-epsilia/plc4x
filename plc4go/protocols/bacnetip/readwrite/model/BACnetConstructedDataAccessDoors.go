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

// BACnetConstructedDataAccessDoors is the corresponding interface of BACnetConstructedDataAccessDoors
type BACnetConstructedDataAccessDoors interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetAccessDoors returns AccessDoors (property field)
	GetAccessDoors() []BACnetDeviceObjectReference
	// GetZero returns Zero (virtual field)
	GetZero() uint64
}

// BACnetConstructedDataAccessDoorsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAccessDoors.
// This is useful for switch cases.
type BACnetConstructedDataAccessDoorsExactly interface {
	BACnetConstructedDataAccessDoors
	isBACnetConstructedDataAccessDoors() bool
}

// _BACnetConstructedDataAccessDoors is the data-structure of this message
type _BACnetConstructedDataAccessDoors struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	AccessDoors          []BACnetDeviceObjectReference
}

var _ BACnetConstructedDataAccessDoors = (*_BACnetConstructedDataAccessDoors)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccessDoors) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAccessDoors) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACCESS_DOORS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccessDoors) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccessDoors) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataAccessDoors) GetAccessDoors() []BACnetDeviceObjectReference {
	return m.AccessDoors
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAccessDoors) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAccessDoors factory function for _BACnetConstructedDataAccessDoors
func NewBACnetConstructedDataAccessDoors(numberOfDataElements BACnetApplicationTagUnsignedInteger, accessDoors []BACnetDeviceObjectReference, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccessDoors {
	_result := &_BACnetConstructedDataAccessDoors{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		AccessDoors:                   accessDoors,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccessDoors(structType any) BACnetConstructedDataAccessDoors {
	if casted, ok := structType.(BACnetConstructedDataAccessDoors); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccessDoors); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccessDoors) GetTypeName() string {
	return "BACnetConstructedDataAccessDoors"
}

func (m *_BACnetConstructedDataAccessDoors) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.AccessDoors) > 0 {
		for _, element := range m.AccessDoors {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataAccessDoors) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAccessDoors) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (_bACnetConstructedDataAccessDoors BACnetConstructedDataAccessDoors, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccessDoors"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccessDoors")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zero, err := ReadVirtualField[uint64](ctx, "zero", (*uint64)(nil), uint64(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zero' field"))
	}
	_ = zero

	_numberOfDataElements, err := ReadOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfDataElements' field"))
	}
	var numberOfDataElements BACnetApplicationTagUnsignedInteger
	if _numberOfDataElements != nil {
		numberOfDataElements = *_numberOfDataElements
	}
	m.NumberOfDataElements = numberOfDataElements

	accessDoors, err := ReadTerminatedArrayField[BACnetDeviceObjectReference](ctx, "accessDoors", ReadComplex[BACnetDeviceObjectReference](BACnetDeviceObjectReferenceParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'accessDoors' field"))
	}
	m.AccessDoors = accessDoors

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccessDoors"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccessDoors")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAccessDoors) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccessDoors) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccessDoors"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccessDoors")
		}
		// Virtual field
		zero := m.GetZero()
		_ = zero
		if _zeroErr := writeBuffer.WriteVirtual(ctx, "zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		if err := WriteOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", GetRef(m.GetNumberOfDataElements()), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'numberOfDataElements' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "accessDoors", m.GetAccessDoors(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'accessDoors' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccessDoors"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccessDoors")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccessDoors) isBACnetConstructedDataAccessDoors() bool {
	return true
}

func (m *_BACnetConstructedDataAccessDoors) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
