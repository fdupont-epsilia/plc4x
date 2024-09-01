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

// SALDataMetering is the corresponding interface of SALDataMetering
type SALDataMetering interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SALData
	// GetMeteringData returns MeteringData (property field)
	GetMeteringData() MeteringData
}

// SALDataMeteringExactly can be used when we want exactly this type and not a type which fulfills SALDataMetering.
// This is useful for switch cases.
type SALDataMeteringExactly interface {
	SALDataMetering
	isSALDataMetering() bool
}

// _SALDataMetering is the data-structure of this message
type _SALDataMetering struct {
	SALDataContract
	MeteringData MeteringData
}

var _ SALDataMetering = (*_SALDataMetering)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataMetering) GetApplicationId() ApplicationId {
	return ApplicationId_METERING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataMetering) GetParent() SALDataContract {
	return m.SALDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataMetering) GetMeteringData() MeteringData {
	return m.MeteringData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSALDataMetering factory function for _SALDataMetering
func NewSALDataMetering(meteringData MeteringData, salData SALData) *_SALDataMetering {
	_result := &_SALDataMetering{
		SALDataContract: NewSALData(salData),
		MeteringData:    meteringData,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastSALDataMetering(structType any) SALDataMetering {
	if casted, ok := structType.(SALDataMetering); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataMetering); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataMetering) GetTypeName() string {
	return "SALDataMetering"
}

func (m *_SALDataMetering) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SALDataContract.(*_SALData).getLengthInBits(ctx))

	// Simple field (meteringData)
	lengthInBits += m.MeteringData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SALDataMetering) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SALDataMetering) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SALData, applicationId ApplicationId) (_sALDataMetering SALDataMetering, err error) {
	m.SALDataContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataMetering"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataMetering")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	meteringData, err := ReadSimpleField[MeteringData](ctx, "meteringData", ReadComplex[MeteringData](MeteringDataParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'meteringData' field"))
	}
	m.MeteringData = meteringData

	if closeErr := readBuffer.CloseContext("SALDataMetering"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataMetering")
	}

	return m, nil
}

func (m *_SALDataMetering) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataMetering) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataMetering"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataMetering")
		}

		if err := WriteSimpleField[MeteringData](ctx, "meteringData", m.GetMeteringData(), WriteComplex[MeteringData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'meteringData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataMetering"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataMetering")
		}
		return nil
	}
	return m.SALDataContract.(*_SALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SALDataMetering) isSALDataMetering() bool {
	return true
}

func (m *_SALDataMetering) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
