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

// BVLCDeleteForeignDeviceTableEntry is the corresponding interface of BVLCDeleteForeignDeviceTableEntry
type BVLCDeleteForeignDeviceTableEntry interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BVLC
	// GetIp returns Ip (property field)
	GetIp() []uint8
	// GetPort returns Port (property field)
	GetPort() uint16
	// IsBVLCDeleteForeignDeviceTableEntry is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBVLCDeleteForeignDeviceTableEntry()
	// CreateBuilder creates a BVLCDeleteForeignDeviceTableEntryBuilder
	CreateBVLCDeleteForeignDeviceTableEntryBuilder() BVLCDeleteForeignDeviceTableEntryBuilder
}

// _BVLCDeleteForeignDeviceTableEntry is the data-structure of this message
type _BVLCDeleteForeignDeviceTableEntry struct {
	BVLCContract
	Ip   []uint8
	Port uint16
}

var _ BVLCDeleteForeignDeviceTableEntry = (*_BVLCDeleteForeignDeviceTableEntry)(nil)
var _ BVLCRequirements = (*_BVLCDeleteForeignDeviceTableEntry)(nil)

// NewBVLCDeleteForeignDeviceTableEntry factory function for _BVLCDeleteForeignDeviceTableEntry
func NewBVLCDeleteForeignDeviceTableEntry(ip []uint8, port uint16) *_BVLCDeleteForeignDeviceTableEntry {
	_result := &_BVLCDeleteForeignDeviceTableEntry{
		BVLCContract: NewBVLC(),
		Ip:           ip,
		Port:         port,
	}
	_result.BVLCContract.(*_BVLC)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BVLCDeleteForeignDeviceTableEntryBuilder is a builder for BVLCDeleteForeignDeviceTableEntry
type BVLCDeleteForeignDeviceTableEntryBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(ip []uint8, port uint16) BVLCDeleteForeignDeviceTableEntryBuilder
	// WithIp adds Ip (property field)
	WithIp(...uint8) BVLCDeleteForeignDeviceTableEntryBuilder
	// WithPort adds Port (property field)
	WithPort(uint16) BVLCDeleteForeignDeviceTableEntryBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BVLCBuilder
	// Build builds the BVLCDeleteForeignDeviceTableEntry or returns an error if something is wrong
	Build() (BVLCDeleteForeignDeviceTableEntry, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BVLCDeleteForeignDeviceTableEntry
}

// NewBVLCDeleteForeignDeviceTableEntryBuilder() creates a BVLCDeleteForeignDeviceTableEntryBuilder
func NewBVLCDeleteForeignDeviceTableEntryBuilder() BVLCDeleteForeignDeviceTableEntryBuilder {
	return &_BVLCDeleteForeignDeviceTableEntryBuilder{_BVLCDeleteForeignDeviceTableEntry: new(_BVLCDeleteForeignDeviceTableEntry)}
}

type _BVLCDeleteForeignDeviceTableEntryBuilder struct {
	*_BVLCDeleteForeignDeviceTableEntry

	parentBuilder *_BVLCBuilder

	err *utils.MultiError
}

var _ (BVLCDeleteForeignDeviceTableEntryBuilder) = (*_BVLCDeleteForeignDeviceTableEntryBuilder)(nil)

func (b *_BVLCDeleteForeignDeviceTableEntryBuilder) setParent(contract BVLCContract) {
	b.BVLCContract = contract
	contract.(*_BVLC)._SubType = b._BVLCDeleteForeignDeviceTableEntry
}

func (b *_BVLCDeleteForeignDeviceTableEntryBuilder) WithMandatoryFields(ip []uint8, port uint16) BVLCDeleteForeignDeviceTableEntryBuilder {
	return b.WithIp(ip...).WithPort(port)
}

func (b *_BVLCDeleteForeignDeviceTableEntryBuilder) WithIp(ip ...uint8) BVLCDeleteForeignDeviceTableEntryBuilder {
	b.Ip = ip
	return b
}

func (b *_BVLCDeleteForeignDeviceTableEntryBuilder) WithPort(port uint16) BVLCDeleteForeignDeviceTableEntryBuilder {
	b.Port = port
	return b
}

func (b *_BVLCDeleteForeignDeviceTableEntryBuilder) Build() (BVLCDeleteForeignDeviceTableEntry, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BVLCDeleteForeignDeviceTableEntry.deepCopy(), nil
}

func (b *_BVLCDeleteForeignDeviceTableEntryBuilder) MustBuild() BVLCDeleteForeignDeviceTableEntry {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BVLCDeleteForeignDeviceTableEntryBuilder) Done() BVLCBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBVLCBuilder().(*_BVLCBuilder)
	}
	return b.parentBuilder
}

