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

// MediaTransportControlDataEnumerationsSize is the corresponding interface of MediaTransportControlDataEnumerationsSize
type MediaTransportControlDataEnumerationsSize interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MediaTransportControlData
	// GetSizeType returns SizeType (property field)
	GetSizeType() byte
	// GetStart returns Start (property field)
	GetStart() uint8
	// GetSize returns Size (property field)
	GetSize() uint8
	// GetIsListCategories returns IsListCategories (virtual field)
	GetIsListCategories() bool
	// GetIsListSelections returns IsListSelections (virtual field)
	GetIsListSelections() bool
	// GetIsListTracks returns IsListTracks (virtual field)
	GetIsListTracks() bool
	// GetIsReserved returns IsReserved (virtual field)
	GetIsReserved() bool
	// IsMediaTransportControlDataEnumerationsSize is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMediaTransportControlDataEnumerationsSize()
}

// _MediaTransportControlDataEnumerationsSize is the data-structure of this message
type _MediaTransportControlDataEnumerationsSize struct {
	MediaTransportControlDataContract
	SizeType byte
	Start    uint8
	Size     uint8
}

var _ MediaTransportControlDataEnumerationsSize = (*_MediaTransportControlDataEnumerationsSize)(nil)
var _ MediaTransportControlDataRequirements = (*_MediaTransportControlDataEnumerationsSize)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MediaTransportControlDataEnumerationsSize) GetParent() MediaTransportControlDataContract {
	return m.MediaTransportControlDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MediaTransportControlDataEnumerationsSize) GetSizeType() byte {
	return m.SizeType
}

func (m *_MediaTransportControlDataEnumerationsSize) GetStart() uint8 {
	return m.Start
}

func (m *_MediaTransportControlDataEnumerationsSize) GetSize() uint8 {
	return m.Size
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_MediaTransportControlDataEnumerationsSize) GetIsListCategories() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetSizeType()) == (0x00)))
}

func (m *_MediaTransportControlDataEnumerationsSize) GetIsListSelections() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetSizeType()) == (0x01)))
}

func (m *_MediaTransportControlDataEnumerationsSize) GetIsListTracks() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetSizeType()) == (0x02)))
}

func (m *_MediaTransportControlDataEnumerationsSize) GetIsReserved() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool(bool(!(m.GetIsListCategories())) && bool(!(m.GetIsListSelections()))) && bool(!(m.GetIsListTracks())))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMediaTransportControlDataEnumerationsSize factory function for _MediaTransportControlDataEnumerationsSize
func NewMediaTransportControlDataEnumerationsSize(sizeType byte, start uint8, size uint8, commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) *_MediaTransportControlDataEnumerationsSize {
	_result := &_MediaTransportControlDataEnumerationsSize{
		MediaTransportControlDataContract: NewMediaTransportControlData(commandTypeContainer, mediaLinkGroup),
		SizeType:                          sizeType,
		Start:                             start,
		Size:                              size,
	}
	_result.MediaTransportControlDataContract.(*_MediaTransportControlData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMediaTransportControlDataEnumerationsSize(structType any) MediaTransportControlDataEnumerationsSize {
	if casted, ok := structType.(MediaTransportControlDataEnumerationsSize); ok {
		return casted
	}
	if casted, ok := structType.(*MediaTransportControlDataEnumerationsSize); ok {
		return *casted
	}
	return nil
}

func (m *_MediaTransportControlDataEnumerationsSize) GetTypeName() string {
	return "MediaTransportControlDataEnumerationsSize"
}

func (m *_MediaTransportControlDataEnumerationsSize) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MediaTransportControlDataContract.(*_MediaTransportControlData).getLengthInBits(ctx))

	// Simple field (sizeType)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (start)
	lengthInBits += 8

	// Simple field (size)
	lengthInBits += 8

	return lengthInBits
}

