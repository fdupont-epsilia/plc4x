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

// CallRequest is the corresponding interface of CallRequest
type CallRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetNoOfMethodsToCall returns NoOfMethodsToCall (property field)
	GetNoOfMethodsToCall() int32
	// GetMethodsToCall returns MethodsToCall (property field)
	GetMethodsToCall() []ExtensionObjectDefinition
}

// CallRequestExactly can be used when we want exactly this type and not a type which fulfills CallRequest.
// This is useful for switch cases.
type CallRequestExactly interface {
	CallRequest
	isCallRequest() bool
}

// _CallRequest is the data-structure of this message
type _CallRequest struct {
	*_ExtensionObjectDefinition
	RequestHeader     ExtensionObjectDefinition
	NoOfMethodsToCall int32
	MethodsToCall     []ExtensionObjectDefinition
}

var _ CallRequest = (*_CallRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CallRequest) GetIdentifier() string {
	return "712"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CallRequest) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_CallRequest) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CallRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_CallRequest) GetNoOfMethodsToCall() int32 {
	return m.NoOfMethodsToCall
}

func (m *_CallRequest) GetMethodsToCall() []ExtensionObjectDefinition {
	return m.MethodsToCall
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCallRequest factory function for _CallRequest
func NewCallRequest(requestHeader ExtensionObjectDefinition, noOfMethodsToCall int32, methodsToCall []ExtensionObjectDefinition) *_CallRequest {
	_result := &_CallRequest{
		RequestHeader:              requestHeader,
		NoOfMethodsToCall:          noOfMethodsToCall,
		MethodsToCall:              methodsToCall,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCallRequest(structType any) CallRequest {
	if casted, ok := structType.(CallRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CallRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CallRequest) GetTypeName() string {
	return "CallRequest"
}

func (m *_CallRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (noOfMethodsToCall)
	lengthInBits += 32

	// Array field
	if len(m.MethodsToCall) > 0 {
		for _curItem, element := range m.MethodsToCall {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.MethodsToCall), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_CallRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CallRequestParse(ctx context.Context, theBytes []byte, identifier string) (CallRequest, error) {
	return CallRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func CallRequestParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (CallRequest, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (CallRequest, error) {
		return CallRequestParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func CallRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (CallRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CallRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CallRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}

	noOfMethodsToCall, err := ReadSimpleField(ctx, "noOfMethodsToCall", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfMethodsToCall' field"))
	}

	methodsToCall, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "methodsToCall", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("706")), readBuffer), uint64(noOfMethodsToCall))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'methodsToCall' field"))
	}

	if closeErr := readBuffer.CloseContext("CallRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CallRequest")
	}

	// Create a partially initialized instance
	_child := &_CallRequest{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		RequestHeader:              requestHeader,
		NoOfMethodsToCall:          noOfMethodsToCall,
		MethodsToCall:              methodsToCall,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_CallRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CallRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CallRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CallRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfMethodsToCall", m.GetNoOfMethodsToCall(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfMethodsToCall' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "methodsToCall", m.GetMethodsToCall(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'methodsToCall' field")
		}

		if popErr := writeBuffer.PopContext("CallRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CallRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CallRequest) isCallRequest() bool {
	return true
}

func (m *_CallRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
