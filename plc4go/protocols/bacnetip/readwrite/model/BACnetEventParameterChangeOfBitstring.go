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

// BACnetEventParameterChangeOfBitstring is the corresponding interface of BACnetEventParameterChangeOfBitstring
type BACnetEventParameterChangeOfBitstring interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetTimeDelay returns TimeDelay (property field)
	GetTimeDelay() BACnetContextTagUnsignedInteger
	// GetBitmask returns Bitmask (property field)
	GetBitmask() BACnetContextTagBitString
	// GetListOfBitstringValues returns ListOfBitstringValues (property field)
	GetListOfBitstringValues() BACnetEventParameterChangeOfBitstringListOfBitstringValues
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetEventParameterChangeOfBitstringExactly can be used when we want exactly this type and not a type which fulfills BACnetEventParameterChangeOfBitstring.
// This is useful for switch cases.
type BACnetEventParameterChangeOfBitstringExactly interface {
	BACnetEventParameterChangeOfBitstring
	isBACnetEventParameterChangeOfBitstring() bool
}

// _BACnetEventParameterChangeOfBitstring is the data-structure of this message
type _BACnetEventParameterChangeOfBitstring struct {
	*_BACnetEventParameter
	OpeningTag            BACnetOpeningTag
	TimeDelay             BACnetContextTagUnsignedInteger
	Bitmask               BACnetContextTagBitString
	ListOfBitstringValues BACnetEventParameterChangeOfBitstringListOfBitstringValues
	ClosingTag            BACnetClosingTag
}

var _ BACnetEventParameterChangeOfBitstring = (*_BACnetEventParameterChangeOfBitstring)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetEventParameterChangeOfBitstring) InitializeParent(parent BACnetEventParameter, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetEventParameterChangeOfBitstring) GetParent() BACnetEventParameter {
	return m._BACnetEventParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfBitstring) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterChangeOfBitstring) GetTimeDelay() BACnetContextTagUnsignedInteger {
	return m.TimeDelay
}

func (m *_BACnetEventParameterChangeOfBitstring) GetBitmask() BACnetContextTagBitString {
	return m.Bitmask
}

func (m *_BACnetEventParameterChangeOfBitstring) GetListOfBitstringValues() BACnetEventParameterChangeOfBitstringListOfBitstringValues {
	return m.ListOfBitstringValues
}

func (m *_BACnetEventParameterChangeOfBitstring) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterChangeOfBitstring factory function for _BACnetEventParameterChangeOfBitstring
func NewBACnetEventParameterChangeOfBitstring(openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, bitmask BACnetContextTagBitString, listOfBitstringValues BACnetEventParameterChangeOfBitstringListOfBitstringValues, closingTag BACnetClosingTag, peekedTagHeader BACnetTagHeader) *_BACnetEventParameterChangeOfBitstring {
	_result := &_BACnetEventParameterChangeOfBitstring{
		OpeningTag:            openingTag,
		TimeDelay:             timeDelay,
		Bitmask:               bitmask,
		ListOfBitstringValues: listOfBitstringValues,
		ClosingTag:            closingTag,
		_BACnetEventParameter: NewBACnetEventParameter(peekedTagHeader),
	}
	_result._BACnetEventParameter._BACnetEventParameterChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterChangeOfBitstring(structType any) BACnetEventParameterChangeOfBitstring {
	if casted, ok := structType.(BACnetEventParameterChangeOfBitstring); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfBitstring); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfBitstring) GetTypeName() string {
	return "BACnetEventParameterChangeOfBitstring"
}

func (m *_BACnetEventParameterChangeOfBitstring) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (timeDelay)
	lengthInBits += m.TimeDelay.GetLengthInBits(ctx)

	// Simple field (bitmask)
	lengthInBits += m.Bitmask.GetLengthInBits(ctx)

	// Simple field (listOfBitstringValues)
	lengthInBits += m.ListOfBitstringValues.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterChangeOfBitstring) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventParameterChangeOfBitstringParse(ctx context.Context, theBytes []byte) (BACnetEventParameterChangeOfBitstring, error) {
	return BACnetEventParameterChangeOfBitstringParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetEventParameterChangeOfBitstringParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterChangeOfBitstring, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterChangeOfBitstring, error) {
		return BACnetEventParameterChangeOfBitstringParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetEventParameterChangeOfBitstringParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterChangeOfBitstring, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfBitstring"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterChangeOfBitstring")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}

	timeDelay, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeDelay", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeDelay' field"))
	}

	bitmask, err := ReadSimpleField[BACnetContextTagBitString](ctx, "bitmask", ReadComplex[BACnetContextTagBitString](BACnetContextTagParseWithBufferProducer[BACnetContextTagBitString]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_BIT_STRING)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bitmask' field"))
	}

	listOfBitstringValues, err := ReadSimpleField[BACnetEventParameterChangeOfBitstringListOfBitstringValues](ctx, "listOfBitstringValues", ReadComplex[BACnetEventParameterChangeOfBitstringListOfBitstringValues](BACnetEventParameterChangeOfBitstringListOfBitstringValuesParseWithBufferProducer((uint8)(uint8(2))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfBitstringValues' field"))
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfBitstring"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterChangeOfBitstring")
	}

	// Create a partially initialized instance
	_child := &_BACnetEventParameterChangeOfBitstring{
		_BACnetEventParameter: &_BACnetEventParameter{},
		OpeningTag:            openingTag,
		TimeDelay:             timeDelay,
		Bitmask:               bitmask,
		ListOfBitstringValues: listOfBitstringValues,
		ClosingTag:            closingTag,
	}
	_child._BACnetEventParameter._BACnetEventParameterChildRequirements = _child
	return _child, nil
}

func (m *_BACnetEventParameterChangeOfBitstring) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterChangeOfBitstring) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfBitstring"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterChangeOfBitstring")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'openingTag' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeDelay", m.GetTimeDelay(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeDelay' field")
		}

		if err := WriteSimpleField[BACnetContextTagBitString](ctx, "bitmask", m.GetBitmask(), WriteComplex[BACnetContextTagBitString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'bitmask' field")
		}

		if err := WriteSimpleField[BACnetEventParameterChangeOfBitstringListOfBitstringValues](ctx, "listOfBitstringValues", m.GetListOfBitstringValues(), WriteComplex[BACnetEventParameterChangeOfBitstringListOfBitstringValues](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'listOfBitstringValues' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfBitstring"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterChangeOfBitstring")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetEventParameterChangeOfBitstring) isBACnetEventParameterChangeOfBitstring() bool {
	return true
}

func (m *_BACnetEventParameterChangeOfBitstring) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
