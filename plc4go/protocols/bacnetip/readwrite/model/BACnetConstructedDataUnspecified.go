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

// BACnetConstructedDataUnspecified is the corresponding interface of BACnetConstructedDataUnspecified
type BACnetConstructedDataUnspecified interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetData returns Data (property field)
	GetData() []BACnetConstructedDataElement
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// IsBACnetConstructedDataUnspecified is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataUnspecified()
	// CreateBuilder creates a BACnetConstructedDataUnspecifiedBuilder
	CreateBACnetConstructedDataUnspecifiedBuilder() BACnetConstructedDataUnspecifiedBuilder
}

// _BACnetConstructedDataUnspecified is the data-structure of this message
type _BACnetConstructedDataUnspecified struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	Data                 []BACnetConstructedDataElement
}

var _ BACnetConstructedDataUnspecified = (*_BACnetConstructedDataUnspecified)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataUnspecified)(nil)

// NewBACnetConstructedDataUnspecified factory function for _BACnetConstructedDataUnspecified
func NewBACnetConstructedDataUnspecified(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfDataElements BACnetApplicationTagUnsignedInteger, data []BACnetConstructedDataElement, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataUnspecified {
	_result := &_BACnetConstructedDataUnspecified{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		Data:                          data,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataUnspecifiedBuilder is a builder for BACnetConstructedDataUnspecified
type BACnetConstructedDataUnspecifiedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(data []BACnetConstructedDataElement) BACnetConstructedDataUnspecifiedBuilder
	// WithNumberOfDataElements adds NumberOfDataElements (property field)
	WithOptionalNumberOfDataElements(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataUnspecifiedBuilder
	// WithOptionalNumberOfDataElementsBuilder adds NumberOfDataElements (property field) which is build by the builder
	WithOptionalNumberOfDataElementsBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataUnspecifiedBuilder
	// WithData adds Data (property field)
	WithData(...BACnetConstructedDataElement) BACnetConstructedDataUnspecifiedBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataUnspecified or returns an error if something is wrong
	Build() (BACnetConstructedDataUnspecified, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataUnspecified
}

// NewBACnetConstructedDataUnspecifiedBuilder() creates a BACnetConstructedDataUnspecifiedBuilder
func NewBACnetConstructedDataUnspecifiedBuilder() BACnetConstructedDataUnspecifiedBuilder {
	return &_BACnetConstructedDataUnspecifiedBuilder{_BACnetConstructedDataUnspecified: new(_BACnetConstructedDataUnspecified)}
}

type _BACnetConstructedDataUnspecifiedBuilder struct {
	*_BACnetConstructedDataUnspecified

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataUnspecifiedBuilder) = (*_BACnetConstructedDataUnspecifiedBuilder)(nil)

func (b *_BACnetConstructedDataUnspecifiedBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataUnspecified
}

func (b *_BACnetConstructedDataUnspecifiedBuilder) WithMandatoryFields(data []BACnetConstructedDataElement) BACnetConstructedDataUnspecifiedBuilder {
	return b.WithData(data...)
}

func (b *_BACnetConstructedDataUnspecifiedBuilder) WithOptionalNumberOfDataElements(numberOfDataElements BACnetApplicationTagUnsignedInteger) BACnetConstructedDataUnspecifiedBuilder {
	b.NumberOfDataElements = numberOfDataElements
	return b
}

func (b *_BACnetConstructedDataUnspecifiedBuilder) WithOptionalNumberOfDataElementsBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataUnspecifiedBuilder {
	builder := builderSupplier(b.NumberOfDataElements.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.NumberOfDataElements, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataUnspecifiedBuilder) WithData(data ...BACnetConstructedDataElement) BACnetConstructedDataUnspecifiedBuilder {
	b.Data = data
	return b
}

func (b *_BACnetConstructedDataUnspecifiedBuilder) Build() (BACnetConstructedDataUnspecified, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataUnspecified.deepCopy(), nil
}

func (b *_BACnetConstructedDataUnspecifiedBuilder) MustBuild() BACnetConstructedDataUnspecified {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataUnspecifiedBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataUnspecifiedBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataUnspecifiedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataUnspecifiedBuilder().(*_BACnetConstructedDataUnspecifiedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataUnspecifiedBuilder creates a BACnetConstructedDataUnspecifiedBuilder
func (b *_BACnetConstructedDataUnspecified) CreateBACnetConstructedDataUnspecifiedBuilder() BACnetConstructedDataUnspecifiedBuilder {
	if b == nil {
		return NewBACnetConstructedDataUnspecifiedBuilder()
	}
	return &_BACnetConstructedDataUnspecifiedBuilder{_BACnetConstructedDataUnspecified: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataUnspecified) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataUnspecified) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataUnspecified) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataUnspecified) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataUnspecified) GetData() []BACnetConstructedDataElement {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataUnspecified) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.GetNumberOfDataElements()
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataUnspecified(structType any) BACnetConstructedDataUnspecified {
	if casted, ok := structType.(BACnetConstructedDataUnspecified); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataUnspecified); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataUnspecified) GetTypeName() string {
	return "BACnetConstructedDataUnspecified"
}

func (m *_BACnetConstructedDataUnspecified) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.Data) > 0 {
		for _, element := range m.Data {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataUnspecified) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataUnspecified) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataUnspecified BACnetConstructedDataUnspecified, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataUnspecified"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataUnspecified")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zero, err := ReadVirtualField[uint64](ctx, "zero", (*uint64)(nil), uint64(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zero' field"))
	}
	_ = zero

	var numberOfDataElements BACnetApplicationTagUnsignedInteger
	_numberOfDataElements, err := ReadOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfDataElements' field"))
	}
	if _numberOfDataElements != nil {
		numberOfDataElements = *_numberOfDataElements
		m.NumberOfDataElements = numberOfDataElements
	}

	data, err := ReadTerminatedArrayField[BACnetConstructedDataElement](ctx, "data", ReadComplex[BACnetConstructedDataElement](BACnetConstructedDataElementParseWithBufferProducer((BACnetObjectType)(objectTypeArgument), (BACnetPropertyIdentifier)(propertyIdentifierArgument), (BACnetTagPayloadUnsignedInteger)(arrayIndexArgument)), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataUnspecified"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataUnspecified")
	}

	return m, nil
}

func (m *_BACnetConstructedDataUnspecified) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataUnspecified) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataUnspecified"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataUnspecified")
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

		if err := WriteComplexTypeArrayField(ctx, "data", m.GetData(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataUnspecified"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataUnspecified")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataUnspecified) IsBACnetConstructedDataUnspecified() {}

func (m *_BACnetConstructedDataUnspecified) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataUnspecified) deepCopy() *_BACnetConstructedDataUnspecified {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataUnspecifiedCopy := &_BACnetConstructedDataUnspecified{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.NumberOfDataElements),
		utils.DeepCopySlice[BACnetConstructedDataElement, BACnetConstructedDataElement](m.Data),
	}
	_BACnetConstructedDataUnspecifiedCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataUnspecifiedCopy
}

func (m *_BACnetConstructedDataUnspecified) String() string {
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
