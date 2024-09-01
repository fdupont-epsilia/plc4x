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

// ModbusPDUMaskWriteHoldingRegisterRequest is the corresponding interface of ModbusPDUMaskWriteHoldingRegisterRequest
type ModbusPDUMaskWriteHoldingRegisterRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetReferenceAddress returns ReferenceAddress (property field)
	GetReferenceAddress() uint16
	// GetAndMask returns AndMask (property field)
	GetAndMask() uint16
	// GetOrMask returns OrMask (property field)
	GetOrMask() uint16
}

// ModbusPDUMaskWriteHoldingRegisterRequestExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUMaskWriteHoldingRegisterRequest.
// This is useful for switch cases.
type ModbusPDUMaskWriteHoldingRegisterRequestExactly interface {
	ModbusPDUMaskWriteHoldingRegisterRequest
	isModbusPDUMaskWriteHoldingRegisterRequest() bool
}

// _ModbusPDUMaskWriteHoldingRegisterRequest is the data-structure of this message
type _ModbusPDUMaskWriteHoldingRegisterRequest struct {
	ModbusPDUContract
	ReferenceAddress uint16
	AndMask          uint16
	OrMask           uint16
}

var _ ModbusPDUMaskWriteHoldingRegisterRequest = (*_ModbusPDUMaskWriteHoldingRegisterRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUMaskWriteHoldingRegisterRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUMaskWriteHoldingRegisterRequest) GetFunctionFlag() uint8 {
	return 0x16
}

func (m *_ModbusPDUMaskWriteHoldingRegisterRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUMaskWriteHoldingRegisterRequest) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUMaskWriteHoldingRegisterRequest) GetReferenceAddress() uint16 {
	return m.ReferenceAddress
}

func (m *_ModbusPDUMaskWriteHoldingRegisterRequest) GetAndMask() uint16 {
	return m.AndMask
}

func (m *_ModbusPDUMaskWriteHoldingRegisterRequest) GetOrMask() uint16 {
	return m.OrMask
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUMaskWriteHoldingRegisterRequest factory function for _ModbusPDUMaskWriteHoldingRegisterRequest
func NewModbusPDUMaskWriteHoldingRegisterRequest(referenceAddress uint16, andMask uint16, orMask uint16) *_ModbusPDUMaskWriteHoldingRegisterRequest {
	_result := &_ModbusPDUMaskWriteHoldingRegisterRequest{
		ModbusPDUContract: NewModbusPDU(),
		ReferenceAddress:  referenceAddress,
		AndMask:           andMask,
		OrMask:            orMask,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUMaskWriteHoldingRegisterRequest(structType any) ModbusPDUMaskWriteHoldingRegisterRequest {
	if casted, ok := structType.(ModbusPDUMaskWriteHoldingRegisterRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUMaskWriteHoldingRegisterRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUMaskWriteHoldingRegisterRequest) GetTypeName() string {
	return "ModbusPDUMaskWriteHoldingRegisterRequest"
}

func (m *_ModbusPDUMaskWriteHoldingRegisterRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	// Simple field (referenceAddress)
	lengthInBits += 16

	// Simple field (andMask)
	lengthInBits += 16

	// Simple field (orMask)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ModbusPDUMaskWriteHoldingRegisterRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUMaskWriteHoldingRegisterRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (_modbusPDUMaskWriteHoldingRegisterRequest ModbusPDUMaskWriteHoldingRegisterRequest, err error) {
	m.ModbusPDUContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUMaskWriteHoldingRegisterRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUMaskWriteHoldingRegisterRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	referenceAddress, err := ReadSimpleField(ctx, "referenceAddress", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'referenceAddress' field"))
	}
	m.ReferenceAddress = referenceAddress

	andMask, err := ReadSimpleField(ctx, "andMask", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'andMask' field"))
	}
	m.AndMask = andMask

	orMask, err := ReadSimpleField(ctx, "orMask", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'orMask' field"))
	}
	m.OrMask = orMask

	if closeErr := readBuffer.CloseContext("ModbusPDUMaskWriteHoldingRegisterRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUMaskWriteHoldingRegisterRequest")
	}

	return m, nil
}

func (m *_ModbusPDUMaskWriteHoldingRegisterRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUMaskWriteHoldingRegisterRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUMaskWriteHoldingRegisterRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUMaskWriteHoldingRegisterRequest")
		}

		if err := WriteSimpleField[uint16](ctx, "referenceAddress", m.GetReferenceAddress(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'referenceAddress' field")
		}

		if err := WriteSimpleField[uint16](ctx, "andMask", m.GetAndMask(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'andMask' field")
		}

		if err := WriteSimpleField[uint16](ctx, "orMask", m.GetOrMask(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'orMask' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUMaskWriteHoldingRegisterRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUMaskWriteHoldingRegisterRequest")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUMaskWriteHoldingRegisterRequest) isModbusPDUMaskWriteHoldingRegisterRequest() bool {
	return true
}

func (m *_ModbusPDUMaskWriteHoldingRegisterRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
