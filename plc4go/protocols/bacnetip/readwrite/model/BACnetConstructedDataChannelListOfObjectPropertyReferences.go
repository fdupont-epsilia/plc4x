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

// BACnetConstructedDataChannelListOfObjectPropertyReferences is the corresponding interface of BACnetConstructedDataChannelListOfObjectPropertyReferences
type BACnetConstructedDataChannelListOfObjectPropertyReferences interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetReferences returns References (property field)
	GetReferences() []BACnetDeviceObjectPropertyReference
	// GetZero returns Zero (virtual field)
	GetZero() uint64
}

// BACnetConstructedDataChannelListOfObjectPropertyReferencesExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataChannelListOfObjectPropertyReferences.
// This is useful for switch cases.
type BACnetConstructedDataChannelListOfObjectPropertyReferencesExactly interface {
	BACnetConstructedDataChannelListOfObjectPropertyReferences
	isBACnetConstructedDataChannelListOfObjectPropertyReferences() bool
}

// _BACnetConstructedDataChannelListOfObjectPropertyReferences is the data-structure of this message
type _BACnetConstructedDataChannelListOfObjectPropertyReferences struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	References           []BACnetDeviceObjectPropertyReference
}

var _ BACnetConstructedDataChannelListOfObjectPropertyReferences = (*_BACnetConstructedDataChannelListOfObjectPropertyReferences)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataChannelListOfObjectPropertyReferences) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_CHANNEL
}

func (m *_BACnetConstructedDataChannelListOfObjectPropertyReferences) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LIST_OF_OBJECT_PROPERTY_REFERENCES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataChannelListOfObjectPropertyReferences) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataChannelListOfObjectPropertyReferences) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataChannelListOfObjectPropertyReferences) GetReferences() []BACnetDeviceObjectPropertyReference {
	return m.References
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataChannelListOfObjectPropertyReferences) GetZero() uint64 {
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

// NewBACnetConstructedDataChannelListOfObjectPropertyReferences factory function for _BACnetConstructedDataChannelListOfObjectPropertyReferences
func NewBACnetConstructedDataChannelListOfObjectPropertyReferences(numberOfDataElements BACnetApplicationTagUnsignedInteger, references []BACnetDeviceObjectPropertyReference, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataChannelListOfObjectPropertyReferences {
	_result := &_BACnetConstructedDataChannelListOfObjectPropertyReferences{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		References:                    references,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataChannelListOfObjectPropertyReferences(structType any) BACnetConstructedDataChannelListOfObjectPropertyReferences {
	if casted, ok := structType.(BACnetConstructedDataChannelListOfObjectPropertyReferences); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataChannelListOfObjectPropertyReferences); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataChannelListOfObjectPropertyReferences) GetTypeName() string {
	return "BACnetConstructedDataChannelListOfObjectPropertyReferences"
}

func (m *_BACnetConstructedDataChannelListOfObjectPropertyReferences) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.References) > 0 {
		for _, element := range m.References {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataChannelListOfObjectPropertyReferences) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataChannelListOfObjectPropertyReferences) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (_bACnetConstructedDataChannelListOfObjectPropertyReferences BACnetConstructedDataChannelListOfObjectPropertyReferences, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataChannelListOfObjectPropertyReferences"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataChannelListOfObjectPropertyReferences")
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

	references, err := ReadTerminatedArrayField[BACnetDeviceObjectPropertyReference](ctx, "references", ReadComplex[BACnetDeviceObjectPropertyReference](BACnetDeviceObjectPropertyReferenceParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'references' field"))
	}
	m.References = references

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataChannelListOfObjectPropertyReferences"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataChannelListOfObjectPropertyReferences")
	}

	return m, nil
}

func (m *_BACnetConstructedDataChannelListOfObjectPropertyReferences) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataChannelListOfObjectPropertyReferences) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataChannelListOfObjectPropertyReferences"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataChannelListOfObjectPropertyReferences")
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

		if err := WriteComplexTypeArrayField(ctx, "references", m.GetReferences(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'references' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataChannelListOfObjectPropertyReferences"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataChannelListOfObjectPropertyReferences")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataChannelListOfObjectPropertyReferences) isBACnetConstructedDataChannelListOfObjectPropertyReferences() bool {
	return true
}

func (m *_BACnetConstructedDataChannelListOfObjectPropertyReferences) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
