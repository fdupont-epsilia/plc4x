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

// BACnetConstructedDataUsesRemaining is the corresponding interface of BACnetConstructedDataUsesRemaining
type BACnetConstructedDataUsesRemaining interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetUsesRemaining returns UsesRemaining (property field)
	GetUsesRemaining() BACnetApplicationTagSignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagSignedInteger
	// IsBACnetConstructedDataUsesRemaining is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataUsesRemaining()
	// CreateBuilder creates a BACnetConstructedDataUsesRemainingBuilder
	CreateBACnetConstructedDataUsesRemainingBuilder() BACnetConstructedDataUsesRemainingBuilder
}

// _BACnetConstructedDataUsesRemaining is the data-structure of this message
type _BACnetConstructedDataUsesRemaining struct {
	BACnetConstructedDataContract
	UsesRemaining BACnetApplicationTagSignedInteger
}

var _ BACnetConstructedDataUsesRemaining = (*_BACnetConstructedDataUsesRemaining)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataUsesRemaining)(nil)

// NewBACnetConstructedDataUsesRemaining factory function for _BACnetConstructedDataUsesRemaining
func NewBACnetConstructedDataUsesRemaining(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, usesRemaining BACnetApplicationTagSignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataUsesRemaining {
	if usesRemaining == nil {
		panic("usesRemaining of type BACnetApplicationTagSignedInteger for BACnetConstructedDataUsesRemaining must not be nil")
	}
	_result := &_BACnetConstructedDataUsesRemaining{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		UsesRemaining:                 usesRemaining,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataUsesRemainingBuilder is a builder for BACnetConstructedDataUsesRemaining
type BACnetConstructedDataUsesRemainingBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(usesRemaining BACnetApplicationTagSignedInteger) BACnetConstructedDataUsesRemainingBuilder
	// WithUsesRemaining adds UsesRemaining (property field)
	WithUsesRemaining(BACnetApplicationTagSignedInteger) BACnetConstructedDataUsesRemainingBuilder
	// WithUsesRemainingBuilder adds UsesRemaining (property field) which is build by the builder
	WithUsesRemainingBuilder(func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetConstructedDataUsesRemainingBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataUsesRemaining or returns an error if something is wrong
	Build() (BACnetConstructedDataUsesRemaining, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataUsesRemaining
}

// NewBACnetConstructedDataUsesRemainingBuilder() creates a BACnetConstructedDataUsesRemainingBuilder
func NewBACnetConstructedDataUsesRemainingBuilder() BACnetConstructedDataUsesRemainingBuilder {
	return &_BACnetConstructedDataUsesRemainingBuilder{_BACnetConstructedDataUsesRemaining: new(_BACnetConstructedDataUsesRemaining)}
}

type _BACnetConstructedDataUsesRemainingBuilder struct {
	*_BACnetConstructedDataUsesRemaining

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataUsesRemainingBuilder) = (*_BACnetConstructedDataUsesRemainingBuilder)(nil)

func (b *_BACnetConstructedDataUsesRemainingBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataUsesRemaining
}

func (b *_BACnetConstructedDataUsesRemainingBuilder) WithMandatoryFields(usesRemaining BACnetApplicationTagSignedInteger) BACnetConstructedDataUsesRemainingBuilder {
	return b.WithUsesRemaining(usesRemaining)
}

func (b *_BACnetConstructedDataUsesRemainingBuilder) WithUsesRemaining(usesRemaining BACnetApplicationTagSignedInteger) BACnetConstructedDataUsesRemainingBuilder {
	b.UsesRemaining = usesRemaining
	return b
}

func (b *_BACnetConstructedDataUsesRemainingBuilder) WithUsesRemainingBuilder(builderSupplier func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetConstructedDataUsesRemainingBuilder {
	builder := builderSupplier(b.UsesRemaining.CreateBACnetApplicationTagSignedIntegerBuilder())
	var err error
	b.UsesRemaining, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagSignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataUsesRemainingBuilder) Build() (BACnetConstructedDataUsesRemaining, error) {
	if b.UsesRemaining == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'usesRemaining' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataUsesRemaining.deepCopy(), nil
}

func (b *_BACnetConstructedDataUsesRemainingBuilder) MustBuild() BACnetConstructedDataUsesRemaining {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataUsesRemainingBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataUsesRemainingBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataUsesRemainingBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataUsesRemainingBuilder().(*_BACnetConstructedDataUsesRemainingBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataUsesRemainingBuilder creates a BACnetConstructedDataUsesRemainingBuilder
func (b *_BACnetConstructedDataUsesRemaining) CreateBACnetConstructedDataUsesRemainingBuilder() BACnetConstructedDataUsesRemainingBuilder {
	if b == nil {
		return NewBACnetConstructedDataUsesRemainingBuilder()
	}
	return &_BACnetConstructedDataUsesRemainingBuilder{_BACnetConstructedDataUsesRemaining: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataUsesRemaining) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataUsesRemaining) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_USES_REMAINING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataUsesRemaining) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataUsesRemaining) GetUsesRemaining() BACnetApplicationTagSignedInteger {
	return m.UsesRemaining
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataUsesRemaining) GetActualValue() BACnetApplicationTagSignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagSignedInteger(m.GetUsesRemaining())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataUsesRemaining(structType any) BACnetConstructedDataUsesRemaining {
	if casted, ok := structType.(BACnetConstructedDataUsesRemaining); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataUsesRemaining); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataUsesRemaining) GetTypeName() string {
	return "BACnetConstructedDataUsesRemaining"
}

func (m *_BACnetConstructedDataUsesRemaining) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (usesRemaining)
	lengthInBits += m.UsesRemaining.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataUsesRemaining) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataUsesRemaining) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataUsesRemaining BACnetConstructedDataUsesRemaining, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataUsesRemaining"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataUsesRemaining")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	usesRemaining, err := ReadSimpleField[BACnetApplicationTagSignedInteger](ctx, "usesRemaining", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagSignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'usesRemaining' field"))
	}
	m.UsesRemaining = usesRemaining

	actualValue, err := ReadVirtualField[BACnetApplicationTagSignedInteger](ctx, "actualValue", (*BACnetApplicationTagSignedInteger)(nil), usesRemaining)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataUsesRemaining"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataUsesRemaining")
	}

	return m, nil
}

func (m *_BACnetConstructedDataUsesRemaining) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataUsesRemaining) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataUsesRemaining"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataUsesRemaining")
		}

		if err := WriteSimpleField[BACnetApplicationTagSignedInteger](ctx, "usesRemaining", m.GetUsesRemaining(), WriteComplex[BACnetApplicationTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'usesRemaining' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataUsesRemaining"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataUsesRemaining")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataUsesRemaining) IsBACnetConstructedDataUsesRemaining() {}

func (m *_BACnetConstructedDataUsesRemaining) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataUsesRemaining) deepCopy() *_BACnetConstructedDataUsesRemaining {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataUsesRemainingCopy := &_BACnetConstructedDataUsesRemaining{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagSignedInteger](m.UsesRemaining),
	}
	_BACnetConstructedDataUsesRemainingCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataUsesRemainingCopy
}

func (m *_BACnetConstructedDataUsesRemaining) String() string {
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
