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

// BACnetConstructedDataBackupPreparationTime is the corresponding interface of BACnetConstructedDataBackupPreparationTime
type BACnetConstructedDataBackupPreparationTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetBackupPreparationTime returns BackupPreparationTime (property field)
	GetBackupPreparationTime() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataBackupPreparationTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataBackupPreparationTime()
	// CreateBuilder creates a BACnetConstructedDataBackupPreparationTimeBuilder
	CreateBACnetConstructedDataBackupPreparationTimeBuilder() BACnetConstructedDataBackupPreparationTimeBuilder
}

// _BACnetConstructedDataBackupPreparationTime is the data-structure of this message
type _BACnetConstructedDataBackupPreparationTime struct {
	BACnetConstructedDataContract
	BackupPreparationTime BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataBackupPreparationTime = (*_BACnetConstructedDataBackupPreparationTime)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataBackupPreparationTime)(nil)

// NewBACnetConstructedDataBackupPreparationTime factory function for _BACnetConstructedDataBackupPreparationTime
func NewBACnetConstructedDataBackupPreparationTime(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, backupPreparationTime BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBackupPreparationTime {
	if backupPreparationTime == nil {
		panic("backupPreparationTime of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataBackupPreparationTime must not be nil")
	}
	_result := &_BACnetConstructedDataBackupPreparationTime{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		BackupPreparationTime:         backupPreparationTime,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataBackupPreparationTimeBuilder is a builder for BACnetConstructedDataBackupPreparationTime
type BACnetConstructedDataBackupPreparationTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(backupPreparationTime BACnetApplicationTagUnsignedInteger) BACnetConstructedDataBackupPreparationTimeBuilder
	// WithBackupPreparationTime adds BackupPreparationTime (property field)
	WithBackupPreparationTime(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataBackupPreparationTimeBuilder
	// WithBackupPreparationTimeBuilder adds BackupPreparationTime (property field) which is build by the builder
	WithBackupPreparationTimeBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataBackupPreparationTimeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataBackupPreparationTime or returns an error if something is wrong
	Build() (BACnetConstructedDataBackupPreparationTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataBackupPreparationTime
}

// NewBACnetConstructedDataBackupPreparationTimeBuilder() creates a BACnetConstructedDataBackupPreparationTimeBuilder
func NewBACnetConstructedDataBackupPreparationTimeBuilder() BACnetConstructedDataBackupPreparationTimeBuilder {
	return &_BACnetConstructedDataBackupPreparationTimeBuilder{_BACnetConstructedDataBackupPreparationTime: new(_BACnetConstructedDataBackupPreparationTime)}
}

type _BACnetConstructedDataBackupPreparationTimeBuilder struct {
	*_BACnetConstructedDataBackupPreparationTime

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataBackupPreparationTimeBuilder) = (*_BACnetConstructedDataBackupPreparationTimeBuilder)(nil)

func (b *_BACnetConstructedDataBackupPreparationTimeBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataBackupPreparationTime
}

func (b *_BACnetConstructedDataBackupPreparationTimeBuilder) WithMandatoryFields(backupPreparationTime BACnetApplicationTagUnsignedInteger) BACnetConstructedDataBackupPreparationTimeBuilder {
	return b.WithBackupPreparationTime(backupPreparationTime)
}

func (b *_BACnetConstructedDataBackupPreparationTimeBuilder) WithBackupPreparationTime(backupPreparationTime BACnetApplicationTagUnsignedInteger) BACnetConstructedDataBackupPreparationTimeBuilder {
	b.BackupPreparationTime = backupPreparationTime
	return b
}

func (b *_BACnetConstructedDataBackupPreparationTimeBuilder) WithBackupPreparationTimeBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataBackupPreparationTimeBuilder {
	builder := builderSupplier(b.BackupPreparationTime.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.BackupPreparationTime, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataBackupPreparationTimeBuilder) Build() (BACnetConstructedDataBackupPreparationTime, error) {
	if b.BackupPreparationTime == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'backupPreparationTime' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataBackupPreparationTime.deepCopy(), nil
}

func (b *_BACnetConstructedDataBackupPreparationTimeBuilder) MustBuild() BACnetConstructedDataBackupPreparationTime {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataBackupPreparationTimeBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataBackupPreparationTimeBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataBackupPreparationTimeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataBackupPreparationTimeBuilder().(*_BACnetConstructedDataBackupPreparationTimeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataBackupPreparationTimeBuilder creates a BACnetConstructedDataBackupPreparationTimeBuilder
func (b *_BACnetConstructedDataBackupPreparationTime) CreateBACnetConstructedDataBackupPreparationTimeBuilder() BACnetConstructedDataBackupPreparationTimeBuilder {
	if b == nil {
		return NewBACnetConstructedDataBackupPreparationTimeBuilder()
	}
	return &_BACnetConstructedDataBackupPreparationTimeBuilder{_BACnetConstructedDataBackupPreparationTime: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBackupPreparationTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataBackupPreparationTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BACKUP_PREPARATION_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBackupPreparationTime) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBackupPreparationTime) GetBackupPreparationTime() BACnetApplicationTagUnsignedInteger {
	return m.BackupPreparationTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBackupPreparationTime) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetBackupPreparationTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBackupPreparationTime(structType any) BACnetConstructedDataBackupPreparationTime {
	if casted, ok := structType.(BACnetConstructedDataBackupPreparationTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBackupPreparationTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBackupPreparationTime) GetTypeName() string {
	return "BACnetConstructedDataBackupPreparationTime"
}

func (m *_BACnetConstructedDataBackupPreparationTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (backupPreparationTime)
	lengthInBits += m.BackupPreparationTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBackupPreparationTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataBackupPreparationTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataBackupPreparationTime BACnetConstructedDataBackupPreparationTime, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBackupPreparationTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBackupPreparationTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	backupPreparationTime, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "backupPreparationTime", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'backupPreparationTime' field"))
	}
	m.BackupPreparationTime = backupPreparationTime

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), backupPreparationTime)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBackupPreparationTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBackupPreparationTime")
	}

	return m, nil
}

func (m *_BACnetConstructedDataBackupPreparationTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBackupPreparationTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBackupPreparationTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBackupPreparationTime")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "backupPreparationTime", m.GetBackupPreparationTime(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'backupPreparationTime' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBackupPreparationTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBackupPreparationTime")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBackupPreparationTime) IsBACnetConstructedDataBackupPreparationTime() {
}

func (m *_BACnetConstructedDataBackupPreparationTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataBackupPreparationTime) deepCopy() *_BACnetConstructedDataBackupPreparationTime {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataBackupPreparationTimeCopy := &_BACnetConstructedDataBackupPreparationTime{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.BackupPreparationTime),
	}
	_BACnetConstructedDataBackupPreparationTimeCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataBackupPreparationTimeCopy
}

func (m *_BACnetConstructedDataBackupPreparationTime) String() string {
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
