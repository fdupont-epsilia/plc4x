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

// BACnetConfirmedServiceRequestConfirmedEventNotification is the corresponding interface of BACnetConfirmedServiceRequestConfirmedEventNotification
type BACnetConfirmedServiceRequestConfirmedEventNotification interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetProcessIdentifier returns ProcessIdentifier (property field)
	GetProcessIdentifier() BACnetContextTagUnsignedInteger
	// GetInitiatingDeviceIdentifier returns InitiatingDeviceIdentifier (property field)
	GetInitiatingDeviceIdentifier() BACnetContextTagObjectIdentifier
	// GetEventObjectIdentifier returns EventObjectIdentifier (property field)
	GetEventObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() BACnetTimeStampEnclosed
	// GetNotificationClass returns NotificationClass (property field)
	GetNotificationClass() BACnetContextTagUnsignedInteger
	// GetPriority returns Priority (property field)
	GetPriority() BACnetContextTagUnsignedInteger
	// GetEventType returns EventType (property field)
	GetEventType() BACnetEventTypeTagged
	// GetMessageText returns MessageText (property field)
	GetMessageText() BACnetContextTagCharacterString
	// GetNotifyType returns NotifyType (property field)
	GetNotifyType() BACnetNotifyTypeTagged
	// GetAckRequired returns AckRequired (property field)
	GetAckRequired() BACnetContextTagBoolean
	// GetFromState returns FromState (property field)
	GetFromState() BACnetEventStateTagged
	// GetToState returns ToState (property field)
	GetToState() BACnetEventStateTagged
	// GetEventValues returns EventValues (property field)
	GetEventValues() BACnetNotificationParameters
}

// BACnetConfirmedServiceRequestConfirmedEventNotificationExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestConfirmedEventNotification.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestConfirmedEventNotificationExactly interface {
	BACnetConfirmedServiceRequestConfirmedEventNotification
	isBACnetConfirmedServiceRequestConfirmedEventNotification() bool
}

// _BACnetConfirmedServiceRequestConfirmedEventNotification is the data-structure of this message
type _BACnetConfirmedServiceRequestConfirmedEventNotification struct {
	*_BACnetConfirmedServiceRequest
	ProcessIdentifier          BACnetContextTagUnsignedInteger
	InitiatingDeviceIdentifier BACnetContextTagObjectIdentifier
	EventObjectIdentifier      BACnetContextTagObjectIdentifier
	Timestamp                  BACnetTimeStampEnclosed
	NotificationClass          BACnetContextTagUnsignedInteger
	Priority                   BACnetContextTagUnsignedInteger
	EventType                  BACnetEventTypeTagged
	MessageText                BACnetContextTagCharacterString
	NotifyType                 BACnetNotifyTypeTagged
	AckRequired                BACnetContextTagBoolean
	FromState                  BACnetEventStateTagged
	ToState                    BACnetEventStateTagged
	EventValues                BACnetNotificationParameters
}

var _ BACnetConfirmedServiceRequestConfirmedEventNotification = (*_BACnetConfirmedServiceRequestConfirmedEventNotification)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedEventNotification) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_CONFIRMED_EVENT_NOTIFICATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedEventNotification) InitializeParent(parent BACnetConfirmedServiceRequest) {
}

func (m *_BACnetConfirmedServiceRequestConfirmedEventNotification) GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedEventNotification) GetProcessIdentifier() BACnetContextTagUnsignedInteger {
	return m.ProcessIdentifier
}

func (m *_BACnetConfirmedServiceRequestConfirmedEventNotification) GetInitiatingDeviceIdentifier() BACnetContextTagObjectIdentifier {
	return m.InitiatingDeviceIdentifier
}

func (m *_BACnetConfirmedServiceRequestConfirmedEventNotification) GetEventObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.EventObjectIdentifier
}

func (m *_BACnetConfirmedServiceRequestConfirmedEventNotification) GetTimestamp() BACnetTimeStampEnclosed {
	return m.Timestamp
}

