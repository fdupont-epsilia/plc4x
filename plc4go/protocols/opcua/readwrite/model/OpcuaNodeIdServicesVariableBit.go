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

	"github.com/apache/plc4x/plc4go/spi/utils"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// OpcuaNodeIdServicesVariableBit is an enum
type OpcuaNodeIdServicesVariableBit int32

type IOpcuaNodeIdServicesVariableBit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableBit_BitFieldType_OptionalFieldNamee_Placeholder OpcuaNodeIdServicesVariableBit = 15014
	OpcuaNodeIdServicesVariableBit_BitFieldType_BitFieldsDefinitions           OpcuaNodeIdServicesVariableBit = 32432
	OpcuaNodeIdServicesVariableBit_BitFieldType_FieldName_Placeholder          OpcuaNodeIdServicesVariableBit = 32433
)

var OpcuaNodeIdServicesVariableBitValues []OpcuaNodeIdServicesVariableBit

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableBitValues = []OpcuaNodeIdServicesVariableBit{
		OpcuaNodeIdServicesVariableBit_BitFieldType_OptionalFieldNamee_Placeholder,
		OpcuaNodeIdServicesVariableBit_BitFieldType_BitFieldsDefinitions,
		OpcuaNodeIdServicesVariableBit_BitFieldType_FieldName_Placeholder,
	}
}

func OpcuaNodeIdServicesVariableBitByValue(value int32) (enum OpcuaNodeIdServicesVariableBit, ok bool) {
	switch value {
	case 15014:
		return OpcuaNodeIdServicesVariableBit_BitFieldType_OptionalFieldNamee_Placeholder, true
	case 32432:
		return OpcuaNodeIdServicesVariableBit_BitFieldType_BitFieldsDefinitions, true
	case 32433:
		return OpcuaNodeIdServicesVariableBit_BitFieldType_FieldName_Placeholder, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableBitByName(value string) (enum OpcuaNodeIdServicesVariableBit, ok bool) {
	switch value {
	case "BitFieldType_OptionalFieldNamee_Placeholder":
		return OpcuaNodeIdServicesVariableBit_BitFieldType_OptionalFieldNamee_Placeholder, true
	case "BitFieldType_BitFieldsDefinitions":
		return OpcuaNodeIdServicesVariableBit_BitFieldType_BitFieldsDefinitions, true
	case "BitFieldType_FieldName_Placeholder":
		return OpcuaNodeIdServicesVariableBit_BitFieldType_FieldName_Placeholder, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableBitKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableBitValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableBit(structType any) OpcuaNodeIdServicesVariableBit {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableBit {
		if sOpcuaNodeIdServicesVariableBit, ok := typ.(OpcuaNodeIdServicesVariableBit); ok {
			return sOpcuaNodeIdServicesVariableBit
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableBit) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableBit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableBitParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableBit, error) {
	return OpcuaNodeIdServicesVariableBitParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableBitParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableBit, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableBit", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableBit")
	}
	if enum, ok := OpcuaNodeIdServicesVariableBitByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableBit")
		return OpcuaNodeIdServicesVariableBit(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableBit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableBit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableBit", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e OpcuaNodeIdServicesVariableBit) GetValue() int32 {
	return int32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableBit) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableBit_BitFieldType_OptionalFieldNamee_Placeholder:
		return "BitFieldType_OptionalFieldNamee_Placeholder"
	case OpcuaNodeIdServicesVariableBit_BitFieldType_BitFieldsDefinitions:
		return "BitFieldType_BitFieldsDefinitions"
	case OpcuaNodeIdServicesVariableBit_BitFieldType_FieldName_Placeholder:
		return "BitFieldType_FieldName_Placeholder"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableBit) String() string {
	return e.PLC4XEnumName()
}
