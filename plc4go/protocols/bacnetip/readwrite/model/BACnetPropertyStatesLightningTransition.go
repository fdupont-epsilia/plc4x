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

// BACnetPropertyStatesLightningTransition is the corresponding interface of BACnetPropertyStatesLightningTransition
type BACnetPropertyStatesLightningTransition interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetLightningTransition returns LightningTransition (property field)
	GetLightningTransition() BACnetLightingTransitionTagged
	// IsBACnetPropertyStatesLightningTransition is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesLightningTransition()
}

// _BACnetPropertyStatesLightningTransition is the data-structure of this message
type _BACnetPropertyStatesLightningTransition struct {
	BACnetPropertyStatesContract
	LightningTransition BACnetLightingTransitionTagged
}

var _ BACnetPropertyStatesLightningTransition = (*_BACnetPropertyStatesLightningTransition)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesLightningTransition)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesLightningTransition) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesLightningTransition) GetLightningTransition() BACnetLightingTransitionTagged {
	return m.LightningTransition
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesLightningTransition factory function for _BACnetPropertyStatesLightningTransition
func NewBACnetPropertyStatesLightningTransition(lightningTransition BACnetLightingTransitionTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesLightningTransition {
	_result := &_BACnetPropertyStatesLightningTransition{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		LightningTransition:          lightningTransition,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesLightningTransition(structType any) BACnetPropertyStatesLightningTransition {
	if casted, ok := structType.(BACnetPropertyStatesLightningTransition); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesLightningTransition); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesLightningTransition) GetTypeName() string {
	return "BACnetPropertyStatesLightningTransition"
}

func (m *_BACnetPropertyStatesLightningTransition) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (lightningTransition)
	lengthInBits += m.LightningTransition.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesLightningTransition) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesLightningTransition) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesLightningTransition BACnetPropertyStatesLightningTransition, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesLightningTransition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesLightningTransition")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lightningTransition, err := ReadSimpleField[BACnetLightingTransitionTagged](ctx, "lightningTransition", ReadComplex[BACnetLightingTransitionTagged](BACnetLightingTransitionTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lightningTransition' field"))
	}
	m.LightningTransition = lightningTransition

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesLightningTransition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesLightningTransition")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesLightningTransition) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesLightningTransition) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesLightningTransition"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesLightningTransition")
		}

		if err := WriteSimpleField[BACnetLightingTransitionTagged](ctx, "lightningTransition", m.GetLightningTransition(), WriteComplex[BACnetLightingTransitionTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lightningTransition' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesLightningTransition"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesLightningTransition")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesLightningTransition) IsBACnetPropertyStatesLightningTransition() {}

func (m *_BACnetPropertyStatesLightningTransition) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
