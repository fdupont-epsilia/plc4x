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

// DeviceConfigurationAck is the corresponding interface of DeviceConfigurationAck
type DeviceConfigurationAck interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	KnxNetIpMessage
	// GetDeviceConfigurationAckDataBlock returns DeviceConfigurationAckDataBlock (property field)
	GetDeviceConfigurationAckDataBlock() DeviceConfigurationAckDataBlock
	// IsDeviceConfigurationAck is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDeviceConfigurationAck()
	// CreateBuilder creates a DeviceConfigurationAckBuilder
	CreateDeviceConfigurationAckBuilder() DeviceConfigurationAckBuilder
}

// _DeviceConfigurationAck is the data-structure of this message
type _DeviceConfigurationAck struct {
	KnxNetIpMessageContract
	DeviceConfigurationAckDataBlock DeviceConfigurationAckDataBlock
}

var _ DeviceConfigurationAck = (*_DeviceConfigurationAck)(nil)
var _ KnxNetIpMessageRequirements = (*_DeviceConfigurationAck)(nil)

// NewDeviceConfigurationAck factory function for _DeviceConfigurationAck
func NewDeviceConfigurationAck(deviceConfigurationAckDataBlock DeviceConfigurationAckDataBlock) *_DeviceConfigurationAck {
	if deviceConfigurationAckDataBlock == nil {
		panic("deviceConfigurationAckDataBlock of type DeviceConfigurationAckDataBlock for DeviceConfigurationAck must not be nil")
	}
	_result := &_DeviceConfigurationAck{
		KnxNetIpMessageContract:         NewKnxNetIpMessage(),
		DeviceConfigurationAckDataBlock: deviceConfigurationAckDataBlock,
	}
	_result.KnxNetIpMessageContract.(*_KnxNetIpMessage)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DeviceConfigurationAckBuilder is a builder for DeviceConfigurationAck
type DeviceConfigurationAckBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(deviceConfigurationAckDataBlock DeviceConfigurationAckDataBlock) DeviceConfigurationAckBuilder
	// WithDeviceConfigurationAckDataBlock adds DeviceConfigurationAckDataBlock (property field)
	WithDeviceConfigurationAckDataBlock(DeviceConfigurationAckDataBlock) DeviceConfigurationAckBuilder
	// WithDeviceConfigurationAckDataBlockBuilder adds DeviceConfigurationAckDataBlock (property field) which is build by the builder
	WithDeviceConfigurationAckDataBlockBuilder(func(DeviceConfigurationAckDataBlockBuilder) DeviceConfigurationAckDataBlockBuilder) DeviceConfigurationAckBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() KnxNetIpMessageBuilder
	// Build builds the DeviceConfigurationAck or returns an error if something is wrong
	Build() (DeviceConfigurationAck, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DeviceConfigurationAck
}

// NewDeviceConfigurationAckBuilder() creates a DeviceConfigurationAckBuilder
func NewDeviceConfigurationAckBuilder() DeviceConfigurationAckBuilder {
	return &_DeviceConfigurationAckBuilder{_DeviceConfigurationAck: new(_DeviceConfigurationAck)}
}

type _DeviceConfigurationAckBuilder struct {
	*_DeviceConfigurationAck

	parentBuilder *_KnxNetIpMessageBuilder

	err *utils.MultiError
}

var _ (DeviceConfigurationAckBuilder) = (*_DeviceConfigurationAckBuilder)(nil)

func (b *_DeviceConfigurationAckBuilder) setParent(contract KnxNetIpMessageContract) {
	b.KnxNetIpMessageContract = contract
	contract.(*_KnxNetIpMessage)._SubType = b._DeviceConfigurationAck
}

func (b *_DeviceConfigurationAckBuilder) WithMandatoryFields(deviceConfigurationAckDataBlock DeviceConfigurationAckDataBlock) DeviceConfigurationAckBuilder {
	return b.WithDeviceConfigurationAckDataBlock(deviceConfigurationAckDataBlock)
}

func (b *_DeviceConfigurationAckBuilder) WithDeviceConfigurationAckDataBlock(deviceConfigurationAckDataBlock DeviceConfigurationAckDataBlock) DeviceConfigurationAckBuilder {
	b.DeviceConfigurationAckDataBlock = deviceConfigurationAckDataBlock
	return b
}

func (b *_DeviceConfigurationAckBuilder) WithDeviceConfigurationAckDataBlockBuilder(builderSupplier func(DeviceConfigurationAckDataBlockBuilder) DeviceConfigurationAckDataBlockBuilder) DeviceConfigurationAckBuilder {
	builder := builderSupplier(b.DeviceConfigurationAckDataBlock.CreateDeviceConfigurationAckDataBlockBuilder())
	var err error
	b.DeviceConfigurationAckDataBlock, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "DeviceConfigurationAckDataBlockBuilder failed"))
	}
	return b
}

