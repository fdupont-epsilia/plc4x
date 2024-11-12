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

// CBusPointToPointToMultiPointCommandStatus is the corresponding interface of CBusPointToPointToMultiPointCommandStatus
type CBusPointToPointToMultiPointCommandStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CBusPointToPointToMultiPointCommand
	// GetStatusRequest returns StatusRequest (property field)
	GetStatusRequest() StatusRequest
	// IsCBusPointToPointToMultiPointCommandStatus is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCBusPointToPointToMultiPointCommandStatus()
	// CreateBuilder creates a CBusPointToPointToMultiPointCommandStatusBuilder
	CreateCBusPointToPointToMultiPointCommandStatusBuilder() CBusPointToPointToMultiPointCommandStatusBuilder
}

// _CBusPointToPointToMultiPointCommandStatus is the data-structure of this message
type _CBusPointToPointToMultiPointCommandStatus struct {
	CBusPointToPointToMultiPointCommandContract
	StatusRequest StatusRequest
	// Reserved Fields
	reservedField0 *byte
}

var _ CBusPointToPointToMultiPointCommandStatus = (*_CBusPointToPointToMultiPointCommandStatus)(nil)
var _ CBusPointToPointToMultiPointCommandRequirements = (*_CBusPointToPointToMultiPointCommandStatus)(nil)

// NewCBusPointToPointToMultiPointCommandStatus factory function for _CBusPointToPointToMultiPointCommandStatus
func NewCBusPointToPointToMultiPointCommandStatus(bridgeAddress BridgeAddress, networkRoute NetworkRoute, peekedApplication byte, statusRequest StatusRequest, cBusOptions CBusOptions) *_CBusPointToPointToMultiPointCommandStatus {
	if statusRequest == nil {
		panic("statusRequest of type StatusRequest for CBusPointToPointToMultiPointCommandStatus must not be nil")
	}
	_result := &_CBusPointToPointToMultiPointCommandStatus{
		CBusPointToPointToMultiPointCommandContract: NewCBusPointToPointToMultiPointCommand(bridgeAddress, networkRoute, peekedApplication, cBusOptions),
		StatusRequest: statusRequest,
	}
	_result.CBusPointToPointToMultiPointCommandContract.(*_CBusPointToPointToMultiPointCommand)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CBusPointToPointToMultiPointCommandStatusBuilder is a builder for CBusPointToPointToMultiPointCommandStatus
type CBusPointToPointToMultiPointCommandStatusBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(statusRequest StatusRequest) CBusPointToPointToMultiPointCommandStatusBuilder
	// WithStatusRequest adds StatusRequest (property field)
	WithStatusRequest(StatusRequest) CBusPointToPointToMultiPointCommandStatusBuilder
	// WithStatusRequestBuilder adds StatusRequest (property field) which is build by the builder
	WithStatusRequestBuilder(func(StatusRequestBuilder) StatusRequestBuilder) CBusPointToPointToMultiPointCommandStatusBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() CBusPointToPointToMultiPointCommandBuilder
	// Build builds the CBusPointToPointToMultiPointCommandStatus or returns an error if something is wrong
	Build() (CBusPointToPointToMultiPointCommandStatus, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CBusPointToPointToMultiPointCommandStatus
}

// NewCBusPointToPointToMultiPointCommandStatusBuilder() creates a CBusPointToPointToMultiPointCommandStatusBuilder
func NewCBusPointToPointToMultiPointCommandStatusBuilder() CBusPointToPointToMultiPointCommandStatusBuilder {
	return &_CBusPointToPointToMultiPointCommandStatusBuilder{_CBusPointToPointToMultiPointCommandStatus: new(_CBusPointToPointToMultiPointCommandStatus)}
}

type _CBusPointToPointToMultiPointCommandStatusBuilder struct {
	*_CBusPointToPointToMultiPointCommandStatus

	parentBuilder *_CBusPointToPointToMultiPointCommandBuilder

	err *utils.MultiError
}

var _ (CBusPointToPointToMultiPointCommandStatusBuilder) = (*_CBusPointToPointToMultiPointCommandStatusBuilder)(nil)

func (b *_CBusPointToPointToMultiPointCommandStatusBuilder) setParent(contract CBusPointToPointToMultiPointCommandContract) {
	b.CBusPointToPointToMultiPointCommandContract = contract
	contract.(*_CBusPointToPointToMultiPointCommand)._SubType = b._CBusPointToPointToMultiPointCommandStatus
}

func (b *_CBusPointToPointToMultiPointCommandStatusBuilder) WithMandatoryFields(statusRequest StatusRequest) CBusPointToPointToMultiPointCommandStatusBuilder {
	return b.WithStatusRequest(statusRequest)
}

func (b *_CBusPointToPointToMultiPointCommandStatusBuilder) WithStatusRequest(statusRequest StatusRequest) CBusPointToPointToMultiPointCommandStatusBuilder {
	b.StatusRequest = statusRequest
	return b
}

func (b *_CBusPointToPointToMultiPointCommandStatusBuilder) WithStatusRequestBuilder(builderSupplier func(StatusRequestBuilder) StatusRequestBuilder) CBusPointToPointToMultiPointCommandStatusBuilder {
	builder := builderSupplier(b.StatusRequest.CreateStatusRequestBuilder())
	var err error
	b.StatusRequest, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "StatusRequestBuilder failed"))
	}
	return b
}

