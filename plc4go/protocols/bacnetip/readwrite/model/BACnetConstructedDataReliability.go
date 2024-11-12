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

// BACnetConstructedDataReliability is the corresponding interface of BACnetConstructedDataReliability
type BACnetConstructedDataReliability interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetReliability returns Reliability (property field)
	GetReliability() BACnetReliabilityTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetReliabilityTagged
	// IsBACnetConstructedDataReliability is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataReliability()
	// CreateBuilder creates a BACnetConstructedDataReliabilityBuilder
	CreateBACnetConstructedDataReliabilityBuilder() BACnetConstructedDataReliabilityBuilder
}

// _BACnetConstructedDataReliability is the data-structure of this message
type _BACnetConstructedDataReliability struct {
	BACnetConstructedDataContract
	Reliability BACnetReliabilityTagged
}

var _ BACnetConstructedDataReliability = (*_BACnetConstructedDataReliability)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataReliability)(nil)

// NewBACnetConstructedDataReliability factory function for _BACnetConstructedDataReliability
func NewBACnetConstructedDataReliability(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, reliability BACnetReliabilityTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataReliability {
	if reliability == nil {
		panic("reliability of type BACnetReliabilityTagged for BACnetConstructedDataReliability must not be nil")
	}
	_result := &_BACnetConstructedDataReliability{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Reliability:                   reliability,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataReliabilityBuilder is a builder for BACnetConstructedDataReliability
type BACnetConstructedDataReliabilityBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(reliability BACnetReliabilityTagged) BACnetConstructedDataReliabilityBuilder
	// WithReliability adds Reliability (property field)
	WithReliability(BACnetReliabilityTagged) BACnetConstructedDataReliabilityBuilder
	// WithReliabilityBuilder adds Reliability (property field) which is build by the builder
	WithReliabilityBuilder(func(BACnetReliabilityTaggedBuilder) BACnetReliabilityTaggedBuilder) BACnetConstructedDataReliabilityBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataReliability or returns an error if something is wrong
	Build() (BACnetConstructedDataReliability, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataReliability
}

// NewBACnetConstructedDataReliabilityBuilder() creates a BACnetConstructedDataReliabilityBuilder
func NewBACnetConstructedDataReliabilityBuilder() BACnetConstructedDataReliabilityBuilder {
	return &_BACnetConstructedDataReliabilityBuilder{_BACnetConstructedDataReliability: new(_BACnetConstructedDataReliability)}
}

type _BACnetConstructedDataReliabilityBuilder struct {
	*_BACnetConstructedDataReliability

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataReliabilityBuilder) = (*_BACnetConstructedDataReliabilityBuilder)(nil)

func (b *_BACnetConstructedDataReliabilityBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataReliability
}

func (b *_BACnetConstructedDataReliabilityBuilder) WithMandatoryFields(reliability BACnetReliabilityTagged) BACnetConstructedDataReliabilityBuilder {
	return b.WithReliability(reliability)
}

func (b *_BACnetConstructedDataReliabilityBuilder) WithReliability(reliability BACnetReliabilityTagged) BACnetConstructedDataReliabilityBuilder {
	b.Reliability = reliability
	return b
}

func (b *_BACnetConstructedDataReliabilityBuilder) WithReliabilityBuilder(builderSupplier func(BACnetReliabilityTaggedBuilder) BACnetReliabilityTaggedBuilder) BACnetConstructedDataReliabilityBuilder {
	builder := builderSupplier(b.Reliability.CreateBACnetReliabilityTaggedBuilder())
	var err error
	b.Reliability, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetReliabilityTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataReliabilityBuilder) Build() (BACnetConstructedDataReliability, error) {
	if b.Reliability == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'reliability' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataReliability.deepCopy(), nil
}

func (b *_BACnetConstructedDataReliabilityBuilder) MustBuild() BACnetConstructedDataReliability {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataReliabilityBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataReliabilityBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataReliabilityBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataReliabilityBuilder().(*_BACnetConstructedDataReliabilityBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataReliabilityBuilder creates a BACnetConstructedDataReliabilityBuilder
func (b *_BACnetConstructedDataReliability) CreateBACnetConstructedDataReliabilityBuilder() BACnetConstructedDataReliabilityBuilder {
	if b == nil {
		return NewBACnetConstructedDataReliabilityBuilder()
	}
	return &_BACnetConstructedDataReliabilityBuilder{_BACnetConstructedDataReliability: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataReliability) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataReliability) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RELIABILITY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataReliability) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataReliability) GetReliability() BACnetReliabilityTagged {
	return m.Reliability
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataReliability) GetActualValue() BACnetReliabilityTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetReliabilityTagged(m.GetReliability())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataReliability(structType any) BACnetConstructedDataReliability {
	if casted, ok := structType.(BACnetConstructedDataReliability); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataReliability); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataReliability) GetTypeName() string {
	return "BACnetConstructedDataReliability"
}

func (m *_BACnetConstructedDataReliability) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (reliability)
	lengthInBits += m.Reliability.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataReliability) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataReliability) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataReliability BACnetConstructedDataReliability, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataReliability"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataReliability")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reliability, err := ReadSimpleField[BACnetReliabilityTagged](ctx, "reliability", ReadComplex[BACnetReliabilityTagged](BACnetReliabilityTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reliability' field"))
	}
	m.Reliability = reliability

	actualValue, err := ReadVirtualField[BACnetReliabilityTagged](ctx, "actualValue", (*BACnetReliabilityTagged)(nil), reliability)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataReliability"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataReliability")
	}

	return m, nil
}

func (m *_BACnetConstructedDataReliability) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataReliability) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataReliability"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataReliability")
		}

		if err := WriteSimpleField[BACnetReliabilityTagged](ctx, "reliability", m.GetReliability(), WriteComplex[BACnetReliabilityTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'reliability' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataReliability"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataReliability")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataReliability) IsBACnetConstructedDataReliability() {}

func (m *_BACnetConstructedDataReliability) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataReliability) deepCopy() *_BACnetConstructedDataReliability {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataReliabilityCopy := &_BACnetConstructedDataReliability{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetReliabilityTagged](m.Reliability),
	}
	_BACnetConstructedDataReliabilityCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataReliabilityCopy
}

func (m *_BACnetConstructedDataReliability) String() string {
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
