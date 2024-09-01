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

// BACnetCOVSubscription is the corresponding interface of BACnetCOVSubscription
type BACnetCOVSubscription interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetRecipient returns Recipient (property field)
	GetRecipient() BACnetRecipientProcessEnclosed
	// GetMonitoredPropertyReference returns MonitoredPropertyReference (property field)
	GetMonitoredPropertyReference() BACnetObjectPropertyReferenceEnclosed
	// GetIssueConfirmedNotifications returns IssueConfirmedNotifications (property field)
	GetIssueConfirmedNotifications() BACnetContextTagBoolean
	// GetTimeRemaining returns TimeRemaining (property field)
	GetTimeRemaining() BACnetContextTagUnsignedInteger
	// GetCovIncrement returns CovIncrement (property field)
	GetCovIncrement() BACnetContextTagReal
}

// BACnetCOVSubscriptionExactly can be used when we want exactly this type and not a type which fulfills BACnetCOVSubscription.
// This is useful for switch cases.
type BACnetCOVSubscriptionExactly interface {
	BACnetCOVSubscription
	isBACnetCOVSubscription() bool
}

// _BACnetCOVSubscription is the data-structure of this message
type _BACnetCOVSubscription struct {
	Recipient                   BACnetRecipientProcessEnclosed
	MonitoredPropertyReference  BACnetObjectPropertyReferenceEnclosed
	IssueConfirmedNotifications BACnetContextTagBoolean
	TimeRemaining               BACnetContextTagUnsignedInteger
	CovIncrement                BACnetContextTagReal
}

var _ BACnetCOVSubscription = (*_BACnetCOVSubscription)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetCOVSubscription) GetRecipient() BACnetRecipientProcessEnclosed {
	return m.Recipient
}

func (m *_BACnetCOVSubscription) GetMonitoredPropertyReference() BACnetObjectPropertyReferenceEnclosed {
	return m.MonitoredPropertyReference
}

func (m *_BACnetCOVSubscription) GetIssueConfirmedNotifications() BACnetContextTagBoolean {
	return m.IssueConfirmedNotifications
}

func (m *_BACnetCOVSubscription) GetTimeRemaining() BACnetContextTagUnsignedInteger {
	return m.TimeRemaining
}

func (m *_BACnetCOVSubscription) GetCovIncrement() BACnetContextTagReal {
	return m.CovIncrement
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetCOVSubscription factory function for _BACnetCOVSubscription
func NewBACnetCOVSubscription(recipient BACnetRecipientProcessEnclosed, monitoredPropertyReference BACnetObjectPropertyReferenceEnclosed, issueConfirmedNotifications BACnetContextTagBoolean, timeRemaining BACnetContextTagUnsignedInteger, covIncrement BACnetContextTagReal) *_BACnetCOVSubscription {
	return &_BACnetCOVSubscription{Recipient: recipient, MonitoredPropertyReference: monitoredPropertyReference, IssueConfirmedNotifications: issueConfirmedNotifications, TimeRemaining: timeRemaining, CovIncrement: covIncrement}
}

// Deprecated: use the interface for direct cast
func CastBACnetCOVSubscription(structType any) BACnetCOVSubscription {
	if casted, ok := structType.(BACnetCOVSubscription); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetCOVSubscription); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetCOVSubscription) GetTypeName() string {
	return "BACnetCOVSubscription"
}

