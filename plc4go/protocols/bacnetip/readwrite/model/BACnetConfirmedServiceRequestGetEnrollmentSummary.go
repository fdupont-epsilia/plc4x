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

// BACnetConfirmedServiceRequestGetEnrollmentSummary is the corresponding interface of BACnetConfirmedServiceRequestGetEnrollmentSummary
type BACnetConfirmedServiceRequestGetEnrollmentSummary interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetAcknowledgmentFilter returns AcknowledgmentFilter (property field)
	GetAcknowledgmentFilter() BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged
	// GetEnrollmentFilter returns EnrollmentFilter (property field)
	GetEnrollmentFilter() BACnetRecipientProcessEnclosed
	// GetEventStateFilter returns EventStateFilter (property field)
	GetEventStateFilter() BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged
	// GetEventTypeFilter returns EventTypeFilter (property field)
	GetEventTypeFilter() BACnetEventTypeTagged
	// GetPriorityFilter returns PriorityFilter (property field)
	GetPriorityFilter() BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter
	// GetNotificationClassFilter returns NotificationClassFilter (property field)
	GetNotificationClassFilter() BACnetContextTagUnsignedInteger
}

// BACnetConfirmedServiceRequestGetEnrollmentSummaryExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestGetEnrollmentSummary.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestGetEnrollmentSummaryExactly interface {
	BACnetConfirmedServiceRequestGetEnrollmentSummary
	isBACnetConfirmedServiceRequestGetEnrollmentSummary() bool
}

// _BACnetConfirmedServiceRequestGetEnrollmentSummary is the data-structure of this message
type _BACnetConfirmedServiceRequestGetEnrollmentSummary struct {
	*_BACnetConfirmedServiceRequest
	AcknowledgmentFilter    BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged
	EnrollmentFilter        BACnetRecipientProcessEnclosed
	EventStateFilter        BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged
	EventTypeFilter         BACnetEventTypeTagged
	PriorityFilter          BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter
	NotificationClassFilter BACnetContextTagUnsignedInteger
}

var _ BACnetConfirmedServiceRequestGetEnrollmentSummary = (*_BACnetConfirmedServiceRequestGetEnrollmentSummary)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummary) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_GET_ENROLLMENT_SUMMARY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummary) InitializeParent(parent BACnetConfirmedServiceRequest) {
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummary) GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummary) GetAcknowledgmentFilter() BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged {
	return m.AcknowledgmentFilter
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummary) GetEnrollmentFilter() BACnetRecipientProcessEnclosed {
	return m.EnrollmentFilter
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummary) GetEventStateFilter() BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged {
	return m.EventStateFilter
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummary) GetEventTypeFilter() BACnetEventTypeTagged {
	return m.EventTypeFilter
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummary) GetPriorityFilter() BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter {
	return m.PriorityFilter
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummary) GetNotificationClassFilter() BACnetContextTagUnsignedInteger {
	return m.NotificationClassFilter
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestGetEnrollmentSummary factory function for _BACnetConfirmedServiceRequestGetEnrollmentSummary
func NewBACnetConfirmedServiceRequestGetEnrollmentSummary(acknowledgmentFilter BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged, enrollmentFilter BACnetRecipientProcessEnclosed, eventStateFilter BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged, eventTypeFilter BACnetEventTypeTagged, priorityFilter BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter, notificationClassFilter BACnetContextTagUnsignedInteger, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestGetEnrollmentSummary {
	_result := &_BACnetConfirmedServiceRequestGetEnrollmentSummary{
		AcknowledgmentFilter:           acknowledgmentFilter,
		EnrollmentFilter:               enrollmentFilter,
		EventStateFilter:               eventStateFilter,
		EventTypeFilter:                eventTypeFilter,
		PriorityFilter:                 priorityFilter,
		NotificationClassFilter:        notificationClassFilter,
		_BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestGetEnrollmentSummary(structType any) BACnetConfirmedServiceRequestGetEnrollmentSummary {
	if casted, ok := structType.(BACnetConfirmedServiceRequestGetEnrollmentSummary); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestGetEnrollmentSummary); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummary) GetTypeName() string {
	return "BACnetConfirmedServiceRequestGetEnrollmentSummary"
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummary) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (acknowledgmentFilter)
	lengthInBits += m.AcknowledgmentFilter.GetLengthInBits(ctx)

	// Optional Field (enrollmentFilter)
	if m.EnrollmentFilter != nil {
		lengthInBits += m.EnrollmentFilter.GetLengthInBits(ctx)
	}

	// Optional Field (eventStateFilter)
	if m.EventStateFilter != nil {
		lengthInBits += m.EventStateFilter.GetLengthInBits(ctx)
	}

	// Optional Field (eventTypeFilter)
	if m.EventTypeFilter != nil {
		lengthInBits += m.EventTypeFilter.GetLengthInBits(ctx)
	}

	// Optional Field (priorityFilter)
	if m.PriorityFilter != nil {
		lengthInBits += m.PriorityFilter.GetLengthInBits(ctx)
	}

	// Optional Field (notificationClassFilter)
	if m.NotificationClassFilter != nil {
		lengthInBits += m.NotificationClassFilter.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummary) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestGetEnrollmentSummaryParse(ctx context.Context, theBytes []byte, serviceRequestLength uint32) (BACnetConfirmedServiceRequestGetEnrollmentSummary, error) {
	return BACnetConfirmedServiceRequestGetEnrollmentSummaryParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), serviceRequestLength)
}

func BACnetConfirmedServiceRequestGetEnrollmentSummaryParseWithBufferProducer(serviceRequestLength uint32) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestGetEnrollmentSummary, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestGetEnrollmentSummary, error) {
		return BACnetConfirmedServiceRequestGetEnrollmentSummaryParseWithBuffer(ctx, readBuffer, serviceRequestLength)
	}
}

