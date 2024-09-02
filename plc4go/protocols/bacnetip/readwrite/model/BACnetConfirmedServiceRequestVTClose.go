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

// BACnetConfirmedServiceRequestVTClose is the corresponding interface of BACnetConfirmedServiceRequestVTClose
type BACnetConfirmedServiceRequestVTClose interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetListOfRemoteVtSessionIdentifiers returns ListOfRemoteVtSessionIdentifiers (property field)
	GetListOfRemoteVtSessionIdentifiers() []BACnetApplicationTagUnsignedInteger
	// IsBACnetConfirmedServiceRequestVTClose is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestVTClose()
}

// _BACnetConfirmedServiceRequestVTClose is the data-structure of this message
type _BACnetConfirmedServiceRequestVTClose struct {
	BACnetConfirmedServiceRequestContract
	ListOfRemoteVtSessionIdentifiers []BACnetApplicationTagUnsignedInteger

	// Arguments.
	ServiceRequestPayloadLength uint32
}

var _ BACnetConfirmedServiceRequestVTClose = (*_BACnetConfirmedServiceRequestVTClose)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestVTClose)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestVTClose) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_VT_CLOSE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestVTClose) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestVTClose) GetListOfRemoteVtSessionIdentifiers() []BACnetApplicationTagUnsignedInteger {
	return m.ListOfRemoteVtSessionIdentifiers
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestVTClose factory function for _BACnetConfirmedServiceRequestVTClose
func NewBACnetConfirmedServiceRequestVTClose(listOfRemoteVtSessionIdentifiers []BACnetApplicationTagUnsignedInteger, serviceRequestPayloadLength uint32, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestVTClose {
	_result := &_BACnetConfirmedServiceRequestVTClose{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		ListOfRemoteVtSessionIdentifiers:      listOfRemoteVtSessionIdentifiers,
	}
	_result.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestVTClose(structType any) BACnetConfirmedServiceRequestVTClose {
	if casted, ok := structType.(BACnetConfirmedServiceRequestVTClose); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestVTClose); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestVTClose) GetTypeName() string {
	return "BACnetConfirmedServiceRequestVTClose"
}

func (m *_BACnetConfirmedServiceRequestVTClose) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Array field
	if len(m.ListOfRemoteVtSessionIdentifiers) > 0 {
		for _, element := range m.ListOfRemoteVtSessionIdentifiers {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestVTClose) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestVTClose) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestPayloadLength uint32, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestVTClose BACnetConfirmedServiceRequestVTClose, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestVTClose"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestVTClose")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	listOfRemoteVtSessionIdentifiers, err := ReadLengthArrayField[BACnetApplicationTagUnsignedInteger](ctx, "listOfRemoteVtSessionIdentifiers", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), int(serviceRequestPayloadLength))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfRemoteVtSessionIdentifiers' field"))
	}
	m.ListOfRemoteVtSessionIdentifiers = listOfRemoteVtSessionIdentifiers

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestVTClose"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestVTClose")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestVTClose) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestVTClose) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestVTClose"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestVTClose")
		}

		if err := WriteComplexTypeArrayField(ctx, "listOfRemoteVtSessionIdentifiers", m.GetListOfRemoteVtSessionIdentifiers(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'listOfRemoteVtSessionIdentifiers' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestVTClose"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestVTClose")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_BACnetConfirmedServiceRequestVTClose) GetServiceRequestPayloadLength() uint32 {
	return m.ServiceRequestPayloadLength
}

//
////

func (m *_BACnetConfirmedServiceRequestVTClose) IsBACnetConfirmedServiceRequestVTClose() {}

func (m *_BACnetConfirmedServiceRequestVTClose) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
