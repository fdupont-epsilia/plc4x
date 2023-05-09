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
package org.apache.plc4x.java.s7.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class S7PayloadUserDataItemCyclicServicesSubscribeRequest extends S7PayloadUserDataItem
    implements Message {

  // Accessors for discriminator values.
  public Byte getCpuFunctionGroup() {
    return (byte) 0x02;
  }

  public Byte getCpuFunctionType() {
    return (byte) 0x04;
  }

  public Short getCpuSubfunction() {
    return (short) 0x01;
  }

  // Properties.
  protected final int itemsCount;
  protected final TimeBase timeBase;
  protected final short timeFactor;
  protected final List<CycServiceItemType> item;

  public S7PayloadUserDataItemCyclicServicesSubscribeRequest(
      DataTransportErrorCode returnCode,
      DataTransportSize transportSize,
      int dataLength,
      int itemsCount,
      TimeBase timeBase,
      short timeFactor,
      List<CycServiceItemType> item) {
    super(returnCode, transportSize, dataLength);
    this.itemsCount = itemsCount;
    this.timeBase = timeBase;
    this.timeFactor = timeFactor;
    this.item = item;
  }

  public int getItemsCount() {
    return itemsCount;
  }

  public TimeBase getTimeBase() {
    return timeBase;
  }

  public short getTimeFactor() {
    return timeFactor;
  }

  public List<CycServiceItemType> getItem() {
    return item;
  }

  @Override
  protected void serializeS7PayloadUserDataItemChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("S7PayloadUserDataItemCyclicServicesSubscribeRequest");

    // Simple Field (itemsCount)
    writeSimpleField("itemsCount", itemsCount, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (timeBase)
    writeSimpleEnumField(
        "timeBase",
        "TimeBase",
        timeBase,
        new DataWriterEnumDefault<>(
            TimeBase::getValue, TimeBase::name, writeUnsignedShort(writeBuffer, 8)));

    // Simple Field (timeFactor)
    writeSimpleField("timeFactor", timeFactor, writeUnsignedShort(writeBuffer, 8));

    // Array Field (item)
    writeComplexTypeArrayField("item", item, writeBuffer);

    writeBuffer.popContext("S7PayloadUserDataItemCyclicServicesSubscribeRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    S7PayloadUserDataItemCyclicServicesSubscribeRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (itemsCount)
    lengthInBits += 16;

    // Simple field (timeBase)
    lengthInBits += 8;

    // Simple field (timeFactor)
    lengthInBits += 8;

    // Array field
    if (item != null) {
      int i = 0;
      for (CycServiceItemType element : item) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= item.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static S7PayloadUserDataItemBuilder staticParseS7PayloadUserDataItemBuilder(
      ReadBuffer readBuffer, Byte cpuFunctionGroup, Byte cpuFunctionType, Short cpuSubfunction)
      throws ParseException {
    readBuffer.pullContext("S7PayloadUserDataItemCyclicServicesSubscribeRequest");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int itemsCount = readSimpleField("itemsCount", readUnsignedInt(readBuffer, 16));

    TimeBase timeBase =
        readEnumField(
            "timeBase",
            "TimeBase",
            new DataReaderEnumDefault<>(TimeBase::enumForValue, readUnsignedShort(readBuffer, 8)));

    short timeFactor = readSimpleField("timeFactor", readUnsignedShort(readBuffer, 8));

    List<CycServiceItemType> item =
        readCountArrayField(
            "item",
            new DataReaderComplexDefault<>(
                () -> CycServiceItemType.staticParse(readBuffer), readBuffer),
            itemsCount);

    readBuffer.closeContext("S7PayloadUserDataItemCyclicServicesSubscribeRequest");
    // Create the instance
    return new S7PayloadUserDataItemCyclicServicesSubscribeRequestBuilderImpl(
        itemsCount, timeBase, timeFactor, item);
  }

  public static class S7PayloadUserDataItemCyclicServicesSubscribeRequestBuilderImpl
      implements S7PayloadUserDataItem.S7PayloadUserDataItemBuilder {
    private final int itemsCount;
    private final TimeBase timeBase;
    private final short timeFactor;
    private final List<CycServiceItemType> item;

    public S7PayloadUserDataItemCyclicServicesSubscribeRequestBuilderImpl(
        int itemsCount, TimeBase timeBase, short timeFactor, List<CycServiceItemType> item) {
      this.itemsCount = itemsCount;
      this.timeBase = timeBase;
      this.timeFactor = timeFactor;
      this.item = item;
    }

    public S7PayloadUserDataItemCyclicServicesSubscribeRequest build(
        DataTransportErrorCode returnCode, DataTransportSize transportSize, int dataLength) {
      S7PayloadUserDataItemCyclicServicesSubscribeRequest
          s7PayloadUserDataItemCyclicServicesSubscribeRequest =
              new S7PayloadUserDataItemCyclicServicesSubscribeRequest(
                  returnCode, transportSize, dataLength, itemsCount, timeBase, timeFactor, item);
      return s7PayloadUserDataItemCyclicServicesSubscribeRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof S7PayloadUserDataItemCyclicServicesSubscribeRequest)) {
      return false;
    }
    S7PayloadUserDataItemCyclicServicesSubscribeRequest that =
        (S7PayloadUserDataItemCyclicServicesSubscribeRequest) o;
    return (getItemsCount() == that.getItemsCount())
        && (getTimeBase() == that.getTimeBase())
        && (getTimeFactor() == that.getTimeFactor())
        && (getItem() == that.getItem())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getItemsCount(), getTimeBase(), getTimeFactor(), getItem());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}