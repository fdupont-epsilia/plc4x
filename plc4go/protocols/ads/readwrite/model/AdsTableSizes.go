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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// AdsTableSizes is the corresponding interface of AdsTableSizes
type AdsTableSizes interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetSymbolCount returns SymbolCount (property field)
	GetSymbolCount() uint32
	// GetSymbolLength returns SymbolLength (property field)
	GetSymbolLength() uint32
	// GetDataTypeCount returns DataTypeCount (property field)
	GetDataTypeCount() uint32
	// GetDataTypeLength returns DataTypeLength (property field)
	GetDataTypeLength() uint32
	// GetExtraCount returns ExtraCount (property field)
	GetExtraCount() uint32
	// GetExtraLength returns ExtraLength (property field)
	GetExtraLength() uint32
}

// AdsTableSizesExactly can be used when we want exactly this type and not a type which fulfills AdsTableSizes.
// This is useful for switch cases.
type AdsTableSizesExactly interface {
	AdsTableSizes
	isAdsTableSizes() bool
}

// _AdsTableSizes is the data-structure of this message
type _AdsTableSizes struct {
	SymbolCount    uint32
	SymbolLength   uint32
	DataTypeCount  uint32
	DataTypeLength uint32
	ExtraCount     uint32
	ExtraLength    uint32
}

var _ AdsTableSizes = (*_AdsTableSizes)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsTableSizes) GetSymbolCount() uint32 {
	return m.SymbolCount
}

func (m *_AdsTableSizes) GetSymbolLength() uint32 {
	return m.SymbolLength
}

func (m *_AdsTableSizes) GetDataTypeCount() uint32 {
	return m.DataTypeCount
}

func (m *_AdsTableSizes) GetDataTypeLength() uint32 {
	return m.DataTypeLength
}

func (m *_AdsTableSizes) GetExtraCount() uint32 {
	return m.ExtraCount
}

func (m *_AdsTableSizes) GetExtraLength() uint32 {
	return m.ExtraLength
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsTableSizes factory function for _AdsTableSizes
func NewAdsTableSizes(symbolCount uint32, symbolLength uint32, dataTypeCount uint32, dataTypeLength uint32, extraCount uint32, extraLength uint32) *_AdsTableSizes {
	return &_AdsTableSizes{SymbolCount: symbolCount, SymbolLength: symbolLength, DataTypeCount: dataTypeCount, DataTypeLength: dataTypeLength, ExtraCount: extraCount, ExtraLength: extraLength}
}

// Deprecated: use the interface for direct cast
func CastAdsTableSizes(structType any) AdsTableSizes {
	if casted, ok := structType.(AdsTableSizes); ok {
		return casted
	}
	if casted, ok := structType.(*AdsTableSizes); ok {
		return *casted
	}
	return nil
}

func (m *_AdsTableSizes) GetTypeName() string {
	return "AdsTableSizes"
}

func (m *_AdsTableSizes) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (symbolCount)
	lengthInBits += 32

	// Simple field (symbolLength)
	lengthInBits += 32

	// Simple field (dataTypeCount)
	lengthInBits += 32

	// Simple field (dataTypeLength)
	lengthInBits += 32

	// Simple field (extraCount)
	lengthInBits += 32

	// Simple field (extraLength)
	lengthInBits += 32

	return lengthInBits
}

func (m *_AdsTableSizes) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsTableSizesParse(ctx context.Context, theBytes []byte) (AdsTableSizes, error) {
	return AdsTableSizesParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.LittleEndian)))
}

func AdsTableSizesParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (AdsTableSizes, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (AdsTableSizes, error) {
		return AdsTableSizesParseWithBuffer(ctx, readBuffer)
	}
}

func AdsTableSizesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AdsTableSizes, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsTableSizes"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsTableSizes")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	symbolCount, err := ReadSimpleField(ctx, "symbolCount", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'symbolCount' field"))
	}

	symbolLength, err := ReadSimpleField(ctx, "symbolLength", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'symbolLength' field"))
	}

	dataTypeCount, err := ReadSimpleField(ctx, "dataTypeCount", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataTypeCount' field"))
	}

	dataTypeLength, err := ReadSimpleField(ctx, "dataTypeLength", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataTypeLength' field"))
	}

	extraCount, err := ReadSimpleField(ctx, "extraCount", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'extraCount' field"))
	}

	extraLength, err := ReadSimpleField(ctx, "extraLength", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'extraLength' field"))
	}

	if closeErr := readBuffer.CloseContext("AdsTableSizes"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsTableSizes")
	}

	// Create the instance
	return &_AdsTableSizes{
		SymbolCount:    symbolCount,
		SymbolLength:   symbolLength,
		DataTypeCount:  dataTypeCount,
		DataTypeLength: dataTypeLength,
		ExtraCount:     extraCount,
		ExtraLength:    extraLength,
	}, nil
}

func (m *_AdsTableSizes) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.LittleEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsTableSizes) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AdsTableSizes"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AdsTableSizes")
	}

	if err := WriteSimpleField[uint32](ctx, "symbolCount", m.GetSymbolCount(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'symbolCount' field")
	}

	if err := WriteSimpleField[uint32](ctx, "symbolLength", m.GetSymbolLength(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'symbolLength' field")
	}

	if err := WriteSimpleField[uint32](ctx, "dataTypeCount", m.GetDataTypeCount(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'dataTypeCount' field")
	}

	if err := WriteSimpleField[uint32](ctx, "dataTypeLength", m.GetDataTypeLength(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'dataTypeLength' field")
	}

	if err := WriteSimpleField[uint32](ctx, "extraCount", m.GetExtraCount(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'extraCount' field")
	}

	if err := WriteSimpleField[uint32](ctx, "extraLength", m.GetExtraLength(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'extraLength' field")
	}

	if popErr := writeBuffer.PopContext("AdsTableSizes"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AdsTableSizes")
	}
	return nil
}

func (m *_AdsTableSizes) isAdsTableSizes() bool {
	return true
}

func (m *_AdsTableSizes) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
