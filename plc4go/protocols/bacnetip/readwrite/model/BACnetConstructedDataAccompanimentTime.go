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

// BACnetConstructedDataAccompanimentTime is the corresponding interface of BACnetConstructedDataAccompanimentTime
type BACnetConstructedDataAccompanimentTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetAccompanimentTime returns AccompanimentTime (property field)
	GetAccompanimentTime() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataAccompanimentTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAccompanimentTime()
	// CreateBuilder creates a BACnetConstructedDataAccompanimentTimeBuilder
	CreateBACnetConstructedDataAccompanimentTimeBuilder() BACnetConstructedDataAccompanimentTimeBuilder
}

// _BACnetConstructedDataAccompanimentTime is the data-structure of this message
type _BACnetConstructedDataAccompanimentTime struct {
	BACnetConstructedDataContract
	AccompanimentTime BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataAccompanimentTime = (*_BACnetConstructedDataAccompanimentTime)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAccompanimentTime)(nil)

// NewBACnetConstructedDataAccompanimentTime factory function for _BACnetConstructedDataAccompanimentTime
func NewBACnetConstructedDataAccompanimentTime(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, accompanimentTime BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccompanimentTime {
	if accompanimentTime == nil {
		panic("accompanimentTime of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataAccompanimentTime must not be nil")
	}
	_result := &_BACnetConstructedDataAccompanimentTime{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		AccompanimentTime:             accompanimentTime,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAccompanimentTimeBuilder is a builder for BACnetConstructedDataAccompanimentTime
type BACnetConstructedDataAccompanimentTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(accompanimentTime BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAccompanimentTimeBuilder
	// WithAccompanimentTime adds AccompanimentTime (property field)
	WithAccompanimentTime(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAccompanimentTimeBuilder
	// WithAccompanimentTimeBuilder adds AccompanimentTime (property field) which is build by the builder
	WithAccompanimentTimeBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataAccompanimentTimeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataAccompanimentTime or returns an error if something is wrong
	Build() (BACnetConstructedDataAccompanimentTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAccompanimentTime
}

// NewBACnetConstructedDataAccompanimentTimeBuilder() creates a BACnetConstructedDataAccompanimentTimeBuilder
func NewBACnetConstructedDataAccompanimentTimeBuilder() BACnetConstructedDataAccompanimentTimeBuilder {
	return &_BACnetConstructedDataAccompanimentTimeBuilder{_BACnetConstructedDataAccompanimentTime: new(_BACnetConstructedDataAccompanimentTime)}
}

type _BACnetConstructedDataAccompanimentTimeBuilder struct {
	*_BACnetConstructedDataAccompanimentTime

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataAccompanimentTimeBuilder) = (*_BACnetConstructedDataAccompanimentTimeBuilder)(nil)

func (b *_BACnetConstructedDataAccompanimentTimeBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataAccompanimentTime
}

func (b *_BACnetConstructedDataAccompanimentTimeBuilder) WithMandatoryFields(accompanimentTime BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAccompanimentTimeBuilder {
	return b.WithAccompanimentTime(accompanimentTime)
}

func (b *_BACnetConstructedDataAccompanimentTimeBuilder) WithAccompanimentTime(accompanimentTime BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAccompanimentTimeBuilder {
	b.AccompanimentTime = accompanimentTime
	return b
}

func (b *_BACnetConstructedDataAccompanimentTimeBuilder) WithAccompanimentTimeBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataAccompanimentTimeBuilder {
	builder := builderSupplier(b.AccompanimentTime.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.AccompanimentTime, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataAccompanimentTimeBuilder) Build() (BACnetConstructedDataAccompanimentTime, error) {
	if b.AccompanimentTime == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'accompanimentTime' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataAccompanimentTime.deepCopy(), nil
}

func (b *_BACnetConstructedDataAccompanimentTimeBuilder) MustBuild() BACnetConstructedDataAccompanimentTime {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataAccompanimentTimeBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataAccompanimentTimeBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataAccompanimentTimeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataAccompanimentTimeBuilder().(*_BACnetConstructedDataAccompanimentTimeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataAccompanimentTimeBuilder creates a BACnetConstructedDataAccompanimentTimeBuilder
func (b *_BACnetConstructedDataAccompanimentTime) CreateBACnetConstructedDataAccompanimentTimeBuilder() BACnetConstructedDataAccompanimentTimeBuilder {
	if b == nil {
		return NewBACnetConstructedDataAccompanimentTimeBuilder()
	}
	return &_BACnetConstructedDataAccompanimentTimeBuilder{_BACnetConstructedDataAccompanimentTime: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccompanimentTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAccompanimentTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACCOMPANIMENT_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccompanimentTime) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccompanimentTime) GetAccompanimentTime() BACnetApplicationTagUnsignedInteger {
	return m.AccompanimentTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAccompanimentTime) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetAccompanimentTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccompanimentTime(structType any) BACnetConstructedDataAccompanimentTime {
	if casted, ok := structType.(BACnetConstructedDataAccompanimentTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccompanimentTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccompanimentTime) GetTypeName() string {
	return "BACnetConstructedDataAccompanimentTime"
}

func (m *_BACnetConstructedDataAccompanimentTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (accompanimentTime)
	lengthInBits += m.AccompanimentTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAccompanimentTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAccompanimentTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAccompanimentTime BACnetConstructedDataAccompanimentTime, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccompanimentTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccompanimentTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	accompanimentTime, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "accompanimentTime", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'accompanimentTime' field"))
	}
	m.AccompanimentTime = accompanimentTime

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), accompanimentTime)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccompanimentTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccompanimentTime")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAccompanimentTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccompanimentTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccompanimentTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccompanimentTime")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "accompanimentTime", m.GetAccompanimentTime(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'accompanimentTime' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccompanimentTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccompanimentTime")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccompanimentTime) IsBACnetConstructedDataAccompanimentTime() {}

func (m *_BACnetConstructedDataAccompanimentTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAccompanimentTime) deepCopy() *_BACnetConstructedDataAccompanimentTime {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAccompanimentTimeCopy := &_BACnetConstructedDataAccompanimentTime{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.AccompanimentTime),
	}
	_BACnetConstructedDataAccompanimentTimeCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAccompanimentTimeCopy
}

func (m *_BACnetConstructedDataAccompanimentTime) String() string {
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
