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

// CBusPointToPointCommandIndirect is the corresponding interface of CBusPointToPointCommandIndirect
type CBusPointToPointCommandIndirect interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CBusPointToPointCommand
	// GetBridgeAddress returns BridgeAddress (property field)
	GetBridgeAddress() BridgeAddress
	// GetNetworkRoute returns NetworkRoute (property field)
	GetNetworkRoute() NetworkRoute
	// GetUnitAddress returns UnitAddress (property field)
	GetUnitAddress() UnitAddress
}

// CBusPointToPointCommandIndirectExactly can be used when we want exactly this type and not a type which fulfills CBusPointToPointCommandIndirect.
// This is useful for switch cases.
type CBusPointToPointCommandIndirectExactly interface {
	CBusPointToPointCommandIndirect
	isCBusPointToPointCommandIndirect() bool
}

// _CBusPointToPointCommandIndirect is the data-structure of this message
type _CBusPointToPointCommandIndirect struct {
	*_CBusPointToPointCommand
	BridgeAddress BridgeAddress
	NetworkRoute  NetworkRoute
	UnitAddress   UnitAddress
}

var _ CBusPointToPointCommandIndirect = (*_CBusPointToPointCommandIndirect)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CBusPointToPointCommandIndirect) InitializeParent(parent CBusPointToPointCommand, bridgeAddressCountPeek uint16, calData CALData) {
	m.BridgeAddressCountPeek = bridgeAddressCountPeek
	m.CalData = calData
}

func (m *_CBusPointToPointCommandIndirect) GetParent() CBusPointToPointCommand {
	return m._CBusPointToPointCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusPointToPointCommandIndirect) GetBridgeAddress() BridgeAddress {
	return m.BridgeAddress
}

func (m *_CBusPointToPointCommandIndirect) GetNetworkRoute() NetworkRoute {
	return m.NetworkRoute
}

func (m *_CBusPointToPointCommandIndirect) GetUnitAddress() UnitAddress {
	return m.UnitAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCBusPointToPointCommandIndirect factory function for _CBusPointToPointCommandIndirect
func NewCBusPointToPointCommandIndirect(bridgeAddress BridgeAddress, networkRoute NetworkRoute, unitAddress UnitAddress, bridgeAddressCountPeek uint16, calData CALData, cBusOptions CBusOptions) *_CBusPointToPointCommandIndirect {
	_result := &_CBusPointToPointCommandIndirect{
		BridgeAddress:            bridgeAddress,
		NetworkRoute:             networkRoute,
		UnitAddress:              unitAddress,
		_CBusPointToPointCommand: NewCBusPointToPointCommand(bridgeAddressCountPeek, calData, cBusOptions),
	}
	_result._CBusPointToPointCommand._CBusPointToPointCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCBusPointToPointCommandIndirect(structType any) CBusPointToPointCommandIndirect {
	if casted, ok := structType.(CBusPointToPointCommandIndirect); ok {
		return casted
	}
	if casted, ok := structType.(*CBusPointToPointCommandIndirect); ok {
		return *casted
	}
	return nil
}

func (m *_CBusPointToPointCommandIndirect) GetTypeName() string {
	return "CBusPointToPointCommandIndirect"
}

func (m *_CBusPointToPointCommandIndirect) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (bridgeAddress)
	lengthInBits += m.BridgeAddress.GetLengthInBits(ctx)

	// Simple field (networkRoute)
	lengthInBits += m.NetworkRoute.GetLengthInBits(ctx)

	// Simple field (unitAddress)
	lengthInBits += m.UnitAddress.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CBusPointToPointCommandIndirect) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CBusPointToPointCommandIndirectParse(ctx context.Context, theBytes []byte, cBusOptions CBusOptions) (CBusPointToPointCommandIndirect, error) {
	return CBusPointToPointCommandIndirectParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), cBusOptions)
}

func CBusPointToPointCommandIndirectParseWithBufferProducer(cBusOptions CBusOptions) func(ctx context.Context, readBuffer utils.ReadBuffer) (CBusPointToPointCommandIndirect, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (CBusPointToPointCommandIndirect, error) {
		return CBusPointToPointCommandIndirectParseWithBuffer(ctx, readBuffer, cBusOptions)
	}
}

func CBusPointToPointCommandIndirectParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions) (CBusPointToPointCommandIndirect, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusPointToPointCommandIndirect"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusPointToPointCommandIndirect")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	bridgeAddress, err := ReadSimpleField[BridgeAddress](ctx, "bridgeAddress", ReadComplex[BridgeAddress](BridgeAddressParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bridgeAddress' field"))
	}

	networkRoute, err := ReadSimpleField[NetworkRoute](ctx, "networkRoute", ReadComplex[NetworkRoute](NetworkRouteParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'networkRoute' field"))
	}

	unitAddress, err := ReadSimpleField[UnitAddress](ctx, "unitAddress", ReadComplex[UnitAddress](UnitAddressParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unitAddress' field"))
	}

	if closeErr := readBuffer.CloseContext("CBusPointToPointCommandIndirect"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusPointToPointCommandIndirect")
	}

	// Create a partially initialized instance
	_child := &_CBusPointToPointCommandIndirect{
		_CBusPointToPointCommand: &_CBusPointToPointCommand{
			CBusOptions: cBusOptions,
		},
		BridgeAddress: bridgeAddress,
		NetworkRoute:  networkRoute,
		UnitAddress:   unitAddress,
	}
	_child._CBusPointToPointCommand._CBusPointToPointCommandChildRequirements = _child
	return _child, nil
}

func (m *_CBusPointToPointCommandIndirect) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CBusPointToPointCommandIndirect) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CBusPointToPointCommandIndirect"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CBusPointToPointCommandIndirect")
		}

		if err := WriteSimpleField[BridgeAddress](ctx, "bridgeAddress", m.GetBridgeAddress(), WriteComplex[BridgeAddress](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'bridgeAddress' field")
		}

		if err := WriteSimpleField[NetworkRoute](ctx, "networkRoute", m.GetNetworkRoute(), WriteComplex[NetworkRoute](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'networkRoute' field")
		}

		if err := WriteSimpleField[UnitAddress](ctx, "unitAddress", m.GetUnitAddress(), WriteComplex[UnitAddress](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'unitAddress' field")
		}

		if popErr := writeBuffer.PopContext("CBusPointToPointCommandIndirect"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CBusPointToPointCommandIndirect")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CBusPointToPointCommandIndirect) isCBusPointToPointCommandIndirect() bool {
	return true
}

func (m *_CBusPointToPointCommandIndirect) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
