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

// BACnetServiceAckVTOpen is the corresponding interface of BACnetServiceAckVTOpen
type BACnetServiceAckVTOpen interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetServiceAck
	// GetRemoteVtSessionIdentifier returns RemoteVtSessionIdentifier (property field)
	GetRemoteVtSessionIdentifier() BACnetApplicationTagUnsignedInteger
}

// BACnetServiceAckVTOpenExactly can be used when we want exactly this type and not a type which fulfills BACnetServiceAckVTOpen.
// This is useful for switch cases.
type BACnetServiceAckVTOpenExactly interface {
	BACnetServiceAckVTOpen
	isBACnetServiceAckVTOpen() bool
}

// _BACnetServiceAckVTOpen is the data-structure of this message
type _BACnetServiceAckVTOpen struct {
	*_BACnetServiceAck
	RemoteVtSessionIdentifier BACnetApplicationTagUnsignedInteger
}

var _ BACnetServiceAckVTOpen = (*_BACnetServiceAckVTOpen)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetServiceAckVTOpen) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_VT_OPEN
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckVTOpen) InitializeParent(parent BACnetServiceAck) {}

func (m *_BACnetServiceAckVTOpen) GetParent() BACnetServiceAck {
	return m._BACnetServiceAck
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckVTOpen) GetRemoteVtSessionIdentifier() BACnetApplicationTagUnsignedInteger {
	return m.RemoteVtSessionIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckVTOpen factory function for _BACnetServiceAckVTOpen
func NewBACnetServiceAckVTOpen(remoteVtSessionIdentifier BACnetApplicationTagUnsignedInteger, serviceAckLength uint32) *_BACnetServiceAckVTOpen {
	_result := &_BACnetServiceAckVTOpen{
		RemoteVtSessionIdentifier: remoteVtSessionIdentifier,
		_BACnetServiceAck:         NewBACnetServiceAck(serviceAckLength),
	}
	_result._BACnetServiceAck._BACnetServiceAckChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckVTOpen(structType any) BACnetServiceAckVTOpen {
	if casted, ok := structType.(BACnetServiceAckVTOpen); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckVTOpen); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckVTOpen) GetTypeName() string {
	return "BACnetServiceAckVTOpen"
}

func (m *_BACnetServiceAckVTOpen) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (remoteVtSessionIdentifier)
	lengthInBits += m.RemoteVtSessionIdentifier.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetServiceAckVTOpen) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetServiceAckVTOpenParse(ctx context.Context, theBytes []byte, serviceAckLength uint32) (BACnetServiceAckVTOpen, error) {
	return BACnetServiceAckVTOpenParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), serviceAckLength)
}

func BACnetServiceAckVTOpenParseWithBufferProducer(serviceAckLength uint32) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetServiceAckVTOpen, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetServiceAckVTOpen, error) {
		return BACnetServiceAckVTOpenParseWithBuffer(ctx, readBuffer, serviceAckLength)
	}
}

func BACnetServiceAckVTOpenParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceAckLength uint32) (BACnetServiceAckVTOpen, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckVTOpen"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckVTOpen")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	remoteVtSessionIdentifier, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "remoteVtSessionIdentifier", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'remoteVtSessionIdentifier' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckVTOpen"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckVTOpen")
	}

	// Create a partially initialized instance
	_child := &_BACnetServiceAckVTOpen{
		_BACnetServiceAck: &_BACnetServiceAck{
			ServiceAckLength: serviceAckLength,
		},
		RemoteVtSessionIdentifier: remoteVtSessionIdentifier,
	}
	_child._BACnetServiceAck._BACnetServiceAckChildRequirements = _child
	return _child, nil
}

func (m *_BACnetServiceAckVTOpen) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetServiceAckVTOpen) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckVTOpen"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckVTOpen")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "remoteVtSessionIdentifier", m.GetRemoteVtSessionIdentifier(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'remoteVtSessionIdentifier' field")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckVTOpen"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckVTOpen")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetServiceAckVTOpen) isBACnetServiceAckVTOpen() bool {
	return true
}

func (m *_BACnetServiceAckVTOpen) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
