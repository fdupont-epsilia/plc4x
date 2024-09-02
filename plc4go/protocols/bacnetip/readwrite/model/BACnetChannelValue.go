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

// BACnetChannelValue is the corresponding interface of BACnetChannelValue
type BACnetChannelValue interface {
	BACnetChannelValueContract
	BACnetChannelValueRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// IsBACnetChannelValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetChannelValue()
}

// BACnetChannelValueContract provides a set of functions which can be overwritten by a sub struct
type BACnetChannelValueContract interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetPeekedIsContextTag returns PeekedIsContextTag (virtual field)
	GetPeekedIsContextTag() bool
	// IsBACnetChannelValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetChannelValue()
}

// BACnetChannelValueRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetChannelValueRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedIsContextTag returns PeekedIsContextTag (discriminator field)
	GetPeekedIsContextTag() bool
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// _BACnetChannelValue is the data-structure of this message
type _BACnetChannelValue struct {
	_SubType        BACnetChannelValue
	PeekedTagHeader BACnetTagHeader
}

var _ BACnetChannelValueContract = (*_BACnetChannelValue)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetChannelValue) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_BACnetChannelValue) GetPeekedTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

func (pm *_BACnetChannelValue) GetPeekedIsContextTag() bool {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetPeekedTagHeader().GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetChannelValue factory function for _BACnetChannelValue
func NewBACnetChannelValue(peekedTagHeader BACnetTagHeader) *_BACnetChannelValue {
	return &_BACnetChannelValue{PeekedTagHeader: peekedTagHeader}
}

// Deprecated: use the interface for direct cast
func CastBACnetChannelValue(structType any) BACnetChannelValue {
	if casted, ok := structType.(BACnetChannelValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetChannelValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetChannelValue) GetTypeName() string {
	return "BACnetChannelValue"
}

func (m *_BACnetChannelValue) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetChannelValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetChannelValueParse[T BACnetChannelValue](ctx context.Context, theBytes []byte) (T, error) {
	return BACnetChannelValueParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetChannelValueParseWithBufferProducer[T BACnetChannelValue]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetChannelValueParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func BACnetChannelValueParseWithBuffer[T BACnetChannelValue](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_BACnetChannelValue{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_BACnetChannelValue) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetChannelValue BACnetChannelValue, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetChannelValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	peekedTagHeader, err := ReadPeekField[BACnetTagHeader](ctx, "peekedTagHeader", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagHeader' field"))
	}
	m.PeekedTagHeader = peekedTagHeader

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}
	_ = peekedTagNumber

	peekedIsContextTag, err := ReadVirtualField[bool](ctx, "peekedIsContextTag", (*bool)(nil), bool((peekedTagHeader.GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedIsContextTag' field"))
	}
	_ = peekedIsContextTag

	// Validation
	if !(bool((!(peekedIsContextTag))) || bool((bool(bool(peekedIsContextTag) && bool(bool((peekedTagHeader.GetLengthValueType()) != (0x6)))) && bool(bool((peekedTagHeader.GetLengthValueType()) != (0x7)))))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "unexpected opening or closing tag"})
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetChannelValue
	switch {
	case peekedTagNumber == 0x0 && peekedIsContextTag == bool(false): // BACnetChannelValueNull
		if _child, err = (&_BACnetChannelValueNull{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetChannelValueNull for type-switch of BACnetChannelValue")
		}
	case peekedTagNumber == 0x4 && peekedIsContextTag == bool(false): // BACnetChannelValueReal
		if _child, err = (&_BACnetChannelValueReal{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetChannelValueReal for type-switch of BACnetChannelValue")
		}
	case peekedTagNumber == 0x9 && peekedIsContextTag == bool(false): // BACnetChannelValueEnumerated
		if _child, err = (&_BACnetChannelValueEnumerated{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetChannelValueEnumerated for type-switch of BACnetChannelValue")
		}
	case peekedTagNumber == 0x2 && peekedIsContextTag == bool(false): // BACnetChannelValueUnsigned
		if _child, err = (&_BACnetChannelValueUnsigned{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetChannelValueUnsigned for type-switch of BACnetChannelValue")
		}
	case peekedTagNumber == 0x1 && peekedIsContextTag == bool(false): // BACnetChannelValueBoolean
		if _child, err = (&_BACnetChannelValueBoolean{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetChannelValueBoolean for type-switch of BACnetChannelValue")
		}
	case peekedTagNumber == 0x3 && peekedIsContextTag == bool(false): // BACnetChannelValueInteger
		if _child, err = (&_BACnetChannelValueInteger{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetChannelValueInteger for type-switch of BACnetChannelValue")
		}
	case peekedTagNumber == 0x5 && peekedIsContextTag == bool(false): // BACnetChannelValueDouble
		if _child, err = (&_BACnetChannelValueDouble{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetChannelValueDouble for type-switch of BACnetChannelValue")
		}
	case peekedTagNumber == 0xB && peekedIsContextTag == bool(false): // BACnetChannelValueTime
		if _child, err = (&_BACnetChannelValueTime{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetChannelValueTime for type-switch of BACnetChannelValue")
		}
	case peekedTagNumber == 0x7 && peekedIsContextTag == bool(false): // BACnetChannelValueCharacterString
		if _child, err = (&_BACnetChannelValueCharacterString{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetChannelValueCharacterString for type-switch of BACnetChannelValue")
		}
	case peekedTagNumber == 0x6 && peekedIsContextTag == bool(false): // BACnetChannelValueOctetString
		if _child, err = (&_BACnetChannelValueOctetString{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetChannelValueOctetString for type-switch of BACnetChannelValue")
		}
	case peekedTagNumber == 0x8 && peekedIsContextTag == bool(false): // BACnetChannelValueBitString
		if _child, err = (&_BACnetChannelValueBitString{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetChannelValueBitString for type-switch of BACnetChannelValue")
		}
	case peekedTagNumber == 0xA && peekedIsContextTag == bool(false): // BACnetChannelValueDate
		if _child, err = (&_BACnetChannelValueDate{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetChannelValueDate for type-switch of BACnetChannelValue")
		}
	case peekedTagNumber == 0xC && peekedIsContextTag == bool(false): // BACnetChannelValueObjectidentifier
		if _child, err = (&_BACnetChannelValueObjectidentifier{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetChannelValueObjectidentifier for type-switch of BACnetChannelValue")
		}
	case peekedTagNumber == uint8(0) && peekedIsContextTag == bool(true): // BACnetChannelValueLightingCommand
		if _child, err = (&_BACnetChannelValueLightingCommand{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetChannelValueLightingCommand for type-switch of BACnetChannelValue")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v, peekedIsContextTag=%v]", peekedTagNumber, peekedIsContextTag)
	}

	if closeErr := readBuffer.CloseContext("BACnetChannelValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetChannelValue")
	}

	return _child, nil
}

func (pm *_BACnetChannelValue) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetChannelValue, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetChannelValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetChannelValue")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}
	// Virtual field
	peekedIsContextTag := m.GetPeekedIsContextTag()
	_ = peekedIsContextTag
	if _peekedIsContextTagErr := writeBuffer.WriteVirtual(ctx, "peekedIsContextTag", m.GetPeekedIsContextTag()); _peekedIsContextTagErr != nil {
		return errors.Wrap(_peekedIsContextTagErr, "Error serializing 'peekedIsContextTag' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetChannelValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetChannelValue")
	}
	return nil
}

func (m *_BACnetChannelValue) IsBACnetChannelValue() {}
