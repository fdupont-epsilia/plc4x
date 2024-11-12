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

// BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate is the corresponding interface of BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate
type BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetNotificationParametersChangeOfDiscreteValueNewValue
	// GetDateValue returns DateValue (property field)
	GetDateValue() BACnetApplicationTagDate
	// IsBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate()
	// CreateBuilder creates a BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder
	CreateBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder() BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder
}

// _BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate is the data-structure of this message
type _BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate struct {
	BACnetNotificationParametersChangeOfDiscreteValueNewValueContract
	DateValue BACnetApplicationTagDate
}

var _ BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate = (*_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate)(nil)
var _ BACnetNotificationParametersChangeOfDiscreteValueNewValueRequirements = (*_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate)(nil)

// NewBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate factory function for _BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate
func NewBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, dateValue BACnetApplicationTagDate, tagNumber uint8) *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate {
	if dateValue == nil {
		panic("dateValue of type BACnetApplicationTagDate for BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate must not be nil")
	}
	_result := &_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate{
		BACnetNotificationParametersChangeOfDiscreteValueNewValueContract: NewBACnetNotificationParametersChangeOfDiscreteValueNewValue(openingTag, peekedTagHeader, closingTag, tagNumber),
		DateValue: dateValue,
	}
	_result.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract.(*_BACnetNotificationParametersChangeOfDiscreteValueNewValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder is a builder for BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate
type BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(dateValue BACnetApplicationTagDate) BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder
	// WithDateValue adds DateValue (property field)
	WithDateValue(BACnetApplicationTagDate) BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder
	// WithDateValueBuilder adds DateValue (property field) which is build by the builder
	WithDateValueBuilder(func(BACnetApplicationTagDateBuilder) BACnetApplicationTagDateBuilder) BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder
	// Build builds the BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate or returns an error if something is wrong
	Build() (BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate
}

// NewBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder() creates a BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder
func NewBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder() BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder {
	return &_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder{_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate: new(_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate)}
}

type _BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder struct {
	*_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate

	parentBuilder *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder

	err *utils.MultiError
}

var _ (BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder) = (*_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder)(nil)

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder) setParent(contract BACnetNotificationParametersChangeOfDiscreteValueNewValueContract) {
	b.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract = contract
	contract.(*_BACnetNotificationParametersChangeOfDiscreteValueNewValue)._SubType = b._BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate
}

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder) WithMandatoryFields(dateValue BACnetApplicationTagDate) BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder {
	return b.WithDateValue(dateValue)
}

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder) WithDateValue(dateValue BACnetApplicationTagDate) BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder {
	b.DateValue = dateValue
	return b
}

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder) WithDateValueBuilder(builderSupplier func(BACnetApplicationTagDateBuilder) BACnetApplicationTagDateBuilder) BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder {
	builder := builderSupplier(b.DateValue.CreateBACnetApplicationTagDateBuilder())
	var err error
	b.DateValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagDateBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder) Build() (BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate, error) {
	if b.DateValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'dateValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate.deepCopy(), nil
}

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder) MustBuild() BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder) Done() BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder().(*_BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder) buildForBACnetNotificationParametersChangeOfDiscreteValueNewValue() (BACnetNotificationParametersChangeOfDiscreteValueNewValue, error) {
	return b.Build()
}

func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder) DeepCopy() any {
	_copy := b.CreateBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder().(*_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder creates a BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder
func (b *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate) CreateBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder() BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder {
	if b == nil {
		return NewBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder()
	}
	return &_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateBuilder{_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate) GetParent() BACnetNotificationParametersChangeOfDiscreteValueNewValueContract {
	return m.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate) GetDateValue() BACnetApplicationTagDate {
	return m.DateValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate(structType any) BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate"
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract.(*_BACnetNotificationParametersChangeOfDiscreteValueNewValue).getLengthInBits(ctx))

	// Simple field (dateValue)
	lengthInBits += m.DateValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetNotificationParametersChangeOfDiscreteValueNewValue, tagNumber uint8) (__bACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate, err error) {
	m.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dateValue, err := ReadSimpleField[BACnetApplicationTagDate](ctx, "dateValue", ReadComplex[BACnetApplicationTagDate](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDate](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dateValue' field"))
	}
	m.DateValue = dateValue

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate")
	}

	return m, nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate")
		}

		if err := WriteSimpleField[BACnetApplicationTagDate](ctx, "dateValue", m.GetDateValue(), WriteComplex[BACnetApplicationTagDate](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dateValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate")
		}
		return nil
	}
	return m.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract.(*_BACnetNotificationParametersChangeOfDiscreteValueNewValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate) IsBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate() {
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate) deepCopy() *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate {
	if m == nil {
		return nil
	}
	_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateCopy := &_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate{
		m.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract.(*_BACnetNotificationParametersChangeOfDiscreteValueNewValue).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagDate](m.DateValue),
	}
	_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateCopy.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract.(*_BACnetNotificationParametersChangeOfDiscreteValueNewValue)._SubType = m
	return _BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateCopy
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate) String() string {
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
