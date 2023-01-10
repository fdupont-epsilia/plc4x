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
package org.apache.plc4x.java.bacnetip.readwrite;

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

public class BACnetPropertyStatesRestartReason extends BACnetPropertyStates implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetRestartReasonTagged restartReason;

  public BACnetPropertyStatesRestartReason(
      BACnetTagHeader peekedTagHeader, BACnetRestartReasonTagged restartReason) {
    super(peekedTagHeader);
    this.restartReason = restartReason;
  }

  public BACnetRestartReasonTagged getRestartReason() {
    return restartReason;
  }

  @Override
  protected void serializeBACnetPropertyStatesChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetPropertyStatesRestartReason");

    // Simple Field (restartReason)
    writeSimpleField("restartReason", restartReason, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetPropertyStatesRestartReason");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetPropertyStatesRestartReason _value = this;

    // Simple field (restartReason)
    lengthInBits += restartReason.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetPropertyStatesRestartReasonBuilder staticParseBuilder(
      ReadBuffer readBuffer, Short peekedTagNumber) throws ParseException {
    readBuffer.pullContext("BACnetPropertyStatesRestartReason");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    BACnetRestartReasonTagged restartReason =
        readSimpleField(
            "restartReason",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetRestartReasonTagged.staticParse(
                        readBuffer,
                        (short) (peekedTagNumber),
                        (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    readBuffer.closeContext("BACnetPropertyStatesRestartReason");
    // Create the instance
    return new BACnetPropertyStatesRestartReasonBuilder(restartReason);
  }

  public static class BACnetPropertyStatesRestartReasonBuilder
      implements BACnetPropertyStates.BACnetPropertyStatesBuilder {
    private final BACnetRestartReasonTagged restartReason;

    public BACnetPropertyStatesRestartReasonBuilder(BACnetRestartReasonTagged restartReason) {

      this.restartReason = restartReason;
    }

    public BACnetPropertyStatesRestartReason build(BACnetTagHeader peekedTagHeader) {
      BACnetPropertyStatesRestartReason bACnetPropertyStatesRestartReason =
          new BACnetPropertyStatesRestartReason(peekedTagHeader, restartReason);
      return bACnetPropertyStatesRestartReason;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetPropertyStatesRestartReason)) {
      return false;
    }
    BACnetPropertyStatesRestartReason that = (BACnetPropertyStatesRestartReason) o;
    return (getRestartReason() == that.getRestartReason()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getRestartReason());
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