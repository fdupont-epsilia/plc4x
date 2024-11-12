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

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// DataSetWriterDataType is the corresponding interface of DataSetWriterDataType
type DataSetWriterDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetName returns Name (property field)
	GetName() PascalString
	// GetEnabled returns Enabled (property field)
	GetEnabled() bool
	// GetDataSetWriterId returns DataSetWriterId (property field)
	GetDataSetWriterId() uint16
	// GetDataSetFieldContentMask returns DataSetFieldContentMask (property field)
	GetDataSetFieldContentMask() DataSetFieldContentMask
	// GetKeyFrameCount returns KeyFrameCount (property field)
	GetKeyFrameCount() uint32
	// GetDataSetName returns DataSetName (property field)
	GetDataSetName() PascalString
	// GetDataSetWriterProperties returns DataSetWriterProperties (property field)
	GetDataSetWriterProperties() []KeyValuePair
	// GetTransportSettings returns TransportSettings (property field)
	GetTransportSettings() ExtensionObject
	// GetMessageSettings returns MessageSettings (property field)
	GetMessageSettings() ExtensionObject
	// IsDataSetWriterDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDataSetWriterDataType()
	// CreateBuilder creates a DataSetWriterDataTypeBuilder
	CreateDataSetWriterDataTypeBuilder() DataSetWriterDataTypeBuilder
}

// _DataSetWriterDataType is the data-structure of this message
type _DataSetWriterDataType struct {
	ExtensionObjectDefinitionContract
	Name                    PascalString
	Enabled                 bool
	DataSetWriterId         uint16
	DataSetFieldContentMask DataSetFieldContentMask
	KeyFrameCount           uint32
	DataSetName             PascalString
	DataSetWriterProperties []KeyValuePair
	TransportSettings       ExtensionObject
	MessageSettings         ExtensionObject
	// Reserved Fields
	reservedField0 *uint8
}

var _ DataSetWriterDataType = (*_DataSetWriterDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_DataSetWriterDataType)(nil)

