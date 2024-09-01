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

// DataSetWriterMessageDataType is the corresponding interface of DataSetWriterMessageDataType
type DataSetWriterMessageDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
}

// DataSetWriterMessageDataTypeExactly can be used when we want exactly this type and not a type which fulfills DataSetWriterMessageDataType.
// This is useful for switch cases.
type DataSetWriterMessageDataTypeExactly interface {
	DataSetWriterMessageDataType
	isDataSetWriterMessageDataType() bool
}

// _DataSetWriterMessageDataType is the data-structure of this message
type _DataSetWriterMessageDataType struct {
	*_ExtensionObjectDefinition
}

var _ DataSetWriterMessageDataType = (*_DataSetWriterMessageDataType)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DataSetWriterMessageDataType) GetIdentifier() string {
	return "15607"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DataSetWriterMessageDataType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_DataSetWriterMessageDataType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

// NewDataSetWriterMessageDataType factory function for _DataSetWriterMessageDataType
func NewDataSetWriterMessageDataType() *_DataSetWriterMessageDataType {
	_result := &_DataSetWriterMessageDataType{
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDataSetWriterMessageDataType(structType any) DataSetWriterMessageDataType {
	if casted, ok := structType.(DataSetWriterMessageDataType); ok {
		return casted
	}
	if casted, ok := structType.(*DataSetWriterMessageDataType); ok {
		return *casted
	}
	return nil
}

func (m *_DataSetWriterMessageDataType) GetTypeName() string {
	return "DataSetWriterMessageDataType"
}

func (m *_DataSetWriterMessageDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_DataSetWriterMessageDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DataSetWriterMessageDataTypeParse(ctx context.Context, theBytes []byte, identifier string) (DataSetWriterMessageDataType, error) {
	return DataSetWriterMessageDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func DataSetWriterMessageDataTypeParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (DataSetWriterMessageDataType, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (DataSetWriterMessageDataType, error) {
		return DataSetWriterMessageDataTypeParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func DataSetWriterMessageDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (DataSetWriterMessageDataType, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DataSetWriterMessageDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DataSetWriterMessageDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("DataSetWriterMessageDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DataSetWriterMessageDataType")
	}

	// Create a partially initialized instance
	_child := &_DataSetWriterMessageDataType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_DataSetWriterMessageDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DataSetWriterMessageDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DataSetWriterMessageDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DataSetWriterMessageDataType")
		}

		if popErr := writeBuffer.PopContext("DataSetWriterMessageDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DataSetWriterMessageDataType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DataSetWriterMessageDataType) isDataSetWriterMessageDataType() bool {
	return true
}

func (m *_DataSetWriterMessageDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
