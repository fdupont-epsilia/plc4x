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

// SecurityDataGasAlarmCleared is the corresponding interface of SecurityDataGasAlarmCleared
type SecurityDataGasAlarmCleared interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SecurityData
	// IsSecurityDataGasAlarmCleared is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataGasAlarmCleared()
	// CreateBuilder creates a SecurityDataGasAlarmClearedBuilder
	CreateSecurityDataGasAlarmClearedBuilder() SecurityDataGasAlarmClearedBuilder
}

// _SecurityDataGasAlarmCleared is the data-structure of this message
type _SecurityDataGasAlarmCleared struct {
	SecurityDataContract
}

var _ SecurityDataGasAlarmCleared = (*_SecurityDataGasAlarmCleared)(nil)
var _ SecurityDataRequirements = (*_SecurityDataGasAlarmCleared)(nil)

// NewSecurityDataGasAlarmCleared factory function for _SecurityDataGasAlarmCleared
func NewSecurityDataGasAlarmCleared(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataGasAlarmCleared {
	_result := &_SecurityDataGasAlarmCleared{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SecurityDataGasAlarmClearedBuilder is a builder for SecurityDataGasAlarmCleared
type SecurityDataGasAlarmClearedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() SecurityDataGasAlarmClearedBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() SecurityDataBuilder
	// Build builds the SecurityDataGasAlarmCleared or returns an error if something is wrong
	Build() (SecurityDataGasAlarmCleared, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SecurityDataGasAlarmCleared
}

// NewSecurityDataGasAlarmClearedBuilder() creates a SecurityDataGasAlarmClearedBuilder
func NewSecurityDataGasAlarmClearedBuilder() SecurityDataGasAlarmClearedBuilder {
	return &_SecurityDataGasAlarmClearedBuilder{_SecurityDataGasAlarmCleared: new(_SecurityDataGasAlarmCleared)}
}

type _SecurityDataGasAlarmClearedBuilder struct {
	*_SecurityDataGasAlarmCleared

	parentBuilder *_SecurityDataBuilder

	err *utils.MultiError
}

var _ (SecurityDataGasAlarmClearedBuilder) = (*_SecurityDataGasAlarmClearedBuilder)(nil)

func (b *_SecurityDataGasAlarmClearedBuilder) setParent(contract SecurityDataContract) {
	b.SecurityDataContract = contract
	contract.(*_SecurityData)._SubType = b._SecurityDataGasAlarmCleared
}

func (b *_SecurityDataGasAlarmClearedBuilder) WithMandatoryFields() SecurityDataGasAlarmClearedBuilder {
	return b
}

func (b *_SecurityDataGasAlarmClearedBuilder) Build() (SecurityDataGasAlarmCleared, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SecurityDataGasAlarmCleared.deepCopy(), nil
}

func (b *_SecurityDataGasAlarmClearedBuilder) MustBuild() SecurityDataGasAlarmCleared {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SecurityDataGasAlarmClearedBuilder) Done() SecurityDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewSecurityDataBuilder().(*_SecurityDataBuilder)
	}
	return b.parentBuilder
}

func (b *_SecurityDataGasAlarmClearedBuilder) buildForSecurityData() (SecurityData, error) {
	return b.Build()
}

func (b *_SecurityDataGasAlarmClearedBuilder) DeepCopy() any {
	_copy := b.CreateSecurityDataGasAlarmClearedBuilder().(*_SecurityDataGasAlarmClearedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSecurityDataGasAlarmClearedBuilder creates a SecurityDataGasAlarmClearedBuilder
func (b *_SecurityDataGasAlarmCleared) CreateSecurityDataGasAlarmClearedBuilder() SecurityDataGasAlarmClearedBuilder {
	if b == nil {
		return NewSecurityDataGasAlarmClearedBuilder()
	}
	return &_SecurityDataGasAlarmClearedBuilder{_SecurityDataGasAlarmCleared: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataGasAlarmCleared) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

// Deprecated: use the interface for direct cast
func CastSecurityDataGasAlarmCleared(structType any) SecurityDataGasAlarmCleared {
	if casted, ok := structType.(SecurityDataGasAlarmCleared); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataGasAlarmCleared); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataGasAlarmCleared) GetTypeName() string {
	return "SecurityDataGasAlarmCleared"
}

func (m *_SecurityDataGasAlarmCleared) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataGasAlarmCleared) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataGasAlarmCleared) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataGasAlarmCleared SecurityDataGasAlarmCleared, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataGasAlarmCleared"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataGasAlarmCleared")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataGasAlarmCleared"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataGasAlarmCleared")
	}

	return m, nil
}

func (m *_SecurityDataGasAlarmCleared) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataGasAlarmCleared) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataGasAlarmCleared"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataGasAlarmCleared")
		}

		if popErr := writeBuffer.PopContext("SecurityDataGasAlarmCleared"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataGasAlarmCleared")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataGasAlarmCleared) IsSecurityDataGasAlarmCleared() {}

func (m *_SecurityDataGasAlarmCleared) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SecurityDataGasAlarmCleared) deepCopy() *_SecurityDataGasAlarmCleared {
	if m == nil {
		return nil
	}
	_SecurityDataGasAlarmClearedCopy := &_SecurityDataGasAlarmCleared{
		m.SecurityDataContract.(*_SecurityData).deepCopy(),
	}
	_SecurityDataGasAlarmClearedCopy.SecurityDataContract.(*_SecurityData)._SubType = m
	return _SecurityDataGasAlarmClearedCopy
}

func (m *_SecurityDataGasAlarmCleared) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
