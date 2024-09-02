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

// CurrencyUnitType is the corresponding interface of CurrencyUnitType
type CurrencyUnitType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNumericCode returns NumericCode (property field)
	GetNumericCode() int16
	// GetExponent returns Exponent (property field)
	GetExponent() int8
	// GetAlphabeticCode returns AlphabeticCode (property field)
	GetAlphabeticCode() PascalString
	// GetCurrency returns Currency (property field)
	GetCurrency() LocalizedText
	// IsCurrencyUnitType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCurrencyUnitType()
}

// _CurrencyUnitType is the data-structure of this message
type _CurrencyUnitType struct {
	ExtensionObjectDefinitionContract
	NumericCode    int16
	Exponent       int8
	AlphabeticCode PascalString
	Currency       LocalizedText
}

var _ CurrencyUnitType = (*_CurrencyUnitType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_CurrencyUnitType)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CurrencyUnitType) GetIdentifier() string {
	return "23500"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CurrencyUnitType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CurrencyUnitType) GetNumericCode() int16 {
	return m.NumericCode
}

func (m *_CurrencyUnitType) GetExponent() int8 {
	return m.Exponent
}

func (m *_CurrencyUnitType) GetAlphabeticCode() PascalString {
	return m.AlphabeticCode
}

func (m *_CurrencyUnitType) GetCurrency() LocalizedText {
	return m.Currency
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCurrencyUnitType factory function for _CurrencyUnitType
func NewCurrencyUnitType(numericCode int16, exponent int8, alphabeticCode PascalString, currency LocalizedText) *_CurrencyUnitType {
	_result := &_CurrencyUnitType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		NumericCode:                       numericCode,
		Exponent:                          exponent,
		AlphabeticCode:                    alphabeticCode,
		Currency:                          currency,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCurrencyUnitType(structType any) CurrencyUnitType {
	if casted, ok := structType.(CurrencyUnitType); ok {
		return casted
	}
	if casted, ok := structType.(*CurrencyUnitType); ok {
		return *casted
	}
	return nil
}

func (m *_CurrencyUnitType) GetTypeName() string {
	return "CurrencyUnitType"
}

func (m *_CurrencyUnitType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (numericCode)
	lengthInBits += 16

	// Simple field (exponent)
	lengthInBits += 8

	// Simple field (alphabeticCode)
	lengthInBits += m.AlphabeticCode.GetLengthInBits(ctx)

	// Simple field (currency)
	lengthInBits += m.Currency.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CurrencyUnitType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CurrencyUnitType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__currencyUnitType CurrencyUnitType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CurrencyUnitType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CurrencyUnitType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	numericCode, err := ReadSimpleField(ctx, "numericCode", ReadSignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numericCode' field"))
	}
	m.NumericCode = numericCode

	exponent, err := ReadSimpleField(ctx, "exponent", ReadSignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'exponent' field"))
	}
	m.Exponent = exponent

	alphabeticCode, err := ReadSimpleField[PascalString](ctx, "alphabeticCode", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'alphabeticCode' field"))
	}
	m.AlphabeticCode = alphabeticCode

	currency, err := ReadSimpleField[LocalizedText](ctx, "currency", ReadComplex[LocalizedText](LocalizedTextParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'currency' field"))
	}
	m.Currency = currency

	if closeErr := readBuffer.CloseContext("CurrencyUnitType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CurrencyUnitType")
	}

	return m, nil
}

func (m *_CurrencyUnitType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CurrencyUnitType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CurrencyUnitType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CurrencyUnitType")
		}

		if err := WriteSimpleField[int16](ctx, "numericCode", m.GetNumericCode(), WriteSignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'numericCode' field")
		}

		if err := WriteSimpleField[int8](ctx, "exponent", m.GetExponent(), WriteSignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'exponent' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "alphabeticCode", m.GetAlphabeticCode(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'alphabeticCode' field")
		}

		if err := WriteSimpleField[LocalizedText](ctx, "currency", m.GetCurrency(), WriteComplex[LocalizedText](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'currency' field")
		}

		if popErr := writeBuffer.PopContext("CurrencyUnitType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CurrencyUnitType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CurrencyUnitType) IsCurrencyUnitType() {}

func (m *_CurrencyUnitType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
