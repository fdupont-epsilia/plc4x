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

// ConnectionTransportDataType is the corresponding interface of ConnectionTransportDataType
type ConnectionTransportDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
}

// ConnectionTransportDataTypeExactly can be used when we want exactly this type and not a type which fulfills ConnectionTransportDataType.
// This is useful for switch cases.
type ConnectionTransportDataTypeExactly interface {
	ConnectionTransportDataType
	isConnectionTransportDataType() bool
}

// _ConnectionTransportDataType is the data-structure of this message
type _ConnectionTransportDataType struct {
	*_ExtensionObjectDefinition
}

var _ ConnectionTransportDataType = (*_ConnectionTransportDataType)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ConnectionTransportDataType) GetIdentifier() string {
	return "15620"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ConnectionTransportDataType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_ConnectionTransportDataType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

// NewConnectionTransportDataType factory function for _ConnectionTransportDataType
func NewConnectionTransportDataType() *_ConnectionTransportDataType {
	_result := &_ConnectionTransportDataType{
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastConnectionTransportDataType(structType any) ConnectionTransportDataType {
	if casted, ok := structType.(ConnectionTransportDataType); ok {
		return casted
	}
	if casted, ok := structType.(*ConnectionTransportDataType); ok {
		return *casted
	}
	return nil
}

func (m *_ConnectionTransportDataType) GetTypeName() string {
	return "ConnectionTransportDataType"
}

func (m *_ConnectionTransportDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_ConnectionTransportDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ConnectionTransportDataTypeParse(ctx context.Context, theBytes []byte, identifier string) (ConnectionTransportDataType, error) {
	return ConnectionTransportDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func ConnectionTransportDataTypeParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (ConnectionTransportDataType, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ConnectionTransportDataType, error) {
		return ConnectionTransportDataTypeParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func ConnectionTransportDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (ConnectionTransportDataType, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ConnectionTransportDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ConnectionTransportDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ConnectionTransportDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ConnectionTransportDataType")
	}

	// Create a partially initialized instance
	_child := &_ConnectionTransportDataType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_ConnectionTransportDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ConnectionTransportDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ConnectionTransportDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ConnectionTransportDataType")
		}

		if popErr := writeBuffer.PopContext("ConnectionTransportDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ConnectionTransportDataType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ConnectionTransportDataType) isConnectionTransportDataType() bool {
	return true
}

func (m *_ConnectionTransportDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