func (b *_DeviceConfigurationAckBuilder) Build() (DeviceConfigurationAck, error) {
	if b.DeviceConfigurationAckDataBlock == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'deviceConfigurationAckDataBlock' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._DeviceConfigurationAck.deepCopy(), nil
}

func (b *_DeviceConfigurationAckBuilder) MustBuild() DeviceConfigurationAck {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_DeviceConfigurationAckBuilder) Done() KnxNetIpMessageBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewKnxNetIpMessageBuilder().(*_KnxNetIpMessageBuilder)
	}
	return b.parentBuilder
}

func (b *_DeviceConfigurationAckBuilder) buildForKnxNetIpMessage() (KnxNetIpMessage, error) {
	return b.Build()
}

func (b *_DeviceConfigurationAckBuilder) DeepCopy() any {
	_copy := b.CreateDeviceConfigurationAckBuilder().(*_DeviceConfigurationAckBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateDeviceConfigurationAckBuilder creates a DeviceConfigurationAckBuilder
func (b *_DeviceConfigurationAck) CreateDeviceConfigurationAckBuilder() DeviceConfigurationAckBuilder {
	if b == nil {
		return NewDeviceConfigurationAckBuilder()
	}
	return &_DeviceConfigurationAckBuilder{_DeviceConfigurationAck: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DeviceConfigurationAck) GetMsgType() uint16 {
	return 0x0311
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DeviceConfigurationAck) GetParent() KnxNetIpMessageContract {
	return m.KnxNetIpMessageContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeviceConfigurationAck) GetDeviceConfigurationAckDataBlock() DeviceConfigurationAckDataBlock {
	return m.DeviceConfigurationAckDataBlock
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDeviceConfigurationAck(structType any) DeviceConfigurationAck {
	if casted, ok := structType.(DeviceConfigurationAck); ok {
		return casted
	}
	if casted, ok := structType.(*DeviceConfigurationAck); ok {
		return *casted
	}
	return nil
}

func (m *_DeviceConfigurationAck) GetTypeName() string {
	return "DeviceConfigurationAck"
}

func (m *_DeviceConfigurationAck) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.KnxNetIpMessageContract.(*_KnxNetIpMessage).getLengthInBits(ctx))

	// Simple field (deviceConfigurationAckDataBlock)
	lengthInBits += m.DeviceConfigurationAckDataBlock.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_DeviceConfigurationAck) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DeviceConfigurationAck) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_KnxNetIpMessage) (__deviceConfigurationAck DeviceConfigurationAck, err error) {
	m.KnxNetIpMessageContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DeviceConfigurationAck"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeviceConfigurationAck")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	deviceConfigurationAckDataBlock, err := ReadSimpleField[DeviceConfigurationAckDataBlock](ctx, "deviceConfigurationAckDataBlock", ReadComplex[DeviceConfigurationAckDataBlock](DeviceConfigurationAckDataBlockParseWithBuffer, readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deviceConfigurationAckDataBlock' field"))
	}
	m.DeviceConfigurationAckDataBlock = deviceConfigurationAckDataBlock

	if closeErr := readBuffer.CloseContext("DeviceConfigurationAck"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeviceConfigurationAck")
	}

	return m, nil
}

func (m *_DeviceConfigurationAck) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DeviceConfigurationAck) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DeviceConfigurationAck"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DeviceConfigurationAck")
		}

		if err := WriteSimpleField[DeviceConfigurationAckDataBlock](ctx, "deviceConfigurationAckDataBlock", m.GetDeviceConfigurationAckDataBlock(), WriteComplex[DeviceConfigurationAckDataBlock](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'deviceConfigurationAckDataBlock' field")
		}

		if popErr := writeBuffer.PopContext("DeviceConfigurationAck"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DeviceConfigurationAck")
		}
		return nil
	}
	return m.KnxNetIpMessageContract.(*_KnxNetIpMessage).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DeviceConfigurationAck) IsDeviceConfigurationAck() {}

func (m *_DeviceConfigurationAck) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DeviceConfigurationAck) deepCopy() *_DeviceConfigurationAck {
	if m == nil {
		return nil
	}
	_DeviceConfigurationAckCopy := &_DeviceConfigurationAck{
		m.KnxNetIpMessageContract.(*_KnxNetIpMessage).deepCopy(),
		utils.DeepCopy[DeviceConfigurationAckDataBlock](m.DeviceConfigurationAckDataBlock),
	}
	_DeviceConfigurationAckCopy.KnxNetIpMessageContract.(*_KnxNetIpMessage)._SubType = m
	return _DeviceConfigurationAckCopy
}

func (m *_DeviceConfigurationAck) String() string {
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
