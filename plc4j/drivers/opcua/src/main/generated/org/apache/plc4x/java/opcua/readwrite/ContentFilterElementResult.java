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
package org.apache.plc4x.java.opcua.readwrite;

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

public class ContentFilterElementResult extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "606";
  }

  // Properties.
  protected final StatusCode statusCode;
  protected final int noOfOperandStatusCodes;
  protected final List<StatusCode> operandStatusCodes;
  protected final int noOfOperandDiagnosticInfos;
  protected final List<DiagnosticInfo> operandDiagnosticInfos;

  public ContentFilterElementResult(
      StatusCode statusCode,
      int noOfOperandStatusCodes,
      List<StatusCode> operandStatusCodes,
      int noOfOperandDiagnosticInfos,
      List<DiagnosticInfo> operandDiagnosticInfos) {
    super();
    this.statusCode = statusCode;
    this.noOfOperandStatusCodes = noOfOperandStatusCodes;
    this.operandStatusCodes = operandStatusCodes;
    this.noOfOperandDiagnosticInfos = noOfOperandDiagnosticInfos;
    this.operandDiagnosticInfos = operandDiagnosticInfos;
  }

  public StatusCode getStatusCode() {
    return statusCode;
  }

  public int getNoOfOperandStatusCodes() {
    return noOfOperandStatusCodes;
  }

  public List<StatusCode> getOperandStatusCodes() {
    return operandStatusCodes;
  }

  public int getNoOfOperandDiagnosticInfos() {
    return noOfOperandDiagnosticInfos;
  }

  public List<DiagnosticInfo> getOperandDiagnosticInfos() {
    return operandDiagnosticInfos;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("ContentFilterElementResult");

    // Simple Field (statusCode)
    writeSimpleField("statusCode", statusCode, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (noOfOperandStatusCodes)
    writeSimpleField(
        "noOfOperandStatusCodes", noOfOperandStatusCodes, writeSignedInt(writeBuffer, 32));

    // Array Field (operandStatusCodes)
    writeComplexTypeArrayField("operandStatusCodes", operandStatusCodes, writeBuffer);

    // Simple Field (noOfOperandDiagnosticInfos)
    writeSimpleField(
        "noOfOperandDiagnosticInfos", noOfOperandDiagnosticInfos, writeSignedInt(writeBuffer, 32));

    // Array Field (operandDiagnosticInfos)
    writeComplexTypeArrayField("operandDiagnosticInfos", operandDiagnosticInfos, writeBuffer);

    writeBuffer.popContext("ContentFilterElementResult");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ContentFilterElementResult _value = this;

    // Simple field (statusCode)
    lengthInBits += statusCode.getLengthInBits();

    // Simple field (noOfOperandStatusCodes)
    lengthInBits += 32;

    // Array field
    if (operandStatusCodes != null) {
      int i = 0;
      for (StatusCode element : operandStatusCodes) {
        boolean last = ++i >= operandStatusCodes.size();
        lengthInBits += element.getLengthInBits();
      }
    }

    // Simple field (noOfOperandDiagnosticInfos)
    lengthInBits += 32;

    // Array field
    if (operandDiagnosticInfos != null) {
      int i = 0;
      for (DiagnosticInfo element : operandDiagnosticInfos) {
        boolean last = ++i >= operandDiagnosticInfos.size();
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ContentFilterElementResultBuilder staticParseBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("ContentFilterElementResult");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    StatusCode statusCode =
        readSimpleField(
            "statusCode",
            new DataReaderComplexDefault<>(() -> StatusCode.staticParse(readBuffer), readBuffer));

    int noOfOperandStatusCodes =
        readSimpleField("noOfOperandStatusCodes", readSignedInt(readBuffer, 32));

    List<StatusCode> operandStatusCodes =
        readCountArrayField(
            "operandStatusCodes",
            new DataReaderComplexDefault<>(() -> StatusCode.staticParse(readBuffer), readBuffer),
            noOfOperandStatusCodes);

    int noOfOperandDiagnosticInfos =
        readSimpleField("noOfOperandDiagnosticInfos", readSignedInt(readBuffer, 32));

    List<DiagnosticInfo> operandDiagnosticInfos =
        readCountArrayField(
            "operandDiagnosticInfos",
            new DataReaderComplexDefault<>(
                () -> DiagnosticInfo.staticParse(readBuffer), readBuffer),
            noOfOperandDiagnosticInfos);

    readBuffer.closeContext("ContentFilterElementResult");
    // Create the instance
    return new ContentFilterElementResultBuilder(
        statusCode,
        noOfOperandStatusCodes,
        operandStatusCodes,
        noOfOperandDiagnosticInfos,
        operandDiagnosticInfos);
  }

  public static class ContentFilterElementResultBuilder
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final StatusCode statusCode;
    private final int noOfOperandStatusCodes;
    private final List<StatusCode> operandStatusCodes;
    private final int noOfOperandDiagnosticInfos;
    private final List<DiagnosticInfo> operandDiagnosticInfos;

    public ContentFilterElementResultBuilder(
        StatusCode statusCode,
        int noOfOperandStatusCodes,
        List<StatusCode> operandStatusCodes,
        int noOfOperandDiagnosticInfos,
        List<DiagnosticInfo> operandDiagnosticInfos) {

      this.statusCode = statusCode;
      this.noOfOperandStatusCodes = noOfOperandStatusCodes;
      this.operandStatusCodes = operandStatusCodes;
      this.noOfOperandDiagnosticInfos = noOfOperandDiagnosticInfos;
      this.operandDiagnosticInfos = operandDiagnosticInfos;
    }

    public ContentFilterElementResult build() {
      ContentFilterElementResult contentFilterElementResult =
          new ContentFilterElementResult(
              statusCode,
              noOfOperandStatusCodes,
              operandStatusCodes,
              noOfOperandDiagnosticInfos,
              operandDiagnosticInfos);
      return contentFilterElementResult;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ContentFilterElementResult)) {
      return false;
    }
    ContentFilterElementResult that = (ContentFilterElementResult) o;
    return (getStatusCode() == that.getStatusCode())
        && (getNoOfOperandStatusCodes() == that.getNoOfOperandStatusCodes())
        && (getOperandStatusCodes() == that.getOperandStatusCodes())
        && (getNoOfOperandDiagnosticInfos() == that.getNoOfOperandDiagnosticInfos())
        && (getOperandDiagnosticInfos() == that.getOperandDiagnosticInfos())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getStatusCode(),
        getNoOfOperandStatusCodes(),
        getOperandStatusCodes(),
        getNoOfOperandDiagnosticInfos(),
        getOperandDiagnosticInfos());
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