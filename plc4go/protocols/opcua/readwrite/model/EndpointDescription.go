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

// EndpointDescription is the corresponding interface of EndpointDescription
type EndpointDescription interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetEndpointUrl returns EndpointUrl (property field)
	GetEndpointUrl() PascalString
	// GetServer returns Server (property field)
	GetServer() ExtensionObjectDefinition
	// GetServerCertificate returns ServerCertificate (property field)
	GetServerCertificate() PascalByteString
	// GetSecurityMode returns SecurityMode (property field)
	GetSecurityMode() MessageSecurityMode
	// GetSecurityPolicyUri returns SecurityPolicyUri (property field)
	GetSecurityPolicyUri() PascalString
	// GetNoOfUserIdentityTokens returns NoOfUserIdentityTokens (property field)
	GetNoOfUserIdentityTokens() int32
	// GetUserIdentityTokens returns UserIdentityTokens (property field)
	GetUserIdentityTokens() []ExtensionObjectDefinition
	// GetTransportProfileUri returns TransportProfileUri (property field)
	GetTransportProfileUri() PascalString
	// GetSecurityLevel returns SecurityLevel (property field)
	GetSecurityLevel() uint8
	// IsEndpointDescription is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsEndpointDescription()
}

// _EndpointDescription is the data-structure of this message
type _EndpointDescription struct {
	ExtensionObjectDefinitionContract
	EndpointUrl            PascalString
	Server                 ExtensionObjectDefinition
	ServerCertificate      PascalByteString
	SecurityMode           MessageSecurityMode
	SecurityPolicyUri      PascalString
	NoOfUserIdentityTokens int32
	UserIdentityTokens     []ExtensionObjectDefinition
	TransportProfileUri    PascalString
	SecurityLevel          uint8
}

var _ EndpointDescription = (*_EndpointDescription)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_EndpointDescription)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_EndpointDescription) GetIdentifier() string {
	return "314"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_EndpointDescription) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_EndpointDescription) GetEndpointUrl() PascalString {
	return m.EndpointUrl
}

func (m *_EndpointDescription) GetServer() ExtensionObjectDefinition {
	return m.Server
}

func (m *_EndpointDescription) GetServerCertificate() PascalByteString {
	return m.ServerCertificate
}

func (m *_EndpointDescription) GetSecurityMode() MessageSecurityMode {
	return m.SecurityMode
}

func (m *_EndpointDescription) GetSecurityPolicyUri() PascalString {
	return m.SecurityPolicyUri
}

func (m *_EndpointDescription) GetNoOfUserIdentityTokens() int32 {
	return m.NoOfUserIdentityTokens
}

func (m *_EndpointDescription) GetUserIdentityTokens() []ExtensionObjectDefinition {
	return m.UserIdentityTokens
}

func (m *_EndpointDescription) GetTransportProfileUri() PascalString {
	return m.TransportProfileUri
}

