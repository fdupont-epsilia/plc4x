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

// BACnetConstructedDataRecordCount is the corresponding interface of BACnetConstructedDataRecordCount
type BACnetConstructedDataRecordCount interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetRecordCount returns RecordCount (property field)
	GetRecordCount() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataRecordCount is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataRecordCount()
	// CreateBuilder creates a BACnetConstructedDataRecordCountBuilder
	CreateBACnetConstructedDataRecordCountBuilder() BACnetConstructedDataRecordCountBuilder
}

// _BACnetConstructedDataRecordCount is the data-structure of this message
type _BACnetConstructedDataRecordCount struct {
	BACnetConstructedDataContract
	RecordCount BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataRecordCount = (*_BACnetConstructedDataRecordCount)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataRecordCount)(nil)

// NewBACnetConstructedDataRecordCount factory function for _BACnetConstructedDataRecordCount
func NewBACnetConstructedDataRecordCount(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, recordCount BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataRecordCount {
	if recordCount == nil {
		panic("recordCount of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataRecordCount must not be nil")
	}
	_result := &_BACnetConstructedDataRecordCount{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		RecordCount:                   recordCount,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataRecordCountBuilder is a builder for BACnetConstructedDataRecordCount
type BACnetConstructedDataRecordCountBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(recordCount BACnetApplicationTagUnsignedInteger) BACnetConstructedDataRecordCountBuilder
	// WithRecordCount adds RecordCount (property field)
	WithRecordCount(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataRecordCountBuilder
	// WithRecordCountBuilder adds RecordCount (property field) which is build by the builder
	WithRecordCountBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataRecordCountBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataRecordCount or returns an error if something is wrong
	Build() (BACnetConstructedDataRecordCount, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataRecordCount
}

// NewBACnetConstructedDataRecordCountBuilder() creates a BACnetConstructedDataRecordCountBuilder
func NewBACnetConstructedDataRecordCountBuilder() BACnetConstructedDataRecordCountBuilder {
	return &_BACnetConstructedDataRecordCountBuilder{_BACnetConstructedDataRecordCount: new(_BACnetConstructedDataRecordCount)}
}

type _BACnetConstructedDataRecordCountBuilder struct {
	*_BACnetConstructedDataRecordCount

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataRecordCountBuilder) = (*_BACnetConstructedDataRecordCountBuilder)(nil)

func (b *_BACnetConstructedDataRecordCountBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataRecordCount
}

func (b *_BACnetConstructedDataRecordCountBuilder) WithMandatoryFields(recordCount BACnetApplicationTagUnsignedInteger) BACnetConstructedDataRecordCountBuilder {
	return b.WithRecordCount(recordCount)
}

func (b *_BACnetConstructedDataRecordCountBuilder) WithRecordCount(recordCount BACnetApplicationTagUnsignedInteger) BACnetConstructedDataRecordCountBuilder {
	b.RecordCount = recordCount
	return b
}

func (b *_BACnetConstructedDataRecordCountBuilder) WithRecordCountBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataRecordCountBuilder {
	builder := builderSupplier(b.RecordCount.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.RecordCount, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataRecordCountBuilder) Build() (BACnetConstructedDataRecordCount, error) {
	if b.RecordCount == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'recordCount' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataRecordCount.deepCopy(), nil
}

func (b *_BACnetConstructedDataRecordCountBuilder) MustBuild() BACnetConstructedDataRecordCount {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataRecordCountBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataRecordCountBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataRecordCountBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataRecordCountBuilder().(*_BACnetConstructedDataRecordCountBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataRecordCountBuilder creates a BACnetConstructedDataRecordCountBuilder
func (b *_BACnetConstructedDataRecordCount) CreateBACnetConstructedDataRecordCountBuilder() BACnetConstructedDataRecordCountBuilder {
	if b == nil {
		return NewBACnetConstructedDataRecordCountBuilder()
	}
	return &_BACnetConstructedDataRecordCountBuilder{_BACnetConstructedDataRecordCount: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataRecordCount) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataRecordCount) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RECORD_COUNT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataRecordCount) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataRecordCount) GetRecordCount() BACnetApplicationTagUnsignedInteger {
	return m.RecordCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataRecordCount) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetRecordCount())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataRecordCount(structType any) BACnetConstructedDataRecordCount {
	if casted, ok := structType.(BACnetConstructedDataRecordCount); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataRecordCount); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataRecordCount) GetTypeName() string {
	return "BACnetConstructedDataRecordCount"
}

func (m *_BACnetConstructedDataRecordCount) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (recordCount)
	lengthInBits += m.RecordCount.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataRecordCount) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataRecordCount) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataRecordCount BACnetConstructedDataRecordCount, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataRecordCount"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataRecordCount")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	recordCount, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "recordCount", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'recordCount' field"))
	}
	m.RecordCount = recordCount

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), recordCount)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataRecordCount"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataRecordCount")
	}

	return m, nil
}

func (m *_BACnetConstructedDataRecordCount) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataRecordCount) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataRecordCount"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataRecordCount")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "recordCount", m.GetRecordCount(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'recordCount' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataRecordCount"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataRecordCount")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataRecordCount) IsBACnetConstructedDataRecordCount() {}

func (m *_BACnetConstructedDataRecordCount) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataRecordCount) deepCopy() *_BACnetConstructedDataRecordCount {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataRecordCountCopy := &_BACnetConstructedDataRecordCount{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.RecordCount),
	}
	_BACnetConstructedDataRecordCountCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataRecordCountCopy
}

func (m *_BACnetConstructedDataRecordCount) String() string {
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
