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

// BACnetConstructedDataLockoutRelinquishTime is the corresponding interface of BACnetConstructedDataLockoutRelinquishTime
type BACnetConstructedDataLockoutRelinquishTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetLockoutRelinquishTime returns LockoutRelinquishTime (property field)
	GetLockoutRelinquishTime() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataLockoutRelinquishTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLockoutRelinquishTime()
	// CreateBuilder creates a BACnetConstructedDataLockoutRelinquishTimeBuilder
	CreateBACnetConstructedDataLockoutRelinquishTimeBuilder() BACnetConstructedDataLockoutRelinquishTimeBuilder
}

// _BACnetConstructedDataLockoutRelinquishTime is the data-structure of this message
type _BACnetConstructedDataLockoutRelinquishTime struct {
	BACnetConstructedDataContract
	LockoutRelinquishTime BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataLockoutRelinquishTime = (*_BACnetConstructedDataLockoutRelinquishTime)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLockoutRelinquishTime)(nil)

// NewBACnetConstructedDataLockoutRelinquishTime factory function for _BACnetConstructedDataLockoutRelinquishTime
func NewBACnetConstructedDataLockoutRelinquishTime(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, lockoutRelinquishTime BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLockoutRelinquishTime {
	if lockoutRelinquishTime == nil {
		panic("lockoutRelinquishTime of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataLockoutRelinquishTime must not be nil")
	}
	_result := &_BACnetConstructedDataLockoutRelinquishTime{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LockoutRelinquishTime:         lockoutRelinquishTime,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLockoutRelinquishTimeBuilder is a builder for BACnetConstructedDataLockoutRelinquishTime
type BACnetConstructedDataLockoutRelinquishTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(lockoutRelinquishTime BACnetApplicationTagUnsignedInteger) BACnetConstructedDataLockoutRelinquishTimeBuilder
	// WithLockoutRelinquishTime adds LockoutRelinquishTime (property field)
	WithLockoutRelinquishTime(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataLockoutRelinquishTimeBuilder
	// WithLockoutRelinquishTimeBuilder adds LockoutRelinquishTime (property field) which is build by the builder
	WithLockoutRelinquishTimeBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataLockoutRelinquishTimeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataLockoutRelinquishTime or returns an error if something is wrong
	Build() (BACnetConstructedDataLockoutRelinquishTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLockoutRelinquishTime
}

// NewBACnetConstructedDataLockoutRelinquishTimeBuilder() creates a BACnetConstructedDataLockoutRelinquishTimeBuilder
func NewBACnetConstructedDataLockoutRelinquishTimeBuilder() BACnetConstructedDataLockoutRelinquishTimeBuilder {
	return &_BACnetConstructedDataLockoutRelinquishTimeBuilder{_BACnetConstructedDataLockoutRelinquishTime: new(_BACnetConstructedDataLockoutRelinquishTime)}
}

type _BACnetConstructedDataLockoutRelinquishTimeBuilder struct {
	*_BACnetConstructedDataLockoutRelinquishTime

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLockoutRelinquishTimeBuilder) = (*_BACnetConstructedDataLockoutRelinquishTimeBuilder)(nil)

func (b *_BACnetConstructedDataLockoutRelinquishTimeBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataLockoutRelinquishTime
}

func (b *_BACnetConstructedDataLockoutRelinquishTimeBuilder) WithMandatoryFields(lockoutRelinquishTime BACnetApplicationTagUnsignedInteger) BACnetConstructedDataLockoutRelinquishTimeBuilder {
	return b.WithLockoutRelinquishTime(lockoutRelinquishTime)
}

func (b *_BACnetConstructedDataLockoutRelinquishTimeBuilder) WithLockoutRelinquishTime(lockoutRelinquishTime BACnetApplicationTagUnsignedInteger) BACnetConstructedDataLockoutRelinquishTimeBuilder {
	b.LockoutRelinquishTime = lockoutRelinquishTime
	return b
}

func (b *_BACnetConstructedDataLockoutRelinquishTimeBuilder) WithLockoutRelinquishTimeBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataLockoutRelinquishTimeBuilder {
	builder := builderSupplier(b.LockoutRelinquishTime.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.LockoutRelinquishTime, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataLockoutRelinquishTimeBuilder) Build() (BACnetConstructedDataLockoutRelinquishTime, error) {
	if b.LockoutRelinquishTime == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'lockoutRelinquishTime' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLockoutRelinquishTime.deepCopy(), nil
}

func (b *_BACnetConstructedDataLockoutRelinquishTimeBuilder) MustBuild() BACnetConstructedDataLockoutRelinquishTime {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataLockoutRelinquishTimeBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLockoutRelinquishTimeBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLockoutRelinquishTimeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLockoutRelinquishTimeBuilder().(*_BACnetConstructedDataLockoutRelinquishTimeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLockoutRelinquishTimeBuilder creates a BACnetConstructedDataLockoutRelinquishTimeBuilder
func (b *_BACnetConstructedDataLockoutRelinquishTime) CreateBACnetConstructedDataLockoutRelinquishTimeBuilder() BACnetConstructedDataLockoutRelinquishTimeBuilder {
	if b == nil {
		return NewBACnetConstructedDataLockoutRelinquishTimeBuilder()
	}
	return &_BACnetConstructedDataLockoutRelinquishTimeBuilder{_BACnetConstructedDataLockoutRelinquishTime: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLockoutRelinquishTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLockoutRelinquishTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOCKOUT_RELINQUISH_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLockoutRelinquishTime) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLockoutRelinquishTime) GetLockoutRelinquishTime() BACnetApplicationTagUnsignedInteger {
	return m.LockoutRelinquishTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLockoutRelinquishTime) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetLockoutRelinquishTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLockoutRelinquishTime(structType any) BACnetConstructedDataLockoutRelinquishTime {
	if casted, ok := structType.(BACnetConstructedDataLockoutRelinquishTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLockoutRelinquishTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLockoutRelinquishTime) GetTypeName() string {
	return "BACnetConstructedDataLockoutRelinquishTime"
}

func (m *_BACnetConstructedDataLockoutRelinquishTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (lockoutRelinquishTime)
	lengthInBits += m.LockoutRelinquishTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLockoutRelinquishTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLockoutRelinquishTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLockoutRelinquishTime BACnetConstructedDataLockoutRelinquishTime, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLockoutRelinquishTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLockoutRelinquishTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lockoutRelinquishTime, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "lockoutRelinquishTime", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lockoutRelinquishTime' field"))
	}
	m.LockoutRelinquishTime = lockoutRelinquishTime

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), lockoutRelinquishTime)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLockoutRelinquishTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLockoutRelinquishTime")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLockoutRelinquishTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLockoutRelinquishTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLockoutRelinquishTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLockoutRelinquishTime")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "lockoutRelinquishTime", m.GetLockoutRelinquishTime(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lockoutRelinquishTime' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLockoutRelinquishTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLockoutRelinquishTime")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLockoutRelinquishTime) IsBACnetConstructedDataLockoutRelinquishTime() {
}

func (m *_BACnetConstructedDataLockoutRelinquishTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLockoutRelinquishTime) deepCopy() *_BACnetConstructedDataLockoutRelinquishTime {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLockoutRelinquishTimeCopy := &_BACnetConstructedDataLockoutRelinquishTime{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.LockoutRelinquishTime),
	}
	_BACnetConstructedDataLockoutRelinquishTimeCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLockoutRelinquishTimeCopy
}

func (m *_BACnetConstructedDataLockoutRelinquishTime) String() string {
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
