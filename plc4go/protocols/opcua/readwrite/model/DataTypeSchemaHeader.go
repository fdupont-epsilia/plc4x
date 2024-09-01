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

// DataTypeSchemaHeader is the corresponding interface of DataTypeSchemaHeader
type DataTypeSchemaHeader interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNoOfNamespaces returns NoOfNamespaces (property field)
	GetNoOfNamespaces() int32
	// GetNamespaces returns Namespaces (property field)
	GetNamespaces() []PascalString
	// GetNoOfStructureDataTypes returns NoOfStructureDataTypes (property field)
	GetNoOfStructureDataTypes() int32
	// GetStructureDataTypes returns StructureDataTypes (property field)
	GetStructureDataTypes() []DataTypeDescription
	// GetNoOfEnumDataTypes returns NoOfEnumDataTypes (property field)
	GetNoOfEnumDataTypes() int32
	// GetEnumDataTypes returns EnumDataTypes (property field)
	GetEnumDataTypes() []DataTypeDescription
	// GetNoOfSimpleDataTypes returns NoOfSimpleDataTypes (property field)
	GetNoOfSimpleDataTypes() int32
	// GetSimpleDataTypes returns SimpleDataTypes (property field)
	GetSimpleDataTypes() []DataTypeDescription
}

// DataTypeSchemaHeaderExactly can be used when we want exactly this type and not a type which fulfills DataTypeSchemaHeader.
// This is useful for switch cases.
type DataTypeSchemaHeaderExactly interface {
	DataTypeSchemaHeader
	isDataTypeSchemaHeader() bool
}

// _DataTypeSchemaHeader is the data-structure of this message
type _DataTypeSchemaHeader struct {
	ExtensionObjectDefinitionContract
	NoOfNamespaces         int32
	Namespaces             []PascalString
	NoOfStructureDataTypes int32
	StructureDataTypes     []DataTypeDescription
	NoOfEnumDataTypes      int32
	EnumDataTypes          []DataTypeDescription
	NoOfSimpleDataTypes    int32
	SimpleDataTypes        []DataTypeDescription
}

var _ DataTypeSchemaHeader = (*_DataTypeSchemaHeader)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DataTypeSchemaHeader) GetIdentifier() string {
	return "15536"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DataTypeSchemaHeader) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DataTypeSchemaHeader) GetNoOfNamespaces() int32 {
	return m.NoOfNamespaces
}

func (m *_DataTypeSchemaHeader) GetNamespaces() []PascalString {
	return m.Namespaces
}

func (m *_DataTypeSchemaHeader) GetNoOfStructureDataTypes() int32 {
	return m.NoOfStructureDataTypes
}

func (m *_DataTypeSchemaHeader) GetStructureDataTypes() []DataTypeDescription {
	return m.StructureDataTypes
}

func (m *_DataTypeSchemaHeader) GetNoOfEnumDataTypes() int32 {
	return m.NoOfEnumDataTypes
}

func (m *_DataTypeSchemaHeader) GetEnumDataTypes() []DataTypeDescription {
	return m.EnumDataTypes
}

func (m *_DataTypeSchemaHeader) GetNoOfSimpleDataTypes() int32 {
	return m.NoOfSimpleDataTypes
}

