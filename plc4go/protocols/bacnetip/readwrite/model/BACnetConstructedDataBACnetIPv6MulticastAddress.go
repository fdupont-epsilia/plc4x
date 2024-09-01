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

// BACnetConstructedDataBACnetIPv6MulticastAddress is the corresponding interface of BACnetConstructedDataBACnetIPv6MulticastAddress
type BACnetConstructedDataBACnetIPv6MulticastAddress interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetIpv6MulticastAddress returns Ipv6MulticastAddress (property field)
	GetIpv6MulticastAddress() BACnetApplicationTagOctetString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagOctetString
}

// BACnetConstructedDataBACnetIPv6MulticastAddressExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataBACnetIPv6MulticastAddress.
// This is useful for switch cases.
type BACnetConstructedDataBACnetIPv6MulticastAddressExactly interface {
	BACnetConstructedDataBACnetIPv6MulticastAddress
	isBACnetConstructedDataBACnetIPv6MulticastAddress() bool
}

// _BACnetConstructedDataBACnetIPv6MulticastAddress is the data-structure of this message
type _BACnetConstructedDataBACnetIPv6MulticastAddress struct {
	*_BACnetConstructedData
	Ipv6MulticastAddress BACnetApplicationTagOctetString
}

var _ BACnetConstructedDataBACnetIPv6MulticastAddress = (*_BACnetConstructedDataBACnetIPv6MulticastAddress)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BACNET_IPV6_MULTICAST_ADDRESS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) GetIpv6MulticastAddress() BACnetApplicationTagOctetString {
	return m.Ipv6MulticastAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) GetActualValue() BACnetApplicationTagOctetString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagOctetString(m.GetIpv6MulticastAddress())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBACnetIPv6MulticastAddress factory function for _BACnetConstructedDataBACnetIPv6MulticastAddress
func NewBACnetConstructedDataBACnetIPv6MulticastAddress(ipv6MulticastAddress BACnetApplicationTagOctetString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBACnetIPv6MulticastAddress {
	_result := &_BACnetConstructedDataBACnetIPv6MulticastAddress{
		Ipv6MulticastAddress:   ipv6MulticastAddress,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBACnetIPv6MulticastAddress(structType any) BACnetConstructedDataBACnetIPv6MulticastAddress {
	if casted, ok := structType.(BACnetConstructedDataBACnetIPv6MulticastAddress); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBACnetIPv6MulticastAddress); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) GetTypeName() string {
	return "BACnetConstructedDataBACnetIPv6MulticastAddress"
}

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (ipv6MulticastAddress)
	lengthInBits += m.Ipv6MulticastAddress.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataBACnetIPv6MulticastAddressParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBACnetIPv6MulticastAddress, error) {
	return BACnetConstructedDataBACnetIPv6MulticastAddressParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataBACnetIPv6MulticastAddressParseWithBufferProducer(tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataBACnetIPv6MulticastAddress, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataBACnetIPv6MulticastAddress, error) {
		return BACnetConstructedDataBACnetIPv6MulticastAddressParseWithBuffer(ctx, readBuffer, tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
	}
}

func BACnetConstructedDataBACnetIPv6MulticastAddressParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBACnetIPv6MulticastAddress, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBACnetIPv6MulticastAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBACnetIPv6MulticastAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	ipv6MulticastAddress, err := ReadSimpleField[BACnetApplicationTagOctetString](ctx, "ipv6MulticastAddress", ReadComplex[BACnetApplicationTagOctetString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagOctetString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ipv6MulticastAddress' field"))
	}

	actualValue, err := ReadVirtualField[BACnetApplicationTagOctetString](ctx, "actualValue", (*BACnetApplicationTagOctetString)(nil), ipv6MulticastAddress)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBACnetIPv6MulticastAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBACnetIPv6MulticastAddress")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataBACnetIPv6MulticastAddress{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		Ipv6MulticastAddress: ipv6MulticastAddress,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBACnetIPv6MulticastAddress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBACnetIPv6MulticastAddress")
		}

		if err := WriteSimpleField[BACnetApplicationTagOctetString](ctx, "ipv6MulticastAddress", m.GetIpv6MulticastAddress(), WriteComplex[BACnetApplicationTagOctetString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'ipv6MulticastAddress' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBACnetIPv6MulticastAddress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBACnetIPv6MulticastAddress")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) isBACnetConstructedDataBACnetIPv6MulticastAddress() bool {
	return true
}

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
