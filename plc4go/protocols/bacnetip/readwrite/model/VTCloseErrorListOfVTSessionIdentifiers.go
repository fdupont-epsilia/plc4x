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

// VTCloseErrorListOfVTSessionIdentifiers is the corresponding interface of VTCloseErrorListOfVTSessionIdentifiers
type VTCloseErrorListOfVTSessionIdentifiers interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfVtSessionIdentifiers returns ListOfVtSessionIdentifiers (property field)
	GetListOfVtSessionIdentifiers() []BACnetApplicationTagUnsignedInteger
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsVTCloseErrorListOfVTSessionIdentifiers is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsVTCloseErrorListOfVTSessionIdentifiers()
}

// _VTCloseErrorListOfVTSessionIdentifiers is the data-structure of this message
type _VTCloseErrorListOfVTSessionIdentifiers struct {
	OpeningTag                 BACnetOpeningTag
	ListOfVtSessionIdentifiers []BACnetApplicationTagUnsignedInteger
	ClosingTag                 BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ VTCloseErrorListOfVTSessionIdentifiers = (*_VTCloseErrorListOfVTSessionIdentifiers)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_VTCloseErrorListOfVTSessionIdentifiers) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_VTCloseErrorListOfVTSessionIdentifiers) GetListOfVtSessionIdentifiers() []BACnetApplicationTagUnsignedInteger {
	return m.ListOfVtSessionIdentifiers
}

func (m *_VTCloseErrorListOfVTSessionIdentifiers) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewVTCloseErrorListOfVTSessionIdentifiers factory function for _VTCloseErrorListOfVTSessionIdentifiers
func NewVTCloseErrorListOfVTSessionIdentifiers(openingTag BACnetOpeningTag, listOfVtSessionIdentifiers []BACnetApplicationTagUnsignedInteger, closingTag BACnetClosingTag, tagNumber uint8) *_VTCloseErrorListOfVTSessionIdentifiers {
	return &_VTCloseErrorListOfVTSessionIdentifiers{OpeningTag: openingTag, ListOfVtSessionIdentifiers: listOfVtSessionIdentifiers, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastVTCloseErrorListOfVTSessionIdentifiers(structType any) VTCloseErrorListOfVTSessionIdentifiers {
	if casted, ok := structType.(VTCloseErrorListOfVTSessionIdentifiers); ok {
		return casted
	}
	if casted, ok := structType.(*VTCloseErrorListOfVTSessionIdentifiers); ok {
		return *casted
	}
	return nil
}

func (m *_VTCloseErrorListOfVTSessionIdentifiers) GetTypeName() string {
	return "VTCloseErrorListOfVTSessionIdentifiers"
}

func (m *_VTCloseErrorListOfVTSessionIdentifiers) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.ListOfVtSessionIdentifiers) > 0 {
		for _, element := range m.ListOfVtSessionIdentifiers {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_VTCloseErrorListOfVTSessionIdentifiers) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func VTCloseErrorListOfVTSessionIdentifiersParse(ctx context.Context, theBytes []byte, tagNumber uint8) (VTCloseErrorListOfVTSessionIdentifiers, error) {
	return VTCloseErrorListOfVTSessionIdentifiersParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func VTCloseErrorListOfVTSessionIdentifiersParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (VTCloseErrorListOfVTSessionIdentifiers, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (VTCloseErrorListOfVTSessionIdentifiers, error) {
		return VTCloseErrorListOfVTSessionIdentifiersParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func VTCloseErrorListOfVTSessionIdentifiersParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (VTCloseErrorListOfVTSessionIdentifiers, error) {
	v, err := (&_VTCloseErrorListOfVTSessionIdentifiers{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_VTCloseErrorListOfVTSessionIdentifiers) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__vTCloseErrorListOfVTSessionIdentifiers VTCloseErrorListOfVTSessionIdentifiers, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("VTCloseErrorListOfVTSessionIdentifiers"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for VTCloseErrorListOfVTSessionIdentifiers")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	listOfVtSessionIdentifiers, err := ReadTerminatedArrayField[BACnetApplicationTagUnsignedInteger](ctx, "listOfVtSessionIdentifiers", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, 1))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfVtSessionIdentifiers' field"))
	}
	m.ListOfVtSessionIdentifiers = listOfVtSessionIdentifiers

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("VTCloseErrorListOfVTSessionIdentifiers"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for VTCloseErrorListOfVTSessionIdentifiers")
	}

	return m, nil
}

func (m *_VTCloseErrorListOfVTSessionIdentifiers) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_VTCloseErrorListOfVTSessionIdentifiers) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("VTCloseErrorListOfVTSessionIdentifiers"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for VTCloseErrorListOfVTSessionIdentifiers")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "listOfVtSessionIdentifiers", m.GetListOfVtSessionIdentifiers(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'listOfVtSessionIdentifiers' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("VTCloseErrorListOfVTSessionIdentifiers"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for VTCloseErrorListOfVTSessionIdentifiers")
	}
	return nil
}

////
// Arguments Getter

func (m *_VTCloseErrorListOfVTSessionIdentifiers) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_VTCloseErrorListOfVTSessionIdentifiers) IsVTCloseErrorListOfVTSessionIdentifiers() {}

func (m *_VTCloseErrorListOfVTSessionIdentifiers) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
