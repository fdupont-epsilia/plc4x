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

// BACnetConstructedDataSupportedSecurityAlgorithms is the corresponding interface of BACnetConstructedDataSupportedSecurityAlgorithms
type BACnetConstructedDataSupportedSecurityAlgorithms interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetSupportedSecurityAlgorithms returns SupportedSecurityAlgorithms (property field)
	GetSupportedSecurityAlgorithms() []BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataSupportedSecurityAlgorithms is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataSupportedSecurityAlgorithms()
	// CreateBuilder creates a BACnetConstructedDataSupportedSecurityAlgorithmsBuilder
	CreateBACnetConstructedDataSupportedSecurityAlgorithmsBuilder() BACnetConstructedDataSupportedSecurityAlgorithmsBuilder
}

// _BACnetConstructedDataSupportedSecurityAlgorithms is the data-structure of this message
type _BACnetConstructedDataSupportedSecurityAlgorithms struct {
	BACnetConstructedDataContract
	SupportedSecurityAlgorithms []BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataSupportedSecurityAlgorithms = (*_BACnetConstructedDataSupportedSecurityAlgorithms)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataSupportedSecurityAlgorithms)(nil)

// NewBACnetConstructedDataSupportedSecurityAlgorithms factory function for _BACnetConstructedDataSupportedSecurityAlgorithms
func NewBACnetConstructedDataSupportedSecurityAlgorithms(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, supportedSecurityAlgorithms []BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataSupportedSecurityAlgorithms {
	_result := &_BACnetConstructedDataSupportedSecurityAlgorithms{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		SupportedSecurityAlgorithms:   supportedSecurityAlgorithms,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataSupportedSecurityAlgorithmsBuilder is a builder for BACnetConstructedDataSupportedSecurityAlgorithms
type BACnetConstructedDataSupportedSecurityAlgorithmsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(supportedSecurityAlgorithms []BACnetApplicationTagUnsignedInteger) BACnetConstructedDataSupportedSecurityAlgorithmsBuilder
	// WithSupportedSecurityAlgorithms adds SupportedSecurityAlgorithms (property field)
	WithSupportedSecurityAlgorithms(...BACnetApplicationTagUnsignedInteger) BACnetConstructedDataSupportedSecurityAlgorithmsBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataSupportedSecurityAlgorithms or returns an error if something is wrong
	Build() (BACnetConstructedDataSupportedSecurityAlgorithms, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataSupportedSecurityAlgorithms
}

// NewBACnetConstructedDataSupportedSecurityAlgorithmsBuilder() creates a BACnetConstructedDataSupportedSecurityAlgorithmsBuilder
func NewBACnetConstructedDataSupportedSecurityAlgorithmsBuilder() BACnetConstructedDataSupportedSecurityAlgorithmsBuilder {
	return &_BACnetConstructedDataSupportedSecurityAlgorithmsBuilder{_BACnetConstructedDataSupportedSecurityAlgorithms: new(_BACnetConstructedDataSupportedSecurityAlgorithms)}
}

type _BACnetConstructedDataSupportedSecurityAlgorithmsBuilder struct {
	*_BACnetConstructedDataSupportedSecurityAlgorithms

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataSupportedSecurityAlgorithmsBuilder) = (*_BACnetConstructedDataSupportedSecurityAlgorithmsBuilder)(nil)

func (b *_BACnetConstructedDataSupportedSecurityAlgorithmsBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataSupportedSecurityAlgorithms
}

func (b *_BACnetConstructedDataSupportedSecurityAlgorithmsBuilder) WithMandatoryFields(supportedSecurityAlgorithms []BACnetApplicationTagUnsignedInteger) BACnetConstructedDataSupportedSecurityAlgorithmsBuilder {
	return b.WithSupportedSecurityAlgorithms(supportedSecurityAlgorithms...)
}

func (b *_BACnetConstructedDataSupportedSecurityAlgorithmsBuilder) WithSupportedSecurityAlgorithms(supportedSecurityAlgorithms ...BACnetApplicationTagUnsignedInteger) BACnetConstructedDataSupportedSecurityAlgorithmsBuilder {
	b.SupportedSecurityAlgorithms = supportedSecurityAlgorithms
	return b
}

func (b *_BACnetConstructedDataSupportedSecurityAlgorithmsBuilder) Build() (BACnetConstructedDataSupportedSecurityAlgorithms, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataSupportedSecurityAlgorithms.deepCopy(), nil
}

func (b *_BACnetConstructedDataSupportedSecurityAlgorithmsBuilder) MustBuild() BACnetConstructedDataSupportedSecurityAlgorithms {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataSupportedSecurityAlgorithmsBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataSupportedSecurityAlgorithmsBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataSupportedSecurityAlgorithmsBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataSupportedSecurityAlgorithmsBuilder().(*_BACnetConstructedDataSupportedSecurityAlgorithmsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataSupportedSecurityAlgorithmsBuilder creates a BACnetConstructedDataSupportedSecurityAlgorithmsBuilder
func (b *_BACnetConstructedDataSupportedSecurityAlgorithms) CreateBACnetConstructedDataSupportedSecurityAlgorithmsBuilder() BACnetConstructedDataSupportedSecurityAlgorithmsBuilder {
	if b == nil {
		return NewBACnetConstructedDataSupportedSecurityAlgorithmsBuilder()
	}
	return &_BACnetConstructedDataSupportedSecurityAlgorithmsBuilder{_BACnetConstructedDataSupportedSecurityAlgorithms: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SUPPORTED_SECURITY_ALGORITHMS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) GetSupportedSecurityAlgorithms() []BACnetApplicationTagUnsignedInteger {
	return m.SupportedSecurityAlgorithms
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataSupportedSecurityAlgorithms(structType any) BACnetConstructedDataSupportedSecurityAlgorithms {
	if casted, ok := structType.(BACnetConstructedDataSupportedSecurityAlgorithms); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSupportedSecurityAlgorithms); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) GetTypeName() string {
	return "BACnetConstructedDataSupportedSecurityAlgorithms"
}

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.SupportedSecurityAlgorithms) > 0 {
		for _, element := range m.SupportedSecurityAlgorithms {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataSupportedSecurityAlgorithms BACnetConstructedDataSupportedSecurityAlgorithms, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSupportedSecurityAlgorithms"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSupportedSecurityAlgorithms")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	supportedSecurityAlgorithms, err := ReadTerminatedArrayField[BACnetApplicationTagUnsignedInteger](ctx, "supportedSecurityAlgorithms", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'supportedSecurityAlgorithms' field"))
	}
	m.SupportedSecurityAlgorithms = supportedSecurityAlgorithms

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSupportedSecurityAlgorithms"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSupportedSecurityAlgorithms")
	}

	return m, nil
}

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSupportedSecurityAlgorithms"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSupportedSecurityAlgorithms")
		}

		if err := WriteComplexTypeArrayField(ctx, "supportedSecurityAlgorithms", m.GetSupportedSecurityAlgorithms(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'supportedSecurityAlgorithms' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSupportedSecurityAlgorithms"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSupportedSecurityAlgorithms")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) IsBACnetConstructedDataSupportedSecurityAlgorithms() {
}

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) deepCopy() *_BACnetConstructedDataSupportedSecurityAlgorithms {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataSupportedSecurityAlgorithmsCopy := &_BACnetConstructedDataSupportedSecurityAlgorithms{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopySlice[BACnetApplicationTagUnsignedInteger, BACnetApplicationTagUnsignedInteger](m.SupportedSecurityAlgorithms),
	}
	_BACnetConstructedDataSupportedSecurityAlgorithmsCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataSupportedSecurityAlgorithmsCopy
}

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) String() string {
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
