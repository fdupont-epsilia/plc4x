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

// SecurityHeader is the corresponding interface of SecurityHeader
type SecurityHeader interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetSecureChannelId returns SecureChannelId (property field)
	GetSecureChannelId() uint32
	// GetSecureTokenId returns SecureTokenId (property field)
	GetSecureTokenId() uint32
	// IsSecurityHeader is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityHeader()
}

// _SecurityHeader is the data-structure of this message
type _SecurityHeader struct {
	SecureChannelId uint32
	SecureTokenId   uint32
}

var _ SecurityHeader = (*_SecurityHeader)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityHeader) GetSecureChannelId() uint32 {
	return m.SecureChannelId
}

func (m *_SecurityHeader) GetSecureTokenId() uint32 {
	return m.SecureTokenId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSecurityHeader factory function for _SecurityHeader
func NewSecurityHeader(secureChannelId uint32, secureTokenId uint32) *_SecurityHeader {
	return &_SecurityHeader{SecureChannelId: secureChannelId, SecureTokenId: secureTokenId}
}

// Deprecated: use the interface for direct cast
func CastSecurityHeader(structType any) SecurityHeader {
	if casted, ok := structType.(SecurityHeader); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityHeader); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityHeader) GetTypeName() string {
	return "SecurityHeader"
}

func (m *_SecurityHeader) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (secureChannelId)
	lengthInBits += 32

	// Simple field (secureTokenId)
	lengthInBits += 32

	return lengthInBits
}

func (m *_SecurityHeader) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SecurityHeaderParse(ctx context.Context, theBytes []byte) (SecurityHeader, error) {
	return SecurityHeaderParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func SecurityHeaderParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityHeader, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityHeader, error) {
		return SecurityHeaderParseWithBuffer(ctx, readBuffer)
	}
}

func SecurityHeaderParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityHeader, error) {
	v, err := (&_SecurityHeader{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_SecurityHeader) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__securityHeader SecurityHeader, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityHeader")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	secureChannelId, err := ReadSimpleField(ctx, "secureChannelId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'secureChannelId' field"))
	}
	m.SecureChannelId = secureChannelId

	secureTokenId, err := ReadSimpleField(ctx, "secureTokenId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'secureTokenId' field"))
	}
	m.SecureTokenId = secureTokenId

	if closeErr := readBuffer.CloseContext("SecurityHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityHeader")
	}

	return m, nil
}

func (m *_SecurityHeader) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityHeader) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("SecurityHeader"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for SecurityHeader")
	}

	if err := WriteSimpleField[uint32](ctx, "secureChannelId", m.GetSecureChannelId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
		return errors.Wrap(err, "Error serializing 'secureChannelId' field")
	}

	if err := WriteSimpleField[uint32](ctx, "secureTokenId", m.GetSecureTokenId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
		return errors.Wrap(err, "Error serializing 'secureTokenId' field")
	}

	if popErr := writeBuffer.PopContext("SecurityHeader"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for SecurityHeader")
	}
	return nil
}

func (m *_SecurityHeader) IsSecurityHeader() {}

func (m *_SecurityHeader) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
