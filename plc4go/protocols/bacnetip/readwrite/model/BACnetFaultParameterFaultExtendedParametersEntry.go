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

// BACnetFaultParameterFaultExtendedParametersEntry is the corresponding interface of BACnetFaultParameterFaultExtendedParametersEntry
type BACnetFaultParameterFaultExtendedParametersEntry interface {
	BACnetFaultParameterFaultExtendedParametersEntryContract
	BACnetFaultParameterFaultExtendedParametersEntryRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// BACnetFaultParameterFaultExtendedParametersEntryContract provides a set of functions which can be overwritten by a sub struct
type BACnetFaultParameterFaultExtendedParametersEntryContract interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetPeekedIsContextTag returns PeekedIsContextTag (virtual field)
	GetPeekedIsContextTag() bool
}

// BACnetFaultParameterFaultExtendedParametersEntryRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetFaultParameterFaultExtendedParametersEntryRequirements interface {
}

// BACnetFaultParameterFaultExtendedParametersEntryExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultExtendedParametersEntry.
// This is useful for switch cases.
type BACnetFaultParameterFaultExtendedParametersEntryExactly interface {
	BACnetFaultParameterFaultExtendedParametersEntry
	isBACnetFaultParameterFaultExtendedParametersEntry() bool
}

// _BACnetFaultParameterFaultExtendedParametersEntry is the data-structure of this message
type _BACnetFaultParameterFaultExtendedParametersEntry struct {
	_BACnetFaultParameterFaultExtendedParametersEntryChildRequirements
	PeekedTagHeader BACnetTagHeader
}

var _ BACnetFaultParameterFaultExtendedParametersEntryContract = (*_BACnetFaultParameterFaultExtendedParametersEntry)(nil)

type _BACnetFaultParameterFaultExtendedParametersEntryChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetPeekedIsContextTag() bool
	GetPeekedTagNumber() uint8
}

