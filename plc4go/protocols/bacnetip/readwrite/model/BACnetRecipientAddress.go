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

// BACnetRecipientAddress is the corresponding interface of BACnetRecipientAddress
type BACnetRecipientAddress interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetRecipient
	// GetAddressValue returns AddressValue (property field)
	GetAddressValue() BACnetAddressEnclosed
	// IsBACnetRecipientAddress is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetRecipientAddress()
	// CreateBuilder creates a BACnetRecipientAddressBuilder
	CreateBACnetRecipientAddressBuilder() BACnetRecipientAddressBuilder
}

// _BACnetRecipientAddress is the data-structure of this message
type _BACnetRecipientAddress struct {
	BACnetRecipientContract
	AddressValue BACnetAddressEnclosed
}

var _ BACnetRecipientAddress = (*_BACnetRecipientAddress)(nil)
var _ BACnetRecipientRequirements = (*_BACnetRecipientAddress)(nil)

// NewBACnetRecipientAddress factory function for _BACnetRecipientAddress
func NewBACnetRecipientAddress(peekedTagHeader BACnetTagHeader, addressValue BACnetAddressEnclosed) *_BACnetRecipientAddress {
	if addressValue == nil {
		panic("addressValue of type BACnetAddressEnclosed for BACnetRecipientAddress must not be nil")
	}
	_result := &_BACnetRecipientAddress{
		BACnetRecipientContract: NewBACnetRecipient(peekedTagHeader),
		AddressValue:            addressValue,
	}
	_result.BACnetRecipientContract.(*_BACnetRecipient)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetRecipientAddressBuilder is a builder for BACnetRecipientAddress
type BACnetRecipientAddressBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(addressValue BACnetAddressEnclosed) BACnetRecipientAddressBuilder
	// WithAddressValue adds AddressValue (property field)
	WithAddressValue(BACnetAddressEnclosed) BACnetRecipientAddressBuilder
	// WithAddressValueBuilder adds AddressValue (property field) which is build by the builder
	WithAddressValueBuilder(func(BACnetAddressEnclosedBuilder) BACnetAddressEnclosedBuilder) BACnetRecipientAddressBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetRecipientBuilder
	// Build builds the BACnetRecipientAddress or returns an error if something is wrong
	Build() (BACnetRecipientAddress, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetRecipientAddress
}

// NewBACnetRecipientAddressBuilder() creates a BACnetRecipientAddressBuilder
func NewBACnetRecipientAddressBuilder() BACnetRecipientAddressBuilder {
	return &_BACnetRecipientAddressBuilder{_BACnetRecipientAddress: new(_BACnetRecipientAddress)}
}

type _BACnetRecipientAddressBuilder struct {
	*_BACnetRecipientAddress

	parentBuilder *_BACnetRecipientBuilder

	err *utils.MultiError
}

var _ (BACnetRecipientAddressBuilder) = (*_BACnetRecipientAddressBuilder)(nil)

func (b *_BACnetRecipientAddressBuilder) setParent(contract BACnetRecipientContract) {
	b.BACnetRecipientContract = contract
	contract.(*_BACnetRecipient)._SubType = b._BACnetRecipientAddress
}

func (b *_BACnetRecipientAddressBuilder) WithMandatoryFields(addressValue BACnetAddressEnclosed) BACnetRecipientAddressBuilder {
	return b.WithAddressValue(addressValue)
}

func (b *_BACnetRecipientAddressBuilder) WithAddressValue(addressValue BACnetAddressEnclosed) BACnetRecipientAddressBuilder {
	b.AddressValue = addressValue
	return b
}

func (b *_BACnetRecipientAddressBuilder) WithAddressValueBuilder(builderSupplier func(BACnetAddressEnclosedBuilder) BACnetAddressEnclosedBuilder) BACnetRecipientAddressBuilder {
	builder := builderSupplier(b.AddressValue.CreateBACnetAddressEnclosedBuilder())
	var err error
	b.AddressValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetAddressEnclosedBuilder failed"))
	}
	return b
}

func (b *_BACnetRecipientAddressBuilder) Build() (BACnetRecipientAddress, error) {
	if b.AddressValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'addressValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetRecipientAddress.deepCopy(), nil
}

func (b *_BACnetRecipientAddressBuilder) MustBuild() BACnetRecipientAddress {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetRecipientAddressBuilder) Done() BACnetRecipientBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetRecipientBuilder().(*_BACnetRecipientBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetRecipientAddressBuilder) buildForBACnetRecipient() (BACnetRecipient, error) {
	return b.Build()
}

func (b *_BACnetRecipientAddressBuilder) DeepCopy() any {
	_copy := b.CreateBACnetRecipientAddressBuilder().(*_BACnetRecipientAddressBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetRecipientAddressBuilder creates a BACnetRecipientAddressBuilder
func (b *_BACnetRecipientAddress) CreateBACnetRecipientAddressBuilder() BACnetRecipientAddressBuilder {
	if b == nil {
		return NewBACnetRecipientAddressBuilder()
	}
	return &_BACnetRecipientAddressBuilder{_BACnetRecipientAddress: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

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

func (m *_BACnetRecipientAddress) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetRecipient) (__bACnetRecipientAddress BACnetRecipientAddress, err error) {
	m.BACnetRecipientContract = parent
	parent._SubType = m
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

func (m *_BACnetRecipientAddress) IsBACnetRecipientAddress() {}

func (m *_BACnetRecipientAddress) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetRecipientAddress) deepCopy() *_BACnetRecipientAddress {
	if m == nil {
		return nil
	}
	_BACnetRecipientAddressCopy := &_BACnetRecipientAddress{
		m.BACnetRecipientContract.(*_BACnetRecipient).deepCopy(),
		utils.DeepCopy[BACnetAddressEnclosed](m.AddressValue),
	}
	_BACnetRecipientAddressCopy.BACnetRecipientContract.(*_BACnetRecipient)._SubType = m
	return _BACnetRecipientAddressCopy
}

func (m *_BACnetRecipientAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
