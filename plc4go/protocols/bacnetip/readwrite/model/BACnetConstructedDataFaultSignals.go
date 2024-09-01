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

// BACnetConstructedDataFaultSignals is the corresponding interface of BACnetConstructedDataFaultSignals
type BACnetConstructedDataFaultSignals interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFaultSignals returns FaultSignals (property field)
	GetFaultSignals() []BACnetLiftFaultTagged
}

// BACnetConstructedDataFaultSignalsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataFaultSignals.
// This is useful for switch cases.
type BACnetConstructedDataFaultSignalsExactly interface {
	BACnetConstructedDataFaultSignals
	isBACnetConstructedDataFaultSignals() bool
}

// _BACnetConstructedDataFaultSignals is the data-structure of this message
type _BACnetConstructedDataFaultSignals struct {
	*_BACnetConstructedData
	FaultSignals []BACnetLiftFaultTagged
}

var _ BACnetConstructedDataFaultSignals = (*_BACnetConstructedDataFaultSignals)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataFaultSignals) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataFaultSignals) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAULT_SIGNALS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataFaultSignals) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataFaultSignals) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataFaultSignals) GetFaultSignals() []BACnetLiftFaultTagged {
	return m.FaultSignals
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataFaultSignals factory function for _BACnetConstructedDataFaultSignals
func NewBACnetConstructedDataFaultSignals(faultSignals []BACnetLiftFaultTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataFaultSignals {
	_result := &_BACnetConstructedDataFaultSignals{
		FaultSignals:           faultSignals,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataFaultSignals(structType any) BACnetConstructedDataFaultSignals {
	if casted, ok := structType.(BACnetConstructedDataFaultSignals); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataFaultSignals); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataFaultSignals) GetTypeName() string {
	return "BACnetConstructedDataFaultSignals"
}

func (m *_BACnetConstructedDataFaultSignals) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.FaultSignals) > 0 {
		for _, element := range m.FaultSignals {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataFaultSignals) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataFaultSignalsParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataFaultSignals, error) {
	return BACnetConstructedDataFaultSignalsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataFaultSignalsParseWithBufferProducer(tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataFaultSignals, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataFaultSignals, error) {
		return BACnetConstructedDataFaultSignalsParseWithBuffer(ctx, readBuffer, tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
	}
}

func BACnetConstructedDataFaultSignalsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataFaultSignals, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataFaultSignals"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataFaultSignals")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	faultSignals, err := ReadTerminatedArrayField[BACnetLiftFaultTagged](ctx, "faultSignals", ReadComplex[BACnetLiftFaultTagged](BACnetLiftFaultTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'faultSignals' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataFaultSignals"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataFaultSignals")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataFaultSignals{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		FaultSignals: faultSignals,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataFaultSignals) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataFaultSignals) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataFaultSignals"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataFaultSignals")
		}

		if err := WriteComplexTypeArrayField(ctx, "faultSignals", m.GetFaultSignals(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'faultSignals' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataFaultSignals"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataFaultSignals")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataFaultSignals) isBACnetConstructedDataFaultSignals() bool {
	return true
}

func (m *_BACnetConstructedDataFaultSignals) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
