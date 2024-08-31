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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetScale is the corresponding interface of BACnetScale
type BACnetScale interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetScaleExactly can be used when we want exactly this type and not a type which fulfills BACnetScale.
// This is useful for switch cases.
type BACnetScaleExactly interface {
	BACnetScale
	isBACnetScale() bool
}

// _BACnetScale is the data-structure of this message
type _BACnetScale struct {
	_BACnetScaleChildRequirements
	PeekedTagHeader BACnetTagHeader
}

type _BACnetScaleChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetPeekedTagNumber() uint8
}

type BACnetScaleParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetScale, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetScaleChild interface {
	utils.Serializable
	InitializeParent(parent BACnetScale, peekedTagHeader BACnetTagHeader)
	GetParent() *BACnetScale

	GetTypeName() string
	BACnetScale
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetScale) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetScale) GetPeekedTagNumber() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetScale factory function for _BACnetScale
func NewBACnetScale(peekedTagHeader BACnetTagHeader) *_BACnetScale {
	return &_BACnetScale{PeekedTagHeader: peekedTagHeader}
}

// Deprecated: use the interface for direct cast
func CastBACnetScale(structType any) BACnetScale {
	if casted, ok := structType.(BACnetScale); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetScale); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetScale) GetTypeName() string {
	return "BACnetScale"
}

func (m *_BACnetScale) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetScale) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetScaleParse(ctx context.Context, theBytes []byte) (BACnetScale, error) {
	return BACnetScaleParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetScaleParseWithBufferProducer[T BACnetScale]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		buffer, err := BACnetScaleParseWithBuffer(ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return buffer.(T), err
	}
}

func BACnetScaleParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetScale, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetScale"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetScale")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (peekedTagHeader)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for peekedTagHeader")
	}
	peekedTagHeader, _ := BACnetTagHeaderParseWithBuffer(ctx, readBuffer)
	readBuffer.Reset(currentPos)

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetScaleChildSerializeRequirement interface {
		BACnetScale
		InitializeParent(BACnetScale, BACnetTagHeader)
		GetParent() BACnetScale
	}
	var _childTemp any
	var _child BACnetScaleChildSerializeRequirement
	var typeSwitchError error
	switch {
	case peekedTagNumber == uint8(0): // BACnetScaleFloatScale
		_childTemp, typeSwitchError = BACnetScaleFloatScaleParseWithBuffer(ctx, readBuffer)
	case peekedTagNumber == uint8(1): // BACnetScaleIntegerScale
		_childTemp, typeSwitchError = BACnetScaleIntegerScaleParseWithBuffer(ctx, readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetScale")
	}
	_child = _childTemp.(BACnetScaleChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("BACnetScale"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetScale")
	}

	// Finish initializing
	_child.InitializeParent(_child, peekedTagHeader)
	return _child, nil
}

func (pm *_BACnetScale) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetScale, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetScale"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetScale")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetScale"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetScale")
	}
	return nil
}

func (m *_BACnetScale) isBACnetScale() bool {
	return true
}

func (m *_BACnetScale) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
