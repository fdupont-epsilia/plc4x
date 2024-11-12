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

// BrokerDataSetWriterTransportDataType is the corresponding interface of BrokerDataSetWriterTransportDataType
type BrokerDataSetWriterTransportDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetQueueName returns QueueName (property field)
	GetQueueName() PascalString
	// GetResourceUri returns ResourceUri (property field)
	GetResourceUri() PascalString
	// GetAuthenticationProfileUri returns AuthenticationProfileUri (property field)
	GetAuthenticationProfileUri() PascalString
	// GetRequestedDeliveryGuarantee returns RequestedDeliveryGuarantee (property field)
	GetRequestedDeliveryGuarantee() BrokerTransportQualityOfService
	// GetMetaDataQueueName returns MetaDataQueueName (property field)
	GetMetaDataQueueName() PascalString
	// GetMetaDataUpdateTime returns MetaDataUpdateTime (property field)
	GetMetaDataUpdateTime() float64
	// IsBrokerDataSetWriterTransportDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBrokerDataSetWriterTransportDataType()
	// CreateBuilder creates a BrokerDataSetWriterTransportDataTypeBuilder
	CreateBrokerDataSetWriterTransportDataTypeBuilder() BrokerDataSetWriterTransportDataTypeBuilder
}

// _BrokerDataSetWriterTransportDataType is the data-structure of this message
type _BrokerDataSetWriterTransportDataType struct {
	ExtensionObjectDefinitionContract
	QueueName                  PascalString
	ResourceUri                PascalString
	AuthenticationProfileUri   PascalString
	RequestedDeliveryGuarantee BrokerTransportQualityOfService
	MetaDataQueueName          PascalString
	MetaDataUpdateTime         float64
}

var _ BrokerDataSetWriterTransportDataType = (*_BrokerDataSetWriterTransportDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_BrokerDataSetWriterTransportDataType)(nil)

