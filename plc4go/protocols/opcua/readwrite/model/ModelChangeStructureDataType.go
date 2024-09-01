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

// ModelChangeStructureDataType is the corresponding interface of ModelChangeStructureDataType
type ModelChangeStructureDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetAffected returns Affected (property field)
	GetAffected() NodeId
	// GetAffectedType returns AffectedType (property field)
	GetAffectedType() NodeId
	// GetVerb returns Verb (property field)
	GetVerb() uint8
}

// ModelChangeStructureDataTypeExactly can be used when we want exactly this type and not a type which fulfills ModelChangeStructureDataType.
// This is useful for switch cases.
type ModelChangeStructureDataTypeExactly interface {
	ModelChangeStructureDataType
	isModelChangeStructureDataType() bool
}

// _ModelChangeStructureDataType is the data-structure of this message
type _ModelChangeStructureDataType struct {
	*_ExtensionObjectDefinition
	Affected     NodeId
	AffectedType NodeId
	Verb         uint8
}

var _ ModelChangeStructureDataType = (*_ModelChangeStructureDataType)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModelChangeStructureDataType) GetIdentifier() string {
	return "879"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModelChangeStructureDataType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_ModelChangeStructureDataType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModelChangeStructureDataType) GetAffected() NodeId {
	return m.Affected
}

func (m *_ModelChangeStructureDataType) GetAffectedType() NodeId {
	return m.AffectedType
}

func (m *_ModelChangeStructureDataType) GetVerb() uint8 {
	return m.Verb
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModelChangeStructureDataType factory function for _ModelChangeStructureDataType
func NewModelChangeStructureDataType(affected NodeId, affectedType NodeId, verb uint8) *_ModelChangeStructureDataType {
	_result := &_ModelChangeStructureDataType{
		Affected:                   affected,
		AffectedType:               affectedType,
		Verb:                       verb,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModelChangeStructureDataType(structType any) ModelChangeStructureDataType {
	if casted, ok := structType.(ModelChangeStructureDataType); ok {
		return casted
	}
	if casted, ok := structType.(*ModelChangeStructureDataType); ok {
		return *casted
	}
	return nil
}

func (m *_ModelChangeStructureDataType) GetTypeName() string {
	return "ModelChangeStructureDataType"
}

func (m *_ModelChangeStructureDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (affected)
	lengthInBits += m.Affected.GetLengthInBits(ctx)

	// Simple field (affectedType)
	lengthInBits += m.AffectedType.GetLengthInBits(ctx)

	// Simple field (verb)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ModelChangeStructureDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModelChangeStructureDataTypeParse(ctx context.Context, theBytes []byte, identifier string) (ModelChangeStructureDataType, error) {
	return ModelChangeStructureDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func ModelChangeStructureDataTypeParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (ModelChangeStructureDataType, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ModelChangeStructureDataType, error) {
		return ModelChangeStructureDataTypeParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func ModelChangeStructureDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (ModelChangeStructureDataType, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModelChangeStructureDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModelChangeStructureDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	affected, err := ReadSimpleField[NodeId](ctx, "affected", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'affected' field"))
	}

	affectedType, err := ReadSimpleField[NodeId](ctx, "affectedType", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'affectedType' field"))
	}

	verb, err := ReadSimpleField(ctx, "verb", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'verb' field"))
	}

	if closeErr := readBuffer.CloseContext("ModelChangeStructureDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModelChangeStructureDataType")
	}

	// Create a partially initialized instance
	_child := &_ModelChangeStructureDataType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		Affected:                   affected,
		AffectedType:               affectedType,
		Verb:                       verb,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_ModelChangeStructureDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModelChangeStructureDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModelChangeStructureDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModelChangeStructureDataType")
		}

		if err := WriteSimpleField[NodeId](ctx, "affected", m.GetAffected(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'affected' field")
		}

		if err := WriteSimpleField[NodeId](ctx, "affectedType", m.GetAffectedType(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'affectedType' field")
		}

		if err := WriteSimpleField[uint8](ctx, "verb", m.GetVerb(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'verb' field")
		}

		if popErr := writeBuffer.PopContext("ModelChangeStructureDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModelChangeStructureDataType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModelChangeStructureDataType) isModelChangeStructureDataType() bool {
	return true
}

func (m *_ModelChangeStructureDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