// NewDataSetWriterDataType factory function for _DataSetWriterDataType
func NewDataSetWriterDataType(name PascalString, enabled bool, dataSetWriterId uint16, dataSetFieldContentMask DataSetFieldContentMask, keyFrameCount uint32, dataSetName PascalString, dataSetWriterProperties []KeyValuePair, transportSettings ExtensionObject, messageSettings ExtensionObject) *_DataSetWriterDataType {
	if name == nil {
		panic("name of type PascalString for DataSetWriterDataType must not be nil")
	}
	if dataSetName == nil {
		panic("dataSetName of type PascalString for DataSetWriterDataType must not be nil")
	}
	if transportSettings == nil {
		panic("transportSettings of type ExtensionObject for DataSetWriterDataType must not be nil")
	}
	if messageSettings == nil {
		panic("messageSettings of type ExtensionObject for DataSetWriterDataType must not be nil")
	}
	_result := &_DataSetWriterDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Name:                              name,
		Enabled:                           enabled,
		DataSetWriterId:                   dataSetWriterId,
		DataSetFieldContentMask:           dataSetFieldContentMask,
		KeyFrameCount:                     keyFrameCount,
		DataSetName:                       dataSetName,
		DataSetWriterProperties:           dataSetWriterProperties,
		TransportSettings:                 transportSettings,
		MessageSettings:                   messageSettings,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DataSetWriterDataTypeBuilder is a builder for DataSetWriterDataType
type DataSetWriterDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(name PascalString, enabled bool, dataSetWriterId uint16, dataSetFieldContentMask DataSetFieldContentMask, keyFrameCount uint32, dataSetName PascalString, dataSetWriterProperties []KeyValuePair, transportSettings ExtensionObject, messageSettings ExtensionObject) DataSetWriterDataTypeBuilder
	// WithName adds Name (property field)
	WithName(PascalString) DataSetWriterDataTypeBuilder
	// WithNameBuilder adds Name (property field) which is build by the builder
	WithNameBuilder(func(PascalStringBuilder) PascalStringBuilder) DataSetWriterDataTypeBuilder
	// WithEnabled adds Enabled (property field)
	WithEnabled(bool) DataSetWriterDataTypeBuilder
	// WithDataSetWriterId adds DataSetWriterId (property field)
	WithDataSetWriterId(uint16) DataSetWriterDataTypeBuilder
	// WithDataSetFieldContentMask adds DataSetFieldContentMask (property field)
	WithDataSetFieldContentMask(DataSetFieldContentMask) DataSetWriterDataTypeBuilder
	// WithKeyFrameCount adds KeyFrameCount (property field)
	WithKeyFrameCount(uint32) DataSetWriterDataTypeBuilder
	// WithDataSetName adds DataSetName (property field)
	WithDataSetName(PascalString) DataSetWriterDataTypeBuilder
	// WithDataSetNameBuilder adds DataSetName (property field) which is build by the builder
	WithDataSetNameBuilder(func(PascalStringBuilder) PascalStringBuilder) DataSetWriterDataTypeBuilder
	// WithDataSetWriterProperties adds DataSetWriterProperties (property field)
	WithDataSetWriterProperties(...KeyValuePair) DataSetWriterDataTypeBuilder
	// WithTransportSettings adds TransportSettings (property field)
	WithTransportSettings(ExtensionObject) DataSetWriterDataTypeBuilder
	// WithTransportSettingsBuilder adds TransportSettings (property field) which is build by the builder
	WithTransportSettingsBuilder(func(ExtensionObjectBuilder) ExtensionObjectBuilder) DataSetWriterDataTypeBuilder
	// WithMessageSettings adds MessageSettings (property field)
	WithMessageSettings(ExtensionObject) DataSetWriterDataTypeBuilder
	// WithMessageSettingsBuilder adds MessageSettings (property field) which is build by the builder
	WithMessageSettingsBuilder(func(ExtensionObjectBuilder) ExtensionObjectBuilder) DataSetWriterDataTypeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the DataSetWriterDataType or returns an error if something is wrong
	Build() (DataSetWriterDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DataSetWriterDataType
}

// NewDataSetWriterDataTypeBuilder() creates a DataSetWriterDataTypeBuilder
func NewDataSetWriterDataTypeBuilder() DataSetWriterDataTypeBuilder {
	return &_DataSetWriterDataTypeBuilder{_DataSetWriterDataType: new(_DataSetWriterDataType)}
}

type _DataSetWriterDataTypeBuilder struct {
	*_DataSetWriterDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (DataSetWriterDataTypeBuilder) = (*_DataSetWriterDataTypeBuilder)(nil)

func (b *_DataSetWriterDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._DataSetWriterDataType
}

func (b *_DataSetWriterDataTypeBuilder) WithMandatoryFields(name PascalString, enabled bool, dataSetWriterId uint16, dataSetFieldContentMask DataSetFieldContentMask, keyFrameCount uint32, dataSetName PascalString, dataSetWriterProperties []KeyValuePair, transportSettings ExtensionObject, messageSettings ExtensionObject) DataSetWriterDataTypeBuilder {
	return b.WithName(name).WithEnabled(enabled).WithDataSetWriterId(dataSetWriterId).WithDataSetFieldContentMask(dataSetFieldContentMask).WithKeyFrameCount(keyFrameCount).WithDataSetName(dataSetName).WithDataSetWriterProperties(dataSetWriterProperties...).WithTransportSettings(transportSettings).WithMessageSettings(messageSettings)
}

func (b *_DataSetWriterDataTypeBuilder) WithName(name PascalString) DataSetWriterDataTypeBuilder {
	b.Name = name
	return b
}

func (b *_DataSetWriterDataTypeBuilder) WithNameBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) DataSetWriterDataTypeBuilder {
	builder := builderSupplier(b.Name.CreatePascalStringBuilder())
	var err error
	b.Name, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_DataSetWriterDataTypeBuilder) WithEnabled(enabled bool) DataSetWriterDataTypeBuilder {
	b.Enabled = enabled
	return b
}

func (b *_DataSetWriterDataTypeBuilder) WithDataSetWriterId(dataSetWriterId uint16) DataSetWriterDataTypeBuilder {
	b.DataSetWriterId = dataSetWriterId
	return b
}

func (b *_DataSetWriterDataTypeBuilder) WithDataSetFieldContentMask(dataSetFieldContentMask DataSetFieldContentMask) DataSetWriterDataTypeBuilder {
	b.DataSetFieldContentMask = dataSetFieldContentMask
	return b
}

func (b *_DataSetWriterDataTypeBuilder) WithKeyFrameCount(keyFrameCount uint32) DataSetWriterDataTypeBuilder {
	b.KeyFrameCount = keyFrameCount
	return b
}

func (b *_DataSetWriterDataTypeBuilder) WithDataSetName(dataSetName PascalString) DataSetWriterDataTypeBuilder {
	b.DataSetName = dataSetName
	return b
}

func (b *_DataSetWriterDataTypeBuilder) WithDataSetNameBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) DataSetWriterDataTypeBuilder {
	builder := builderSupplier(b.DataSetName.CreatePascalStringBuilder())
	var err error
	b.DataSetName, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_DataSetWriterDataTypeBuilder) WithDataSetWriterProperties(dataSetWriterProperties ...KeyValuePair) DataSetWriterDataTypeBuilder {
	b.DataSetWriterProperties = dataSetWriterProperties
	return b
}

