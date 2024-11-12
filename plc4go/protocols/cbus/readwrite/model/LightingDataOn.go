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

// LightingDataOn is the corresponding interface of LightingDataOn
type LightingDataOn interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	LightingData
	// GetGroup returns Group (property field)
	GetGroup() byte
	// IsLightingDataOn is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsLightingDataOn()
	// CreateBuilder creates a LightingDataOnBuilder
	CreateLightingDataOnBuilder() LightingDataOnBuilder
}

// _LightingDataOn is the data-structure of this message
type _LightingDataOn struct {
	LightingDataContract
	Group byte
}

var _ LightingDataOn = (*_LightingDataOn)(nil)
var _ LightingDataRequirements = (*_LightingDataOn)(nil)

// NewLightingDataOn factory function for _LightingDataOn
func NewLightingDataOn(commandTypeContainer LightingCommandTypeContainer, group byte) *_LightingDataOn {
	_result := &_LightingDataOn{
		LightingDataContract: NewLightingData(commandTypeContainer),
		Group:                group,
	}
	_result.LightingDataContract.(*_LightingData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// LightingDataOnBuilder is a builder for LightingDataOn
type LightingDataOnBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(group byte) LightingDataOnBuilder
	// WithGroup adds Group (property field)
	WithGroup(byte) LightingDataOnBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() LightingDataBuilder
	// Build builds the LightingDataOn or returns an error if something is wrong
	Build() (LightingDataOn, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() LightingDataOn
}

// NewLightingDataOnBuilder() creates a LightingDataOnBuilder
func NewLightingDataOnBuilder() LightingDataOnBuilder {
	return &_LightingDataOnBuilder{_LightingDataOn: new(_LightingDataOn)}
}

type _LightingDataOnBuilder struct {
	*_LightingDataOn

	parentBuilder *_LightingDataBuilder

	err *utils.MultiError
}

var _ (LightingDataOnBuilder) = (*_LightingDataOnBuilder)(nil)

func (b *_LightingDataOnBuilder) setParent(contract LightingDataContract) {
	b.LightingDataContract = contract
	contract.(*_LightingData)._SubType = b._LightingDataOn
}

func (b *_LightingDataOnBuilder) WithMandatoryFields(group byte) LightingDataOnBuilder {
	return b.WithGroup(group)
}

func (b *_LightingDataOnBuilder) WithGroup(group byte) LightingDataOnBuilder {
	b.Group = group
	return b
}

func (b *_LightingDataOnBuilder) Build() (LightingDataOn, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._LightingDataOn.deepCopy(), nil
}

func (b *_LightingDataOnBuilder) MustBuild() LightingDataOn {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_LightingDataOnBuilder) Done() LightingDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewLightingDataBuilder().(*_LightingDataBuilder)
	}
	return b.parentBuilder
}

func (b *_LightingDataOnBuilder) buildForLightingData() (LightingData, error) {
	return b.Build()
}

func (b *_LightingDataOnBuilder) DeepCopy() any {
	_copy := b.CreateLightingDataOnBuilder().(*_LightingDataOnBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateLightingDataOnBuilder creates a LightingDataOnBuilder
func (b *_LightingDataOn) CreateLightingDataOnBuilder() LightingDataOnBuilder {
	if b == nil {
		return NewLightingDataOnBuilder()
	}
	return &_LightingDataOnBuilder{_LightingDataOn: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LightingDataOn) GetParent() LightingDataContract {
	return m.LightingDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LightingDataOn) GetGroup() byte {
	return m.Group
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastLightingDataOn(structType any) LightingDataOn {
	if casted, ok := structType.(LightingDataOn); ok {
		return casted
	}
	if casted, ok := structType.(*LightingDataOn); ok {
		return *casted
	}
	return nil
}

func (m *_LightingDataOn) GetTypeName() string {
	return "LightingDataOn"
}

func (m *_LightingDataOn) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.LightingDataContract.(*_LightingData).getLengthInBits(ctx))

	// Simple field (group)
	lengthInBits += 8

	return lengthInBits
}

func (m *_LightingDataOn) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_LightingDataOn) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_LightingData) (__lightingDataOn LightingDataOn, err error) {
	m.LightingDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LightingDataOn"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LightingDataOn")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	group, err := ReadSimpleField(ctx, "group", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'group' field"))
	}
	m.Group = group

	if closeErr := readBuffer.CloseContext("LightingDataOn"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LightingDataOn")
	}

	return m, nil
}

func (m *_LightingDataOn) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LightingDataOn) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LightingDataOn"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LightingDataOn")
		}

		if err := WriteSimpleField[byte](ctx, "group", m.GetGroup(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'group' field")
		}

		if popErr := writeBuffer.PopContext("LightingDataOn"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LightingDataOn")
		}
		return nil
	}
	return m.LightingDataContract.(*_LightingData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_LightingDataOn) IsLightingDataOn() {}

func (m *_LightingDataOn) DeepCopy() any {
	return m.deepCopy()
}

func (m *_LightingDataOn) deepCopy() *_LightingDataOn {
	if m == nil {
		return nil
	}
	_LightingDataOnCopy := &_LightingDataOn{
		m.LightingDataContract.(*_LightingData).deepCopy(),
		m.Group,
	}
	_LightingDataOnCopy.LightingDataContract.(*_LightingData)._SubType = m
	return _LightingDataOnCopy
}

func (m *_LightingDataOn) String() string {
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