func (m *_MediaTransportControlDataEnumerationsSize) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MediaTransportControlDataEnumerationsSize) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MediaTransportControlData) (__mediaTransportControlDataEnumerationsSize MediaTransportControlDataEnumerationsSize, err error) {
	m.MediaTransportControlDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MediaTransportControlDataEnumerationsSize"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlDataEnumerationsSize")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	sizeType, err := ReadSimpleField(ctx, "sizeType", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sizeType' field"))
	}
	m.SizeType = sizeType

	isListCategories, err := ReadVirtualField[bool](ctx, "isListCategories", (*bool)(nil), bool((sizeType) == (0x00)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isListCategories' field"))
	}
	_ = isListCategories

	isListSelections, err := ReadVirtualField[bool](ctx, "isListSelections", (*bool)(nil), bool((sizeType) == (0x01)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isListSelections' field"))
	}
	_ = isListSelections

	isListTracks, err := ReadVirtualField[bool](ctx, "isListTracks", (*bool)(nil), bool((sizeType) == (0x02)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isListTracks' field"))
	}
	_ = isListTracks

	isReserved, err := ReadVirtualField[bool](ctx, "isReserved", (*bool)(nil), bool(bool(!(isListCategories)) && bool(!(isListSelections))) && bool(!(isListTracks)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isReserved' field"))
	}
	_ = isReserved

	start, err := ReadSimpleField(ctx, "start", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'start' field"))
	}
	m.Start = start

	size, err := ReadSimpleField(ctx, "size", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'size' field"))
	}
	m.Size = size

	if closeErr := readBuffer.CloseContext("MediaTransportControlDataEnumerationsSize"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlDataEnumerationsSize")
	}

	return m, nil
}

func (m *_MediaTransportControlDataEnumerationsSize) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MediaTransportControlDataEnumerationsSize) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MediaTransportControlDataEnumerationsSize"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MediaTransportControlDataEnumerationsSize")
		}

		if err := WriteSimpleField[byte](ctx, "sizeType", m.GetSizeType(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'sizeType' field")
		}
		// Virtual field
		isListCategories := m.GetIsListCategories()
		_ = isListCategories
		if _isListCategoriesErr := writeBuffer.WriteVirtual(ctx, "isListCategories", m.GetIsListCategories()); _isListCategoriesErr != nil {
			return errors.Wrap(_isListCategoriesErr, "Error serializing 'isListCategories' field")
		}
		// Virtual field
		isListSelections := m.GetIsListSelections()
		_ = isListSelections
		if _isListSelectionsErr := writeBuffer.WriteVirtual(ctx, "isListSelections", m.GetIsListSelections()); _isListSelectionsErr != nil {
			return errors.Wrap(_isListSelectionsErr, "Error serializing 'isListSelections' field")
		}
		// Virtual field
		isListTracks := m.GetIsListTracks()
		_ = isListTracks
		if _isListTracksErr := writeBuffer.WriteVirtual(ctx, "isListTracks", m.GetIsListTracks()); _isListTracksErr != nil {
			return errors.Wrap(_isListTracksErr, "Error serializing 'isListTracks' field")
		}
		// Virtual field
		isReserved := m.GetIsReserved()
		_ = isReserved
		if _isReservedErr := writeBuffer.WriteVirtual(ctx, "isReserved", m.GetIsReserved()); _isReservedErr != nil {
			return errors.Wrap(_isReservedErr, "Error serializing 'isReserved' field")
		}

		if err := WriteSimpleField[uint8](ctx, "start", m.GetStart(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'start' field")
		}

		if err := WriteSimpleField[uint8](ctx, "size", m.GetSize(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'size' field")
		}

		if popErr := writeBuffer.PopContext("MediaTransportControlDataEnumerationsSize"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MediaTransportControlDataEnumerationsSize")
		}
		return nil
	}
	return m.MediaTransportControlDataContract.(*_MediaTransportControlData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MediaTransportControlDataEnumerationsSize) IsMediaTransportControlDataEnumerationsSize() {}

func (m *_MediaTransportControlDataEnumerationsSize) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
