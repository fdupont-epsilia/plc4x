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

// CartesianCoordinates is the corresponding interface of CartesianCoordinates
type CartesianCoordinates interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// IsCartesianCoordinates is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCartesianCoordinates()
}

// _CartesianCoordinates is the data-structure of this message
type _CartesianCoordinates struct {
	ExtensionObjectDefinitionContract
}

var _ CartesianCoordinates = (*_CartesianCoordinates)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_CartesianCoordinates)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CartesianCoordinates) GetIdentifier() string {
	return "18811"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CartesianCoordinates) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

// NewCartesianCoordinates factory function for _CartesianCoordinates
func NewCartesianCoordinates() *_CartesianCoordinates {
	_result := &_CartesianCoordinates{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCartesianCoordinates(structType any) CartesianCoordinates {
	if casted, ok := structType.(CartesianCoordinates); ok {
		return casted
	}
	if casted, ok := structType.(*CartesianCoordinates); ok {
		return *casted
	}
	return nil
}

func (m *_CartesianCoordinates) GetTypeName() string {
	return "CartesianCoordinates"
}

func (m *_CartesianCoordinates) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_CartesianCoordinates) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CartesianCoordinates) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__cartesianCoordinates CartesianCoordinates, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CartesianCoordinates"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CartesianCoordinates")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("CartesianCoordinates"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CartesianCoordinates")
	}

	return m, nil
}

func (m *_CartesianCoordinates) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CartesianCoordinates) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CartesianCoordinates"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CartesianCoordinates")
		}

		if popErr := writeBuffer.PopContext("CartesianCoordinates"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CartesianCoordinates")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CartesianCoordinates) IsCartesianCoordinates() {}

func (m *_CartesianCoordinates) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
