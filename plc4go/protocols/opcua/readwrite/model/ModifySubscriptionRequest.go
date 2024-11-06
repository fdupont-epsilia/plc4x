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

// ModifySubscriptionRequest is the corresponding interface of ModifySubscriptionRequest
type ModifySubscriptionRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() RequestHeader
	// GetSubscriptionId returns SubscriptionId (property field)
	GetSubscriptionId() uint32
	// GetRequestedPublishingInterval returns RequestedPublishingInterval (property field)
	GetRequestedPublishingInterval() float64
	// GetRequestedLifetimeCount returns RequestedLifetimeCount (property field)
	GetRequestedLifetimeCount() uint32
	// GetRequestedMaxKeepAliveCount returns RequestedMaxKeepAliveCount (property field)
	GetRequestedMaxKeepAliveCount() uint32
	// GetMaxNotificationsPerPublish returns MaxNotificationsPerPublish (property field)
	GetMaxNotificationsPerPublish() uint32
	// GetPriority returns Priority (property field)
	GetPriority() uint8
	// IsModifySubscriptionRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModifySubscriptionRequest()
	// CreateBuilder creates a ModifySubscriptionRequestBuilder
	CreateModifySubscriptionRequestBuilder() ModifySubscriptionRequestBuilder
}

// _ModifySubscriptionRequest is the data-structure of this message
type _ModifySubscriptionRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader               RequestHeader
	SubscriptionId              uint32
	RequestedPublishingInterval float64
	RequestedLifetimeCount      uint32
	RequestedMaxKeepAliveCount  uint32
	MaxNotificationsPerPublish  uint32
	Priority                    uint8
}

var _ ModifySubscriptionRequest = (*_ModifySubscriptionRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ModifySubscriptionRequest)(nil)

// NewModifySubscriptionRequest factory function for _ModifySubscriptionRequest
func NewModifySubscriptionRequest(requestHeader RequestHeader, subscriptionId uint32, requestedPublishingInterval float64, requestedLifetimeCount uint32, requestedMaxKeepAliveCount uint32, maxNotificationsPerPublish uint32, priority uint8) *_ModifySubscriptionRequest {
	if requestHeader == nil {
		panic("requestHeader of type RequestHeader for ModifySubscriptionRequest must not be nil")
	}
	_result := &_ModifySubscriptionRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		SubscriptionId:                    subscriptionId,
		RequestedPublishingInterval:       requestedPublishingInterval,
		RequestedLifetimeCount:            requestedLifetimeCount,
		RequestedMaxKeepAliveCount:        requestedMaxKeepAliveCount,
		MaxNotificationsPerPublish:        maxNotificationsPerPublish,
		Priority:                          priority,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ModifySubscriptionRequestBuilder is a builder for ModifySubscriptionRequest
type ModifySubscriptionRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestHeader RequestHeader, subscriptionId uint32, requestedPublishingInterval float64, requestedLifetimeCount uint32, requestedMaxKeepAliveCount uint32, maxNotificationsPerPublish uint32, priority uint8) ModifySubscriptionRequestBuilder
	// WithRequestHeader adds RequestHeader (property field)
	WithRequestHeader(RequestHeader) ModifySubscriptionRequestBuilder
	// WithRequestHeaderBuilder adds RequestHeader (property field) which is build by the builder
	WithRequestHeaderBuilder(func(RequestHeaderBuilder) RequestHeaderBuilder) ModifySubscriptionRequestBuilder
	// WithSubscriptionId adds SubscriptionId (property field)
	WithSubscriptionId(uint32) ModifySubscriptionRequestBuilder
	// WithRequestedPublishingInterval adds RequestedPublishingInterval (property field)
	WithRequestedPublishingInterval(float64) ModifySubscriptionRequestBuilder
	// WithRequestedLifetimeCount adds RequestedLifetimeCount (property field)
	WithRequestedLifetimeCount(uint32) ModifySubscriptionRequestBuilder
	// WithRequestedMaxKeepAliveCount adds RequestedMaxKeepAliveCount (property field)
	WithRequestedMaxKeepAliveCount(uint32) ModifySubscriptionRequestBuilder
	// WithMaxNotificationsPerPublish adds MaxNotificationsPerPublish (property field)
	WithMaxNotificationsPerPublish(uint32) ModifySubscriptionRequestBuilder
	// WithPriority adds Priority (property field)
	WithPriority(uint8) ModifySubscriptionRequestBuilder
	// Build builds the ModifySubscriptionRequest or returns an error if something is wrong
	Build() (ModifySubscriptionRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ModifySubscriptionRequest
}

