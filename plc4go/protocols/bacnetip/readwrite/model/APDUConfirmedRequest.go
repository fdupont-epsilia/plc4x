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

// APDUConfirmedRequest is the corresponding interface of APDUConfirmedRequest
type APDUConfirmedRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	APDU
	// GetSegmentedMessage returns SegmentedMessage (property field)
	GetSegmentedMessage() bool
	// GetMoreFollows returns MoreFollows (property field)
	GetMoreFollows() bool
	// GetSegmentedResponseAccepted returns SegmentedResponseAccepted (property field)
	GetSegmentedResponseAccepted() bool
	// GetMaxSegmentsAccepted returns MaxSegmentsAccepted (property field)
	GetMaxSegmentsAccepted() MaxSegmentsAccepted
	// GetMaxApduLengthAccepted returns MaxApduLengthAccepted (property field)
	GetMaxApduLengthAccepted() MaxApduLengthAccepted
	// GetInvokeId returns InvokeId (property field)
	GetInvokeId() uint8
	// GetSequenceNumber returns SequenceNumber (property field)
	GetSequenceNumber() *uint8
	// GetProposedWindowSize returns ProposedWindowSize (property field)
	GetProposedWindowSize() *uint8
	// GetServiceRequest returns ServiceRequest (property field)
	GetServiceRequest() BACnetConfirmedServiceRequest
	// GetSegmentServiceChoice returns SegmentServiceChoice (property field)
	GetSegmentServiceChoice() *BACnetConfirmedServiceChoice
	// GetSegment returns Segment (property field)
	GetSegment() []byte
	// GetApduHeaderReduction returns ApduHeaderReduction (virtual field)
	GetApduHeaderReduction() uint16
	// GetSegmentReduction returns SegmentReduction (virtual field)
	GetSegmentReduction() uint16
}

// APDUConfirmedRequestExactly can be used when we want exactly this type and not a type which fulfills APDUConfirmedRequest.
// This is useful for switch cases.
type APDUConfirmedRequestExactly interface {
	APDUConfirmedRequest
	isAPDUConfirmedRequest() bool
}

// _APDUConfirmedRequest is the data-structure of this message
type _APDUConfirmedRequest struct {
	*_APDU
	SegmentedMessage          bool
	MoreFollows               bool
	SegmentedResponseAccepted bool
	MaxSegmentsAccepted       MaxSegmentsAccepted
	MaxApduLengthAccepted     MaxApduLengthAccepted
	InvokeId                  uint8
	SequenceNumber            *uint8
	ProposedWindowSize        *uint8
	ServiceRequest            BACnetConfirmedServiceRequest
	SegmentServiceChoice      *BACnetConfirmedServiceChoice
	Segment                   []byte
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_APDUConfirmedRequest) GetApduType() ApduType {
	return ApduType_CONFIRMED_REQUEST_PDU
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_APDUConfirmedRequest) InitializeParent(parent APDU) {}

func (m *_APDUConfirmedRequest) GetParent() APDU {
	return m._APDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_APDUConfirmedRequest) GetSegmentedMessage() bool {
	return m.SegmentedMessage
}

func (m *_APDUConfirmedRequest) GetMoreFollows() bool {
	return m.MoreFollows
}

func (m *_APDUConfirmedRequest) GetSegmentedResponseAccepted() bool {
	return m.SegmentedResponseAccepted
}

func (m *_APDUConfirmedRequest) GetMaxSegmentsAccepted() MaxSegmentsAccepted {
	return m.MaxSegmentsAccepted
}

func (m *_APDUConfirmedRequest) GetMaxApduLengthAccepted() MaxApduLengthAccepted {
	return m.MaxApduLengthAccepted
}

func (m *_APDUConfirmedRequest) GetInvokeId() uint8 {
	return m.InvokeId
}

func (m *_APDUConfirmedRequest) GetSequenceNumber() *uint8 {
	return m.SequenceNumber
}

func (m *_APDUConfirmedRequest) GetProposedWindowSize() *uint8 {
	return m.ProposedWindowSize
}

func (m *_APDUConfirmedRequest) GetServiceRequest() BACnetConfirmedServiceRequest {
	return m.ServiceRequest
}