// NewBrokerDataSetWriterTransportDataType factory function for _BrokerDataSetWriterTransportDataType
func NewBrokerDataSetWriterTransportDataType(queueName PascalString, resourceUri PascalString, authenticationProfileUri PascalString, requestedDeliveryGuarantee BrokerTransportQualityOfService, metaDataQueueName PascalString, metaDataUpdateTime float64) *_BrokerDataSetWriterTransportDataType {
	if queueName == nil {
		panic("queueName of type PascalString for BrokerDataSetWriterTransportDataType must not be nil")
	}
	if resourceUri == nil {
		panic("resourceUri of type PascalString for BrokerDataSetWriterTransportDataType must not be nil")
	}
	if authenticationProfileUri == nil {
		panic("authenticationProfileUri of type PascalString for BrokerDataSetWriterTransportDataType must not be nil")
	}
	if metaDataQueueName == nil {
		panic("metaDataQueueName of type PascalString for BrokerDataSetWriterTransportDataType must not be nil")
	}
	_result := &_BrokerDataSetWriterTransportDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		QueueName:                         queueName,
		ResourceUri:                       resourceUri,
		AuthenticationProfileUri:          authenticationProfileUri,
		RequestedDeliveryGuarantee:        requestedDeliveryGuarantee,
		MetaDataQueueName:                 metaDataQueueName,
		MetaDataUpdateTime:                metaDataUpdateTime,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BrokerDataSetWriterTransportDataTypeBuilder is a builder for BrokerDataSetWriterTransportDataType
type BrokerDataSetWriterTransportDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(queueName PascalString, resourceUri PascalString, authenticationProfileUri PascalString, requestedDeliveryGuarantee BrokerTransportQualityOfService, metaDataQueueName PascalString, metaDataUpdateTime float64) BrokerDataSetWriterTransportDataTypeBuilder
	// WithQueueName adds QueueName (property field)
	WithQueueName(PascalString) BrokerDataSetWriterTransportDataTypeBuilder
	// WithQueueNameBuilder adds QueueName (property field) which is build by the builder
	WithQueueNameBuilder(func(PascalStringBuilder) PascalStringBuilder) BrokerDataSetWriterTransportDataTypeBuilder
	// WithResourceUri adds ResourceUri (property field)
	WithResourceUri(PascalString) BrokerDataSetWriterTransportDataTypeBuilder
	// WithResourceUriBuilder adds ResourceUri (property field) which is build by the builder
	WithResourceUriBuilder(func(PascalStringBuilder) PascalStringBuilder) BrokerDataSetWriterTransportDataTypeBuilder
	// WithAuthenticationProfileUri adds AuthenticationProfileUri (property field)
	WithAuthenticationProfileUri(PascalString) BrokerDataSetWriterTransportDataTypeBuilder
	// WithAuthenticationProfileUriBuilder adds AuthenticationProfileUri (property field) which is build by the builder
	WithAuthenticationProfileUriBuilder(func(PascalStringBuilder) PascalStringBuilder) BrokerDataSetWriterTransportDataTypeBuilder
	// WithRequestedDeliveryGuarantee adds RequestedDeliveryGuarantee (property field)
	WithRequestedDeliveryGuarantee(BrokerTransportQualityOfService) BrokerDataSetWriterTransportDataTypeBuilder
	// WithMetaDataQueueName adds MetaDataQueueName (property field)
	WithMetaDataQueueName(PascalString) BrokerDataSetWriterTransportDataTypeBuilder
	// WithMetaDataQueueNameBuilder adds MetaDataQueueName (property field) which is build by the builder
	WithMetaDataQueueNameBuilder(func(PascalStringBuilder) PascalStringBuilder) BrokerDataSetWriterTransportDataTypeBuilder
	// WithMetaDataUpdateTime adds MetaDataUpdateTime (property field)
	WithMetaDataUpdateTime(float64) BrokerDataSetWriterTransportDataTypeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the BrokerDataSetWriterTransportDataType or returns an error if something is wrong
	Build() (BrokerDataSetWriterTransportDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BrokerDataSetWriterTransportDataType
}

// NewBrokerDataSetWriterTransportDataTypeBuilder() creates a BrokerDataSetWriterTransportDataTypeBuilder
func NewBrokerDataSetWriterTransportDataTypeBuilder() BrokerDataSetWriterTransportDataTypeBuilder {
	return &_BrokerDataSetWriterTransportDataTypeBuilder{_BrokerDataSetWriterTransportDataType: new(_BrokerDataSetWriterTransportDataType)}
}

type _BrokerDataSetWriterTransportDataTypeBuilder struct {
	*_BrokerDataSetWriterTransportDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (BrokerDataSetWriterTransportDataTypeBuilder) = (*_BrokerDataSetWriterTransportDataTypeBuilder)(nil)

func (b *_BrokerDataSetWriterTransportDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._BrokerDataSetWriterTransportDataType
}

func (b *_BrokerDataSetWriterTransportDataTypeBuilder) WithMandatoryFields(queueName PascalString, resourceUri PascalString, authenticationProfileUri PascalString, requestedDeliveryGuarantee BrokerTransportQualityOfService, metaDataQueueName PascalString, metaDataUpdateTime float64) BrokerDataSetWriterTransportDataTypeBuilder {
	return b.WithQueueName(queueName).WithResourceUri(resourceUri).WithAuthenticationProfileUri(authenticationProfileUri).WithRequestedDeliveryGuarantee(requestedDeliveryGuarantee).WithMetaDataQueueName(metaDataQueueName).WithMetaDataUpdateTime(metaDataUpdateTime)
}

func (b *_BrokerDataSetWriterTransportDataTypeBuilder) WithQueueName(queueName PascalString) BrokerDataSetWriterTransportDataTypeBuilder {
	b.QueueName = queueName
	return b
}

func (b *_BrokerDataSetWriterTransportDataTypeBuilder) WithQueueNameBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) BrokerDataSetWriterTransportDataTypeBuilder {
	builder := builderSupplier(b.QueueName.CreatePascalStringBuilder())
	var err error
	b.QueueName, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_BrokerDataSetWriterTransportDataTypeBuilder) WithResourceUri(resourceUri PascalString) BrokerDataSetWriterTransportDataTypeBuilder {
	b.ResourceUri = resourceUri
	return b
}

func (b *_BrokerDataSetWriterTransportDataTypeBuilder) WithResourceUriBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) BrokerDataSetWriterTransportDataTypeBuilder {
	builder := builderSupplier(b.ResourceUri.CreatePascalStringBuilder())
	var err error
	b.ResourceUri, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_BrokerDataSetWriterTransportDataTypeBuilder) WithAuthenticationProfileUri(authenticationProfileUri PascalString) BrokerDataSetWriterTransportDataTypeBuilder {
	b.AuthenticationProfileUri = authenticationProfileUri
	return b
}

