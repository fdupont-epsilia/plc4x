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

// MonitoredSALReply is the corresponding interface of MonitoredSALReply
type MonitoredSALReply interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	EncodedReply
	// GetMonitoredSAL returns MonitoredSAL (property field)
	GetMonitoredSAL() MonitoredSAL
}

// MonitoredSALReplyExactly can be used when we want exactly this type and not a type which fulfills MonitoredSALReply.
// This is useful for switch cases.
type MonitoredSALReplyExactly interface {
	MonitoredSALReply
	isMonitoredSALReply() bool
}

// _MonitoredSALReply is the data-structure of this message
type _MonitoredSALReply struct {
	*_EncodedReply
	MonitoredSAL MonitoredSAL
}

var _ MonitoredSALReply = (*_MonitoredSALReply)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MonitoredSALReply) InitializeParent(parent EncodedReply, peekedByte byte) {
	m.PeekedByte = peekedByte
}

func (m *_MonitoredSALReply) GetParent() EncodedReply {
	return m._EncodedReply
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MonitoredSALReply) GetMonitoredSAL() MonitoredSAL {
	return m.MonitoredSAL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMonitoredSALReply factory function for _MonitoredSALReply
func NewMonitoredSALReply(monitoredSAL MonitoredSAL, peekedByte byte, cBusOptions CBusOptions, requestContext RequestContext) *_MonitoredSALReply {
	_result := &_MonitoredSALReply{
		MonitoredSAL:  monitoredSAL,
		_EncodedReply: NewEncodedReply(peekedByte, cBusOptions, requestContext),
	}
	_result._EncodedReply._EncodedReplyChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMonitoredSALReply(structType any) MonitoredSALReply {
	if casted, ok := structType.(MonitoredSALReply); ok {
		return casted
	}
	if casted, ok := structType.(*MonitoredSALReply); ok {
		return *casted
	}
	return nil
}

func (m *_MonitoredSALReply) GetTypeName() string {
	return "MonitoredSALReply"
}

func (m *_MonitoredSALReply) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (monitoredSAL)
	lengthInBits += m.MonitoredSAL.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_MonitoredSALReply) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MonitoredSALReplyParse(ctx context.Context, theBytes []byte, cBusOptions CBusOptions, requestContext RequestContext) (MonitoredSALReply, error) {
	return MonitoredSALReplyParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), cBusOptions, requestContext)
}

func MonitoredSALReplyParseWithBufferProducer(cBusOptions CBusOptions, requestContext RequestContext) func(ctx context.Context, readBuffer utils.ReadBuffer) (MonitoredSALReply, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (MonitoredSALReply, error) {
		return MonitoredSALReplyParseWithBuffer(ctx, readBuffer, cBusOptions, requestContext)
	}
}

func MonitoredSALReplyParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions, requestContext RequestContext) (MonitoredSALReply, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MonitoredSALReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MonitoredSALReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	monitoredSAL, err := ReadSimpleField[MonitoredSAL](ctx, "monitoredSAL", ReadComplex[MonitoredSAL](MonitoredSALParseWithBufferProducer[MonitoredSAL]((CBusOptions)(cBusOptions)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'monitoredSAL' field"))
	}

	if closeErr := readBuffer.CloseContext("MonitoredSALReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MonitoredSALReply")
	}

	// Create a partially initialized instance
	_child := &_MonitoredSALReply{
		_EncodedReply: &_EncodedReply{
			CBusOptions:    cBusOptions,
			RequestContext: requestContext,
		},
		MonitoredSAL: monitoredSAL,
	}
	_child._EncodedReply._EncodedReplyChildRequirements = _child
	return _child, nil
}

func (m *_MonitoredSALReply) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MonitoredSALReply) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MonitoredSALReply"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MonitoredSALReply")
		}

		if err := WriteSimpleField[MonitoredSAL](ctx, "monitoredSAL", m.GetMonitoredSAL(), WriteComplex[MonitoredSAL](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'monitoredSAL' field")
		}

		if popErr := writeBuffer.PopContext("MonitoredSALReply"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MonitoredSALReply")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MonitoredSALReply) isMonitoredSALReply() bool {
	return true
}

func (m *_MonitoredSALReply) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
