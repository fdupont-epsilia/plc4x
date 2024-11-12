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

// BACnetConstructedDataRepresents is the corresponding interface of BACnetConstructedDataRepresents
type BACnetConstructedDataRepresents interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetRepresents returns Represents (property field)
	GetRepresents() BACnetDeviceObjectReference
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDeviceObjectReference
	// IsBACnetConstructedDataRepresents is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataRepresents()
	// CreateBuilder creates a BACnetConstructedDataRepresentsBuilder
	CreateBACnetConstructedDataRepresentsBuilder() BACnetConstructedDataRepresentsBuilder
}

// _BACnetConstructedDataRepresents is the data-structure of this message
type _BACnetConstructedDataRepresents struct {
	BACnetConstructedDataContract
	Represents BACnetDeviceObjectReference
}

var _ BACnetConstructedDataRepresents = (*_BACnetConstructedDataRepresents)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataRepresents)(nil)

// NewBACnetConstructedDataRepresents factory function for _BACnetConstructedDataRepresents
func NewBACnetConstructedDataRepresents(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, represents BACnetDeviceObjectReference, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataRepresents {
	if represents == nil {
		panic("represents of type BACnetDeviceObjectReference for BACnetConstructedDataRepresents must not be nil")
	}
	_result := &_BACnetConstructedDataRepresents{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Represents:                    represents,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataRepresentsBuilder is a builder for BACnetConstructedDataRepresents
type BACnetConstructedDataRepresentsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(represents BACnetDeviceObjectReference) BACnetConstructedDataRepresentsBuilder
	// WithRepresents adds Represents (property field)
	WithRepresents(BACnetDeviceObjectReference) BACnetConstructedDataRepresentsBuilder
	// WithRepresentsBuilder adds Represents (property field) which is build by the builder
	WithRepresentsBuilder(func(BACnetDeviceObjectReferenceBuilder) BACnetDeviceObjectReferenceBuilder) BACnetConstructedDataRepresentsBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataRepresents or returns an error if something is wrong
	Build() (BACnetConstructedDataRepresents, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataRepresents
}

// NewBACnetConstructedDataRepresentsBuilder() creates a BACnetConstructedDataRepresentsBuilder
func NewBACnetConstructedDataRepresentsBuilder() BACnetConstructedDataRepresentsBuilder {
	return &_BACnetConstructedDataRepresentsBuilder{_BACnetConstructedDataRepresents: new(_BACnetConstructedDataRepresents)}
}

type _BACnetConstructedDataRepresentsBuilder struct {
	*_BACnetConstructedDataRepresents

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataRepresentsBuilder) = (*_BACnetConstructedDataRepresentsBuilder)(nil)

func (b *_BACnetConstructedDataRepresentsBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataRepresents
}

func (b *_BACnetConstructedDataRepresentsBuilder) WithMandatoryFields(represents BACnetDeviceObjectReference) BACnetConstructedDataRepresentsBuilder {
	return b.WithRepresents(represents)
}

func (b *_BACnetConstructedDataRepresentsBuilder) WithRepresents(represents BACnetDeviceObjectReference) BACnetConstructedDataRepresentsBuilder {
	b.Represents = represents
	return b
}

func (b *_BACnetConstructedDataRepresentsBuilder) WithRepresentsBuilder(builderSupplier func(BACnetDeviceObjectReferenceBuilder) BACnetDeviceObjectReferenceBuilder) BACnetConstructedDataRepresentsBuilder {
	builder := builderSupplier(b.Represents.CreateBACnetDeviceObjectReferenceBuilder())
	var err error
	b.Represents, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDeviceObjectReferenceBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataRepresentsBuilder) Build() (BACnetConstructedDataRepresents, error) {
	if b.Represents == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'represents' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataRepresents.deepCopy(), nil
}

func (b *_BACnetConstructedDataRepresentsBuilder) MustBuild() BACnetConstructedDataRepresents {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataRepresentsBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataRepresentsBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataRepresentsBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataRepresentsBuilder().(*_BACnetConstructedDataRepresentsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataRepresentsBuilder creates a BACnetConstructedDataRepresentsBuilder
func (b *_BACnetConstructedDataRepresents) CreateBACnetConstructedDataRepresentsBuilder() BACnetConstructedDataRepresentsBuilder {
	if b == nil {
		return NewBACnetConstructedDataRepresentsBuilder()
	}
	return &_BACnetConstructedDataRepresentsBuilder{_BACnetConstructedDataRepresents: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataRepresents) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataRepresents) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_REPRESENTS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataRepresents) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataRepresents) GetRepresents() BACnetDeviceObjectReference {
	return m.Represents
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataRepresents) GetActualValue() BACnetDeviceObjectReference {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDeviceObjectReference(m.GetRepresents())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataRepresents(structType any) BACnetConstructedDataRepresents {
	if casted, ok := structType.(BACnetConstructedDataRepresents); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataRepresents); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataRepresents) GetTypeName() string {
	return "BACnetConstructedDataRepresents"
}

func (m *_BACnetConstructedDataRepresents) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (represents)
	lengthInBits += m.Represents.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataRepresents) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataRepresents) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataRepresents BACnetConstructedDataRepresents, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataRepresents"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataRepresents")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	represents, err := ReadSimpleField[BACnetDeviceObjectReference](ctx, "represents", ReadComplex[BACnetDeviceObjectReference](BACnetDeviceObjectReferenceParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'represents' field"))
	}
	m.Represents = represents

	actualValue, err := ReadVirtualField[BACnetDeviceObjectReference](ctx, "actualValue", (*BACnetDeviceObjectReference)(nil), represents)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataRepresents"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataRepresents")
	}

	return m, nil
}

func (m *_BACnetConstructedDataRepresents) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataRepresents) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataRepresents"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataRepresents")
		}

		if err := WriteSimpleField[BACnetDeviceObjectReference](ctx, "represents", m.GetRepresents(), WriteComplex[BACnetDeviceObjectReference](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'represents' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataRepresents"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataRepresents")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataRepresents) IsBACnetConstructedDataRepresents() {}

func (m *_BACnetConstructedDataRepresents) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataRepresents) deepCopy() *_BACnetConstructedDataRepresents {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataRepresentsCopy := &_BACnetConstructedDataRepresents{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetDeviceObjectReference](m.Represents),
	}
	_BACnetConstructedDataRepresentsCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataRepresentsCopy
}

func (m *_BACnetConstructedDataRepresents) String() string {
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
