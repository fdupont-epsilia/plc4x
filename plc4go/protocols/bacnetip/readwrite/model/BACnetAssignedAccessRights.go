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

// BACnetAssignedAccessRights is the corresponding interface of BACnetAssignedAccessRights
type BACnetAssignedAccessRights interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetAssignedAccessRights returns AssignedAccessRights (property field)
	GetAssignedAccessRights() BACnetDeviceObjectReferenceEnclosed
	// GetEnable returns Enable (property field)
	GetEnable() BACnetContextTagBoolean
	// IsBACnetAssignedAccessRights is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetAssignedAccessRights()
}

// _BACnetAssignedAccessRights is the data-structure of this message
type _BACnetAssignedAccessRights struct {
	AssignedAccessRights BACnetDeviceObjectReferenceEnclosed
	Enable               BACnetContextTagBoolean
}

var _ BACnetAssignedAccessRights = (*_BACnetAssignedAccessRights)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAssignedAccessRights) GetAssignedAccessRights() BACnetDeviceObjectReferenceEnclosed {
	return m.AssignedAccessRights
}

func (m *_BACnetAssignedAccessRights) GetEnable() BACnetContextTagBoolean {
	return m.Enable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetAssignedAccessRights factory function for _BACnetAssignedAccessRights
func NewBACnetAssignedAccessRights(assignedAccessRights BACnetDeviceObjectReferenceEnclosed, enable BACnetContextTagBoolean) *_BACnetAssignedAccessRights {
	return &_BACnetAssignedAccessRights{AssignedAccessRights: assignedAccessRights, Enable: enable}
}

// Deprecated: use the interface for direct cast
func CastBACnetAssignedAccessRights(structType any) BACnetAssignedAccessRights {
	if casted, ok := structType.(BACnetAssignedAccessRights); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAssignedAccessRights); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAssignedAccessRights) GetTypeName() string {
	return "BACnetAssignedAccessRights"
}

func (m *_BACnetAssignedAccessRights) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (assignedAccessRights)
	lengthInBits += m.AssignedAccessRights.GetLengthInBits(ctx)

	// Simple field (enable)
	lengthInBits += m.Enable.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetAssignedAccessRights) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAssignedAccessRightsParse(ctx context.Context, theBytes []byte) (BACnetAssignedAccessRights, error) {
	return BACnetAssignedAccessRightsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetAssignedAccessRightsParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAssignedAccessRights, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAssignedAccessRights, error) {
		return BACnetAssignedAccessRightsParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetAssignedAccessRightsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAssignedAccessRights, error) {
	v, err := (&_BACnetAssignedAccessRights{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetAssignedAccessRights) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetAssignedAccessRights BACnetAssignedAccessRights, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAssignedAccessRights"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAssignedAccessRights")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	assignedAccessRights, err := ReadSimpleField[BACnetDeviceObjectReferenceEnclosed](ctx, "assignedAccessRights", ReadComplex[BACnetDeviceObjectReferenceEnclosed](BACnetDeviceObjectReferenceEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'assignedAccessRights' field"))
	}
	m.AssignedAccessRights = assignedAccessRights

	enable, err := ReadSimpleField[BACnetContextTagBoolean](ctx, "enable", ReadComplex[BACnetContextTagBoolean](BACnetContextTagParseWithBufferProducer[BACnetContextTagBoolean]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_BOOLEAN)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'enable' field"))
	}
	m.Enable = enable

	if closeErr := readBuffer.CloseContext("BACnetAssignedAccessRights"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAssignedAccessRights")
	}

	return m, nil
}

func (m *_BACnetAssignedAccessRights) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAssignedAccessRights) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAssignedAccessRights"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAssignedAccessRights")
	}

	if err := WriteSimpleField[BACnetDeviceObjectReferenceEnclosed](ctx, "assignedAccessRights", m.GetAssignedAccessRights(), WriteComplex[BACnetDeviceObjectReferenceEnclosed](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'assignedAccessRights' field")
	}

	if err := WriteSimpleField[BACnetContextTagBoolean](ctx, "enable", m.GetEnable(), WriteComplex[BACnetContextTagBoolean](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'enable' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAssignedAccessRights"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAssignedAccessRights")
	}
	return nil
}

func (m *_BACnetAssignedAccessRights) IsBACnetAssignedAccessRights() {}

func (m *_BACnetAssignedAccessRights) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
