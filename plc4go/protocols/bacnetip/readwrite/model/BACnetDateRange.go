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

// BACnetDateRange is the corresponding interface of BACnetDateRange
type BACnetDateRange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetStartDate returns StartDate (property field)
	GetStartDate() BACnetApplicationTagDate
	// GetEndDate returns EndDate (property field)
	GetEndDate() BACnetApplicationTagDate
	// IsBACnetDateRange is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetDateRange()
}

// _BACnetDateRange is the data-structure of this message
type _BACnetDateRange struct {
	StartDate BACnetApplicationTagDate
	EndDate   BACnetApplicationTagDate
}

var _ BACnetDateRange = (*_BACnetDateRange)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetDateRange) GetStartDate() BACnetApplicationTagDate {
	return m.StartDate
}

func (m *_BACnetDateRange) GetEndDate() BACnetApplicationTagDate {
	return m.EndDate
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetDateRange factory function for _BACnetDateRange
func NewBACnetDateRange(startDate BACnetApplicationTagDate, endDate BACnetApplicationTagDate) *_BACnetDateRange {
	return &_BACnetDateRange{StartDate: startDate, EndDate: endDate}
}

// Deprecated: use the interface for direct cast
func CastBACnetDateRange(structType any) BACnetDateRange {
	if casted, ok := structType.(BACnetDateRange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetDateRange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetDateRange) GetTypeName() string {
	return "BACnetDateRange"
}

func (m *_BACnetDateRange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (startDate)
	lengthInBits += m.StartDate.GetLengthInBits(ctx)

	// Simple field (endDate)
	lengthInBits += m.EndDate.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetDateRange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetDateRangeParse(ctx context.Context, theBytes []byte) (BACnetDateRange, error) {
	return BACnetDateRangeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetDateRangeParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDateRange, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDateRange, error) {
		return BACnetDateRangeParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetDateRangeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDateRange, error) {
	v, err := (&_BACnetDateRange{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetDateRange) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetDateRange BACnetDateRange, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetDateRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetDateRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	startDate, err := ReadSimpleField[BACnetApplicationTagDate](ctx, "startDate", ReadComplex[BACnetApplicationTagDate](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDate](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'startDate' field"))
	}
	m.StartDate = startDate

	endDate, err := ReadSimpleField[BACnetApplicationTagDate](ctx, "endDate", ReadComplex[BACnetApplicationTagDate](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDate](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'endDate' field"))
	}
	m.EndDate = endDate

	if closeErr := readBuffer.CloseContext("BACnetDateRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetDateRange")
	}

	return m, nil
}

func (m *_BACnetDateRange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetDateRange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetDateRange"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetDateRange")
	}

	if err := WriteSimpleField[BACnetApplicationTagDate](ctx, "startDate", m.GetStartDate(), WriteComplex[BACnetApplicationTagDate](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'startDate' field")
	}

	if err := WriteSimpleField[BACnetApplicationTagDate](ctx, "endDate", m.GetEndDate(), WriteComplex[BACnetApplicationTagDate](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'endDate' field")
	}

	if popErr := writeBuffer.PopContext("BACnetDateRange"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetDateRange")
	}
	return nil
}

func (m *_BACnetDateRange) IsBACnetDateRange() {}

func (m *_BACnetDateRange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
