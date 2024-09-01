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

// BACnetPropertyStatesZoneOccupanyState is the corresponding interface of BACnetPropertyStatesZoneOccupanyState
type BACnetPropertyStatesZoneOccupanyState interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetZoneOccupanyState returns ZoneOccupanyState (property field)
	GetZoneOccupanyState() BACnetAccessZoneOccupancyStateTagged
}

// BACnetPropertyStatesZoneOccupanyStateExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesZoneOccupanyState.
// This is useful for switch cases.
type BACnetPropertyStatesZoneOccupanyStateExactly interface {
	BACnetPropertyStatesZoneOccupanyState
	isBACnetPropertyStatesZoneOccupanyState() bool
}

// _BACnetPropertyStatesZoneOccupanyState is the data-structure of this message
type _BACnetPropertyStatesZoneOccupanyState struct {
	*_BACnetPropertyStates
	ZoneOccupanyState BACnetAccessZoneOccupancyStateTagged
}

var _ BACnetPropertyStatesZoneOccupanyState = (*_BACnetPropertyStatesZoneOccupanyState)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesZoneOccupanyState) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesZoneOccupanyState) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesZoneOccupanyState) GetZoneOccupanyState() BACnetAccessZoneOccupancyStateTagged {
	return m.ZoneOccupanyState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesZoneOccupanyState factory function for _BACnetPropertyStatesZoneOccupanyState
func NewBACnetPropertyStatesZoneOccupanyState(zoneOccupanyState BACnetAccessZoneOccupancyStateTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesZoneOccupanyState {
	_result := &_BACnetPropertyStatesZoneOccupanyState{
		ZoneOccupanyState:     zoneOccupanyState,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesZoneOccupanyState(structType any) BACnetPropertyStatesZoneOccupanyState {
	if casted, ok := structType.(BACnetPropertyStatesZoneOccupanyState); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesZoneOccupanyState); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesZoneOccupanyState) GetTypeName() string {
	return "BACnetPropertyStatesZoneOccupanyState"
}

func (m *_BACnetPropertyStatesZoneOccupanyState) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (zoneOccupanyState)
	lengthInBits += m.ZoneOccupanyState.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesZoneOccupanyState) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyStatesZoneOccupanyStateParse(ctx context.Context, theBytes []byte, peekedTagNumber uint8) (BACnetPropertyStatesZoneOccupanyState, error) {
	return BACnetPropertyStatesZoneOccupanyStateParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), peekedTagNumber)
}

func BACnetPropertyStatesZoneOccupanyStateParseWithBufferProducer(peekedTagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPropertyStatesZoneOccupanyState, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPropertyStatesZoneOccupanyState, error) {
		return BACnetPropertyStatesZoneOccupanyStateParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	}
}

func BACnetPropertyStatesZoneOccupanyStateParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesZoneOccupanyState, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesZoneOccupanyState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesZoneOccupanyState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zoneOccupanyState, err := ReadSimpleField[BACnetAccessZoneOccupancyStateTagged](ctx, "zoneOccupanyState", ReadComplex[BACnetAccessZoneOccupancyStateTagged](BACnetAccessZoneOccupancyStateTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneOccupanyState' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesZoneOccupanyState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesZoneOccupanyState")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesZoneOccupanyState{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		ZoneOccupanyState:     zoneOccupanyState,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesZoneOccupanyState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesZoneOccupanyState) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesZoneOccupanyState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesZoneOccupanyState")
		}

		if err := WriteSimpleField[BACnetAccessZoneOccupancyStateTagged](ctx, "zoneOccupanyState", m.GetZoneOccupanyState(), WriteComplex[BACnetAccessZoneOccupancyStateTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneOccupanyState' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesZoneOccupanyState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesZoneOccupanyState")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesZoneOccupanyState) isBACnetPropertyStatesZoneOccupanyState() bool {
	return true
}

func (m *_BACnetPropertyStatesZoneOccupanyState) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
