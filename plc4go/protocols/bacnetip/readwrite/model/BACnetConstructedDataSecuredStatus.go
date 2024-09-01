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

// BACnetConstructedDataSecuredStatus is the corresponding interface of BACnetConstructedDataSecuredStatus
type BACnetConstructedDataSecuredStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetSecuredStatus returns SecuredStatus (property field)
	GetSecuredStatus() BACnetDoorSecuredStatusTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDoorSecuredStatusTagged
}

// BACnetConstructedDataSecuredStatusExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataSecuredStatus.
// This is useful for switch cases.
type BACnetConstructedDataSecuredStatusExactly interface {
	BACnetConstructedDataSecuredStatus
	isBACnetConstructedDataSecuredStatus() bool
}

// _BACnetConstructedDataSecuredStatus is the data-structure of this message
type _BACnetConstructedDataSecuredStatus struct {
	*_BACnetConstructedData
	SecuredStatus BACnetDoorSecuredStatusTagged
}

var _ BACnetConstructedDataSecuredStatus = (*_BACnetConstructedDataSecuredStatus)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSecuredStatus) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataSecuredStatus) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SECURED_STATUS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSecuredStatus) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataSecuredStatus) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSecuredStatus) GetSecuredStatus() BACnetDoorSecuredStatusTagged {
	return m.SecuredStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataSecuredStatus) GetActualValue() BACnetDoorSecuredStatusTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDoorSecuredStatusTagged(m.GetSecuredStatus())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataSecuredStatus factory function for _BACnetConstructedDataSecuredStatus
func NewBACnetConstructedDataSecuredStatus(securedStatus BACnetDoorSecuredStatusTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataSecuredStatus {
	_result := &_BACnetConstructedDataSecuredStatus{
		SecuredStatus:          securedStatus,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataSecuredStatus(structType any) BACnetConstructedDataSecuredStatus {
	if casted, ok := structType.(BACnetConstructedDataSecuredStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSecuredStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSecuredStatus) GetTypeName() string {
	return "BACnetConstructedDataSecuredStatus"
}

func (m *_BACnetConstructedDataSecuredStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (securedStatus)
	lengthInBits += m.SecuredStatus.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataSecuredStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataSecuredStatusParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataSecuredStatus, error) {
	return BACnetConstructedDataSecuredStatusParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataSecuredStatusParseWithBufferProducer(tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataSecuredStatus, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataSecuredStatus, error) {
		return BACnetConstructedDataSecuredStatusParseWithBuffer(ctx, readBuffer, tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
	}
}

func BACnetConstructedDataSecuredStatusParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataSecuredStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSecuredStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSecuredStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	securedStatus, err := ReadSimpleField[BACnetDoorSecuredStatusTagged](ctx, "securedStatus", ReadComplex[BACnetDoorSecuredStatusTagged](BACnetDoorSecuredStatusTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securedStatus' field"))
	}

	actualValue, err := ReadVirtualField[BACnetDoorSecuredStatusTagged](ctx, "actualValue", (*BACnetDoorSecuredStatusTagged)(nil), securedStatus)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSecuredStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSecuredStatus")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataSecuredStatus{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		SecuredStatus: securedStatus,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataSecuredStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataSecuredStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSecuredStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSecuredStatus")
		}

		if err := WriteSimpleField[BACnetDoorSecuredStatusTagged](ctx, "securedStatus", m.GetSecuredStatus(), WriteComplex[BACnetDoorSecuredStatusTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'securedStatus' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSecuredStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSecuredStatus")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataSecuredStatus) isBACnetConstructedDataSecuredStatus() bool {
	return true
}

func (m *_BACnetConstructedDataSecuredStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