func (m *_BACnetCOVSubscription) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (recipient)
	lengthInBits += m.Recipient.GetLengthInBits(ctx)

	// Simple field (monitoredPropertyReference)
	lengthInBits += m.MonitoredPropertyReference.GetLengthInBits(ctx)

	// Simple field (issueConfirmedNotifications)
	lengthInBits += m.IssueConfirmedNotifications.GetLengthInBits(ctx)

	// Simple field (timeRemaining)
	lengthInBits += m.TimeRemaining.GetLengthInBits(ctx)

	// Optional Field (covIncrement)
	if m.CovIncrement != nil {
		lengthInBits += m.CovIncrement.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetCOVSubscription) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetCOVSubscriptionParse(ctx context.Context, theBytes []byte) (BACnetCOVSubscription, error) {
	return BACnetCOVSubscriptionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetCOVSubscriptionParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCOVSubscription, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCOVSubscription, error) {
		return BACnetCOVSubscriptionParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetCOVSubscriptionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCOVSubscription, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetCOVSubscription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetCOVSubscription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	recipient, err := ReadSimpleField[BACnetRecipientProcessEnclosed](ctx, "recipient", ReadComplex[BACnetRecipientProcessEnclosed](BACnetRecipientProcessEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'recipient' field"))
	}

	monitoredPropertyReference, err := ReadSimpleField[BACnetObjectPropertyReferenceEnclosed](ctx, "monitoredPropertyReference", ReadComplex[BACnetObjectPropertyReferenceEnclosed](BACnetObjectPropertyReferenceEnclosedParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'monitoredPropertyReference' field"))
	}

	issueConfirmedNotifications, err := ReadSimpleField[BACnetContextTagBoolean](ctx, "issueConfirmedNotifications", ReadComplex[BACnetContextTagBoolean](BACnetContextTagParseWithBufferProducer[BACnetContextTagBoolean]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_BOOLEAN)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'issueConfirmedNotifications' field"))
	}

	timeRemaining, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeRemaining", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeRemaining' field"))
	}

	_covIncrement, err := ReadOptionalField[BACnetContextTagReal](ctx, "covIncrement", ReadComplex[BACnetContextTagReal](BACnetContextTagParseWithBufferProducer[BACnetContextTagReal]((uint8)(uint8(4)), (BACnetDataType)(BACnetDataType_REAL)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'covIncrement' field"))
	}
	var covIncrement BACnetContextTagReal
	if _covIncrement != nil {
		covIncrement = *_covIncrement
	}

	if closeErr := readBuffer.CloseContext("BACnetCOVSubscription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetCOVSubscription")
	}

	// Create the instance
	return &_BACnetCOVSubscription{
		Recipient:                   recipient,
		MonitoredPropertyReference:  monitoredPropertyReference,
		IssueConfirmedNotifications: issueConfirmedNotifications,
		TimeRemaining:               timeRemaining,
		CovIncrement:                covIncrement,
	}, nil
}

func (m *_BACnetCOVSubscription) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetCOVSubscription) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetCOVSubscription"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetCOVSubscription")
	}

	if err := WriteSimpleField[BACnetRecipientProcessEnclosed](ctx, "recipient", m.GetRecipient(), WriteComplex[BACnetRecipientProcessEnclosed](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'recipient' field")
	}

	if err := WriteSimpleField[BACnetObjectPropertyReferenceEnclosed](ctx, "monitoredPropertyReference", m.GetMonitoredPropertyReference(), WriteComplex[BACnetObjectPropertyReferenceEnclosed](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'monitoredPropertyReference' field")
	}

	if err := WriteSimpleField[BACnetContextTagBoolean](ctx, "issueConfirmedNotifications", m.GetIssueConfirmedNotifications(), WriteComplex[BACnetContextTagBoolean](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'issueConfirmedNotifications' field")
	}

	if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeRemaining", m.GetTimeRemaining(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'timeRemaining' field")
	}

	if err := WriteOptionalField[BACnetContextTagReal](ctx, "covIncrement", GetRef(m.GetCovIncrement()), WriteComplex[BACnetContextTagReal](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'covIncrement' field")
	}

	if popErr := writeBuffer.PopContext("BACnetCOVSubscription"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetCOVSubscription")
	}
	return nil
}

func (m *_BACnetCOVSubscription) isBACnetCOVSubscription() bool {
	return true
}

func (m *_BACnetCOVSubscription) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
