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

// MonitoredItemModifyResult is the corresponding interface of MonitoredItemModifyResult
type MonitoredItemModifyResult interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetStatusCode returns StatusCode (property field)
	GetStatusCode() StatusCode
	// GetRevisedSamplingInterval returns RevisedSamplingInterval (property field)
	GetRevisedSamplingInterval() float64
	// GetRevisedQueueSize returns RevisedQueueSize (property field)
	GetRevisedQueueSize() uint32
	// GetFilterResult returns FilterResult (property field)
	GetFilterResult() ExtensionObject
}

// MonitoredItemModifyResultExactly can be used when we want exactly this type and not a type which fulfills MonitoredItemModifyResult.
// This is useful for switch cases.
type MonitoredItemModifyResultExactly interface {
	MonitoredItemModifyResult
	isMonitoredItemModifyResult() bool
}

// _MonitoredItemModifyResult is the data-structure of this message
type _MonitoredItemModifyResult struct {
	ExtensionObjectDefinitionContract
	StatusCode              StatusCode
	RevisedSamplingInterval float64
	RevisedQueueSize        uint32
	FilterResult            ExtensionObject
}

var _ MonitoredItemModifyResult = (*_MonitoredItemModifyResult)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MonitoredItemModifyResult) GetIdentifier() string {
	return "760"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MonitoredItemModifyResult) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MonitoredItemModifyResult) GetStatusCode() StatusCode {
	return m.StatusCode
}

func (m *_MonitoredItemModifyResult) GetRevisedSamplingInterval() float64 {
	return m.RevisedSamplingInterval
}

func (m *_MonitoredItemModifyResult) GetRevisedQueueSize() uint32 {
	return m.RevisedQueueSize
}

func (m *_MonitoredItemModifyResult) GetFilterResult() ExtensionObject {
	return m.FilterResult
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMonitoredItemModifyResult factory function for _MonitoredItemModifyResult
func NewMonitoredItemModifyResult(statusCode StatusCode, revisedSamplingInterval float64, revisedQueueSize uint32, filterResult ExtensionObject) *_MonitoredItemModifyResult {
	_result := &_MonitoredItemModifyResult{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		StatusCode:                        statusCode,
		RevisedSamplingInterval:           revisedSamplingInterval,
		RevisedQueueSize:                  revisedQueueSize,
		FilterResult:                      filterResult,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastMonitoredItemModifyResult(structType any) MonitoredItemModifyResult {
	if casted, ok := structType.(MonitoredItemModifyResult); ok {
		return casted
	}
	if casted, ok := structType.(*MonitoredItemModifyResult); ok {
		return *casted
	}
	return nil
}

func (m *_MonitoredItemModifyResult) GetTypeName() string {
	return "MonitoredItemModifyResult"
}

func (m *_MonitoredItemModifyResult) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (statusCode)
	lengthInBits += m.StatusCode.GetLengthInBits(ctx)

	// Simple field (revisedSamplingInterval)
	lengthInBits += 64

	// Simple field (revisedQueueSize)
	lengthInBits += 32

	// Simple field (filterResult)
	lengthInBits += m.FilterResult.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_MonitoredItemModifyResult) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MonitoredItemModifyResult) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (_monitoredItemModifyResult MonitoredItemModifyResult, err error) {
	m.ExtensionObjectDefinitionContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MonitoredItemModifyResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MonitoredItemModifyResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	statusCode, err := ReadSimpleField[StatusCode](ctx, "statusCode", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusCode' field"))
	}
	m.StatusCode = statusCode

	revisedSamplingInterval, err := ReadSimpleField(ctx, "revisedSamplingInterval", ReadDouble(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'revisedSamplingInterval' field"))
	}
	m.RevisedSamplingInterval = revisedSamplingInterval

	revisedQueueSize, err := ReadSimpleField(ctx, "revisedQueueSize", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'revisedQueueSize' field"))
	}
	m.RevisedQueueSize = revisedQueueSize

	filterResult, err := ReadSimpleField[ExtensionObject](ctx, "filterResult", ReadComplex[ExtensionObject](ExtensionObjectParseWithBufferProducer((bool)(bool(true))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'filterResult' field"))
	}
	m.FilterResult = filterResult

	if closeErr := readBuffer.CloseContext("MonitoredItemModifyResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MonitoredItemModifyResult")
	}

	return m, nil
}

func (m *_MonitoredItemModifyResult) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MonitoredItemModifyResult) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MonitoredItemModifyResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MonitoredItemModifyResult")
		}

		if err := WriteSimpleField[StatusCode](ctx, "statusCode", m.GetStatusCode(), WriteComplex[StatusCode](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'statusCode' field")
		}

		if err := WriteSimpleField[float64](ctx, "revisedSamplingInterval", m.GetRevisedSamplingInterval(), WriteDouble(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'revisedSamplingInterval' field")
		}

		if err := WriteSimpleField[uint32](ctx, "revisedQueueSize", m.GetRevisedQueueSize(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'revisedQueueSize' field")
		}

		if err := WriteSimpleField[ExtensionObject](ctx, "filterResult", m.GetFilterResult(), WriteComplex[ExtensionObject](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'filterResult' field")
		}

		if popErr := writeBuffer.PopContext("MonitoredItemModifyResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MonitoredItemModifyResult")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MonitoredItemModifyResult) isMonitoredItemModifyResult() bool {
	return true
}

func (m *_MonitoredItemModifyResult) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
