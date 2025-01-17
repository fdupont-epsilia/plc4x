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

// CIPAttributes is the corresponding interface of CIPAttributes
type CIPAttributes interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetClassId returns ClassId (property field)
	GetClassId() []uint16
	// GetNumberAvailable returns NumberAvailable (property field)
	GetNumberAvailable() *uint16
	// GetNumberActive returns NumberActive (property field)
	GetNumberActive() *uint16
	// GetData returns Data (property field)
	GetData() []byte
	// IsCIPAttributes is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCIPAttributes()
	// CreateBuilder creates a CIPAttributesBuilder
	CreateCIPAttributesBuilder() CIPAttributesBuilder
}

// _CIPAttributes is the data-structure of this message
type _CIPAttributes struct {
	ClassId         []uint16
	NumberAvailable *uint16
	NumberActive    *uint16
	Data            []byte

	// Arguments.
	PacketLength uint16
}

var _ CIPAttributes = (*_CIPAttributes)(nil)

// NewCIPAttributes factory function for _CIPAttributes
func NewCIPAttributes(classId []uint16, numberAvailable *uint16, numberActive *uint16, data []byte, packetLength uint16) *_CIPAttributes {
	return &_CIPAttributes{ClassId: classId, NumberAvailable: numberAvailable, NumberActive: numberActive, Data: data, PacketLength: packetLength}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CIPAttributesBuilder is a builder for CIPAttributes
type CIPAttributesBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(classId []uint16, data []byte) CIPAttributesBuilder
	// WithClassId adds ClassId (property field)
	WithClassId(...uint16) CIPAttributesBuilder
	// WithNumberAvailable adds NumberAvailable (property field)
	WithOptionalNumberAvailable(uint16) CIPAttributesBuilder
	// WithNumberActive adds NumberActive (property field)
	WithOptionalNumberActive(uint16) CIPAttributesBuilder
	// WithData adds Data (property field)
	WithData(...byte) CIPAttributesBuilder
	// WithArgPacketLength sets a parser argument
	WithArgPacketLength(uint16) CIPAttributesBuilder
	// Build builds the CIPAttributes or returns an error if something is wrong
	Build() (CIPAttributes, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CIPAttributes
}

// NewCIPAttributesBuilder() creates a CIPAttributesBuilder
func NewCIPAttributesBuilder() CIPAttributesBuilder {
	return &_CIPAttributesBuilder{_CIPAttributes: new(_CIPAttributes)}
}

type _CIPAttributesBuilder struct {
	*_CIPAttributes

	err *utils.MultiError
}

var _ (CIPAttributesBuilder) = (*_CIPAttributesBuilder)(nil)

func (b *_CIPAttributesBuilder) WithMandatoryFields(classId []uint16, data []byte) CIPAttributesBuilder {
	return b.WithClassId(classId...).WithData(data...)
}

func (b *_CIPAttributesBuilder) WithClassId(classId ...uint16) CIPAttributesBuilder {
	b.ClassId = classId
	return b
}

func (b *_CIPAttributesBuilder) WithOptionalNumberAvailable(numberAvailable uint16) CIPAttributesBuilder {
	b.NumberAvailable = &numberAvailable
	return b
}

func (b *_CIPAttributesBuilder) WithOptionalNumberActive(numberActive uint16) CIPAttributesBuilder {
	b.NumberActive = &numberActive
	return b
}

func (b *_CIPAttributesBuilder) WithData(data ...byte) CIPAttributesBuilder {
	b.Data = data
	return b
}

func (b *_CIPAttributesBuilder) WithArgPacketLength(packetLength uint16) CIPAttributesBuilder {
	b.PacketLength = packetLength
	return b
}

func (b *_CIPAttributesBuilder) Build() (CIPAttributes, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CIPAttributes.deepCopy(), nil
}

func (b *_CIPAttributesBuilder) MustBuild() CIPAttributes {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CIPAttributesBuilder) DeepCopy() any {
	_copy := b.CreateCIPAttributesBuilder().(*_CIPAttributesBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCIPAttributesBuilder creates a CIPAttributesBuilder
func (b *_CIPAttributes) CreateCIPAttributesBuilder() CIPAttributesBuilder {
	if b == nil {
		return NewCIPAttributesBuilder()
	}
	return &_CIPAttributesBuilder{_CIPAttributes: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CIPAttributes) GetClassId() []uint16 {
	return m.ClassId
}

func (m *_CIPAttributes) GetNumberAvailable() *uint16 {
	return m.NumberAvailable
}

func (m *_CIPAttributes) GetNumberActive() *uint16 {
	return m.NumberActive
}

func (m *_CIPAttributes) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCIPAttributes(structType any) CIPAttributes {
	if casted, ok := structType.(CIPAttributes); ok {
		return casted
	}
	if casted, ok := structType.(*CIPAttributes); ok {
		return *casted
	}
	return nil
}

func (m *_CIPAttributes) GetTypeName() string {
	return "CIPAttributes"
}

func (m *_CIPAttributes) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (numberOfClasses)
	lengthInBits += 16

	// Array field
	if len(m.ClassId) > 0 {
		lengthInBits += 16 * uint16(len(m.ClassId))
	}

	// Optional Field (numberAvailable)
	if m.NumberAvailable != nil {
		lengthInBits += 16
	}

	// Optional Field (numberActive)
	if m.NumberActive != nil {
		lengthInBits += 16
	}

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_CIPAttributes) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CIPAttributesParse(ctx context.Context, theBytes []byte, packetLength uint16) (CIPAttributes, error) {
	return CIPAttributesParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), packetLength)
}

func CIPAttributesParseWithBufferProducer(packetLength uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (CIPAttributes, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (CIPAttributes, error) {
		return CIPAttributesParseWithBuffer(ctx, readBuffer, packetLength)
	}
}

func CIPAttributesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, packetLength uint16) (CIPAttributes, error) {
	v, err := (&_CIPAttributes{PacketLength: packetLength}).parse(ctx, readBuffer, packetLength)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_CIPAttributes) parse(ctx context.Context, readBuffer utils.ReadBuffer, packetLength uint16) (__cIPAttributes CIPAttributes, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CIPAttributes"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CIPAttributes")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	numberOfClasses, err := ReadImplicitField[uint16](ctx, "numberOfClasses", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfClasses' field"))
	}
	_ = numberOfClasses

	classId, err := ReadCountArrayField[uint16](ctx, "classId", ReadUnsignedShort(readBuffer, uint8(16)), uint64(numberOfClasses))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'classId' field"))
	}
	m.ClassId = classId

	var numberAvailable *uint16
	numberAvailable, err = ReadOptionalField[uint16](ctx, "numberAvailable", ReadUnsignedShort(readBuffer, uint8(16)), bool((packetLength) >= (((numberOfClasses)*(2))+(4))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberAvailable' field"))
	}
	m.NumberAvailable = numberAvailable

	var numberActive *uint16
	numberActive, err = ReadOptionalField[uint16](ctx, "numberActive", ReadUnsignedShort(readBuffer, uint8(16)), bool((packetLength) >= (((numberOfClasses)*(2))+(6))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberActive' field"))
	}
	m.NumberActive = numberActive

	data, err := readBuffer.ReadByteArray("data", int(utils.InlineIf((bool((packetLength) > (((numberOfClasses)*(2))+(6)))), func() any {
		return int32(int32(packetLength) - int32((int32((int32(numberOfClasses) * int32(int32(2)))) + int32(int32(6)))))
	}, func() any { return int32(int32(0)) }).(int32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("CIPAttributes"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CIPAttributes")
	}

	return m, nil
}

func (m *_CIPAttributes) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CIPAttributes) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("CIPAttributes"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CIPAttributes")
	}
	numberOfClasses := uint16(uint16(len(m.GetClassId())))
	if err := WriteImplicitField(ctx, "numberOfClasses", numberOfClasses, WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'numberOfClasses' field")
	}

	if err := WriteSimpleTypeArrayField(ctx, "classId", m.GetClassId(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'classId' field")
	}

	if err := WriteOptionalField[uint16](ctx, "numberAvailable", m.GetNumberAvailable(), WriteUnsignedShort(writeBuffer, 16), true); err != nil {
		return errors.Wrap(err, "Error serializing 'numberAvailable' field")
	}

	if err := WriteOptionalField[uint16](ctx, "numberActive", m.GetNumberActive(), WriteUnsignedShort(writeBuffer, 16), true); err != nil {
		return errors.Wrap(err, "Error serializing 'numberActive' field")
	}

	if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'data' field")
	}

	if popErr := writeBuffer.PopContext("CIPAttributes"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CIPAttributes")
	}
	return nil
}

////
// Arguments Getter

func (m *_CIPAttributes) GetPacketLength() uint16 {
	return m.PacketLength
}

//
////

func (m *_CIPAttributes) IsCIPAttributes() {}

func (m *_CIPAttributes) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CIPAttributes) deepCopy() *_CIPAttributes {
	if m == nil {
		return nil
	}
	_CIPAttributesCopy := &_CIPAttributes{
		utils.DeepCopySlice[uint16, uint16](m.ClassId),
		utils.CopyPtr[uint16](m.NumberAvailable),
		utils.CopyPtr[uint16](m.NumberActive),
		utils.DeepCopySlice[byte, byte](m.Data),
		m.PacketLength,
	}
	return _CIPAttributesCopy
}

func (m *_CIPAttributes) String() string {
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
