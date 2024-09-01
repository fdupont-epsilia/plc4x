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

// AccessControlDataAccessPointLeftOpen is the corresponding interface of AccessControlDataAccessPointLeftOpen
type AccessControlDataAccessPointLeftOpen interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AccessControlData
}

// AccessControlDataAccessPointLeftOpenExactly can be used when we want exactly this type and not a type which fulfills AccessControlDataAccessPointLeftOpen.
// This is useful for switch cases.
type AccessControlDataAccessPointLeftOpenExactly interface {
	AccessControlDataAccessPointLeftOpen
	isAccessControlDataAccessPointLeftOpen() bool
}

// _AccessControlDataAccessPointLeftOpen is the data-structure of this message
type _AccessControlDataAccessPointLeftOpen struct {
	AccessControlDataContract
}

var _ AccessControlDataAccessPointLeftOpen = (*_AccessControlDataAccessPointLeftOpen)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AccessControlDataAccessPointLeftOpen) GetParent() AccessControlDataContract {
	return m.AccessControlDataContract
}

// NewAccessControlDataAccessPointLeftOpen factory function for _AccessControlDataAccessPointLeftOpen
func NewAccessControlDataAccessPointLeftOpen(commandTypeContainer AccessControlCommandTypeContainer, networkId byte, accessPointId byte) *_AccessControlDataAccessPointLeftOpen {
	_result := &_AccessControlDataAccessPointLeftOpen{
		AccessControlDataContract: NewAccessControlData(commandTypeContainer, networkId, accessPointId),
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastAccessControlDataAccessPointLeftOpen(structType any) AccessControlDataAccessPointLeftOpen {
	if casted, ok := structType.(AccessControlDataAccessPointLeftOpen); ok {
		return casted
	}
	if casted, ok := structType.(*AccessControlDataAccessPointLeftOpen); ok {
		return *casted
	}
	return nil
}

func (m *_AccessControlDataAccessPointLeftOpen) GetTypeName() string {
	return "AccessControlDataAccessPointLeftOpen"
}

func (m *_AccessControlDataAccessPointLeftOpen) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AccessControlDataContract.(*_AccessControlData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_AccessControlDataAccessPointLeftOpen) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AccessControlDataAccessPointLeftOpen) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AccessControlData) (_accessControlDataAccessPointLeftOpen AccessControlDataAccessPointLeftOpen, err error) {
	m.AccessControlDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AccessControlDataAccessPointLeftOpen"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AccessControlDataAccessPointLeftOpen")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("AccessControlDataAccessPointLeftOpen"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AccessControlDataAccessPointLeftOpen")
	}

	return m, nil
}

func (m *_AccessControlDataAccessPointLeftOpen) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AccessControlDataAccessPointLeftOpen) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AccessControlDataAccessPointLeftOpen"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AccessControlDataAccessPointLeftOpen")
		}

		if popErr := writeBuffer.PopContext("AccessControlDataAccessPointLeftOpen"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AccessControlDataAccessPointLeftOpen")
		}
		return nil
	}
	return m.AccessControlDataContract.(*_AccessControlData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AccessControlDataAccessPointLeftOpen) isAccessControlDataAccessPointLeftOpen() bool {
	return true
}

func (m *_AccessControlDataAccessPointLeftOpen) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
