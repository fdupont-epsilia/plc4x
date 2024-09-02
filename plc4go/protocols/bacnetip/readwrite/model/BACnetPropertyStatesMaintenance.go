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

// BACnetPropertyStatesMaintenance is the corresponding interface of BACnetPropertyStatesMaintenance
type BACnetPropertyStatesMaintenance interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetMaintenance returns Maintenance (property field)
	GetMaintenance() BACnetMaintenanceTagged
	// IsBACnetPropertyStatesMaintenance is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesMaintenance()
}

// _BACnetPropertyStatesMaintenance is the data-structure of this message
type _BACnetPropertyStatesMaintenance struct {
	BACnetPropertyStatesContract
	Maintenance BACnetMaintenanceTagged
}

var _ BACnetPropertyStatesMaintenance = (*_BACnetPropertyStatesMaintenance)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesMaintenance)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesMaintenance) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesMaintenance) GetMaintenance() BACnetMaintenanceTagged {
	return m.Maintenance
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesMaintenance factory function for _BACnetPropertyStatesMaintenance
func NewBACnetPropertyStatesMaintenance(maintenance BACnetMaintenanceTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesMaintenance {
	_result := &_BACnetPropertyStatesMaintenance{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		Maintenance:                  maintenance,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesMaintenance(structType any) BACnetPropertyStatesMaintenance {
	if casted, ok := structType.(BACnetPropertyStatesMaintenance); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesMaintenance); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesMaintenance) GetTypeName() string {
	return "BACnetPropertyStatesMaintenance"
}

func (m *_BACnetPropertyStatesMaintenance) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (maintenance)
	lengthInBits += m.Maintenance.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesMaintenance) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesMaintenance) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesMaintenance BACnetPropertyStatesMaintenance, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesMaintenance"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesMaintenance")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	maintenance, err := ReadSimpleField[BACnetMaintenanceTagged](ctx, "maintenance", ReadComplex[BACnetMaintenanceTagged](BACnetMaintenanceTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maintenance' field"))
	}
	m.Maintenance = maintenance

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesMaintenance"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesMaintenance")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesMaintenance) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesMaintenance) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesMaintenance"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesMaintenance")
		}

		if err := WriteSimpleField[BACnetMaintenanceTagged](ctx, "maintenance", m.GetMaintenance(), WriteComplex[BACnetMaintenanceTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'maintenance' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesMaintenance"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesMaintenance")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesMaintenance) IsBACnetPropertyStatesMaintenance() {}

func (m *_BACnetPropertyStatesMaintenance) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
