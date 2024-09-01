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

// BACnetConstructedDataLocalForwardingOnly is the corresponding interface of BACnetConstructedDataLocalForwardingOnly
type BACnetConstructedDataLocalForwardingOnly interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLocalForwardingOnly returns LocalForwardingOnly (property field)
	GetLocalForwardingOnly() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataLocalForwardingOnlyExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLocalForwardingOnly.
// This is useful for switch cases.
type BACnetConstructedDataLocalForwardingOnlyExactly interface {
	BACnetConstructedDataLocalForwardingOnly
	isBACnetConstructedDataLocalForwardingOnly() bool
}

// _BACnetConstructedDataLocalForwardingOnly is the data-structure of this message
type _BACnetConstructedDataLocalForwardingOnly struct {
	BACnetConstructedDataContract
	LocalForwardingOnly BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataLocalForwardingOnly = (*_BACnetConstructedDataLocalForwardingOnly)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLocalForwardingOnly) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLocalForwardingOnly) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOCAL_FORWARDING_ONLY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLocalForwardingOnly) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLocalForwardingOnly) GetLocalForwardingOnly() BACnetApplicationTagBoolean {
	return m.LocalForwardingOnly
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLocalForwardingOnly) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetLocalForwardingOnly())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLocalForwardingOnly factory function for _BACnetConstructedDataLocalForwardingOnly
func NewBACnetConstructedDataLocalForwardingOnly(localForwardingOnly BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLocalForwardingOnly {
	_result := &_BACnetConstructedDataLocalForwardingOnly{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LocalForwardingOnly:           localForwardingOnly,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLocalForwardingOnly(structType any) BACnetConstructedDataLocalForwardingOnly {
	if casted, ok := structType.(BACnetConstructedDataLocalForwardingOnly); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLocalForwardingOnly); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLocalForwardingOnly) GetTypeName() string {
	return "BACnetConstructedDataLocalForwardingOnly"
}

func (m *_BACnetConstructedDataLocalForwardingOnly) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (localForwardingOnly)
	lengthInBits += m.LocalForwardingOnly.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLocalForwardingOnly) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLocalForwardingOnly) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (_bACnetConstructedDataLocalForwardingOnly BACnetConstructedDataLocalForwardingOnly, err error) {
	m.BACnetConstructedDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLocalForwardingOnly"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLocalForwardingOnly")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	localForwardingOnly, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "localForwardingOnly", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'localForwardingOnly' field"))
	}
	m.LocalForwardingOnly = localForwardingOnly

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), localForwardingOnly)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLocalForwardingOnly"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLocalForwardingOnly")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLocalForwardingOnly) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLocalForwardingOnly) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLocalForwardingOnly"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLocalForwardingOnly")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "localForwardingOnly", m.GetLocalForwardingOnly(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'localForwardingOnly' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLocalForwardingOnly"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLocalForwardingOnly")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLocalForwardingOnly) isBACnetConstructedDataLocalForwardingOnly() bool {
	return true
}

func (m *_BACnetConstructedDataLocalForwardingOnly) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
