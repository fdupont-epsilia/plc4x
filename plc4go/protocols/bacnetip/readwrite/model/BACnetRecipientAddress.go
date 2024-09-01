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

// BACnetRecipientAddress is the corresponding interface of BACnetRecipientAddress
type BACnetRecipientAddress interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetRecipient
	// GetAddressValue returns AddressValue (property field)
	GetAddressValue() BACnetAddressEnclosed
}

// BACnetRecipientAddressExactly can be used when we want exactly this type and not a type which fulfills BACnetRecipientAddress.
// This is useful for switch cases.
type BACnetRecipientAddressExactly interface {
	BACnetRecipientAddress
	isBACnetRecipientAddress() bool
}

// _BACnetRecipientAddress is the data-structure of this message
type _BACnetRecipientAddress struct {
	BACnetRecipientContract
	AddressValue BACnetAddressEnclosed
}

var _ BACnetRecipientAddress = (*_BACnetRecipientAddress)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetRecipientAddress) GetParent() BACnetRecipientContract {
	return m.BACnetRecipientContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetRecipientAddress) GetAddressValue() BACnetAddressEnclosed {
	return m.AddressValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetRecipientAddress factory function for _BACnetRecipientAddress
func NewBACnetRecipientAddress(addressValue BACnetAddressEnclosed, peekedTagHeader BACnetTagHeader) *_BACnetRecipientAddress {
	_result := &_BACnetRecipientAddress{
		BACnetRecipientContract: NewBACnetRecipient(peekedTagHeader),
		AddressValue:            addressValue,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetRecipientAddress(structType any) BACnetRecipientAddress {
	if casted, ok := structType.(BACnetRecipientAddress); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetRecipientAddress); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetRecipientAddress) GetTypeName() string {
	return "BACnetRecipientAddress"
}

func (m *_BACnetRecipientAddress) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetRecipientContract.(*_BACnetRecipient).getLengthInBits(ctx))

	// Simple field (addressValue)
	lengthInBits += m.AddressValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetRecipientAddress) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetRecipientAddress) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetRecipient) (_bACnetRecipientAddress BACnetRecipientAddress, err error) {
	m.BACnetRecipientContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetRecipientAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetRecipientAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	addressValue, err := ReadSimpleField[BACnetAddressEnclosed](ctx, "addressValue", ReadComplex[BACnetAddressEnclosed](BACnetAddressEnclosedParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'addressValue' field"))
	}
	m.AddressValue = addressValue

	if closeErr := readBuffer.CloseContext("BACnetRecipientAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetRecipientAddress")
	}

	return m, nil
}

func (m *_BACnetRecipientAddress) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetRecipientAddress) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetRecipientAddress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetRecipientAddress")
		}

		if err := WriteSimpleField[BACnetAddressEnclosed](ctx, "addressValue", m.GetAddressValue(), WriteComplex[BACnetAddressEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'addressValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetRecipientAddress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetRecipientAddress")
		}
		return nil
	}
	return m.BACnetRecipientContract.(*_BACnetRecipient).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetRecipientAddress) isBACnetRecipientAddress() bool {
	return true
}

func (m *_BACnetRecipientAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
