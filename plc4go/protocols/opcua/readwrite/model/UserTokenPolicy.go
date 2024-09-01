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

// UserTokenPolicy is the corresponding interface of UserTokenPolicy
type UserTokenPolicy interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetPolicyId returns PolicyId (property field)
	GetPolicyId() PascalString
	// GetTokenType returns TokenType (property field)
	GetTokenType() UserTokenType
	// GetIssuedTokenType returns IssuedTokenType (property field)
	GetIssuedTokenType() PascalString
	// GetIssuerEndpointUrl returns IssuerEndpointUrl (property field)
	GetIssuerEndpointUrl() PascalString
	// GetSecurityPolicyUri returns SecurityPolicyUri (property field)
	GetSecurityPolicyUri() PascalString
}

// UserTokenPolicyExactly can be used when we want exactly this type and not a type which fulfills UserTokenPolicy.
// This is useful for switch cases.
type UserTokenPolicyExactly interface {
	UserTokenPolicy
	isUserTokenPolicy() bool
}

// _UserTokenPolicy is the data-structure of this message
type _UserTokenPolicy struct {
	*_ExtensionObjectDefinition
	PolicyId          PascalString
	TokenType         UserTokenType
	IssuedTokenType   PascalString
	IssuerEndpointUrl PascalString
	SecurityPolicyUri PascalString
}

var _ UserTokenPolicy = (*_UserTokenPolicy)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_UserTokenPolicy) GetIdentifier() string {
	return "306"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_UserTokenPolicy) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_UserTokenPolicy) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_UserTokenPolicy) GetPolicyId() PascalString {
	return m.PolicyId
}

func (m *_UserTokenPolicy) GetTokenType() UserTokenType {
	return m.TokenType
}

func (m *_UserTokenPolicy) GetIssuedTokenType() PascalString {
	return m.IssuedTokenType
}

func (m *_UserTokenPolicy) GetIssuerEndpointUrl() PascalString {
	return m.IssuerEndpointUrl
}

func (m *_UserTokenPolicy) GetSecurityPolicyUri() PascalString {
	return m.SecurityPolicyUri
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewUserTokenPolicy factory function for _UserTokenPolicy
func NewUserTokenPolicy(policyId PascalString, tokenType UserTokenType, issuedTokenType PascalString, issuerEndpointUrl PascalString, securityPolicyUri PascalString) *_UserTokenPolicy {
	_result := &_UserTokenPolicy{
		PolicyId:                   policyId,
		TokenType:                  tokenType,
		IssuedTokenType:            issuedTokenType,
		IssuerEndpointUrl:          issuerEndpointUrl,
		SecurityPolicyUri:          securityPolicyUri,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastUserTokenPolicy(structType any) UserTokenPolicy {
	if casted, ok := structType.(UserTokenPolicy); ok {
		return casted
	}
	if casted, ok := structType.(*UserTokenPolicy); ok {
		return *casted
	}
	return nil
}

func (m *_UserTokenPolicy) GetTypeName() string {
	return "UserTokenPolicy"
}

func (m *_UserTokenPolicy) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (policyId)
	lengthInBits += m.PolicyId.GetLengthInBits(ctx)

	// Simple field (tokenType)
	lengthInBits += 32

	// Simple field (issuedTokenType)
	lengthInBits += m.IssuedTokenType.GetLengthInBits(ctx)

	// Simple field (issuerEndpointUrl)
	lengthInBits += m.IssuerEndpointUrl.GetLengthInBits(ctx)

	// Simple field (securityPolicyUri)
	lengthInBits += m.SecurityPolicyUri.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_UserTokenPolicy) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func UserTokenPolicyParse(ctx context.Context, theBytes []byte, identifier string) (UserTokenPolicy, error) {
	return UserTokenPolicyParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func UserTokenPolicyParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (UserTokenPolicy, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (UserTokenPolicy, error) {
		return UserTokenPolicyParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func UserTokenPolicyParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (UserTokenPolicy, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("UserTokenPolicy"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for UserTokenPolicy")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	policyId, err := ReadSimpleField[PascalString](ctx, "policyId", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'policyId' field"))
	}

	tokenType, err := ReadEnumField[UserTokenType](ctx, "tokenType", "UserTokenType", ReadEnum(UserTokenTypeByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'tokenType' field"))
	}

	issuedTokenType, err := ReadSimpleField[PascalString](ctx, "issuedTokenType", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'issuedTokenType' field"))
	}

	issuerEndpointUrl, err := ReadSimpleField[PascalString](ctx, "issuerEndpointUrl", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'issuerEndpointUrl' field"))
	}

	securityPolicyUri, err := ReadSimpleField[PascalString](ctx, "securityPolicyUri", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityPolicyUri' field"))
	}

	if closeErr := readBuffer.CloseContext("UserTokenPolicy"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for UserTokenPolicy")
	}

	// Create a partially initialized instance
	_child := &_UserTokenPolicy{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		PolicyId:                   policyId,
		TokenType:                  tokenType,
		IssuedTokenType:            issuedTokenType,
		IssuerEndpointUrl:          issuerEndpointUrl,
		SecurityPolicyUri:          securityPolicyUri,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_UserTokenPolicy) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_UserTokenPolicy) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("UserTokenPolicy"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for UserTokenPolicy")
		}

		if err := WriteSimpleField[PascalString](ctx, "policyId", m.GetPolicyId(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'policyId' field")
		}

		if err := WriteSimpleEnumField[UserTokenType](ctx, "tokenType", "UserTokenType", m.GetTokenType(), WriteEnum[UserTokenType, uint32](UserTokenType.GetValue, UserTokenType.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'tokenType' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "issuedTokenType", m.GetIssuedTokenType(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'issuedTokenType' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "issuerEndpointUrl", m.GetIssuerEndpointUrl(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'issuerEndpointUrl' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "securityPolicyUri", m.GetSecurityPolicyUri(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'securityPolicyUri' field")
		}

		if popErr := writeBuffer.PopContext("UserTokenPolicy"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for UserTokenPolicy")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_UserTokenPolicy) isUserTokenPolicy() bool {
	return true
}

func (m *_UserTokenPolicy) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
