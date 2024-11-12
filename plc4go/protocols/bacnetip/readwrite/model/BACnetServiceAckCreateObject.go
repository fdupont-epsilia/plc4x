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

// BACnetServiceAckCreateObject is the corresponding interface of BACnetServiceAckCreateObject
type BACnetServiceAckCreateObject interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetServiceAck
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetApplicationTagObjectIdentifier
	// IsBACnetServiceAckCreateObject is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetServiceAckCreateObject()
	// CreateBuilder creates a BACnetServiceAckCreateObjectBuilder
	CreateBACnetServiceAckCreateObjectBuilder() BACnetServiceAckCreateObjectBuilder
}

// _BACnetServiceAckCreateObject is the data-structure of this message
type _BACnetServiceAckCreateObject struct {
	BACnetServiceAckContract
	ObjectIdentifier BACnetApplicationTagObjectIdentifier
}

var _ BACnetServiceAckCreateObject = (*_BACnetServiceAckCreateObject)(nil)
var _ BACnetServiceAckRequirements = (*_BACnetServiceAckCreateObject)(nil)

// NewBACnetServiceAckCreateObject factory function for _BACnetServiceAckCreateObject
func NewBACnetServiceAckCreateObject(objectIdentifier BACnetApplicationTagObjectIdentifier, serviceAckLength uint32) *_BACnetServiceAckCreateObject {
	if objectIdentifier == nil {
		panic("objectIdentifier of type BACnetApplicationTagObjectIdentifier for BACnetServiceAckCreateObject must not be nil")
	}
	_result := &_BACnetServiceAckCreateObject{
		BACnetServiceAckContract: NewBACnetServiceAck(serviceAckLength),
		ObjectIdentifier:         objectIdentifier,
	}
	_result.BACnetServiceAckContract.(*_BACnetServiceAck)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetServiceAckCreateObjectBuilder is a builder for BACnetServiceAckCreateObject
type BACnetServiceAckCreateObjectBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(objectIdentifier BACnetApplicationTagObjectIdentifier) BACnetServiceAckCreateObjectBuilder
	// WithObjectIdentifier adds ObjectIdentifier (property field)
	WithObjectIdentifier(BACnetApplicationTagObjectIdentifier) BACnetServiceAckCreateObjectBuilder
	// WithObjectIdentifierBuilder adds ObjectIdentifier (property field) which is build by the builder
	WithObjectIdentifierBuilder(func(BACnetApplicationTagObjectIdentifierBuilder) BACnetApplicationTagObjectIdentifierBuilder) BACnetServiceAckCreateObjectBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetServiceAckBuilder
	// Build builds the BACnetServiceAckCreateObject or returns an error if something is wrong
	Build() (BACnetServiceAckCreateObject, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetServiceAckCreateObject
}

// NewBACnetServiceAckCreateObjectBuilder() creates a BACnetServiceAckCreateObjectBuilder
func NewBACnetServiceAckCreateObjectBuilder() BACnetServiceAckCreateObjectBuilder {
	return &_BACnetServiceAckCreateObjectBuilder{_BACnetServiceAckCreateObject: new(_BACnetServiceAckCreateObject)}
}

type _BACnetServiceAckCreateObjectBuilder struct {
	*_BACnetServiceAckCreateObject

	parentBuilder *_BACnetServiceAckBuilder

	err *utils.MultiError
}

var _ (BACnetServiceAckCreateObjectBuilder) = (*_BACnetServiceAckCreateObjectBuilder)(nil)

func (b *_BACnetServiceAckCreateObjectBuilder) setParent(contract BACnetServiceAckContract) {
	b.BACnetServiceAckContract = contract
	contract.(*_BACnetServiceAck)._SubType = b._BACnetServiceAckCreateObject
}

func (b *_BACnetServiceAckCreateObjectBuilder) WithMandatoryFields(objectIdentifier BACnetApplicationTagObjectIdentifier) BACnetServiceAckCreateObjectBuilder {
	return b.WithObjectIdentifier(objectIdentifier)
}

func (b *_BACnetServiceAckCreateObjectBuilder) WithObjectIdentifier(objectIdentifier BACnetApplicationTagObjectIdentifier) BACnetServiceAckCreateObjectBuilder {
	b.ObjectIdentifier = objectIdentifier
	return b
}

func (b *_BACnetServiceAckCreateObjectBuilder) WithObjectIdentifierBuilder(builderSupplier func(BACnetApplicationTagObjectIdentifierBuilder) BACnetApplicationTagObjectIdentifierBuilder) BACnetServiceAckCreateObjectBuilder {
	builder := builderSupplier(b.ObjectIdentifier.CreateBACnetApplicationTagObjectIdentifierBuilder())
	var err error
	b.ObjectIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagObjectIdentifierBuilder failed"))
	}
	return b
}

