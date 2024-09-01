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

// BACnetConfirmedServiceRequestAtomicReadFileRecord is the corresponding interface of BACnetConfirmedServiceRequestAtomicReadFileRecord
type BACnetConfirmedServiceRequestAtomicReadFileRecord interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord
	// GetFileStartRecord returns FileStartRecord (property field)
	GetFileStartRecord() BACnetApplicationTagSignedInteger
	// GetRequestRecordCount returns RequestRecordCount (property field)
	GetRequestRecordCount() BACnetApplicationTagUnsignedInteger
}

// BACnetConfirmedServiceRequestAtomicReadFileRecordExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestAtomicReadFileRecord.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestAtomicReadFileRecordExactly interface {
	BACnetConfirmedServiceRequestAtomicReadFileRecord
	isBACnetConfirmedServiceRequestAtomicReadFileRecord() bool
}

// _BACnetConfirmedServiceRequestAtomicReadFileRecord is the data-structure of this message
type _BACnetConfirmedServiceRequestAtomicReadFileRecord struct {
	BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract
	FileStartRecord    BACnetApplicationTagSignedInteger
	RequestRecordCount BACnetApplicationTagUnsignedInteger
}

var _ BACnetConfirmedServiceRequestAtomicReadFileRecord = (*_BACnetConfirmedServiceRequestAtomicReadFileRecord)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) GetParent() BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract {
	return m.BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) GetFileStartRecord() BACnetApplicationTagSignedInteger {
	return m.FileStartRecord
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) GetRequestRecordCount() BACnetApplicationTagUnsignedInteger {
	return m.RequestRecordCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestAtomicReadFileRecord factory function for _BACnetConfirmedServiceRequestAtomicReadFileRecord
func NewBACnetConfirmedServiceRequestAtomicReadFileRecord(fileStartRecord BACnetApplicationTagSignedInteger, requestRecordCount BACnetApplicationTagUnsignedInteger, peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, closingTag BACnetClosingTag) *_BACnetConfirmedServiceRequestAtomicReadFileRecord {
	_result := &_BACnetConfirmedServiceRequestAtomicReadFileRecord{
		BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract: NewBACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord(peekedTagHeader, openingTag, closingTag),
		FileStartRecord:    fileStartRecord,
		RequestRecordCount: requestRecordCount,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestAtomicReadFileRecord(structType any) BACnetConfirmedServiceRequestAtomicReadFileRecord {
	if casted, ok := structType.(BACnetConfirmedServiceRequestAtomicReadFileRecord); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestAtomicReadFileRecord); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) GetTypeName() string {
	return "BACnetConfirmedServiceRequestAtomicReadFileRecord"
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract.(*_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord).getLengthInBits(ctx))

	// Simple field (fileStartRecord)
	lengthInBits += m.FileStartRecord.GetLengthInBits(ctx)

	// Simple field (requestRecordCount)
	lengthInBits += m.RequestRecordCount.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord) (_bACnetConfirmedServiceRequestAtomicReadFileRecord BACnetConfirmedServiceRequestAtomicReadFileRecord, err error) {
	m.BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestAtomicReadFileRecord"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestAtomicReadFileRecord")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	fileStartRecord, err := ReadSimpleField[BACnetApplicationTagSignedInteger](ctx, "fileStartRecord", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagSignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fileStartRecord' field"))
	}
	m.FileStartRecord = fileStartRecord

	requestRecordCount, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "requestRecordCount", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestRecordCount' field"))
	}
	m.RequestRecordCount = requestRecordCount

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestAtomicReadFileRecord"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestAtomicReadFileRecord")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestAtomicReadFileRecord"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestAtomicReadFileRecord")
		}

		if err := WriteSimpleField[BACnetApplicationTagSignedInteger](ctx, "fileStartRecord", m.GetFileStartRecord(), WriteComplex[BACnetApplicationTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'fileStartRecord' field")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "requestRecordCount", m.GetRequestRecordCount(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestRecordCount' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestAtomicReadFileRecord"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestAtomicReadFileRecord")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract.(*_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) isBACnetConfirmedServiceRequestAtomicReadFileRecord() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
