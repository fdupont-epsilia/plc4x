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

// BACnetEventParameterExtended is the corresponding interface of BACnetEventParameterExtended
type BACnetEventParameterExtended interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetVendorId returns VendorId (property field)
	GetVendorId() BACnetVendorIdTagged
	// GetExtendedEventType returns ExtendedEventType (property field)
	GetExtendedEventType() BACnetContextTagUnsignedInteger
	// GetParameters returns Parameters (property field)
	GetParameters() BACnetEventParameterExtendedParameters
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetEventParameterExtendedExactly can be used when we want exactly this type and not a type which fulfills BACnetEventParameterExtended.
// This is useful for switch cases.
type BACnetEventParameterExtendedExactly interface {
	BACnetEventParameterExtended
	isBACnetEventParameterExtended() bool
}

// _BACnetEventParameterExtended is the data-structure of this message
type _BACnetEventParameterExtended struct {
	BACnetEventParameterContract
	OpeningTag        BACnetOpeningTag
	VendorId          BACnetVendorIdTagged
	ExtendedEventType BACnetContextTagUnsignedInteger
	Parameters        BACnetEventParameterExtendedParameters
	ClosingTag        BACnetClosingTag
}

var _ BACnetEventParameterExtended = (*_BACnetEventParameterExtended)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetEventParameterExtended) GetParent() BACnetEventParameterContract {
	return m.BACnetEventParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterExtended) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterExtended) GetVendorId() BACnetVendorIdTagged {
	return m.VendorId
}

func (m *_BACnetEventParameterExtended) GetExtendedEventType() BACnetContextTagUnsignedInteger {
	return m.ExtendedEventType
}

func (m *_BACnetEventParameterExtended) GetParameters() BACnetEventParameterExtendedParameters {
	return m.Parameters
}

func (m *_BACnetEventParameterExtended) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterExtended factory function for _BACnetEventParameterExtended
func NewBACnetEventParameterExtended(openingTag BACnetOpeningTag, vendorId BACnetVendorIdTagged, extendedEventType BACnetContextTagUnsignedInteger, parameters BACnetEventParameterExtendedParameters, closingTag BACnetClosingTag, peekedTagHeader BACnetTagHeader) *_BACnetEventParameterExtended {
	_result := &_BACnetEventParameterExtended{
		BACnetEventParameterContract: NewBACnetEventParameter(peekedTagHeader),
		OpeningTag:                   openingTag,
		VendorId:                     vendorId,
		ExtendedEventType:            extendedEventType,
		Parameters:                   parameters,
		ClosingTag:                   closingTag,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterExtended(structType any) BACnetEventParameterExtended {
	if casted, ok := structType.(BACnetEventParameterExtended); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterExtended); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterExtended) GetTypeName() string {
	return "BACnetEventParameterExtended"
}

func (m *_BACnetEventParameterExtended) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetEventParameterContract.(*_BACnetEventParameter).getLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (vendorId)
	lengthInBits += m.VendorId.GetLengthInBits(ctx)

	// Simple field (extendedEventType)
	lengthInBits += m.ExtendedEventType.GetLengthInBits(ctx)

	// Simple field (parameters)
	lengthInBits += m.Parameters.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterExtended) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetEventParameterExtended) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetEventParameter) (_bACnetEventParameterExtended BACnetEventParameterExtended, err error) {
	m.BACnetEventParameterContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterExtended"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterExtended")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(9))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	vendorId, err := ReadSimpleField[BACnetVendorIdTagged](ctx, "vendorId", ReadComplex[BACnetVendorIdTagged](BACnetVendorIdTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'vendorId' field"))
	}
	m.VendorId = vendorId

	extendedEventType, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "extendedEventType", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'extendedEventType' field"))
	}
	m.ExtendedEventType = extendedEventType

	parameters, err := ReadSimpleField[BACnetEventParameterExtendedParameters](ctx, "parameters", ReadComplex[BACnetEventParameterExtendedParameters](BACnetEventParameterExtendedParametersParseWithBufferProducer((uint8)(uint8(2))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'parameters' field"))
	}
	m.Parameters = parameters

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(9))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetEventParameterExtended"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterExtended")
	}

	return m, nil
}

func (m *_BACnetEventParameterExtended) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterExtended) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterExtended"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterExtended")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'openingTag' field")
		}

		if err := WriteSimpleField[BACnetVendorIdTagged](ctx, "vendorId", m.GetVendorId(), WriteComplex[BACnetVendorIdTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'vendorId' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "extendedEventType", m.GetExtendedEventType(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'extendedEventType' field")
		}

		if err := WriteSimpleField[BACnetEventParameterExtendedParameters](ctx, "parameters", m.GetParameters(), WriteComplex[BACnetEventParameterExtendedParameters](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'parameters' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventParameterExtended"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterExtended")
		}
		return nil
	}
	return m.BACnetEventParameterContract.(*_BACnetEventParameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetEventParameterExtended) isBACnetEventParameterExtended() bool {
	return true
}

func (m *_BACnetEventParameterExtended) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
