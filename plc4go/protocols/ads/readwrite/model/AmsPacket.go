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
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const AmsPacket_INITCOMMAND bool = bool(false)
const AmsPacket_UPDCOMMAND bool = bool(false)
const AmsPacket_TIMESTAMPADDED bool = bool(false)
const AmsPacket_HIGHPRIORITYCOMMAND bool = bool(false)
const AmsPacket_SYSTEMCOMMAND bool = bool(false)
const AmsPacket_ADSCOMMAND bool = bool(true)
const AmsPacket_NORETURN bool = bool(false)
const AmsPacket_BROADCAST bool = bool(false)

// AmsPacket is the corresponding interface of AmsPacket
type AmsPacket interface {
	utils.LengthAware
	utils.Serializable
	// GetCommandId returns CommandId (discriminator field)
	GetCommandId() CommandId
	// GetResponse returns Response (discriminator field)
	GetResponse() bool
	// GetTargetAmsNetId returns TargetAmsNetId (property field)
	GetTargetAmsNetId() AmsNetId
	// GetTargetAmsPort returns TargetAmsPort (property field)
	GetTargetAmsPort() uint16
	// GetSourceAmsNetId returns SourceAmsNetId (property field)
	GetSourceAmsNetId() AmsNetId
	// GetSourceAmsPort returns SourceAmsPort (property field)
	GetSourceAmsPort() uint16
	// GetErrorCode returns ErrorCode (property field)
	GetErrorCode() uint32
	// GetInvokeId returns InvokeId (property field)
	GetInvokeId() uint32
}

// AmsPacketExactly can be used when we want exactly this type and not a type which fulfills AmsPacket.
// This is useful for switch cases.
type AmsPacketExactly interface {
	AmsPacket
	isAmsPacket() bool
}

// _AmsPacket is the data-structure of this message
type _AmsPacket struct {
	_AmsPacketChildRequirements
	TargetAmsNetId AmsNetId
	TargetAmsPort  uint16
	SourceAmsNetId AmsNetId
	SourceAmsPort  uint16
	ErrorCode      uint32
	InvokeId       uint32
	// Reserved Fields
	reservedField0 *int8
}

type _AmsPacketChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
	GetCommandId() CommandId
	GetResponse() bool
}

type AmsPacketParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child AmsPacket, serializeChildFunction func() error) error
	GetTypeName() string
}

type AmsPacketChild interface {
	utils.Serializable
	InitializeParent(parent AmsPacket, targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32)
	GetParent() *AmsPacket

	GetTypeName() string
	AmsPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AmsPacket) GetTargetAmsNetId() AmsNetId {
	return m.TargetAmsNetId
}

func (m *_AmsPacket) GetTargetAmsPort() uint16 {
	return m.TargetAmsPort
}

func (m *_AmsPacket) GetSourceAmsNetId() AmsNetId {
	return m.SourceAmsNetId
}

func (m *_AmsPacket) GetSourceAmsPort() uint16 {
	return m.SourceAmsPort
}

func (m *_AmsPacket) GetErrorCode() uint32 {
	return m.ErrorCode
}

func (m *_AmsPacket) GetInvokeId() uint32 {
	return m.InvokeId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_AmsPacket) GetInitCommand() bool {
	return AmsPacket_INITCOMMAND
}

func (m *_AmsPacket) GetUpdCommand() bool {
	return AmsPacket_UPDCOMMAND
}

func (m *_AmsPacket) GetTimestampAdded() bool {
	return AmsPacket_TIMESTAMPADDED
}

func (m *_AmsPacket) GetHighPriorityCommand() bool {
	return AmsPacket_HIGHPRIORITYCOMMAND
}

func (m *_AmsPacket) GetSystemCommand() bool {
	return AmsPacket_SYSTEMCOMMAND
}

func (m *_AmsPacket) GetAdsCommand() bool {
	return AmsPacket_ADSCOMMAND
}

func (m *_AmsPacket) GetNoReturn() bool {
	return AmsPacket_NORETURN
}

