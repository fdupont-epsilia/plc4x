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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// ParameterChangeReply is the corresponding interface of ParameterChangeReply
type ParameterChangeReply interface {
	utils.LengthAware
	utils.Serializable
	Reply
	// GetParameterChange returns ParameterChange (property field)
	GetParameterChange() ParameterChange
}

// ParameterChangeReplyExactly can be used when we want exactly this type and not a type which fulfills ParameterChangeReply.
// This is useful for switch cases.
type ParameterChangeReplyExactly interface {
	ParameterChangeReply
	isParameterChangeReply() bool
}

// _ParameterChangeReply is the data-structure of this message
type _ParameterChangeReply struct {
	*_Reply
	ParameterChange ParameterChange
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ParameterChangeReply) InitializeParent(parent Reply, peekedByte byte) {
	m.PeekedByte = peekedByte
}

func (m *_ParameterChangeReply) GetParent() Reply {
	return m._Reply
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ParameterChangeReply) GetParameterChange() ParameterChange {
	return m.ParameterChange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewParameterChangeReply factory function for _ParameterChangeReply
func NewParameterChangeReply(parameterChange ParameterChange, peekedByte byte, cBusOptions CBusOptions, requestContext RequestContext) *_ParameterChangeReply {
	_result := &_ParameterChangeReply{
		ParameterChange: parameterChange,
		_Reply:          NewReply(peekedByte, cBusOptions, requestContext),
	}
	_result._Reply._ReplyChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastParameterChangeReply(structType interface{}) ParameterChangeReply {
	if casted, ok := structType.(ParameterChangeReply); ok {
		return casted
	}
	if casted, ok := structType.(*ParameterChangeReply); ok {
		return *casted
	}
	return nil
}

func (m *_ParameterChangeReply) GetTypeName() string {
	return "ParameterChangeReply"
}

func (m *_ParameterChangeReply) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ParameterChangeReply) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (parameterChange)
	lengthInBits += m.ParameterChange.GetLengthInBits()

	return lengthInBits
}

func (m *_ParameterChangeReply) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ParameterChangeReplyParse(readBuffer utils.ReadBuffer, cBusOptions CBusOptions, requestContext RequestContext) (ParameterChangeReply, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ParameterChangeReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ParameterChangeReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (parameterChange)
	if pullErr := readBuffer.PullContext("parameterChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for parameterChange")
	}
	_parameterChange, _parameterChangeErr := ParameterChangeParse(readBuffer)
	if _parameterChangeErr != nil {
		return nil, errors.Wrap(_parameterChangeErr, "Error parsing 'parameterChange' field of ParameterChangeReply")
	}
	parameterChange := _parameterChange.(ParameterChange)
	if closeErr := readBuffer.CloseContext("parameterChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for parameterChange")
	}

	if closeErr := readBuffer.CloseContext("ParameterChangeReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ParameterChangeReply")
	}

	// Create a partially initialized instance
	_child := &_ParameterChangeReply{
		_Reply: &_Reply{
			CBusOptions:    cBusOptions,
			RequestContext: requestContext,
		},
		ParameterChange: parameterChange,
	}
	_child._Reply._ReplyChildRequirements = _child
	return _child, nil
}

func (m *_ParameterChangeReply) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ParameterChangeReply) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ParameterChangeReply"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ParameterChangeReply")
		}

		// Simple Field (parameterChange)
		if pushErr := writeBuffer.PushContext("parameterChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for parameterChange")
		}
		_parameterChangeErr := writeBuffer.WriteSerializable(m.GetParameterChange())
		if popErr := writeBuffer.PopContext("parameterChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for parameterChange")
		}
		if _parameterChangeErr != nil {
			return errors.Wrap(_parameterChangeErr, "Error serializing 'parameterChange' field")
		}

		if popErr := writeBuffer.PopContext("ParameterChangeReply"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ParameterChangeReply")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ParameterChangeReply) isParameterChangeReply() bool {
	return true
}

func (m *_ParameterChangeReply) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}