func (b *_BrokerDataSetWriterTransportDataTypeBuilder) WithAuthenticationProfileUriBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) BrokerDataSetWriterTransportDataTypeBuilder {
	builder := builderSupplier(b.AuthenticationProfileUri.CreatePascalStringBuilder())
	var err error
	b.AuthenticationProfileUri, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_BrokerDataSetWriterTransportDataTypeBuilder) WithRequestedDeliveryGuarantee(requestedDeliveryGuarantee BrokerTransportQualityOfService) BrokerDataSetWriterTransportDataTypeBuilder {
	b.RequestedDeliveryGuarantee = requestedDeliveryGuarantee
	return b
}

func (b *_BrokerDataSetWriterTransportDataTypeBuilder) WithMetaDataQueueName(metaDataQueueName PascalString) BrokerDataSetWriterTransportDataTypeBuilder {
	b.MetaDataQueueName = metaDataQueueName
	return b
}

func (b *_BrokerDataSetWriterTransportDataTypeBuilder) WithMetaDataQueueNameBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) BrokerDataSetWriterTransportDataTypeBuilder {
	builder := builderSupplier(b.MetaDataQueueName.CreatePascalStringBuilder())
	var err error
	b.MetaDataQueueName, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_BrokerDataSetWriterTransportDataTypeBuilder) WithMetaDataUpdateTime(metaDataUpdateTime float64) BrokerDataSetWriterTransportDataTypeBuilder {
	b.MetaDataUpdateTime = metaDataUpdateTime
	return b
}

func (b *_BrokerDataSetWriterTransportDataTypeBuilder) Build() (BrokerDataSetWriterTransportDataType, error) {
	if b.QueueName == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'queueName' not set"))
	}
	if b.ResourceUri == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'resourceUri' not set"))
	}
	if b.AuthenticationProfileUri == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'authenticationProfileUri' not set"))
	}
	if b.MetaDataQueueName == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'metaDataQueueName' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BrokerDataSetWriterTransportDataType.deepCopy(), nil
}

