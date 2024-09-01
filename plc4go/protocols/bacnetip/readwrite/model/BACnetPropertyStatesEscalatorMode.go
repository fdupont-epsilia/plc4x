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

// BACnetPropertyStatesEscalatorMode is the corresponding interface of BACnetPropertyStatesEscalatorMode
type BACnetPropertyStatesEscalatorMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetEscalatorMode returns EscalatorMode (property field)
	GetEscalatorMode() BACnetEscalatorModeTagged
}

// BACnetPropertyStatesEscalatorModeExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesEscalatorMode.
// This is useful for switch cases.
type BACnetPropertyStatesEscalatorModeExactly interface {
	BACnetPropertyStatesEscalatorMode
	isBACnetPropertyStatesEscalatorMode() bool
}

// _BACnetPropertyStatesEscalatorMode is the data-structure of this message
type _BACnetPropertyStatesEscalatorMode struct {
	*_BACnetPropertyStates
	EscalatorMode BACnetEscalatorModeTagged
}

var _ BACnetPropertyStatesEscalatorMode = (*_BACnetPropertyStatesEscalatorMode)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesEscalatorMode) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesEscalatorMode) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesEscalatorMode) GetEscalatorMode() BACnetEscalatorModeTagged {
	return m.EscalatorMode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesEscalatorMode factory function for _BACnetPropertyStatesEscalatorMode
func NewBACnetPropertyStatesEscalatorMode(escalatorMode BACnetEscalatorModeTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesEscalatorMode {
	_result := &_BACnetPropertyStatesEscalatorMode{
		EscalatorMode:         escalatorMode,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesEscalatorMode(structType any) BACnetPropertyStatesEscalatorMode {
	if casted, ok := structType.(BACnetPropertyStatesEscalatorMode); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesEscalatorMode); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesEscalatorMode) GetTypeName() string {
	return "BACnetPropertyStatesEscalatorMode"
}

func (m *_BACnetPropertyStatesEscalatorMode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (escalatorMode)
	lengthInBits += m.EscalatorMode.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesEscalatorMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyStatesEscalatorModeParse(ctx context.Context, theBytes []byte, peekedTagNumber uint8) (BACnetPropertyStatesEscalatorMode, error) {
	return BACnetPropertyStatesEscalatorModeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), peekedTagNumber)
}

func BACnetPropertyStatesEscalatorModeParseWithBufferProducer(peekedTagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPropertyStatesEscalatorMode, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPropertyStatesEscalatorMode, error) {
		return BACnetPropertyStatesEscalatorModeParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	}
}

func BACnetPropertyStatesEscalatorModeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesEscalatorMode, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesEscalatorMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesEscalatorMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	escalatorMode, err := ReadSimpleField[BACnetEscalatorModeTagged](ctx, "escalatorMode", ReadComplex[BACnetEscalatorModeTagged](BACnetEscalatorModeTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'escalatorMode' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesEscalatorMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesEscalatorMode")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesEscalatorMode{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		EscalatorMode:         escalatorMode,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesEscalatorMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesEscalatorMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesEscalatorMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesEscalatorMode")
		}

		if err := WriteSimpleField[BACnetEscalatorModeTagged](ctx, "escalatorMode", m.GetEscalatorMode(), WriteComplex[BACnetEscalatorModeTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'escalatorMode' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesEscalatorMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesEscalatorMode")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesEscalatorMode) isBACnetPropertyStatesEscalatorMode() bool {
	return true
}

func (m *_BACnetPropertyStatesEscalatorMode) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
