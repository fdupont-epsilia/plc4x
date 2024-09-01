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

// TriggerControlDataLabel is the corresponding interface of TriggerControlDataLabel
type TriggerControlDataLabel interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	TriggerControlData
	// GetTriggerControlOptions returns TriggerControlOptions (property field)
	GetTriggerControlOptions() TriggerControlLabelOptions
	// GetActionSelector returns ActionSelector (property field)
	GetActionSelector() byte
	// GetLanguage returns Language (property field)
	GetLanguage() *Language
	// GetData returns Data (property field)
	GetData() []byte
}

// TriggerControlDataLabelExactly can be used when we want exactly this type and not a type which fulfills TriggerControlDataLabel.
// This is useful for switch cases.
type TriggerControlDataLabelExactly interface {
	TriggerControlDataLabel
	isTriggerControlDataLabel() bool
}

// _TriggerControlDataLabel is the data-structure of this message
type _TriggerControlDataLabel struct {
	*_TriggerControlData
	TriggerControlOptions TriggerControlLabelOptions
	ActionSelector        byte
	Language              *Language
	Data                  []byte
}

var _ TriggerControlDataLabel = (*_TriggerControlDataLabel)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TriggerControlDataLabel) InitializeParent(parent TriggerControlData, commandTypeContainer TriggerControlCommandTypeContainer, triggerGroup byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.TriggerGroup = triggerGroup
}

func (m *_TriggerControlDataLabel) GetParent() TriggerControlData {
	return m._TriggerControlData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TriggerControlDataLabel) GetTriggerControlOptions() TriggerControlLabelOptions {
	return m.TriggerControlOptions
}

func (m *_TriggerControlDataLabel) GetActionSelector() byte {
	return m.ActionSelector
}

func (m *_TriggerControlDataLabel) GetLanguage() *Language {
	return m.Language
}

func (m *_TriggerControlDataLabel) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTriggerControlDataLabel factory function for _TriggerControlDataLabel
func NewTriggerControlDataLabel(triggerControlOptions TriggerControlLabelOptions, actionSelector byte, language *Language, data []byte, commandTypeContainer TriggerControlCommandTypeContainer, triggerGroup byte) *_TriggerControlDataLabel {
	_result := &_TriggerControlDataLabel{
		TriggerControlOptions: triggerControlOptions,
		ActionSelector:        actionSelector,
		Language:              language,
		Data:                  data,
		_TriggerControlData:   NewTriggerControlData(commandTypeContainer, triggerGroup),
	}
	_result._TriggerControlData._TriggerControlDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastTriggerControlDataLabel(structType any) TriggerControlDataLabel {
	if casted, ok := structType.(TriggerControlDataLabel); ok {
		return casted
	}
	if casted, ok := structType.(*TriggerControlDataLabel); ok {
		return *casted
	}
	return nil
}

func (m *_TriggerControlDataLabel) GetTypeName() string {
	return "TriggerControlDataLabel"
}

func (m *_TriggerControlDataLabel) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (triggerControlOptions)
	lengthInBits += m.TriggerControlOptions.GetLengthInBits(ctx)

	// Simple field (actionSelector)
	lengthInBits += 8

	// Optional Field (language)
	if m.Language != nil {
		lengthInBits += 8
	}

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_TriggerControlDataLabel) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TriggerControlDataLabelParse(ctx context.Context, theBytes []byte, commandTypeContainer TriggerControlCommandTypeContainer) (TriggerControlDataLabel, error) {
	return TriggerControlDataLabelParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), commandTypeContainer)
}

func TriggerControlDataLabelParseWithBufferProducer(commandTypeContainer TriggerControlCommandTypeContainer) func(ctx context.Context, readBuffer utils.ReadBuffer) (TriggerControlDataLabel, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (TriggerControlDataLabel, error) {
		return TriggerControlDataLabelParseWithBuffer(ctx, readBuffer, commandTypeContainer)
	}
}

func TriggerControlDataLabelParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, commandTypeContainer TriggerControlCommandTypeContainer) (TriggerControlDataLabel, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TriggerControlDataLabel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TriggerControlDataLabel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	triggerControlOptions, err := ReadSimpleField[TriggerControlLabelOptions](ctx, "triggerControlOptions", ReadComplex[TriggerControlLabelOptions](TriggerControlLabelOptionsParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'triggerControlOptions' field"))
	}

	actionSelector, err := ReadSimpleField(ctx, "actionSelector", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actionSelector' field"))
	}

	language, err := ReadOptionalField[Language](ctx, "language", ReadEnum(LanguageByValue, ReadUnsignedByte(readBuffer, uint8(8))), bool((triggerControlOptions.GetLabelType()) != (TriggerControlLabelType_LOAD_DYNAMIC_ICON)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'language' field"))
	}

	data, err := readBuffer.ReadByteArray("data", int((int32(commandTypeContainer.NumBytes()) - int32((utils.InlineIf((bool((triggerControlOptions.GetLabelType()) != (TriggerControlLabelType_LOAD_DYNAMIC_ICON))), func() any { return int32((int32(4))) }, func() any { return int32((int32(3))) }).(int32))))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}

	if closeErr := readBuffer.CloseContext("TriggerControlDataLabel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TriggerControlDataLabel")
	}

	// Create a partially initialized instance
	_child := &_TriggerControlDataLabel{
		_TriggerControlData:   &_TriggerControlData{},
		TriggerControlOptions: triggerControlOptions,
		ActionSelector:        actionSelector,
		Language:              language,
		Data:                  data,
	}
	_child._TriggerControlData._TriggerControlDataChildRequirements = _child
	return _child, nil
}

func (m *_TriggerControlDataLabel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TriggerControlDataLabel) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TriggerControlDataLabel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TriggerControlDataLabel")
		}

		if err := WriteSimpleField[TriggerControlLabelOptions](ctx, "triggerControlOptions", m.GetTriggerControlOptions(), WriteComplex[TriggerControlLabelOptions](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'triggerControlOptions' field")
		}

		if err := WriteSimpleField[byte](ctx, "actionSelector", m.GetActionSelector(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'actionSelector' field")
		}

		if err := WriteOptionalEnumField[Language](ctx, "language", "Language", m.GetLanguage(), WriteEnum[Language, uint8](Language.GetValue, Language.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8)), bool((m.GetTriggerControlOptions().GetLabelType()) != (TriggerControlLabelType_LOAD_DYNAMIC_ICON))); err != nil {
			return errors.Wrap(err, "Error serializing 'language' field")
		}

		if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("TriggerControlDataLabel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TriggerControlDataLabel")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TriggerControlDataLabel) isTriggerControlDataLabel() bool {
	return true
}

func (m *_TriggerControlDataLabel) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