// NewModifySubscriptionRequestBuilder() creates a ModifySubscriptionRequestBuilder
func NewModifySubscriptionRequestBuilder() ModifySubscriptionRequestBuilder {
	return &_ModifySubscriptionRequestBuilder{_ModifySubscriptionRequest: new(_ModifySubscriptionRequest)}
}

type _ModifySubscriptionRequestBuilder struct {
	*_ModifySubscriptionRequest

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (ModifySubscriptionRequestBuilder) = (*_ModifySubscriptionRequestBuilder)(nil)

func (b *_ModifySubscriptionRequestBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_ModifySubscriptionRequestBuilder) WithMandatoryFields(requestHeader RequestHeader, subscriptionId uint32, requestedPublishingInterval float64, requestedLifetimeCount uint32, requestedMaxKeepAliveCount uint32, maxNotificationsPerPublish uint32, priority uint8) ModifySubscriptionRequestBuilder {
	return b.WithRequestHeader(requestHeader).WithSubscriptionId(subscriptionId).WithRequestedPublishingInterval(requestedPublishingInterval).WithRequestedLifetimeCount(requestedLifetimeCount).WithRequestedMaxKeepAliveCount(requestedMaxKeepAliveCount).WithMaxNotificationsPerPublish(maxNotificationsPerPublish).WithPriority(priority)
}

func (b *_ModifySubscriptionRequestBuilder) WithRequestHeader(requestHeader RequestHeader) ModifySubscriptionRequestBuilder {
	b.RequestHeader = requestHeader
	return b
}

func (b *_ModifySubscriptionRequestBuilder) WithRequestHeaderBuilder(builderSupplier func(RequestHeaderBuilder) RequestHeaderBuilder) ModifySubscriptionRequestBuilder {
	builder := builderSupplier(b.RequestHeader.CreateRequestHeaderBuilder())
	var err error
	b.RequestHeader, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "RequestHeaderBuilder failed"))
	}
	return b
}

func (b *_ModifySubscriptionRequestBuilder) WithSubscriptionId(subscriptionId uint32) ModifySubscriptionRequestBuilder {
	b.SubscriptionId = subscriptionId
	return b
}

func (b *_ModifySubscriptionRequestBuilder) WithRequestedPublishingInterval(requestedPublishingInterval float64) ModifySubscriptionRequestBuilder {
	b.RequestedPublishingInterval = requestedPublishingInterval
	return b
}

func (b *_ModifySubscriptionRequestBuilder) WithRequestedLifetimeCount(requestedLifetimeCount uint32) ModifySubscriptionRequestBuilder {
	b.RequestedLifetimeCount = requestedLifetimeCount
	return b
}

func (b *_ModifySubscriptionRequestBuilder) WithRequestedMaxKeepAliveCount(requestedMaxKeepAliveCount uint32) ModifySubscriptionRequestBuilder {
	b.RequestedMaxKeepAliveCount = requestedMaxKeepAliveCount
	return b
}

func (b *_ModifySubscriptionRequestBuilder) WithMaxNotificationsPerPublish(maxNotificationsPerPublish uint32) ModifySubscriptionRequestBuilder {
	b.MaxNotificationsPerPublish = maxNotificationsPerPublish
	return b
}

func (b *_ModifySubscriptionRequestBuilder) WithPriority(priority uint8) ModifySubscriptionRequestBuilder {
	b.Priority = priority
	return b
}

func (b *_ModifySubscriptionRequestBuilder) Build() (ModifySubscriptionRequest, error) {
	if b.RequestHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'requestHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ModifySubscriptionRequest.deepCopy(), nil
}

