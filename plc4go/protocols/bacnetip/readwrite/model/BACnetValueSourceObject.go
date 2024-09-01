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

// BACnetValueSourceObject is the corresponding interface of BACnetValueSourceObject
type BACnetValueSourceObject interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetValueSource
	// GetObject returns Object (property field)
	GetObject() BACnetDeviceObjectReferenceEnclosed
}

// BACnetValueSourceObjectExactly can be used when we want exactly this type and not a type which fulfills BACnetValueSourceObject.
// This is useful for switch cases.
type BACnetValueSourceObjectExactly interface {
	BACnetValueSourceObject
	isBACnetValueSourceObject() bool
}

// _BACnetValueSourceObject is the data-structure of this message
type _BACnetValueSourceObject struct {
	*_BACnetValueSource
	Object BACnetDeviceObjectReferenceEnclosed
}

var _ BACnetValueSourceObject = (*_BACnetValueSourceObject)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetValueSourceObject) InitializeParent(parent BACnetValueSource, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetValueSourceObject) GetParent() BACnetValueSource {
	return m._BACnetValueSource
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetValueSourceObject) GetObject() BACnetDeviceObjectReferenceEnclosed {
	return m.Object
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetValueSourceObject factory function for _BACnetValueSourceObject
func NewBACnetValueSourceObject(object BACnetDeviceObjectReferenceEnclosed, peekedTagHeader BACnetTagHeader) *_BACnetValueSourceObject {
	_result := &_BACnetValueSourceObject{
		Object:             object,
		_BACnetValueSource: NewBACnetValueSource(peekedTagHeader),
	}
	_result._BACnetValueSource._BACnetValueSourceChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetValueSourceObject(structType any) BACnetValueSourceObject {
	if casted, ok := structType.(BACnetValueSourceObject); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetValueSourceObject); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetValueSourceObject) GetTypeName() string {
	return "BACnetValueSourceObject"
}

func (m *_BACnetValueSourceObject) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (object)
	lengthInBits += m.Object.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetValueSourceObject) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetValueSourceObjectParse(ctx context.Context, theBytes []byte) (BACnetValueSourceObject, error) {
	return BACnetValueSourceObjectParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetValueSourceObjectParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetValueSourceObject, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetValueSourceObject, error) {
		return BACnetValueSourceObjectParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetValueSourceObjectParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetValueSourceObject, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetValueSourceObject"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetValueSourceObject")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	object, err := ReadSimpleField[BACnetDeviceObjectReferenceEnclosed](ctx, "object", ReadComplex[BACnetDeviceObjectReferenceEnclosed](BACnetDeviceObjectReferenceEnclosedParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'object' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetValueSourceObject"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetValueSourceObject")
	}

	// Create a partially initialized instance
	_child := &_BACnetValueSourceObject{
		_BACnetValueSource: &_BACnetValueSource{},
		Object:             object,
	}
	_child._BACnetValueSource._BACnetValueSourceChildRequirements = _child
	return _child, nil
}

func (m *_BACnetValueSourceObject) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetValueSourceObject) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetValueSourceObject"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetValueSourceObject")
		}

		if err := WriteSimpleField[BACnetDeviceObjectReferenceEnclosed](ctx, "object", m.GetObject(), WriteComplex[BACnetDeviceObjectReferenceEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'object' field")
		}

		if popErr := writeBuffer.PopContext("BACnetValueSourceObject"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetValueSourceObject")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetValueSourceObject) isBACnetValueSourceObject() bool {
	return true
}

func (m *_BACnetValueSourceObject) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
