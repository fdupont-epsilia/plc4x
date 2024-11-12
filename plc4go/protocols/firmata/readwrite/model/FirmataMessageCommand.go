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

// FirmataMessageCommand is the corresponding interface of FirmataMessageCommand
type FirmataMessageCommand interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	FirmataMessage
	// GetCommand returns Command (property field)
	GetCommand() FirmataCommand
	// IsFirmataMessageCommand is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsFirmataMessageCommand()
	// CreateBuilder creates a FirmataMessageCommandBuilder
	CreateFirmataMessageCommandBuilder() FirmataMessageCommandBuilder
}

// _FirmataMessageCommand is the data-structure of this message
type _FirmataMessageCommand struct {
	FirmataMessageContract
	Command FirmataCommand
}

var _ FirmataMessageCommand = (*_FirmataMessageCommand)(nil)
var _ FirmataMessageRequirements = (*_FirmataMessageCommand)(nil)

// NewFirmataMessageCommand factory function for _FirmataMessageCommand
func NewFirmataMessageCommand(command FirmataCommand, response bool) *_FirmataMessageCommand {
	if command == nil {
		panic("command of type FirmataCommand for FirmataMessageCommand must not be nil")
	}
	_result := &_FirmataMessageCommand{
		FirmataMessageContract: NewFirmataMessage(response),
		Command:                command,
	}
	_result.FirmataMessageContract.(*_FirmataMessage)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// FirmataMessageCommandBuilder is a builder for FirmataMessageCommand
type FirmataMessageCommandBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(command FirmataCommand) FirmataMessageCommandBuilder
	// WithCommand adds Command (property field)
	WithCommand(FirmataCommand) FirmataMessageCommandBuilder
	// WithCommandBuilder adds Command (property field) which is build by the builder
	WithCommandBuilder(func(FirmataCommandBuilder) FirmataCommandBuilder) FirmataMessageCommandBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() FirmataMessageBuilder
	// Build builds the FirmataMessageCommand or returns an error if something is wrong
	Build() (FirmataMessageCommand, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() FirmataMessageCommand
}

// NewFirmataMessageCommandBuilder() creates a FirmataMessageCommandBuilder
func NewFirmataMessageCommandBuilder() FirmataMessageCommandBuilder {
	return &_FirmataMessageCommandBuilder{_FirmataMessageCommand: new(_FirmataMessageCommand)}
}

type _FirmataMessageCommandBuilder struct {
	*_FirmataMessageCommand

	parentBuilder *_FirmataMessageBuilder

	err *utils.MultiError
}

var _ (FirmataMessageCommandBuilder) = (*_FirmataMessageCommandBuilder)(nil)

func (b *_FirmataMessageCommandBuilder) setParent(contract FirmataMessageContract) {
	b.FirmataMessageContract = contract
	contract.(*_FirmataMessage)._SubType = b._FirmataMessageCommand
}

func (b *_FirmataMessageCommandBuilder) WithMandatoryFields(command FirmataCommand) FirmataMessageCommandBuilder {
	return b.WithCommand(command)
}

func (b *_FirmataMessageCommandBuilder) WithCommand(command FirmataCommand) FirmataMessageCommandBuilder {
	b.Command = command
	return b
}

func (b *_FirmataMessageCommandBuilder) WithCommandBuilder(builderSupplier func(FirmataCommandBuilder) FirmataCommandBuilder) FirmataMessageCommandBuilder {
	builder := builderSupplier(b.Command.CreateFirmataCommandBuilder())
	var err error
	b.Command, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "FirmataCommandBuilder failed"))
	}
	return b
}

func (b *_FirmataMessageCommandBuilder) Build() (FirmataMessageCommand, error) {
	if b.Command == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'command' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._FirmataMessageCommand.deepCopy(), nil
}

func (b *_FirmataMessageCommandBuilder) MustBuild() FirmataMessageCommand {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_FirmataMessageCommandBuilder) Done() FirmataMessageBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewFirmataMessageBuilder().(*_FirmataMessageBuilder)
	}
	return b.parentBuilder
}

func (b *_FirmataMessageCommandBuilder) buildForFirmataMessage() (FirmataMessage, error) {
	return b.Build()
}

func (b *_FirmataMessageCommandBuilder) DeepCopy() any {
	_copy := b.CreateFirmataMessageCommandBuilder().(*_FirmataMessageCommandBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateFirmataMessageCommandBuilder creates a FirmataMessageCommandBuilder
func (b *_FirmataMessageCommand) CreateFirmataMessageCommandBuilder() FirmataMessageCommandBuilder {
	if b == nil {
		return NewFirmataMessageCommandBuilder()
	}
	return &_FirmataMessageCommandBuilder{_FirmataMessageCommand: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FirmataMessageCommand) GetMessageType() uint8 {
	return 0xF
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FirmataMessageCommand) GetParent() FirmataMessageContract {
	return m.FirmataMessageContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_FirmataMessageCommand) GetCommand() FirmataCommand {
	return m.Command
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastFirmataMessageCommand(structType any) FirmataMessageCommand {
	if casted, ok := structType.(FirmataMessageCommand); ok {
		return casted
	}
	if casted, ok := structType.(*FirmataMessageCommand); ok {
		return *casted
	}
	return nil
}

func (m *_FirmataMessageCommand) GetTypeName() string {
	return "FirmataMessageCommand"
}

func (m *_FirmataMessageCommand) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.FirmataMessageContract.(*_FirmataMessage).getLengthInBits(ctx))

	// Simple field (command)
	lengthInBits += m.Command.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_FirmataMessageCommand) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_FirmataMessageCommand) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_FirmataMessage, response bool) (__firmataMessageCommand FirmataMessageCommand, err error) {
	m.FirmataMessageContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FirmataMessageCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FirmataMessageCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	command, err := ReadSimpleField[FirmataCommand](ctx, "command", ReadComplex[FirmataCommand](FirmataCommandParseWithBufferProducer[FirmataCommand]((bool)(response)), readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'command' field"))
	}
	m.Command = command

	if closeErr := readBuffer.CloseContext("FirmataMessageCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FirmataMessageCommand")
	}

	return m, nil
}

func (m *_FirmataMessageCommand) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FirmataMessageCommand) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FirmataMessageCommand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FirmataMessageCommand")
		}

		if err := WriteSimpleField[FirmataCommand](ctx, "command", m.GetCommand(), WriteComplex[FirmataCommand](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'command' field")
		}

		if popErr := writeBuffer.PopContext("FirmataMessageCommand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FirmataMessageCommand")
		}
		return nil
	}
	return m.FirmataMessageContract.(*_FirmataMessage).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_FirmataMessageCommand) IsFirmataMessageCommand() {}

func (m *_FirmataMessageCommand) DeepCopy() any {
	return m.deepCopy()
}

func (m *_FirmataMessageCommand) deepCopy() *_FirmataMessageCommand {
	if m == nil {
		return nil
	}
	_FirmataMessageCommandCopy := &_FirmataMessageCommand{
		m.FirmataMessageContract.(*_FirmataMessage).deepCopy(),
		utils.DeepCopy[FirmataCommand](m.Command),
	}
	_FirmataMessageCommandCopy.FirmataMessageContract.(*_FirmataMessage)._SubType = m
	return _FirmataMessageCommandCopy
}

func (m *_FirmataMessageCommand) String() string {
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
