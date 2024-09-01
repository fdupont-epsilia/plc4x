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

// BACnetConfirmedServiceRequestWritePropertyMultiple is the corresponding interface of BACnetConfirmedServiceRequestWritePropertyMultiple
type BACnetConfirmedServiceRequestWritePropertyMultiple interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetData returns Data (property field)
	GetData() []BACnetWriteAccessSpecification
}

// BACnetConfirmedServiceRequestWritePropertyMultipleExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestWritePropertyMultiple.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestWritePropertyMultipleExactly interface {
	BACnetConfirmedServiceRequestWritePropertyMultiple
	isBACnetConfirmedServiceRequestWritePropertyMultiple() bool
}

// _BACnetConfirmedServiceRequestWritePropertyMultiple is the data-structure of this message
type _BACnetConfirmedServiceRequestWritePropertyMultiple struct {
	*_BACnetConfirmedServiceRequest
	Data []BACnetWriteAccessSpecification

	// Arguments.
	ServiceRequestPayloadLength uint32
}

var _ BACnetConfirmedServiceRequestWritePropertyMultiple = (*_BACnetConfirmedServiceRequestWritePropertyMultiple)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_WRITE_PROPERTY_MULTIPLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) InitializeParent(parent BACnetConfirmedServiceRequest) {
}

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) GetData() []BACnetWriteAccessSpecification {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestWritePropertyMultiple factory function for _BACnetConfirmedServiceRequestWritePropertyMultiple
func NewBACnetConfirmedServiceRequestWritePropertyMultiple(data []BACnetWriteAccessSpecification, serviceRequestPayloadLength uint32, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestWritePropertyMultiple {
	_result := &_BACnetConfirmedServiceRequestWritePropertyMultiple{
		Data:                           data,
		_BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestWritePropertyMultiple(structType any) BACnetConfirmedServiceRequestWritePropertyMultiple {
	if casted, ok := structType.(BACnetConfirmedServiceRequestWritePropertyMultiple); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestWritePropertyMultiple); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) GetTypeName() string {
	return "BACnetConfirmedServiceRequestWritePropertyMultiple"
}

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.Data) > 0 {
		for _, element := range m.Data {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestWritePropertyMultipleParse(ctx context.Context, theBytes []byte, serviceRequestPayloadLength uint32, serviceRequestLength uint32) (BACnetConfirmedServiceRequestWritePropertyMultiple, error) {
	return BACnetConfirmedServiceRequestWritePropertyMultipleParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), serviceRequestPayloadLength, serviceRequestLength)
}

func BACnetConfirmedServiceRequestWritePropertyMultipleParseWithBufferProducer(serviceRequestPayloadLength uint32, serviceRequestLength uint32) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestWritePropertyMultiple, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestWritePropertyMultiple, error) {
		return BACnetConfirmedServiceRequestWritePropertyMultipleParseWithBuffer(ctx, readBuffer, serviceRequestPayloadLength, serviceRequestLength)
	}
}

func BACnetConfirmedServiceRequestWritePropertyMultipleParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceRequestPayloadLength uint32, serviceRequestLength uint32) (BACnetConfirmedServiceRequestWritePropertyMultiple, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestWritePropertyMultiple"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestWritePropertyMultiple")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	data, err := ReadLengthArrayField[BACnetWriteAccessSpecification](ctx, "data", ReadComplex[BACnetWriteAccessSpecification](BACnetWriteAccessSpecificationParseWithBuffer, readBuffer), int(serviceRequestPayloadLength))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestWritePropertyMultiple"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestWritePropertyMultiple")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestWritePropertyMultiple{
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		Data: data,
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestWritePropertyMultiple"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestWritePropertyMultiple")
		}

		if err := WriteComplexTypeArrayField(ctx, "data", m.GetData(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestWritePropertyMultiple"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestWritePropertyMultiple")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) GetServiceRequestPayloadLength() uint32 {
	return m.ServiceRequestPayloadLength
}

//
////

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) isBACnetConfirmedServiceRequestWritePropertyMultiple() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
