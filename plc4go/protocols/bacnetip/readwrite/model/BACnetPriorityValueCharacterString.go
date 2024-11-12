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

// BACnetPriorityValueCharacterString is the corresponding interface of BACnetPriorityValueCharacterString
type BACnetPriorityValueCharacterString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPriorityValue
	// GetCharacterStringValue returns CharacterStringValue (property field)
	GetCharacterStringValue() BACnetApplicationTagCharacterString
	// IsBACnetPriorityValueCharacterString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPriorityValueCharacterString()
	// CreateBuilder creates a BACnetPriorityValueCharacterStringBuilder
	CreateBACnetPriorityValueCharacterStringBuilder() BACnetPriorityValueCharacterStringBuilder
}

// _BACnetPriorityValueCharacterString is the data-structure of this message
type _BACnetPriorityValueCharacterString struct {
	BACnetPriorityValueContract
	CharacterStringValue BACnetApplicationTagCharacterString
}

var _ BACnetPriorityValueCharacterString = (*_BACnetPriorityValueCharacterString)(nil)
var _ BACnetPriorityValueRequirements = (*_BACnetPriorityValueCharacterString)(nil)

// NewBACnetPriorityValueCharacterString factory function for _BACnetPriorityValueCharacterString
func NewBACnetPriorityValueCharacterString(peekedTagHeader BACnetTagHeader, characterStringValue BACnetApplicationTagCharacterString, objectTypeArgument BACnetObjectType) *_BACnetPriorityValueCharacterString {
	if characterStringValue == nil {
		panic("characterStringValue of type BACnetApplicationTagCharacterString for BACnetPriorityValueCharacterString must not be nil")
	}
	_result := &_BACnetPriorityValueCharacterString{
		BACnetPriorityValueContract: NewBACnetPriorityValue(peekedTagHeader, objectTypeArgument),
		CharacterStringValue:        characterStringValue,
	}
	_result.BACnetPriorityValueContract.(*_BACnetPriorityValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPriorityValueCharacterStringBuilder is a builder for BACnetPriorityValueCharacterString
type BACnetPriorityValueCharacterStringBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(characterStringValue BACnetApplicationTagCharacterString) BACnetPriorityValueCharacterStringBuilder
	// WithCharacterStringValue adds CharacterStringValue (property field)
	WithCharacterStringValue(BACnetApplicationTagCharacterString) BACnetPriorityValueCharacterStringBuilder
	// WithCharacterStringValueBuilder adds CharacterStringValue (property field) which is build by the builder
	WithCharacterStringValueBuilder(func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetPriorityValueCharacterStringBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetPriorityValueBuilder
	// Build builds the BACnetPriorityValueCharacterString or returns an error if something is wrong
	Build() (BACnetPriorityValueCharacterString, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPriorityValueCharacterString
}

// NewBACnetPriorityValueCharacterStringBuilder() creates a BACnetPriorityValueCharacterStringBuilder
func NewBACnetPriorityValueCharacterStringBuilder() BACnetPriorityValueCharacterStringBuilder {
	return &_BACnetPriorityValueCharacterStringBuilder{_BACnetPriorityValueCharacterString: new(_BACnetPriorityValueCharacterString)}
}

type _BACnetPriorityValueCharacterStringBuilder struct {
	*_BACnetPriorityValueCharacterString

	parentBuilder *_BACnetPriorityValueBuilder

	err *utils.MultiError
}

var _ (BACnetPriorityValueCharacterStringBuilder) = (*_BACnetPriorityValueCharacterStringBuilder)(nil)

func (b *_BACnetPriorityValueCharacterStringBuilder) setParent(contract BACnetPriorityValueContract) {
	b.BACnetPriorityValueContract = contract
	contract.(*_BACnetPriorityValue)._SubType = b._BACnetPriorityValueCharacterString
}

func (b *_BACnetPriorityValueCharacterStringBuilder) WithMandatoryFields(characterStringValue BACnetApplicationTagCharacterString) BACnetPriorityValueCharacterStringBuilder {
	return b.WithCharacterStringValue(characterStringValue)
}

func (b *_BACnetPriorityValueCharacterStringBuilder) WithCharacterStringValue(characterStringValue BACnetApplicationTagCharacterString) BACnetPriorityValueCharacterStringBuilder {
	b.CharacterStringValue = characterStringValue
	return b
}

func (b *_BACnetPriorityValueCharacterStringBuilder) WithCharacterStringValueBuilder(builderSupplier func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetPriorityValueCharacterStringBuilder {
	builder := builderSupplier(b.CharacterStringValue.CreateBACnetApplicationTagCharacterStringBuilder())
	var err error
	b.CharacterStringValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagCharacterStringBuilder failed"))
	}
	return b
}

func (b *_BACnetPriorityValueCharacterStringBuilder) Build() (BACnetPriorityValueCharacterString, error) {
	if b.CharacterStringValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'characterStringValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPriorityValueCharacterString.deepCopy(), nil
}

func (b *_BACnetPriorityValueCharacterStringBuilder) MustBuild() BACnetPriorityValueCharacterString {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetPriorityValueCharacterStringBuilder) Done() BACnetPriorityValueBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetPriorityValueBuilder().(*_BACnetPriorityValueBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetPriorityValueCharacterStringBuilder) buildForBACnetPriorityValue() (BACnetPriorityValue, error) {
	return b.Build()
}

func (b *_BACnetPriorityValueCharacterStringBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPriorityValueCharacterStringBuilder().(*_BACnetPriorityValueCharacterStringBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPriorityValueCharacterStringBuilder creates a BACnetPriorityValueCharacterStringBuilder
func (b *_BACnetPriorityValueCharacterString) CreateBACnetPriorityValueCharacterStringBuilder() BACnetPriorityValueCharacterStringBuilder {
	if b == nil {
		return NewBACnetPriorityValueCharacterStringBuilder()
	}
	return &_BACnetPriorityValueCharacterStringBuilder{_BACnetPriorityValueCharacterString: b.deepCopy()}
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

func (m *_BACnetPriorityValueCharacterString) GetParent() BACnetPriorityValueContract {
	return m.BACnetPriorityValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPriorityValueCharacterString) GetCharacterStringValue() BACnetApplicationTagCharacterString {
	return m.CharacterStringValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPriorityValueCharacterString(structType any) BACnetPriorityValueCharacterString {
	if casted, ok := structType.(BACnetPriorityValueCharacterString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPriorityValueCharacterString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPriorityValueCharacterString) GetTypeName() string {
	return "BACnetPriorityValueCharacterString"
}

func (m *_BACnetPriorityValueCharacterString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPriorityValueContract.(*_BACnetPriorityValue).getLengthInBits(ctx))

	// Simple field (characterStringValue)
	lengthInBits += m.CharacterStringValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPriorityValueCharacterString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPriorityValueCharacterString) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPriorityValue, objectTypeArgument BACnetObjectType) (__bACnetPriorityValueCharacterString BACnetPriorityValueCharacterString, err error) {
	m.BACnetPriorityValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityValueCharacterString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValueCharacterString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	characterStringValue, err := ReadSimpleField[BACnetApplicationTagCharacterString](ctx, "characterStringValue", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'characterStringValue' field"))
	}
	m.CharacterStringValue = characterStringValue

	if closeErr := readBuffer.CloseContext("BACnetPriorityValueCharacterString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValueCharacterString")
	}

	return m, nil
}

func (m *_BACnetPriorityValueCharacterString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPriorityValueCharacterString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPriorityValueCharacterString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValueCharacterString")
		}

		if err := WriteSimpleField[BACnetApplicationTagCharacterString](ctx, "characterStringValue", m.GetCharacterStringValue(), WriteComplex[BACnetApplicationTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'characterStringValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPriorityValueCharacterString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPriorityValueCharacterString")
		}
		return nil
	}
	return m.BACnetPriorityValueContract.(*_BACnetPriorityValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPriorityValueCharacterString) IsBACnetPriorityValueCharacterString() {}

func (m *_BACnetPriorityValueCharacterString) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPriorityValueCharacterString) deepCopy() *_BACnetPriorityValueCharacterString {
	if m == nil {
		return nil
	}
	_BACnetPriorityValueCharacterStringCopy := &_BACnetPriorityValueCharacterString{
		m.BACnetPriorityValueContract.(*_BACnetPriorityValue).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagCharacterString](m.CharacterStringValue),
	}
	_BACnetPriorityValueCharacterStringCopy.BACnetPriorityValueContract.(*_BACnetPriorityValue)._SubType = m
	return _BACnetPriorityValueCharacterStringCopy
}

func (m *_BACnetPriorityValueCharacterString) String() string {
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
