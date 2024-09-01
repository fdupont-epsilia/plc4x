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

// BACnetPropertyStatesLifeSafetyOperations is the corresponding interface of BACnetPropertyStatesLifeSafetyOperations
type BACnetPropertyStatesLifeSafetyOperations interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetLifeSafetyOperations returns LifeSafetyOperations (property field)
	GetLifeSafetyOperations() BACnetLifeSafetyOperationTagged
}

// BACnetPropertyStatesLifeSafetyOperationsExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesLifeSafetyOperations.
// This is useful for switch cases.
type BACnetPropertyStatesLifeSafetyOperationsExactly interface {
	BACnetPropertyStatesLifeSafetyOperations
	isBACnetPropertyStatesLifeSafetyOperations() bool
}

// _BACnetPropertyStatesLifeSafetyOperations is the data-structure of this message
type _BACnetPropertyStatesLifeSafetyOperations struct {
	*_BACnetPropertyStates
	LifeSafetyOperations BACnetLifeSafetyOperationTagged
}

var _ BACnetPropertyStatesLifeSafetyOperations = (*_BACnetPropertyStatesLifeSafetyOperations)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesLifeSafetyOperations) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesLifeSafetyOperations) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesLifeSafetyOperations) GetLifeSafetyOperations() BACnetLifeSafetyOperationTagged {
	return m.LifeSafetyOperations
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesLifeSafetyOperations factory function for _BACnetPropertyStatesLifeSafetyOperations
func NewBACnetPropertyStatesLifeSafetyOperations(lifeSafetyOperations BACnetLifeSafetyOperationTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesLifeSafetyOperations {
	_result := &_BACnetPropertyStatesLifeSafetyOperations{
		LifeSafetyOperations:  lifeSafetyOperations,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesLifeSafetyOperations(structType any) BACnetPropertyStatesLifeSafetyOperations {
	if casted, ok := structType.(BACnetPropertyStatesLifeSafetyOperations); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesLifeSafetyOperations); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesLifeSafetyOperations) GetTypeName() string {
	return "BACnetPropertyStatesLifeSafetyOperations"
}

func (m *_BACnetPropertyStatesLifeSafetyOperations) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (lifeSafetyOperations)
	lengthInBits += m.LifeSafetyOperations.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesLifeSafetyOperations) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyStatesLifeSafetyOperationsParse(ctx context.Context, theBytes []byte, peekedTagNumber uint8) (BACnetPropertyStatesLifeSafetyOperations, error) {
	return BACnetPropertyStatesLifeSafetyOperationsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), peekedTagNumber)
}

func BACnetPropertyStatesLifeSafetyOperationsParseWithBufferProducer(peekedTagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPropertyStatesLifeSafetyOperations, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPropertyStatesLifeSafetyOperations, error) {
		return BACnetPropertyStatesLifeSafetyOperationsParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	}
}

func BACnetPropertyStatesLifeSafetyOperationsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesLifeSafetyOperations, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesLifeSafetyOperations"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesLifeSafetyOperations")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lifeSafetyOperations, err := ReadSimpleField[BACnetLifeSafetyOperationTagged](ctx, "lifeSafetyOperations", ReadComplex[BACnetLifeSafetyOperationTagged](BACnetLifeSafetyOperationTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lifeSafetyOperations' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesLifeSafetyOperations"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesLifeSafetyOperations")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesLifeSafetyOperations{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		LifeSafetyOperations:  lifeSafetyOperations,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesLifeSafetyOperations) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesLifeSafetyOperations) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesLifeSafetyOperations"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesLifeSafetyOperations")
		}

		if err := WriteSimpleField[BACnetLifeSafetyOperationTagged](ctx, "lifeSafetyOperations", m.GetLifeSafetyOperations(), WriteComplex[BACnetLifeSafetyOperationTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lifeSafetyOperations' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesLifeSafetyOperations"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesLifeSafetyOperations")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesLifeSafetyOperations) isBACnetPropertyStatesLifeSafetyOperations() bool {
	return true
}

func (m *_BACnetPropertyStatesLifeSafetyOperations) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
