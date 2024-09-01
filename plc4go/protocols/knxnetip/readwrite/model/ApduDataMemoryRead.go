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

// ApduDataMemoryRead is the corresponding interface of ApduDataMemoryRead
type ApduDataMemoryRead interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ApduData
	// GetNumBytes returns NumBytes (property field)
	GetNumBytes() uint8
	// GetAddress returns Address (property field)
	GetAddress() uint16
}

// ApduDataMemoryReadExactly can be used when we want exactly this type and not a type which fulfills ApduDataMemoryRead.
// This is useful for switch cases.
type ApduDataMemoryReadExactly interface {
	ApduDataMemoryRead
	isApduDataMemoryRead() bool
}

// _ApduDataMemoryRead is the data-structure of this message
type _ApduDataMemoryRead struct {
	*_ApduData
	NumBytes uint8
	Address  uint16
}

var _ ApduDataMemoryRead = (*_ApduDataMemoryRead)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataMemoryRead) GetApciType() uint8 {
	return 0x8
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataMemoryRead) InitializeParent(parent ApduData) {}

func (m *_ApduDataMemoryRead) GetParent() ApduData {
	return m._ApduData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ApduDataMemoryRead) GetNumBytes() uint8 {
	return m.NumBytes
}

func (m *_ApduDataMemoryRead) GetAddress() uint16 {
	return m.Address
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewApduDataMemoryRead factory function for _ApduDataMemoryRead
func NewApduDataMemoryRead(numBytes uint8, address uint16, dataLength uint8) *_ApduDataMemoryRead {
	_result := &_ApduDataMemoryRead{
		NumBytes:  numBytes,
		Address:   address,
		_ApduData: NewApduData(dataLength),
	}
	_result._ApduData._ApduDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataMemoryRead(structType any) ApduDataMemoryRead {
	if casted, ok := structType.(ApduDataMemoryRead); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataMemoryRead); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataMemoryRead) GetTypeName() string {
	return "ApduDataMemoryRead"
}

func (m *_ApduDataMemoryRead) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (numBytes)
	lengthInBits += 6

	// Simple field (address)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ApduDataMemoryRead) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ApduDataMemoryReadParse(ctx context.Context, theBytes []byte, dataLength uint8) (ApduDataMemoryRead, error) {
	return ApduDataMemoryReadParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), dataLength)
}

func ApduDataMemoryReadParseWithBufferProducer(dataLength uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (ApduDataMemoryRead, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ApduDataMemoryRead, error) {
		return ApduDataMemoryReadParseWithBuffer(ctx, readBuffer, dataLength)
	}
}

func ApduDataMemoryReadParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, dataLength uint8) (ApduDataMemoryRead, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataMemoryRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataMemoryRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	numBytes, err := ReadSimpleField(ctx, "numBytes", ReadUnsignedByte(readBuffer, uint8(6)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numBytes' field"))
	}

	address, err := ReadSimpleField(ctx, "address", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'address' field"))
	}

	if closeErr := readBuffer.CloseContext("ApduDataMemoryRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataMemoryRead")
	}

	// Create a partially initialized instance
	_child := &_ApduDataMemoryRead{
		_ApduData: &_ApduData{
			DataLength: dataLength,
		},
		NumBytes: numBytes,
		Address:  address,
	}
	_child._ApduData._ApduDataChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataMemoryRead) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataMemoryRead) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataMemoryRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataMemoryRead")
		}

		if err := WriteSimpleField[uint8](ctx, "numBytes", m.GetNumBytes(), WriteUnsignedByte(writeBuffer, 6)); err != nil {
			return errors.Wrap(err, "Error serializing 'numBytes' field")
		}

		if err := WriteSimpleField[uint16](ctx, "address", m.GetAddress(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'address' field")
		}

		if popErr := writeBuffer.PopContext("ApduDataMemoryRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataMemoryRead")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataMemoryRead) isApduDataMemoryRead() bool {
	return true
}

func (m *_ApduDataMemoryRead) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