func (b *_DataSetWriterDataTypeBuilder) WithTransportSettings(transportSettings ExtensionObject) DataSetWriterDataTypeBuilder {
	b.TransportSettings = transportSettings
	return b
}

func (b *_DataSetWriterDataTypeBuilder) WithTransportSettingsBuilder(builderSupplier func(ExtensionObjectBuilder) ExtensionObjectBuilder) DataSetWriterDataTypeBuilder {
	builder := builderSupplier(b.TransportSettings.CreateExtensionObjectBuilder())
	var err error
	b.TransportSettings, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExtensionObjectBuilder failed"))
	}
	return b
}

func (b *_DataSetWriterDataTypeBuilder) WithMessageSettings(messageSettings ExtensionObject) DataSetWriterDataTypeBuilder {
	b.MessageSettings = messageSettings
	return b
}

func (b *_DataSetWriterDataTypeBuilder) WithMessageSettingsBuilder(builderSupplier func(ExtensionObjectBuilder) ExtensionObjectBuilder) DataSetWriterDataTypeBuilder {
	builder := builderSupplier(b.MessageSettings.CreateExtensionObjectBuilder())
	var err error
	b.MessageSettings, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExtensionObjectBuilder failed"))
	}
	return b
}

func (b *_DataSetWriterDataTypeBuilder) Build() (DataSetWriterDataType, error) {
	if b.Name == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'name' not set"))
	}
	if b.DataSetName == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'dataSetName' not set"))
	}
	if b.TransportSettings == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'transportSettings' not set"))
	}
	if b.MessageSettings == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'messageSettings' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._DataSetWriterDataType.deepCopy(), nil
}

