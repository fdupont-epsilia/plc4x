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

// BACnetConstructedDataMaximumValueTimestamp is the corresponding interface of BACnetConstructedDataMaximumValueTimestamp
type BACnetConstructedDataMaximumValueTimestamp interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetMaximumValueTimestamp returns MaximumValueTimestamp (property field)
	GetMaximumValueTimestamp() BACnetDateTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDateTime
	// IsBACnetConstructedDataMaximumValueTimestamp is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataMaximumValueTimestamp()
	// CreateBuilder creates a BACnetConstructedDataMaximumValueTimestampBuilder
	CreateBACnetConstructedDataMaximumValueTimestampBuilder() BACnetConstructedDataMaximumValueTimestampBuilder
}

// _BACnetConstructedDataMaximumValueTimestamp is the data-structure of this message
type _BACnetConstructedDataMaximumValueTimestamp struct {
	BACnetConstructedDataContract
	MaximumValueTimestamp BACnetDateTime
}

var _ BACnetConstructedDataMaximumValueTimestamp = (*_BACnetConstructedDataMaximumValueTimestamp)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataMaximumValueTimestamp)(nil)

// NewBACnetConstructedDataMaximumValueTimestamp factory function for _BACnetConstructedDataMaximumValueTimestamp
func NewBACnetConstructedDataMaximumValueTimestamp(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, maximumValueTimestamp BACnetDateTime, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMaximumValueTimestamp {
	if maximumValueTimestamp == nil {
		panic("maximumValueTimestamp of type BACnetDateTime for BACnetConstructedDataMaximumValueTimestamp must not be nil")
	}
	_result := &_BACnetConstructedDataMaximumValueTimestamp{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		MaximumValueTimestamp:         maximumValueTimestamp,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataMaximumValueTimestampBuilder is a builder for BACnetConstructedDataMaximumValueTimestamp
type BACnetConstructedDataMaximumValueTimestampBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(maximumValueTimestamp BACnetDateTime) BACnetConstructedDataMaximumValueTimestampBuilder
	// WithMaximumValueTimestamp adds MaximumValueTimestamp (property field)
	WithMaximumValueTimestamp(BACnetDateTime) BACnetConstructedDataMaximumValueTimestampBuilder
	// WithMaximumValueTimestampBuilder adds MaximumValueTimestamp (property field) which is build by the builder
	WithMaximumValueTimestampBuilder(func(BACnetDateTimeBuilder) BACnetDateTimeBuilder) BACnetConstructedDataMaximumValueTimestampBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataMaximumValueTimestamp or returns an error if something is wrong
	Build() (BACnetConstructedDataMaximumValueTimestamp, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataMaximumValueTimestamp
}

// NewBACnetConstructedDataMaximumValueTimestampBuilder() creates a BACnetConstructedDataMaximumValueTimestampBuilder
func NewBACnetConstructedDataMaximumValueTimestampBuilder() BACnetConstructedDataMaximumValueTimestampBuilder {
	return &_BACnetConstructedDataMaximumValueTimestampBuilder{_BACnetConstructedDataMaximumValueTimestamp: new(_BACnetConstructedDataMaximumValueTimestamp)}
}

type _BACnetConstructedDataMaximumValueTimestampBuilder struct {
	*_BACnetConstructedDataMaximumValueTimestamp

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataMaximumValueTimestampBuilder) = (*_BACnetConstructedDataMaximumValueTimestampBuilder)(nil)

func (b *_BACnetConstructedDataMaximumValueTimestampBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataMaximumValueTimestamp
}

func (b *_BACnetConstructedDataMaximumValueTimestampBuilder) WithMandatoryFields(maximumValueTimestamp BACnetDateTime) BACnetConstructedDataMaximumValueTimestampBuilder {
	return b.WithMaximumValueTimestamp(maximumValueTimestamp)
}

func (b *_BACnetConstructedDataMaximumValueTimestampBuilder) WithMaximumValueTimestamp(maximumValueTimestamp BACnetDateTime) BACnetConstructedDataMaximumValueTimestampBuilder {
	b.MaximumValueTimestamp = maximumValueTimestamp
	return b
}

func (b *_BACnetConstructedDataMaximumValueTimestampBuilder) WithMaximumValueTimestampBuilder(builderSupplier func(BACnetDateTimeBuilder) BACnetDateTimeBuilder) BACnetConstructedDataMaximumValueTimestampBuilder {
	builder := builderSupplier(b.MaximumValueTimestamp.CreateBACnetDateTimeBuilder())
	var err error
	b.MaximumValueTimestamp, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDateTimeBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataMaximumValueTimestampBuilder) Build() (BACnetConstructedDataMaximumValueTimestamp, error) {
	if b.MaximumValueTimestamp == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'maximumValueTimestamp' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataMaximumValueTimestamp.deepCopy(), nil
}

func (b *_BACnetConstructedDataMaximumValueTimestampBuilder) MustBuild() BACnetConstructedDataMaximumValueTimestamp {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataMaximumValueTimestampBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataMaximumValueTimestampBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataMaximumValueTimestampBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataMaximumValueTimestampBuilder().(*_BACnetConstructedDataMaximumValueTimestampBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataMaximumValueTimestampBuilder creates a BACnetConstructedDataMaximumValueTimestampBuilder
func (b *_BACnetConstructedDataMaximumValueTimestamp) CreateBACnetConstructedDataMaximumValueTimestampBuilder() BACnetConstructedDataMaximumValueTimestampBuilder {
	if b == nil {
		return NewBACnetConstructedDataMaximumValueTimestampBuilder()
	}
	return &_BACnetConstructedDataMaximumValueTimestampBuilder{_BACnetConstructedDataMaximumValueTimestamp: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMaximumValueTimestamp) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMaximumValueTimestamp) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAXIMUM_VALUE_TIMESTAMP
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMaximumValueTimestamp) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMaximumValueTimestamp) GetMaximumValueTimestamp() BACnetDateTime {
	return m.MaximumValueTimestamp
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMaximumValueTimestamp) GetActualValue() BACnetDateTime {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDateTime(m.GetMaximumValueTimestamp())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMaximumValueTimestamp(structType any) BACnetConstructedDataMaximumValueTimestamp {
	if casted, ok := structType.(BACnetConstructedDataMaximumValueTimestamp); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMaximumValueTimestamp); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMaximumValueTimestamp) GetTypeName() string {
	return "BACnetConstructedDataMaximumValueTimestamp"
}

func (m *_BACnetConstructedDataMaximumValueTimestamp) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (maximumValueTimestamp)
	lengthInBits += m.MaximumValueTimestamp.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMaximumValueTimestamp) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataMaximumValueTimestamp) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataMaximumValueTimestamp BACnetConstructedDataMaximumValueTimestamp, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMaximumValueTimestamp"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMaximumValueTimestamp")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	maximumValueTimestamp, err := ReadSimpleField[BACnetDateTime](ctx, "maximumValueTimestamp", ReadComplex[BACnetDateTime](BACnetDateTimeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maximumValueTimestamp' field"))
	}
	m.MaximumValueTimestamp = maximumValueTimestamp

	actualValue, err := ReadVirtualField[BACnetDateTime](ctx, "actualValue", (*BACnetDateTime)(nil), maximumValueTimestamp)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMaximumValueTimestamp"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMaximumValueTimestamp")
	}

	return m, nil
}

func (m *_BACnetConstructedDataMaximumValueTimestamp) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMaximumValueTimestamp) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMaximumValueTimestamp"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMaximumValueTimestamp")
		}

		if err := WriteSimpleField[BACnetDateTime](ctx, "maximumValueTimestamp", m.GetMaximumValueTimestamp(), WriteComplex[BACnetDateTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'maximumValueTimestamp' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMaximumValueTimestamp"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMaximumValueTimestamp")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMaximumValueTimestamp) IsBACnetConstructedDataMaximumValueTimestamp() {
}

func (m *_BACnetConstructedDataMaximumValueTimestamp) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataMaximumValueTimestamp) deepCopy() *_BACnetConstructedDataMaximumValueTimestamp {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataMaximumValueTimestampCopy := &_BACnetConstructedDataMaximumValueTimestamp{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetDateTime](m.MaximumValueTimestamp),
	}
	_BACnetConstructedDataMaximumValueTimestampCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataMaximumValueTimestampCopy
}

func (m *_BACnetConstructedDataMaximumValueTimestamp) String() string {
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