func (m *_BACnetConfirmedServiceRequestConfirmedEventNotification) GetNotificationClass() BACnetContextTagUnsignedInteger {
	return m.NotificationClass
}

func (m *_BACnetConfirmedServiceRequestConfirmedEventNotification) GetPriority() BACnetContextTagUnsignedInteger {
	return m.Priority
}

func (m *_BACnetConfirmedServiceRequestConfirmedEventNotification) GetEventType() BACnetEventTypeTagged {
	return m.EventType
}

func (m *_BACnetConfirmedServiceRequestConfirmedEventNotification) GetMessageText() BACnetContextTagCharacterString {
	return m.MessageText
}

func (m *_BACnetConfirmedServiceRequestConfirmedEventNotification) GetNotifyType() BACnetNotifyTypeTagged {
	return m.NotifyType
}

func (m *_BACnetConfirmedServiceRequestConfirmedEventNotification) GetAckRequired() BACnetContextTagBoolean {
	return m.AckRequired
}

func (m *_BACnetConfirmedServiceRequestConfirmedEventNotification) GetFromState() BACnetEventStateTagged {
	return m.FromState
}

func (m *_BACnetConfirmedServiceRequestConfirmedEventNotification) GetToState() BACnetEventStateTagged {
	return m.ToState
}

func (m *_BACnetConfirmedServiceRequestConfirmedEventNotification) GetEventValues() BACnetNotificationParameters {
	return m.EventValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestConfirmedEventNotification factory function for _BACnetConfirmedServiceRequestConfirmedEventNotification
func NewBACnetConfirmedServiceRequestConfirmedEventNotification(processIdentifier BACnetContextTagUnsignedInteger, initiatingDeviceIdentifier BACnetContextTagObjectIdentifier, eventObjectIdentifier BACnetContextTagObjectIdentifier, timestamp BACnetTimeStampEnclosed, notificationClass BACnetContextTagUnsignedInteger, priority BACnetContextTagUnsignedInteger, eventType BACnetEventTypeTagged, messageText BACnetContextTagCharacterString, notifyType BACnetNotifyTypeTagged, ackRequired BACnetContextTagBoolean, fromState BACnetEventStateTagged, toState BACnetEventStateTagged, eventValues BACnetNotificationParameters, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestConfirmedEventNotification {
	_result := &_BACnetConfirmedServiceRequestConfirmedEventNotification{
		ProcessIdentifier:              processIdentifier,
		InitiatingDeviceIdentifier:     initiatingDeviceIdentifier,
		EventObjectIdentifier:          eventObjectIdentifier,
		Timestamp:                      timestamp,
		NotificationClass:              notificationClass,
		Priority:                       priority,
		EventType:                      eventType,
		MessageText:                    messageText,
		NotifyType:                     notifyType,
		AckRequired:                    ackRequired,
		FromState:                      fromState,
		ToState:                        toState,
		EventValues:                    eventValues,
		_BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestConfirmedEventNotification(structType any) BACnetConfirmedServiceRequestConfirmedEventNotification {
	if casted, ok := structType.(BACnetConfirmedServiceRequestConfirmedEventNotification); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestConfirmedEventNotification); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedEventNotification) GetTypeName() string {
	return "BACnetConfirmedServiceRequestConfirmedEventNotification"
}

func (m *_BACnetConfirmedServiceRequestConfirmedEventNotification) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (processIdentifier)
	lengthInBits += m.ProcessIdentifier.GetLengthInBits(ctx)

	// Simple field (initiatingDeviceIdentifier)
	lengthInBits += m.InitiatingDeviceIdentifier.GetLengthInBits(ctx)

	// Simple field (eventObjectIdentifier)
	lengthInBits += m.EventObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (timestamp)
	lengthInBits += m.Timestamp.GetLengthInBits(ctx)

	// Simple field (notificationClass)
	lengthInBits += m.NotificationClass.GetLengthInBits(ctx)

	// Simple field (priority)
	lengthInBits += m.Priority.GetLengthInBits(ctx)

	// Simple field (eventType)
	lengthInBits += m.EventType.GetLengthInBits(ctx)

	// Optional Field (messageText)
	if m.MessageText != nil {
		lengthInBits += m.MessageText.GetLengthInBits(ctx)
	}

	// Simple field (notifyType)
	lengthInBits += m.NotifyType.GetLengthInBits(ctx)

	// Optional Field (ackRequired)
	if m.AckRequired != nil {
		lengthInBits += m.AckRequired.GetLengthInBits(ctx)
	}

	// Optional Field (fromState)
	if m.FromState != nil {
		lengthInBits += m.FromState.GetLengthInBits(ctx)
	}

	// Simple field (toState)
	lengthInBits += m.ToState.GetLengthInBits(ctx)

	// Optional Field (eventValues)
	if m.EventValues != nil {
		lengthInBits += m.EventValues.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestConfirmedEventNotification) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestConfirmedEventNotificationParse(ctx context.Context, theBytes []byte, serviceRequestLength uint32) (BACnetConfirmedServiceRequestConfirmedEventNotification, error) {
	return BACnetConfirmedServiceRequestConfirmedEventNotificationParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), serviceRequestLength)
}

func BACnetConfirmedServiceRequestConfirmedEventNotificationParseWithBufferProducer(serviceRequestLength uint32) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestConfirmedEventNotification, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestConfirmedEventNotification, error) {
		return BACnetConfirmedServiceRequestConfirmedEventNotificationParseWithBuffer(ctx, readBuffer, serviceRequestLength)
	}
}

