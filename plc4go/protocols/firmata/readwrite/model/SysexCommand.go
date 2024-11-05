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

// SysexCommand is the corresponding interface of SysexCommand
type SysexCommand interface {
	SysexCommandContract
	SysexCommandRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsSysexCommand is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSysexCommand()
	// CreateBuilder creates a SysexCommandBuilder
	CreateSysexCommandBuilder() SysexCommandBuilder
}

// SysexCommandContract provides a set of functions which can be overwritten by a sub struct
type SysexCommandContract interface {
	// IsSysexCommand is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSysexCommand()
	// CreateBuilder creates a SysexCommandBuilder
	CreateSysexCommandBuilder() SysexCommandBuilder
}

// SysexCommandRequirements provides a set of functions which need to be implemented by a sub struct
type SysexCommandRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetCommandType returns CommandType (discriminator field)
	GetCommandType() uint8
	// GetResponse returns Response (discriminator field)
	GetResponse() bool
}

// _SysexCommand is the data-structure of this message
type _SysexCommand struct {
	_SubType interface {
		SysexCommandContract
		SysexCommandRequirements
	}
}

var _ SysexCommandContract = (*_SysexCommand)(nil)

// NewSysexCommand factory function for _SysexCommand
func NewSysexCommand() *_SysexCommand {
	return &_SysexCommand{}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SysexCommandBuilder is a builder for SysexCommand
type SysexCommandBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() SysexCommandBuilder
	// AsSysexCommandExtendedId converts this build to a subType of SysexCommand. It is always possible to return to current builder using Done()
	AsSysexCommandExtendedId() interface {
		SysexCommandExtendedIdBuilder
		Done() SysexCommandBuilder
	}
	// AsSysexCommandAnalogMappingQueryRequest converts this build to a subType of SysexCommand. It is always possible to return to current builder using Done()
	AsSysexCommandAnalogMappingQueryRequest() interface {
		SysexCommandAnalogMappingQueryRequestBuilder
		Done() SysexCommandBuilder
	}
	// AsSysexCommandAnalogMappingQueryResponse converts this build to a subType of SysexCommand. It is always possible to return to current builder using Done()
	AsSysexCommandAnalogMappingQueryResponse() interface {
		SysexCommandAnalogMappingQueryResponseBuilder
		Done() SysexCommandBuilder
	}
	// AsSysexCommandAnalogMappingResponse converts this build to a subType of SysexCommand. It is always possible to return to current builder using Done()
	AsSysexCommandAnalogMappingResponse() interface {
		SysexCommandAnalogMappingResponseBuilder
		Done() SysexCommandBuilder
	}
	// AsSysexCommandCapabilityQuery converts this build to a subType of SysexCommand. It is always possible to return to current builder using Done()
	AsSysexCommandCapabilityQuery() interface {
		SysexCommandCapabilityQueryBuilder
		Done() SysexCommandBuilder
	}
	// AsSysexCommandCapabilityResponse converts this build to a subType of SysexCommand. It is always possible to return to current builder using Done()
	AsSysexCommandCapabilityResponse() interface {
		SysexCommandCapabilityResponseBuilder
		Done() SysexCommandBuilder
	}
	// AsSysexCommandPinStateQuery converts this build to a subType of SysexCommand. It is always possible to return to current builder using Done()
	AsSysexCommandPinStateQuery() interface {
		SysexCommandPinStateQueryBuilder
		Done() SysexCommandBuilder
	}
	// AsSysexCommandPinStateResponse converts this build to a subType of SysexCommand. It is always possible to return to current builder using Done()
	AsSysexCommandPinStateResponse() interface {
		SysexCommandPinStateResponseBuilder
		Done() SysexCommandBuilder
	}
	// AsSysexCommandExtendedAnalog converts this build to a subType of SysexCommand. It is always possible to return to current builder using Done()
	AsSysexCommandExtendedAnalog() interface {
		SysexCommandExtendedAnalogBuilder
		Done() SysexCommandBuilder
	}
	// AsSysexCommandStringData converts this build to a subType of SysexCommand. It is always possible to return to current builder using Done()
	AsSysexCommandStringData() interface {
		SysexCommandStringDataBuilder
		Done() SysexCommandBuilder
	}
	// AsSysexCommandReportFirmwareRequest converts this build to a subType of SysexCommand. It is always possible to return to current builder using Done()
	AsSysexCommandReportFirmwareRequest() interface {
		SysexCommandReportFirmwareRequestBuilder
		Done() SysexCommandBuilder
	}
	// AsSysexCommandReportFirmwareResponse converts this build to a subType of SysexCommand. It is always possible to return to current builder using Done()
	AsSysexCommandReportFirmwareResponse() interface {
		SysexCommandReportFirmwareResponseBuilder
		Done() SysexCommandBuilder
	}
	// AsSysexCommandSamplingInterval converts this build to a subType of SysexCommand. It is always possible to return to current builder using Done()
	AsSysexCommandSamplingInterval() interface {
		SysexCommandSamplingIntervalBuilder
		Done() SysexCommandBuilder
	}
	// AsSysexCommandSysexNonRealtime converts this build to a subType of SysexCommand. It is always possible to return to current builder using Done()
	AsSysexCommandSysexNonRealtime() interface {
		SysexCommandSysexNonRealtimeBuilder
		Done() SysexCommandBuilder
	}
	// AsSysexCommandSysexRealtime converts this build to a subType of SysexCommand. It is always possible to return to current builder using Done()
	AsSysexCommandSysexRealtime() interface {
		SysexCommandSysexRealtimeBuilder
		Done() SysexCommandBuilder
	}
	// Build builds the SysexCommand or returns an error if something is wrong
	PartialBuild() (SysexCommandContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() SysexCommandContract
	// Build builds the SysexCommand or returns an error if something is wrong
	Build() (SysexCommand, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SysexCommand
}

// NewSysexCommandBuilder() creates a SysexCommandBuilder
func NewSysexCommandBuilder() SysexCommandBuilder {
	return &_SysexCommandBuilder{_SysexCommand: new(_SysexCommand)}
}

type _SysexCommandChildBuilder interface {
	utils.Copyable
	setParent(SysexCommandContract)
	buildForSysexCommand() (SysexCommand, error)
}

type _SysexCommandBuilder struct {
	*_SysexCommand

	childBuilder _SysexCommandChildBuilder

	err *utils.MultiError
}

var _ (SysexCommandBuilder) = (*_SysexCommandBuilder)(nil)

func (b *_SysexCommandBuilder) WithMandatoryFields() SysexCommandBuilder {
	return b
}

func (b *_SysexCommandBuilder) PartialBuild() (SysexCommandContract, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SysexCommand.deepCopy(), nil
}

func (b *_SysexCommandBuilder) PartialMustBuild() SysexCommandContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SysexCommandBuilder) AsSysexCommandExtendedId() interface {
	SysexCommandExtendedIdBuilder
	Done() SysexCommandBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		SysexCommandExtendedIdBuilder
		Done() SysexCommandBuilder
	}); ok {
		return cb
	}
	cb := NewSysexCommandExtendedIdBuilder().(*_SysexCommandExtendedIdBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_SysexCommandBuilder) AsSysexCommandAnalogMappingQueryRequest() interface {
	SysexCommandAnalogMappingQueryRequestBuilder
	Done() SysexCommandBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		SysexCommandAnalogMappingQueryRequestBuilder
		Done() SysexCommandBuilder
	}); ok {
		return cb
	}
	cb := NewSysexCommandAnalogMappingQueryRequestBuilder().(*_SysexCommandAnalogMappingQueryRequestBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_SysexCommandBuilder) AsSysexCommandAnalogMappingQueryResponse() interface {
	SysexCommandAnalogMappingQueryResponseBuilder
	Done() SysexCommandBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		SysexCommandAnalogMappingQueryResponseBuilder
		Done() SysexCommandBuilder
	}); ok {
		return cb
	}
	cb := NewSysexCommandAnalogMappingQueryResponseBuilder().(*_SysexCommandAnalogMappingQueryResponseBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_SysexCommandBuilder) AsSysexCommandAnalogMappingResponse() interface {
	SysexCommandAnalogMappingResponseBuilder
	Done() SysexCommandBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		SysexCommandAnalogMappingResponseBuilder
		Done() SysexCommandBuilder
	}); ok {
		return cb
	}
	cb := NewSysexCommandAnalogMappingResponseBuilder().(*_SysexCommandAnalogMappingResponseBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_SysexCommandBuilder) AsSysexCommandCapabilityQuery() interface {
	SysexCommandCapabilityQueryBuilder
	Done() SysexCommandBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		SysexCommandCapabilityQueryBuilder
		Done() SysexCommandBuilder
	}); ok {
		return cb
	}
	cb := NewSysexCommandCapabilityQueryBuilder().(*_SysexCommandCapabilityQueryBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_SysexCommandBuilder) AsSysexCommandCapabilityResponse() interface {
	SysexCommandCapabilityResponseBuilder
	Done() SysexCommandBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		SysexCommandCapabilityResponseBuilder
		Done() SysexCommandBuilder
	}); ok {
		return cb
	}
	cb := NewSysexCommandCapabilityResponseBuilder().(*_SysexCommandCapabilityResponseBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_SysexCommandBuilder) AsSysexCommandPinStateQuery() interface {
	SysexCommandPinStateQueryBuilder
	Done() SysexCommandBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		SysexCommandPinStateQueryBuilder
		Done() SysexCommandBuilder
	}); ok {
		return cb
	}
	cb := NewSysexCommandPinStateQueryBuilder().(*_SysexCommandPinStateQueryBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_SysexCommandBuilder) AsSysexCommandPinStateResponse() interface {
	SysexCommandPinStateResponseBuilder
	Done() SysexCommandBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		SysexCommandPinStateResponseBuilder
		Done() SysexCommandBuilder
	}); ok {
		return cb
	}
	cb := NewSysexCommandPinStateResponseBuilder().(*_SysexCommandPinStateResponseBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_SysexCommandBuilder) AsSysexCommandExtendedAnalog() interface {
	SysexCommandExtendedAnalogBuilder
	Done() SysexCommandBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		SysexCommandExtendedAnalogBuilder
		Done() SysexCommandBuilder
	}); ok {
		return cb
	}
	cb := NewSysexCommandExtendedAnalogBuilder().(*_SysexCommandExtendedAnalogBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_SysexCommandBuilder) AsSysexCommandStringData() interface {
	SysexCommandStringDataBuilder
	Done() SysexCommandBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		SysexCommandStringDataBuilder
		Done() SysexCommandBuilder
	}); ok {
		return cb
	}
	cb := NewSysexCommandStringDataBuilder().(*_SysexCommandStringDataBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_SysexCommandBuilder) AsSysexCommandReportFirmwareRequest() interface {
	SysexCommandReportFirmwareRequestBuilder
	Done() SysexCommandBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		SysexCommandReportFirmwareRequestBuilder
		Done() SysexCommandBuilder
	}); ok {
		return cb
	}
	cb := NewSysexCommandReportFirmwareRequestBuilder().(*_SysexCommandReportFirmwareRequestBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_SysexCommandBuilder) AsSysexCommandReportFirmwareResponse() interface {
	SysexCommandReportFirmwareResponseBuilder
	Done() SysexCommandBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		SysexCommandReportFirmwareResponseBuilder
		Done() SysexCommandBuilder
	}); ok {
		return cb
	}
	cb := NewSysexCommandReportFirmwareResponseBuilder().(*_SysexCommandReportFirmwareResponseBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_SysexCommandBuilder) AsSysexCommandSamplingInterval() interface {
	SysexCommandSamplingIntervalBuilder
	Done() SysexCommandBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		SysexCommandSamplingIntervalBuilder
		Done() SysexCommandBuilder
	}); ok {
		return cb
	}
	cb := NewSysexCommandSamplingIntervalBuilder().(*_SysexCommandSamplingIntervalBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_SysexCommandBuilder) AsSysexCommandSysexNonRealtime() interface {
	SysexCommandSysexNonRealtimeBuilder
	Done() SysexCommandBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		SysexCommandSysexNonRealtimeBuilder
		Done() SysexCommandBuilder
	}); ok {
		return cb
	}
	cb := NewSysexCommandSysexNonRealtimeBuilder().(*_SysexCommandSysexNonRealtimeBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_SysexCommandBuilder) AsSysexCommandSysexRealtime() interface {
	SysexCommandSysexRealtimeBuilder
	Done() SysexCommandBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		SysexCommandSysexRealtimeBuilder
		Done() SysexCommandBuilder
	}); ok {
		return cb
	}
	cb := NewSysexCommandSysexRealtimeBuilder().(*_SysexCommandSysexRealtimeBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_SysexCommandBuilder) Build() (SysexCommand, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForSysexCommand()
}

