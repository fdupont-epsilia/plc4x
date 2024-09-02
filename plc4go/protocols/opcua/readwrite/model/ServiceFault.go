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

// ServiceFault is the corresponding interface of ServiceFault
type ServiceFault interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ExtensionObjectDefinition
	// IsServiceFault is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsServiceFault()
}

// _ServiceFault is the data-structure of this message
type _ServiceFault struct {
	ExtensionObjectDefinitionContract
	ResponseHeader ExtensionObjectDefinition
}

var _ ServiceFault = (*_ServiceFault)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ServiceFault)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ServiceFault) GetIdentifier() string {
	return "397"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ServiceFault) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ServiceFault) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewServiceFault factory function for _ServiceFault
func NewServiceFault(responseHeader ExtensionObjectDefinition) *_ServiceFault {
	_result := &_ServiceFault{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ResponseHeader:                    responseHeader,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastServiceFault(structType any) ServiceFault {
	if casted, ok := structType.(ServiceFault); ok {
		return casted
	}
	if casted, ok := structType.(*ServiceFault); ok {
		return *casted
	}
	return nil
}

func (m *_ServiceFault) GetTypeName() string {
	return "ServiceFault"
}

func (m *_ServiceFault) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ServiceFault) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ServiceFault) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__serviceFault ServiceFault, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ServiceFault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ServiceFault")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	responseHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("394")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'responseHeader' field"))
	}
	m.ResponseHeader = responseHeader

	if closeErr := readBuffer.CloseContext("ServiceFault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ServiceFault")
	}

	return m, nil
}

func (m *_ServiceFault) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ServiceFault) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ServiceFault"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ServiceFault")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", m.GetResponseHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'responseHeader' field")
		}

		if popErr := writeBuffer.PopContext("ServiceFault"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ServiceFault")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ServiceFault) IsServiceFault() {}

func (m *_ServiceFault) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
