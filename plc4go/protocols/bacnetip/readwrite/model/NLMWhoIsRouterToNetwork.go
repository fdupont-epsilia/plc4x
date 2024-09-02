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

// NLMWhoIsRouterToNetwork is the corresponding interface of NLMWhoIsRouterToNetwork
type NLMWhoIsRouterToNetwork interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NLM
	// GetDestinationNetworkAddress returns DestinationNetworkAddress (property field)
	GetDestinationNetworkAddress() *uint16
	// IsNLMWhoIsRouterToNetwork is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNLMWhoIsRouterToNetwork()
}

// _NLMWhoIsRouterToNetwork is the data-structure of this message
type _NLMWhoIsRouterToNetwork struct {
	NLMContract
	DestinationNetworkAddress *uint16
}

var _ NLMWhoIsRouterToNetwork = (*_NLMWhoIsRouterToNetwork)(nil)
var _ NLMRequirements = (*_NLMWhoIsRouterToNetwork)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMWhoIsRouterToNetwork) GetMessageType() uint8 {
	return 0x00
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMWhoIsRouterToNetwork) GetParent() NLMContract {
	return m.NLMContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMWhoIsRouterToNetwork) GetDestinationNetworkAddress() *uint16 {
	return m.DestinationNetworkAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLMWhoIsRouterToNetwork factory function for _NLMWhoIsRouterToNetwork
func NewNLMWhoIsRouterToNetwork(destinationNetworkAddress *uint16, apduLength uint16) *_NLMWhoIsRouterToNetwork {
	_result := &_NLMWhoIsRouterToNetwork{
		NLMContract:               NewNLM(apduLength),
		DestinationNetworkAddress: destinationNetworkAddress,
	}
	_result.NLMContract.(*_NLM)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNLMWhoIsRouterToNetwork(structType any) NLMWhoIsRouterToNetwork {
	if casted, ok := structType.(NLMWhoIsRouterToNetwork); ok {
		return casted
	}
	if casted, ok := structType.(*NLMWhoIsRouterToNetwork); ok {
		return *casted
	}
	return nil
}

func (m *_NLMWhoIsRouterToNetwork) GetTypeName() string {
	return "NLMWhoIsRouterToNetwork"
}

func (m *_NLMWhoIsRouterToNetwork) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.NLMContract.(*_NLM).getLengthInBits(ctx))

	// Optional Field (destinationNetworkAddress)
	if m.DestinationNetworkAddress != nil {
		lengthInBits += 16
	}

	return lengthInBits
}

func (m *_NLMWhoIsRouterToNetwork) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NLMWhoIsRouterToNetwork) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_NLM, apduLength uint16) (__nLMWhoIsRouterToNetwork NLMWhoIsRouterToNetwork, err error) {
	m.NLMContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMWhoIsRouterToNetwork"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMWhoIsRouterToNetwork")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var destinationNetworkAddress *uint16
	destinationNetworkAddress, err = ReadOptionalField[uint16](ctx, "destinationNetworkAddress", ReadUnsignedShort(readBuffer, uint8(16)), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'destinationNetworkAddress' field"))
	}
	m.DestinationNetworkAddress = destinationNetworkAddress

	if closeErr := readBuffer.CloseContext("NLMWhoIsRouterToNetwork"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMWhoIsRouterToNetwork")
	}

	return m, nil
}

func (m *_NLMWhoIsRouterToNetwork) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMWhoIsRouterToNetwork) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMWhoIsRouterToNetwork"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMWhoIsRouterToNetwork")
		}

		if err := WriteOptionalField[uint16](ctx, "destinationNetworkAddress", m.GetDestinationNetworkAddress(), WriteUnsignedShort(writeBuffer, 16), true); err != nil {
			return errors.Wrap(err, "Error serializing 'destinationNetworkAddress' field")
		}

		if popErr := writeBuffer.PopContext("NLMWhoIsRouterToNetwork"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMWhoIsRouterToNetwork")
		}
		return nil
	}
	return m.NLMContract.(*_NLM).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NLMWhoIsRouterToNetwork) IsNLMWhoIsRouterToNetwork() {}

func (m *_NLMWhoIsRouterToNetwork) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
