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

// SecurityDataLineCutAlarmCleared is the corresponding interface of SecurityDataLineCutAlarmCleared
type SecurityDataLineCutAlarmCleared interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SecurityData
	// IsSecurityDataLineCutAlarmCleared is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataLineCutAlarmCleared()
}

// _SecurityDataLineCutAlarmCleared is the data-structure of this message
type _SecurityDataLineCutAlarmCleared struct {
	SecurityDataContract
}

var _ SecurityDataLineCutAlarmCleared = (*_SecurityDataLineCutAlarmCleared)(nil)
var _ SecurityDataRequirements = (*_SecurityDataLineCutAlarmCleared)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataLineCutAlarmCleared) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

// NewSecurityDataLineCutAlarmCleared factory function for _SecurityDataLineCutAlarmCleared
func NewSecurityDataLineCutAlarmCleared(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataLineCutAlarmCleared {
	_result := &_SecurityDataLineCutAlarmCleared{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataLineCutAlarmCleared(structType any) SecurityDataLineCutAlarmCleared {
	if casted, ok := structType.(SecurityDataLineCutAlarmCleared); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataLineCutAlarmCleared); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataLineCutAlarmCleared) GetTypeName() string {
	return "SecurityDataLineCutAlarmCleared"
}

func (m *_SecurityDataLineCutAlarmCleared) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataLineCutAlarmCleared) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataLineCutAlarmCleared) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataLineCutAlarmCleared SecurityDataLineCutAlarmCleared, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataLineCutAlarmCleared"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataLineCutAlarmCleared")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataLineCutAlarmCleared"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataLineCutAlarmCleared")
	}

	return m, nil
}

func (m *_SecurityDataLineCutAlarmCleared) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataLineCutAlarmCleared) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataLineCutAlarmCleared"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataLineCutAlarmCleared")
		}

		if popErr := writeBuffer.PopContext("SecurityDataLineCutAlarmCleared"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataLineCutAlarmCleared")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataLineCutAlarmCleared) IsSecurityDataLineCutAlarmCleared() {}

func (m *_SecurityDataLineCutAlarmCleared) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