func (m *_EndpointDescription) GetSecurityLevel() uint8 {
	return m.SecurityLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewEndpointDescription factory function for _EndpointDescription
func NewEndpointDescription(endpointUrl PascalString, server ExtensionObjectDefinition, serverCertificate PascalByteString, securityMode MessageSecurityMode, securityPolicyUri PascalString, noOfUserIdentityTokens int32, userIdentityTokens []ExtensionObjectDefinition, transportProfileUri PascalString, securityLevel uint8) *_EndpointDescription {
	_result := &_EndpointDescription{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		EndpointUrl:                       endpointUrl,
		Server:                            server,
		ServerCertificate:                 serverCertificate,
		SecurityMode:                      securityMode,
		SecurityPolicyUri:                 securityPolicyUri,
		NoOfUserIdentityTokens:            noOfUserIdentityTokens,
		UserIdentityTokens:                userIdentityTokens,
		TransportProfileUri:               transportProfileUri,
		SecurityLevel:                     securityLevel,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastEndpointDescription(structType any) EndpointDescription {
	if casted, ok := structType.(EndpointDescription); ok {
		return casted
	}
	if casted, ok := structType.(*EndpointDescription); ok {
		return *casted
	}
	return nil
}

func (m *_EndpointDescription) GetTypeName() string {
	return "EndpointDescription"
}

func (m *_EndpointDescription) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (endpointUrl)
	lengthInBits += m.EndpointUrl.GetLengthInBits(ctx)

	// Simple field (server)
	lengthInBits += m.Server.GetLengthInBits(ctx)

	// Simple field (serverCertificate)
	lengthInBits += m.ServerCertificate.GetLengthInBits(ctx)

	// Simple field (securityMode)
	lengthInBits += 32

	// Simple field (securityPolicyUri)
	lengthInBits += m.SecurityPolicyUri.GetLengthInBits(ctx)

	// Simple field (noOfUserIdentityTokens)
	lengthInBits += 32

	// Array field
	if len(m.UserIdentityTokens) > 0 {
		for _curItem, element := range m.UserIdentityTokens {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.UserIdentityTokens), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (transportProfileUri)
	lengthInBits += m.TransportProfileUri.GetLengthInBits(ctx)

	// Simple field (securityLevel)
	lengthInBits += 8

	return lengthInBits
}

func (m *_EndpointDescription) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_EndpointDescription) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__endpointDescription EndpointDescription, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("EndpointDescription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EndpointDescription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	endpointUrl, err := ReadSimpleField[PascalString](ctx, "endpointUrl", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'endpointUrl' field"))
	}
	m.EndpointUrl = endpointUrl

	server, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "server", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("310")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'server' field"))
	}
	m.Server = server

	serverCertificate, err := ReadSimpleField[PascalByteString](ctx, "serverCertificate", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serverCertificate' field"))
	}
	m.ServerCertificate = serverCertificate

	securityMode, err := ReadEnumField[MessageSecurityMode](ctx, "securityMode", "MessageSecurityMode", ReadEnum(MessageSecurityModeByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityMode' field"))
	}
	m.SecurityMode = securityMode

	securityPolicyUri, err := ReadSimpleField[PascalString](ctx, "securityPolicyUri", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityPolicyUri' field"))
	}
	m.SecurityPolicyUri = securityPolicyUri

	noOfUserIdentityTokens, err := ReadSimpleField(ctx, "noOfUserIdentityTokens", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfUserIdentityTokens' field"))
	}
	m.NoOfUserIdentityTokens = noOfUserIdentityTokens

	userIdentityTokens, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "userIdentityTokens", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("306")), readBuffer), uint64(noOfUserIdentityTokens))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'userIdentityTokens' field"))
	}
	m.UserIdentityTokens = userIdentityTokens

	transportProfileUri, err := ReadSimpleField[PascalString](ctx, "transportProfileUri", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transportProfileUri' field"))
	}
	m.TransportProfileUri = transportProfileUri

	securityLevel, err := ReadSimpleField(ctx, "securityLevel", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityLevel' field"))
	}
	m.SecurityLevel = securityLevel

	if closeErr := readBuffer.CloseContext("EndpointDescription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EndpointDescription")
	}

	return m, nil
}

func (m *_EndpointDescription) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_EndpointDescription) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("EndpointDescription"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for EndpointDescription")
		}

		if err := WriteSimpleField[PascalString](ctx, "endpointUrl", m.GetEndpointUrl(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'endpointUrl' field")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "server", m.GetServer(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'server' field")
		}

		if err := WriteSimpleField[PascalByteString](ctx, "serverCertificate", m.GetServerCertificate(), WriteComplex[PascalByteString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'serverCertificate' field")
		}

		if err := WriteSimpleEnumField[MessageSecurityMode](ctx, "securityMode", "MessageSecurityMode", m.GetSecurityMode(), WriteEnum[MessageSecurityMode, uint32](MessageSecurityMode.GetValue, MessageSecurityMode.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'securityMode' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "securityPolicyUri", m.GetSecurityPolicyUri(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'securityPolicyUri' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfUserIdentityTokens", m.GetNoOfUserIdentityTokens(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfUserIdentityTokens' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "userIdentityTokens", m.GetUserIdentityTokens(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'userIdentityTokens' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "transportProfileUri", m.GetTransportProfileUri(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'transportProfileUri' field")
		}

		if err := WriteSimpleField[uint8](ctx, "securityLevel", m.GetSecurityLevel(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'securityLevel' field")
		}

		if popErr := writeBuffer.PopContext("EndpointDescription"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for EndpointDescription")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_EndpointDescription) IsEndpointDescription() {}

func (m *_EndpointDescription) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
