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

// BACnetClientCOV is the corresponding interface of BACnetClientCOV
type BACnetClientCOV interface {
	BACnetClientCOVContract
	BACnetClientCOVRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// IsBACnetClientCOV is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetClientCOV()
}

// BACnetClientCOVContract provides a set of functions which can be overwritten by a sub struct
type BACnetClientCOVContract interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// IsBACnetClientCOV is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetClientCOV()
}

// BACnetClientCOVRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetClientCOVRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// _BACnetClientCOV is the data-structure of this message
type _BACnetClientCOV struct {
	_SubType        BACnetClientCOV
	PeekedTagHeader BACnetTagHeader
}

var _ BACnetClientCOVContract = (*_BACnetClientCOV)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetClientCOV) GetPeekedTagHeader() BACnetTagHeader {
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

func (pm *_BACnetClientCOV) GetPeekedTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetClientCOV factory function for _BACnetClientCOV
func NewBACnetClientCOV(peekedTagHeader BACnetTagHeader) *_BACnetClientCOV {
	return &_BACnetClientCOV{PeekedTagHeader: peekedTagHeader}
}

// Deprecated: use the interface for direct cast
func CastBACnetClientCOV(structType any) BACnetClientCOV {
	if casted, ok := structType.(BACnetClientCOV); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetClientCOV); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetClientCOV) GetTypeName() string {
	return "BACnetClientCOV"
}

func (m *_BACnetClientCOV) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetClientCOV) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetClientCOVParse[T BACnetClientCOV](ctx context.Context, theBytes []byte) (T, error) {
	return BACnetClientCOVParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetClientCOVParseWithBufferProducer[T BACnetClientCOV]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetClientCOVParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func BACnetClientCOVParseWithBuffer[T BACnetClientCOV](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_BACnetClientCOV{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_BACnetClientCOV) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetClientCOV BACnetClientCOV, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetClientCOV"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetClientCOV")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	peekedTagHeader, err := ReadPeekField[BACnetTagHeader](ctx, "peekedTagHeader", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagHeader' field"))
	}
	m.PeekedTagHeader = peekedTagHeader

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetClientCOV
	switch {
	case peekedTagNumber == 0x4: // BACnetClientCOVObject
		if _child, err = (&_BACnetClientCOVObject{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetClientCOVObject for type-switch of BACnetClientCOV")
		}
	case peekedTagNumber == 0x0: // BACnetClientCOVNone
		if _child, err = (&_BACnetClientCOVNone{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetClientCOVNone for type-switch of BACnetClientCOV")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	if closeErr := readBuffer.CloseContext("BACnetClientCOV"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetClientCOV")
	}

	return _child, nil
}

func (pm *_BACnetClientCOV) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetClientCOV, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetClientCOV"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetClientCOV")
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

	if popErr := writeBuffer.PopContext("BACnetClientCOV"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetClientCOV")
	}
	return nil
}

func (m *_BACnetClientCOV) IsBACnetClientCOV() {}