func (b *_CBusPointToPointToMultiPointCommandStatusBuilder) Build() (CBusPointToPointToMultiPointCommandStatus, error) {
	if b.StatusRequest == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'statusRequest' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CBusPointToPointToMultiPointCommandStatus.deepCopy(), nil
}

func (b *_CBusPointToPointToMultiPointCommandStatusBuilder) MustBuild() CBusPointToPointToMultiPointCommandStatus {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CBusPointToPointToMultiPointCommandStatusBuilder) Done() CBusPointToPointToMultiPointCommandBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewCBusPointToPointToMultiPointCommandBuilder().(*_CBusPointToPointToMultiPointCommandBuilder)
	}
	return b.parentBuilder
}

func (b *_CBusPointToPointToMultiPointCommandStatusBuilder) buildForCBusPointToPointToMultiPointCommand() (CBusPointToPointToMultiPointCommand, error) {
	return b.Build()
}

func (b *_CBusPointToPointToMultiPointCommandStatusBuilder) DeepCopy() any {
	_copy := b.CreateCBusPointToPointToMultiPointCommandStatusBuilder().(*_CBusPointToPointToMultiPointCommandStatusBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCBusPointToPointToMultiPointCommandStatusBuilder creates a CBusPointToPointToMultiPointCommandStatusBuilder
func (b *_CBusPointToPointToMultiPointCommandStatus) CreateCBusPointToPointToMultiPointCommandStatusBuilder() CBusPointToPointToMultiPointCommandStatusBuilder {
	if b == nil {
		return NewCBusPointToPointToMultiPointCommandStatusBuilder()
	}
	return &_CBusPointToPointToMultiPointCommandStatusBuilder{_CBusPointToPointToMultiPointCommandStatus: b.deepCopy()}
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

func (m *_CBusPointToPointToMultiPointCommandStatus) GetParent() CBusPointToPointToMultiPointCommandContract {
	return m.CBusPointToPointToMultiPointCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusPointToPointToMultiPointCommandStatus) GetStatusRequest() StatusRequest {
	return m.StatusRequest
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCBusPointToPointToMultiPointCommandStatus(structType any) CBusPointToPointToMultiPointCommandStatus {
	if casted, ok := structType.(CBusPointToPointToMultiPointCommandStatus); ok {
		return casted
	}
	if casted, ok := structType.(*CBusPointToPointToMultiPointCommandStatus); ok {
		return *casted
	}
	return nil
}

func (m *_CBusPointToPointToMultiPointCommandStatus) GetTypeName() string {
	return "CBusPointToPointToMultiPointCommandStatus"
}

func (m *_CBusPointToPointToMultiPointCommandStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CBusPointToPointToMultiPointCommandContract.(*_CBusPointToPointToMultiPointCommand).getLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (statusRequest)
	lengthInBits += m.StatusRequest.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CBusPointToPointToMultiPointCommandStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CBusPointToPointToMultiPointCommandStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CBusPointToPointToMultiPointCommand, cBusOptions CBusOptions) (__cBusPointToPointToMultiPointCommandStatus CBusPointToPointToMultiPointCommandStatus, err error) {
	m.CBusPointToPointToMultiPointCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusPointToPointToMultiPointCommandStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusPointToPointToMultiPointCommandStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadByte(readBuffer, 8), byte(0xFF))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	statusRequest, err := ReadSimpleField[StatusRequest](ctx, "statusRequest", ReadComplex[StatusRequest](StatusRequestParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusRequest' field"))
	}
	m.StatusRequest = statusRequest

	if closeErr := readBuffer.CloseContext("CBusPointToPointToMultiPointCommandStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusPointToPointToMultiPointCommandStatus")
	}

	return m, nil
}

func (m *_CBusPointToPointToMultiPointCommandStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CBusPointToPointToMultiPointCommandStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CBusPointToPointToMultiPointCommandStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CBusPointToPointToMultiPointCommandStatus")
		}

		if err := WriteReservedField[byte](ctx, "reserved", byte(0xFF), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[StatusRequest](ctx, "statusRequest", m.GetStatusRequest(), WriteComplex[StatusRequest](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'statusRequest' field")
		}

		if popErr := writeBuffer.PopContext("CBusPointToPointToMultiPointCommandStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CBusPointToPointToMultiPointCommandStatus")
		}
		return nil
	}
	return m.CBusPointToPointToMultiPointCommandContract.(*_CBusPointToPointToMultiPointCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CBusPointToPointToMultiPointCommandStatus) IsCBusPointToPointToMultiPointCommandStatus() {}

func (m *_CBusPointToPointToMultiPointCommandStatus) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CBusPointToPointToMultiPointCommandStatus) deepCopy() *_CBusPointToPointToMultiPointCommandStatus {
	if m == nil {
		return nil
	}
	_CBusPointToPointToMultiPointCommandStatusCopy := &_CBusPointToPointToMultiPointCommandStatus{
		m.CBusPointToPointToMultiPointCommandContract.(*_CBusPointToPointToMultiPointCommand).deepCopy(),
		utils.DeepCopy[StatusRequest](m.StatusRequest),
		m.reservedField0,
	}
	_CBusPointToPointToMultiPointCommandStatusCopy.CBusPointToPointToMultiPointCommandContract.(*_CBusPointToPointToMultiPointCommand)._SubType = m
	return _CBusPointToPointToMultiPointCommandStatusCopy
}

func (m *_CBusPointToPointToMultiPointCommandStatus) String() string {
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
