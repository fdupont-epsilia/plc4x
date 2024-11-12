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

// BACnetConfirmedServiceRequestDeviceCommunicationControl is the corresponding interface of BACnetConfirmedServiceRequestDeviceCommunicationControl
type BACnetConfirmedServiceRequestDeviceCommunicationControl interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConfirmedServiceRequest
	// GetTimeDuration returns TimeDuration (property field)
	GetTimeDuration() BACnetContextTagUnsignedInteger
	// GetEnableDisable returns EnableDisable (property field)
	GetEnableDisable() BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged
	// GetPassword returns Password (property field)
	GetPassword() BACnetContextTagCharacterString
	// IsBACnetConfirmedServiceRequestDeviceCommunicationControl is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestDeviceCommunicationControl()
	// CreateBuilder creates a BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder
	CreateBACnetConfirmedServiceRequestDeviceCommunicationControlBuilder() BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder
}

// _BACnetConfirmedServiceRequestDeviceCommunicationControl is the data-structure of this message
type _BACnetConfirmedServiceRequestDeviceCommunicationControl struct {
	BACnetConfirmedServiceRequestContract
	TimeDuration  BACnetContextTagUnsignedInteger
	EnableDisable BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged
	Password      BACnetContextTagCharacterString
}

var _ BACnetConfirmedServiceRequestDeviceCommunicationControl = (*_BACnetConfirmedServiceRequestDeviceCommunicationControl)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestDeviceCommunicationControl)(nil)

