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

// BACnetDaysOfWeek is an enum
type BACnetDaysOfWeek uint8

type IBACnetDaysOfWeek interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetDaysOfWeek_MONDAY    BACnetDaysOfWeek = 0
	BACnetDaysOfWeek_TUESDAY   BACnetDaysOfWeek = 1
	BACnetDaysOfWeek_WEDNESDAY BACnetDaysOfWeek = 2
	BACnetDaysOfWeek_THURSDAY  BACnetDaysOfWeek = 3
	BACnetDaysOfWeek_FRIDAY    BACnetDaysOfWeek = 4
	BACnetDaysOfWeek_SATURDAY  BACnetDaysOfWeek = 5
	BACnetDaysOfWeek_SUNDAY    BACnetDaysOfWeek = 6
)

var BACnetDaysOfWeekValues []BACnetDaysOfWeek

func init() {
	_ = errors.New
	BACnetDaysOfWeekValues = []BACnetDaysOfWeek{
		BACnetDaysOfWeek_MONDAY,
		BACnetDaysOfWeek_TUESDAY,
		BACnetDaysOfWeek_WEDNESDAY,
		BACnetDaysOfWeek_THURSDAY,
		BACnetDaysOfWeek_FRIDAY,
		BACnetDaysOfWeek_SATURDAY,
		BACnetDaysOfWeek_SUNDAY,
	}
}

func BACnetDaysOfWeekByValue(value uint8) (enum BACnetDaysOfWeek, ok bool) {
	switch value {
	case 0:
		return BACnetDaysOfWeek_MONDAY, true
	case 1:
		return BACnetDaysOfWeek_TUESDAY, true
	case 2:
		return BACnetDaysOfWeek_WEDNESDAY, true
	case 3:
		return BACnetDaysOfWeek_THURSDAY, true
	case 4:
		return BACnetDaysOfWeek_FRIDAY, true
	case 5:
		return BACnetDaysOfWeek_SATURDAY, true
	case 6:
		return BACnetDaysOfWeek_SUNDAY, true
	}
	return 0, false
}

func BACnetDaysOfWeekByName(value string) (enum BACnetDaysOfWeek, ok bool) {
	switch value {
	case "MONDAY":
		return BACnetDaysOfWeek_MONDAY, true
	case "TUESDAY":
		return BACnetDaysOfWeek_TUESDAY, true
	case "WEDNESDAY":
		return BACnetDaysOfWeek_WEDNESDAY, true
	case "THURSDAY":
		return BACnetDaysOfWeek_THURSDAY, true
	case "FRIDAY":
		return BACnetDaysOfWeek_FRIDAY, true
	case "SATURDAY":
		return BACnetDaysOfWeek_SATURDAY, true
	case "SUNDAY":
		return BACnetDaysOfWeek_SUNDAY, true
	}
	return 0, false
}

func BACnetDaysOfWeekKnows(value uint8) bool {
	for _, typeValue := range BACnetDaysOfWeekValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetDaysOfWeek(structType any) BACnetDaysOfWeek {
	castFunc := func(typ any) BACnetDaysOfWeek {
		if sBACnetDaysOfWeek, ok := typ.(BACnetDaysOfWeek); ok {
			return sBACnetDaysOfWeek
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetDaysOfWeek) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetDaysOfWeek) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetDaysOfWeekParse(ctx context.Context, theBytes []byte) (BACnetDaysOfWeek, error) {
	return BACnetDaysOfWeekParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetDaysOfWeekParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDaysOfWeek, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("BACnetDaysOfWeek", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetDaysOfWeek")
	}
	if enum, ok := BACnetDaysOfWeekByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetDaysOfWeek")
		return BACnetDaysOfWeek(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetDaysOfWeek) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetDaysOfWeek) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint8("BACnetDaysOfWeek", 8, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e BACnetDaysOfWeek) GetValue() uint8 {
	return uint8(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetDaysOfWeek) PLC4XEnumName() string {
	switch e {
	case BACnetDaysOfWeek_MONDAY:
		return "MONDAY"
	case BACnetDaysOfWeek_TUESDAY:
		return "TUESDAY"
	case BACnetDaysOfWeek_WEDNESDAY:
		return "WEDNESDAY"
	case BACnetDaysOfWeek_THURSDAY:
		return "THURSDAY"
	case BACnetDaysOfWeek_FRIDAY:
		return "FRIDAY"
	case BACnetDaysOfWeek_SATURDAY:
		return "SATURDAY"
	case BACnetDaysOfWeek_SUNDAY:
		return "SUNDAY"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e BACnetDaysOfWeek) String() string {
	return e.PLC4XEnumName()
}