func (b *_ModifySubscriptionRequestBuilder) MustBuild() ModifySubscriptionRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_ModifySubscriptionRequestBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_ModifySubscriptionRequestBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_ModifySubscriptionRequestBuilder) DeepCopy() any {
	_copy := b.CreateModifySubscriptionRequestBuilder().(*_ModifySubscriptionRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateModifySubscriptionRequestBuilder creates a ModifySubscriptionRequestBuilder
func (b *_ModifySubscriptionRequest) CreateModifySubscriptionRequestBuilder() ModifySubscriptionRequestBuilder {
	if b == nil {
		return NewModifySubscriptionRequestBuilder()
	}
	return &_ModifySubscriptionRequestBuilder{_ModifySubscriptionRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModifySubscriptionRequest) GetExtensionId() int32 {
	return int32(793)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModifySubscriptionRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModifySubscriptionRequest) GetRequestHeader() RequestHeader {
	return m.RequestHeader
}

func (m *_ModifySubscriptionRequest) GetSubscriptionId() uint32 {
	return m.SubscriptionId
}

func (m *_ModifySubscriptionRequest) GetRequestedPublishingInterval() float64 {
	return m.RequestedPublishingInterval
}

func (m *_ModifySubscriptionRequest) GetRequestedLifetimeCount() uint32 {
	return m.RequestedLifetimeCount
}

func (m *_ModifySubscriptionRequest) GetRequestedMaxKeepAliveCount() uint32 {
	return m.RequestedMaxKeepAliveCount
}

func (m *_ModifySubscriptionRequest) GetMaxNotificationsPerPublish() uint32 {
	return m.MaxNotificationsPerPublish
}

func (m *_ModifySubscriptionRequest) GetPriority() uint8 {
	return m.Priority
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastModifySubscriptionRequest(structType any) ModifySubscriptionRequest {
	if casted, ok := structType.(ModifySubscriptionRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ModifySubscriptionRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ModifySubscriptionRequest) GetTypeName() string {
	return "ModifySubscriptionRequest"
}

func (m *_ModifySubscriptionRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (subscriptionId)
	lengthInBits += 32

	// Simple field (requestedPublishingInterval)
	lengthInBits += 64

	// Simple field (requestedLifetimeCount)
	lengthInBits += 32

	// Simple field (requestedMaxKeepAliveCount)
	lengthInBits += 32

	// Simple field (maxNotificationsPerPublish)
	lengthInBits += 32

	// Simple field (priority)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ModifySubscriptionRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModifySubscriptionRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__modifySubscriptionRequest ModifySubscriptionRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModifySubscriptionRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModifySubscriptionRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[RequestHeader](ctx, "requestHeader", ReadComplex[RequestHeader](ExtensionObjectDefinitionParseWithBufferProducer[RequestHeader]((int32)(int32(391))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	subscriptionId, err := ReadSimpleField(ctx, "subscriptionId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subscriptionId' field"))
	}
	m.SubscriptionId = subscriptionId

	requestedPublishingInterval, err := ReadSimpleField(ctx, "requestedPublishingInterval", ReadDouble(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestedPublishingInterval' field"))
	}
	m.RequestedPublishingInterval = requestedPublishingInterval

	requestedLifetimeCount, err := ReadSimpleField(ctx, "requestedLifetimeCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestedLifetimeCount' field"))
	}
	m.RequestedLifetimeCount = requestedLifetimeCount

	requestedMaxKeepAliveCount, err := ReadSimpleField(ctx, "requestedMaxKeepAliveCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestedMaxKeepAliveCount' field"))
	}
	m.RequestedMaxKeepAliveCount = requestedMaxKeepAliveCount

	maxNotificationsPerPublish, err := ReadSimpleField(ctx, "maxNotificationsPerPublish", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxNotificationsPerPublish' field"))
	}
	m.MaxNotificationsPerPublish = maxNotificationsPerPublish

	priority, err := ReadSimpleField(ctx, "priority", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priority' field"))
	}
	m.Priority = priority

	if closeErr := readBuffer.CloseContext("ModifySubscriptionRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModifySubscriptionRequest")
	}

	return m, nil
}

func (m *_ModifySubscriptionRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModifySubscriptionRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModifySubscriptionRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModifySubscriptionRequest")
		}

		if err := WriteSimpleField[RequestHeader](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[RequestHeader](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[uint32](ctx, "subscriptionId", m.GetSubscriptionId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'subscriptionId' field")
		}

		if err := WriteSimpleField[float64](ctx, "requestedPublishingInterval", m.GetRequestedPublishingInterval(), WriteDouble(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestedPublishingInterval' field")
		}

		if err := WriteSimpleField[uint32](ctx, "requestedLifetimeCount", m.GetRequestedLifetimeCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestedLifetimeCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "requestedMaxKeepAliveCount", m.GetRequestedMaxKeepAliveCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestedMaxKeepAliveCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "maxNotificationsPerPublish", m.GetMaxNotificationsPerPublish(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxNotificationsPerPublish' field")
		}

		if err := WriteSimpleField[uint8](ctx, "priority", m.GetPriority(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'priority' field")
		}

		if popErr := writeBuffer.PopContext("ModifySubscriptionRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModifySubscriptionRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModifySubscriptionRequest) IsModifySubscriptionRequest() {}

func (m *_ModifySubscriptionRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModifySubscriptionRequest) deepCopy() *_ModifySubscriptionRequest {
	if m == nil {
		return nil
	}
	_ModifySubscriptionRequestCopy := &_ModifySubscriptionRequest{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.RequestHeader.DeepCopy().(RequestHeader),
		m.SubscriptionId,
		m.RequestedPublishingInterval,
		m.RequestedLifetimeCount,
		m.RequestedMaxKeepAliveCount,
		m.MaxNotificationsPerPublish,
		m.Priority,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _ModifySubscriptionRequestCopy
}

func (m *_ModifySubscriptionRequest) String() string {
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
