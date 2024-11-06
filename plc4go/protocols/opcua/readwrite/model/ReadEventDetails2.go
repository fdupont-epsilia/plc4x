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

// ReadEventDetails2 is the corresponding interface of ReadEventDetails2
type ReadEventDetails2 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetNumValuesPerNode returns NumValuesPerNode (property field)
	GetNumValuesPerNode() uint32
	// GetStartTime returns StartTime (property field)
	GetStartTime() int64
	// GetEndTime returns EndTime (property field)
	GetEndTime() int64
	// GetFilter returns Filter (property field)
	GetFilter() EventFilter
	// GetReadModified returns ReadModified (property field)
	GetReadModified() bool
	// IsReadEventDetails2 is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsReadEventDetails2()
	// CreateBuilder creates a ReadEventDetails2Builder
	CreateReadEventDetails2Builder() ReadEventDetails2Builder
}

// _ReadEventDetails2 is the data-structure of this message
type _ReadEventDetails2 struct {
	ExtensionObjectDefinitionContract
	NumValuesPerNode uint32
	StartTime        int64
	EndTime          int64
	Filter           EventFilter
	ReadModified     bool
	// Reserved Fields
	reservedField0 *uint8
}

var _ ReadEventDetails2 = (*_ReadEventDetails2)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ReadEventDetails2)(nil)

// NewReadEventDetails2 factory function for _ReadEventDetails2
func NewReadEventDetails2(numValuesPerNode uint32, startTime int64, endTime int64, filter EventFilter, readModified bool) *_ReadEventDetails2 {
	if filter == nil {
		panic("filter of type EventFilter for ReadEventDetails2 must not be nil")
	}
	_result := &_ReadEventDetails2{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		NumValuesPerNode:                  numValuesPerNode,
		StartTime:                         startTime,
		EndTime:                           endTime,
		Filter:                            filter,
		ReadModified:                      readModified,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ReadEventDetails2Builder is a builder for ReadEventDetails2
type ReadEventDetails2Builder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(numValuesPerNode uint32, startTime int64, endTime int64, filter EventFilter, readModified bool) ReadEventDetails2Builder
	// WithNumValuesPerNode adds NumValuesPerNode (property field)
	WithNumValuesPerNode(uint32) ReadEventDetails2Builder
	// WithStartTime adds StartTime (property field)
	WithStartTime(int64) ReadEventDetails2Builder
	// WithEndTime adds EndTime (property field)
	WithEndTime(int64) ReadEventDetails2Builder
	// WithFilter adds Filter (property field)
	WithFilter(EventFilter) ReadEventDetails2Builder
	// WithFilterBuilder adds Filter (property field) which is build by the builder
	WithFilterBuilder(func(EventFilterBuilder) EventFilterBuilder) ReadEventDetails2Builder
	// WithReadModified adds ReadModified (property field)
	WithReadModified(bool) ReadEventDetails2Builder
	// Build builds the ReadEventDetails2 or returns an error if something is wrong
	Build() (ReadEventDetails2, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ReadEventDetails2
}

// NewReadEventDetails2Builder() creates a ReadEventDetails2Builder
func NewReadEventDetails2Builder() ReadEventDetails2Builder {
	return &_ReadEventDetails2Builder{_ReadEventDetails2: new(_ReadEventDetails2)}
}

type _ReadEventDetails2Builder struct {
	*_ReadEventDetails2

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (ReadEventDetails2Builder) = (*_ReadEventDetails2Builder)(nil)

func (b *_ReadEventDetails2Builder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_ReadEventDetails2Builder) WithMandatoryFields(numValuesPerNode uint32, startTime int64, endTime int64, filter EventFilter, readModified bool) ReadEventDetails2Builder {
	return b.WithNumValuesPerNode(numValuesPerNode).WithStartTime(startTime).WithEndTime(endTime).WithFilter(filter).WithReadModified(readModified)
}

func (b *_ReadEventDetails2Builder) WithNumValuesPerNode(numValuesPerNode uint32) ReadEventDetails2Builder {
	b.NumValuesPerNode = numValuesPerNode
	return b
}

func (b *_ReadEventDetails2Builder) WithStartTime(startTime int64) ReadEventDetails2Builder {
	b.StartTime = startTime
	return b
}

func (b *_ReadEventDetails2Builder) WithEndTime(endTime int64) ReadEventDetails2Builder {
	b.EndTime = endTime
	return b
}

func (b *_ReadEventDetails2Builder) WithFilter(filter EventFilter) ReadEventDetails2Builder {
	b.Filter = filter
	return b
}

func (b *_ReadEventDetails2Builder) WithFilterBuilder(builderSupplier func(EventFilterBuilder) EventFilterBuilder) ReadEventDetails2Builder {
	builder := builderSupplier(b.Filter.CreateEventFilterBuilder())
	var err error
	b.Filter, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "EventFilterBuilder failed"))
	}
	return b
}

