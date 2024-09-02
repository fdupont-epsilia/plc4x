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

// BACnetConstructedDataLoggingRecord is the corresponding interface of BACnetConstructedDataLoggingRecord
type BACnetConstructedDataLoggingRecord interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLoggingRecord returns LoggingRecord (property field)
	GetLoggingRecord() BACnetAccumulatorRecord
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetAccumulatorRecord
	// IsBACnetConstructedDataLoggingRecord is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLoggingRecord()
}

// _BACnetConstructedDataLoggingRecord is the data-structure of this message
type _BACnetConstructedDataLoggingRecord struct {
	BACnetConstructedDataContract
	LoggingRecord BACnetAccumulatorRecord
}

var _ BACnetConstructedDataLoggingRecord = (*_BACnetConstructedDataLoggingRecord)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLoggingRecord)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLoggingRecord) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLoggingRecord) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOGGING_RECORD
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLoggingRecord) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLoggingRecord) GetLoggingRecord() BACnetAccumulatorRecord {
	return m.LoggingRecord
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLoggingRecord) GetActualValue() BACnetAccumulatorRecord {
	ctx := context.Background()
	_ = ctx
	return CastBACnetAccumulatorRecord(m.GetLoggingRecord())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLoggingRecord factory function for _BACnetConstructedDataLoggingRecord
func NewBACnetConstructedDataLoggingRecord(loggingRecord BACnetAccumulatorRecord, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLoggingRecord {
	_result := &_BACnetConstructedDataLoggingRecord{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LoggingRecord:                 loggingRecord,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLoggingRecord(structType any) BACnetConstructedDataLoggingRecord {
	if casted, ok := structType.(BACnetConstructedDataLoggingRecord); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLoggingRecord); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLoggingRecord) GetTypeName() string {
	return "BACnetConstructedDataLoggingRecord"
}

func (m *_BACnetConstructedDataLoggingRecord) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (loggingRecord)
	lengthInBits += m.LoggingRecord.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLoggingRecord) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLoggingRecord) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLoggingRecord BACnetConstructedDataLoggingRecord, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLoggingRecord"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLoggingRecord")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	loggingRecord, err := ReadSimpleField[BACnetAccumulatorRecord](ctx, "loggingRecord", ReadComplex[BACnetAccumulatorRecord](BACnetAccumulatorRecordParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'loggingRecord' field"))
	}
	m.LoggingRecord = loggingRecord

	actualValue, err := ReadVirtualField[BACnetAccumulatorRecord](ctx, "actualValue", (*BACnetAccumulatorRecord)(nil), loggingRecord)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLoggingRecord"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLoggingRecord")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLoggingRecord) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLoggingRecord) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLoggingRecord"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLoggingRecord")
		}

		if err := WriteSimpleField[BACnetAccumulatorRecord](ctx, "loggingRecord", m.GetLoggingRecord(), WriteComplex[BACnetAccumulatorRecord](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'loggingRecord' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLoggingRecord"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLoggingRecord")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLoggingRecord) IsBACnetConstructedDataLoggingRecord() {}

func (m *_BACnetConstructedDataLoggingRecord) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
