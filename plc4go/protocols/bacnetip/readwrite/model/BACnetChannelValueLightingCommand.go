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

// BACnetChannelValueLightingCommand is the corresponding interface of BACnetChannelValueLightingCommand
type BACnetChannelValueLightingCommand interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetChannelValue
	// GetLigthingCommandValue returns LigthingCommandValue (property field)
	GetLigthingCommandValue() BACnetLightingCommandEnclosed
	// IsBACnetChannelValueLightingCommand is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetChannelValueLightingCommand()
}

// _BACnetChannelValueLightingCommand is the data-structure of this message
type _BACnetChannelValueLightingCommand struct {
	BACnetChannelValueContract
	LigthingCommandValue BACnetLightingCommandEnclosed
}

var _ BACnetChannelValueLightingCommand = (*_BACnetChannelValueLightingCommand)(nil)
var _ BACnetChannelValueRequirements = (*_BACnetChannelValueLightingCommand)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetChannelValueLightingCommand) GetParent() BACnetChannelValueContract {
	return m.BACnetChannelValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetChannelValueLightingCommand) GetLigthingCommandValue() BACnetLightingCommandEnclosed {
	return m.LigthingCommandValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetChannelValueLightingCommand factory function for _BACnetChannelValueLightingCommand
func NewBACnetChannelValueLightingCommand(ligthingCommandValue BACnetLightingCommandEnclosed, peekedTagHeader BACnetTagHeader) *_BACnetChannelValueLightingCommand {
	_result := &_BACnetChannelValueLightingCommand{
		BACnetChannelValueContract: NewBACnetChannelValue(peekedTagHeader),
		LigthingCommandValue:       ligthingCommandValue,
	}
	_result.BACnetChannelValueContract.(*_BACnetChannelValue)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetChannelValueLightingCommand(structType any) BACnetChannelValueLightingCommand {
	if casted, ok := structType.(BACnetChannelValueLightingCommand); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetChannelValueLightingCommand); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetChannelValueLightingCommand) GetTypeName() string {
	return "BACnetChannelValueLightingCommand"
}

func (m *_BACnetChannelValueLightingCommand) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetChannelValueContract.(*_BACnetChannelValue).getLengthInBits(ctx))

	// Simple field (ligthingCommandValue)
	lengthInBits += m.LigthingCommandValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetChannelValueLightingCommand) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetChannelValueLightingCommand) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetChannelValue) (__bACnetChannelValueLightingCommand BACnetChannelValueLightingCommand, err error) {
	m.BACnetChannelValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValueLightingCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetChannelValueLightingCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	ligthingCommandValue, err := ReadSimpleField[BACnetLightingCommandEnclosed](ctx, "ligthingCommandValue", ReadComplex[BACnetLightingCommandEnclosed](BACnetLightingCommandEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ligthingCommandValue' field"))
	}
	m.LigthingCommandValue = ligthingCommandValue

	if closeErr := readBuffer.CloseContext("BACnetChannelValueLightingCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetChannelValueLightingCommand")
	}

	return m, nil
}

func (m *_BACnetChannelValueLightingCommand) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetChannelValueLightingCommand) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueLightingCommand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetChannelValueLightingCommand")
		}

		if err := WriteSimpleField[BACnetLightingCommandEnclosed](ctx, "ligthingCommandValue", m.GetLigthingCommandValue(), WriteComplex[BACnetLightingCommandEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'ligthingCommandValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetChannelValueLightingCommand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetChannelValueLightingCommand")
		}
		return nil
	}
	return m.BACnetChannelValueContract.(*_BACnetChannelValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetChannelValueLightingCommand) IsBACnetChannelValueLightingCommand() {}

func (m *_BACnetChannelValueLightingCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
