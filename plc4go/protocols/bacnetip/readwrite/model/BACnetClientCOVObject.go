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

// BACnetClientCOVObject is the corresponding interface of BACnetClientCOVObject
type BACnetClientCOVObject interface {
	utils.LengthAware
	utils.Serializable
	BACnetClientCOV
	// GetRealIncrement returns RealIncrement (property field)
	GetRealIncrement() BACnetApplicationTagReal
}

// BACnetClientCOVObjectExactly can be used when we want exactly this type and not a type which fulfills BACnetClientCOVObject.
// This is useful for switch cases.
type BACnetClientCOVObjectExactly interface {
	BACnetClientCOVObject
	isBACnetClientCOVObject() bool
}

// _BACnetClientCOVObject is the data-structure of this message
type _BACnetClientCOVObject struct {
	*_BACnetClientCOV
	RealIncrement BACnetApplicationTagReal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetClientCOVObject) InitializeParent(parent BACnetClientCOV, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetClientCOVObject) GetParent() BACnetClientCOV {
	return m._BACnetClientCOV
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetClientCOVObject) GetRealIncrement() BACnetApplicationTagReal {
	return m.RealIncrement
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetClientCOVObject factory function for _BACnetClientCOVObject
func NewBACnetClientCOVObject(realIncrement BACnetApplicationTagReal, peekedTagHeader BACnetTagHeader) *_BACnetClientCOVObject {
	_result := &_BACnetClientCOVObject{
		RealIncrement:    realIncrement,
		_BACnetClientCOV: NewBACnetClientCOV(peekedTagHeader),
	}
	_result._BACnetClientCOV._BACnetClientCOVChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetClientCOVObject(structType interface{}) BACnetClientCOVObject {
	if casted, ok := structType.(BACnetClientCOVObject); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetClientCOVObject); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetClientCOVObject) GetTypeName() string {
	return "BACnetClientCOVObject"
}

func (m *_BACnetClientCOVObject) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetClientCOVObject) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (realIncrement)
	lengthInBits += m.RealIncrement.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetClientCOVObject) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetClientCOVObjectParse(readBuffer utils.ReadBuffer) (BACnetClientCOVObject, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetClientCOVObject"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetClientCOVObject")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (realIncrement)
	if pullErr := readBuffer.PullContext("realIncrement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for realIncrement")
	}
	_realIncrement, _realIncrementErr := BACnetApplicationTagParse(readBuffer)
	if _realIncrementErr != nil {
		return nil, errors.Wrap(_realIncrementErr, "Error parsing 'realIncrement' field of BACnetClientCOVObject")
	}
	realIncrement := _realIncrement.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("realIncrement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for realIncrement")
	}

	if closeErr := readBuffer.CloseContext("BACnetClientCOVObject"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetClientCOVObject")
	}

	// Create a partially initialized instance
	_child := &_BACnetClientCOVObject{
		_BACnetClientCOV: &_BACnetClientCOV{},
		RealIncrement:    realIncrement,
	}
	_child._BACnetClientCOV._BACnetClientCOVChildRequirements = _child
	return _child, nil
}

func (m *_BACnetClientCOVObject) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetClientCOVObject) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetClientCOVObject"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetClientCOVObject")
		}

		// Simple Field (realIncrement)
		if pushErr := writeBuffer.PushContext("realIncrement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for realIncrement")
		}
		_realIncrementErr := writeBuffer.WriteSerializable(m.GetRealIncrement())
		if popErr := writeBuffer.PopContext("realIncrement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for realIncrement")
		}
		if _realIncrementErr != nil {
			return errors.Wrap(_realIncrementErr, "Error serializing 'realIncrement' field")
		}

		if popErr := writeBuffer.PopContext("BACnetClientCOVObject"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetClientCOVObject")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetClientCOVObject) isBACnetClientCOVObject() bool {
	return true
}

func (m *_BACnetClientCOVObject) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}