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

// BACnetPrescale is the corresponding interface of BACnetPrescale
type BACnetPrescale interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetMultiplier returns Multiplier (property field)
	GetMultiplier() BACnetContextTagUnsignedInteger
	// GetModuloDivide returns ModuloDivide (property field)
	GetModuloDivide() BACnetContextTagUnsignedInteger
}

// BACnetPrescaleExactly can be used when we want exactly this type and not a type which fulfills BACnetPrescale.
// This is useful for switch cases.
type BACnetPrescaleExactly interface {
	BACnetPrescale
	isBACnetPrescale() bool
}

// _BACnetPrescale is the data-structure of this message
type _BACnetPrescale struct {
	Multiplier   BACnetContextTagUnsignedInteger
	ModuloDivide BACnetContextTagUnsignedInteger
}

var _ BACnetPrescale = (*_BACnetPrescale)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPrescale) GetMultiplier() BACnetContextTagUnsignedInteger {
	return m.Multiplier
}

func (m *_BACnetPrescale) GetModuloDivide() BACnetContextTagUnsignedInteger {
	return m.ModuloDivide
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPrescale factory function for _BACnetPrescale
func NewBACnetPrescale(multiplier BACnetContextTagUnsignedInteger, moduloDivide BACnetContextTagUnsignedInteger) *_BACnetPrescale {
	return &_BACnetPrescale{Multiplier: multiplier, ModuloDivide: moduloDivide}
}

// Deprecated: use the interface for direct cast
func CastBACnetPrescale(structType any) BACnetPrescale {
	if casted, ok := structType.(BACnetPrescale); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPrescale); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPrescale) GetTypeName() string {
	return "BACnetPrescale"
}

func (m *_BACnetPrescale) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (multiplier)
	lengthInBits += m.Multiplier.GetLengthInBits(ctx)

	// Simple field (moduloDivide)
	lengthInBits += m.ModuloDivide.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPrescale) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPrescaleParse(ctx context.Context, theBytes []byte) (BACnetPrescale, error) {
	return BACnetPrescaleParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetPrescaleParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPrescale, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPrescale, error) {
		return BACnetPrescaleParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetPrescaleParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPrescale, error) {
	v, err := (&_BACnetPrescale{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetPrescale) parse(ctx context.Context, readBuffer utils.ReadBuffer) (_bACnetPrescale BACnetPrescale, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPrescale"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPrescale")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	multiplier, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "multiplier", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'multiplier' field"))
	}
	m.Multiplier = multiplier

	moduloDivide, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "moduloDivide", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'moduloDivide' field"))
	}
	m.ModuloDivide = moduloDivide

	if closeErr := readBuffer.CloseContext("BACnetPrescale"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPrescale")
	}

	return m, nil
}

func (m *_BACnetPrescale) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPrescale) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetPrescale"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetPrescale")
	}

	if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "multiplier", m.GetMultiplier(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'multiplier' field")
	}

	if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "moduloDivide", m.GetModuloDivide(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'moduloDivide' field")
	}

	if popErr := writeBuffer.PopContext("BACnetPrescale"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetPrescale")
	}
	return nil
}

func (m *_BACnetPrescale) isBACnetPrescale() bool {
	return true
}

func (m *_BACnetPrescale) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
