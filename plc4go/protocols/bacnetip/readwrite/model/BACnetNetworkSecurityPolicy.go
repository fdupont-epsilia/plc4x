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

// BACnetNetworkSecurityPolicy is the corresponding interface of BACnetNetworkSecurityPolicy
type BACnetNetworkSecurityPolicy interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetPortId returns PortId (property field)
	GetPortId() BACnetContextTagUnsignedInteger
	// GetSecurityLevel returns SecurityLevel (property field)
	GetSecurityLevel() BACnetSecurityPolicyTagged
}

// BACnetNetworkSecurityPolicyExactly can be used when we want exactly this type and not a type which fulfills BACnetNetworkSecurityPolicy.
// This is useful for switch cases.
type BACnetNetworkSecurityPolicyExactly interface {
	BACnetNetworkSecurityPolicy
	isBACnetNetworkSecurityPolicy() bool
}

// _BACnetNetworkSecurityPolicy is the data-structure of this message
type _BACnetNetworkSecurityPolicy struct {
	PortId        BACnetContextTagUnsignedInteger
	SecurityLevel BACnetSecurityPolicyTagged
}

var _ BACnetNetworkSecurityPolicy = (*_BACnetNetworkSecurityPolicy)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNetworkSecurityPolicy) GetPortId() BACnetContextTagUnsignedInteger {
	return m.PortId
}

func (m *_BACnetNetworkSecurityPolicy) GetSecurityLevel() BACnetSecurityPolicyTagged {
	return m.SecurityLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNetworkSecurityPolicy factory function for _BACnetNetworkSecurityPolicy
func NewBACnetNetworkSecurityPolicy(portId BACnetContextTagUnsignedInteger, securityLevel BACnetSecurityPolicyTagged) *_BACnetNetworkSecurityPolicy {
	return &_BACnetNetworkSecurityPolicy{PortId: portId, SecurityLevel: securityLevel}
}

// Deprecated: use the interface for direct cast
func CastBACnetNetworkSecurityPolicy(structType any) BACnetNetworkSecurityPolicy {
	if casted, ok := structType.(BACnetNetworkSecurityPolicy); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNetworkSecurityPolicy); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNetworkSecurityPolicy) GetTypeName() string {
	return "BACnetNetworkSecurityPolicy"
}

func (m *_BACnetNetworkSecurityPolicy) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (portId)
	lengthInBits += m.PortId.GetLengthInBits(ctx)

	// Simple field (securityLevel)
	lengthInBits += m.SecurityLevel.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNetworkSecurityPolicy) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetNetworkSecurityPolicyParse(ctx context.Context, theBytes []byte) (BACnetNetworkSecurityPolicy, error) {
	return BACnetNetworkSecurityPolicyParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetNetworkSecurityPolicyParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetNetworkSecurityPolicy, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetNetworkSecurityPolicy, error) {
		return BACnetNetworkSecurityPolicyParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetNetworkSecurityPolicyParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetNetworkSecurityPolicy, error) {
	v, err := (&_BACnetNetworkSecurityPolicy{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetNetworkSecurityPolicy) parse(ctx context.Context, readBuffer utils.ReadBuffer) (_bACnetNetworkSecurityPolicy BACnetNetworkSecurityPolicy, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNetworkSecurityPolicy"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNetworkSecurityPolicy")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	portId, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "portId", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'portId' field"))
	}
	m.PortId = portId

	securityLevel, err := ReadSimpleField[BACnetSecurityPolicyTagged](ctx, "securityLevel", ReadComplex[BACnetSecurityPolicyTagged](BACnetSecurityPolicyTaggedParseWithBufferProducer((uint8)(uint8(1)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityLevel' field"))
	}
	m.SecurityLevel = securityLevel

	if closeErr := readBuffer.CloseContext("BACnetNetworkSecurityPolicy"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNetworkSecurityPolicy")
	}

	return m, nil
}

func (m *_BACnetNetworkSecurityPolicy) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNetworkSecurityPolicy) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetNetworkSecurityPolicy"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetNetworkSecurityPolicy")
	}

	if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "portId", m.GetPortId(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'portId' field")
	}

	if err := WriteSimpleField[BACnetSecurityPolicyTagged](ctx, "securityLevel", m.GetSecurityLevel(), WriteComplex[BACnetSecurityPolicyTagged](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'securityLevel' field")
	}

	if popErr := writeBuffer.PopContext("BACnetNetworkSecurityPolicy"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetNetworkSecurityPolicy")
	}
	return nil
}

func (m *_BACnetNetworkSecurityPolicy) isBACnetNetworkSecurityPolicy() bool {
	return true
}

func (m *_BACnetNetworkSecurityPolicy) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