func (m *_AmsPacket) GetBroadcast() bool {
	return AmsPacket_BROADCAST
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAmsPacket factory function for _AmsPacket
func NewAmsPacket(targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32) *_AmsPacket {
	return &_AmsPacket{TargetAmsNetId: targetAmsNetId, TargetAmsPort: targetAmsPort, SourceAmsNetId: sourceAmsNetId, SourceAmsPort: sourceAmsPort, ErrorCode: errorCode, InvokeId: invokeId}
}

// Deprecated: use the interface for direct cast
func CastAmsPacket(structType interface{}) AmsPacket {
	if casted, ok := structType.(AmsPacket); ok {
		return casted
	}
	if casted, ok := structType.(*AmsPacket); ok {
		return *casted
	}
	return nil
}

func (m *_AmsPacket) GetTypeName() string {
	return "AmsPacket"
}

func (m *_AmsPacket) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (targetAmsNetId)
	lengthInBits += m.TargetAmsNetId.GetLengthInBits()

	// Simple field (targetAmsPort)
	lengthInBits += 16

	// Simple field (sourceAmsNetId)
	lengthInBits += m.SourceAmsNetId.GetLengthInBits()

	// Simple field (sourceAmsPort)
	lengthInBits += 16
	// Discriminator Field (commandId)
	lengthInBits += 16

	// Const Field (initCommand)
	lengthInBits += 1

	// Const Field (updCommand)
	lengthInBits += 1

	// Const Field (timestampAdded)
	lengthInBits += 1

	// Const Field (highPriorityCommand)
	lengthInBits += 1

	// Const Field (systemCommand)
	lengthInBits += 1

	// Const Field (adsCommand)
	lengthInBits += 1

	// Const Field (noReturn)
	lengthInBits += 1
	// Discriminator Field (response)
	lengthInBits += 1

	// Const Field (broadcast)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 7

	// Implicit Field (length)
	lengthInBits += 32

	// Simple field (errorCode)
	lengthInBits += 32

	// Simple field (invokeId)
	lengthInBits += 32

	return lengthInBits
}

