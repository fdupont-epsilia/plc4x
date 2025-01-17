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

// LDataReq is the corresponding interface of LDataReq
type LDataReq interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CEMI
	// GetAdditionalInformationLength returns AdditionalInformationLength (property field)
	GetAdditionalInformationLength() uint8
	// GetAdditionalInformation returns AdditionalInformation (property field)
	GetAdditionalInformation() []CEMIAdditionalInformation
	// GetDataFrame returns DataFrame (property field)
	GetDataFrame() LDataFrame
	// IsLDataReq is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsLDataReq()
	// CreateBuilder creates a LDataReqBuilder
	CreateLDataReqBuilder() LDataReqBuilder
}

// _LDataReq is the data-structure of this message
type _LDataReq struct {
	CEMIContract
	AdditionalInformationLength uint8
	AdditionalInformation       []CEMIAdditionalInformation
	DataFrame                   LDataFrame
}

var _ LDataReq = (*_LDataReq)(nil)
var _ CEMIRequirements = (*_LDataReq)(nil)

// NewLDataReq factory function for _LDataReq
func NewLDataReq(additionalInformationLength uint8, additionalInformation []CEMIAdditionalInformation, dataFrame LDataFrame, size uint16) *_LDataReq {
	if dataFrame == nil {
		panic("dataFrame of type LDataFrame for LDataReq must not be nil")
	}
	_result := &_LDataReq{
		CEMIContract:                NewCEMI(size),
		AdditionalInformationLength: additionalInformationLength,
		AdditionalInformation:       additionalInformation,
		DataFrame:                   dataFrame,
	}
	_result.CEMIContract.(*_CEMI)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// LDataReqBuilder is a builder for LDataReq
type LDataReqBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(additionalInformationLength uint8, additionalInformation []CEMIAdditionalInformation, dataFrame LDataFrame) LDataReqBuilder
	// WithAdditionalInformationLength adds AdditionalInformationLength (property field)
	WithAdditionalInformationLength(uint8) LDataReqBuilder
	// WithAdditionalInformation adds AdditionalInformation (property field)
	WithAdditionalInformation(...CEMIAdditionalInformation) LDataReqBuilder
	// WithDataFrame adds DataFrame (property field)
	WithDataFrame(LDataFrame) LDataReqBuilder
	// WithDataFrameBuilder adds DataFrame (property field) which is build by the builder
	WithDataFrameBuilder(func(LDataFrameBuilder) LDataFrameBuilder) LDataReqBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() CEMIBuilder
	// Build builds the LDataReq or returns an error if something is wrong
	Build() (LDataReq, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() LDataReq
}

// NewLDataReqBuilder() creates a LDataReqBuilder
func NewLDataReqBuilder() LDataReqBuilder {
	return &_LDataReqBuilder{_LDataReq: new(_LDataReq)}
}

type _LDataReqBuilder struct {
	*_LDataReq

	parentBuilder *_CEMIBuilder

	err *utils.MultiError
}

var _ (LDataReqBuilder) = (*_LDataReqBuilder)(nil)

func (b *_LDataReqBuilder) setParent(contract CEMIContract) {
	b.CEMIContract = contract
	contract.(*_CEMI)._SubType = b._LDataReq
}

func (b *_LDataReqBuilder) WithMandatoryFields(additionalInformationLength uint8, additionalInformation []CEMIAdditionalInformation, dataFrame LDataFrame) LDataReqBuilder {
	return b.WithAdditionalInformationLength(additionalInformationLength).WithAdditionalInformation(additionalInformation...).WithDataFrame(dataFrame)
}

func (b *_LDataReqBuilder) WithAdditionalInformationLength(additionalInformationLength uint8) LDataReqBuilder {
	b.AdditionalInformationLength = additionalInformationLength
	return b
}

func (b *_LDataReqBuilder) WithAdditionalInformation(additionalInformation ...CEMIAdditionalInformation) LDataReqBuilder {
	b.AdditionalInformation = additionalInformation
	return b
}

func (b *_LDataReqBuilder) WithDataFrame(dataFrame LDataFrame) LDataReqBuilder {
	b.DataFrame = dataFrame
	return b
}

func (b *_LDataReqBuilder) WithDataFrameBuilder(builderSupplier func(LDataFrameBuilder) LDataFrameBuilder) LDataReqBuilder {
	builder := builderSupplier(b.DataFrame.CreateLDataFrameBuilder())
	var err error
	b.DataFrame, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "LDataFrameBuilder failed"))
	}
	return b
}

