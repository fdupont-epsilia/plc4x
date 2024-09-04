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

// Code generated by "plc4xGenerator -type=Configuration"; DO NOT EDIT.

package cbus

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

var _ = fmt.Printf

func (d *Configuration) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := d.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (d *Configuration) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	if err := writeBuffer.PushContext("Configuration"); err != nil {
		return err
	}

	if err := writeBuffer.WriteBit("srchk", d.Srchk); err != nil {
		return err
	}

	if err := writeBuffer.WriteBit("exstat", d.Exstat); err != nil {
		return err
	}

	if err := writeBuffer.WriteBit("pun", d.Pun); err != nil {
		return err
	}

	if err := writeBuffer.WriteBit("localSal", d.LocalSal); err != nil {
		return err
	}

	if err := writeBuffer.WriteBit("pcn", d.Pcn); err != nil {
		return err
	}

	if err := writeBuffer.WriteBit("idmon", d.Idmon); err != nil {
		return err
	}

	if err := writeBuffer.WriteBit("monitor", d.Monitor); err != nil {
		return err
	}

	if err := writeBuffer.WriteBit("smart", d.Smart); err != nil {
		return err
	}

	if err := writeBuffer.WriteBit("xonXoff", d.XonXoff); err != nil {
		return err
	}

	if err := writeBuffer.WriteBit("connect", d.Connect); err != nil {
		return err
	}

	if err := writeBuffer.WriteByte("monitoredApplication1", d.MonitoredApplication1); err != nil {
		return err
	}

	if err := writeBuffer.WriteByte("monitoredApplication2", d.MonitoredApplication2); err != nil {
		return err
	}
	if err := writeBuffer.PopContext("Configuration"); err != nil {
		return err
	}
	return nil
}

func (d *Configuration) String() string {
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), d); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
