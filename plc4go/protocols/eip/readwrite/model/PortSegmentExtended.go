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

// PortSegmentExtended is the corresponding interface of PortSegmentExtended
type PortSegmentExtended interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	PortSegmentType
	// GetPort returns Port (property field)
	GetPort() uint8
	// GetLinkAddressSize returns LinkAddressSize (property field)
	GetLinkAddressSize() uint8
	// GetAddress returns Address (property field)
	GetAddress() string
	// GetPaddingByte returns PaddingByte (virtual field)
	GetPaddingByte() uint8
}

// PortSegmentExtendedExactly can be used when we want exactly this type and not a type which fulfills PortSegmentExtended.
// This is useful for switch cases.
type PortSegmentExtendedExactly interface {
	PortSegmentExtended
	isPortSegmentExtended() bool
}

// _PortSegmentExtended is the data-structure of this message
type _PortSegmentExtended struct {
	PortSegmentTypeContract
	Port            uint8
	LinkAddressSize uint8
	Address         string
}

var _ PortSegmentExtended = (*_PortSegmentExtended)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_PortSegmentExtended) GetExtendedLinkAddress() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_PortSegmentExtended) GetParent() PortSegmentTypeContract {
	return m.PortSegmentTypeContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PortSegmentExtended) GetPort() uint8 {
	return m.Port
}

func (m *_PortSegmentExtended) GetLinkAddressSize() uint8 {
	return m.LinkAddressSize
}

func (m *_PortSegmentExtended) GetAddress() string {
	return m.Address
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_PortSegmentExtended) GetPaddingByte() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(uint8(m.GetLinkAddressSize()) % uint8(uint8(2)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewPortSegmentExtended factory function for _PortSegmentExtended
func NewPortSegmentExtended(port uint8, linkAddressSize uint8, address string) *_PortSegmentExtended {
	_result := &_PortSegmentExtended{
		PortSegmentTypeContract: NewPortSegmentType(),
		Port:                    port,
		LinkAddressSize:         linkAddressSize,
		Address:                 address,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastPortSegmentExtended(structType any) PortSegmentExtended {
	if casted, ok := structType.(PortSegmentExtended); ok {
		return casted
	}
	if casted, ok := structType.(*PortSegmentExtended); ok {
		return *casted
	}
	return nil
}

func (m *_PortSegmentExtended) GetTypeName() string {
	return "PortSegmentExtended"
}

func (m *_PortSegmentExtended) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.PortSegmentTypeContract.(*_PortSegmentType).getLengthInBits(ctx))

	// Simple field (port)
	lengthInBits += 4

	// Simple field (linkAddressSize)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (address)
	lengthInBits += uint16(int32((int32(m.GetLinkAddressSize()) * int32(int32(8)))) + int32((int32(m.GetPaddingByte()) * int32(int32(8)))))

	return lengthInBits
}

func (m *_PortSegmentExtended) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_PortSegmentExtended) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_PortSegmentType) (_portSegmentExtended PortSegmentExtended, err error) {
	m.PortSegmentTypeContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("PortSegmentExtended"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PortSegmentExtended")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	port, err := ReadSimpleField(ctx, "port", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'port' field"))
	}
	m.Port = port

	linkAddressSize, err := ReadSimpleField(ctx, "linkAddressSize", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'linkAddressSize' field"))
	}
	m.LinkAddressSize = linkAddressSize

	paddingByte, err := ReadVirtualField[uint8](ctx, "paddingByte", (*uint8)(nil), uint8(linkAddressSize)%uint8(uint8(2)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'paddingByte' field"))
	}
	_ = paddingByte

	address, err := ReadSimpleField(ctx, "address", ReadString(readBuffer, uint32(int32((int32(linkAddressSize)*int32(int32(8))))+int32((int32(paddingByte)*int32(int32(8)))))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'address' field"))
	}
	m.Address = address

	if closeErr := readBuffer.CloseContext("PortSegmentExtended"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PortSegmentExtended")
	}

	return m, nil
}

func (m *_PortSegmentExtended) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PortSegmentExtended) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PortSegmentExtended"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PortSegmentExtended")
		}

		if err := WriteSimpleField[uint8](ctx, "port", m.GetPort(), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'port' field")
		}

		if err := WriteSimpleField[uint8](ctx, "linkAddressSize", m.GetLinkAddressSize(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'linkAddressSize' field")
		}
		// Virtual field
		paddingByte := m.GetPaddingByte()
		_ = paddingByte
		if _paddingByteErr := writeBuffer.WriteVirtual(ctx, "paddingByte", m.GetPaddingByte()); _paddingByteErr != nil {
			return errors.Wrap(_paddingByteErr, "Error serializing 'paddingByte' field")
		}

		if err := WriteSimpleField[string](ctx, "address", m.GetAddress(), WriteString(writeBuffer, int32(int32((int32(m.GetLinkAddressSize())*int32(int32(8))))+int32((int32(m.GetPaddingByte())*int32(int32(8))))))); err != nil {
			return errors.Wrap(err, "Error serializing 'address' field")
		}

		if popErr := writeBuffer.PopContext("PortSegmentExtended"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PortSegmentExtended")
		}
		return nil
	}
	return m.PortSegmentTypeContract.(*_PortSegmentType).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_PortSegmentExtended) isPortSegmentExtended() bool {
	return true
}

func (m *_PortSegmentExtended) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
