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

// BACnetConstructedDataProtocolLevel is the corresponding interface of BACnetConstructedDataProtocolLevel
type BACnetConstructedDataProtocolLevel interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetProtocolLevel returns ProtocolLevel (property field)
	GetProtocolLevel() BACnetProtocolLevelTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetProtocolLevelTagged
}

// BACnetConstructedDataProtocolLevelExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataProtocolLevel.
// This is useful for switch cases.
type BACnetConstructedDataProtocolLevelExactly interface {
	BACnetConstructedDataProtocolLevel
	isBACnetConstructedDataProtocolLevel() bool
}

// _BACnetConstructedDataProtocolLevel is the data-structure of this message
type _BACnetConstructedDataProtocolLevel struct {
	*_BACnetConstructedData
	ProtocolLevel BACnetProtocolLevelTagged
}

var _ BACnetConstructedDataProtocolLevel = (*_BACnetConstructedDataProtocolLevel)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataProtocolLevel) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataProtocolLevel) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PROTOCOL_LEVEL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataProtocolLevel) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataProtocolLevel) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataProtocolLevel) GetProtocolLevel() BACnetProtocolLevelTagged {
	return m.ProtocolLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataProtocolLevel) GetActualValue() BACnetProtocolLevelTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetProtocolLevelTagged(m.GetProtocolLevel())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataProtocolLevel factory function for _BACnetConstructedDataProtocolLevel
func NewBACnetConstructedDataProtocolLevel(protocolLevel BACnetProtocolLevelTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataProtocolLevel {
	_result := &_BACnetConstructedDataProtocolLevel{
		ProtocolLevel:          protocolLevel,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataProtocolLevel(structType any) BACnetConstructedDataProtocolLevel {
	if casted, ok := structType.(BACnetConstructedDataProtocolLevel); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataProtocolLevel); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataProtocolLevel) GetTypeName() string {
	return "BACnetConstructedDataProtocolLevel"
}

func (m *_BACnetConstructedDataProtocolLevel) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (protocolLevel)
	lengthInBits += m.ProtocolLevel.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataProtocolLevel) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataProtocolLevelParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataProtocolLevel, error) {
	return BACnetConstructedDataProtocolLevelParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataProtocolLevelParseWithBufferProducer(tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataProtocolLevel, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataProtocolLevel, error) {
		return BACnetConstructedDataProtocolLevelParseWithBuffer(ctx, readBuffer, tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
	}
}

func BACnetConstructedDataProtocolLevelParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataProtocolLevel, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataProtocolLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataProtocolLevel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	protocolLevel, err := ReadSimpleField[BACnetProtocolLevelTagged](ctx, "protocolLevel", ReadComplex[BACnetProtocolLevelTagged](BACnetProtocolLevelTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'protocolLevel' field"))
	}

	actualValue, err := ReadVirtualField[BACnetProtocolLevelTagged](ctx, "actualValue", (*BACnetProtocolLevelTagged)(nil), protocolLevel)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataProtocolLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataProtocolLevel")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataProtocolLevel{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ProtocolLevel: protocolLevel,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataProtocolLevel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataProtocolLevel) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataProtocolLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataProtocolLevel")
		}

		if err := WriteSimpleField[BACnetProtocolLevelTagged](ctx, "protocolLevel", m.GetProtocolLevel(), WriteComplex[BACnetProtocolLevelTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'protocolLevel' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataProtocolLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataProtocolLevel")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataProtocolLevel) isBACnetConstructedDataProtocolLevel() bool {
	return true
}

func (m *_BACnetConstructedDataProtocolLevel) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
