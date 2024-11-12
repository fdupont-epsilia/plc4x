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

// ConnectionResponseDataBlockDeviceManagement is the corresponding interface of ConnectionResponseDataBlockDeviceManagement
type ConnectionResponseDataBlockDeviceManagement interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ConnectionResponseDataBlock
	// IsConnectionResponseDataBlockDeviceManagement is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsConnectionResponseDataBlockDeviceManagement()
	// CreateBuilder creates a ConnectionResponseDataBlockDeviceManagementBuilder
	CreateConnectionResponseDataBlockDeviceManagementBuilder() ConnectionResponseDataBlockDeviceManagementBuilder
}

// _ConnectionResponseDataBlockDeviceManagement is the data-structure of this message
type _ConnectionResponseDataBlockDeviceManagement struct {
	ConnectionResponseDataBlockContract
}

var _ ConnectionResponseDataBlockDeviceManagement = (*_ConnectionResponseDataBlockDeviceManagement)(nil)
var _ ConnectionResponseDataBlockRequirements = (*_ConnectionResponseDataBlockDeviceManagement)(nil)

// NewConnectionResponseDataBlockDeviceManagement factory function for _ConnectionResponseDataBlockDeviceManagement
func NewConnectionResponseDataBlockDeviceManagement() *_ConnectionResponseDataBlockDeviceManagement {
	_result := &_ConnectionResponseDataBlockDeviceManagement{
		ConnectionResponseDataBlockContract: NewConnectionResponseDataBlock(),
	}
	_result.ConnectionResponseDataBlockContract.(*_ConnectionResponseDataBlock)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ConnectionResponseDataBlockDeviceManagementBuilder is a builder for ConnectionResponseDataBlockDeviceManagement
type ConnectionResponseDataBlockDeviceManagementBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ConnectionResponseDataBlockDeviceManagementBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ConnectionResponseDataBlockBuilder
	// Build builds the ConnectionResponseDataBlockDeviceManagement or returns an error if something is wrong
	Build() (ConnectionResponseDataBlockDeviceManagement, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ConnectionResponseDataBlockDeviceManagement
}

// NewConnectionResponseDataBlockDeviceManagementBuilder() creates a ConnectionResponseDataBlockDeviceManagementBuilder
func NewConnectionResponseDataBlockDeviceManagementBuilder() ConnectionResponseDataBlockDeviceManagementBuilder {
	return &_ConnectionResponseDataBlockDeviceManagementBuilder{_ConnectionResponseDataBlockDeviceManagement: new(_ConnectionResponseDataBlockDeviceManagement)}
}

type _ConnectionResponseDataBlockDeviceManagementBuilder struct {
	*_ConnectionResponseDataBlockDeviceManagement

	parentBuilder *_ConnectionResponseDataBlockBuilder

	err *utils.MultiError
}

var _ (ConnectionResponseDataBlockDeviceManagementBuilder) = (*_ConnectionResponseDataBlockDeviceManagementBuilder)(nil)

func (b *_ConnectionResponseDataBlockDeviceManagementBuilder) setParent(contract ConnectionResponseDataBlockContract) {
	b.ConnectionResponseDataBlockContract = contract
	contract.(*_ConnectionResponseDataBlock)._SubType = b._ConnectionResponseDataBlockDeviceManagement
}

func (b *_ConnectionResponseDataBlockDeviceManagementBuilder) WithMandatoryFields() ConnectionResponseDataBlockDeviceManagementBuilder {
	return b
}

func (b *_ConnectionResponseDataBlockDeviceManagementBuilder) Build() (ConnectionResponseDataBlockDeviceManagement, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ConnectionResponseDataBlockDeviceManagement.deepCopy(), nil
}

func (b *_ConnectionResponseDataBlockDeviceManagementBuilder) MustBuild() ConnectionResponseDataBlockDeviceManagement {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ConnectionResponseDataBlockDeviceManagementBuilder) Done() ConnectionResponseDataBlockBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewConnectionResponseDataBlockBuilder().(*_ConnectionResponseDataBlockBuilder)
	}
	return b.parentBuilder
}

