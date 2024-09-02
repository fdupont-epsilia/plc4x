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
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// QueryFirstResponse is the corresponding interface of QueryFirstResponse
type QueryFirstResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ExtensionObjectDefinition
	// GetNoOfQueryDataSets returns NoOfQueryDataSets (property field)
	GetNoOfQueryDataSets() int32
	// GetQueryDataSets returns QueryDataSets (property field)
	GetQueryDataSets() []ExtensionObjectDefinition
	// GetContinuationPoint returns ContinuationPoint (property field)
	GetContinuationPoint() PascalByteString
	// GetNoOfParsingResults returns NoOfParsingResults (property field)
	GetNoOfParsingResults() int32
	// GetParsingResults returns ParsingResults (property field)
	GetParsingResults() []ExtensionObjectDefinition
	// GetNoOfDiagnosticInfos returns NoOfDiagnosticInfos (property field)
	GetNoOfDiagnosticInfos() int32
	// GetDiagnosticInfos returns DiagnosticInfos (property field)
	GetDiagnosticInfos() []DiagnosticInfo
	// GetFilterResult returns FilterResult (property field)
	GetFilterResult() ExtensionObjectDefinition
	// IsQueryFirstResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsQueryFirstResponse()
}

// _QueryFirstResponse is the data-structure of this message
type _QueryFirstResponse struct {
	ExtensionObjectDefinitionContract
	ResponseHeader      ExtensionObjectDefinition
	NoOfQueryDataSets   int32
	QueryDataSets       []ExtensionObjectDefinition
	ContinuationPoint   PascalByteString
	NoOfParsingResults  int32
	ParsingResults      []ExtensionObjectDefinition
	NoOfDiagnosticInfos int32
	DiagnosticInfos     []DiagnosticInfo
	FilterResult        ExtensionObjectDefinition
}

var _ QueryFirstResponse = (*_QueryFirstResponse)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_QueryFirstResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_QueryFirstResponse) GetIdentifier() string {
	return "618"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_QueryFirstResponse) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_QueryFirstResponse) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

func (m *_QueryFirstResponse) GetNoOfQueryDataSets() int32 {
	return m.NoOfQueryDataSets
}

func (m *_QueryFirstResponse) GetQueryDataSets() []ExtensionObjectDefinition {
	return m.QueryDataSets
}

func (m *_QueryFirstResponse) GetContinuationPoint() PascalByteString {
	return m.ContinuationPoint
}

func (m *_QueryFirstResponse) GetNoOfParsingResults() int32 {
	return m.NoOfParsingResults
}

func (m *_QueryFirstResponse) GetParsingResults() []ExtensionObjectDefinition {
	return m.ParsingResults
}

func (m *_QueryFirstResponse) GetNoOfDiagnosticInfos() int32 {
	return m.NoOfDiagnosticInfos
}

func (m *_QueryFirstResponse) GetDiagnosticInfos() []DiagnosticInfo {
	return m.DiagnosticInfos
}

