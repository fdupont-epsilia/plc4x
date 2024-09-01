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

// WriteRequest is the corresponding interface of WriteRequest
type WriteRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetNoOfNodesToWrite returns NoOfNodesToWrite (property field)
	GetNoOfNodesToWrite() int32
	// GetNodesToWrite returns NodesToWrite (property field)
	GetNodesToWrite() []ExtensionObjectDefinition
}

// WriteRequestExactly can be used when we want exactly this type and not a type which fulfills WriteRequest.
// This is useful for switch cases.
type WriteRequestExactly interface {
	WriteRequest
	isWriteRequest() bool
}

// _WriteRequest is the data-structure of this message
type _WriteRequest struct {
	*_ExtensionObjectDefinition
	RequestHeader    ExtensionObjectDefinition
	NoOfNodesToWrite int32
	NodesToWrite     []ExtensionObjectDefinition
}

var _ WriteRequest = (*_WriteRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_WriteRequest) GetIdentifier() string {
	return "673"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_WriteRequest) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_WriteRequest) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_WriteRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_WriteRequest) GetNoOfNodesToWrite() int32 {
	return m.NoOfNodesToWrite
}

func (m *_WriteRequest) GetNodesToWrite() []ExtensionObjectDefinition {
	return m.NodesToWrite
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewWriteRequest factory function for _WriteRequest
func NewWriteRequest(requestHeader ExtensionObjectDefinition, noOfNodesToWrite int32, nodesToWrite []ExtensionObjectDefinition) *_WriteRequest {
	_result := &_WriteRequest{
		RequestHeader:              requestHeader,
		NoOfNodesToWrite:           noOfNodesToWrite,
		NodesToWrite:               nodesToWrite,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastWriteRequest(structType any) WriteRequest {
	if casted, ok := structType.(WriteRequest); ok {
		return casted
	}
	if casted, ok := structType.(*WriteRequest); ok {
		return *casted
	}
	return nil
}

func (m *_WriteRequest) GetTypeName() string {
	return "WriteRequest"
}

func (m *_WriteRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (noOfNodesToWrite)
	lengthInBits += 32

	// Array field
	if len(m.NodesToWrite) > 0 {
		for _curItem, element := range m.NodesToWrite {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.NodesToWrite), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_WriteRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func WriteRequestParse(ctx context.Context, theBytes []byte, identifier string) (WriteRequest, error) {
	return WriteRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func WriteRequestParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (WriteRequest, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (WriteRequest, error) {
		return WriteRequestParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func WriteRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (WriteRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("WriteRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for WriteRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}

	noOfNodesToWrite, err := ReadSimpleField(ctx, "noOfNodesToWrite", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfNodesToWrite' field"))
	}

	nodesToWrite, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "nodesToWrite", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("670")), readBuffer), uint64(noOfNodesToWrite))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodesToWrite' field"))
	}

	if closeErr := readBuffer.CloseContext("WriteRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for WriteRequest")
	}

	// Create a partially initialized instance
	_child := &_WriteRequest{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		RequestHeader:              requestHeader,
		NoOfNodesToWrite:           noOfNodesToWrite,
		NodesToWrite:               nodesToWrite,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_WriteRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_WriteRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("WriteRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for WriteRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfNodesToWrite", m.GetNoOfNodesToWrite(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfNodesToWrite' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "nodesToWrite", m.GetNodesToWrite(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'nodesToWrite' field")
		}

		if popErr := writeBuffer.PopContext("WriteRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for WriteRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_WriteRequest) isWriteRequest() bool {
	return true
}

func (m *_WriteRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
