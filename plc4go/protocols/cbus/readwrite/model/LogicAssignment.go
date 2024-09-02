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

// LogicAssignment is the corresponding interface of LogicAssignment
type LogicAssignment interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetGreaterOfOrLogic returns GreaterOfOrLogic (property field)
	GetGreaterOfOrLogic() bool
	// GetReStrikeDelay returns ReStrikeDelay (property field)
	GetReStrikeDelay() bool
	// GetAssignedToGav16 returns AssignedToGav16 (property field)
	GetAssignedToGav16() bool
	// GetAssignedToGav15 returns AssignedToGav15 (property field)
	GetAssignedToGav15() bool
	// GetAssignedToGav14 returns AssignedToGav14 (property field)
	GetAssignedToGav14() bool
	// GetAssignedToGav13 returns AssignedToGav13 (property field)
	GetAssignedToGav13() bool
	// IsLogicAssignment is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsLogicAssignment()
}

// _LogicAssignment is the data-structure of this message
type _LogicAssignment struct {
	GreaterOfOrLogic bool
	ReStrikeDelay    bool
	AssignedToGav16  bool
	AssignedToGav15  bool
	AssignedToGav14  bool
	AssignedToGav13  bool
	// Reserved Fields
	reservedField0 *bool
	reservedField1 *bool
}

var _ LogicAssignment = (*_LogicAssignment)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LogicAssignment) GetGreaterOfOrLogic() bool {
	return m.GreaterOfOrLogic
}

func (m *_LogicAssignment) GetReStrikeDelay() bool {
	return m.ReStrikeDelay
}

func (m *_LogicAssignment) GetAssignedToGav16() bool {
	return m.AssignedToGav16
}

func (m *_LogicAssignment) GetAssignedToGav15() bool {
	return m.AssignedToGav15
}

func (m *_LogicAssignment) GetAssignedToGav14() bool {
	return m.AssignedToGav14
}

func (m *_LogicAssignment) GetAssignedToGav13() bool {
	return m.AssignedToGav13
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewLogicAssignment factory function for _LogicAssignment
func NewLogicAssignment(greaterOfOrLogic bool, reStrikeDelay bool, assignedToGav16 bool, assignedToGav15 bool, assignedToGav14 bool, assignedToGav13 bool) *_LogicAssignment {
	return &_LogicAssignment{GreaterOfOrLogic: greaterOfOrLogic, ReStrikeDelay: reStrikeDelay, AssignedToGav16: assignedToGav16, AssignedToGav15: assignedToGav15, AssignedToGav14: assignedToGav14, AssignedToGav13: assignedToGav13}
}

// Deprecated: use the interface for direct cast
func CastLogicAssignment(structType any) LogicAssignment {
	if casted, ok := structType.(LogicAssignment); ok {
		return casted
	}
	if casted, ok := structType.(*LogicAssignment); ok {
		return *casted
	}
	return nil
}

func (m *_LogicAssignment) GetTypeName() string {
	return "LogicAssignment"
}

func (m *_LogicAssignment) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (greaterOfOrLogic)
	lengthInBits += 1

	// Simple field (reStrikeDelay)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (assignedToGav16)
	lengthInBits += 1

	// Simple field (assignedToGav15)
	lengthInBits += 1

	// Simple field (assignedToGav14)
	lengthInBits += 1

	// Simple field (assignedToGav13)
	lengthInBits += 1

	return lengthInBits
}

func (m *_LogicAssignment) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func LogicAssignmentParse(ctx context.Context, theBytes []byte) (LogicAssignment, error) {
	return LogicAssignmentParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func LogicAssignmentParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (LogicAssignment, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (LogicAssignment, error) {
		return LogicAssignmentParseWithBuffer(ctx, readBuffer)
	}
}

func LogicAssignmentParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (LogicAssignment, error) {
	v, err := (&_LogicAssignment{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_LogicAssignment) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__logicAssignment LogicAssignment, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LogicAssignment"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LogicAssignment")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	greaterOfOrLogic, err := ReadSimpleField(ctx, "greaterOfOrLogic", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'greaterOfOrLogic' field"))
	}
	m.GreaterOfOrLogic = greaterOfOrLogic

	reStrikeDelay, err := ReadSimpleField(ctx, "reStrikeDelay", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reStrikeDelay' field"))
	}
	m.ReStrikeDelay = reStrikeDelay

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	reservedField1, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField1 = reservedField1

	assignedToGav16, err := ReadSimpleField(ctx, "assignedToGav16", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'assignedToGav16' field"))
	}
	m.AssignedToGav16 = assignedToGav16

	assignedToGav15, err := ReadSimpleField(ctx, "assignedToGav15", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'assignedToGav15' field"))
	}
	m.AssignedToGav15 = assignedToGav15

	assignedToGav14, err := ReadSimpleField(ctx, "assignedToGav14", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'assignedToGav14' field"))
	}
	m.AssignedToGav14 = assignedToGav14

	assignedToGav13, err := ReadSimpleField(ctx, "assignedToGav13", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'assignedToGav13' field"))
	}
	m.AssignedToGav13 = assignedToGav13

	if closeErr := readBuffer.CloseContext("LogicAssignment"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LogicAssignment")
	}

	return m, nil
}

func (m *_LogicAssignment) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LogicAssignment) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("LogicAssignment"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for LogicAssignment")
	}

	if err := WriteSimpleField[bool](ctx, "greaterOfOrLogic", m.GetGreaterOfOrLogic(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'greaterOfOrLogic' field")
	}

	if err := WriteSimpleField[bool](ctx, "reStrikeDelay", m.GetReStrikeDelay(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reStrikeDelay' field")
	}

	if err := WriteReservedField[bool](ctx, "reserved", bool(false), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 1")
	}

	if err := WriteReservedField[bool](ctx, "reserved", bool(false), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 2")
	}

	if err := WriteSimpleField[bool](ctx, "assignedToGav16", m.GetAssignedToGav16(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'assignedToGav16' field")
	}

	if err := WriteSimpleField[bool](ctx, "assignedToGav15", m.GetAssignedToGav15(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'assignedToGav15' field")
	}

	if err := WriteSimpleField[bool](ctx, "assignedToGav14", m.GetAssignedToGav14(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'assignedToGav14' field")
	}

	if err := WriteSimpleField[bool](ctx, "assignedToGav13", m.GetAssignedToGav13(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'assignedToGav13' field")
	}

	if popErr := writeBuffer.PopContext("LogicAssignment"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for LogicAssignment")
	}
	return nil
}

func (m *_LogicAssignment) IsLogicAssignment() {}

func (m *_LogicAssignment) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
