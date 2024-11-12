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

// HistoryUpdateRequest is the corresponding interface of HistoryUpdateRequest
type HistoryUpdateRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() RequestHeader
	// GetHistoryUpdateDetails returns HistoryUpdateDetails (property field)
	GetHistoryUpdateDetails() []ExtensionObject
	// IsHistoryUpdateRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsHistoryUpdateRequest()
	// CreateBuilder creates a HistoryUpdateRequestBuilder
	CreateHistoryUpdateRequestBuilder() HistoryUpdateRequestBuilder
}

// _HistoryUpdateRequest is the data-structure of this message
type _HistoryUpdateRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader        RequestHeader
	HistoryUpdateDetails []ExtensionObject
}

var _ HistoryUpdateRequest = (*_HistoryUpdateRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_HistoryUpdateRequest)(nil)

// NewHistoryUpdateRequest factory function for _HistoryUpdateRequest
func NewHistoryUpdateRequest(requestHeader RequestHeader, historyUpdateDetails []ExtensionObject) *_HistoryUpdateRequest {
	if requestHeader == nil {
		panic("requestHeader of type RequestHeader for HistoryUpdateRequest must not be nil")
	}
	_result := &_HistoryUpdateRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		HistoryUpdateDetails:              historyUpdateDetails,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// HistoryUpdateRequestBuilder is a builder for HistoryUpdateRequest
type HistoryUpdateRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestHeader RequestHeader, historyUpdateDetails []ExtensionObject) HistoryUpdateRequestBuilder
	// WithRequestHeader adds RequestHeader (property field)
	WithRequestHeader(RequestHeader) HistoryUpdateRequestBuilder
	// WithRequestHeaderBuilder adds RequestHeader (property field) which is build by the builder
	WithRequestHeaderBuilder(func(RequestHeaderBuilder) RequestHeaderBuilder) HistoryUpdateRequestBuilder
	// WithHistoryUpdateDetails adds HistoryUpdateDetails (property field)
	WithHistoryUpdateDetails(...ExtensionObject) HistoryUpdateRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the HistoryUpdateRequest or returns an error if something is wrong
	Build() (HistoryUpdateRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() HistoryUpdateRequest
}

// NewHistoryUpdateRequestBuilder() creates a HistoryUpdateRequestBuilder
func NewHistoryUpdateRequestBuilder() HistoryUpdateRequestBuilder {
	return &_HistoryUpdateRequestBuilder{_HistoryUpdateRequest: new(_HistoryUpdateRequest)}
}