func (m *_AmsPacket) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AmsPacketParse(theBytes []byte) (AmsPacket, error) {
	return AmsPacketParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func AmsPacketParseWithBuffer(readBuffer utils.ReadBuffer) (AmsPacket, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AmsPacket"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AmsPacket")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (targetAmsNetId)
	if pullErr := readBuffer.PullContext("targetAmsNetId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for targetAmsNetId")
	}
	_targetAmsNetId, _targetAmsNetIdErr := AmsNetIdParseWithBuffer(readBuffer)
	if _targetAmsNetIdErr != nil {
		return nil, errors.Wrap(_targetAmsNetIdErr, "Error parsing 'targetAmsNetId' field of AmsPacket")
	}
	targetAmsNetId := _targetAmsNetId.(AmsNetId)
	if closeErr := readBuffer.CloseContext("targetAmsNetId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for targetAmsNetId")
	}

	// Simple Field (targetAmsPort)
	_targetAmsPort, _targetAmsPortErr := readBuffer.ReadUint16("targetAmsPort", 16)
	if _targetAmsPortErr != nil {
		return nil, errors.Wrap(_targetAmsPortErr, "Error parsing 'targetAmsPort' field of AmsPacket")
	}
	targetAmsPort := _targetAmsPort

	// Simple Field (sourceAmsNetId)
	if pullErr := readBuffer.PullContext("sourceAmsNetId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for sourceAmsNetId")
	}
	_sourceAmsNetId, _sourceAmsNetIdErr := AmsNetIdParseWithBuffer(readBuffer)
	if _sourceAmsNetIdErr != nil {
		return nil, errors.Wrap(_sourceAmsNetIdErr, "Error parsing 'sourceAmsNetId' field of AmsPacket")
	}
	sourceAmsNetId := _sourceAmsNetId.(AmsNetId)
	if closeErr := readBuffer.CloseContext("sourceAmsNetId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for sourceAmsNetId")
	}

	// Simple Field (sourceAmsPort)
	_sourceAmsPort, _sourceAmsPortErr := readBuffer.ReadUint16("sourceAmsPort", 16)
	if _sourceAmsPortErr != nil {
		return nil, errors.Wrap(_sourceAmsPortErr, "Error parsing 'sourceAmsPort' field of AmsPacket")
	}
	sourceAmsPort := _sourceAmsPort

	// Discriminator Field (commandId) (Used as input to a switch field)
	if pullErr := readBuffer.PullContext("commandId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for commandId")
	}
	commandId_temp, _commandIdErr := CommandIdParseWithBuffer(readBuffer)
	var commandId CommandId = commandId_temp
	if closeErr := readBuffer.CloseContext("commandId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for commandId")
	}
	if _commandIdErr != nil {
		return nil, errors.Wrap(_commandIdErr, "Error parsing 'commandId' field of AmsPacket")
	}

	// Const Field (initCommand)
	initCommand, _initCommandErr := readBuffer.ReadBit("initCommand")
	if _initCommandErr != nil {
		return nil, errors.Wrap(_initCommandErr, "Error parsing 'initCommand' field of AmsPacket")
	}
	if initCommand != AmsPacket_INITCOMMAND {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%t", AmsPacket_INITCOMMAND) + " but got " + fmt.Sprintf("%t", initCommand))
	}

	// Const Field (updCommand)
	updCommand, _updCommandErr := readBuffer.ReadBit("updCommand")
	if _updCommandErr != nil {
		return nil, errors.Wrap(_updCommandErr, "Error parsing 'updCommand' field of AmsPacket")
	}
	if updCommand != AmsPacket_UPDCOMMAND {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%t", AmsPacket_UPDCOMMAND) + " but got " + fmt.Sprintf("%t", updCommand))
	}

	// Const Field (timestampAdded)
	timestampAdded, _timestampAddedErr := readBuffer.ReadBit("timestampAdded")
	if _timestampAddedErr != nil {
		return nil, errors.Wrap(_timestampAddedErr, "Error parsing 'timestampAdded' field of AmsPacket")
	}
	if timestampAdded != AmsPacket_TIMESTAMPADDED {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%t", AmsPacket_TIMESTAMPADDED) + " but got " + fmt.Sprintf("%t", timestampAdded))
	}

	// Const Field (highPriorityCommand)
	highPriorityCommand, _highPriorityCommandErr := readBuffer.ReadBit("highPriorityCommand")
	if _highPriorityCommandErr != nil {
		return nil, errors.Wrap(_highPriorityCommandErr, "Error parsing 'highPriorityCommand' field of AmsPacket")
	}
	if highPriorityCommand != AmsPacket_HIGHPRIORITYCOMMAND {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%t", AmsPacket_HIGHPRIORITYCOMMAND) + " but got " + fmt.Sprintf("%t", highPriorityCommand))
	}

	// Const Field (systemCommand)
	systemCommand, _systemCommandErr := readBuffer.ReadBit("systemCommand")
	if _systemCommandErr != nil {
		return nil, errors.Wrap(_systemCommandErr, "Error parsing 'systemCommand' field of AmsPacket")
	}
	if systemCommand != AmsPacket_SYSTEMCOMMAND {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%t", AmsPacket_SYSTEMCOMMAND) + " but got " + fmt.Sprintf("%t", systemCommand))
	}

	// Const Field (adsCommand)
	adsCommand, _adsCommandErr := readBuffer.ReadBit("adsCommand")
	if _adsCommandErr != nil {
		return nil, errors.Wrap(_adsCommandErr, "Error parsing 'adsCommand' field of AmsPacket")
	}
	if adsCommand != AmsPacket_ADSCOMMAND {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%t", AmsPacket_ADSCOMMAND) + " but got " + fmt.Sprintf("%t", adsCommand))
	}

	// Const Field (noReturn)
	noReturn, _noReturnErr := readBuffer.ReadBit("noReturn")
	if _noReturnErr != nil {
		return nil, errors.Wrap(_noReturnErr, "Error parsing 'noReturn' field of AmsPacket")
	}
	if noReturn != AmsPacket_NORETURN {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%t", AmsPacket_NORETURN) + " but got " + fmt.Sprintf("%t", noReturn))
	}

	// Discriminator Field (response) (Used as input to a switch field)
	response, _responseErr := readBuffer.ReadBit("response")
	if _responseErr != nil {
		return nil, errors.Wrap(_responseErr, "Error parsing 'response' field of AmsPacket")
	}

	// Const Field (broadcast)
	broadcast, _broadcastErr := readBuffer.ReadBit("broadcast")
	if _broadcastErr != nil {
		return nil, errors.Wrap(_broadcastErr, "Error parsing 'broadcast' field of AmsPacket")
	}
	if broadcast != AmsPacket_BROADCAST {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%t", AmsPacket_BROADCAST) + " but got " + fmt.Sprintf("%t", broadcast))
	}

	var reservedField0 *int8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadInt8("reserved", 7)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of AmsPacket")
		}
		if reserved != int8(0x0) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": int8(0x0),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Implicit Field (length) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	length, _lengthErr := readBuffer.ReadUint32("length", 32)
	_ = length
	if _lengthErr != nil {
		return nil, errors.Wrap(_lengthErr, "Error parsing 'length' field of AmsPacket")
	}

	// Simple Field (errorCode)
	_errorCode, _errorCodeErr := readBuffer.ReadUint32("errorCode", 32)
	if _errorCodeErr != nil {
		return nil, errors.Wrap(_errorCodeErr, "Error parsing 'errorCode' field of AmsPacket")
	}
	errorCode := _errorCode

	// Simple Field (invokeId)
	_invokeId, _invokeIdErr := readBuffer.ReadUint32("invokeId", 32)
	if _invokeIdErr != nil {
		return nil, errors.Wrap(_invokeIdErr, "Error parsing 'invokeId' field of AmsPacket")
	}
	invokeId := _invokeId

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type AmsPacketChildSerializeRequirement interface {
		AmsPacket
		InitializeParent(AmsPacket, AmsNetId, uint16, AmsNetId, uint16, uint32, uint32)
		GetParent() AmsPacket
	}
	var _childTemp interface{}
	var _child AmsPacketChildSerializeRequirement
	var typeSwitchError error
	switch {
	case errorCode == 0x00000000 && commandId == CommandId_INVALID && response == bool(false): // AdsInvalidRequest
		_childTemp, typeSwitchError = AdsInvalidRequestParseWithBuffer(readBuffer)
	case errorCode == 0x00000000 && commandId == CommandId_INVALID && response == bool(true): // AdsInvalidResponse
		_childTemp, typeSwitchError = AdsInvalidResponseParseWithBuffer(readBuffer)
	case errorCode == 0x00000000 && commandId == CommandId_ADS_READ_DEVICE_INFO && response == bool(false): // AdsReadDeviceInfoRequest
		_childTemp, typeSwitchError = AdsReadDeviceInfoRequestParseWithBuffer(readBuffer)
	case errorCode == 0x00000000 && commandId == CommandId_ADS_READ_DEVICE_INFO && response == bool(true): // AdsReadDeviceInfoResponse
		_childTemp, typeSwitchError = AdsReadDeviceInfoResponseParseWithBuffer(readBuffer)
	case errorCode == 0x00000000 && commandId == CommandId_ADS_READ && response == bool(false): // AdsReadRequest
		_childTemp, typeSwitchError = AdsReadRequestParseWithBuffer(readBuffer)
	case errorCode == 0x00000000 && commandId == CommandId_ADS_READ && response == bool(true): // AdsReadResponse
		_childTemp, typeSwitchError = AdsReadResponseParseWithBuffer(readBuffer)
	case errorCode == 0x00000000 && commandId == CommandId_ADS_WRITE && response == bool(false): // AdsWriteRequest
		_childTemp, typeSwitchError = AdsWriteRequestParseWithBuffer(readBuffer)
	case errorCode == 0x00000000 && commandId == CommandId_ADS_WRITE && response == bool(true): // AdsWriteResponse
		_childTemp, typeSwitchError = AdsWriteResponseParseWithBuffer(readBuffer)
	case errorCode == 0x00000000 && commandId == CommandId_ADS_READ_STATE && response == bool(false): // AdsReadStateRequest
		_childTemp, typeSwitchError = AdsReadStateRequestParseWithBuffer(readBuffer)
	case errorCode == 0x00000000 && commandId == CommandId_ADS_READ_STATE && response == bool(true): // AdsReadStateResponse
		_childTemp, typeSwitchError = AdsReadStateResponseParseWithBuffer(readBuffer)
	case errorCode == 0x00000000 && commandId == CommandId_ADS_WRITE_CONTROL && response == bool(false): // AdsWriteControlRequest
		_childTemp, typeSwitchError = AdsWriteControlRequestParseWithBuffer(readBuffer)
	case errorCode == 0x00000000 && commandId == CommandId_ADS_WRITE_CONTROL && response == bool(true): // AdsWriteControlResponse
		_childTemp, typeSwitchError = AdsWriteControlResponseParseWithBuffer(readBuffer)
	case errorCode == 0x00000000 && commandId == CommandId_ADS_ADD_DEVICE_NOTIFICATION && response == bool(false): // AdsAddDeviceNotificationRequest
		_childTemp, typeSwitchError = AdsAddDeviceNotificationRequestParseWithBuffer(readBuffer)
	case errorCode == 0x00000000 && commandId == CommandId_ADS_ADD_DEVICE_NOTIFICATION && response == bool(true): // AdsAddDeviceNotificationResponse
		_childTemp, typeSwitchError = AdsAddDeviceNotificationResponseParseWithBuffer(readBuffer)
	case errorCode == 0x00000000 && commandId == CommandId_ADS_DELETE_DEVICE_NOTIFICATION && response == bool(false): // AdsDeleteDeviceNotificationRequest
		_childTemp, typeSwitchError = AdsDeleteDeviceNotificationRequestParseWithBuffer(readBuffer)
	case errorCode == 0x00000000 && commandId == CommandId_ADS_DELETE_DEVICE_NOTIFICATION && response == bool(true): // AdsDeleteDeviceNotificationResponse
		_childTemp, typeSwitchError = AdsDeleteDeviceNotificationResponseParseWithBuffer(readBuffer)
	case errorCode == 0x00000000 && commandId == CommandId_ADS_DEVICE_NOTIFICATION && response == bool(false): // AdsDeviceNotificationRequest
		_childTemp, typeSwitchError = AdsDeviceNotificationRequestParseWithBuffer(readBuffer)
	case errorCode == 0x00000000 && commandId == CommandId_ADS_DEVICE_NOTIFICATION && response == bool(true): // AdsDeviceNotificationResponse
		_childTemp, typeSwitchError = AdsDeviceNotificationResponseParseWithBuffer(readBuffer)
	case errorCode == 0x00000000 && commandId == CommandId_ADS_READ_WRITE && response == bool(false): // AdsReadWriteRequest
		_childTemp, typeSwitchError = AdsReadWriteRequestParseWithBuffer(readBuffer)
	case errorCode == 0x00000000 && commandId == CommandId_ADS_READ_WRITE && response == bool(true): // AdsReadWriteResponse
		_childTemp, typeSwitchError = AdsReadWriteResponseParseWithBuffer(readBuffer)
	case true: // ErrorResponse
		_childTemp, typeSwitchError = ErrorResponseParseWithBuffer(readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [errorCode=%v, commandId=%v, response=%v]", errorCode, commandId, response)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of AmsPacket")
	}
	_child = _childTemp.(AmsPacketChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("AmsPacket"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AmsPacket")
	}

	// Finish initializing
	_child.InitializeParent(_child, targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId)
	_child.GetParent().(*_AmsPacket).reservedField0 = reservedField0
	return _child, nil
}

