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

// BACnetConstructedDataIntegerValueLowLimit is the corresponding interface of BACnetConstructedDataIntegerValueLowLimit
type BACnetConstructedDataIntegerValueLowLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetLowLimit returns LowLimit (property field)
	GetLowLimit() BACnetApplicationTagSignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagSignedInteger
	// IsBACnetConstructedDataIntegerValueLowLimit is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataIntegerValueLowLimit()
	// CreateBuilder creates a BACnetConstructedDataIntegerValueLowLimitBuilder
	CreateBACnetConstructedDataIntegerValueLowLimitBuilder() BACnetConstructedDataIntegerValueLowLimitBuilder
}

// _BACnetConstructedDataIntegerValueLowLimit is the data-structure of this message
type _BACnetConstructedDataIntegerValueLowLimit struct {
	BACnetConstructedDataContract
	LowLimit BACnetApplicationTagSignedInteger
}

var _ BACnetConstructedDataIntegerValueLowLimit = (*_BACnetConstructedDataIntegerValueLowLimit)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataIntegerValueLowLimit)(nil)

// NewBACnetConstructedDataIntegerValueLowLimit factory function for _BACnetConstructedDataIntegerValueLowLimit
func NewBACnetConstructedDataIntegerValueLowLimit(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, lowLimit BACnetApplicationTagSignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIntegerValueLowLimit {
	if lowLimit == nil {
		panic("lowLimit of type BACnetApplicationTagSignedInteger for BACnetConstructedDataIntegerValueLowLimit must not be nil")
	}
	_result := &_BACnetConstructedDataIntegerValueLowLimit{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LowLimit:                      lowLimit,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataIntegerValueLowLimitBuilder is a builder for BACnetConstructedDataIntegerValueLowLimit
type BACnetConstructedDataIntegerValueLowLimitBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(lowLimit BACnetApplicationTagSignedInteger) BACnetConstructedDataIntegerValueLowLimitBuilder
	// WithLowLimit adds LowLimit (property field)
	WithLowLimit(BACnetApplicationTagSignedInteger) BACnetConstructedDataIntegerValueLowLimitBuilder
	// WithLowLimitBuilder adds LowLimit (property field) which is build by the builder
	WithLowLimitBuilder(func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetConstructedDataIntegerValueLowLimitBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataIntegerValueLowLimit or returns an error if something is wrong
	Build() (BACnetConstructedDataIntegerValueLowLimit, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataIntegerValueLowLimit
}

// NewBACnetConstructedDataIntegerValueLowLimitBuilder() creates a BACnetConstructedDataIntegerValueLowLimitBuilder
func NewBACnetConstructedDataIntegerValueLowLimitBuilder() BACnetConstructedDataIntegerValueLowLimitBuilder {
	return &_BACnetConstructedDataIntegerValueLowLimitBuilder{_BACnetConstructedDataIntegerValueLowLimit: new(_BACnetConstructedDataIntegerValueLowLimit)}
}

type _BACnetConstructedDataIntegerValueLowLimitBuilder struct {
	*_BACnetConstructedDataIntegerValueLowLimit

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataIntegerValueLowLimitBuilder) = (*_BACnetConstructedDataIntegerValueLowLimitBuilder)(nil)

func (b *_BACnetConstructedDataIntegerValueLowLimitBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataIntegerValueLowLimit
}

func (b *_BACnetConstructedDataIntegerValueLowLimitBuilder) WithMandatoryFields(lowLimit BACnetApplicationTagSignedInteger) BACnetConstructedDataIntegerValueLowLimitBuilder {
	return b.WithLowLimit(lowLimit)
}

func (b *_BACnetConstructedDataIntegerValueLowLimitBuilder) WithLowLimit(lowLimit BACnetApplicationTagSignedInteger) BACnetConstructedDataIntegerValueLowLimitBuilder {
	b.LowLimit = lowLimit
	return b
}

func (b *_BACnetConstructedDataIntegerValueLowLimitBuilder) WithLowLimitBuilder(builderSupplier func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetConstructedDataIntegerValueLowLimitBuilder {
	builder := builderSupplier(b.LowLimit.CreateBACnetApplicationTagSignedIntegerBuilder())
	var err error
	b.LowLimit, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagSignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataIntegerValueLowLimitBuilder) Build() (BACnetConstructedDataIntegerValueLowLimit, error) {
	if b.LowLimit == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'lowLimit' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataIntegerValueLowLimit.deepCopy(), nil
}

func (b *_BACnetConstructedDataIntegerValueLowLimitBuilder) MustBuild() BACnetConstructedDataIntegerValueLowLimit {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataIntegerValueLowLimitBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataIntegerValueLowLimitBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataIntegerValueLowLimitBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataIntegerValueLowLimitBuilder().(*_BACnetConstructedDataIntegerValueLowLimitBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataIntegerValueLowLimitBuilder creates a BACnetConstructedDataIntegerValueLowLimitBuilder
func (b *_BACnetConstructedDataIntegerValueLowLimit) CreateBACnetConstructedDataIntegerValueLowLimitBuilder() BACnetConstructedDataIntegerValueLowLimitBuilder {
	if b == nil {
		return NewBACnetConstructedDataIntegerValueLowLimitBuilder()
	}
	return &_BACnetConstructedDataIntegerValueLowLimitBuilder{_BACnetConstructedDataIntegerValueLowLimit: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueLowLimit) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_INTEGER_VALUE
}

func (m *_BACnetConstructedDataIntegerValueLowLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOW_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIntegerValueLowLimit) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueLowLimit) GetLowLimit() BACnetApplicationTagSignedInteger {
	return m.LowLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueLowLimit) GetActualValue() BACnetApplicationTagSignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagSignedInteger(m.GetLowLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIntegerValueLowLimit(structType any) BACnetConstructedDataIntegerValueLowLimit {
	if casted, ok := structType.(BACnetConstructedDataIntegerValueLowLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIntegerValueLowLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIntegerValueLowLimit) GetTypeName() string {
	return "BACnetConstructedDataIntegerValueLowLimit"
}

func (m *_BACnetConstructedDataIntegerValueLowLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (lowLimit)
	lengthInBits += m.LowLimit.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIntegerValueLowLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataIntegerValueLowLimit) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataIntegerValueLowLimit BACnetConstructedDataIntegerValueLowLimit, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIntegerValueLowLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIntegerValueLowLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lowLimit, err := ReadSimpleField[BACnetApplicationTagSignedInteger](ctx, "lowLimit", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagSignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lowLimit' field"))
	}
	m.LowLimit = lowLimit

	actualValue, err := ReadVirtualField[BACnetApplicationTagSignedInteger](ctx, "actualValue", (*BACnetApplicationTagSignedInteger)(nil), lowLimit)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIntegerValueLowLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIntegerValueLowLimit")
	}

	return m, nil
}

func (m *_BACnetConstructedDataIntegerValueLowLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIntegerValueLowLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIntegerValueLowLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIntegerValueLowLimit")
		}

		if err := WriteSimpleField[BACnetApplicationTagSignedInteger](ctx, "lowLimit", m.GetLowLimit(), WriteComplex[BACnetApplicationTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lowLimit' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIntegerValueLowLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIntegerValueLowLimit")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIntegerValueLowLimit) IsBACnetConstructedDataIntegerValueLowLimit() {}

func (m *_BACnetConstructedDataIntegerValueLowLimit) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataIntegerValueLowLimit) deepCopy() *_BACnetConstructedDataIntegerValueLowLimit {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataIntegerValueLowLimitCopy := &_BACnetConstructedDataIntegerValueLowLimit{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagSignedInteger](m.LowLimit),
	}
	_BACnetConstructedDataIntegerValueLowLimitCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataIntegerValueLowLimitCopy
}

func (m *_BACnetConstructedDataIntegerValueLowLimit) String() string {
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
