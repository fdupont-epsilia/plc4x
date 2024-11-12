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

// BACnetConstructedDataDateTimePatternValuePresentValue is the corresponding interface of BACnetConstructedDataDateTimePatternValuePresentValue
type BACnetConstructedDataDateTimePatternValuePresentValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() BACnetDateTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDateTime
	// IsBACnetConstructedDataDateTimePatternValuePresentValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataDateTimePatternValuePresentValue()
	// CreateBuilder creates a BACnetConstructedDataDateTimePatternValuePresentValueBuilder
	CreateBACnetConstructedDataDateTimePatternValuePresentValueBuilder() BACnetConstructedDataDateTimePatternValuePresentValueBuilder
}

// _BACnetConstructedDataDateTimePatternValuePresentValue is the data-structure of this message
type _BACnetConstructedDataDateTimePatternValuePresentValue struct {
	BACnetConstructedDataContract
	PresentValue BACnetDateTime
}

var _ BACnetConstructedDataDateTimePatternValuePresentValue = (*_BACnetConstructedDataDateTimePatternValuePresentValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataDateTimePatternValuePresentValue)(nil)

// NewBACnetConstructedDataDateTimePatternValuePresentValue factory function for _BACnetConstructedDataDateTimePatternValuePresentValue
func NewBACnetConstructedDataDateTimePatternValuePresentValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, presentValue BACnetDateTime, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDateTimePatternValuePresentValue {
	if presentValue == nil {
		panic("presentValue of type BACnetDateTime for BACnetConstructedDataDateTimePatternValuePresentValue must not be nil")
	}
	_result := &_BACnetConstructedDataDateTimePatternValuePresentValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		PresentValue:                  presentValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataDateTimePatternValuePresentValueBuilder is a builder for BACnetConstructedDataDateTimePatternValuePresentValue
type BACnetConstructedDataDateTimePatternValuePresentValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(presentValue BACnetDateTime) BACnetConstructedDataDateTimePatternValuePresentValueBuilder
	// WithPresentValue adds PresentValue (property field)
	WithPresentValue(BACnetDateTime) BACnetConstructedDataDateTimePatternValuePresentValueBuilder
	// WithPresentValueBuilder adds PresentValue (property field) which is build by the builder
	WithPresentValueBuilder(func(BACnetDateTimeBuilder) BACnetDateTimeBuilder) BACnetConstructedDataDateTimePatternValuePresentValueBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataDateTimePatternValuePresentValue or returns an error if something is wrong
	Build() (BACnetConstructedDataDateTimePatternValuePresentValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataDateTimePatternValuePresentValue
}

// NewBACnetConstructedDataDateTimePatternValuePresentValueBuilder() creates a BACnetConstructedDataDateTimePatternValuePresentValueBuilder
func NewBACnetConstructedDataDateTimePatternValuePresentValueBuilder() BACnetConstructedDataDateTimePatternValuePresentValueBuilder {
	return &_BACnetConstructedDataDateTimePatternValuePresentValueBuilder{_BACnetConstructedDataDateTimePatternValuePresentValue: new(_BACnetConstructedDataDateTimePatternValuePresentValue)}
}

type _BACnetConstructedDataDateTimePatternValuePresentValueBuilder struct {
	*_BACnetConstructedDataDateTimePatternValuePresentValue

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataDateTimePatternValuePresentValueBuilder) = (*_BACnetConstructedDataDateTimePatternValuePresentValueBuilder)(nil)

func (b *_BACnetConstructedDataDateTimePatternValuePresentValueBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataDateTimePatternValuePresentValue
}

func (b *_BACnetConstructedDataDateTimePatternValuePresentValueBuilder) WithMandatoryFields(presentValue BACnetDateTime) BACnetConstructedDataDateTimePatternValuePresentValueBuilder {
	return b.WithPresentValue(presentValue)
}

func (b *_BACnetConstructedDataDateTimePatternValuePresentValueBuilder) WithPresentValue(presentValue BACnetDateTime) BACnetConstructedDataDateTimePatternValuePresentValueBuilder {
	b.PresentValue = presentValue
	return b
}

func (b *_BACnetConstructedDataDateTimePatternValuePresentValueBuilder) WithPresentValueBuilder(builderSupplier func(BACnetDateTimeBuilder) BACnetDateTimeBuilder) BACnetConstructedDataDateTimePatternValuePresentValueBuilder {
	builder := builderSupplier(b.PresentValue.CreateBACnetDateTimeBuilder())
	var err error
	b.PresentValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDateTimeBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataDateTimePatternValuePresentValueBuilder) Build() (BACnetConstructedDataDateTimePatternValuePresentValue, error) {
	if b.PresentValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'presentValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataDateTimePatternValuePresentValue.deepCopy(), nil
}

func (b *_BACnetConstructedDataDateTimePatternValuePresentValueBuilder) MustBuild() BACnetConstructedDataDateTimePatternValuePresentValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataDateTimePatternValuePresentValueBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataDateTimePatternValuePresentValueBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataDateTimePatternValuePresentValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataDateTimePatternValuePresentValueBuilder().(*_BACnetConstructedDataDateTimePatternValuePresentValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataDateTimePatternValuePresentValueBuilder creates a BACnetConstructedDataDateTimePatternValuePresentValueBuilder
func (b *_BACnetConstructedDataDateTimePatternValuePresentValue) CreateBACnetConstructedDataDateTimePatternValuePresentValueBuilder() BACnetConstructedDataDateTimePatternValuePresentValueBuilder {
	if b == nil {
		return NewBACnetConstructedDataDateTimePatternValuePresentValueBuilder()
	}
	return &_BACnetConstructedDataDateTimePatternValuePresentValueBuilder{_BACnetConstructedDataDateTimePatternValuePresentValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDateTimePatternValuePresentValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_DATETIMEPATTERN_VALUE
}

func (m *_BACnetConstructedDataDateTimePatternValuePresentValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRESENT_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDateTimePatternValuePresentValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDateTimePatternValuePresentValue) GetPresentValue() BACnetDateTime {
	return m.PresentValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDateTimePatternValuePresentValue) GetActualValue() BACnetDateTime {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDateTime(m.GetPresentValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDateTimePatternValuePresentValue(structType any) BACnetConstructedDataDateTimePatternValuePresentValue {
	if casted, ok := structType.(BACnetConstructedDataDateTimePatternValuePresentValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDateTimePatternValuePresentValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDateTimePatternValuePresentValue) GetTypeName() string {
	return "BACnetConstructedDataDateTimePatternValuePresentValue"
}

func (m *_BACnetConstructedDataDateTimePatternValuePresentValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDateTimePatternValuePresentValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDateTimePatternValuePresentValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataDateTimePatternValuePresentValue BACnetConstructedDataDateTimePatternValuePresentValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDateTimePatternValuePresentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDateTimePatternValuePresentValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	presentValue, err := ReadSimpleField[BACnetDateTime](ctx, "presentValue", ReadComplex[BACnetDateTime](BACnetDateTimeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'presentValue' field"))
	}
	m.PresentValue = presentValue

	actualValue, err := ReadVirtualField[BACnetDateTime](ctx, "actualValue", (*BACnetDateTime)(nil), presentValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDateTimePatternValuePresentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDateTimePatternValuePresentValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDateTimePatternValuePresentValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDateTimePatternValuePresentValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDateTimePatternValuePresentValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDateTimePatternValuePresentValue")
		}

		if err := WriteSimpleField[BACnetDateTime](ctx, "presentValue", m.GetPresentValue(), WriteComplex[BACnetDateTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'presentValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDateTimePatternValuePresentValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDateTimePatternValuePresentValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDateTimePatternValuePresentValue) IsBACnetConstructedDataDateTimePatternValuePresentValue() {
}

func (m *_BACnetConstructedDataDateTimePatternValuePresentValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataDateTimePatternValuePresentValue) deepCopy() *_BACnetConstructedDataDateTimePatternValuePresentValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataDateTimePatternValuePresentValueCopy := &_BACnetConstructedDataDateTimePatternValuePresentValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetDateTime](m.PresentValue),
	}
	_BACnetConstructedDataDateTimePatternValuePresentValueCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataDateTimePatternValuePresentValueCopy
}

func (m *_BACnetConstructedDataDateTimePatternValuePresentValue) String() string {
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
