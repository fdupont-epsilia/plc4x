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

// Structure is the corresponding interface of Structure
type Structure interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// StructureExactly can be used when we want exactly this type and not a type which fulfills Structure.
// This is useful for switch cases.
type StructureExactly interface {
	Structure
	isStructure() bool
}

// _Structure is the data-structure of this message
type _Structure struct {
}

var _ Structure = (*_Structure)(nil)

// NewStructure factory function for _Structure
func NewStructure() *_Structure {
	return &_Structure{}
}

// Deprecated: use the interface for direct cast
func CastStructure(structType any) Structure {
	if casted, ok := structType.(Structure); ok {
		return casted
	}
	if casted, ok := structType.(*Structure); ok {
		return *casted
	}
	return nil
}

func (m *_Structure) GetTypeName() string {
	return "Structure"
}

func (m *_Structure) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_Structure) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func StructureParse(ctx context.Context, theBytes []byte) (Structure, error) {
	return StructureParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func StructureParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (Structure, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (Structure, error) {
		return StructureParseWithBuffer(ctx, readBuffer)
	}
}

func StructureParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (Structure, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("Structure"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Structure")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("Structure"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Structure")
	}

	// Create the instance
	return &_Structure{}, nil
}

func (m *_Structure) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_Structure) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("Structure"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for Structure")
	}

	if popErr := writeBuffer.PopContext("Structure"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for Structure")
	}
	return nil
}

func (m *_Structure) isStructure() bool {
	return true
}

func (m *_Structure) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
