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

// NLMRequestKeyUpdate is the corresponding interface of NLMRequestKeyUpdate
type NLMRequestKeyUpdate interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NLM
	// GetSet1KeyRevision returns Set1KeyRevision (property field)
	GetSet1KeyRevision() byte
	// GetSet1ActivationTime returns Set1ActivationTime (property field)
	GetSet1ActivationTime() uint32
	// GetSet1ExpirationTime returns Set1ExpirationTime (property field)
	GetSet1ExpirationTime() uint32
	// GetSet2KeyRevision returns Set2KeyRevision (property field)
	GetSet2KeyRevision() byte
	// GetSet2ActivationTime returns Set2ActivationTime (property field)
	GetSet2ActivationTime() uint32
	// GetSet2ExpirationTime returns Set2ExpirationTime (property field)
	GetSet2ExpirationTime() uint32
	// GetDistributionKeyRevision returns DistributionKeyRevision (property field)
	GetDistributionKeyRevision() byte
	// IsNLMRequestKeyUpdate is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNLMRequestKeyUpdate()
}

// _NLMRequestKeyUpdate is the data-structure of this message
type _NLMRequestKeyUpdate struct {
	NLMContract
	Set1KeyRevision         byte
	Set1ActivationTime      uint32
	Set1ExpirationTime      uint32
	Set2KeyRevision         byte
	Set2ActivationTime      uint32
	Set2ExpirationTime      uint32
	DistributionKeyRevision byte
}

var _ NLMRequestKeyUpdate = (*_NLMRequestKeyUpdate)(nil)
var _ NLMRequirements = (*_NLMRequestKeyUpdate)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMRequestKeyUpdate) GetMessageType() uint8 {
	return 0x0D
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMRequestKeyUpdate) GetParent() NLMContract {
	return m.NLMContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMRequestKeyUpdate) GetSet1KeyRevision() byte {
	return m.Set1KeyRevision
}

func (m *_NLMRequestKeyUpdate) GetSet1ActivationTime() uint32 {
	return m.Set1ActivationTime
}

func (m *_NLMRequestKeyUpdate) GetSet1ExpirationTime() uint32 {
	return m.Set1ExpirationTime
}

func (m *_NLMRequestKeyUpdate) GetSet2KeyRevision() byte {
	return m.Set2KeyRevision
}

func (m *_NLMRequestKeyUpdate) GetSet2ActivationTime() uint32 {
	return m.Set2ActivationTime
}

func (m *_NLMRequestKeyUpdate) GetSet2ExpirationTime() uint32 {
	return m.Set2ExpirationTime
}

func (m *_NLMRequestKeyUpdate) GetDistributionKeyRevision() byte {
	return m.DistributionKeyRevision
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLMRequestKeyUpdate factory function for _NLMRequestKeyUpdate
func NewNLMRequestKeyUpdate(set1KeyRevision byte, set1ActivationTime uint32, set1ExpirationTime uint32, set2KeyRevision byte, set2ActivationTime uint32, set2ExpirationTime uint32, distributionKeyRevision byte, apduLength uint16) *_NLMRequestKeyUpdate {
	_result := &_NLMRequestKeyUpdate{
		NLMContract:             NewNLM(apduLength),
		Set1KeyRevision:         set1KeyRevision,
		Set1ActivationTime:      set1ActivationTime,
		Set1ExpirationTime:      set1ExpirationTime,
		Set2KeyRevision:         set2KeyRevision,
		Set2ActivationTime:      set2ActivationTime,
		Set2ExpirationTime:      set2ExpirationTime,
		DistributionKeyRevision: distributionKeyRevision,
	}
	_result.NLMContract.(*_NLM)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNLMRequestKeyUpdate(structType any) NLMRequestKeyUpdate {
	if casted, ok := structType.(NLMRequestKeyUpdate); ok {
		return casted
	}
	if casted, ok := structType.(*NLMRequestKeyUpdate); ok {
		return *casted
	}
	return nil
}

func (m *_NLMRequestKeyUpdate) GetTypeName() string {
	return "NLMRequestKeyUpdate"
}

func (m *_NLMRequestKeyUpdate) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.NLMContract.(*_NLM).getLengthInBits(ctx))

	// Simple field (set1KeyRevision)
	lengthInBits += 8

	// Simple field (set1ActivationTime)
	lengthInBits += 32

	// Simple field (set1ExpirationTime)
	lengthInBits += 32

	// Simple field (set2KeyRevision)
	lengthInBits += 8

	// Simple field (set2ActivationTime)
	lengthInBits += 32

	// Simple field (set2ExpirationTime)
	lengthInBits += 32

	// Simple field (distributionKeyRevision)
	lengthInBits += 8

	return lengthInBits
}

