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

// SecurityDataPanicCleared is the corresponding interface of SecurityDataPanicCleared
type SecurityDataPanicCleared interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SecurityData
	// IsSecurityDataPanicCleared is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataPanicCleared()
}

// _SecurityDataPanicCleared is the data-structure of this message
type _SecurityDataPanicCleared struct {
	SecurityDataContract
}

var _ SecurityDataPanicCleared = (*_SecurityDataPanicCleared)(nil)
var _ SecurityDataRequirements = (*_SecurityDataPanicCleared)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataPanicCleared) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

// NewSecurityDataPanicCleared factory function for _SecurityDataPanicCleared
func NewSecurityDataPanicCleared(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataPanicCleared {
	_result := &_SecurityDataPanicCleared{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataPanicCleared(structType any) SecurityDataPanicCleared {
	if casted, ok := structType.(SecurityDataPanicCleared); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataPanicCleared); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataPanicCleared) GetTypeName() string {
	return "SecurityDataPanicCleared"
}

func (m *_SecurityDataPanicCleared) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataPanicCleared) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataPanicCleared) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataPanicCleared SecurityDataPanicCleared, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataPanicCleared"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataPanicCleared")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataPanicCleared"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataPanicCleared")
	}

	return m, nil
}

func (m *_SecurityDataPanicCleared) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataPanicCleared) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataPanicCleared"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataPanicCleared")
		}

		if popErr := writeBuffer.PopContext("SecurityDataPanicCleared"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataPanicCleared")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataPanicCleared) IsSecurityDataPanicCleared() {}

func (m *_SecurityDataPanicCleared) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
