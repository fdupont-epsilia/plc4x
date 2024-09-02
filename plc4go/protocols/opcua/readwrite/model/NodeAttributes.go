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

// NodeAttributes is the corresponding interface of NodeAttributes
type NodeAttributes interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetSpecifiedAttributes returns SpecifiedAttributes (property field)
	GetSpecifiedAttributes() uint32
	// GetDisplayName returns DisplayName (property field)
	GetDisplayName() LocalizedText
	// GetDescription returns Description (property field)
	GetDescription() LocalizedText
	// GetWriteMask returns WriteMask (property field)
	GetWriteMask() uint32
	// GetUserWriteMask returns UserWriteMask (property field)
	GetUserWriteMask() uint32
	// IsNodeAttributes is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNodeAttributes()
}

// _NodeAttributes is the data-structure of this message
type _NodeAttributes struct {
	ExtensionObjectDefinitionContract
	SpecifiedAttributes uint32
	DisplayName         LocalizedText
	Description         LocalizedText
	WriteMask           uint32
	UserWriteMask       uint32
}

var _ NodeAttributes = (*_NodeAttributes)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_NodeAttributes)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NodeAttributes) GetIdentifier() string {
	return "351"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NodeAttributes) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NodeAttributes) GetSpecifiedAttributes() uint32 {
	return m.SpecifiedAttributes
}

func (m *_NodeAttributes) GetDisplayName() LocalizedText {
	return m.DisplayName
}

func (m *_NodeAttributes) GetDescription() LocalizedText {
	return m.Description
}

func (m *_NodeAttributes) GetWriteMask() uint32 {
	return m.WriteMask
}

func (m *_NodeAttributes) GetUserWriteMask() uint32 {
	return m.UserWriteMask
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNodeAttributes factory function for _NodeAttributes
func NewNodeAttributes(specifiedAttributes uint32, displayName LocalizedText, description LocalizedText, writeMask uint32, userWriteMask uint32) *_NodeAttributes {
	_result := &_NodeAttributes{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		SpecifiedAttributes:               specifiedAttributes,
		DisplayName:                       displayName,
		Description:                       description,
		WriteMask:                         writeMask,
		UserWriteMask:                     userWriteMask,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNodeAttributes(structType any) NodeAttributes {
	if casted, ok := structType.(NodeAttributes); ok {
		return casted
	}
	if casted, ok := structType.(*NodeAttributes); ok {
		return *casted
	}
	return nil
}

func (m *_NodeAttributes) GetTypeName() string {
	return "NodeAttributes"
}

func (m *_NodeAttributes) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (specifiedAttributes)
	lengthInBits += 32

	// Simple field (displayName)
	lengthInBits += m.DisplayName.GetLengthInBits(ctx)

	// Simple field (description)
	lengthInBits += m.Description.GetLengthInBits(ctx)

	// Simple field (writeMask)
	lengthInBits += 32

	// Simple field (userWriteMask)
	lengthInBits += 32

	return lengthInBits
}

func (m *_NodeAttributes) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NodeAttributes) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__nodeAttributes NodeAttributes, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NodeAttributes"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NodeAttributes")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	specifiedAttributes, err := ReadSimpleField(ctx, "specifiedAttributes", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'specifiedAttributes' field"))
	}
	m.SpecifiedAttributes = specifiedAttributes

	displayName, err := ReadSimpleField[LocalizedText](ctx, "displayName", ReadComplex[LocalizedText](LocalizedTextParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'displayName' field"))
	}
	m.DisplayName = displayName

	description, err := ReadSimpleField[LocalizedText](ctx, "description", ReadComplex[LocalizedText](LocalizedTextParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'description' field"))
	}
	m.Description = description

	writeMask, err := ReadSimpleField(ctx, "writeMask", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'writeMask' field"))
	}
	m.WriteMask = writeMask

	userWriteMask, err := ReadSimpleField(ctx, "userWriteMask", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'userWriteMask' field"))
	}
	m.UserWriteMask = userWriteMask

	if closeErr := readBuffer.CloseContext("NodeAttributes"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NodeAttributes")
	}

	return m, nil
}

func (m *_NodeAttributes) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NodeAttributes) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NodeAttributes"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NodeAttributes")
		}

		if err := WriteSimpleField[uint32](ctx, "specifiedAttributes", m.GetSpecifiedAttributes(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'specifiedAttributes' field")
		}

		if err := WriteSimpleField[LocalizedText](ctx, "displayName", m.GetDisplayName(), WriteComplex[LocalizedText](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'displayName' field")
		}

		if err := WriteSimpleField[LocalizedText](ctx, "description", m.GetDescription(), WriteComplex[LocalizedText](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'description' field")
		}

		if err := WriteSimpleField[uint32](ctx, "writeMask", m.GetWriteMask(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'writeMask' field")
		}

		if err := WriteSimpleField[uint32](ctx, "userWriteMask", m.GetUserWriteMask(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'userWriteMask' field")
		}

		if popErr := writeBuffer.PopContext("NodeAttributes"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NodeAttributes")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NodeAttributes) IsNodeAttributes() {}

func (m *_NodeAttributes) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
