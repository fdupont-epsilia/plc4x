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

// PortableQualifiedName is the corresponding interface of PortableQualifiedName
type PortableQualifiedName interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNamespaceUri returns NamespaceUri (property field)
	GetNamespaceUri() PascalString
	// GetName returns Name (property field)
	GetName() PascalString
	// IsPortableQualifiedName is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsPortableQualifiedName()
}

// _PortableQualifiedName is the data-structure of this message
type _PortableQualifiedName struct {
	ExtensionObjectDefinitionContract
	NamespaceUri PascalString
	Name         PascalString
}

var _ PortableQualifiedName = (*_PortableQualifiedName)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_PortableQualifiedName)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_PortableQualifiedName) GetIdentifier() string {
	return "24107"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_PortableQualifiedName) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PortableQualifiedName) GetNamespaceUri() PascalString {
	return m.NamespaceUri
}

func (m *_PortableQualifiedName) GetName() PascalString {
	return m.Name
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewPortableQualifiedName factory function for _PortableQualifiedName
func NewPortableQualifiedName(namespaceUri PascalString, name PascalString) *_PortableQualifiedName {
	_result := &_PortableQualifiedName{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		NamespaceUri:                      namespaceUri,
		Name:                              name,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastPortableQualifiedName(structType any) PortableQualifiedName {
	if casted, ok := structType.(PortableQualifiedName); ok {
		return casted
	}
	if casted, ok := structType.(*PortableQualifiedName); ok {
		return *casted
	}
	return nil
}

func (m *_PortableQualifiedName) GetTypeName() string {
	return "PortableQualifiedName"
}

func (m *_PortableQualifiedName) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (namespaceUri)
	lengthInBits += m.NamespaceUri.GetLengthInBits(ctx)

	// Simple field (name)
	lengthInBits += m.Name.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_PortableQualifiedName) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_PortableQualifiedName) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__portableQualifiedName PortableQualifiedName, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("PortableQualifiedName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PortableQualifiedName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	namespaceUri, err := ReadSimpleField[PascalString](ctx, "namespaceUri", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'namespaceUri' field"))
	}
	m.NamespaceUri = namespaceUri

	name, err := ReadSimpleField[PascalString](ctx, "name", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'name' field"))
	}
	m.Name = name

	if closeErr := readBuffer.CloseContext("PortableQualifiedName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PortableQualifiedName")
	}

	return m, nil
}

func (m *_PortableQualifiedName) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PortableQualifiedName) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PortableQualifiedName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PortableQualifiedName")
		}

		if err := WriteSimpleField[PascalString](ctx, "namespaceUri", m.GetNamespaceUri(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'namespaceUri' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "name", m.GetName(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'name' field")
		}

		if popErr := writeBuffer.PopContext("PortableQualifiedName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PortableQualifiedName")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_PortableQualifiedName) IsPortableQualifiedName() {}

func (m *_PortableQualifiedName) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
