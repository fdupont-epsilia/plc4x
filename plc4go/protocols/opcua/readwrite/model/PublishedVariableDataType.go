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

// PublishedVariableDataType is the corresponding interface of PublishedVariableDataType
type PublishedVariableDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetPublishedVariable returns PublishedVariable (property field)
	GetPublishedVariable() NodeId
	// GetAttributeId returns AttributeId (property field)
	GetAttributeId() uint32
	// GetSamplingIntervalHint returns SamplingIntervalHint (property field)
	GetSamplingIntervalHint() float64
	// GetDeadbandType returns DeadbandType (property field)
	GetDeadbandType() uint32
	// GetDeadbandValue returns DeadbandValue (property field)
	GetDeadbandValue() float64
	// GetIndexRange returns IndexRange (property field)
	GetIndexRange() PascalString
	// GetSubstituteValue returns SubstituteValue (property field)
	GetSubstituteValue() Variant
	// GetNoOfMetaDataProperties returns NoOfMetaDataProperties (property field)
	GetNoOfMetaDataProperties() int32
	// GetMetaDataProperties returns MetaDataProperties (property field)
	GetMetaDataProperties() []QualifiedName
	// IsPublishedVariableDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsPublishedVariableDataType()
}

// _PublishedVariableDataType is the data-structure of this message
type _PublishedVariableDataType struct {
	ExtensionObjectDefinitionContract
	PublishedVariable      NodeId
	AttributeId            uint32
	SamplingIntervalHint   float64
	DeadbandType           uint32
	DeadbandValue          float64
	IndexRange             PascalString
	SubstituteValue        Variant
	NoOfMetaDataProperties int32
	MetaDataProperties     []QualifiedName
}

var _ PublishedVariableDataType = (*_PublishedVariableDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_PublishedVariableDataType)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_PublishedVariableDataType) GetIdentifier() string {
	return "14275"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_PublishedVariableDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PublishedVariableDataType) GetPublishedVariable() NodeId {
	return m.PublishedVariable
}

func (m *_PublishedVariableDataType) GetAttributeId() uint32 {
	return m.AttributeId
}

func (m *_PublishedVariableDataType) GetSamplingIntervalHint() float64 {
	return m.SamplingIntervalHint
}

func (m *_PublishedVariableDataType) GetDeadbandType() uint32 {
	return m.DeadbandType
}

func (m *_PublishedVariableDataType) GetDeadbandValue() float64 {
	return m.DeadbandValue
}

func (m *_PublishedVariableDataType) GetIndexRange() PascalString {
	return m.IndexRange
}

func (m *_PublishedVariableDataType) GetSubstituteValue() Variant {
	return m.SubstituteValue
}

func (m *_PublishedVariableDataType) GetNoOfMetaDataProperties() int32 {
	return m.NoOfMetaDataProperties
}

