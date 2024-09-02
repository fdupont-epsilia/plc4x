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

// BACnetProcessIdSelectionValue is the corresponding interface of BACnetProcessIdSelectionValue
type BACnetProcessIdSelectionValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetProcessIdSelection
	// GetProcessIdentifier returns ProcessIdentifier (property field)
	GetProcessIdentifier() BACnetApplicationTagUnsignedInteger
	// IsBACnetProcessIdSelectionValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetProcessIdSelectionValue()
}

// _BACnetProcessIdSelectionValue is the data-structure of this message
type _BACnetProcessIdSelectionValue struct {
	BACnetProcessIdSelectionContract
	ProcessIdentifier BACnetApplicationTagUnsignedInteger
}

var _ BACnetProcessIdSelectionValue = (*_BACnetProcessIdSelectionValue)(nil)
var _ BACnetProcessIdSelectionRequirements = (*_BACnetProcessIdSelectionValue)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetProcessIdSelectionValue) GetParent() BACnetProcessIdSelectionContract {
	return m.BACnetProcessIdSelectionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetProcessIdSelectionValue) GetProcessIdentifier() BACnetApplicationTagUnsignedInteger {
	return m.ProcessIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetProcessIdSelectionValue factory function for _BACnetProcessIdSelectionValue
func NewBACnetProcessIdSelectionValue(processIdentifier BACnetApplicationTagUnsignedInteger, peekedTagHeader BACnetTagHeader) *_BACnetProcessIdSelectionValue {
	_result := &_BACnetProcessIdSelectionValue{
		BACnetProcessIdSelectionContract: NewBACnetProcessIdSelection(peekedTagHeader),
		ProcessIdentifier:                processIdentifier,
	}
	_result.BACnetProcessIdSelectionContract.(*_BACnetProcessIdSelection)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetProcessIdSelectionValue(structType any) BACnetProcessIdSelectionValue {
	if casted, ok := structType.(BACnetProcessIdSelectionValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetProcessIdSelectionValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetProcessIdSelectionValue) GetTypeName() string {
	return "BACnetProcessIdSelectionValue"
}

func (m *_BACnetProcessIdSelectionValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetProcessIdSelectionContract.(*_BACnetProcessIdSelection).getLengthInBits(ctx))

	// Simple field (processIdentifier)
	lengthInBits += m.ProcessIdentifier.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetProcessIdSelectionValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetProcessIdSelectionValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetProcessIdSelection) (__bACnetProcessIdSelectionValue BACnetProcessIdSelectionValue, err error) {
	m.BACnetProcessIdSelectionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetProcessIdSelectionValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetProcessIdSelectionValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	processIdentifier, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "processIdentifier", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'processIdentifier' field"))
	}
	m.ProcessIdentifier = processIdentifier

	if closeErr := readBuffer.CloseContext("BACnetProcessIdSelectionValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetProcessIdSelectionValue")
	}

	return m, nil
}

func (m *_BACnetProcessIdSelectionValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetProcessIdSelectionValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetProcessIdSelectionValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetProcessIdSelectionValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "processIdentifier", m.GetProcessIdentifier(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'processIdentifier' field")
		}

		if popErr := writeBuffer.PopContext("BACnetProcessIdSelectionValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetProcessIdSelectionValue")
		}
		return nil
	}
	return m.BACnetProcessIdSelectionContract.(*_BACnetProcessIdSelection).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetProcessIdSelectionValue) IsBACnetProcessIdSelectionValue() {}

func (m *_BACnetProcessIdSelectionValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