func (b *_BrokerDataSetWriterTransportDataTypeBuilder) MustBuild() BrokerDataSetWriterTransportDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BrokerDataSetWriterTransportDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_BrokerDataSetWriterTransportDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_BrokerDataSetWriterTransportDataTypeBuilder) DeepCopy() any {
	_copy := b.CreateBrokerDataSetWriterTransportDataTypeBuilder().(*_BrokerDataSetWriterTransportDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBrokerDataSetWriterTransportDataTypeBuilder creates a BrokerDataSetWriterTransportDataTypeBuilder
func (b *_BrokerDataSetWriterTransportDataType) CreateBrokerDataSetWriterTransportDataTypeBuilder() BrokerDataSetWriterTransportDataTypeBuilder {
	if b == nil {
		return NewBrokerDataSetWriterTransportDataTypeBuilder()
	}
	return &_BrokerDataSetWriterTransportDataTypeBuilder{_BrokerDataSetWriterTransportDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BrokerDataSetWriterTransportDataType) GetExtensionId() int32 {
	return int32(15671)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BrokerDataSetWriterTransportDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BrokerDataSetWriterTransportDataType) GetQueueName() PascalString {
	return m.QueueName
}

func (m *_BrokerDataSetWriterTransportDataType) GetResourceUri() PascalString {
	return m.ResourceUri
}

func (m *_BrokerDataSetWriterTransportDataType) GetAuthenticationProfileUri() PascalString {
	return m.AuthenticationProfileUri
}

func (m *_BrokerDataSetWriterTransportDataType) GetRequestedDeliveryGuarantee() BrokerTransportQualityOfService {
	return m.RequestedDeliveryGuarantee
}

func (m *_BrokerDataSetWriterTransportDataType) GetMetaDataQueueName() PascalString {
	return m.MetaDataQueueName
}

func (m *_BrokerDataSetWriterTransportDataType) GetMetaDataUpdateTime() float64 {
	return m.MetaDataUpdateTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBrokerDataSetWriterTransportDataType(structType any) BrokerDataSetWriterTransportDataType {
	if casted, ok := structType.(BrokerDataSetWriterTransportDataType); ok {
		return casted
	}
	if casted, ok := structType.(*BrokerDataSetWriterTransportDataType); ok {
		return *casted
	}
	return nil
}

func (m *_BrokerDataSetWriterTransportDataType) GetTypeName() string {
	return "BrokerDataSetWriterTransportDataType"
}

func (m *_BrokerDataSetWriterTransportDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (queueName)
	lengthInBits += m.QueueName.GetLengthInBits(ctx)

	// Simple field (resourceUri)
	lengthInBits += m.ResourceUri.GetLengthInBits(ctx)

	// Simple field (authenticationProfileUri)
	lengthInBits += m.AuthenticationProfileUri.GetLengthInBits(ctx)

	// Simple field (requestedDeliveryGuarantee)
	lengthInBits += 32

	// Simple field (metaDataQueueName)
	lengthInBits += m.MetaDataQueueName.GetLengthInBits(ctx)

	// Simple field (metaDataUpdateTime)
	lengthInBits += 64

	return lengthInBits
}

func (m *_BrokerDataSetWriterTransportDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BrokerDataSetWriterTransportDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__brokerDataSetWriterTransportDataType BrokerDataSetWriterTransportDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BrokerDataSetWriterTransportDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BrokerDataSetWriterTransportDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	queueName, err := ReadSimpleField[PascalString](ctx, "queueName", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'queueName' field"))
	}
	m.QueueName = queueName

	resourceUri, err := ReadSimpleField[PascalString](ctx, "resourceUri", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'resourceUri' field"))
	}
	m.ResourceUri = resourceUri

	authenticationProfileUri, err := ReadSimpleField[PascalString](ctx, "authenticationProfileUri", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'authenticationProfileUri' field"))
	}
	m.AuthenticationProfileUri = authenticationProfileUri

	requestedDeliveryGuarantee, err := ReadEnumField[BrokerTransportQualityOfService](ctx, "requestedDeliveryGuarantee", "BrokerTransportQualityOfService", ReadEnum(BrokerTransportQualityOfServiceByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestedDeliveryGuarantee' field"))
	}
	m.RequestedDeliveryGuarantee = requestedDeliveryGuarantee

	metaDataQueueName, err := ReadSimpleField[PascalString](ctx, "metaDataQueueName", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'metaDataQueueName' field"))
	}
	m.MetaDataQueueName = metaDataQueueName

	metaDataUpdateTime, err := ReadSimpleField(ctx, "metaDataUpdateTime", ReadDouble(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'metaDataUpdateTime' field"))
	}
	m.MetaDataUpdateTime = metaDataUpdateTime

	if closeErr := readBuffer.CloseContext("BrokerDataSetWriterTransportDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BrokerDataSetWriterTransportDataType")
	}

	return m, nil
}

func (m *_BrokerDataSetWriterTransportDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BrokerDataSetWriterTransportDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BrokerDataSetWriterTransportDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BrokerDataSetWriterTransportDataType")
		}

		if err := WriteSimpleField[PascalString](ctx, "queueName", m.GetQueueName(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'queueName' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "resourceUri", m.GetResourceUri(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'resourceUri' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "authenticationProfileUri", m.GetAuthenticationProfileUri(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'authenticationProfileUri' field")
		}

		if err := WriteSimpleEnumField[BrokerTransportQualityOfService](ctx, "requestedDeliveryGuarantee", "BrokerTransportQualityOfService", m.GetRequestedDeliveryGuarantee(), WriteEnum[BrokerTransportQualityOfService, uint32](BrokerTransportQualityOfService.GetValue, BrokerTransportQualityOfService.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'requestedDeliveryGuarantee' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "metaDataQueueName", m.GetMetaDataQueueName(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'metaDataQueueName' field")
		}

		if err := WriteSimpleField[float64](ctx, "metaDataUpdateTime", m.GetMetaDataUpdateTime(), WriteDouble(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'metaDataUpdateTime' field")
		}

		if popErr := writeBuffer.PopContext("BrokerDataSetWriterTransportDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BrokerDataSetWriterTransportDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BrokerDataSetWriterTransportDataType) IsBrokerDataSetWriterTransportDataType() {}

func (m *_BrokerDataSetWriterTransportDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BrokerDataSetWriterTransportDataType) deepCopy() *_BrokerDataSetWriterTransportDataType {
	if m == nil {
		return nil
	}
	_BrokerDataSetWriterTransportDataTypeCopy := &_BrokerDataSetWriterTransportDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[PascalString](m.QueueName),
		utils.DeepCopy[PascalString](m.ResourceUri),
		utils.DeepCopy[PascalString](m.AuthenticationProfileUri),
		m.RequestedDeliveryGuarantee,
		utils.DeepCopy[PascalString](m.MetaDataQueueName),
		m.MetaDataUpdateTime,
	}
	_BrokerDataSetWriterTransportDataTypeCopy.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _BrokerDataSetWriterTransportDataTypeCopy
}

func (m *_BrokerDataSetWriterTransportDataType) String() string {
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
