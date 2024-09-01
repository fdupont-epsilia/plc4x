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

// BACnetFaultParameterFaultState is the corresponding interface of BACnetFaultParameterFaultState
type BACnetFaultParameterFaultState interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetFaultParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfFaultValues returns ListOfFaultValues (property field)
	GetListOfFaultValues() BACnetFaultParameterFaultStateListOfFaultValues
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetFaultParameterFaultStateExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultState.
// This is useful for switch cases.
type BACnetFaultParameterFaultStateExactly interface {
	BACnetFaultParameterFaultState
	isBACnetFaultParameterFaultState() bool
}

// _BACnetFaultParameterFaultState is the data-structure of this message
type _BACnetFaultParameterFaultState struct {
	BACnetFaultParameterContract
	OpeningTag        BACnetOpeningTag
	ListOfFaultValues BACnetFaultParameterFaultStateListOfFaultValues
	ClosingTag        BACnetClosingTag
}

var _ BACnetFaultParameterFaultState = (*_BACnetFaultParameterFaultState)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetFaultParameterFaultState) GetParent() BACnetFaultParameterContract {
	return m.BACnetFaultParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultState) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetFaultParameterFaultState) GetListOfFaultValues() BACnetFaultParameterFaultStateListOfFaultValues {
	return m.ListOfFaultValues
}

func (m *_BACnetFaultParameterFaultState) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultState factory function for _BACnetFaultParameterFaultState
func NewBACnetFaultParameterFaultState(openingTag BACnetOpeningTag, listOfFaultValues BACnetFaultParameterFaultStateListOfFaultValues, closingTag BACnetClosingTag, peekedTagHeader BACnetTagHeader) *_BACnetFaultParameterFaultState {
	_result := &_BACnetFaultParameterFaultState{
		BACnetFaultParameterContract: NewBACnetFaultParameter(peekedTagHeader),
		OpeningTag:                   openingTag,
		ListOfFaultValues:            listOfFaultValues,
		ClosingTag:                   closingTag,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultState(structType any) BACnetFaultParameterFaultState {
	if casted, ok := structType.(BACnetFaultParameterFaultState); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultState); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultState) GetTypeName() string {
	return "BACnetFaultParameterFaultState"
}

func (m *_BACnetFaultParameterFaultState) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetFaultParameterContract.(*_BACnetFaultParameter).getLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (listOfFaultValues)
	lengthInBits += m.ListOfFaultValues.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultState) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetFaultParameterFaultState) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetFaultParameter) (_bACnetFaultParameterFaultState BACnetFaultParameterFaultState, err error) {
	m.BACnetFaultParameterContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(4))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	listOfFaultValues, err := ReadSimpleField[BACnetFaultParameterFaultStateListOfFaultValues](ctx, "listOfFaultValues", ReadComplex[BACnetFaultParameterFaultStateListOfFaultValues](BACnetFaultParameterFaultStateListOfFaultValuesParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfFaultValues' field"))
	}
	m.ListOfFaultValues = listOfFaultValues

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(4))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultState")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultState) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultState")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'openingTag' field")
		}

		if err := WriteSimpleField[BACnetFaultParameterFaultStateListOfFaultValues](ctx, "listOfFaultValues", m.GetListOfFaultValues(), WriteComplex[BACnetFaultParameterFaultStateListOfFaultValues](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'listOfFaultValues' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultState")
		}
		return nil
	}
	return m.BACnetFaultParameterContract.(*_BACnetFaultParameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultState) isBACnetFaultParameterFaultState() bool {
	return true
}

func (m *_BACnetFaultParameterFaultState) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
