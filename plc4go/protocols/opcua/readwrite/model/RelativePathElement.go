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

// RelativePathElement is the corresponding interface of RelativePathElement
type RelativePathElement interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetReferenceTypeId returns ReferenceTypeId (property field)
	GetReferenceTypeId() NodeId
	// GetIncludeSubtypes returns IncludeSubtypes (property field)
	GetIncludeSubtypes() bool
	// GetIsInverse returns IsInverse (property field)
	GetIsInverse() bool
	// GetTargetName returns TargetName (property field)
	GetTargetName() QualifiedName
}

// RelativePathElementExactly can be used when we want exactly this type and not a type which fulfills RelativePathElement.
// This is useful for switch cases.
type RelativePathElementExactly interface {
	RelativePathElement
	isRelativePathElement() bool
}

// _RelativePathElement is the data-structure of this message
type _RelativePathElement struct {
	*_ExtensionObjectDefinition
	ReferenceTypeId NodeId
	IncludeSubtypes bool
	IsInverse       bool
	TargetName      QualifiedName
	// Reserved Fields
	reservedField0 *uint8
}

var _ RelativePathElement = (*_RelativePathElement)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_RelativePathElement) GetIdentifier() string {
	return "539"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RelativePathElement) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_RelativePathElement) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RelativePathElement) GetReferenceTypeId() NodeId {
	return m.ReferenceTypeId
}

func (m *_RelativePathElement) GetIncludeSubtypes() bool {
	return m.IncludeSubtypes
}

func (m *_RelativePathElement) GetIsInverse() bool {
	return m.IsInverse
}

func (m *_RelativePathElement) GetTargetName() QualifiedName {
	return m.TargetName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewRelativePathElement factory function for _RelativePathElement
func NewRelativePathElement(referenceTypeId NodeId, includeSubtypes bool, isInverse bool, targetName QualifiedName) *_RelativePathElement {
	_result := &_RelativePathElement{
		ReferenceTypeId:            referenceTypeId,
		IncludeSubtypes:            includeSubtypes,
		IsInverse:                  isInverse,
		TargetName:                 targetName,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastRelativePathElement(structType any) RelativePathElement {
	if casted, ok := structType.(RelativePathElement); ok {
		return casted
	}
	if casted, ok := structType.(*RelativePathElement); ok {
		return *casted
	}
	return nil
}

func (m *_RelativePathElement) GetTypeName() string {
	return "RelativePathElement"
}

func (m *_RelativePathElement) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (referenceTypeId)
	lengthInBits += m.ReferenceTypeId.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 6

	// Simple field (includeSubtypes)
	lengthInBits += 1

	// Simple field (isInverse)
	lengthInBits += 1

	// Simple field (targetName)
	lengthInBits += m.TargetName.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_RelativePathElement) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func RelativePathElementParse(ctx context.Context, theBytes []byte, identifier string) (RelativePathElement, error) {
	return RelativePathElementParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func RelativePathElementParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (RelativePathElement, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (RelativePathElement, error) {
		return RelativePathElementParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func RelativePathElementParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (RelativePathElement, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RelativePathElement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RelativePathElement")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	referenceTypeId, err := ReadSimpleField[NodeId](ctx, "referenceTypeId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'referenceTypeId' field"))
	}

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(6)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	includeSubtypes, err := ReadSimpleField(ctx, "includeSubtypes", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'includeSubtypes' field"))
	}

	isInverse, err := ReadSimpleField(ctx, "isInverse", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isInverse' field"))
	}

	targetName, err := ReadSimpleField[QualifiedName](ctx, "targetName", ReadComplex[QualifiedName](QualifiedNameParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'targetName' field"))
	}

	if closeErr := readBuffer.CloseContext("RelativePathElement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RelativePathElement")
	}

	// Create a partially initialized instance
	_child := &_RelativePathElement{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		ReferenceTypeId:            referenceTypeId,
		IncludeSubtypes:            includeSubtypes,
		IsInverse:                  isInverse,
		TargetName:                 targetName,
		reservedField0:             reservedField0,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_RelativePathElement) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RelativePathElement) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RelativePathElement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RelativePathElement")
		}

		if err := WriteSimpleField[NodeId](ctx, "referenceTypeId", m.GetReferenceTypeId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'referenceTypeId' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 6)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "includeSubtypes", m.GetIncludeSubtypes(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'includeSubtypes' field")
		}

		if err := WriteSimpleField[bool](ctx, "isInverse", m.GetIsInverse(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'isInverse' field")
		}

		if err := WriteSimpleField[QualifiedName](ctx, "targetName", m.GetTargetName(), WriteComplex[QualifiedName](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'targetName' field")
		}

		if popErr := writeBuffer.PopContext("RelativePathElement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RelativePathElement")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RelativePathElement) isRelativePathElement() bool {
	return true
}

func (m *_RelativePathElement) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
