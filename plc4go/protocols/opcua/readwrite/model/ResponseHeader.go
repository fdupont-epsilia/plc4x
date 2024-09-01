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

// ResponseHeader is the corresponding interface of ResponseHeader
type ResponseHeader interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() int64
	// GetRequestHandle returns RequestHandle (property field)
	GetRequestHandle() uint32
	// GetServiceResult returns ServiceResult (property field)
	GetServiceResult() StatusCode
	// GetServiceDiagnostics returns ServiceDiagnostics (property field)
	GetServiceDiagnostics() DiagnosticInfo
	// GetNoOfStringTable returns NoOfStringTable (property field)
	GetNoOfStringTable() int32
	// GetStringTable returns StringTable (property field)
	GetStringTable() []PascalString
	// GetAdditionalHeader returns AdditionalHeader (property field)
	GetAdditionalHeader() ExtensionObject
}

// ResponseHeaderExactly can be used when we want exactly this type and not a type which fulfills ResponseHeader.
// This is useful for switch cases.
type ResponseHeaderExactly interface {
	ResponseHeader
	isResponseHeader() bool
}

// _ResponseHeader is the data-structure of this message
type _ResponseHeader struct {
	*_ExtensionObjectDefinition
	Timestamp          int64
	RequestHandle      uint32
	ServiceResult      StatusCode
	ServiceDiagnostics DiagnosticInfo
	NoOfStringTable    int32
	StringTable        []PascalString
	AdditionalHeader   ExtensionObject
}

var _ ResponseHeader = (*_ResponseHeader)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ResponseHeader) GetIdentifier() string {
	return "394"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ResponseHeader) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_ResponseHeader) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ResponseHeader) GetTimestamp() int64 {
	return m.Timestamp
}

func (m *_ResponseHeader) GetRequestHandle() uint32 {
	return m.RequestHandle
}

func (m *_ResponseHeader) GetServiceResult() StatusCode {
	return m.ServiceResult
}

func (m *_ResponseHeader) GetServiceDiagnostics() DiagnosticInfo {
	return m.ServiceDiagnostics
}

func (m *_ResponseHeader) GetNoOfStringTable() int32 {
	return m.NoOfStringTable
}

func (m *_ResponseHeader) GetStringTable() []PascalString {
	return m.StringTable
}

func (m *_ResponseHeader) GetAdditionalHeader() ExtensionObject {
	return m.AdditionalHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewResponseHeader factory function for _ResponseHeader
func NewResponseHeader(timestamp int64, requestHandle uint32, serviceResult StatusCode, serviceDiagnostics DiagnosticInfo, noOfStringTable int32, stringTable []PascalString, additionalHeader ExtensionObject) *_ResponseHeader {
	_result := &_ResponseHeader{
		Timestamp:                  timestamp,
		RequestHandle:              requestHandle,
		ServiceResult:              serviceResult,
		ServiceDiagnostics:         serviceDiagnostics,
		NoOfStringTable:            noOfStringTable,
		StringTable:                stringTable,
		AdditionalHeader:           additionalHeader,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastResponseHeader(structType any) ResponseHeader {
	if casted, ok := structType.(ResponseHeader); ok {
		return casted
	}
	if casted, ok := structType.(*ResponseHeader); ok {
		return *casted
	}
	return nil
}

func (m *_ResponseHeader) GetTypeName() string {
	return "ResponseHeader"
}

func (m *_ResponseHeader) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (timestamp)
	lengthInBits += 64

	// Simple field (requestHandle)
	lengthInBits += 32

	// Simple field (serviceResult)
	lengthInBits += m.ServiceResult.GetLengthInBits(ctx)

	// Simple field (serviceDiagnostics)
	lengthInBits += m.ServiceDiagnostics.GetLengthInBits(ctx)

	// Simple field (noOfStringTable)
	lengthInBits += 32

	// Array field
	if len(m.StringTable) > 0 {
		for _curItem, element := range m.StringTable {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.StringTable), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (additionalHeader)
	lengthInBits += m.AdditionalHeader.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ResponseHeader) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ResponseHeaderParse(ctx context.Context, theBytes []byte, identifier string) (ResponseHeader, error) {
	return ResponseHeaderParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func ResponseHeaderParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (ResponseHeader, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ResponseHeader, error) {
		return ResponseHeaderParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func ResponseHeaderParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (ResponseHeader, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ResponseHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ResponseHeader")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timestamp, err := ReadSimpleField(ctx, "timestamp", ReadSignedLong(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timestamp' field"))
	}

	requestHandle, err := ReadSimpleField(ctx, "requestHandle", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHandle' field"))
	}

	serviceResult, err := ReadSimpleField[StatusCode](ctx, "serviceResult", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serviceResult' field"))
	}

	serviceDiagnostics, err := ReadSimpleField[DiagnosticInfo](ctx, "serviceDiagnostics", ReadComplex[DiagnosticInfo](DiagnosticInfoParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serviceDiagnostics' field"))
	}

	noOfStringTable, err := ReadSimpleField(ctx, "noOfStringTable", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfStringTable' field"))
	}

	stringTable, err := ReadCountArrayField[PascalString](ctx, "stringTable", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer), uint64(noOfStringTable))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'stringTable' field"))
	}

	additionalHeader, err := ReadSimpleField[ExtensionObject](ctx, "additionalHeader", ReadComplex[ExtensionObject](ExtensionObjectParseWithBufferProducer((bool)(bool(true))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'additionalHeader' field"))
	}

	if closeErr := readBuffer.CloseContext("ResponseHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ResponseHeader")
	}

	// Create a partially initialized instance
	_child := &_ResponseHeader{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		Timestamp:                  timestamp,
		RequestHandle:              requestHandle,
		ServiceResult:              serviceResult,
		ServiceDiagnostics:         serviceDiagnostics,
		NoOfStringTable:            noOfStringTable,
		StringTable:                stringTable,
		AdditionalHeader:           additionalHeader,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_ResponseHeader) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ResponseHeader) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ResponseHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ResponseHeader")
		}

		if err := WriteSimpleField[int64](ctx, "timestamp", m.GetTimestamp(), WriteSignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'timestamp' field")
		}

		if err := WriteSimpleField[uint32](ctx, "requestHandle", m.GetRequestHandle(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHandle' field")
		}

		if err := WriteSimpleField[StatusCode](ctx, "serviceResult", m.GetServiceResult(), WriteComplex[StatusCode](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'serviceResult' field")
		}

		if err := WriteSimpleField[DiagnosticInfo](ctx, "serviceDiagnostics", m.GetServiceDiagnostics(), WriteComplex[DiagnosticInfo](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'serviceDiagnostics' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfStringTable", m.GetNoOfStringTable(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfStringTable' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "stringTable", m.GetStringTable(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'stringTable' field")
		}

		if err := WriteSimpleField[ExtensionObject](ctx, "additionalHeader", m.GetAdditionalHeader(), WriteComplex[ExtensionObject](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'additionalHeader' field")
		}

		if popErr := writeBuffer.PopContext("ResponseHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ResponseHeader")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ResponseHeader) isResponseHeader() bool {
	return true
}

func (m *_ResponseHeader) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