func (b *_BVLCDeleteForeignDeviceTableEntryBuilder) buildForBVLC() (BVLC, error) {
	return b.Build()
}

func (b *_BVLCDeleteForeignDeviceTableEntryBuilder) DeepCopy() any {
	_copy := b.CreateBVLCDeleteForeignDeviceTableEntryBuilder().(*_BVLCDeleteForeignDeviceTableEntryBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBVLCDeleteForeignDeviceTableEntryBuilder creates a BVLCDeleteForeignDeviceTableEntryBuilder
func (b *_BVLCDeleteForeignDeviceTableEntry) CreateBVLCDeleteForeignDeviceTableEntryBuilder() BVLCDeleteForeignDeviceTableEntryBuilder {
	if b == nil {
		return NewBVLCDeleteForeignDeviceTableEntryBuilder()
	}
	return &_BVLCDeleteForeignDeviceTableEntryBuilder{_BVLCDeleteForeignDeviceTableEntry: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BVLCDeleteForeignDeviceTableEntry) GetBvlcFunction() uint8 {
	return 0x08
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BVLCDeleteForeignDeviceTableEntry) GetParent() BVLCContract {
	return m.BVLCContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BVLCDeleteForeignDeviceTableEntry) GetIp() []uint8 {
	return m.Ip
}

func (m *_BVLCDeleteForeignDeviceTableEntry) GetPort() uint16 {
	return m.Port
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBVLCDeleteForeignDeviceTableEntry(structType any) BVLCDeleteForeignDeviceTableEntry {
	if casted, ok := structType.(BVLCDeleteForeignDeviceTableEntry); ok {
		return casted
	}
	if casted, ok := structType.(*BVLCDeleteForeignDeviceTableEntry); ok {
		return *casted
	}
	return nil
}

func (m *_BVLCDeleteForeignDeviceTableEntry) GetTypeName() string {
	return "BVLCDeleteForeignDeviceTableEntry"
}

func (m *_BVLCDeleteForeignDeviceTableEntry) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BVLCContract.(*_BVLC).getLengthInBits(ctx))

	// Array field
	if len(m.Ip) > 0 {
		lengthInBits += 8 * uint16(len(m.Ip))
	}

	// Simple field (port)
	lengthInBits += 16

	return lengthInBits
}

func (m *_BVLCDeleteForeignDeviceTableEntry) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BVLCDeleteForeignDeviceTableEntry) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BVLC) (__bVLCDeleteForeignDeviceTableEntry BVLCDeleteForeignDeviceTableEntry, err error) {
	m.BVLCContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BVLCDeleteForeignDeviceTableEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BVLCDeleteForeignDeviceTableEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	ip, err := ReadCountArrayField[uint8](ctx, "ip", ReadUnsignedByte(readBuffer, uint8(8)), uint64(int32(4)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ip' field"))
	}
	m.Ip = ip

	port, err := ReadSimpleField(ctx, "port", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'port' field"))
	}
	m.Port = port

	if closeErr := readBuffer.CloseContext("BVLCDeleteForeignDeviceTableEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BVLCDeleteForeignDeviceTableEntry")
	}

	return m, nil
}

func (m *_BVLCDeleteForeignDeviceTableEntry) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BVLCDeleteForeignDeviceTableEntry) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCDeleteForeignDeviceTableEntry"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BVLCDeleteForeignDeviceTableEntry")
		}

		if err := WriteSimpleTypeArrayField(ctx, "ip", m.GetIp(), WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'ip' field")
		}

		if err := WriteSimpleField[uint16](ctx, "port", m.GetPort(), WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'port' field")
		}

		if popErr := writeBuffer.PopContext("BVLCDeleteForeignDeviceTableEntry"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BVLCDeleteForeignDeviceTableEntry")
		}
		return nil
	}
	return m.BVLCContract.(*_BVLC).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BVLCDeleteForeignDeviceTableEntry) IsBVLCDeleteForeignDeviceTableEntry() {}

func (m *_BVLCDeleteForeignDeviceTableEntry) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BVLCDeleteForeignDeviceTableEntry) deepCopy() *_BVLCDeleteForeignDeviceTableEntry {
	if m == nil {
		return nil
	}
	_BVLCDeleteForeignDeviceTableEntryCopy := &_BVLCDeleteForeignDeviceTableEntry{
		m.BVLCContract.(*_BVLC).deepCopy(),
		utils.DeepCopySlice[uint8, uint8](m.Ip),
		m.Port,
	}
	_BVLCDeleteForeignDeviceTableEntryCopy.BVLCContract.(*_BVLC)._SubType = m
	return _BVLCDeleteForeignDeviceTableEntryCopy
}

func (m *_BVLCDeleteForeignDeviceTableEntry) String() string {
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
