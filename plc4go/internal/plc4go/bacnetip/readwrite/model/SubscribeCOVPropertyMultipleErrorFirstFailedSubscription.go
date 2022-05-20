/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// SubscribeCOVPropertyMultipleErrorFirstFailedSubscription is the data-structure of this message
type SubscribeCOVPropertyMultipleErrorFirstFailedSubscription struct {
	OpeningTag                 *BACnetOpeningTag
	MonitoredObjectIdentifier  *BACnetContextTagObjectIdentifier
	MonitoredPropertyReference *BACnetPropertyReferenceEnclosed
	ErrorType                  *ErrorEnclosed
	ClosingTag                 *BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

// ISubscribeCOVPropertyMultipleErrorFirstFailedSubscription is the corresponding interface of SubscribeCOVPropertyMultipleErrorFirstFailedSubscription
type ISubscribeCOVPropertyMultipleErrorFirstFailedSubscription interface {
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetMonitoredObjectIdentifier returns MonitoredObjectIdentifier (property field)
	GetMonitoredObjectIdentifier() *BACnetContextTagObjectIdentifier
	// GetMonitoredPropertyReference returns MonitoredPropertyReference (property field)
	GetMonitoredPropertyReference() *BACnetPropertyReferenceEnclosed
	// GetErrorType returns ErrorType (property field)
	GetErrorType() *ErrorEnclosed
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() *BACnetClosingTag
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetMonitoredObjectIdentifier() *BACnetContextTagObjectIdentifier {
	return m.MonitoredObjectIdentifier
}

func (m *SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetMonitoredPropertyReference() *BACnetPropertyReferenceEnclosed {
	return m.MonitoredPropertyReference
}

func (m *SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetErrorType() *ErrorEnclosed {
	return m.ErrorType
}

func (m *SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSubscribeCOVPropertyMultipleErrorFirstFailedSubscription factory function for SubscribeCOVPropertyMultipleErrorFirstFailedSubscription
func NewSubscribeCOVPropertyMultipleErrorFirstFailedSubscription(openingTag *BACnetOpeningTag, monitoredObjectIdentifier *BACnetContextTagObjectIdentifier, monitoredPropertyReference *BACnetPropertyReferenceEnclosed, errorType *ErrorEnclosed, closingTag *BACnetClosingTag, tagNumber uint8) *SubscribeCOVPropertyMultipleErrorFirstFailedSubscription {
	return &SubscribeCOVPropertyMultipleErrorFirstFailedSubscription{OpeningTag: openingTag, MonitoredObjectIdentifier: monitoredObjectIdentifier, MonitoredPropertyReference: monitoredPropertyReference, ErrorType: errorType, ClosingTag: closingTag, TagNumber: tagNumber}
}

func CastSubscribeCOVPropertyMultipleErrorFirstFailedSubscription(structType interface{}) *SubscribeCOVPropertyMultipleErrorFirstFailedSubscription {
	if casted, ok := structType.(SubscribeCOVPropertyMultipleErrorFirstFailedSubscription); ok {
		return &casted
	}
	if casted, ok := structType.(*SubscribeCOVPropertyMultipleErrorFirstFailedSubscription); ok {
		return casted
	}
	return nil
}

func (m *SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetTypeName() string {
	return "SubscribeCOVPropertyMultipleErrorFirstFailedSubscription"
}

func (m *SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (monitoredObjectIdentifier)
	lengthInBits += m.MonitoredObjectIdentifier.GetLengthInBits()

	// Simple field (monitoredPropertyReference)
	lengthInBits += m.MonitoredPropertyReference.GetLengthInBits()

	// Simple field (errorType)
	lengthInBits += m.ErrorType.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*SubscribeCOVPropertyMultipleErrorFirstFailedSubscription, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SubscribeCOVPropertyMultipleErrorFirstFailedSubscription"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, pullErr
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (monitoredObjectIdentifier)
	if pullErr := readBuffer.PullContext("monitoredObjectIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_monitoredObjectIdentifier, _monitoredObjectIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _monitoredObjectIdentifierErr != nil {
		return nil, errors.Wrap(_monitoredObjectIdentifierErr, "Error parsing 'monitoredObjectIdentifier' field")
	}
	monitoredObjectIdentifier := CastBACnetContextTagObjectIdentifier(_monitoredObjectIdentifier)
	if closeErr := readBuffer.CloseContext("monitoredObjectIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (monitoredPropertyReference)
	if pullErr := readBuffer.PullContext("monitoredPropertyReference"); pullErr != nil {
		return nil, pullErr
	}
	_monitoredPropertyReference, _monitoredPropertyReferenceErr := BACnetPropertyReferenceEnclosedParse(readBuffer, uint8(uint8(1)))
	if _monitoredPropertyReferenceErr != nil {
		return nil, errors.Wrap(_monitoredPropertyReferenceErr, "Error parsing 'monitoredPropertyReference' field")
	}
	monitoredPropertyReference := CastBACnetPropertyReferenceEnclosed(_monitoredPropertyReference)
	if closeErr := readBuffer.CloseContext("monitoredPropertyReference"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (errorType)
	if pullErr := readBuffer.PullContext("errorType"); pullErr != nil {
		return nil, pullErr
	}
	_errorType, _errorTypeErr := ErrorEnclosedParse(readBuffer, uint8(uint8(2)))
	if _errorTypeErr != nil {
		return nil, errors.Wrap(_errorTypeErr, "Error parsing 'errorType' field")
	}
	errorType := CastErrorEnclosed(_errorType)
	if closeErr := readBuffer.CloseContext("errorType"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, pullErr
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("SubscribeCOVPropertyMultipleErrorFirstFailedSubscription"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewSubscribeCOVPropertyMultipleErrorFirstFailedSubscription(openingTag, monitoredObjectIdentifier, monitoredPropertyReference, errorType, closingTag, tagNumber), nil
}

func (m *SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("SubscribeCOVPropertyMultipleErrorFirstFailedSubscription"); pushErr != nil {
		return pushErr
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return pushErr
	}
	_openingTagErr := m.OpeningTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return popErr
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Simple Field (monitoredObjectIdentifier)
	if pushErr := writeBuffer.PushContext("monitoredObjectIdentifier"); pushErr != nil {
		return pushErr
	}
	_monitoredObjectIdentifierErr := m.MonitoredObjectIdentifier.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("monitoredObjectIdentifier"); popErr != nil {
		return popErr
	}
	if _monitoredObjectIdentifierErr != nil {
		return errors.Wrap(_monitoredObjectIdentifierErr, "Error serializing 'monitoredObjectIdentifier' field")
	}

	// Simple Field (monitoredPropertyReference)
	if pushErr := writeBuffer.PushContext("monitoredPropertyReference"); pushErr != nil {
		return pushErr
	}
	_monitoredPropertyReferenceErr := m.MonitoredPropertyReference.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("monitoredPropertyReference"); popErr != nil {
		return popErr
	}
	if _monitoredPropertyReferenceErr != nil {
		return errors.Wrap(_monitoredPropertyReferenceErr, "Error serializing 'monitoredPropertyReference' field")
	}

	// Simple Field (errorType)
	if pushErr := writeBuffer.PushContext("errorType"); pushErr != nil {
		return pushErr
	}
	_errorTypeErr := m.ErrorType.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("errorType"); popErr != nil {
		return popErr
	}
	if _errorTypeErr != nil {
		return errors.Wrap(_errorTypeErr, "Error serializing 'errorType' field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return pushErr
	}
	_closingTagErr := m.ClosingTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return popErr
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("SubscribeCOVPropertyMultipleErrorFirstFailedSubscription"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