func (b *_BACnetServiceAckCreateObjectBuilder) Build() (BACnetServiceAckCreateObject, error) {
	if b.ObjectIdentifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'objectIdentifier' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetServiceAckCreateObject.deepCopy(), nil
}

func (b *_BACnetServiceAckCreateObjectBuilder) MustBuild() BACnetServiceAckCreateObject {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetServiceAckCreateObjectBuilder) Done() BACnetServiceAckBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetServiceAckBuilder().(*_BACnetServiceAckBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetServiceAckCreateObjectBuilder) buildForBACnetServiceAck() (BACnetServiceAck, error) {
	return b.Build()
}

func (b *_BACnetServiceAckCreateObjectBuilder) DeepCopy() any {
	_copy := b.CreateBACnetServiceAckCreateObjectBuilder().(*_BACnetServiceAckCreateObjectBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetServiceAckCreateObjectBuilder creates a BACnetServiceAckCreateObjectBuilder
func (b *_BACnetServiceAckCreateObject) CreateBACnetServiceAckCreateObjectBuilder() BACnetServiceAckCreateObjectBuilder {
	if b == nil {
		return NewBACnetServiceAckCreateObjectBuilder()
	}
	return &_BACnetServiceAckCreateObjectBuilder{_BACnetServiceAckCreateObject: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetServiceAckCreateObject) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_CREATE_OBJECT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckCreateObject) GetParent() BACnetServiceAckContract {
	return m.BACnetServiceAckContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckCreateObject) GetObjectIdentifier() BACnetApplicationTagObjectIdentifier {
	return m.ObjectIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckCreateObject(structType any) BACnetServiceAckCreateObject {
	if casted, ok := structType.(BACnetServiceAckCreateObject); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckCreateObject); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckCreateObject) GetTypeName() string {
	return "BACnetServiceAckCreateObject"
}

func (m *_BACnetServiceAckCreateObject) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetServiceAckContract.(*_BACnetServiceAck).getLengthInBits(ctx))

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetServiceAckCreateObject) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetServiceAckCreateObject) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetServiceAck, serviceAckLength uint32) (__bACnetServiceAckCreateObject BACnetServiceAckCreateObject, err error) {
	m.BACnetServiceAckContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckCreateObject"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckCreateObject")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectIdentifier, err := ReadSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "objectIdentifier", ReadComplex[BACnetApplicationTagObjectIdentifier](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagObjectIdentifier](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectIdentifier' field"))
	}
	m.ObjectIdentifier = objectIdentifier

	if closeErr := readBuffer.CloseContext("BACnetServiceAckCreateObject"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckCreateObject")
	}

	return m, nil
}

func (m *_BACnetServiceAckCreateObject) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetServiceAckCreateObject) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckCreateObject"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckCreateObject")
		}

		if err := WriteSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "objectIdentifier", m.GetObjectIdentifier(), WriteComplex[BACnetApplicationTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'objectIdentifier' field")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckCreateObject"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckCreateObject")
		}
		return nil
	}
	return m.BACnetServiceAckContract.(*_BACnetServiceAck).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetServiceAckCreateObject) IsBACnetServiceAckCreateObject() {}

func (m *_BACnetServiceAckCreateObject) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetServiceAckCreateObject) deepCopy() *_BACnetServiceAckCreateObject {
	if m == nil {
		return nil
	}
	_BACnetServiceAckCreateObjectCopy := &_BACnetServiceAckCreateObject{
		m.BACnetServiceAckContract.(*_BACnetServiceAck).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagObjectIdentifier](m.ObjectIdentifier),
	}
	_BACnetServiceAckCreateObjectCopy.BACnetServiceAckContract.(*_BACnetServiceAck)._SubType = m
	return _BACnetServiceAckCreateObjectCopy
}

func (m *_BACnetServiceAckCreateObject) String() string {
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
