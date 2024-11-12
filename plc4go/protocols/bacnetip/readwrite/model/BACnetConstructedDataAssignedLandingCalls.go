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

// BACnetConstructedDataAssignedLandingCalls is the corresponding interface of BACnetConstructedDataAssignedLandingCalls
type BACnetConstructedDataAssignedLandingCalls interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetAssignedLandingCalls returns AssignedLandingCalls (property field)
	GetAssignedLandingCalls() []BACnetAssignedLandingCalls
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// IsBACnetConstructedDataAssignedLandingCalls is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAssignedLandingCalls()
	// CreateBuilder creates a BACnetConstructedDataAssignedLandingCallsBuilder
	CreateBACnetConstructedDataAssignedLandingCallsBuilder() BACnetConstructedDataAssignedLandingCallsBuilder
}

// _BACnetConstructedDataAssignedLandingCalls is the data-structure of this message
type _BACnetConstructedDataAssignedLandingCalls struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	AssignedLandingCalls []BACnetAssignedLandingCalls
}

var _ BACnetConstructedDataAssignedLandingCalls = (*_BACnetConstructedDataAssignedLandingCalls)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAssignedLandingCalls)(nil)

// NewBACnetConstructedDataAssignedLandingCalls factory function for _BACnetConstructedDataAssignedLandingCalls
func NewBACnetConstructedDataAssignedLandingCalls(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfDataElements BACnetApplicationTagUnsignedInteger, assignedLandingCalls []BACnetAssignedLandingCalls, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAssignedLandingCalls {
	_result := &_BACnetConstructedDataAssignedLandingCalls{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		AssignedLandingCalls:          assignedLandingCalls,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAssignedLandingCallsBuilder is a builder for BACnetConstructedDataAssignedLandingCalls
type BACnetConstructedDataAssignedLandingCallsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(assignedLandingCalls []BACnetAssignedLandingCalls) BACnetConstructedDataAssignedLandingCallsBuilder
	// WithNumberOfDataElements adds NumberOfDataElements (property field)
	WithOptionalNumberOfDataElements(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAssignedLandingCallsBuilder
	// WithOptionalNumberOfDataElementsBuilder adds NumberOfDataElements (property field) which is build by the builder
	WithOptionalNumberOfDataElementsBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataAssignedLandingCallsBuilder
	// WithAssignedLandingCalls adds AssignedLandingCalls (property field)
	WithAssignedLandingCalls(...BACnetAssignedLandingCalls) BACnetConstructedDataAssignedLandingCallsBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataAssignedLandingCalls or returns an error if something is wrong
	Build() (BACnetConstructedDataAssignedLandingCalls, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAssignedLandingCalls
}

// NewBACnetConstructedDataAssignedLandingCallsBuilder() creates a BACnetConstructedDataAssignedLandingCallsBuilder
func NewBACnetConstructedDataAssignedLandingCallsBuilder() BACnetConstructedDataAssignedLandingCallsBuilder {
	return &_BACnetConstructedDataAssignedLandingCallsBuilder{_BACnetConstructedDataAssignedLandingCalls: new(_BACnetConstructedDataAssignedLandingCalls)}
}

type _BACnetConstructedDataAssignedLandingCallsBuilder struct {
	*_BACnetConstructedDataAssignedLandingCalls

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataAssignedLandingCallsBuilder) = (*_BACnetConstructedDataAssignedLandingCallsBuilder)(nil)

func (b *_BACnetConstructedDataAssignedLandingCallsBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataAssignedLandingCalls
}

func (b *_BACnetConstructedDataAssignedLandingCallsBuilder) WithMandatoryFields(assignedLandingCalls []BACnetAssignedLandingCalls) BACnetConstructedDataAssignedLandingCallsBuilder {
	return b.WithAssignedLandingCalls(assignedLandingCalls...)
}

func (b *_BACnetConstructedDataAssignedLandingCallsBuilder) WithOptionalNumberOfDataElements(numberOfDataElements BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAssignedLandingCallsBuilder {
	b.NumberOfDataElements = numberOfDataElements
	return b
}

func (b *_BACnetConstructedDataAssignedLandingCallsBuilder) WithOptionalNumberOfDataElementsBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataAssignedLandingCallsBuilder {
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

func (b *_BACnetConstructedDataAssignedLandingCallsBuilder) WithAssignedLandingCalls(assignedLandingCalls ...BACnetAssignedLandingCalls) BACnetConstructedDataAssignedLandingCallsBuilder {
	b.AssignedLandingCalls = assignedLandingCalls
	return b
}

func (b *_BACnetConstructedDataAssignedLandingCallsBuilder) Build() (BACnetConstructedDataAssignedLandingCalls, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataAssignedLandingCalls.deepCopy(), nil
}

func (b *_BACnetConstructedDataAssignedLandingCallsBuilder) MustBuild() BACnetConstructedDataAssignedLandingCalls {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataAssignedLandingCallsBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataAssignedLandingCallsBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataAssignedLandingCallsBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataAssignedLandingCallsBuilder().(*_BACnetConstructedDataAssignedLandingCallsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataAssignedLandingCallsBuilder creates a BACnetConstructedDataAssignedLandingCallsBuilder
func (b *_BACnetConstructedDataAssignedLandingCalls) CreateBACnetConstructedDataAssignedLandingCallsBuilder() BACnetConstructedDataAssignedLandingCallsBuilder {
	if b == nil {
		return NewBACnetConstructedDataAssignedLandingCallsBuilder()
	}
	return &_BACnetConstructedDataAssignedLandingCallsBuilder{_BACnetConstructedDataAssignedLandingCalls: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAssignedLandingCalls) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAssignedLandingCalls) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ASSIGNED_LANDING_CALLS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAssignedLandingCalls) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAssignedLandingCalls) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataAssignedLandingCalls) GetAssignedLandingCalls() []BACnetAssignedLandingCalls {
	return m.AssignedLandingCalls
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAssignedLandingCalls) GetZero() uint64 {
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
func CastBACnetConstructedDataAssignedLandingCalls(structType any) BACnetConstructedDataAssignedLandingCalls {
	if casted, ok := structType.(BACnetConstructedDataAssignedLandingCalls); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAssignedLandingCalls); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAssignedLandingCalls) GetTypeName() string {
	return "BACnetConstructedDataAssignedLandingCalls"
}

func (m *_BACnetConstructedDataAssignedLandingCalls) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.AssignedLandingCalls) > 0 {
		for _, element := range m.AssignedLandingCalls {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataAssignedLandingCalls) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAssignedLandingCalls) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAssignedLandingCalls BACnetConstructedDataAssignedLandingCalls, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAssignedLandingCalls"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAssignedLandingCalls")
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

	assignedLandingCalls, err := ReadTerminatedArrayField[BACnetAssignedLandingCalls](ctx, "assignedLandingCalls", ReadComplex[BACnetAssignedLandingCalls](BACnetAssignedLandingCallsParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'assignedLandingCalls' field"))
	}
	m.AssignedLandingCalls = assignedLandingCalls

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAssignedLandingCalls"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAssignedLandingCalls")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAssignedLandingCalls) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAssignedLandingCalls) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAssignedLandingCalls"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAssignedLandingCalls")
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

		if err := WriteComplexTypeArrayField(ctx, "assignedLandingCalls", m.GetAssignedLandingCalls(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'assignedLandingCalls' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAssignedLandingCalls"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAssignedLandingCalls")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAssignedLandingCalls) IsBACnetConstructedDataAssignedLandingCalls() {}

func (m *_BACnetConstructedDataAssignedLandingCalls) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAssignedLandingCalls) deepCopy() *_BACnetConstructedDataAssignedLandingCalls {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAssignedLandingCallsCopy := &_BACnetConstructedDataAssignedLandingCalls{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.NumberOfDataElements),
		utils.DeepCopySlice[BACnetAssignedLandingCalls, BACnetAssignedLandingCalls](m.AssignedLandingCalls),
	}
	_BACnetConstructedDataAssignedLandingCallsCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAssignedLandingCallsCopy
}

func (m *_BACnetConstructedDataAssignedLandingCalls) String() string {
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
