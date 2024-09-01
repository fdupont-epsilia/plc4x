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

// BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry is the corresponding interface of BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry
type BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetMonitoredObjectIdentifier returns MonitoredObjectIdentifier (property field)
	GetMonitoredObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetListOfCovReferences returns ListOfCovReferences (property field)
	GetListOfCovReferences() BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences
}

// BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryExactly can be used when we want exactly this type and not a type which fulfills BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry.
// This is useful for switch cases.
type BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryExactly interface {
	BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry
	isBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry() bool
}

// _BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry is the data-structure of this message
type _BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry struct {
	MonitoredObjectIdentifier BACnetContextTagObjectIdentifier
	ListOfCovReferences       BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences
}

var _ BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry = (*_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) GetMonitoredObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.MonitoredObjectIdentifier
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) GetListOfCovReferences() BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences {
	return m.ListOfCovReferences
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry factory function for _BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry
func NewBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry(monitoredObjectIdentifier BACnetContextTagObjectIdentifier, listOfCovReferences BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences) *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry {
	return &_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry{MonitoredObjectIdentifier: monitoredObjectIdentifier, ListOfCovReferences: listOfCovReferences}
}

// Deprecated: use the interface for direct cast
func CastBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry(structType any) BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry {
	if casted, ok := structType.(BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) GetTypeName() string {
	return "BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry"
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (monitoredObjectIdentifier)
	lengthInBits += m.MonitoredObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (listOfCovReferences)
	lengthInBits += m.ListOfCovReferences.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryParse(ctx context.Context, theBytes []byte) (BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry, error) {
	return BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry, error) {
		return BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry, error) {
	v, err := (&_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) parse(ctx context.Context, readBuffer utils.ReadBuffer) (_bACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	monitoredObjectIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "monitoredObjectIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'monitoredObjectIdentifier' field"))
	}
	m.MonitoredObjectIdentifier = monitoredObjectIdentifier

	listOfCovReferences, err := ReadSimpleField[BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences](ctx, "listOfCovReferences", ReadComplex[BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences](BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferencesParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfCovReferences' field"))
	}
	m.ListOfCovReferences = listOfCovReferences

	if closeErr := readBuffer.CloseContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry")
	}

	return m, nil
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry")
	}

	if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "monitoredObjectIdentifier", m.GetMonitoredObjectIdentifier(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'monitoredObjectIdentifier' field")
	}

	if err := WriteSimpleField[BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences](ctx, "listOfCovReferences", m.GetListOfCovReferences(), WriteComplex[BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryListOfCovReferences](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'listOfCovReferences' field")
	}

	if popErr := writeBuffer.PopContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry")
	}
	return nil
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) isBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry() bool {
	return true
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
