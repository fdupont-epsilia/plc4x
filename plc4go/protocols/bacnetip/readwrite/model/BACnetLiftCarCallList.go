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

// BACnetLiftCarCallList is the corresponding interface of BACnetLiftCarCallList
type BACnetLiftCarCallList interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetFloorNumbers returns FloorNumbers (property field)
	GetFloorNumbers() BACnetLiftCarCallListFloorList
}

// BACnetLiftCarCallListExactly can be used when we want exactly this type and not a type which fulfills BACnetLiftCarCallList.
// This is useful for switch cases.
type BACnetLiftCarCallListExactly interface {
	BACnetLiftCarCallList
	isBACnetLiftCarCallList() bool
}

// _BACnetLiftCarCallList is the data-structure of this message
type _BACnetLiftCarCallList struct {
	FloorNumbers BACnetLiftCarCallListFloorList
}

var _ BACnetLiftCarCallList = (*_BACnetLiftCarCallList)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLiftCarCallList) GetFloorNumbers() BACnetLiftCarCallListFloorList {
	return m.FloorNumbers
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLiftCarCallList factory function for _BACnetLiftCarCallList
func NewBACnetLiftCarCallList(floorNumbers BACnetLiftCarCallListFloorList) *_BACnetLiftCarCallList {
	return &_BACnetLiftCarCallList{FloorNumbers: floorNumbers}
}

// Deprecated: use the interface for direct cast
func CastBACnetLiftCarCallList(structType any) BACnetLiftCarCallList {
	if casted, ok := structType.(BACnetLiftCarCallList); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLiftCarCallList); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLiftCarCallList) GetTypeName() string {
	return "BACnetLiftCarCallList"
}

func (m *_BACnetLiftCarCallList) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (floorNumbers)
	lengthInBits += m.FloorNumbers.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLiftCarCallList) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLiftCarCallListParse(ctx context.Context, theBytes []byte) (BACnetLiftCarCallList, error) {
	return BACnetLiftCarCallListParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetLiftCarCallListParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLiftCarCallList, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLiftCarCallList, error) {
		return BACnetLiftCarCallListParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetLiftCarCallListParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLiftCarCallList, error) {
	v, err := (&_BACnetLiftCarCallList{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetLiftCarCallList) parse(ctx context.Context, readBuffer utils.ReadBuffer) (_bACnetLiftCarCallList BACnetLiftCarCallList, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLiftCarCallList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLiftCarCallList")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	floorNumbers, err := ReadSimpleField[BACnetLiftCarCallListFloorList](ctx, "floorNumbers", ReadComplex[BACnetLiftCarCallListFloorList](BACnetLiftCarCallListFloorListParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'floorNumbers' field"))
	}
	m.FloorNumbers = floorNumbers

	if closeErr := readBuffer.CloseContext("BACnetLiftCarCallList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLiftCarCallList")
	}

	return m, nil
}

func (m *_BACnetLiftCarCallList) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLiftCarCallList) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetLiftCarCallList"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLiftCarCallList")
	}

	if err := WriteSimpleField[BACnetLiftCarCallListFloorList](ctx, "floorNumbers", m.GetFloorNumbers(), WriteComplex[BACnetLiftCarCallListFloorList](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'floorNumbers' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLiftCarCallList"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLiftCarCallList")
	}
	return nil
}

func (m *_BACnetLiftCarCallList) isBACnetLiftCarCallList() bool {
	return true
}

func (m *_BACnetLiftCarCallList) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
