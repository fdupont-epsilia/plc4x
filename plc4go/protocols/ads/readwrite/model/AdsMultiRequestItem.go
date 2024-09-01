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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// AdsMultiRequestItem is the corresponding interface of AdsMultiRequestItem
type AdsMultiRequestItem interface {
	AdsMultiRequestItemContract
	AdsMultiRequestItemRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// AdsMultiRequestItemContract provides a set of functions which can be overwritten by a sub struct
type AdsMultiRequestItemContract interface {
}

// AdsMultiRequestItemRequirements provides a set of functions which need to be implemented by a sub struct
type AdsMultiRequestItemRequirements interface {
	// GetIndexGroup returns IndexGroup (discriminator field)
	GetIndexGroup() uint32
}

// AdsMultiRequestItemExactly can be used when we want exactly this type and not a type which fulfills AdsMultiRequestItem.
// This is useful for switch cases.
type AdsMultiRequestItemExactly interface {
	AdsMultiRequestItem
	isAdsMultiRequestItem() bool
}

// _AdsMultiRequestItem is the data-structure of this message
type _AdsMultiRequestItem struct {
	_AdsMultiRequestItemChildRequirements
}

var _ AdsMultiRequestItemContract = (*_AdsMultiRequestItem)(nil)

type _AdsMultiRequestItemChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetIndexGroup() uint32
}

type AdsMultiRequestItemParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child AdsMultiRequestItem, serializeChildFunction func() error) error
	GetTypeName() string
}

type AdsMultiRequestItemChild interface {
	utils.Serializable
	InitializeParent(parent AdsMultiRequestItem)
	GetParent() *AdsMultiRequestItem

	GetTypeName() string
	AdsMultiRequestItem
}

// NewAdsMultiRequestItem factory function for _AdsMultiRequestItem
func NewAdsMultiRequestItem() *_AdsMultiRequestItem {
	return &_AdsMultiRequestItem{}
}

// Deprecated: use the interface for direct cast
func CastAdsMultiRequestItem(structType any) AdsMultiRequestItem {
	if casted, ok := structType.(AdsMultiRequestItem); ok {
		return casted
	}
	if casted, ok := structType.(*AdsMultiRequestItem); ok {
		return *casted
	}
	return nil
}

func (m *_AdsMultiRequestItem) GetTypeName() string {
	return "AdsMultiRequestItem"
}

func (m *_AdsMultiRequestItem) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_AdsMultiRequestItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsMultiRequestItemParse(ctx context.Context, theBytes []byte, indexGroup uint32) (AdsMultiRequestItem, error) {
	return AdsMultiRequestItemParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), indexGroup)
}

func AdsMultiRequestItemParseWithBufferProducer[T AdsMultiRequestItem](indexGroup uint32) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		buffer, err := AdsMultiRequestItemParseWithBuffer(ctx, readBuffer, indexGroup)
		if err != nil {
			var zero T
			return zero, err
		}
		return buffer.(T), err
	}
}

func AdsMultiRequestItemParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, indexGroup uint32) (AdsMultiRequestItem, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsMultiRequestItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsMultiRequestItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type AdsMultiRequestItemChildSerializeRequirement interface {
		AdsMultiRequestItem
		InitializeParent(AdsMultiRequestItem)
		GetParent() AdsMultiRequestItem
	}
	var _childTemp any
	var _child AdsMultiRequestItemChildSerializeRequirement
	var typeSwitchError error
	switch {
	case indexGroup == uint32(61568): // AdsMultiRequestItemRead
		_childTemp, typeSwitchError = AdsMultiRequestItemReadParseWithBuffer(ctx, readBuffer, indexGroup)
	case indexGroup == uint32(61569): // AdsMultiRequestItemWrite
		_childTemp, typeSwitchError = AdsMultiRequestItemWriteParseWithBuffer(ctx, readBuffer, indexGroup)
	case indexGroup == uint32(61570): // AdsMultiRequestItemReadWrite
		_childTemp, typeSwitchError = AdsMultiRequestItemReadWriteParseWithBuffer(ctx, readBuffer, indexGroup)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [indexGroup=%v]", indexGroup)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of AdsMultiRequestItem")
	}
	_child = _childTemp.(AdsMultiRequestItemChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("AdsMultiRequestItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsMultiRequestItem")
	}

	// Finish initializing
	_child.InitializeParent(_child)
	return _child, nil
}

func (pm *_AdsMultiRequestItem) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child AdsMultiRequestItem, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AdsMultiRequestItem"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AdsMultiRequestItem")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("AdsMultiRequestItem"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AdsMultiRequestItem")
	}
	return nil
}

func (m *_AdsMultiRequestItem) isAdsMultiRequestItem() bool {
	return true
}

func (m *_AdsMultiRequestItem) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