type _HistoryUpdateRequestBuilder struct {
	*_HistoryUpdateRequest

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (HistoryUpdateRequestBuilder) = (*_HistoryUpdateRequestBuilder)(nil)

func (b *_HistoryUpdateRequestBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._HistoryUpdateRequest
}

func (b *_HistoryUpdateRequestBuilder) WithMandatoryFields(requestHeader RequestHeader, historyUpdateDetails []ExtensionObject) HistoryUpdateRequestBuilder {
	return b.WithRequestHeader(requestHeader).WithHistoryUpdateDetails(historyUpdateDetails...)
}

func (b *_HistoryUpdateRequestBuilder) WithRequestHeader(requestHeader RequestHeader) HistoryUpdateRequestBuilder {
	b.RequestHeader = requestHeader
	return b
}

func (b *_HistoryUpdateRequestBuilder) WithRequestHeaderBuilder(builderSupplier func(RequestHeaderBuilder) RequestHeaderBuilder) HistoryUpdateRequestBuilder {
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

func (b *_HistoryUpdateRequestBuilder) WithHistoryUpdateDetails(historyUpdateDetails ...ExtensionObject) HistoryUpdateRequestBuilder {
	b.HistoryUpdateDetails = historyUpdateDetails
	return b
}

func (b *_HistoryUpdateRequestBuilder) Build() (HistoryUpdateRequest, error) {
	if b.RequestHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'requestHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._HistoryUpdateRequest.deepCopy(), nil
}

func (b *_HistoryUpdateRequestBuilder) MustBuild() HistoryUpdateRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_HistoryUpdateRequestBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_HistoryUpdateRequestBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_HistoryUpdateRequestBuilder) DeepCopy() any {
	_copy := b.CreateHistoryUpdateRequestBuilder().(*_HistoryUpdateRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateHistoryUpdateRequestBuilder creates a HistoryUpdateRequestBuilder
func (b *_HistoryUpdateRequest) CreateHistoryUpdateRequestBuilder() HistoryUpdateRequestBuilder {
	if b == nil {
		return NewHistoryUpdateRequestBuilder()
	}
	return &_HistoryUpdateRequestBuilder{_HistoryUpdateRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_HistoryUpdateRequest) GetExtensionId() int32 {
	return int32(700)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_HistoryUpdateRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HistoryUpdateRequest) GetRequestHeader() RequestHeader {
	return m.RequestHeader
}

func (m *_HistoryUpdateRequest) GetHistoryUpdateDetails() []ExtensionObject {
	return m.HistoryUpdateDetails
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastHistoryUpdateRequest(structType any) HistoryUpdateRequest {
	if casted, ok := structType.(HistoryUpdateRequest); ok {
		return casted
	}
	if casted, ok := structType.(*HistoryUpdateRequest); ok {
		return *casted
	}
	return nil
}

func (m *_HistoryUpdateRequest) GetTypeName() string {
	return "HistoryUpdateRequest"
}

func (m *_HistoryUpdateRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Implicit Field (noOfHistoryUpdateDetails)
	lengthInBits += 32

	// Array field
	if len(m.HistoryUpdateDetails) > 0 {
		for _curItem, element := range m.HistoryUpdateDetails {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.HistoryUpdateDetails), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_HistoryUpdateRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_HistoryUpdateRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__historyUpdateRequest HistoryUpdateRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("HistoryUpdateRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HistoryUpdateRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[RequestHeader](ctx, "requestHeader", ReadComplex[RequestHeader](ExtensionObjectDefinitionParseWithBufferProducer[RequestHeader]((int32)(int32(391))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	noOfHistoryUpdateDetails, err := ReadImplicitField[int32](ctx, "noOfHistoryUpdateDetails", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfHistoryUpdateDetails' field"))
	}
	_ = noOfHistoryUpdateDetails

	historyUpdateDetails, err := ReadCountArrayField[ExtensionObject](ctx, "historyUpdateDetails", ReadComplex[ExtensionObject](ExtensionObjectParseWithBufferProducer[ExtensionObject]((bool)(bool(true))), readBuffer), uint64(noOfHistoryUpdateDetails))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'historyUpdateDetails' field"))
	}
	m.HistoryUpdateDetails = historyUpdateDetails

	if closeErr := readBuffer.CloseContext("HistoryUpdateRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HistoryUpdateRequest")
	}

	return m, nil
}

func (m *_HistoryUpdateRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HistoryUpdateRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("HistoryUpdateRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for HistoryUpdateRequest")
		}

		if err := WriteSimpleField[RequestHeader](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[RequestHeader](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}
		noOfHistoryUpdateDetails := int32(utils.InlineIf(bool((m.GetHistoryUpdateDetails()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetHistoryUpdateDetails()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfHistoryUpdateDetails", noOfHistoryUpdateDetails, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfHistoryUpdateDetails' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "historyUpdateDetails", m.GetHistoryUpdateDetails(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'historyUpdateDetails' field")
		}

		if popErr := writeBuffer.PopContext("HistoryUpdateRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for HistoryUpdateRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_HistoryUpdateRequest) IsHistoryUpdateRequest() {}

func (m *_HistoryUpdateRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_HistoryUpdateRequest) deepCopy() *_HistoryUpdateRequest {
	if m == nil {
		return nil
	}
	_HistoryUpdateRequestCopy := &_HistoryUpdateRequest{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[RequestHeader](m.RequestHeader),
		utils.DeepCopySlice[ExtensionObject, ExtensionObject](m.HistoryUpdateDetails),
	}
	_HistoryUpdateRequestCopy.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _HistoryUpdateRequestCopy
}

func (m *_HistoryUpdateRequest) String() string {
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
