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
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// SearchResponse is the corresponding interface of SearchResponse
type SearchResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	KnxNetIpMessage
	// GetHpaiControlEndpoint returns HpaiControlEndpoint (property field)
	GetHpaiControlEndpoint() HPAIControlEndpoint
	// GetDibDeviceInfo returns DibDeviceInfo (property field)
	GetDibDeviceInfo() DIBDeviceInfo
	// GetDibSuppSvcFamilies returns DibSuppSvcFamilies (property field)
	GetDibSuppSvcFamilies() DIBSuppSvcFamilies
}

// SearchResponseExactly can be used when we want exactly this type and not a type which fulfills SearchResponse.
// This is useful for switch cases.
type SearchResponseExactly interface {
	SearchResponse
	isSearchResponse() bool
}

// _SearchResponse is the data-structure of this message
type _SearchResponse struct {
	KnxNetIpMessageContract
	HpaiControlEndpoint HPAIControlEndpoint
	DibDeviceInfo       DIBDeviceInfo
	DibSuppSvcFamilies  DIBSuppSvcFamilies
}

var _ SearchResponse = (*_SearchResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SearchResponse) GetMsgType() uint16 {
	return 0x0202
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SearchResponse) GetParent() KnxNetIpMessageContract {
	return m.KnxNetIpMessageContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SearchResponse) GetHpaiControlEndpoint() HPAIControlEndpoint {
	return m.HpaiControlEndpoint
}

func (m *_SearchResponse) GetDibDeviceInfo() DIBDeviceInfo {
	return m.DibDeviceInfo
}

func (m *_SearchResponse) GetDibSuppSvcFamilies() DIBSuppSvcFamilies {
	return m.DibSuppSvcFamilies
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSearchResponse factory function for _SearchResponse
func NewSearchResponse(hpaiControlEndpoint HPAIControlEndpoint, dibDeviceInfo DIBDeviceInfo, dibSuppSvcFamilies DIBSuppSvcFamilies) *_SearchResponse {
	_result := &_SearchResponse{
		KnxNetIpMessageContract: NewKnxNetIpMessage(),
		HpaiControlEndpoint:     hpaiControlEndpoint,
		DibDeviceInfo:           dibDeviceInfo,
		DibSuppSvcFamilies:      dibSuppSvcFamilies,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastSearchResponse(structType any) SearchResponse {
	if casted, ok := structType.(SearchResponse); ok {
		return casted
	}
	if casted, ok := structType.(*SearchResponse); ok {
		return *casted
	}
	return nil
}

func (m *_SearchResponse) GetTypeName() string {
	return "SearchResponse"
}

func (m *_SearchResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.KnxNetIpMessageContract.(*_KnxNetIpMessage).getLengthInBits(ctx))

	// Simple field (hpaiControlEndpoint)
	lengthInBits += m.HpaiControlEndpoint.GetLengthInBits(ctx)

	// Simple field (dibDeviceInfo)
	lengthInBits += m.DibDeviceInfo.GetLengthInBits(ctx)

	// Simple field (dibSuppSvcFamilies)
	lengthInBits += m.DibSuppSvcFamilies.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SearchResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SearchResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_KnxNetIpMessage) (_searchResponse SearchResponse, err error) {
	m.KnxNetIpMessageContract = parent
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SearchResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SearchResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	hpaiControlEndpoint, err := ReadSimpleField[HPAIControlEndpoint](ctx, "hpaiControlEndpoint", ReadComplex[HPAIControlEndpoint](HPAIControlEndpointParseWithBuffer, readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hpaiControlEndpoint' field"))
	}
	m.HpaiControlEndpoint = hpaiControlEndpoint

	dibDeviceInfo, err := ReadSimpleField[DIBDeviceInfo](ctx, "dibDeviceInfo", ReadComplex[DIBDeviceInfo](DIBDeviceInfoParseWithBuffer, readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dibDeviceInfo' field"))
	}
	m.DibDeviceInfo = dibDeviceInfo

	dibSuppSvcFamilies, err := ReadSimpleField[DIBSuppSvcFamilies](ctx, "dibSuppSvcFamilies", ReadComplex[DIBSuppSvcFamilies](DIBSuppSvcFamiliesParseWithBuffer, readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dibSuppSvcFamilies' field"))
	}
	m.DibSuppSvcFamilies = dibSuppSvcFamilies

	if closeErr := readBuffer.CloseContext("SearchResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SearchResponse")
	}

	return m, nil
}

func (m *_SearchResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SearchResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SearchResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SearchResponse")
		}

		if err := WriteSimpleField[HPAIControlEndpoint](ctx, "hpaiControlEndpoint", m.GetHpaiControlEndpoint(), WriteComplex[HPAIControlEndpoint](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'hpaiControlEndpoint' field")
		}

		if err := WriteSimpleField[DIBDeviceInfo](ctx, "dibDeviceInfo", m.GetDibDeviceInfo(), WriteComplex[DIBDeviceInfo](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'dibDeviceInfo' field")
		}

		if err := WriteSimpleField[DIBSuppSvcFamilies](ctx, "dibSuppSvcFamilies", m.GetDibSuppSvcFamilies(), WriteComplex[DIBSuppSvcFamilies](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'dibSuppSvcFamilies' field")
		}

		if popErr := writeBuffer.PopContext("SearchResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SearchResponse")
		}
		return nil
	}
	return m.KnxNetIpMessageContract.(*_KnxNetIpMessage).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SearchResponse) isSearchResponse() bool {
	return true
}

func (m *_SearchResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