func (m *_NLMRequestKeyUpdate) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NLMRequestKeyUpdate) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_NLM, apduLength uint16) (__nLMRequestKeyUpdate NLMRequestKeyUpdate, err error) {
	m.NLMContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMRequestKeyUpdate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMRequestKeyUpdate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	set1KeyRevision, err := ReadSimpleField(ctx, "set1KeyRevision", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'set1KeyRevision' field"))
	}
	m.Set1KeyRevision = set1KeyRevision

	set1ActivationTime, err := ReadSimpleField(ctx, "set1ActivationTime", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'set1ActivationTime' field"))
	}
	m.Set1ActivationTime = set1ActivationTime

	set1ExpirationTime, err := ReadSimpleField(ctx, "set1ExpirationTime", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'set1ExpirationTime' field"))
	}
	m.Set1ExpirationTime = set1ExpirationTime

	set2KeyRevision, err := ReadSimpleField(ctx, "set2KeyRevision", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'set2KeyRevision' field"))
	}
	m.Set2KeyRevision = set2KeyRevision

	set2ActivationTime, err := ReadSimpleField(ctx, "set2ActivationTime", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'set2ActivationTime' field"))
	}
	m.Set2ActivationTime = set2ActivationTime

	set2ExpirationTime, err := ReadSimpleField(ctx, "set2ExpirationTime", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'set2ExpirationTime' field"))
	}
	m.Set2ExpirationTime = set2ExpirationTime

	distributionKeyRevision, err := ReadSimpleField(ctx, "distributionKeyRevision", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'distributionKeyRevision' field"))
	}
	m.DistributionKeyRevision = distributionKeyRevision

	if closeErr := readBuffer.CloseContext("NLMRequestKeyUpdate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMRequestKeyUpdate")
	}

	return m, nil
}

func (m *_NLMRequestKeyUpdate) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMRequestKeyUpdate) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMRequestKeyUpdate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMRequestKeyUpdate")
		}

		if err := WriteSimpleField[byte](ctx, "set1KeyRevision", m.GetSet1KeyRevision(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'set1KeyRevision' field")
		}

		if err := WriteSimpleField[uint32](ctx, "set1ActivationTime", m.GetSet1ActivationTime(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'set1ActivationTime' field")
		}

		if err := WriteSimpleField[uint32](ctx, "set1ExpirationTime", m.GetSet1ExpirationTime(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'set1ExpirationTime' field")
		}

		if err := WriteSimpleField[byte](ctx, "set2KeyRevision", m.GetSet2KeyRevision(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'set2KeyRevision' field")
		}

		if err := WriteSimpleField[uint32](ctx, "set2ActivationTime", m.GetSet2ActivationTime(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'set2ActivationTime' field")
		}

		if err := WriteSimpleField[uint32](ctx, "set2ExpirationTime", m.GetSet2ExpirationTime(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'set2ExpirationTime' field")
		}

		if err := WriteSimpleField[byte](ctx, "distributionKeyRevision", m.GetDistributionKeyRevision(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'distributionKeyRevision' field")
		}

		if popErr := writeBuffer.PopContext("NLMRequestKeyUpdate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMRequestKeyUpdate")
		}
		return nil
	}
	return m.NLMContract.(*_NLM).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NLMRequestKeyUpdate) IsNLMRequestKeyUpdate() {}

func (m *_NLMRequestKeyUpdate) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
