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

// BACnetApplicationTagTime is the corresponding interface of BACnetApplicationTagTime
type BACnetApplicationTagTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetApplicationTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadTime
}

// BACnetApplicationTagTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetApplicationTagTime.
// This is useful for switch cases.
type BACnetApplicationTagTimeExactly interface {
	BACnetApplicationTagTime
	isBACnetApplicationTagTime() bool
}

// _BACnetApplicationTagTime is the data-structure of this message
type _BACnetApplicationTagTime struct {
	*_BACnetApplicationTag
	Payload BACnetTagPayloadTime
}

var _ BACnetApplicationTagTime = (*_BACnetApplicationTagTime)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetApplicationTagTime) InitializeParent(parent BACnetApplicationTag, header BACnetTagHeader) {
	m.Header = header
}

func (m *_BACnetApplicationTagTime) GetParent() BACnetApplicationTag {
	return m._BACnetApplicationTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetApplicationTagTime) GetPayload() BACnetTagPayloadTime {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetApplicationTagTime factory function for _BACnetApplicationTagTime
func NewBACnetApplicationTagTime(payload BACnetTagPayloadTime, header BACnetTagHeader) *_BACnetApplicationTagTime {
	_result := &_BACnetApplicationTagTime{
		Payload:               payload,
		_BACnetApplicationTag: NewBACnetApplicationTag(header),
	}
	_result._BACnetApplicationTag._BACnetApplicationTagChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetApplicationTagTime(structType any) BACnetApplicationTagTime {
	if casted, ok := structType.(BACnetApplicationTagTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetApplicationTagTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetApplicationTagTime) GetTypeName() string {
	return "BACnetApplicationTagTime"
}

func (m *_BACnetApplicationTagTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetApplicationTagTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetApplicationTagTimeParse(ctx context.Context, theBytes []byte) (BACnetApplicationTagTime, error) {
	return BACnetApplicationTagTimeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetApplicationTagTimeParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetApplicationTagTime, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetApplicationTagTime, error) {
		return BACnetApplicationTagTimeParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetApplicationTagTimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetApplicationTagTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetApplicationTagTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetApplicationTagTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	payload, err := ReadSimpleField[BACnetTagPayloadTime](ctx, "payload", ReadComplex[BACnetTagPayloadTime](BACnetTagPayloadTimeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetApplicationTagTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetApplicationTagTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetApplicationTagTime{
		_BACnetApplicationTag: &_BACnetApplicationTag{},
		Payload:               payload,
	}
	_child._BACnetApplicationTag._BACnetApplicationTagChildRequirements = _child
	return _child, nil
}

func (m *_BACnetApplicationTagTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetApplicationTagTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetApplicationTagTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetApplicationTagTime")
		}

		if err := WriteSimpleField[BACnetTagPayloadTime](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'payload' field")
		}

		if popErr := writeBuffer.PopContext("BACnetApplicationTagTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetApplicationTagTime")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetApplicationTagTime) isBACnetApplicationTagTime() bool {
	return true
}

func (m *_BACnetApplicationTagTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
