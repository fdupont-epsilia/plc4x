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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// AddNodesResult is the corresponding interface of AddNodesResult
type AddNodesResult interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetStatusCode returns StatusCode (property field)
	GetStatusCode() StatusCode
	// GetAddedNodeId returns AddedNodeId (property field)
	GetAddedNodeId() NodeId
}

// AddNodesResultExactly can be used when we want exactly this type and not a type which fulfills AddNodesResult.
// This is useful for switch cases.
type AddNodesResultExactly interface {
	AddNodesResult
	isAddNodesResult() bool
}

// _AddNodesResult is the data-structure of this message
type _AddNodesResult struct {
	*_ExtensionObjectDefinition
	StatusCode  StatusCode
	AddedNodeId NodeId
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AddNodesResult) GetIdentifier() string {
	return "485"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AddNodesResult) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_AddNodesResult) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AddNodesResult) GetStatusCode() StatusCode {
	return m.StatusCode
}

func (m *_AddNodesResult) GetAddedNodeId() NodeId {
	return m.AddedNodeId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAddNodesResult factory function for _AddNodesResult
func NewAddNodesResult(statusCode StatusCode, addedNodeId NodeId) *_AddNodesResult {
	_result := &_AddNodesResult{
		StatusCode:                 statusCode,
		AddedNodeId:                addedNodeId,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAddNodesResult(structType any) AddNodesResult {
	if casted, ok := structType.(AddNodesResult); ok {
		return casted
	}
	if casted, ok := structType.(*AddNodesResult); ok {
		return *casted
	}
	return nil
}

func (m *_AddNodesResult) GetTypeName() string {
	return "AddNodesResult"
}

func (m *_AddNodesResult) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (statusCode)
	lengthInBits += m.StatusCode.GetLengthInBits(ctx)

	// Simple field (addedNodeId)
	lengthInBits += m.AddedNodeId.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_AddNodesResult) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AddNodesResultParse(ctx context.Context, theBytes []byte, identifier string) (AddNodesResult, error) {
	return AddNodesResultParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func AddNodesResultParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (AddNodesResult, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("AddNodesResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AddNodesResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (statusCode)
	if pullErr := readBuffer.PullContext("statusCode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for statusCode")
	}
	_statusCode, _statusCodeErr := StatusCodeParseWithBuffer(ctx, readBuffer)
	if _statusCodeErr != nil {
		return nil, errors.Wrap(_statusCodeErr, "Error parsing 'statusCode' field of AddNodesResult")
	}
	statusCode := _statusCode.(StatusCode)
	if closeErr := readBuffer.CloseContext("statusCode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for statusCode")
	}

	// Simple Field (addedNodeId)
	if pullErr := readBuffer.PullContext("addedNodeId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for addedNodeId")
	}
	_addedNodeId, _addedNodeIdErr := NodeIdParseWithBuffer(ctx, readBuffer)
	if _addedNodeIdErr != nil {
		return nil, errors.Wrap(_addedNodeIdErr, "Error parsing 'addedNodeId' field of AddNodesResult")
	}
	addedNodeId := _addedNodeId.(NodeId)
	if closeErr := readBuffer.CloseContext("addedNodeId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for addedNodeId")
	}

	if closeErr := readBuffer.CloseContext("AddNodesResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AddNodesResult")
	}

	// Create a partially initialized instance
	_child := &_AddNodesResult{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		StatusCode:                 statusCode,
		AddedNodeId:                addedNodeId,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_AddNodesResult) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AddNodesResult) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AddNodesResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AddNodesResult")
		}

		// Simple Field (statusCode)
		if pushErr := writeBuffer.PushContext("statusCode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for statusCode")
		}
		_statusCodeErr := writeBuffer.WriteSerializable(ctx, m.GetStatusCode())
		if popErr := writeBuffer.PopContext("statusCode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for statusCode")
		}
		if _statusCodeErr != nil {
			return errors.Wrap(_statusCodeErr, "Error serializing 'statusCode' field")
		}

		// Simple Field (addedNodeId)
		if pushErr := writeBuffer.PushContext("addedNodeId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for addedNodeId")
		}
		_addedNodeIdErr := writeBuffer.WriteSerializable(ctx, m.GetAddedNodeId())
		if popErr := writeBuffer.PopContext("addedNodeId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for addedNodeId")
		}
		if _addedNodeIdErr != nil {
			return errors.Wrap(_addedNodeIdErr, "Error serializing 'addedNodeId' field")
		}

		if popErr := writeBuffer.PopContext("AddNodesResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AddNodesResult")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AddNodesResult) isAddNodesResult() bool {
	return true
}

func (m *_AddNodesResult) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}