func (m *_APDUConfirmedRequest) GetSegmentServiceChoice() *BACnetConfirmedServiceChoice {
	return m.SegmentServiceChoice
}

func (m *_APDUConfirmedRequest) GetSegment() []byte {
	return m.Segment
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_APDUConfirmedRequest) GetApduHeaderReduction() uint16 {
	ctx := context.Background()
	_ = ctx
	sequenceNumber := m.SequenceNumber
	_ = sequenceNumber
	proposedWindowSize := m.ProposedWindowSize
	_ = proposedWindowSize
	serviceRequest := m.ServiceRequest
	_ = serviceRequest
	segmentServiceChoice := m.SegmentServiceChoice
	_ = segmentServiceChoice
	return uint16(uint16(uint16(3)) + uint16((utils.InlineIf(m.GetSegmentedMessage(), func() any { return uint16(uint16(2)) }, func() any { return uint16(uint16(0)) }).(uint16))))
}

func (m *_APDUConfirmedRequest) GetSegmentReduction() uint16 {
	ctx := context.Background()
	_ = ctx
	sequenceNumber := m.SequenceNumber
	_ = sequenceNumber
	proposedWindowSize := m.ProposedWindowSize
	_ = proposedWindowSize
	serviceRequest := m.ServiceRequest
	_ = serviceRequest
	segmentServiceChoice := m.SegmentServiceChoice
	_ = segmentServiceChoice
	return uint16(utils.InlineIf((bool((m.GetSegmentServiceChoice()) != (nil))), func() any { return uint16((uint16(m.GetApduHeaderReduction()) + uint16(uint16(1)))) }, func() any { return uint16(m.GetApduHeaderReduction()) }).(uint16))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAPDUConfirmedRequest factory function for _APDUConfirmedRequest
func NewAPDUConfirmedRequest(segmentedMessage bool, moreFollows bool, segmentedResponseAccepted bool, maxSegmentsAccepted MaxSegmentsAccepted, maxApduLengthAccepted MaxApduLengthAccepted, invokeId uint8, sequenceNumber *uint8, proposedWindowSize *uint8, serviceRequest BACnetConfirmedServiceRequest, segmentServiceChoice *BACnetConfirmedServiceChoice, segment []byte, apduLength uint16) *_APDUConfirmedRequest {
	_result := &_APDUConfirmedRequest{
		SegmentedMessage:          segmentedMessage,
		MoreFollows:               moreFollows,
		SegmentedResponseAccepted: segmentedResponseAccepted,
		MaxSegmentsAccepted:       maxSegmentsAccepted,
		MaxApduLengthAccepted:     maxApduLengthAccepted,
		InvokeId:                  invokeId,
		SequenceNumber:            sequenceNumber,
		ProposedWindowSize:        proposedWindowSize,
		ServiceRequest:            serviceRequest,
		SegmentServiceChoice:      segmentServiceChoice,
		Segment:                   segment,
		_APDU:                     NewAPDU(apduLength),
	}
	_result._APDU._APDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAPDUConfirmedRequest(structType any) APDUConfirmedRequest {
	if casted, ok := structType.(APDUConfirmedRequest); ok {
		return casted
	}
	if casted, ok := structType.(*APDUConfirmedRequest); ok {
		return *casted
	}
	return nil
}

func (m *_APDUConfirmedRequest) GetTypeName() string {
	return "APDUConfirmedRequest"
}

func (m *_APDUConfirmedRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (segmentedMessage)
	lengthInBits += 1

	// Simple field (moreFollows)
	lengthInBits += 1

	// Simple field (segmentedResponseAccepted)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 2

	// Simple field (maxSegmentsAccepted)
	lengthInBits += 3

	// Simple field (maxApduLengthAccepted)
	lengthInBits += 4

	// Simple field (invokeId)
	lengthInBits += 8

	// Optional Field (sequenceNumber)
	if m.SequenceNumber != nil {
		lengthInBits += 8
	}

	// Optional Field (proposedWindowSize)
	if m.ProposedWindowSize != nil {
		lengthInBits += 8
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (serviceRequest)
	if m.ServiceRequest != nil {
		lengthInBits += m.ServiceRequest.GetLengthInBits(ctx)
	}

	// Optional Field (segmentServiceChoice)
	if m.SegmentServiceChoice != nil {
		lengthInBits += 8
	}

	// A virtual field doesn't have any in- or output.

	// Array field
	if len(m.Segment) > 0 {
		lengthInBits += 8 * uint16(len(m.Segment))
	}

	return lengthInBits
}

func (m *_APDUConfirmedRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func APDUConfirmedRequestParse(ctx context.Context, theBytes []byte, apduLength uint16) (APDUConfirmedRequest, error) {
	return APDUConfirmedRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), apduLength)
}

func APDUConfirmedRequestParseWithBufferProducer(apduLength uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (APDUConfirmedRequest, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (APDUConfirmedRequest, error) {
		return APDUConfirmedRequestParseWithBuffer(ctx, readBuffer, apduLength)
	}
}

func APDUConfirmedRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, apduLength uint16) (APDUConfirmedRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("APDUConfirmedRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for APDUConfirmedRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	segmentedMessage, err := ReadSimpleField(ctx, "segmentedMessage", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'segmentedMessage' field"))
	}

	moreFollows, err := ReadSimpleField(ctx, "moreFollows", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'moreFollows' field"))
	}

	segmentedResponseAccepted, err := ReadSimpleField(ctx, "segmentedResponseAccepted", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'segmentedResponseAccepted' field"))
	}

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(2)), uint8(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	maxSegmentsAccepted, err := ReadEnumField[MaxSegmentsAccepted](ctx, "maxSegmentsAccepted", "MaxSegmentsAccepted", ReadEnum(MaxSegmentsAcceptedByValue, ReadUnsignedByte(readBuffer, uint8(3))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxSegmentsAccepted' field"))
	}

	maxApduLengthAccepted, err := ReadEnumField[MaxApduLengthAccepted](ctx, "maxApduLengthAccepted", "MaxApduLengthAccepted", ReadEnum(MaxApduLengthAcceptedByValue, ReadUnsignedByte(readBuffer, uint8(4))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxApduLengthAccepted' field"))
	}

	invokeId, err := ReadSimpleField(ctx, "invokeId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'invokeId' field"))
	}

	sequenceNumber, err := ReadOptionalField[uint8](ctx, "sequenceNumber", ReadUnsignedByte(readBuffer, uint8(8)), segmentedMessage)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sequenceNumber' field"))
	}

	proposedWindowSize, err := ReadOptionalField[uint8](ctx, "proposedWindowSize", ReadUnsignedByte(readBuffer, uint8(8)), segmentedMessage)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proposedWindowSize' field"))
	}

	apduHeaderReduction, err := ReadVirtualField[uint16](ctx, "apduHeaderReduction", (*uint16)(nil), uint16(uint16(3))+uint16((utils.InlineIf(segmentedMessage, func() any { return uint16(uint16(2)) }, func() any { return uint16(uint16(0)) }).(uint16))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'apduHeaderReduction' field"))
	}

	_serviceRequest, err := ReadOptionalField[BACnetConfirmedServiceRequest](ctx, "serviceRequest", ReadComplex[BACnetConfirmedServiceRequest](BACnetConfirmedServiceRequestParseWithBufferProducer[BACnetConfirmedServiceRequest]((uint32)(uint32(apduLength)-uint32(apduHeaderReduction))), readBuffer), !(segmentedMessage))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serviceRequest' field"))
	}
	var serviceRequest BACnetConfirmedServiceRequest
	if _serviceRequest != nil {
		serviceRequest = *_serviceRequest
	}

	// Validation
	if !(bool((bool(!(segmentedMessage)) && bool(bool((serviceRequest) != (nil))))) || bool(segmentedMessage)) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "service request should be set"})
	}

	segmentServiceChoice, err := ReadOptionalField[BACnetConfirmedServiceChoice](ctx, "segmentServiceChoice", ReadEnum(BACnetConfirmedServiceChoiceByValue, ReadUnsignedByte(readBuffer, uint8(8))), bool(segmentedMessage) && bool(bool((*sequenceNumber) != (0))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'segmentServiceChoice' field"))
	}

	segmentReduction, err := ReadVirtualField[uint16](ctx, "segmentReduction", (*uint16)(nil), utils.InlineIf((bool((segmentServiceChoice) != (nil))), func() any { return uint16((uint16(apduHeaderReduction) + uint16(uint16(1)))) }, func() any { return uint16(apduHeaderReduction) }).(uint16))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'segmentReduction' field"))
	}

	segment, err := readBuffer.ReadByteArray("segment", int(utils.InlineIf(segmentedMessage, func() any {
		return int32((utils.InlineIf((bool((apduLength) > (0))), func() any { return int32((int32(apduLength) - int32(segmentReduction))) }, func() any { return int32(int32(0)) }).(int32)))
	}, func() any { return int32(int32(0)) }).(int32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'segment' field"))
	}

	if closeErr := readBuffer.CloseContext("APDUConfirmedRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for APDUConfirmedRequest")
	}

	// Create a partially initialized instance
	_child := &_APDUConfirmedRequest{
		_APDU: &_APDU{
			ApduLength: apduLength,
		},
		SegmentedMessage:          segmentedMessage,
		MoreFollows:               moreFollows,
		SegmentedResponseAccepted: segmentedResponseAccepted,
		MaxSegmentsAccepted:       maxSegmentsAccepted,
		MaxApduLengthAccepted:     maxApduLengthAccepted,
		InvokeId:                  invokeId,
		SequenceNumber:            sequenceNumber,
		ProposedWindowSize:        proposedWindowSize,
		ServiceRequest:            serviceRequest,
		SegmentServiceChoice:      segmentServiceChoice,
		Segment:                   segment,
		reservedField0:            reservedField0,
	}
	_child._APDU._APDUChildRequirements = _child
	return _child, nil
}

