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

// AccessControlDataRequestToExit is the corresponding interface of AccessControlDataRequestToExit
type AccessControlDataRequestToExit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AccessControlData
	// IsAccessControlDataRequestToExit is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAccessControlDataRequestToExit()
	// CreateBuilder creates a AccessControlDataRequestToExitBuilder
	CreateAccessControlDataRequestToExitBuilder() AccessControlDataRequestToExitBuilder
}

// _AccessControlDataRequestToExit is the data-structure of this message
type _AccessControlDataRequestToExit struct {
	AccessControlDataContract
}

var _ AccessControlDataRequestToExit = (*_AccessControlDataRequestToExit)(nil)
var _ AccessControlDataRequirements = (*_AccessControlDataRequestToExit)(nil)

// NewAccessControlDataRequestToExit factory function for _AccessControlDataRequestToExit
func NewAccessControlDataRequestToExit(commandTypeContainer AccessControlCommandTypeContainer, networkId byte, accessPointId byte) *_AccessControlDataRequestToExit {
	_result := &_AccessControlDataRequestToExit{
		AccessControlDataContract: NewAccessControlData(commandTypeContainer, networkId, accessPointId),
	}
	_result.AccessControlDataContract.(*_AccessControlData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AccessControlDataRequestToExitBuilder is a builder for AccessControlDataRequestToExit
type AccessControlDataRequestToExitBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() AccessControlDataRequestToExitBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() AccessControlDataBuilder
	// Build builds the AccessControlDataRequestToExit or returns an error if something is wrong
	Build() (AccessControlDataRequestToExit, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AccessControlDataRequestToExit
}

// NewAccessControlDataRequestToExitBuilder() creates a AccessControlDataRequestToExitBuilder
func NewAccessControlDataRequestToExitBuilder() AccessControlDataRequestToExitBuilder {
	return &_AccessControlDataRequestToExitBuilder{_AccessControlDataRequestToExit: new(_AccessControlDataRequestToExit)}
}

type _AccessControlDataRequestToExitBuilder struct {
	*_AccessControlDataRequestToExit

	parentBuilder *_AccessControlDataBuilder

	err *utils.MultiError
}

var _ (AccessControlDataRequestToExitBuilder) = (*_AccessControlDataRequestToExitBuilder)(nil)

func (b *_AccessControlDataRequestToExitBuilder) setParent(contract AccessControlDataContract) {
	b.AccessControlDataContract = contract
	contract.(*_AccessControlData)._SubType = b._AccessControlDataRequestToExit
}

func (b *_AccessControlDataRequestToExitBuilder) WithMandatoryFields() AccessControlDataRequestToExitBuilder {
	return b
}

func (b *_AccessControlDataRequestToExitBuilder) Build() (AccessControlDataRequestToExit, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AccessControlDataRequestToExit.deepCopy(), nil
}

func (b *_AccessControlDataRequestToExitBuilder) MustBuild() AccessControlDataRequestToExit {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_AccessControlDataRequestToExitBuilder) Done() AccessControlDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewAccessControlDataBuilder().(*_AccessControlDataBuilder)
	}
	return b.parentBuilder
}

func (b *_AccessControlDataRequestToExitBuilder) buildForAccessControlData() (AccessControlData, error) {
	return b.Build()
}

func (b *_AccessControlDataRequestToExitBuilder) DeepCopy() any {
	_copy := b.CreateAccessControlDataRequestToExitBuilder().(*_AccessControlDataRequestToExitBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAccessControlDataRequestToExitBuilder creates a AccessControlDataRequestToExitBuilder
func (b *_AccessControlDataRequestToExit) CreateAccessControlDataRequestToExitBuilder() AccessControlDataRequestToExitBuilder {
	if b == nil {
		return NewAccessControlDataRequestToExitBuilder()
	}
	return &_AccessControlDataRequestToExitBuilder{_AccessControlDataRequestToExit: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AccessControlDataRequestToExit) GetParent() AccessControlDataContract {
	return m.AccessControlDataContract
}

// Deprecated: use the interface for direct cast
func CastAccessControlDataRequestToExit(structType any) AccessControlDataRequestToExit {
	if casted, ok := structType.(AccessControlDataRequestToExit); ok {
		return casted
	}
	if casted, ok := structType.(*AccessControlDataRequestToExit); ok {
		return *casted
	}
	return nil
}

func (m *_AccessControlDataRequestToExit) GetTypeName() string {
	return "AccessControlDataRequestToExit"
}

func (m *_AccessControlDataRequestToExit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AccessControlDataContract.(*_AccessControlData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_AccessControlDataRequestToExit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AccessControlDataRequestToExit) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AccessControlData) (__accessControlDataRequestToExit AccessControlDataRequestToExit, err error) {
	m.AccessControlDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AccessControlDataRequestToExit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AccessControlDataRequestToExit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("AccessControlDataRequestToExit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AccessControlDataRequestToExit")
	}

	return m, nil
}

func (m *_AccessControlDataRequestToExit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AccessControlDataRequestToExit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AccessControlDataRequestToExit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AccessControlDataRequestToExit")
		}

		if popErr := writeBuffer.PopContext("AccessControlDataRequestToExit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AccessControlDataRequestToExit")
		}
		return nil
	}
	return m.AccessControlDataContract.(*_AccessControlData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AccessControlDataRequestToExit) IsAccessControlDataRequestToExit() {}

func (m *_AccessControlDataRequestToExit) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AccessControlDataRequestToExit) deepCopy() *_AccessControlDataRequestToExit {
	if m == nil {
		return nil
	}
	_AccessControlDataRequestToExitCopy := &_AccessControlDataRequestToExit{
		m.AccessControlDataContract.(*_AccessControlData).deepCopy(),
	}
	_AccessControlDataRequestToExitCopy.AccessControlDataContract.(*_AccessControlData)._SubType = m
	return _AccessControlDataRequestToExitCopy
}

func (m *_AccessControlDataRequestToExit) String() string {
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
