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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// UserIdentityTokenDefinition is the corresponding interface of UserIdentityTokenDefinition
type UserIdentityTokenDefinition interface {
	UserIdentityTokenDefinitionContract
	UserIdentityTokenDefinitionRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// IsUserIdentityTokenDefinition is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsUserIdentityTokenDefinition()
}

// UserIdentityTokenDefinitionContract provides a set of functions which can be overwritten by a sub struct
type UserIdentityTokenDefinitionContract interface {
	// IsUserIdentityTokenDefinition is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsUserIdentityTokenDefinition()
}

// UserIdentityTokenDefinitionRequirements provides a set of functions which need to be implemented by a sub struct
type UserIdentityTokenDefinitionRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetIdentifier returns Identifier (discriminator field)
	GetIdentifier() string
}

// _UserIdentityTokenDefinition is the data-structure of this message
type _UserIdentityTokenDefinition struct {
	_SubType UserIdentityTokenDefinition
}

var _ UserIdentityTokenDefinitionContract = (*_UserIdentityTokenDefinition)(nil)

// NewUserIdentityTokenDefinition factory function for _UserIdentityTokenDefinition
func NewUserIdentityTokenDefinition() *_UserIdentityTokenDefinition {
	return &_UserIdentityTokenDefinition{}
}

// Deprecated: use the interface for direct cast
func CastUserIdentityTokenDefinition(structType any) UserIdentityTokenDefinition {
	if casted, ok := structType.(UserIdentityTokenDefinition); ok {
		return casted
	}
	if casted, ok := structType.(*UserIdentityTokenDefinition); ok {
		return *casted
	}
	return nil
}

func (m *_UserIdentityTokenDefinition) GetTypeName() string {
	return "UserIdentityTokenDefinition"
}

func (m *_UserIdentityTokenDefinition) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_UserIdentityTokenDefinition) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func UserIdentityTokenDefinitionParse[T UserIdentityTokenDefinition](ctx context.Context, theBytes []byte, identifier string) (T, error) {
	return UserIdentityTokenDefinitionParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func UserIdentityTokenDefinitionParseWithBufferProducer[T UserIdentityTokenDefinition](identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := UserIdentityTokenDefinitionParseWithBuffer[T](ctx, readBuffer, identifier)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func UserIdentityTokenDefinitionParseWithBuffer[T UserIdentityTokenDefinition](ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (T, error) {
	v, err := (&_UserIdentityTokenDefinition{}).parse(ctx, readBuffer, identifier)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_UserIdentityTokenDefinition) parse(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (__userIdentityTokenDefinition UserIdentityTokenDefinition, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("UserIdentityTokenDefinition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for UserIdentityTokenDefinition")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child UserIdentityTokenDefinition
	switch {
	case identifier == "anonymous": // AnonymousIdentityToken
		if _child, err = (&_AnonymousIdentityToken{}).parse(ctx, readBuffer, m, identifier); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type AnonymousIdentityToken for type-switch of UserIdentityTokenDefinition")
		}
	case identifier == "username": // UserNameIdentityToken
		if _child, err = (&_UserNameIdentityToken{}).parse(ctx, readBuffer, m, identifier); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type UserNameIdentityToken for type-switch of UserIdentityTokenDefinition")
		}
	case identifier == "certificate": // X509IdentityToken
		if _child, err = (&_X509IdentityToken{}).parse(ctx, readBuffer, m, identifier); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type X509IdentityToken for type-switch of UserIdentityTokenDefinition")
		}
	case identifier == "identity": // IssuedIdentityToken
		if _child, err = (&_IssuedIdentityToken{}).parse(ctx, readBuffer, m, identifier); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type IssuedIdentityToken for type-switch of UserIdentityTokenDefinition")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [identifier=%v]", identifier)
	}

	if closeErr := readBuffer.CloseContext("UserIdentityTokenDefinition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for UserIdentityTokenDefinition")
	}

	return _child, nil
}

func (pm *_UserIdentityTokenDefinition) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child UserIdentityTokenDefinition, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("UserIdentityTokenDefinition"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for UserIdentityTokenDefinition")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("UserIdentityTokenDefinition"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for UserIdentityTokenDefinition")
	}
	return nil
}

func (m *_UserIdentityTokenDefinition) IsUserIdentityTokenDefinition() {}
