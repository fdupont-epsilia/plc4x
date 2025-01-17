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

// CBusCommandPointToPointToMultiPoint is the corresponding interface of CBusCommandPointToPointToMultiPoint
type CBusCommandPointToPointToMultiPoint interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CBusCommand
	// GetCommand returns Command (property field)
	GetCommand() CBusPointToPointToMultiPointCommand
	// IsCBusCommandPointToPointToMultiPoint is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCBusCommandPointToPointToMultiPoint()
	// CreateBuilder creates a CBusCommandPointToPointToMultiPointBuilder
	CreateCBusCommandPointToPointToMultiPointBuilder() CBusCommandPointToPointToMultiPointBuilder
}

// _CBusCommandPointToPointToMultiPoint is the data-structure of this message
type _CBusCommandPointToPointToMultiPoint struct {
	CBusCommandContract
	Command CBusPointToPointToMultiPointCommand
}

var _ CBusCommandPointToPointToMultiPoint = (*_CBusCommandPointToPointToMultiPoint)(nil)
var _ CBusCommandRequirements = (*_CBusCommandPointToPointToMultiPoint)(nil)

// NewCBusCommandPointToPointToMultiPoint factory function for _CBusCommandPointToPointToMultiPoint
func NewCBusCommandPointToPointToMultiPoint(header CBusHeader, command CBusPointToPointToMultiPointCommand, cBusOptions CBusOptions) *_CBusCommandPointToPointToMultiPoint {
	if command == nil {
		panic("command of type CBusPointToPointToMultiPointCommand for CBusCommandPointToPointToMultiPoint must not be nil")
	}
	_result := &_CBusCommandPointToPointToMultiPoint{
		CBusCommandContract: NewCBusCommand(header, cBusOptions),
		Command:             command,
	}
	_result.CBusCommandContract.(*_CBusCommand)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CBusCommandPointToPointToMultiPointBuilder is a builder for CBusCommandPointToPointToMultiPoint
type CBusCommandPointToPointToMultiPointBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(command CBusPointToPointToMultiPointCommand) CBusCommandPointToPointToMultiPointBuilder
	// WithCommand adds Command (property field)
	WithCommand(CBusPointToPointToMultiPointCommand) CBusCommandPointToPointToMultiPointBuilder
	// WithCommandBuilder adds Command (property field) which is build by the builder
	WithCommandBuilder(func(CBusPointToPointToMultiPointCommandBuilder) CBusPointToPointToMultiPointCommandBuilder) CBusCommandPointToPointToMultiPointBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() CBusCommandBuilder
	// Build builds the CBusCommandPointToPointToMultiPoint or returns an error if something is wrong
	Build() (CBusCommandPointToPointToMultiPoint, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CBusCommandPointToPointToMultiPoint
}

// NewCBusCommandPointToPointToMultiPointBuilder() creates a CBusCommandPointToPointToMultiPointBuilder
func NewCBusCommandPointToPointToMultiPointBuilder() CBusCommandPointToPointToMultiPointBuilder {
	return &_CBusCommandPointToPointToMultiPointBuilder{_CBusCommandPointToPointToMultiPoint: new(_CBusCommandPointToPointToMultiPoint)}
}

type _CBusCommandPointToPointToMultiPointBuilder struct {
	*_CBusCommandPointToPointToMultiPoint

	parentBuilder *_CBusCommandBuilder

	err *utils.MultiError
}

var _ (CBusCommandPointToPointToMultiPointBuilder) = (*_CBusCommandPointToPointToMultiPointBuilder)(nil)

func (b *_CBusCommandPointToPointToMultiPointBuilder) setParent(contract CBusCommandContract) {
	b.CBusCommandContract = contract
	contract.(*_CBusCommand)._SubType = b._CBusCommandPointToPointToMultiPoint
}

func (b *_CBusCommandPointToPointToMultiPointBuilder) WithMandatoryFields(command CBusPointToPointToMultiPointCommand) CBusCommandPointToPointToMultiPointBuilder {
	return b.WithCommand(command)
}

func (b *_CBusCommandPointToPointToMultiPointBuilder) WithCommand(command CBusPointToPointToMultiPointCommand) CBusCommandPointToPointToMultiPointBuilder {
	b.Command = command
	return b
}

func (b *_CBusCommandPointToPointToMultiPointBuilder) WithCommandBuilder(builderSupplier func(CBusPointToPointToMultiPointCommandBuilder) CBusPointToPointToMultiPointCommandBuilder) CBusCommandPointToPointToMultiPointBuilder {
	builder := builderSupplier(b.Command.CreateCBusPointToPointToMultiPointCommandBuilder())
	var err error
	b.Command, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "CBusPointToPointToMultiPointCommandBuilder failed"))
	}
	return b
}

