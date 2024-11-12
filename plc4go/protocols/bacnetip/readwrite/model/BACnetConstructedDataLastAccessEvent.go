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

// BACnetConstructedDataLastAccessEvent is the corresponding interface of BACnetConstructedDataLastAccessEvent
type BACnetConstructedDataLastAccessEvent interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetLastAccessEvent returns LastAccessEvent (property field)
	GetLastAccessEvent() BACnetAccessEventTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetAccessEventTagged
	// IsBACnetConstructedDataLastAccessEvent is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLastAccessEvent()
	// CreateBuilder creates a BACnetConstructedDataLastAccessEventBuilder
	CreateBACnetConstructedDataLastAccessEventBuilder() BACnetConstructedDataLastAccessEventBuilder
}

// _BACnetConstructedDataLastAccessEvent is the data-structure of this message
type _BACnetConstructedDataLastAccessEvent struct {
	BACnetConstructedDataContract
	LastAccessEvent BACnetAccessEventTagged
}

var _ BACnetConstructedDataLastAccessEvent = (*_BACnetConstructedDataLastAccessEvent)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLastAccessEvent)(nil)

// NewBACnetConstructedDataLastAccessEvent factory function for _BACnetConstructedDataLastAccessEvent
func NewBACnetConstructedDataLastAccessEvent(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, lastAccessEvent BACnetAccessEventTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLastAccessEvent {
	if lastAccessEvent == nil {
		panic("lastAccessEvent of type BACnetAccessEventTagged for BACnetConstructedDataLastAccessEvent must not be nil")
	}
	_result := &_BACnetConstructedDataLastAccessEvent{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LastAccessEvent:               lastAccessEvent,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLastAccessEventBuilder is a builder for BACnetConstructedDataLastAccessEvent
type BACnetConstructedDataLastAccessEventBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(lastAccessEvent BACnetAccessEventTagged) BACnetConstructedDataLastAccessEventBuilder
	// WithLastAccessEvent adds LastAccessEvent (property field)
	WithLastAccessEvent(BACnetAccessEventTagged) BACnetConstructedDataLastAccessEventBuilder
	// WithLastAccessEventBuilder adds LastAccessEvent (property field) which is build by the builder
	WithLastAccessEventBuilder(func(BACnetAccessEventTaggedBuilder) BACnetAccessEventTaggedBuilder) BACnetConstructedDataLastAccessEventBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataLastAccessEvent or returns an error if something is wrong
	Build() (BACnetConstructedDataLastAccessEvent, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLastAccessEvent
}

// NewBACnetConstructedDataLastAccessEventBuilder() creates a BACnetConstructedDataLastAccessEventBuilder
func NewBACnetConstructedDataLastAccessEventBuilder() BACnetConstructedDataLastAccessEventBuilder {
	return &_BACnetConstructedDataLastAccessEventBuilder{_BACnetConstructedDataLastAccessEvent: new(_BACnetConstructedDataLastAccessEvent)}
}

type _BACnetConstructedDataLastAccessEventBuilder struct {
	*_BACnetConstructedDataLastAccessEvent

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLastAccessEventBuilder) = (*_BACnetConstructedDataLastAccessEventBuilder)(nil)

func (b *_BACnetConstructedDataLastAccessEventBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataLastAccessEvent
}

func (b *_BACnetConstructedDataLastAccessEventBuilder) WithMandatoryFields(lastAccessEvent BACnetAccessEventTagged) BACnetConstructedDataLastAccessEventBuilder {
	return b.WithLastAccessEvent(lastAccessEvent)
}

func (b *_BACnetConstructedDataLastAccessEventBuilder) WithLastAccessEvent(lastAccessEvent BACnetAccessEventTagged) BACnetConstructedDataLastAccessEventBuilder {
	b.LastAccessEvent = lastAccessEvent
	return b
}

func (b *_BACnetConstructedDataLastAccessEventBuilder) WithLastAccessEventBuilder(builderSupplier func(BACnetAccessEventTaggedBuilder) BACnetAccessEventTaggedBuilder) BACnetConstructedDataLastAccessEventBuilder {
	builder := builderSupplier(b.LastAccessEvent.CreateBACnetAccessEventTaggedBuilder())
	var err error
	b.LastAccessEvent, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetAccessEventTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataLastAccessEventBuilder) Build() (BACnetConstructedDataLastAccessEvent, error) {
	if b.LastAccessEvent == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'lastAccessEvent' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLastAccessEvent.deepCopy(), nil
}

func (b *_BACnetConstructedDataLastAccessEventBuilder) MustBuild() BACnetConstructedDataLastAccessEvent {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataLastAccessEventBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLastAccessEventBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLastAccessEventBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLastAccessEventBuilder().(*_BACnetConstructedDataLastAccessEventBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLastAccessEventBuilder creates a BACnetConstructedDataLastAccessEventBuilder
func (b *_BACnetConstructedDataLastAccessEvent) CreateBACnetConstructedDataLastAccessEventBuilder() BACnetConstructedDataLastAccessEventBuilder {
	if b == nil {
		return NewBACnetConstructedDataLastAccessEventBuilder()
	}
	return &_BACnetConstructedDataLastAccessEventBuilder{_BACnetConstructedDataLastAccessEvent: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLastAccessEvent) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLastAccessEvent) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LAST_ACCESS_EVENT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLastAccessEvent) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLastAccessEvent) GetLastAccessEvent() BACnetAccessEventTagged {
	return m.LastAccessEvent
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLastAccessEvent) GetActualValue() BACnetAccessEventTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetAccessEventTagged(m.GetLastAccessEvent())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLastAccessEvent(structType any) BACnetConstructedDataLastAccessEvent {
	if casted, ok := structType.(BACnetConstructedDataLastAccessEvent); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLastAccessEvent); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLastAccessEvent) GetTypeName() string {
	return "BACnetConstructedDataLastAccessEvent"
}

func (m *_BACnetConstructedDataLastAccessEvent) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (lastAccessEvent)
	lengthInBits += m.LastAccessEvent.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLastAccessEvent) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLastAccessEvent) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLastAccessEvent BACnetConstructedDataLastAccessEvent, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLastAccessEvent"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLastAccessEvent")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lastAccessEvent, err := ReadSimpleField[BACnetAccessEventTagged](ctx, "lastAccessEvent", ReadComplex[BACnetAccessEventTagged](BACnetAccessEventTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastAccessEvent' field"))
	}
	m.LastAccessEvent = lastAccessEvent

	actualValue, err := ReadVirtualField[BACnetAccessEventTagged](ctx, "actualValue", (*BACnetAccessEventTagged)(nil), lastAccessEvent)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLastAccessEvent"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLastAccessEvent")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLastAccessEvent) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLastAccessEvent) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLastAccessEvent"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLastAccessEvent")
		}

		if err := WriteSimpleField[BACnetAccessEventTagged](ctx, "lastAccessEvent", m.GetLastAccessEvent(), WriteComplex[BACnetAccessEventTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lastAccessEvent' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLastAccessEvent"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLastAccessEvent")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLastAccessEvent) IsBACnetConstructedDataLastAccessEvent() {}

func (m *_BACnetConstructedDataLastAccessEvent) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLastAccessEvent) deepCopy() *_BACnetConstructedDataLastAccessEvent {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLastAccessEventCopy := &_BACnetConstructedDataLastAccessEvent{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetAccessEventTagged](m.LastAccessEvent),
	}
	_BACnetConstructedDataLastAccessEventCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLastAccessEventCopy
}

func (m *_BACnetConstructedDataLastAccessEvent) String() string {
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
