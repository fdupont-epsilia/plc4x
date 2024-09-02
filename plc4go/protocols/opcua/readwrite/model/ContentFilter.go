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

// ContentFilter is the corresponding interface of ContentFilter
type ContentFilter interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNoOfElements returns NoOfElements (property field)
	GetNoOfElements() int32
	// GetElements returns Elements (property field)
	GetElements() []ExtensionObjectDefinition
	// IsContentFilter is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsContentFilter()
}

// _ContentFilter is the data-structure of this message
type _ContentFilter struct {
	ExtensionObjectDefinitionContract
	NoOfElements int32
	Elements     []ExtensionObjectDefinition
}

var _ ContentFilter = (*_ContentFilter)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ContentFilter)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ContentFilter) GetIdentifier() string {
	return "588"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ContentFilter) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ContentFilter) GetNoOfElements() int32 {
	return m.NoOfElements
}

func (m *_ContentFilter) GetElements() []ExtensionObjectDefinition {
	return m.Elements
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewContentFilter factory function for _ContentFilter
func NewContentFilter(noOfElements int32, elements []ExtensionObjectDefinition) *_ContentFilter {
	_result := &_ContentFilter{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		NoOfElements:                      noOfElements,
		Elements:                          elements,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastContentFilter(structType any) ContentFilter {
	if casted, ok := structType.(ContentFilter); ok {
		return casted
	}
	if casted, ok := structType.(*ContentFilter); ok {
		return *casted
	}
	return nil
}

func (m *_ContentFilter) GetTypeName() string {
	return "ContentFilter"
}

func (m *_ContentFilter) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (noOfElements)
	lengthInBits += 32

	// Array field
	if len(m.Elements) > 0 {
		for _curItem, element := range m.Elements {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Elements), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_ContentFilter) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ContentFilter) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__contentFilter ContentFilter, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ContentFilter"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ContentFilter")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	noOfElements, err := ReadSimpleField(ctx, "noOfElements", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfElements' field"))
	}
	m.NoOfElements = noOfElements

	elements, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "elements", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("585")), readBuffer), uint64(noOfElements))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'elements' field"))
	}
	m.Elements = elements

	if closeErr := readBuffer.CloseContext("ContentFilter"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ContentFilter")
	}

	return m, nil
}

func (m *_ContentFilter) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ContentFilter) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ContentFilter"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ContentFilter")
		}

		if err := WriteSimpleField[int32](ctx, "noOfElements", m.GetNoOfElements(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfElements' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "elements", m.GetElements(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'elements' field")
		}

		if popErr := writeBuffer.PopContext("ContentFilter"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ContentFilter")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ContentFilter) IsContentFilter() {}

func (m *_ContentFilter) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
