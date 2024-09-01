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

// Constant values.
const AdsDiscoveryConstants_ADSDISCOVERYUDPDEFAULTPORT uint16 = uint16(48899)

// AdsDiscoveryConstants is the corresponding interface of AdsDiscoveryConstants
type AdsDiscoveryConstants interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// AdsDiscoveryConstantsExactly can be used when we want exactly this type and not a type which fulfills AdsDiscoveryConstants.
// This is useful for switch cases.
type AdsDiscoveryConstantsExactly interface {
	AdsDiscoveryConstants
	isAdsDiscoveryConstants() bool
}

// _AdsDiscoveryConstants is the data-structure of this message
type _AdsDiscoveryConstants struct {
}

var _ AdsDiscoveryConstants = (*_AdsDiscoveryConstants)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_AdsDiscoveryConstants) GetAdsDiscoveryUdpDefaultPort() uint16 {
	return AdsDiscoveryConstants_ADSDISCOVERYUDPDEFAULTPORT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsDiscoveryConstants factory function for _AdsDiscoveryConstants
func NewAdsDiscoveryConstants() *_AdsDiscoveryConstants {
	return &_AdsDiscoveryConstants{}
}

// Deprecated: use the interface for direct cast
func CastAdsDiscoveryConstants(structType any) AdsDiscoveryConstants {
	if casted, ok := structType.(AdsDiscoveryConstants); ok {
		return casted
	}
	if casted, ok := structType.(*AdsDiscoveryConstants); ok {
		return *casted
	}
	return nil
}

func (m *_AdsDiscoveryConstants) GetTypeName() string {
	return "AdsDiscoveryConstants"
}

func (m *_AdsDiscoveryConstants) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Const Field (adsDiscoveryUdpDefaultPort)
	lengthInBits += 16

	return lengthInBits
}

func (m *_AdsDiscoveryConstants) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsDiscoveryConstantsParse(ctx context.Context, theBytes []byte) (AdsDiscoveryConstants, error) {
	return AdsDiscoveryConstantsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AdsDiscoveryConstantsParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (AdsDiscoveryConstants, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (AdsDiscoveryConstants, error) {
		return AdsDiscoveryConstantsParseWithBuffer(ctx, readBuffer)
	}
}

func AdsDiscoveryConstantsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AdsDiscoveryConstants, error) {
	v, err := (&_AdsDiscoveryConstants{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_AdsDiscoveryConstants) parse(ctx context.Context, readBuffer utils.ReadBuffer) (_adsDiscoveryConstants AdsDiscoveryConstants, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsDiscoveryConstants"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsDiscoveryConstants")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	adsDiscoveryUdpDefaultPort, err := ReadConstField[uint16](ctx, "adsDiscoveryUdpDefaultPort", ReadUnsignedShort(readBuffer, uint8(16)), AdsDiscoveryConstants_ADSDISCOVERYUDPDEFAULTPORT)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'adsDiscoveryUdpDefaultPort' field"))
	}
	_ = adsDiscoveryUdpDefaultPort

	if closeErr := readBuffer.CloseContext("AdsDiscoveryConstants"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsDiscoveryConstants")
	}

	return m, nil
}

func (m *_AdsDiscoveryConstants) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsDiscoveryConstants) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AdsDiscoveryConstants"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AdsDiscoveryConstants")
	}

	if err := WriteConstField(ctx, "adsDiscoveryUdpDefaultPort", AdsDiscoveryConstants_ADSDISCOVERYUDPDEFAULTPORT, WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'adsDiscoveryUdpDefaultPort' field")
	}

	if popErr := writeBuffer.PopContext("AdsDiscoveryConstants"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AdsDiscoveryConstants")
	}
	return nil
}

func (m *_AdsDiscoveryConstants) isAdsDiscoveryConstants() bool {
	return true
}

func (m *_AdsDiscoveryConstants) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
