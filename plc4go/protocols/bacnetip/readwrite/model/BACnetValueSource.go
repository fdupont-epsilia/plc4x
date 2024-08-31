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

// BACnetValueSource is the corresponding interface of BACnetValueSource
type BACnetValueSource interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetValueSourceExactly can be used when we want exactly this type and not a type which fulfills BACnetValueSource.
// This is useful for switch cases.
type BACnetValueSourceExactly interface {
	BACnetValueSource
	isBACnetValueSource() bool
}

// _BACnetValueSource is the data-structure of this message
type _BACnetValueSource struct {
	_BACnetValueSourceChildRequirements
	PeekedTagHeader BACnetTagHeader
}

type _BACnetValueSourceChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetPeekedTagNumber() uint8
}

type BACnetValueSourceParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetValueSource, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetValueSourceChild interface {
	utils.Serializable
	InitializeParent(parent BACnetValueSource, peekedTagHeader BACnetTagHeader)
	GetParent() *BACnetValueSource

	GetTypeName() string
	BACnetValueSource
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetValueSource) GetPeekedTagHeader() BACnetTagHeader {
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

func (m *_BACnetValueSource) GetPeekedTagNumber() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetValueSource factory function for _BACnetValueSource
func NewBACnetValueSource(peekedTagHeader BACnetTagHeader) *_BACnetValueSource {
	return &_BACnetValueSource{PeekedTagHeader: peekedTagHeader}
}

// Deprecated: use the interface for direct cast
func CastBACnetValueSource(structType any) BACnetValueSource {
	if casted, ok := structType.(BACnetValueSource); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetValueSource); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetValueSource) GetTypeName() string {
	return "BACnetValueSource"
}

func (m *_BACnetValueSource) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetValueSource) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetValueSourceParse(ctx context.Context, theBytes []byte) (BACnetValueSource, error) {
	return BACnetValueSourceParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetValueSourceParseWithBufferProducer[T BACnetValueSource]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		buffer, err := BACnetValueSourceParseWithBuffer(ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return buffer.(T), err
	}
}

func BACnetValueSourceParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetValueSource, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetValueSource"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetValueSource")
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
	type BACnetValueSourceChildSerializeRequirement interface {
		BACnetValueSource
		InitializeParent(BACnetValueSource, BACnetTagHeader)
		GetParent() BACnetValueSource
	}
	var _childTemp any
	var _child BACnetValueSourceChildSerializeRequirement
	var typeSwitchError error
	switch {
	case peekedTagNumber == uint8(0): // BACnetValueSourceNone
		_childTemp, typeSwitchError = BACnetValueSourceNoneParseWithBuffer(ctx, readBuffer)
	case peekedTagNumber == uint8(1): // BACnetValueSourceObject
		_childTemp, typeSwitchError = BACnetValueSourceObjectParseWithBuffer(ctx, readBuffer)
	case peekedTagNumber == uint8(2): // BACnetValueSourceAddress
		_childTemp, typeSwitchError = BACnetValueSourceAddressParseWithBuffer(ctx, readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetValueSource")
	}
	_child = _childTemp.(BACnetValueSourceChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("BACnetValueSource"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetValueSource")
	}

	// Finish initializing
	_child.InitializeParent(_child, peekedTagHeader)
	return _child, nil
}

func (pm *_BACnetValueSource) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetValueSource, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetValueSource"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetValueSource")
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

	if popErr := writeBuffer.PopContext("BACnetValueSource"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetValueSource")
	}
	return nil
}

func (m *_BACnetValueSource) isBACnetValueSource() bool {
	return true
}

func (m *_BACnetValueSource) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
