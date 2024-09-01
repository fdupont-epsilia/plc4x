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

// DIBSuppSvcFamilies is the corresponding interface of DIBSuppSvcFamilies
type DIBSuppSvcFamilies interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetDescriptionType returns DescriptionType (property field)
	GetDescriptionType() uint8
	// GetServiceIds returns ServiceIds (property field)
	GetServiceIds() []ServiceId
}

// DIBSuppSvcFamiliesExactly can be used when we want exactly this type and not a type which fulfills DIBSuppSvcFamilies.
// This is useful for switch cases.
type DIBSuppSvcFamiliesExactly interface {
	DIBSuppSvcFamilies
	isDIBSuppSvcFamilies() bool
}

// _DIBSuppSvcFamilies is the data-structure of this message
type _DIBSuppSvcFamilies struct {
	DescriptionType uint8
	ServiceIds      []ServiceId
}

var _ DIBSuppSvcFamilies = (*_DIBSuppSvcFamilies)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DIBSuppSvcFamilies) GetDescriptionType() uint8 {
	return m.DescriptionType
}

func (m *_DIBSuppSvcFamilies) GetServiceIds() []ServiceId {
	return m.ServiceIds
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDIBSuppSvcFamilies factory function for _DIBSuppSvcFamilies
func NewDIBSuppSvcFamilies(descriptionType uint8, serviceIds []ServiceId) *_DIBSuppSvcFamilies {
	return &_DIBSuppSvcFamilies{DescriptionType: descriptionType, ServiceIds: serviceIds}
}

// Deprecated: use the interface for direct cast
func CastDIBSuppSvcFamilies(structType any) DIBSuppSvcFamilies {
	if casted, ok := structType.(DIBSuppSvcFamilies); ok {
		return casted
	}
	if casted, ok := structType.(*DIBSuppSvcFamilies); ok {
		return *casted
	}
	return nil
}

func (m *_DIBSuppSvcFamilies) GetTypeName() string {
	return "DIBSuppSvcFamilies"
}

func (m *_DIBSuppSvcFamilies) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (structureLength)
	lengthInBits += 8

	// Simple field (descriptionType)
	lengthInBits += 8

	// Array field
	if len(m.ServiceIds) > 0 {
		for _, element := range m.ServiceIds {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_DIBSuppSvcFamilies) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DIBSuppSvcFamiliesParse(ctx context.Context, theBytes []byte) (DIBSuppSvcFamilies, error) {
	return DIBSuppSvcFamiliesParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func DIBSuppSvcFamiliesParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (DIBSuppSvcFamilies, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (DIBSuppSvcFamilies, error) {
		return DIBSuppSvcFamiliesParseWithBuffer(ctx, readBuffer)
	}
}

func DIBSuppSvcFamiliesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DIBSuppSvcFamilies, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DIBSuppSvcFamilies"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DIBSuppSvcFamilies")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	structureLength, err := ReadImplicitField[uint8](ctx, "structureLength", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'structureLength' field"))
	}
	_ = structureLength

	descriptionType, err := ReadSimpleField(ctx, "descriptionType", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'descriptionType' field"))
	}

	serviceIds, err := ReadLengthArrayField[ServiceId](ctx, "serviceIds", ReadComplex[ServiceId](ServiceIdParseWithBuffer, readBuffer), int(int32(structureLength)-int32(int32(2))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serviceIds' field"))
	}

	if closeErr := readBuffer.CloseContext("DIBSuppSvcFamilies"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DIBSuppSvcFamilies")
	}

	// Create the instance
	return &_DIBSuppSvcFamilies{
		DescriptionType: descriptionType,
		ServiceIds:      serviceIds,
	}, nil
}

func (m *_DIBSuppSvcFamilies) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DIBSuppSvcFamilies) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("DIBSuppSvcFamilies"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for DIBSuppSvcFamilies")
	}
	structureLength := uint8(uint8(m.GetLengthInBytes(ctx)))
	if err := WriteImplicitField(ctx, "structureLength", structureLength, WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'structureLength' field")
	}

	if err := WriteSimpleField[uint8](ctx, "descriptionType", m.GetDescriptionType(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'descriptionType' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "serviceIds", m.GetServiceIds(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'serviceIds' field")
	}

	if popErr := writeBuffer.PopContext("DIBSuppSvcFamilies"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for DIBSuppSvcFamilies")
	}
	return nil
}

func (m *_DIBSuppSvcFamilies) isDIBSuppSvcFamilies() bool {
	return true
}

func (m *_DIBSuppSvcFamilies) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
