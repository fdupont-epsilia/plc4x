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

// BACnetConstructedDataNetworkAccessSecurityPolicies is the corresponding interface of BACnetConstructedDataNetworkAccessSecurityPolicies
type BACnetConstructedDataNetworkAccessSecurityPolicies interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetNetworkAccessSecurityPolicies returns NetworkAccessSecurityPolicies (property field)
	GetNetworkAccessSecurityPolicies() []BACnetNetworkSecurityPolicy
	// GetZero returns Zero (virtual field)
	GetZero() uint64
}

// BACnetConstructedDataNetworkAccessSecurityPoliciesExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataNetworkAccessSecurityPolicies.
// This is useful for switch cases.
type BACnetConstructedDataNetworkAccessSecurityPoliciesExactly interface {
	BACnetConstructedDataNetworkAccessSecurityPolicies
	isBACnetConstructedDataNetworkAccessSecurityPolicies() bool
}

// _BACnetConstructedDataNetworkAccessSecurityPolicies is the data-structure of this message
type _BACnetConstructedDataNetworkAccessSecurityPolicies struct {
	BACnetConstructedDataContract
	NumberOfDataElements          BACnetApplicationTagUnsignedInteger
	NetworkAccessSecurityPolicies []BACnetNetworkSecurityPolicy
}

var _ BACnetConstructedDataNetworkAccessSecurityPolicies = (*_BACnetConstructedDataNetworkAccessSecurityPolicies)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataNetworkAccessSecurityPolicies) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataNetworkAccessSecurityPolicies) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_NETWORK_ACCESS_SECURITY_POLICIES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataNetworkAccessSecurityPolicies) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataNetworkAccessSecurityPolicies) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataNetworkAccessSecurityPolicies) GetNetworkAccessSecurityPolicies() []BACnetNetworkSecurityPolicy {
	return m.NetworkAccessSecurityPolicies
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataNetworkAccessSecurityPolicies) GetZero() uint64 {
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

// NewBACnetConstructedDataNetworkAccessSecurityPolicies factory function for _BACnetConstructedDataNetworkAccessSecurityPolicies
func NewBACnetConstructedDataNetworkAccessSecurityPolicies(numberOfDataElements BACnetApplicationTagUnsignedInteger, networkAccessSecurityPolicies []BACnetNetworkSecurityPolicy, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataNetworkAccessSecurityPolicies {
	_result := &_BACnetConstructedDataNetworkAccessSecurityPolicies{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		NetworkAccessSecurityPolicies: networkAccessSecurityPolicies,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataNetworkAccessSecurityPolicies(structType any) BACnetConstructedDataNetworkAccessSecurityPolicies {
	if casted, ok := structType.(BACnetConstructedDataNetworkAccessSecurityPolicies); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataNetworkAccessSecurityPolicies); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataNetworkAccessSecurityPolicies) GetTypeName() string {
	return "BACnetConstructedDataNetworkAccessSecurityPolicies"
}

func (m *_BACnetConstructedDataNetworkAccessSecurityPolicies) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.NetworkAccessSecurityPolicies) > 0 {
		for _, element := range m.NetworkAccessSecurityPolicies {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataNetworkAccessSecurityPolicies) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataNetworkAccessSecurityPolicies) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (_bACnetConstructedDataNetworkAccessSecurityPolicies BACnetConstructedDataNetworkAccessSecurityPolicies, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataNetworkAccessSecurityPolicies"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataNetworkAccessSecurityPolicies")
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

	networkAccessSecurityPolicies, err := ReadTerminatedArrayField[BACnetNetworkSecurityPolicy](ctx, "networkAccessSecurityPolicies", ReadComplex[BACnetNetworkSecurityPolicy](BACnetNetworkSecurityPolicyParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'networkAccessSecurityPolicies' field"))
	}
	m.NetworkAccessSecurityPolicies = networkAccessSecurityPolicies

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataNetworkAccessSecurityPolicies"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataNetworkAccessSecurityPolicies")
	}

	return m, nil
}

func (m *_BACnetConstructedDataNetworkAccessSecurityPolicies) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataNetworkAccessSecurityPolicies) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataNetworkAccessSecurityPolicies"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataNetworkAccessSecurityPolicies")
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

		if err := WriteComplexTypeArrayField(ctx, "networkAccessSecurityPolicies", m.GetNetworkAccessSecurityPolicies(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'networkAccessSecurityPolicies' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataNetworkAccessSecurityPolicies"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataNetworkAccessSecurityPolicies")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataNetworkAccessSecurityPolicies) isBACnetConstructedDataNetworkAccessSecurityPolicies() bool {
	return true
}

func (m *_BACnetConstructedDataNetworkAccessSecurityPolicies) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