func (b *_SysexCommandBuilder) MustBuild() SysexCommand {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SysexCommandBuilder) DeepCopy() any {
	_copy := b.CreateSysexCommandBuilder().(*_SysexCommandBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_SysexCommandChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSysexCommandBuilder creates a SysexCommandBuilder
func (b *_SysexCommand) CreateSysexCommandBuilder() SysexCommandBuilder {
	if b == nil {
		return NewSysexCommandBuilder()
	}
	return &_SysexCommandBuilder{_SysexCommand: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSysexCommand(structType any) SysexCommand {
	if casted, ok := structType.(SysexCommand); ok {
		return casted
	}
	if casted, ok := structType.(*SysexCommand); ok {
		return *casted
	}
	return nil
}

func (m *_SysexCommand) GetTypeName() string {
	return "SysexCommand"
}

func (m *_SysexCommand) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (commandType)
	lengthInBits += 8

	return lengthInBits
}

func (m *_SysexCommand) GetLengthInBits(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx)
}

func (m *_SysexCommand) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func SysexCommandParse[T SysexCommand](ctx context.Context, theBytes []byte, response bool) (T, error) {
	return SysexCommandParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func SysexCommandParseWithBufferProducer[T SysexCommand](response bool) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := SysexCommandParseWithBuffer[T](ctx, readBuffer, response)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func SysexCommandParseWithBuffer[T SysexCommand](ctx context.Context, readBuffer utils.ReadBuffer, response bool) (T, error) {
	v, err := (&_SysexCommand{}).parse(ctx, readBuffer, response)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_SysexCommand) parse(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (__sysexCommand SysexCommand, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	commandType, err := ReadDiscriminatorField[uint8](ctx, "commandType", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandType' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child SysexCommand
	switch {
	case commandType == 0x00: // SysexCommandExtendedId
		if _child, err = new(_SysexCommandExtendedId).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SysexCommandExtendedId for type-switch of SysexCommand")
		}
	case commandType == 0x69 && response == bool(false): // SysexCommandAnalogMappingQueryRequest
		if _child, err = new(_SysexCommandAnalogMappingQueryRequest).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SysexCommandAnalogMappingQueryRequest for type-switch of SysexCommand")
		}
	case commandType == 0x69 && response == bool(true): // SysexCommandAnalogMappingQueryResponse
		if _child, err = new(_SysexCommandAnalogMappingQueryResponse).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SysexCommandAnalogMappingQueryResponse for type-switch of SysexCommand")
		}
	case commandType == 0x6A: // SysexCommandAnalogMappingResponse
		if _child, err = new(_SysexCommandAnalogMappingResponse).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SysexCommandAnalogMappingResponse for type-switch of SysexCommand")
		}
	case commandType == 0x6B: // SysexCommandCapabilityQuery
		if _child, err = new(_SysexCommandCapabilityQuery).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SysexCommandCapabilityQuery for type-switch of SysexCommand")
		}
	case commandType == 0x6C: // SysexCommandCapabilityResponse
		if _child, err = new(_SysexCommandCapabilityResponse).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SysexCommandCapabilityResponse for type-switch of SysexCommand")
		}
	case commandType == 0x6D: // SysexCommandPinStateQuery
		if _child, err = new(_SysexCommandPinStateQuery).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SysexCommandPinStateQuery for type-switch of SysexCommand")
		}
	case commandType == 0x6E: // SysexCommandPinStateResponse
		if _child, err = new(_SysexCommandPinStateResponse).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SysexCommandPinStateResponse for type-switch of SysexCommand")
		}
	case commandType == 0x6F: // SysexCommandExtendedAnalog
		if _child, err = new(_SysexCommandExtendedAnalog).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SysexCommandExtendedAnalog for type-switch of SysexCommand")
		}
	case commandType == 0x71: // SysexCommandStringData
		if _child, err = new(_SysexCommandStringData).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SysexCommandStringData for type-switch of SysexCommand")
		}
	case commandType == 0x79 && response == bool(false): // SysexCommandReportFirmwareRequest
		if _child, err = new(_SysexCommandReportFirmwareRequest).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SysexCommandReportFirmwareRequest for type-switch of SysexCommand")
		}
	case commandType == 0x79 && response == bool(true): // SysexCommandReportFirmwareResponse
		if _child, err = new(_SysexCommandReportFirmwareResponse).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SysexCommandReportFirmwareResponse for type-switch of SysexCommand")
		}
	case commandType == 0x7A: // SysexCommandSamplingInterval
		if _child, err = new(_SysexCommandSamplingInterval).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SysexCommandSamplingInterval for type-switch of SysexCommand")
		}
	case commandType == 0x7E: // SysexCommandSysexNonRealtime
		if _child, err = new(_SysexCommandSysexNonRealtime).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SysexCommandSysexNonRealtime for type-switch of SysexCommand")
		}
	case commandType == 0x7F: // SysexCommandSysexRealtime
		if _child, err = new(_SysexCommandSysexRealtime).parse(ctx, readBuffer, m, response); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SysexCommandSysexRealtime for type-switch of SysexCommand")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [commandType=%v, response=%v]", commandType, response)
	}

	if closeErr := readBuffer.CloseContext("SysexCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommand")
	}

	return _child, nil
}

func (pm *_SysexCommand) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child SysexCommand, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("SysexCommand"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for SysexCommand")
	}

	if err := WriteDiscriminatorField(ctx, "commandType", m.GetCommandType(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'commandType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("SysexCommand"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for SysexCommand")
	}
	return nil
}

func (m *_SysexCommand) IsSysexCommand() {}

func (m *_SysexCommand) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SysexCommand) deepCopy() *_SysexCommand {
	if m == nil {
		return nil
	}
	_SysexCommandCopy := &_SysexCommand{
		nil, // will be set by child
	}
	return _SysexCommandCopy
}
