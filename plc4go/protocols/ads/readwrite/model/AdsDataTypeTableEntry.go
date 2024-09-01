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
	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const AdsDataTypeTableEntry_DATATYPENAMETERMINATOR uint8 = 0x00
const AdsDataTypeTableEntry_SIMPLETYPENAMETERMINATOR uint8 = 0x00
const AdsDataTypeTableEntry_COMMENTTERMINATOR uint8 = 0x00

// AdsDataTypeTableEntry is the corresponding interface of AdsDataTypeTableEntry
type AdsDataTypeTableEntry interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetEntryLength returns EntryLength (property field)
	GetEntryLength() uint32
	// GetVersion returns Version (property field)
	GetVersion() uint32
	// GetHashValue returns HashValue (property field)
	GetHashValue() uint32
	// GetTypeHashValue returns TypeHashValue (property field)
	GetTypeHashValue() uint32
	// GetSize returns Size (property field)
	GetSize() uint32
	// GetOffset returns Offset (property field)
	GetOffset() uint32
	// GetDataType returns DataType (property field)
	GetDataType() uint32
	// GetFlags returns Flags (property field)
	GetFlags() uint32
	// GetArrayDimensions returns ArrayDimensions (property field)
	GetArrayDimensions() uint16
	// GetNumChildren returns NumChildren (property field)
	GetNumChildren() uint16
	// GetDataTypeName returns DataTypeName (property field)
	GetDataTypeName() string
	// GetSimpleTypeName returns SimpleTypeName (property field)
	GetSimpleTypeName() string
	// GetComment returns Comment (property field)
	GetComment() string
	// GetArrayInfo returns ArrayInfo (property field)
	GetArrayInfo() []AdsDataTypeArrayInfo
	// GetChildren returns Children (property field)
	GetChildren() []AdsDataTypeTableChildEntry
	// GetRest returns Rest (property field)
	GetRest() []byte
}

// AdsDataTypeTableEntryExactly can be used when we want exactly this type and not a type which fulfills AdsDataTypeTableEntry.
// This is useful for switch cases.
type AdsDataTypeTableEntryExactly interface {
	AdsDataTypeTableEntry
	isAdsDataTypeTableEntry() bool
}

// _AdsDataTypeTableEntry is the data-structure of this message
type _AdsDataTypeTableEntry struct {
	EntryLength     uint32
	Version         uint32
	HashValue       uint32
	TypeHashValue   uint32
	Size            uint32
	Offset          uint32
	DataType        uint32
	Flags           uint32
	ArrayDimensions uint16
	NumChildren     uint16
	DataTypeName    string
	SimpleTypeName  string
	Comment         string
	ArrayInfo       []AdsDataTypeArrayInfo
	Children        []AdsDataTypeTableChildEntry
	Rest            []byte
}

var _ AdsDataTypeTableEntry = (*_AdsDataTypeTableEntry)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsDataTypeTableEntry) GetEntryLength() uint32 {
	return m.EntryLength
}

func (m *_AdsDataTypeTableEntry) GetVersion() uint32 {
	return m.Version
}

func (m *_AdsDataTypeTableEntry) GetHashValue() uint32 {
	return m.HashValue
}

func (m *_AdsDataTypeTableEntry) GetTypeHashValue() uint32 {
	return m.TypeHashValue
}

func (m *_AdsDataTypeTableEntry) GetSize() uint32 {
	return m.Size
}

func (m *_AdsDataTypeTableEntry) GetOffset() uint32 {
	return m.Offset
}

func (m *_AdsDataTypeTableEntry) GetDataType() uint32 {
	return m.DataType
}

func (m *_AdsDataTypeTableEntry) GetFlags() uint32 {
	return m.Flags
}

func (m *_AdsDataTypeTableEntry) GetArrayDimensions() uint16 {
	return m.ArrayDimensions
}

func (m *_AdsDataTypeTableEntry) GetNumChildren() uint16 {
	return m.NumChildren
}

func (m *_AdsDataTypeTableEntry) GetDataTypeName() string {
	return m.DataTypeName
}

func (m *_AdsDataTypeTableEntry) GetSimpleTypeName() string {
	return m.SimpleTypeName
}

func (m *_AdsDataTypeTableEntry) GetComment() string {
	return m.Comment
}

func (m *_AdsDataTypeTableEntry) GetArrayInfo() []AdsDataTypeArrayInfo {
	return m.ArrayInfo
}