func BACnetConfirmedServiceRequestConfirmedEventNotificationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceRequestLength uint32) (BACnetConfirmedServiceRequestConfirmedEventNotification, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestConfirmedEventNotification"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestConfirmedEventNotification")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	processIdentifier, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "processIdentifier", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'processIdentifier' field"))
	}

	initiatingDeviceIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "initiatingDeviceIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'initiatingDeviceIdentifier' field"))
	}

	eventObjectIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "eventObjectIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventObjectIdentifier' field"))
	}

	timestamp, err := ReadSimpleField[BACnetTimeStampEnclosed](ctx, "timestamp", ReadComplex[BACnetTimeStampEnclosed](BACnetTimeStampEnclosedParseWithBufferProducer((uint8)(uint8(3))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timestamp' field"))
	}

	notificationClass, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "notificationClass", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(4)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'notificationClass' field"))
	}

	priority, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "priority", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(5)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priority' field"))
	}

	eventType, err := ReadSimpleField[BACnetEventTypeTagged](ctx, "eventType", ReadComplex[BACnetEventTypeTagged](BACnetEventTypeTaggedParseWithBufferProducer((uint8)(uint8(6)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventType' field"))
	}

	_messageText, err := ReadOptionalField[BACnetContextTagCharacterString](ctx, "messageText", ReadComplex[BACnetContextTagCharacterString](BACnetContextTagParseWithBufferProducer[BACnetContextTagCharacterString]((uint8)(uint8(7)), (BACnetDataType)(BACnetDataType_CHARACTER_STRING)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'messageText' field"))
	}
	var messageText BACnetContextTagCharacterString
	if _messageText != nil {
		messageText = *_messageText
	}

	notifyType, err := ReadSimpleField[BACnetNotifyTypeTagged](ctx, "notifyType", ReadComplex[BACnetNotifyTypeTagged](BACnetNotifyTypeTaggedParseWithBufferProducer((uint8)(uint8(8)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'notifyType' field"))
	}

	_ackRequired, err := ReadOptionalField[BACnetContextTagBoolean](ctx, "ackRequired", ReadComplex[BACnetContextTagBoolean](BACnetContextTagParseWithBufferProducer[BACnetContextTagBoolean]((uint8)(uint8(9)), (BACnetDataType)(BACnetDataType_BOOLEAN)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ackRequired' field"))
	}
	var ackRequired BACnetContextTagBoolean
	if _ackRequired != nil {
		ackRequired = *_ackRequired
	}

	_fromState, err := ReadOptionalField[BACnetEventStateTagged](ctx, "fromState", ReadComplex[BACnetEventStateTagged](BACnetEventStateTaggedParseWithBufferProducer((uint8)(uint8(10)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fromState' field"))
	}
	var fromState BACnetEventStateTagged
	if _fromState != nil {
		fromState = *_fromState
	}

	toState, err := ReadSimpleField[BACnetEventStateTagged](ctx, "toState", ReadComplex[BACnetEventStateTagged](BACnetEventStateTaggedParseWithBufferProducer((uint8)(uint8(11)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toState' field"))
	}

	_eventValues, err := ReadOptionalField[BACnetNotificationParameters](ctx, "eventValues", ReadComplex[BACnetNotificationParameters](BACnetNotificationParametersParseWithBufferProducer[BACnetNotificationParameters]((uint8)(uint8(12)), (BACnetObjectType)(eventObjectIdentifier.GetObjectType())), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventValues' field"))
	}
	var eventValues BACnetNotificationParameters
	if _eventValues != nil {
		eventValues = *_eventValues
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestConfirmedEventNotification"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestConfirmedEventNotification")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestConfirmedEventNotification{
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		ProcessIdentifier:          processIdentifier,
		InitiatingDeviceIdentifier: initiatingDeviceIdentifier,
		EventObjectIdentifier:      eventObjectIdentifier,
		Timestamp:                  timestamp,
		NotificationClass:          notificationClass,
		Priority:                   priority,
		EventType:                  eventType,
		MessageText:                messageText,
		NotifyType:                 notifyType,
		AckRequired:                ackRequired,
		FromState:                  fromState,
		ToState:                    toState,
		EventValues:                eventValues,
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedEventNotification) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedEventNotification) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestConfirmedEventNotification"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestConfirmedEventNotification")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "processIdentifier", m.GetProcessIdentifier(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'processIdentifier' field")
		}

		if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "initiatingDeviceIdentifier", m.GetInitiatingDeviceIdentifier(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'initiatingDeviceIdentifier' field")
		}

		if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "eventObjectIdentifier", m.GetEventObjectIdentifier(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'eventObjectIdentifier' field")
		}

		if err := WriteSimpleField[BACnetTimeStampEnclosed](ctx, "timestamp", m.GetTimestamp(), WriteComplex[BACnetTimeStampEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timestamp' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "notificationClass", m.GetNotificationClass(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'notificationClass' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "priority", m.GetPriority(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'priority' field")
		}

		if err := WriteSimpleField[BACnetEventTypeTagged](ctx, "eventType", m.GetEventType(), WriteComplex[BACnetEventTypeTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'eventType' field")
		}

		if err := WriteOptionalField[BACnetContextTagCharacterString](ctx, "messageText", GetRef(m.GetMessageText()), WriteComplex[BACnetContextTagCharacterString](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'messageText' field")
		}

		if err := WriteSimpleField[BACnetNotifyTypeTagged](ctx, "notifyType", m.GetNotifyType(), WriteComplex[BACnetNotifyTypeTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'notifyType' field")
		}

		if err := WriteOptionalField[BACnetContextTagBoolean](ctx, "ackRequired", GetRef(m.GetAckRequired()), WriteComplex[BACnetContextTagBoolean](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'ackRequired' field")
		}

		if err := WriteOptionalField[BACnetEventStateTagged](ctx, "fromState", GetRef(m.GetFromState()), WriteComplex[BACnetEventStateTagged](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'fromState' field")
		}

		if err := WriteSimpleField[BACnetEventStateTagged](ctx, "toState", m.GetToState(), WriteComplex[BACnetEventStateTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'toState' field")
		}

		if err := WriteOptionalField[BACnetNotificationParameters](ctx, "eventValues", GetRef(m.GetEventValues()), WriteComplex[BACnetNotificationParameters](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'eventValues' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestConfirmedEventNotification"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestConfirmedEventNotification")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestConfirmedEventNotification) isBACnetConfirmedServiceRequestConfirmedEventNotification() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestConfirmedEventNotification) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
