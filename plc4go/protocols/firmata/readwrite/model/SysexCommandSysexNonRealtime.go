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

// SysexCommandSysexNonRealtime is the corresponding interface of SysexCommandSysexNonRealtime
type SysexCommandSysexNonRealtime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SysexCommand
}

// SysexCommandSysexNonRealtimeExactly can be used when we want exactly this type and not a type which fulfills SysexCommandSysexNonRealtime.
// This is useful for switch cases.
type SysexCommandSysexNonRealtimeExactly interface {
	SysexCommandSysexNonRealtime
	isSysexCommandSysexNonRealtime() bool
}

// _SysexCommandSysexNonRealtime is the data-structure of this message
type _SysexCommandSysexNonRealtime struct {
	*_SysexCommand
}

var _ SysexCommandSysexNonRealtime = (*_SysexCommandSysexNonRealtime)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SysexCommandSysexNonRealtime) GetCommandType() uint8 {
	return 0x7E
}

func (m *_SysexCommandSysexNonRealtime) GetResponse() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SysexCommandSysexNonRealtime) InitializeParent(parent SysexCommand) {}

func (m *_SysexCommandSysexNonRealtime) GetParent() SysexCommand {
	return m._SysexCommand
}

// NewSysexCommandSysexNonRealtime factory function for _SysexCommandSysexNonRealtime
func NewSysexCommandSysexNonRealtime() *_SysexCommandSysexNonRealtime {
	_result := &_SysexCommandSysexNonRealtime{
		_SysexCommand: NewSysexCommand(),
	}
	_result._SysexCommand._SysexCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSysexCommandSysexNonRealtime(structType any) SysexCommandSysexNonRealtime {
	if casted, ok := structType.(SysexCommandSysexNonRealtime); ok {
		return casted
	}
	if casted, ok := structType.(*SysexCommandSysexNonRealtime); ok {
		return *casted
	}
	return nil
}

func (m *_SysexCommandSysexNonRealtime) GetTypeName() string {
	return "SysexCommandSysexNonRealtime"
}

func (m *_SysexCommandSysexNonRealtime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_SysexCommandSysexNonRealtime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SysexCommandSysexNonRealtimeParse(ctx context.Context, theBytes []byte, response bool) (SysexCommandSysexNonRealtime, error) {
	return SysexCommandSysexNonRealtimeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func SysexCommandSysexNonRealtimeParseWithBufferProducer(response bool) func(ctx context.Context, readBuffer utils.ReadBuffer) (SysexCommandSysexNonRealtime, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (SysexCommandSysexNonRealtime, error) {
		return SysexCommandSysexNonRealtimeParseWithBuffer(ctx, readBuffer, response)
	}
}

func SysexCommandSysexNonRealtimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (SysexCommandSysexNonRealtime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommandSysexNonRealtime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandSysexNonRealtime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SysexCommandSysexNonRealtime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandSysexNonRealtime")
	}

	// Create a partially initialized instance
	_child := &_SysexCommandSysexNonRealtime{
		_SysexCommand: &_SysexCommand{},
	}
	_child._SysexCommand._SysexCommandChildRequirements = _child
	return _child, nil
}

func (m *_SysexCommandSysexNonRealtime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SysexCommandSysexNonRealtime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandSysexNonRealtime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandSysexNonRealtime")
		}

		if popErr := writeBuffer.PopContext("SysexCommandSysexNonRealtime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandSysexNonRealtime")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SysexCommandSysexNonRealtime) isSysexCommandSysexNonRealtime() bool {
	return true
}

func (m *_SysexCommandSysexNonRealtime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
