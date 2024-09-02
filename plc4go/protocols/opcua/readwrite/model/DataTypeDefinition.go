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

// DataTypeDefinition is the corresponding interface of DataTypeDefinition
type DataTypeDefinition interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// IsDataTypeDefinition is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDataTypeDefinition()
}

// _DataTypeDefinition is the data-structure of this message
type _DataTypeDefinition struct {
	ExtensionObjectDefinitionContract
}

var _ DataTypeDefinition = (*_DataTypeDefinition)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_DataTypeDefinition)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DataTypeDefinition) GetIdentifier() string {
	return "99"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DataTypeDefinition) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

// NewDataTypeDefinition factory function for _DataTypeDefinition
func NewDataTypeDefinition() *_DataTypeDefinition {
	_result := &_DataTypeDefinition{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDataTypeDefinition(structType any) DataTypeDefinition {
	if casted, ok := structType.(DataTypeDefinition); ok {
		return casted
	}
	if casted, ok := structType.(*DataTypeDefinition); ok {
		return *casted
	}
	return nil
}

func (m *_DataTypeDefinition) GetTypeName() string {
	return "DataTypeDefinition"
}

func (m *_DataTypeDefinition) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_DataTypeDefinition) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DataTypeDefinition) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__dataTypeDefinition DataTypeDefinition, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DataTypeDefinition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DataTypeDefinition")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("DataTypeDefinition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DataTypeDefinition")
	}

	return m, nil
}

func (m *_DataTypeDefinition) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DataTypeDefinition) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DataTypeDefinition"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DataTypeDefinition")
		}

		if popErr := writeBuffer.PopContext("DataTypeDefinition"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DataTypeDefinition")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DataTypeDefinition) IsDataTypeDefinition() {}

func (m *_DataTypeDefinition) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