func (b *_DataSetWriterDataTypeBuilder) MustBuild() DataSetWriterDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_DataSetWriterDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_DataSetWriterDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_DataSetWriterDataTypeBuilder) DeepCopy() any {
	_copy := b.CreateDataSetWriterDataTypeBuilder().(*_DataSetWriterDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateDataSetWriterDataTypeBuilder creates a DataSetWriterDataTypeBuilder
func (b *_DataSetWriterDataType) CreateDataSetWriterDataTypeBuilder() DataSetWriterDataTypeBuilder {
	if b == nil {
		return NewDataSetWriterDataTypeBuilder()
	}
	return &_DataSetWriterDataTypeBuilder{_DataSetWriterDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DataSetWriterDataType) GetExtensionId() int32 {
	return int32(15599)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DataSetWriterDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DataSetWriterDataType) GetName() PascalString {
	return m.Name
}

func (m *_DataSetWriterDataType) GetEnabled() bool {
	return m.Enabled
}

func (m *_DataSetWriterDataType) GetDataSetWriterId() uint16 {
	return m.DataSetWriterId
}

func (m *_DataSetWriterDataType) GetDataSetFieldContentMask() DataSetFieldContentMask {
	return m.DataSetFieldContentMask
}

func (m *_DataSetWriterDataType) GetKeyFrameCount() uint32 {
	return m.KeyFrameCount
}

func (m *_DataSetWriterDataType) GetDataSetName() PascalString {
	return m.DataSetName
}

func (m *_DataSetWriterDataType) GetDataSetWriterProperties() []KeyValuePair {
	return m.DataSetWriterProperties
}

func (m *_DataSetWriterDataType) GetTransportSettings() ExtensionObject {
	return m.TransportSettings
}

func (m *_DataSetWriterDataType) GetMessageSettings() ExtensionObject {
	return m.MessageSettings
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDataSetWriterDataType(structType any) DataSetWriterDataType {
	if casted, ok := structType.(DataSetWriterDataType); ok {
		return casted
	}
	if casted, ok := structType.(*DataSetWriterDataType); ok {
		return *casted
	}
	return nil
}

func (m *_DataSetWriterDataType) GetTypeName() string {
	return "DataSetWriterDataType"
}

func (m *_DataSetWriterDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (name)
	lengthInBits += m.Name.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (enabled)
	lengthInBits += 1

	// Simple field (dataSetWriterId)
	lengthInBits += 16

	// Simple field (dataSetFieldContentMask)
	lengthInBits += 32

	// Simple field (keyFrameCount)
	lengthInBits += 32

	// Simple field (dataSetName)
	lengthInBits += m.DataSetName.GetLengthInBits(ctx)

	// Implicit Field (noOfDataSetWriterProperties)
	lengthInBits += 32

	// Array field
	if len(m.DataSetWriterProperties) > 0 {
		for _curItem, element := range m.DataSetWriterProperties {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.DataSetWriterProperties), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (transportSettings)
	lengthInBits += m.TransportSettings.GetLengthInBits(ctx)

	// Simple field (messageSettings)
	lengthInBits += m.MessageSettings.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_DataSetWriterDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DataSetWriterDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__dataSetWriterDataType DataSetWriterDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DataSetWriterDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DataSetWriterDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	name, err := ReadSimpleField[PascalString](ctx, "name", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'name' field"))
	}
	m.Name = name

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	enabled, err := ReadSimpleField(ctx, "enabled", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'enabled' field"))
	}
	m.Enabled = enabled

	dataSetWriterId, err := ReadSimpleField(ctx, "dataSetWriterId", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataSetWriterId' field"))
	}
	m.DataSetWriterId = dataSetWriterId

	dataSetFieldContentMask, err := ReadEnumField[DataSetFieldContentMask](ctx, "dataSetFieldContentMask", "DataSetFieldContentMask", ReadEnum(DataSetFieldContentMaskByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataSetFieldContentMask' field"))
	}
	m.DataSetFieldContentMask = dataSetFieldContentMask

	keyFrameCount, err := ReadSimpleField(ctx, "keyFrameCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'keyFrameCount' field"))
	}
	m.KeyFrameCount = keyFrameCount

	dataSetName, err := ReadSimpleField[PascalString](ctx, "dataSetName", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataSetName' field"))
	}
	m.DataSetName = dataSetName

	noOfDataSetWriterProperties, err := ReadImplicitField[int32](ctx, "noOfDataSetWriterProperties", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfDataSetWriterProperties' field"))
	}
	_ = noOfDataSetWriterProperties

	dataSetWriterProperties, err := ReadCountArrayField[KeyValuePair](ctx, "dataSetWriterProperties", ReadComplex[KeyValuePair](ExtensionObjectDefinitionParseWithBufferProducer[KeyValuePair]((int32)(int32(14535))), readBuffer), uint64(noOfDataSetWriterProperties))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataSetWriterProperties' field"))
	}
	m.DataSetWriterProperties = dataSetWriterProperties

	transportSettings, err := ReadSimpleField[ExtensionObject](ctx, "transportSettings", ReadComplex[ExtensionObject](ExtensionObjectParseWithBufferProducer[ExtensionObject]((bool)(bool(true))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transportSettings' field"))
	}
	m.TransportSettings = transportSettings

	messageSettings, err := ReadSimpleField[ExtensionObject](ctx, "messageSettings", ReadComplex[ExtensionObject](ExtensionObjectParseWithBufferProducer[ExtensionObject]((bool)(bool(true))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'messageSettings' field"))
	}
	m.MessageSettings = messageSettings

	if closeErr := readBuffer.CloseContext("DataSetWriterDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DataSetWriterDataType")
	}

	return m, nil
}

func (m *_DataSetWriterDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DataSetWriterDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DataSetWriterDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DataSetWriterDataType")
		}

		if err := WriteSimpleField[PascalString](ctx, "name", m.GetName(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'name' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "enabled", m.GetEnabled(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'enabled' field")
		}

		if err := WriteSimpleField[uint16](ctx, "dataSetWriterId", m.GetDataSetWriterId(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'dataSetWriterId' field")
		}

		if err := WriteSimpleEnumField[DataSetFieldContentMask](ctx, "dataSetFieldContentMask", "DataSetFieldContentMask", m.GetDataSetFieldContentMask(), WriteEnum[DataSetFieldContentMask, uint32](DataSetFieldContentMask.GetValue, DataSetFieldContentMask.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'dataSetFieldContentMask' field")
		}

		if err := WriteSimpleField[uint32](ctx, "keyFrameCount", m.GetKeyFrameCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'keyFrameCount' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "dataSetName", m.GetDataSetName(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dataSetName' field")
		}
		noOfDataSetWriterProperties := int32(utils.InlineIf(bool((m.GetDataSetWriterProperties()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetDataSetWriterProperties()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfDataSetWriterProperties", noOfDataSetWriterProperties, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfDataSetWriterProperties' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "dataSetWriterProperties", m.GetDataSetWriterProperties(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'dataSetWriterProperties' field")
		}

		if err := WriteSimpleField[ExtensionObject](ctx, "transportSettings", m.GetTransportSettings(), WriteComplex[ExtensionObject](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'transportSettings' field")
		}

		if err := WriteSimpleField[ExtensionObject](ctx, "messageSettings", m.GetMessageSettings(), WriteComplex[ExtensionObject](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'messageSettings' field")
		}

		if popErr := writeBuffer.PopContext("DataSetWriterDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DataSetWriterDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DataSetWriterDataType) IsDataSetWriterDataType() {}

func (m *_DataSetWriterDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DataSetWriterDataType) deepCopy() *_DataSetWriterDataType {
	if m == nil {
		return nil
	}
	_DataSetWriterDataTypeCopy := &_DataSetWriterDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[PascalString](m.Name),
		m.Enabled,
		m.DataSetWriterId,
		m.DataSetFieldContentMask,
		m.KeyFrameCount,
		utils.DeepCopy[PascalString](m.DataSetName),
		utils.DeepCopySlice[KeyValuePair, KeyValuePair](m.DataSetWriterProperties),
		utils.DeepCopy[ExtensionObject](m.TransportSettings),
		utils.DeepCopy[ExtensionObject](m.MessageSettings),
		m.reservedField0,
	}
	_DataSetWriterDataTypeCopy.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _DataSetWriterDataTypeCopy
}

func (m *_DataSetWriterDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
