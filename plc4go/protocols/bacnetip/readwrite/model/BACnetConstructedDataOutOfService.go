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

// BACnetConstructedDataOutOfService is the corresponding interface of BACnetConstructedDataOutOfService
type BACnetConstructedDataOutOfService interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetOutOfService returns OutOfService (property field)
	GetOutOfService() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
	// IsBACnetConstructedDataOutOfService is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataOutOfService()
	// CreateBuilder creates a BACnetConstructedDataOutOfServiceBuilder
	CreateBACnetConstructedDataOutOfServiceBuilder() BACnetConstructedDataOutOfServiceBuilder
}

// _BACnetConstructedDataOutOfService is the data-structure of this message
type _BACnetConstructedDataOutOfService struct {
	BACnetConstructedDataContract
	OutOfService BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataOutOfService = (*_BACnetConstructedDataOutOfService)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataOutOfService)(nil)

// NewBACnetConstructedDataOutOfService factory function for _BACnetConstructedDataOutOfService
func NewBACnetConstructedDataOutOfService(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, outOfService BACnetApplicationTagBoolean, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataOutOfService {
	if outOfService == nil {
		panic("outOfService of type BACnetApplicationTagBoolean for BACnetConstructedDataOutOfService must not be nil")
	}
	_result := &_BACnetConstructedDataOutOfService{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		OutOfService:                  outOfService,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataOutOfServiceBuilder is a builder for BACnetConstructedDataOutOfService
type BACnetConstructedDataOutOfServiceBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(outOfService BACnetApplicationTagBoolean) BACnetConstructedDataOutOfServiceBuilder
	// WithOutOfService adds OutOfService (property field)
	WithOutOfService(BACnetApplicationTagBoolean) BACnetConstructedDataOutOfServiceBuilder
	// WithOutOfServiceBuilder adds OutOfService (property field) which is build by the builder
	WithOutOfServiceBuilder(func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataOutOfServiceBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataOutOfService or returns an error if something is wrong
	Build() (BACnetConstructedDataOutOfService, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataOutOfService
}

// NewBACnetConstructedDataOutOfServiceBuilder() creates a BACnetConstructedDataOutOfServiceBuilder
func NewBACnetConstructedDataOutOfServiceBuilder() BACnetConstructedDataOutOfServiceBuilder {
	return &_BACnetConstructedDataOutOfServiceBuilder{_BACnetConstructedDataOutOfService: new(_BACnetConstructedDataOutOfService)}
}

type _BACnetConstructedDataOutOfServiceBuilder struct {
	*_BACnetConstructedDataOutOfService

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataOutOfServiceBuilder) = (*_BACnetConstructedDataOutOfServiceBuilder)(nil)

func (b *_BACnetConstructedDataOutOfServiceBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataOutOfService
}

func (b *_BACnetConstructedDataOutOfServiceBuilder) WithMandatoryFields(outOfService BACnetApplicationTagBoolean) BACnetConstructedDataOutOfServiceBuilder {
	return b.WithOutOfService(outOfService)
}

func (b *_BACnetConstructedDataOutOfServiceBuilder) WithOutOfService(outOfService BACnetApplicationTagBoolean) BACnetConstructedDataOutOfServiceBuilder {
	b.OutOfService = outOfService
	return b
}

func (b *_BACnetConstructedDataOutOfServiceBuilder) WithOutOfServiceBuilder(builderSupplier func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataOutOfServiceBuilder {
	builder := builderSupplier(b.OutOfService.CreateBACnetApplicationTagBooleanBuilder())
	var err error
	b.OutOfService, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagBooleanBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataOutOfServiceBuilder) Build() (BACnetConstructedDataOutOfService, error) {
	if b.OutOfService == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'outOfService' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataOutOfService.deepCopy(), nil
}

func (b *_BACnetConstructedDataOutOfServiceBuilder) MustBuild() BACnetConstructedDataOutOfService {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataOutOfServiceBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataOutOfServiceBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataOutOfServiceBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataOutOfServiceBuilder().(*_BACnetConstructedDataOutOfServiceBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataOutOfServiceBuilder creates a BACnetConstructedDataOutOfServiceBuilder
func (b *_BACnetConstructedDataOutOfService) CreateBACnetConstructedDataOutOfServiceBuilder() BACnetConstructedDataOutOfServiceBuilder {
	if b == nil {
		return NewBACnetConstructedDataOutOfServiceBuilder()
	}
	return &_BACnetConstructedDataOutOfServiceBuilder{_BACnetConstructedDataOutOfService: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataOutOfService) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataOutOfService) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_OUT_OF_SERVICE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataOutOfService) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataOutOfService) GetOutOfService() BACnetApplicationTagBoolean {
	return m.OutOfService
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataOutOfService) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetOutOfService())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataOutOfService(structType any) BACnetConstructedDataOutOfService {
	if casted, ok := structType.(BACnetConstructedDataOutOfService); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataOutOfService); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataOutOfService) GetTypeName() string {
	return "BACnetConstructedDataOutOfService"
}

func (m *_BACnetConstructedDataOutOfService) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (outOfService)
	lengthInBits += m.OutOfService.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataOutOfService) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataOutOfService) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataOutOfService BACnetConstructedDataOutOfService, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataOutOfService"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataOutOfService")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	outOfService, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "outOfService", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'outOfService' field"))
	}
	m.OutOfService = outOfService

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), outOfService)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataOutOfService"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataOutOfService")
	}

	return m, nil
}

func (m *_BACnetConstructedDataOutOfService) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataOutOfService) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataOutOfService"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataOutOfService")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "outOfService", m.GetOutOfService(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'outOfService' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataOutOfService"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataOutOfService")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataOutOfService) IsBACnetConstructedDataOutOfService() {}

func (m *_BACnetConstructedDataOutOfService) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataOutOfService) deepCopy() *_BACnetConstructedDataOutOfService {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataOutOfServiceCopy := &_BACnetConstructedDataOutOfService{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagBoolean](m.OutOfService),
	}
	_BACnetConstructedDataOutOfServiceCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataOutOfServiceCopy
}

func (m *_BACnetConstructedDataOutOfService) String() string {
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
