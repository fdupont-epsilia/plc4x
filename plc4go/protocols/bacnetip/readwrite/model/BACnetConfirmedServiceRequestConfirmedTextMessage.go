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

// BACnetConfirmedServiceRequestConfirmedTextMessage is the corresponding interface of BACnetConfirmedServiceRequestConfirmedTextMessage
type BACnetConfirmedServiceRequestConfirmedTextMessage interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetTextMessageSourceDevice returns TextMessageSourceDevice (property field)
	GetTextMessageSourceDevice() BACnetContextTagObjectIdentifier
	// GetMessageClass returns MessageClass (property field)
	GetMessageClass() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
	// GetMessagePriority returns MessagePriority (property field)
	GetMessagePriority() BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged
	// GetMessage returns Message (property field)
	GetMessage() BACnetContextTagCharacterString
}

// BACnetConfirmedServiceRequestConfirmedTextMessageExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestConfirmedTextMessage.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestConfirmedTextMessageExactly interface {
	BACnetConfirmedServiceRequestConfirmedTextMessage
	isBACnetConfirmedServiceRequestConfirmedTextMessage() bool
}

// _BACnetConfirmedServiceRequestConfirmedTextMessage is the data-structure of this message
type _BACnetConfirmedServiceRequestConfirmedTextMessage struct {
	*_BACnetConfirmedServiceRequest
	TextMessageSourceDevice BACnetContextTagObjectIdentifier
	MessageClass            BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
	MessagePriority         BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged
	Message                 BACnetContextTagCharacterString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessage) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_CONFIRMED_TEXT_MESSAGE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessage) InitializeParent(parent BACnetConfirmedServiceRequest) {
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessage) GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessage) GetTextMessageSourceDevice() BACnetContextTagObjectIdentifier {
	return m.TextMessageSourceDevice
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessage) GetMessageClass() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass {
	return m.MessageClass
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessage) GetMessagePriority() BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged {
	return m.MessagePriority
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessage) GetMessage() BACnetContextTagCharacterString {
	return m.Message
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestConfirmedTextMessage factory function for _BACnetConfirmedServiceRequestConfirmedTextMessage
func NewBACnetConfirmedServiceRequestConfirmedTextMessage(textMessageSourceDevice BACnetContextTagObjectIdentifier, messageClass BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass, messagePriority BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged, message BACnetContextTagCharacterString, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestConfirmedTextMessage {
	_result := &_BACnetConfirmedServiceRequestConfirmedTextMessage{
		TextMessageSourceDevice:        textMessageSourceDevice,
		MessageClass:                   messageClass,
		MessagePriority:                messagePriority,
		Message:                        message,
		_BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestConfirmedTextMessage(structType any) BACnetConfirmedServiceRequestConfirmedTextMessage {
	if casted, ok := structType.(BACnetConfirmedServiceRequestConfirmedTextMessage); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestConfirmedTextMessage); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessage) GetTypeName() string {
	return "BACnetConfirmedServiceRequestConfirmedTextMessage"
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessage) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (textMessageSourceDevice)
	lengthInBits += m.TextMessageSourceDevice.GetLengthInBits(ctx)

	// Optional Field (messageClass)
	if m.MessageClass != nil {
		lengthInBits += m.MessageClass.GetLengthInBits(ctx)
	}

	// Simple field (messagePriority)
	lengthInBits += m.MessagePriority.GetLengthInBits(ctx)

	// Simple field (message)
	lengthInBits += m.Message.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessage) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestConfirmedTextMessageParse(ctx context.Context, theBytes []byte, serviceRequestLength uint32) (BACnetConfirmedServiceRequestConfirmedTextMessage, error) {
	return BACnetConfirmedServiceRequestConfirmedTextMessageParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), serviceRequestLength)
}

func BACnetConfirmedServiceRequestConfirmedTextMessageParseWithBufferProducer(serviceRequestLength uint32) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestConfirmedTextMessage, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestConfirmedTextMessage, error) {
		return BACnetConfirmedServiceRequestConfirmedTextMessageParseWithBuffer(ctx, readBuffer, serviceRequestLength)
	}
}

func BACnetConfirmedServiceRequestConfirmedTextMessageParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceRequestLength uint32) (BACnetConfirmedServiceRequestConfirmedTextMessage, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestConfirmedTextMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestConfirmedTextMessage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	textMessageSourceDevice, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "textMessageSourceDevice", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'textMessageSourceDevice' field"))
	}

	_messageClass, err := ReadOptionalField[BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass](ctx, "messageClass", ReadComplex[BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass](BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassParseWithBufferProducer[BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass]((uint8)(uint8(1))), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'messageClass' field"))
	}
	var messageClass BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
	if _messageClass != nil {
		messageClass = *_messageClass
	}

	messagePriority, err := ReadSimpleField[BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged](ctx, "messagePriority", ReadComplex[BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged](BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedParseWithBufferProducer((uint8)(uint8(2)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'messagePriority' field"))
	}

	message, err := ReadSimpleField[BACnetContextTagCharacterString](ctx, "message", ReadComplex[BACnetContextTagCharacterString](BACnetContextTagParseWithBufferProducer((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_CHARACTER_STRING)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'message' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestConfirmedTextMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestConfirmedTextMessage")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestConfirmedTextMessage{
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		TextMessageSourceDevice: textMessageSourceDevice,
		MessageClass:            messageClass,
		MessagePriority:         messagePriority,
		Message:                 message,
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessage) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessage) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestConfirmedTextMessage"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestConfirmedTextMessage")
		}

		// Simple Field (textMessageSourceDevice)
		if pushErr := writeBuffer.PushContext("textMessageSourceDevice"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for textMessageSourceDevice")
		}
		_textMessageSourceDeviceErr := writeBuffer.WriteSerializable(ctx, m.GetTextMessageSourceDevice())
		if popErr := writeBuffer.PopContext("textMessageSourceDevice"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for textMessageSourceDevice")
		}
		if _textMessageSourceDeviceErr != nil {
			return errors.Wrap(_textMessageSourceDeviceErr, "Error serializing 'textMessageSourceDevice' field")
		}

		// Optional Field (messageClass) (Can be skipped, if the value is null)
		var messageClass BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass = nil
		if m.GetMessageClass() != nil {
			if pushErr := writeBuffer.PushContext("messageClass"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for messageClass")
			}
			messageClass = m.GetMessageClass()
			_messageClassErr := writeBuffer.WriteSerializable(ctx, messageClass)
			if popErr := writeBuffer.PopContext("messageClass"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for messageClass")
			}
			if _messageClassErr != nil {
				return errors.Wrap(_messageClassErr, "Error serializing 'messageClass' field")
			}
		}

		// Simple Field (messagePriority)
		if pushErr := writeBuffer.PushContext("messagePriority"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for messagePriority")
		}
		_messagePriorityErr := writeBuffer.WriteSerializable(ctx, m.GetMessagePriority())
		if popErr := writeBuffer.PopContext("messagePriority"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for messagePriority")
		}
		if _messagePriorityErr != nil {
			return errors.Wrap(_messagePriorityErr, "Error serializing 'messagePriority' field")
		}

		// Simple Field (message)
		if pushErr := writeBuffer.PushContext("message"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for message")
		}
		_messageErr := writeBuffer.WriteSerializable(ctx, m.GetMessage())
		if popErr := writeBuffer.PopContext("message"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for message")
		}
		if _messageErr != nil {
			return errors.Wrap(_messageErr, "Error serializing 'message' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestConfirmedTextMessage"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestConfirmedTextMessage")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessage) isBACnetConfirmedServiceRequestConfirmedTextMessage() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessage) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
