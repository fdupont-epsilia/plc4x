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

// BACnetConstructedDataStructuredObjectList is the corresponding interface of BACnetConstructedDataStructuredObjectList
type BACnetConstructedDataStructuredObjectList interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetStructuredObjectList returns StructuredObjectList (property field)
	GetStructuredObjectList() []BACnetApplicationTagObjectIdentifier
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// IsBACnetConstructedDataStructuredObjectList is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataStructuredObjectList()
}

// _BACnetConstructedDataStructuredObjectList is the data-structure of this message
type _BACnetConstructedDataStructuredObjectList struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	StructuredObjectList []BACnetApplicationTagObjectIdentifier
}

var _ BACnetConstructedDataStructuredObjectList = (*_BACnetConstructedDataStructuredObjectList)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataStructuredObjectList)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataStructuredObjectList) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataStructuredObjectList) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_STRUCTURED_OBJECT_LIST
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataStructuredObjectList) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataStructuredObjectList) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataStructuredObjectList) GetStructuredObjectList() []BACnetApplicationTagObjectIdentifier {
	return m.StructuredObjectList
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataStructuredObjectList) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.GetNumberOfDataElements()
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataStructuredObjectList factory function for _BACnetConstructedDataStructuredObjectList
func NewBACnetConstructedDataStructuredObjectList(numberOfDataElements BACnetApplicationTagUnsignedInteger, structuredObjectList []BACnetApplicationTagObjectIdentifier, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataStructuredObjectList {
	_result := &_BACnetConstructedDataStructuredObjectList{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		StructuredObjectList:          structuredObjectList,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataStructuredObjectList(structType any) BACnetConstructedDataStructuredObjectList {
	if casted, ok := structType.(BACnetConstructedDataStructuredObjectList); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataStructuredObjectList); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataStructuredObjectList) GetTypeName() string {
	return "BACnetConstructedDataStructuredObjectList"
}

func (m *_BACnetConstructedDataStructuredObjectList) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.StructuredObjectList) > 0 {
		for _, element := range m.StructuredObjectList {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataStructuredObjectList) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataStructuredObjectList) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataStructuredObjectList BACnetConstructedDataStructuredObjectList, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataStructuredObjectList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataStructuredObjectList")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zero, err := ReadVirtualField[uint64](ctx, "zero", (*uint64)(nil), uint64(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zero' field"))
	}
	_ = zero

	var numberOfDataElements BACnetApplicationTagUnsignedInteger
	_numberOfDataElements, err := ReadOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfDataElements' field"))
	}
	if _numberOfDataElements != nil {
		numberOfDataElements = *_numberOfDataElements
		m.NumberOfDataElements = numberOfDataElements
	}

	structuredObjectList, err := ReadTerminatedArrayField[BACnetApplicationTagObjectIdentifier](ctx, "structuredObjectList", ReadComplex[BACnetApplicationTagObjectIdentifier](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagObjectIdentifier](), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'structuredObjectList' field"))
	}
	m.StructuredObjectList = structuredObjectList

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataStructuredObjectList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataStructuredObjectList")
	}

	return m, nil
}

func (m *_BACnetConstructedDataStructuredObjectList) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataStructuredObjectList) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataStructuredObjectList"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataStructuredObjectList")
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

		if err := WriteComplexTypeArrayField(ctx, "structuredObjectList", m.GetStructuredObjectList(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'structuredObjectList' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataStructuredObjectList"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataStructuredObjectList")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataStructuredObjectList) IsBACnetConstructedDataStructuredObjectList() {}

func (m *_BACnetConstructedDataStructuredObjectList) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
