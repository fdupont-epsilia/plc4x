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

// SecurityDataLowBatteryCorrected is the corresponding interface of SecurityDataLowBatteryCorrected
type SecurityDataLowBatteryCorrected interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SecurityData
}

// SecurityDataLowBatteryCorrectedExactly can be used when we want exactly this type and not a type which fulfills SecurityDataLowBatteryCorrected.
// This is useful for switch cases.
type SecurityDataLowBatteryCorrectedExactly interface {
	SecurityDataLowBatteryCorrected
	isSecurityDataLowBatteryCorrected() bool
}

// _SecurityDataLowBatteryCorrected is the data-structure of this message
type _SecurityDataLowBatteryCorrected struct {
	*_SecurityData
}

var _ SecurityDataLowBatteryCorrected = (*_SecurityDataLowBatteryCorrected)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataLowBatteryCorrected) InitializeParent(parent SecurityData, commandTypeContainer SecurityCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataLowBatteryCorrected) GetParent() SecurityData {
	return m._SecurityData
}

// NewSecurityDataLowBatteryCorrected factory function for _SecurityDataLowBatteryCorrected
func NewSecurityDataLowBatteryCorrected(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataLowBatteryCorrected {
	_result := &_SecurityDataLowBatteryCorrected{
		_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataLowBatteryCorrected(structType any) SecurityDataLowBatteryCorrected {
	if casted, ok := structType.(SecurityDataLowBatteryCorrected); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataLowBatteryCorrected); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataLowBatteryCorrected) GetTypeName() string {
	return "SecurityDataLowBatteryCorrected"
}

func (m *_SecurityDataLowBatteryCorrected) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataLowBatteryCorrected) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SecurityDataLowBatteryCorrectedParse(ctx context.Context, theBytes []byte) (SecurityDataLowBatteryCorrected, error) {
	return SecurityDataLowBatteryCorrectedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func SecurityDataLowBatteryCorrectedParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityDataLowBatteryCorrected, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityDataLowBatteryCorrected, error) {
		return SecurityDataLowBatteryCorrectedParseWithBuffer(ctx, readBuffer)
	}
}

func SecurityDataLowBatteryCorrectedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityDataLowBatteryCorrected, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataLowBatteryCorrected"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataLowBatteryCorrected")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataLowBatteryCorrected"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataLowBatteryCorrected")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataLowBatteryCorrected{
		_SecurityData: &_SecurityData{},
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataLowBatteryCorrected) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataLowBatteryCorrected) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataLowBatteryCorrected"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataLowBatteryCorrected")
		}

		if popErr := writeBuffer.PopContext("SecurityDataLowBatteryCorrected"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataLowBatteryCorrected")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataLowBatteryCorrected) isSecurityDataLowBatteryCorrected() bool {
	return true
}

func (m *_SecurityDataLowBatteryCorrected) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
