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

// ParameterValueInterfaceOptions1PowerUpSettings is the corresponding interface of ParameterValueInterfaceOptions1PowerUpSettings
type ParameterValueInterfaceOptions1PowerUpSettings interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ParameterValue
	// GetValue returns Value (property field)
	GetValue() InterfaceOptions1PowerUpSettings
	// IsParameterValueInterfaceOptions1PowerUpSettings is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsParameterValueInterfaceOptions1PowerUpSettings()
	// CreateBuilder creates a ParameterValueInterfaceOptions1PowerUpSettingsBuilder
	CreateParameterValueInterfaceOptions1PowerUpSettingsBuilder() ParameterValueInterfaceOptions1PowerUpSettingsBuilder
}

// _ParameterValueInterfaceOptions1PowerUpSettings is the data-structure of this message
type _ParameterValueInterfaceOptions1PowerUpSettings struct {
	ParameterValueContract
	Value InterfaceOptions1PowerUpSettings
}

var _ ParameterValueInterfaceOptions1PowerUpSettings = (*_ParameterValueInterfaceOptions1PowerUpSettings)(nil)
var _ ParameterValueRequirements = (*_ParameterValueInterfaceOptions1PowerUpSettings)(nil)

// NewParameterValueInterfaceOptions1PowerUpSettings factory function for _ParameterValueInterfaceOptions1PowerUpSettings
func NewParameterValueInterfaceOptions1PowerUpSettings(value InterfaceOptions1PowerUpSettings, numBytes uint8) *_ParameterValueInterfaceOptions1PowerUpSettings {
	if value == nil {
		panic("value of type InterfaceOptions1PowerUpSettings for ParameterValueInterfaceOptions1PowerUpSettings must not be nil")
	}
	_result := &_ParameterValueInterfaceOptions1PowerUpSettings{
		ParameterValueContract: NewParameterValue(numBytes),
		Value:                  value,
	}
	_result.ParameterValueContract.(*_ParameterValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ParameterValueInterfaceOptions1PowerUpSettingsBuilder is a builder for ParameterValueInterfaceOptions1PowerUpSettings
type ParameterValueInterfaceOptions1PowerUpSettingsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(value InterfaceOptions1PowerUpSettings) ParameterValueInterfaceOptions1PowerUpSettingsBuilder
	// WithValue adds Value (property field)
	WithValue(InterfaceOptions1PowerUpSettings) ParameterValueInterfaceOptions1PowerUpSettingsBuilder
	// WithValueBuilder adds Value (property field) which is build by the builder
	WithValueBuilder(func(InterfaceOptions1PowerUpSettingsBuilder) InterfaceOptions1PowerUpSettingsBuilder) ParameterValueInterfaceOptions1PowerUpSettingsBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ParameterValueBuilder
	// Build builds the ParameterValueInterfaceOptions1PowerUpSettings or returns an error if something is wrong
	Build() (ParameterValueInterfaceOptions1PowerUpSettings, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ParameterValueInterfaceOptions1PowerUpSettings
}

// NewParameterValueInterfaceOptions1PowerUpSettingsBuilder() creates a ParameterValueInterfaceOptions1PowerUpSettingsBuilder
func NewParameterValueInterfaceOptions1PowerUpSettingsBuilder() ParameterValueInterfaceOptions1PowerUpSettingsBuilder {
	return &_ParameterValueInterfaceOptions1PowerUpSettingsBuilder{_ParameterValueInterfaceOptions1PowerUpSettings: new(_ParameterValueInterfaceOptions1PowerUpSettings)}
}

type _ParameterValueInterfaceOptions1PowerUpSettingsBuilder struct {
	*_ParameterValueInterfaceOptions1PowerUpSettings

	parentBuilder *_ParameterValueBuilder

	err *utils.MultiError
}

var _ (ParameterValueInterfaceOptions1PowerUpSettingsBuilder) = (*_ParameterValueInterfaceOptions1PowerUpSettingsBuilder)(nil)

func (b *_ParameterValueInterfaceOptions1PowerUpSettingsBuilder) setParent(contract ParameterValueContract) {
	b.ParameterValueContract = contract
	contract.(*_ParameterValue)._SubType = b._ParameterValueInterfaceOptions1PowerUpSettings
}

func (b *_ParameterValueInterfaceOptions1PowerUpSettingsBuilder) WithMandatoryFields(value InterfaceOptions1PowerUpSettings) ParameterValueInterfaceOptions1PowerUpSettingsBuilder {
	return b.WithValue(value)
}

func (b *_ParameterValueInterfaceOptions1PowerUpSettingsBuilder) WithValue(value InterfaceOptions1PowerUpSettings) ParameterValueInterfaceOptions1PowerUpSettingsBuilder {
	b.Value = value
	return b
}

func (b *_ParameterValueInterfaceOptions1PowerUpSettingsBuilder) WithValueBuilder(builderSupplier func(InterfaceOptions1PowerUpSettingsBuilder) InterfaceOptions1PowerUpSettingsBuilder) ParameterValueInterfaceOptions1PowerUpSettingsBuilder {
	builder := builderSupplier(b.Value.CreateInterfaceOptions1PowerUpSettingsBuilder())
	var err error
	b.Value, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "InterfaceOptions1PowerUpSettingsBuilder failed"))
	}
	return b
}

