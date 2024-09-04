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

// Code generated by "plc4xGenerator -type=WhoIsIAmServices -prefix=device_"; DO NOT EDIT.

package bacgopes

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

var _ = fmt.Printf

func (d *WhoIsIAmServices) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := d.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (d *WhoIsIAmServices) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	if err := writeBuffer.PushContext("WhoIsIAmServices"); err != nil {
		return err
	}
	{
		_value := fmt.Sprintf("%v", d._requirements)

		if err := writeBuffer.WriteString("_requirements", uint32(len(_value)*8), _value); err != nil {
			return err
		}
	}
	if err := d.Capability.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
		return err
	}
	{
		_value := fmt.Sprintf("%v", d.localDevice)

		if err := writeBuffer.WriteString("localDevice", uint32(len(_value)*8), _value); err != nil {
			return err
		}
	}
	if err := writeBuffer.PopContext("WhoIsIAmServices"); err != nil {
		return err
	}
	return nil
}

func (d *WhoIsIAmServices) String() string {
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), d); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
