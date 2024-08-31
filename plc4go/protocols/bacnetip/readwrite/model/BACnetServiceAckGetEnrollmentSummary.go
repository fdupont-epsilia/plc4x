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

// BACnetServiceAckGetEnrollmentSummary is the corresponding interface of BACnetServiceAckGetEnrollmentSummary
type BACnetServiceAckGetEnrollmentSummary interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetServiceAck
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetApplicationTagObjectIdentifier
	// GetEventType returns EventType (property field)
	GetEventType() BACnetEventTypeTagged
	// GetEventState returns EventState (property field)
	GetEventState() BACnetEventStateTagged
	// GetPriority returns Priority (property field)
	GetPriority() BACnetApplicationTagUnsignedInteger
	// GetNotificationClass returns NotificationClass (property field)
	GetNotificationClass() BACnetApplicationTagUnsignedInteger
}

// BACnetServiceAckGetEnrollmentSummaryExactly can be used when we want exactly this type and not a type which fulfills BACnetServiceAckGetEnrollmentSummary.
// This is useful for switch cases.
type BACnetServiceAckGetEnrollmentSummaryExactly interface {
	BACnetServiceAckGetEnrollmentSummary
	isBACnetServiceAckGetEnrollmentSummary() bool
}

// _BACnetServiceAckGetEnrollmentSummary is the data-structure of this message
type _BACnetServiceAckGetEnrollmentSummary struct {
	*_BACnetServiceAck
	ObjectIdentifier  BACnetApplicationTagObjectIdentifier
	EventType         BACnetEventTypeTagged
	EventState        BACnetEventStateTagged
	Priority          BACnetApplicationTagUnsignedInteger
	NotificationClass BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetServiceAckGetEnrollmentSummary) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_GET_ENROLLMENT_SUMMARY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckGetEnrollmentSummary) InitializeParent(parent BACnetServiceAck) {}

func (m *_BACnetServiceAckGetEnrollmentSummary) GetParent() BACnetServiceAck {
	return m._BACnetServiceAck
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckGetEnrollmentSummary) GetObjectIdentifier() BACnetApplicationTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetServiceAckGetEnrollmentSummary) GetEventType() BACnetEventTypeTagged {
	return m.EventType
}

func (m *_BACnetServiceAckGetEnrollmentSummary) GetEventState() BACnetEventStateTagged {
	return m.EventState
}

func (m *_BACnetServiceAckGetEnrollmentSummary) GetPriority() BACnetApplicationTagUnsignedInteger {
	return m.Priority
}

func (m *_BACnetServiceAckGetEnrollmentSummary) GetNotificationClass() BACnetApplicationTagUnsignedInteger {
	return m.NotificationClass
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckGetEnrollmentSummary factory function for _BACnetServiceAckGetEnrollmentSummary
func NewBACnetServiceAckGetEnrollmentSummary(objectIdentifier BACnetApplicationTagObjectIdentifier, eventType BACnetEventTypeTagged, eventState BACnetEventStateTagged, priority BACnetApplicationTagUnsignedInteger, notificationClass BACnetApplicationTagUnsignedInteger, serviceAckLength uint32) *_BACnetServiceAckGetEnrollmentSummary {
	_result := &_BACnetServiceAckGetEnrollmentSummary{
		ObjectIdentifier:  objectIdentifier,
		EventType:         eventType,
		EventState:        eventState,
		Priority:          priority,
		NotificationClass: notificationClass,
		_BACnetServiceAck: NewBACnetServiceAck(serviceAckLength),
	}
	_result._BACnetServiceAck._BACnetServiceAckChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckGetEnrollmentSummary(structType any) BACnetServiceAckGetEnrollmentSummary {
	if casted, ok := structType.(BACnetServiceAckGetEnrollmentSummary); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckGetEnrollmentSummary); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckGetEnrollmentSummary) GetTypeName() string {
	return "BACnetServiceAckGetEnrollmentSummary"
}

func (m *_BACnetServiceAckGetEnrollmentSummary) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (eventType)
	lengthInBits += m.EventType.GetLengthInBits(ctx)

	// Simple field (eventState)
	lengthInBits += m.EventState.GetLengthInBits(ctx)

	// Simple field (priority)
	lengthInBits += m.Priority.GetLengthInBits(ctx)

	// Optional Field (notificationClass)
	if m.NotificationClass != nil {
		lengthInBits += m.NotificationClass.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetServiceAckGetEnrollmentSummary) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetServiceAckGetEnrollmentSummaryParse(ctx context.Context, theBytes []byte, serviceAckLength uint32) (BACnetServiceAckGetEnrollmentSummary, error) {
	return BACnetServiceAckGetEnrollmentSummaryParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), serviceAckLength)
}

