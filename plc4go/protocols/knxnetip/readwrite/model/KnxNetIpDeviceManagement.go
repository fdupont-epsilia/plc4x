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

// KnxNetIpDeviceManagement is the corresponding interface of KnxNetIpDeviceManagement
type KnxNetIpDeviceManagement interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ServiceId
	// GetVersion returns Version (property field)
	GetVersion() uint8
}

// KnxNetIpDeviceManagementExactly can be used when we want exactly this type and not a type which fulfills KnxNetIpDeviceManagement.
// This is useful for switch cases.
type KnxNetIpDeviceManagementExactly interface {
	KnxNetIpDeviceManagement
	isKnxNetIpDeviceManagement() bool
}

// _KnxNetIpDeviceManagement is the data-structure of this message
type _KnxNetIpDeviceManagement struct {
	ServiceIdContract
	Version uint8
}

var _ KnxNetIpDeviceManagement = (*_KnxNetIpDeviceManagement)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_KnxNetIpDeviceManagement) GetServiceType() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_KnxNetIpDeviceManagement) GetParent() ServiceIdContract {
	return m.ServiceIdContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_KnxNetIpDeviceManagement) GetVersion() uint8 {
	return m.Version
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewKnxNetIpDeviceManagement factory function for _KnxNetIpDeviceManagement
func NewKnxNetIpDeviceManagement(version uint8) *_KnxNetIpDeviceManagement {
	_result := &_KnxNetIpDeviceManagement{
		ServiceIdContract: NewServiceId(),
		Version:           version,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastKnxNetIpDeviceManagement(structType any) KnxNetIpDeviceManagement {
	if casted, ok := structType.(KnxNetIpDeviceManagement); ok {
		return casted
	}
	if casted, ok := structType.(*KnxNetIpDeviceManagement); ok {
		return *casted
	}
	return nil
}

func (m *_KnxNetIpDeviceManagement) GetTypeName() string {
	return "KnxNetIpDeviceManagement"
}

func (m *_KnxNetIpDeviceManagement) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ServiceIdContract.(*_ServiceId).getLengthInBits(ctx))

	// Simple field (version)
	lengthInBits += 8

	return lengthInBits
}

func (m *_KnxNetIpDeviceManagement) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_KnxNetIpDeviceManagement) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ServiceId) (_knxNetIpDeviceManagement KnxNetIpDeviceManagement, err error) {
	m.ServiceIdContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("KnxNetIpDeviceManagement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for KnxNetIpDeviceManagement")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	version, err := ReadSimpleField(ctx, "version", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'version' field"))
	}
	m.Version = version

	if closeErr := readBuffer.CloseContext("KnxNetIpDeviceManagement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for KnxNetIpDeviceManagement")
	}

	return m, nil
}

func (m *_KnxNetIpDeviceManagement) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_KnxNetIpDeviceManagement) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("KnxNetIpDeviceManagement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for KnxNetIpDeviceManagement")
		}

		if err := WriteSimpleField[uint8](ctx, "version", m.GetVersion(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'version' field")
		}

		if popErr := writeBuffer.PopContext("KnxNetIpDeviceManagement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for KnxNetIpDeviceManagement")
		}
		return nil
	}
	return m.ServiceIdContract.(*_ServiceId).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_KnxNetIpDeviceManagement) isKnxNetIpDeviceManagement() bool {
	return true
}

func (m *_KnxNetIpDeviceManagement) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
