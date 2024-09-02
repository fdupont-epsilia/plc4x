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

// BACnetLogDataLogStatus is the corresponding interface of BACnetLogDataLogStatus
type BACnetLogDataLogStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetLogData
	// GetLogStatus returns LogStatus (property field)
	GetLogStatus() BACnetLogStatusTagged
	// IsBACnetLogDataLogStatus is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLogDataLogStatus()
}

// _BACnetLogDataLogStatus is the data-structure of this message
type _BACnetLogDataLogStatus struct {
	BACnetLogDataContract
	LogStatus BACnetLogStatusTagged
}

var _ BACnetLogDataLogStatus = (*_BACnetLogDataLogStatus)(nil)
var _ BACnetLogDataRequirements = (*_BACnetLogDataLogStatus)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetLogDataLogStatus) GetParent() BACnetLogDataContract {
	return m.BACnetLogDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogDataLogStatus) GetLogStatus() BACnetLogStatusTagged {
	return m.LogStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogDataLogStatus factory function for _BACnetLogDataLogStatus
func NewBACnetLogDataLogStatus(logStatus BACnetLogStatusTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetLogDataLogStatus {
	_result := &_BACnetLogDataLogStatus{
		BACnetLogDataContract: NewBACnetLogData(openingTag, peekedTagHeader, closingTag, tagNumber),
		LogStatus:             logStatus,
	}
	_result.BACnetLogDataContract.(*_BACnetLogData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetLogDataLogStatus(structType any) BACnetLogDataLogStatus {
	if casted, ok := structType.(BACnetLogDataLogStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogDataLogStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogDataLogStatus) GetTypeName() string {
	return "BACnetLogDataLogStatus"
}

func (m *_BACnetLogDataLogStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetLogDataContract.(*_BACnetLogData).getLengthInBits(ctx))

	// Simple field (logStatus)
	lengthInBits += m.LogStatus.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLogDataLogStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetLogDataLogStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetLogData, tagNumber uint8) (__bACnetLogDataLogStatus BACnetLogDataLogStatus, err error) {
	m.BACnetLogDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogDataLogStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogDataLogStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	logStatus, err := ReadSimpleField[BACnetLogStatusTagged](ctx, "logStatus", ReadComplex[BACnetLogStatusTagged](BACnetLogStatusTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'logStatus' field"))
	}
	m.LogStatus = logStatus

	if closeErr := readBuffer.CloseContext("BACnetLogDataLogStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogDataLogStatus")
	}

	return m, nil
}

func (m *_BACnetLogDataLogStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogDataLogStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogDataLogStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogDataLogStatus")
		}

		if err := WriteSimpleField[BACnetLogStatusTagged](ctx, "logStatus", m.GetLogStatus(), WriteComplex[BACnetLogStatusTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'logStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogDataLogStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogDataLogStatus")
		}
		return nil
	}
	return m.BACnetLogDataContract.(*_BACnetLogData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetLogDataLogStatus) IsBACnetLogDataLogStatus() {}

func (m *_BACnetLogDataLogStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