func BACnetServiceAckGetEnrollmentSummaryParseWithBufferProducer(serviceAckLength uint32) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetServiceAckGetEnrollmentSummary, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetServiceAckGetEnrollmentSummary, error) {
		return BACnetServiceAckGetEnrollmentSummaryParseWithBuffer(ctx, readBuffer, serviceAckLength)
	}
}

func BACnetServiceAckGetEnrollmentSummaryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceAckLength uint32) (BACnetServiceAckGetEnrollmentSummary, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetServiceAckGetEnrollmentSummary"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckGetEnrollmentSummary")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectIdentifier, err := ReadSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "objectIdentifier", ReadComplex[BACnetApplicationTagObjectIdentifier](BACnetApplicationTagParseWithBufferProducer(), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectIdentifier' field"))
	}

	eventType, err := ReadSimpleField[BACnetEventTypeTagged](ctx, "eventType", ReadComplex[BACnetEventTypeTagged](BACnetEventTypeTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventType' field"))
	}

	eventState, err := ReadSimpleField[BACnetEventStateTagged](ctx, "eventState", ReadComplex[BACnetEventStateTagged](BACnetEventStateTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventState' field"))
	}

	priority, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "priority", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer(), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priority' field"))
	}

	_notificationClass, err := ReadOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "notificationClass", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer(), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'notificationClass' field"))
	}
	var notificationClass BACnetApplicationTagUnsignedInteger
	if _notificationClass != nil {
		notificationClass = *_notificationClass
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckGetEnrollmentSummary"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckGetEnrollmentSummary")
	}

	// Create a partially initialized instance
	_child := &_BACnetServiceAckGetEnrollmentSummary{
		_BACnetServiceAck: &_BACnetServiceAck{
			ServiceAckLength: serviceAckLength,
		},
		ObjectIdentifier:  objectIdentifier,
		EventType:         eventType,
		EventState:        eventState,
		Priority:          priority,
		NotificationClass: notificationClass,
	}
	_child._BACnetServiceAck._BACnetServiceAckChildRequirements = _child
	return _child, nil
}

func (m *_BACnetServiceAckGetEnrollmentSummary) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetServiceAckGetEnrollmentSummary) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckGetEnrollmentSummary"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckGetEnrollmentSummary")
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

		// Simple Field (eventType)
		if pushErr := writeBuffer.PushContext("eventType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for eventType")
		}
		_eventTypeErr := writeBuffer.WriteSerializable(ctx, m.GetEventType())
		if popErr := writeBuffer.PopContext("eventType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for eventType")
		}
		if _eventTypeErr != nil {
			return errors.Wrap(_eventTypeErr, "Error serializing 'eventType' field")
		}

		// Simple Field (eventState)
		if pushErr := writeBuffer.PushContext("eventState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for eventState")
		}
		_eventStateErr := writeBuffer.WriteSerializable(ctx, m.GetEventState())
		if popErr := writeBuffer.PopContext("eventState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for eventState")
		}
		if _eventStateErr != nil {
			return errors.Wrap(_eventStateErr, "Error serializing 'eventState' field")
		}

		// Simple Field (priority)
		if pushErr := writeBuffer.PushContext("priority"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for priority")
		}
		_priorityErr := writeBuffer.WriteSerializable(ctx, m.GetPriority())
		if popErr := writeBuffer.PopContext("priority"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for priority")
		}
		if _priorityErr != nil {
			return errors.Wrap(_priorityErr, "Error serializing 'priority' field")
		}

		// Optional Field (notificationClass) (Can be skipped, if the value is null)
		var notificationClass BACnetApplicationTagUnsignedInteger = nil
		if m.GetNotificationClass() != nil {
			if pushErr := writeBuffer.PushContext("notificationClass"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for notificationClass")
			}
			notificationClass = m.GetNotificationClass()
			_notificationClassErr := writeBuffer.WriteSerializable(ctx, notificationClass)
			if popErr := writeBuffer.PopContext("notificationClass"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for notificationClass")
			}
			if _notificationClassErr != nil {
				return errors.Wrap(_notificationClassErr, "Error serializing 'notificationClass' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckGetEnrollmentSummary"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckGetEnrollmentSummary")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetServiceAckGetEnrollmentSummary) isBACnetServiceAckGetEnrollmentSummary() bool {
	return true
}

func (m *_BACnetServiceAckGetEnrollmentSummary) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
