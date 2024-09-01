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

// BrowsePathResult is the corresponding interface of BrowsePathResult
type BrowsePathResult interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetStatusCode returns StatusCode (property field)
	GetStatusCode() StatusCode
	// GetNoOfTargets returns NoOfTargets (property field)
	GetNoOfTargets() int32
	// GetTargets returns Targets (property field)
	GetTargets() []ExtensionObjectDefinition
}

// BrowsePathResultExactly can be used when we want exactly this type and not a type which fulfills BrowsePathResult.
// This is useful for switch cases.
type BrowsePathResultExactly interface {
	BrowsePathResult
	isBrowsePathResult() bool
}

// _BrowsePathResult is the data-structure of this message
type _BrowsePathResult struct {
	*_ExtensionObjectDefinition
	StatusCode  StatusCode
	NoOfTargets int32
	Targets     []ExtensionObjectDefinition
}

var _ BrowsePathResult = (*_BrowsePathResult)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BrowsePathResult) GetIdentifier() string {
	return "551"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BrowsePathResult) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_BrowsePathResult) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BrowsePathResult) GetStatusCode() StatusCode {
	return m.StatusCode
}

func (m *_BrowsePathResult) GetNoOfTargets() int32 {
	return m.NoOfTargets
}

func (m *_BrowsePathResult) GetTargets() []ExtensionObjectDefinition {
	return m.Targets
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBrowsePathResult factory function for _BrowsePathResult
func NewBrowsePathResult(statusCode StatusCode, noOfTargets int32, targets []ExtensionObjectDefinition) *_BrowsePathResult {
	_result := &_BrowsePathResult{
		StatusCode:                 statusCode,
		NoOfTargets:                noOfTargets,
		Targets:                    targets,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBrowsePathResult(structType any) BrowsePathResult {
	if casted, ok := structType.(BrowsePathResult); ok {
		return casted
	}
	if casted, ok := structType.(*BrowsePathResult); ok {
		return *casted
	}
	return nil
}

func (m *_BrowsePathResult) GetTypeName() string {
	return "BrowsePathResult"
}

func (m *_BrowsePathResult) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (statusCode)
	lengthInBits += m.StatusCode.GetLengthInBits(ctx)

	// Simple field (noOfTargets)
	lengthInBits += 32

	// Array field
	if len(m.Targets) > 0 {
		for _curItem, element := range m.Targets {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Targets), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_BrowsePathResult) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BrowsePathResultParse(ctx context.Context, theBytes []byte, identifier string) (BrowsePathResult, error) {
	return BrowsePathResultParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func BrowsePathResultParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (BrowsePathResult, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BrowsePathResult, error) {
		return BrowsePathResultParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func BrowsePathResultParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (BrowsePathResult, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BrowsePathResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BrowsePathResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	statusCode, err := ReadSimpleField[StatusCode](ctx, "statusCode", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusCode' field"))
	}

	noOfTargets, err := ReadSimpleField(ctx, "noOfTargets", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfTargets' field"))
	}

	targets, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "targets", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("548")), readBuffer), uint64(noOfTargets))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'targets' field"))
	}

	if closeErr := readBuffer.CloseContext("BrowsePathResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BrowsePathResult")
	}

	// Create a partially initialized instance
	_child := &_BrowsePathResult{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		StatusCode:                 statusCode,
		NoOfTargets:                noOfTargets,
		Targets:                    targets,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_BrowsePathResult) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BrowsePathResult) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BrowsePathResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BrowsePathResult")
		}

		if err := WriteSimpleField[StatusCode](ctx, "statusCode", m.GetStatusCode(), WriteComplex[StatusCode](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'statusCode' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfTargets", m.GetNoOfTargets(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfTargets' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "targets", m.GetTargets(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'targets' field")
		}

		if popErr := writeBuffer.PopContext("BrowsePathResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BrowsePathResult")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BrowsePathResult) isBrowsePathResult() bool {
	return true
}

func (m *_BrowsePathResult) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
