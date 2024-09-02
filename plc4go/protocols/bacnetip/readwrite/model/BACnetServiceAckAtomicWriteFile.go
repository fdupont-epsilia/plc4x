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

// BACnetServiceAckAtomicWriteFile is the corresponding interface of BACnetServiceAckAtomicWriteFile
type BACnetServiceAckAtomicWriteFile interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetServiceAck
	// GetFileStartPosition returns FileStartPosition (property field)
	GetFileStartPosition() BACnetContextTagSignedInteger
	// IsBACnetServiceAckAtomicWriteFile is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetServiceAckAtomicWriteFile()
}

// _BACnetServiceAckAtomicWriteFile is the data-structure of this message
type _BACnetServiceAckAtomicWriteFile struct {
	BACnetServiceAckContract
	FileStartPosition BACnetContextTagSignedInteger
}

var _ BACnetServiceAckAtomicWriteFile = (*_BACnetServiceAckAtomicWriteFile)(nil)
var _ BACnetServiceAckRequirements = (*_BACnetServiceAckAtomicWriteFile)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetServiceAckAtomicWriteFile) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_ATOMIC_WRITE_FILE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckAtomicWriteFile) GetParent() BACnetServiceAckContract {
	return m.BACnetServiceAckContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckAtomicWriteFile) GetFileStartPosition() BACnetContextTagSignedInteger {
	return m.FileStartPosition
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckAtomicWriteFile factory function for _BACnetServiceAckAtomicWriteFile
func NewBACnetServiceAckAtomicWriteFile(fileStartPosition BACnetContextTagSignedInteger, serviceAckLength uint32) *_BACnetServiceAckAtomicWriteFile {
	_result := &_BACnetServiceAckAtomicWriteFile{
		BACnetServiceAckContract: NewBACnetServiceAck(serviceAckLength),
		FileStartPosition:        fileStartPosition,
	}
	_result.BACnetServiceAckContract.(*_BACnetServiceAck)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckAtomicWriteFile(structType any) BACnetServiceAckAtomicWriteFile {
	if casted, ok := structType.(BACnetServiceAckAtomicWriteFile); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckAtomicWriteFile); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckAtomicWriteFile) GetTypeName() string {
	return "BACnetServiceAckAtomicWriteFile"
}

func (m *_BACnetServiceAckAtomicWriteFile) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetServiceAckContract.(*_BACnetServiceAck).getLengthInBits(ctx))

	// Simple field (fileStartPosition)
	lengthInBits += m.FileStartPosition.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetServiceAckAtomicWriteFile) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetServiceAckAtomicWriteFile) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetServiceAck, serviceAckLength uint32) (__bACnetServiceAckAtomicWriteFile BACnetServiceAckAtomicWriteFile, err error) {
	m.BACnetServiceAckContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckAtomicWriteFile"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckAtomicWriteFile")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	fileStartPosition, err := ReadSimpleField[BACnetContextTagSignedInteger](ctx, "fileStartPosition", ReadComplex[BACnetContextTagSignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagSignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_SIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fileStartPosition' field"))
	}
	m.FileStartPosition = fileStartPosition

	if closeErr := readBuffer.CloseContext("BACnetServiceAckAtomicWriteFile"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckAtomicWriteFile")
	}

	return m, nil
}

func (m *_BACnetServiceAckAtomicWriteFile) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetServiceAckAtomicWriteFile) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckAtomicWriteFile"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckAtomicWriteFile")
		}

		if err := WriteSimpleField[BACnetContextTagSignedInteger](ctx, "fileStartPosition", m.GetFileStartPosition(), WriteComplex[BACnetContextTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'fileStartPosition' field")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckAtomicWriteFile"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckAtomicWriteFile")
		}
		return nil
	}
	return m.BACnetServiceAckContract.(*_BACnetServiceAck).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetServiceAckAtomicWriteFile) IsBACnetServiceAckAtomicWriteFile() {}

func (m *_BACnetServiceAckAtomicWriteFile) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
