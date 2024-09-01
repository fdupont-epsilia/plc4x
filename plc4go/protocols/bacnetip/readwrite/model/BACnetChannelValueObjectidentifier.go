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

// BACnetChannelValueObjectidentifier is the corresponding interface of BACnetChannelValueObjectidentifier
type BACnetChannelValueObjectidentifier interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetChannelValue
	// GetObjectidentifierValue returns ObjectidentifierValue (property field)
	GetObjectidentifierValue() BACnetApplicationTagObjectIdentifier
}

// BACnetChannelValueObjectidentifierExactly can be used when we want exactly this type and not a type which fulfills BACnetChannelValueObjectidentifier.
// This is useful for switch cases.
type BACnetChannelValueObjectidentifierExactly interface {
	BACnetChannelValueObjectidentifier
	isBACnetChannelValueObjectidentifier() bool
}

// _BACnetChannelValueObjectidentifier is the data-structure of this message
type _BACnetChannelValueObjectidentifier struct {
	BACnetChannelValueContract
	ObjectidentifierValue BACnetApplicationTagObjectIdentifier
}

var _ BACnetChannelValueObjectidentifier = (*_BACnetChannelValueObjectidentifier)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetChannelValueObjectidentifier) GetParent() BACnetChannelValueContract {
	return m.BACnetChannelValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetChannelValueObjectidentifier) GetObjectidentifierValue() BACnetApplicationTagObjectIdentifier {
	return m.ObjectidentifierValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetChannelValueObjectidentifier factory function for _BACnetChannelValueObjectidentifier
func NewBACnetChannelValueObjectidentifier(objectidentifierValue BACnetApplicationTagObjectIdentifier, peekedTagHeader BACnetTagHeader) *_BACnetChannelValueObjectidentifier {
	_result := &_BACnetChannelValueObjectidentifier{
		BACnetChannelValueContract: NewBACnetChannelValue(peekedTagHeader),
		ObjectidentifierValue:      objectidentifierValue,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetChannelValueObjectidentifier(structType any) BACnetChannelValueObjectidentifier {
	if casted, ok := structType.(BACnetChannelValueObjectidentifier); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetChannelValueObjectidentifier); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetChannelValueObjectidentifier) GetTypeName() string {
	return "BACnetChannelValueObjectidentifier"
}

func (m *_BACnetChannelValueObjectidentifier) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetChannelValueContract.(*_BACnetChannelValue).getLengthInBits(ctx))

	// Simple field (objectidentifierValue)
	lengthInBits += m.ObjectidentifierValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetChannelValueObjectidentifier) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetChannelValueObjectidentifier) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetChannelValue) (_bACnetChannelValueObjectidentifier BACnetChannelValueObjectidentifier, err error) {
	m.BACnetChannelValueContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValueObjectidentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetChannelValueObjectidentifier")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectidentifierValue, err := ReadSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "objectidentifierValue", ReadComplex[BACnetApplicationTagObjectIdentifier](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagObjectIdentifier](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectidentifierValue' field"))
	}
	m.ObjectidentifierValue = objectidentifierValue

	if closeErr := readBuffer.CloseContext("BACnetChannelValueObjectidentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetChannelValueObjectidentifier")
	}

	return m, nil
}

func (m *_BACnetChannelValueObjectidentifier) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetChannelValueObjectidentifier) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueObjectidentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetChannelValueObjectidentifier")
		}

		if err := WriteSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "objectidentifierValue", m.GetObjectidentifierValue(), WriteComplex[BACnetApplicationTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'objectidentifierValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetChannelValueObjectidentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetChannelValueObjectidentifier")
		}
		return nil
	}
	return m.BACnetChannelValueContract.(*_BACnetChannelValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetChannelValueObjectidentifier) isBACnetChannelValueObjectidentifier() bool {
	return true
}

func (m *_BACnetChannelValueObjectidentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
