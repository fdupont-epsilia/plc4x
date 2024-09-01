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

// ComObjectTable is the corresponding interface of ComObjectTable
type ComObjectTable interface {
	ComObjectTableContract
	ComObjectTableRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// ComObjectTableContract provides a set of functions which can be overwritten by a sub struct
type ComObjectTableContract interface {
}

// ComObjectTableRequirements provides a set of functions which need to be implemented by a sub struct
type ComObjectTableRequirements interface {
	// GetFirmwareType returns FirmwareType (discriminator field)
	GetFirmwareType() FirmwareType
}

// ComObjectTableExactly can be used when we want exactly this type and not a type which fulfills ComObjectTable.
// This is useful for switch cases.
type ComObjectTableExactly interface {
	ComObjectTable
	isComObjectTable() bool
}

// _ComObjectTable is the data-structure of this message
type _ComObjectTable struct {
	_ComObjectTableChildRequirements
}

var _ ComObjectTableContract = (*_ComObjectTable)(nil)

type _ComObjectTableChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetFirmwareType() FirmwareType
}

type ComObjectTableParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ComObjectTable, serializeChildFunction func() error) error
	GetTypeName() string
}

type ComObjectTableChild interface {
	utils.Serializable
	InitializeParent(parent ComObjectTable)
	GetParent() *ComObjectTable

	GetTypeName() string
	ComObjectTable
}

// NewComObjectTable factory function for _ComObjectTable
func NewComObjectTable() *_ComObjectTable {
	return &_ComObjectTable{}
}

// Deprecated: use the interface for direct cast
func CastComObjectTable(structType any) ComObjectTable {
	if casted, ok := structType.(ComObjectTable); ok {
		return casted
	}
	if casted, ok := structType.(*ComObjectTable); ok {
		return *casted
	}
	return nil
}

func (m *_ComObjectTable) GetTypeName() string {
	return "ComObjectTable"
}

func (m *_ComObjectTable) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_ComObjectTable) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ComObjectTableParse(ctx context.Context, theBytes []byte, firmwareType FirmwareType) (ComObjectTable, error) {
	return ComObjectTableParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), firmwareType)
}

func ComObjectTableParseWithBufferProducer[T ComObjectTable](firmwareType FirmwareType) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		buffer, err := ComObjectTableParseWithBuffer(ctx, readBuffer, firmwareType)
		if err != nil {
			var zero T
			return zero, err
		}
		return buffer.(T), err
	}
}

func ComObjectTableParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, firmwareType FirmwareType) (ComObjectTable, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ComObjectTable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ComObjectTable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type ComObjectTableChildSerializeRequirement interface {
		ComObjectTable
		InitializeParent(ComObjectTable)
		GetParent() ComObjectTable
	}
	var _childTemp any
	var _child ComObjectTableChildSerializeRequirement
	var typeSwitchError error
	switch {
	case firmwareType == FirmwareType_SYSTEM_1: // ComObjectTableRealisationType1
		_childTemp, typeSwitchError = ComObjectTableRealisationType1ParseWithBuffer(ctx, readBuffer, firmwareType)
	case firmwareType == FirmwareType_SYSTEM_2: // ComObjectTableRealisationType2
		_childTemp, typeSwitchError = ComObjectTableRealisationType2ParseWithBuffer(ctx, readBuffer, firmwareType)
	case firmwareType == FirmwareType_SYSTEM_300: // ComObjectTableRealisationType6
		_childTemp, typeSwitchError = ComObjectTableRealisationType6ParseWithBuffer(ctx, readBuffer, firmwareType)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [firmwareType=%v]", firmwareType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of ComObjectTable")
	}
	_child = _childTemp.(ComObjectTableChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("ComObjectTable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ComObjectTable")
	}

	// Finish initializing
	_child.InitializeParent(_child)
	return _child, nil
}

func (pm *_ComObjectTable) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ComObjectTable, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ComObjectTable"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ComObjectTable")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ComObjectTable"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ComObjectTable")
	}
	return nil
}

func (m *_ComObjectTable) isComObjectTable() bool {
	return true
}

func (m *_ComObjectTable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
