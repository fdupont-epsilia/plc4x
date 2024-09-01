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

// NLMRejectMessageToNetwork is the corresponding interface of NLMRejectMessageToNetwork
type NLMRejectMessageToNetwork interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NLM
	// GetRejectReason returns RejectReason (property field)
	GetRejectReason() NLMRejectMessageToNetworkRejectReason
	// GetDestinationNetworkAddress returns DestinationNetworkAddress (property field)
	GetDestinationNetworkAddress() uint16
}

// NLMRejectMessageToNetworkExactly can be used when we want exactly this type and not a type which fulfills NLMRejectMessageToNetwork.
// This is useful for switch cases.
type NLMRejectMessageToNetworkExactly interface {
	NLMRejectMessageToNetwork
	isNLMRejectMessageToNetwork() bool
}

// _NLMRejectMessageToNetwork is the data-structure of this message
type _NLMRejectMessageToNetwork struct {
	*_NLM
	RejectReason              NLMRejectMessageToNetworkRejectReason
	DestinationNetworkAddress uint16
}

var _ NLMRejectMessageToNetwork = (*_NLMRejectMessageToNetwork)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMRejectMessageToNetwork) GetMessageType() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMRejectMessageToNetwork) InitializeParent(parent NLM) {}

func (m *_NLMRejectMessageToNetwork) GetParent() NLM {
	return m._NLM
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMRejectMessageToNetwork) GetRejectReason() NLMRejectMessageToNetworkRejectReason {
	return m.RejectReason
}

func (m *_NLMRejectMessageToNetwork) GetDestinationNetworkAddress() uint16 {
	return m.DestinationNetworkAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLMRejectMessageToNetwork factory function for _NLMRejectMessageToNetwork
func NewNLMRejectMessageToNetwork(rejectReason NLMRejectMessageToNetworkRejectReason, destinationNetworkAddress uint16, apduLength uint16) *_NLMRejectMessageToNetwork {
	_result := &_NLMRejectMessageToNetwork{
		RejectReason:              rejectReason,
		DestinationNetworkAddress: destinationNetworkAddress,
		_NLM:                      NewNLM(apduLength),
	}
	_result._NLM._NLMChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNLMRejectMessageToNetwork(structType any) NLMRejectMessageToNetwork {
	if casted, ok := structType.(NLMRejectMessageToNetwork); ok {
		return casted
	}
	if casted, ok := structType.(*NLMRejectMessageToNetwork); ok {
		return *casted
	}
	return nil
}

func (m *_NLMRejectMessageToNetwork) GetTypeName() string {
	return "NLMRejectMessageToNetwork"
}

func (m *_NLMRejectMessageToNetwork) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (rejectReason)
	lengthInBits += 8

	// Simple field (destinationNetworkAddress)
	lengthInBits += 16

	return lengthInBits
}

func (m *_NLMRejectMessageToNetwork) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NLMRejectMessageToNetworkParse(ctx context.Context, theBytes []byte, apduLength uint16) (NLMRejectMessageToNetwork, error) {
	return NLMRejectMessageToNetworkParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), apduLength)
}

func NLMRejectMessageToNetworkParseWithBufferProducer(apduLength uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (NLMRejectMessageToNetwork, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (NLMRejectMessageToNetwork, error) {
		return NLMRejectMessageToNetworkParseWithBuffer(ctx, readBuffer, apduLength)
	}
}

func NLMRejectMessageToNetworkParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, apduLength uint16) (NLMRejectMessageToNetwork, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMRejectMessageToNetwork"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMRejectMessageToNetwork")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	rejectReason, err := ReadEnumField[NLMRejectMessageToNetworkRejectReason](ctx, "rejectReason", "NLMRejectMessageToNetworkRejectReason", ReadEnum(NLMRejectMessageToNetworkRejectReasonByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'rejectReason' field"))
	}

	destinationNetworkAddress, err := ReadSimpleField(ctx, "destinationNetworkAddress", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'destinationNetworkAddress' field"))
	}

	if closeErr := readBuffer.CloseContext("NLMRejectMessageToNetwork"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMRejectMessageToNetwork")
	}

	// Create a partially initialized instance
	_child := &_NLMRejectMessageToNetwork{
		_NLM: &_NLM{
			ApduLength: apduLength,
		},
		RejectReason:              rejectReason,
		DestinationNetworkAddress: destinationNetworkAddress,
	}
	_child._NLM._NLMChildRequirements = _child
	return _child, nil
}

func (m *_NLMRejectMessageToNetwork) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMRejectMessageToNetwork) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMRejectMessageToNetwork"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMRejectMessageToNetwork")
		}

		if err := WriteSimpleEnumField[NLMRejectMessageToNetworkRejectReason](ctx, "rejectReason", "NLMRejectMessageToNetworkRejectReason", m.GetRejectReason(), WriteEnum[NLMRejectMessageToNetworkRejectReason, uint8](NLMRejectMessageToNetworkRejectReason.GetValue, NLMRejectMessageToNetworkRejectReason.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'rejectReason' field")
		}

		if err := WriteSimpleField[uint16](ctx, "destinationNetworkAddress", m.GetDestinationNetworkAddress(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'destinationNetworkAddress' field")
		}

		if popErr := writeBuffer.PopContext("NLMRejectMessageToNetwork"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMRejectMessageToNetwork")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NLMRejectMessageToNetwork) isNLMRejectMessageToNetwork() bool {
	return true
}

func (m *_NLMRejectMessageToNetwork) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
