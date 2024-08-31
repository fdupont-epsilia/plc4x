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

// BACnetServiceAckAtomicReadFileStreamOrRecord is the corresponding interface of BACnetServiceAckAtomicReadFileStreamOrRecord
type BACnetServiceAckAtomicReadFileStreamOrRecord interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetServiceAckAtomicReadFileStreamOrRecordExactly can be used when we want exactly this type and not a type which fulfills BACnetServiceAckAtomicReadFileStreamOrRecord.
// This is useful for switch cases.
type BACnetServiceAckAtomicReadFileStreamOrRecordExactly interface {
	BACnetServiceAckAtomicReadFileStreamOrRecord
	isBACnetServiceAckAtomicReadFileStreamOrRecord() bool
}

// _BACnetServiceAckAtomicReadFileStreamOrRecord is the data-structure of this message
type _BACnetServiceAckAtomicReadFileStreamOrRecord struct {
	_BACnetServiceAckAtomicReadFileStreamOrRecordChildRequirements
	PeekedTagHeader BACnetTagHeader
	OpeningTag      BACnetOpeningTag
	ClosingTag      BACnetClosingTag
}

type _BACnetServiceAckAtomicReadFileStreamOrRecordChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetPeekedTagNumber() uint8
}

type BACnetServiceAckAtomicReadFileStreamOrRecordParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetServiceAckAtomicReadFileStreamOrRecord, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetServiceAckAtomicReadFileStreamOrRecordChild interface {
	utils.Serializable
	InitializeParent(parent BACnetServiceAckAtomicReadFileStreamOrRecord, peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, closingTag BACnetClosingTag)
	GetParent() *BACnetServiceAckAtomicReadFileStreamOrRecord

	GetTypeName() string
	BACnetServiceAckAtomicReadFileStreamOrRecord
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckAtomicReadFileStreamOrRecord) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *_BACnetServiceAckAtomicReadFileStreamOrRecord) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetServiceAckAtomicReadFileStreamOrRecord) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetServiceAckAtomicReadFileStreamOrRecord) GetPeekedTagNumber() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckAtomicReadFileStreamOrRecord factory function for _BACnetServiceAckAtomicReadFileStreamOrRecord
func NewBACnetServiceAckAtomicReadFileStreamOrRecord(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, closingTag BACnetClosingTag) *_BACnetServiceAckAtomicReadFileStreamOrRecord {
	return &_BACnetServiceAckAtomicReadFileStreamOrRecord{PeekedTagHeader: peekedTagHeader, OpeningTag: openingTag, ClosingTag: closingTag}
}

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckAtomicReadFileStreamOrRecord(structType any) BACnetServiceAckAtomicReadFileStreamOrRecord {
	if casted, ok := structType.(BACnetServiceAckAtomicReadFileStreamOrRecord); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckAtomicReadFileStreamOrRecord); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckAtomicReadFileStreamOrRecord) GetTypeName() string {
	return "BACnetServiceAckAtomicReadFileStreamOrRecord"
}

func (m *_BACnetServiceAckAtomicReadFileStreamOrRecord) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetServiceAckAtomicReadFileStreamOrRecord) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetServiceAckAtomicReadFileStreamOrRecordParse(ctx context.Context, theBytes []byte) (BACnetServiceAckAtomicReadFileStreamOrRecord, error) {
	return BACnetServiceAckAtomicReadFileStreamOrRecordParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetServiceAckAtomicReadFileStreamOrRecordParseWithBufferProducer[T BACnetServiceAckAtomicReadFileStreamOrRecord]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		buffer, err := BACnetServiceAckAtomicReadFileStreamOrRecordParseWithBuffer(ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return buffer.(T), err
	}
}

func BACnetServiceAckAtomicReadFileStreamOrRecordParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetServiceAckAtomicReadFileStreamOrRecord, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetServiceAckAtomicReadFileStreamOrRecord"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckAtomicReadFileStreamOrRecord")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (peekedTagHeader)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for peekedTagHeader")
	}
	peekedTagHeader, _ := BACnetTagHeaderParseWithBuffer(ctx, readBuffer)
	readBuffer.Reset(currentPos)

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(peekedTagHeader.GetActualTagNumber())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetServiceAckAtomicReadFileStreamOrRecordChildSerializeRequirement interface {
		BACnetServiceAckAtomicReadFileStreamOrRecord
		InitializeParent(BACnetServiceAckAtomicReadFileStreamOrRecord, BACnetTagHeader, BACnetOpeningTag, BACnetClosingTag)
		GetParent() BACnetServiceAckAtomicReadFileStreamOrRecord
	}
	var _childTemp any
	var _child BACnetServiceAckAtomicReadFileStreamOrRecordChildSerializeRequirement
	var typeSwitchError error
	switch {
	case peekedTagNumber == 0x0: // BACnetServiceAckAtomicReadFileStream
		_childTemp, typeSwitchError = BACnetServiceAckAtomicReadFileStreamParseWithBuffer(ctx, readBuffer)
	case peekedTagNumber == 0x1: // BACnetServiceAckAtomicReadFileRecord
		_childTemp, typeSwitchError = BACnetServiceAckAtomicReadFileRecordParseWithBuffer(ctx, readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetServiceAckAtomicReadFileStreamOrRecord")
	}
	_child = _childTemp.(BACnetServiceAckAtomicReadFileStreamOrRecordChildSerializeRequirement)

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(peekedTagHeader.GetActualTagNumber())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckAtomicReadFileStreamOrRecord"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckAtomicReadFileStreamOrRecord")
	}

	// Finish initializing
	_child.InitializeParent(_child, peekedTagHeader, openingTag, closingTag)
	return _child, nil
}

func (pm *_BACnetServiceAckAtomicReadFileStreamOrRecord) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetServiceAckAtomicReadFileStreamOrRecord, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetServiceAckAtomicReadFileStreamOrRecord"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckAtomicReadFileStreamOrRecord")
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
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
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

	if popErr := writeBuffer.PopContext("BACnetServiceAckAtomicReadFileStreamOrRecord"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetServiceAckAtomicReadFileStreamOrRecord")
	}
	return nil
}

func (m *_BACnetServiceAckAtomicReadFileStreamOrRecord) isBACnetServiceAckAtomicReadFileStreamOrRecord() bool {
	return true
}

func (m *_BACnetServiceAckAtomicReadFileStreamOrRecord) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
