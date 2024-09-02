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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// VariantNull is the corresponding interface of VariantNull
type VariantNull interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	Variant
	// IsVariantNull is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsVariantNull()
}

// _VariantNull is the data-structure of this message
type _VariantNull struct {
	VariantContract
}

var _ VariantNull = (*_VariantNull)(nil)
var _ VariantRequirements = (*_VariantNull)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_VariantNull) GetVariantType() uint8 {
	return uint8(0)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_VariantNull) GetParent() VariantContract {
	return m.VariantContract
}

// NewVariantNull factory function for _VariantNull
func NewVariantNull(arrayLengthSpecified bool, arrayDimensionsSpecified bool, noOfArrayDimensions *int32, arrayDimensions []bool) *_VariantNull {
	_result := &_VariantNull{
		VariantContract: NewVariant(arrayLengthSpecified, arrayDimensionsSpecified, noOfArrayDimensions, arrayDimensions),
	}
	_result.VariantContract.(*_Variant)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastVariantNull(structType any) VariantNull {
	if casted, ok := structType.(VariantNull); ok {
		return casted
	}
	if casted, ok := structType.(*VariantNull); ok {
		return *casted
	}
	return nil
}

func (m *_VariantNull) GetTypeName() string {
	return "VariantNull"
}

func (m *_VariantNull) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.VariantContract.(*_Variant).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_VariantNull) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_VariantNull) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_Variant) (__variantNull VariantNull, err error) {
	m.VariantContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("VariantNull"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for VariantNull")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("VariantNull"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for VariantNull")
	}

	return m, nil
}

func (m *_VariantNull) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_VariantNull) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("VariantNull"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for VariantNull")
		}

		if popErr := writeBuffer.PopContext("VariantNull"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for VariantNull")
		}
		return nil
	}
	return m.VariantContract.(*_Variant).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_VariantNull) IsVariantNull() {}

func (m *_VariantNull) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