func (m *_DataTypeSchemaHeader) GetSimpleDataTypes() []DataTypeDescription {
	return m.SimpleDataTypes
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDataTypeSchemaHeader factory function for _DataTypeSchemaHeader
func NewDataTypeSchemaHeader(noOfNamespaces int32, namespaces []PascalString, noOfStructureDataTypes int32, structureDataTypes []DataTypeDescription, noOfEnumDataTypes int32, enumDataTypes []DataTypeDescription, noOfSimpleDataTypes int32, simpleDataTypes []DataTypeDescription) *_DataTypeSchemaHeader {
	_result := &_DataTypeSchemaHeader{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		NoOfNamespaces:                    noOfNamespaces,
		Namespaces:                        namespaces,
		NoOfStructureDataTypes:            noOfStructureDataTypes,
		StructureDataTypes:                structureDataTypes,
		NoOfEnumDataTypes:                 noOfEnumDataTypes,
		EnumDataTypes:                     enumDataTypes,
		NoOfSimpleDataTypes:               noOfSimpleDataTypes,
		SimpleDataTypes:                   simpleDataTypes,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastDataTypeSchemaHeader(structType any) DataTypeSchemaHeader {
	if casted, ok := structType.(DataTypeSchemaHeader); ok {
		return casted
	}
	if casted, ok := structType.(*DataTypeSchemaHeader); ok {
		return *casted
	}
	return nil
}

func (m *_DataTypeSchemaHeader) GetTypeName() string {
	return "DataTypeSchemaHeader"
}

func (m *_DataTypeSchemaHeader) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (noOfNamespaces)
	lengthInBits += 32

	// Array field
	if len(m.Namespaces) > 0 {
		for _curItem, element := range m.Namespaces {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Namespaces), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfStructureDataTypes)
	lengthInBits += 32

	// Array field
	if len(m.StructureDataTypes) > 0 {
		for _curItem, element := range m.StructureDataTypes {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.StructureDataTypes), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfEnumDataTypes)
	lengthInBits += 32

	// Array field
	if len(m.EnumDataTypes) > 0 {
		for _curItem, element := range m.EnumDataTypes {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.EnumDataTypes), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfSimpleDataTypes)
	lengthInBits += 32

	// Array field
	if len(m.SimpleDataTypes) > 0 {
		for _curItem, element := range m.SimpleDataTypes {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.SimpleDataTypes), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_DataTypeSchemaHeader) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DataTypeSchemaHeader) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (_dataTypeSchemaHeader DataTypeSchemaHeader, err error) {
	m.ExtensionObjectDefinitionContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DataTypeSchemaHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DataTypeSchemaHeader")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	noOfNamespaces, err := ReadSimpleField(ctx, "noOfNamespaces", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfNamespaces' field"))
	}
	m.NoOfNamespaces = noOfNamespaces

	namespaces, err := ReadCountArrayField[PascalString](ctx, "namespaces", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer), uint64(noOfNamespaces))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'namespaces' field"))
	}
	m.Namespaces = namespaces

	noOfStructureDataTypes, err := ReadSimpleField(ctx, "noOfStructureDataTypes", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfStructureDataTypes' field"))
	}
	m.NoOfStructureDataTypes = noOfStructureDataTypes

	structureDataTypes, err := ReadCountArrayField[DataTypeDescription](ctx, "structureDataTypes", ReadComplex[DataTypeDescription](ExtensionObjectDefinitionParseWithBufferProducer[DataTypeDescription]((string)("14525")), readBuffer), uint64(noOfStructureDataTypes))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'structureDataTypes' field"))
	}
	m.StructureDataTypes = structureDataTypes

	noOfEnumDataTypes, err := ReadSimpleField(ctx, "noOfEnumDataTypes", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfEnumDataTypes' field"))
	}
	m.NoOfEnumDataTypes = noOfEnumDataTypes

	enumDataTypes, err := ReadCountArrayField[DataTypeDescription](ctx, "enumDataTypes", ReadComplex[DataTypeDescription](ExtensionObjectDefinitionParseWithBufferProducer[DataTypeDescription]((string)("14525")), readBuffer), uint64(noOfEnumDataTypes))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'enumDataTypes' field"))
	}
	m.EnumDataTypes = enumDataTypes

	noOfSimpleDataTypes, err := ReadSimpleField(ctx, "noOfSimpleDataTypes", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfSimpleDataTypes' field"))
	}
	m.NoOfSimpleDataTypes = noOfSimpleDataTypes

	simpleDataTypes, err := ReadCountArrayField[DataTypeDescription](ctx, "simpleDataTypes", ReadComplex[DataTypeDescription](ExtensionObjectDefinitionParseWithBufferProducer[DataTypeDescription]((string)("14525")), readBuffer), uint64(noOfSimpleDataTypes))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'simpleDataTypes' field"))
	}
	m.SimpleDataTypes = simpleDataTypes

	if closeErr := readBuffer.CloseContext("DataTypeSchemaHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DataTypeSchemaHeader")
	}

	return m, nil
}

func (m *_DataTypeSchemaHeader) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DataTypeSchemaHeader) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DataTypeSchemaHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DataTypeSchemaHeader")
		}

		if err := WriteSimpleField[int32](ctx, "noOfNamespaces", m.GetNoOfNamespaces(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfNamespaces' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "namespaces", m.GetNamespaces(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'namespaces' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfStructureDataTypes", m.GetNoOfStructureDataTypes(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfStructureDataTypes' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "structureDataTypes", m.GetStructureDataTypes(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'structureDataTypes' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfEnumDataTypes", m.GetNoOfEnumDataTypes(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfEnumDataTypes' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "enumDataTypes", m.GetEnumDataTypes(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'enumDataTypes' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfSimpleDataTypes", m.GetNoOfSimpleDataTypes(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfSimpleDataTypes' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "simpleDataTypes", m.GetSimpleDataTypes(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'simpleDataTypes' field")
		}

		if popErr := writeBuffer.PopContext("DataTypeSchemaHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DataTypeSchemaHeader")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DataTypeSchemaHeader) isDataTypeSchemaHeader() bool {
	return true
}

func (m *_DataTypeSchemaHeader) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
