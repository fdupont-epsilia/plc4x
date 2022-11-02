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
	"encoding/binary"

	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// DataTransportSize is an enum
type DataTransportSize uint8

type IDataTransportSize interface {
	utils.Serializable
	SizeInBits() bool
}

const (
	DataTransportSize_NULL            DataTransportSize = 0x00
	DataTransportSize_BIT             DataTransportSize = 0x03
	DataTransportSize_BYTE_WORD_DWORD DataTransportSize = 0x04
	DataTransportSize_INTEGER         DataTransportSize = 0x05
	DataTransportSize_DINTEGER        DataTransportSize = 0x06
	DataTransportSize_REAL            DataTransportSize = 0x07
	DataTransportSize_OCTET_STRING    DataTransportSize = 0x09
)

var DataTransportSizeValues []DataTransportSize

func init() {
	_ = errors.New
	DataTransportSizeValues = []DataTransportSize{
		DataTransportSize_NULL,
		DataTransportSize_BIT,
		DataTransportSize_BYTE_WORD_DWORD,
		DataTransportSize_INTEGER,
		DataTransportSize_DINTEGER,
		DataTransportSize_REAL,
		DataTransportSize_OCTET_STRING,
	}
}

func (e DataTransportSize) SizeInBits() bool {
	switch e {
	case 0x00:
		{ /* '0x00' */
			return false
		}
	case 0x03:
		{ /* '0x03' */
			return true
		}
	case 0x04:
		{ /* '0x04' */
			return true
		}
	case 0x05:
		{ /* '0x05' */
			return true
		}
	case 0x06:
		{ /* '0x06' */
			return false
		}
	case 0x07:
		{ /* '0x07' */
			return false
		}
	case 0x09:
		{ /* '0x09' */
			return false
		}
	default:
		{
			return false
		}
	}
}

func DataTransportSizeFirstEnumForFieldSizeInBits(value bool) (DataTransportSize, error) {
	for _, sizeValue := range DataTransportSizeValues {
		if sizeValue.SizeInBits() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing SizeInBits not found", value)
}
func DataTransportSizeByValue(value uint8) (enum DataTransportSize, ok bool) {
	switch value {
	case 0x00:
		return DataTransportSize_NULL, true
	case 0x03:
		return DataTransportSize_BIT, true
	case 0x04:
		return DataTransportSize_BYTE_WORD_DWORD, true
	case 0x05:
		return DataTransportSize_INTEGER, true
	case 0x06:
		return DataTransportSize_DINTEGER, true
	case 0x07:
		return DataTransportSize_REAL, true
	case 0x09:
		return DataTransportSize_OCTET_STRING, true
	}
	return 0, false
}

func DataTransportSizeByName(value string) (enum DataTransportSize, ok bool) {
	switch value {
	case "NULL":
		return DataTransportSize_NULL, true
	case "BIT":
		return DataTransportSize_BIT, true
	case "BYTE_WORD_DWORD":
		return DataTransportSize_BYTE_WORD_DWORD, true
	case "INTEGER":
		return DataTransportSize_INTEGER, true
	case "DINTEGER":
		return DataTransportSize_DINTEGER, true
	case "REAL":
		return DataTransportSize_REAL, true
	case "OCTET_STRING":
		return DataTransportSize_OCTET_STRING, true
	}
	return 0, false
}

func DataTransportSizeKnows(value uint8) bool {
	for _, typeValue := range DataTransportSizeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastDataTransportSize(structType interface{}) DataTransportSize {
	castFunc := func(typ interface{}) DataTransportSize {
		if sDataTransportSize, ok := typ.(DataTransportSize); ok {
			return sDataTransportSize
		}
		return 0
	}
	return castFunc(structType)
}

func (m DataTransportSize) GetLengthInBits() uint16 {
	return 8
}

func (m DataTransportSize) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func DataTransportSizeParse(readBuffer utils.ReadBuffer) (DataTransportSize, error) {
	val, err := readBuffer.ReadUint8("DataTransportSize", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading DataTransportSize")
	}
	if enum, ok := DataTransportSizeByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return DataTransportSize(val), nil
	} else {
		return enum, nil
	}
}

func (e DataTransportSize) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e DataTransportSize) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("DataTransportSize", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e DataTransportSize) PLC4XEnumName() string {
	switch e {
	case DataTransportSize_NULL:
		return "NULL"
	case DataTransportSize_BIT:
		return "BIT"
	case DataTransportSize_BYTE_WORD_DWORD:
		return "BYTE_WORD_DWORD"
	case DataTransportSize_INTEGER:
		return "INTEGER"
	case DataTransportSize_DINTEGER:
		return "DINTEGER"
	case DataTransportSize_REAL:
		return "REAL"
	case DataTransportSize_OCTET_STRING:
		return "OCTET_STRING"
	}
	return ""
}

func (e DataTransportSize) String() string {
	return e.PLC4XEnumName()
}