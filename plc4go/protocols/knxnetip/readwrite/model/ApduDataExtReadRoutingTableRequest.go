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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// ApduDataExtReadRoutingTableRequest is the corresponding interface of ApduDataExtReadRoutingTableRequest
type ApduDataExtReadRoutingTableRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ApduDataExt
}

// ApduDataExtReadRoutingTableRequestExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtReadRoutingTableRequest.
// This is useful for switch cases.
type ApduDataExtReadRoutingTableRequestExactly interface {
	ApduDataExtReadRoutingTableRequest
	isApduDataExtReadRoutingTableRequest() bool
}

// _ApduDataExtReadRoutingTableRequest is the data-structure of this message
type _ApduDataExtReadRoutingTableRequest struct {
	*_ApduDataExt
}

var _ ApduDataExtReadRoutingTableRequest = (*_ApduDataExtReadRoutingTableRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtReadRoutingTableRequest) GetExtApciType() uint8 {
	return 0x01
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtReadRoutingTableRequest) InitializeParent(parent ApduDataExt) {}

func (m *_ApduDataExtReadRoutingTableRequest) GetParent() ApduDataExt {
	return m._ApduDataExt
}

// NewApduDataExtReadRoutingTableRequest factory function for _ApduDataExtReadRoutingTableRequest
func NewApduDataExtReadRoutingTableRequest(length uint8) *_ApduDataExtReadRoutingTableRequest {
	_result := &_ApduDataExtReadRoutingTableRequest{
		_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtReadRoutingTableRequest(structType any) ApduDataExtReadRoutingTableRequest {
	if casted, ok := structType.(ApduDataExtReadRoutingTableRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtReadRoutingTableRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtReadRoutingTableRequest) GetTypeName() string {
	return "ApduDataExtReadRoutingTableRequest"
}

func (m *_ApduDataExtReadRoutingTableRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtReadRoutingTableRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ApduDataExtReadRoutingTableRequestParse(ctx context.Context, theBytes []byte, length uint8) (ApduDataExtReadRoutingTableRequest, error) {
	return ApduDataExtReadRoutingTableRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), length)
}

func ApduDataExtReadRoutingTableRequestParseWithBufferProducer(length uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (ApduDataExtReadRoutingTableRequest, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ApduDataExtReadRoutingTableRequest, error) {
		return ApduDataExtReadRoutingTableRequestParseWithBuffer(ctx, readBuffer, length)
	}
}

func ApduDataExtReadRoutingTableRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, length uint8) (ApduDataExtReadRoutingTableRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtReadRoutingTableRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtReadRoutingTableRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtReadRoutingTableRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtReadRoutingTableRequest")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtReadRoutingTableRequest{
		_ApduDataExt: &_ApduDataExt{
			Length: length,
		},
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtReadRoutingTableRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtReadRoutingTableRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtReadRoutingTableRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtReadRoutingTableRequest")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtReadRoutingTableRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtReadRoutingTableRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtReadRoutingTableRequest) isApduDataExtReadRoutingTableRequest() bool {
	return true
}

func (m *_ApduDataExtReadRoutingTableRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
