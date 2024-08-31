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

// BACnetContextTag is the corresponding interface of BACnetContextTag
type BACnetContextTag interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetDataType returns DataType (discriminator field)
	GetDataType() BACnetDataType
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetTagNumber returns TagNumber (virtual field)
	GetTagNumber() uint8
	// GetActualLength returns ActualLength (virtual field)
	GetActualLength() uint32
}

// BACnetContextTagExactly can be used when we want exactly this type and not a type which fulfills BACnetContextTag.
// This is useful for switch cases.
type BACnetContextTagExactly interface {
	BACnetContextTag
	isBACnetContextTag() bool
}

// _BACnetContextTag is the data-structure of this message
type _BACnetContextTag struct {
	_BACnetContextTagChildRequirements
	Header BACnetTagHeader

	// Arguments.
	TagNumberArgument uint8
}

type _BACnetContextTagChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetDataType() BACnetDataType
}

type BACnetContextTagParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetContextTag, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetContextTagChild interface {
	utils.Serializable
	InitializeParent(parent BACnetContextTag, header BACnetTagHeader)
	GetParent() *BACnetContextTag

	GetTypeName() string
	BACnetContextTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetContextTag) GetHeader() BACnetTagHeader {
	return m.Header
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetContextTag) GetTagNumber() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetHeader().GetTagNumber())
}

func (m *_BACnetContextTag) GetActualLength() uint32 {
	ctx := context.Background()
	_ = ctx
	return uint32(m.GetHeader().GetActualLength())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetContextTag factory function for _BACnetContextTag
func NewBACnetContextTag(header BACnetTagHeader, tagNumberArgument uint8) *_BACnetContextTag {
	return &_BACnetContextTag{Header: header, TagNumberArgument: tagNumberArgument}
}

// Deprecated: use the interface for direct cast
func CastBACnetContextTag(structType any) BACnetContextTag {
	if casted, ok := structType.(BACnetContextTag); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetContextTag); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetContextTag) GetTypeName() string {
	return "BACnetContextTag"
}

func (m *_BACnetContextTag) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetContextTag) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetContextTagParse(ctx context.Context, theBytes []byte, tagNumberArgument uint8, dataType BACnetDataType) (BACnetContextTag, error) {
	return BACnetContextTagParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumberArgument, dataType)
}

func BACnetContextTagParseWithBufferProducer[T BACnetContextTag](tagNumberArgument uint8, dataType BACnetDataType) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		buffer, err := BACnetContextTagParseWithBuffer(ctx, readBuffer, tagNumberArgument, dataType)
		if err != nil {
			var zero T
			return zero, err
		}
		return buffer.(T), err
	}
}

func BACnetContextTagParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType) (BACnetContextTag, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetContextTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTag")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := ReadSimpleField[BACnetTagHeader](ctx, "header", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}

	// Validation
	if !(bool((header.GetActualTagNumber()) == (tagNumberArgument))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "tagnumber doesn't match"})
	}

	// Validation
	if !(bool((header.GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "should be a context tag"})
	}

	tagNumber, err := ReadVirtualField[uint8](ctx, "tagNumber", (*uint8)(nil), header.GetTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'tagNumber' field"))
	}

	actualLength, err := ReadVirtualField[uint32](ctx, "actualLength", (*uint32)(nil), header.GetActualLength())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualLength' field"))
	}

	// Validation
	if !(bool(bool((header.GetLengthValueType()) != (6))) && bool(bool((header.GetLengthValueType()) != (7)))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "length 6 and 7 reserved for opening and closing tag"})
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetContextTagChildSerializeRequirement interface {
		BACnetContextTag
		InitializeParent(BACnetContextTag, BACnetTagHeader)
		GetParent() BACnetContextTag
	}
	var _childTemp any
	var _child BACnetContextTagChildSerializeRequirement
	var typeSwitchError error
	switch {
	case dataType == BACnetDataType_NULL: // BACnetContextTagNull
		_childTemp, typeSwitchError = BACnetContextTagNullParseWithBuffer(ctx, readBuffer, header, tagNumberArgument, dataType)
	case dataType == BACnetDataType_BOOLEAN: // BACnetContextTagBoolean
		_childTemp, typeSwitchError = BACnetContextTagBooleanParseWithBuffer(ctx, readBuffer, header, tagNumberArgument, dataType)
	case dataType == BACnetDataType_UNSIGNED_INTEGER: // BACnetContextTagUnsignedInteger
		_childTemp, typeSwitchError = BACnetContextTagUnsignedIntegerParseWithBuffer(ctx, readBuffer, header, tagNumberArgument, dataType)
	case dataType == BACnetDataType_SIGNED_INTEGER: // BACnetContextTagSignedInteger
		_childTemp, typeSwitchError = BACnetContextTagSignedIntegerParseWithBuffer(ctx, readBuffer, header, tagNumberArgument, dataType)
	case dataType == BACnetDataType_REAL: // BACnetContextTagReal
		_childTemp, typeSwitchError = BACnetContextTagRealParseWithBuffer(ctx, readBuffer, tagNumberArgument, dataType)
	case dataType == BACnetDataType_DOUBLE: // BACnetContextTagDouble
		_childTemp, typeSwitchError = BACnetContextTagDoubleParseWithBuffer(ctx, readBuffer, tagNumberArgument, dataType)
	case dataType == BACnetDataType_OCTET_STRING: // BACnetContextTagOctetString
		_childTemp, typeSwitchError = BACnetContextTagOctetStringParseWithBuffer(ctx, readBuffer, header, tagNumberArgument, dataType)
	case dataType == BACnetDataType_CHARACTER_STRING: // BACnetContextTagCharacterString
		_childTemp, typeSwitchError = BACnetContextTagCharacterStringParseWithBuffer(ctx, readBuffer, header, tagNumberArgument, dataType)
	case dataType == BACnetDataType_BIT_STRING: // BACnetContextTagBitString
		_childTemp, typeSwitchError = BACnetContextTagBitStringParseWithBuffer(ctx, readBuffer, header, tagNumberArgument, dataType)
	case dataType == BACnetDataType_ENUMERATED: // BACnetContextTagEnumerated
		_childTemp, typeSwitchError = BACnetContextTagEnumeratedParseWithBuffer(ctx, readBuffer, header, tagNumberArgument, dataType)
	case dataType == BACnetDataType_DATE: // BACnetContextTagDate
		_childTemp, typeSwitchError = BACnetContextTagDateParseWithBuffer(ctx, readBuffer, tagNumberArgument, dataType)
	case dataType == BACnetDataType_TIME: // BACnetContextTagTime
		_childTemp, typeSwitchError = BACnetContextTagTimeParseWithBuffer(ctx, readBuffer, tagNumberArgument, dataType)
	case dataType == BACnetDataType_BACNET_OBJECT_IDENTIFIER: // BACnetContextTagObjectIdentifier
		_childTemp, typeSwitchError = BACnetContextTagObjectIdentifierParseWithBuffer(ctx, readBuffer, tagNumberArgument, dataType)
	case dataType == BACnetDataType_UNKNOWN: // BACnetContextTagUnknown
		_childTemp, typeSwitchError = BACnetContextTagUnknownParseWithBuffer(ctx, readBuffer, actualLength, tagNumberArgument, dataType)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [dataType=%v]", dataType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetContextTag")
	}
	_child = _childTemp.(BACnetContextTagChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("BACnetContextTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTag")
	}

	// Finish initializing
	_child.InitializeParent(_child, header)
	return _child, nil
}

func (pm *_BACnetContextTag) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetContextTag, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetContextTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetContextTag")
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for header")
	}
	_headerErr := writeBuffer.WriteSerializable(ctx, m.GetHeader())
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for header")
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
	}
	// Virtual field
	tagNumber := m.GetTagNumber()
	_ = tagNumber
	if _tagNumberErr := writeBuffer.WriteVirtual(ctx, "tagNumber", m.GetTagNumber()); _tagNumberErr != nil {
		return errors.Wrap(_tagNumberErr, "Error serializing 'tagNumber' field")
	}
	// Virtual field
	actualLength := m.GetActualLength()
	_ = actualLength
	if _actualLengthErr := writeBuffer.WriteVirtual(ctx, "actualLength", m.GetActualLength()); _actualLengthErr != nil {
		return errors.Wrap(_actualLengthErr, "Error serializing 'actualLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetContextTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetContextTag")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetContextTag) GetTagNumberArgument() uint8 {
	return m.TagNumberArgument
}

//
////

func (m *_BACnetContextTag) isBACnetContextTag() bool {
	return true
}

func (m *_BACnetContextTag) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