func (pm *_AmsPacket) SerializeParent(writeBuffer utils.WriteBuffer, child AmsPacket, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("AmsPacket"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AmsPacket")
	}

	// Simple Field (targetAmsNetId)
	if pushErr := writeBuffer.PushContext("targetAmsNetId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for targetAmsNetId")
	}
	_targetAmsNetIdErr := writeBuffer.WriteSerializable(m.GetTargetAmsNetId())
	if popErr := writeBuffer.PopContext("targetAmsNetId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for targetAmsNetId")
	}
	if _targetAmsNetIdErr != nil {
		return errors.Wrap(_targetAmsNetIdErr, "Error serializing 'targetAmsNetId' field")
	}

	// Simple Field (targetAmsPort)
	targetAmsPort := uint16(m.GetTargetAmsPort())
	_targetAmsPortErr := writeBuffer.WriteUint16("targetAmsPort", 16, (targetAmsPort))
	if _targetAmsPortErr != nil {
		return errors.Wrap(_targetAmsPortErr, "Error serializing 'targetAmsPort' field")
	}

	// Simple Field (sourceAmsNetId)
	if pushErr := writeBuffer.PushContext("sourceAmsNetId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for sourceAmsNetId")
	}
	_sourceAmsNetIdErr := writeBuffer.WriteSerializable(m.GetSourceAmsNetId())
	if popErr := writeBuffer.PopContext("sourceAmsNetId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for sourceAmsNetId")
	}
	if _sourceAmsNetIdErr != nil {
		return errors.Wrap(_sourceAmsNetIdErr, "Error serializing 'sourceAmsNetId' field")
	}

	// Simple Field (sourceAmsPort)
	sourceAmsPort := uint16(m.GetSourceAmsPort())
	_sourceAmsPortErr := writeBuffer.WriteUint16("sourceAmsPort", 16, (sourceAmsPort))
	if _sourceAmsPortErr != nil {
		return errors.Wrap(_sourceAmsPortErr, "Error serializing 'sourceAmsPort' field")
	}

	// Discriminator Field (commandId) (Used as input to a switch field)
	commandId := CommandId(child.GetCommandId())
	if pushErr := writeBuffer.PushContext("commandId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for commandId")
	}
	_commandIdErr := writeBuffer.WriteSerializable(commandId)
	if popErr := writeBuffer.PopContext("commandId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for commandId")
	}

	if _commandIdErr != nil {
		return errors.Wrap(_commandIdErr, "Error serializing 'commandId' field")
	}

	// Const Field (initCommand)
	_initCommandErr := writeBuffer.WriteBit("initCommand", false)
	if _initCommandErr != nil {
		return errors.Wrap(_initCommandErr, "Error serializing 'initCommand' field")
	}

	// Const Field (updCommand)
	_updCommandErr := writeBuffer.WriteBit("updCommand", false)
	if _updCommandErr != nil {
		return errors.Wrap(_updCommandErr, "Error serializing 'updCommand' field")
	}

	// Const Field (timestampAdded)
	_timestampAddedErr := writeBuffer.WriteBit("timestampAdded", false)
	if _timestampAddedErr != nil {
		return errors.Wrap(_timestampAddedErr, "Error serializing 'timestampAdded' field")
	}

	// Const Field (highPriorityCommand)
	_highPriorityCommandErr := writeBuffer.WriteBit("highPriorityCommand", false)
	if _highPriorityCommandErr != nil {
		return errors.Wrap(_highPriorityCommandErr, "Error serializing 'highPriorityCommand' field")
	}

	// Const Field (systemCommand)
	_systemCommandErr := writeBuffer.WriteBit("systemCommand", false)
	if _systemCommandErr != nil {
		return errors.Wrap(_systemCommandErr, "Error serializing 'systemCommand' field")
	}

	// Const Field (adsCommand)
	_adsCommandErr := writeBuffer.WriteBit("adsCommand", true)
	if _adsCommandErr != nil {
		return errors.Wrap(_adsCommandErr, "Error serializing 'adsCommand' field")
	}

	// Const Field (noReturn)
	_noReturnErr := writeBuffer.WriteBit("noReturn", false)
	if _noReturnErr != nil {
		return errors.Wrap(_noReturnErr, "Error serializing 'noReturn' field")
	}

	// Discriminator Field (response) (Used as input to a switch field)
	response := bool(child.GetResponse())
	_responseErr := writeBuffer.WriteBit("response", (response))

	if _responseErr != nil {
		return errors.Wrap(_responseErr, "Error serializing 'response' field")
	}

	// Const Field (broadcast)
	_broadcastErr := writeBuffer.WriteBit("broadcast", false)
	if _broadcastErr != nil {
		return errors.Wrap(_broadcastErr, "Error serializing 'broadcast' field")
	}

	// Reserved Field (reserved)
	{
		var reserved int8 = int8(0x0)
		if pm.reservedField0 != nil {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": int8(0x0),
				"got value":      reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *pm.reservedField0
		}
		_err := writeBuffer.WriteInt8("reserved", 7, reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Implicit Field (length) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	length := uint32(uint32(uint32(m.GetLengthInBytes())) - uint32(uint32(32)))
	_lengthErr := writeBuffer.WriteUint32("length", 32, (length))
	if _lengthErr != nil {
		return errors.Wrap(_lengthErr, "Error serializing 'length' field")
	}

	// Simple Field (errorCode)
	errorCode := uint32(m.GetErrorCode())
	_errorCodeErr := writeBuffer.WriteUint32("errorCode", 32, (errorCode))
	if _errorCodeErr != nil {
		return errors.Wrap(_errorCodeErr, "Error serializing 'errorCode' field")
	}

	// Simple Field (invokeId)
	invokeId := uint32(m.GetInvokeId())
	_invokeIdErr := writeBuffer.WriteUint32("invokeId", 32, (invokeId))
	if _invokeIdErr != nil {
		return errors.Wrap(_invokeIdErr, "Error serializing 'invokeId' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("AmsPacket"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AmsPacket")
	}
	return nil
}

func (m *_AmsPacket) isAmsPacket() bool {
	return true
}

func (m *_AmsPacket) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
