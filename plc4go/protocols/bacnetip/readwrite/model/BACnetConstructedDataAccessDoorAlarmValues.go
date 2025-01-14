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

// BACnetConstructedDataAccessDoorAlarmValues is the corresponding interface of BACnetConstructedDataAccessDoorAlarmValues
type BACnetConstructedDataAccessDoorAlarmValues interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetAlarmValues returns AlarmValues (property field)
	GetAlarmValues() []BACnetDoorAlarmStateTagged
	// IsBACnetConstructedDataAccessDoorAlarmValues is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAccessDoorAlarmValues()
	// CreateBuilder creates a BACnetConstructedDataAccessDoorAlarmValuesBuilder
	CreateBACnetConstructedDataAccessDoorAlarmValuesBuilder() BACnetConstructedDataAccessDoorAlarmValuesBuilder
}

// _BACnetConstructedDataAccessDoorAlarmValues is the data-structure of this message
type _BACnetConstructedDataAccessDoorAlarmValues struct {
	BACnetConstructedDataContract
	AlarmValues []BACnetDoorAlarmStateTagged
}

var _ BACnetConstructedDataAccessDoorAlarmValues = (*_BACnetConstructedDataAccessDoorAlarmValues)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAccessDoorAlarmValues)(nil)

// NewBACnetConstructedDataAccessDoorAlarmValues factory function for _BACnetConstructedDataAccessDoorAlarmValues
func NewBACnetConstructedDataAccessDoorAlarmValues(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, alarmValues []BACnetDoorAlarmStateTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccessDoorAlarmValues {
	_result := &_BACnetConstructedDataAccessDoorAlarmValues{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		AlarmValues:                   alarmValues,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAccessDoorAlarmValuesBuilder is a builder for BACnetConstructedDataAccessDoorAlarmValues
type BACnetConstructedDataAccessDoorAlarmValuesBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(alarmValues []BACnetDoorAlarmStateTagged) BACnetConstructedDataAccessDoorAlarmValuesBuilder
	// WithAlarmValues adds AlarmValues (property field)
	WithAlarmValues(...BACnetDoorAlarmStateTagged) BACnetConstructedDataAccessDoorAlarmValuesBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataAccessDoorAlarmValues or returns an error if something is wrong
	Build() (BACnetConstructedDataAccessDoorAlarmValues, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAccessDoorAlarmValues
}

// NewBACnetConstructedDataAccessDoorAlarmValuesBuilder() creates a BACnetConstructedDataAccessDoorAlarmValuesBuilder
func NewBACnetConstructedDataAccessDoorAlarmValuesBuilder() BACnetConstructedDataAccessDoorAlarmValuesBuilder {
	return &_BACnetConstructedDataAccessDoorAlarmValuesBuilder{_BACnetConstructedDataAccessDoorAlarmValues: new(_BACnetConstructedDataAccessDoorAlarmValues)}
}

type _BACnetConstructedDataAccessDoorAlarmValuesBuilder struct {
	*_BACnetConstructedDataAccessDoorAlarmValues

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataAccessDoorAlarmValuesBuilder) = (*_BACnetConstructedDataAccessDoorAlarmValuesBuilder)(nil)

func (b *_BACnetConstructedDataAccessDoorAlarmValuesBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataAccessDoorAlarmValues
}

func (b *_BACnetConstructedDataAccessDoorAlarmValuesBuilder) WithMandatoryFields(alarmValues []BACnetDoorAlarmStateTagged) BACnetConstructedDataAccessDoorAlarmValuesBuilder {
	return b.WithAlarmValues(alarmValues...)
}

func (b *_BACnetConstructedDataAccessDoorAlarmValuesBuilder) WithAlarmValues(alarmValues ...BACnetDoorAlarmStateTagged) BACnetConstructedDataAccessDoorAlarmValuesBuilder {
	b.AlarmValues = alarmValues
	return b
}

func (b *_BACnetConstructedDataAccessDoorAlarmValuesBuilder) Build() (BACnetConstructedDataAccessDoorAlarmValues, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataAccessDoorAlarmValues.deepCopy(), nil
}

func (b *_BACnetConstructedDataAccessDoorAlarmValuesBuilder) MustBuild() BACnetConstructedDataAccessDoorAlarmValues {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataAccessDoorAlarmValuesBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataAccessDoorAlarmValuesBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataAccessDoorAlarmValuesBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataAccessDoorAlarmValuesBuilder().(*_BACnetConstructedDataAccessDoorAlarmValuesBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataAccessDoorAlarmValuesBuilder creates a BACnetConstructedDataAccessDoorAlarmValuesBuilder
func (b *_BACnetConstructedDataAccessDoorAlarmValues) CreateBACnetConstructedDataAccessDoorAlarmValuesBuilder() BACnetConstructedDataAccessDoorAlarmValuesBuilder {
	if b == nil {
		return NewBACnetConstructedDataAccessDoorAlarmValuesBuilder()
	}
	return &_BACnetConstructedDataAccessDoorAlarmValuesBuilder{_BACnetConstructedDataAccessDoorAlarmValues: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccessDoorAlarmValues) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ACCESS_DOOR
}

func (m *_BACnetConstructedDataAccessDoorAlarmValues) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALARM_VALUES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccessDoorAlarmValues) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccessDoorAlarmValues) GetAlarmValues() []BACnetDoorAlarmStateTagged {
	return m.AlarmValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccessDoorAlarmValues(structType any) BACnetConstructedDataAccessDoorAlarmValues {
	if casted, ok := structType.(BACnetConstructedDataAccessDoorAlarmValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccessDoorAlarmValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccessDoorAlarmValues) GetTypeName() string {
	return "BACnetConstructedDataAccessDoorAlarmValues"
}

func (m *_BACnetConstructedDataAccessDoorAlarmValues) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.AlarmValues) > 0 {
		for _, element := range m.AlarmValues {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataAccessDoorAlarmValues) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAccessDoorAlarmValues) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAccessDoorAlarmValues BACnetConstructedDataAccessDoorAlarmValues, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccessDoorAlarmValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccessDoorAlarmValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	alarmValues, err := ReadTerminatedArrayField[BACnetDoorAlarmStateTagged](ctx, "alarmValues", ReadComplex[BACnetDoorAlarmStateTagged](BACnetDoorAlarmStateTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'alarmValues' field"))
	}
	m.AlarmValues = alarmValues

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccessDoorAlarmValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccessDoorAlarmValues")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAccessDoorAlarmValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccessDoorAlarmValues) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccessDoorAlarmValues"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccessDoorAlarmValues")
		}

		if err := WriteComplexTypeArrayField(ctx, "alarmValues", m.GetAlarmValues(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'alarmValues' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccessDoorAlarmValues"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccessDoorAlarmValues")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccessDoorAlarmValues) IsBACnetConstructedDataAccessDoorAlarmValues() {
}

func (m *_BACnetConstructedDataAccessDoorAlarmValues) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAccessDoorAlarmValues) deepCopy() *_BACnetConstructedDataAccessDoorAlarmValues {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAccessDoorAlarmValuesCopy := &_BACnetConstructedDataAccessDoorAlarmValues{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopySlice[BACnetDoorAlarmStateTagged, BACnetDoorAlarmStateTagged](m.AlarmValues),
	}
	_BACnetConstructedDataAccessDoorAlarmValuesCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAccessDoorAlarmValuesCopy
}

func (m *_BACnetConstructedDataAccessDoorAlarmValues) String() string {
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
