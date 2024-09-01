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

// TransferResult is the corresponding interface of TransferResult
type TransferResult interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetStatusCode returns StatusCode (property field)
	GetStatusCode() StatusCode
	// GetNoOfAvailableSequenceNumbers returns NoOfAvailableSequenceNumbers (property field)
	GetNoOfAvailableSequenceNumbers() int32
	// GetAvailableSequenceNumbers returns AvailableSequenceNumbers (property field)
	GetAvailableSequenceNumbers() []uint32
}

// TransferResultExactly can be used when we want exactly this type and not a type which fulfills TransferResult.
// This is useful for switch cases.
type TransferResultExactly interface {
	TransferResult
	isTransferResult() bool
}

// _TransferResult is the data-structure of this message
type _TransferResult struct {
	*_ExtensionObjectDefinition
	StatusCode                   StatusCode
	NoOfAvailableSequenceNumbers int32
	AvailableSequenceNumbers     []uint32
}

var _ TransferResult = (*_TransferResult)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_TransferResult) GetIdentifier() string {
	return "838"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TransferResult) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_TransferResult) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TransferResult) GetStatusCode() StatusCode {
	return m.StatusCode
}

func (m *_TransferResult) GetNoOfAvailableSequenceNumbers() int32 {
	return m.NoOfAvailableSequenceNumbers
}

func (m *_TransferResult) GetAvailableSequenceNumbers() []uint32 {
	return m.AvailableSequenceNumbers
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTransferResult factory function for _TransferResult
func NewTransferResult(statusCode StatusCode, noOfAvailableSequenceNumbers int32, availableSequenceNumbers []uint32) *_TransferResult {
	_result := &_TransferResult{
		StatusCode:                   statusCode,
		NoOfAvailableSequenceNumbers: noOfAvailableSequenceNumbers,
		AvailableSequenceNumbers:     availableSequenceNumbers,
		_ExtensionObjectDefinition:   NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastTransferResult(structType any) TransferResult {
	if casted, ok := structType.(TransferResult); ok {
		return casted
	}
	if casted, ok := structType.(*TransferResult); ok {
		return *casted
	}
	return nil
}

func (m *_TransferResult) GetTypeName() string {
	return "TransferResult"
}

func (m *_TransferResult) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (statusCode)
	lengthInBits += m.StatusCode.GetLengthInBits(ctx)

	// Simple field (noOfAvailableSequenceNumbers)
	lengthInBits += 32

	// Array field
	if len(m.AvailableSequenceNumbers) > 0 {
		lengthInBits += 32 * uint16(len(m.AvailableSequenceNumbers))
	}

	return lengthInBits
}

func (m *_TransferResult) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TransferResultParse(ctx context.Context, theBytes []byte, identifier string) (TransferResult, error) {
	return TransferResultParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func TransferResultParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (TransferResult, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (TransferResult, error) {
		return TransferResultParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func TransferResultParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (TransferResult, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TransferResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TransferResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	statusCode, err := ReadSimpleField[StatusCode](ctx, "statusCode", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusCode' field"))
	}

	noOfAvailableSequenceNumbers, err := ReadSimpleField(ctx, "noOfAvailableSequenceNumbers", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfAvailableSequenceNumbers' field"))
	}

	availableSequenceNumbers, err := ReadCountArrayField[uint32](ctx, "availableSequenceNumbers", ReadUnsignedInt(readBuffer, uint8(32)), uint64(noOfAvailableSequenceNumbers))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'availableSequenceNumbers' field"))
	}

	if closeErr := readBuffer.CloseContext("TransferResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TransferResult")
	}

	// Create a partially initialized instance
	_child := &_TransferResult{
		_ExtensionObjectDefinition:   &_ExtensionObjectDefinition{},
		StatusCode:                   statusCode,
		NoOfAvailableSequenceNumbers: noOfAvailableSequenceNumbers,
		AvailableSequenceNumbers:     availableSequenceNumbers,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_TransferResult) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TransferResult) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TransferResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TransferResult")
		}

		if err := WriteSimpleField[StatusCode](ctx, "statusCode", m.GetStatusCode(), WriteComplex[StatusCode](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'statusCode' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfAvailableSequenceNumbers", m.GetNoOfAvailableSequenceNumbers(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfAvailableSequenceNumbers' field")
		}

		if err := WriteSimpleTypeArrayField(ctx, "availableSequenceNumbers", m.GetAvailableSequenceNumbers(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'availableSequenceNumbers' field")
		}

		if popErr := writeBuffer.PopContext("TransferResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TransferResult")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TransferResult) isTransferResult() bool {
	return true
}

func (m *_TransferResult) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
