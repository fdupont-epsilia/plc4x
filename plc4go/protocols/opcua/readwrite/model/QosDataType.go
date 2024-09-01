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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// QosDataType is the corresponding interface of QosDataType
type QosDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
}

// QosDataTypeExactly can be used when we want exactly this type and not a type which fulfills QosDataType.
// This is useful for switch cases.
type QosDataTypeExactly interface {
	QosDataType
	isQosDataType() bool
}

// _QosDataType is the data-structure of this message
type _QosDataType struct {
	*_ExtensionObjectDefinition
}

var _ QosDataType = (*_QosDataType)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_QosDataType) GetIdentifier() string {
	return "23605"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_QosDataType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_QosDataType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

// NewQosDataType factory function for _QosDataType
func NewQosDataType() *_QosDataType {
	_result := &_QosDataType{
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastQosDataType(structType any) QosDataType {
	if casted, ok := structType.(QosDataType); ok {
		return casted
	}
	if casted, ok := structType.(*QosDataType); ok {
		return *casted
	}
	return nil
}

func (m *_QosDataType) GetTypeName() string {
	return "QosDataType"
}

func (m *_QosDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_QosDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func QosDataTypeParse(ctx context.Context, theBytes []byte, identifier string) (QosDataType, error) {
	return QosDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func QosDataTypeParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (QosDataType, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (QosDataType, error) {
		return QosDataTypeParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func QosDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (QosDataType, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("QosDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for QosDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("QosDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for QosDataType")
	}

	// Create a partially initialized instance
	_child := &_QosDataType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_QosDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_QosDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("QosDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for QosDataType")
		}

		if popErr := writeBuffer.PopContext("QosDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for QosDataType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_QosDataType) isQosDataType() bool {
	return true
}

func (m *_QosDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
