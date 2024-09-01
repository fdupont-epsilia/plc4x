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

// OpcuaNodeIdServicesVariableWrite is an enum
type OpcuaNodeIdServicesVariableWrite int32

type IOpcuaNodeIdServicesVariableWrite interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableWrite_WriteMethodType_InputArguments OpcuaNodeIdServicesVariableWrite = 11747
)

var OpcuaNodeIdServicesVariableWriteValues []OpcuaNodeIdServicesVariableWrite

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableWriteValues = []OpcuaNodeIdServicesVariableWrite{
		OpcuaNodeIdServicesVariableWrite_WriteMethodType_InputArguments,
	}
}

func OpcuaNodeIdServicesVariableWriteByValue(value int32) (enum OpcuaNodeIdServicesVariableWrite, ok bool) {
	switch value {
	case 11747:
		return OpcuaNodeIdServicesVariableWrite_WriteMethodType_InputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableWriteByName(value string) (enum OpcuaNodeIdServicesVariableWrite, ok bool) {
	switch value {
	case "WriteMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableWrite_WriteMethodType_InputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableWriteKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableWriteValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableWrite(structType any) OpcuaNodeIdServicesVariableWrite {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableWrite {
		if sOpcuaNodeIdServicesVariableWrite, ok := typ.(OpcuaNodeIdServicesVariableWrite); ok {
			return sOpcuaNodeIdServicesVariableWrite
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableWrite) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableWrite) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableWriteParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableWrite, error) {
	return OpcuaNodeIdServicesVariableWriteParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableWriteParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableWrite, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableWrite", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableWrite")
	}
	if enum, ok := OpcuaNodeIdServicesVariableWriteByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableWrite")
		return OpcuaNodeIdServicesVariableWrite(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableWrite) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableWrite) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableWrite", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e OpcuaNodeIdServicesVariableWrite) GetValue() int32 {
	return int32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableWrite) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableWrite_WriteMethodType_InputArguments:
		return "WriteMethodType_InputArguments"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableWrite) String() string {
	return e.PLC4XEnumName()
}
