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

// BACnetConstructedDataLastKeyServer is the corresponding interface of BACnetConstructedDataLastKeyServer
type BACnetConstructedDataLastKeyServer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLastKeyServer returns LastKeyServer (property field)
	GetLastKeyServer() BACnetAddressBinding
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetAddressBinding
	// IsBACnetConstructedDataLastKeyServer is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLastKeyServer()
}

// _BACnetConstructedDataLastKeyServer is the data-structure of this message
type _BACnetConstructedDataLastKeyServer struct {
	BACnetConstructedDataContract
	LastKeyServer BACnetAddressBinding
}

var _ BACnetConstructedDataLastKeyServer = (*_BACnetConstructedDataLastKeyServer)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLastKeyServer)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLastKeyServer) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLastKeyServer) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LAST_KEY_SERVER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLastKeyServer) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLastKeyServer) GetLastKeyServer() BACnetAddressBinding {
	return m.LastKeyServer
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLastKeyServer) GetActualValue() BACnetAddressBinding {
	ctx := context.Background()
	_ = ctx
	return CastBACnetAddressBinding(m.GetLastKeyServer())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLastKeyServer factory function for _BACnetConstructedDataLastKeyServer
func NewBACnetConstructedDataLastKeyServer(lastKeyServer BACnetAddressBinding, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLastKeyServer {
	_result := &_BACnetConstructedDataLastKeyServer{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LastKeyServer:                 lastKeyServer,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLastKeyServer(structType any) BACnetConstructedDataLastKeyServer {
	if casted, ok := structType.(BACnetConstructedDataLastKeyServer); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLastKeyServer); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLastKeyServer) GetTypeName() string {
	return "BACnetConstructedDataLastKeyServer"
}

func (m *_BACnetConstructedDataLastKeyServer) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (lastKeyServer)
	lengthInBits += m.LastKeyServer.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLastKeyServer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLastKeyServer) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLastKeyServer BACnetConstructedDataLastKeyServer, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLastKeyServer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLastKeyServer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lastKeyServer, err := ReadSimpleField[BACnetAddressBinding](ctx, "lastKeyServer", ReadComplex[BACnetAddressBinding](BACnetAddressBindingParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastKeyServer' field"))
	}
	m.LastKeyServer = lastKeyServer

	actualValue, err := ReadVirtualField[BACnetAddressBinding](ctx, "actualValue", (*BACnetAddressBinding)(nil), lastKeyServer)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLastKeyServer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLastKeyServer")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLastKeyServer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLastKeyServer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLastKeyServer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLastKeyServer")
		}

		if err := WriteSimpleField[BACnetAddressBinding](ctx, "lastKeyServer", m.GetLastKeyServer(), WriteComplex[BACnetAddressBinding](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lastKeyServer' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLastKeyServer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLastKeyServer")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLastKeyServer) IsBACnetConstructedDataLastKeyServer() {}

func (m *_BACnetConstructedDataLastKeyServer) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
