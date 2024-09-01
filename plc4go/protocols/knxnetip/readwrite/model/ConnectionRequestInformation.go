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
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// ConnectionRequestInformation is the corresponding interface of ConnectionRequestInformation
type ConnectionRequestInformation interface {
	ConnectionRequestInformationContract
	ConnectionRequestInformationRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// ConnectionRequestInformationContract provides a set of functions which can be overwritten by a sub struct
type ConnectionRequestInformationContract interface {
}

// ConnectionRequestInformationRequirements provides a set of functions which need to be implemented by a sub struct
type ConnectionRequestInformationRequirements interface {
	// GetConnectionType returns ConnectionType (discriminator field)
	GetConnectionType() uint8
}

// ConnectionRequestInformationExactly can be used when we want exactly this type and not a type which fulfills ConnectionRequestInformation.
// This is useful for switch cases.
type ConnectionRequestInformationExactly interface {
	ConnectionRequestInformation
	isConnectionRequestInformation() bool
}

// _ConnectionRequestInformation is the data-structure of this message
type _ConnectionRequestInformation struct {
	_ConnectionRequestInformationChildRequirements
}

var _ ConnectionRequestInformationContract = (*_ConnectionRequestInformation)(nil)

type _ConnectionRequestInformationChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetConnectionType() uint8
}

type ConnectionRequestInformationChild interface {
	utils.Serializable

	GetParent() *ConnectionRequestInformation

	GetTypeName() string
	ConnectionRequestInformation
}

// NewConnectionRequestInformation factory function for _ConnectionRequestInformation
func NewConnectionRequestInformation() *_ConnectionRequestInformation {
	return &_ConnectionRequestInformation{}
}

// Deprecated: use the interface for direct cast
func CastConnectionRequestInformation(structType any) ConnectionRequestInformation {
	if casted, ok := structType.(ConnectionRequestInformation); ok {
		return casted
	}
	if casted, ok := structType.(*ConnectionRequestInformation); ok {
		return *casted
	}
	return nil
}

func (m *_ConnectionRequestInformation) GetTypeName() string {
	return "ConnectionRequestInformation"
}

func (m *_ConnectionRequestInformation) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (structureLength)
	lengthInBits += 8
	// Discriminator Field (connectionType)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ConnectionRequestInformation) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ConnectionRequestInformationParse[T ConnectionRequestInformation](ctx context.Context, theBytes []byte) (T, error) {
	return ConnectionRequestInformationParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func ConnectionRequestInformationParseWithBufferProducer[T ConnectionRequestInformation]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := ConnectionRequestInformationParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func ConnectionRequestInformationParseWithBuffer[T ConnectionRequestInformation](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_ConnectionRequestInformation{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_ConnectionRequestInformation) parse(ctx context.Context, readBuffer utils.ReadBuffer) (_connectionRequestInformation ConnectionRequestInformation, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ConnectionRequestInformation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ConnectionRequestInformation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	structureLength, err := ReadImplicitField[uint8](ctx, "structureLength", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'structureLength' field"))
	}
	_ = structureLength

	connectionType, err := ReadDiscriminatorField[uint8](ctx, "connectionType", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'connectionType' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child ConnectionRequestInformation
	switch {
	case connectionType == 0x03: // ConnectionRequestInformationDeviceManagement
		if _child, err = (&_ConnectionRequestInformationDeviceManagement{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ConnectionRequestInformationDeviceManagement for type-switch of ConnectionRequestInformation")
		}
	case connectionType == 0x04: // ConnectionRequestInformationTunnelConnection
		if _child, err = (&_ConnectionRequestInformationTunnelConnection{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ConnectionRequestInformationTunnelConnection for type-switch of ConnectionRequestInformation")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [connectionType=%v]", connectionType)
	}

	if closeErr := readBuffer.CloseContext("ConnectionRequestInformation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ConnectionRequestInformation")
	}

	return _child, nil
}

func (pm *_ConnectionRequestInformation) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ConnectionRequestInformation, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ConnectionRequestInformation"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ConnectionRequestInformation")
	}
	structureLength := uint8(uint8(m.GetLengthInBytes(ctx)))
	if err := WriteImplicitField(ctx, "structureLength", structureLength, WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'structureLength' field")
	}

	if err := WriteDiscriminatorField(ctx, "connectionType", m.GetConnectionType(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'connectionType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ConnectionRequestInformation"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ConnectionRequestInformation")
	}
	return nil
}

func (m *_ConnectionRequestInformation) isConnectionRequestInformation() bool {
	return true
}
