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

// MeteringDataMeasureDrinkingWater is the corresponding interface of MeteringDataMeasureDrinkingWater
type MeteringDataMeasureDrinkingWater interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MeteringData
}

// MeteringDataMeasureDrinkingWaterExactly can be used when we want exactly this type and not a type which fulfills MeteringDataMeasureDrinkingWater.
// This is useful for switch cases.
type MeteringDataMeasureDrinkingWaterExactly interface {
	MeteringDataMeasureDrinkingWater
	isMeteringDataMeasureDrinkingWater() bool
}

// _MeteringDataMeasureDrinkingWater is the data-structure of this message
type _MeteringDataMeasureDrinkingWater struct {
	*_MeteringData
}

var _ MeteringDataMeasureDrinkingWater = (*_MeteringDataMeasureDrinkingWater)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MeteringDataMeasureDrinkingWater) InitializeParent(parent MeteringData, commandTypeContainer MeteringCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_MeteringDataMeasureDrinkingWater) GetParent() MeteringData {
	return m._MeteringData
}

// NewMeteringDataMeasureDrinkingWater factory function for _MeteringDataMeasureDrinkingWater
func NewMeteringDataMeasureDrinkingWater(commandTypeContainer MeteringCommandTypeContainer, argument byte) *_MeteringDataMeasureDrinkingWater {
	_result := &_MeteringDataMeasureDrinkingWater{
		_MeteringData: NewMeteringData(commandTypeContainer, argument),
	}
	_result._MeteringData._MeteringDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMeteringDataMeasureDrinkingWater(structType any) MeteringDataMeasureDrinkingWater {
	if casted, ok := structType.(MeteringDataMeasureDrinkingWater); ok {
		return casted
	}
	if casted, ok := structType.(*MeteringDataMeasureDrinkingWater); ok {
		return *casted
	}
	return nil
}

func (m *_MeteringDataMeasureDrinkingWater) GetTypeName() string {
	return "MeteringDataMeasureDrinkingWater"
}

func (m *_MeteringDataMeasureDrinkingWater) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_MeteringDataMeasureDrinkingWater) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MeteringDataMeasureDrinkingWaterParse(ctx context.Context, theBytes []byte) (MeteringDataMeasureDrinkingWater, error) {
	return MeteringDataMeasureDrinkingWaterParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func MeteringDataMeasureDrinkingWaterParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (MeteringDataMeasureDrinkingWater, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (MeteringDataMeasureDrinkingWater, error) {
		return MeteringDataMeasureDrinkingWaterParseWithBuffer(ctx, readBuffer)
	}
}

func MeteringDataMeasureDrinkingWaterParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (MeteringDataMeasureDrinkingWater, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MeteringDataMeasureDrinkingWater"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MeteringDataMeasureDrinkingWater")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MeteringDataMeasureDrinkingWater"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MeteringDataMeasureDrinkingWater")
	}

	// Create a partially initialized instance
	_child := &_MeteringDataMeasureDrinkingWater{
		_MeteringData: &_MeteringData{},
	}
	_child._MeteringData._MeteringDataChildRequirements = _child
	return _child, nil
}

func (m *_MeteringDataMeasureDrinkingWater) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MeteringDataMeasureDrinkingWater) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MeteringDataMeasureDrinkingWater"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MeteringDataMeasureDrinkingWater")
		}

		if popErr := writeBuffer.PopContext("MeteringDataMeasureDrinkingWater"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MeteringDataMeasureDrinkingWater")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MeteringDataMeasureDrinkingWater) isMeteringDataMeasureDrinkingWater() bool {
	return true
}

func (m *_MeteringDataMeasureDrinkingWater) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
