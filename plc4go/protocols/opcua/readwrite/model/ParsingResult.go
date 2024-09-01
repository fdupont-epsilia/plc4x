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

// ParsingResult is the corresponding interface of ParsingResult
type ParsingResult interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetStatusCode returns StatusCode (property field)
	GetStatusCode() StatusCode
	// GetNoOfDataStatusCodes returns NoOfDataStatusCodes (property field)
	GetNoOfDataStatusCodes() int32
	// GetDataStatusCodes returns DataStatusCodes (property field)
	GetDataStatusCodes() []StatusCode
	// GetNoOfDataDiagnosticInfos returns NoOfDataDiagnosticInfos (property field)
	GetNoOfDataDiagnosticInfos() int32
	// GetDataDiagnosticInfos returns DataDiagnosticInfos (property field)
	GetDataDiagnosticInfos() []DiagnosticInfo
}

// ParsingResultExactly can be used when we want exactly this type and not a type which fulfills ParsingResult.
// This is useful for switch cases.
type ParsingResultExactly interface {
	ParsingResult
	isParsingResult() bool
}

// _ParsingResult is the data-structure of this message
type _ParsingResult struct {
	*_ExtensionObjectDefinition
	StatusCode              StatusCode
	NoOfDataStatusCodes     int32
	DataStatusCodes         []StatusCode
	NoOfDataDiagnosticInfos int32
	DataDiagnosticInfos     []DiagnosticInfo
}

var _ ParsingResult = (*_ParsingResult)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ParsingResult) GetIdentifier() string {
	return "612"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ParsingResult) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_ParsingResult) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ParsingResult) GetStatusCode() StatusCode {
	return m.StatusCode
}

func (m *_ParsingResult) GetNoOfDataStatusCodes() int32 {
	return m.NoOfDataStatusCodes
}

func (m *_ParsingResult) GetDataStatusCodes() []StatusCode {
	return m.DataStatusCodes
}

func (m *_ParsingResult) GetNoOfDataDiagnosticInfos() int32 {
	return m.NoOfDataDiagnosticInfos
}

func (m *_ParsingResult) GetDataDiagnosticInfos() []DiagnosticInfo {
	return m.DataDiagnosticInfos
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewParsingResult factory function for _ParsingResult
func NewParsingResult(statusCode StatusCode, noOfDataStatusCodes int32, dataStatusCodes []StatusCode, noOfDataDiagnosticInfos int32, dataDiagnosticInfos []DiagnosticInfo) *_ParsingResult {
	_result := &_ParsingResult{
		StatusCode:                 statusCode,
		NoOfDataStatusCodes:        noOfDataStatusCodes,
		DataStatusCodes:            dataStatusCodes,
		NoOfDataDiagnosticInfos:    noOfDataDiagnosticInfos,
		DataDiagnosticInfos:        dataDiagnosticInfos,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastParsingResult(structType any) ParsingResult {
	if casted, ok := structType.(ParsingResult); ok {
		return casted
	}
	if casted, ok := structType.(*ParsingResult); ok {
		return *casted
	}
	return nil
}

func (m *_ParsingResult) GetTypeName() string {
	return "ParsingResult"
}

func (m *_ParsingResult) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (statusCode)
	lengthInBits += m.StatusCode.GetLengthInBits(ctx)

	// Simple field (noOfDataStatusCodes)
	lengthInBits += 32

	// Array field
	if len(m.DataStatusCodes) > 0 {
		for _curItem, element := range m.DataStatusCodes {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.DataStatusCodes), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfDataDiagnosticInfos)
	lengthInBits += 32

	// Array field
	if len(m.DataDiagnosticInfos) > 0 {
		for _curItem, element := range m.DataDiagnosticInfos {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.DataDiagnosticInfos), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_ParsingResult) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ParsingResultParse(ctx context.Context, theBytes []byte, identifier string) (ParsingResult, error) {
	return ParsingResultParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func ParsingResultParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (ParsingResult, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ParsingResult, error) {
		return ParsingResultParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func ParsingResultParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (ParsingResult, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ParsingResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ParsingResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	statusCode, err := ReadSimpleField[StatusCode](ctx, "statusCode", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusCode' field"))
	}

	noOfDataStatusCodes, err := ReadSimpleField(ctx, "noOfDataStatusCodes", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfDataStatusCodes' field"))
	}

	dataStatusCodes, err := ReadCountArrayField[StatusCode](ctx, "dataStatusCodes", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer), uint64(noOfDataStatusCodes))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataStatusCodes' field"))
	}

	noOfDataDiagnosticInfos, err := ReadSimpleField(ctx, "noOfDataDiagnosticInfos", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfDataDiagnosticInfos' field"))
	}

	dataDiagnosticInfos, err := ReadCountArrayField[DiagnosticInfo](ctx, "dataDiagnosticInfos", ReadComplex[DiagnosticInfo](DiagnosticInfoParseWithBuffer, readBuffer), uint64(noOfDataDiagnosticInfos))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataDiagnosticInfos' field"))
	}

	if closeErr := readBuffer.CloseContext("ParsingResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ParsingResult")
	}

	// Create a partially initialized instance
	_child := &_ParsingResult{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		StatusCode:                 statusCode,
		NoOfDataStatusCodes:        noOfDataStatusCodes,
		DataStatusCodes:            dataStatusCodes,
		NoOfDataDiagnosticInfos:    noOfDataDiagnosticInfos,
		DataDiagnosticInfos:        dataDiagnosticInfos,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_ParsingResult) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ParsingResult) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ParsingResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ParsingResult")
		}

		if err := WriteSimpleField[StatusCode](ctx, "statusCode", m.GetStatusCode(), WriteComplex[StatusCode](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'statusCode' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfDataStatusCodes", m.GetNoOfDataStatusCodes(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfDataStatusCodes' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "dataStatusCodes", m.GetDataStatusCodes(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'dataStatusCodes' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfDataDiagnosticInfos", m.GetNoOfDataDiagnosticInfos(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfDataDiagnosticInfos' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "dataDiagnosticInfos", m.GetDataDiagnosticInfos(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'dataDiagnosticInfos' field")
		}

		if popErr := writeBuffer.PopContext("ParsingResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ParsingResult")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ParsingResult) isParsingResult() bool {
	return true
}

func (m *_ParsingResult) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
