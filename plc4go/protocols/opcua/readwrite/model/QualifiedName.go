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

// QualifiedName is the corresponding interface of QualifiedName
type QualifiedName interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetNamespaceIndex returns NamespaceIndex (property field)
	GetNamespaceIndex() uint16
	// GetName returns Name (property field)
	GetName() PascalString
}

// QualifiedNameExactly can be used when we want exactly this type and not a type which fulfills QualifiedName.
// This is useful for switch cases.
type QualifiedNameExactly interface {
	QualifiedName
	isQualifiedName() bool
}

// _QualifiedName is the data-structure of this message
type _QualifiedName struct {
	NamespaceIndex uint16
	Name           PascalString
}

var _ QualifiedName = (*_QualifiedName)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_QualifiedName) GetNamespaceIndex() uint16 {
	return m.NamespaceIndex
}

func (m *_QualifiedName) GetName() PascalString {
	return m.Name
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewQualifiedName factory function for _QualifiedName
func NewQualifiedName(namespaceIndex uint16, name PascalString) *_QualifiedName {
	return &_QualifiedName{NamespaceIndex: namespaceIndex, Name: name}
}

// Deprecated: use the interface for direct cast
func CastQualifiedName(structType any) QualifiedName {
	if casted, ok := structType.(QualifiedName); ok {
		return casted
	}
	if casted, ok := structType.(*QualifiedName); ok {
		return *casted
	}
	return nil
}

func (m *_QualifiedName) GetTypeName() string {
	return "QualifiedName"
}

func (m *_QualifiedName) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (namespaceIndex)
	lengthInBits += 16

	// Simple field (name)
	lengthInBits += m.Name.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_QualifiedName) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func QualifiedNameParse(ctx context.Context, theBytes []byte) (QualifiedName, error) {
	return QualifiedNameParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func QualifiedNameParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (QualifiedName, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (QualifiedName, error) {
		return QualifiedNameParseWithBuffer(ctx, readBuffer)
	}
}

func QualifiedNameParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (QualifiedName, error) {
	v, err := (&_QualifiedName{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_QualifiedName) parse(ctx context.Context, readBuffer utils.ReadBuffer) (_qualifiedName QualifiedName, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("QualifiedName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for QualifiedName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	namespaceIndex, err := ReadSimpleField(ctx, "namespaceIndex", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'namespaceIndex' field"))
	}
	m.NamespaceIndex = namespaceIndex

	name, err := ReadSimpleField[PascalString](ctx, "name", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'name' field"))
	}
	m.Name = name

	if closeErr := readBuffer.CloseContext("QualifiedName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for QualifiedName")
	}

	return m, nil
}

func (m *_QualifiedName) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_QualifiedName) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("QualifiedName"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for QualifiedName")
	}

	if err := WriteSimpleField[uint16](ctx, "namespaceIndex", m.GetNamespaceIndex(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'namespaceIndex' field")
	}

	if err := WriteSimpleField[PascalString](ctx, "name", m.GetName(), WriteComplex[PascalString](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'name' field")
	}

	if popErr := writeBuffer.PopContext("QualifiedName"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for QualifiedName")
	}
	return nil
}

func (m *_QualifiedName) isQualifiedName() bool {
	return true
}

func (m *_QualifiedName) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
