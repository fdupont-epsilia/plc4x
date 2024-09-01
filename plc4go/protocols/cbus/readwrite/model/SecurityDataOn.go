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

// SecurityDataOn is the corresponding interface of SecurityDataOn
type SecurityDataOn interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SecurityData
	// GetData returns Data (property field)
	GetData() []byte
}

// SecurityDataOnExactly can be used when we want exactly this type and not a type which fulfills SecurityDataOn.
// This is useful for switch cases.
type SecurityDataOnExactly interface {
	SecurityDataOn
	isSecurityDataOn() bool
}

// _SecurityDataOn is the data-structure of this message
type _SecurityDataOn struct {
	*_SecurityData
	Data []byte
}

var _ SecurityDataOn = (*_SecurityDataOn)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataOn) InitializeParent(parent SecurityData, commandTypeContainer SecurityCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataOn) GetParent() SecurityData {
	return m._SecurityData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataOn) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSecurityDataOn factory function for _SecurityDataOn
func NewSecurityDataOn(data []byte, commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataOn {
	_result := &_SecurityDataOn{
		Data:          data,
		_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataOn(structType any) SecurityDataOn {
	if casted, ok := structType.(SecurityDataOn); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataOn); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataOn) GetTypeName() string {
	return "SecurityDataOn"
}

func (m *_SecurityDataOn) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_SecurityDataOn) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SecurityDataOnParse(ctx context.Context, theBytes []byte, commandTypeContainer SecurityCommandTypeContainer) (SecurityDataOn, error) {
	return SecurityDataOnParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), commandTypeContainer)
}

func SecurityDataOnParseWithBufferProducer(commandTypeContainer SecurityCommandTypeContainer) func(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityDataOn, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityDataOn, error) {
		return SecurityDataOnParseWithBuffer(ctx, readBuffer, commandTypeContainer)
	}
}

func SecurityDataOnParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, commandTypeContainer SecurityCommandTypeContainer) (SecurityDataOn, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataOn"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataOn")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	data, err := readBuffer.ReadByteArray("data", int(int32(commandTypeContainer.NumBytes())-int32(int32(1))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}

	if closeErr := readBuffer.CloseContext("SecurityDataOn"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataOn")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataOn{
		_SecurityData: &_SecurityData{},
		Data:          data,
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataOn) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataOn) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataOn"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataOn")
		}

		if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("SecurityDataOn"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataOn")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataOn) isSecurityDataOn() bool {
	return true
}

func (m *_SecurityDataOn) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
