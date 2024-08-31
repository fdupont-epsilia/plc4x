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

// BACnetWriteAccessSpecification is the corresponding interface of BACnetWriteAccessSpecification
type BACnetWriteAccessSpecification interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfPropertyWriteDefinition returns ListOfPropertyWriteDefinition (property field)
	GetListOfPropertyWriteDefinition() []BACnetPropertyWriteDefinition
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetWriteAccessSpecificationExactly can be used when we want exactly this type and not a type which fulfills BACnetWriteAccessSpecification.
// This is useful for switch cases.
type BACnetWriteAccessSpecificationExactly interface {
	BACnetWriteAccessSpecification
	isBACnetWriteAccessSpecification() bool
}

// _BACnetWriteAccessSpecification is the data-structure of this message
type _BACnetWriteAccessSpecification struct {
	ObjectIdentifier              BACnetContextTagObjectIdentifier
	OpeningTag                    BACnetOpeningTag
	ListOfPropertyWriteDefinition []BACnetPropertyWriteDefinition
	ClosingTag                    BACnetClosingTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetWriteAccessSpecification) GetObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetWriteAccessSpecification) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetWriteAccessSpecification) GetListOfPropertyWriteDefinition() []BACnetPropertyWriteDefinition {
	return m.ListOfPropertyWriteDefinition
}

func (m *_BACnetWriteAccessSpecification) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetWriteAccessSpecification factory function for _BACnetWriteAccessSpecification
func NewBACnetWriteAccessSpecification(objectIdentifier BACnetContextTagObjectIdentifier, openingTag BACnetOpeningTag, listOfPropertyWriteDefinition []BACnetPropertyWriteDefinition, closingTag BACnetClosingTag) *_BACnetWriteAccessSpecification {
	return &_BACnetWriteAccessSpecification{ObjectIdentifier: objectIdentifier, OpeningTag: openingTag, ListOfPropertyWriteDefinition: listOfPropertyWriteDefinition, ClosingTag: closingTag}
}

// Deprecated: use the interface for direct cast
func CastBACnetWriteAccessSpecification(structType any) BACnetWriteAccessSpecification {
	if casted, ok := structType.(BACnetWriteAccessSpecification); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetWriteAccessSpecification); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetWriteAccessSpecification) GetTypeName() string {
	return "BACnetWriteAccessSpecification"
}

func (m *_BACnetWriteAccessSpecification) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.ListOfPropertyWriteDefinition) > 0 {
		for _, element := range m.ListOfPropertyWriteDefinition {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetWriteAccessSpecification) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetWriteAccessSpecificationParse(ctx context.Context, theBytes []byte) (BACnetWriteAccessSpecification, error) {
	return BACnetWriteAccessSpecificationParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetWriteAccessSpecificationParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetWriteAccessSpecification, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetWriteAccessSpecification, error) {
		return BACnetWriteAccessSpecificationParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetWriteAccessSpecificationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetWriteAccessSpecification, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetWriteAccessSpecification"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetWriteAccessSpecification")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "objectIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectIdentifier' field"))
	}

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}

	listOfPropertyWriteDefinition, err := ReadTerminatedArrayField[BACnetPropertyWriteDefinition](ctx, "listOfPropertyWriteDefinition", ReadComplex[BACnetPropertyWriteDefinition](BACnetPropertyWriteDefinitionParseWithBufferProducer((BACnetObjectType)(objectIdentifier.GetObjectType())), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, 1))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfPropertyWriteDefinition' field"))
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetWriteAccessSpecification"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetWriteAccessSpecification")
	}

	// Create the instance
	return &_BACnetWriteAccessSpecification{
		ObjectIdentifier:              objectIdentifier,
		OpeningTag:                    openingTag,
		ListOfPropertyWriteDefinition: listOfPropertyWriteDefinition,
		ClosingTag:                    closingTag,
	}, nil
}

func (m *_BACnetWriteAccessSpecification) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetWriteAccessSpecification) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetWriteAccessSpecification"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetWriteAccessSpecification")
	}

	// Simple Field (objectIdentifier)
	if pushErr := writeBuffer.PushContext("objectIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for objectIdentifier")
	}
	_objectIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetObjectIdentifier())
	if popErr := writeBuffer.PopContext("objectIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for objectIdentifier")
	}
	if _objectIdentifierErr != nil {
		return errors.Wrap(_objectIdentifierErr, "Error serializing 'objectIdentifier' field")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(ctx, m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Array Field (listOfPropertyWriteDefinition)
	if pushErr := writeBuffer.PushContext("listOfPropertyWriteDefinition", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for listOfPropertyWriteDefinition")
	}
	for _curItem, _element := range m.GetListOfPropertyWriteDefinition() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetListOfPropertyWriteDefinition()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'listOfPropertyWriteDefinition' field")
		}
	}
	if popErr := writeBuffer.PopContext("listOfPropertyWriteDefinition", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for listOfPropertyWriteDefinition")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(ctx, m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetWriteAccessSpecification"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetWriteAccessSpecification")
	}
	return nil
}

func (m *_BACnetWriteAccessSpecification) isBACnetWriteAccessSpecification() bool {
	return true
}

func (m *_BACnetWriteAccessSpecification) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
