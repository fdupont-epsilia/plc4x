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

// BACnetTimeStampsEnclosed is the corresponding interface of BACnetTimeStampsEnclosed
type BACnetTimeStampsEnclosed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetTimestamps returns Timestamps (property field)
	GetTimestamps() []BACnetTimeStamp
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetTimeStampsEnclosed is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTimeStampsEnclosed()
	// CreateBuilder creates a BACnetTimeStampsEnclosedBuilder
	CreateBACnetTimeStampsEnclosedBuilder() BACnetTimeStampsEnclosedBuilder
}

// _BACnetTimeStampsEnclosed is the data-structure of this message
type _BACnetTimeStampsEnclosed struct {
	OpeningTag BACnetOpeningTag
	Timestamps []BACnetTimeStamp
	ClosingTag BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetTimeStampsEnclosed = (*_BACnetTimeStampsEnclosed)(nil)

// NewBACnetTimeStampsEnclosed factory function for _BACnetTimeStampsEnclosed
func NewBACnetTimeStampsEnclosed(openingTag BACnetOpeningTag, timestamps []BACnetTimeStamp, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetTimeStampsEnclosed {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetTimeStampsEnclosed must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetTimeStampsEnclosed must not be nil")
	}
	return &_BACnetTimeStampsEnclosed{OpeningTag: openingTag, Timestamps: timestamps, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetTimeStampsEnclosedBuilder is a builder for BACnetTimeStampsEnclosed
type BACnetTimeStampsEnclosedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, timestamps []BACnetTimeStamp, closingTag BACnetClosingTag) BACnetTimeStampsEnclosedBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetTimeStampsEnclosedBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetTimeStampsEnclosedBuilder
	// WithTimestamps adds Timestamps (property field)
	WithTimestamps(...BACnetTimeStamp) BACnetTimeStampsEnclosedBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetTimeStampsEnclosedBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetTimeStampsEnclosedBuilder
	// WithArgTagNumber sets a parser argument
	WithArgTagNumber(uint8) BACnetTimeStampsEnclosedBuilder
	// Build builds the BACnetTimeStampsEnclosed or returns an error if something is wrong
	Build() (BACnetTimeStampsEnclosed, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetTimeStampsEnclosed
}

// NewBACnetTimeStampsEnclosedBuilder() creates a BACnetTimeStampsEnclosedBuilder
func NewBACnetTimeStampsEnclosedBuilder() BACnetTimeStampsEnclosedBuilder {
	return &_BACnetTimeStampsEnclosedBuilder{_BACnetTimeStampsEnclosed: new(_BACnetTimeStampsEnclosed)}
}

type _BACnetTimeStampsEnclosedBuilder struct {
	*_BACnetTimeStampsEnclosed

	err *utils.MultiError
}

var _ (BACnetTimeStampsEnclosedBuilder) = (*_BACnetTimeStampsEnclosedBuilder)(nil)

func (b *_BACnetTimeStampsEnclosedBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, timestamps []BACnetTimeStamp, closingTag BACnetClosingTag) BACnetTimeStampsEnclosedBuilder {
	return b.WithOpeningTag(openingTag).WithTimestamps(timestamps...).WithClosingTag(closingTag)
}

func (b *_BACnetTimeStampsEnclosedBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetTimeStampsEnclosedBuilder {
	b.OpeningTag = openingTag
	return b
}

func (b *_BACnetTimeStampsEnclosedBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetTimeStampsEnclosedBuilder {
	builder := builderSupplier(b.OpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	b.OpeningTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return b
}

func (b *_BACnetTimeStampsEnclosedBuilder) WithTimestamps(timestamps ...BACnetTimeStamp) BACnetTimeStampsEnclosedBuilder {
	b.Timestamps = timestamps
	return b
}

func (b *_BACnetTimeStampsEnclosedBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetTimeStampsEnclosedBuilder {
	b.ClosingTag = closingTag
	return b
}

func (b *_BACnetTimeStampsEnclosedBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetTimeStampsEnclosedBuilder {
	builder := builderSupplier(b.ClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	b.ClosingTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return b
}

func (b *_BACnetTimeStampsEnclosedBuilder) WithArgTagNumber(tagNumber uint8) BACnetTimeStampsEnclosedBuilder {
	b.TagNumber = tagNumber
	return b
}

func (b *_BACnetTimeStampsEnclosedBuilder) Build() (BACnetTimeStampsEnclosed, error) {
	if b.OpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if b.ClosingTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'closingTag' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetTimeStampsEnclosed.deepCopy(), nil
}

func (b *_BACnetTimeStampsEnclosedBuilder) MustBuild() BACnetTimeStampsEnclosed {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetTimeStampsEnclosedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetTimeStampsEnclosedBuilder().(*_BACnetTimeStampsEnclosedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetTimeStampsEnclosedBuilder creates a BACnetTimeStampsEnclosedBuilder
func (b *_BACnetTimeStampsEnclosed) CreateBACnetTimeStampsEnclosedBuilder() BACnetTimeStampsEnclosedBuilder {
	if b == nil {
		return NewBACnetTimeStampsEnclosedBuilder()
	}
	return &_BACnetTimeStampsEnclosedBuilder{_BACnetTimeStampsEnclosed: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimeStampsEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetTimeStampsEnclosed) GetTimestamps() []BACnetTimeStamp {
	return m.Timestamps
}

func (m *_BACnetTimeStampsEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetTimeStampsEnclosed(structType any) BACnetTimeStampsEnclosed {
	if casted, ok := structType.(BACnetTimeStampsEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimeStampsEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimeStampsEnclosed) GetTypeName() string {
	return "BACnetTimeStampsEnclosed"
}

func (m *_BACnetTimeStampsEnclosed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.Timestamps) > 0 {
		for _, element := range m.Timestamps {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetTimeStampsEnclosed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTimeStampsEnclosedParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetTimeStampsEnclosed, error) {
	return BACnetTimeStampsEnclosedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetTimeStampsEnclosedParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTimeStampsEnclosed, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTimeStampsEnclosed, error) {
		return BACnetTimeStampsEnclosedParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetTimeStampsEnclosedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetTimeStampsEnclosed, error) {
	v, err := (&_BACnetTimeStampsEnclosed{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetTimeStampsEnclosed) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetTimeStampsEnclosed BACnetTimeStampsEnclosed, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimeStampsEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimeStampsEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	timestamps, err := ReadTerminatedArrayField[BACnetTimeStamp](ctx, "timestamps", ReadComplex[BACnetTimeStamp](BACnetTimeStampParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timestamps' field"))
	}
	m.Timestamps = timestamps

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetTimeStampsEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimeStampsEnclosed")
	}

	return m, nil
}

func (m *_BACnetTimeStampsEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimeStampsEnclosed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetTimeStampsEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTimeStampsEnclosed")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "timestamps", m.GetTimestamps(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'timestamps' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTimeStampsEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTimeStampsEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetTimeStampsEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetTimeStampsEnclosed) IsBACnetTimeStampsEnclosed() {}

func (m *_BACnetTimeStampsEnclosed) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetTimeStampsEnclosed) deepCopy() *_BACnetTimeStampsEnclosed {
	if m == nil {
		return nil
	}
	_BACnetTimeStampsEnclosedCopy := &_BACnetTimeStampsEnclosed{
		utils.DeepCopy[BACnetOpeningTag](m.OpeningTag),
		utils.DeepCopySlice[BACnetTimeStamp, BACnetTimeStamp](m.Timestamps),
		utils.DeepCopy[BACnetClosingTag](m.ClosingTag),
		m.TagNumber,
	}
	return _BACnetTimeStampsEnclosedCopy
}

func (m *_BACnetTimeStampsEnclosed) String() string {
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
