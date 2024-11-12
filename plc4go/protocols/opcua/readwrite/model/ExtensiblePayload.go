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

// ExtensiblePayload is the corresponding interface of ExtensiblePayload
type ExtensiblePayload interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	Payload
	// GetPayload returns Payload (property field)
	GetPayload() RootExtensionObject
	// IsExtensiblePayload is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsExtensiblePayload()
	// CreateBuilder creates a ExtensiblePayloadBuilder
	CreateExtensiblePayloadBuilder() ExtensiblePayloadBuilder
}

// _ExtensiblePayload is the data-structure of this message
type _ExtensiblePayload struct {
	PayloadContract
	Payload RootExtensionObject
}

var _ ExtensiblePayload = (*_ExtensiblePayload)(nil)
var _ PayloadRequirements = (*_ExtensiblePayload)(nil)

// NewExtensiblePayload factory function for _ExtensiblePayload
func NewExtensiblePayload(sequenceHeader SequenceHeader, payload RootExtensionObject, byteCount uint32) *_ExtensiblePayload {
	if payload == nil {
		panic("payload of type RootExtensionObject for ExtensiblePayload must not be nil")
	}
	_result := &_ExtensiblePayload{
		PayloadContract: NewPayload(sequenceHeader, byteCount),
		Payload:         payload,
	}
	_result.PayloadContract.(*_Payload)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ExtensiblePayloadBuilder is a builder for ExtensiblePayload
type ExtensiblePayloadBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(payload RootExtensionObject) ExtensiblePayloadBuilder
	// WithPayload adds Payload (property field)
	WithPayload(RootExtensionObject) ExtensiblePayloadBuilder
	// WithPayloadBuilder adds Payload (property field) which is build by the builder
	WithPayloadBuilder(func(RootExtensionObjectBuilder) RootExtensionObjectBuilder) ExtensiblePayloadBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() PayloadBuilder
	// Build builds the ExtensiblePayload or returns an error if something is wrong
	Build() (ExtensiblePayload, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ExtensiblePayload
}

// NewExtensiblePayloadBuilder() creates a ExtensiblePayloadBuilder
func NewExtensiblePayloadBuilder() ExtensiblePayloadBuilder {
	return &_ExtensiblePayloadBuilder{_ExtensiblePayload: new(_ExtensiblePayload)}
}

type _ExtensiblePayloadBuilder struct {
	*_ExtensiblePayload

	parentBuilder *_PayloadBuilder

	err *utils.MultiError
}

var _ (ExtensiblePayloadBuilder) = (*_ExtensiblePayloadBuilder)(nil)

func (b *_ExtensiblePayloadBuilder) setParent(contract PayloadContract) {
	b.PayloadContract = contract
	contract.(*_Payload)._SubType = b._ExtensiblePayload
}

func (b *_ExtensiblePayloadBuilder) WithMandatoryFields(payload RootExtensionObject) ExtensiblePayloadBuilder {
	return b.WithPayload(payload)
}

func (b *_ExtensiblePayloadBuilder) WithPayload(payload RootExtensionObject) ExtensiblePayloadBuilder {
	b.Payload = payload
	return b
}

func (b *_ExtensiblePayloadBuilder) WithPayloadBuilder(builderSupplier func(RootExtensionObjectBuilder) RootExtensionObjectBuilder) ExtensiblePayloadBuilder {
	builder := builderSupplier(b.Payload.CreateRootExtensionObjectBuilder())
	var err error
	b.Payload, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "RootExtensionObjectBuilder failed"))
	}
	return b
}

func (b *_ExtensiblePayloadBuilder) Build() (ExtensiblePayload, error) {
	if b.Payload == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'payload' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ExtensiblePayload.deepCopy(), nil
}

func (b *_ExtensiblePayloadBuilder) MustBuild() ExtensiblePayload {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ExtensiblePayloadBuilder) Done() PayloadBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewPayloadBuilder().(*_PayloadBuilder)
	}
	return b.parentBuilder
}

func (b *_ExtensiblePayloadBuilder) buildForPayload() (Payload, error) {
	return b.Build()
}

func (b *_ExtensiblePayloadBuilder) DeepCopy() any {
	_copy := b.CreateExtensiblePayloadBuilder().(*_ExtensiblePayloadBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateExtensiblePayloadBuilder creates a ExtensiblePayloadBuilder
func (b *_ExtensiblePayload) CreateExtensiblePayloadBuilder() ExtensiblePayloadBuilder {
	if b == nil {
		return NewExtensiblePayloadBuilder()
	}
	return &_ExtensiblePayloadBuilder{_ExtensiblePayload: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ExtensiblePayload) GetBinary() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ExtensiblePayload) GetParent() PayloadContract {
	return m.PayloadContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ExtensiblePayload) GetPayload() RootExtensionObject {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastExtensiblePayload(structType any) ExtensiblePayload {
	if casted, ok := structType.(ExtensiblePayload); ok {
		return casted
	}
	if casted, ok := structType.(*ExtensiblePayload); ok {
		return *casted
	}
	return nil
}

func (m *_ExtensiblePayload) GetTypeName() string {
	return "ExtensiblePayload"
}

func (m *_ExtensiblePayload) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.PayloadContract.(*_Payload).getLengthInBits(ctx))

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ExtensiblePayload) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ExtensiblePayload) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_Payload, binary bool, byteCount uint32) (__extensiblePayload ExtensiblePayload, err error) {
	m.PayloadContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ExtensiblePayload"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ExtensiblePayload")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	payload, err := ReadSimpleField[RootExtensionObject](ctx, "payload", ReadComplex[RootExtensionObject](ExtensionObjectParseWithBufferProducer[RootExtensionObject]((bool)(bool(false))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	if closeErr := readBuffer.CloseContext("ExtensiblePayload"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ExtensiblePayload")
	}

	return m, nil
}

func (m *_ExtensiblePayload) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ExtensiblePayload) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ExtensiblePayload"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ExtensiblePayload")
		}

		if err := WriteSimpleField[RootExtensionObject](ctx, "payload", m.GetPayload(), WriteComplex[RootExtensionObject](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'payload' field")
		}

		if popErr := writeBuffer.PopContext("ExtensiblePayload"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ExtensiblePayload")
		}
		return nil
	}
	return m.PayloadContract.(*_Payload).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ExtensiblePayload) IsExtensiblePayload() {}

func (m *_ExtensiblePayload) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ExtensiblePayload) deepCopy() *_ExtensiblePayload {
	if m == nil {
		return nil
	}
	_ExtensiblePayloadCopy := &_ExtensiblePayload{
		m.PayloadContract.(*_Payload).deepCopy(),
		utils.DeepCopy[RootExtensionObject](m.Payload),
	}
	_ExtensiblePayloadCopy.PayloadContract.(*_Payload)._SubType = m
	return _ExtensiblePayloadCopy
}

func (m *_ExtensiblePayload) String() string {
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