func (b *_ConnectionResponseDataBlockDeviceManagementBuilder) buildForConnectionResponseDataBlock() (ConnectionResponseDataBlock, error) {
	return b.Build()
}

func (b *_ConnectionResponseDataBlockDeviceManagementBuilder) DeepCopy() any {
	_copy := b.CreateConnectionResponseDataBlockDeviceManagementBuilder().(*_ConnectionResponseDataBlockDeviceManagementBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateConnectionResponseDataBlockDeviceManagementBuilder creates a ConnectionResponseDataBlockDeviceManagementBuilder
func (b *_ConnectionResponseDataBlockDeviceManagement) CreateConnectionResponseDataBlockDeviceManagementBuilder() ConnectionResponseDataBlockDeviceManagementBuilder {
	if b == nil {
		return NewConnectionResponseDataBlockDeviceManagementBuilder()
	}
	return &_ConnectionResponseDataBlockDeviceManagementBuilder{_ConnectionResponseDataBlockDeviceManagement: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ConnectionResponseDataBlockDeviceManagement) GetConnectionType() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ConnectionResponseDataBlockDeviceManagement) GetParent() ConnectionResponseDataBlockContract {
	return m.ConnectionResponseDataBlockContract
}

// Deprecated: use the interface for direct cast
func CastConnectionResponseDataBlockDeviceManagement(structType any) ConnectionResponseDataBlockDeviceManagement {
	if casted, ok := structType.(ConnectionResponseDataBlockDeviceManagement); ok {
		return casted
	}
	if casted, ok := structType.(*ConnectionResponseDataBlockDeviceManagement); ok {
		return *casted
	}
	return nil
}

func (m *_ConnectionResponseDataBlockDeviceManagement) GetTypeName() string {
	return "ConnectionResponseDataBlockDeviceManagement"
}

func (m *_ConnectionResponseDataBlockDeviceManagement) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ConnectionResponseDataBlockContract.(*_ConnectionResponseDataBlock).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ConnectionResponseDataBlockDeviceManagement) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ConnectionResponseDataBlockDeviceManagement) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ConnectionResponseDataBlock) (__connectionResponseDataBlockDeviceManagement ConnectionResponseDataBlockDeviceManagement, err error) {
	m.ConnectionResponseDataBlockContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ConnectionResponseDataBlockDeviceManagement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ConnectionResponseDataBlockDeviceManagement")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ConnectionResponseDataBlockDeviceManagement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ConnectionResponseDataBlockDeviceManagement")
	}

	return m, nil
}

func (m *_ConnectionResponseDataBlockDeviceManagement) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ConnectionResponseDataBlockDeviceManagement) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ConnectionResponseDataBlockDeviceManagement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ConnectionResponseDataBlockDeviceManagement")
		}

		if popErr := writeBuffer.PopContext("ConnectionResponseDataBlockDeviceManagement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ConnectionResponseDataBlockDeviceManagement")
		}
		return nil
	}
	return m.ConnectionResponseDataBlockContract.(*_ConnectionResponseDataBlock).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ConnectionResponseDataBlockDeviceManagement) IsConnectionResponseDataBlockDeviceManagement() {
}

func (m *_ConnectionResponseDataBlockDeviceManagement) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ConnectionResponseDataBlockDeviceManagement) deepCopy() *_ConnectionResponseDataBlockDeviceManagement {
	if m == nil {
		return nil
	}
	_ConnectionResponseDataBlockDeviceManagementCopy := &_ConnectionResponseDataBlockDeviceManagement{
		m.ConnectionResponseDataBlockContract.(*_ConnectionResponseDataBlock).deepCopy(),
	}
	_ConnectionResponseDataBlockDeviceManagementCopy.ConnectionResponseDataBlockContract.(*_ConnectionResponseDataBlock)._SubType = m
	return _ConnectionResponseDataBlockDeviceManagementCopy
}

func (m *_ConnectionResponseDataBlockDeviceManagement) String() string {
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
