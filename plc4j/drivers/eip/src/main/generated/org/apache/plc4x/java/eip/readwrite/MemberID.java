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
package org.apache.plc4x.java.eip.readwrite;

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

public class MemberID extends LogicalSegmentType implements Message {

  // Accessors for discriminator values.
  public Byte getLogicalSegmentType() {
    return (byte) 0x02;
  }

  // Properties.
  protected final byte format;
  protected final short instance;

  // Arguments.
  protected final IntegerEncoding order;

  public MemberID(byte format, short instance, IntegerEncoding order) {
    super(order);
    this.format = format;
    this.instance = instance;
    this.order = order;
  }

  public byte getFormat() {
    return format;
  }

  public short getInstance() {
    return instance;
  }

  @Override
  protected void serializeLogicalSegmentTypeChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("MemberID");

    // Simple Field (format)
    writeSimpleField(
        "format",
        format,
        writeUnsignedByte(writeBuffer, 2),
        WithOption.WithByteOrder(
            (((order) == (IntegerEncoding.BIG_ENDIAN))
                ? ByteOrder.BIG_ENDIAN
                : ByteOrder.LITTLE_ENDIAN)));

    // Simple Field (instance)
    writeSimpleField(
        "instance",
        instance,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(
            (((order) == (IntegerEncoding.BIG_ENDIAN))
                ? ByteOrder.BIG_ENDIAN
                : ByteOrder.LITTLE_ENDIAN)));

    writeBuffer.popContext("MemberID");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    MemberID _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (format)
    lengthInBits += 2;

    // Simple field (instance)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static LogicalSegmentTypeBuilder staticParseLogicalSegmentTypeBuilder(
      ReadBuffer readBuffer, IntegerEncoding order) throws ParseException {
    readBuffer.pullContext("MemberID");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte format =
        readSimpleField(
            "format",
            readUnsignedByte(readBuffer, 2),
            WithOption.WithByteOrder(
                (((order) == (IntegerEncoding.BIG_ENDIAN))
                    ? ByteOrder.BIG_ENDIAN
                    : ByteOrder.LITTLE_ENDIAN)));

    short instance =
        readSimpleField(
            "instance",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithByteOrder(
                (((order) == (IntegerEncoding.BIG_ENDIAN))
                    ? ByteOrder.BIG_ENDIAN
                    : ByteOrder.LITTLE_ENDIAN)));

    readBuffer.closeContext("MemberID");
    // Create the instance
    return new MemberIDBuilderImpl(format, instance, order);
  }

  public static class MemberIDBuilderImpl implements LogicalSegmentType.LogicalSegmentTypeBuilder {
    private final byte format;
    private final short instance;
    private final IntegerEncoding order;

    public MemberIDBuilderImpl(byte format, short instance, IntegerEncoding order) {
      this.format = format;
      this.instance = instance;
      this.order = order;
    }

    public MemberID build(IntegerEncoding order) {

      MemberID memberID = new MemberID(format, instance, order);
      return memberID;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof MemberID)) {
      return false;
    }
    MemberID that = (MemberID) o;
    return (getFormat() == that.getFormat())
        && (getInstance() == that.getInstance())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getFormat(), getInstance());
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