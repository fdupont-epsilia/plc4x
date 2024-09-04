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

// Code generated by "plc4xgenerator -type=DefaultPlcWriteResponse"; DO NOT EDIT.

package model

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

var _ = fmt.Printf

func (d *DefaultPlcWriteResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := d.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (d *DefaultPlcWriteResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	if err := writeBuffer.PushContext("PlcWriteResponse"); err != nil {
		return err
	}

	if any(d.request) != nil {
		if serializableField, ok := any(d.request).(utils.Serializable); ok {
			if err := writeBuffer.PushContext("request"); err != nil {
				return err
			}
			if err := serializableField.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext("request"); err != nil {
				return err
			}
		} else {
			stringValue := fmt.Sprintf("%v", d.request)
			if err := writeBuffer.WriteString("request", uint32(len(stringValue)*8), stringValue); err != nil {
				return err
			}
		}
	}
	if err := writeBuffer.PushContext("responseCodes", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	for _name, elem := range d.responseCodes {
		name := _name

		var elem any = elem
		if serializable, ok := elem.(utils.Serializable); ok {
			if err := writeBuffer.PushContext(name); err != nil {
				return err
			}
			if err := serializable.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext(name); err != nil {
				return err
			}
		} else {
			elemAsString := fmt.Sprintf("%v", elem)
			if err := writeBuffer.WriteString(name, uint32(len(elemAsString)*8), elemAsString); err != nil {
				return err
			}
		}
	}
	if err := writeBuffer.PopContext("responseCodes", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	if err := writeBuffer.PopContext("PlcWriteResponse"); err != nil {
		return err
	}
	return nil
}

func (d *DefaultPlcWriteResponse) String() string {
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), d); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
