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

// NLMEstablishConnectionToNetwork is the corresponding interface of NLMEstablishConnectionToNetwork
type NLMEstablishConnectionToNetwork interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	NLM
	// GetDestinationNetworkAddress returns DestinationNetworkAddress (property field)
	GetDestinationNetworkAddress() uint16
	// GetTerminationTime returns TerminationTime (property field)
	GetTerminationTime() uint8
	// IsNLMEstablishConnectionToNetwork is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNLMEstablishConnectionToNetwork()
	// CreateBuilder creates a NLMEstablishConnectionToNetworkBuilder
	CreateNLMEstablishConnectionToNetworkBuilder() NLMEstablishConnectionToNetworkBuilder
}

// _NLMEstablishConnectionToNetwork is the data-structure of this message
type _NLMEstablishConnectionToNetwork struct {
	NLMContract
	DestinationNetworkAddress uint16
	TerminationTime           uint8
}

var _ NLMEstablishConnectionToNetwork = (*_NLMEstablishConnectionToNetwork)(nil)
var _ NLMRequirements = (*_NLMEstablishConnectionToNetwork)(nil)

// NewNLMEstablishConnectionToNetwork factory function for _NLMEstablishConnectionToNetwork
func NewNLMEstablishConnectionToNetwork(destinationNetworkAddress uint16, terminationTime uint8, apduLength uint16) *_NLMEstablishConnectionToNetwork {
	_result := &_NLMEstablishConnectionToNetwork{
		NLMContract:               NewNLM(apduLength),
		DestinationNetworkAddress: destinationNetworkAddress,
		TerminationTime:           terminationTime,
	}
	_result.NLMContract.(*_NLM)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// NLMEstablishConnectionToNetworkBuilder is a builder for NLMEstablishConnectionToNetwork
type NLMEstablishConnectionToNetworkBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(destinationNetworkAddress uint16, terminationTime uint8) NLMEstablishConnectionToNetworkBuilder
	// WithDestinationNetworkAddress adds DestinationNetworkAddress (property field)
	WithDestinationNetworkAddress(uint16) NLMEstablishConnectionToNetworkBuilder
	// WithTerminationTime adds TerminationTime (property field)
	WithTerminationTime(uint8) NLMEstablishConnectionToNetworkBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() NLMBuilder
	// Build builds the NLMEstablishConnectionToNetwork or returns an error if something is wrong
	Build() (NLMEstablishConnectionToNetwork, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() NLMEstablishConnectionToNetwork
}

// NewNLMEstablishConnectionToNetworkBuilder() creates a NLMEstablishConnectionToNetworkBuilder
func NewNLMEstablishConnectionToNetworkBuilder() NLMEstablishConnectionToNetworkBuilder {
	return &_NLMEstablishConnectionToNetworkBuilder{_NLMEstablishConnectionToNetwork: new(_NLMEstablishConnectionToNetwork)}
}

type _NLMEstablishConnectionToNetworkBuilder struct {
	*_NLMEstablishConnectionToNetwork

	parentBuilder *_NLMBuilder

	err *utils.MultiError
}

var _ (NLMEstablishConnectionToNetworkBuilder) = (*_NLMEstablishConnectionToNetworkBuilder)(nil)

func (b *_NLMEstablishConnectionToNetworkBuilder) setParent(contract NLMContract) {
	b.NLMContract = contract
	contract.(*_NLM)._SubType = b._NLMEstablishConnectionToNetwork
}

func (b *_NLMEstablishConnectionToNetworkBuilder) WithMandatoryFields(destinationNetworkAddress uint16, terminationTime uint8) NLMEstablishConnectionToNetworkBuilder {
	return b.WithDestinationNetworkAddress(destinationNetworkAddress).WithTerminationTime(terminationTime)
}

func (b *_NLMEstablishConnectionToNetworkBuilder) WithDestinationNetworkAddress(destinationNetworkAddress uint16) NLMEstablishConnectionToNetworkBuilder {
	b.DestinationNetworkAddress = destinationNetworkAddress
	return b
}

func (b *_NLMEstablishConnectionToNetworkBuilder) WithTerminationTime(terminationTime uint8) NLMEstablishConnectionToNetworkBuilder {
	b.TerminationTime = terminationTime
	return b
}

func (b *_NLMEstablishConnectionToNetworkBuilder) Build() (NLMEstablishConnectionToNetwork, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._NLMEstablishConnectionToNetwork.deepCopy(), nil
}

func (b *_NLMEstablishConnectionToNetworkBuilder) MustBuild() NLMEstablishConnectionToNetwork {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_NLMEstablishConnectionToNetworkBuilder) Done() NLMBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewNLMBuilder().(*_NLMBuilder)
	}
	return b.parentBuilder
}