func (m *_QueryFirstResponse) GetFilterResult() ExtensionObjectDefinition {
	return m.FilterResult
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewQueryFirstResponse factory function for _QueryFirstResponse
func NewQueryFirstResponse(responseHeader ExtensionObjectDefinition, noOfQueryDataSets int32, queryDataSets []ExtensionObjectDefinition, continuationPoint PascalByteString, noOfParsingResults int32, parsingResults []ExtensionObjectDefinition, noOfDiagnosticInfos int32, diagnosticInfos []DiagnosticInfo, filterResult ExtensionObjectDefinition) *_QueryFirstResponse {
	_result := &_QueryFirstResponse{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ResponseHeader:                    responseHeader,
		NoOfQueryDataSets:                 noOfQueryDataSets,
		QueryDataSets:                     queryDataSets,
		ContinuationPoint:                 continuationPoint,
		NoOfParsingResults:                noOfParsingResults,
		ParsingResults:                    parsingResults,
		NoOfDiagnosticInfos:               noOfDiagnosticInfos,
		DiagnosticInfos:                   diagnosticInfos,
		FilterResult:                      filterResult,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastQueryFirstResponse(structType any) QueryFirstResponse {
	if casted, ok := structType.(QueryFirstResponse); ok {
		return casted
	}
	if casted, ok := structType.(*QueryFirstResponse); ok {
		return *casted
	}
	return nil
}

func (m *_QueryFirstResponse) GetTypeName() string {
	return "QueryFirstResponse"
}

func (m *_QueryFirstResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Simple field (noOfQueryDataSets)
	lengthInBits += 32

	// Array field
	if len(m.QueryDataSets) > 0 {
		for _curItem, element := range m.QueryDataSets {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.QueryDataSets), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (continuationPoint)
	lengthInBits += m.ContinuationPoint.GetLengthInBits(ctx)

	// Simple field (noOfParsingResults)
	lengthInBits += 32

	// Array field
	if len(m.ParsingResults) > 0 {
		for _curItem, element := range m.ParsingResults {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ParsingResults), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfDiagnosticInfos)
	lengthInBits += 32

	// Array field
	if len(m.DiagnosticInfos) > 0 {
		for _curItem, element := range m.DiagnosticInfos {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.DiagnosticInfos), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (filterResult)
	lengthInBits += m.FilterResult.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_QueryFirstResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_QueryFirstResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__queryFirstResponse QueryFirstResponse, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("QueryFirstResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for QueryFirstResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	responseHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("394")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'responseHeader' field"))
	}
	m.ResponseHeader = responseHeader

	noOfQueryDataSets, err := ReadSimpleField(ctx, "noOfQueryDataSets", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfQueryDataSets' field"))
	}
	m.NoOfQueryDataSets = noOfQueryDataSets

	queryDataSets, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "queryDataSets", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("579")), readBuffer), uint64(noOfQueryDataSets))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'queryDataSets' field"))
	}
	m.QueryDataSets = queryDataSets

	continuationPoint, err := ReadSimpleField[PascalByteString](ctx, "continuationPoint", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'continuationPoint' field"))
	}
	m.ContinuationPoint = continuationPoint

	noOfParsingResults, err := ReadSimpleField(ctx, "noOfParsingResults", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfParsingResults' field"))
	}
	m.NoOfParsingResults = noOfParsingResults

	parsingResults, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "parsingResults", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("612")), readBuffer), uint64(noOfParsingResults))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'parsingResults' field"))
	}
	m.ParsingResults = parsingResults

	noOfDiagnosticInfos, err := ReadSimpleField(ctx, "noOfDiagnosticInfos", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfDiagnosticInfos' field"))
	}
	m.NoOfDiagnosticInfos = noOfDiagnosticInfos

	diagnosticInfos, err := ReadCountArrayField[DiagnosticInfo](ctx, "diagnosticInfos", ReadComplex[DiagnosticInfo](DiagnosticInfoParseWithBuffer, readBuffer), uint64(noOfDiagnosticInfos))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'diagnosticInfos' field"))
	}
	m.DiagnosticInfos = diagnosticInfos

	filterResult, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "filterResult", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("609")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'filterResult' field"))
	}
	m.FilterResult = filterResult

	if closeErr := readBuffer.CloseContext("QueryFirstResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for QueryFirstResponse")
	}

	return m, nil
}

func (m *_QueryFirstResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_QueryFirstResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("QueryFirstResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for QueryFirstResponse")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", m.GetResponseHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'responseHeader' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfQueryDataSets", m.GetNoOfQueryDataSets(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfQueryDataSets' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "queryDataSets", m.GetQueryDataSets(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'queryDataSets' field")
		}

		if err := WriteSimpleField[PascalByteString](ctx, "continuationPoint", m.GetContinuationPoint(), WriteComplex[PascalByteString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'continuationPoint' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfParsingResults", m.GetNoOfParsingResults(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfParsingResults' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "parsingResults", m.GetParsingResults(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'parsingResults' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfDiagnosticInfos", m.GetNoOfDiagnosticInfos(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfDiagnosticInfos' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "diagnosticInfos", m.GetDiagnosticInfos(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'diagnosticInfos' field")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "filterResult", m.GetFilterResult(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'filterResult' field")
		}

		if popErr := writeBuffer.PopContext("QueryFirstResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for QueryFirstResponse")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_QueryFirstResponse) IsQueryFirstResponse() {}

func (m *_QueryFirstResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
