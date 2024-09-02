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

// DataSetWriterTransportDataType is the corresponding interface of DataSetWriterTransportDataType
type DataSetWriterTransportDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// IsDataSetWriterTransportDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDataSetWriterTransportDataType()
}

// _DataSetWriterTransportDataType is the data-structure of this message
type _DataSetWriterTransportDataType struct {
	ExtensionObjectDefinitionContract
}

var _ DataSetWriterTransportDataType = (*_DataSetWriterTransportDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_DataSetWriterTransportDataType)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DataSetWriterTransportDataType) GetIdentifier() string {
	return "15600"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DataSetWriterTransportDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

// NewDataSetWriterTransportDataType factory function for _DataSetWriterTransportDataType
func NewDataSetWriterTransportDataType() *_DataSetWriterTransportDataType {
	_result := &_DataSetWriterTransportDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDataSetWriterTransportDataType(structType any) DataSetWriterTransportDataType {
	if casted, ok := structType.(DataSetWriterTransportDataType); ok {
		return casted
	}
	if casted, ok := structType.(*DataSetWriterTransportDataType); ok {
		return *casted
	}
	return nil
}

func (m *_DataSetWriterTransportDataType) GetTypeName() string {
	return "DataSetWriterTransportDataType"
}

func (m *_DataSetWriterTransportDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_DataSetWriterTransportDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DataSetWriterTransportDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__dataSetWriterTransportDataType DataSetWriterTransportDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DataSetWriterTransportDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DataSetWriterTransportDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("DataSetWriterTransportDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DataSetWriterTransportDataType")
	}

	return m, nil
}

func (m *_DataSetWriterTransportDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DataSetWriterTransportDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DataSetWriterTransportDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DataSetWriterTransportDataType")
		}

		if popErr := writeBuffer.PopContext("DataSetWriterTransportDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DataSetWriterTransportDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DataSetWriterTransportDataType) IsDataSetWriterTransportDataType() {}

func (m *_DataSetWriterTransportDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
