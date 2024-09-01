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

// BACnetHostNPort is the corresponding interface of BACnetHostNPort
type BACnetHostNPort interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetHost returns Host (property field)
	GetHost() BACnetHostAddressEnclosed
	// GetPort returns Port (property field)
	GetPort() BACnetContextTagUnsignedInteger
}

// BACnetHostNPortExactly can be used when we want exactly this type and not a type which fulfills BACnetHostNPort.
// This is useful for switch cases.
type BACnetHostNPortExactly interface {
	BACnetHostNPort
	isBACnetHostNPort() bool
}

// _BACnetHostNPort is the data-structure of this message
type _BACnetHostNPort struct {
	Host BACnetHostAddressEnclosed
	Port BACnetContextTagUnsignedInteger
}

var _ BACnetHostNPort = (*_BACnetHostNPort)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetHostNPort) GetHost() BACnetHostAddressEnclosed {
	return m.Host
}

func (m *_BACnetHostNPort) GetPort() BACnetContextTagUnsignedInteger {
	return m.Port
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetHostNPort factory function for _BACnetHostNPort
func NewBACnetHostNPort(host BACnetHostAddressEnclosed, port BACnetContextTagUnsignedInteger) *_BACnetHostNPort {
	return &_BACnetHostNPort{Host: host, Port: port}
}

// Deprecated: use the interface for direct cast
func CastBACnetHostNPort(structType any) BACnetHostNPort {
	if casted, ok := structType.(BACnetHostNPort); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetHostNPort); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetHostNPort) GetTypeName() string {
	return "BACnetHostNPort"
}

func (m *_BACnetHostNPort) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (host)
	lengthInBits += m.Host.GetLengthInBits(ctx)

	// Simple field (port)
	lengthInBits += m.Port.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetHostNPort) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetHostNPortParse(ctx context.Context, theBytes []byte) (BACnetHostNPort, error) {
	return BACnetHostNPortParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetHostNPortParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetHostNPort, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetHostNPort, error) {
		return BACnetHostNPortParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetHostNPortParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetHostNPort, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetHostNPort"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetHostNPort")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	host, err := ReadSimpleField[BACnetHostAddressEnclosed](ctx, "host", ReadComplex[BACnetHostAddressEnclosed](BACnetHostAddressEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'host' field"))
	}

	port, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "port", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'port' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetHostNPort"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetHostNPort")
	}

	// Create the instance
	return &_BACnetHostNPort{
		Host: host,
		Port: port,
	}, nil
}

func (m *_BACnetHostNPort) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetHostNPort) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetHostNPort"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetHostNPort")
	}

	if err := WriteSimpleField[BACnetHostAddressEnclosed](ctx, "host", m.GetHost(), WriteComplex[BACnetHostAddressEnclosed](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'host' field")
	}

	if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "port", m.GetPort(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'port' field")
	}

	if popErr := writeBuffer.PopContext("BACnetHostNPort"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetHostNPort")
	}
	return nil
}

func (m *_BACnetHostNPort) isBACnetHostNPort() bool {
	return true
}

func (m *_BACnetHostNPort) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