func (m *_APDUConfirmedRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_APDUConfirmedRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUConfirmedRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for APDUConfirmedRequest")
		}

		// Simple Field (segmentedMessage)
		segmentedMessage := bool(m.GetSegmentedMessage())
		_segmentedMessageErr := /*TODO: migrate me*/ writeBuffer.WriteBit("segmentedMessage", (segmentedMessage))
		if _segmentedMessageErr != nil {
			return errors.Wrap(_segmentedMessageErr, "Error serializing 'segmentedMessage' field")
		}

		// Simple Field (moreFollows)
		moreFollows := bool(m.GetMoreFollows())
		_moreFollowsErr := /*TODO: migrate me*/ writeBuffer.WriteBit("moreFollows", (moreFollows))
		if _moreFollowsErr != nil {
			return errors.Wrap(_moreFollowsErr, "Error serializing 'moreFollows' field")
		}

		// Simple Field (segmentedResponseAccepted)
		segmentedResponseAccepted := bool(m.GetSegmentedResponseAccepted())
		_segmentedResponseAcceptedErr := /*TODO: migrate me*/ writeBuffer.WriteBit("segmentedResponseAccepted", (segmentedResponseAccepted))
		if _segmentedResponseAcceptedErr != nil {
			return errors.Wrap(_segmentedResponseAcceptedErr, "Error serializing 'segmentedResponseAccepted' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0)
			if m.reservedField0 != nil {
				log.Info().Fields(map[string]any{
					"expected value": uint8(0),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := /*TODO: migrate me*/ writeBuffer.WriteUint8("reserved", 2, uint8(reserved))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (maxSegmentsAccepted)
		if pushErr := writeBuffer.PushContext("maxSegmentsAccepted"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for maxSegmentsAccepted")
		}
		_maxSegmentsAcceptedErr := writeBuffer.WriteSerializable(ctx, m.GetMaxSegmentsAccepted())
		if popErr := writeBuffer.PopContext("maxSegmentsAccepted"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for maxSegmentsAccepted")
		}
		if _maxSegmentsAcceptedErr != nil {
			return errors.Wrap(_maxSegmentsAcceptedErr, "Error serializing 'maxSegmentsAccepted' field")
		}

		// Simple Field (maxApduLengthAccepted)
		if pushErr := writeBuffer.PushContext("maxApduLengthAccepted"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for maxApduLengthAccepted")
		}
		_maxApduLengthAcceptedErr := writeBuffer.WriteSerializable(ctx, m.GetMaxApduLengthAccepted())
		if popErr := writeBuffer.PopContext("maxApduLengthAccepted"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for maxApduLengthAccepted")
		}
		if _maxApduLengthAcceptedErr != nil {
			return errors.Wrap(_maxApduLengthAcceptedErr, "Error serializing 'maxApduLengthAccepted' field")
		}

		// Simple Field (invokeId)
		invokeId := uint8(m.GetInvokeId())
		_invokeIdErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("invokeId", 8, uint8((invokeId)))
		if _invokeIdErr != nil {
			return errors.Wrap(_invokeIdErr, "Error serializing 'invokeId' field")
		}

		// Optional Field (sequenceNumber) (Can be skipped, if the value is null)
		var sequenceNumber *uint8 = nil
		if m.GetSequenceNumber() != nil {
			sequenceNumber = m.GetSequenceNumber()
			_sequenceNumberErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("sequenceNumber", 8, uint8(*(sequenceNumber)))
			if _sequenceNumberErr != nil {
				return errors.Wrap(_sequenceNumberErr, "Error serializing 'sequenceNumber' field")
			}
		}

		// Optional Field (proposedWindowSize) (Can be skipped, if the value is null)
		var proposedWindowSize *uint8 = nil
		if m.GetProposedWindowSize() != nil {
			proposedWindowSize = m.GetProposedWindowSize()
			_proposedWindowSizeErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("proposedWindowSize", 8, uint8(*(proposedWindowSize)))
			if _proposedWindowSizeErr != nil {
				return errors.Wrap(_proposedWindowSizeErr, "Error serializing 'proposedWindowSize' field")
			}
		}
		// Virtual field
		apduHeaderReduction := m.GetApduHeaderReduction()
		_ = apduHeaderReduction
		if _apduHeaderReductionErr := writeBuffer.WriteVirtual(ctx, "apduHeaderReduction", m.GetApduHeaderReduction()); _apduHeaderReductionErr != nil {
			return errors.Wrap(_apduHeaderReductionErr, "Error serializing 'apduHeaderReduction' field")
		}

		// Optional Field (serviceRequest) (Can be skipped, if the value is null)
		var serviceRequest BACnetConfirmedServiceRequest = nil
		if m.GetServiceRequest() != nil {
			if pushErr := writeBuffer.PushContext("serviceRequest"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for serviceRequest")
			}
			serviceRequest = m.GetServiceRequest()
			_serviceRequestErr := writeBuffer.WriteSerializable(ctx, serviceRequest)
			if popErr := writeBuffer.PopContext("serviceRequest"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for serviceRequest")
			}
			if _serviceRequestErr != nil {
				return errors.Wrap(_serviceRequestErr, "Error serializing 'serviceRequest' field")
			}
		}

		// Optional Field (segmentServiceChoice) (Can be skipped, if the value is null)
		var segmentServiceChoice *BACnetConfirmedServiceChoice = nil
		if m.GetSegmentServiceChoice() != nil {
			if pushErr := writeBuffer.PushContext("segmentServiceChoice"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for segmentServiceChoice")
			}
			segmentServiceChoice = m.GetSegmentServiceChoice()
			_segmentServiceChoiceErr := writeBuffer.WriteSerializable(ctx, segmentServiceChoice)
			if popErr := writeBuffer.PopContext("segmentServiceChoice"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for segmentServiceChoice")
			}
			if _segmentServiceChoiceErr != nil {
				return errors.Wrap(_segmentServiceChoiceErr, "Error serializing 'segmentServiceChoice' field")
			}
		}
		// Virtual field
		segmentReduction := m.GetSegmentReduction()
		_ = segmentReduction
		if _segmentReductionErr := writeBuffer.WriteVirtual(ctx, "segmentReduction", m.GetSegmentReduction()); _segmentReductionErr != nil {
			return errors.Wrap(_segmentReductionErr, "Error serializing 'segmentReduction' field")
		}

		// Array Field (segment)
		// Byte Array field (segment)
		if err := writeBuffer.WriteByteArray("segment", m.GetSegment()); err != nil {
			return errors.Wrap(err, "Error serializing 'segment' field")
		}

		if popErr := writeBuffer.PopContext("APDUConfirmedRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for APDUConfirmedRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_APDUConfirmedRequest) isAPDUConfirmedRequest() bool {
	return true
}

func (m *_APDUConfirmedRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