func (m *_PublishedVariableDataType) GetMetaDataProperties() []QualifiedName {
	return m.MetaDataProperties
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewPublishedVariableDataType factory function for _PublishedVariableDataType
func NewPublishedVariableDataType(publishedVariable NodeId, attributeId uint32, samplingIntervalHint float64, deadbandType uint32, deadbandValue float64, indexRange PascalString, substituteValue Variant, noOfMetaDataProperties int32, metaDataProperties []QualifiedName) *_PublishedVariableDataType {
	_result := &_PublishedVariableDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		PublishedVariable:                 publishedVariable,
		AttributeId:                       attributeId,
		SamplingIntervalHint:              samplingIntervalHint,
		DeadbandType:                      deadbandType,
		DeadbandValue:                     deadbandValue,
		IndexRange:                        indexRange,
		SubstituteValue:                   substituteValue,
		NoOfMetaDataProperties:            noOfMetaDataProperties,
		MetaDataProperties:                metaDataProperties,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastPublishedVariableDataType(structType any) PublishedVariableDataType {
	if casted, ok := structType.(PublishedVariableDataType); ok {
		return casted
	}
	if casted, ok := structType.(*PublishedVariableDataType); ok {
		return *casted
	}
	return nil
}

func (m *_PublishedVariableDataType) GetTypeName() string {
	return "PublishedVariableDataType"
}

func (m *_PublishedVariableDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (publishedVariable)
	lengthInBits += m.PublishedVariable.GetLengthInBits(ctx)

	// Simple field (attributeId)
	lengthInBits += 32

	// Simple field (samplingIntervalHint)
	lengthInBits += 64

	// Simple field (deadbandType)
	lengthInBits += 32

	// Simple field (deadbandValue)
	lengthInBits += 64

	// Simple field (indexRange)
	lengthInBits += m.IndexRange.GetLengthInBits(ctx)

	// Simple field (substituteValue)
	lengthInBits += m.SubstituteValue.GetLengthInBits(ctx)

	// Simple field (noOfMetaDataProperties)
	lengthInBits += 32

	// Array field
	if len(m.MetaDataProperties) > 0 {
		for _curItem, element := range m.MetaDataProperties {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.MetaDataProperties), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_PublishedVariableDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_PublishedVariableDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__publishedVariableDataType PublishedVariableDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("PublishedVariableDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PublishedVariableDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	publishedVariable, err := ReadSimpleField[NodeId](ctx, "publishedVariable", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'publishedVariable' field"))
	}
	m.PublishedVariable = publishedVariable

	attributeId, err := ReadSimpleField(ctx, "attributeId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'attributeId' field"))
	}
	m.AttributeId = attributeId

	samplingIntervalHint, err := ReadSimpleField(ctx, "samplingIntervalHint", ReadDouble(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'samplingIntervalHint' field"))
	}
	m.SamplingIntervalHint = samplingIntervalHint

	deadbandType, err := ReadSimpleField(ctx, "deadbandType", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deadbandType' field"))
	}
	m.DeadbandType = deadbandType

	deadbandValue, err := ReadSimpleField(ctx, "deadbandValue", ReadDouble(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deadbandValue' field"))
	}
	m.DeadbandValue = deadbandValue

	indexRange, err := ReadSimpleField[PascalString](ctx, "indexRange", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'indexRange' field"))
	}
	m.IndexRange = indexRange

	substituteValue, err := ReadSimpleField[Variant](ctx, "substituteValue", ReadComplex[Variant](VariantParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'substituteValue' field"))
	}
	m.SubstituteValue = substituteValue

	noOfMetaDataProperties, err := ReadSimpleField(ctx, "noOfMetaDataProperties", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfMetaDataProperties' field"))
	}
	m.NoOfMetaDataProperties = noOfMetaDataProperties

	metaDataProperties, err := ReadCountArrayField[QualifiedName](ctx, "metaDataProperties", ReadComplex[QualifiedName](QualifiedNameParseWithBuffer, readBuffer), uint64(noOfMetaDataProperties))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'metaDataProperties' field"))
	}
	m.MetaDataProperties = metaDataProperties

	if closeErr := readBuffer.CloseContext("PublishedVariableDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PublishedVariableDataType")
	}

	return m, nil
}

func (m *_PublishedVariableDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PublishedVariableDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PublishedVariableDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PublishedVariableDataType")
		}

		if err := WriteSimpleField[NodeId](ctx, "publishedVariable", m.GetPublishedVariable(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'publishedVariable' field")
		}

		if err := WriteSimpleField[uint32](ctx, "attributeId", m.GetAttributeId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'attributeId' field")
		}

		if err := WriteSimpleField[float64](ctx, "samplingIntervalHint", m.GetSamplingIntervalHint(), WriteDouble(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'samplingIntervalHint' field")
		}

		if err := WriteSimpleField[uint32](ctx, "deadbandType", m.GetDeadbandType(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'deadbandType' field")
		}

		if err := WriteSimpleField[float64](ctx, "deadbandValue", m.GetDeadbandValue(), WriteDouble(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'deadbandValue' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "indexRange", m.GetIndexRange(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'indexRange' field")
		}

		if err := WriteSimpleField[Variant](ctx, "substituteValue", m.GetSubstituteValue(), WriteComplex[Variant](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'substituteValue' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfMetaDataProperties", m.GetNoOfMetaDataProperties(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfMetaDataProperties' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "metaDataProperties", m.GetMetaDataProperties(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'metaDataProperties' field")
		}

		if popErr := writeBuffer.PopContext("PublishedVariableDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PublishedVariableDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_PublishedVariableDataType) IsPublishedVariableDataType() {}

func (m *_PublishedVariableDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
