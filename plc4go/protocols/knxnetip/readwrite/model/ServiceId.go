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

// ServiceId is the corresponding interface of ServiceId
type ServiceId interface {
	ServiceIdContract
	ServiceIdRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// ServiceIdContract provides a set of functions which can be overwritten by a sub struct
type ServiceIdContract interface {
}

// ServiceIdRequirements provides a set of functions which need to be implemented by a sub struct
type ServiceIdRequirements interface {
	// GetServiceType returns ServiceType (discriminator field)
	GetServiceType() uint8
}

// ServiceIdExactly can be used when we want exactly this type and not a type which fulfills ServiceId.
// This is useful for switch cases.
type ServiceIdExactly interface {
	ServiceId
	isServiceId() bool
}

// _ServiceId is the data-structure of this message
type _ServiceId struct {
	_ServiceIdChildRequirements
}

var _ ServiceIdContract = (*_ServiceId)(nil)

type _ServiceIdChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetServiceType() uint8
}

type ServiceIdChild interface {
	utils.Serializable

	GetParent() *ServiceId

	GetTypeName() string
	ServiceId
}

// NewServiceId factory function for _ServiceId
func NewServiceId() *_ServiceId {
	return &_ServiceId{}
}

// Deprecated: use the interface for direct cast
func CastServiceId(structType any) ServiceId {
	if casted, ok := structType.(ServiceId); ok {
		return casted
	}
	if casted, ok := structType.(*ServiceId); ok {
		return *casted
	}
	return nil
}

func (m *_ServiceId) GetTypeName() string {
	return "ServiceId"
}

func (m *_ServiceId) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (serviceType)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ServiceId) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ServiceIdParse[T ServiceId](ctx context.Context, theBytes []byte) (T, error) {
	return ServiceIdParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func ServiceIdParseWithBufferProducer[T ServiceId]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := ServiceIdParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func ServiceIdParseWithBuffer[T ServiceId](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_ServiceId{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_ServiceId) parse(ctx context.Context, readBuffer utils.ReadBuffer) (_serviceId ServiceId, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ServiceId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ServiceId")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	serviceType, err := ReadDiscriminatorField[uint8](ctx, "serviceType", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serviceType' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child ServiceId
	switch {
	case serviceType == 0x02: // KnxNetIpCore
		if _child, err = (&_KnxNetIpCore{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type KnxNetIpCore for type-switch of ServiceId")
		}
	case serviceType == 0x03: // KnxNetIpDeviceManagement
		if _child, err = (&_KnxNetIpDeviceManagement{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type KnxNetIpDeviceManagement for type-switch of ServiceId")
		}
	case serviceType == 0x04: // KnxNetIpTunneling
		if _child, err = (&_KnxNetIpTunneling{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type KnxNetIpTunneling for type-switch of ServiceId")
		}
	case serviceType == 0x05: // KnxNetIpRouting
		if _child, err = (&_KnxNetIpRouting{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type KnxNetIpRouting for type-switch of ServiceId")
		}
	case serviceType == 0x06: // KnxNetRemoteLogging
		if _child, err = (&_KnxNetRemoteLogging{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type KnxNetRemoteLogging for type-switch of ServiceId")
		}
	case serviceType == 0x07: // KnxNetRemoteConfigurationAndDiagnosis
		if _child, err = (&_KnxNetRemoteConfigurationAndDiagnosis{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type KnxNetRemoteConfigurationAndDiagnosis for type-switch of ServiceId")
		}
	case serviceType == 0x08: // KnxNetObjectServer
		if _child, err = (&_KnxNetObjectServer{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type KnxNetObjectServer for type-switch of ServiceId")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [serviceType=%v]", serviceType)
	}

	if closeErr := readBuffer.CloseContext("ServiceId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ServiceId")
	}

	return _child, nil
}

func (pm *_ServiceId) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ServiceId, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ServiceId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ServiceId")
	}

	if err := WriteDiscriminatorField(ctx, "serviceType", m.GetServiceType(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'serviceType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ServiceId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ServiceId")
	}
	return nil
}

func (m *_ServiceId) isServiceId() bool {
	return true
}
