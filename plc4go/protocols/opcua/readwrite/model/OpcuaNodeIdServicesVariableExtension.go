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

// OpcuaNodeIdServicesVariableExtension is an enum
type OpcuaNodeIdServicesVariableExtension int32

type IOpcuaNodeIdServicesVariableExtension interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableExtension_ExtensionFieldsType_ExtensionFieldName_Placeholder      OpcuaNodeIdServicesVariableExtension = 15490
	OpcuaNodeIdServicesVariableExtension_ExtensionFieldsType_AddExtensionField_InputArguments    OpcuaNodeIdServicesVariableExtension = 15492
	OpcuaNodeIdServicesVariableExtension_ExtensionFieldsType_AddExtensionField_OutputArguments   OpcuaNodeIdServicesVariableExtension = 15493
	OpcuaNodeIdServicesVariableExtension_ExtensionFieldsType_RemoveExtensionField_InputArguments OpcuaNodeIdServicesVariableExtension = 15495
)

var OpcuaNodeIdServicesVariableExtensionValues []OpcuaNodeIdServicesVariableExtension

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableExtensionValues = []OpcuaNodeIdServicesVariableExtension{
		OpcuaNodeIdServicesVariableExtension_ExtensionFieldsType_ExtensionFieldName_Placeholder,
		OpcuaNodeIdServicesVariableExtension_ExtensionFieldsType_AddExtensionField_InputArguments,
		OpcuaNodeIdServicesVariableExtension_ExtensionFieldsType_AddExtensionField_OutputArguments,
		OpcuaNodeIdServicesVariableExtension_ExtensionFieldsType_RemoveExtensionField_InputArguments,
	}
}

func OpcuaNodeIdServicesVariableExtensionByValue(value int32) (enum OpcuaNodeIdServicesVariableExtension, ok bool) {
	switch value {
	case 15490:
		return OpcuaNodeIdServicesVariableExtension_ExtensionFieldsType_ExtensionFieldName_Placeholder, true
	case 15492:
		return OpcuaNodeIdServicesVariableExtension_ExtensionFieldsType_AddExtensionField_InputArguments, true
	case 15493:
		return OpcuaNodeIdServicesVariableExtension_ExtensionFieldsType_AddExtensionField_OutputArguments, true
	case 15495:
		return OpcuaNodeIdServicesVariableExtension_ExtensionFieldsType_RemoveExtensionField_InputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableExtensionByName(value string) (enum OpcuaNodeIdServicesVariableExtension, ok bool) {
	switch value {
	case "ExtensionFieldsType_ExtensionFieldName_Placeholder":
		return OpcuaNodeIdServicesVariableExtension_ExtensionFieldsType_ExtensionFieldName_Placeholder, true
	case "ExtensionFieldsType_AddExtensionField_InputArguments":
		return OpcuaNodeIdServicesVariableExtension_ExtensionFieldsType_AddExtensionField_InputArguments, true
	case "ExtensionFieldsType_AddExtensionField_OutputArguments":
		return OpcuaNodeIdServicesVariableExtension_ExtensionFieldsType_AddExtensionField_OutputArguments, true
	case "ExtensionFieldsType_RemoveExtensionField_InputArguments":
		return OpcuaNodeIdServicesVariableExtension_ExtensionFieldsType_RemoveExtensionField_InputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableExtensionKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableExtensionValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableExtension(structType any) OpcuaNodeIdServicesVariableExtension {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableExtension {
		if sOpcuaNodeIdServicesVariableExtension, ok := typ.(OpcuaNodeIdServicesVariableExtension); ok {
			return sOpcuaNodeIdServicesVariableExtension
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableExtension) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableExtension) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableExtensionParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableExtension, error) {
	return OpcuaNodeIdServicesVariableExtensionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableExtensionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableExtension, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableExtension", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableExtension")
	}
	if enum, ok := OpcuaNodeIdServicesVariableExtensionByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableExtension")
		return OpcuaNodeIdServicesVariableExtension(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableExtension) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableExtension) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableExtension", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e OpcuaNodeIdServicesVariableExtension) GetValue() int32 {
	return int32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableExtension) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableExtension_ExtensionFieldsType_ExtensionFieldName_Placeholder:
		return "ExtensionFieldsType_ExtensionFieldName_Placeholder"
	case OpcuaNodeIdServicesVariableExtension_ExtensionFieldsType_AddExtensionField_InputArguments:
		return "ExtensionFieldsType_AddExtensionField_InputArguments"
	case OpcuaNodeIdServicesVariableExtension_ExtensionFieldsType_AddExtensionField_OutputArguments:
		return "ExtensionFieldsType_AddExtensionField_OutputArguments"
	case OpcuaNodeIdServicesVariableExtension_ExtensionFieldsType_RemoveExtensionField_InputArguments:
		return "ExtensionFieldsType_RemoveExtensionField_InputArguments"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableExtension) String() string {
	return e.PLC4XEnumName()
}