type BACnetFaultParameterFaultExtendedParametersEntryChild interface {
	utils.Serializable

	GetParent() *BACnetFaultParameterFaultExtendedParametersEntry

	GetTypeName() string
	BACnetFaultParameterFaultExtendedParametersEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultExtendedParametersEntry) GetPeekedTagHeader() BACnetTagHeader {
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

func (m *_BACnetFaultParameterFaultExtendedParametersEntry) GetPeekedTagNumber() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntry) GetPeekedIsContextTag() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetPeekedTagHeader().GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultExtendedParametersEntry factory function for _BACnetFaultParameterFaultExtendedParametersEntry
func NewBACnetFaultParameterFaultExtendedParametersEntry(peekedTagHeader BACnetTagHeader) *_BACnetFaultParameterFaultExtendedParametersEntry {
	return &_BACnetFaultParameterFaultExtendedParametersEntry{PeekedTagHeader: peekedTagHeader}
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultExtendedParametersEntry(structType any) BACnetFaultParameterFaultExtendedParametersEntry {
	if casted, ok := structType.(BACnetFaultParameterFaultExtendedParametersEntry); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtendedParametersEntry); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntry) GetTypeName() string {
	return "BACnetFaultParameterFaultExtendedParametersEntry"
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntry) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntry) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetFaultParameterFaultExtendedParametersEntryParse[T BACnetFaultParameterFaultExtendedParametersEntry](ctx context.Context, theBytes []byte) (T, error) {
	return BACnetFaultParameterFaultExtendedParametersEntryParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetFaultParameterFaultExtendedParametersEntryParseWithBufferProducer[T BACnetFaultParameterFaultExtendedParametersEntry]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetFaultParameterFaultExtendedParametersEntryParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func BACnetFaultParameterFaultExtendedParametersEntryParseWithBuffer[T BACnetFaultParameterFaultExtendedParametersEntry](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_BACnetFaultParameterFaultExtendedParametersEntry{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntry) parse(ctx context.Context, readBuffer utils.ReadBuffer) (_bACnetFaultParameterFaultExtendedParametersEntry BACnetFaultParameterFaultExtendedParametersEntry, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultExtendedParametersEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultExtendedParametersEntry")
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
	var _child BACnetFaultParameterFaultExtendedParametersEntry
	switch {
	case peekedTagNumber == 0x0 && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryNull
		if _child, err = (&_BACnetFaultParameterFaultExtendedParametersEntryNull{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetFaultParameterFaultExtendedParametersEntryNull for type-switch of BACnetFaultParameterFaultExtendedParametersEntry")
		}
	case peekedTagNumber == 0x4 && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryReal
		if _child, err = (&_BACnetFaultParameterFaultExtendedParametersEntryReal{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetFaultParameterFaultExtendedParametersEntryReal for type-switch of BACnetFaultParameterFaultExtendedParametersEntry")
		}
	case peekedTagNumber == 0x2 && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryUnsigned
		if _child, err = (&_BACnetFaultParameterFaultExtendedParametersEntryUnsigned{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetFaultParameterFaultExtendedParametersEntryUnsigned for type-switch of BACnetFaultParameterFaultExtendedParametersEntry")
		}
	case peekedTagNumber == 0x1 && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryBoolean
		if _child, err = (&_BACnetFaultParameterFaultExtendedParametersEntryBoolean{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetFaultParameterFaultExtendedParametersEntryBoolean for type-switch of BACnetFaultParameterFaultExtendedParametersEntry")
		}
	case peekedTagNumber == 0x3 && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryInteger
		if _child, err = (&_BACnetFaultParameterFaultExtendedParametersEntryInteger{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetFaultParameterFaultExtendedParametersEntryInteger for type-switch of BACnetFaultParameterFaultExtendedParametersEntry")
		}
	case peekedTagNumber == 0x5 && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryDouble
		if _child, err = (&_BACnetFaultParameterFaultExtendedParametersEntryDouble{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetFaultParameterFaultExtendedParametersEntryDouble for type-switch of BACnetFaultParameterFaultExtendedParametersEntry")
		}
	case peekedTagNumber == 0x6 && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryOctetString
		if _child, err = (&_BACnetFaultParameterFaultExtendedParametersEntryOctetString{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetFaultParameterFaultExtendedParametersEntryOctetString for type-switch of BACnetFaultParameterFaultExtendedParametersEntry")
		}
	case peekedTagNumber == 0x7 && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryCharacterString
		if _child, err = (&_BACnetFaultParameterFaultExtendedParametersEntryCharacterString{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetFaultParameterFaultExtendedParametersEntryCharacterString for type-switch of BACnetFaultParameterFaultExtendedParametersEntry")
		}
	case peekedTagNumber == 0x8 && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryBitString
		if _child, err = (&_BACnetFaultParameterFaultExtendedParametersEntryBitString{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetFaultParameterFaultExtendedParametersEntryBitString for type-switch of BACnetFaultParameterFaultExtendedParametersEntry")
		}
	case peekedTagNumber == 0x9 && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryEnumerated
		if _child, err = (&_BACnetFaultParameterFaultExtendedParametersEntryEnumerated{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetFaultParameterFaultExtendedParametersEntryEnumerated for type-switch of BACnetFaultParameterFaultExtendedParametersEntry")
		}
	case peekedTagNumber == 0xA && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryDate
		if _child, err = (&_BACnetFaultParameterFaultExtendedParametersEntryDate{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetFaultParameterFaultExtendedParametersEntryDate for type-switch of BACnetFaultParameterFaultExtendedParametersEntry")
		}
	case peekedTagNumber == 0xB && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryTime
		if _child, err = (&_BACnetFaultParameterFaultExtendedParametersEntryTime{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetFaultParameterFaultExtendedParametersEntryTime for type-switch of BACnetFaultParameterFaultExtendedParametersEntry")
		}
	case peekedTagNumber == 0xC && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryObjectidentifier
		if _child, err = (&_BACnetFaultParameterFaultExtendedParametersEntryObjectidentifier{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetFaultParameterFaultExtendedParametersEntryObjectidentifier for type-switch of BACnetFaultParameterFaultExtendedParametersEntry")
		}
	case peekedTagNumber == uint8(0) && peekedIsContextTag == bool(true): // BACnetFaultParameterFaultExtendedParametersEntryReference
		if _child, err = (&_BACnetFaultParameterFaultExtendedParametersEntryReference{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetFaultParameterFaultExtendedParametersEntryReference for type-switch of BACnetFaultParameterFaultExtendedParametersEntry")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v, peekedIsContextTag=%v]", peekedTagNumber, peekedIsContextTag)
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultExtendedParametersEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultExtendedParametersEntry")
	}

	return _child, nil
}

func (pm *_BACnetFaultParameterFaultExtendedParametersEntry) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetFaultParameterFaultExtendedParametersEntry, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultExtendedParametersEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultExtendedParametersEntry")
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

	if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultExtendedParametersEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultExtendedParametersEntry")
	}
	return nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntry) isBACnetFaultParameterFaultExtendedParametersEntry() bool {
	return true
}
