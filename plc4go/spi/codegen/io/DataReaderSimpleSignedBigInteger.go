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

package io

import (
	"context"
	"math/big"

	"github.com/apache/plc4x/plc4go/spi/utils"
)

type DataReaderSimpleSignedBigInteger struct {
	*DataReaderSimpleBase[*big.Int]
}

var _ DataReader[*big.Int] = (*DataReaderSimpleSignedBigInteger)(nil)

func NewDataReaderSimpleSignedBigInteger(readBuffer utils.ReadBuffer, bitLength uint8) *DataReaderSimpleSignedBigInteger {
	return &DataReaderSimpleSignedBigInteger{
		DataReaderSimpleBase: NewDataReaderSimpleBase[*big.Int](readBuffer, uint(bitLength)),
	}
}

func (d *DataReaderSimpleSignedBigInteger) Read(ctx context.Context, logicalName string, readerArgs ...utils.WithReaderArgs) (*big.Int, error) {
	return d.readBuffer.ReadBigInt(logicalName, uint64(d.bitLength), readerArgs...)
}