func (b *_CBusCommandPointToPointToMultiPointBuilder) Build() (CBusCommandPointToPointToMultiPoint, error) {
	if b.Command == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'command' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CBusCommandPointToPointToMultiPoint.deepCopy(), nil
}

func (b *_CBusCommandPointToPointToMultiPointBuilder) MustBuild() CBusCommandPointToPointToMultiPoint {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CBusCommandPointToPointToMultiPointBuilder) Done() CBusCommandBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewCBusCommandBuilder().(*_CBusCommandBuilder)
	}
	return b.parentBuilder
}

func (b *_CBusCommandPointToPointToMultiPointBuilder) buildForCBusCommand() (CBusCommand, error) {
	return b.Build()
}

func (b *_CBusCommandPointToPointToMultiPointBuilder) DeepCopy() any {
	_copy := b.CreateCBusCommandPointToPointToMultiPointBuilder().(*_CBusCommandPointToPointToMultiPointBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCBusCommandPointToPointToMultiPointBuilder creates a CBusCommandPointToPointToMultiPointBuilder
func (b *_CBusCommandPointToPointToMultiPoint) CreateCBusCommandPointToPointToMultiPointBuilder() CBusCommandPointToPointToMultiPointBuilder {
	if b == nil {
		return NewCBusCommandPointToPointToMultiPointBuilder()
	}
	return &_CBusCommandPointToPointToMultiPointBuilder{_CBusCommandPointToPointToMultiPoint: b.deepCopy()}
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

func (m *_CBusCommandPointToPointToMultiPoint) GetParent() CBusCommandContract {
	return m.CBusCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusCommandPointToPointToMultiPoint) GetCommand() CBusPointToPointToMultiPointCommand {
	return m.Command
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCBusCommandPointToPointToMultiPoint(structType any) CBusCommandPointToPointToMultiPoint {
	if casted, ok := structType.(CBusCommandPointToPointToMultiPoint); ok {
		return casted
	}
	if casted, ok := structType.(*CBusCommandPointToPointToMultiPoint); ok {
		return *casted
	}
	return nil
}

func (m *_CBusCommandPointToPointToMultiPoint) GetTypeName() string {
	return "CBusCommandPointToPointToMultiPoint"
}

func (m *_CBusCommandPointToPointToMultiPoint) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CBusCommandContract.(*_CBusCommand).getLengthInBits(ctx))

	// Simple field (command)
	lengthInBits += m.Command.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CBusCommandPointToPointToMultiPoint) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CBusCommandPointToPointToMultiPoint) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CBusCommand, cBusOptions CBusOptions) (__cBusCommandPointToPointToMultiPoint CBusCommandPointToPointToMultiPoint, err error) {
	m.CBusCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusCommandPointToPointToMultiPoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusCommandPointToPointToMultiPoint")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	command, err := ReadSimpleField[CBusPointToPointToMultiPointCommand](ctx, "command", ReadComplex[CBusPointToPointToMultiPointCommand](CBusPointToPointToMultiPointCommandParseWithBufferProducer[CBusPointToPointToMultiPointCommand]((CBusOptions)(cBusOptions)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'command' field"))
	}
	m.Command = command

	if closeErr := readBuffer.CloseContext("CBusCommandPointToPointToMultiPoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusCommandPointToPointToMultiPoint")
	}

	return m, nil
}

func (m *_CBusCommandPointToPointToMultiPoint) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CBusCommandPointToPointToMultiPoint) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CBusCommandPointToPointToMultiPoint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CBusCommandPointToPointToMultiPoint")
		}

		if err := WriteSimpleField[CBusPointToPointToMultiPointCommand](ctx, "command", m.GetCommand(), WriteComplex[CBusPointToPointToMultiPointCommand](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'command' field")
		}

		if popErr := writeBuffer.PopContext("CBusCommandPointToPointToMultiPoint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CBusCommandPointToPointToMultiPoint")
		}
		return nil
	}
	return m.CBusCommandContract.(*_CBusCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CBusCommandPointToPointToMultiPoint) IsCBusCommandPointToPointToMultiPoint() {}

func (m *_CBusCommandPointToPointToMultiPoint) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CBusCommandPointToPointToMultiPoint) deepCopy() *_CBusCommandPointToPointToMultiPoint {
	if m == nil {
		return nil
	}
	_CBusCommandPointToPointToMultiPointCopy := &_CBusCommandPointToPointToMultiPoint{
		m.CBusCommandContract.(*_CBusCommand).deepCopy(),
		utils.DeepCopy[CBusPointToPointToMultiPointCommand](m.Command),
	}
	_CBusCommandPointToPointToMultiPointCopy.CBusCommandContract.(*_CBusCommand)._SubType = m
	return _CBusCommandPointToPointToMultiPointCopy
}

func (m *_CBusCommandPointToPointToMultiPoint) String() string {
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
