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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// SecurityDataOtherAlarmRaised is the corresponding interface of SecurityDataOtherAlarmRaised
type SecurityDataOtherAlarmRaised interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SecurityData
}

// SecurityDataOtherAlarmRaisedExactly can be used when we want exactly this type and not a type which fulfills SecurityDataOtherAlarmRaised.
// This is useful for switch cases.
type SecurityDataOtherAlarmRaisedExactly interface {
	SecurityDataOtherAlarmRaised
	isSecurityDataOtherAlarmRaised() bool
}

// _SecurityDataOtherAlarmRaised is the data-structure of this message
type _SecurityDataOtherAlarmRaised struct {
	*_SecurityData
}

var _ SecurityDataOtherAlarmRaised = (*_SecurityDataOtherAlarmRaised)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataOtherAlarmRaised) InitializeParent(parent SecurityData, commandTypeContainer SecurityCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataOtherAlarmRaised) GetParent() SecurityData {
	return m._SecurityData
}

// NewSecurityDataOtherAlarmRaised factory function for _SecurityDataOtherAlarmRaised
func NewSecurityDataOtherAlarmRaised(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataOtherAlarmRaised {
	_result := &_SecurityDataOtherAlarmRaised{
		_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataOtherAlarmRaised(structType any) SecurityDataOtherAlarmRaised {
	if casted, ok := structType.(SecurityDataOtherAlarmRaised); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataOtherAlarmRaised); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataOtherAlarmRaised) GetTypeName() string {
	return "SecurityDataOtherAlarmRaised"
}

func (m *_SecurityDataOtherAlarmRaised) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataOtherAlarmRaised) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SecurityDataOtherAlarmRaisedParse(ctx context.Context, theBytes []byte) (SecurityDataOtherAlarmRaised, error) {
	return SecurityDataOtherAlarmRaisedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func SecurityDataOtherAlarmRaisedParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityDataOtherAlarmRaised, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityDataOtherAlarmRaised, error) {
		return SecurityDataOtherAlarmRaisedParseWithBuffer(ctx, readBuffer)
	}
}

func SecurityDataOtherAlarmRaisedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityDataOtherAlarmRaised, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataOtherAlarmRaised"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataOtherAlarmRaised")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataOtherAlarmRaised"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataOtherAlarmRaised")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataOtherAlarmRaised{
		_SecurityData: &_SecurityData{},
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataOtherAlarmRaised) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataOtherAlarmRaised) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataOtherAlarmRaised"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataOtherAlarmRaised")
		}

		if popErr := writeBuffer.PopContext("SecurityDataOtherAlarmRaised"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataOtherAlarmRaised")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataOtherAlarmRaised) isSecurityDataOtherAlarmRaised() bool {
	return true
}

func (m *_SecurityDataOtherAlarmRaised) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