func (b *_ParameterValueInterfaceOptions1PowerUpSettingsBuilder) Build() (ParameterValueInterfaceOptions1PowerUpSettings, error) {
	if b.Value == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'value' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ParameterValueInterfaceOptions1PowerUpSettings.deepCopy(), nil
}

func (b *_ParameterValueInterfaceOptions1PowerUpSettingsBuilder) MustBuild() ParameterValueInterfaceOptions1PowerUpSettings {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ParameterValueInterfaceOptions1PowerUpSettingsBuilder) Done() ParameterValueBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewParameterValueBuilder().(*_ParameterValueBuilder)
	}
	return b.parentBuilder
}

func (b *_ParameterValueInterfaceOptions1PowerUpSettingsBuilder) buildForParameterValue() (ParameterValue, error) {
	return b.Build()
}

func (b *_ParameterValueInterfaceOptions1PowerUpSettingsBuilder) DeepCopy() any {
	_copy := b.CreateParameterValueInterfaceOptions1PowerUpSettingsBuilder().(*_ParameterValueInterfaceOptions1PowerUpSettingsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateParameterValueInterfaceOptions1PowerUpSettingsBuilder creates a ParameterValueInterfaceOptions1PowerUpSettingsBuilder
func (b *_ParameterValueInterfaceOptions1PowerUpSettings) CreateParameterValueInterfaceOptions1PowerUpSettingsBuilder() ParameterValueInterfaceOptions1PowerUpSettingsBuilder {
	if b == nil {
		return NewParameterValueInterfaceOptions1PowerUpSettingsBuilder()
	}
	return &_ParameterValueInterfaceOptions1PowerUpSettingsBuilder{_ParameterValueInterfaceOptions1PowerUpSettings: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ParameterValueInterfaceOptions1PowerUpSettings) GetParameterType() ParameterType {
	return ParameterType_INTERFACE_OPTIONS_1_POWER_UP_SETTINGS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ParameterValueInterfaceOptions1PowerUpSettings) GetParent() ParameterValueContract {
	return m.ParameterValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ParameterValueInterfaceOptions1PowerUpSettings) GetValue() InterfaceOptions1PowerUpSettings {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastParameterValueInterfaceOptions1PowerUpSettings(structType any) ParameterValueInterfaceOptions1PowerUpSettings {
	if casted, ok := structType.(ParameterValueInterfaceOptions1PowerUpSettings); ok {
		return casted
	}
	if casted, ok := structType.(*ParameterValueInterfaceOptions1PowerUpSettings); ok {
		return *casted
	}
	return nil
}

func (m *_ParameterValueInterfaceOptions1PowerUpSettings) GetTypeName() string {
	return "ParameterValueInterfaceOptions1PowerUpSettings"
}

func (m *_ParameterValueInterfaceOptions1PowerUpSettings) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ParameterValueContract.(*_ParameterValue).getLengthInBits(ctx))

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ParameterValueInterfaceOptions1PowerUpSettings) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ParameterValueInterfaceOptions1PowerUpSettings) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ParameterValue, parameterType ParameterType, numBytes uint8) (__parameterValueInterfaceOptions1PowerUpSettings ParameterValueInterfaceOptions1PowerUpSettings, err error) {
	m.ParameterValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ParameterValueInterfaceOptions1PowerUpSettings"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ParameterValueInterfaceOptions1PowerUpSettings")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((numBytes) >= (1))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "InterfaceOptions1PowerUpSettings has exactly one byte"})
	}

	value, err := ReadSimpleField[InterfaceOptions1PowerUpSettings](ctx, "value", ReadComplex[InterfaceOptions1PowerUpSettings](InterfaceOptions1PowerUpSettingsParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("ParameterValueInterfaceOptions1PowerUpSettings"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ParameterValueInterfaceOptions1PowerUpSettings")
	}

	return m, nil
}

func (m *_ParameterValueInterfaceOptions1PowerUpSettings) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ParameterValueInterfaceOptions1PowerUpSettings) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ParameterValueInterfaceOptions1PowerUpSettings"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ParameterValueInterfaceOptions1PowerUpSettings")
		}

		if err := WriteSimpleField[InterfaceOptions1PowerUpSettings](ctx, "value", m.GetValue(), WriteComplex[InterfaceOptions1PowerUpSettings](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("ParameterValueInterfaceOptions1PowerUpSettings"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ParameterValueInterfaceOptions1PowerUpSettings")
		}
		return nil
	}
	return m.ParameterValueContract.(*_ParameterValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ParameterValueInterfaceOptions1PowerUpSettings) IsParameterValueInterfaceOptions1PowerUpSettings() {
}

func (m *_ParameterValueInterfaceOptions1PowerUpSettings) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ParameterValueInterfaceOptions1PowerUpSettings) deepCopy() *_ParameterValueInterfaceOptions1PowerUpSettings {
	if m == nil {
		return nil
	}
	_ParameterValueInterfaceOptions1PowerUpSettingsCopy := &_ParameterValueInterfaceOptions1PowerUpSettings{
		m.ParameterValueContract.(*_ParameterValue).deepCopy(),
		utils.DeepCopy[InterfaceOptions1PowerUpSettings](m.Value),
	}
	_ParameterValueInterfaceOptions1PowerUpSettingsCopy.ParameterValueContract.(*_ParameterValue)._SubType = m
	return _ParameterValueInterfaceOptions1PowerUpSettingsCopy
}

func (m *_ParameterValueInterfaceOptions1PowerUpSettings) String() string {
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