func BACnetConfirmedServiceRequestGetEnrollmentSummaryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceRequestLength uint32) (BACnetConfirmedServiceRequestGetEnrollmentSummary, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestGetEnrollmentSummary"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestGetEnrollmentSummary")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	acknowledgmentFilter, err := ReadSimpleField[BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged](ctx, "acknowledgmentFilter", ReadComplex[BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged](BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'acknowledgmentFilter' field"))
	}

	_enrollmentFilter, err := ReadOptionalField[BACnetRecipientProcessEnclosed](ctx, "enrollmentFilter", ReadComplex[BACnetRecipientProcessEnclosed](BACnetRecipientProcessEnclosedParseWithBufferProducer((uint8)(uint8(1))), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'enrollmentFilter' field"))
	}
	var enrollmentFilter BACnetRecipientProcessEnclosed
	if _enrollmentFilter != nil {
		enrollmentFilter = *_enrollmentFilter
	}

	_eventStateFilter, err := ReadOptionalField[BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged](ctx, "eventStateFilter", ReadComplex[BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged](BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTaggedParseWithBufferProducer((uint8)(uint8(2)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventStateFilter' field"))
	}
	var eventStateFilter BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged
	if _eventStateFilter != nil {
		eventStateFilter = *_eventStateFilter
	}

	_eventTypeFilter, err := ReadOptionalField[BACnetEventTypeTagged](ctx, "eventTypeFilter", ReadComplex[BACnetEventTypeTagged](BACnetEventTypeTaggedParseWithBufferProducer((uint8)(uint8(3)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventTypeFilter' field"))
	}
	var eventTypeFilter BACnetEventTypeTagged
	if _eventTypeFilter != nil {
		eventTypeFilter = *_eventTypeFilter
	}

	_priorityFilter, err := ReadOptionalField[BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter](ctx, "priorityFilter", ReadComplex[BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter](BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilterParseWithBufferProducer((uint8)(uint8(4))), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priorityFilter' field"))
	}
	var priorityFilter BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter
	if _priorityFilter != nil {
		priorityFilter = *_priorityFilter
	}

	_notificationClassFilter, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "notificationClassFilter", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(5)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'notificationClassFilter' field"))
	}
	var notificationClassFilter BACnetContextTagUnsignedInteger
	if _notificationClassFilter != nil {
		notificationClassFilter = *_notificationClassFilter
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestGetEnrollmentSummary"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestGetEnrollmentSummary")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestGetEnrollmentSummary{
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		AcknowledgmentFilter:    acknowledgmentFilter,
		EnrollmentFilter:        enrollmentFilter,
		EventStateFilter:        eventStateFilter,
		EventTypeFilter:         eventTypeFilter,
		PriorityFilter:          priorityFilter,
		NotificationClassFilter: notificationClassFilter,
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummary) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummary) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestGetEnrollmentSummary"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestGetEnrollmentSummary")
		}

		if err := WriteSimpleField[BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged](ctx, "acknowledgmentFilter", m.GetAcknowledgmentFilter(), WriteComplex[BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'acknowledgmentFilter' field")
		}

		if err := WriteOptionalField[BACnetRecipientProcessEnclosed](ctx, "enrollmentFilter", GetRef(m.GetEnrollmentFilter()), WriteComplex[BACnetRecipientProcessEnclosed](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'enrollmentFilter' field")
		}

		if err := WriteOptionalField[BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged](ctx, "eventStateFilter", GetRef(m.GetEventStateFilter()), WriteComplex[BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'eventStateFilter' field")
		}

		if err := WriteOptionalField[BACnetEventTypeTagged](ctx, "eventTypeFilter", GetRef(m.GetEventTypeFilter()), WriteComplex[BACnetEventTypeTagged](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'eventTypeFilter' field")
		}

		if err := WriteOptionalField[BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter](ctx, "priorityFilter", GetRef(m.GetPriorityFilter()), WriteComplex[BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'priorityFilter' field")
		}

		if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "notificationClassFilter", GetRef(m.GetNotificationClassFilter()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'notificationClassFilter' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestGetEnrollmentSummary"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestGetEnrollmentSummary")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummary) isBACnetConfirmedServiceRequestGetEnrollmentSummary() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummary) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