func (m *_AdsDataTypeTableEntry) GetChildren() []AdsDataTypeTableChildEntry {
	return m.Children
}

func (m *_AdsDataTypeTableEntry) GetRest() []byte {
	return m.Rest
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_AdsDataTypeTableEntry) GetDataTypeNameTerminator() uint8 {
	return AdsDataTypeTableEntry_DATATYPENAMETERMINATOR
}

func (m *_AdsDataTypeTableEntry) GetSimpleTypeNameTerminator() uint8 {
	return AdsDataTypeTableEntry_SIMPLETYPENAMETERMINATOR
}

func (m *_AdsDataTypeTableEntry) GetCommentTerminator() uint8 {
	return AdsDataTypeTableEntry_COMMENTTERMINATOR
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsDataTypeTableEntry factory function for _AdsDataTypeTableEntry
func NewAdsDataTypeTableEntry(entryLength uint32, version uint32, hashValue uint32, typeHashValue uint32, size uint32, offset uint32, dataType uint32, flags uint32, arrayDimensions uint16, numChildren uint16, dataTypeName string, simpleTypeName string, comment string, arrayInfo []AdsDataTypeArrayInfo, children []AdsDataTypeTableChildEntry, rest []byte) *_AdsDataTypeTableEntry {
	return &_AdsDataTypeTableEntry{EntryLength: entryLength, Version: version, HashValue: hashValue, TypeHashValue: typeHashValue, Size: size, Offset: offset, DataType: dataType, Flags: flags, ArrayDimensions: arrayDimensions, NumChildren: numChildren, DataTypeName: dataTypeName, SimpleTypeName: simpleTypeName, Comment: comment, ArrayInfo: arrayInfo, Children: children, Rest: rest}
}

// Deprecated: use the interface for direct cast
func CastAdsDataTypeTableEntry(structType any) AdsDataTypeTableEntry {
	if casted, ok := structType.(AdsDataTypeTableEntry); ok {
		return casted
	}
	if casted, ok := structType.(*AdsDataTypeTableEntry); ok {
		return *casted
	}
	return nil
}

func (m *_AdsDataTypeTableEntry) GetTypeName() string {
	return "AdsDataTypeTableEntry"
}

func (m *_AdsDataTypeTableEntry) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (entryLength)
	lengthInBits += 32

	// Simple field (version)
	lengthInBits += 32

	// Simple field (hashValue)
	lengthInBits += 32

	// Simple field (typeHashValue)
	lengthInBits += 32

	// Simple field (size)
	lengthInBits += 32

	// Simple field (offset)
	lengthInBits += 32

	// Simple field (dataType)
	lengthInBits += 32

	// Simple field (flags)
	lengthInBits += 32

	// Implicit Field (dataTypeNameLength)
	lengthInBits += 16

	// Implicit Field (simpleTypeNameLength)
	lengthInBits += 16

	// Implicit Field (commentLength)
	lengthInBits += 16

	// Simple field (arrayDimensions)
	lengthInBits += 16

	// Simple field (numChildren)
	lengthInBits += 16

	// Simple field (dataTypeName)
	lengthInBits += uint16(int32(uint16(len(m.GetDataTypeName()))) * int32(int32(8)))

	// Const Field (dataTypeNameTerminator)
	lengthInBits += 8

	// Simple field (simpleTypeName)
	lengthInBits += uint16(int32(uint16(len(m.GetSimpleTypeName()))) * int32(int32(8)))

	// Const Field (simpleTypeNameTerminator)
	lengthInBits += 8

	// Simple field (comment)
	lengthInBits += uint16(int32(uint16(len(m.GetComment()))) * int32(int32(8)))

	// Const Field (commentTerminator)
	lengthInBits += 8

	// Array field
	if len(m.ArrayInfo) > 0 {
		for _curItem, element := range m.ArrayInfo {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ArrayInfo), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Array field
	if len(m.Children) > 0 {
		for _curItem, element := range m.Children {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Children), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Array field
	if len(m.Rest) > 0 {
		lengthInBits += 8 * uint16(len(m.Rest))
	}

	return lengthInBits
}

func (m *_AdsDataTypeTableEntry) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsDataTypeTableEntryParse(ctx context.Context, theBytes []byte) (AdsDataTypeTableEntry, error) {
	return AdsDataTypeTableEntryParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.LittleEndian)))
}

func AdsDataTypeTableEntryParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (AdsDataTypeTableEntry, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (AdsDataTypeTableEntry, error) {
		return AdsDataTypeTableEntryParseWithBuffer(ctx, readBuffer)
	}
}

func AdsDataTypeTableEntryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AdsDataTypeTableEntry, error) {
	v, err := (&_AdsDataTypeTableEntry{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_AdsDataTypeTableEntry) parse(ctx context.Context, readBuffer utils.ReadBuffer) (_adsDataTypeTableEntry AdsDataTypeTableEntry, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsDataTypeTableEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsDataTypeTableEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	var startPos = positionAware.GetPos()
	_ = startPos

	entryLength, err := ReadSimpleField(ctx, "entryLength", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'entryLength' field"))
	}
	m.EntryLength = entryLength

	version, err := ReadSimpleField(ctx, "version", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'version' field"))
	}
	m.Version = version

	hashValue, err := ReadSimpleField(ctx, "hashValue", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hashValue' field"))
	}
	m.HashValue = hashValue

	typeHashValue, err := ReadSimpleField(ctx, "typeHashValue", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'typeHashValue' field"))
	}
	m.TypeHashValue = typeHashValue

	size, err := ReadSimpleField(ctx, "size", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'size' field"))
	}
	m.Size = size

	offset, err := ReadSimpleField(ctx, "offset", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'offset' field"))
	}
	m.Offset = offset

	dataType, err := ReadSimpleField(ctx, "dataType", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataType' field"))
	}
	m.DataType = dataType

	flags, err := ReadSimpleField(ctx, "flags", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'flags' field"))
	}
	m.Flags = flags

	dataTypeNameLength, err := ReadImplicitField[uint16](ctx, "dataTypeNameLength", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataTypeNameLength' field"))
	}
	_ = dataTypeNameLength

	simpleTypeNameLength, err := ReadImplicitField[uint16](ctx, "simpleTypeNameLength", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'simpleTypeNameLength' field"))
	}
	_ = simpleTypeNameLength

	commentLength, err := ReadImplicitField[uint16](ctx, "commentLength", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commentLength' field"))
	}
	_ = commentLength

	arrayDimensions, err := ReadSimpleField(ctx, "arrayDimensions", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayDimensions' field"))
	}
	m.ArrayDimensions = arrayDimensions

	numChildren, err := ReadSimpleField(ctx, "numChildren", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numChildren' field"))
	}
	m.NumChildren = numChildren

	dataTypeName, err := ReadSimpleField(ctx, "dataTypeName", ReadString(readBuffer, uint32(int32(dataTypeNameLength)*int32(int32(8)))), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataTypeName' field"))
	}
	m.DataTypeName = dataTypeName

	dataTypeNameTerminator, err := ReadConstField[uint8](ctx, "dataTypeNameTerminator", ReadUnsignedByte(readBuffer, uint8(8)), AdsDataTypeTableEntry_DATATYPENAMETERMINATOR, codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataTypeNameTerminator' field"))
	}
	_ = dataTypeNameTerminator

	simpleTypeName, err := ReadSimpleField(ctx, "simpleTypeName", ReadString(readBuffer, uint32(int32(simpleTypeNameLength)*int32(int32(8)))), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'simpleTypeName' field"))
	}
	m.SimpleTypeName = simpleTypeName

	simpleTypeNameTerminator, err := ReadConstField[uint8](ctx, "simpleTypeNameTerminator", ReadUnsignedByte(readBuffer, uint8(8)), AdsDataTypeTableEntry_SIMPLETYPENAMETERMINATOR, codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'simpleTypeNameTerminator' field"))
	}
	_ = simpleTypeNameTerminator

	comment, err := ReadSimpleField(ctx, "comment", ReadString(readBuffer, uint32(int32(commentLength)*int32(int32(8)))), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'comment' field"))
	}
	m.Comment = comment

	commentTerminator, err := ReadConstField[uint8](ctx, "commentTerminator", ReadUnsignedByte(readBuffer, uint8(8)), AdsDataTypeTableEntry_COMMENTTERMINATOR, codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commentTerminator' field"))
	}
	_ = commentTerminator

	arrayInfo, err := ReadCountArrayField[AdsDataTypeArrayInfo](ctx, "arrayInfo", ReadComplex[AdsDataTypeArrayInfo](AdsDataTypeArrayInfoParseWithBuffer, readBuffer), uint64(arrayDimensions), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayInfo' field"))
	}
	m.ArrayInfo = arrayInfo

	children, err := ReadCountArrayField[AdsDataTypeTableChildEntry](ctx, "children", ReadComplex[AdsDataTypeTableChildEntry](AdsDataTypeTableChildEntryParseWithBuffer, readBuffer), uint64(numChildren), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'children' field"))
	}
	m.Children = children

	rest, err := readBuffer.ReadByteArray("rest", int(int32(entryLength)-int32((positionAware.GetPos()-startPos))), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'rest' field"))
	}
	m.Rest = rest

	if closeErr := readBuffer.CloseContext("AdsDataTypeTableEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsDataTypeTableEntry")
	}

	return m, nil
}

func (m *_AdsDataTypeTableEntry) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.LittleEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsDataTypeTableEntry) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AdsDataTypeTableEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AdsDataTypeTableEntry")
	}

	if err := WriteSimpleField[uint32](ctx, "entryLength", m.GetEntryLength(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'entryLength' field")
	}

	if err := WriteSimpleField[uint32](ctx, "version", m.GetVersion(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'version' field")
	}

	if err := WriteSimpleField[uint32](ctx, "hashValue", m.GetHashValue(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'hashValue' field")
	}

	if err := WriteSimpleField[uint32](ctx, "typeHashValue", m.GetTypeHashValue(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'typeHashValue' field")
	}

	if err := WriteSimpleField[uint32](ctx, "size", m.GetSize(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'size' field")
	}

	if err := WriteSimpleField[uint32](ctx, "offset", m.GetOffset(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'offset' field")
	}

	if err := WriteSimpleField[uint32](ctx, "dataType", m.GetDataType(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'dataType' field")
	}

	if err := WriteSimpleField[uint32](ctx, "flags", m.GetFlags(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'flags' field")
	}
	dataTypeNameLength := uint16(uint16(len(m.GetDataTypeName())))
	if err := WriteImplicitField(ctx, "dataTypeNameLength", dataTypeNameLength, WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'dataTypeNameLength' field")
	}
	simpleTypeNameLength := uint16(uint16(len(m.GetSimpleTypeName())))
	if err := WriteImplicitField(ctx, "simpleTypeNameLength", simpleTypeNameLength, WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'simpleTypeNameLength' field")
	}
	commentLength := uint16(uint16(len(m.GetComment())))
	if err := WriteImplicitField(ctx, "commentLength", commentLength, WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'commentLength' field")
	}

	if err := WriteSimpleField[uint16](ctx, "arrayDimensions", m.GetArrayDimensions(), WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'arrayDimensions' field")
	}

	if err := WriteSimpleField[uint16](ctx, "numChildren", m.GetNumChildren(), WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'numChildren' field")
	}

	if err := WriteSimpleField[string](ctx, "dataTypeName", m.GetDataTypeName(), WriteString(writeBuffer, int32(int32(uint16(len(m.GetDataTypeName())))*int32(int32(8)))), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'dataTypeName' field")
	}

	if err := WriteConstField(ctx, "dataTypeNameTerminator", AdsDataTypeTableEntry_DATATYPENAMETERMINATOR, WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'dataTypeNameTerminator' field")
	}

	if err := WriteSimpleField[string](ctx, "simpleTypeName", m.GetSimpleTypeName(), WriteString(writeBuffer, int32(int32(uint16(len(m.GetSimpleTypeName())))*int32(int32(8)))), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'simpleTypeName' field")
	}

	if err := WriteConstField(ctx, "simpleTypeNameTerminator", AdsDataTypeTableEntry_SIMPLETYPENAMETERMINATOR, WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'simpleTypeNameTerminator' field")
	}

	if err := WriteSimpleField[string](ctx, "comment", m.GetComment(), WriteString(writeBuffer, int32(int32(uint16(len(m.GetComment())))*int32(int32(8)))), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'comment' field")
	}

	if err := WriteConstField(ctx, "commentTerminator", AdsDataTypeTableEntry_COMMENTTERMINATOR, WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'commentTerminator' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "arrayInfo", m.GetArrayInfo(), writeBuffer, codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'arrayInfo' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "children", m.GetChildren(), writeBuffer, codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'children' field")
	}

	if err := WriteByteArrayField(ctx, "rest", m.GetRest(), WriteByteArray(writeBuffer, 8), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'rest' field")
	}

	if popErr := writeBuffer.PopContext("AdsDataTypeTableEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AdsDataTypeTableEntry")
	}
	return nil
}

func (m *_AdsDataTypeTableEntry) isAdsDataTypeTableEntry() bool {
	return true
}

func (m *_AdsDataTypeTableEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
