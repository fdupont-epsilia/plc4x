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

// SysexCommandReportFirmwareResponse is the corresponding interface of SysexCommandReportFirmwareResponse
type SysexCommandReportFirmwareResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SysexCommand
	// GetMajorVersion returns MajorVersion (property field)
	GetMajorVersion() uint8
	// GetMinorVersion returns MinorVersion (property field)
	GetMinorVersion() uint8
	// GetFileName returns FileName (property field)
	GetFileName() []byte
}

// SysexCommandReportFirmwareResponseExactly can be used when we want exactly this type and not a type which fulfills SysexCommandReportFirmwareResponse.
// This is useful for switch cases.
type SysexCommandReportFirmwareResponseExactly interface {
	SysexCommandReportFirmwareResponse
	isSysexCommandReportFirmwareResponse() bool
}

// _SysexCommandReportFirmwareResponse is the data-structure of this message
type _SysexCommandReportFirmwareResponse struct {
	*_SysexCommand
	MajorVersion uint8
	MinorVersion uint8
	FileName     []byte
}

var _ SysexCommandReportFirmwareResponse = (*_SysexCommandReportFirmwareResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SysexCommandReportFirmwareResponse) GetCommandType() uint8 {
	return 0x79
}

func (m *_SysexCommandReportFirmwareResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SysexCommandReportFirmwareResponse) InitializeParent(parent SysexCommand) {}

func (m *_SysexCommandReportFirmwareResponse) GetParent() SysexCommand {
	return m._SysexCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SysexCommandReportFirmwareResponse) GetMajorVersion() uint8 {
	return m.MajorVersion
}

func (m *_SysexCommandReportFirmwareResponse) GetMinorVersion() uint8 {
	return m.MinorVersion
}

func (m *_SysexCommandReportFirmwareResponse) GetFileName() []byte {
	return m.FileName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSysexCommandReportFirmwareResponse factory function for _SysexCommandReportFirmwareResponse
func NewSysexCommandReportFirmwareResponse(majorVersion uint8, minorVersion uint8, fileName []byte) *_SysexCommandReportFirmwareResponse {
	_result := &_SysexCommandReportFirmwareResponse{
		MajorVersion:  majorVersion,
		MinorVersion:  minorVersion,
		FileName:      fileName,
		_SysexCommand: NewSysexCommand(),
	}
	_result._SysexCommand._SysexCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSysexCommandReportFirmwareResponse(structType any) SysexCommandReportFirmwareResponse {
	if casted, ok := structType.(SysexCommandReportFirmwareResponse); ok {
		return casted
	}
	if casted, ok := structType.(*SysexCommandReportFirmwareResponse); ok {
		return *casted
	}
	return nil
}

func (m *_SysexCommandReportFirmwareResponse) GetTypeName() string {
	return "SysexCommandReportFirmwareResponse"
}

func (m *_SysexCommandReportFirmwareResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (majorVersion)
	lengthInBits += 8

	// Simple field (minorVersion)
	lengthInBits += 8

	// Manual Array Field (fileName)
	lengthInBits += uint16(LengthSysexString(ctx, m.GetFileName()))

	return lengthInBits
}

func (m *_SysexCommandReportFirmwareResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SysexCommandReportFirmwareResponseParse(ctx context.Context, theBytes []byte, response bool) (SysexCommandReportFirmwareResponse, error) {
	return SysexCommandReportFirmwareResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func SysexCommandReportFirmwareResponseParseWithBufferProducer(response bool) func(ctx context.Context, readBuffer utils.ReadBuffer) (SysexCommandReportFirmwareResponse, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (SysexCommandReportFirmwareResponse, error) {
		return SysexCommandReportFirmwareResponseParseWithBuffer(ctx, readBuffer, response)
	}
}

func SysexCommandReportFirmwareResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (SysexCommandReportFirmwareResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommandReportFirmwareResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandReportFirmwareResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	majorVersion, err := ReadSimpleField(ctx, "majorVersion", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'majorVersion' field"))
	}

	minorVersion, err := ReadSimpleField(ctx, "minorVersion", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'minorVersion' field"))
	}

	fileName, err := ReadManualByteArrayField(ctx, "fileName", readBuffer, IsSysexEnd(ctx, readBuffer), ParseSysexString(ctx, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fileName' field"))
	}

	if closeErr := readBuffer.CloseContext("SysexCommandReportFirmwareResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandReportFirmwareResponse")
	}

	// Create a partially initialized instance
	_child := &_SysexCommandReportFirmwareResponse{
		_SysexCommand: &_SysexCommand{},
		MajorVersion:  majorVersion,
		MinorVersion:  minorVersion,
		FileName:      fileName,
	}
	_child._SysexCommand._SysexCommandChildRequirements = _child
	return _child, nil
}

func (m *_SysexCommandReportFirmwareResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SysexCommandReportFirmwareResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandReportFirmwareResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandReportFirmwareResponse")
		}

		if err := WriteSimpleField[uint8](ctx, "majorVersion", m.GetMajorVersion(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'majorVersion' field")
		}

		if err := WriteSimpleField[uint8](ctx, "minorVersion", m.GetMinorVersion(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'minorVersion' field")
		}

		if err := WriteManualArrayField[byte](ctx, "fileName", m.GetFileName(), func(ctx context.Context, writeBuffer utils.WriteBuffer, m byte) error {
			return SerializeSysexString(ctx, writeBuffer, m)
		}, writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'fileName' field")
		}

		if popErr := writeBuffer.PopContext("SysexCommandReportFirmwareResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandReportFirmwareResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SysexCommandReportFirmwareResponse) isSysexCommandReportFirmwareResponse() bool {
	return true
}

func (m *_SysexCommandReportFirmwareResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
