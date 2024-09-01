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

// BrowsePathTarget is the corresponding interface of BrowsePathTarget
type BrowsePathTarget interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetTargetId returns TargetId (property field)
	GetTargetId() ExpandedNodeId
	// GetRemainingPathIndex returns RemainingPathIndex (property field)
	GetRemainingPathIndex() uint32
}

// BrowsePathTargetExactly can be used when we want exactly this type and not a type which fulfills BrowsePathTarget.
// This is useful for switch cases.
type BrowsePathTargetExactly interface {
	BrowsePathTarget
	isBrowsePathTarget() bool
}

// _BrowsePathTarget is the data-structure of this message
type _BrowsePathTarget struct {
	*_ExtensionObjectDefinition
	TargetId           ExpandedNodeId
	RemainingPathIndex uint32
}

var _ BrowsePathTarget = (*_BrowsePathTarget)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BrowsePathTarget) GetIdentifier() string {
	return "548"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BrowsePathTarget) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_BrowsePathTarget) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BrowsePathTarget) GetTargetId() ExpandedNodeId {
	return m.TargetId
}

func (m *_BrowsePathTarget) GetRemainingPathIndex() uint32 {
	return m.RemainingPathIndex
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBrowsePathTarget factory function for _BrowsePathTarget
func NewBrowsePathTarget(targetId ExpandedNodeId, remainingPathIndex uint32) *_BrowsePathTarget {
	_result := &_BrowsePathTarget{
		TargetId:                   targetId,
		RemainingPathIndex:         remainingPathIndex,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBrowsePathTarget(structType any) BrowsePathTarget {
	if casted, ok := structType.(BrowsePathTarget); ok {
		return casted
	}
	if casted, ok := structType.(*BrowsePathTarget); ok {
		return *casted
	}
	return nil
}

func (m *_BrowsePathTarget) GetTypeName() string {
	return "BrowsePathTarget"
}

func (m *_BrowsePathTarget) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (targetId)
	lengthInBits += m.TargetId.GetLengthInBits(ctx)

	// Simple field (remainingPathIndex)
	lengthInBits += 32

	return lengthInBits
}

func (m *_BrowsePathTarget) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BrowsePathTargetParse(ctx context.Context, theBytes []byte, identifier string) (BrowsePathTarget, error) {
	return BrowsePathTargetParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func BrowsePathTargetParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (BrowsePathTarget, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BrowsePathTarget, error) {
		return BrowsePathTargetParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func BrowsePathTargetParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (BrowsePathTarget, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BrowsePathTarget"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BrowsePathTarget")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	targetId, err := ReadSimpleField[ExpandedNodeId](ctx, "targetId", ReadComplex[ExpandedNodeId](ExpandedNodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'targetId' field"))
	}

	remainingPathIndex, err := ReadSimpleField(ctx, "remainingPathIndex", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'remainingPathIndex' field"))
	}

	if closeErr := readBuffer.CloseContext("BrowsePathTarget"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BrowsePathTarget")
	}

	// Create a partially initialized instance
	_child := &_BrowsePathTarget{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		TargetId:                   targetId,
		RemainingPathIndex:         remainingPathIndex,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_BrowsePathTarget) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BrowsePathTarget) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BrowsePathTarget"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BrowsePathTarget")
		}

		if err := WriteSimpleField[ExpandedNodeId](ctx, "targetId", m.GetTargetId(), WriteComplex[ExpandedNodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'targetId' field")
		}

		if err := WriteSimpleField[uint32](ctx, "remainingPathIndex", m.GetRemainingPathIndex(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'remainingPathIndex' field")
		}

		if popErr := writeBuffer.PopContext("BrowsePathTarget"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BrowsePathTarget")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BrowsePathTarget) isBrowsePathTarget() bool {
	return true
}

func (m *_BrowsePathTarget) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
