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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataCommandAll is the corresponding interface of BACnetConstructedDataCommandAll
type BACnetConstructedDataCommandAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataCommandAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataCommandAll()
	// CreateBuilder creates a BACnetConstructedDataCommandAllBuilder
	CreateBACnetConstructedDataCommandAllBuilder() BACnetConstructedDataCommandAllBuilder
}

// _BACnetConstructedDataCommandAll is the data-structure of this message
type _BACnetConstructedDataCommandAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataCommandAll = (*_BACnetConstructedDataCommandAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataCommandAll)(nil)

// NewBACnetConstructedDataCommandAll factory function for _BACnetConstructedDataCommandAll
func NewBACnetConstructedDataCommandAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCommandAll {
	_result := &_BACnetConstructedDataCommandAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataCommandAllBuilder is a builder for BACnetConstructedDataCommandAll
type BACnetConstructedDataCommandAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataCommandAllBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataCommandAll or returns an error if something is wrong
	Build() (BACnetConstructedDataCommandAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataCommandAll
}

// NewBACnetConstructedDataCommandAllBuilder() creates a BACnetConstructedDataCommandAllBuilder
func NewBACnetConstructedDataCommandAllBuilder() BACnetConstructedDataCommandAllBuilder {
	return &_BACnetConstructedDataCommandAllBuilder{_BACnetConstructedDataCommandAll: new(_BACnetConstructedDataCommandAll)}
}

type _BACnetConstructedDataCommandAllBuilder struct {
	*_BACnetConstructedDataCommandAll

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataCommandAllBuilder) = (*_BACnetConstructedDataCommandAllBuilder)(nil)

func (b *_BACnetConstructedDataCommandAllBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataCommandAll
}

func (b *_BACnetConstructedDataCommandAllBuilder) WithMandatoryFields() BACnetConstructedDataCommandAllBuilder {
	return b
}

func (b *_BACnetConstructedDataCommandAllBuilder) Build() (BACnetConstructedDataCommandAll, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataCommandAll.deepCopy(), nil
}

func (b *_BACnetConstructedDataCommandAllBuilder) MustBuild() BACnetConstructedDataCommandAll {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataCommandAllBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataCommandAllBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataCommandAllBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataCommandAllBuilder().(*_BACnetConstructedDataCommandAllBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataCommandAllBuilder creates a BACnetConstructedDataCommandAllBuilder
func (b *_BACnetConstructedDataCommandAll) CreateBACnetConstructedDataCommandAllBuilder() BACnetConstructedDataCommandAllBuilder {
	if b == nil {
		return NewBACnetConstructedDataCommandAllBuilder()
	}
	return &_BACnetConstructedDataCommandAllBuilder{_BACnetConstructedDataCommandAll: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCommandAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_COMMAND
}

func (m *_BACnetConstructedDataCommandAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCommandAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCommandAll(structType any) BACnetConstructedDataCommandAll {
	if casted, ok := structType.(BACnetConstructedDataCommandAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCommandAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCommandAll) GetTypeName() string {
	return "BACnetConstructedDataCommandAll"
}

func (m *_BACnetConstructedDataCommandAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataCommandAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataCommandAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataCommandAll BACnetConstructedDataCommandAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCommandAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCommandAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCommandAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCommandAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataCommandAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCommandAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCommandAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCommandAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCommandAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCommandAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCommandAll) IsBACnetConstructedDataCommandAll() {}

func (m *_BACnetConstructedDataCommandAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataCommandAll) deepCopy() *_BACnetConstructedDataCommandAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataCommandAllCopy := &_BACnetConstructedDataCommandAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	_BACnetConstructedDataCommandAllCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataCommandAllCopy
}

func (m *_BACnetConstructedDataCommandAll) String() string {
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