func (b *_NLMEstablishConnectionToNetworkBuilder) buildForNLM() (NLM, error) {
	return b.Build()
}

func (b *_NLMEstablishConnectionToNetworkBuilder) DeepCopy() any {
	_copy := b.CreateNLMEstablishConnectionToNetworkBuilder().(*_NLMEstablishConnectionToNetworkBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateNLMEstablishConnectionToNetworkBuilder creates a NLMEstablishConnectionToNetworkBuilder
func (b *_NLMEstablishConnectionToNetwork) CreateNLMEstablishConnectionToNetworkBuilder() NLMEstablishConnectionToNetworkBuilder {
	if b == nil {
		return NewNLMEstablishConnectionToNetworkBuilder()
	}
	return &_NLMEstablishConnectionToNetworkBuilder{_NLMEstablishConnectionToNetwork: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMEstablishConnectionToNetwork) GetMessageType() uint8 {
	return 0x08
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMEstablishConnectionToNetwork) GetParent() NLMContract {
	return m.NLMContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMEstablishConnectionToNetwork) GetDestinationNetworkAddress() uint16 {
	return m.DestinationNetworkAddress
}

func (m *_NLMEstablishConnectionToNetwork) GetTerminationTime() uint8 {
	return m.TerminationTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastNLMEstablishConnectionToNetwork(structType any) NLMEstablishConnectionToNetwork {
	if casted, ok := structType.(NLMEstablishConnectionToNetwork); ok {
		return casted
	}
	if casted, ok := structType.(*NLMEstablishConnectionToNetwork); ok {
		return *casted
	}
	return nil
}

func (m *_NLMEstablishConnectionToNetwork) GetTypeName() string {
	return "NLMEstablishConnectionToNetwork"
}

func (m *_NLMEstablishConnectionToNetwork) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.NLMContract.(*_NLM).getLengthInBits(ctx))

	// Simple field (destinationNetworkAddress)
	lengthInBits += 16

	// Simple field (terminationTime)
	lengthInBits += 8

	return lengthInBits
}

func (m *_NLMEstablishConnectionToNetwork) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NLMEstablishConnectionToNetwork) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_NLM, apduLength uint16) (__nLMEstablishConnectionToNetwork NLMEstablishConnectionToNetwork, err error) {
	m.NLMContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMEstablishConnectionToNetwork"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMEstablishConnectionToNetwork")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	destinationNetworkAddress, err := ReadSimpleField(ctx, "destinationNetworkAddress", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'destinationNetworkAddress' field"))
	}
	m.DestinationNetworkAddress = destinationNetworkAddress

	terminationTime, err := ReadSimpleField(ctx, "terminationTime", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'terminationTime' field"))
	}
	m.TerminationTime = terminationTime

	if closeErr := readBuffer.CloseContext("NLMEstablishConnectionToNetwork"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMEstablishConnectionToNetwork")
	}

	return m, nil
}

func (m *_NLMEstablishConnectionToNetwork) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMEstablishConnectionToNetwork) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMEstablishConnectionToNetwork"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMEstablishConnectionToNetwork")
		}

		if err := WriteSimpleField[uint16](ctx, "destinationNetworkAddress", m.GetDestinationNetworkAddress(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'destinationNetworkAddress' field")
		}

		if err := WriteSimpleField[uint8](ctx, "terminationTime", m.GetTerminationTime(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'terminationTime' field")
		}

		if popErr := writeBuffer.PopContext("NLMEstablishConnectionToNetwork"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMEstablishConnectionToNetwork")
		}
		return nil
	}
	return m.NLMContract.(*_NLM).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NLMEstablishConnectionToNetwork) IsNLMEstablishConnectionToNetwork() {}

func (m *_NLMEstablishConnectionToNetwork) DeepCopy() any {
	return m.deepCopy()
}

func (m *_NLMEstablishConnectionToNetwork) deepCopy() *_NLMEstablishConnectionToNetwork {
	if m == nil {
		return nil
	}
	_NLMEstablishConnectionToNetworkCopy := &_NLMEstablishConnectionToNetwork{
		m.NLMContract.(*_NLM).deepCopy(),
		m.DestinationNetworkAddress,
		m.TerminationTime,
	}
	_NLMEstablishConnectionToNetworkCopy.NLMContract.(*_NLM)._SubType = m
	return _NLMEstablishConnectionToNetworkCopy
}

func (m *_NLMEstablishConnectionToNetwork) String() string {
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
