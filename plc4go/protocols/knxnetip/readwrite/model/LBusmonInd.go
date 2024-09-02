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
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// LBusmonInd is the corresponding interface of LBusmonInd
type LBusmonInd interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CEMI
	// GetAdditionalInformationLength returns AdditionalInformationLength (property field)
	GetAdditionalInformationLength() uint8
	// GetAdditionalInformation returns AdditionalInformation (property field)
	GetAdditionalInformation() []CEMIAdditionalInformation
	// GetDataFrame returns DataFrame (property field)
	GetDataFrame() LDataFrame
	// GetCrc returns Crc (property field)
	GetCrc() *uint8
	// IsLBusmonInd is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsLBusmonInd()
}

// _LBusmonInd is the data-structure of this message
type _LBusmonInd struct {
	CEMIContract
	AdditionalInformationLength uint8
	AdditionalInformation       []CEMIAdditionalInformation
	DataFrame                   LDataFrame
	Crc                         *uint8
}

var _ LBusmonInd = (*_LBusmonInd)(nil)
var _ CEMIRequirements = (*_LBusmonInd)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_LBusmonInd) GetMessageCode() uint8 {
	return 0x2B
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LBusmonInd) GetParent() CEMIContract {
	return m.CEMIContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LBusmonInd) GetAdditionalInformationLength() uint8 {
	return m.AdditionalInformationLength
}

func (m *_LBusmonInd) GetAdditionalInformation() []CEMIAdditionalInformation {
	return m.AdditionalInformation
}

func (m *_LBusmonInd) GetDataFrame() LDataFrame {
	return m.DataFrame
}

func (m *_LBusmonInd) GetCrc() *uint8 {
	return m.Crc
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewLBusmonInd factory function for _LBusmonInd
func NewLBusmonInd(additionalInformationLength uint8, additionalInformation []CEMIAdditionalInformation, dataFrame LDataFrame, crc *uint8, size uint16) *_LBusmonInd {
	_result := &_LBusmonInd{
		CEMIContract:                NewCEMI(size),
		AdditionalInformationLength: additionalInformationLength,
		AdditionalInformation:       additionalInformation,
		DataFrame:                   dataFrame,
		Crc:                         crc,
	}
	_result.CEMIContract.(*_CEMI)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastLBusmonInd(structType any) LBusmonInd {
	if casted, ok := structType.(LBusmonInd); ok {
		return casted
	}
	if casted, ok := structType.(*LBusmonInd); ok {
		return *casted
	}
	return nil
}

func (m *_LBusmonInd) GetTypeName() string {
	return "LBusmonInd"
}

func (m *_LBusmonInd) GetLengthInBits(ctx context.Context) uint16 {
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

	// Optional Field (crc)
	if m.Crc != nil {
		lengthInBits += 8
	}

	return lengthInBits
}

func (m *_LBusmonInd) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_LBusmonInd) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CEMI, size uint16) (__lBusmonInd LBusmonInd, err error) {
	m.CEMIContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LBusmonInd"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LBusmonInd")
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

	var crc *uint8
	crc, err = ReadOptionalField[uint8](ctx, "crc", ReadUnsignedByte(readBuffer, uint8(8)), CastLDataFrame(dataFrame).GetNotAckFrame())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'crc' field"))
	}
	m.Crc = crc

	if closeErr := readBuffer.CloseContext("LBusmonInd"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LBusmonInd")
	}

	return m, nil
}

func (m *_LBusmonInd) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LBusmonInd) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LBusmonInd"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LBusmonInd")
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

		if err := WriteOptionalField[uint8](ctx, "crc", m.GetCrc(), WriteUnsignedByte(writeBuffer, 8), true); err != nil {
			return errors.Wrap(err, "Error serializing 'crc' field")
		}

		if popErr := writeBuffer.PopContext("LBusmonInd"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LBusmonInd")
		}
		return nil
	}
	return m.CEMIContract.(*_CEMI).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_LBusmonInd) IsLBusmonInd() {}

func (m *_LBusmonInd) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
