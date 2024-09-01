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

// KnxNetObjectServer is the corresponding interface of KnxNetObjectServer
type KnxNetObjectServer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ServiceId
	// GetVersion returns Version (property field)
	GetVersion() uint8
}

// KnxNetObjectServerExactly can be used when we want exactly this type and not a type which fulfills KnxNetObjectServer.
// This is useful for switch cases.
type KnxNetObjectServerExactly interface {
	KnxNetObjectServer
	isKnxNetObjectServer() bool
}

// _KnxNetObjectServer is the data-structure of this message
type _KnxNetObjectServer struct {
	*_ServiceId
	Version uint8
}

var _ KnxNetObjectServer = (*_KnxNetObjectServer)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_KnxNetObjectServer) GetServiceType() uint8 {
	return 0x08
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_KnxNetObjectServer) InitializeParent(parent ServiceId) {}

func (m *_KnxNetObjectServer) GetParent() ServiceId {
	return m._ServiceId
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_KnxNetObjectServer) GetVersion() uint8 {
	return m.Version
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewKnxNetObjectServer factory function for _KnxNetObjectServer
func NewKnxNetObjectServer(version uint8) *_KnxNetObjectServer {
	_result := &_KnxNetObjectServer{
		Version:    version,
		_ServiceId: NewServiceId(),
	}
	_result._ServiceId._ServiceIdChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastKnxNetObjectServer(structType any) KnxNetObjectServer {
	if casted, ok := structType.(KnxNetObjectServer); ok {
		return casted
	}
	if casted, ok := structType.(*KnxNetObjectServer); ok {
		return *casted
	}
	return nil
}

func (m *_KnxNetObjectServer) GetTypeName() string {
	return "KnxNetObjectServer"
}

func (m *_KnxNetObjectServer) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (version)
	lengthInBits += 8

	return lengthInBits
}

func (m *_KnxNetObjectServer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func KnxNetObjectServerParse(ctx context.Context, theBytes []byte) (KnxNetObjectServer, error) {
	return KnxNetObjectServerParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func KnxNetObjectServerParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (KnxNetObjectServer, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (KnxNetObjectServer, error) {
		return KnxNetObjectServerParseWithBuffer(ctx, readBuffer)
	}
}

func KnxNetObjectServerParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (KnxNetObjectServer, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("KnxNetObjectServer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for KnxNetObjectServer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	version, err := ReadSimpleField(ctx, "version", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'version' field"))
	}

	if closeErr := readBuffer.CloseContext("KnxNetObjectServer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for KnxNetObjectServer")
	}

	// Create a partially initialized instance
	_child := &_KnxNetObjectServer{
		_ServiceId: &_ServiceId{},
		Version:    version,
	}
	_child._ServiceId._ServiceIdChildRequirements = _child
	return _child, nil
}

func (m *_KnxNetObjectServer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_KnxNetObjectServer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("KnxNetObjectServer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for KnxNetObjectServer")
		}

		if err := WriteSimpleField[uint8](ctx, "version", m.GetVersion(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'version' field")
		}

		if popErr := writeBuffer.PopContext("KnxNetObjectServer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for KnxNetObjectServer")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_KnxNetObjectServer) isKnxNetObjectServer() bool {
	return true
}

func (m *_KnxNetObjectServer) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
