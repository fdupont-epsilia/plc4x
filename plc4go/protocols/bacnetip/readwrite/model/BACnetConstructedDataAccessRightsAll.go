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

// BACnetConstructedDataAccessRightsAll is the corresponding interface of BACnetConstructedDataAccessRightsAll
type BACnetConstructedDataAccessRightsAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataAccessRightsAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAccessRightsAll()
	// CreateBuilder creates a BACnetConstructedDataAccessRightsAllBuilder
	CreateBACnetConstructedDataAccessRightsAllBuilder() BACnetConstructedDataAccessRightsAllBuilder
}

// _BACnetConstructedDataAccessRightsAll is the data-structure of this message
type _BACnetConstructedDataAccessRightsAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataAccessRightsAll = (*_BACnetConstructedDataAccessRightsAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAccessRightsAll)(nil)

// NewBACnetConstructedDataAccessRightsAll factory function for _BACnetConstructedDataAccessRightsAll
func NewBACnetConstructedDataAccessRightsAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccessRightsAll {
	_result := &_BACnetConstructedDataAccessRightsAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAccessRightsAllBuilder is a builder for BACnetConstructedDataAccessRightsAll
type BACnetConstructedDataAccessRightsAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataAccessRightsAllBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataAccessRightsAll or returns an error if something is wrong
	Build() (BACnetConstructedDataAccessRightsAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAccessRightsAll
}

// NewBACnetConstructedDataAccessRightsAllBuilder() creates a BACnetConstructedDataAccessRightsAllBuilder
func NewBACnetConstructedDataAccessRightsAllBuilder() BACnetConstructedDataAccessRightsAllBuilder {
	return &_BACnetConstructedDataAccessRightsAllBuilder{_BACnetConstructedDataAccessRightsAll: new(_BACnetConstructedDataAccessRightsAll)}
}

type _BACnetConstructedDataAccessRightsAllBuilder struct {
	*_BACnetConstructedDataAccessRightsAll

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataAccessRightsAllBuilder) = (*_BACnetConstructedDataAccessRightsAllBuilder)(nil)

func (b *_BACnetConstructedDataAccessRightsAllBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataAccessRightsAll
}

func (b *_BACnetConstructedDataAccessRightsAllBuilder) WithMandatoryFields() BACnetConstructedDataAccessRightsAllBuilder {
	return b
}

func (b *_BACnetConstructedDataAccessRightsAllBuilder) Build() (BACnetConstructedDataAccessRightsAll, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataAccessRightsAll.deepCopy(), nil
}

func (b *_BACnetConstructedDataAccessRightsAllBuilder) MustBuild() BACnetConstructedDataAccessRightsAll {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataAccessRightsAllBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataAccessRightsAllBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataAccessRightsAllBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataAccessRightsAllBuilder().(*_BACnetConstructedDataAccessRightsAllBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataAccessRightsAllBuilder creates a BACnetConstructedDataAccessRightsAllBuilder
func (b *_BACnetConstructedDataAccessRightsAll) CreateBACnetConstructedDataAccessRightsAllBuilder() BACnetConstructedDataAccessRightsAllBuilder {
	if b == nil {
		return NewBACnetConstructedDataAccessRightsAllBuilder()
	}
	return &_BACnetConstructedDataAccessRightsAllBuilder{_BACnetConstructedDataAccessRightsAll: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccessRightsAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ACCESS_RIGHTS
}

func (m *_BACnetConstructedDataAccessRightsAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccessRightsAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccessRightsAll(structType any) BACnetConstructedDataAccessRightsAll {
	if casted, ok := structType.(BACnetConstructedDataAccessRightsAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccessRightsAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccessRightsAll) GetTypeName() string {
	return "BACnetConstructedDataAccessRightsAll"
}

func (m *_BACnetConstructedDataAccessRightsAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataAccessRightsAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAccessRightsAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAccessRightsAll BACnetConstructedDataAccessRightsAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccessRightsAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccessRightsAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccessRightsAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccessRightsAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAccessRightsAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccessRightsAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccessRightsAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccessRightsAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccessRightsAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccessRightsAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccessRightsAll) IsBACnetConstructedDataAccessRightsAll() {}

func (m *_BACnetConstructedDataAccessRightsAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAccessRightsAll) deepCopy() *_BACnetConstructedDataAccessRightsAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAccessRightsAllCopy := &_BACnetConstructedDataAccessRightsAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	_BACnetConstructedDataAccessRightsAllCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAccessRightsAllCopy
}

func (m *_BACnetConstructedDataAccessRightsAll) String() string {
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