func (b *_LDataReqBuilder) Build() (LDataReq, error) {
	if b.DataFrame == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'dataFrame' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._LDataReq.deepCopy(), nil
}

func (b *_LDataReqBuilder) MustBuild() LDataReq {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_LDataReqBuilder) Done() CEMIBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewCEMIBuilder().(*_CEMIBuilder)
	}
	return b.parentBuilder
}

func (b *_LDataReqBuilder) buildForCEMI() (CEMI, error) {
	return b.Build()
}

func (b *_LDataReqBuilder) DeepCopy() any {
	_copy := b.CreateLDataReqBuilder().(*_LDataReqBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateLDataReqBuilder creates a LDataReqBuilder
func (b *_LDataReq) CreateLDataReqBuilder() LDataReqBuilder {
	if b == nil {
		return NewLDataReqBuilder()
	}
	return &_LDataReqBuilder{_LDataReq: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_LDataReq) GetMessageCode() uint8 {
	return 0x11
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LDataReq) GetParent() CEMIContract {
	return m.CEMIContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LDataReq) GetAdditionalInformationLength() uint8 {
	return m.AdditionalInformationLength
}

func (m *_LDataReq) GetAdditionalInformation() []CEMIAdditionalInformation {
	return m.AdditionalInformation
}

func (m *_LDataReq) GetDataFrame() LDataFrame {
	return m.DataFrame
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastLDataReq(structType any) LDataReq {
	if casted, ok := structType.(LDataReq); ok {
		return casted
	}
	if casted, ok := structType.(*LDataReq); ok {
		return *casted
	}
	return nil
}

func (m *_LDataReq) GetTypeName() string {
	return "LDataReq"
}

func (m *_LDataReq) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CEMIContract.(*_CEMI).getLengthInBits(ctx))

	// Simple field (additionalInformationLength)
	lengthInBits += 8

	// Array field
	if len(m.AdditionalInformation) > 0 {
		for _, element := range m.AdditionalInformation {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (dataFrame)
	lengthInBits += m.DataFrame.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_LDataReq) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_LDataReq) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CEMI, size uint16) (__lDataReq LDataReq, err error) {
	m.CEMIContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LDataReq"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LDataReq")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	additionalInformationLength, err := ReadSimpleField(ctx, "additionalInformationLength", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'additionalInformationLength' field"))
	}
	m.AdditionalInformationLength = additionalInformationLength

	additionalInformation, err := ReadLengthArrayField[CEMIAdditionalInformation](ctx, "additionalInformation", ReadComplex[CEMIAdditionalInformation](CEMIAdditionalInformationParseWithBuffer, readBuffer), int(additionalInformationLength))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'additionalInformation' field"))
	}
	m.AdditionalInformation = additionalInformation

	dataFrame, err := ReadSimpleField[LDataFrame](ctx, "dataFrame", ReadComplex[LDataFrame](LDataFrameParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataFrame' field"))
	}
	m.DataFrame = dataFrame

	if closeErr := readBuffer.CloseContext("LDataReq"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LDataReq")
	}

	return m, nil
}

func (m *_LDataReq) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LDataReq) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LDataReq"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LDataReq")
		}

		if err := WriteSimpleField[uint8](ctx, "additionalInformationLength", m.GetAdditionalInformationLength(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'additionalInformationLength' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "additionalInformation", m.GetAdditionalInformation(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'additionalInformation' field")
		}

		if err := WriteSimpleField[LDataFrame](ctx, "dataFrame", m.GetDataFrame(), WriteComplex[LDataFrame](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dataFrame' field")
		}

		if popErr := writeBuffer.PopContext("LDataReq"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LDataReq")
		}
		return nil
	}
	return m.CEMIContract.(*_CEMI).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_LDataReq) IsLDataReq() {}

func (m *_LDataReq) DeepCopy() any {
	return m.deepCopy()
}

func (m *_LDataReq) deepCopy() *_LDataReq {
	if m == nil {
		return nil
	}
	_LDataReqCopy := &_LDataReq{
		m.CEMIContract.(*_CEMI).deepCopy(),
		m.AdditionalInformationLength,
		utils.DeepCopySlice[CEMIAdditionalInformation, CEMIAdditionalInformation](m.AdditionalInformation),
		utils.DeepCopy[LDataFrame](m.DataFrame),
	}
	_LDataReqCopy.CEMIContract.(*_CEMI)._SubType = m
	return _LDataReqCopy
}

func (m *_LDataReq) String() string {
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
