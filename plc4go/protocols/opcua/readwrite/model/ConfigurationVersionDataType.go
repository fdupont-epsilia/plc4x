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

// ConfigurationVersionDataType is the corresponding interface of ConfigurationVersionDataType
type ConfigurationVersionDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetMajorVersion returns MajorVersion (property field)
	GetMajorVersion() uint32
	// GetMinorVersion returns MinorVersion (property field)
	GetMinorVersion() uint32
}

// ConfigurationVersionDataTypeExactly can be used when we want exactly this type and not a type which fulfills ConfigurationVersionDataType.
// This is useful for switch cases.
type ConfigurationVersionDataTypeExactly interface {
	ConfigurationVersionDataType
	isConfigurationVersionDataType() bool
}

// _ConfigurationVersionDataType is the data-structure of this message
type _ConfigurationVersionDataType struct {
	*_ExtensionObjectDefinition
	MajorVersion uint32
	MinorVersion uint32
}

var _ ConfigurationVersionDataType = (*_ConfigurationVersionDataType)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ConfigurationVersionDataType) GetIdentifier() string {
	return "14595"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ConfigurationVersionDataType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_ConfigurationVersionDataType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ConfigurationVersionDataType) GetMajorVersion() uint32 {
	return m.MajorVersion
}

func (m *_ConfigurationVersionDataType) GetMinorVersion() uint32 {
	return m.MinorVersion
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewConfigurationVersionDataType factory function for _ConfigurationVersionDataType
func NewConfigurationVersionDataType(majorVersion uint32, minorVersion uint32) *_ConfigurationVersionDataType {
	_result := &_ConfigurationVersionDataType{
		MajorVersion:               majorVersion,
		MinorVersion:               minorVersion,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastConfigurationVersionDataType(structType any) ConfigurationVersionDataType {
	if casted, ok := structType.(ConfigurationVersionDataType); ok {
		return casted
	}
	if casted, ok := structType.(*ConfigurationVersionDataType); ok {
		return *casted
	}
	return nil
}

func (m *_ConfigurationVersionDataType) GetTypeName() string {
	return "ConfigurationVersionDataType"
}

func (m *_ConfigurationVersionDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (majorVersion)
	lengthInBits += 32

	// Simple field (minorVersion)
	lengthInBits += 32

	return lengthInBits
}

func (m *_ConfigurationVersionDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ConfigurationVersionDataTypeParse(ctx context.Context, theBytes []byte, identifier string) (ConfigurationVersionDataType, error) {
	return ConfigurationVersionDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func ConfigurationVersionDataTypeParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (ConfigurationVersionDataType, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ConfigurationVersionDataType, error) {
		return ConfigurationVersionDataTypeParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func ConfigurationVersionDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (ConfigurationVersionDataType, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ConfigurationVersionDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ConfigurationVersionDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	majorVersion, err := ReadSimpleField(ctx, "majorVersion", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'majorVersion' field"))
	}

	minorVersion, err := ReadSimpleField(ctx, "minorVersion", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'minorVersion' field"))
	}

	if closeErr := readBuffer.CloseContext("ConfigurationVersionDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ConfigurationVersionDataType")
	}

	// Create a partially initialized instance
	_child := &_ConfigurationVersionDataType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		MajorVersion:               majorVersion,
		MinorVersion:               minorVersion,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_ConfigurationVersionDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ConfigurationVersionDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ConfigurationVersionDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ConfigurationVersionDataType")
		}

		if err := WriteSimpleField[uint32](ctx, "majorVersion", m.GetMajorVersion(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'majorVersion' field")
		}

		if err := WriteSimpleField[uint32](ctx, "minorVersion", m.GetMinorVersion(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'minorVersion' field")
		}

		if popErr := writeBuffer.PopContext("ConfigurationVersionDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ConfigurationVersionDataType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ConfigurationVersionDataType) isConfigurationVersionDataType() bool {
	return true
}

func (m *_ConfigurationVersionDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
