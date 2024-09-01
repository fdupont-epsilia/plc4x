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

// RepublishRequest is the corresponding interface of RepublishRequest
type RepublishRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetSubscriptionId returns SubscriptionId (property field)
	GetSubscriptionId() uint32
	// GetRetransmitSequenceNumber returns RetransmitSequenceNumber (property field)
	GetRetransmitSequenceNumber() uint32
}

// RepublishRequestExactly can be used when we want exactly this type and not a type which fulfills RepublishRequest.
// This is useful for switch cases.
type RepublishRequestExactly interface {
	RepublishRequest
	isRepublishRequest() bool
}

// _RepublishRequest is the data-structure of this message
type _RepublishRequest struct {
	*_ExtensionObjectDefinition
	RequestHeader            ExtensionObjectDefinition
	SubscriptionId           uint32
	RetransmitSequenceNumber uint32
}

var _ RepublishRequest = (*_RepublishRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_RepublishRequest) GetIdentifier() string {
	return "832"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RepublishRequest) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_RepublishRequest) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RepublishRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_RepublishRequest) GetSubscriptionId() uint32 {
	return m.SubscriptionId
}

func (m *_RepublishRequest) GetRetransmitSequenceNumber() uint32 {
	return m.RetransmitSequenceNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewRepublishRequest factory function for _RepublishRequest
func NewRepublishRequest(requestHeader ExtensionObjectDefinition, subscriptionId uint32, retransmitSequenceNumber uint32) *_RepublishRequest {
	_result := &_RepublishRequest{
		RequestHeader:              requestHeader,
		SubscriptionId:             subscriptionId,
		RetransmitSequenceNumber:   retransmitSequenceNumber,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastRepublishRequest(structType any) RepublishRequest {
	if casted, ok := structType.(RepublishRequest); ok {
		return casted
	}
	if casted, ok := structType.(*RepublishRequest); ok {
		return *casted
	}
	return nil
}

func (m *_RepublishRequest) GetTypeName() string {
	return "RepublishRequest"
}

func (m *_RepublishRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (subscriptionId)
	lengthInBits += 32

	// Simple field (retransmitSequenceNumber)
	lengthInBits += 32

	return lengthInBits
}

func (m *_RepublishRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func RepublishRequestParse(ctx context.Context, theBytes []byte, identifier string) (RepublishRequest, error) {
	return RepublishRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func RepublishRequestParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (RepublishRequest, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (RepublishRequest, error) {
		return RepublishRequestParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func RepublishRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (RepublishRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RepublishRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RepublishRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}

	subscriptionId, err := ReadSimpleField(ctx, "subscriptionId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subscriptionId' field"))
	}

	retransmitSequenceNumber, err := ReadSimpleField(ctx, "retransmitSequenceNumber", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'retransmitSequenceNumber' field"))
	}

	if closeErr := readBuffer.CloseContext("RepublishRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RepublishRequest")
	}

	// Create a partially initialized instance
	_child := &_RepublishRequest{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		RequestHeader:              requestHeader,
		SubscriptionId:             subscriptionId,
		RetransmitSequenceNumber:   retransmitSequenceNumber,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_RepublishRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RepublishRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RepublishRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RepublishRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[uint32](ctx, "subscriptionId", m.GetSubscriptionId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'subscriptionId' field")
		}

		if err := WriteSimpleField[uint32](ctx, "retransmitSequenceNumber", m.GetRetransmitSequenceNumber(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'retransmitSequenceNumber' field")
		}

		if popErr := writeBuffer.PopContext("RepublishRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RepublishRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RepublishRequest) isRepublishRequest() bool {
	return true
}

func (m *_RepublishRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