func (b *_ReadEventDetails2Builder) WithReadModified(readModified bool) ReadEventDetails2Builder {
	b.ReadModified = readModified
	return b
}

func (b *_ReadEventDetails2Builder) Build() (ReadEventDetails2, error) {
	if b.Filter == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'filter' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ReadEventDetails2.deepCopy(), nil
}

func (b *_ReadEventDetails2Builder) MustBuild() ReadEventDetails2 {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_ReadEventDetails2Builder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_ReadEventDetails2Builder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_ReadEventDetails2Builder) DeepCopy() any {
	_copy := b.CreateReadEventDetails2Builder().(*_ReadEventDetails2Builder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateReadEventDetails2Builder creates a ReadEventDetails2Builder
func (b *_ReadEventDetails2) CreateReadEventDetails2Builder() ReadEventDetails2Builder {
	if b == nil {
		return NewReadEventDetails2Builder()
	}
	return &_ReadEventDetails2Builder{_ReadEventDetails2: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ReadEventDetails2) GetExtensionId() int32 {
	return int32(32801)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ReadEventDetails2) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ReadEventDetails2) GetNumValuesPerNode() uint32 {
	return m.NumValuesPerNode
}

func (m *_ReadEventDetails2) GetStartTime() int64 {
	return m.StartTime
}

func (m *_ReadEventDetails2) GetEndTime() int64 {
	return m.EndTime
}

func (m *_ReadEventDetails2) GetFilter() EventFilter {
	return m.Filter
}

func (m *_ReadEventDetails2) GetReadModified() bool {
	return m.ReadModified
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastReadEventDetails2(structType any) ReadEventDetails2 {
	if casted, ok := structType.(ReadEventDetails2); ok {
		return casted
	}
	if casted, ok := structType.(*ReadEventDetails2); ok {
		return *casted
	}
	return nil
}

func (m *_ReadEventDetails2) GetTypeName() string {
	return "ReadEventDetails2"
}

func (m *_ReadEventDetails2) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (numValuesPerNode)
	lengthInBits += 32

	// Simple field (startTime)
	lengthInBits += 64

	// Simple field (endTime)
	lengthInBits += 64

	// Simple field (filter)
	lengthInBits += m.Filter.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (readModified)
	lengthInBits += 1

	return lengthInBits
}

func (m *_ReadEventDetails2) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ReadEventDetails2) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__readEventDetails2 ReadEventDetails2, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ReadEventDetails2"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ReadEventDetails2")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	numValuesPerNode, err := ReadSimpleField(ctx, "numValuesPerNode", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numValuesPerNode' field"))
	}
	m.NumValuesPerNode = numValuesPerNode

	startTime, err := ReadSimpleField(ctx, "startTime", ReadSignedLong(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'startTime' field"))
	}
	m.StartTime = startTime

	endTime, err := ReadSimpleField(ctx, "endTime", ReadSignedLong(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'endTime' field"))
	}
	m.EndTime = endTime

	filter, err := ReadSimpleField[EventFilter](ctx, "filter", ReadComplex[EventFilter](ExtensionObjectDefinitionParseWithBufferProducer[EventFilter]((int32)(int32(727))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'filter' field"))
	}
	m.Filter = filter

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	readModified, err := ReadSimpleField(ctx, "readModified", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'readModified' field"))
	}
	m.ReadModified = readModified

	if closeErr := readBuffer.CloseContext("ReadEventDetails2"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ReadEventDetails2")
	}

	return m, nil
}

func (m *_ReadEventDetails2) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ReadEventDetails2) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ReadEventDetails2"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ReadEventDetails2")
		}

		if err := WriteSimpleField[uint32](ctx, "numValuesPerNode", m.GetNumValuesPerNode(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'numValuesPerNode' field")
		}

		if err := WriteSimpleField[int64](ctx, "startTime", m.GetStartTime(), WriteSignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'startTime' field")
		}

		if err := WriteSimpleField[int64](ctx, "endTime", m.GetEndTime(), WriteSignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'endTime' field")
		}

		if err := WriteSimpleField[EventFilter](ctx, "filter", m.GetFilter(), WriteComplex[EventFilter](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'filter' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "readModified", m.GetReadModified(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'readModified' field")
		}

		if popErr := writeBuffer.PopContext("ReadEventDetails2"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ReadEventDetails2")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ReadEventDetails2) IsReadEventDetails2() {}

func (m *_ReadEventDetails2) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ReadEventDetails2) deepCopy() *_ReadEventDetails2 {
	if m == nil {
		return nil
	}
	_ReadEventDetails2Copy := &_ReadEventDetails2{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.NumValuesPerNode,
		m.StartTime,
		m.EndTime,
		m.Filter.DeepCopy().(EventFilter),
		m.ReadModified,
		m.reservedField0,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _ReadEventDetails2Copy
}

func (m *_ReadEventDetails2) String() string {
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