// NewBACnetConfirmedServiceRequestDeviceCommunicationControl factory function for _BACnetConfirmedServiceRequestDeviceCommunicationControl
func NewBACnetConfirmedServiceRequestDeviceCommunicationControl(timeDuration BACnetContextTagUnsignedInteger, enableDisable BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged, password BACnetContextTagCharacterString, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestDeviceCommunicationControl {
	if enableDisable == nil {
		panic("enableDisable of type BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged for BACnetConfirmedServiceRequestDeviceCommunicationControl must not be nil")
	}
	_result := &_BACnetConfirmedServiceRequestDeviceCommunicationControl{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		TimeDuration:                          timeDuration,
		EnableDisable:                         enableDisable,
		Password:                              password,
	}
	_result.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder is a builder for BACnetConfirmedServiceRequestDeviceCommunicationControl
type BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(enableDisable BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged) BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder
	// WithTimeDuration adds TimeDuration (property field)
	WithOptionalTimeDuration(BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder
	// WithOptionalTimeDurationBuilder adds TimeDuration (property field) which is build by the builder
	WithOptionalTimeDurationBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder
	// WithEnableDisable adds EnableDisable (property field)
	WithEnableDisable(BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged) BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder
	// WithEnableDisableBuilder adds EnableDisable (property field) which is build by the builder
	WithEnableDisableBuilder(func(BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTaggedBuilder) BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTaggedBuilder) BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder
	// WithPassword adds Password (property field)
	WithOptionalPassword(BACnetContextTagCharacterString) BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder
	// WithOptionalPasswordBuilder adds Password (property field) which is build by the builder
	WithOptionalPasswordBuilder(func(BACnetContextTagCharacterStringBuilder) BACnetContextTagCharacterStringBuilder) BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConfirmedServiceRequestBuilder
	// Build builds the BACnetConfirmedServiceRequestDeviceCommunicationControl or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestDeviceCommunicationControl, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestDeviceCommunicationControl
}

// NewBACnetConfirmedServiceRequestDeviceCommunicationControlBuilder() creates a BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder
func NewBACnetConfirmedServiceRequestDeviceCommunicationControlBuilder() BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder {
	return &_BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder{_BACnetConfirmedServiceRequestDeviceCommunicationControl: new(_BACnetConfirmedServiceRequestDeviceCommunicationControl)}
}

type _BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder struct {
	*_BACnetConfirmedServiceRequestDeviceCommunicationControl

	parentBuilder *_BACnetConfirmedServiceRequestBuilder

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder) = (*_BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder)(nil)

func (b *_BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder) setParent(contract BACnetConfirmedServiceRequestContract) {
	b.BACnetConfirmedServiceRequestContract = contract
	contract.(*_BACnetConfirmedServiceRequest)._SubType = b._BACnetConfirmedServiceRequestDeviceCommunicationControl
}

func (b *_BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder) WithMandatoryFields(enableDisable BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged) BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder {
	return b.WithEnableDisable(enableDisable)
}

func (b *_BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder) WithOptionalTimeDuration(timeDuration BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder {
	b.TimeDuration = timeDuration
	return b
}

func (b *_BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder) WithOptionalTimeDurationBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder {
	builder := builderSupplier(b.TimeDuration.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.TimeDuration, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder) WithEnableDisable(enableDisable BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged) BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder {
	b.EnableDisable = enableDisable
	return b
}

func (b *_BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder) WithEnableDisableBuilder(builderSupplier func(BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTaggedBuilder) BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTaggedBuilder) BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder {
	builder := builderSupplier(b.EnableDisable.CreateBACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTaggedBuilder())
	var err error
	b.EnableDisable, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder) WithOptionalPassword(password BACnetContextTagCharacterString) BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder {
	b.Password = password
	return b
}

func (b *_BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder) WithOptionalPasswordBuilder(builderSupplier func(BACnetContextTagCharacterStringBuilder) BACnetContextTagCharacterStringBuilder) BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder {
	builder := builderSupplier(b.Password.CreateBACnetContextTagCharacterStringBuilder())
	var err error
	b.Password, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagCharacterStringBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder) Build() (BACnetConfirmedServiceRequestDeviceCommunicationControl, error) {
	if b.EnableDisable == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'enableDisable' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConfirmedServiceRequestDeviceCommunicationControl.deepCopy(), nil
}

func (b *_BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder) MustBuild() BACnetConfirmedServiceRequestDeviceCommunicationControl {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder) Done() BACnetConfirmedServiceRequestBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConfirmedServiceRequestBuilder().(*_BACnetConfirmedServiceRequestBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder) buildForBACnetConfirmedServiceRequest() (BACnetConfirmedServiceRequest, error) {
	return b.Build()
}

func (b *_BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConfirmedServiceRequestDeviceCommunicationControlBuilder().(*_BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConfirmedServiceRequestDeviceCommunicationControlBuilder creates a BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder
func (b *_BACnetConfirmedServiceRequestDeviceCommunicationControl) CreateBACnetConfirmedServiceRequestDeviceCommunicationControlBuilder() BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder {
	if b == nil {
		return NewBACnetConfirmedServiceRequestDeviceCommunicationControlBuilder()
	}
	return &_BACnetConfirmedServiceRequestDeviceCommunicationControlBuilder{_BACnetConfirmedServiceRequestDeviceCommunicationControl: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_DEVICE_COMMUNICATION_CONTROL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) GetTimeDuration() BACnetContextTagUnsignedInteger {
	return m.TimeDuration
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) GetEnableDisable() BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged {
	return m.EnableDisable
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) GetPassword() BACnetContextTagCharacterString {
	return m.Password
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestDeviceCommunicationControl(structType any) BACnetConfirmedServiceRequestDeviceCommunicationControl {
	if casted, ok := structType.(BACnetConfirmedServiceRequestDeviceCommunicationControl); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestDeviceCommunicationControl); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) GetTypeName() string {
	return "BACnetConfirmedServiceRequestDeviceCommunicationControl"
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Optional Field (timeDuration)
	if m.TimeDuration != nil {
		lengthInBits += m.TimeDuration.GetLengthInBits(ctx)
	}

	// Simple field (enableDisable)
	lengthInBits += m.EnableDisable.GetLengthInBits(ctx)

	// Optional Field (password)
	if m.Password != nil {
		lengthInBits += m.Password.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestDeviceCommunicationControl BACnetConfirmedServiceRequestDeviceCommunicationControl, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestDeviceCommunicationControl"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestDeviceCommunicationControl")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var timeDuration BACnetContextTagUnsignedInteger
	_timeDuration, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "timeDuration", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeDuration' field"))
	}
	if _timeDuration != nil {
		timeDuration = *_timeDuration
		m.TimeDuration = timeDuration
	}

	enableDisable, err := ReadSimpleField[BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged](ctx, "enableDisable", ReadComplex[BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged](BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTaggedParseWithBufferProducer((uint8)(uint8(1)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'enableDisable' field"))
	}
	m.EnableDisable = enableDisable

	var password BACnetContextTagCharacterString
	_password, err := ReadOptionalField[BACnetContextTagCharacterString](ctx, "password", ReadComplex[BACnetContextTagCharacterString](BACnetContextTagParseWithBufferProducer[BACnetContextTagCharacterString]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_CHARACTER_STRING)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'password' field"))
	}
	if _password != nil {
		password = *_password
		m.Password = password
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestDeviceCommunicationControl"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestDeviceCommunicationControl")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestDeviceCommunicationControl"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestDeviceCommunicationControl")
		}

		if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "timeDuration", GetRef(m.GetTimeDuration()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'timeDuration' field")
		}

		if err := WriteSimpleField[BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged](ctx, "enableDisable", m.GetEnableDisable(), WriteComplex[BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'enableDisable' field")
		}

		if err := WriteOptionalField[BACnetContextTagCharacterString](ctx, "password", GetRef(m.GetPassword()), WriteComplex[BACnetContextTagCharacterString](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'password' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestDeviceCommunicationControl"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestDeviceCommunicationControl")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) IsBACnetConfirmedServiceRequestDeviceCommunicationControl() {
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) deepCopy() *_BACnetConfirmedServiceRequestDeviceCommunicationControl {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestDeviceCommunicationControlCopy := &_BACnetConfirmedServiceRequestDeviceCommunicationControl{
		m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).deepCopy(),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.TimeDuration),
		utils.DeepCopy[BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged](m.EnableDisable),
		utils.DeepCopy[BACnetContextTagCharacterString](m.Password),
	}
	_BACnetConfirmedServiceRequestDeviceCommunicationControlCopy.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = m
	return _BACnetConfirmedServiceRequestDeviceCommunicationControlCopy
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) String() string {
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
