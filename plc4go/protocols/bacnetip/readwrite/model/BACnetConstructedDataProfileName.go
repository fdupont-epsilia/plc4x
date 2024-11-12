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

// BACnetConstructedDataProfileName is the corresponding interface of BACnetConstructedDataProfileName
type BACnetConstructedDataProfileName interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetProfileName returns ProfileName (property field)
	GetProfileName() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
	// IsBACnetConstructedDataProfileName is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataProfileName()
	// CreateBuilder creates a BACnetConstructedDataProfileNameBuilder
	CreateBACnetConstructedDataProfileNameBuilder() BACnetConstructedDataProfileNameBuilder
}

// _BACnetConstructedDataProfileName is the data-structure of this message
type _BACnetConstructedDataProfileName struct {
	BACnetConstructedDataContract
	ProfileName BACnetApplicationTagCharacterString
}

var _ BACnetConstructedDataProfileName = (*_BACnetConstructedDataProfileName)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataProfileName)(nil)

// NewBACnetConstructedDataProfileName factory function for _BACnetConstructedDataProfileName
func NewBACnetConstructedDataProfileName(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, profileName BACnetApplicationTagCharacterString, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataProfileName {
	if profileName == nil {
		panic("profileName of type BACnetApplicationTagCharacterString for BACnetConstructedDataProfileName must not be nil")
	}
	_result := &_BACnetConstructedDataProfileName{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ProfileName:                   profileName,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataProfileNameBuilder is a builder for BACnetConstructedDataProfileName
type BACnetConstructedDataProfileNameBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(profileName BACnetApplicationTagCharacterString) BACnetConstructedDataProfileNameBuilder
	// WithProfileName adds ProfileName (property field)
	WithProfileName(BACnetApplicationTagCharacterString) BACnetConstructedDataProfileNameBuilder
	// WithProfileNameBuilder adds ProfileName (property field) which is build by the builder
	WithProfileNameBuilder(func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetConstructedDataProfileNameBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataProfileName or returns an error if something is wrong
	Build() (BACnetConstructedDataProfileName, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataProfileName
}

// NewBACnetConstructedDataProfileNameBuilder() creates a BACnetConstructedDataProfileNameBuilder
func NewBACnetConstructedDataProfileNameBuilder() BACnetConstructedDataProfileNameBuilder {
	return &_BACnetConstructedDataProfileNameBuilder{_BACnetConstructedDataProfileName: new(_BACnetConstructedDataProfileName)}
}

type _BACnetConstructedDataProfileNameBuilder struct {
	*_BACnetConstructedDataProfileName

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataProfileNameBuilder) = (*_BACnetConstructedDataProfileNameBuilder)(nil)

func (b *_BACnetConstructedDataProfileNameBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataProfileName
}

func (b *_BACnetConstructedDataProfileNameBuilder) WithMandatoryFields(profileName BACnetApplicationTagCharacterString) BACnetConstructedDataProfileNameBuilder {
	return b.WithProfileName(profileName)
}

func (b *_BACnetConstructedDataProfileNameBuilder) WithProfileName(profileName BACnetApplicationTagCharacterString) BACnetConstructedDataProfileNameBuilder {
	b.ProfileName = profileName
	return b
}

func (b *_BACnetConstructedDataProfileNameBuilder) WithProfileNameBuilder(builderSupplier func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetConstructedDataProfileNameBuilder {
	builder := builderSupplier(b.ProfileName.CreateBACnetApplicationTagCharacterStringBuilder())
	var err error
	b.ProfileName, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagCharacterStringBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataProfileNameBuilder) Build() (BACnetConstructedDataProfileName, error) {
	if b.ProfileName == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'profileName' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataProfileName.deepCopy(), nil
}

func (b *_BACnetConstructedDataProfileNameBuilder) MustBuild() BACnetConstructedDataProfileName {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataProfileNameBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataProfileNameBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataProfileNameBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataProfileNameBuilder().(*_BACnetConstructedDataProfileNameBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataProfileNameBuilder creates a BACnetConstructedDataProfileNameBuilder
func (b *_BACnetConstructedDataProfileName) CreateBACnetConstructedDataProfileNameBuilder() BACnetConstructedDataProfileNameBuilder {
	if b == nil {
		return NewBACnetConstructedDataProfileNameBuilder()
	}
	return &_BACnetConstructedDataProfileNameBuilder{_BACnetConstructedDataProfileName: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataProfileName) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataProfileName) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PROFILE_NAME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataProfileName) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataProfileName) GetProfileName() BACnetApplicationTagCharacterString {
	return m.ProfileName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataProfileName) GetActualValue() BACnetApplicationTagCharacterString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagCharacterString(m.GetProfileName())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataProfileName(structType any) BACnetConstructedDataProfileName {
	if casted, ok := structType.(BACnetConstructedDataProfileName); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataProfileName); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataProfileName) GetTypeName() string {
	return "BACnetConstructedDataProfileName"
}

func (m *_BACnetConstructedDataProfileName) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (profileName)
	lengthInBits += m.ProfileName.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataProfileName) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataProfileName) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataProfileName BACnetConstructedDataProfileName, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataProfileName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataProfileName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	profileName, err := ReadSimpleField[BACnetApplicationTagCharacterString](ctx, "profileName", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'profileName' field"))
	}
	m.ProfileName = profileName

	actualValue, err := ReadVirtualField[BACnetApplicationTagCharacterString](ctx, "actualValue", (*BACnetApplicationTagCharacterString)(nil), profileName)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataProfileName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataProfileName")
	}

	return m, nil
}

func (m *_BACnetConstructedDataProfileName) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataProfileName) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataProfileName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataProfileName")
		}

		if err := WriteSimpleField[BACnetApplicationTagCharacterString](ctx, "profileName", m.GetProfileName(), WriteComplex[BACnetApplicationTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'profileName' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataProfileName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataProfileName")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataProfileName) IsBACnetConstructedDataProfileName() {}

func (m *_BACnetConstructedDataProfileName) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataProfileName) deepCopy() *_BACnetConstructedDataProfileName {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataProfileNameCopy := &_BACnetConstructedDataProfileName{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagCharacterString](m.ProfileName),
	}
	_BACnetConstructedDataProfileNameCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataProfileNameCopy
}

func (m *_BACnetConstructedDataProfileName) String() string {
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
