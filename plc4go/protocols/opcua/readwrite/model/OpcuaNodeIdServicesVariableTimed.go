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

// OpcuaNodeIdServicesVariableTimed is an enum
type OpcuaNodeIdServicesVariableTimed int32

type IOpcuaNodeIdServicesVariableTimed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableTimed_TimedShelveMethodType_InputArguments  OpcuaNodeIdServicesVariableTimed = 6103
	OpcuaNodeIdServicesVariableTimed_TimedShelve2MethodType_InputArguments OpcuaNodeIdServicesVariableTimed = 25158
)

var OpcuaNodeIdServicesVariableTimedValues []OpcuaNodeIdServicesVariableTimed

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableTimedValues = []OpcuaNodeIdServicesVariableTimed{
		OpcuaNodeIdServicesVariableTimed_TimedShelveMethodType_InputArguments,
		OpcuaNodeIdServicesVariableTimed_TimedShelve2MethodType_InputArguments,
	}
}

func OpcuaNodeIdServicesVariableTimedByValue(value int32) (enum OpcuaNodeIdServicesVariableTimed, ok bool) {
	switch value {
	case 25158:
		return OpcuaNodeIdServicesVariableTimed_TimedShelve2MethodType_InputArguments, true
	case 6103:
		return OpcuaNodeIdServicesVariableTimed_TimedShelveMethodType_InputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableTimedByName(value string) (enum OpcuaNodeIdServicesVariableTimed, ok bool) {
	switch value {
	case "TimedShelve2MethodType_InputArguments":
		return OpcuaNodeIdServicesVariableTimed_TimedShelve2MethodType_InputArguments, true
	case "TimedShelveMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableTimed_TimedShelveMethodType_InputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableTimedKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableTimedValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableTimed(structType any) OpcuaNodeIdServicesVariableTimed {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableTimed {
		if sOpcuaNodeIdServicesVariableTimed, ok := typ.(OpcuaNodeIdServicesVariableTimed); ok {
			return sOpcuaNodeIdServicesVariableTimed
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableTimed) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableTimed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableTimedParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableTimed, error) {
	return OpcuaNodeIdServicesVariableTimedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableTimedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableTimed, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableTimed", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableTimed")
	}
	if enum, ok := OpcuaNodeIdServicesVariableTimedByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableTimed")
		return OpcuaNodeIdServicesVariableTimed(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableTimed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableTimed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableTimed", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e OpcuaNodeIdServicesVariableTimed) GetValue() int32 {
	return int32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableTimed) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableTimed_TimedShelve2MethodType_InputArguments:
		return "TimedShelve2MethodType_InputArguments"
	case OpcuaNodeIdServicesVariableTimed_TimedShelveMethodType_InputArguments:
		return "TimedShelveMethodType_InputArguments"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableTimed) String() string {
	return e.PLC4XEnumName()
}
